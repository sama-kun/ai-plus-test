package dto

type ErrorResponse struct {
	StatusCode int    `json:"statusCode" example:"400"`
	Error      string `json:"error" example:"invalid input"`
	Details    string `json:"details,omitempty" example:"name is required"`
}
