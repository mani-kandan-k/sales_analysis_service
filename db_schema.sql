
-- 1) CUSTOMER TABLE
CREATE TABLE CustomerDetails (
    id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id VARCHAR(20) ,
    customer_name VARCHAR(100),
    customer_email VARCHAR(100),
    customer_address VARCHAR(50),
    createdDate DATETIME,
    updatedDate DATETIME
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- 2) PRODUCT TABLE
CREATE TABLE ProductDetails (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id VARCHAR(20) ,
    product_name VARCHAR(100),
    unit_price DECIMAL(10,2),
    discount DECIMAL(10,2),
    createdDate DATETIME,
    updatedDate DATETIME
)ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

-- 3) ORDER TABLE
CREATE TABLE OrderDetails (
    id int AUTO_INCREMENT PRIMARY KEY,
    order_id VARCHAR(100),
    product_id int,
    customer_id int,
    region VARCHAR(50),
    category VARCHAR(50),
    date_of_sale DATE,
    quantity_sold INT,
    shipping_cost DECIMAL(10,2),
    payment_method VARCHAR(30), 
    createdDate DATETIME,
    updatedDate DATETIME,
    CONSTRAINT fk_OrderDetails_customer_id FOREIGN KEY (customer_id) REFERENCES CustomerDetails(id) ON DELETE CASCADE,
    CONSTRAINT fk_OrderDetails_product_id FOREIGN KEY (product_id) REFERENCES ProductDetails(id) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;