# Step 1: Build the Go application
FROM golang:1.20-alpine AS BUILD_IMAGE

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY ../src/go.mod .
COPY ../src/main.go .

# Build the Go application
# RUN go mod init example.com/mymodule
RUN go build -o bin .

ENTRYPOINT [ "/app/bin"] 

# # Step 2: Create a lightweight image to run the application
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /root/

# # Copy the compiled Go binary from the builder stage
# COPY --from=BUILD_IMAGE /app/main .

# # Expose the port the application will run on
# EXPOSE 3000

# # Command to run the application
# CMD ["make", "app"]