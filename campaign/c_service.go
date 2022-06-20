package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	// if params user id tidak sama dengan kosong maka tampilkan campaign yg dimiliki userid tersebut
	if userID != 0 {
		campaigns, err := s.r.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	// else tampilkan semua campaign
	campaigns, err := s.r.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil

}

//mendapatkan data campaign berdasarkan id
func (s *service) GetCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.r.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %v", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil

}
