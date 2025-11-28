CREATE TABLE public.users (
	user_id bigserial NOT NULL,
	"name" varchar(150) NOT NULL,
	email varchar(150) NOT NULL,
	"password" text NOT NULL,
	"role" varchar(50) DEFAULT 'customer'::character varying NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE public.favorites (
	user_id int4 NOT NULL,
	vehicle_id int4 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT favorites_pkey PRIMARY KEY (user_id, vehicle_id),
	CONSTRAINT fk_favorites_user FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT fk_favorites_vehicle FOREIGN KEY (vehicle_id) REFERENCES public.vehicles(vehicle_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE public.vehicles (
	vehicle_id bigserial NOT NULL,
	type_id int8 NOT NULL,
	brand_id int8 NOT NULL,
	"name" varchar(150) NOT NULL,
	fuel_type varchar(50) NULL,
	transmission varchar(50) NULL,
	price numeric(12, 2) NOT NULL,
	stock int4 DEFAULT 0 NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT vehicles_pkey PRIMARY KEY (vehicle_id),
	CONSTRAINT fk_vehicle_brand FOREIGN KEY (brand_id) REFERENCES public.brands(brand_id) ON DELETE RESTRICT ON UPDATE CASCADE,
	CONSTRAINT fk_vehicle_type FOREIGN KEY (type_id) REFERENCES public.vehicle_types(type_id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE public.vehicle_types (
	type_id bigserial NOT NULL,
	type_name varchar(100) NOT NULL,
	CONSTRAINT vehicle_types_name_unique UNIQUE (type_name),
	CONSTRAINT vehicle_types_pkey PRIMARY KEY (type_id)
);

CREATE TABLE public.brands (
	brand_id bigserial NOT NULL,
	brand_name varchar(100) NOT NULL,
	CONSTRAINT brands_name_unique UNIQUE (brand_name),
	CONSTRAINT brands_pkey PRIMARY KEY (brand_id)
);

CREATE TABLE public.transactions (
	transaction_id bigserial NOT NULL,
	user_id int8 NOT NULL,
	total_amount numeric(12, 2) NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT transactions_pkey PRIMARY KEY (transaction_id),
	CONSTRAINT fk_transactions_user FOREIGN KEY (user_id) REFERENCES public.users(user_id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE public.transaction_items (
	detail_id bigserial NOT NULL,
	transaction_id int8 NOT NULL,
	vehicle_id int8 NOT NULL,
	quantity int4 NOT NULL,
	price numeric(12, 2) NOT NULL,
	CONSTRAINT transaction_items_pkey PRIMARY KEY (detail_id),
	CONSTRAINT fk_items_transaction FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT fk_items_vehicle FOREIGN KEY (vehicle_id) REFERENCES public.vehicles(vehicle_id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE public.payment_detail (
	payment_id bigserial NOT NULL,
	transaction_id int8 NOT NULL,
	payment_method varchar(50) NOT NULL,
	status varchar(50) NOT NULL,
	paid_at timestamp NULL,
	note text NULL,
	CONSTRAINT payment_detail_pkey PRIMARY KEY (payment_id),
	CONSTRAINT payment_detail_transaction_id_key UNIQUE (transaction_id),
	CONSTRAINT fk_payment_transaction FOREIGN KEY (transaction_id) REFERENCES public.transactions(transaction_id) ON DELETE CASCADE ON UPDATE CASCADE
);