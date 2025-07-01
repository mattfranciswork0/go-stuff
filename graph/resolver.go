//go:generate go run github.com/99designs/gqlgen -v
package graph

import (
	"context"

	"example.com/this/package/graph/model"
	"github.com/mattfranciswork0/go-p2/models"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Replace with your DB logic
var albumList = []*model.Album{
	{ID: 1, Title: "Kind of Blue", Artist: "Miles Davis", Price: 12.99},
}

func (r *queryResolver) Albums(ctx context.Context) ([]*models.Album, error) {
	// If you have a real database:
	// rows, err := db.Query("SELECT id, title, artist, price FROM albums")
	// return rows.Scan() etc...

	return albumList, nil
}

type Resolver struct{}
