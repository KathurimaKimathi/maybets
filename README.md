![Lint and Tests](https://github.com/KathurimaKimathi/maybets/actions/workflows/ci.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/KathurimaKimathi/maybets/badge.svg?branch=main&kill_cache=1)](https://coveralls.io/github/KathurimaKimathi/maybets?branch=main)

## Overview
Maybets is a high-performance Go application designed for processing large volumes of betting transactions in real-time and providing analytical insights. It efficiently ingests transaction data, stores it, and exposes APIs for retrieving key betting analytics.

## Features
- **Data Ingestion**: Accepts betting transactions in JSON format from a file (`bets.json`) or an API endpoint.
- **Processing & Storage**: Utilizes Go's in-memory data structures and SQLite for efficient transaction handling.
- **Analytics APIs**: Provides insights into user betting statistics and detects anomalies.
- **Performance Optimization**: Uses goroutines for concurrent processing and ensures at least 10,000 bets per second throughput.
- **CLI Support**: Includes a command-line interface for batch processing.

## Setup & Installation
### Prerequisites
- Go 1.20+
- Redis (MANDATORY)

### Installation & Running
#### 1. Clone the Repository
```sh
git clone https://github.com/KathurimaKimathi/maybets.git
cd maybets
```
#### 2. Install Dependencies
```sh
go mod tidy
```
#### 3. Start the Redis Server (if caching is enabled)
Ensure Redis is installed and running:
```sh
redis-server
```
#### 4. Run the Server
**Environmental Variables**
```sh
export ENVIRONMENT="LOCAL"
export REDIS_URL="redis://localhost:6379/0"
export JAEGER_ENDPOINT="localhost:4318"
export PORT="8080"
export SQLITE_URL="/path/to/your/sqlite/file"
```
**Method 1: Using CLI**
```sh
cd cmd
go run cmd.go runserver
```
**Method 2: Direct Execution**
```sh
go run server.go
```
#### 5. Generate and Load Test Data
Generate test betting data:
```sh
cd cmd
go run cmd.go generate --betdata 10000 bets.json
```
Load test data into SQLite:
```sh
go run cmd.go process bets.json
```

## API Reference
Each betting transaction follows this JSON structure:
```json
{
  "bet_id": "string",
  "user_id": "string",
  "amount": "float64",
  "odds": "float64",
  "outcome": "win" | "lose",
  "timestamp": "RFC3339 format"
}
```

### Endpoints
#### 1. Get Total Bets by User
```sh
curl --location '<BASEURL>:<PORT>/api/v1/analytics/total_bets?user_id={user_id}'
```
#### 2. Get Total Winnings by User
```sh
curl --location '<BASEURL>:<PORT>/api/v1/analytics/total_winnings?user_id={user_id}'
```
#### 3. Get Top 5 Users by Betting Volume
```sh
curl --location '<BASEURL>:<PORT>/api/v1/analytics/top_users'
```
#### 4. Get Users with Anomalous Betting Activity
```sh
curl --location '<BASEURL>:<PORT>/api/v1/analytics/anomalies'
```

## Contribution
Feel free to submit issues, feature requests, or pull requests to improve the project.

## License
MIT License.
