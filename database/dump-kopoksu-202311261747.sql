--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

-- Started on 2023-11-26 17:47:45

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

DROP DATABASE kopoksu;
--
-- TOC entry 3029 (class 1262 OID 26889)
-- Name: kopoksu; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE kopoksu WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';


ALTER DATABASE kopoksu OWNER TO postgres;

\connect kopoksu

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
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3030 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 202 (class 1259 OID 26908)
-- Name: cart; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart (
    id uuid NOT NULL,
    total integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.cart OWNER TO postgres;

--
-- TOC entry 200 (class 1259 OID 26890)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 26913)
-- Name: detail_cart; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detail_cart (
    id uuid NOT NULL,
    product_id uuid NOT NULL,
    cart_id uuid NOT NULL,
    quantity integer NOT NULL,
    total integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.detail_cart OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 26928)
-- Name: offline_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.offline_order (
    id uuid NOT NULL,
    cart_id uuid NOT NULL,
    name character varying NOT NULL,
    phone_number character varying NOT NULL,
    total integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.offline_order OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 26941)
-- Name: online_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.online_order (
    id uuid NOT NULL,
    cart_id uuid NOT NULL,
    name character varying NOT NULL,
    address character varying NOT NULL,
    city character varying NOT NULL,
    province character varying NOT NULL,
    phone_number character varying NOT NULL,
    post_code integer NOT NULL,
    total integer NOT NULL,
    cost integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.online_order OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 26895)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying NOT NULL,
    description character varying NOT NULL,
    quantity integer NOT NULL,
    price integer NOT NULL,
    image character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    category_id uuid NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 3020 (class 0 OID 26908)
-- Dependencies: 202
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3018 (class 0 OID 26890)
-- Dependencies: 200
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3021 (class 0 OID 26913)
-- Dependencies: 203
-- Data for Name: detail_cart; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3022 (class 0 OID 26928)
-- Dependencies: 204
-- Data for Name: offline_order; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3023 (class 0 OID 26941)
-- Dependencies: 205
-- Data for Name: online_order; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3019 (class 0 OID 26895)
-- Dependencies: 201
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 2876 (class 2606 OID 26912)
-- Name: cart cart_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart
    ADD CONSTRAINT cart_pk PRIMARY KEY (id);


--
-- TOC entry 2872 (class 2606 OID 26894)
-- Name: categories categories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pk PRIMARY KEY (id);


--
-- TOC entry 2878 (class 2606 OID 26917)
-- Name: detail_cart detail_cart_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_cart
    ADD CONSTRAINT detail_cart_pk PRIMARY KEY (id);


--
-- TOC entry 2880 (class 2606 OID 26935)
-- Name: offline_order offline_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offline_order
    ADD CONSTRAINT offline_order_pk PRIMARY KEY (id);


--
-- TOC entry 2882 (class 2606 OID 26948)
-- Name: online_order online_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.online_order
    ADD CONSTRAINT online_order_pk PRIMARY KEY (id);


--
-- TOC entry 2874 (class 2606 OID 26902)
-- Name: products products_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pk PRIMARY KEY (id);


--
-- TOC entry 2884 (class 2606 OID 26918)
-- Name: detail_cart detail_cart_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_cart
    ADD CONSTRAINT detail_cart_fk FOREIGN KEY (cart_id) REFERENCES public.cart(id);


--
-- TOC entry 2885 (class 2606 OID 26923)
-- Name: detail_cart detail_cart_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_cart
    ADD CONSTRAINT detail_cart_fk_1 FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 2886 (class 2606 OID 26936)
-- Name: offline_order offline_order_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.offline_order
    ADD CONSTRAINT offline_order_fk FOREIGN KEY (cart_id) REFERENCES public.cart(id);


--
-- TOC entry 2887 (class 2606 OID 26949)
-- Name: online_order online_order_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.online_order
    ADD CONSTRAINT online_order_fk FOREIGN KEY (cart_id) REFERENCES public.cart(id);


--
-- TOC entry 2883 (class 2606 OID 26903)
-- Name: products products_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_fk FOREIGN KEY (category_id) REFERENCES public.categories(id);


-- Completed on 2023-11-26 17:47:45

--
-- PostgreSQL database dump complete
--

