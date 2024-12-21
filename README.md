# Commiter CLI

Commiter is a command-line interface (CLI) tool written in Go that fetches staged changes from a Git repository and generates commit messages using GitHub models.

## Features

- Fetches staged changes from the Git repository.
- Generates commit messages based on the fetched changes using an external API.

## Installation

To install Commiter, clone the repository and build the project:

```bash
git clone https://github.com/Paranoia8972/commiter.git
cd commiter
go build -o commiter ./cmd
```

## Usage

After building the project, you can use the CLI tool as follows:

```bash
./commiter
```

This command will fetch the staged changes and generate a commit message.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
