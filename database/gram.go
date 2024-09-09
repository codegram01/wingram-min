package database

import (
	"database/sql"

	"github.com/codegram01/wingram-min/gram"
)

func (db *Db) GramCreate(gramReq *gram.Gram) (*gram.Gram, error) {
	var gram gram.Gram

	queryStr := `
		INSERT INTO gram 
			(title, description, content)
		VALUES 
			($1, $2, $3) 
		RETURNING id, title, description, content
	`
	row := db.QueryRow(
		queryStr,
		gramReq.Title,
		gramReq.Description,
		gramReq.Content,
	)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
	)

	return &gram, err
}

func (db *Db) GramUpdate(id int64, gramReq *gram.Gram) (*gram.Gram, error) {
	var gram gram.Gram

	queryStr := `
		UPDATE gram
		SET
			title = $2,
			description = $3,
			content = $4
		WHERE id = $1
		RETURNING id, title, description, content
	`
	row := db.QueryRow(
		queryStr,
		id,
		gramReq.Title,
		gramReq.Description,
		gramReq.Content,
	)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
	)

	return &gram, err
}

func (db *Db) GramDetail(id int64) (*gram.Gram, error) {
	var gram gram.Gram

	queryStr := `
		SELECT
			id, title, description, content
		FROM gram
		WHERE id = $1
	`
	row := db.QueryRow(queryStr, id)

	err := row.Scan(
		&gram.Id,
		&gram.Title,
		&gram.Description,
		&gram.Content,
	)

	return &gram, err
}

func (db *Db) GramDelete(id int64) error {
	queryStr := `DELETE FROM gram WHERE id = $1`
	_, err := db.Exec(queryStr, id)

	return err
}

func (db *Db) GramList() ([]*gram.Gram, error) {
	var grams []*gram.Gram

	query := `
		SELECT
			id, title, description, content
		FROM
			gram
	`
	err := db.Query(query, func(rows *sql.Rows) error {
		var gram gram.Gram

		err := rows.Scan(
			&gram.Id,
			&gram.Title,
			&gram.Description,
			&gram.Content,
		)

		if err != nil {
			return err
		}
		grams = append(grams, &gram)
		return nil
	})

	return grams, err
}
