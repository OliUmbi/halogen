package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type color struct {
	Name string
	Value string
}

type breakpoint struct {
	Name string
	Value string
}

type pseudo struct {
	Name string
	Value string
}

type variant struct {
	Name string
	Value string
}

type attribute struct {
	Name string
	Variants []variant
}

type style struct {
	Colors []color
	Breakpoints []breakpoint
	Pseudos []pseudo
	Attributes []attribute
}

func main() {

	data, err := os.ReadFile("./definitions.json")
	if err != nil {
			panic(err)
	}

	var style style
	if err := json.Unmarshal(data, &style); err != nil {
			panic(err)
	}

	var results []string

	for _, attribute := range style.Attributes {
		for _, variant := range attribute.Variants {
			results = append(results, "." + attribute.Name + "-" + variant.Name + "{" + variant.Value + "}")

			for _, breakpoint := range style.Breakpoints {
				results = append(results, "." + attribute.Name + "-" + variant.Name + "-" + breakpoint.Name + "{@media(width<=" + breakpoint.Value + "){" + variant.Value + "}}")
			}

			for _, pseudo := range style.Pseudos {
				results = append(results, "." + attribute.Name + "-" + variant.Name + "-" + pseudo.Name + ":" + pseudo.Value + "{" + variant.Value + "}")
			}
		}
	}

	fmt.Println(utf8.RuneCountInString(strings.Join(results, "")))

}

