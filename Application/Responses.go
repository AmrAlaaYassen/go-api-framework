package Application

import (
	"github.com/bykovme/gotrans"
)

func (req Request) Ok(body interface{}) {
	req.Response(200, responseBuilder(body, gotrans.T("ok"), 200, nil))
}

func (req Request) Created(body interface{}) {
	req.Response(201, responseBuilder(body, gotrans.T("created"), 200, nil))
}

func (req Request) NotAuth() {

	req.Response(401, responseBuilder(nil, gotrans.T("unauthorized"), 401, nil))
}
func (req Request) BadRequest(err interface{}) {
	req.Response(422, responseBuilder(nil, gotrans.T("went_wrong"), 422, err))
}

func (req Request) UserNotFound() {
	req.Response(404, responseBuilder(nil, gotrans.T("not_found"), 404, nil))
}

func responseBuilder(payload interface{}, message string, code int, err interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["payload"] = payload
	response["message"] = message
	response["code"] = code
	response["errors"] = err
	return response
}
