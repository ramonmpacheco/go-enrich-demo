package com.example.piiapi

import jakarta.persistence.Entity
import jakarta.persistence.Id
import jakarta.persistence.Table
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.data.jpa.repository.JpaRepository
import org.springframework.stereotype.Repository
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
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

data class PiiRequest(@param:JsonProperty("customer_code") val customerCode: UUID)

@RestController
class PiiController(val repository: PiiRepository) {

    @PostMapping("/pii")
    fun getPii(@RequestBody request: PiiRequest): Pii? {
        return repository.findById(request.customerCode).orElse(null)
    }
}
