package main.java.com.quota.backend.model;

import jakarta.persistence.*;
import java.time.LocalDateTime;
import java.util.List;

@Entity
public class Gruppo {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @Column(nullable = false, length = 100)
    private String nome;

    @Column(nullable = false, unique = true)
    private String linkInvito;

    @Column(nullable = false, updatable = false)
    private LocalDateTime createdAt = LocalDateTime.now();

    @OneToMany(mappedBy = "gruppo", cascade = CascadeType.ALL, orphanRemoval = true)
    private List<Spesa> spese;

    // Getters and Setters
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public String getLinkInvito() {
        return linkInvito;
    }

    public void setLinkInvito(String linkInvito) {
        this.linkInvito = linkInvito;
    }

    public LocalDateTime getCreatedAt() {
        return createdAt;
    }

    public List<Spesa> getSpese() {
        return spese;
    }

    public void setSpese(List<Spesa> spese) {
        this.spese = spese;
    }
}
