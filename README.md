# go-config-service

![Logo](https://i.imgur.com/YqPKnqI.png)


It is a config server developed for other services or programs to use. Data can be kept according to the data type.😎. 

> Config operations is abstracted from the whole system. You don't have to think ✌️.

### Swagger
http://localhost:4000/docs/index.html

Each endpoint is protected with jwt. It must be entered with default id and pass .env.

### Tech-Stack

go-config-server uses a number of open source projects to work properly:
* [Golang] - for runnig server
* [Mongo Dv] - for storing configs
* [Redis] - for caching

### Installation

Rabbitmq-mail-consumer-server requires [Go](https://golang.org/) to run.

Install the dependencies and start the server.

```sh
$ cd go-config-server
$ go run main.go
```
For Make A Docker Image:

```sh
$ docker build -t  go-config-server .
$ docker run --env-file .env -p 4000:4000 go-config-server
```

### Env Variables

| Env        | Example           
| ------------- |:-------------:
| MONGO_DB_CONNECTION_STRING      | connection string
| MONGO_DB_NAME   | db name      
| ADMIN_USER_ID | admin user id
| ADMIN_USER_PASS| admin user pass
| JWT_SECRET   | jwt secret      
| REDIS_HOST | redis host
| REDIS_PASSWORD      | redis password 
