package main

import(
    "log"
    "net/http"
    "encoding/json"
)

func main() {

    log.Println("Waiting for connections...")
    http.HandleFunc("/", mainController)

    log.Fatal(http.ListenAndServe(":8000", nil))
}

func mainController(res http.ResponseWriter, req *http.Request) {
    var message interface{};

    if err := json.NewDecoder(req.Body).Decode(&message); err != nil {
        log.Printf("Error decoding request body as json: %v", err)
    }

    log.Printf("Request decoded: %v", message)
}
