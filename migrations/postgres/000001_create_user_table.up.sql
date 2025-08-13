CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.users (
   id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
   email VARCHAR(40) NOT NULL UNIQUE ,
   password_hash VARCHAR(255) NOT NULL ,
   created_at TIMESTAMP DEFAULT now()
);