package appl

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mindlessdrone/go-podcasts/model"
)

type SQLRepository struct {
	DefaultRepository
	db *sql.DB
}

func NewSQLRepository(dbName string) (*SQLRepository, error) {
	sqlRepo := &SQLRepository{}

	db, err := sql.Open("sqlite3", "podcasts.db")
	if err != nil {
		return nil, err
	}

	// create tables if they do not already exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS feeds(
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			title		TEXT NOT NULL,
			description	TEXT NOT NULL,
			icon_url	TEXT NOT NULL,
			url			TEXT NOT NULL,
			author		TEXT NOT NULL,
			updated		DATETIME NOT NULL,
			feedUpdated DATETIME NOT NULL
		);`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS episodes(
			id 			INTEGER PRIMARY KEY AUTOINCREMENT,
			feed_id 	INTEGER NOT NULL,
			title		TEXT NOT NULL,
			description	TEXT NOT NULL,
			published	DATETIME NOT NULL,
			guid		TEXT NOT NULL,
			played		INTEGER NOT NULL,
			length		INTEGER NOT NULL,
			url			TEXT NOT NULL,
			FOREIGN KEY(feed_id) REFERENCES feeds(id)
		);`)
	if err != nil {
		return nil, err
	}

	sqlRepo.db = db
	return sqlRepo, nil
}

func (repo SQLRepository) Add(feed *model.Feed) error {

	updatedString := feed.Updated.Format("2006-01-02 15:04")
	feedUpdatedString := feed.FeedUpdated.Format("2006-01-02 15:04")
	insertStmt, err := repo.db.Prepare(`
		INSERT INTO feeds(id, title, description, icon_url, url, author, updated, feedUpdated)
		SELECT null, ?, ?, ?, ?, ?, datetime(?), datetime(?)
		WHERE NOT EXISTS (SELECT * FROM feeds WHERE url = ?);`)
	if err != nil {
		return err
	}

	insertResult, err := insertStmt.Exec(
		feed.Title,
		feed.Description,
		feed.IconURL,
		feed.FeedURL,
		feed.Author,
		updatedString,
		feedUpdatedString, feed.FeedURL)

	if err != nil {
		return err
	}

	rowsInserted, _ := insertResult.RowsAffected()
	if rowsInserted == 0 {
		return errors.New("Feed already exists")
	}

	// insert episodes
	episodeStmt, err := repo.db.Prepare("INSERT INTO episodes VALUES (null, ?, ?, ?, datetime(?), ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	for i, episode := range feed.Episodes {
		if i != 0 {
			episode.SetPlayed()
		}
		publishedString := episode.Published().Format("2006-01-02 15:04")
		feedID, _ := insertResult.LastInsertId()
		_, err = episodeStmt.Exec(
			feedID,
			episode.Title(),
			episode.Description(),
			publishedString,
			episode.GUID(),
			episode.Played(),
			episode.Length(),
			episode.URL())
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo SQLRepository) Query(id int) (*model.Feed, error) {
	return nil, nil
}

func (repo SQLRepository) QueryAll() ([]*model.Feed, error) {
	return nil, errors.New("Not implemented")
}
