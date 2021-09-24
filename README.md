# A GO rest api boilerplate with gin framework, authentication, migrations and some sample routes

# Running

**Locally**

go mod install
go run main.go

**Docker**

docker build --tag go-gin-rest-boilerplate .
docker run -it --rm --name go-gin-rest-boilerplate go-gin-rest-boilerplate