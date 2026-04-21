package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PaginationMeta struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

type PaginatedProductsResponse struct {
	Products   []Product      `json:"products"`
	Pagination PaginationMeta `json:"pagination"`
}

// POST /products
func createProduct(c *gin.Context) {
	var p Product

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.Price > 100000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price too large"})
		return
	}

	var exists int
	err := DB.QueryRow("SELECT COUNT(*) FROM products WHERE product_code = ?", p.ProductCode).Scan(&exists)

	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code already exists"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	query := `
	INSERT INTO products 
	(product_code, name, price, quantity, category, description, created_at) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := DB.Exec(query,
		p.ProductCode,
		p.Name,
		p.Price,
		p.Quantity,
		p.Category,
		p.Description,
		time.Now(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	p.ID = int64(id)
	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": p})
}

// GET /products
func getProducts(c *gin.Context) {
	rows, err := DB.Query(`
		SELECT id, product_code, name, price, quantity, category, description, created_at 
		FROM products`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		rows.Scan(
			&p.ID,
			&p.ProductCode,
			&p.Name,
			&p.Price,
			&p.Quantity,
			&p.Category,
			&p.Description,
			&p.CreatedAt,
		)
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

// GET /products/:id
func getProductById(c *gin.Context) {
	id := c.Param("id")

	var p Product
	err := DB.QueryRow(`SELECT id, product_code, name, price, quantity, category, description, created_at FROM products WHERE id = ?`, id).
		Scan(&p.ID, &p.ProductCode, &p.Name, &p.Price,
			&p.Quantity, &p.Category, &p.Description, &p.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, p)
}

// PUT /products/:id
func updateProductById(c *gin.Context) {
	id := c.Param("id")

	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if p.ProductCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code cannot be empty"})
		return
	}

	if p.Price > 100000000 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price too large"})
		return
	}

	query := `
	UPDATE products 
	SET product_code=?, name=?, price=?, quantity=?, category=?, description=?
	WHERE id=?`

	res, err := DB.Exec(query,
		p.ProductCode,
		p.Name,
		p.Price,
		p.Quantity,
		p.Category,
		p.Description,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

// DELETE /products/:id
func deleteProductById(c *gin.Context) {
	id := c.Param("id")

	res, err := DB.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}

// GET /products/category/
func getAllCategories(c *gin.Context) {
	rows, err := DB.Query(`SELECT DISTINCT category FROM products`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var cat string
		if err := rows.Scan(&cat); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

// GET /products/category/:category
func getProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	rows, err := DB.Query(`
	SELECT id, product_code, name, price, quantity, category, description, created_at FROM products WHERE category = ?`, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Price,
			&p.Quantity, &p.Category, &p.Description, &p.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

func getProductsPagination(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}

	offset := (page - 1) * limit

	rows, err := DB.Query(`
		SELECT id, product_code, name, price, quantity, category, description, created_at FROM products LIMIT ? OFFSET ?`, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Name, &p.Price,
			&p.Quantity, &p.Category, &p.Description, &p.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	var total int
	err = DB.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	// Calculate pagination metadata
	totalPages := (total + limit - 1) / limit

	hasNext := page < totalPages
	hasPrev := page > 1

	pagination := PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}

	c.JSON(http.StatusOK, PaginatedProductsResponse{
		Products:   products,
		Pagination: pagination,
	})
}

// GET /products/filter
func filterProducts(c *gin.Context) {
	// Query params
	keyword := c.Query("q")
	category := c.Query("category")
	minPrice := c.Query("minPrice")
	maxPrice := c.Query("maxPrice")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if page < 1 {
		page = 1
	}
	if limit <= 0 {
		limit = 5
	}
	offset := (page - 1) * limit

	// Base query
	query := `SELECT id, product_code, name, price, quantity, category, description, created_at 
	          FROM products WHERE 1=1`
	args := []interface{}{}

	// Dynamic conditions
	if keyword != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+keyword+"%")
	}

	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	if minPrice != "" && maxPrice != "" {
		query += " AND price BETWEEN ? AND ?"
		args = append(args, minPrice, maxPrice)
	} else if minPrice != "" {
		query += " AND price >= ?"
		args = append(args, minPrice)
	} else if maxPrice != "" {
		query += " AND price <= ?"
		args = append(args, maxPrice)
	}

	// Sorting by price
	sort := c.DefaultQuery("sort", "asc") // default asc

	if sort != "asc" && sort != "desc" {
		sort = "asc" // fallback to safe value
	}

	query += " ORDER BY price " + sort

	// Pagination
	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID,
			&p.ProductCode,
			&p.Name,
			&p.Price,
			&p.Quantity,
			&p.Category,
			&p.Description,
			&p.CreatedAt,
		); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}

	c.JSON(http.StatusOK, products)
}

func main() {
	InitDB()

	r := gin.Default()

	r.POST("/products", createProduct)
	r.GET("/products", getProductsPagination)
	r.GET("/products/:id", getProductById)
	r.PUT("/products/:id", updateProductById)
	r.DELETE("/products/:id", deleteProductById)

	r.GET("/categories", getAllCategories)
	r.GET("/category/:category", getProductsByCategory)
	r.GET("/products/filter", filterProducts)

	r.Run(":8080")
}
