# Personal Blog

A simple blog application built with Go that allows users to create, read, update, and delete blog posts. The application includes an admin panel with authentication to manage content.

This project is part of the [roadmap.sh](https://roadmap.sh/projects/personal-blog)

## Features

- **Public Blog Interface**
  - View all blog posts
  - Read individual blog posts
  
- **Admin Panel**
  - Secure login system
  - Create new blog posts
  - Edit existing blog posts
  - Delete blog posts
  - Session-based authentication

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS
- **Storage**: JSON file-based storage
- **Authentication**: Custom session-based authentication

## Project Structure

```
personal-blog/
├── controllers/       # Request handlers and business logic
├── data/              # Data storage and persistence
├── static/            # Static assets (CSS, JS, images)
│   └── css/
├── templates/         # HTML templates
├── types/             # Type definitions
├── main.go            # Application entry point
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── posts.json         # Data storage file
```

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/personal-blog.git
   cd personal-blog
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the application:
   ```
   go run main.go
   ```

4. Access the application:
   - Blog: http://localhost:8080/
   - Admin panel: http://localhost:8080/admin
   - Login: http://localhost:8080/login

### Admin Credentials

- Username: `admin`
- Password: `password123`

## API Routes

- `GET /` - Home page with all blog posts
- `GET /article/{id}` - View a specific blog post
- `GET /login` - Login page
- `POST /login` - Process login
- `GET /logout` - Logout

### Protected Routes (require authentication)

- `GET /admin` - Admin dashboard
- `GET /new` - Create new blog post form
- `POST /new` - Submit new blog post
- `GET /edit/{id}` - Edit blog post form
- `POST /edit/{id}` - Submit blog post edits
- `GET /delete/{id}` - Delete a blog post

## Security Features

- Session-based authentication
- Secure session ID generation using crypto/rand
- HttpOnly cookies to prevent XSS attacks
- Authentication middleware for protected routes

## Future Improvements

- Add user registration
- Implement password hashing
- Add categories and tags for blog posts
- Add search functionality
- Implement pagination
- Add image upload support
- Switch to a database for storage

## License

This project is licensed under the MIT License - see the LICENSE file for details.
