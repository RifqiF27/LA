--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

-- Started on 2024-11-10 23:42:52

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

DROP DATABASE "Cinema";
--
-- TOC entry 4921 (class 1262 OID 37006)
-- Name: Cinema; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "Cinema" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';


ALTER DATABASE "Cinema" OWNER TO postgres;

\connect "Cinema"

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

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: pg_database_owner
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO pg_database_owner;

--
-- TOC entry 4922 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 226 (class 1259 OID 37312)
-- Name: bookings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bookings (
    id integer NOT NULL,
    user_id integer,
    cinema_id integer,
    seat_id integer,
    booking_date date NOT NULL,
    booking_time time without time zone NOT NULL,
    payment_method character varying(50),
    status character varying(20) DEFAULT 'pending'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.bookings OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 37311)
-- Name: bookings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bookings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bookings_id_seq OWNER TO postgres;

--
-- TOC entry 4923 (class 0 OID 0)
-- Dependencies: 225
-- Name: bookings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bookings_id_seq OWNED BY public.bookings.id;


--
-- TOC entry 222 (class 1259 OID 37289)
-- Name: cinemas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cinemas (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    location text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.cinemas OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 37288)
-- Name: cinemas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cinemas_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cinemas_id_seq OWNER TO postgres;

--
-- TOC entry 4924 (class 0 OID 0)
-- Dependencies: 221
-- Name: cinemas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cinemas_id_seq OWNED BY public.cinemas.id;


--
-- TOC entry 216 (class 1259 OID 37070)
-- Name: payment_methods; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_methods (
    id integer NOT NULL,
    method_name character varying(50) NOT NULL
);


ALTER TABLE public.payment_methods OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 37069)
-- Name: payment_methods_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.payment_methods_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.payment_methods_id_seq OWNER TO postgres;

--
-- TOC entry 4925 (class 0 OID 0)
-- Dependencies: 215
-- Name: payment_methods_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.payment_methods_id_seq OWNED BY public.payment_methods.id;


--
-- TOC entry 224 (class 1259 OID 37299)
-- Name: seats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.seats (
    id integer NOT NULL,
    cinema_id integer,
    date date NOT NULL,
    "time" time without time zone NOT NULL,
    status character varying(50) NOT NULL,
    CONSTRAINT seats_status_check CHECK (((status)::text = ANY ((ARRAY['available'::character varying, 'booked'::character varying])::text[])))
);


ALTER TABLE public.seats OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 37298)
-- Name: seats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.seats_id_seq OWNER TO postgres;

--
-- TOC entry 4926 (class 0 OID 0)
-- Dependencies: 223
-- Name: seats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.seats_id_seq OWNED BY public.seats.id;


--
-- TOC entry 220 (class 1259 OID 37272)
-- Name: sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sessions (
    id integer NOT NULL,
    user_id integer,
    token text NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.sessions OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 37271)
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sessions_id_seq OWNER TO postgres;

--
-- TOC entry 4927 (class 0 OID 0)
-- Dependencies: 219
-- Name: sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sessions_id_seq OWNED BY public.sessions.id;


--
-- TOC entry 228 (class 1259 OID 37336)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    booking_id integer,
    transaction_type character varying(50) NOT NULL,
    amount numeric(10,2),
    transaction_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    payment_method_id integer,
    status character varying(20) DEFAULT 'completed'::character varying
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 37335)
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transactions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.transactions_id_seq OWNER TO postgres;

--
-- TOC entry 4928 (class 0 OID 0)
-- Dependencies: 227
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 218 (class 1259 OID 37261)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(100) NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 37260)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 4929 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 4725 (class 2604 OID 37315)
-- Name: bookings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings ALTER COLUMN id SET DEFAULT nextval('public.bookings_id_seq'::regclass);


--
-- TOC entry 4722 (class 2604 OID 37292)
-- Name: cinemas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas ALTER COLUMN id SET DEFAULT nextval('public.cinemas_id_seq'::regclass);


--
-- TOC entry 4718 (class 2604 OID 37073)
-- Name: payment_methods id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods ALTER COLUMN id SET DEFAULT nextval('public.payment_methods_id_seq'::regclass);


--
-- TOC entry 4724 (class 2604 OID 37302)
-- Name: seats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats ALTER COLUMN id SET DEFAULT nextval('public.seats_id_seq'::regclass);


--
-- TOC entry 4720 (class 2604 OID 37275)
-- Name: sessions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions ALTER COLUMN id SET DEFAULT nextval('public.sessions_id_seq'::regclass);


--
-- TOC entry 4728 (class 2604 OID 37339)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 4719 (class 2604 OID 37264)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 4913 (class 0 OID 37312)
-- Dependencies: 226
-- Data for Name: bookings; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 4909 (class 0 OID 37289)
-- Dependencies: 222
-- Data for Name: cinemas; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.cinemas VALUES (1, 'Cinema 1', '123 Main St, City A', '2024-11-10 22:18:04.177652');
INSERT INTO public.cinemas VALUES (2, 'Cinema 2', '456 Elm St, City B', '2024-11-10 22:18:04.177652');


--
-- TOC entry 4903 (class 0 OID 37070)
-- Dependencies: 216
-- Data for Name: payment_methods; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 4911 (class 0 OID 37299)
-- Dependencies: 224
-- Data for Name: seats; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.seats VALUES (1, 1, '2024-11-15', '15:00:00', 'available');
INSERT INTO public.seats VALUES (2, 1, '2024-11-15', '15:00:00', 'available');
INSERT INTO public.seats VALUES (3, 1, '2024-11-15', '17:00:00', 'booked');
INSERT INTO public.seats VALUES (4, 1, '2024-11-16', '15:00:00', 'booked');


--
-- TOC entry 4907 (class 0 OID 37272)
-- Dependencies: 220
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.sessions VALUES (1, 1, '3650f4179cb527c330cf7c39fb234ce36241cf399c488c985404fa06f485485f', '2024-11-11 22:28:06.766276', '2024-11-10 22:28:06.767421');


--
-- TOC entry 4915 (class 0 OID 37336)
-- Dependencies: 228
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 4905 (class 0 OID 37261)
-- Dependencies: 218
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES (1, 'admin', 'admin', 'admin1@mail.com');


--
-- TOC entry 4930 (class 0 OID 0)
-- Dependencies: 225
-- Name: bookings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bookings_id_seq', 1, false);


--
-- TOC entry 4931 (class 0 OID 0)
-- Dependencies: 221
-- Name: cinemas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cinemas_id_seq', 2, true);


--
-- TOC entry 4932 (class 0 OID 0)
-- Dependencies: 215
-- Name: payment_methods_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.payment_methods_id_seq', 1, false);


--
-- TOC entry 4933 (class 0 OID 0)
-- Dependencies: 223
-- Name: seats_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seats_id_seq', 4, true);


--
-- TOC entry 4934 (class 0 OID 0)
-- Dependencies: 219
-- Name: sessions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sessions_id_seq', 1, true);


--
-- TOC entry 4935 (class 0 OID 0)
-- Dependencies: 227
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 1, false);


--
-- TOC entry 4936 (class 0 OID 0)
-- Dependencies: 217
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- TOC entry 4749 (class 2606 OID 37319)
-- Name: bookings bookings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_pkey PRIMARY KEY (id);


--
-- TOC entry 4745 (class 2606 OID 37297)
-- Name: cinemas cinemas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cinemas
    ADD CONSTRAINT cinemas_pkey PRIMARY KEY (id);


--
-- TOC entry 4733 (class 2606 OID 37075)
-- Name: payment_methods payment_methods_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_methods
    ADD CONSTRAINT payment_methods_pkey PRIMARY KEY (id);


--
-- TOC entry 4747 (class 2606 OID 37305)
-- Name: seats seats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_pkey PRIMARY KEY (id);


--
-- TOC entry 4741 (class 2606 OID 37280)
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- TOC entry 4743 (class 2606 OID 37282)
-- Name: sessions sessions_token_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_token_key UNIQUE (token);


--
-- TOC entry 4751 (class 2606 OID 37343)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 4735 (class 2606 OID 37270)
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- TOC entry 4737 (class 2606 OID 37266)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 4739 (class 2606 OID 37268)
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- TOC entry 4754 (class 2606 OID 37325)
-- Name: bookings bookings_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- TOC entry 4755 (class 2606 OID 37330)
-- Name: bookings bookings_seat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_seat_id_fkey FOREIGN KEY (seat_id) REFERENCES public.seats(id) ON DELETE CASCADE;


--
-- TOC entry 4756 (class 2606 OID 37320)
-- Name: bookings bookings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bookings
    ADD CONSTRAINT bookings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4753 (class 2606 OID 37306)
-- Name: seats seats_cinema_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.seats
    ADD CONSTRAINT seats_cinema_id_fkey FOREIGN KEY (cinema_id) REFERENCES public.cinemas(id) ON DELETE CASCADE;


--
-- TOC entry 4752 (class 2606 OID 37283)
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- TOC entry 4757 (class 2606 OID 37344)
-- Name: transactions transactions_booking_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_booking_id_fkey FOREIGN KEY (booking_id) REFERENCES public.bookings(id) ON DELETE CASCADE;


--
-- TOC entry 4758 (class 2606 OID 37349)
-- Name: transactions transactions_payment_method_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_payment_method_id_fkey FOREIGN KEY (payment_method_id) REFERENCES public.payment_methods(id) ON DELETE SET NULL;


-- Completed on 2024-11-10 23:42:52

--
-- PostgreSQL database dump complete
--

