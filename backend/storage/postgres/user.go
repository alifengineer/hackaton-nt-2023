package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/customerrors"
	"github.com/dilmurodov/xakaton_nt/pkg/models"
	_ "github.com/lib/pq"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db: db}
}

func (u *userRepo) GetUserByID(ctx context.Context, req *models.GetUserByIDRequest) (resp *models.GetUserByIDResponse, err error) {

	resp = &models.GetUserByIDResponse{}

	params := []interface{}{}

	query := `
		SELECT
			guid,
			first_name, 
			last_name, 
			email, 
			created_at, 
			updated_at
		FROM "users"
		`

	filter := `WHERE guid = $1 AND deleted_at = 0`
	params = append(params, req.UserId)

	if req.Email != "" {
		filter += " AND email = $2"
		params = append(params, req.Email)
	}

	query += filter

	row := u.db.QueryRowContext(ctx, query, params...)

	user := &models.User{}
	err = row.Scan(
		&user.Guid,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, &customerrors.UserNotFoundError{Guid: req.UserId}
	} else if err != nil {
		return nil, &customerrors.InternalServerError{Message: err.Error()}
	}

	resp.User = user

	return resp, nil
}

// Get User By Email

func (u *userRepo) GetUserPasswordByEmail(ctx context.Context, email string) (resp *models.User, err error) {

	query := `
		SELECT
			guid,
			first_name,
			last_name,
			email,
			password,
			created_at,
			updated_at
		FROM "users"
		WHERE email = $1 AND deleted_at = 0
	`

	row := u.db.QueryRowContext(ctx, query, email)
	resp = &models.User{}
	err = row.Scan(
		&resp.Guid,
		&resp.FirstName,
		&resp.LastName,
		&resp.Email,
		&resp.Password,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	if err != nil && err == sql.ErrNoRows {
		return nil, &customerrors.UserNotFoundWithEmailError{Emal: email}
	} else if err != nil {
		return nil, &customerrors.InternalServerError{Message: err.Error()}
	}

	return resp, nil
}

func (u *userRepo) CreateUser(ctx context.Context, req *models.CreateUserRequest) (resp *models.User, err error) {
	resp = &models.User{}

	query := `
		INSERT INTO "users" (
			first_name,
			last_name,
			email,
			password
		) VALUES (
			$1, $2, $3, $4
		) RETURNING guid, first_name, last_name, email, created_at, updated_at
	`

	row := u.db.QueryRowContext(ctx, query,
		req.User.FirstName,
		req.User.LastName,
		req.User.Email,
		req.User.Password,
	)

	err = row.Scan(
		&resp.Guid,
		&resp.FirstName,
		&resp.LastName,
		&resp.Email,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		return nil, &customerrors.InternalServerError{Message: err.Error()}
	}

	return resp, nil
}
