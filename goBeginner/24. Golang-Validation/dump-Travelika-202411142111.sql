--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

-- Started on 2024-11-14 21:11:39

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
-- TOC entry 4888 (class 1262 OID 37596)
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
-- TOC entry 4889 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 39013)
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
-- TOC entry 215 (class 1259 OID 39012)
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
-- TOC entry 4890 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.destinations_id_seq OWNED BY public.destinations.id;


--
-- TOC entry 220 (class 1259 OID 39036)
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
-- TOC entry 219 (class 1259 OID 39035)
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
-- TOC entry 4891 (class 0 OID 0)
-- Dependencies: 219
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- TOC entry 218 (class 1259 OID 39022)
-- Name: gallery; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.gallery (
    id integer NOT NULL,
    destination_id integer NOT NULL,
    image_url character varying(255) NOT NULL,
    description character varying(255) NOT NULL
);


ALTER TABLE public.gallery OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 39021)
-- Name: gallery_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.gallery_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.gallery_id_seq OWNER TO postgres;

--
-- TOC entry 4892 (class 0 OID 0)
-- Dependencies: 217
-- Name: gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gallery_id_seq OWNED BY public.gallery.id;


--
-- TOC entry 224 (class 1259 OID 39061)
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
-- TOC entry 223 (class 1259 OID 39060)
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
-- TOC entry 4893 (class 0 OID 0)
-- Dependencies: 223
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- TOC entry 222 (class 1259 OID 39048)
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    phone character varying(13) NOT NULL,
    comment character varying(255),
    event_id integer NOT NULL,
    status boolean,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 39047)
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
-- TOC entry 4894 (class 0 OID 0)
-- Dependencies: 221
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 4708 (class 2604 OID 39016)
-- Name: destinations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations ALTER COLUMN id SET DEFAULT nextval('public.destinations_id_seq'::regclass);


--
-- TOC entry 4710 (class 2604 OID 39039)
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- TOC entry 4709 (class 2604 OID 39025)
-- Name: gallery id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery ALTER COLUMN id SET DEFAULT nextval('public.gallery_id_seq'::regclass);


--
-- TOC entry 4713 (class 2604 OID 39064)
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- TOC entry 4711 (class 2604 OID 39051)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 4874 (class 0 OID 39013)
-- Dependencies: 216
-- Data for Name: destinations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.destinations VALUES (1, 'Bali, Indonesia', 'https://example.com/bali.jpg', 'Tropical island with beautiful beaches and unique Balinese culture.');
INSERT INTO public.destinations VALUES (2, 'Kyoto, Japan', 'https://example.com/kyoto.jpg', 'Traditional Japanese city with many temples and gardens.');
INSERT INTO public.destinations VALUES (3, 'Paris, France', 'https://example.com/paris.jpg', 'Romantic city with the Eiffel Tower and rich art.');
INSERT INTO public.destinations VALUES (4, 'Rome, Italy', 'https://example.com/rome.jpg', 'City with ancient history, Colosseum, and famous museums.');
INSERT INTO public.destinations VALUES (5, 'Sydney, Australia', 'https://example.com/sydney.jpg', 'Coastal city with the Opera House and Bondi Beach.');
INSERT INTO public.destinations VALUES (6, 'New York, USA', 'https://example.com/ny.jpg', 'Modern city with skyscrapers and Central Park.');
INSERT INTO public.destinations VALUES (7, 'London, UK', 'https://example.com/london.jpg', 'City of history and art with Big Ben and Buckingham Palace.');
INSERT INTO public.destinations VALUES (8, 'Dubai, UAE', 'https://example.com/dubai.jpg', 'Futuristic city with the world''s tallest building, Burj Khalifa.');
INSERT INTO public.destinations VALUES (9, 'Bangkok, Thailand', 'https://example.com/bangkok.jpg', 'Southeast Asian city with temples and traditional markets.');
INSERT INTO public.destinations VALUES (10, 'Istanbul, Turkey', 'https://example.com/istanbul.jpg', 'City blending European and Asian cultures with historical mosques.');


--
-- TOC entry 4878 (class 0 OID 39036)
-- Dependencies: 220
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
-- TOC entry 4876 (class 0 OID 39022)
-- Dependencies: 218
-- Data for Name: gallery; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gallery VALUES (1, 1, 'https://example.com/bali_beach.jpg', 'Sunset on a beautiful beach in Bali.');
INSERT INTO public.gallery VALUES (2, 1, 'https://example.com/bali_rice.jpg', 'Terraced rice fields in Ubud, Bali.');
INSERT INTO public.gallery VALUES (3, 2, 'https://example.com/kyoto_temple.jpg', 'Historic temple surrounded by cherry blossoms in Kyoto.');
INSERT INTO public.gallery VALUES (4, 2, 'https://example.com/kyoto_garden.jpg', 'Serene Japanese garden in Kyoto.');
INSERT INTO public.gallery VALUES (5, 3, 'https://example.com/paris_eiffel.jpg', 'The Eiffel Tower lit up at night in Paris.');
INSERT INTO public.gallery VALUES (6, 3, 'https://example.com/paris_louvre.jpg', 'The Louvre Museum with its famous glass pyramid.');
INSERT INTO public.gallery VALUES (7, 4, 'https://example.com/rome_colosseum.jpg', 'The majestic Colosseum in Rome.');
INSERT INTO public.gallery VALUES (8, 4, 'https://example.com/rome_vatican.jpg', 'View of St. Peter’s Basilica in the Vatican.');
INSERT INTO public.gallery VALUES (9, 5, 'https://example.com/sydney_opera.jpg', 'The Sydney Opera House illuminated in the evening.');
INSERT INTO public.gallery VALUES (10, 5, 'https://example.com/sydney_beach.jpg', 'Bondi Beach with surfers catching waves.');
INSERT INTO public.gallery VALUES (11, 6, 'https://example.com/ny_skyline.jpg', 'New York City skyline with the Empire State Building.');
INSERT INTO public.gallery VALUES (12, 6, 'https://example.com/ny_central_park.jpg', 'Central Park in autumn with vibrant foliage.');
INSERT INTO public.gallery VALUES (13, 7, 'https://example.com/london_bigben.jpg', 'Big Ben and the Houses of Parliament in London.');
INSERT INTO public.gallery VALUES (14, 7, 'https://example.com/london_bridge.jpg', 'Tower Bridge over the River Thames in London.');
INSERT INTO public.gallery VALUES (15, 8, 'https://example.com/dubai_burj.jpg', 'The iconic Burj Khalifa reaching into the sky in Dubai.');
INSERT INTO public.gallery VALUES (16, 8, 'https://example.com/dubai_marina.jpg', 'Dubai Marina with modern skyscrapers and yachts.');
INSERT INTO public.gallery VALUES (17, 9, 'https://example.com/bangkok_temple.jpg', 'The Golden Temple of Dawn on the Chao Phraya River.');
INSERT INTO public.gallery VALUES (18, 9, 'https://example.com/bangkok_market.jpg', 'Floating market in Bangkok with boats full of fruits and goods.');
INSERT INTO public.gallery VALUES (19, 10, 'https://example.com/istanbul_mosque.jpg', 'The Blue Mosque with its beautiful domes and minarets in Istanbul.');
INSERT INTO public.gallery VALUES (20, 10, 'https://example.com/istanbul_bosphorus.jpg', 'Bosphorus Strait with Istanbul’s cityscape in the background.');


--
-- TOC entry 4882 (class 0 OID 39061)
-- Dependencies: 224
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
-- TOC entry 4880 (class 0 OID 39048)
-- Dependencies: 222
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES (1, 'John Doe', 'john.doe@example.com', '1234567890', 'Looking forward to it!', 1, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (2, 'Jane Smith', 'jane.smith@example.com', '2345678901', 'Can''t wait to visit.', 2, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (3, 'Carlos Rodriguez', 'carlos.rod@example.com', '3456789012', 'Excited for the art tour.', 3, false, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (4, 'Mia Chen', 'mia.chen@example.com', '4567890123', 'Hope it''s worth the price.', 4, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (5, 'Sam Wilson', 'sam.wilson@example.com', '5678901234', 'Ready for an adventure!', 5, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (6, 'Emily Johnson', 'emily.john@example.com', '6789012345', 'Always wanted to see a show.', 6, false, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (7, 'Daniel Brown', 'dan.brown@example.com', '7890123456', 'History lover here.', 7, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (8, 'Sara Lee', 'sara.lee@example.com', '8901234567', 'First time in the desert!', 8, true, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (9, 'Tom Davis', 'tom.davis@example.com', '9012345678', 'Food and culture.', 9, false, '2024-11-14 21:10:08.178728');
INSERT INTO public.transactions VALUES (10, 'Alex Kim', 'alex.kim@example.com', '0123456789', 'Heard so much about it!', 10, true, '2024-11-14 21:10:08.178728');


--
-- TOC entry 4895 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.destinations_id_seq', 10, true);


--
-- TOC entry 4896 (class 0 OID 0)
-- Dependencies: 219
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 10, true);


--
-- TOC entry 4897 (class 0 OID 0)
-- Dependencies: 217
-- Name: gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gallery_id_seq', 20, true);


--
-- TOC entry 4898 (class 0 OID 0)
-- Dependencies: 223
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 10, true);


--
-- TOC entry 4899 (class 0 OID 0)
-- Dependencies: 221
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 10, true);


--
-- TOC entry 4716 (class 2606 OID 39020)
-- Name: destinations destinations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations
    ADD CONSTRAINT destinations_pkey PRIMARY KEY (id);


--
-- TOC entry 4720 (class 2606 OID 39041)
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- TOC entry 4718 (class 2606 OID 39029)
-- Name: gallery gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_pkey PRIMARY KEY (id);


--
-- TOC entry 4724 (class 2606 OID 39067)
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- TOC entry 4722 (class 2606 OID 39054)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 4726 (class 2606 OID 39042)
-- Name: events events_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4725 (class 2606 OID 39030)
-- Name: gallery gallery_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4728 (class 2606 OID 39068)
-- Name: reviews reviews_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4729 (class 2606 OID 39073)
-- Name: reviews reviews_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- TOC entry 4727 (class 2606 OID 39055)
-- Name: transactions transactions_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


-- Completed on 2024-11-14 21:11:39

--
-- PostgreSQL database dump complete
--

