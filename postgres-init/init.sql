CREATE TABLE addresses (
    customer_code UUID PRIMARY KEY,
    street VARCHAR(255),
    number VARCHAR(50),
    city VARCHAR(100),
    state VARCHAR(100),
    country VARCHAR(100)
);

CREATE TABLE pii (
    customer_code UUID PRIMARY KEY,
    name VARCHAR(255),
    cpf VARCHAR(14),
    email VARCHAR(255)
);

CREATE TABLE suggestions (
    id SERIAL PRIMARY KEY,
    customer_code UUID,
    restaurant_name VARCHAR(255),
    promotion_description VARCHAR(255)
);

-- Data for customer 1 (Complete)
INSERT INTO addresses (customer_code, street, number, city, state, country)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'Rua das Flores', '123', 'São Paulo', 'SP', 'Brasil');

INSERT INTO pii (customer_code, name, cpf, email)
VALUES ('123e4567-e89b-12d3-a456-426614174000', 'João Silva', '123.456.789-00', 'joao.silva@example.com');

INSERT INTO suggestions (customer_code, restaurant_name, promotion_description)
VALUES 
    ('123e4567-e89b-12d3-a456-426614174000', 'Restaurante A', '10% off'),
    ('123e4567-e89b-12d3-a456-426614174000', 'Restaurante B', 'Compre 1 Leve 2');

-- Data for customer 2 (Missing Address)
INSERT INTO pii (customer_code, name, cpf, email)
VALUES ('123e4567-e89b-12d3-a456-426614174001', 'Maria Souza', '987.654.321-00', 'maria.souza@example.com');

INSERT INTO suggestions (customer_code, restaurant_name, promotion_description)
VALUES 
    ('123e4567-e89b-12d3-a456-426614174001', 'Restaurante C', 'Frete Grátis');

-- Data for customer 3 (Missing Suggestions)
INSERT INTO addresses (customer_code, street, number, city, state, country)
VALUES ('123e4567-e89b-12d3-a456-426614174002', 'Av Paulista', '1000', 'São Paulo', 'SP', 'Brasil');

INSERT INTO pii (customer_code, name, cpf, email)
VALUES ('123e4567-e89b-12d3-a456-426614174002', 'Pedro Santos', '111.222.333-44', 'pedro.santos@example.com');
