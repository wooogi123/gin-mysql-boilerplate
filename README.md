Gin(golang) + Mariadb(mysql) Boilerplate
=======================================

### Installation
```
$ go get github.com/wooogi123/gin-mysql-boilerplate
$ cd $GOPATH/src/github.com/wooogi123/gin-mysql-boilerplate
```

### Running Application
```
$ go run *.go
```

### Build Application
```
$ go build -v
$ ./gin-mysql-boilerplate
```

### Requirement
 - Environment on Linux
 ```
 export DB_USER=YOUR_DB_USERNAME
 export DB_PASSWORD=YOUR_DB_PASSWORD
 export DB_NAME=YOUR_DB_NAME
 export DB_HOST=YOUR_DB_HOST
 export DB_PORT=YOUR_DB_PORT
 ```

 - Database on Linux
 ```
 mysql -uYOUR_DB_USERNAME-p < db/database.sql
 ```
