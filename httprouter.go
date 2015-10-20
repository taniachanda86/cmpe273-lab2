
package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
)

type Request struct{
    Name string `json:"name"`
}

type Response struct{
    Greeting   string   `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func helloPost(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {  

    var request Request
    var response Response

    json.NewDecoder(req.Body).Decode(&request)
    response.Greeting = "Hello " + request.Name
    responsejson, _ := json.Marshal(response)

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", responsejson)
}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello", helloPost)

    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}