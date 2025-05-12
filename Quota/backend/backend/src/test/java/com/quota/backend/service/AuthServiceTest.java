package com.quota.backend.service;

import com.quota.backend.model.Utente;
import com.quota.backend.repository.UtenteRepository;
import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

class AuthServiceTest {

    @Mock
    private UtenteRepository utenteRepository;

    @Mock
    private BCryptPasswordEncoder passwordEncoder;

    @InjectMocks
    private AuthService authService;

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testRegisterUser_Success() {
        String email = "test@example.com";
        String password = "password123";
        String nome = "Mario";
        String cognome = "Rossi";

        when(utenteRepository.findByEmail(email)).thenReturn(Optional.empty());
        when(passwordEncoder.encode(password)).thenReturn("encodedPassword");
        when(utenteRepository.save(any(Utente.class))).thenAnswer(invocation -> invocation.getArgument(0));

        Utente result = authService.registerUser(email, password, nome, cognome);

        assertNotNull(result);
        assertEquals(email, result.getEmail());
        assertEquals("encodedPassword", result.getPassword());
        assertEquals(nome, result.getNome());
        assertEquals(cognome, result.getCognome());

        verify(utenteRepository, times(1)).findByEmail(email);
        verify(passwordEncoder, times(1)).encode(password);
        verify(utenteRepository, times(1)).save(any(Utente.class));
    }

    @Test
    void testRegisterUser_EmailAlreadyExists() {
        String email = "test@example.com";
        String password = "password123";
        String nome = "Mario";
        String cognome = "Rossi";

        when(utenteRepository.findByEmail(email)).thenReturn(Optional.of(new Utente()));

        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () ->
                authService.registerUser(email, password, nome, cognome));

        assertEquals("Email già in uso", exception.getMessage());

        verify(utenteRepository, times(1)).findByEmail(email);
        verify(passwordEncoder, never()).encode(password);
        verify(utenteRepository, never()).save(any(Utente.class));
    }

    @Test
    void testLoginUser_Success() {
        String email = "test@example.com";
        String password = "password123";
        String encodedPassword = "encodedPassword";

        Utente utente = new Utente();
        utente.setEmail(email);
        utente.setPassword(encodedPassword);
        utente.setNome("Mario");
        utente.setCognome("Rossi");

        when(utenteRepository.findByEmail(email)).thenReturn(Optional.of(utente));
        when(passwordEncoder.matches(password, encodedPassword)).thenReturn(true);

        String token = authService.loginUser(email, password);

        assertNotNull(token);

        // Decodifica del token JWT per verificare i claim
        Claims claims = Jwts.parserBuilder()
                .setSigningKey(authService.getSecretKey())
                .build()
                .parseClaimsJws(token)
                .getBody();

        assertEquals(email, claims.getSubject());
        assertEquals("Mario", claims.get("nome"));
        assertEquals("Rossi", claims.get("cognome"));

        verify(utenteRepository, times(1)).findByEmail(email);
        verify(passwordEncoder, times(1)).matches(password, encodedPassword);
    }

    @Test
    void testLoginUser_InvalidCredentials() {
        String email = "test@example.com";
        String password = "password123";

        when(utenteRepository.findByEmail(email)).thenReturn(Optional.empty());

        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () ->
                authService.loginUser(email, password));

        assertEquals("Credenziali non valide", exception.getMessage());

        verify(utenteRepository, times(1)).findByEmail(email);
        verify(passwordEncoder, never()).matches(anyString(), anyString());
    }
}
