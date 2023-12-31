CREATE TABLE m_vouchers (
    id BIGINT NOT NULL AUTO_INCREMENT,
    description VARCHAR(1000) NOT NULL,
    min_price BIGINT NOT NULL,
    quantity_all INT NOT NULL,
    quantity_user INT NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    user_id BIGINT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (id),
    PRIMARY KEY (id)
);

ALTER TABLE m_vouchers
ADD CONSTRAINT FK1_Users_Vouchers FOREIGN KEY (user_id) REFERENCES m_users(id);