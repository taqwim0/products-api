## Products API Application

### Overview
This repository contains a simple Products API built with Go and PostgreSQL. The application is containerized using Docker and can be easily set up and run locally using Docker Compose.

### Prerequisites
* Go (Golang) 1.23 or later
* Docker and Docker Compose (for running PostgreSQL database)
* Git (for version control)

### Setup
#### 1. Clone the Repository
```
git clone https://github.com/taqwim0/products-api
cd products-api
```

#### 2. Install Dependencies
```
go mod tidy
```

#### 3. Configure Environment Variables
Create an `app.env` file in the repository and set the following environment variables:
```
DB_HOST=localhost
DB_USER=userdb
DB_PASSWORD=passworddb
DB_NAME=products_db
DB_PORT=5432
```

#### 4. Run PostgreSQL with Docker
Run the following command to start the database:
```
docker-compose up -d
```

Open postgreSQL application (DBeaver / pgAdmin) to establish database connection configuration on your local for table creation & data samples insertion. You can check the SQL queries on:
```
files/table_creation.sql (for table creation)
files/sample_products.sql (for data samples insertion)
```

#### 5. Run the Backend Application with Docker
```
docker-compose up --build
``` 

### Running the Application 
Ensure the backend is running on `http://localhost:8080`. You can check & import the API list on your local Postman application from postman collection file located at
```
files/products-api.postman_collection.json
```

### API Documentation

#### Get All Products

```http
  GET localhost:8080/products
```

#### Get Product by ID

```http
  GET localhost:8080/product/detail/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of product to fetch. |

#### Create Product

```http
  POST localhost:8080/product/add/

  Request Body JSON Sample:
  {
    "name": "Wireless Charger Electric",
    "description": "Fast wireless charger for smartphones.",
    "price": 200000,
    "variety": "Accessories",
    "rating": 4.5,
    "stock": 200
   }

```
| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Product name. |
| `description`      | `string` | **Required**. Product description. |
| `price`      | `float` | **Required**. Product price. |
| `variety`      | `string` | **Required**. Product variety. |
| `rating`      | `float` | **Required**. Product rating. |
| `stock`      | `int` | **Required**. Product stock. |

#### Update Product

```http
  PUT localhost:8080/product/update/{id}

  Request Body JSON Sample:
  {
    "name": "Wireless Charger Electric for All",
    "description": "Fast wireless charger for smartphones.",
    "price": 200000,
    "variety": "Accessories",
    "rating": 4.5,
    "stock": 200
   }

```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of product to update. |

| Body | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Product name. |
| `description`      | `string` | **Required**. Product description. |
| `price`      | `float` | **Required**. Product price. |
| `variety`      | `string` | **Required**. Product variety. |
| `rating`      | `float` | **Required**. Product rating. |
| `stock`      | `int` | **Required**. Product stock. |

#### Delete Product

```http
  DELETE localhost:8080/product/delete/{id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. ID of product to delete. |


### Troubleshooting
* **Database Connection Errors**: Ensure that the database credentials in docker-compose.yaml match those set in PostgreSQL.
* **Port Conflicts**: Make sure no other service is using ports 5432 (PostgreSQL) or 8080 (API) on your machine.