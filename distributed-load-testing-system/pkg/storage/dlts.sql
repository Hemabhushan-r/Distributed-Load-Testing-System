--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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
-- Name: scenario_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scenario_requests (
    id uuid NOT NULL,
    scenario_id uuid,
    method character varying(10) NOT NULL,
    url text NOT NULL,
    headers jsonb,
    body text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.scenario_requests OWNER TO postgres;

--
-- Name: test_configurations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_configurations (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    description text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.test_configurations OWNER TO postgres;

--
-- Name: test_scenarios; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.test_scenarios (
    id uuid NOT NULL,
    config_id uuid,
    name character varying(255) NOT NULL,
    request_rate integer NOT NULL,
    concurrency_level integer NOT NULL,
    duration integer NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.test_scenarios OWNER TO postgres;

--
-- Data for Name: scenario_requests; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scenario_requests (id, scenario_id, method, url, headers, body, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: test_configurations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test_configurations (id, name, description, created_at, updated_at) FROM stdin;
5c8aaffa-b0b6-471d-af1d-a5d286381978	Go Server Test	Testing Dummy Go HTTP server	2024-07-08 19:56:42.472499	2024-07-08 19:56:42.472499
\.


--
-- Data for Name: test_scenarios; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.test_scenarios (id, config_id, name, request_rate, concurrency_level, duration, created_at, updated_at) FROM stdin;
\.


--
-- Name: scenario_requests scenario_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scenario_requests
    ADD CONSTRAINT scenario_requests_pkey PRIMARY KEY (id);


--
-- Name: test_configurations test_configurations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_configurations
    ADD CONSTRAINT test_configurations_pkey PRIMARY KEY (id);


--
-- Name: test_scenarios test_scenarios_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_scenarios
    ADD CONSTRAINT test_scenarios_pkey PRIMARY KEY (id);


--
-- Name: scenario_requests scenario_requests_scenario_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scenario_requests
    ADD CONSTRAINT scenario_requests_scenario_id_fkey FOREIGN KEY (scenario_id) REFERENCES public.test_scenarios(id) ON DELETE CASCADE;


--
-- Name: test_scenarios test_scenarios_config_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.test_scenarios
    ADD CONSTRAINT test_scenarios_config_id_fkey FOREIGN KEY (config_id) REFERENCES public.test_configurations(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

