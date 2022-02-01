package helpers

type ResponseApi struct {
	Message string      `json:"message"`
	Code    int32       `json:"code"`
	Data    interface{} `json:"data"`
}

func ApiResponse(msg string, code int32, data interface{}) *ResponseApi {

	res := &ResponseApi{
		Message: msg,
		Code:    code,
		Data:    data,
	}

	return res

}
