(database name cat_adoption_db)

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Definisi enum
CREATE TYPE gender_enum AS ENUM ('male', 'female');
CREATE TYPE vaccination_status_enum AS ENUM ('vaccinated', 'not_vaccinated');
CREATE TYPE role_enum AS ENUM ('admin', 'user');

-- Tabel m_cat
CREATE TABLE IF NOT EXISTS public.m_cat
(
    cat_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    name character varying(100) COLLATE pg_catalog."default",
    breed character varying(100) COLLATE pg_catalog."default",
    age integer,
    color character varying(50) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    adopted boolean DEFAULT false,
    latitude numeric(10, 8),
    longitude numeric(11, 8),
    location_name character varying(255) COLLATE pg_catalog."default",
    photo_url character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    gender gender_enum,
    vaccination_status vaccination_status_enum,
    CONSTRAINT m_cat_pkey PRIMARY KEY (cat_id)
);

-- Tabel m_user
CREATE TABLE IF NOT EXISTS public.m_user
(
    user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username character varying(100) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    name character varying(100) COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    phone_number character varying(20) COLLATE pg_catalog."default",
    address character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    role role_enum,
    CONSTRAINT m_user_pkey PRIMARY KEY (user_id),
    CONSTRAINT m_user_username_key UNIQUE (username)
);

-- Tabel t_review
CREATE TABLE IF NOT EXISTS public.t_review
(
    review_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid,
    cat_id uuid,
    rating integer,
    comment text COLLATE pg_catalog."default",
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT t_review_pkey PRIMARY KEY (review_id),
    CONSTRAINT t_review_cat_id_fkey FOREIGN KEY (cat_id)
        REFERENCES public.m_cat (cat_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT t_review_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.m_user (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

-- Tabel tx_adoption
CREATE TABLE IF NOT EXISTS public.tx_adoption
(
    adoption_id uuid NOT NULL DEFAULT uuid_generate_v4(),
    user_id uuid,
    cat_id uuid,
    adopted_date date,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT tx_adoption_pkey PRIMARY KEY (adoption_id),
    CONSTRAINT tx_adoption_cat_id_fkey FOREIGN KEY (cat_id)
        REFERENCES public.m_cat (cat_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT tx_adoption_user_id_fkey FOREIGN KEY (user_id)
        REFERENCES public.m_user (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
