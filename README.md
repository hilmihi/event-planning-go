<!-- ABOUT THE PROJECT -->

## About The Project

GRAPHQL Project for Event Planning

Building the project with layered architecture, and clean code approach for the structure, with the intention of simplicity when the app is scaling up and ease of maintenance

<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

This project structure is built using

- [Golang]
- [Mysql]
- [Labstack/Echo]
- [qraphql]

<p align="right">(<a href="#top">back to top</a>)</p>

### Features

- USERS CRUD
- EVENTS CRUD
- COMMENTS CR
- CATEGORY R

### Folder Structure

```
├── app                             # Main.go
├── addMiddleware/                  # Create middleware
├── config/                         # Configuration to connect to database
├── entities/                       # Create entities for category, comment, event, participant, and user
├── graph/                          # Create schema and resolver for category, comment, event, participant, and user
├── helper/                         # Create request, response, and helper for category, comment, event, participant, and user
├── repository/                     # Get all required data from database for category, comment, event, participant, and user
├── service/                        # Create service for handle the data from repository of category, comment, event, participant, and user

```

<!-- GETTING STARTED -->

## Getting Started

To start project, just clone this repo

### Installation

1. Clone the repo
   ```bash
   git clone https://github.com/hilmihi/event-planning-go.git
   ```
2. Create .env file in main directory
   ```bash
   touch .env
   ```
3. Write the following example environment
   ```
   export DB_CONNECTION_STRING='root:[fillpasswordhere]@/[schema name]?charset=utf8&parseTime=True&loc=Local'
   ```
4. Run the server
   ```bash
   source .env && go run main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>
