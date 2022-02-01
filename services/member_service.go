package services

import (
	"time"
	"xendit/models"
	"xendit/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jaswdr/faker"
)

type IMemberService interface {
	CreateMember(member models.Member, param string) error
	FindAllMemberByOrg(param string) ([]models.Member, error)
}

type MemberService struct {
	repoMember *repository.MemberRepository
	repoOrg    *repository.OrganizationRepo
}

func NewMemberService(repoMember *repository.MemberRepository, repoOrg *repository.OrganizationRepo) *MemberService {
	return &MemberService{repoMember, repoOrg}
}

func (s *MemberService) CreateMember(member models.Member, param string) error {

	org, _ := s.repoOrg.FindByName(param)

	if org.Id.String() == "" {
		return fiber.ErrNotFound
	}

	fake := faker.New()

	member.Organization = org
	member.Avatar = "https://robohash.org/" + fake.Hash().MD5()
	member.Followers = fake.Int64()
	member.Following = fake.Int64()
	member.CreatedAt = time.Now().UTC()
	member.UpdatedAt = time.Now().UTC()

	if err := s.repoMember.Save(member); err != nil {
		return err
	}

	return nil
}

func (s *MemberService) FindAllMemberByOrg(param string) ([]models.Member, error) {

	org, _ := s.repoOrg.FindByName(param)

	if org.Id.String() == "" {
		return nil, fiber.ErrNotImplemented
	}

	res, err := s.repoMember.FindAll(org)

	if err != nil {
		return nil, err
	}

	return res, nil

}
