package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(goCmd)
	rootCmd.AddCommand(rustCmd)
	rootCmd.AddCommand(nodeCmd)
	rootCmd.AddCommand(javaCmd)
	rootCmd.AddCommand(pythonCmd)
	rootCmd.AddCommand(phpCmd)
	rootCmd.AddCommand(versionCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Initialize a Go Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Go project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("go")
	},
}

var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "Initialize a Rust Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Rust project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("rust")
	},
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Initialize a Node.js Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Node.js project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("node")
	},
}

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Initialize a Java Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Java project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("java")
	},
}

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Initialize a Python Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a Python project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("python")
	},
}

var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "Initialize a PHP Docker project",
	Long:  `This command sets up the Docker and Docker Compose configuration files for a PHP project.`,
	Run: func(cmd *cobra.Command, args []string) {
		createFiles("php")
	},
}

func createFiles(projectType string) {
	createFile(".dockerignore", "node_modules\n*.log\n")
	createFile("Docker.md", fmt.Sprintf("## Docker setup for %s project\n", projectType))
	createFile("docker-compose.yaml", dockerComposeContent(projectType))
	createFile("Dockerfile", dockerfileContent(projectType))

	fmt.Println("Project initialized successfully.")
}

func createFile(name, content string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Printf("Error creating %s: %v\n", name, err)
		return
	}
	defer file.Close()

	file.WriteString(content)
}

func dockerComposeContent(projectType string) string {
	return fmt.Sprintf(`version: '3'
services:
  app:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - .:/app
    command: %s
`, dockerComposeCommand(projectType))
}

func dockerComposeCommand(projectType string) string {
	switch projectType {
	case "php":
		return "php -S 0.0.0.0:8080 -t /app"
	case "go":
		return "./main"
	case "java":
		return "java -jar /app/app.jar"
	case "python":
		return "python app.py"
	case "node":
		return "node app.js"
	case "rust":
		return "./app"
	default:
		return ""
	}
}

func dockerfileContent(projectType string) string {
	switch projectType {
	case "php":
		return `FROM php:7.4-cli
WORKDIR /app
COPY . .
	`
	case "go":
		return `FROM golang:1.16
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["./main"]
	`
	case "java":
		return `FROM openjdk:11
WORKDIR /app
COPY . .
RUN javac App.java
CMD ["java", "App"]
	`
	case "python":
		return `FROM python:3.9
WORKDIR /app
COPY . .
CMD ["python", "app.py"]
	`
	case "node":
		return `FROM node:14
WORKDIR /app
COPY . .
RUN npm install
CMD ["node", "app.js"]
	`
	case "rust":
		return `FROM rust:1.50
WORKDIR /app
COPY . .
RUN cargo build --release
CMD ["./target/release/app"]
	`
	default:
		return ""
	}
}
