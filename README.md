# Golang Echo REST API

This repository contains a REST API developed using the Go programming language and the Echo framework.

## Features

- Lightweight and fast HTTP router with middleware support
- Easy-to-use request handling and routing
- Customizable middleware stack for authentication, logging, etc.
- JSON serialization and deserialization
- Error handling and validation

## Prerequisites

The development of this project is done using the following versions:

- Go (version 1.20)
- Echo Framework (version 4.9.0)
- PostgreSQL (version 15.2)
- Logrus for Logging (version 1.9.0)
- Crypto (Argon2) for Improving System Security (version 0.8.0)
- Go Open Telemetry for Improving System Observability (version 1.10.0)

To ensure a smooth execution of the program, it is recommended to have the corresponding versions installed. However, you can adjust the versions based on your specific needs and requirements.

## Execution

#### 1. Clone the repository:

```shell
git clone https://github.com/dzikrurrohmani/golang-echo-rest-api.git
```

#### 2. Install the dependencies:

```shell
go mod download
```

#### 3. Configure the database connection in the `config/config.go` file.

#### 4. Run the application:

```shell
API_URL=some_host:some_port DB_HOST=some_host DB_PORT=some_port DB_USER=some_user DB_PASS=some_password DB_NAME=some_name SIGN_KEY=some_secret_key go run ./cmd/main.go
```

## API Endpoints:

- `GET /menu`: Get menu list
- `POST /order`: Create an order
- `GET /order/:id`: Get a specific order
- `POST /user/register`: Create a user
- `POST /user/login`: login into user account

## Contributing:

We welcome contributions from the community. If you find any issues or have suggestions for improvements, please submit a pull request or open an issue.

## License:

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT). Feel free to use and modify the code according to the terms of the license.

## Contact:

For any inquiries or support, please feel free to contact me at dzikrurrohmanizrmh@gmail.com
