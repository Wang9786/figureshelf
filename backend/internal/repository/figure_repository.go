package repository

import (
	"context"
	"database/sql"

	"figureshelf-backend/internal/model"
)

type FigureRepository struct {
	db *sql.DB
}

func NewFigureRepository(db *sql.DB) *FigureRepository {
	return &FigureRepository{
		db: db,
	}
}

func (r *FigureRepository) Create(ctx context.Context, figure *model.Figure) (*model.Figure, error) {
	query := `
		INSERT INTO figures (
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
			$9, $10, $11, $12, $13, $14, $15, $16,
			$17, $18
		)
		RETURNING
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
	`

	var created model.Figure

	err := r.db.QueryRowContext(
		ctx,
		query,
		figure.UserID,
		figure.Name,
		figure.CharacterName,
		figure.SeriesName,
		figure.Manufacturer,
		figure.FigureType,
		figure.Scale,
		figure.Status,
		figure.Price,
		figure.Deposit,
		figure.Balance,
		figure.PreorderStartDate,
		figure.PreorderDeadline,
		figure.ReleaseDate,
		figure.PaymentDueDate,
		figure.ArrivalDate,
		figure.ShopName,
		figure.Note,
	).Scan(
		&created.ID,
		&created.UserID,
		&created.Name,
		&created.CharacterName,
		&created.SeriesName,
		&created.Manufacturer,
		&created.FigureType,
		&created.Scale,
		&created.Status,
		&created.Price,
		&created.Deposit,
		&created.Balance,
		&created.PreorderStartDate,
		&created.PreorderDeadline,
		&created.ReleaseDate,
		&created.PaymentDueDate,
		&created.ArrivalDate,
		&created.ShopName,
		&created.Note,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *FigureRepository) ListByUserID(ctx context.Context, userID string) ([]model.Figure, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
		FROM figures
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	figures := make([]model.Figure, 0)

	for rows.Next() {
		var figure model.Figure

		err := rows.Scan(
			&figure.ID,
			&figure.UserID,
			&figure.Name,
			&figure.CharacterName,
			&figure.SeriesName,
			&figure.Manufacturer,
			&figure.FigureType,
			&figure.Scale,
			&figure.Status,
			&figure.Price,
			&figure.Deposit,
			&figure.Balance,
			&figure.PreorderStartDate,
			&figure.PreorderDeadline,
			&figure.ReleaseDate,
			&figure.PaymentDueDate,
			&figure.ArrivalDate,
			&figure.ShopName,
			&figure.Note,
			&figure.CreatedAt,
			&figure.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		figures = append(figures, figure)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return figures, nil
}

func (r *FigureRepository) FindByIDAndUserID(ctx context.Context, id string, userID string) (*model.Figure, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
		FROM figures
		WHERE id = $1
		  AND user_id = $2
	`

	var figure model.Figure

	err := r.db.QueryRowContext(ctx, query, id, userID).Scan(
		&figure.ID,
		&figure.UserID,
		&figure.Name,
		&figure.CharacterName,
		&figure.SeriesName,
		&figure.Manufacturer,
		&figure.FigureType,
		&figure.Scale,
		&figure.Status,
		&figure.Price,
		&figure.Deposit,
		&figure.Balance,
		&figure.PreorderStartDate,
		&figure.PreorderDeadline,
		&figure.ReleaseDate,
		&figure.PaymentDueDate,
		&figure.ArrivalDate,
		&figure.ShopName,
		&figure.Note,
		&figure.CreatedAt,
		&figure.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &figure, nil
}

func (r *FigureRepository) Update(ctx context.Context, figure *model.Figure) (*model.Figure, error) {
	query := `
		UPDATE figures
		SET
			name = $1,
			character_name = $2,
			series_name = $3,
			manufacturer = $4,
			figure_type = $5,
			scale = $6,
			status = $7,
			price = $8,
			deposit = $9,
			balance = $10,
			preorder_start_date = $11,
			preorder_deadline = $12,
			release_date = $13,
			payment_due_date = $14,
			arrival_date = $15,
			shop_name = $16,
			note = $17,
			updated_at = NOW()
		WHERE id = $18
		  AND user_id = $19
		RETURNING
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
	`

	var updated model.Figure

	err := r.db.QueryRowContext(
		ctx,
		query,
		figure.Name,
		figure.CharacterName,
		figure.SeriesName,
		figure.Manufacturer,
		figure.FigureType,
		figure.Scale,
		figure.Status,
		figure.Price,
		figure.Deposit,
		figure.Balance,
		figure.PreorderStartDate,
		figure.PreorderDeadline,
		figure.ReleaseDate,
		figure.PaymentDueDate,
		figure.ArrivalDate,
		figure.ShopName,
		figure.Note,
		figure.ID,
		figure.UserID,
	).Scan(
		&updated.ID,
		&updated.UserID,
		&updated.Name,
		&updated.CharacterName,
		&updated.SeriesName,
		&updated.Manufacturer,
		&updated.FigureType,
		&updated.Scale,
		&updated.Status,
		&updated.Price,
		&updated.Deposit,
		&updated.Balance,
		&updated.PreorderStartDate,
		&updated.PreorderDeadline,
		&updated.ReleaseDate,
		&updated.PaymentDueDate,
		&updated.ArrivalDate,
		&updated.ShopName,
		&updated.Note,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &updated, nil
}

func (r *FigureRepository) Delete(ctx context.Context, id string, userID string) (bool, error) {
	query := `
		DELETE FROM figures
		WHERE id = $1
		  AND user_id = $2
	`

	result, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *FigureRepository) ListUpcomingPayments(ctx context.Context, userID string, days int) ([]model.Figure, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
		FROM figures
		WHERE user_id = $1
		  AND payment_due_date IS NOT NULL
		  AND payment_due_date BETWEEN CURRENT_DATE AND CURRENT_DATE + ($2 * INTERVAL '1 day')
		  AND status NOT IN ('arrived', 'cancelled', 'sold')
		ORDER BY payment_due_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query, userID, days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	figures := make([]model.Figure, 0)

	for rows.Next() {
		var figure model.Figure

		err := rows.Scan(
			&figure.ID,
			&figure.UserID,
			&figure.Name,
			&figure.CharacterName,
			&figure.SeriesName,
			&figure.Manufacturer,
			&figure.FigureType,
			&figure.Scale,
			&figure.Status,
			&figure.Price,
			&figure.Deposit,
			&figure.Balance,
			&figure.PreorderStartDate,
			&figure.PreorderDeadline,
			&figure.ReleaseDate,
			&figure.PaymentDueDate,
			&figure.ArrivalDate,
			&figure.ShopName,
			&figure.Note,
			&figure.CreatedAt,
			&figure.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		figures = append(figures, figure)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return figures, nil
}

func (r *FigureRepository) ListUpcomingReleases(ctx context.Context, userID string, days int) ([]model.Figure, error) {
	query := `
		SELECT
			id,
			user_id,
			name,
			character_name,
			series_name,
			manufacturer,
			figure_type,
			scale,
			status,
			price,
			deposit,
			balance,
			preorder_start_date,
			preorder_deadline,
			release_date,
			payment_due_date,
			arrival_date,
			shop_name,
			note,
			created_at,
			updated_at
		FROM figures
		WHERE user_id = $1
		  AND release_date IS NOT NULL
		  AND release_date BETWEEN CURRENT_DATE AND CURRENT_DATE + ($2 * INTERVAL '1 day')
		  AND status NOT IN ('arrived', 'cancelled', 'sold')
		ORDER BY release_date ASC
	`

	rows, err := r.db.QueryContext(ctx, query, userID, days)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	figures := make([]model.Figure, 0)

	for rows.Next() {
		var figure model.Figure

		err := rows.Scan(
			&figure.ID,
			&figure.UserID,
			&figure.Name,
			&figure.CharacterName,
			&figure.SeriesName,
			&figure.Manufacturer,
			&figure.FigureType,
			&figure.Scale,
			&figure.Status,
			&figure.Price,
			&figure.Deposit,
			&figure.Balance,
			&figure.PreorderStartDate,
			&figure.PreorderDeadline,
			&figure.ReleaseDate,
			&figure.PaymentDueDate,
			&figure.ArrivalDate,
			&figure.ShopName,
			&figure.Note,
			&figure.CreatedAt,
			&figure.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		figures = append(figures, figure)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return figures, nil
}

func (r *FigureRepository) GetDashboardSummary(ctx context.Context, userID string) (map[string]interface{}, error) {
	query := `
		SELECT
			COUNT(*) AS total_figures,
			COUNT(*) FILTER (WHERE status = 'wishlist') AS wishlist_count,
			COUNT(*) FILTER (WHERE status IN ('preordered', 'deposit_paid', 'balance_due', 'paid', 'shipped')) AS preordered_count,
			COUNT(*) FILTER (WHERE status = 'arrived') AS arrived_count,
			COALESCE(SUM(price), 0) AS total_price,
			COALESCE(SUM(deposit), 0) AS total_deposit,
			COALESCE(SUM(balance), 0) AS total_balance,
			COUNT(*) FILTER (
				WHERE payment_due_date IS NOT NULL
				  AND payment_due_date BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '30 days'
				  AND status NOT IN ('arrived', 'cancelled', 'sold')
			) AS upcoming_payments_count,
			COUNT(*) FILTER (
				WHERE release_date IS NOT NULL
				  AND release_date BETWEEN CURRENT_DATE AND CURRENT_DATE + INTERVAL '60 days'
				  AND status NOT IN ('arrived', 'cancelled', 'sold')
			) AS upcoming_releases_count
		FROM figures
		WHERE user_id = $1
	`

	var totalFigures int
	var wishlistCount int
	var preorderedCount int
	var arrivedCount int
	var totalPrice float64
	var totalDeposit float64
	var totalBalance float64
	var upcomingPaymentsCount int
	var upcomingReleasesCount int

	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&totalFigures,
		&wishlistCount,
		&preorderedCount,
		&arrivedCount,
		&totalPrice,
		&totalDeposit,
		&totalBalance,
		&upcomingPaymentsCount,
		&upcomingReleasesCount,
	)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"total_figures":             totalFigures,
		"wishlist_count":            wishlistCount,
		"preordered_count":          preorderedCount,
		"arrived_count":             arrivedCount,
		"total_price":               totalPrice,
		"total_deposit":             totalDeposit,
		"total_balance":             totalBalance,
		"upcoming_payments_count":   upcomingPaymentsCount,
		"upcoming_releases_count":   upcomingReleasesCount,
	}, nil
}