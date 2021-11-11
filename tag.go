package ingresso

import (
	"fmt"
)

// Tag type represents Ingresso.com's Database Tag
type Tag struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	GroupingActive bool   `json:"groupingActive"`
	Priority       int    `json:"priority"`
	ShowAllEnabled bool   `json:"showAllEnabled"`
	URLKey         string `json:"urlKey"`
}

// GetTags Get all tags
// https://api-content.ingresso.com/v0/tags/getAllTags
func (ing *Ingresso) GetTags() ([]Tag, error) {
	var tags []Tag
	url := fmt.Sprintf("%stags/getAllTags", ing.BaseURL)
	_, err := ing.getIngresso(url, &tags)
	return tags, err
}
