# MVC Architecture Lecture - Go Implementation

## Steps to Setup

1. Install Go dependencies:

   ```bash
   go mod tidy
   ```

2. `cp .env.sample .env`

3. Set up the MySQL database using the SQL file:

   ```bash
   mysql -u your_username -p < config/db.sql
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on http://localhost:8787
