package usecase

import (
	"github.com/Federico191/freepass-2024/internal/entity/token"
	"github.com/Federico191/freepass-2024/internal/repository"
	"time"
)

type UseCase struct {
	User                 UserUseCaseInterface
	Post                 PostUseCaseInterface
	ElectionPeriod       ElectionPeriodUseCaseInterface
	Candidate            CandidateUseCaseInterface
	CandidatePost        CandidatePostUseCaseInterface
	CandidateInformation CandidateInformationInterface
	AccountVote          AccountVoteInterface
	Account              AccountUseCaseInterface
	AccountComment       AccountCommentUseCaseInterface
}

func Init(repo *repository.Repository, timeOut time.Duration, tokenMaker token.Maker) *UseCase {
	return &UseCase{
		User:                 NewUserUseCase(repo.User, timeOut, tokenMaker),
		Post:                 NewPostUseCase(repo.Post, timeOut),
		ElectionPeriod:       NewElectionPeriodUseCase(repo.ElectionPeriod, timeOut),
		Candidate:            NewCandidateUseCase(repo.Candidate, timeOut),
		CandidatePost:        NewCandidatePostUseCase(repo.CandidatePost, timeOut),
		CandidateInformation: NewCandidateInformationUseCase(repo.CandidateInformation, timeOut),
		AccountVote:          NewAccountVoteUseCase(repo.AccountVote, timeOut),
		Account:              NewAccountUseCase(repo.Account, timeOut),
		AccountComment:       NewAccountCommentUseCase(repo.AccountComment, timeOut),
	}
}
