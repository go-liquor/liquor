package templates

import (
	"fmt"
	"html/template"
	"os"
)

// ParseTemplate processes a given template string and writes the output to a specified file.
//
// This function takes a raw template string, compiles it using Go's text/template package,
// and applies the provided data to generate the final content. The result is written to the specified output file.
//
// Parameters:
//   - content: A string containing the raw template content.
//   - outputFile: The name of the file where the processed template will be saved.
//   - data: Any data structure that will be injected into the template.
//
// Returns:
//   - An `error` if any issue occurs while creating the file, parsing the template, or executing the template.
//
// Example:
//
//	content := "Hello, {{.Name}}!"
//	outputFile := "output.txt"
//	data := struct{ Name string }{Name: "Alice"}
//
//	if err := ParseTemplate(content, outputFile, data); err != nil {
//		fmt.Println("Error processing template:", err)
//	}
func ParseTemplate(content string, outputFile string, data any) error {
	newFile, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", outputFile, err.Error())
	}
	tpl, err := template.New("template").Parse(content)
	if err != nil {
		return fmt.Errorf("failed to parse file template: %v", err.Error())
	}
	if err := tpl.Execute(newFile, data); err != nil {
		return fmt.Errorf("failed to execute template write: %v", err.Error())
	}
	return nil
}

// ParseTemplates parses and processes multiple template files with the provided data.
//
// It takes a map where the keys represent the output file names and the values contain the template content.
// Each template is processed using the `ParseTemplate` function, applying the given `data`.
//
// Parameters:
//   - files: A map where each key is an output file name and the corresponding value is the template content as a string.
//   - data: Any data structure that will be injected into the templates.
//
// Returns:
//   - An `error` if any template fails to process; otherwise, `nil`.
//
// Example:
//
//	files := map[string]string{
//		"index.html": "<html><body>{{.Title}}</body></html>",
//	}
//	data := struct{ Title string }{Title: "Hello, World!"}
//
//	if err := ParseTemplates(files, data); err != nil {
//		fmt.Println("Error parsing templates:", err)
//	}
func ParseTemplates(files map[string]string, data any) error {
	for outputFile, content := range files {
		if err := ParseTemplate(content, outputFile, data); err != nil {
			return err
		}
	}
	return nil
}
