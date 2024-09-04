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

## Getting Started

To get started with the Task Management API, follow these steps:

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/your-repository/task-management-api.git

2. **Navigate to the Project Directory**
   ```bash
   cd task-management-api

3. **Install Dependencies**
   ```bash
   go mod tidy

4. **Run Docker**
   ```bash
   sudo docker-compose up

5. **Run the Application**
   ```bash
   go run main.go

6. **Access the Swagger Documentation**
   ```bash
   [sudo docker-compose up](http://localhost:8000/swagger/index.html
)



# API Endpoints Documentation

This documentation provides an overview of the API endpoints available in the system.

| Resource | Route              | Method | Description                                                                 |
|----------|--------------------|--------|-----------------------------------------------------------------------------|
| Comment  | `/comment`          | GET    | Retrieves all comments available in the system.                             |
| Comment  | `/comment/{id}`     | GET    | Retrieves a specific comment by its ID.                                     |
| Comment  | `/comment`          | POST   | Creates a new comment with the provided details, including optional image upload. |
| Comment  | `/comment/{id}`     | PUT    | Updates an existing comment's content and optional image URL by its ID.     |
| Comment  | `/comment/{id}`     | DELETE | Deletes an existing comment from the system by its ID.                      |
| User     | `/user`             | GET    | Retrieves a list of all users.                                              |
| User     | `/user/{id}`        | GET    | Retrieves a specific user by their ID.                                      |
| User     | `/user`             | POST   | Creates a new user with the provided details.                               |
| User     | `/user/{id}`        | PUT    | Updates an existing user's details by their ID.                             |
| User     | `/user/{id}`        | DELETE | Deletes an existing user from the system by their ID.                       |
| Project  | `/project`          | GET    | Retrieves a list of all projects.                                           |
| Project  | `/project/{id}`     | GET    | Retrieves a specific project by its ID.                                     |
| Project  | `/project`          | POST   | Creates a new project with the provided details.                            |
| Project  | `/project/{id}`     | PUT    | Updates an existing project's details by its ID.                            |
| Project  | `/project/{id}`     | DELETE | Deletes an existing project from the system by its ID.                      |
| Section  | `/section`          | GET    | Retrieves a list of all sections.                                           |
| Section  | `/section/{id}`     | GET    | Retrieves a specific section by its ID.                                     |
| Section  | `/section`          | POST   | Creates a new section with the provided details.                            |
| Section  | `/section/{id}`     | PUT    | Updates an existing section's details by its ID.                            |
| Section  | `/section/{id}`     | DELETE | Deletes an existing section from the system by its ID.                      |
| Task     | `/task`             | GET    | Retrieves a list of all tasks.                                              |
| Task     | `/task/{id}`        | GET    | Retrieves a specific task by its ID.                                        |
| Task     | `/task`             | POST   | Creates a new task with the provided details.                               |
| Task     | `/task/{id}`        | PUT    | Updates an existing task's details by its ID.                               |
| Task     | `/task/{id}`        | DELETE | Deletes an existing task from the system by its ID.                         |
| Task     | `/task/{task_id}/label` | POST   | Assigns labels to a specific task by its ID.                               |
| Subtask  | `/subtask`          | GET    | Retrieves a list of all subtasks.                                           |
| Subtask  | `/subtask/{id}`     | GET    | Retrieves a specific subtask by its ID.                                     |
| Subtask  | `/subtask`          | POST   | Creates a new subtask with the provided details.                            |
| Subtask  | `/subtask/{id}`     | PUT    | Updates an existing subtask's details by its ID.                            |
| Subtask  | `/subtask/{id}`     | DELETE | Deletes an existing subtask from the system by its ID.                      |
| Label    | `/label`            | GET    | Retrieves a list of all labels.                                             |
| Label    | `/label/{id}`       | GET    | Retrieves a specific label by its ID.                                       |
| Label    | `/label`            | POST   | Creates a new label with the provided details.                              |
| Label    | `/label/{id}`       | PUT    | Updates an existing label's details by its ID.                              |
| Label    | `/label/{id}`       | DELETE | Deletes an existing label from the system by its ID.                        |
