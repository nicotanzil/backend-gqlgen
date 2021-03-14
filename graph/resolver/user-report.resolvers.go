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

func (r *mutationResolver) CreateUserReport(ctx context.Context, reporterID int, reportedID int, description string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	report := model.UserReport{
		Reported:    &model.User{ID: reportedID},
		Reporter:    &model.User{ID: reporterID},
		Description: description,
	}

	db.Create(&report)

	var reportCount int64

	db.Model(model.UserReport{}).
		Where("reported_id = ?", reportedID).
		Where("DATE_PART('day', CURRENT_DATE) - DATE_PART('day', created_at) < 7").
		Count(&reportCount)

	if reportCount > 5 {
		// Suspend user
		var user model.User
		db.Where("id = ?", reportedID).First(&user)
		user.IsSuspend = true
		db.Save(&user)
	}

	return true, nil
}

func (r *queryResolver) Reports(ctx context.Context) ([]*model.UserReport, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var reports []*model.UserReport

	db.Preload(clause.Associations).Find(&reports)

	return reports, nil
}

func (r *queryResolver) GetReportByReported(ctx context.Context, id int) ([]*model.UserReport, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var reports []*model.UserReport

	db.Preload(clause.Associations).Where("reported_id = ?", id).Find(&reports)
	fmt.Println(len(reports))
	return reports, nil
}
