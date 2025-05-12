package com.quota.backend.controller;

import com.quota.backend.model.Utente;
import com.quota.backend.service.AuthService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/auth")
public class AuthController {

    @Autowired
    private AuthService authService;

    @PostMapping("/register")
    public ResponseEntity<?> registerUser(@RequestBody RegisterRequest request) {
        try {
            Utente utente = authService.registerUser(request.getEmail(), request.getPassword(), request.getNome(), request.getCognome());
            return ResponseEntity.ok(new RegisterResponse("Utente registrato con successo", utente.getId()));
        } catch (IllegalArgumentException e) {
            return ResponseEntity.badRequest().body(e.getMessage());
        }
    }

    // Classe per la richiesta di registrazione
    public static class RegisterRequest {
        private String email;
        private String password;
        private String nome;
        private String cognome;

        // Getter e Setter
        public String getEmail() {
            return email;
        }

        public void setEmail(String email) {
            this.email = email;
        }

        public String getPassword() {
            return password;
        }

        public void setPassword(String password) {
            this.password = password;
        }

        public String getNome() {
            return nome;
        }

        public void setNome(String nome) {
            this.nome = nome;
        }

        public String getCognome() {
            return cognome;
        }

        public void setCognome(String cognome) {
            this.cognome = cognome;
        }
    }

    // Classe per la risposta di registrazione
    public static class RegisterResponse {
        private String message;
        private Long userId;

        public RegisterResponse(String message, Long userId) {
            this.message = message;
            this.userId = userId;
        }

        // Getter
        public String getMessage() {
            return message;
        }

        public Long getUserId() {
            return userId;
        }
    }
}
