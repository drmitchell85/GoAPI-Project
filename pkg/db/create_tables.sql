-- public.article_id_sq definition

-- DROP SEQUENCE public.article_id_sq;

CREATE SEQUENCE IF NOT EXISTS public.article_id_sq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1
	CACHE 1
	NO CYCLE;


-- public.article definition

-- Drop table

-- DROP TABLE public.article;

CREATE TABLE IF NOT EXISTS public.article (
	article_id text NOT NULL DEFAULT nextval('article_id_sq'::regclass),
	articletitle text NULL,
	articledesc text NULL,
	articlecontent text NULL
);