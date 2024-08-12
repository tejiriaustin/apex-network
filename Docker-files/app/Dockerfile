# Step 1: Build the Go application
FROM golang:1.20-alpine AS build_image

LABEL "Project"="Apex_Network"
LABEL "Author"="Tejiri"
# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY src/go.mod .
COPY src/*.go .
COPY .env .

# Download go dependencies
RUN go mod download

# Fetch the missing module (optional but recommended)
RUN go get github.com/tejiriaustin/apex-network/cmd

# Build the Go application
RUN go build -o /app/apex_network



# Step 2: Create a minimal container to run the application
FROM alpine:latest

# Set a working directory in the new container
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=build_image /apex_network /apex_network

# Copy .env file if it's needed
# COPY .env /app/.env

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD [ "/app/apex_network" ]

