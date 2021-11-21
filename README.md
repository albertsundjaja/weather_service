# Weather Service

## Installation

clone the repo

```
git clone https://github.com/albertsundjaja/weather_service.git
```

build the app

```
go build main.go
```

Run the app (windows)

```
./main.exe
```

or (linux)

```
./main
```

test the app either using curl or postman, the app is running on port `8080`,
make sure no other service is running on this port

```
curl http://localhost:8080/v1/weather?city=sydney
```

## Design Tradeoffs

For ease of testing, `.env` is committed into the repo with a working key.
This is not a good practice for security reason.

For better scalability and reliability in caching, 
an external caching service such as Redis should be used. 
In this way, when we have multiple microservices, the cache can be accessed by all the of the app
which further reduce the need to call the external API.

The config, controller, model and service should have been implemented using 
Inversion of Control (dependency injection) so that they are testable.

More logging should have been implemented to have better debugging capability.