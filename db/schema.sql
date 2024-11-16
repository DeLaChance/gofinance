
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    date DATE, 
    description VARCHAR(64) NOT NULL, 
    account VARCHAR (64) NOT NULL, 
    amount NUMERIC(13,4) NOT NULL, 
    remarks VARCHAR (128) NULL
);
