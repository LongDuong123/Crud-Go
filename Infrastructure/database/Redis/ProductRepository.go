package redis

import (
	"context"
	domain "crud/Domain"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type ProductRepositoryByRedis struct {
	Databaseredis *databaseRedis
}

func NewProductRepositoryByRedis(dbr *databaseRedis) domain.ProductRepository {
	return &ProductRepositoryByRedis{Databaseredis: dbr}
}

func (prredis *ProductRepositoryByRedis) GetByID(id int) (*domain.Product, error) {
	val, err := prredis.Databaseredis.redis.Get(context.Background(), strconv.Itoa(id)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	var book domain.Product
	err = json.Unmarshal([]byte(val), &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (prredis *ProductRepositoryByRedis) GetByPage(start int, end int) ([]domain.Product, error) {
	return nil, errors.New("something went wrong")
}

func (prredis *ProductRepositoryByRedis) Create(Product *domain.Product) error {
	productJson, err := json.Marshal(Product)
	if err != nil {
		return err
	}
	err = prredis.Databaseredis.redis.Set(context.Background(), strconv.Itoa(Product.ID), productJson, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (prredis *ProductRepositoryByRedis) UpdateByID(id int, Product *domain.Product) error {
	return errors.New("something went wrong")
}

func (prredis *ProductRepositoryByRedis) DeleteByID(id int) error {
	prredis.Databaseredis.redis.Del(context.Background(), strconv.Itoa(id))
	return nil
}
