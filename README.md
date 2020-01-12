Gin(golang) + Mariadb(mysql) Boilerplate
=======================================

### Installation
```
$ git clone https://github.com/wooogi123/gin-mysql-boilerplate.git
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
 export DB_USER=__YOUR_DB_USERNAME__
 export DB_PASSWORD=`<strong>YOUR_DB_PASSWORD</strong>`
 export DB_NAME=**YOUR_DB_NAME**
 export DB_HOST=**YOUR_DB_HOST**
 export DB_PORT=**YOUR_DB_PORT**
 ```

 - Database on Linux
 ```
 mysql -u**YOUR_DB_USERNAME**-p < db/database.sql
 ```
