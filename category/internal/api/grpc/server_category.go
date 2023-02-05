package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nkolosov/mentor-109/internal/entity"
	categoryv1 "github.com/nkolosov/mentor-109/pkg/api/grpc/gen/auction/category/category/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type CategoryRepository interface {
	Create(ctx context.Context, id entity.CategoryId, name string) (*entity.Category, error)
	Update(ctx context.Context, id entity.CategoryId, name string) (*entity.Category, error)
	Delete(ctx context.Context, id entity.CategoryId) error
	Filter(ctx context.Context, ids []entity.CategoryId) ([]*entity.Category, error)
}

type CategoryServer struct {
	repository     CategoryRepository
	protobufMapper *ToProtobufMapper
}

func NewCategoryServer(
	repository CategoryRepository,
	protobufMapper *ToProtobufMapper,
) *CategoryServer {
	return &CategoryServer{
		repository:     repository,
		protobufMapper: protobufMapper,
	}
}

func (s *CategoryServer) Create(ctx context.Context, request *categoryv1.CreateRequest) (*categoryv1.CreateResponse, error) {
	category, err := s.repository.Create(ctx, entity.CategoryId(uuid.New()), request.GetName())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to create category: %s", err))
	}
	return &categoryv1.CreateResponse{
		Category: s.protobufMapper.mapCategory(category),
	}, nil
}

func (s *CategoryServer) Update(ctx context.Context, request *categoryv1.UpdateRequest) (*categoryv1.UpdateResponse, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid id: %s: %s", request.GetId(), err))
	}

	category, err := s.repository.Update(ctx, entity.CategoryId(id), request.GetName())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update category: %s", err))
	}
	return &categoryv1.UpdateResponse{
		Category: s.protobufMapper.mapCategory(category),
	}, nil
}

func (s *CategoryServer) Delete(ctx context.Context, request *categoryv1.DeleteRequest) (*categoryv1.DeleteResponse, error) {
	id, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid id: %s: %s", request.GetId(), err))
	}

	err = s.repository.Delete(ctx, entity.CategoryId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to delete category: %s", err))
	}
	return &categoryv1.DeleteResponse{}, nil
}

func (s *CategoryServer) Filter(ctx context.Context, request *categoryv1.FilterRequest) (*categoryv1.FilterResponse, error) {
	ids := make([]entity.CategoryId, 0, len(request.GetIds()))
	for _, id := range request.GetIds() {
		id, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid id: %s: %s", strings.Join(request.GetIds(), ","), err))
		}
		ids = append(ids, entity.CategoryId(id))
	}

	categories, err := s.repository.Filter(ctx, ids)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to filter categories: %s", err))
	}

	return &categoryv1.FilterResponse{
		Categories: s.protobufMapper.mapCategories(categories),
	}, nil
}
