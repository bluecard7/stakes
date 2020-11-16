CREATE TABLE clock_records (
    id serial PRIMARY KEY,
    user_email varchar(128),
    clock_interval tsrange,
    created_at date
);
