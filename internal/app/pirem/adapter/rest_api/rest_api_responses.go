package adapter

type RestApiResponseStatus string

const (
	StatusError   = RestApiResponseStatus("error")
	StatusSuccess = RestApiResponseStatus("success")
)

type RestApiResponse struct {
	Status  RestApiResponseStatus `json:"status"`
	Content interface{}           `json:"content"`
}

func NewRestApiSuccessResponse(content interface{}) RestApiResponse {
	return RestApiResponse{
		Status:  StatusSuccess,
		Content: content,
	}
}

func NewRestApiErrorResponse(err error) RestApiResponse {
	return RestApiResponse{
		Status:  StatusError,
		Content: err,
	}
}
