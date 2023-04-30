package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type emailRepo struct {
	db *sql.DB
}

func NewEmailRepo(db *sql.DB) storage.EmailRepoI {
	return &emailRepo{
		db: db,
	}
}

func (e *emailRepo) Create(ctx context.Context, input *models.Email) (*models.Email, error) {

	id := uuid.NewString()

	query := `
		INSERT INTO "email_sms" (
			id,
			email,
			otp,
			expires_at
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)`

	_, err := e.db.ExecContext(ctx, query,
		id,
		input.Email,
		input.Otp,
		input.ExpiresAt,
	)
	if err != nil {
		return nil, err
	}
	input.Id = id

	return input, nil
}

func (e *emailRepo) GetByPK(ctx context.Context, pKey *models.GetEmailOtpByPK) (res *models.Email, err error) {
	res = &models.Email{}
	query := `
		SELECT
			id,
			email,
			otp
		FROM
			"email_sms"
		WHERE
			id = $1 AND expires_at > NOW()`

	err = e.db.QueryRowContext(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.Email,
		&res.Otp,
	)
	if err == sql.ErrNoRows {
		err := errors.New("Otp has been expired")
		return nil, err
	} else if err != nil {
		return res, err
	}

	return res, nil
}
