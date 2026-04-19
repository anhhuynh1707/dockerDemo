# 🛒 Product Management API (Golang + Gin + MySQL + Docker)

## 📌 Overview

This project is a **RESTful Product Management API** built using:

* **Golang (Gin framework)**
* **MySQL**
* **Docker & Docker Compose**

It supports full **CRUD operations**, along with advanced features like:

* Pagination
* Filtering (keyword, category, price range)
* Sorting (price ASC/DESC)
* Category management

## 🚀 Features

### ✅ Core Features

* Create product
* Get all products (with pagination)
* Get product by ID
* Update product
* Delete product

### 🔍 Advanced Features

* Search by keyword
* Filter by category
* Filter by price range (minPrice, maxPrice)
* Sorting by price (ASC / DESC)
* Get all categories
* Get products by category

## 🏗️ Tech Stack

* **Backend:** Golang (Gin)
* **Database:** MySQL
* **Containerization:** Docker, Docker Compose
* **API Testing:** Postman

## 📂 Project Structure

```
product_management
├── main.go
├── database.go
├── model.go
├── db/
│   └── product_management.sql
├── docker-compose.yml
├── Dockerfile
├── .env
├── go.mod
├── go.sum
└── README.md
```

## ⚙️ Setup Instructions

### 1️⃣ Clone the repository

```
git clone https://github.com/anhhuynh1707/dockerDemo.git
cd product-management
```


### 2️⃣ Create `.env` file

```
MYSQL_ROOT_PASSWORD=123456
MYSQL_DATABASE=product_management
```

### 3️⃣ Run with Docker

```
docker-compose build
docker-compose up
```

### 4️⃣ API Base URL

```
http://localhost:8080
```

## 📬 API Endpoints

### 🔹 Create Product

```
POST /products
```

### 🔹 Get All Products (Pagination)

```
GET /products
GET /products?page=1&limit=5
```

### 🔹 Get Product by ID

```
GET /products/:id
```


### 🔹 Update Product

```
PUT /products/:id
```


### 🔹 Delete Product

```
DELETE /products/:id
```

### 🔹 Get All Categories

```
GET /categories
```

### 🔹 Get Products by Category

```
GET /category/:category
```

### 🔹 Advanced Filter (Search + Pagination + Sorting)

```
GET /products/filter
```

#### Query Parameters:

| Param    | Description        |
| -------- | ------------------ |
| q        | Search keyword     |
| category | Filter by category |
| minPrice | Minimum price      |
| maxPrice | Maximum price      |
| page     | Page number        |
| limit    | Items per page     |
| sort     | asc / desc         |
