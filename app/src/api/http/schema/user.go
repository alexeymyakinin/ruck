package schema

import (
	"github.com/labstack/echo/v4"
	"net/url"
)

type (
	UserSimpleResponse struct {
		ID       uint64   `json:"id"`
		Username string   `json:"username"`
		ImageURL *url.URL `json:"imageURL"`
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
		User *UserDetailResponse `json:"user"`
	}
)

func NewUserSimpleResponse(id uint64, username string, imageURL *url.URL) *UserSimpleResponse {
	return &UserSimpleResponse{ID: id, Username: username, ImageURL: imageURL}
}

func NewUserDetailResponse(id uint64, username string, imageURL *url.URL) *UserDetailResponse {
	return &UserDetailResponse{ID: id, Username: username, ImageURL: imageURL}
}

func GetUserCreateRequest(c echo.Context) (*UserCreateRequest, error) {
	var req UserCreateRequest
	err := c.Bind(&req)

	if err != nil {
		return nil, err
	}
	return &req, nil
}
