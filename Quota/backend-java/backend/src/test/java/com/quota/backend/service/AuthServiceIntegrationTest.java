package com.quota.backend.service;

import com.quota.backend.model.Utente;
import com.quota.backend.repository.UtenteRepository;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.annotation.DirtiesContext;

import static org.junit.jupiter.api.Assertions.*;

@SpringBootTest
@DirtiesContext(classMode = DirtiesContext.ClassMode.AFTER_EACH_TEST_METHOD)
class AuthServiceIntegrationTest {

    @Autowired
    private UtenteRepository utenteRepository;

    @Autowired
    private AuthService authService;

    @Test
    void testRegisterUserIntegration() {
        String email = "integration@example.com";
        String password = "password123";
        String nome = "Integration";
        String cognome = "Test";

        // Verifica che l'utente non esista prima della registrazione
        assertFalse(utenteRepository.findByEmail(email).isPresent());

        // Registra un nuovo utente
        Utente utente = authService.registerUser(email, password, nome, cognome);

        // Verifica che l'utente sia stato salvato correttamente
        assertNotNull(utente);
        assertEquals(email, utente.getEmail());
        assertEquals(nome, utente.getNome());
        assertEquals(cognome, utente.getCognome());

        // Verifica che l'utente sia presente nel database
        assertTrue(utenteRepository.findByEmail(email).isPresent());
    }
}