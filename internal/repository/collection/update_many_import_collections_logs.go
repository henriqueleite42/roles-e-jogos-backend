package collection_repository

import (
	"context"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository/queries"
)

func (self *collectionRepositoryImplementation) UpdateManyImportCollectionsLogs(ctx context.Context, i *UpdateManyImportCollectionsLogsInput) error {
	db, err := self.getSlcQueries(ctx)
	if err != nil {
		return err
	}

	ids := make([]int32, len(i.Ids))
	for k, v := range i.Ids {
		ids[k] = int32(v)
	}

	err = db.UpdateManyImportCollectionsLogs(ctx, queries.UpdateManyImportCollectionsLogsParams{
		Column1: ids,
		Status:  queries.CollectionImportStatusEnum(i.Status),
	})
	if err != nil {
		return err
	}

	return nil
}
