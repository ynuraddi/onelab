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
-- Name: book_borrowing_history; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.book_borrowing_history (
    borrow_id integer NOT NULL,
    book_id integer NOT NULL,
    user_id integer NOT NULL,
    borrow_date timestamp without time zone DEFAULT now() NOT NULL,
    return_date timestamp without time zone
);


ALTER TABLE public.book_borrowing_history OWNER TO onelab;

--
-- Name: book_borrowing_history_borrow_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.book_borrowing_history_borrow_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_borrowing_history_borrow_id_seq OWNER TO onelab;

--
-- Name: book_borrowing_history_borrow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.book_borrowing_history_borrow_id_seq OWNED BY public.book_borrowing_history.borrow_id;


--
-- Name: books; Type: TABLE; Schema: public; Owner: onelab
--

CREATE TABLE public.books (
    book_id integer NOT NULL,
    title character varying(50) NOT NULL,
    author character varying(50) NOT NULL
);


ALTER TABLE public.books OWNER TO onelab;

--
-- Name: books_book_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.books_book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.books_book_id_seq OWNER TO onelab;

--
-- Name: books_book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.books_book_id_seq OWNED BY public.books.book_id;


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
    user_id integer NOT NULL,
    user_name character varying(50) NOT NULL,
    login character varying(50) NOT NULL,
    password character varying(72) NOT NULL
);


ALTER TABLE public.users OWNER TO onelab;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: onelab
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO onelab;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: onelab
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: book_borrowing_history borrow_id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrowing_history ALTER COLUMN borrow_id SET DEFAULT nextval('public.book_borrowing_history_borrow_id_seq'::regclass);


--
-- Name: books book_id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.books ALTER COLUMN book_id SET DEFAULT nextval('public.books_book_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: book_borrowing_history; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.book_borrowing_history (borrow_id, book_id, user_id, borrow_date, return_date) FROM stdin;
\.


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: onelab
--

COPY public.books (book_id, title, author) FROM stdin;
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

COPY public.users (user_id, user_name, login, password) FROM stdin;
1	John Doe	johndoe	$2a$14$0yyjIqiWGVWtbi3.nVS6weHM.iOOYyM7B9fLBK6w9u9ElBKfizKUy
\.


--
-- Name: book_borrowing_history_borrow_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.book_borrowing_history_borrow_id_seq', 1, false);


--
-- Name: books_book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.books_book_id_seq', 1, false);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: onelab
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, true);


--
-- Name: book_borrowing_history book_borrowing_history_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrowing_history
    ADD CONSTRAINT book_borrowing_history_pkey PRIMARY KEY (borrow_id);


--
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (book_id);


--
-- Name: books books_title_author_key; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_title_author_key UNIQUE (title, author);


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
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: users users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);


--
-- Name: book_borrowing_history book_borrowing_history_book_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrowing_history
    ADD CONSTRAINT book_borrowing_history_book_id_fkey FOREIGN KEY (book_id) REFERENCES public.books(book_id);


--
-- Name: book_borrowing_history book_borrowing_history_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: onelab
--

ALTER TABLE ONLY public.book_borrowing_history
    ADD CONSTRAINT book_borrowing_history_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- PostgreSQL database dump complete
--

