package mysql

import (
	"1_assignment/pkg/models"
	"database/sql"
	"strconv"
)

type ArticleModel struct {
	DB *sql.DB
}

func (m *ArticleModel) Insert(category, author, readership, title, description, content string) (int, error) {
	stmt := `INSERT INTO articles (category, author, readership, title, description, content, created) VALUES (?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`
	result, err := m.DB.Exec(stmt, category, author, readership, title, description, content)
	if err != nil {
		return 0, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err

	}
	return int(id), nil
}

func (m *ArticleModel) Latest() ([]*models.Article, error) {
	stmt := `SELECT id, category, author, readership, title, description, content, created FROM articles LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*models.Article{}

	for rows.Next() {
		a := &models.Article{}

		err = rows.Scan(&a.ID, &a.Category, &a.Author, &a.Readership, &a.Title, &a.Description, &a.Content, &a.PublishedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *ArticleModel) GetByCategory(readership string) ([]*models.Article, error) {
	stmt := `SELECT id, category, author, readership, title, description, content, created FROM articles WHERE readership = ? LIMIT 10`
	rows, err := m.DB.Query(stmt, readership)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*models.Article{}

	for rows.Next() {
		a := &models.Article{}

		err = rows.Scan(&a.ID, &a.Category, &a.Author, &a.Readership, &a.Title, &a.Description, &a.Content, &a.PublishedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *ArticleModel) Delete(id int) (bool, error) {
	stmt := `DELETE FROM articles WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *ArticleModel) GetById(id int) (*models.Article, error) {
	stmt := `SELECT id, category, author, readership, title, description, content, created FROM articles WHERE id = ?`
	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	a := &models.Article{}

	for rows.Next() {
		err := rows.Scan(&a.ID, &a.Category, &a.Author, &a.Readership, &a.Title, &a.Description, &a.Content, &a.PublishedAt)
		if err != nil {
			return nil, err
		}
	}

	return a, nil
}

func (m *ArticleModel) Update(id, category, author, title, description, content, readership string) (bool, error) {
	stmt := `UPDATE articles SET category = ?, author = ?, title = ?, description = ?, content = ?, readership = ?, created = UTC_TIMESTAMP() WHERE id = ?`
	num, _ := strconv.Atoi(id)
	_, err := m.DB.Exec(stmt, category, author, title, description, content, readership, num)
	if err != nil {
		return false, err
	}
	return true, nil
}
