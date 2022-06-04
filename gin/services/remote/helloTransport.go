package remote

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewHelloTransport() http.Handler {

	endpoint := MakeHelloEndPoint(NewHelloSerivce())

	return httptransport.NewServer(endpoint, decodeHelloRequest, endcodeHelloResponse)
}

// type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)
func decodeHelloRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request HelloRequest

	request.Name = r.FormValue("name")
	//Name: r.URL.Query().Get("name"),
	return request, nil

}

// type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
func endcodeHelloResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {

	return json.NewEncoder(w).Encode(resp)
}
