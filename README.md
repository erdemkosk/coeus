# Coeus

 ![Logo](https://i.imgur.com/Jgf7swu.gif)
## Documentation

Coeus is a custom flag and config service creation tool. Its main purpose is to make business decisions used in many projects or to turn off some features when requested, or to keep translations for applications that serve in more than one language. You can easily access the necessary data by cloning and using the npm client. It's free, the only fee is your own server fee. For this reason, it can be changed individually.

>▪️ It can be used to turn on or off some features in the system.
▪️Information that can change at any time, such as a phone number, can be stored.
▪️It can be used as a translation storage tool for multiple languages.


In addition, it has a simple panel that is protected by the password you set. In this way, you can easily enter, update or delete your data.
### Swagger
localhost:port/swagger/index.html

![Logo](https://i.imgur.com/sAmY4b8.png)

Each endpoint is protected with jwt. It must be entered with default id and pass .env.
## Deployment
### With Heroku
[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/erdemkosk/coeus)
### With Docker
For Make A Docker Image:
```sh
$ docker build -t  go-config-server .
$ docker run --env-file .env -p 4000:4000 go-config-server
```

## Tech
* [Golang] - for runnig server
* [Mongo Dv] - for storing configs
* [Redis] - for caching
* [HTML] - for panel
* [Vue Js] - for panel operations

## Installation
Coeus requires [Go](https://golang.org/) to run.
Install the dependencies and start the server.

```sh
$ go mod vendor
$ go get -u github.com/swaggo/swag/cmd/swag
$ $HOME/go/bin/swag init // for creating swagger files
```

```sh
$ cd go-config-server
$ go run main.go
```
## Plugins

The npm client wrapper for javascript is currently in development.

```sh
$ npm i coeus-client
```

### Example of Usage

```sh
const coeus = require('coeus-client');
const config = require('../../config');

coeus.connect({
  url: config.coeus.url,
  identity: config.coeus.identity,
  password: config.coeus.password,
  interval: config.coeus.interval,
});

coeus.client.addKeys({ keys: ['co:email', 'co:json', 'co:NUMBER_OF_GOOD_TIMES'] });

// this is updated event on every interval
coeus.client.on('updated', (configs) => {
  console.log(configs);
});
// this error event for when getting error on server
coeus.client.on('error', (value) => {
  console.log(value);
});

const express = require('express')
const app = express()
const port = 3000

app.get('/', (req, res) => {
    res.send(coeus.client.getConfig({ key: 'co:NUMBER_OF_GOOD_TIMES' }))
})

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})
```

### Env Variables✨

| Env        | Example           
| ------------- |:-------------:
| MONGO_DB_CONNECTION_STRING      | connection string
| MONGO_DB_NAME   | db name      
| ADMIN_USER_ID | admin user id
| ADMIN_USER_PASS| admin user pass
| JWT_SECRET   | jwt secret      
| REDIS_HOST | redis host
| REDIS_PASSWORD      | redis password 

## License
Copyright (c) 2022 Mustafa Erdem Köşk <erdemkosk@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.



  
