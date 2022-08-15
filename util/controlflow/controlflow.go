package controlflow

import (
    "log"
)

func PanicOnErr(e error) {
    if e != nil {
        panic(e)
    }
}

func Fatal(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}

