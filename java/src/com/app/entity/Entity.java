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
    private Long id;  // Equivalent to int64 in Go

    @Column(name = "address_line1")
    @JsonProperty("address_line1")  // Jackson annotation for JSON mapping
    private String addressLine1;

    @Column(name = "address_line2")
    @JsonProperty("address_line2")  // Jackson annotation for JSON mapping
    private String addressLine2;

    @Column(name = "city")
    @JsonProperty("city")
    private String city;

    @Column(name = "state")
    @JsonProperty("state")
    private String state;

    @Column(name = "zip_code")
    @JsonProperty("zip_code")
    private String zipCode;

    @Column(name = "country")
    @JsonProperty("country")
    private String country;

    @Column(name = "created_at", updatable = false)
    @JsonProperty("created_at")  // Jackson annotation for JSON mapping
    private LocalDateTime createdAt;

    // You can also use @PrePersist to automatically set the createdAt field before saving
    @PrePersist
    public void prePersist() {
        this.createdAt = LocalDateTime.now();
    }
}
