# MyShop-Webservice
**A workshop from "The World of Back-end Development" short course by CS-Infinite KMITL
about Go web service using Gin framework.**

## Installation
※ In this project I use MySQL
```sh
$ go get github.com/gin-gonic/gin
$ go get github.com/spf13/viper
$ go get github.com/go-sql-driver/mysql
```
## Try it out
※ Clone this project
####
※ Create database and table `products` and `product_types`
```
CREATE TABLE product_types (
    product_type_id     INT PRIMARY KEY,
    product_type_name   VARCHAR(50)
);
```
```
CREATE TABLE products (
    product_id          INT PRIMARY KEY,
    product_name        VARCHAR(50),
    product_quantity    INT,
    product_price       DECIMAL(10,4),
    product_product_type_id INT REFERENCES product_types
);
```
※ Create a file `development.yaml` at the project root path, inside is config about your database and server
*example*
```
database:
  user: dbuser
  password: 12345678
  host: localhost
  port: 3306
  initdb: example

server:
  port: 4000
```
※ Run the project using
```sh
go run main.go
```
※ Go to `localhost:{port}/`
※ Services in this project
**Ping**
```
GET     /ping       --> to check server status
```
**Product**
```
POST    /product    --> to create product
GET     /product    --> to read product by product_id (pass by query string)
GET     /products   --> to read all products
PUT     /update     --> to update product by product_id (pass by query string)
DELETE  /delete     --> to delete product by product_id (pass by query string)
```
## File Structure
```
webservice/
    ├─── config/
    │       └─── config.go      read config from development.yaml
    ├─── controller/            the logic of each service
    ├─── database/ 
    │       └─── database.go    connect to database
    ├─── form/                  model(struct) of each service
    ├─── model/                 the logic about database of each service
    ├─── server/
    │       ├─── router.go      route the request to the function
    │       └─── server.go      start the server
    ├─── development.yaml       config about database and port
    └─── main.go                run and init the project
```
## Logic
**POST**    /product
*example* `localhost:4000/product`
- Bind request body to `model`
- Insert product to database

**GET**     /product
*example* `localhost:4000/product?product_id=1`
- Parse query string to `Uint`
- Get product by `product_id` from database

**GET**     /products
*example* `localhost:4000/products`
- Get all products in database

**PUT**     /update
*example* `localhost:4000/update?product_id=1`
- Parse query string to `Uint`
- Check if there is product belong to this `product_id`
- Bind request body to `model`
- Update product by `product_id` in database if the variable != zero value

**DELETE**  /delete
*example* `localhost:4000/delete?product_id=1`
- Parse query string to `Uint`
- Check if there is product belong to this `product_id`
- Delete product by `product_id` in database
