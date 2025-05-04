package account_repository

import (
	"context"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (self *accountRepositoryImplementation) GetOtp(ctx context.Context, i *GetOtpInput) (*GetOtpOutput, error) {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return nil, err
	}

	row, err := db.GetOtp(ctx, queries.GetOtpParams{
		AccountID: int32(i.AccountId),
		Code:      i.Code,
		Purpose:   queries.OtpPurposeEnum(i.Purpose),
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("not found")
		}

		return nil, err
	}
	if !row.Valid {
		return nil, fmt.Errorf("not found")
	}

	return &GetOtpOutput{
		CreatedAt: row.Time,
	}, nil
}
