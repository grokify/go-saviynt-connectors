package main

import (
	"fmt"
	"log"

	gosaviyntconnectors "github.com/grokify/go-saviynt-connectors"
	"github.com/grokify/mogo/fmt/fmtutil"
)

func main() {
	dir := "/Users/johnwang/go/src/github.com/saviynt/saviynt-connectors/src/connectors"

	set := gosaviyntconnectors.NewConnectionSet()
	err := set.ReadDir(dir, true)
	fmtutil.MustPrintJSON(err)

	for connKey, conn := range set.Map {
		fmt.Printf("PROC (%s)\n", connKey)
		fmtutil.PrintJSON(conn)
		names := conn.ExternalAttrs.Names(true, true, true, true)
		fmtutil.MustPrintJSON(names)
		break
	}

	err = set.ExtendedAttributesWriteXLSX("community-connectors.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
