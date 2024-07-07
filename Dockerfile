# Base image with Go and common dependencies
FROM golang:1.22 AS base
WORKDIR /app

ENV HOST=0.0.0.0
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies for Debian-based image
RUN apt-get update \
  && apt-get install -y \
  ca-certificates \
  git \
  wget \
  bash \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Install Air for hot reloading and Delve for debugging
RUN go install github.com/air-verse/air@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest

# Expose ports for Air and Delve
EXPOSE 8080
EXPOSE 2345  

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Ensure Air is available in the PATH
ENV PATH="/go/bin:${PATH}"

# Define the command to run Air
ENTRYPOINT ["air"]

# Ensure that Air is the default command
CMD ["air"]
