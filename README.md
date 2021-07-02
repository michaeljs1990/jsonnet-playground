Jsonnet Playground
==================

This is largely inspired by the work done on golang playground 
and takes advantage of the go-jsonnet program to give you
similar features. In memory and persistance via mysql are
supported.

### Hosted Site

Unable to host the site right now but hope to have a hosted version back soon.

### Dockerfile

A dockerfile is included that can be built as you would any other
docker file.

```
docker build -t jsonnet-playground -- .
```

Additionally a prebuilt image can be pulled down to run like this

```
docker run -p 8080:8080 michaeljs1990/jsonnet-playground:latest
```

Only the latest tag is currently supported.

### Quick MySQL Testing

For quick mysql seting you can spin up a docker container with the
following command.

```
$ docker run --net=host --name jsonnet --rm -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=jsonnet -d mysql:5.7
$ JSONNET_MYSQL_CONN="root:secret@tcp(127.0.0.1:3306)/jsonnet" ./jsonnet-playground -sql
```

### Local Testing

Build and run

```
./jsonnet-playground -in-memory
```
