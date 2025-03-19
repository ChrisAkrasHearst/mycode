/* RZFeeser | Alta3 Research
Exploring XML             */
package main
 
import (
    "encoding/xml"
    "fmt"
    "os"
)
 
type Notes struct {
    To      string `xml:"to"`
    From    string `xml:"from"`
    Subject string `xml:"subject"`
    Body    string `xml:"body"`
}
 
func main() {
    data, _ := os.ReadFile("avengers.xml")    // open the file "avengers.xml"
 
    // declade our note from the structure Notes
    // we will initalize with xml.Unmarshal
    var note Notes
 
    // this only returns an error, which we are throwing away with the blank identifier
    _ = xml.Unmarshal([]byte(data), &note)
 
    fmt.Println(note.To)                    // display value of "to"
    fmt.Println(note.From)                  // display value from "from"
    fmt.Println(note.Subject)               // display value from "subject"
    fmt.Println(note.Body)                  // display value from "body"
}

