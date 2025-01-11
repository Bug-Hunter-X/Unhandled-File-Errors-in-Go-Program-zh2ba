package main

import (


        "fmt"

        "os"

)

func main() {

        file, err := os.Open("myFile.txt")

        if err != nil {

                fmt.Println("Error opening file:", err)

                return

        }

        defer file.Close()

        //Process the file

        // ... your file processing logic ...

}