package entity

type AccountVote struct {
	AccountId           uint    `gorm:"primaryKey"`
	Account             Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	IsVoted             *bool   `gorm:"default:false"`
	VotedElectionNumber *uint
	Candidate           []Candidate `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
