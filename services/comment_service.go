package services

import (
	"reflect"
	"time"
	"xendit/models"
	"xendit/repository"

	"github.com/gofiber/fiber/v2"
)

type ICommentService interface {
	CreateComment(comment models.InputComment, param string) error
	GetAllCommentByOrg(param string) ([]models.Comment, error)
}

type CommentService struct {
	repo    *repository.CommentRepository
	orgrepo *repository.OrganizationRepo
}

func NewCommentService(repo *repository.CommentRepository, orgrepo *repository.OrganizationRepo) *CommentService {
	return &CommentService{repo, orgrepo}
}

func (s *CommentService) CreateCommentService(comment models.InputComment, param string) error {

	org, err := s.orgrepo.FindByName(param)

	if reflect.ValueOf(org).IsNil() {
		return fiber.NewError(fiber.StatusNotFound, "Organization Data Not Found")
	}

	if err != nil {
		return err
	}

	var data models.Comment

	data.Message = comment.Message
	data.Organization.Id = org.Id
	data.Organization.Name = org.Name
	data.Organization.CreatedAt = org.CreatedAt
	data.Organization.UpdatedAt = org.UpdatedAt
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	if err := s.repo.Save(data); err != nil {
		return err
	}

	return nil

}

func (s *CommentService) GetAllCommentByOrg(param string) ([]models.Comment, error) {

	org, err := s.orgrepo.FindByName(param)

	if reflect.ValueOf(org).IsNil() {
		return nil, fiber.NewError(fiber.StatusNotFound, "Organization Data Not Found")
	}

	if err != nil {
		return nil, err
	}

	res, err := s.repo.FindAll(org)

	if err != nil {
		return nil, err
	}

	return res, nil

}
