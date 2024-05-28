DROP SCHEMA IF EXISTS complaint CASCADE;
CREATE SCHEMA complaint AUTHORIZATION postgres
 CREATE TABLE complaints(
    id UUID PRIMARY KEY NOT NULL,
    author_id VARCHAR(255) NOT NULL,
    receiver_id VARCHAR(255) NOT NULL,
    complaint_status VARCHAR(15) NOT NULL,
    title VARCHAR(80) NOT NULL,
    descriptionn VARCHAR(120) NOT NULL,
    body VARCHAR(250) NOT NULL,
    rating_rate INTEGER,
    rating_comment VARCHAR(250),
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL
) CREATE TABLE replies(
    id UUID PRIMARY KEY NOT NULL,
    complaint_id UUID NOT NULL,
    sender_id VARCHAR(255) NOT NULL,
    sender_img VARCHAR(255) NOT NULL,
    sender_name VARCHAR(255) NOT NULL,
    body VARCHAR(250) NOT NULL,
    read_status BOOLEAN NOT NULL,
    read_at VARCHAR(255) NOT NULL,
    created_at VARCHAR(255) NOT NULL,
    updated_at VARCHAR(255) NOT NULL
);

