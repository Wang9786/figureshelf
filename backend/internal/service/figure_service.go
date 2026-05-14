package service

import (
	"context"
	"errors"
	"time"

	"figureshelf-backend/internal/model"
	"figureshelf-backend/internal/repository"
)

type FigureService struct {
	figureRepo       *repository.FigureRepository
	dashboardService *DashboardService
}

func NewFigureService(figureRepo *repository.FigureRepository, dashboardService *DashboardService) *FigureService {
	return &FigureService{
		figureRepo:       figureRepo,
		dashboardService: dashboardService,
	}
}

func (s *FigureService) Create(ctx context.Context, userID string, req model.CreateFigureRequest) (*model.FigureResponse, error) {
	status := req.Status
	if status == "" {
		status = "wishlist"
	}

	if !isValidFigureStatus(status) {
		return nil, errors.New("invalid figure status")
	}

	figure := &model.Figure{
		UserID:            userID,
		Name:              req.Name,
		CharacterName:     req.CharacterName,
		SeriesName:        req.SeriesName,
		Manufacturer:      req.Manufacturer,
		FigureType:         req.FigureType,
		Scale:             req.Scale,
		Status:            status,
		Price:             req.Price,
		Deposit:           req.Deposit,
		Balance:           req.Balance,
		PreorderStartDate: parseDatePtr(req.PreorderStartDate),
		PreorderDeadline:  parseDatePtr(req.PreorderDeadline),
		ReleaseDate:       parseDatePtr(req.ReleaseDate),
		PaymentDueDate:    parseDatePtr(req.PaymentDueDate),
		ArrivalDate:       parseDatePtr(req.ArrivalDate),
		ShopName:          req.ShopName,
		Note:              req.Note,
	}

	created, err := s.figureRepo.Create(ctx, figure)
	if err != nil {
		return nil, err
	}

	if s.dashboardService != nil {
		_ = s.dashboardService.ClearSummaryCache(ctx, userID)
	}

	return toFigureResponse(created), nil
}

func (s *FigureService) List(ctx context.Context, userID string) ([]model.FigureResponse, error) {
	figures, err := s.figureRepo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	responses := make([]model.FigureResponse, 0, len(figures))

	for _, figure := range figures {
		responses = append(responses, *toFigureResponse(&figure))
	}

	return responses, nil
}

func (s *FigureService) GetByID(ctx context.Context, userID string, figureID string) (*model.FigureResponse, error) {
	figure, err := s.figureRepo.FindByIDAndUserID(ctx, figureID, userID)
	if err != nil {
		return nil, err
	}

	if figure == nil {
		return nil, errors.New("figure not found")
	}

	return toFigureResponse(figure), nil
}

func (s *FigureService) Update(ctx context.Context, userID string, figureID string, req model.UpdateFigureRequest) (*model.FigureResponse, error) {
	if !isValidFigureStatus(req.Status) {
		return nil, errors.New("invalid figure status")
	}

	figure := &model.Figure{
		ID:                figureID,
		UserID:            userID,
		Name:              req.Name,
		CharacterName:     req.CharacterName,
		SeriesName:        req.SeriesName,
		Manufacturer:      req.Manufacturer,
		FigureType:         req.FigureType,
		Scale:             req.Scale,
		Status:            req.Status,
		Price:             req.Price,
		Deposit:           req.Deposit,
		Balance:           req.Balance,
		PreorderStartDate: parseDatePtr(req.PreorderStartDate),
		PreorderDeadline:  parseDatePtr(req.PreorderDeadline),
		ReleaseDate:       parseDatePtr(req.ReleaseDate),
		PaymentDueDate:    parseDatePtr(req.PaymentDueDate),
		ArrivalDate:       parseDatePtr(req.ArrivalDate),
		ShopName:          req.ShopName,
		Note:              req.Note,
	}

	updated, err := s.figureRepo.Update(ctx, figure)
	if err != nil {
		return nil, err
	}

	if updated == nil {
		return nil, errors.New("figure not found")
	}

	if s.dashboardService != nil {
		_ = s.dashboardService.ClearSummaryCache(ctx, userID)
	}

	return toFigureResponse(updated), nil
}

func (s *FigureService) Delete(ctx context.Context, userID string, figureID string) error {
	deleted, err := s.figureRepo.Delete(ctx, figureID, userID)
	if err != nil {
		return err
	}

	if !deleted {
		return errors.New("figure not found")
	}

	if s.dashboardService != nil {
		_ = s.dashboardService.ClearSummaryCache(ctx, userID)
	}

	return nil
}

func (s *FigureService) ListUpcomingPayments(ctx context.Context, userID string, days int) ([]model.FigureResponse, error) {
	if days <= 0 {
		days = 30
	}

	if days > 365 {
		days = 365
	}

	figures, err := s.figureRepo.ListUpcomingPayments(ctx, userID, days)
	if err != nil {
		return nil, err
	}

	responses := make([]model.FigureResponse, 0, len(figures))

	for _, figure := range figures {
		responses = append(responses, *toFigureResponse(&figure))
	}

	return responses, nil
}

func (s *FigureService) ListUpcomingReleases(ctx context.Context, userID string, days int) ([]model.FigureResponse, error) {
	if days <= 0 {
		days = 60
	}

	if days > 365 {
		days = 365
	}

	figures, err := s.figureRepo.ListUpcomingReleases(ctx, userID, days)
	if err != nil {
		return nil, err
	}

	responses := make([]model.FigureResponse, 0, len(figures))

	for _, figure := range figures {
		responses = append(responses, *toFigureResponse(&figure))
	}

	return responses, nil
}

func (s *FigureService) GetDashboardSummary(ctx context.Context, userID string) (map[string]interface{}, error) {
	return s.figureRepo.GetDashboardSummary(ctx, userID)
}

func parseDatePtr(value *string) *time.Time {
	if value == nil || *value == "" {
		return nil
	}

	parsed, err := time.Parse("2006-01-02", *value)
	if err != nil {
		return nil
	}

	return &parsed
}

func isValidFigureStatus(status string) bool {
	validStatuses := map[string]bool{
		"wishlist":     true,
		"preordered":   true,
		"deposit_paid": true,
		"balance_due":  true,
		"paid":         true,
		"shipped":      true,
		"arrived":      true,
		"cancelled":    true,
		"sold":         true,
	}

	return validStatuses[status]
}

func toFigureResponse(figure *model.Figure) *model.FigureResponse {
	return &model.FigureResponse{
		ID:                figure.ID,
		Name:              figure.Name,
		CharacterName:     figure.CharacterName,
		SeriesName:        figure.SeriesName,
		Manufacturer:      figure.Manufacturer,
		FigureType:         figure.FigureType,
		Scale:             figure.Scale,
		Status:            figure.Status,
		Price:             figure.Price,
		Deposit:           figure.Deposit,
		Balance:           figure.Balance,
		PreorderStartDate: figure.PreorderStartDate,
		PreorderDeadline:  figure.PreorderDeadline,
		ReleaseDate:       figure.ReleaseDate,
		PaymentDueDate:    figure.PaymentDueDate,
		ArrivalDate:       figure.ArrivalDate,
		ShopName:          figure.ShopName,
		Note:              figure.Note,
		CreatedAt:         figure.CreatedAt,
		UpdatedAt:         figure.UpdatedAt,
	}
}