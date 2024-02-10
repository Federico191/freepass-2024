package handler

import "github.com/Federico191/freepass-2024/internal/usecase"

type Handler struct {
	User                 UserHandler
	ElectionPeriod       ElectionPeriodHandler
	CandidatePost        CandidatePostHandler
	CandidateInformation CandidateInformationHandler
	Candidate            CandidateHandler
	AccountVote          AccountVoteHandler
	Account              AccountHandler
	AccountComment       AccountCommentHandler
}

func Init(useCase *usecase.UseCase) *Handler {
	return &Handler{
		User:                 *NewUserHandler(useCase.User),
		ElectionPeriod:       *NewElectionPeriodHandler(useCase.ElectionPeriod),
		CandidatePost:        *NewCandidatePostHandler(useCase.CandidatePost),
		CandidateInformation: *NewCandidateInformationHandler(useCase.CandidateInformation),
		Candidate:            *NewCandidateHandler(useCase.Candidate),
		AccountVote:          *NewAccountVoteHandler(useCase.AccountVote),
		Account:              *NewAccountHandler(useCase.Account),
		AccountComment:       *NewAccountCommentHandler(useCase.AccountComment, useCase.Account),
	}
}
