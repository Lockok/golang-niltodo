package users_transport_http

import "github.com/Lockok/golang-niltodo/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"			 example:"10"`
	Version     int     `json:"version"		 example:"1"`
	FullName    string  `json:"full_name"	 example:"Vanek Vankov"`
	PhoneNumber *string `json:"phone_number" example"+79380012233"`
}

func userDtoFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDtoFromDomain(user)
	}

	return usersDTO
}