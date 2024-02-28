package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	patch "github.com/evanphx/json-patch/v5"
)

// patch file name
var filename string = "patch-age.json"

func init() {
	flag.StringVar(&filename, "p", filename, "patch file name")
}

type User struct {
	Name  string `json:"name,omitempty"`
	Age   int    `json:"age,omitempty"`
	Sleep bool   `json:"sleep,omitempty"`
}

func main() {
	flag.Parse()

	u := &User{Name: "Flandre", Age: 495}
	printStructAsJSON("Original", u)
	if err := patchStruct(u, filename); err != nil {
		panic(err)
	}
	printStructAsJSON("Modified", u)
}

func printStructAsJSON(prefix string, target any) {
	j, _ := json.Marshal(target)
	fmt.Printf("%s: %s\n", prefix, j)
}

func patchStruct(target any, filename string) error {
	originalJSON, err := json.Marshal(target)
	if err != nil {
		return fmt.Errorf("encode original struct fail: %w", err)
	}

	patchJSON, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read patch file fail: %w", err)
	}

	p, err := patch.DecodePatch(patchJSON)
	if err != nil {
		return fmt.Errorf("decode patch json fail: %w", err)
	}

	patchedJSON, err := p.Apply(originalJSON)
	// patchedJSON, err := patch.MergePatch(originalJSON, patchJSON)
	if err != nil {
		return fmt.Errorf("patch json fail: %w", err)
	}

	if err := json.Unmarshal(patchedJSON, target); err != nil {
		return fmt.Errorf("decode patched json fail: %w", err)
	}
	return nil
}
