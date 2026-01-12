# Todo CLI

A simple command-line TODO manager written in Go using Cobra.

## Features

- Add new tasks with colorful confirmation
- List all tasks with visual indicators (✔ for done, • for pending)
- Mark tasks as done
- Clear completed tasks
- Colored output for better user experience
- Shell completion support

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Fatih Color](https://github.com/fatih/color) - Terminal colors

## Installation

```bash
go build -o todo
```

## Usage

```bash
# Add a new task
./todo add "Your task description"

# List all tasks
./todo list

# Mark a task as done
./todo done <task_id>

# Clear completed tasks
./todo clear

# Generate shell completion
./todo completion [bash|zsh|fish|powershell]

# Get help
./todo --help
./todo [command] --help
```

## File Storage

Tasks are stored in a `tasks.json` file in the `~/.todo/` directory in your home folder. The directory is automatically created if it doesn't exist.
