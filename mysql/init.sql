/* Inits ramenGo database container with this migrations */

CREATE TABLE broths (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    imageInactive VARCHAR(510),
    imageActive VARCHAR(510),
    name VARCHAR(255),
    description VARCHAR(510),
    price DECIMAL(16,14)
);

CREATE TABLE proteins (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    imageInactive VARCHAR(510),
    imageActive VARCHAR(510),
    name VARCHAR(255),
    description VARCHAR(510),
    price DECIMAL(16,14)

);


CREATE TABLE orders (
    id VARCHAR(255) NOT NULL PRIMARY KEY, 
    description VARCHAR(510),
    image VARCHAR(510),
    brothId VARCHAR(255),
    CONSTRAINT FK_BrothId FOREIGN KEY (brothId) REFERENCES broths(id),
    proteinId VARCHAR(255),
    CONSTRAINT FK_ProteinId FOREIGN KEY (proteinId) REFERENCES proteins(id)
);