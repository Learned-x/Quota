package com.quota.backend.service;

import com.quota.backend.model.Utente;
import com.quota.backend.repository.UtenteRepository;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import java.security.Key;
import java.util.Date;

@Service
public class AuthService {

    @Autowired
    private UtenteRepository utenteRepository;

    @Autowired
    private BCryptPasswordEncoder passwordEncoder;

    private static final Key SECRET_KEY = Keys.secretKeyFor(SignatureAlgorithm.HS256);

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
        // Logica per invalidare il token JWT
    }

    public static Key getSecretKey() {
        return SECRET_KEY;
    }
}
