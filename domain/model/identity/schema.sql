DROP SCHEMA IF EXISTS identityandaccess CASCADE;
CREATE TABLE users(
    id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    register_date VARCHAR(255) NOT NULL,
    is_confirmed boolean not null
)
CREATE TABLE roles(
    role_name VARCHAR(255) PRIMARY KEY
)
CREATE TABLE user_roles(
    user_id UUID NOT NULL,
    role_id VARCHAR(255) NOT NULL,
    enterprise_id UUID NOT NULL,
    PRIMARY KEY(user_id, role_id, enterprise_id)
)
CREATE TABLE person(
    id UUID PRIMARY KEY,
    email VARCHAR(255) PRIMARY KEY,
    profile_img VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    pronoun VARCHAR(255) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    birth_date VARCHAR(255) NOT NULL,
    phone VARCHAR(21) NOT NULL,
    address_id UUID NOT NULL,
);
CREATE TABLE address(
    id UUID PRIMARY KEY,
    country_id INTEGER NOT NULL,
    country_state_id INTEGER NOT NULL,
    state_cities_id INTEGER NOT NULL
);
INSERT INTO identityandaccess.roles(role_name)
VALUES('ASSISTANT');
INSERT INTO identityandaccess.roles(role_name)
VALUES('MANAGER');
INSERT INTO identityandaccess.roles(role_name)
VALUES('OWNER');