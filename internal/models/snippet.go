package models

import "time"

type Snippet struct {
    ID        int
    Title     string
    Tags      []string
    CreatedAt time.Time
    Content   string
}
