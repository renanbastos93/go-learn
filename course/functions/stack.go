package main

import (
    "runtime/debug"
)

func f3() {
    debug.PrintStack()
}

func f2() {
    f3()
}

func f1() {
    f2()
}


// Funcao pura e algo que nao mexe em nada externo

func main() {
    f1()
}