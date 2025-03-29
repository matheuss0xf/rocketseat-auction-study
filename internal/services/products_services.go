package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/matheuss0xf/gobid/internal/store/pgstore"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var ErrProductNotFound = errors.New("product not found")

func (ps *ProductService) CreateProduct(ctx context.Context, sellerId string, productName string, description string, base_price float64, auctionEnd time.Time) (string, error) {

	IdStr := uuid.New().String()
	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		ID:          IdStr,
		SellerID:    sellerId,
		ProductName: productName,
		Description: description,
		BasePrice:   base_price,
		AuctionEnd:  auctionEnd,
	})

	if err != nil {
		return uuid.UUID{}.String(), err
	}

	return id, nil
}

func (ps *ProductService) GetProductById(ctx context.Context, productId string) (pgstore.Product, error) {
	product, err := ps.queries.GetProductById(ctx, productId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Product{}, ErrProductNotFound
		}
		return pgstore.Product{}, err
	}

	return product, nil
}
