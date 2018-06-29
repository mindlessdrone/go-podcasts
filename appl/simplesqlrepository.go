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
				id			INTEGER PRIMARY KEY AUTOINCREMENT,
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
		INSERT INTO feeds(id, url, feed_data)
		SELECT null, ?, ?
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

func (repo SimpleSQLRepository) QueryAll() ([]*model.Feed, error) {
	feeds := make([]*model.Feed, 0)

	rows, err := repo.db.Query("SELECT feed_data FROM feeds;")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var feedData []byte
		var feed model.Feed

		err = rows.Scan(&feedData)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(feedData, &feed)
		if err != nil {
			return nil, err
		}

		feeds = append(feeds, &feed)
	}
	return feeds, nil
}

func (repo SimpleSQLRepository) ItemIDs() ([]int, error) {
	ids := make([]int, 0)

	rows, err := repo.db.Query("SELECT id FROM feeds;")
	if err != nil {
		return nil, err
	}

	// gather all ids into slice
	var id int
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}
