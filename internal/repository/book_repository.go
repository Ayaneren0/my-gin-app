package repository

import (
	"bookstore/internal/models"
	"database/sql"
	"fmt"
	"strings"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `
        INSERT INTO books (title, author, price, stock, created_at, updated_at)
        VALUES (?, ?, ?, ?, NOW(), NOW())
    `
	result, err := r.db.Exec(query, book.Title, book.Author, book.Price, book.Stock)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.ID = id
	return nil
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	query := "SELECT id, title, author, price, stock, created_at, updated_at FROM books"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.Stock,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (r *BookRepository) GetByID(id int64) (*models.Book, error) {
	query := `
        SELECT id, title, author, price, stock, created_at, updated_at
        FROM books WHERE id = ?
    `
	var book models.Book
	err := r.db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Price,
		&book.Stock,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `
        UPDATE books
        SET title = ?, author = ?, price = ?, stock = ?, updated_at = NOW()
        WHERE id = ?
    `
	_, err := r.db.Exec(query, book.Title, book.Author, book.Price, book.Stock, book.ID)
	return err
}

func (r *BookRepository) Delete(id int64) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *BookRepository) GetAllWithFilters(query *models.PaginationQuery) (*models.PaginatedResponse, error) {
	// Base query
	baseQuery := `SELECT id, title, author, price, stock, created_at, updated_at FROM books WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM books WHERE 1=1`

	// Initialize parameters for prepared statement
	params := []interface{}{}

	// Build where clause
	whereClause := ""

	// Search functionality
	if query.Search != "" {
		whereClause += ` AND (title LIKE ? OR author LIKE ?)`
		searchTerm := "%" + query.Search + "%"
		params = append(params, searchTerm, searchTerm)
	}

	// Price range filter
	if query.MinPrice > 0 {
		whereClause += ` AND price >= ?`
		params = append(params, query.MinPrice)
	}
	if query.MaxPrice > 0 {
		whereClause += ` AND price <= ?`
		params = append(params, query.MaxPrice)
	}

	// Author filter
	if query.Author != "" {
		whereClause += ` AND author = ?`
		params = append(params, query.Author)
	}

	// Add where clause to queries
	baseQuery += whereClause
	countQuery += whereClause

	// Sorting
	if query.SortBy != "" {
		direction := "ASC"
		if strings.ToUpper(query.SortDir) == "DESC" {
			direction = "DESC"
		}
		// Validate sort column to prevent SQL injection
		allowedColumns := map[string]bool{
			"title": true, "author": true, "price": true,
			"stock": true, "created_at": true,
		}
		if allowedColumns[query.SortBy] {
			baseQuery += fmt.Sprintf(" ORDER BY %s %s", query.SortBy, direction)
		}
	}

	// Count total records
	var total int
	err := r.db.QueryRow(countQuery, params...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Pagination
	if query.PageSize <= 0 {
		query.PageSize = 10 // Default page size
	}
	if query.Page <= 0 {
		query.Page = 1 // Default page
	}
	offset := (query.Page - 1) * query.PageSize

	baseQuery += " LIMIT ? OFFSET ?"
	params = append(params, query.PageSize, offset)

	// Execute final query
	rows, err := r.db.Query(baseQuery, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.Stock,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	totalPages := (total + query.PageSize - 1) / query.PageSize

	return &models.PaginatedResponse{
		Data:       books,
		Total:      total,
		Page:       query.Page,
		PageSize:   query.PageSize,
		TotalPages: totalPages,
	}, nil
}
