# 🌟 Golang Fiber & GORM Project 🌟

Welcome to the **Golang Fiber & GORM** project! This application is built with Go, using the powerful Fiber framework for web development and GORM for ORM, making it a robust choice for building RESTful APIs. 🚀

## 📚 Features

- **User Authentication** 🔑: Secure login and user management.
- **Dynamic Book Management** 📚: Create, read, and manage books effortlessly.
- **Post Handling** 📝: Manage user-generated content with ease.

## 🛠️ Technologies Used

- **Go**: The programming language used.
- **Fiber**: Fast web framework for Go.
- **GORM**: ORM library for Golang, simplifies database interactions.
- **MySQL**: Relational database management system for storing data.

## 📄 API Endpoints

- **Login**: `POST /login`
- **User Management**:
  - `GET /user` - Get all users
  - `GET /user/:id` - Get user by ID
  - `POST /user` - Create a new user
  - `PUT /user/:id` - Update user by ID
  - `PUT /user/:id/update-email` - Update user email
  - `DELETE /user/:id` - Delete user by ID
  - `GET /user-post` - Get posts of users
- **Book Management**:
  - `GET /book` - Get all books
  - `POST /book` - Create a new book
- **Post Management**:
  - `GET /post` - Get all posts
  - `POST /post` - Create a new post

## 📦 Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/golang-fiber-gorm.git
   ```
2. Navigate to the project directory:
   ```bash
   cd golang-fiber-gorm
   ```
3. Install the dependencies:
   ```bash
   go mod tidy
   ```
4. Start the server:
   ```bash
   go run main.go
   ```

## ⚡ Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements. 💬

Thank you for checking out the Golang Fiber & GORM project! Happy coding! 🎉
