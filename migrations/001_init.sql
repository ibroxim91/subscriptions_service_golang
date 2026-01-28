CREATE TABLE public.subscriptions (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    service_name character varying(255) NOT NULL,
    price bigint NOT NULL,
    user_id uuid NOT NULL,
    start_date timestamp with time zone NOT NULL,
    end_date timestamp with time zone,
    CONSTRAINT chk_subscriptions_price CHECK ((price >= 0))
);



CREATE SEQUENCE public.subscriptions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;




ALTER SEQUENCE public.subscriptions_id_seq OWNED BY public.subscriptions.id;



ALTER TABLE ONLY public.subscriptions ALTER COLUMN id SET DEFAULT nextval('public.subscriptions_id_seq'::regclass);

ALTER TABLE ONLY public.subscriptions
    ADD CONSTRAINT subscriptions_pkey PRIMARY KEY (id);



CREATE INDEX idx_subscriptions_deleted_at ON public.subscriptions USING btree (deleted_at);



