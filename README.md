## Project Management API Server

RESTful API Server built with Go for managing projects and tasks. It features user authentication using JWT tokens, integration with a MySQL database and also includes Docker support for containerization. 

### Getting Started

1. Clone this repository:
    ```sh
    git clone https://github.com/neerrrajj/go-rest-api.git
    cd go-rest-api
    ```

2. Set up your MySQL database and update the configuration in `config.go`

3. Install the dependencies:
    ```sh
    go mod download
    ```

4. Build and run the project:
    ```sh
    go run main.go
    ```

### Docker support

1. Build the docker image:
    ```sh
    docker build -t go-rest-api .
    ``` 

- Use Docker Compose to run a MySQL Server and the API Server or follow the below steps to run the containers separately and link them.  

2. Run MySQL container:
    ```sh
    docker run --name project-manager-db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=project-manager -p 3306:3306 -d mysql
    ```

3. Run the API Server container:
    ```sh
    docker run --name go-rest-api --link project-manager-db:db -e DB_HOST=db -e DB_PORT=3306 -e DB_USER=root -e DB_PASSWORD=password -e DB_NAME=project-manager -p 8080:8080 go-rest-api
    ```
