package com.app.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.time.LocalDateTime;

@Entity
@Data
@NoArgsConstructor
public class Address {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String addressLine1; // Default column name will be address_line1
    private String addressLine2; // Default column name will be address_line2

    private String city; // Default column name will be city
    private String state; // Default column name will be state
    private String zipCode; // Default column name will be zip_code
    private String country; // Default column name will be country

    private LocalDateTime createdAt; // Default column name will be created_at

    // PrePersist method to automatically set the createdAt field
    @PrePersist
    public void prePersist() {
        this.createdAt = LocalDateTime.now();
    }
}
