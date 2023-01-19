## Introducion
Creating RESTful API with [Go](https://go.dev/) and [Gin](https://gin-gonic.com/) to Read and write data in Database.

Using [Postgresql](https://www.postgresql.org/) database to handle all the data.

Using [Postman](https://www.postman.com/) to handle all Api requests.

Create a Server to send Api requests in [main.go](https://github.com/Ume-habiba9/Api/blob/master/main.go)

Function [DBConnect()] in [db.go](https://github.com/Ume-habiba9/Api/blob/master/db/db.go) connects server with Postgres Database.

In [Modules.go](https://github.com/Ume-habiba9/Api/blob/master/Modules/Modules.go) Update Api Responses , by using GET, POST, PUT, DELETE methods.

In [dbfunc.go](https://github.com/Ume-habiba9/Api/blob/master/db/dbfunc.go), update data in database. To update,delete or get single item in database, passed id as a param to only update or get required item.


## Getting Started

 Run the development server:
```go run ./main.go```
 

 Request Url to Send Api request via Postman is "localhost:8080/Cars".

 Check terminal to see the requests response.
