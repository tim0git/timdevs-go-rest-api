# eVe Vehicle API

This repository contains the code for a RESTful API for managing vehicle information.

## API Endpoints

### The Vehicle API exposes the following endpoints:

    GET /health: Returns the status of the service.

    GET /swagger/*any: Serves the Swagger UI for API documentation.

    POST /vehicle: Registers a new vehicle.

    GET /vehicle/:vin: Retrieves a specific vehicle based on its Vehicle Identification Number (VIN).

    PATCH /vehicle/:vin: Updates an existing vehicle based on its VIN.

## Graceful Shutdown

The server gracefully handles shutdown signals (SIGINT and SIGTERM). When the server receives a shutdown signal, it initiates the graceful shutdown process.

During shutdown, the server stops accepting new requests and waits for the ongoing requests to complete. If the process exceeds the predefined timeout period (5 seconds in this case), the server is forced to shut down.

## Dependencies

- Docker 20.10.7
- AWS CLI v2
- Go 1.20.4
- Newman 5.2.2
- Dive 0.10.0
- Make 4.3

## Setup

>Note: Mock values have been used for AWS credentials and DynamoDB endpoint. Whilst running in cloud environment these will be injected via RBAC and therefore do not need to be set.

To run this application locally, follow these steps:

1. Start the application using Docker Compose:
    ```
    make compose_up
    ```
2. Access the API at http://localhost:8443/

# Table commands

To start the DynamoDB container, use the following command:

```
make start_db
```

To delete the DynamoDB container, use the following command:

```
make delete_db
```

To create the Vehicles table in DynamoDB, use the following command:

```
make create_table
```

# Testing commands

To run tests, use the following command:

```
make test
```

To run tests with verbose output, use the following command:

```
make debug
```

To run tests with coverage and generate an HTML report, use the following command:

```
make coverage
```

To run integration tests using Docker Compose, use the following commands:

```
make integration
```

To run Newman tests using the Postman collection, use the following command:

```
make newman
```

# Go binary commands

To run the application locally in development mode, use the following command:

```
make dev
```

To build the Go binary, use the following command:

```
make build
```

To run the Go binary, use the following command:

```
make run
```

# Docker commands

To package the application in a Docker image, use the following command:

```
make package
```

To run the Docker container with the packaged application, use the following command:

```
make run_package
```

To analyze the Docker image using Dive, use the following command:

```
make dive
```

To run the application using Docker Compose, use the following commands:

```
make compose_up
```

To stop and remove the Docker containers created by Docker Compose, use the following command:

```
make compose_down
```

To clean up the environment and remove all generated files and containers, use the following command:

```
make clean
```

# Documentation commands

To generate Swagger documentation, use the following command:

```
make generate_swagger_docs
```

## Swagger Documentation

The API is documented using Swagger. The following are the details of the API:

    Base Path: /
    Host: localhost:8443
    Version: 1
    Description: This is the eVe API for vehicle management

Definitions

    handlers.HealthStatus:
        Properties:
            status: string
        Type: object

    vehicle.Capacity:
        Properties:
            unit: string
            value: integer
        Required: unit, value
        Type: object

    vehicle.Update:
        Properties:
            capacity: reference to vehicle.Capacity
            color: string
            license_plate: string
            manufacturer: string
            model: string
            year: integer
        Required: capacity, color, license_plate, manufacturer, model, year
        Type: object

    vehicle.Vehicle:
        Properties:
            capacity: reference to vehicle.Capacity
            color: string
            license_plate: string
            manufacturer: string
            model: string
            vin: string
            year: integer
        Required: capacity, color, manufacturer, model, vin, year
        Type: object

API Endpoints

    /health:
        Method: GET
        Description: Returns the status of the service
        Consumes: application/json
        Produces: application/json
        Responses:
            200: OK
                Schema: reference to handlers.HealthStatus
        Summary: Health check endpoint
        Tags: health

    /vehicle:
        Method: POST
        Description: Register a new vehicle
        Consumes: application/json
        Parameters:
            request (body): Vehicle information (reference to vehicle.Vehicle)
        Produces: application/json
        Responses:
            201: Created
        Summary: Register a new vehicle
        Tags: vehicle

    /vehicle/{vin}:
        Method: GET
        Description: Retrieve a vehicle
        Consumes: application/json
        Parameters:
            vin (path): Vehicle identification number (string)
        Produces: application/json
        Responses:
            200: OK
                Schema: reference to vehicle.Vehicle
        Summary: Retrieve a vehicle
        Tags: vehicle

    /vehicle/{vin}:
        Method: PATCH
        Description: Update an existing vehicle
        Consumes: application/json
        Parameters:
            vin (path): Vehicle identification number (string)
            request (body): Vehicle information to update (reference to vehicle.Update)
        Produces: application/json
        Responses:
            200: OK
        Summary: Update an existing vehicle
        Tags: vehicle

Note: The API produces responses in application/json format.
Running the API

After following the setup instructions to start the application locally, you can access the API endpoints using the specified base path and host.

## Contributing

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.
## License

This project is licensed under the MIT License.
## Acknowledgments

Special thanks to the authors and contributors of the Gin and Swagger packages for their valuable contributions.
