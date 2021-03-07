package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreatePromo(ctx context.Context, input model.NewPromo) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var newPromo model.Promo

	newPromo = model.Promo{
		DiscountPercentage: input.DiscountPercentage,
		ValidUntil:         input.ValidUntil,
	}

	db.Create(&newPromo)

	return &newPromo, nil
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, input model.NewPromo, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promo model.Promo

	db.Where("id = ?", id).First(&promo)
	promo.ValidUntil = input.ValidUntil
	promo.DiscountPercentage = input.DiscountPercentage

	db.Save(&promo)
	return &promo, nil
}

func (r *mutationResolver) DeletePromo(ctx context.Context, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promo model.Promo

	db.Preload(clause.Associations).Where("id = ?", id).First(&promo)

	db.Delete(&model.Promo{}, id)
	return &promo, nil
}

func (r *queryResolver) Promos(ctx context.Context) ([]*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promos []*model.Promo

	db.Preload(clause.Associations).Find(&promos)

	return promos, nil
}

func (r *queryResolver) GetTotalPromo(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(&model.Promo{}).Count(&count)

	return int(count), nil
}

func (r *queryResolver) GetPromoByID(ctx context.Context, id int) (*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promo model.Promo

	db.Preload(clause.Associations).Where("id = ?", id).First(&promo)

	return &promo, nil
}

func (r *queryResolver) GetPromoPaginationAdmin(ctx context.Context, page *int) ([]*model.Promo, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var promos []*model.Promo

	db.Preload(clause.Associations).Limit(providers.ADMIN_PROMO_PAGINATION).Offset(providers.ADMIN_PROMO_PAGINATION * (*page - 1)).Find(&promos)

	return promos, nil
}
