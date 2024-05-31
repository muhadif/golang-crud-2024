package product_category

import (
	"context"
	"encoding/json"
	"golang-crud-2024/core/entity"
	"golang-crud-2024/core/repository"
	"golang-crud-2024/pkg/redis"
)

const (
	productCategoryKey = "product-category"
)

func NewProductCategoryRedisRepository(redis *redis.RedisClient) repository.ProductCategoryRedisRepository {
	return &repoRedis{
		redis: redis,
	}
}

type repoRedis struct {
	redis *redis.RedisClient
}

func (r repoRedis) GetProductCategory(ctx context.Context) ([]*entity.ProductCategory, error) {
	raw, err := r.redis.Get(productCategoryKey)
	if err != nil {
		return nil, err
	}

	var resp []*entity.ProductCategory
	err = json.Unmarshal([]byte(raw), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r repoRedis) SetProductCategory(ctx context.Context, req []*entity.ProductCategory) error {
	raw, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return r.redis.Set(productCategoryKey, raw)
}
