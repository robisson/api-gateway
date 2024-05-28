# Open Source API Gateway

## Objective
The Open Source API Gateway is designed to provide a simple and efficient way to manage and route API requests to various backend services. It uses the OpenAPI 3.0 specification for dynamic configuration of endpoints, HTTP methods, and backend destinations. This project aims to facilitate easy API management with minimal configuration and high performance.

## Tech Stack
- **Language**: Go (Golang) 1.20
- **Containers**: Docker, Docker Compose
- **Tools**: Alpine Linux, Git
- **Testing**: Testify

## Getting Started

### Prerequisites
Ensure you have the following installed:
- Docker
- Docker Compose

### Running Locally

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/open-source-api-gateway.git
   cd open-source-api-gateway

2. **Build containers:**
   ```sh
   docker-compose up -d

3. **Run units tests:**
    ```sh
    go test ./... -v 
    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html 

4. **How to Use openapi.yaml to Configure API Gateway:**
    ```yaml
    openapi: 3.0.0
    info:
    title: Payment API
    version: 1.0.0
    description: API to manage payments

    paths:
    /payments:
        post:
        summary: Create a new payment
        operationId: createPayment
        x-backend-url: "http://localhost:8081"
        requestBody:
            description: Data to create a new payment
            required: true
            content:
            application/json:
                schema:
                type: object
                properties:
                    account_id:
                    type: string
                    description: Account ID
                    example: "123456"
                    amount:
                    type: number
                    format: float
                    description: Payment amount
                    example: 100.50
        responses:
            '201':
            description: Payment created successfully
            content:
                application/json:
                schema:
                    type: object
                    properties:
                    payment_id:
                        type: string
                        format: uuid
                        description: Payment ID
                        example: "550e8400-e29b-41d4-a716-446655440000"
                    amount:
                        type: number
                        format: float
                        description: Payment amount
                        example: 100.50
                    created_at:
                        type: string
                        format: date-time
                        description: Payment creation date
                        example: "2024-05-27T14:53:00Z"
                    created_by:
                        type: string
                        description: Account ID that created the payment
                        example: "123456"
                    status:
                        type: string
                        enum: [Authorized, Unauthorized]
                        description: Payment status
                        example: "Authorized"

    /payments/{payment_id}:
        get:
        summary: Get payment details
        operationId: getPayment
        x-backend-url: "http://localhost:8081"
        parameters:
            - name: payment_id
            in: path
            required: true
            description: Payment ID
            schema:
                type: string
                format: uuid
                example: "550e8400-e29b-41d4-a716-446655440000"
        responses:
            '200':
            description: Payment details
            content:
                application/json:
                schema:
                    type: object
                    properties:
                    payment_id:
                        type: string
                        format: uuid
                        description: Payment ID
                        example: "550e8400-e29b-41d4-a716-446655440000"
                    amount:
                        type: number
                        format: float
                        description: Payment amount
                        example: 100.50
                    created_at:
                        type: string
                        format: date-time
                        description: Payment creation date
                        example: "2024-05-27T14:53:00Z"
                    created_by:
                        type: string
                        description: Account ID that created the payment
                        example: "123456"
                    status:
                        type: string
                        enum: [Authorized, Unauthorized]
                        description: Payment status
                        example: "Authorized"
            '404':
            description: Payment not found

