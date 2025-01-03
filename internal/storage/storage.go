package storage

import "log"

var functionStore = make(map[string]Function)

type Function struct {
    ID   string
    Name string
}

func SaveFunction(fn Function) {
    functionStore[fn.ID] = fn
    log.Printf("Function saved: %s", fn.Name)
}

func GetFunction(id string) (Function, bool) {
    fn, exists := functionStore[id]
    return fn, exists
}
