package appl

import (
	"database/sql"
	"encoding/json"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mindlessdrone/go-podcasts/model"
)

type SimpleSQLRepository struct {
	DefaultRepository
	db *sql.DB
}

func NewSimpleSQLRepository(dbName string) (*SimpleSQLRepository, error) {
	var repo SimpleSQLRepository
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	repo.db = db

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS feeds(
				url 		TEXT NOT NULL,
				feed_data 	BLOB NOT NULL
			);`)
	if err != nil {
		return nil, err
	}

	return &repo, nil
}

func (repo SimpleSQLRepository) Add(feed *model.Feed) error {
	insert, err := repo.db.Prepare(`
		INSERT INTO feeds(url, feed_data)
		SELECT ?, ?
		WHERE NOT EXISTS(SELECT * FROM feeds WHERE url = ?);`)

	if err != nil {
		return err
	}

	feedData, err := json.Marshal(feed)
	if err != nil {
		return err
	}

	result, err := insert.Exec(feed.FeedURL, feedData, feed.FeedURL)
	if err != nil {
		return err
	}

	rowsInserted, _ := result.RowsAffected()
	if rowsInserted == 0 {
		return errors.New("Feed already exists")
	}

	return nil
}
