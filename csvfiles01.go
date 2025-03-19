/* RZFeeser | Alta3 Research
   Reading CSV file to stdout - Read()  */

package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
)

func main() {

    // open "vendorData.csv"
    f, err := os.Open("vendorData.csv")

    // check to see if opening the file caused an error
    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(f)

    for {

        /* r.Read() will return for us a slice containing all of the seperated values
           for example the csv data...
           just,an,example
           ...would become...
           []string{'just','an','example'}   */
        record, err := r.Read()   // record is now a slice of a single line of CSV data

        // if the err contains a notice about End Of File, we need to break out of the infinite loop
        if err == io.EOF {
            break
        }

        // check to see if the err contains an error
        if err != nil {
            log.Fatal(err)
        }

        // _ (position), value (the actual value stored in the slice)
        for _, value := range record {
            fmt.Printf("%s\n", value)
        }
    }
}

