DROP SCHEMA IF EXISTS identityandaccess CASCADE;
CREATE SCHEMA identityandaccess AUTHORIZATION postgres CREATE TABLE users(
    email VARCHAR(255) PRIMARY KEY,
    profile_img VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    register_date VARCHAR(255) NOT NULL
) 
CREATE TABLE roles(role_name VARCHAR(255) PRIMARY KEY)
CREATE TABLE user_roles(
        user_id VARCHAR(255) NOT NULL,
        role_id VARCHAR(255) NOT NULL,
        enterprise_id VARCHAR(255) NOT NULL,
        PRIMARY KEY(user_id, role_id, enterprise_id)
	 )
CREATE TABLE persons(
        email VARCHAR(255) PRIMARY KEY,
        first_name VARCHAR(50) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        birth_date VARCHAR(255) NOT NULL,
        phone VARCHAR(21) NOT NULL,
        country VARCHAR(255) NOT NULL,
        county VARCHAR(255) NOT NULL,
        city VARCHAR(255) NOT NULL
    );
INSERT INTO identityandaccess.roles(role_name)
VALUES('ASSISTANT');
INSERT INTO  identityandaccess.roles(role_name)
VALUES('MANAGER');
INSERT INTO  identityandaccess.roles(role_name)
VALUES('OWNER');