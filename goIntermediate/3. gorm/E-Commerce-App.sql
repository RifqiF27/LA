--
-- PostgreSQL database dump
--

-- Dumped from database version 17.2
-- Dumped by pg_dump version 17.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: history_deliveries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.history_deliveries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_shipping_id bigint NOT NULL,
    status character varying(20) DEFAULT 'pending'::character varying NOT NULL,
    location character varying(100) NOT NULL,
    CONSTRAINT chk_history_deliveries_status CHECK (((status)::text = ANY ((ARRAY['pending'::character varying, 'shipped'::character varying, 'completed'::character varying, 'canceled'::character varying])::text[])))
);


ALTER TABLE public.history_deliveries OWNER TO postgres;

--
-- Name: history_deliveries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.history_deliveries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.history_deliveries_id_seq OWNER TO postgres;

--
-- Name: history_deliveries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.history_deliveries_id_seq OWNED BY public.history_deliveries.id;


--
-- Name: order_shippings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_shippings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id character varying(100),
    shipping_id bigint NOT NULL,
    origin_lat_long character varying(100) NOT NULL,
    destination_lat_long character varying(100) NOT NULL,
    total_payment numeric(10,2) NOT NULL
);


ALTER TABLE public.order_shippings OWNER TO postgres;

--
-- Name: order_shippings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.order_shippings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.order_shippings_id_seq OWNER TO postgres;

--
-- Name: order_shippings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.order_shippings_id_seq OWNED BY public.order_shippings.id;


--
-- Name: shippings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shippings (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    price numeric(10,2) NOT NULL
);


ALTER TABLE public.shippings OWNER TO postgres;

--
-- Name: shippings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shippings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.shippings_id_seq OWNER TO postgres;

--
-- Name: shippings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shippings_id_seq OWNED BY public.shippings.id;


--
-- Name: history_deliveries id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.history_deliveries ALTER COLUMN id SET DEFAULT nextval('public.history_deliveries_id_seq'::regclass);


--
-- Name: order_shippings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_shippings ALTER COLUMN id SET DEFAULT nextval('public.order_shippings_id_seq'::regclass);


--
-- Name: shippings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shippings ALTER COLUMN id SET DEFAULT nextval('public.shippings_id_seq'::regclass);


--
-- Data for Name: history_deliveries; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.history_deliveries (id, created_at, updated_at, deleted_at, order_shipping_id, status, location) FROM stdin;
\.


--
-- Data for Name: order_shippings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_shippings (id, created_at, updated_at, deleted_at, order_id, shipping_id, origin_lat_long, destination_lat_long, total_payment) FROM stdin;
\.


--
-- Data for Name: shippings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shippings (id, created_at, updated_at, deleted_at, name, price) FROM stdin;
1	2024-11-28 22:55:55.994919+07	2024-11-28 22:55:55.994919+07	\N	Standard Shipping	5.00
2	2024-11-28 22:55:55.994919+07	2024-11-28 22:55:55.994919+07	\N	Express Shipping	15.00
3	2024-11-28 22:55:55.994919+07	2024-11-28 22:55:55.994919+07	\N	Overnight Shipping	25.00
\.


--
-- Name: history_deliveries_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.history_deliveries_id_seq', 1, false);


--
-- Name: order_shippings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.order_shippings_id_seq', 1, false);


--
-- Name: shippings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shippings_id_seq', 3, true);


--
-- Name: history_deliveries history_deliveries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.history_deliveries
    ADD CONSTRAINT history_deliveries_pkey PRIMARY KEY (id);


--
-- Name: order_shippings order_shippings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_shippings
    ADD CONSTRAINT order_shippings_pkey PRIMARY KEY (id);


--
-- Name: shippings shippings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shippings
    ADD CONSTRAINT shippings_pkey PRIMARY KEY (id);


--
-- Name: idx_history_deliveries_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_history_deliveries_deleted_at ON public.history_deliveries USING btree (deleted_at);


--
-- Name: idx_order_shippings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_order_shippings_deleted_at ON public.order_shippings USING btree (deleted_at);


--
-- Name: idx_shippings_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shippings_deleted_at ON public.shippings USING btree (deleted_at);


--
-- Name: history_deliveries fk_history_deliveries_order_shipping; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.history_deliveries
    ADD CONSTRAINT fk_history_deliveries_order_shipping FOREIGN KEY (order_shipping_id) REFERENCES public.order_shippings(id);


--
-- Name: order_shippings fk_order_shippings_shipping; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_shippings
    ADD CONSTRAINT fk_order_shippings_shipping FOREIGN KEY (shipping_id) REFERENCES public.shippings(id);


--
-- PostgreSQL database dump complete
--

