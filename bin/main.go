package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

func main() {
    fmt.Println("hello world")
    // TODO register hooks
    y := readFile("example.yml")
    fmt.Println(y)
    mainHookPrototype(y.MainBlock)
}


type DoitConfig struct {
    MainBlock           map[string][]interface{} "__MAIN__"
    MetadataBlock       map[string][]interface{} "__METADATA__"
}


func readFile(filename string) DoitConfig {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("FileOpen Error")
    }
    defer file.Close()

    b, err := ioutil.ReadAll(file)

    var y DoitConfig

    err2 := yaml.Unmarshal(b, &y)
    if err2 != nil { ; }

    return y
}

func mainHookPrototype(b map[string][]interface{}) {
    //TODO this will be abstracted to hooks/
    fmt.Println("Running __MAIN__", b)
    for k, v := range b {
        fmt.Println("it is", k)
        for _, cmd := range v {
            fmt.Println("exec", cmd)
        }
    }
}
