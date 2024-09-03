package helpers_service

import (
	session_model "github.com/ManuelTello/veterinary/internal/models/session"
)

type HelpersService struct {
	sessionModel session_model.SessionModel
}

func (service HelpersService) SearchIfEmailExists(email string) (bool, error) {
	emailExists, searchErr := service.sessionModel.DoesEmailExists(email)
	if searchErr != nil {
		return false, searchErr
	}

	if emailExists == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func New(session session_model.SessionModel) HelpersService {
	return HelpersService{
		sessionModel: session,
	}
}
