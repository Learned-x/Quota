package com.quota.backend.service;

import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import io.jsonwebtoken.security.Keys;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.MockitoAnnotations;

import java.security.Key;
import java.util.Date;

import static org.junit.jupiter.api.Assertions.*;

class AuthServiceLogoutTest {

    @InjectMocks
    private AuthService authService;

    private static final Key SECRET_KEY = AuthService.getSecretKey(); // Usa la chiave segreta del servizio

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
    }

    @Test
    void testLogoutUser_ValidToken() {
        // Genera un token JWT valido
        String token = Jwts.builder()
                .setSubject("test@example.com")
                .setIssuedAt(new Date())
                .setExpiration(new Date(System.currentTimeMillis() + 3600000)) // 1 ora
                .signWith(SECRET_KEY)
                .compact();

        // Logout
        authService.logoutUser(token);

        // Verifica che il token sia invalidato
        assertFalse(authService.isTokenValid(token));
    }

    @Test
    void testLogoutUser_ExpiredToken() {
        // Genera un token JWT scaduto
        String token = Jwts.builder()
                .setSubject("test@example.com")
                .setIssuedAt(new Date(System.currentTimeMillis() - 7200000)) // 2 ore fa
                .setExpiration(new Date(System.currentTimeMillis() - 3600000)) // 1 ora fa
                .signWith(SECRET_KEY)
                .compact();

        // Verifica che il logout lanci un'eccezione
        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () ->
                authService.logoutUser(token));

        assertEquals("Token scaduto", exception.getMessage());
    }

    @Test
    void testLogoutUser_InvalidToken() {
        // Token non valido
        String token = "invalid.token.here";

        // Verifica che il logout lanci un'eccezione
        IllegalArgumentException exception = assertThrows(IllegalArgumentException.class, () ->
                authService.logoutUser(token));

        assertEquals("Token non valido", exception.getMessage());
    }
}
