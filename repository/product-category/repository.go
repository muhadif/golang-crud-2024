package product_category

import (
	"context"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/redis"
	"gorm.io/gorm"
	"log"
)

func NewProductCategoryRepository(db *gorm.DB, redis *redis.RedisClient) repository.ProductCategoryRepository {
	return &repo{
		repoRedis: NewProductCategoryRedisRepository(redis),
		repoMysql: NewProductCategoryMysqlRepository(db),
	}
}

type repo struct {
	repoMysql repository.ProductCategoryRepository
	repoRedis repository.ProductCategoryRedisRepository
}

func (r repo) GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error) {
	resp, err := r.repoRedis.GetProductCategory(ctx)
	if err == nil && resp != nil {
		return resp, nil
	}

	resp, err = r.repoMysql.GetProductCategory(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := r.repoRedis.SetProductCategory(context.Background(), resp); err != nil {
			log.Fatalf("error SetProductCategory : %v", err)
		}
	}()

	return resp, nil
}
