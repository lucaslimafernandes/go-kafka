# Credit Card Transaction Simulation with Kafka and PostgreSQL in Go

This project simulates credit card transactions using Kafka as the messaging middleware and PostgreSQL as the database. The system generates financial transactions, sends them via Kafka, and the Kafka consumer receives and validates these transactions by checking the users' balance in the PostgreSQL database.

## Features

- **Transaction Simulation**: Generate random credit card transactions.
- **Kafka Messaging**: Transactions are sent to a specific Kafka topic.
- **Balance Validation**: The Kafka consumer reads transactions and checks the PostgreSQL database to verify if the user has sufficient balance.
- **Balance Update**: After validation, the user's balance is updated in the database.

## Architecture

- **Go**: The main programming language for generating transactions and managing the Kafka and PostgreSQL flow.
- **Kafka**: Message broker used to transmit transactions.
- **PostgreSQL**: Database used to store user balances and transaction logs.
- **pgx**: Go PostgreSQL driver used for database communication.
- **confluent-kafka-go**: Kafka library for Go used in this project.

## Requirements

- **Go** (1.16+)
- **Docker** (optional, for running Kafka and PostgreSQL)
- **Kafka** (broker and zookeeper)
- **PostgreSQL**

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   ```

2. Install necessary dependencies:
   ```bash
   go mod tidy
   ```

3. Set up and start the required services (Kafka and PostgreSQL). Example using Docker:
   ```bash
    sudo docker run --name pg_kafka -d -e POSTGRES_PASSWORD=password -v pg_vol:/var/lib/postgresql/data -p 5432:5432 postgres:14-alpine

    sudo docker run --name broker -d -p 9092:9092 apache/kafka:latest
   ```

4. Create the necessary database and tables:
   The application automatically creates the tables on startup if `Migrate()` is configured correctly.

## How to Use

### Available Flags

- `--new_users`: Generates and inserts new fictitious users into the database.
- `--run`: Starts the credit card transaction simulation and processes these transactions via Kafka.

### Example Usage

1. To create new users in the database:
   ```bash
   go run main.go --new_users
   ```

2. To start generating transactions and the Kafka consumer:
   ```bash
   go run main.go --run
   ```

## File Structure

- `main.go`: Main entry point that starts the program. Configures Kafka consumers and producers and manages the transaction simulation.
- `models/`: Directory containing logic for PostgreSQL connection, such as `ConnectDB()` and `Migrate()`.
- `pay_utils/`: Directory containing functions related to transaction generation and processing, like `PassingCards()` and `Validat()`.
- `consume.go`: Kafka consumer that reads transaction messages and validates the user's balance in PostgreSQL.
- `produce.go`: Kafka producer that generates transactions and sends them to the Kafka topic.
- `insert_users.go`: Script for inserting new fictitious users into the database.
- `sells.go`: Contains the sales structure and functions related to sales transactions.
- `faker.go`: Generates fictitious data using the `Faker` library to simulate credit card transactions.

## Workflow

1. **User Creation**: The system generates a set of fictitious users with random balances.
2. **Transaction Simulation**: Sales transactions are generated based on the created users.
3. **Kafka Messaging**: These transactions are sent to a specific Kafka topic.
4. **Transaction Consumption**: The Kafka consumer receives the transactions, checks the user’s balance in PostgreSQL, and validates if there are enough funds to complete the purchase.
5. **Balance Update**: If the user has sufficient balance, the balance is updated in the database, and the transaction is logged.

## Technologies Used

- **Go**: Programming language used to build the application.
- **Kafka**: Messaging middleware for asynchronous data transmission.
- **PostgreSQL**: Relational database used to store user and transaction information.
- **pgx**: PostgreSQL library for communication with the database.
- **confluent-kafka-go**: Go library for Kafka integration.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to your branch (`git push origin feature/new-feature`).
5. Open a Pull Request.
