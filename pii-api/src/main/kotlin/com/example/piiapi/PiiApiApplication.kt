package com.example.piiapi

import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController
import com.fasterxml.jackson.annotation.JsonProperty
import java.util.UUID

@SpringBootApplication
class PiiApiApplication

fun main(args: Array<String>) {
    runApplication<PiiApiApplication>(*args)
}

@Entity
@Table(name = "pii")
data class Pii(
    @Id
    val customerCode: UUID,
    val name: String,
    val cpf: String,
    val email: String
)

@Repository
interface PiiRepository : JpaRepository<Pii, UUID>

@RestController
class PiiController(val repository: PiiRepository) {

    @GetMapping("/pii")
    fun getPii(@RequestParam("customer_code") customerCode: UUID): Pii? {
        return repository.findById(customerCode).orElse(null)
    }
}
