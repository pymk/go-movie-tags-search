package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pymk/go-movie-tags-search/internal/models"

	_ "github.com/marcboeker/go-duckdb"
)

func Connect(dbPath string) (*sql.DB, error) {
	return sql.Open("duckdb", dbPath)
}

func GetTopTags(db *sql.DB, limit int) ([]models.Tag, error) {
	query := fmt.Sprintf(`
		SELECT 		tag, COUNT(tag) AS count
		FROM 		tags
		GROUP BY ALL
		ORDER BY 	count DESC
		LIMIT 		%d
		;
		`, limit)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var t models.Tag
		if err = rows.Scan(&t.Name, &t.Count); err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}
		tags = append(tags, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reading row: %w", err)
	}

	return tags, nil
}

func SearchMovies(db *sql.DB, tags []string) ([]models.Movie, error) {
	conditions := make([]string, len(tags))

	// TODO: use args with prepared statements in
	// the PIVOT clause to prevent SQL injection.
	// The strings.Join() function takes a slice (not []interface)
	// so I'd need to convert it to a slice first.
	args := make([]interface{}, len(tags))

	for i, t := range tags {
		conditions[i] = fmt.Sprintf("t.%s = 1", t)
	}

	whereClause := strings.Join(conditions, " AND ")
	pivotClause := strings.Join(tags, ",")

	query := fmt.Sprintf(`
   	WITH 			temp AS (PIVOT tags ON tag IN (%s))
   	SELECT DISTINCT t.movieId,
   					m.title,
   					m.genres
   	FROM 			temp t
   	LEFT JOIN 		movies m ON m.movieId = t.movieId
   	WHERE 			%s
   	;
   	`, pivotClause, whereClause)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("executing query: %w", err)
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err = rows.Scan(&m.ID, &m.Title, &m.Genres); err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}
		movies = append(movies, m)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("reading row: %w", err)
	}

	return movies, nil
}
