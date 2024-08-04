package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ComponentDefinition represents the structure of the input YAML file.
type ComponentDefinition struct {
	ID       string    `yaml:"id"`
	Title    string    `yaml:"title"`
	VERSION  string    `yaml:"VERSION"`
	Controls []Control `yaml:"controls"`
}

// Control represents the structure of each control within the YAML file.
type Control struct {
	ID               string              `yaml:"id"`
	FeatureID        string              `yaml:"feature_id"`
	Title            string              `yaml:"title"`
	Objective        string              `yaml:"objective"`
	NISTCSF          string              `yaml:"nist_csf"`
	MITREAttack      string              `yaml:"mitre_attack"`
	ControlMappings  map[string][]string `yaml:"control_mappings"`
	TestRequirements map[string]string   `yaml:"test_requirements"`
}

// NucleiTemplate represents the structure of the Nuclei template file.
type NucleiTemplate struct {
	ID            string `yaml:"id"`
	Info          Info   `yaml:"info"`
	Code          []Code `yaml:"code"`
	SelfContained bool   `yaml:"self-contained"`
}

// Info represents the structure of the info section within the Nuclei template file.
type Info struct {
	Name     string `yaml:"name"`
	Severity string `yaml:"severity"`
	Author   string `yaml:"author"`
}

// Code represents the structure of the code section within the Nuclei template file.
type Code struct {
	Engine   []string `yaml:"engine"`
	Source   string   `yaml:"source"`
	Matchers []struct {
		Type  string   `yaml:"type"`
		Words []string `yaml:"words"`
	} `yaml:"matchers"`
}

// NucleiProfile represents the structure of the Nuclei profile file.
type NucleiProfile struct {
	Code      bool     `yaml:"code"`
	Templates []string `yaml:"templates"`
	Var       []string `yaml:"var"`
}

var functionNames = []string{}
var SOURCE_FILE string
var VERSION string
var OUTPUT_DIR string
var SERVICE_NAME string

func main() {
	fmt.Println("")
	log.Printf("Beginning Generation Process\n\n")
	// exit with a warning if no arguments are provided
	if len(os.Args) < 2 {
		log.Fatalf("[ERROR] Please provide the YAML component definition file as the first argument.")
	} else if len(os.Args) < 3 {
		log.Fatalf("[ERROR] Please provide the abbreviated name of the CSP as the second argument.")
	} else if len(os.Args[2]) > 8 {
		log.Fatalf("[ERROR] The CSP shortname must be shorter than 8 characters.")
	} else if len(os.Args) < 4 {
		log.Fatalf("[ERROR] Please provide the abbreviated name of the service as the third argument.")
	} else if len(os.Args[2]) > 8 {
		log.Fatalf("[ERROR] The service name must be shorter than 8 characters.")
	} else if len(os.Args) < 5 {
		log.Fatalf("[ERROR] Please provide the VERSION of the service as the fourth argument.")
	} else if len(os.Args[4]) > 8 {
		log.Fatalf("[ERROR] The VERSION must be shorter than 8 characters.")
	}

	setConstants()

	source_data := readData()
	templateFiles := generateNucleiTemplates(source_data)
	createNucleiProfile(templateFiles)
	createGoFiles(functionNames)

	fmt.Println("")
	log.Printf("Generation Process Complete\n\n")
}

func createGoFiles(functionNames []string) {
	createGoModFile()
	createMainGoFile(functionNames)
	createSecurityGoFile(functionNames)
	// goModTidy()
}

func goModTidy() {
	cmd := exec.Command("go", "mod", "tidy")

	// Set the output for the command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run go mod tidy command: %s", err)
	}
}

func createGoModFile() {
	data := struct {
		ServiceName string
	}{ServiceName: SERVICE_NAME}

	tmpl, err := template.ParseFiles("templates/go.mod")
	if err != nil {
		log.Fatalf("Error reading template file: %v", err)
	}

	outFile, err := os.Create(filepath.Join(OUTPUT_DIR, "src", "go.mod"))
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	// Execute the template with data and write to the output file
	err = tmpl.Execute(outFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Println("Writing go.mod file for service")
}

func createMainGoFile(functionNames []string) {
	data := struct {
		FunctionNames []string
	}{FunctionNames: functionNames}

	tmpl, err := template.ParseFiles("templates/main.txt")
	if err != nil {
		log.Fatalf("Error reading template file: %v", err)
	}

	outFile, err := os.Create(filepath.Join(OUTPUT_DIR, "src", "main.go"))
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	// Execute the template with data and write to the output file
	err = tmpl.Execute(outFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Println("Writing main function wireframe to main.go")
}

func createSecurityGoFile(functionNames []string) {
	tmpl, err := template.ParseFiles("templates/security.txt")
	if err != nil {
		log.Fatalf("Error reading template file: %v", err)
	}

	data := struct {
		FunctionNames []string
	}{FunctionNames: functionNames}

	outFile, err := os.Create(filepath.Join(OUTPUT_DIR, "src", "security.go"))
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outFile.Close()

	// Execute the template with data and write to the output file
	err = tmpl.Execute(outFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	log.Println("Writing security functions wireframe to security.go")
}

func setConstants() {
	SOURCE_FILE = os.Args[1]
	SERVICE_NAME = os.Args[3]
	VERSION = os.Args[4]
	OUTPUT_DIR = setupOutputDirs()

	fmt.Printf("Generating Nuclei templates for %s %s from %s\n", SERVICE_NAME, VERSION, SOURCE_FILE)
	fmt.Printf("Profile and templates will be generated in ./%s\n", OUTPUT_DIR)
}

func setupOutputDirs() string {
	provider := os.Args[2]
	outputDir := filepath.Join(provider, SERVICE_NAME)
	// Create the output directories if they don't exist
	create_dirs := filepath.Join(outputDir, "security-templates")
	err := os.MkdirAll(create_dirs, 0755)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	create_dirs = filepath.Join(outputDir, "src")
	err = os.MkdirAll(create_dirs, 0755)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return outputDir
}

func readData() ComponentDefinition {
	if strings.HasPrefix(SOURCE_FILE, "http") {
		return readYAMLURL()
	} else {
		return readYAMLFile()
	}
}

func readYAMLURL() ComponentDefinition {
	resp, err := http.Get(SOURCE_FILE)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch URL: %v", resp.Status)
	}

	var data ComponentDefinition
	decoder := yaml.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatalf("Failed to decode YAML from URL: %v", err)
	}

	return data
}

func readYAMLFile() ComponentDefinition {
	yamlFile, err := ioutil.ReadFile(SOURCE_FILE)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var data ComponentDefinition
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML file: %v", err)
	}

	return data
}

func generateNucleiTemplates(data ComponentDefinition) []string {
	var profileRefNames []string
	var buffer bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2)

	// Create a separate template file for each control
	for _, control := range data.Controls {
		nucleiTemplate := NucleiTemplate{
			Info: Info{
				Severity: "info",
				Author:   "FINOS",
			},
			SelfContained: true,
			Code:          []Code{},
		}

		controlID := generateControlID(control.ID)

		// Override the template data with the control data
		nucleiTemplate.ID = controlID
		nucleiTemplate.Info.Name = fmt.Sprintf("%s %s: %s", SERVICE_NAME, data.Title, control.Title)

		// Create a code section for each test requirement
		for testID := range control.TestRequirements {
			code := createCodeSection(controlID, testID)
			nucleiTemplate.Code = append(nucleiTemplate.Code, code)
		}

		profileRefName := writeNucleiTemplateToFile(controlID, nucleiTemplate, yamlEncoder, &buffer)
		profileRefNames = append(profileRefNames, profileRefName)
	}

	return profileRefNames
}

func generateControlID(controlID string) string {
	return fmt.Sprintf("%s_%s", os.Args[3], strings.Replace(controlID, ".", "_", -1))
}

func createCodeSection(controlID, testID string) Code {
	// the compiled executable will eventually live at this path
	binSrc := filepath.Join("src", "test-exec_"+VERSION)
	functionName := fmt.Sprintf("%s_TR%s", controlID, testID)
	functionNames = append(functionNames, functionName)

	return Code{
		Engine: []string{"zsh"}, // This might need to change depending on what the runners are using
		Source: fmt.Sprintf("\n%s %s", binSrc, functionName),
		Matchers: []struct {
			Type  string   `yaml:"type"`
			Words []string `yaml:"words"`
		}{
			{
				Type:  "word",
				Words: []string{"FAIL", "ERROR"},
			},
		},
	}
}

func writeNucleiTemplateToFile(controlID string, nucleiTemplate NucleiTemplate, yamlEncoder *yaml.Encoder, buffer *bytes.Buffer) string {
	profileRefName := filepath.Join("security-templates", controlID+".yaml")
	filename := filepath.Join(OUTPUT_DIR, profileRefName)
	fmt.Printf("Writing Nuclei template to %s\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	buffer.Reset()
	yamlEncoder.Encode(nucleiTemplate)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	signNuclei(filename)

	return profileRefName
}

func signNuclei(filepath string) {
	cmd := exec.Command("nuclei", "-t", filepath, "-code", "-sign")

	// Set the output for the command
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run nuclei sign command: %s", err)
	}
}

func createNucleiProfile(templateFiles []string) {
	profile := NucleiProfile{
		Code:      true,
		Templates: templateFiles,
		Var:       []string{"region=us-east-1"},
	}
	var buffer bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&buffer)
	yamlEncoder.SetIndent(2)

	profileFile := filepath.Join(OUTPUT_DIR, "security-profile.yaml")
	pfFile, err := os.Create(profileFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer pfFile.Close()

	yamlEncoder.Encode(profile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	_, err = pfFile.Write(buffer.Bytes())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
