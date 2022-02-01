package services

import (
	"time"
	"xendit/models"
	"xendit/repository"

	"github.com/gofiber/fiber/v2"
)

type ICommentService interface {
	CreateComment(comment models.InputComment, param string) error
	GetAllCommentByOrg(param string) ([]models.Comment, error)
	DeleteAllByOrg(org string) error
}

type CommentService struct {
	repo    *repository.CommentRepository
	orgrepo *repository.OrganizationRepo
}

func NewCommentService(repo *repository.CommentRepository, orgrepo *repository.OrganizationRepo) *CommentService {
	return &CommentService{repo, orgrepo}
}

func (s *CommentService) CreateCommentService(comment models.InputComment, param string) error {

	org, _ := s.orgrepo.FindByName(param)

	if org.Id.String() == "" {
		return fiber.NewError(fiber.StatusNotFound, "Organization Data Not Found")
	}

	var data models.Comment

	data.Message = comment.Message
	data.Organization = org
	data.CreatedAt = time.Now().UTC()
	data.UpdatedAt = time.Now().UTC()

	if err := s.repo.Save(data); err != nil {
		return err
	}

	return nil

}

func (s *CommentService) GetAllCommentByOrg(param string) ([]models.Comment, error) {

	org, _ := s.orgrepo.FindByName(param)

	if org.Id.String() == "" {
		return nil, fiber.NewError(fiber.StatusNotFound, "Organization Data Not Found")
	}

	res, err := s.repo.FindAll(org)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *CommentService) DeleteAllByOrg(org string) error {

	res, _ := s.orgrepo.FindByName(org)

	if res.Id.String() == "" {
		return fiber.ErrNotFound
	}

	comment, err := s.repo.FindAll(res)

	if err != nil {
		return err
	}

	var dump []interface{}

	for _, v := range comment {
		dump = append(dump, v)
	}

	if err := s.repo.SaveDump(dump); err != nil {
		return err
	}

	if err := s.repo.DeleteAll(res); err != nil {
		return err
	}

	return nil
}
