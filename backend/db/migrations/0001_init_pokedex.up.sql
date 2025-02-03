CREATE TABLE pokedex (
    id VARCHAR(255) PRIMARY KEY,
    formId VARCHAR(255) NOT NULL,
    dex_num INT NOT NULL,
    generation INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    primary_type VARCHAR(255) NOT NULL,
    secondary_type VARCHAR(255)
);