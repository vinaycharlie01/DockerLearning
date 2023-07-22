# Use a lightweight base image
FROM golang:alpine AS build

# Set the working directory in the image
WORKDIR /app

# Copy the Go source code to the image
COPY . .

# Build the Go application
RUN go build -o main /app/go/*.go

# Use a smaller image for the final image
FROM alpine

# or
# # Use a minimal base image
# FROM scratch


# # Install the necessary packages
# RUN apk add --no-cache ca-certificates



# Copy the compiled Go binary from the build stage
COPY --from=build /app/main /app/main

# Set the command to run the Go binary
CMD ["/app/main"]
