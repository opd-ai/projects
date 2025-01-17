package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func generateHTML(inputFile, outputFile string) error {
	mdContent, err := readFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading input file: %v\n", err)
	}

	htmlContent := convertMarkdownToHTML(mdContent)

	// Extract title and description from markdown
	lines := strings.Split(string(mdContent), "\n")
	title := extractTitle(lines)
	description := extractDescription(lines)

	// Create the HTML content with template
	finalOutput := createHTMLTemplate(title, description, string(htmlContent))

	err = writeFile(outputFile, finalOutput)
	if err != nil {
		return fmt.Errorf("Error writing output file: %v\n", err)
	}

	fmt.Printf("Generated HTML file: %s\n", outputFile)
	return nil
}

func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0o644)
}

func convertMarkdownToHTML(mdContent []byte) []byte {
	return blackfriday.Run(mdContent)
}

func extractTitle(lines []string) string {
	if len(lines) > 0 && strings.HasPrefix(lines[0], "# ") {
		return strings.TrimPrefix(lines[0], "# ")
	}
	return "Untitled"
}

func extractDescription(lines []string) string {
	if len(lines) > 1 && strings.HasPrefix(lines[1], "> ") {
		return strings.TrimPrefix(lines[1], "> ")
	}
	return ""
}

func createHTMLTemplate(title, description, body string) string {
	htmlTemplate := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <meta name="description" content="%s"/>
    
    <!-- Google Fonts -->
    <link href="https://fonts.googleapis.com/css2?family=MedievalSharp&family=Crimson+Text:ital,wght@0,400;0,700;1,400&display=swap" rel="stylesheet">
    
    <link rel="stylesheet" href="/assets/css/style.css">
    
    <!-- Syntax highlighting -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/styles/github-dark.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.9.0/highlight.min.js"></script>
    <script>hljs.highlightAll();</script>
</head>
<body>
    <header class="hero">
        <div class="hero-content">
            <h1 id="hero-headline">%s</h1>
            <div id="hero-intro">%s</div>
        </div>
    </header>
    <main class="main-content">
        <div class="content-display">
            %s
        </div>
    </main>
</body>
</html>`
	return fmt.Sprintf(htmlTemplate, title, description, title, description, body)
}
