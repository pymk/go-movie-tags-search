# go-movie-tags-search

A CLI tool to search for movies using tags.

This is a practice project for me to use Go and duckdb.

Initially, everything was in the `main.go` file, but I split it based on [this](https://go.dev/doc/modules/layout) and this [repository](https://github.com/golang-standards/project-layout).

## Usage

The `movies.duckdb` dataset should be created first (see Database section below).

```sh
cd go-movie-tags-search/cmd/go-movie-tags-search/main.go

# Display all available tags
go run . -tags
go run . -tags -limit=20

# Display movies with tags of interest
go run . -search=sci-fi
go run . -search=sci-fi,comedy
go run . -search=sci-fi,comedy -limit=20
```

## Datasets

The program's `search` command uses the "tags" and "movies" tables:

```
> duckdb movies.duckdb "
WITH temp AS    (PIVOT tags ON tag IN (action, assassin))
SELECT DISTINCT t.movieId, m.title, m.genres
FROM            temp t
LEFT JOIN       movies m ON m.movieId = t.movieId
WHERE           t.assassin = 1 AND action = 1
LIMIT           10
;
"
┌─────────┬────────────────────┬─────────────────────────────────┐
│ movieId │       title        │             genres              │
│  int64  │      varchar       │             varchar             │
├─────────┼────────────────────┼─────────────────────────────────┤
│  166480 │ Eliminators (2016) │ Action|Thriller                 │
│   89087 │ Colombiana (2011)  │ Action|Adventure|Drama|Thriller │
│  115149 │ John Wick (2014)   │ Action|Thriller                 │
└─────────┴────────────────────┴─────────────────────────────────┘
```

The `tags` command uses the "tags" table:

```
> duckdb movies.duckdb "
SELECT       tag, COUNT(tag) AS count
FROM         tags
GROUP BY ALL
ORDER BY     count DESC
LIMIT        20
;
"
┌────────────────────┬───────┐
│        tag         │ count │
│      varchar       │ int64 │
├────────────────────┼───────┤
│ sci-fi             │ 10996 │
│ atmospheric        │  9589 │
│ action             │  8473 │
│ comedy             │  8139 │
│ funny              │  7467 │
│ surreal            │  7231 │
│ visually appealing │  7090 │
│ based on a book    │  6617 │
│ twist ending       │  6521 │
│ dark comedy        │  6053 │
│ thought-provoking  │  6000 │
│ dystopia           │  5646 │
│ violence           │  5389 │
│ cinematography     │  5277 │
│ romance            │  5271 │
│ murder             │  5180 │
│ social commentary  │  5152 │
│ stylized           │  4917 │
│ fantasy            │  4844 │
│ psychology         │  4827 │
├────────────────────┴───────┤
│ 20 rows          2 columns │
└────────────────────────────┘
```

Note that there are duckdb-specific commands in these queries.

The `movies.duckdb` was created using the datasets from [here](https://grouplens.org/datasets/movielens/):

```sql
CREATE TABLE movies (movieId int64, title varchar, genres varchar); copy movies from 'movies.csv';
CREATE TABLE links (movieId int64, imdbId varchar, tmdbId int64); COPY links FROM 'links.csv';
CREATE TABLE tags (userId int64, movieId int64, tag varchar, "timestamp" int64); copy tags from 'tags.csv';
CREATE TABLE ratings (userId int64, movieId int64, rating double, "timestamp" int64); copy ratings from 'ratings.csv';
```
