package response

const (
	GENERIC_MESSAGE           string = "Something went wrong"
	INVALID_REQUEST           string = "Invalid request body"
	COUNTRY_NOT_FOUND         string = "Country not found"
	FOOTBALL_TEAM_NOT_FOUND   string = "Football Team not found"
	FOOTBALL_LEAGUE_NOT_FOUND string = "Football League not found"
)

type _APIResponse struct {
	Result interface{}        `json:"result"`
	Error  *_APIErrorResponse `json:"error"`
}

type _APIErrorResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAPIResponse(result interface{}) _APIResponse {
	return _APIResponse{
		Result: result,
	}
}

func NewAPIErrorResponse(message string, data interface{}) _APIResponse {
	return _APIResponse{
		Result: nil,
		Error: &_APIErrorResponse{
			Message: message,
			Data:    data,
		},
	}
}
