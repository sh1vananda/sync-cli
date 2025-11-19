package scanner

import (
	"strings"
)

// ScorePackage calculates a development score (0-100) for a package
func ScorePackage(name, publisher, path, description string) int {
	score := 0
	nameLower := strings.ToLower(name)
	publisherLower := strings.ToLower(publisher)
	pathLower := strings.ToLower(path)
	descLower := strings.ToLower(description)

	// 1. Exact Match Analysis (High Confidence)
	exactKeywords := map[string]int{
		"go": 40, "git": 40, "node": 40, "npm": 40, "pip": 40,
		"python": 40, "java": 40, "rust": 40, "cargo": 40,
		"docker": 40, "kubectl": 40, "helm": 40, "terraform": 40,
		"code": 30, "vim": 30, "nvim": 30, "emacs": 30,
	}

	if val, ok := exactKeywords[nameLower]; ok {
		score += val
	}

	// 2. Keyword Analysis (Name & Description)
	// Use word boundaries or longer terms to avoid false positives
	devKeywords := []string{
		"sdk", "cli", "compiler", "debugger", "editor", "terminal",
		"framework", "library", "api", "build", "lint", "test",
		"kubernetes", "aws", "azure", "gcp", "ansible",
		"postgres", "mysql", "redis", "mongo", "sqlite",
		"clang", "gcc", "cmake", "make", "gradle", "maven",
	}

	for _, kw := range devKeywords {
		if strings.Contains(nameLower, kw) {
			score += 20
		}
		if strings.Contains(descLower, kw) {
			score += 10
		}
	}

	// 3. Publisher Analysis
	devPublishers := []string{
		"microsoft", "jetbrains", "hashicorp", "canonical", "openjs",
		"python software foundation", "google", "amazon", "oracle",
		"mozilla", "apache", "github", "docker",
	}

	for _, pub := range devPublishers {
		if strings.Contains(publisherLower, pub) {
			score += 15
		}
	}

	// 4. Path Analysis
	if strings.Contains(pathLower, "bin") || strings.Contains(pathLower, "cmd") {
		score += 10
	}
	if strings.Contains(pathLower, "tools") {
		score += 10
	}
	// Go specific path heuristic
	if strings.Contains(pathLower, "go/bin") || strings.Contains(pathLower, "go\\bin") {
		score += 20
	}

	// 5. Negative Signals
	negativeKeywords := []string{
		"game", "player", "driver", "update", "service", "redistributable",
		"runtime", "viewer", "media", "audio", "video", "install", "setup",
		"background", "helper", "assistant",
	}

	for _, kw := range negativeKeywords {
		if strings.Contains(nameLower, kw) {
			score -= 30
		}
	}

	// Cap score
	if score > 100 {
		score = 100
	}
	if score < 0 {
		score = 0
	}

	return score
}
