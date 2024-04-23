package domain

type success struct {
	Data interface{} `json:"data"`
}

type failed struct {
	Message interface{} `json:"message"`
}

func ResponseSuccess(data interface{}) success {
	return success{
		Data: data,
	}
}

func ResponseFailed(message interface{}) failed {
	return failed{
		Message: message,
	}
}
