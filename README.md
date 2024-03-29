# gRPC application

## Goals

This project is a gRPC-based TODO application that allows users to manage their tasks. It supports CRUD operations (Create, Read, Update, Delete) for both users and tasks entities using the gRPC framework.

## Application assumptions

1. User Management
    - The application provides gRPC service methods to create, retrieve, update, and delete user profiles.
    - Users is able to provide details such as their name, email, and any other relevant information through gRPC requests.
2. Task Management
    - The application provides a gRPC service methods to create, retrieve, update, and delete tasks.
    - Tasks have properties such as a title, description, due date, status.
3. User-Task Relationship
    - Each task is associated with a specific user.
    - Users are able to view their own tasks as well as tasks assigned to other users.
4. gRPC Service Definition
    - The application define a gRPC service with appropriate methods for user and task management.
    - The service definition include the necessary message types for user profiles and tasks.
5. Error Handling
    - The application handle errors gracefully and provide informative error responses to gRPC clients.
    - It validates user inputs, ensuring they meet the required format and constraints.
6. Data Storage
    - The application persist user and task data in a file.
    - Users and tasks are stored separately and linked appropriately.
7. Response Formats
    - The gRPC service use protocol buffers for defining the request and response message formats.
    - Messages should be defined using protocol buffer language syntax.
8. gRPC Client
    - A gRPC client is implemented to interact with the gRPC service.
    - The client is separate application a separate module within the same application.
