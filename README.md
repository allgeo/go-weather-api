## Install go dependency

```
go mod download
```

## Define your environment variables in .env file

<pre>
API_PORT=8082
WEATHER_BASE_URL=https://api.openweathermap.org
WEATHER_API_KEY=

</pre>

## Build go package

```
go build main.go
```

## Run the service

```
go run main.go
```
