package helper

import (
	"net/http"

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
		newError := exception.NewInternalErr(100, err.Error())
		msg := model.NewErrorMessage(newError.Error())
		newError.LogError()
		Write(writer, http.StatusInternalServerError, msg)
		return
	}
	writer.WriteHeader(status)
	writer.Write(cont)
}
