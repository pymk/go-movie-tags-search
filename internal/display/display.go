package display

import (
	"fmt"

	"github.com/pymk/go-movie-tags-search/internal/models"
)

func Tags(tags []models.Tag) {
	for _, t := range tags {
		fmt.Println(t)
	}
}

func Movies(movies []models.Movie, limit int) {
	printInfo := false
	displayed := movies

	if len(movies) > limit {
		printInfo = true
		displayed = movies[:limit]
	}

	for _, m := range displayed {
		fmt.Println(m)
	}

	if printInfo {
		fmt.Printf("** Result: %d | Display: %d (use `limit` flag to adjust) **\n",
			len(movies), limit)
	}
}
