package helper

import (
	"net/http"

	"github.com/dixonwille/PokeGoSlack/env"
	"github.com/dixonwille/PokeGoSlack/exception"
	"github.com/dixonwille/PokeGoSlack/model"
)

//Write is used to write content to the writer with the status.
//Accepts a model and will turn it into json before writing.
func Write(writer http.ResponseWriter, status int, content interface{}) {
	obj, ok := content.(model.Publicer)
	if ok {
		content = obj.Public()
	}
	cont, err := model.Jsonify(content)
	if err != nil {
		newError := exception.NewInternalError(err.Error())
		WriteError(writer, newError)
		return
	}
	writer.WriteHeader(status)
	writer.Write(cont)
}

//WriteError is used to write errors back to slack
func WriteError(w http.ResponseWriter, err error) {
	errMsg := new(model.Response)
	if e, ok := exception.IsException(err); ok {
		errMsg = model.NewErrorMessage(e.Error())
		e.LogError()
		Write(w, e.Code, errMsg)
		return
	}
	errMsg = model.NewErrorMessage(err.Error())
	env.Logger.Println(err.Error())
	Write(w, http.StatusInternalServerError, errMsg)
	return
}
