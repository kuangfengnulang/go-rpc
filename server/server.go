package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
type Input struct{ Name string }
type Output struct{ Msg string }

func handle(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	serviceName := r.Header.Get("sparrow-service")
	methodName := r.Header.Get("sparrow-service-method")

	service :=
	input := &Input{}
	_ = json.Unmarshal(data, input)

	fmt.Printf("%#v\n", input)
	out, _ := json.Marshal(&Output{
		Msg: "hello " + input.Name,
	})
	fmt.Fprintf(w, "%s", string(out))

}
func main() {
	AddService(&helloService{})
	AddService(&userService{})

	http.HandleFunc("/", handle)
	http.ListenAndServe(":9000", nil)
}
