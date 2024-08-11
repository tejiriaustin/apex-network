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
# RUN go mod init example.com/mymodule
RUN go build -o bin /apex_network

EXPOSE 3000

CMD [ "/apex_network apex_network_api" ]

# ENTRYPOINT [ "/app/bin"]

# # Set the working directory inside the container
# WORKDIR /root/

# # Copy the compiled Go binary from the builder stage
# COPY --from=BUILD_IMAGE /app/main .

# # Expose the port the application will run on
# EXPOSE 3000

# # Command to run the application
#  CMD ["make", "app"]
