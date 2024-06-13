CREATE TABLE execution_time (
    id SERIAL PRIMARY KEY,
    parameter VARCHAR,
    test VARCHAR,
    value NUMERIC,
    deviation NUMERIC
);
