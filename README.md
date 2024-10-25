# mnc-go-test2

## Pre-Requisite
### 1. Initiate tables
```
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	phone_number varchar NOT NULL,
	address varchar NULL,
	pin varchar NOT NULL,
	created_dt timestamptz NOT NULL,
	updated_dt timestamptz NULL,
	CONSTRAINT users_phone_no_un UNIQUE (phone_number),
	CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE TABLE public.accounts (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	user_id uuid NOT NULL,
	"type" varchar NOT NULL,
	balance numeric(20, 3) DEFAULT 0 NOT NULL,
	created_dt timestamptz NOT NULL,
	update_dt timestamptz NULL,
	CONSTRAINT accounts_pk PRIMARY KEY (id),
	CONSTRAINT accounts_user_id_type_un UNIQUE (user_id, type),
	CONSTRAINT accounts_users_fk FOREIGN KEY (user_id) REFERENCES public.users(id)
);

CREATE TABLE public.transactions (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	account_id uuid NOT NULL,
	"type" varchar NOT NULL,
	cr bool NOT NULL,
	amount numeric(20, 3) NOT NULL,
	balance numeric(20, 3) NOT NULL,
	created_dt timestamptz NOT NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (id),
	CONSTRAINT transactions_accounts_fk FOREIGN KEY (account_id) REFERENCES public.accounts(id)
);
```
### 2. Adjust database configuration in 
```conf/app.yaml```

```
database:
  playground:
    driver: postgresql
    address: localhost
    database: playground
    username: postgres
    password: 2628
```

---
## API Specification
### 1. Register
```
curl --location 'http://localhost:9092/register' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "Guntur",
    "last_name": "Saputro",
    "phone_number": "0811255501",
    "address": "Jl. Kebon Sirih No. 1",
    "pin": "123456"
}'
```

### 2. Login
```
curl --location 'http://localhost:9092/login' \
--header 'Content-Type: application/json' \
--data '{
    "phone_number": "0811255501",
    "pin": "123456"
}'
```

### 3. Topup
```
curl --location 'http://localhost:9092/topup' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1NjMxYjRjZC1lOWI1LTRiYjAtYTliMS03NGMwM2NkZDY0YTYiLCJwaG9uZV9udW1iZXIiOiIwODExMjU1NTAxIiwiZXhwIjoxNzI5ODk4ODI3fQ.qgWeTri2d1iQAMxzO2z4_7cpYD3BRTebKhp_SWhNTlY' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 500000
}'
```