package transaction

import (
	"crowdfund-go/campaign"
	"crowdfund-go/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	User       user.User
	Campaign   campaign.Campaign
	PaymentURL string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
