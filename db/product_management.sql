CREATE DATABASE IF NOT EXISTS product_management;
USE product_management;

CREATE TABLE IF NOT EXISTS products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    product_code VARCHAR(20) UNIQUE NOT NULL,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    quantity INT DEFAULT 0,
    category VARCHAR(50),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Sample data
INSERT INTO products (product_code, name, price, quantity, category, description) VALUES
('P001', 'Laptop Dell XPS 13', 1299.99, 10, 'Electronics', 'High-performance laptop'),
('P002', 'iPhone 15 Pro', 999.99, 25, 'Electronics', 'Latest iPhone model'),
('P003', 'Office Chair', 199.99, 50, 'Furniture', 'Ergonomic office chair'),
('P004', 'Samsung Galaxy S24 Ultra', 1199.99, 20, 'Electronics', 'Flagship Android smartphone with AI features'),
('P005', 'Apple iPad Air 2024', 699.99, 18, 'Electronics', 'Lightweight tablet with M2 chip'),
('P006', 'Sony WH-1000XM5 Headphones', 349.99, 30, 'Electronics', 'Premium noise-cancelling wireless headphones'),
('P007', 'Logitech MX Master 3S Mouse', 99.99, 40, 'Electronics', 'Advanced ergonomic wireless mouse'),
('P008', '4K Webcam Pro', 149.99, 15, 'Electronics', 'Ultra HD video conferencing camera'),
('P009', 'Bookshelf Oak 5-Tier', 159.99, 22, 'Furniture', 'Modern oak wood bookshelf'),
('P010', 'Recliner Sofa Chair', 459.99, 10, 'Furniture', 'Comfortable reclining leather chair'),
('P011', 'Standing Desk Adjustable', 329.99, 12, 'Furniture', 'Electric height adjustable standing desk'),
('P012', 'Men Winter Jacket', 89.99, 35, 'Clothing', 'Warm and waterproof winter jacket'),
('P013', 'Women Sports Leggings', 29.99, 60, 'Clothing', 'Stretchable quick-dry leggings'),
('P014', 'Unisex Hoodie Classic', 39.99, 55, 'Clothing', 'Cotton fleece hoodie'),
('P015', 'Machine Learning with Python', 54.99, 25, 'Books', 'Hands-on ML guide with examples'),
('P016', 'Clean Code', 49.99, 28, 'Books', 'Software craftsmanship principles & practices'),
('P017', 'Organic Honey 500g', 14.99, 80, 'Food', 'Pure natural honey from organic farms'),
('P018', 'Premium Green Tea 200g', 12.99, 90, 'Food', 'High-quality loose-leaf green tea')
ON DUPLICATE KEY UPDATE
name = VALUES(name),
price = VALUES(price),
quantity = VALUES(quantity),
category = VALUES(category),
description = VALUES(description);
