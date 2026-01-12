# Todo CLI

A simple command-line TODO manager written in Go using Cobra.

## Features

- Add new tasks
- List all tasks
- Mark tasks as done
- Clear completed tasks

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
```

## File Storage

Tasks are stored in a `tasks.json` file in the current directory.