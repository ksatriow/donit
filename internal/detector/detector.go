package detector

import (
	"os"
	"path/filepath"
	"strings"
)

type ProjectType string

const (
	ProjectTypeGo      ProjectType = "go"
	ProjectTypeNode    ProjectType = "node"
	ProjectTypePython  ProjectType = "python"
	ProjectTypeJava    ProjectType = "java"
	ProjectTypeRust    ProjectType = "rust"
	ProjectTypePHP     ProjectType = "php"
	ProjectTypeUnknown ProjectType = "unknown"
)

type Detector struct {
	projectPath string
}

func NewDetector(projectPath string) *Detector {
	return &Detector{projectPath: projectPath}
}

func (d *Detector) Detect() (ProjectType, error) {
	indicators := map[ProjectType][]string{
		ProjectTypeGo:     {"go.mod", "go.sum", "main.go"},
		ProjectTypeNode:   {"package.json", "package-lock.json", "yarn.lock", "pnpm-lock.yaml"},
		ProjectTypePython: {"requirements.txt", "pyproject.toml", "setup.py", "Pipfile", "poetry.lock"},
		ProjectTypeJava:   {"pom.xml", "build.gradle", "build.gradle.kts", "settings.gradle"},
		ProjectTypeRust:   {"Cargo.toml", "Cargo.lock"},
		ProjectTypePHP:    {"composer.json", "composer.lock"},
	}

	scores := make(map[ProjectType]int)

	err := filepath.Walk(d.projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories (except current directory)
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") && info.Name() != "." {
			return filepath.SkipDir
		}

		// Check file indicators
		for projectType, files := range indicators {
			for _, indicator := range files {
				if strings.HasSuffix(path, indicator) {
					scores[projectType]++
				}
			}
		}

		return nil
	})

	if err != nil {
		return ProjectTypeUnknown, err
	}

	// Find project type with highest score
	maxScore := 0
	var detectedType ProjectType = ProjectTypeUnknown

	for projectType, score := range scores {
		if score > maxScore {
			maxScore = score
			detectedType = projectType
		}
	}

	return detectedType, nil
}

