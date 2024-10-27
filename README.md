# mnc-go-test2

## Note
```
Having a queue task that will transfer money to another user (bonus point).
```
Above request is contradictive with the requested API spec for transfer especially the response, cause when we are using queue, then the process will become async.

We can block the main thread process to wait async response via chan, but i don't think it's a good design.

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
	status varchar NOT NULL,
	cr bool NOT NULL,
	amount numeric(20, 3) NOT NULL,
	balance_before numeric(20, 3) NOT NULL,
	balance_after numeric(20, 3) NOT NULL,
	remark varchar NULL,
	created_dt timestamptz NOT NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (id),
	CONSTRAINT transactions_accounts_fk FOREIGN KEY (account_id) REFERENCES public.accounts(id)
);

INSERT INTO public.users (id,first_name,last_name,phone_number,address,pin,created_dt,updated_dt) VALUES
	 ('0f707cca-0abd-4f11-8bf4-449fc3daef90'::uuid,'Rosa','Pratama','085722955088','THE RIVER Parung Panjang','7c4a8d09ca3762af61e59520943dc26494f8941b','2024-10-27 03:18:53.856',NULL),
	 ('c3c8254d-e38a-4ddd-8f33-1e928f7083cf'::uuid,'Guntur','Saputro','0811255501','Jl. Kebon Sirih No. 1','7c4a8d09ca3762af61e59520943dc26494f8941b','2024-10-27 03:15:49.331','2024-10-27 03:46:32.252'),
	 ('a8b0331d-e53d-4972-926a-849471ae20d8'::uuid,'Ahmad','Wijaya','087437052790','Perumahan Green Hills','ea60957c751e0247371dac948ac738129efa4c27','2024-10-27 03:51:39.547',NULL),
	 ('6cdde176-73ac-4ded-ad78-8bf64260d2af'::uuid,'Teguh','Rahmawati','084568175760','Jl. Thamrin','973a6a06644e168b55a3442640fdf39f625a3a24','2024-10-27 03:51:46.259',NULL),
	 ('49200fbe-a47a-43a0-a50d-961465a26130'::uuid,'Arif','Fadilah','086371102940','Jl. Malioboro','2c9391b9e6de35afcdde4b059abef6c71c9a5245','2024-10-27 03:51:52.181',NULL),
	 ('5d1b85d0-1365-4696-ba9a-ccf2105cb70e'::uuid,'Dewi','Purwanto','088474578477','Jl. Gatot Subroto','6eb9ecc44088ca95c1d3b31f564e2f8d47394543','2024-10-27 03:51:58.326',NULL),
	 ('d0787641-e76e-4539-9d8d-8e55aa151806'::uuid,'Wulan','Fadilah','085691838197','Jl. Malioboro','5ee5f62d6f21c1359432e652c85219830fff8139','2024-10-27 03:52:04.725',NULL),
	 ('b6f0df29-5ad1-4b2b-bb29-534fb9213378'::uuid,'Siti','Fadilah','087619212952','Jl. Sudirman','db8806690b08514cf220521cb3a70211182bbdb9','2024-10-27 03:52:10.550',NULL),
	 ('ec1520bd-3372-43b5-acb1-56f24c116751'::uuid,'Indra','Santoso','088118457175','Jl. Asia Afrika','f152b8b21d7e87b14664ae3bf01612cdb14820b8','2024-10-27 03:52:16.460',NULL),
	 ('672a5bbd-866f-4800-bd5f-21009ce70a80'::uuid,'Ahmad','Nugraha','087714524262','Jl. Kebon Sirih No. 1','caaae94785fd7b78e4ac854293ef76d0e1b99acb','2024-10-27 03:52:23.455',NULL);
INSERT INTO public.users (id,first_name,last_name,phone_number,address,pin,created_dt,updated_dt) VALUES
	 ('0ddb82b6-3745-42c0-bb75-e757203eeec6'::uuid,'Wulan','Pratama','085346356732','Jl. Malioboro','d2d017cb9f820c5a904d0e44115efa3d85d3c013','2024-10-27 03:52:29.727',NULL),
	 ('54e4a576-5a03-486a-987b-202ebb39e676'::uuid,'Siti','Yulianti','089168254177','THE RIVER Parung Panjang','727e2fc628dea3c706b826d1df7eaa763e0029c0','2024-10-27 03:52:35.641',NULL);
INSERT INTO public.accounts (id,user_id,"type",balance,created_dt,update_dt) VALUES
	 ('e26959e1-92fe-48a9-90f2-7b36488fbb46'::uuid,'c3c8254d-e38a-4ddd-8f33-1e928f7083cf'::uuid,'SAVING',260001.000,'2024-10-27 03:15:49.334','2024-10-27 03:23:00.045'),
	 ('c02c19ca-2eb8-4ff1-a2f2-a5a7bac42238'::uuid,'0f707cca-0abd-4f11-8bf4-449fc3daef90'::uuid,'SAVING',40000.000,'2024-10-27 03:18:53.857','2024-10-27 03:23:00.045'),
	 ('8047c22f-b991-48fe-acb9-22dea5a35bb9'::uuid,'a8b0331d-e53d-4972-926a-849471ae20d8'::uuid,'SAVING',0.000,'2024-10-27 03:51:39.548',NULL),
	 ('633d8bff-ddd8-4f9a-b0ca-5d47d1722edb'::uuid,'6cdde176-73ac-4ded-ad78-8bf64260d2af'::uuid,'SAVING',0.000,'2024-10-27 03:51:46.259',NULL),
	 ('e1015e5f-beb2-4017-b9c1-37f36a98091a'::uuid,'49200fbe-a47a-43a0-a50d-961465a26130'::uuid,'SAVING',0.000,'2024-10-27 03:51:52.182',NULL),
	 ('3eb95dfb-5a6d-4eed-8be2-631e83629d63'::uuid,'5d1b85d0-1365-4696-ba9a-ccf2105cb70e'::uuid,'SAVING',0.000,'2024-10-27 03:51:58.326',NULL),
	 ('e331db91-179a-4365-91a4-f2749b4463af'::uuid,'d0787641-e76e-4539-9d8d-8e55aa151806'::uuid,'SAVING',0.000,'2024-10-27 03:52:04.726',NULL),
	 ('dbf3c2a6-52e5-4adc-9bab-09e37f7b0982'::uuid,'b6f0df29-5ad1-4b2b-bb29-534fb9213378'::uuid,'SAVING',0.000,'2024-10-27 03:52:10.550',NULL),
	 ('7d0ea0d9-50b2-4f5c-a999-42874bdb9d5d'::uuid,'ec1520bd-3372-43b5-acb1-56f24c116751'::uuid,'SAVING',0.000,'2024-10-27 03:52:16.460',NULL),
	 ('c67b892c-0829-4a21-8325-f49a28415f56'::uuid,'672a5bbd-866f-4800-bd5f-21009ce70a80'::uuid,'SAVING',0.000,'2024-10-27 03:52:23.456',NULL);
INSERT INTO public.accounts (id,user_id,"type",balance,created_dt,update_dt) VALUES
	 ('fb7558a2-c923-4837-913c-264d37e862f8'::uuid,'0ddb82b6-3745-42c0-bb75-e757203eeec6'::uuid,'SAVING',0.000,'2024-10-27 03:52:29.727',NULL),
	 ('c3916b95-b8aa-4126-bab5-4b4a7c902ad8'::uuid,'54e4a576-5a03-486a-987b-202ebb39e676'::uuid,'SAVING',0.000,'2024-10-27 03:52:35.641',NULL);
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
### 3. Run application
```
.\mnc-go-test2.exe (windows)
./mnc-go-test2 (linux)
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
Login first to get access token, then use it as Authorization Bearer token
```
curl --location 'http://localhost:9092/topup' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1NjMxYjRjZC1lOWI1LTRiYjAtYTliMS03NGMwM2NkZDY0YTYiLCJwaG9uZV9udW1iZXIiOiIwODExMjU1NTAxIiwiZXhwIjoxNzI5ODk4ODI3fQ.qgWeTri2d1iQAMxzO2z4_7cpYD3BRTebKhp_SWhNTlY' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 500000
}'
```

### 4. Payment
Login first to get access token, then use it as Authorization Bearer token
```
curl --location 'http://localhost:9092/pay' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1NjMxYjRjZC1lOWI1LTRiYjAtYTliMS03NGMwM2NkZDY0YTYiLCJwaG9uZV9udW1iZXIiOiIwODExMjU1NTAxIiwiZXhwIjoxNzI5OTcwNDc0fQ.b8b16YuXq7Sw0gmjURPKwtYc-uSUVuk_MsIJUHytmm8' \
--data '{
    "amount": 100000,
    "remarks": "Pulsa Telkomsel 100k"
}'
```

### 5. Transfer
Login first to get access token, then use it as Authorization Bearer token

Create new user first for destination user if not have
```
curl --location 'http://localhost:9092/register' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "Rosa",
    "last_name": "Pratama",
    "phone_number": "085722955088",
    "address": "THE RIVER Parung Panjang",
    "pin": "123456"
}'
```
Transfer funds to destination user
```
curl --location 'http://localhost:9092/transfer' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI1NjMxYjRjZC1lOWI1LTRiYjAtYTliMS03NGMwM2NkZDY0YTYiLCJwaG9uZV9udW1iZXIiOiIwODExMjU1NTAxIiwiZXhwIjoxNzI5OTczNTcxfQ.YWxehAPRVK_FQK02wKc4uOz_lyliQgD-HJyYQibKI5E' \
--data '{
    "target_user": "cdb7f504-ef54-4236-abeb-93c25f7dec99",
    "amount": 20000,
    "remarks": "Hadiah Ultah"
}'
```

### 6. Transactions
Login first to get access token, then use it as Authorization Bearer token

```
curl --location 'http://localhost:9092/transactions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJjM2M4MjU0ZC1lMzhhLTRkZGQtOGYzMy0xZTkyOGY3MDgzY2YiLCJwaG9uZV9udW1iZXIiOiIwODExMjU1NTAxIiwiZXhwIjoxNzI5OTc2NDkyfQ.XPEITRo9xw13spomHtOEpTXMoOV4DzSHabNHHwXqZbA'
```