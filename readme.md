# Event Manager REST API

This is a simple REST API for managing events that I created as a learning project while studying Golang. The API allows users to:

- Sign up and log in with JWT-based authentication.
- Create, update, and delete events.
- Register for events and cancel registrations.
- View all events or fetch details for a specific event.

## Features

- User authentication with JWT
- CRUD operations for events
- Event registration and cancellation
- Database integration using SQL

## Tech Stack

- **Programming Language**: Go
- **Framework**: Gin
- **Database**: SQLite
- **Authentication**: JWT

## Getting Started

### Prerequisites

- Go installed (version >= 1.18)
- SQLite installed (optional; the database will be auto-created)
- Git installed

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/samankeyhan/event-manager-api.git
   cd event-manager-api
   ```
2. Install dependencies:
   go mod tidy

3. Run the application:
   go run main.go

Access the API at http://localhost:8080.

#### 5. **API Endpoints**

```markdown
## API Endpoints

### Public Endpoints

- `POST /signup`: Create a new user.
- `POST /login`: Log in and get a JWT token.

### Protected Endpoints

- `GET /events`: Get all events.
- `GET /events/:id`: Get details of a specific event.
- `POST /events`: Create a new event.
- `PUT /events/:id`: Update an event.
- `DELETE /events/:id`: Delete an event.
- `POST /events/:id/register`: Register for an event.
- `DELETE /events/:id/register`: Cancel registration for an event.
```

## Future Improvements

- Add more advanced features like pagination, filtering, and search.
- Implement role-based access control.
- Use an ORM like GORM for database operations.
- Enhance error handling and validation.
- and so on...
