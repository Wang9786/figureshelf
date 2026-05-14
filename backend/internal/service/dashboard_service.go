package service

import (
	"context"
	"encoding/json"
	"time"

	"figureshelf-backend/internal/repository"

	"github.com/redis/go-redis/v9"
)

type DashboardService struct {
	figureRepo  *repository.FigureRepository
	redisClient *redis.Client
}

func NewDashboardService(figureRepo *repository.FigureRepository, redisClient *redis.Client) *DashboardService {
	return &DashboardService{
		figureRepo:  figureRepo,
		redisClient: redisClient,
	}
}

func (s *DashboardService) GetSummary(ctx context.Context, userID string) (map[string]interface{}, error) {
	cacheKey := "dashboard:summary:user:" + userID

	cachedValue, err := s.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		var cachedSummary map[string]interface{}

		if err := json.Unmarshal([]byte(cachedValue), &cachedSummary); err == nil {
			cachedSummary["cache"] = "hit"
			return cachedSummary, nil
		}
	}

	summary, err := s.figureRepo.GetDashboardSummary(ctx, userID)
	if err != nil {
		return nil, err
	}

	cacheBytes, err := json.Marshal(summary)
	if err == nil {
		s.redisClient.Set(ctx, cacheKey, cacheBytes, 5*time.Minute)
	}

	summary["cache"] = "miss"

	return summary, nil
}

func (s *DashboardService) ClearSummaryCache(ctx context.Context, userID string) error {
	cacheKey := "dashboard:summary:user:" + userID

	return s.redisClient.Del(ctx, cacheKey).Err()
}