-- +migrate Up

DROP TABLE IF EXISTS public.subscribers;
DROP TABLE IF EXISTS public.campaigns;

-- Table: public.campaigns

CREATE TABLE IF NOT EXISTS public.campaigns
(
    id character varying(25) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    status character varying(15) COLLATE pg_catalog."default" NOT NULL,
    "from" character varying(100) COLLATE pg_catalog."default" NOT NULL,
    subject character varying(150) COLLATE pg_catalog."default" NOT NULL,
    body text COLLATE pg_catalog."default" NOT NULL,
    user_id character varying(25) COLLATE pg_catalog."default",
    CONSTRAINT campaigns_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.campaigns
    OWNER to "user";

-- Table: public.subscribers

CREATE TABLE IF NOT EXISTS public.subscribers
(
    id character varying(25) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL,
    campaign_id character varying(25) COLLATE pg_catalog."default" NOT NULL,
    contact_id character varying(25) COLLATE pg_catalog."default" NOT NULL,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    status character varying(15) COLLATE pg_catalog."default" NOT NULL,
    provider_email_key text COLLATE pg_catalog."default",
    CONSTRAINT subscribers_pkey PRIMARY KEY (id),
    CONSTRAINT fk_subscribers_campaign FOREIGN KEY (campaign_id)
        REFERENCES public.campaigns (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.subscribers
    OWNER to "user";

