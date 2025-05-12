package main.java.com.quota.backend.model;

import jakarta.persistence.*;
import java.math.BigDecimal;

@Entity
public class PartecipanteSpesa {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne
    @JoinColumn(name = "id_spesa", nullable = false)
    private Spesa spesa;

    @ManyToOne
    @JoinColumn(name = "id_partecipante", nullable = false)
    private Utente partecipante;

    @Column(nullable = false, precision = 10, scale = 2)
    private BigDecimal quota;

    // Getters and Setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Spesa getSpesa() {
        return spesa;
    }

    public void setSpesa(Spesa spesa) {
        this.spesa = spesa;
    }

    public Utente getPartecipante() {
        return partecipante;
    }

    public void setPartecipante(Utente partecipante) {
        this.partecipante = partecipante;
    }

    public BigDecimal getQuota() {
        return quota;
    }

    public void setQuota(BigDecimal quota) {
        this.quota = quota;
    }
}
