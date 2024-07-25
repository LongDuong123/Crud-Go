CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;

-- Tạo bảng User
CREATE TABLE IF NOT EXISTS users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    age INT,
    gender VARCHAR(10),
    email VARCHAR(255),
    role VARCHAR(10) DEFAULT 'user',
    password VARCHAR(255)
);

-- Tạo bảng Product
CREATE TABLE IF NOT EXISTS product (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    image_url VARCHAR(255),
    price DECIMAL(10,2),
    is_created_by INT
);

-- Chèn dữ liệu người dùng John vào bảng User
INSERT INTO users (id, name, age, gender, email, role, password)
VALUES (1, 'John', 18, 'male', 'john@example.com', 'user', '$2a$10$sWfmvprYwVt.u5sQlm02aeQTO0HakmNlrGPzwFs05E3BiQeKlHRmq');

-- Chèn dữ liệu người dùng Duong vào bảng User
INSERT INTO users (id, name, email, role, password)
VALUES (4, 'Duong', 'Duong@example.com', 'user', '$2a$10$Mch6rD2qYIodLXrm3x5t.ubfTg3ahdKj7K7mVZ.AyYsVUyfwvEhsW');

-- Chèn dữ liệu sản phẩm vào bảng Product
INSERT INTO product (id, name, image_url, price, is_created_by)
VALUES
(2, 'Chuoi', 'chuoi.jpg', 12000.00, 1),
(3, 'Sau Rieng', 'cam.jpg', 50000.00, 1),
(4, 'Xoai', 'cam.jpg', 12000.00, 1),
(5, 'Man', 'Man.jpg', 6000.00, 1),
(6, 'Oi', 'Oi.jpg', 20000.00, 1);
