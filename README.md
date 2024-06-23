# Gin-JWT implementation 🚀

This is a user management API built with Go. It provides endpoints for registering, logging in, managing and logging out users.

## Features💎

- User registration: Allows users to create an account with a username, email, and password.
- User login: Provides a way for users to log in to the API.
- User retrieval: Allows retrieving information about a specific user.
- User deletion: Provides a way to delete a user.
- User logout: Provides a way for users to log out of the API and blacklist the token.

## Packages Used🧩

- **Gin** (Web Framework): [Gin](https://github.com/gin-gonic/gin) is a high-performance web framework for Go that provides routing, middleware, and more to build web applications. 
- **GORM** (ORM): [GORM](https://github.com/go-gorm/gorm) is an Object-Relational Mapping library for Go that simplifies database interactions by providing a higher-level abstraction over SQL queries.
- **golang-jwt v5** (JWT implementation): [golang-jwt](https://github.com/golang-jwt/jwt) v5 is a module for implementing JSON Web Tokens (JWT) in Go, providing functionality for creating, validating, and parsing JWTs.

## Usage💻 

1. Clone the repository:

   ```sh
   git clone https://github.com/oscar503sv/gin-jwt.git 
   ```

2. Install dependencies:

   ```sh
   go mod download
   ```

3. Set up the database:

- Create a SQLite database file named `user.db`. You can configure any database engine in the `database.go` file.
- Run the `config.ConnectDB()` function to connect to the database.

4. Run the API:

   ```sh
   go run main.go
   ```

5. Make API requests:

- Use a tool like cURL or Postman to send HTTP requests to the API endpoints.
- The API endpoints are as follows:
  - Register a user: `POST /register`
  - Login: `POST /login`
  - Logout: `POST /logout`
  - Retrieve all users: `GET /users`
  - Retrieve a user: `GET /users/{id}`
  - Update a user: `PUT /users/{id}`
  - Delete a user: `DELETE /users/{id}`

## Contributing🤝

- If you want to contribute a new feature or improve the current implementation, please open an issue first.

## License⚖️

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Website😊

Oscar's [website](https://ozkar.codes).