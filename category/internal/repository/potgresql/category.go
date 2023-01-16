package potgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nkolosov/mentor-109/internal/entity"
	"go.uber.org/zap"
	"time"
)

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
	queryCtx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		`insert into category (id, name)VALUES ($1, $2)`, uuid.UUID(id).String(), name)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return nil, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id entity.CategoryId, name string) (*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		"UPDATE category SET name=$1 WHERE id=$2", name, uuid.UUID(id).String())
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return nil, nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id entity.CategoryId) error {
	queryCtx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	defer cancel()

	_, err := r.db.ExecContext(
		queryCtx,
		`DELETE FROM category WHERE id=$1`, uuid.UUID(id).String())
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	return nil
}

func (r *CategoryRepository) Filter(ctx context.Context, ids []entity.CategoryId) ([]*entity.Category, error) {
	queryCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
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
			r.logger.Error("failed to filter categories in", zap.Error(err))
			return nil, errors.New("failed to filter categories, see log for more info")
		}
		query = r.db.Rebind(query)
	}

	rows, err := r.db.QueryxContext(queryCtx, query, args...)
	if err != nil {
		r.logger.Error("failed to filter categories", zap.Error(err))
		return nil, errors.New("failed to filter categories, see log for more info")
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

		categories = append(categories, &entity.Category{
			Id:               category.Id,
			Name:             category.Name,
			CreateDate:       category.CreatedAt,
			ModificationDate: category.UpdatedAt,
			DeleteDate:       category.DeletedAt.Time,
		})
	}

	if rows.Err() != nil {
		r.logger.Error("db error on rows iteration", zap.Error(rows.Err()))
		return nil, errors.New("db error, see log for more info")
	}

	return categories, nil
}
