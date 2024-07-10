# DONIT CLI

Donit is a command-line interface (CLI) tool for initializing Docker and Docker Compose files for various programming languages. It simplifies the setup process for Docker projects by generating the necessary configuration files for you.

## Features

- Initialize Docker projects for multiple programming languages.
- Generates `.dockerignore`, `Dockerfile`, `docker-compose.yaml`, and `Docker.md`.
- Supports Go, Rust, Node.js, Java, Python, and PHP.
- Provides version information and help commands.

## Installation

### Prerequisites

- Go 1.16 or later installed on your machine.
- Docker installed on your machine.

### Building from Source

1. Clone the repository:

    ```sh
    git clone https://github.com/ksatriow/donit
    cd donit
    ```

2. Build the binary:

    ```sh
    make build
    ```

3. Install the binary:

    ```sh
    sudo make install
    ```

This will place the `donit` binary in `/usr/local/bin`, making it accessible from anywhere on your system.

## Usage

### Basic Commands

- **Initialize a Go project:**

    ```sh
    donit go
    ```

- **Initialize a Rust project:**

    ```sh
    donit rust
    ```

- **Initialize a Node.js project:**

    ```sh
    donit node
    ```

- **Initialize a Java project:**

    ```sh
    donit java
    ```

- **Initialize a Python project:**

    ```sh
    donit python
    ```

- **Initialize a PHP project:**

    ```sh
    donit php
    ```

### Version Command

To check the version of `donit`:

```sh
donit version
```

### Help Command

To get help with `donit`:

```sh
donit help
```

You can also get help for specific commands:

```sh
donit help go
```

## Examples

### Initializing a Go Project

1. Navigate to your project directory:

    ```sh
    cd /path/to/your/go/project
    ```

2. Run the `donit go` command:

    ```sh
    donit go
    ```

3. The following files will be created in your project directory:

    - `.dockerignore`
    - `Dockerfile`
    - `docker-compose.yaml`
    - `Docker.md`

### Initializing a Node.js Project

1. Navigate to your project directory:

    ```sh
    cd /path/to/your/node/project
    ```

2. Run the `donit node` command:

    ```sh
    donit node
    ```

3. The following files will be created in your project directory:

    - `.dockerignore`
    - `Dockerfile`
    - `docker-compose.yaml`
    - `Docker.md`

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or bug fixes.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or suggestions, feel free to reach out at [kukuhsatriowibowo@gmail.com](mailto:kukuhsatriowibowo@gmail.com).



**Happy Coding!**
