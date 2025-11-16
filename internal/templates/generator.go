package templates

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed all:templates
var templateFS embed.FS

type TemplateData struct {
	ProjectName string
	Port        int
	BaseImage   string
	BuildArgs   map[string]string
	Language    string
	Version     string
}

type Generator struct {
	templates map[string]*template.Template
}

func NewGenerator() (*Generator, error) {
	g := &Generator{
		templates: make(map[string]*template.Template),
	}

	// Load templates for each language
	languages := []string{"go", "node", "python", "java", "rust", "php"}
	for _, lang := range languages {
		if err := g.loadTemplates(lang); err != nil {
			return nil, fmt.Errorf("failed to load templates for %s: %w", lang, err)
		}
	}

	return g, nil
}

func (g *Generator) loadTemplates(lang string) error {
	templateFiles := []string{
		"Dockerfile.tmpl",
		"docker-compose.yaml.tmpl",
		"dockerignore.tmpl",
		"Docker.md.tmpl",
	}

	for _, filename := range templateFiles {
		path := fmt.Sprintf("templates/%s/%s", lang, filename)
		content, err := templateFS.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		tmpl, err := template.New(filename).Parse(string(content))
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		key := fmt.Sprintf("%s:%s", lang, filename)
		g.templates[key] = tmpl
	}

	return nil
}

func (g *Generator) Generate(lang, filename string, data TemplateData) ([]byte, error) {
	key := fmt.Sprintf("%s:%s", lang, filename)
	tmpl, exists := g.templates[key]
	if !exists {
		return nil, fmt.Errorf("template not found: %s", key)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.Bytes(), nil
}

func (g *Generator) GenerateFiles(lang string, data TemplateData, outputDir string) error {
	files := map[string]string{
		"Dockerfile.tmpl":           "Dockerfile",
		"docker-compose.yaml.tmpl":  "docker-compose.yaml",
		"dockerignore.tmpl":          ".dockerignore",
		"Docker.md.tmpl":             "Docker.md",
	}

	for tmplName, outputName := range files {
		content, err := g.Generate(lang, tmplName, data)
		if err != nil {
			return fmt.Errorf("failed to generate %s: %w", outputName, err)
		}

		outputPath := filepath.Join(outputDir, outputName)
		if err := os.WriteFile(outputPath, content, 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", outputPath, err)
		}
	}

	return nil
}

