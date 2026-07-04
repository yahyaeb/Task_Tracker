# Task Tracker CLI

A simple command-line interface (CLI) application written in Go to track and manage your tasks — what you need to do, what you're currently working on, and what you've completed.

This project was built as a practical exercise in working with the filesystem, handling positional CLI arguments, and building a self-contained CLI tool using only Go's standard library (no external dependencies).

## Features

- Add, update, and delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List all tasks
- Filter tasks by status: `done`, `todo`, `in-progress`
- Persistent storage in a local JSON file
- Automatic creation of the JSON file if it doesn't exist
- Graceful error handling for invalid input, missing tasks, and malformed data

## Requirements

- Go (1.18+ recommended)
- No external libraries or frameworks — built entirely with Go's standard library (`encoding/json`, `os`, `fmt`, etc.)

## Installation / Build

Clone the repository and build the binary:

```bash
git clone <your-repo-url>
cd task-tracker
go build -o task-cli
```

This produces an executable named `task-cli` in the project directory. You can move it to a directory on your `PATH` if you want to run it from anywhere:

```bash
mv task-cli /usr/local/bin/
```

## Usage

Tasks are stored in a `tasks.json` file created automatically in the current working directory the first time you run the CLI.

### Add a task

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Update a task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task

```bash
task-cli delete 1
```

### Mark a task's status

```bash
task-cli mark-in-progress 1
task-cli mark-done 1
```

### List tasks

```bash
# List all tasks
task-cli list

# List tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## Task Properties

Each task stored in `tasks.json` has the following properties:

| Property      | Description                                      |
|---------------|---------------------------------------------------|
| `id`          | Unique identifier for the task                    |
| `description` | Short description of the task                     |
| `status`      | One of `todo`, `in-progress`, `done`              |
| `createdAt`   | Timestamp when the task was created               |
| `updatedAt`   | Timestamp when the task was last updated          |

Example entry in `tasks.json`:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2026-07-04T10:15:00Z",
  "updatedAt": "2026-07-04T10:15:00Z"
}
```

## Data Storage

- Tasks are stored in a `tasks.json` file in the current directory.
- The file is created automatically on first use if it doesn't already exist.
- All reads/writes go through Go's native `os` and `encoding/json` packages — no external database or library is used.

## Error Handling

The CLI handles common edge cases gracefully, including:

- Missing or invalid command arguments
- Attempting to update, delete, or mark a task ID that doesn't exist
- An empty or corrupted `tasks.json` file
- Invalid status filters passed to `list`

## Project Structure

```
task-tracker/
├── main.go        # Entry point, argument parsing and command dispatch
├── tasks.json      # Generated at runtime — stores task data
└── README.md
```

## Notes

This project is based on the [roadmap.sh Task Tracker CLI project](https://roadmap.sh/projects/task-tracker) and was implemented in Go as a way to practice CLI design, JSON file handling, and standard-library-only development.

## License

MIT