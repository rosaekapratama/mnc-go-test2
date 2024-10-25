# mnc-go-test2

Init tables
```
CREATE TABLE public.users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	first_name varchar NULL,
	last_name varchar NULL,
	phone_number varchar NULL,
	address varchar NULL,
	pin varchar NOT NULL,
	created_dt timestamptz NOT NULL,
	updated_dt timestamptz NULL,
	CONSTRAINT users_phone_no_un UNIQUE (phone_number),
	CONSTRAINT users_pk PRIMARY KEY (id)
);
```

Adjust database configuration in ```conf/app.yaml```
```
database:
  playground:
    driver: postgresql
    address: localhost
    database: playground
    username: postgres
    password: 2628
```


Register
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