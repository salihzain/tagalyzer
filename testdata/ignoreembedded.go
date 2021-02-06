package testdata

import "time"

type EmbedMe struct {
	CreatedAt time.Time `json:"createdAt"`
}

type IgnoreEmbedded struct {
	EmbedMe
}
