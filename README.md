# REST - gRPC

Examples of REST vs gRPC API workflows. 

## Go Swagger

[go-swagger](https://github.com/go-swagger/go-swagger)

### Install

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
```

### Generate swagger.json from annotations and serve

[Documentation](https://goswagger.io/use/spec.html).

```bash
swagger generate spec -o ./swagger.json --scan-models
swagger serve -F=redoc swagger.json
```

### Server from your Go code

Copy `dist` folder of swagger-ui and place your `swagger.json` in the same folder.

```bash
git clone https://github.com/swagger-api/swagger-ui.git
cp -r swagger-ui/dist clus-2019/rest-api/
```

Edit `index.html`.

```bash
url: "https://petstore.swagger.io/v2/swagger.json",
-> url: "./swagger.json",
```

### Generate an API server from swagger.json

To generate a [server for a swagger spec](https://goswagger.io/generate/server.html) document:

```bash
swagger generate server [-f ./swagger.json] -A [application-name]
```

### Generate an API client from swagger.json

To generate a [client for a swagger spec](https://goswagger.io/generate/client.html) document:

```bash
swagger generate client [-f ./swagger.json] -A [application-name]
```
