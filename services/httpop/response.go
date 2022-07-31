package httpop

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Response struct {
	Status  string
	Data    interface{}
	Message string
}

func (this *Response) SendResponse(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(this)
}

func (this *Response) GenerateErrorResponse(data interface{}, msg string) {
	this.Status = "error"
	this.Data = data
	this.Message = msg
}

func (this *Response) GenerateErrorAuthResponse(data interface{}, msg string) {
	this.Status = "auth"
	this.Data = data
	this.Message = msg
}

func (this *Response) GenerateErrorAccessDeniedResponse(data interface{}, msg string) {
	this.Status = "ac"
	this.Data = data
	this.Message = msg
}

func (this *Response) GenerateOkResponse(data interface{}, msg string) {
	this.Status = "ok"
	this.Data = data
	this.Message = msg
}

func (this *Response) ValidateError(result *gorm.DB, w http.ResponseWriter, msg string) bool {
	if result.Error != nil || result.RowsAffected < 1 {
		w.WriteHeader(http.StatusBadRequest)
		this.GenerateErrorResponse(nil, msg)
		return false
	}
	return true
}
