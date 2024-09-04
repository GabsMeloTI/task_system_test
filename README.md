# Task Management API

## Description

This is the API for the backend of the task management system. The main goal of this API is to provide a robust and efficient solution for managing tasks, users, projects, and other related entities in a task management system.

With this API, you can perform complete CRUD (Create, Read, Update, Delete) operations for the following entities:

- **Users**: Manage the system users, allowing the creation, updating, deletion, and retrieval of information about them.
- **Projects**: Create and manage projects that group tasks and other sections.
- **Sections**: Organize tasks into sections within a project.
- **Tasks**: Add, modify, remove, and query tasks.
- **Subtasks**: Break tasks into subtasks for better organization and tracking.
- **Comments**: Add comments to tasks to collaborate with other users.
- **Tags**: Classify tasks with tags to facilitate organization and search.

## Features

- **Authentication and Authorization**: Support for JWT authentication to ensure that only authenticated users can access and manipulate sensitive data.
- **User Management**: Functions to register new users, log in, and update information.
- **Project and Task Management**: Creation and manipulation of projects and tasks, including the ability to organize tasks into sections and assign tags.
- **Comments and Collaboration**: Add and manage comments on tasks to facilitate collaboration and communication between users.
- **Database Migration**: Automatic database schema setup with support for migration using GORM.

## Technologies Used

- **Golang**: Main language for API development, providing high performance and efficiency.
- **GORM**: ORM (Object-Relational Mapping) for interaction with PostgreSQL database.
- **PostgreSQL**: Relational database management system used to store application data.
- **Gorilla Mux**: HTTP routing framework for Go, used for defining and managing API routes.
- **JWT (JSON Web Tokens)**: Authentication technology used to secure routes and ensure that only authorized users can access certain operations.

## Project Structure

- **`main.go`**: Entry point of the application, where the database setup, migrations, and route configurations are performed.
- **`configs`**: Package responsible for loading and managing application configurations.
- **`db`**: Package for configuring and managing the database connection.
- **`docs`**: Package for configuring and managing Swagger documentation in the project.
- **`models`**: Contains definitions of database entities and their relationships.
- **`repository`**: Package containing the logic for accessing the database.
- **`service`**: Contains the business logic of the application and interacts with the repositories.
- **`controllers`**: Defines the controllers responsible for handling HTTP requests and interacting with services.
- **`routes`**: API route configuration and registration of routes with the appropriate controllers.
- **`utils`**: Defines authentication and data security for the project.

## Entity-Relationship Diagram (ERD)
![Entity-Relationship Diagram](https://github.com/user-attachments/assets/84829514-62bf-4af1-a651-cfdd42dea34a)

## Swagger
Swagger documentation can be found at the following address: [http://localhost:/swagger-ui.html](http://localhost:8000/swagger/index.html#/)
Note: If running the project on a different port, adjust the URL accordingly.
