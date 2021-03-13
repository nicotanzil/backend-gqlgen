package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/generated"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
)

func (r *itemTransactionResolver) DeletedAt(ctx context.Context, obj *model.ItemTransaction) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPreviousTransactionData(ctx context.Context, typeID int) ([]*model.GraphData, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	dbClose, _ := db.DB()
	defer dbClose.Close()

	//var transactions []*model.ItemTransaction

	var graphDatas []*model.GraphData

	//db.Order("created_at asc").
	//	Where("item_id = ?", typeID).
	//	Where("created_at BETWEEN (now() - '1 week'::interval) AND now()").
	//	Preload(clause.Associations).
	//	Find(&transactions)

	db.Raw("SELECT AVG(price)::numeric(10,0) AS price, DATE(created_at) as date"+
		"\nFROM item_transactions it JOIN items i ON it.item_id = i.id"+
		"\nWHERE created_at BETWEEN (now() - '1 week'::interval) AND now()"+
		"\nAND item_type_id = ?"+
		"\nGROUP BY DATE(created_at)"+
		"\nORDER BY DATE(created_at) asc", typeID).
		Scan(&graphDatas)

	return graphDatas, nil
}

// ItemTransaction returns generated.ItemTransactionResolver implementation.
func (r *Resolver) ItemTransaction() generated.ItemTransactionResolver {
	return &itemTransactionResolver{r}
}

type itemTransactionResolver struct{ *Resolver }
