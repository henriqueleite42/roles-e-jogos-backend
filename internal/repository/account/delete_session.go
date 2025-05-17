package account_repository

import (
	"context"
)

func (self *accountRepositoryImplementation) DeleteSession(ctx context.Context, i *DeleteSessionInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	err = db.DeleteSession(ctx, i.SessionId)
	if err != nil {
		return err
	}

	return nil
}
