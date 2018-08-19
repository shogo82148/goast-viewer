// +build js,wasm

package main

import (
	"encoding/json"
	"log"
	"strings"
	"syscall/js"
)

type Result struct {
	*Ast   `json:"ast"`
	Source string `json:"source"`
	Dump   string `json:"dump"`
}

func main() {
	js.Global().Get("document").Call("getElementById", "btn-parse").Set("disabled", js.ValueOf(false))

	js.Global().Set("GoASTParse", js.NewCallback(func(args []js.Value) {
		source := args[0].String()
		source = strings.Replace(source, "\r", "", -1)
		ast, dump, err := Parse("", source)
		if err != nil {
			log.Printf("handleParse : Failed to convert Ast to json %+v\n", err)
			return
		}
		result := Result{Ast: ast, Source: source, Dump: dump}
		body, err := json.Marshal(result)
		if err != nil {
			log.Printf("handleParse: Failed to marshal Ast %+v\n", err)
			return
		}
		args[1].Invoke(string(body))
	}))

	select {}
}
