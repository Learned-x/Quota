package com.quota.backend.service;

import com.quota.backend.model.Utente;
import com.quota.backend.repository.UtenteRepository;
import io.jsonwebtoken.ExpiredJwtException;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.MalformedJwtException;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import io.jsonwebtoken.security.SignatureException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.security.Key;
import java.util.Date;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

@Service
public class AuthService {

    @Autowired
    private UtenteRepository utenteRepository;

    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    private static final Key SECRET_KEY = Keys.secretKeyFor(SignatureAlgorithm.HS256);

    private final Set<String> invalidatedTokens = ConcurrentHashMap.newKeySet();

    private static final Logger logger = LoggerFactory.getLogger(AuthService.class);

    public Utente registerUser(String email, String password, String nome, String cognome) {
        // Validazione input
        if (utenteRepository.findByEmail(email).isPresent()) {
            throw new IllegalArgumentException("Email già in uso");
        }

        // Crittografia password
        String encodedPassword = passwordEncoder.encode(password);

        // Creazione utente
        Utente utente = new Utente();
        utente.setEmail(email);
        utente.setPassword(encodedPassword);
        utente.setNome(nome);
        utente.setCognome(cognome);
        utente.setProvider("email"); // Imposta il provider predefinito come 'email'
        utente.setUsername(email); // Imposta il valore predefinito per il campo `username` come l'email dell'utente

        return utenteRepository.save(utente);
    }

    // Metodo per il login e generazione del token JWT
    public String loginUser(String email, String password) {
        Utente utente = utenteRepository.findByEmail(email)
                .orElseThrow(() -> new IllegalArgumentException("Credenziali non valide"));

        if (!passwordEncoder.matches(password, utente.getPassword())) {
            throw new IllegalArgumentException("Credenziali non valide");
        }

        // Generazione del token JWT
        return Jwts.builder()
                .setSubject(utente.getEmail())
                .claim("email", utente.getEmail()) // Aggiunta esplicita dell'email come claim
                .claim("nome", utente.getNome())
                .claim("cognome", utente.getCognome())
                .setIssuedAt(new Date())
                .setExpiration(new Date(System.currentTimeMillis() + 3600000)) // 1 ora
                .signWith(SECRET_KEY)
                .compact();
    }

    // Metodo per il logout (invalida il token JWT, da implementare)
    public void logoutUser(String token) {
        try {
            // Verifica se il token è valido prima di invalidarlo
            Jwts.parserBuilder().setSigningKey(SECRET_KEY).build().parseClaimsJws(token);
            invalidatedTokens.add(token);
            logger.info("Token invalidato con successo: {}", token);
        } catch (ExpiredJwtException e) {
            logger.warn("Tentativo di logout con token scaduto: {}", token);
            throw new IllegalArgumentException("Token scaduto"); // Corretto per lanciare il messaggio atteso
        } catch (MalformedJwtException | SignatureException e) {
            logger.error("Tentativo di logout con token non valido: {}", token);
            throw new IllegalArgumentException("Token non valido");
        }
    }

    public boolean isTokenValid(String token) {
        try {
            // Verifica se il token è valido e non nella blacklist
            Jwts.parserBuilder().setSigningKey(SECRET_KEY).build().parseClaimsJws(token);
            return !invalidatedTokens.contains(token);
        } catch (ExpiredJwtException e) {
            logger.warn("Token scaduto: {}", token);
            return false;
        } catch (MalformedJwtException | SignatureException e) {
            logger.error("Token non valido: {}", token);
            return false;
        }
    }

    public static Key getSecretKey() {
        return SECRET_KEY;
    }
}
