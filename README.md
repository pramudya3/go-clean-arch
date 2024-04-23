# Go Clean Architecture Project

This project follows the principles of Clean Architecture in Go, separating concerns into layers and promoting testability, maintainability, and scalability.

## Folder Structure

```
- cmd/
  - main.go
- config/
  - config.go
  - server.go
  - app.go
- domain/
  - user.go
  - ...
- internal/
  - api/
    - controller/
      - user.go
      - healthcheck.go
    - route/
      - routes.go
      - user.go
      -healthcheck.go
    - middleware/
      - authentication.go
  - usecase/
    - user.go
  - repository/
    - user.go
- pkg/
  - mongodb/
    - mongo.go
```

## API Specifications

1. **Health Check**
   - Endpoint: `/api/v1/healthcheck`
   - Method: GET
   - Response: HTTP 200 OK

2. **Signup**
   - Endpoint: `/api/v1/users/signup`
   - Method: POST
   - Request Body:
     ```json
     {
         "name": "example",
         "email": "example@example.com",
         "password": "password123"
     }
     ```
   - Response Body:
     ```json
     {
         "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
         "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
     }
     ```

3. **Signin**
   - Endpoint: `/api/v1/users/signin`
   - Method: POST
   - Request Body:
     ```json
     {
         "email": "example@example.com",
         "password": "password123"
     }
     ```
   - Response Body:
     ```json
     {
         "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
         "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
     }
     ```

4. **Profile**
   - Endpoint: `/api/v1/users/profile`
   - Method: GET
   - Request Headers:
     - Authorization: Bearer \<access_token\>
   - Response Body:
     ```json
     {
         "username": "example",
         "email": "example@example.com"
     }
     ```

## TODO

1. Add controller for signin and refresh token.
2. Implement Role-Based Access Control (RBAC).
3. Implement other necessary functionalities.

---
