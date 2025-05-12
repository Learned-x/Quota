package com.quota.backend.repository;

import com.quota.backend.model.Spesa;
import org.springframework.data.jpa.repository.JpaRepository;

public interface SpesaRepository extends JpaRepository<Spesa, Long> {
}
