# 🔐 GoLang Authentication API

A backend authentication API built with **Go (Golang)** using **Gin**, **MongoDB**, and **JWT**.

This project is focused on building a secure and maintainable authentication system while following a modular backend architecture.

## 🚀 Features

* User authentication
* JWT based authentication
* JWT token validation
* Role based claims
* Password/user management structure
* MongoDB database integration
* Authentication middleware
* Role middleware
* Environment based configuration
* Modular project architecture
* HTTP API server
* 
## 🛠️ Tech Stack
* **Go** — Backend programming language
* **Gin** — HTTP web framework
* **MongoDB** — Database
* **JWT** — Authentication and authorization
* **Godotenv** — Environment variable management

## 📁 Project Structure
```textGoLang-Authentication/
│
├── internaal/
│   ├── auth/
│   │   └── jwt.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── db/
│   │   └── mongo.go
│   │
│   ├── httpserver/
│   │   ├── health.go
│   │   └── router.go
│   │
│   ├── middleware/
│   │   ├── auth.go
│   │   └── roles.go
│   │
│   └── user/
│       ├── handler.go
│       ├── model.go
│       ├── repo.go
│       └── services.go
│
├── tmp/
│   ├── api.exe
│   └── build-errors.log
│
├── .air.toml
├── .gitignore
├── go.mod
└── go.sum
```

The repository is organized around application startup, authentication, database access, HTTP handling, middleware, configuration, and user functionality.

## ⚙️ Requirements

Before running the project, make sure you have:

* Go 1.26.5 or a compatible Go version
* MongoDB
* Git

## 📥 Installation

Clone the repository:

```bash
git clone https://github.com/ngwebwork/GoLang-Authentication.git
```

Enter the project:

```bash
cd GoLang-Authentication
```

Download the dependencies:

```bash
go mod download
```

## 🔐 Environment Variables

Create a `.env` file in the project root.

Example:

```env
MONGO_URI=your_mongodb_connection_string
JWT_SECRET=your_super_secret_key
```

> Never commit your `.env` file or expose your JWT secret publicly.

Make sure your `.gitignore` contains:

```gitignore
.env
```

## ▶️ Running the API

Start the server with:

```bash
go run ./cmd/api
```

The API server runs on:

```text
http://localhost:8000
```

The application entry point creates the application, initializes the router, and starts an HTTP server on port `8000`.

## 🔄 Development

If you use Air for live reloading:

```bash
air
```

The repository includes an `.air.toml` configuration file for development.

## 🧱 Architecture

The project separates responsibilities into different packages:

### `cmd/api`

Application entry point.

Responsible for:

* Starting the application
* Creating the router
* Starting the HTTP server
* Handling application shutdown

### `internaal/auth`

Handles authentication logic, including JWT creation and token parsing.

### `internaal/db`

Contains MongoDB database functionality.

### `internaal/middleware`

Contains authentication and role-based middleware.

### `internaal/user`

Contains user-related:

* Models
* Handlers
* Repository logic
* Services

### `internaal/httpserver`

Responsible for HTTP routing and server configuration.

### `internaal/config`

Handles application configuration.

## 🔒 Security

This project uses JWT-based authentication with HS256 signing.

For production:

* Use a strong randomly generated JWT secret.
* Store secrets in environment variables.
* Never commit `.env` files.
* Use HTTPS.
* Validate JWT signing algorithms.
* Implement appropriate token expiration and refresh strategies.
* Validate and sanitize user input.
* Use secure password hashing rather than storing plaintext passwords.

## 🧪 Testing

Run the project's tests with:

```bash
go test ./...
```

Run tests with additional output:

```bash
go test -v ./...
```

## 🔧 Build

Build the API:

```bash
go build -o api ./cmd/api
```

Run the compiled application:

```bash
./api
```

On Windows:

```bash
api.exe
```

## 📌 Current Project Status

This project is actively being developed as a learning and backend engineering project focused on building authentication services with Go.

Future improvements can include:

* Refresh tokens
* Password hashing
* Email verification
* Password reset
* Account verification
* Improved validation
* Automated tests
* API documentation
* Docker support
* Production deployment
* Rate limiting

## 👨‍💻 Author

**ngwebwork**

GitHub:
https://github.com/ngwebwork

## ⭐ Support

If you find this project useful or are following the development journey, consider giving the repository a ⭐ on GitHub.

---

**Built with Go ❤️**
