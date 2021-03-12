package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/nicotanzil/backend-gqlgen/app/providers"
	"github.com/nicotanzil/backend-gqlgen/database"
	"github.com/nicotanzil/backend-gqlgen/graph/model"
	"gorm.io/gorm/clause"
)

func (r *mutationResolver) CreateGame(ctx context.Context, input model.NewGame) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	genres := []*model.Genre{}
	tags := []*model.Tag{}
	developers := []*model.Developer{}

	for i := 0; i < len(input.Genres); i++ {
		curr := model.Genre{ID: input.Genres[i].ID}
		genres = append(genres, &curr)
	}

	for i := 0; i < len(input.Tags); i++ {
		curr := model.Tag{ID: input.Tags[i].ID}
		tags = append(tags, &curr)
	}

	for i := 0; i < len(input.Developers); i++ {
		curr := model.Developer{ID: input.Developers[i].ID}
		developers = append(developers, &curr)
	}

	var newGame model.Game

	newGame = model.Game{
		Name:          input.Name,
		Description:   input.Description,
		ReleaseDate:   time.Now(),
		Genres:        genres,
		Tags:          tags,
		OriginalPrice: input.OriginalPrice,
		PromoID:       providers.DEFAULT_PROMO_ID,
		GamePlayHour:  0,
		GameReviews:   nil,
		Developers:    developers,
		PublisherID:   input.PublisherID,
		SystemID:      input.SystemID,
		Users:         nil,
	}

	db.Create(&newGame)
	return &newGame, nil
}

func (r *mutationResolver) UpdateGame(ctx context.Context, id int, input model.NewGame) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	genres := []*model.Genre{}
	tags := []*model.Tag{}
	developers := []*model.Developer{}

	for i := 0; i < len(input.Genres); i++ {
		curr := model.Genre{ID: input.Genres[i].ID}
		genres = append(genres, &curr)
	}

	for i := 0; i < len(input.Tags); i++ {
		curr := model.Tag{ID: input.Tags[i].ID}
		tags = append(tags, &curr)
	}

	for i := 0; i < len(input.Developers); i++ {
		curr := model.Developer{ID: input.Developers[i].ID}
		developers = append(developers, &curr)
	}

	var game model.Game

	db.Where("id = ?", id).First(&game)

	game.Name = input.Name
	game.Description = input.Description
	game.Genres = genres
	game.Tags = tags
	game.OriginalPrice = input.OriginalPrice
	game.Developers = developers
	game.PublisherID = input.PublisherID
	game.SystemID = input.SystemID

	db.Save(&game)

	return &game, nil
}

func (r *mutationResolver) DeleteGame(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	db.Exec("DELETE FROM game_images WHERE game_id = ?", id)
	db.Exec("DELETE FROM game_videos WHERE game_id = ?", id)

	var game model.Game

	db.Where("id = ?", id).First(&game)

	db.Delete(&model.Game{}, id)
	return &game, nil
}

func (r *mutationResolver) InsertGameBanner(ctx context.Context, id int, link string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var game model.Game

	db.Where("id = ?", id).First(&game)

	game.Banner = link
	db.Save(&game)
	return true, nil
}

func (r *mutationResolver) SetGamePromo(ctx context.Context, gameID int, promoID int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var game model.Game

	db.Where("id = ?", gameID).First(&game)
	game.Promo = &model.Promo{ID: promoID}
	db.Save(&game)

	var users []model.User
	db.Joins("JOIN wishlists ON wishlists.user_id = users.id").
		Where("wishlists.game_id = ?", game.ID).
		Find(&users)

	// INSERT FIND GAME LOGIC AND USER
	// SEND EMAIL
	for _, user := range users {
		fmt.Println("Sending email to ", user.Email)
		model.SendEmailPromo(game, user)
	}

	return &game, nil
}

func (r *queryResolver) Games(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	db.Preload(clause.Associations).Find(&games)

	return games, nil
}

func (r *queryResolver) GameByID(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var game model.Game

	db.Preload(clause.Associations).Where("id = ?", id).First(&game)

	return &game, nil
}

func (r *queryResolver) GameByMultipleID(ctx context.Context, ids []int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	for _, id := range ids {
		var temp model.Game
		db.Preload(clause.Associations).First(&temp, "id = ?", id)
		games = append(games, &temp)
	}

	return games, nil
}

func (r *queryResolver) GetGameByPromoID(ctx context.Context, id int) (*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var game model.Game

	db.Preload(clause.Associations).Where("promo_id = ?", id).First(&game)
	return &game, nil
}

func (r *queryResolver) GetGamePaginationAdmin(ctx context.Context, page *int) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	db.Preload(clause.Associations).Limit(providers.ADMIN_GAME_PAGINATION).Offset(providers.ADMIN_GAME_PAGINATION * (*page - 1)).Find(&games)

	return games, nil
}

func (r *queryResolver) GetTotalGame(ctx context.Context) (int, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var count int64

	db.Model(&model.Game{}).Count(&count)

	return int(count), nil
}

func (r *queryResolver) GameSearch(ctx context.Context, keyword string) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	keyword = "%" + keyword + "%"
	db.Preload(clause.Associations).Limit(providers.GAME_SEARCH_LIMIT).Where("LOWER(name) LIKE LOWER(?)", keyword).Find(&games)

	return games, nil
}

func (r *queryResolver) GameSearchPage(ctx context.Context, keyword string, page int, price int, tag []*model.InputTag) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game
	var tagsId []int

	if len(tag) > 0 {
		for i := 0; i < len(tag); i++ {
			tagsId = append(tagsId, tag[i].ID)
		}

		keyword = "%" + keyword + "%"
		db.Preload(clause.Associations).
			Limit(providers.GAME_SEARCH_PAGE_LIMIT).
			Offset(providers.GAME_SEARCH_PAGE_LIMIT*(page-1)).
			Joins("JOIN game_tags ON game_tags.game_id = games.id").
			Where("LOWER(name) LIKE LOWER(?)", keyword).
			Where("original_price <= ?", price).
			Where("game_tags.tag_id IN ?", tagsId).
			Find(&games)
	} else {
		keyword = "%" + keyword + "%"
		db.Preload(clause.Associations).
			Limit(providers.GAME_SEARCH_PAGE_LIMIT).
			Offset(providers.GAME_SEARCH_PAGE_LIMIT*(page-1)).
			Where("LOWER(name) LIKE LOWER(?)", keyword).
			Where("original_price <= ?", price).
			Find(&games)
	}

	return games, nil
}

func (r *queryResolver) GetSpecialOfferGame(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	db.Preload(clause.Associations).Where("promo_id <> ?", 1).Find(&games)

	return games, nil
}

func (r *queryResolver) GetNewTrendingGame(ctx context.Context) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game

	db.Preload(clause.Associations).Order("created_at desc").Limit(providers.DEFAULT_HOME_CATEGORY_ITEM).Find(&games)

	return games, nil
}

func (r *queryResolver) GetGamesForDiscussions(ctx context.Context, keyword string) ([]*model.Game, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var games []*model.Game
	keyword = "%" + keyword + "%"
	db.Preload("Discussions.User").
		Preload("Discussions.Replies").
		Where("LOWER(name) LIKE LOWER(?)", keyword).
		Find(&games)

	return games, nil
}

func (r *queryResolver) GetTopSellerGames(ctx context.Context) ([]*model.TopSellerGame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var results []*model.TopSellerGame

	/**
	SELECT game_id AS gameId, g.name AS game_name, g.original_price AS gamePrice, p.discount_percentage AS gameDiscount, g.banner AS gameBanner,COUNT(game_id) AS purchaseCount
	FROM transaction_details td
	JOIN transaction_headers th ON td.transaction_header_id = th.id
	JOIN games g ON g.id = td.game_id
	JOIN promos p ON g.promo_id = p.id
	WHERE th.created_at BETWEEN(now() - '1 week'::interval) AND now()
	GROUP BY game_id, g.name, g.original_price, g.banner, p.discount_percentage
	ORDER BY COUNT(game_id) desc
	*/

	db.Raw("SELECT td.game_id AS game_id, g.name AS game_name, g.original_price AS game_price, p.discount_percentage AS game_discount, g.banner AS game_banner,COUNT(td.game_id) AS purchase_count" +
		" FROM transaction_details td " +
		" JOIN transaction_headers th ON td.transaction_header_id = th.id" +
		" JOIN games g ON g.id = td.game_id" +
		" JOIN promos p ON g.promo_id = p.id" +
		" WHERE th.created_at BETWEEN(now() - '1 week'::interval) AND now()" +
		" GROUP BY game_id, g.name, g.original_price, g.banner, p.discount_percentage" +
		" ORDER BY COUNT(game_id) desc").
		Limit(5).
		Scan(&results)

	return results, nil
}

func (r *queryResolver) GetTopReviewGames(ctx context.Context) ([]*model.TopReviewGame, error) {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	dbClose, _ := db.DB()
	defer dbClose.Close()

	var results []*model.TopReviewGame

	db.Raw("SELECT g.id AS game_id, g.name AS game_name, g.original_price AS game_price, g.banner AS game_banner, " +
		"\np.discount_percentage AS game_discount, COUNT(is_recommended) AS review_count" +
		"\nFROM community_game_reviews cgr " +
		"\nJOIN games g ON cgr.game_id = g.id" +
		"\nJOIN promos p ON g.promo_id = p.id" +
		"\nWHERE cgr.is_recommended = true" +
		"\nGROUP BY g.id, g.name, g.original_price, g.banner, p.discount_percentage" +
		"\nORDER BY COUNT(is_recommended) DESC").
		Limit(5).
		Scan(&results)

	return results, nil
}
