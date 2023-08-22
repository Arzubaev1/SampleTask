package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/cast"
	"github.com/user/api/models"
	"github.com/user/pkg/helper"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		id    = uuid.New().String()
		query string
	)
	query = `
		INSERT INTO users(id, login, password, name, age)
		VALUES ($1, $2, $3, $4, $5)
	`
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return "", err
	}
	_, err = r.db.Exec(ctx, query,
		id,
		req.Login,
		cast.ToString(bytes),
		req.Name,
		req.Age,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var whereField = "id"
	if len(req.Login) > 0 {
		whereField = "login"
		req.Id = req.Login
	}

	var (
		query string

		id        sql.NullString
		login     sql.NullString
		password  sql.NullString
		name      sql.NullString
		age       int
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query = `
		SELECT
			id,
			login,
			password,
			name,
			age,
			created_at,
			updated_at
		FROM users
		WHERE ` + whereField + ` = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&login,
		&password,
		&name,
		&age,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        id.String,
		Login:     login.String,
		Password:  password.String,
		Name:      name.String,
		Age:       age,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			login,
			password,
			name,
			age,
			created_at,
			updated_at
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id        sql.NullString
			login     sql.NullString
			password  sql.NullString
			name      sql.NullString
			age       int
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&login,
			&password,
			&name,
			&age,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &models.User{
			Id:        id.String,
			Login:     login.String,
			Password:  password.String,
			Name:      name.String,
			Age:       age,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			users
		SET
			login = :login,
			password = :password,
			name = :name,
			age = :age,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":       req.Id,
		"login":    req.Login,
		"password": req.Password,
		"name":     req.Name,
		"age":      req.Age,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) error {

	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
