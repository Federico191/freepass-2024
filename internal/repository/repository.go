package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	User                 UserRepositoryInterface
	Account              AccountRepositoryInterface
	Post                 PostRepositoryInterface
	Candidate            CandidateRepositoryInterface
	CandidatePost        CandidatePostRepositoryInterface
	CandidateInformation CandidateInformationRepositoryInterface
	AccountVote          AccountVoteRepositoryInterface
	AccountComment       AccountCommentRepositoryInterface
	ElectionPeriod       ElectionPeriodRepositoryInterface
}

func Init(db *gorm.DB, log *logrus.Logger) *Repository {
	return &Repository{
		User:                 NewUserRepository(db, log),
		Account:              NewAccountRepository(db, log),
		Post:                 NewPostRepository(db, log),
		Candidate:            NewCandidateRepository(db, log),
		CandidatePost:        NewCandidatePostRepository(db, log),
		CandidateInformation: NewCandidateInformationRepository(db, log),
		AccountVote:          NewAccountVoteRepository(db, log),
		AccountComment:       NewAccountCommentRepository(db, log),
		ElectionPeriod:       NewElectionPeriodRepository(db, log),
	}
}
