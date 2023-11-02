CREATE TABLE t_order_details (
    id BIGINT NOT NULL AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    price BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (id),
    PRIMARY KEY (id)
);

ALTER TABLE t_order_details
ADD CONSTRAINT FK_Orders_OrderDetails FOREIGN KEY (order_id) REFERENCES t_orders(id),
ADD CONSTRAINT FK_Products_OrderDetails FOREIGN KEY (product_id) REFERENCES m_products(id);