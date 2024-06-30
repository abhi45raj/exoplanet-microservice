# Exoplanet Service

A Go microservice for managing exoplanets. This service supports adding, retrieving, updating, and deleting exoplanets, as well as estimating fuel requirements for voyages.

## Features

- Add an Exoplanet
- List Exoplanets
- Get Exoplanet by ID
- Update Exoplanet
- Delete Exoplanet
- Fuel Estimation for space voyages

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16+
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerization)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) (for version control)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/your-username/exoplanet-service.git
   cd exoplanet-service
   
### API

### 1. Add an Exoplanet
curl -X POST http://localhost:8080/exoplanets -H "Content-Type: application/json" -d '{
"name": "Earth-like",
"description": "A planet similar to Earth",
"distance": 50,
"radius": 1.1,
"mass": 0.9,
"type": "Terrestrial"
}'

### 2. List Exoplanets
curl -X GET http://localhost:8080/exoplanets

### 3. Get Exoplanet by ID
curl -X GET http://localhost:8080/exoplanets/{id}

### 4. Update Exoplanet
curl -X PUT http://localhost:8080/exoplanets/{id} -H "Content-Type: application/json" -d '{
"name": "Earth-like Updated",
"description": "An updated description",
"distance": 55,
"radius": 1.2,
"mass": 1.0,
"type": "Terrestrial"
}'

### 5. Delete Exoplanet
curl -X DELETE http://localhost:8080/exoplanets/{id}

### 6. Fuel Estimation
curl -X GET "http://localhost:8080/exoplanets/{id}/fuel?crewCapacity=5"



