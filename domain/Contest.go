package domain

type Contest struct {
	ContestID  uint64 `gorm:"primaryKey;not null" json:"contest_id"`
	CornerRed  string `gorm:"not null" json:"corner_red"`
	CornerBlue string `gorm:"not null" json:"corner_blue"`
}

type ContestRepository interface {
	GetContest(string) (*Contest, error)
	GetAllContest() ([]Contest, error)
	CreateContest(*Contest) error
}

type ContestUsecase interface {
	GetContest(string) (*Contest, error)
	GetAllContest() ([]Contest, error)
	Contest(string, string) (map[string]interface{}, error)
}
