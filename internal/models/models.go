package models

import "fmt"

type Movie struct {
	ID     int
	Title  string
	Genres string
}

type Tag struct {
	Name  string
	Count int
}

func (m Movie) String() string {
	return fmt.Sprintf("Title: %s\nGenres: %s\n",
		m.Title, m.Genres)
}

func (t Tag) String() string {
	return fmt.Sprintf("%s", t.Name)
}
