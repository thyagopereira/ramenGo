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

-- Inserts into table

INSERT INTO broths (id, imageInactive, imageActive, name, description, price) VALUES ("619104a1-8c7d-42b6-9875-f312fc709980", "somepic", "somepic", "Chicken Broth", "Chicken broth for sick ppl.", 10.2);
INSERT INTO broths (id, imageInactive, imageActive, name, description, price) VALUES ("87495b95-1c7f-4038-ae55-ab36ed6a9411", "somepic2", "somepic2", "Fish Broth", "Fish Broth.", 10.25);
INSERT INTO broths (id, imageInactive, imageActive, name, description, price) VALUES ("02c08bd5-7ec2-45d9-8f27-8c3422927b6a", "somepic3", "somepic3", "Capirira Chicken broth", "tasty", 11.32);

INSERT INTO proteins (id, imageInactive, imageActive, name, description, price) VALUES ("3b0e603c-4580-442e-addf-497ef3d2f895", "somepic4", "somepic4", "Cow protein", "The best", 25.3);
INSERT INTO proteins (id, imageInactive, imageActive, name, description, price) VALUES ("1e1e53bc-5c29-4dad-9274-eb208a7d81c1", "somepic5", "somepic5", "Vegan protein", "No animal stuff related.", 40.2);

