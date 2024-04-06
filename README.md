# Golang DoNextApp

This project is a simple To-Do application written in Go, designed as a learning exercise to understand the fundamentals of Go programming, including its syntax, standard libraries, and app structure. The initial version of the application provides a CLI for managing tasks, with plans for future enhancements to incorporate REST APIs and database integration for learning.

## Features

- **CLI Operations**: Users can add, view, and delete tasks via the command line.
- **File Persistence**: Tasks are persisted to a file in JSON format, allowing them to be retained across application restarts.
- **Simple and Extensible Codebase**: The project is structured to facilitate understanding and further development.

## Getting Started

### Prerequisites

- Go (Version 1.16 or later recommended)

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/hariscats/DoNextApp.git
   ```
2. Navigate to the project directory:
   ```
   cd DoNextApp
   ```
3. Run the application:
   ```
   go run main.go
   ```

## Usage

The application supports the following commands:

- `add`: Add a new task. You will be prompted to enter the task description.
- `view`: View all tasks.
- `delete`: Delete a task. You will be prompted to enter the task ID.

## Future Enhancements

- **REST API**: Refactor the application to serve tasks over a RESTful API, enabling integration with frontend technologies and other systems.
- **Database Integration**: Migrate from file-based persistence to a relational or NoSQL database for scalable and efficient data storage and retrieval.
- **Authentication**: Implement user authentication and authorization to support multiple users with private task lists.
- **Containerize**: Containerize the application for easy deployment and scalability.