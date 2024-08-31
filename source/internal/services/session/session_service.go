package session_service

import (
	session_dto "github.com/ManuelTello/veterinary/internal/dto/session"
	audit_model "github.com/ManuelTello/veterinary/internal/models/audit"
	session_model "github.com/ManuelTello/veterinary/internal/models/session"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	sessionModel session_model.SessionModel
	auditModel   audit_model.AuditModel
}

func (service SessionService) ProcessSignUp(dto session_dto.IncomingSignUp) error {
	user, searchErr := service.sessionModel.SearchAccountByUsername(dto.Username)
	if searchErr != nil {
		return searchErr
	}

	if user != nil {

	} else {
		passwordHashed, hashErr := bcrypt.GenerateFromPassword([]byte(dto.Password), 7)
		if hashErr != nil {
			return hashErr
		}

		insertErr := service.sessionModel.InsertNewUser(dto.Username, string(passwordHashed), dto.FirstName, dto.LastName, dto.Email, dto.DateCreated, dto.PhoneNumber)
		if insertErr != nil {
			return insertErr
		}
	}

	return nil
}

func (service SessionService) ProcessSignIn(dto session_dto.IncomingSignIn) error {
	return nil
}

func New(session session_model.SessionModel, audit audit_model.AuditModel) SessionService {
	service := SessionService{
		sessionModel: session,
		auditModel:   audit,
	}

	return service
}
