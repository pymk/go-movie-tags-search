package config

import (
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	Limit    int
	ShowTags bool
	Tags     []string
}

func ParseFlags() (*Config, error) {
	limit := flag.Int("limit", 10, "number of results to display")
	showTags := flag.Bool("tags", false, "display top used tags")
	tagList := flag.String("search", "", "search movies by tags (comma-separated)")

	flag.Parse()

	if *showTags && *tagList != "" {
		return nil, fmt.Errorf("cannot use both --tags and --search flags")
	}

	if !*showTags && *tagList == "" {
		return nil, fmt.Errorf("must specify either --tags or --search flag")
	}

	config := &Config{
		Limit:    *limit,
		ShowTags: *showTags,
	}

	if *tagList != "" {
		config.Tags = strings.Split(*tagList, ",")
		for i, t := range config.Tags {
			config.Tags[i] = fmt.Sprintf("\"%s\"", strings.TrimSpace(t))
		}
	}

	return config, nil
}
