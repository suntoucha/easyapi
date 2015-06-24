package easyapi

const (
	errStatus = "error"
)

//По гайдлайну лучше назвать APIError (GoLINT)

//Ошибка, которую может отдать API
type ApiError struct {
	Status  string `json:"status"`
	Code    string `json:"error_code,omitempty"`
	Text    string `json:"error_text,omitempty"`
	Details string `json:"error_details,omitempty"`
}

//Для совместимости с стандартным error
func (ae *ApiError) Error() string {
	return ae.Code
}

//Создаем ошибку
func NewError(code string, text string, details string) (err *ApiError) {
	err = new(ApiError)

	err.Status = errStatus
	err.Code = code
	err.Text = text
	err.Details = details

	return err
}
