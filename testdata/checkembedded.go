package testdata

import "time"

type EmbedMeAndCheck struct {
	CreatedAt time.Time `json:"createdAt"`
}

type IgnoreEmbedded struct {
	EmbedMeAndCheck // want "field:EmbedMeAndCheck is missing tag:json"
}
