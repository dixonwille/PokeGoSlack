package helper

import (
	"bytes"
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

//RespondingLater does a post to a specific url
func RespondingLater(url string, content interface{}) {
	obj, ok := content.(model.Publicer)
	if ok {
		content = obj.Public()
	}
	cont, err := model.Jsonify(content)
	if err != nil {
		newError := exception.NewInternalError(err.Error())
		RespondingLater(url, newError)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(cont))
	if err != nil {
		panic(exception.NewInternalError("Could not make post request: " + err.Error()))
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(exception.NewInternalError("Could not get a response from post request: " + err.Error()))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(exception.NewInternalError("Slack did not like request: " + err.Error()))
	}
}

func RespondingLaterError(url string, err error) {
	errMsg := new(model.Response)
	if e, ok := exception.IsException(err); ok {
		errMsg = model.NewErrorMessage(e.Error())
		e.LogError()
		RespondingLater(url, errMsg)
		return
	}
	errMsg = model.NewErrorMessage(err.Error())
	env.Logger.Println(err.Error())
	RespondingLater(url, errMsg)
}
