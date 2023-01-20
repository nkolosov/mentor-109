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

func (r *CategoryRepository) Create(ctx context.Context, id entity.CategoryId, name string) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		`insert into category (id, name)VALUES ($1, $2)`, uuid.UUID(id).String(), name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create category")
	}

	return nil, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id entity.CategoryId, name string) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		"UPDATE category SET name=$1 WHERE id=$2", name, uuid.UUID(id).String())
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
		var category struct {
			Id        string       `db:"id"`
			Name      string       `db:"name"`
			CreatedAt time.Time    `db:"created_at"`
			DeletedAt sql.NullTime `db:"deleted_at"`
			UpdatedAt time.Time    `db:"updated_at"`
		}
		if err = rows.StructScan(&category); err != nil {
			r.logger.Error("can't scan category to struct from db", zap.Error(err))
			continue
		}

		id, err := uuid.Parse(category.Id)
		if err != nil {
			r.logger.Error("invalid id", zap.Error(err))
			continue
		}

		categories = append(categories, &entity.Category{
			Id:               entity.CategoryId(id),
			Name:             category.Name,
			CreateDate:       category.CreatedAt,
			ModificationDate: category.UpdatedAt,
			DeleteDate:       category.DeletedAt.Time,
		})
	}

	if rows.Err() != nil {
		return nil, errors.Wrapf(rows.Err(), "db error on rows iteration")
	}

	return categories, nil
}
