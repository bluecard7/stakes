CREATE TABLE clock_records (
    id serial PRIMARY KEY,
    email varchar(128) UNIQUE NOT NULL,
    clockIn timestamp,
    clockOut timestamp
);
