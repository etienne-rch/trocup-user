package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pagination structure
type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

// PaginatedResult structure to return the paginated data and total count
type PaginatedResult struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

// Paginate function for reusability
func Paginate(collection *mongo.Collection, filter interface{}, pagination Pagination, result interface{}) (PaginatedResult, error) {
	skip := (pagination.Page - 1) * pagination.Limit
	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(pagination.Limit)

	// Query to count the total number of documents
	total, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return PaginatedResult{}, err
	}

	// Query to get the paginated result
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return PaginatedResult{}, err
	}
	defer cursor.Close(context.TODO())

	// Decode the documents into the result
	if err := cursor.All(context.TODO(), result); err != nil {
		return PaginatedResult{}, err
	}

	return PaginatedResult{
		Total: total,
		Data:  result,
	}, nil
}
