package schema

type (
	ChatSimpleResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	ChatDetailResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	ChatCreateRequest struct {
		Name string `json:"name"`
	}

	ChatCreateResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}
)

func NewChatSimpleResponse(id uint64, name string) *ChatSimpleResponse {
	return &ChatSimpleResponse{ID: id, Name: name}
}

func NewChatCreateResponse(id uint64, name string) *ChatCreateResponse {
	return &ChatCreateResponse{ID: id, Name: name}
}
