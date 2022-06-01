package schema

import (
	"net/url"
)

const (
	UserOffline UserOnlineStatus = "offline"
)

type (
	UserOnlineStatus string

	UserSimpleResponse struct {
		ID           uint64           `json:"id"`
		Username     string           `json:"username"`
		ImageURL     *url.URL         `json:"imageURL"`
		OnlineStatus UserOnlineStatus `json:"onlineStatus"`
	}

	UserDetailResponse struct {
		ID       uint64   `json:"id"`
		Username string   `json:"username"`
		ImageURL *url.URL `json:"imageURL"`
	}

	UserCreateRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserCreateResponse struct {
		ID       uint64   `json:"id"`
		Username string   `json:"username"`
		ImageURL *url.URL `json:"imageURL"`
	}
)

func NewUserSimpleResponse(id uint64, username string, imageURL *url.URL, status UserOnlineStatus) *UserSimpleResponse {
	return &UserSimpleResponse{
		ID:           id,
		Username:     username,
		ImageURL:     imageURL,
		OnlineStatus: status,
	}
}

func NewUserCreateResponse(id uint64, username string, imageURL *url.URL) *UserCreateResponse {
	return &UserCreateResponse{id, username, imageURL}
}
