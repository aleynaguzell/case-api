package record

import "time"

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int64     `json:"totalCount"`
}
