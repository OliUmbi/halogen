package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type color struct {
	Name  string
	Value string
}

type breakpoint struct {
	Name  string
	Value string
}

type pseudo struct {
	Name  string
	Value string
}

type variant struct {
	Name  string
	Value string
}

type attribute struct {
	Name       string
	Breakable  bool
	Pseudoable bool
	Variants   []variant
}

type style struct {
	Colors      []color
	Breakpoints []breakpoint
	Pseudos     []pseudo
	Attributes  []attribute
}

func main() {

	var data, err = os.ReadFile("./definitions.json")
	if err != nil {
		panic(err)
	}

	var style style
	if err := json.Unmarshal(data, &style); err != nil {
		panic(err)
	}

	var results []string
	var breakpoints map[string][]string = map[string][]string{}

	results = append(results, ":root{")
	for _, color := range style.Colors {
		results = append(results, "--"+color.Name+":"+color.Value+";")
	}
	results = append(results, "}\n")

	for _, attribute := range style.Attributes {
		for _, variant := range attribute.Variants {

			var classes []string
			classes = append(classes, "."+attribute.Name+"\\="+variant.Name)

			if attribute.Pseudoable {
				for _, pseudo := range style.Pseudos {
					classes = append(classes, "."+attribute.Name+"\\:"+pseudo.Name+"\\="+variant.Name+":"+pseudo.Value)
				}
			}

			results = append(results, strings.Join(classes, ",")+"{"+variant.Value+"}")

			if attribute.Breakable {
				for _, breakpoint := range style.Breakpoints {
					breakpoints[breakpoint.Value] = append(breakpoints[breakpoint.Value], "."+attribute.Name+"\\:"+breakpoint.Name+"\\="+variant.Name+"{"+variant.Value+"}")
				}
			}
		}
	}

	for value, data := range breakpoints {
		results = append(results, "\n@media(width<="+value+"){"+strings.Join(data, "")+"}")
	}

	fmt.Println(utf8.RuneCountInString(strings.Join(results, "")))
	fmt.Println(utf8.RuneCountInString(strings.Join(results, "")) / 8)

	os.WriteFile("./output.css", []byte(strings.Join(results, "")), 0644)
}
