--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

-- Started on 2024-11-13 17:27:13

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

DROP DATABASE "Travelika";
--
-- TOC entry 4876 (class 1262 OID 37596)
-- Name: Travelika; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "Travelika" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';


ALTER DATABASE "Travelika" OWNER TO postgres;

\connect "Travelika"

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
-- TOC entry 4877 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 37700)
-- Name: destinations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.destinations (
    id integer NOT NULL,
    location character varying(100) NOT NULL,
    image_url character varying(255) NOT NULL,
    description character varying(255) NOT NULL
);


ALTER TABLE public.destinations OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 37699)
-- Name: destinations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.destinations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.destinations_id_seq OWNER TO postgres;

--
-- TOC entry 4878 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.destinations_id_seq OWNED BY public.destinations.id;


--
-- TOC entry 218 (class 1259 OID 37709)
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.events (
    id integer NOT NULL,
    destination_id integer NOT NULL,
    name character varying(100) NOT NULL,
    schedule date NOT NULL,
    price numeric(10,2) NOT NULL
);


ALTER TABLE public.events OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 37708)
-- Name: events_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.events_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.events_id_seq OWNER TO postgres;

--
-- TOC entry 4879 (class 0 OID 0)
-- Dependencies: 217
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- TOC entry 222 (class 1259 OID 37733)
-- Name: reviews; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.reviews (
    id integer NOT NULL,
    destination_id integer NOT NULL,
    transaction_id integer NOT NULL,
    rating integer,
    comment character varying(255),
    CONSTRAINT reviews_rating_check CHECK (((rating >= 1) AND (rating <= 5)))
);


ALTER TABLE public.reviews OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 37732)
-- Name: reviews_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.reviews_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.reviews_id_seq OWNER TO postgres;

--
-- TOC entry 4880 (class 0 OID 0)
-- Dependencies: 221
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- TOC entry 220 (class 1259 OID 37721)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    event_id integer NOT NULL,
    status boolean
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 37720)
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
-- TOC entry 4881 (class 0 OID 0)
-- Dependencies: 219
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 4703 (class 2604 OID 37703)
-- Name: destinations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations ALTER COLUMN id SET DEFAULT nextval('public.destinations_id_seq'::regclass);


--
-- TOC entry 4704 (class 2604 OID 37712)
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- TOC entry 4706 (class 2604 OID 37736)
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- TOC entry 4705 (class 2604 OID 37724)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 4864 (class 0 OID 37700)
-- Dependencies: 216
-- Data for Name: destinations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.destinations VALUES (1, 'Bali, Indonesia', 'https://example.com/bali.jpg', 'Pulau tropis dengan pantai indah dan budaya Bali yang unik.');
INSERT INTO public.destinations VALUES (2, 'Kyoto, Japan', 'https://example.com/kyoto.jpg', 'Kota tradisional Jepang dengan banyak kuil dan taman.');
INSERT INTO public.destinations VALUES (3, 'Paris, France', 'https://example.com/paris.jpg', 'Kota romantis dengan Menara Eiffel dan seni yang kaya.');
INSERT INTO public.destinations VALUES (4, 'Rome, Italy', 'https://example.com/rome.jpg', 'Kota dengan sejarah kuno, Colosseum, dan museum terkenal.');
INSERT INTO public.destinations VALUES (5, 'Sydney, Australia', 'https://example.com/sydney.jpg', 'Kota pesisir dengan Opera House dan pantai Bondi.');
INSERT INTO public.destinations VALUES (6, 'New York, USA', 'https://example.com/ny.jpg', 'Kota modern dengan gedung pencakar langit dan Central Park.');
INSERT INTO public.destinations VALUES (7, 'London, UK', 'https://example.com/london.jpg', 'Kota sejarah dan seni dengan Big Ben dan Buckingham Palace.');
INSERT INTO public.destinations VALUES (8, 'Dubai, UAE', 'https://example.com/dubai.jpg', 'Kota futuristik dengan gedung tertinggi di dunia, Burj Khalifa.');
INSERT INTO public.destinations VALUES (9, 'Bangkok, Thailand', 'https://example.com/bangkok.jpg', 'Kota budaya Asia Tenggara dengan kuil dan pasar tradisional.');
INSERT INTO public.destinations VALUES (10, 'Istanbul, Turkey', 'https://example.com/istanbul.jpg', 'Kota dengan perpaduan budaya Eropa dan Asia serta masjid bersejarah.');


--
-- TOC entry 4866 (class 0 OID 37709)
-- Dependencies: 218
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.events VALUES (1, 1, 'Bali Cultural Tour', '2024-12-01', 150.00);
INSERT INTO public.events VALUES (2, 2, 'Kyoto Temple Visit', '2024-11-20', 120.00);
INSERT INTO public.events VALUES (3, 3, 'Paris Art & Wine', '2024-12-15', 200.00);
INSERT INTO public.events VALUES (4, 4, 'Rome Ancient Tour', '2025-01-10', 180.00);
INSERT INTO public.events VALUES (5, 5, 'Sydney Harbour Cruise', '2024-12-05', 250.00);
INSERT INTO public.events VALUES (6, 6, 'NYC Broadway Show', '2024-12-22', 300.00);
INSERT INTO public.events VALUES (7, 7, 'London History Walk', '2024-11-25', 90.00);
INSERT INTO public.events VALUES (8, 8, 'Dubai Desert Safari', '2024-12-18', 220.00);
INSERT INTO public.events VALUES (9, 9, 'Bangkok Food Tour', '2024-11-30', 80.00);
INSERT INTO public.events VALUES (10, 10, 'Istanbul Bazaar Walk', '2024-12-08', 70.00);


--
-- TOC entry 4870 (class 0 OID 37733)
-- Dependencies: 222
-- Data for Name: reviews; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.reviews VALUES (1, 1, 1, 5, 'Amazing experience, highly recommended!');
INSERT INTO public.reviews VALUES (2, 2, 2, 4, 'Beautiful temples, but a bit crowded.');
INSERT INTO public.reviews VALUES (3, 3, 3, 3, 'Interesting, but not worth the price.');
INSERT INTO public.reviews VALUES (4, 4, 4, 5, 'Incredible history and architecture!');
INSERT INTO public.reviews VALUES (5, 5, 5, 4, 'Great cruise with stunning views.');
INSERT INTO public.reviews VALUES (6, 6, 6, 3, 'Good show, but too pricey.');
INSERT INTO public.reviews VALUES (7, 7, 7, 4, 'Informative and fun walking tour.');
INSERT INTO public.reviews VALUES (8, 8, 8, 5, 'A must-do experience in the desert!');
INSERT INTO public.reviews VALUES (9, 9, 9, 3, 'Food was good, but expected more variety.');
INSERT INTO public.reviews VALUES (10, 10, 10, 4, 'Unique experience with friendly locals.');


--
-- TOC entry 4868 (class 0 OID 37721)
-- Dependencies: 220
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES (1, 1, true);
INSERT INTO public.transactions VALUES (2, 2, true);
INSERT INTO public.transactions VALUES (3, 3, false);
INSERT INTO public.transactions VALUES (4, 4, true);
INSERT INTO public.transactions VALUES (5, 5, true);
INSERT INTO public.transactions VALUES (6, 6, false);
INSERT INTO public.transactions VALUES (7, 7, true);
INSERT INTO public.transactions VALUES (8, 8, true);
INSERT INTO public.transactions VALUES (9, 9, false);
INSERT INTO public.transactions VALUES (10, 10, true);


--
-- TOC entry 4882 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.destinations_id_seq', 10, true);


--
-- TOC entry 4883 (class 0 OID 0)
-- Dependencies: 217
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 10, true);


--
-- TOC entry 4884 (class 0 OID 0)
-- Dependencies: 221
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 10, true);


--
-- TOC entry 4885 (class 0 OID 0)
-- Dependencies: 219
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 10, true);


--
-- TOC entry 4709 (class 2606 OID 37707)
-- Name: destinations destinations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations
    ADD CONSTRAINT destinations_pkey PRIMARY KEY (id);


--
-- TOC entry 4711 (class 2606 OID 37714)
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- TOC entry 4715 (class 2606 OID 37739)
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- TOC entry 4713 (class 2606 OID 37726)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 4716 (class 2606 OID 37715)
-- Name: events events_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4718 (class 2606 OID 37740)
-- Name: reviews reviews_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4719 (class 2606 OID 37745)
-- Name: reviews reviews_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- TOC entry 4717 (class 2606 OID 37727)
-- Name: transactions transactions_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


-- Completed on 2024-11-13 17:27:14

--
-- PostgreSQL database dump complete
--

