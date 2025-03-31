package store

import (
	"database/sql"

	"github.com/tbytes404/clipboard/model"
)

type BlobsStore struct {
	db *sql.DB
}

func NewBlobsStore(db *sql.DB) *BlobsStore {
	return &BlobsStore{db}
}

func (s *BlobsStore) FilterBlobs() (model.Blobs, error) {
	rows, err := s.db.Query("SELECT id, text FROM blobs")
	if err != nil {
		return nil, err
	}
	var blobs model.Blobs
	for rows.Next() {
		var blob model.Blob
		rows.Scan(&blob.ID, &blob.Text)
		blobs = append(blobs, blob)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return blobs, nil
}

func (s *BlobsStore) AddBlob(text string) (model.Blob, error) {
	res, err := s.db.Exec("INSERT INTO blobs ( text ) VALUES ( ? )", text)
	if err != nil {
		return model.Blob{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return model.Blob{}, err
	}
	return model.Blob{ID: id, Text: text}, nil
}

func (s *BlobsStore) DeleteBlob(id int64) error {
	_, err := s.db.Exec("DELETE FROM blobs WHERE id = ?", id)
	return err
}
