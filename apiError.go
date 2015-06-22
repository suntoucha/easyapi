package easyapi



type ApiError struct {
	Status string `json:"status"`
	Code string `json:"error_code,omitempty"`
	Text string `json:"error_text,omitempty"`
	Details string `json:"error_details,omitempty"`
}



func NewError(code string, text string, details string) (*ApiError) {
	var (
		e ApiError
		)

	e.Status = "error"
	e.Code = code
	e.Text = text
	e.Details = details

	return &e
}