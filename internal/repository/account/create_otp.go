package account_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *accountRepositoryImplementation) CreateOtp(ctx context.Context, i *CreateOtpInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	err = db.CreateOtp(ctx, queries.CreateOtpParams{
		AccountID: int32(i.AccountId),
		Code:      i.Code,
		Purpose:   queries.OtpPurposeEnum(i.Purpose),
	})
	if err != nil {
		return err
	}

	return nil
}
