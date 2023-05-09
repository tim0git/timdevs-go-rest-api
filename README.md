# README

This repository contains the code for a RESTful API for managing vehicle information.

## Dependencies

- Docker
- AWS CLI
- Go

## Setup

To run this application locally, follow these steps:

1. Start the DynamoDB container:
```
make start_db
```

2. Create the Vehicles table in DynamoDB:
```
make create_table
```

3. Build and run the application:
```
make run
```

4. Access the API at http://localhost:8443/

## Testing

To run tests and generate coverage reports, use the following commands:

- Run tests:
```
make test
```

- Run tests with verbose output:
```
make debug
```

- Run tests with coverage and generate HTML report:
```
make coverage
```

## Building and Running the Binary

To build the binary and run it, use the following commands:

- Build the binary:
```
make build
```

- Run the binary:
```
make run_build
```

## Packaging and Running in Docker

To package the application and run it in a Docker container, use the following commands:

- Build the Docker image:
```
make package
```

- Run the Docker container:
```
make run_package
```

## Compose and Dive

To run the application using Docker Compose or analyze the Docker image using Dive, use the following commands:

- Run Docker Compose:
```
make compose
```

- Analyze the Docker image using Dive:
```
make dive
```

## Cleaning Up

To clean up the environment and remove all containers and images, use the following command:

```
make clean
```
