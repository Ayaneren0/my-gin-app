package repository

import (
	"bookstore/internal/models"
	"database/sql"
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
