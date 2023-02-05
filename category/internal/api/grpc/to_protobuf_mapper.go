package grpc

import (
	"github.com/google/uuid"
	"github.com/nkolosov/mentor-109/internal/entity"
	categoryv1 "github.com/nkolosov/mentor-109/pkg/api/grpc/gen/auction/category/category/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToDomainMapper мапер в grpc структуры
type ToProtobufMapper struct{}

// NewToDomainMapper создает новый мапер в grpc структуры
func NewToProtobufMapper() *ToProtobufMapper {
	return &ToProtobufMapper{}
}

func (m *ToProtobufMapper) mapCategory(category *entity.Category) *categoryv1.Category {
	if category == nil {
		return nil
	}

	return &categoryv1.Category{
		Id:         uuid.UUID(category.Id()).String(),
		Name:       category.Name(),
		CreateTime: timestamppb.New(category.CreateDate()),
		ModifyTime: timestamppb.New(category.ModificationDate()),
		DeleteTime: timestamppb.New(category.DeleteDate()),
	}
}

func (m *ToProtobufMapper) mapCategories(categories []*entity.Category) []*categoryv1.Category {
	categoriesPb := make([]*categoryv1.Category, 0, len(categories))

	for _, category := range categories {
		categoriesPb = append(categoriesPb, m.mapCategory(category))
	}

	return categoriesPb
}
