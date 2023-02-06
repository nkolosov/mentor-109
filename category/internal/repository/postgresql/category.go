package postgresql

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nkolosov/mentor-109/internal/entity"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

const timeout = 250 * time.Millisecond

type categoryDb struct {
	Id        string       `db:"id"`
	Name      string       `db:"name"`
	CreatedAt time.Time    `db:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

type CategoryRepository struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewCategoryRepository(
	db *sqlx.DB,
	logger *zap.Logger,
) *CategoryRepository {
	return &CategoryRepository{
		db:     db,
		logger: logger,
	}
}

func (r *CategoryRepository) Create(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		`insert into category (id, name,created_at,updated_at)VALUES ($1, $2, $3, $4)`,
		uuid.UUID(category.Id()).String(), category.Name(), category.CreateDate(), category.ModificationDate())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create category")
	}

	return nil, nil
}

func (r *CategoryRepository) Update(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		"UPDATE category SET name=$1 WHERE id=$2", category.Name(), uuid.UUID(category.Id()).String())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to update category")
	}

	return nil, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id entity.CategoryId) error {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		`DELETE FROM category WHERE id=$1`, uuid.UUID(id).String())
	if err != nil {
		return errors.Wrapf(err, "failed to delete category")
	}

	return nil
}

func (r *CategoryRepository) Filter(ctx context.Context, ids []entity.CategoryId) ([]*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	categoryIds := make([]string, 0, len(ids))
	for _, id := range ids {
		categoryIds = append(categoryIds, uuid.UUID(id).String())
	}

	query := "select * from category"
	var args []interface{}
	var err error
	if len(ids) != 0 {
		query, args, err = sqlx.In("select * from category where id IN (?);", categoryIds)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to filter categories in")
		}
		query = r.db.Rebind(query)
	}

	rows, err := r.db.QueryxContext(queryCtx, query, args...)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to filter categorie")
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			r.logger.Error("failed to close rows db", zap.Error(err))
		}
	}()

	categories := make([]*entity.Category, 0)
	for rows.Next() {
		var categoryDb categoryDb
		if err = rows.StructScan(&categoryDb); err != nil {
			r.logger.Error("can't scan category to struct from db", zap.Error(err))
			continue
		}

		category, err := r.toDomain(categoryDb)
		if err = rows.StructScan(&categoryDb); err != nil {
			return nil, errors.Wrapf(err, "failed convert to domain")
		}

		categories = append(categories, category)
	}

	if rows.Err() != nil {
		return nil, errors.Wrapf(rows.Err(), "db error on rows iteration")
	}

	return categories, nil
}

func (r *CategoryRepository) Get(ctx context.Context, id entity.CategoryId) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := r.db.QueryxContext(
		queryCtx,
		`select * from category WHERE id=$1`, uuid.UUID(id).String())
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get query")
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			r.logger.Error("failed to close rows db", zap.Error(err))
		}
	}()

	for rows.Next() {
		var categoryDb categoryDb
		if err = rows.StructScan(&categoryDb); err != nil {
			return nil, errors.Wrapf(err, "failed to get scan categorie")
		}

		category, err := r.toDomain(categoryDb)
		if err = rows.StructScan(&categoryDb); err != nil {
			return nil, errors.Wrapf(err, "failed convert to domain")
		}

		return category, nil
	}

	return nil, nil
}

func (r *CategoryRepository) toDomain(categoryDb categoryDb) (*entity.Category, error) {
	id, err := uuid.Parse(categoryDb.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get scan categorie")
	}

	return entity.NewCategory(
		entity.CategoryId(id),
		categoryDb.Name,
		categoryDb.CreatedAt,
		categoryDb.UpdatedAt,
		categoryDb.DeletedAt.Time,
	), nil
}
