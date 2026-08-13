# Task Management API with JWT

## Authentication
- **POST `/register`**: Create a new account. Provide `username`, `password`, and optionally `role` ("admin" or "user").
- **POST `/login`**: Authenticate and receive a JWT token.

## Tasks (Protected Routes - Require Header `Authorization: Bearer <token>`)
- **GET `/tasks`**: Retrieve all tasks. (All authenticated users)
- **GET `/tasks/:id`**: Retrieve a task by ID. (All authenticated users)

## Admin Tasks (Require Role "admin")
- **POST `/tasks`**: Create a new task.
- **PUT `/tasks/:id`**: Update an existing task.
- **DELETE `/tasks/:id`**: Delete a task.
