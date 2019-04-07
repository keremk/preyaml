package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type templateConfig struct {
	Vars map[string]string
}

func Generate(templateFile string, data string, output string) {
	templateContents := readTemplate(templateFile)
	templateConfig := readData(data)

	t, err := template.New("foo").Parse(templateContents)
	if err != nil {
		fmt.Println("Can not parse template ", err)
		os.Exit(1)
	}

	f, err := os.Create(output)
	if err != nil {
		fmt.Println("Can not create output file ", err)
		os.Exit(1)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	err = t.Execute(w, templateConfig)
	if err != nil {
		fmt.Println("Can not execute template ", err)
		os.Exit(1)
	}
	w.Flush()
}

func readTemplate(template string) string {
	templateContents, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Println("Cannot read template ", err)
		os.Exit(1)
	}
	return string(templateContents)
}

func readData(data string) *templateConfig {
	variables := make(map[string]string)
	f, err := os.Open(data)
	if err != nil {
		fmt.Print("Cannot open data file", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f) // f is the *os.File
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		pairs := strings.Split(scanner.Text(), "=")
		variables[pairs[0]] = pairs[1]
	}
	if err := scanner.Err(); err != nil {
		// handle error
		fmt.Println("Cannot read data file ", err)
		os.Exit(1)
	}
	return &templateConfig{Vars: variables}
}
