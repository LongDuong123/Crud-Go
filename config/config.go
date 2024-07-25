package config

import (
	domain "crud/Domain"
	mysql "crud/Infrastructure/database/MySQL"
	redis "crud/Infrastructure/database/Redis"
)

type Config struct {
	RepositoryUser         domain.UserRepository
	RepositoryProduct      domain.ProductRepository
	RepositoryProductRedis domain.ProductRepository
}

func InitializeAppConfig() (*Config, error) {
	databaseMySql, err := mysql.ConnnectMySql()
	if err != nil {
		return nil, err
	}
	databaseRedis := redis.ConnnectRedis()
	repositoryUser := mysql.NewUserRepository(databaseMySql)
	repositoryProduct := mysql.NewProductRepository(databaseMySql)
	repositoryProductRedis := redis.NewProductRepositoryByRedis(databaseRedis)
	return &Config{RepositoryUser: repositoryUser, RepositoryProduct: repositoryProduct, RepositoryProductRedis: repositoryProductRedis}, nil
}
