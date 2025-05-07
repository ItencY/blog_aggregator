package main

import (
	"github.com/itency/blog_aggregator/internal/config"
	"github.com/itency/blog_aggregator/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}
