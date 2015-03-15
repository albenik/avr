package main

import (
    "bytes"
    "fmt"
    "go/format"
    "log"
    "os"
    "strings"
)

type Kind int

const (
    Flat Kind = iota
)

type Generator struct {
    buf bytes.Buffer
}

func (g *Generator) Printf(format string, args ...interface{}) {
    fmt.Fprintf(&g.buf, format, args...)
}

func (g *Generator) Generate(pkgName string, kind Kind) {
    g.Printf("// automatically generated by avrmakedec %s\n", strings.Join(os.Args[1:], " "))
    g.Printf("// DO NOT EDIT\n")
    g.Printf("package %s\n", pkgName)
    g.Printf("import \"github.com/kierdavis/avr\"\n")
    g.Printf("func Decode(word uint16, reducedCore bool) avr.Instruction {\n")
    switch kind {
    case Flat:
        g.GenerateFlat()
    default:
        panic("Generator.Generate: bad Kind")
    }
    g.Printf("}\n")
}

func (g *Generator) Format() []byte {
    src, err := format.Source(g.buf.Bytes())
    if err != nil {
        log.Printf("warning: invalid Go generated: %s\n", err)
        log.Printf("warning: compile the package to analyse the error\n")
        return g.buf.Bytes()
    }
    return src
}

