package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Repository struct {
	Name        string `json:"name"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
}

func fetchRepositories(username string) ([]Repository, error) {
	client := resty.New()
	resp, err := client.R().
		SetResult([]Repository{}).
		Get(fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username))
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*[]Repository), nil
}

func generateMarkdown(username string, repos []Repository) string {
	markdown := fmt.Sprintf("# %s's Public Repositories\n\n", username)
	for _, repo := range repos {
		description := repo.Description
		if description == "" {
			description = "No description provided."
		}
		markdown += fmt.Sprintf("- [%s](%s): %s\n", repo.Name, repo.HTMLURL, description)
	}
	return markdown
}

func saveMarkdown(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

var (
	username = flag.String("user", "octocat", "user to generate a listing for")
	markdown = flag.String("mdoverride", "", "use an existing markdown input")
)

func main() {
	flag.Parse()
	md := ""
	ht := ""
	if *markdown == "" {
		repos, err := fetchRepositories(*username)
		if err != nil {
			log.Fatalf("Error fetching repositories: %v", err)
		}
		markdown := generateMarkdown(*username, repos)
		md = fmt.Sprintf("%s.md", *username)
		ht = fmt.Sprintf("%s.html", *username)
		err = saveMarkdown(md, markdown)
		if err != nil {
			log.Fatalf("Error saving markdown file: %v", err)
		}
		fmt.Printf("Markdown document %s.md generated successfully.\n", *username)
	} else {
		md = *markdown
		if _, err := os.Stat(md); err != nil {
			log.Fatalf("Error reading markdown file: %s", err)
		}
		trmd := strings.TrimSuffix(md, filepath.Ext(md))
		ht = fmt.Sprintf("%s.html", trmd)
	}
	err := generateHTML(md, ht)
	if err != nil {
		log.Fatalf("Error saving markdown file: %v", err)
	}
}
