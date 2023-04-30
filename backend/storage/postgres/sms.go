package postgres

import (
	"context"
	"database/sql"

	"github.com/dilmurodov/xakaton_nt/pkg/models"
	"github.com/dilmurodov/xakaton_nt/storage"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type smsRepo struct {
	db *sql.DB
}

// NewSmsRepo ...
func NewSmsRepo(db *sql.DB) storage.SmsRepoI {
	return &smsRepo{db: db}
}

// MakeSent ...
func (cm *smsRepo) UpdateSMSAsSent(ctx context.Context, SMS_ID string) error {
	makesent := `UPDATE sms_send SET sent_at = CURRENT_TIMESTAMP where id = $1`
	_, err := cm.db.ExecContext(ctx, makesent, SMS_ID)

	return err
}

func (cm *smsRepo) IncrementSendCount(ctx context.Context, ID string) error {
	query := `UPDATE sms_send SET send_count = send_count + 1 where id = $1`
	_, err := cm.db.ExecContext(ctx, query, ID)
	return err
}

func (cm *smsRepo) SendOtp(ctx context.Context, req *models.Sms) (*models.SendOtpResponse, error) {
	var (
		resp = &models.SendOtpResponse{}
	)
	sendID, err := uuid.NewRandom()
	if err != nil {
		return resp, err
	}

	query := `INSERT INTO
		sms_send
		(
			id,
			text,
			phone_number,
			otp,
			expires_at
		)
		values($1, $2, $3, $4, now() + interval '10 minute')`

	_, err = cm.db.ExecContext(ctx, query, sendID, req.Text, req.Recipient, req.Otp)
	if err != nil {
		return resp, err
	}

	resp.SmsID = sendID.String()

	return resp, nil
}

func (cm *smsRepo) ConfirmOtp(ctx context.Context, req *models.ConfirmOtpRequest) (err error) {
	var (
		exist bool
	)

	query := `
		SELECT EXISTS(SELECT 1 FROM sms_send WHERE id = $1 AND otp = $2 AND expires_at >= CURRENT_TIMESTAMP)
	`

	err = cm.db.QueryRowContext(ctx, query, req.SmsID, req.Otp).Scan(&exist)
	if err != nil {
		return errors.Wrap(err, "Error while scanning exists")
	}

	if !exist {
		return errors.New("wrong otp or expired")
	}

	return
}

// GetNotSent ...
func (cm *smsRepo) GetNotSent(ctx context.Context) ([]*models.Sms, error) {
	var smss []*models.Sms

	query := `
		SELECT  
			id,
			phone_number,
			text,
			send_count,
			otp,
			expires_at
		FROM sms_send
		WHERE sent_at IS NULL and send_count < 4`
	rows, err := cm.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			sms       models.Sms
			expiresAt sql.NullString
		)
		if err := rows.Scan(
			&sms.Id,
			&sms.PhoneNumber,
			&sms.Text,
			&sms.SendCount,
			&sms.Otp,
			&expiresAt,
		); err != nil {
			return nil, err
		}

		if expiresAt.Valid {
			sms.ExpiresAt = expiresAt.String
		}

		smss = append(smss, &sms)
	}
	return smss, err
}
