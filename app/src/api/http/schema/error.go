package schema

func NewError(err error, message string) map[string]any {
	return map[string]any{
		"message": message,
		"details": err.Error(),
	}
}
