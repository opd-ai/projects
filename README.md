# Projects

This repository contains a tool that leverages the GitHub API to fetch and compile a list of a user's public repositories, presenting them in a markdown document with links and descriptions for easy reference. It automates the process of creating a comprehensive overview of a user's GitHub projects.

## Features

- Fetches all public repositories for a specified GitHub user.
- Generates a markdown document with an unordered list of repository names linked to their GitHub pages.
- Includes detailed descriptions for each repository.
- Easy-to-use and customizable.

## Getting Started

### Prerequisites

- Go (for building and running the tool)
- GitHub account and access token (if needed for higher rate limits)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/opd-ai/projects.git
   cd projects
   ```

2. Install dependencies:

   ```sh
   go get
   ```

### Usage

1. Open the `main.go` file and replace `username := "octocat"` with the desired GitHub username.

2. Build and run the tool:

   ```sh
   go build
   ./projects
   ```

3. The tool will generate a markdown file named `{username}_repos.md` in the current directory.

### Example

For example, if you set the username to `octocat`, the tool will generate a file named `octocat_repos.md` with the following content:

```markdown
# octocat's Public Repositories

- [repo1](https://github.com/octocat/repo1): Description for repo1.
- [repo2](https://github.com/octocat/repo2): Description for repo2.
- ...
```

## Contributing

We welcome contributions! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License.

## Contact

For any questions or feedback, please feel free to reach out via GitHub Issues or contact us directly.