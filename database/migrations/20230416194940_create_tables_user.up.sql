create extension citext;
CREATE TABLE public.user (
    user_id serial PRIMARY KEY,
    email VARCHAR ( 50 ) UNIQUE NOT NULL,
    name VARCHAR (50) NOT NULL,
    user_verified_at BOOLEAN,
    password VARCHAR (255) NOT NULL,
    last_login TIMESTAMP,
    reset_token citext,
    created_on TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
);