# Blog Aggregator

This project is a blog aggregator application built with Go. It allows users to collect and manage blog posts from various sources.

## Requirements

- **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/).
- **PostgreSQL**: The application requires a PostgreSQL database. Install and configure PostgreSQL before running the application.

## Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/itency/blog_aggregator.git
    cd blog_aggregator
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure the database:
    - Create a PostgreSQL database.
    - Update the database connection settings in the configuration file.

4. Run the application:
    ```bash
    go run main.go
    ```

## Features

- Collect blog posts from multiple sources.
- Manage and organize blog content.
- User-friendly interface for browsing posts.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.