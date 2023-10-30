# mitrais_test

# prepare docker database
1. Install Docker: If you haven't already, you need to install Docker on your system. You can download and install Docker from the official website: https://docs.docker.com/get-docker/
2. Pull the PostgreSQL Docker Image: Open a terminal and run the following command to pull the official PostgreSQL Docker image. You can specify the version you want, but for this example, we'll use the latest version. Using this command 'docker pull postgres'
3. 'docker run -d --name dbcontainer -p 5432:5432 -e POSTGRES_PASSWORD=pass123 postgres' (to create new database on local container)
4. run docker ps (to make sure the DB container already created)

# Run Golang Apps on local
1. go mod tidy
2. go build
3. go run main.go

# Run with docker
1. Make sure your device already have docker
2. open project from IDE
3. open terminal
4. run docker-compose up --build
5. execute create employee table
6. we can use localhost:8000 to test all the API