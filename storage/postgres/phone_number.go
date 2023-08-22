package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/user/api/models"
	"github.com/user/pkg/helper"
)

type phoneNumberRepo struct {
	db *pgxpool.Pool
}

func NewphoneNumberRepo(db *pgxpool.Pool) *phoneNumberRepo {
	return &phoneNumberRepo{
		db: db,
	}
}

func (r *phoneNumberRepo) Create(ctx context.Context, req *models.CreatePhoneNumber) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO phone_number(id, user_id, phone, is_fax, description)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Userid,
		req.Phone,
		req.Isfax,
		req.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *phoneNumberRepo) GetByID(ctx context.Context, req *models.PhoneNumberPrimaryKey) (*models.PhoneNumber, error) {

	var whereField = "id"
	var (
		query string

		id          sql.NullString
		user_id     sql.NullString
		phone       sql.NullString
		is_fax      bool
		description sql.NullString
	)

	query = `
		SELECT
			id,
			user_id,
			phone,
			is_fax,
			description
		FROM phone_number
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&user_id,
		&phone,
		&is_fax,
		&description,
	)

	if err != nil {
		return nil, err
	}

	return &models.PhoneNumber{
		Id:          id.String,
		UserId:      user_id.String,
		Phone:       phone.String,
		Isfax:       is_fax,
		Description: description.String,
	}, nil
}

func (r *phoneNumberRepo) GetList(ctx context.Context, req *models.PhoneNumberGetListRequest) (*models.PhoneNumberGetListResponse, error) {

	var (
		resp   = &models.PhoneNumberGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			user_id,
			phone,
			is_fax,
			description
		FROM phone_number
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND phone ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          sql.NullString
			user_id     sql.NullString
			phone       sql.NullString
			is_fax      bool
			description sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&user_id,
			&phone,
			&is_fax,
			&description,
		)

		if err != nil {
			return nil, err
		}

		resp.PhoneNumbers = append(resp.PhoneNumbers, &models.PhoneNumber{
			Id:          id.String,
			UserId:      user_id.String,
			Phone:       phone.String,
			Isfax:       is_fax,
			Description: description.String,
		})
	}

	return resp, nil
}

func (r *phoneNumberRepo) Update(ctx context.Context, req *models.UpdatePhoneNumber) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			phone_number
		SET
			user_id = :user_id,
			phone = :phone,
			is_fax = :is_fax,
			description = :description
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"user_id":     req.Userid,
		"phone":       req.Phone,
		"is_fax":      req.Isfax,
		"description": req.Description,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *phoneNumberRepo) Delete(ctx context.Context, req *models.PhoneNumberPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM phone_number WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
