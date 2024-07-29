DROP SCHEMA IF EXISTS event CASCADE;
CREATE SCHEMA event AUTHORIZATION postgres CREATE TABLE events(
    event_id UUID PRIMARY KEY,
    event_body JSON NOT NULL,
    ocurred_on VARCHAR(255) NOT NULL,
    type_name VARCHAR(255) NOT NULL,
);
-- Table: event.events_log
-- DROP TABLE IF EXISTS event.events_log;
CREATE TABLE IF NOT EXISTS event.events_log (
    event_id uuid PRIMARY KEY,
    event_body json NOT NULL,
    occurred_on character varying(255) COLLATE pg_catalog."default" NOT NULL,
    type_name character varying(255) COLLATE pg_catalog."default" NOT NULL
) TABLESPACE pg_default;
ALTER TABLE IF EXISTS event.events_log OWNER to postgres;

        CREATE TABLE IF NOT EXISTS PUBLIC.NOTIFICATIONS (
            ID UUID PRIMARY KEY,
            OWNER_ID UUID NOT NULL,
            SENDER_ID UUID NOT NULL,
            TITLE VARCHAR(255) NOT NULL,
            CONTENT VARCHAR(255) NOT NULL,
            LINK VARCHAR(255) NOT NULL,
            OCCURRED_ON VARCHAR(255) NOT NULL,
            SEEN BOOLEAN NOT NULL
        );
    
