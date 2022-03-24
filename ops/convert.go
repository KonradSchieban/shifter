/*
copyright 2019 google llc
licensed under the apache license, version 2.0 (the "license");
you may not use this file except in compliance with the license.
you may obtain a copy of the license at
    http://www.apache.org/licenses/license-2.0
unless required by applicable law or agreed to in writing, software
distributed under the license is distributed on an "as is" basis,
without warranties or conditions of any kind, either express or implied.
see the license for the specific language governing permissions and
limitations under the license.
*/

package ops

import (
	"fmt"
	"log"
	"path/filepath"
	generators "shifter/generators"
	inputs "shifter/inputs"

	"github.com/google/uuid"
)

// Input Types
const YAML string = "YAML"
const TEMPLATE string = "template"

// Create New Converter
func NewConverter(inputType string, sourcePath string, generator string, outputPath string, flags map[string]string) *Converter {
	// Create New Instance of Converter
	converter := &Converter{}

	// Create UUID for Converter
	converter.UUID = uuid.New().String()

	// Set all the Variables for the Converter
	converter.InputType = inputType
	converter.SourcePath = sourcePath
	converter.Generator = generator
	converter.OutputPath = outputPath
	converter.Flags = flags

	// Process the Path and Create Array of File Objects
	files, err := ProcessPath(converter.SourcePath)
	if err != nil {
		log.Println(err)
	}

	// Set Converter Files
	converter.SourceFiles = files

	return converter
}

func (converter *Converter) WriteSourceFiles() {
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		file.WriteFile()
	}
}

func (converter *Converter) LoadSourceFiles() {
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		file.LoadFile()
	}
}

func (converter *Converter) ListSourceFiles() {
	// Process Input Objects
	for _, file := range converter.SourceFiles {
		file.Meta()
	}
}

func (converter *Converter) ListOutputFiles() {
	// Process Input Objects
	for _, file := range converter.OutputFiles {
		file.Meta()
	}
}

func (converter *Converter) ConvertFiles() {
	// Process Input Objects
	for idx, file := range converter.SourceFiles {

		// Run Conversion..... HERE
		// Store Return Buffer in New File and Write File

		fileObj := &FileObject{
			StorageType:   file.StorageType,
			SourcePath:    (converter.OutputPath + "/" + fmt.Sprint(idx) + " - " + filepath.Ext(file.SourcePath)),
			Ext:           filepath.Ext(file.SourcePath),
			Content:       file.Content,
			ContentLength: file.ContentLength,
		}

		// Write Converted File to Storage
		fileObj.WriteFile()
		// Add Converted File Object to Converter
		converter.OutputFiles = append(converter.OutputFiles, fileObj)
	}
}

/*
	TODO
	- Add Errors Handling to Convert,
	- Catch Convert Errors,
	- Return error struct on Errors
*/
func Convert(inputType string, sourcePath string, generator string, outputPath string, flags map[string]string) {

	con := NewConverter(inputType, sourcePath, generator, outputPath, flags)
	//con.ListSourceFiles()
	con.LoadSourceFiles()
	con.ConvertFiles()
	//con.ListSourceFiles()

	switch inputType {
	case "template":
		t, p, n := inputs.Template(sourcePath, flags)
		switch generator {
		case "helm":
			generators.Helm(outputPath, t, p, n)
		}
	case "yaml":
		t := inputs.Yaml(sourcePath, flags)
		switch generator {
		case "yaml":
			generators.Yaml(outputPath, t, "gcs")
		}
	case "cluster":
		log.Fatal("Openshift resources have not been implemented yet!")
	}
	log.Println("Conversion completed")
}
