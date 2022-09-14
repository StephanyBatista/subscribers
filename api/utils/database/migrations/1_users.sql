-- +migrate Up
DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    id character varying(25) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(125) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to "user";
-- Index: idx_users_email

DROP INDEX IF EXISTS public.idx_users_email;

CREATE INDEX IF NOT EXISTS idx_users_email
    ON public.users USING btree
    (email COLLATE pg_catalog."default" ASC NULLS LAST)
    TABLESPACE pg_default;