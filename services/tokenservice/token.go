package tokenservice

import "time"

type Token struct {
	ID          int64
	UserID      int64
	Token       string
	IsValid     bool
	CreatedDate time.Time
}
