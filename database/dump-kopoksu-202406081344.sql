--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4
-- Dumped by pg_dump version 13.4

-- Started on 2024-06-08 13:44:46

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
-- TOC entry 3030 (class 1262 OID 26889)
-- Name: kopoksu; Type: DATABASE; Schema: -; Owner: postgres
--

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

ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3031 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 200 (class 1259 OID 26890)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name character varying NOT NULL
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 26990)
-- Name: detail_online_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detail_online_orders (
    id uuid NOT NULL,
    online_order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    amount integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.detail_online_orders OWNER TO postgres;

--
-- TOC entry 204 (class 1259 OID 26957)
-- Name: detail_pickup_online_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.detail_pickup_online_orders (
    id uuid NOT NULL,
    pickup_online_order_id uuid NOT NULL,
    product_id uuid NOT NULL,
    amount integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone NOT NULL
);


ALTER TABLE public.detail_pickup_online_orders OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 26941)
-- Name: online_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.online_orders (
    id uuid NOT NULL,
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


ALTER TABLE public.online_orders OWNER TO postgres;

--
-- TOC entry 202 (class 1259 OID 26928)
-- Name: pickup_online_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pickup_online_orders (
    id uuid NOT NULL,
    name character varying NOT NULL,
    phone_number character varying NOT NULL,
    total integer NOT NULL,
    status character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    pickup_date timestamp without time zone
);


ALTER TABLE public.pickup_online_orders OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 26895)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid NOT NULL,
    name character varying NOT NULL,
    description text NOT NULL,
    quantity integer NOT NULL,
    price integer NOT NULL,
    image character varying NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    category_id uuid NOT NULL,
    buy_price integer,
    weight integer
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 3019 (class 0 OID 26890)
-- Dependencies: 200
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.categories VALUES ('ea600c63-283a-415e-8ed1-b10d12c544a0', '2023-11-27 20:03:53.232198+07', '2023-11-27 20:03:53.232198+07', 'Susu Formula');
INSERT INTO public.categories VALUES ('981464fb-3241-4a33-97ae-33b110e2d4aa', '2023-11-27 20:04:59.08845+07', '2023-11-27 20:04:59.08845+07', 'Popok Bayi');
INSERT INTO public.categories VALUES ('f5976ce9-7496-4fd2-8322-3beaef36e4d8', '2023-11-27 20:04:59.089719+07', '2023-11-27 20:04:59.089719+07', 'Popok Dewasa');


--
-- TOC entry 3024 (class 0 OID 26990)
-- Dependencies: 205
-- Data for Name: detail_online_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.detail_online_orders VALUES ('fac21ee7-029b-4012-938a-faab9ba9fd92', 'f43cd927-60c5-433d-a4a2-69649f895be3', '2b0d4f72-2639-4fd7-8e83-3d3b3d8354bd', 1, '2024-06-08 12:28:11.738377+07', '2024-06-08 12:28:11.738377+07');
INSERT INTO public.detail_online_orders VALUES ('380ff052-71cf-4276-a517-bdea1b05b802', 'f43cd927-60c5-433d-a4a2-69649f895be3', 'b81b6268-472c-465e-8b3f-6ff1cfbef3d3', 1, '2024-06-08 12:28:11.734+07', '2024-06-08 12:28:11.734324+07');


--
-- TOC entry 3023 (class 0 OID 26957)
-- Dependencies: 204
-- Data for Name: detail_pickup_online_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.detail_pickup_online_orders VALUES ('172d45dd-fcaa-464e-957e-7fc89ea4e115', '89c3793d-9e74-47ed-b243-c7115b855237', '54a5e212-2b78-4478-9c3f-b9a2fdd6e3c3', 1, '2024-06-08 13:33:18.197463+07', '2024-06-08 13:33:18.197463+07');
INSERT INTO public.detail_pickup_online_orders VALUES ('4141c40c-286c-44ea-ab01-8b91b72f3cea', '89c3793d-9e74-47ed-b243-c7115b855237', '59d6c6d2-296a-44f5-a173-2e72e15c355b', 1, '2024-06-08 13:33:18.200958+07', '2024-06-08 13:33:18.200958+07');


--
-- TOC entry 3022 (class 0 OID 26941)
-- Dependencies: 203
-- Data for Name: online_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.online_orders VALUES ('f43cd927-60c5-433d-a4a2-69649f895be3', 'Muchammad Gema Akbar', 'Jl. Raya Dramaga Kampus IPB Dramaga Bogor 16680 West Java, Indonesia', 'Surabaya', 'Jawa Barat', '082237436363', 16680, 218988, 88000, 'Pengiriman', '2024-06-08 12:28:11.732816+07', '2024-06-08 12:29:35.54328+07');


--
-- TOC entry 3021 (class 0 OID 26928)
-- Dependencies: 202
-- Data for Name: pickup_online_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.pickup_online_orders VALUES ('89c3793d-9e74-47ed-b243-c7115b855237', 'Muchammad Gema Akbar', '082237436363', 102025, 'Menunggu pengambilan', '2024-06-08 13:33:18.194601+07', '2024-06-08 13:34:21.73941+07', '2024-06-08 13:33:00');


--
-- TOC entry 3020 (class 0 OID 26895)
-- Dependencies: 201
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.products VALUES ('2b0d4f72-2639-4fd7-8e83-3d3b3d8354bd', 'MAKUKU DRY CARE PANTS M48', 'MAKUKU Dry Care merupakan popok hemat yang dilengkapi dengan 3 lapisan sirkulasi udara yang terdapat di lapisan atas, lapisan penyerap dan lapisan bawah. Perlu diketahui bahwa sirkulasi udara pada popok bayi itu berperan penting', 99, 74900, '/static/images/uploads/9b5ced2b-604f-4274-af7a-fae2e78b9524.jpg', '2024-02-08 16:28:20.521969+07', '2024-06-08 12:28:11.739163+07', '981464fb-3241-4a33-97ae-33b110e2d4aa', 71000, 1000);
INSERT INTO public.products VALUES ('54a5e212-2b78-4478-9c3f-b9a2fdd6e3c3', 'OTO PREMIUM ADULT DIAPERS XL6', 'OTO PREMIUM Diapers for Adult / Popok Dewasa Premium ( Daya Serap Lebih Banyak ) model Perekat ukuran XL isi 6 pcs
Type Premium, daya serap lebih banyak.
Berdaya serap tinggi
Permukaan popok sangat lembut serta tidak panas karena adanya sirkulasi udara
Pelindung perekat tidak merusak popok sehingga dapat direkatkan berulang kali
Model : Perekat', 99, 44000, '/static/images/uploads/179aed8a-96bb-46b4-9639-7c9455e7431e.jpg', '2024-02-08 17:12:37.822182+07', '2024-06-08 13:33:18.200445+07', 'f5976ce9-7496-4fd2-8322-3beaef36e4d8', 40000, 1000);
INSERT INTO public.products VALUES ('e37e5270-0e1d-4ccf-8d71-34c9a786ccf9', 'WECARE ADULT DIAPERS M10', 'WeCare Adult Diapers adalah popok dewasa tipe perekat yang efektif mengontrol kuman dalam air seni dan kotoran, serta mengontrol dan menetralisir bau tidak sedap. Popok ini aktif mengubah kebasahan ke dalam gel sehingga tidak akan terasa sangat basah karena ada super absorben pad. Terdapat indikator kebasahan guna untuk mendeteksi kapan popok harus diganti.', 100, 39500, '/static/images/uploads/5f007b3b-5b09-4949-8f57-b671f5cfca80.jpg', '2024-02-08 17:10:28.911886+07', '2024-06-08 12:25:55.873615+07', 'f5976ce9-7496-4fd2-8322-3beaef36e4d8', 37000, 1000);
INSERT INTO public.products VALUES ('b81b6268-472c-465e-8b3f-6ff1cfbef3d3', 'ZEE PLATINUM COKLAT 350 GR', 'Zee Platinum merupakan susu pertumbuhan untuk anak usia 3-12 tahun dibuat dari susu bubuk berkualitas dengan Nutriprocomplex + Promune Formula. Dilengkapi dengan Minyak Ikan, DHA, Omega 3 & 6, Kolin, Serat Pangan Inulin, Protein, 9 Asam Amino Esensial, Tinggi Kalsium, 12 Vitamin & 5 Mineral, sebagai nutrisi tepat untuk mendukung pertumbuhan dan menjaga daya tahan tubuh anak serta keluarga.', 99, 56000, '/static/images/uploads/5e379348-c88b-4b3a-ad55-e6ff5228c815.jpg', '2024-02-08 17:01:53.615006+07', '2024-06-08 12:28:11.73779+07', 'ea600c63-283a-415e-8ed1-b10d12c544a0', 25000, 1000);
INSERT INTO public.products VALUES ('59d6c6d2-296a-44f5-a173-2e72e15c355b', 'SENSI ADULT DIAPERS XL8', 'Sensi Adult Diapers atau Popok Dewasa merupakan popok untuk orang dewasa yang memiliki permukaan lembut serta sirkulasi udara sehingga tidak panas saat digunakan. Popok dewasa ini juga memiliki frontal tape atau perekat pelindung yang tidak merusak sehingga dapat direkatkan berulang-ulang. Popok ini cocok digunakan untuk manula, orang sakit, pasien operasi, wasir, diare, Ibu melahirkan, dan saat sedang diperjalanan.', 99, 58000, '/static/images/uploads/fef74bb4-def5-46bc-87f2-cc5053852f5c.jpg', '2024-02-08 17:11:28.759478+07', '2024-06-08 13:33:18.202222+07', 'f5976ce9-7496-4fd2-8322-3beaef36e4d8', 55000, 1000);
INSERT INTO public.products VALUES ('8ce2faae-05b7-41b1-8c04-ceaafa1a7858', 'FF PRIMAGRO 1+ MADU 3KG', 'Susu pertumbuhan FRISIAN FLAG PRIMAGRO 1+ diformulasikan khusus untuk anak usia 1-3 tahun. Hadir dengan 9 AAE dan DHA 4x Lebih Tinggi maksimalkan tumbuh kembang si Kecil sehingga si Kecil berakal kreatif, tangkas, dan berani! Kini lebih sehat dengan gula lebih rendah.', 100, 309000, '/static/images/uploads/40dd3798-87f8-4a3b-bf5e-a0c9d8c5aa41.jpg', '2024-02-08 16:56:22.132266+07', '2024-06-08 12:25:04.352781+07', 'ea600c63-283a-415e-8ed1-b10d12c544a0', 300000, 1000);
INSERT INTO public.products VALUES ('98ebdb47-671c-4913-875f-2adc15597776', 'SWEETY GOLD PANTS S36', 'Sweety Gold Pants Popok Bayi S 36 Popok bayi sekali pakai dengan model celana yang mudah digunakan dan mudah untuk dibersihkan. Sweety Gold Pants ini memiliki banyak keunggulan yang membuat ayah dan bunda tidak perlu khawatir.', 100, 69000, '/static/images/uploads/d4d1859b-9c9f-4066-99d8-6460ee768286.jpg', '2024-02-08 16:43:38.475041+07', '2024-02-13 17:30:21.173239+07', '981464fb-3241-4a33-97ae-33b110e2d4aa', 66000, 1000);
INSERT INTO public.products VALUES ('123be81d-38be-4a7f-a7b9-63bfdae497af', 'CONFIDENCE CLASSIC DAY M8', 'Confidence Classic Day adalah produk Confidence dengan inovasi terbaru dan pertama di Indonesia dengan perlindungan di siang hari yang dapat menyerap extra cepat sehingga permukaan cepat kering. Confidence Classic Day memiliki 4x Daya Serap yang dapat digunakan untuk penderita Inkontinensia Berat dan Ultra Soft Cover yang membuat permukaan extra lembut dan tetap kering ketika dipakai. Confidence Classic Day juga memiliki Frontal Tape & Resealable Tape yang dapat merekat erat, dapat dilepas dan dipasang berulang kali.', 100, 38000, '/static/images/uploads/4820cad8-4db8-47e2-ab6b-16eec84f44c2.jpg', '2024-02-08 17:13:10.782027+07', '2024-06-08 12:26:00.65673+07', 'f5976ce9-7496-4fd2-8322-3beaef36e4d8', 35000, 500);
INSERT INTO public.products VALUES ('42f0f0f6-2b85-4003-9c7e-f9f8f5721ee7', 'MAKUKU SAP DIAPERS GROW CARE XXXXL12', 'Popok New Generation SAP Pertama di Indonesia, dengan struktur yang stabil membuat popok Makuku #AntiGumpal sehingga membuat bayi tidak merasa berat saat bergerak. Bahan SAP yang cepat menyerap sekaligus mengunci cairan dengan merata, dapat melindungi pantat dari iritasi dan ruam popok dengan ukuran JUMBO XXXXL untuk berat 25-45 kg.', 100, 99900, '/static/images/uploads/3965cc19-3760-4500-a81e-85636f4a5dd3.jpg', '2024-02-08 16:37:00.663483+07', '2024-06-08 12:25:34.072216+07', '981464fb-3241-4a33-97ae-33b110e2d4aa', 95000, 500);
INSERT INTO public.products VALUES ('eb75f779-ca85-4b66-8d8e-2e07df477d5a', 'GOON SMILE BABY COMFORT FIT XXL 24', 'GOO.N Smile Baby Comfort Fit XXL-24. Popok sekali pakai berbentuk celana ini ekonomis namun sangat nyaman dan mudah untuk digunakan.

Memiliki daya tampung hingga 12 Jam penggunaan untuk siang dan malam
Ukuran popok lebih besar, disesuaikan dengan bentuk tubuh si kecil
Popok lebih fit, sehingga mudah dan nyaman dipakai', 100, 55000, '/static/images/uploads/a09c4f28-07eb-4893-bb62-53e3b1821112.jpg', '2024-02-08 16:42:47.516925+07', '2024-02-14 14:21:56.033875+07', '981464fb-3241-4a33-97ae-33b110e2d4aa', 50000, 100);
INSERT INTO public.products VALUES ('6fde8684-4d2c-41e9-9a03-18f2c6956216', 'MAKUKU SAP DIAPERS SLIM CARE PANTS S34', 'ANTI GUMPAL DAN TIPIS

Popok New Generation SAP Pertama di Indonesia
Struktur yang stabil membuat popok Makuku AntiGumpal sehingga membuat bayi tidak merasa berat saat bergerak.
Bahan SAP yang cepat menyerap sekaligus mengunci cairan dengan merata, dapat melindungi pantat bayi dari iritasi dan AntiRuamPopok', 100, 57500, '/static/images/uploads/3e9a2106-4662-40d2-b20d-4e37cdfa2111.jpg', '2024-02-08 16:41:14.59654+07', '2024-06-08 12:25:41.498359+07', '981464fb-3241-4a33-97ae-33b110e2d4aa', 53000, 1000);
INSERT INTO public.products VALUES ('212aa15a-c0dd-469a-a640-6c39b3a53429', 'S26 PROMIL GOLD PHPRO 1 400GR', 'S-26 Promil GOLD pHPro Tahap 1 dengan Protein Terhidrolisis sebagian adalah susu formula untuk bayi usia 0-6 bulan. Diperkaya dengan nutrisi DHA yang memegang peran penting dalam pembentukan otak, retina mata dan sistem saraf pusat. DHA dalam sel tubuh anak dapat membantu perkembangan kemampuan motorik serta meningkatkan fokus. Kandungan DHA dalam susu pertumbuhan ini dapat membantu perkembangan daya pikir dan daya ingat si kecil yang sangat penting untuk masa depannya.', 100, 213000, '/static/images/uploads/9ba6b466-d606-4aea-8139-50d46a244825.jpg', '2024-02-08 16:57:03.378729+07', '2024-06-08 12:25:13.24545+07', 'ea600c63-283a-415e-8ed1-b10d12c544a0', 209000, 700);
INSERT INTO public.products VALUES ('5979bc0a-3bee-4035-995f-306516943bdb', 'FF PRIMAMIL 6-12 BULAN 1500GR', 'Frisian Baby PRIMAMIL 6-12 Bulan adalah susu formula lanjutan untuk bayi usia 6-12 Bulan. Frisian Baby Primamil 6-12 Bulan dibuat sebagai makanan pendamping ASI. Mengandung banyak nutrisi penting seperti DHA, 9AAE, AA, Prebiotik FOS, dan 28 Vitamin & Mineral untuk mendukung pertumbuhan dan perkembangan si kecil.', 100, 176000, '/static/images/uploads/b999255e-2226-4d38-9c62-0b81449032a5.jpg', '2024-02-08 16:55:33.858281+07', '2024-06-08 12:25:17.604043+07', 'ea600c63-283a-415e-8ed1-b10d12c544a0', 173000, 1000);
INSERT INTO public.products VALUES ('295ac270-c0d4-4432-a7b9-3f4e4dd69130', 'LIFREE POPOK CELANA L16', 'Lifree Popok Perekat L-16 merupakan popok perekat dengan inovasi baru dan teknologi Jepang yakni berdaya serap tinggi dan cepat kering. Popok dewasa yang dibuat dengan sirkulasi udara yang baik agar tidak terasa panas, terutama untuk penggunaan sehari-hari.', 100, 137000, '/static/images/uploads/6dcb184b-1e05-44b9-a5b6-20613e08d3bd.jpg', '2024-02-08 17:14:07.655774+07', '2024-02-14 14:22:13.526947+07', 'f5976ce9-7496-4fd2-8322-3beaef36e4d8', 135000, 700);


--
-- TOC entry 2873 (class 2606 OID 26894)
-- Name: categories categories_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pk PRIMARY KEY (id);


--
-- TOC entry 2883 (class 2606 OID 26996)
-- Name: detail_online_orders detail_online_orders_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_online_orders
    ADD CONSTRAINT detail_online_orders_pk PRIMARY KEY (id);


--
-- TOC entry 2881 (class 2606 OID 26961)
-- Name: detail_pickup_online_orders detail_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_pickup_online_orders
    ADD CONSTRAINT detail_order_pk PRIMARY KEY (id);


--
-- TOC entry 2877 (class 2606 OID 26935)
-- Name: pickup_online_orders offline_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pickup_online_orders
    ADD CONSTRAINT offline_order_pk PRIMARY KEY (id);


--
-- TOC entry 2879 (class 2606 OID 26948)
-- Name: online_orders online_order_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.online_orders
    ADD CONSTRAINT online_order_pk PRIMARY KEY (id);


--
-- TOC entry 2875 (class 2606 OID 26902)
-- Name: products products_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pk PRIMARY KEY (id);


--
-- TOC entry 2887 (class 2606 OID 26997)
-- Name: detail_online_orders detail_online_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_online_orders
    ADD CONSTRAINT detail_online_orders_fk FOREIGN KEY (online_order_id) REFERENCES public.online_orders(id);


--
-- TOC entry 2888 (class 2606 OID 27002)
-- Name: detail_online_orders detail_online_orders_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_online_orders
    ADD CONSTRAINT detail_online_orders_fk_1 FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 2885 (class 2606 OID 26972)
-- Name: detail_pickup_online_orders detail_order_fk_1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_pickup_online_orders
    ADD CONSTRAINT detail_order_fk_1 FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- TOC entry 2886 (class 2606 OID 35371)
-- Name: detail_pickup_online_orders detail_pickup_online_orders_pickup_online_orders_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.detail_pickup_online_orders
    ADD CONSTRAINT detail_pickup_online_orders_pickup_online_orders_fk FOREIGN KEY (pickup_online_order_id) REFERENCES public.pickup_online_orders(id);


--
-- TOC entry 2884 (class 2606 OID 26903)
-- Name: products products_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_fk FOREIGN KEY (category_id) REFERENCES public.categories(id);


-- Completed on 2024-06-08 13:44:47

--
-- PostgreSQL database dump complete
--

