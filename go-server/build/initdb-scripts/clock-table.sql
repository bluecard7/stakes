CREATE TABLE clock_records (
    id uuid PRIMARY KEY,
    email varchar(128) NOT NULL,
    clockIn timestamp,
    clockOut timestamp
);
