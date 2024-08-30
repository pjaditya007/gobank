## Sample Bank - Golang API

A simple bank API built with Golang.

Features,
User account creation,
Login with account number and password,
Account retrieval,
JWT-based authentication

Prerequisites
Go (version 1.18 or later)
PostgreSQL

go mod download

Set environment variable:

Set the JWT_SECRET environment variable with a strong secret key for JWT signing.

export JWT_SECRET=your_secret_key

Create a PostgreSQL database:

Create a PostgreSQL database and user with appropriate permissions.

Running the application
Configure database connection:

Update the connStr variable in storage.go with your PostgreSQL connection details.

Run the application:

go run main.go

API Endpoints
/login (POST): Login with account number and password. Returns a JWT token on success.
/account (GET): Retrieve all accounts (Admin Only). (POST): Create a new account.
/account/{id} (GET): Retrieve account details by ID (requires JWT token). (DELETE): Delete account by ID (requires JWT token).
/transfer (POST): Transfer funds between accounts (not implemented yet).
Note: All endpoints requiring JWT token need to include the token in the header x-jwt-token.
