package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *queryResolver) Reports(ctx context.Context) ([]*model.UserReport, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var reports []*model.UserReport

	db.Preload(clause.Associations).Find(&reports)

	return reports, nil
}

func (r *queryResolver) GetReportByReported(ctx context.Context, id int) ([]*model.UserReport, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	var reports []*model.UserReport

	db.Preload(clause.Associations).Where("reported_id = ?", id).Find(&reports)
	fmt.Println(len(reports))
	return reports, nil
}
