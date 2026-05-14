package model

import "time"

type Figure struct {
	ID                string     `json:"id"`
	UserID            string     `json:"user_id"`
	Name              string     `json:"name"`
	CharacterName     *string    `json:"character_name"`
	SeriesName        *string    `json:"series_name"`
	Manufacturer      *string    `json:"manufacturer"`
	FigureType         *string    `json:"figure_type"`
	Scale             *string    `json:"scale"`
	Status            string     `json:"status"`
	Price             float64    `json:"price"`
	Deposit           float64    `json:"deposit"`
	Balance           float64    `json:"balance"`
	PreorderStartDate *time.Time `json:"preorder_start_date"`
	PreorderDeadline  *time.Time `json:"preorder_deadline"`
	ReleaseDate       *time.Time `json:"release_date"`
	PaymentDueDate    *time.Time `json:"payment_due_date"`
	ArrivalDate       *time.Time `json:"arrival_date"`
	ShopName          *string    `json:"shop_name"`
	Note              *string    `json:"note"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

type CreateFigureRequest struct {
	Name              string   `json:"name" binding:"required"`
	CharacterName     *string  `json:"character_name"`
	SeriesName        *string  `json:"series_name"`
	Manufacturer      *string  `json:"manufacturer"`
	FigureType         *string  `json:"figure_type"`
	Scale             *string  `json:"scale"`
	Status            string   `json:"status"`
	Price             float64  `json:"price"`
	Deposit           float64  `json:"deposit"`
	Balance           float64  `json:"balance"`
	PreorderStartDate *string  `json:"preorder_start_date"`
	PreorderDeadline  *string  `json:"preorder_deadline"`
	ReleaseDate       *string  `json:"release_date"`
	PaymentDueDate    *string  `json:"payment_due_date"`
	ArrivalDate       *string  `json:"arrival_date"`
	ShopName          *string  `json:"shop_name"`
	Note              *string  `json:"note"`
}

type UpdateFigureRequest struct {
	Name              string   `json:"name" binding:"required"`
	CharacterName     *string  `json:"character_name"`
	SeriesName        *string  `json:"series_name"`
	Manufacturer      *string  `json:"manufacturer"`
	FigureType         *string  `json:"figure_type"`
	Scale             *string  `json:"scale"`
	Status            string   `json:"status" binding:"required"`
	Price             float64  `json:"price"`
	Deposit           float64  `json:"deposit"`
	Balance           float64  `json:"balance"`
	PreorderStartDate *string  `json:"preorder_start_date"`
	PreorderDeadline  *string  `json:"preorder_deadline"`
	ReleaseDate       *string  `json:"release_date"`
	PaymentDueDate    *string  `json:"payment_due_date"`
	ArrivalDate       *string  `json:"arrival_date"`
	ShopName          *string  `json:"shop_name"`
	Note              *string  `json:"note"`
}

type FigureResponse struct {
	ID                string     `json:"id"`
	Name              string     `json:"name"`
	CharacterName     *string    `json:"character_name"`
	SeriesName        *string    `json:"series_name"`
	Manufacturer      *string    `json:"manufacturer"`
	FigureType         *string    `json:"figure_type"`
	Scale             *string    `json:"scale"`
	Status            string     `json:"status"`
	Price             float64    `json:"price"`
	Deposit           float64    `json:"deposit"`
	Balance           float64    `json:"balance"`
	PreorderStartDate *time.Time `json:"preorder_start_date"`
	PreorderDeadline  *time.Time `json:"preorder_deadline"`
	ReleaseDate       *time.Time `json:"release_date"`
	PaymentDueDate    *time.Time `json:"payment_due_date"`
	ArrivalDate       *time.Time `json:"arrival_date"`
	ShopName          *string    `json:"shop_name"`
	Note              *string    `json:"note"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}