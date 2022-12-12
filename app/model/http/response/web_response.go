package response

type (
	Success struct {
		Data interface{} `json:"data"`
	}

	Failure struct {
		Errors interface{} `json:"errors"`
	}

	Message struct {
		Message string `json:"message"`
	}
)
