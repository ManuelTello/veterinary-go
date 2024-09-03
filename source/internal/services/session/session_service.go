package session_service

import (
	session_dto "github.com/ManuelTello/veterinary/internal/dto/session"
	audit_model "github.com/ManuelTello/veterinary/internal/models/audit"
	role_model "github.com/ManuelTello/veterinary/internal/models/role"
	session_model "github.com/ManuelTello/veterinary/internal/models/session"
	bcrypt "golang.org/x/crypto/bcrypt"
)

type SessionService struct {
	sessionModel session_model.SessionModel
	auditModel   audit_model.AuditModel
	roleModel    role_model.RoleModel
}

func (service SessionService) ProcessSignUp(dto session_dto.IncomingSignUp) error {
	passwordHashed, hashErr := bcrypt.GenerateFromPassword([]byte(dto.Password), 7)
	if hashErr != nil {
		return hashErr
	}

	lastInserted, insertErr := service.sessionModel.InsertNewUser(string(passwordHashed), dto.FirstName, dto.LastName, dto.Email, dto.DateCreated, dto.PhoneNumber, dto.AlternativeNumber)
	if insertErr != nil {
		return insertErr
	}

	linkErr := service.roleModel.LinkAccountWithRole(lastInserted, 2)
	if linkErr != nil {
		return linkErr
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
