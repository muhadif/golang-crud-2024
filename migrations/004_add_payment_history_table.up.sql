
CREATE TABLE IF NOT EXISTS `payment_history`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    serial VARCHAR(256) NOT NULL,
    user_serial VARCHAR(255) NOT NULL,
    open_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expired_time TIMESTAMP NULL DEFAULT NULL,
    status VARCHAR(50) NOT NULL,
    payment_method VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    CONSTRAINT UC_serial UNIQUE (serial)
    );

CREATE TABLE IF NOT EXISTS `payment_history_item`
(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    payment_history_serial VARCHAR(256) NOT NULL,
    product_serial VARCHAR(256) NOT NULL,
    price DECIMAL(10, 2) DEFAULT 0.0,
    quantity INT NOT NULL
)