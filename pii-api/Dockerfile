FROM gradle:8.5-jdk17 AS builder
WORKDIR /app
COPY --chown=gradle:gradle . .
# Skip tests to speed up build and avoid DB dependency during build
RUN gradle build --no-daemon -x test

FROM eclipse-temurin:17-jre
WORKDIR /app
COPY --from=builder /app/build/libs/*.jar app.jar
CMD ["java", "-jar", "app.jar"]
