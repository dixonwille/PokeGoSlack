package adapter

import "net/http"

//Adapter is used to apply certain attributes to a handler
type Adapter func(http.Handler) http.Handler

//Adapt adds all the adapters to the handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
