#####################
# DEVELOPMENT IMAGE #
#####################

# Pull base image
FROM golang:1.14 AS builder

# Set environment variables
# none here

# Set work directory
WORKDIR /code

# Copy project
COPY ./ /code/

# Build main.go
RUN go build main.go

# Expose the application on port 8080
EXPOSE 8080

# Run starting command
CMD ["./main"]

#####################
# PRODUCTION IMAGE #
#####################

# Pull base image
FROM golang:1.14 AS prod

# Set environment variables
# env variables should usually come from the deployment process

# Set work directory
WORKDIR /code

# Copy executable file
COPY --from=builder /code/main /code/

# Expose the application on port 8080
EXPOSE 8080

# Run starting command
CMD ["./main"]
