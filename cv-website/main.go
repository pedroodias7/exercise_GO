// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)


func check(e error)  {
    if e != nil {
        fmt.Println(e)
        os.Exit(1)

    }
}

func main()  {
    out, err := os.Create("index.html")
    defer out.Close()
    check(e)




    data := map[string]interface{}{}
    file, err := ioutil.ReadFile("data.json")
    check(err)


    err = json.Unmarshal(file, &data)
    check(err)

    t, err := template.ParseGlob("template/*")
    check(err)

    t.Execute(out, data)


}
