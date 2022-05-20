package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `jsin:"status"`
}

func APIresponse(message string, code int, status string, data interface{}) Response {
	meta_res := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response_json := Response{
		Meta: meta_res,
		Data: data,
	}

	return response_json
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
