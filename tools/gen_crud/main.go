package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ModelData struct {
	ModelName            string
	LowerModelName       string
	ModelNamePlural      string
	LowerModelNamePlural string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <ModelName>")
	}
	modelName := os.Args[1]

	data := ModelData{
		ModelName:            strings.Title(modelName),
		LowerModelName:       strings.ToLower(modelName),
		ModelNamePlural:      strings.Title(modelName) + "s",
		LowerModelNamePlural: strings.ToLower(modelName) + "s",
	}

	templates := map[string]string{
		"internal/model/{{.LowerModelName}}.go":                   "tools/gen_crud/templates/model.go.tpl",
		"internal/request/{{.LowerModelName}}.go":                 "tools/gen_crud/templates/request.go.tpl",
		"internal/repository/{{.LowerModelName}}_repository.go": "tools/gen_crud/templates/repository.go.tpl",
		"internal/service/{{.LowerModelName}}_service.go":         "tools/gen_crud/templates/service.go.tpl",
		"internal/handler/{{.LowerModelName}}_handler.go":         "tools/gen_crud/templates/handler.go.tpl",
	}

	for outputPathTpl, tplFile := range templates {
		// Create output path
		var outputPathBuf bytes.Buffer
		pathTpl, err := template.New("path").Parse(outputPathTpl)
		if err != nil {
			log.Fatalf("Failed to parse output path template: %v", err)
		}
		if err := pathTpl.Execute(&outputPathBuf, data); err != nil {
			log.Fatalf("Failed to execute output path template: %v", err)
		}
		outputPath := outputPathBuf.String()

		// Create directory if not exists
		if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
			log.Fatalf("Failed to create directory for %s: %v", outputPath, err)
		}

		// Read template
		tpl, err := template.ParseFiles(tplFile)
		if err != nil {
			log.Fatalf("Failed to parse template file %s: %v", tplFile, err)
		}

		// Create output file
		outputFile, err := os.Create(outputPath)
		if err != nil {
			log.Fatalf("Failed to create output file %s: %v", outputPath, err)
		}
		defer outputFile.Close()

		// Execute template
		if err := tpl.Execute(outputFile, data); err != nil {
			log.Fatalf("Failed to execute template for %s: %v", outputPath, err)
		}

		fmt.Printf("Generated file: %s\n", outputPath)
	}

	// Inject routes into main.go
	injectRoutes(data)
}

func injectRoutes(data ModelData) {
	mainGoPath := "cmd/server/main.go"
	content, err := os.ReadFile(mainGoPath)
	if err != nil {
		log.Fatalf("Failed to read main.go: %v", err)
	}

	// Prepare route snippet
	routeSnippetTpl := `
	// Init {{.ModelName}}
	{{.LowerModelName}}Repo := repository.New{{.ModelName}}Repository(db)
	{{.LowerModelName}}Service := service.New{{.ModelName}}Service({{.LowerModelName}}Repo)
	{{.LowerModelName}}Handler := handler.New{{.ModelName}}Handler({{.LowerModelName}}Service)

	// {{.ModelName}} routes
	{{.LowerModelName}}Group := r.Group("/api/{{.LowerModelNamePlural}}")
	{{.LowerModelName}}Group.Use(middleware.AuthMiddleware())
	{
		{{.LowerModelName}}Group.POST("/", {{.LowerModelName}}Handler.Create)
		{{.LowerModelName}}Group.GET("/:id", {{.LowerModelName}}Handler.Get)
		{{.LowerModelName}}Group.PUT("/:id", {{.LowerModelName}}Handler.Update)
		{{.LowerModelName}}Group.DELETE("/:id", {{.LowerModelName}}Handler.Delete)
		{{.LowerModelName}}Group.GET("/", {{.LowerModelName}}Handler.List)
	}
`
	ttpl, err := template.New("routes").Parse(routeSnippetTpl)
	if err != nil {
		log.Fatalf("Failed to parse route snippet template: %v", err)
	}
	var snippetBuf bytes.Buffer
	if err := tpl.Execute(&snippetBuf, data); err != nil {
		log.Fatalf("Failed to execute route snippet template: %v", err)
	}
	routeSnippet := snippetBuf.String()

	// Find injection point and insert
	injectionPoint := "// Protected routes"
	newContent := strings.Replace(string(content), injectionPoint, injectionPoint+"\n"+routeSnippet, 1)

	if err := os.WriteFile(mainGoPath, []byte(newContent), 0644); err != nil {
		log.Fatalf("Failed to write updated main.go: %v", err)
	}

	fmt.Println("Successfully injected routes into cmd/server/main.go")
}
