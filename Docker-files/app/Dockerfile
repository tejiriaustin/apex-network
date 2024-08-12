# Step 1: Build the Go application
FROM golang:1.20-alpine AS BUILD_IMAGE

LABEL "Project"="Apex_Network"
LABEL "Author"="Tejiri"
# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY ../src/go.mod .
COPY *.go ../src/
COPY .env ../src/

# Download go dependencies
RUN go mod download

# Build the Go application
RUN go build -o bin /apex_network



# # Step 2: Create a minimal container to run the application
# FROM alpine:3.18

# # Copy the built Go binary from the builder stage
# COPY --from=BUILD_IMAGE /apex_network /apex_network

# # Copy .env file if it's needed
# COPY .env /app/.env

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD [ "/apex_network apex_network_api" ]

