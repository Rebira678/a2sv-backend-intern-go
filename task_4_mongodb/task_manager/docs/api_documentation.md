# Task Management REST API - MongoDB Integration

## Overview
This API provides task management capabilities backed by MongoDB for persistent data storage.

## Configuration
* **Default Connection URI**: `mongodb://rebik:246810Re%3B%3A%3A%3A@localhost:27017/?authSource=admin`
* **Custom Connection URI**: Set the `MONGO_URI` environment variable before launching:
  ```bash
  export MONGO_URI="mongodb://username:password@localhost:27017/?authSource=admin"
  ```

## API Endpoints
* **GET** `/tasks` - Retrieve all tasks
* **GET** `/tasks/:id` - Retrieve a specific task by MongoDB Hex ID
* **POST** `/tasks` - Create a new task
* **PUT** `/tasks/:id` - Update an existing task
* **DELETE** `/tasks/:id` - Delete a task
