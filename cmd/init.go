package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"donit/internal/detector"
	"donit/internal/templates"

	"github.com/spf13/cobra"
)

var (
	port      int
	outputDir string
	force     bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all supported programming languages",
	Long:  `List all supported programming languages that can be initialized with donit.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported languages:")
		fmt.Println("  • go     - Go (Golang)")
		fmt.Println("  • node   - Node.js")
		fmt.Println("  • python - Python")
		fmt.Println("  • java   - Java")
		fmt.Println("  • rust   - Rust")
		fmt.Println("  • php    - PHP")
		fmt.Println("\nUsage: donit <language> or donit init [language]")
	},
}

func init() {
	rootCmd.AddCommand(goCmd)
	rootCmd.AddCommand(rustCmd)
	rootCmd.AddCommand(nodeCmd)
	rootCmd.AddCommand(javaCmd)
	rootCmd.AddCommand(pythonCmd)
	rootCmd.AddCommand(phpCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(listCmd)

	// Add flags to language commands
	for _, cmd := range []*cobra.Command{goCmd, rustCmd, nodeCmd, javaCmd, pythonCmd, phpCmd, initCmd} {
		cmd.Flags().IntVarP(&port, "port", "p", 8080, "Port number for the application")
		cmd.Flags().StringVarP(&outputDir, "output", "o", ".", "Output directory")
		cmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing files without prompting")
	}
}

var initCmd = &cobra.Command{
	Use:   "init [language]",
	Short: "Initialize Docker files for a project",
	Long: `Initialize Docker and Docker Compose files for a project.

If language is not specified, donit will attempt to auto-detect the project type.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Initialize a Go Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Go project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("go")
	},
}

var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "Initialize a Rust Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Rust project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("rust")
	},
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Initialize a Node.js Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Node.js project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("node")
	},
}

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Initialize a Java Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Java project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("java")
	},
}

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Initialize a Python Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Python project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("python")
	},
}

var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "Initialize a PHP Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a PHP project.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInitForLanguage("php")
	},
}

func runInit(cmd *cobra.Command, args []string) error {
	var language string

	// Detect or use provided language
	if len(args) > 0 {
		language = args[0]
	} else {
		det := detector.NewDetector(".")
		detected, detectErr := det.Detect()
		if detectErr != nil {
			return fmt.Errorf("failed to detect project type: %w", detectErr)
		}
		if detected == detector.ProjectTypeUnknown {
			return fmt.Errorf("could not detect project type. Please specify language: donit init <language>")
		}
		language = string(detected)
		fmt.Printf("✓ Detected project type: %s\n", language)
	}

	return runInitForLanguage(language)
}

func runInitForLanguage(language string) error {
	// Validate language
	supportedLanguages := map[string]bool{
		"go":     true,
		"node":   true,
		"python": true,
		"java":   true,
		"rust":   true,
		"php":    true,
	}

	if !supportedLanguages[language] {
		return fmt.Errorf("unsupported language: %s. Supported languages: go, node, python, java, rust, php", language)
	}

	// Validate port
	if port < 1 || port > 65535 {
		return fmt.Errorf("invalid port number: %d. Port must be between 1 and 65535", port)
	}

	// Validate output directory
	if outputDir != "." {
		info, err := os.Stat(outputDir)
		if err != nil {
			if os.IsNotExist(err) {
				// Create directory if it doesn't exist
				if err := os.MkdirAll(outputDir, 0755); err != nil {
					return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
				}
			} else {
				return fmt.Errorf("failed to access output directory %s: %w", outputDir, err)
			}
		} else if !info.IsDir() {
			return fmt.Errorf("output path %s is not a directory", outputDir)
		}
	}

	// Check for existing files
	if !force {
		if err := checkExistingFiles(outputDir); err != nil {
			return err
		}
	}

	// Get project name
	projectName := getProjectName()

	// Prepare template data
	data := templates.TemplateData{
		ProjectName: projectName,
		Port:        port,
		BaseImage:   getBaseImage(language),
		Language:    language,
		Version:     getDefaultVersion(language),
		BuildArgs:   make(map[string]string),
	}

	// Initialize generator
	generator, err := templates.NewGenerator()
	if err != nil {
		return fmt.Errorf("failed to initialize generator: %w", err)
	}

	// Generate files
	if err := generator.GenerateFiles(language, data, outputDir); err != nil {
		return fmt.Errorf("failed to generate files: %w", err)
	}

	fmt.Printf("✓ Successfully initialized Docker files for %s project!\n", language)
	fmt.Printf("📁 Files created in: %s\n", outputDir)
	fmt.Printf("🚀 Run 'docker-compose up --build' to start your application\n")

	return nil
}

func checkExistingFiles(dir string) error {
	files := []string{"Dockerfile", "docker-compose.yaml", ".dockerignore", "Docker.md"}
	var existing []string

	for _, file := range files {
		path := filepath.Join(dir, file)
		if _, err := os.Stat(path); err == nil {
			existing = append(existing, file)
		}
	}

	if len(existing) > 0 {
		fmt.Printf("⚠️  The following files already exist: %v\n", existing)
		fmt.Print("Do you want to overwrite them? (y/N): ")

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		response = strings.TrimSpace(strings.ToLower(response))
		if response != "y" && response != "yes" {
			return fmt.Errorf("operation cancelled by user")
		}
		fmt.Println("✓ Proceeding with overwrite...")
	}

	return nil
}

func getProjectName() string {
	wd, err := os.Getwd()
	if err != nil {
		return "app"
	}
	return filepath.Base(wd)
}

func getBaseImage(lang string) string {
	defaults := map[string]string{
		"go":     "golang:1.21-alpine",
		"node":   "node:20-alpine",
		"python": "python:3.11-alpine",
		"java":   "eclipse-temurin:17-jre-alpine",
		"rust":   "rust:1.75-alpine",
		"php":    "php:8.2-fpm-alpine",
	}
	return defaults[lang]
}

func getDefaultVersion(lang string) string {
	versions := map[string]string{
		"go":     "1.21",
		"node":   "20",
		"python": "3.11",
		"java":   "17",
		"rust":   "1.75",
		"php":    "8.2",
	}
	return versions[lang]
}
