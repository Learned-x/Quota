package com.quota.backend.model;

import jakarta.persistence.*;
import java.time.LocalDateTime;

@Entity
public class Avatar {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Lob
    @Column(nullable = false)
    private String immagineBase64;

    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt = LocalDateTime.now();

    // Getters and Setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getImmagineBase64() {
        return immagineBase64;
    }

    public void setImmagineBase64(String immagineBase64) {
        this.immagineBase64 = immagineBase64;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }
}
