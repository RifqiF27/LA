--
-- PostgreSQL database dump
--

-- Dumped from database version 16rc1
-- Dumped by pg_dump version 16rc1

-- Started on 2024-11-15 21:08:03

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
-- TOC entry 4913 (class 1262 OID 37596)
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
-- TOC entry 4914 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: pg_database_owner
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 216 (class 1259 OID 39550)
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
-- TOC entry 215 (class 1259 OID 39549)
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
-- TOC entry 4915 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.destinations_id_seq OWNED BY public.destinations.id;


--
-- TOC entry 222 (class 1259 OID 39587)
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
-- TOC entry 221 (class 1259 OID 39586)
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
-- TOC entry 4916 (class 0 OID 0)
-- Dependencies: 221
-- Name: events_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.events_id_seq OWNED BY public.events.id;


--
-- TOC entry 220 (class 1259 OID 39573)
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
-- TOC entry 219 (class 1259 OID 39572)
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
-- TOC entry 4917 (class 0 OID 0)
-- Dependencies: 219
-- Name: gallery_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.gallery_id_seq OWNED BY public.gallery.id;


--
-- TOC entry 218 (class 1259 OID 39559)
-- Name: locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.locations (
    id integer NOT NULL,
    destination_id integer NOT NULL,
    summary character varying(255) NOT NULL,
    longlat point NOT NULL,
    detail text NOT NULL
);


ALTER TABLE public.locations OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 39558)
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.locations_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.locations_id_seq OWNER TO postgres;

--
-- TOC entry 4918 (class 0 OID 0)
-- Dependencies: 217
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


--
-- TOC entry 228 (class 1259 OID 39633)
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
-- TOC entry 227 (class 1259 OID 39632)
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
-- TOC entry 4919 (class 0 OID 0)
-- Dependencies: 227
-- Name: reviews_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.reviews_id_seq OWNED BY public.reviews.id;


--
-- TOC entry 224 (class 1259 OID 39599)
-- Name: tour_plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tour_plans (
    id integer NOT NULL,
    destination_id integer NOT NULL,
    event_id integer NOT NULL,
    day integer NOT NULL,
    activity character varying(255) NOT NULL,
    facilities jsonb DEFAULT '{}'::jsonb,
    CONSTRAINT tour_plans_day_check CHECK ((day > 0))
);


ALTER TABLE public.tour_plans OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 39598)
-- Name: tour_plans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tour_plans_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tour_plans_id_seq OWNER TO postgres;

--
-- TOC entry 4920 (class 0 OID 0)
-- Dependencies: 223
-- Name: tour_plans_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tour_plans_id_seq OWNED BY public.tour_plans.id;


--
-- TOC entry 226 (class 1259 OID 39620)
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
-- TOC entry 225 (class 1259 OID 39619)
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
-- TOC entry 4921 (class 0 OID 0)
-- Dependencies: 225
-- Name: transactions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transactions_id_seq OWNED BY public.transactions.id;


--
-- TOC entry 4718 (class 2604 OID 39553)
-- Name: destinations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations ALTER COLUMN id SET DEFAULT nextval('public.destinations_id_seq'::regclass);


--
-- TOC entry 4721 (class 2604 OID 39590)
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.events_id_seq'::regclass);


--
-- TOC entry 4720 (class 2604 OID 39576)
-- Name: gallery id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery ALTER COLUMN id SET DEFAULT nextval('public.gallery_id_seq'::regclass);


--
-- TOC entry 4719 (class 2604 OID 39562)
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- TOC entry 4726 (class 2604 OID 39636)
-- Name: reviews id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews ALTER COLUMN id SET DEFAULT nextval('public.reviews_id_seq'::regclass);


--
-- TOC entry 4722 (class 2604 OID 39602)
-- Name: tour_plans id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tour_plans ALTER COLUMN id SET DEFAULT nextval('public.tour_plans_id_seq'::regclass);


--
-- TOC entry 4724 (class 2604 OID 39623)
-- Name: transactions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions ALTER COLUMN id SET DEFAULT nextval('public.transactions_id_seq'::regclass);


--
-- TOC entry 4895 (class 0 OID 39550)
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
-- TOC entry 4901 (class 0 OID 39587)
-- Dependencies: 222
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
-- TOC entry 4899 (class 0 OID 39573)
-- Dependencies: 220
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
-- TOC entry 4897 (class 0 OID 39559)
-- Dependencies: 218
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.locations VALUES (1, 1, 'Bali - A Tropical Paradise of Beautiful Beaches, Lush Green Rice Fields, and Vibrant Balinese Culture, Ideal for Relaxation and Adventure', '(115.216667,-8.65)', 'Bali is a tropical island that offers an unforgettable escape with its pristine beaches, terraced rice fields, and temples that reveal the island''s spiritual heritage. Bali has long been known as a paradise for tourists, thanks to its balance of relaxation and adventure.

From the lively atmosphere of Kuta Beach to the serene beauty of Ubud, Bali provides travelers with diverse experiences. Its unique culture, rich in traditions such as dance, art, and cuisine, makes it a vibrant destination for those seeking both relaxation and cultural exploration.');
INSERT INTO public.locations VALUES (2, 2, 'Kyoto - A City Steeped in Tradition with Beautiful Temples, Tranquil Gardens, and Historic Wooden Houses, Offering a Step Back in Time', '(135.7681,35.0116)', 'Kyoto is one of Japan''s most culturally significant cities, known for its stunning temples, shrines, and well-preserved traditional wooden houses. This city, which once served as Japan''s capital, remains a place where history and culture come alive in every street and corner.

Visitors can explore iconic locations like the Golden Pavilion, Fushimi Inari Shrine, and the Arashiyama Bamboo Grove. Kyoto also offers tranquil tea ceremonies, beautiful Zen gardens, and seasonal cherry blossoms, making it a place where the old and new blend harmoniously.');
INSERT INTO public.locations VALUES (3, 3, 'Paris - The City of Romance, Art, and Iconic Landmarks, Full of Rich Culture, Stunning Architecture, and Unmatched Beauty', '(2.3522,48.8566)', 'Paris is the ultimate destination for lovers of art, history, and romance. Known as the City of Light, Paris is home to world-renowned landmarks like the Eiffel Tower, the Louvre Museum, and the historic Notre-Dame Cathedral. The city''s allure lies in its ability to captivate travelers with both its famous sights and hidden gems.

Beyond its iconic landmarks, Paris is a hub for artistic expression, from the works of the French Impressionists to contemporary art in the Le Marais district. Strolling along the Seine River, indulging in delicious pastries, and visiting Montmartre’s bohemian streets are just a few of the experiences that make Paris so irresistible.');
INSERT INTO public.locations VALUES (4, 4, 'Rome - An Open-Air Museum with Ancient Ruins, Magnificent Palaces, and Stunning Architecture, Offering a Deep Dive into Roman History', '(12.4964,41.9028)', 'Rome is one of the most historically significant cities in the world, a true open-air museum with monuments dating back thousands of years. The city''s iconic landmarks, such as the Colosseum, Roman Forum, and Pantheon, offer visitors a glimpse into the grandeur of the Roman Empire.

Rome is more than just its ancient ruins. It is a city full of vibrant neighborhoods, bustling piazzas, and delicious food. The Vatican, home to St. Peter’s Basilica and the Sistine Chapel, adds a spiritual layer to the city’s cultural richness, making it a must-visit for history buffs and spiritual seekers alike.');
INSERT INTO public.locations VALUES (5, 5, 'Sydney - A Coastal Metropolis with Stunning Views, Beautiful Beaches, and Modern Landmarks Like the Opera House and Harbour Bridge', '(151.2093,-33.8688)', 'Sydney is one of the most vibrant and scenic cities in the world. Set along the stunning coastline of Australia, Sydney is famous for its sparkling beaches, bustling harbors, and iconic landmarks. Visitors flock to the Sydney Opera House, the Sydney Harbour Bridge, and Bondi Beach, all of which embody the spirit of this lively city.

But Sydney is more than just its famous sights. It is a melting pot of cultures, with neighborhoods like Chinatown, Darlinghurst, and Newtown offering diverse experiences. Sydney’s natural beauty, combined with its dynamic urban culture, makes it a must-visit destination for travelers.');
INSERT INTO public.locations VALUES (6, 6, 'New York - The City That Never Sleeps, Filled with Iconic Landmarks, Cultural Diversity, and a Constant Energy That Inspires Millions', '(-74.006,40.7128)', 'New York City is a global hub of culture, finance, and entertainment. Known as “The City That Never Sleeps,” it is a place where opportunities and dreams abound. From Times Square’s neon lights to the tranquility of Central Park, New York offers a unique contrast between the fast-paced urban landscape and moments of calm.

The city is home to famous landmarks like the Empire State Building, Statue of Liberty, and Broadway theatres, making it a mecca for tourists and residents alike. Each of its boroughs – from Manhattan to Brooklyn – offers a distinct experience, making New York an ever-evolving metropolis that never loses its charm.');
INSERT INTO public.locations VALUES (7, 7, 'London - A Historical and Cultural Capital with Iconic Monuments, Museums, and Royal Heritage, Where Tradition Meets Modernity', '(-0.1276,51.5074)', 'London is a city where the past and the present exist side by side. From the iconic Tower of London to Buckingham Palace, the city is rich in history and culture, offering visitors a chance to immerse themselves in the British way of life. Whether you’re walking through the grand halls of the British Museum or catching a West End show, London provides an endless list of experiences.

Beyond its historical landmarks, London is a cosmopolitan city with diverse neighborhoods, from the trendy shops of Soho to the historic streets of Covent Garden. The city’s arts scene is world-renowned, and its green spaces, like Hyde Park, offer peaceful retreats from the hustle and bustle of city life.');
INSERT INTO public.locations VALUES (8, 8, 'Dubai - A Modern Metropolis in the Heart of the Desert with Futuristic Architecture, Luxury Shopping, and World-Class Entertainment', '(55.2708,25.2048)', 'Dubai is a city that blends innovation with tradition. Known for its skyline filled with towering skyscrapers like the Burj Khalifa and luxurious shopping malls, Dubai offers a taste of the future in the heart of the desert. The city is a playground for those who love high-end living, thrilling activities, and cutting-edge architecture.

Dubai’s diversity is also evident in its cultural offerings, from the bustling souks in the old town to the modern art galleries in Alserkal Avenue. Visitors can explore the nearby desert with thrilling safaris or relax on the beaches along the Arabian Gulf.');
INSERT INTO public.locations VALUES (9, 9, 'Bangkok - A Bustling City with Temples, Markets, and a Thriving Street Food Scene, Blending Tradition and Modern Life', '(100.5018,13.7563)', 'Bangkok is a city of contrasts, where ancient traditions coexist with modern life. The city is famous for its golden temples, bustling markets, and vibrant nightlife. A visit to the Grand Palace, Wat Arun, or the floating markets offers a glimpse into Thailand’s rich cultural heritage.

But Bangkok is also a modern metropolis, with high-rise buildings, luxury malls, and trendy cafes. The city’s street food scene is legendary, with vendors offering everything from spicy curries to sweet treats. Whether you’re exploring the markets by boat or enjoying a rooftop cocktail, Bangkok’s energy is infectious.');
INSERT INTO public.locations VALUES (10, 10, 'Istanbul - A City at the Crossroads of Two Continents with Rich History, Magnificent Mosques, and a Unique Blend of Cultures', '(28.9784,41.0082)', 'Istanbul is a city where East meets West, where European and Asian influences blend into a unique cultural tapestry. From the grandeur of the Hagia Sophia to the vibrant bazaars, Istanbul offers a fascinating mix of history, art, and cuisine.

The city is a melting pot of cultures, with its stunning mosques like the Blue Mosque, intricate palaces, and delicious food. Visitors can take a boat ride along the Bosphorus, shop in the Grand Bazaar, or enjoy a traditional Turkish bath. Istanbul’s charm lies in its ability to offer something for everyone, whether you’re exploring its ancient sites or enjoying its modern cafes.');


--
-- TOC entry 4907 (class 0 OID 39633)
-- Dependencies: 228
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
-- TOC entry 4903 (class 0 OID 39599)
-- Dependencies: 224
-- Data for Name: tour_plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.tour_plans VALUES (1, 1, 1, 1, 'Arrival and check-in', '{"meals": "Dinner", "transport": "Private Car", "accommodation": "5 Star Resort"}');
INSERT INTO public.tour_plans VALUES (2, 1, 1, 2, 'Beach day and surfing', '{"guide": "Surfing Instructor", "meals": "Lunch and Dinner", "activities": ["Surfing", "Beach volleyball"]}');
INSERT INTO public.tour_plans VALUES (3, 2, 2, 1, 'Temple visit', '{"meals": "Lunch", "transport": "Private Bus", "accommodation": "Traditional Ryokan"}');
INSERT INTO public.tour_plans VALUES (4, 2, 2, 2, 'Cherry blossom tour', '{"meals": "Breakfast and Dinner", "transport": "Private Walking Tour", "activities": ["Photography", "Walking"]}');
INSERT INTO public.tour_plans VALUES (5, 3, 3, 1, 'Seine River cruise', '{"meals": "Lunch", "transport": "Private Boat", "accommodation": "Boutique Hotel"}');
INSERT INTO public.tour_plans VALUES (6, 3, 3, 2, 'Art museum tour', '{"meals": "Lunch", "transport": "City Metro", "activities": ["Museum Tour", "Photography"]}');
INSERT INTO public.tour_plans VALUES (7, 4, 4, 1, 'Colosseum visit', '{"guide": "History Expert", "meals": "Breakfast", "transport": "Private Coach"}');
INSERT INTO public.tour_plans VALUES (8, 4, 4, 2, 'Vatican Museums tour', '{"meals": "Dinner", "transport": "Shuttle", "accommodation": "Luxury Hotel"}');
INSERT INTO public.tour_plans VALUES (9, 5, 5, 1, 'Sydney Opera House tour', '{"meals": "Lunch and Dinner", "activities": ["Opera House Tour", "Local Exploration"], "accommodation": "Seaside Hotel"}');
INSERT INTO public.tour_plans VALUES (10, 5, 5, 2, 'Harbour cruise', '{"meals": "Dinner", "transport": "Luxury Yacht", "activities": ["Sightseeing", "Dinner on cruise"]}');
INSERT INTO public.tour_plans VALUES (11, 6, 6, 1, 'Broadway Show in NYC', '{"meals": "Dinner", "activities": ["Show Viewing"], "accommodation": "Central Park Hotel"}');
INSERT INTO public.tour_plans VALUES (12, 6, 6, 2, 'Central Park Exploration', '{"meals": "Lunch", "transport": "Bicycle Rental", "activities": ["Walking", "Cycling"]}');
INSERT INTO public.tour_plans VALUES (13, 7, 7, 1, 'Big Ben and Tower of London Tour', '{"meals": "Lunch", "transport": "Private Tour Bus", "accommodation": "City Hotel"}');
INSERT INTO public.tour_plans VALUES (14, 7, 7, 2, 'River Thames Cruise', '{"meals": "Dinner", "transport": "Private Boat", "activities": ["Boat Cruise"]}');
INSERT INTO public.tour_plans VALUES (15, 8, 8, 1, 'Desert Safari', '{"meals": "Lunch and Dinner", "activities": ["Sandboarding", "Camel Ride"], "accommodation": "Desert Camp"}');
INSERT INTO public.tour_plans VALUES (16, 8, 8, 2, 'Burj Khalifa Observation Deck', '{"meals": "Dinner", "transport": "Private Transfer", "activities": ["Sightseeing", "Photography"]}');
INSERT INTO public.tour_plans VALUES (17, 9, 9, 1, 'Floating Market Exploration', '{"meals": "Breakfast and Lunch", "activities": ["Boating", "Market Tour"], "accommodation": "Local Guesthouse"}');
INSERT INTO public.tour_plans VALUES (18, 9, 9, 2, 'Grand Palace Visit', '{"meals": "Lunch", "transport": "Private Van", "activities": ["Sightseeing", "Photography"]}');
INSERT INTO public.tour_plans VALUES (19, 10, 10, 1, 'Hagia Sophia Visit', '{"meals": "Lunch", "activities": ["Cultural Tour", "Photography"], "accommodation": "Boutique Hotel"}');
INSERT INTO public.tour_plans VALUES (20, 10, 10, 2, 'Bosphorus Cruise', '{"meals": "Dinner", "transport": "Private Yacht", "activities": ["Sightseeing", "Boat Tour"]}');


--
-- TOC entry 4905 (class 0 OID 39620)
-- Dependencies: 226
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transactions VALUES (1, 'John Doe', 'john.doe@example.com', '1234567890', 'Looking forward to it!', 1, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (2, 'Jane Smith', 'jane.smith@example.com', '2345678901', 'Can''t wait to visit.', 2, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (3, 'Carlos Rodriguez', 'carlos.rod@example.com', '3456789012', 'Excited for the art tour.', 3, false, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (4, 'Mia Chen', 'mia.chen@example.com', '4567890123', 'Hope it''s worth the price.', 4, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (5, 'Sam Wilson', 'sam.wilson@example.com', '5678901234', 'Ready for an adventure!', 5, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (6, 'Emily Johnson', 'emily.john@example.com', '6789012345', 'Always wanted to see a show.', 6, false, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (7, 'Daniel Brown', 'dan.brown@example.com', '7890123456', 'History lover here.', 7, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (8, 'Sara Lee', 'sara.lee@example.com', '8901234567', 'First time in the desert!', 8, true, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (9, 'Tom Davis', 'tom.davis@example.com', '9012345678', 'Food and culture.', 9, false, '2024-11-15 21:06:36.034746');
INSERT INTO public.transactions VALUES (10, 'Alex Kim', 'alex.kim@example.com', '0123456789', 'Heard so much about it!', 10, true, '2024-11-15 21:06:36.034746');


--
-- TOC entry 4922 (class 0 OID 0)
-- Dependencies: 215
-- Name: destinations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.destinations_id_seq', 10, true);


--
-- TOC entry 4923 (class 0 OID 0)
-- Dependencies: 221
-- Name: events_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.events_id_seq', 10, true);


--
-- TOC entry 4924 (class 0 OID 0)
-- Dependencies: 219
-- Name: gallery_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gallery_id_seq', 20, true);


--
-- TOC entry 4925 (class 0 OID 0)
-- Dependencies: 217
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_id_seq', 10, true);


--
-- TOC entry 4926 (class 0 OID 0)
-- Dependencies: 227
-- Name: reviews_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.reviews_id_seq', 10, true);


--
-- TOC entry 4927 (class 0 OID 0)
-- Dependencies: 223
-- Name: tour_plans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tour_plans_id_seq', 20, true);


--
-- TOC entry 4928 (class 0 OID 0)
-- Dependencies: 225
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 10, true);


--
-- TOC entry 4730 (class 2606 OID 39557)
-- Name: destinations destinations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.destinations
    ADD CONSTRAINT destinations_pkey PRIMARY KEY (id);


--
-- TOC entry 4736 (class 2606 OID 39592)
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);


--
-- TOC entry 4734 (class 2606 OID 39580)
-- Name: gallery gallery_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_pkey PRIMARY KEY (id);


--
-- TOC entry 4732 (class 2606 OID 39566)
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- TOC entry 4742 (class 2606 OID 39639)
-- Name: reviews reviews_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_pkey PRIMARY KEY (id);


--
-- TOC entry 4738 (class 2606 OID 39608)
-- Name: tour_plans tour_plans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tour_plans
    ADD CONSTRAINT tour_plans_pkey PRIMARY KEY (id);


--
-- TOC entry 4740 (class 2606 OID 39626)
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- TOC entry 4745 (class 2606 OID 39593)
-- Name: events events_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4744 (class 2606 OID 39581)
-- Name: gallery gallery_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.gallery
    ADD CONSTRAINT gallery_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4743 (class 2606 OID 39567)
-- Name: locations locations_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4749 (class 2606 OID 39640)
-- Name: reviews reviews_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4750 (class 2606 OID 39645)
-- Name: reviews reviews_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.reviews
    ADD CONSTRAINT reviews_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id);


--
-- TOC entry 4746 (class 2606 OID 39609)
-- Name: tour_plans tour_plans_destination_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tour_plans
    ADD CONSTRAINT tour_plans_destination_id_fkey FOREIGN KEY (destination_id) REFERENCES public.destinations(id);


--
-- TOC entry 4747 (class 2606 OID 39614)
-- Name: tour_plans tour_plans_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tour_plans
    ADD CONSTRAINT tour_plans_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


--
-- TOC entry 4748 (class 2606 OID 39627)
-- Name: transactions transactions_event_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id);


-- Completed on 2024-11-15 21:08:03

--
-- PostgreSQL database dump complete
--

