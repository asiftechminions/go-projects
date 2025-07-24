Introduces a Dockerfile to containerize a Go web application.

Uses a multi-stage build to compile the application with the Golang image and a lightweight Alpine image for deployment. Sets up a basic HTTP server that responds with a greeting and runs on port 8080.

Includes initial Go module setup and main application logic.
