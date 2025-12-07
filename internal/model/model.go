package shortener

import "time"

type ShortenURLModel struct {
	ID        int
	Original  string
	Short     string
	CreatedAt time.Time
	Views     int
}
