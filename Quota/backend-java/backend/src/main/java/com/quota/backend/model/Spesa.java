package com.quota.backend.model;

import jakarta.persistence.*;
import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.List;
import com.quota.backend.model.Utente;

@Entity
public class Spesa {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, precision = 10, scale = 2)
    private BigDecimal importo;

    @Column(nullable = false, length = 50)
    private String categoria;

    @Column(nullable = false)
    private LocalDate data;

    @ManyToOne
    @JoinColumn(name = "id_pagante", nullable = false)
    private Utente pagante;

    @ManyToOne
    @JoinColumn(name = "id_gruppo", nullable = false)
    private Gruppo gruppo;

    @OneToMany(mappedBy = "spesa", cascade = CascadeType.ALL, orphanRemoval = true)
    private List<PartecipanteSpesa> partecipanti;

    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt = LocalDateTime.now();

    // Getters and Setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public BigDecimal getImporto() {
        return importo;
    }

    public void setImporto(BigDecimal importo) {
        this.importo = importo;
    }

    public String getCategoria() {
        return categoria;
    }

    public void setCategoria(String categoria) {
        this.categoria = categoria;
    }

    public LocalDate getData() {
        return data;
    }

    public void setData(LocalDate data) {
        this.data = data;
    }

    public Utente getPagante() {
        return pagante;
    }

    public void setPagante(Utente pagante) {
        this.pagante = pagante;
    }

    public Gruppo getGruppo() {
        return gruppo;
    }

    public void setGruppo(Gruppo gruppo) {
        this.gruppo = gruppo;
    }

    public List<PartecipanteSpesa> getPartecipanti() {
        return partecipanti;
    }

    public void setPartecipanti(List<PartecipanteSpesa> partecipanti) {
        this.partecipanti = partecipanti;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }
}
