-- Create the "survey" database
CREATE DATABASE survey;

-- Switch to the "survey" database
\c survey;

-- Create the "forms" table
CREATE TABLE forms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    gender CHAR(2) NOT NULL,
    age INT NOT NULL
);

-- Insert data into the "forms" table
INSERT INTO forms (name, gender, age) VALUES
    ('Dobby', 'M', 30),
    ('Misty', 'F', 25),
    ('Taz', 'M', 40);