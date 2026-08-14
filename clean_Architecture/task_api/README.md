# Task Management API: A Beginner's Guide to Clean Architecture

Welcome to the Task Management API! This project is built using a design pattern called **Clean Architecture**. If you're a beginner, the codebase might look like it has a lot of folders, but don't worry! This guide will explain exactly how everything works together in plain English.

---

## 1. What is Clean Architecture?

Imagine building a house. You want the foundation and the walls (the core) to be solid, no matter what kind of paint or furniture (the outside) you use. 

Clean Architecture is like an onion with different layers:
- **Inner Layers (The Core):** These are the pure rules of your application (like what a "Task" or a "User" is). They know **nothing** about databases, web frameworks, or the internet.
- **Outer Layers (The Outside World):** These handle things like saving to MongoDB or responding to HTTP web requests. 

Because the inside doesn't know about the outside, you can easily swap out your database or web framework in the future without breaking the core of your app!

---

## 2. The Flow: Journey of a Web Request

What happens when a user wants to create a new task? Here is the step-by-step journey of that request through our folders:

1. **The User** sends a request (e.g., "Create a task for me") to a specific URL.
2. **Delivery/routers** acts like a map. It sees the URL and says, "Ah, this goes to the Task Controller."
3. **Delivery/controllers** acts like a waiter in a restaurant. It takes the user's JSON data, checks if it looks okay, and hands it to the Usecase (the kitchen).
4. **Usecases** is the chef. It does the actual "business logic." It adds a `CreatedAt` timestamp, sets the status to `pending`, and tells the Repository to save it.
5. **Repositories** is the storage room. It knows exactly how to talk to MongoDB. It saves the task to the database and says "Done!" back to the Usecase.
6. **Usecases** tells the Controller, "The food is ready!"
7. **Delivery/controllers** hands the finished task back to the User as a JSON response.

---

## 3. Explaining Every Folder and File

Here is a breakdown of every piece of the puzzle:

### 📁 Domain (`Domain/`)
**The Blueprint of the Application.**
This folder is the absolute center of the onion. It doesn't know about MongoDB, web servers, or security. It just defines the shapes of our data and the "contracts" (Interfaces) for the rest of the app.
*   **`domain.go`**: Contains the structs for `User` and `Task`. It also contains **Interfaces** (like `TaskRepository` or `TaskUsecase`). Interfaces are simply lists of promises saying, "Somewhere in this app, there will be code to Create, Delete, and Update Tasks."

### 📁 Usecases (`Usecases/`)
**The Business Rules (The Brain).**
This layer contains the actual logic of the app.
*   **`task_usecases.go`**: Handles task logic. When creating a task, it sets the creation date. When getting a task, it asks the repository to fetch it.
*   **`user_usecases.go`**: Handles user logic. It checks if a username already exists during registration, and verifies passwords during login.

### 📁 Repositories (`Repositories/`)
**The Database Translators.**
This is the only folder that knows we are using MongoDB. If we ever switch to PostgreSQL, we only have to change files in this folder!
*   **`task_repository.go`**: Takes the Go `Task` struct and inserts, updates, deletes, or reads it from the MongoDB database.
*   **`user_repository.go`**: Does the same thing, but for `User` data.

### 📁 Infrastructure (`Infrastructure/`)
**The External Tools and Helpers.**
This folder holds external technical details like security and passwords.
*   **`password_service.go`**: Takes a plain-text password (like "password123") and scrambles (hashes) it so hackers can't read it.
*   **`jwt_service.go`**: Generates and checks "JSON Web Tokens" (JWT). Think of a JWT as a digital ID card or wristband given to a user after they log in.
*   **`auth_middleWare.go`**: The bouncer at the club. Before a user can visit a protected URL (like creating a task), this file checks if they have a valid JWT "wristband". If not, it rejects them.

### 📁 Delivery (`Delivery/`)
**The Front Desk.**
This folder handles the web server, HTTP requests, and connecting all the pieces together.
*   **`routers/router.go`**: The directory. It maps URLs like `/login` or `/tasks` to specific controller functions.
*   **`controllers/controller.go`**: The translator. It translates raw internet HTTP requests into Go data, passes it to the Usecases, and then translates the result back into an HTTP response (like a `200 OK` or `404 Not Found`).
*   **`main.go`**: The starting engine. This file connects to the MongoDB database, creates all the repositories, usecases, and controllers, wires them together, and starts the web server on port 8080.

### 📄 `app`
**The Compiled Executable.**
This is a binary executable file generated when you ran the `go build -o app Delivery/main.go` command. It takes all the separate `.go` files mentioned above, translates them into machine code that your computer understands, and bundles them into one single file. You can run the entire server just by executing `./app` in your terminal!

---

## 4. Summary of the Flow (from what to what)

1. User sends an HTTP request -> **`router.go`**
2. Router routes request to -> **`controller.go`** (Checks authorization via **`auth_middleWare.go`**)
3. Controller parses JSON and calls -> **`task_usecases.go`** (or `user_usecases.go`)
4. Usecase applies business logic and calls -> **`task_repository.go`** (or `user_repository.go`)
5. Repository writes/reads from -> **MongoDB**
6. Repository returns data back to -> **Usecase**
7. Usecase returns data back to -> **Controller**
8. Controller sends JSON response back to -> **User**

This strict one-way street ensures your code is extremely easy to test, read, and maintain!
