package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/henriqueleite42/roles-e-jogos-backend/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetDbInCtx(db *pgxpool.Pool, ctx context.Context) (context.Context, error) {
	existentTx := ctx.Value("tx")
	if existentTx != nil {
		return ctx, nil
	}

	return context.WithValue(ctx, "tx", db), nil
}

func SetTxInCtx(db *pgxpool.Pool, ctx context.Context) (pgx.Tx, context.Context, error) {
	existentTx := ctx.Value("tx")
	if existentTx != nil {
		tx, ok := existentTx.(*pgxpool.Tx)
		if !ok {
			return nil, nil, errors.New("existent \"tx\" isn't a transaction")
		}
		return tx, ctx, nil
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, nil, errors.New("fail to create transaction")
	}
	ctx = context.WithValue(ctx, "tx", tx)

	return tx, ctx, nil
}

func GetDbFromCtx(ctx context.Context) (repository.Sql, error) {
	dbAny := ctx.Value("tx")
	if dbAny == nil {
		return nil, fmt.Errorf("missing context value: \"tx\"")
	}
	db, ok := dbAny.(repository.Sql)
	if !ok {
		return nil, fmt.Errorf("fail to parse \"dbAny\" to \"repository.Sql\"")
	}

	return db, nil
}
