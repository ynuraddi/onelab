--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2 (Debian 15.2-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: book_borrows; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.book_borrows (
    id integer NOT NULL,
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    borrow_date timestamp without time zone DEFAULT now() NOT NULL,
    return_date timestamp without time zone,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.book_borrows OWNER TO onelab;

--
-- Name: book_borrows_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.book_borrows_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_borrows_id_seq OWNER TO onelab;

--
-- Name: book_borrows_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.book_borrows_id_seq OWNED BY public.book_borrows.id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.books (
    id integer NOT NULL,
    title character varying NOT NULL,
    author character varying NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.books OWNER TO onelab;

--
-- Name: books_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.books_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_id_seq OWNER TO onelab;

--
-- Name: books_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.books_id_seq OWNED BY public.books.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO onelab;

--
-- Name: users; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying NOT NULL,
    login character varying NOT NULL,
    password character varying NOT NULL,
    is_active boolean NOT NULL,
    version integer DEFAULT 1 NOT NULL
);


ALTER TABLE public.users OWNER TO onelab;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO onelab;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: book_borrows id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrows ALTER COLUMN id SET DEFAULT nextval('public.book_borrows_id_seq'::regclass);


--
-- Name: books id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.books ALTER COLUMN id SET DEFAULT nextval('public.books_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: book_borrows; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.book_borrows (id, book_id, user_id, borrow_date, return_date, version) FROM stdin;
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.books (id, title, author, version) FROM stdin;
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.schema_migrations (version, dirty) FROM stdin;
3	f
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.users (id, name, login, password, is_active, version) FROM stdin;
\.


--
-- Name: book_borrows_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.book_borrows_id_seq', 1, false);


--
-- Name: books_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.books_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: book_borrows book_borrows_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrows
    ADD CONSTRAINT book_borrows_pkey PRIMARY KEY (id);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_login_key; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- Name: users users_name_key; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_name_key UNIQUE (name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: book_borrows book_borrows_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrows
    ADD CONSTRAINT book_borrows_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.books(id);


--
-- Name: book_borrows book_borrows_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrows
    ADD CONSTRAINT book_borrows_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

