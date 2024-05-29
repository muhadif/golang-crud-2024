CREATE TABLE IF NOT EXISTS `user`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    serial VARCHAR (255) NOT NULL,
    username VARCHAR (255) NOT NULL,
    full_name VARCHAR (255) NOT NULL,
    role VARCHAR (40) NOT NULL,
    email VARCHAR (255) NOT NUlL,
    password VARCHAR (255) NOT NULL,
    status VARCHAR (40) NOT NULL,
    access_status VARCHAR (40) NOT NULL,
    registration_otp VARCHAR (16) NOT NULL,
    forgot_password_token VARCHAR (16) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT UC_SERIAL UNIQUE (serial),
    CONSTRAINT UC_EMAIL UNIQUE (email)
    );

CREATE TABLE IF NOT EXISTS `user_token`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_serial VARCHAR (255) NOT NULL,
    refresh_token VARCHAR (255) NOT NULL,
    CONSTRAINT UC_user_serial UNIQUE (user_serial)
);

CREATE TABLE IF NOT EXISTS `product_category`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    serial VARCHAR (255) NOT NULL,
    name VARCHAR (255) NOT NULL,
    CONSTRAINT UC_user_serial UNIQUE (serial)
);

CREATE TABLE IF NOT EXISTS `product`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    serial VARCHAR (255) NOT NULL,
    name VARCHAR (255) NOT NULL,
    price DECIMAL(10, 2) DEFAULT 0.0,
    stock INT DEFAULT 0,
    description TEXT,
    CONSTRAINT UC_serial UNIQUE (serial)
);

CREATE TABLE IF NOT EXISTS `product_product_category`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    product_serial VARCHAR (255) NOT NULL,
    product_category_serial VARCHAR (255) NOT NULL,
    CONSTRAINT UC_product_product_category UNIQUE (product_serial, product_category_serial)
);


-- Seeding
INSERT INTO product_category (serial, name) VALUES
                                                ('001', 'Electronics'),
                                                ('002', 'Clothing'),
                                                ('003', 'Books'),
                                                ('004', 'Home and Kitchen');

-- Insert seed data into the product table
INSERT INTO product (serial, name, price, stock, description) VALUES
                                                                  ('P001', 'Laptop', 999.99, 50, 'High-performance laptop with SSD storage.'),
                                                                  ('P002', 'T-shirt', 19.99, 100, 'Comfortable cotton t-shirt in various colors.'),
                                                                  ('P003', 'The Great Gatsby', 12.99, 75, 'Classic novel by F. Scott Fitzgerald.'),
                                                                  ('P004', 'Coffee Maker', 49.99, 30, 'Automatic drip coffee maker with programmable timer.');

INSERT INTO product_product_category (product_serial, product_category_serial) VALUES
                                                                                   ('P001', '001'), -- Laptop belongs to Electronics
                                                                                   ('P002', '002'), -- T-shirt belongs to Clothing
                                                                                   ('P003', '003'), -- The Great Gatsby belongs to Books
                                                                                   ('P004', '004'); -- Coffee Maker belongs to Home and Kitchen