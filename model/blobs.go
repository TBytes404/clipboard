package model

type Blob struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`
}

type Blobs = []Blob
