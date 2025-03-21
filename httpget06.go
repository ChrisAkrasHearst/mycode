/* Alta3 Research | RZFeeser
   HTTP GET requests  */
   
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    //"encoding/json"
)

func main() {

    // create a GET request to the page
    resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/pikachu")

    // check for errors
    if err != nil {
        log.Fatal(err)
    }

    // the client must close the response body when finished
    defer resp.Body.Close()

    // read the content of the body with "ReadAll()"
    body, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    // print received data to the console
    fmt.Println(string(body))

    //out, err := json.Marchal(body)

    //if err != nil {
        //panic(err)
    //}

    //os.Stdout.Write(out)
}

