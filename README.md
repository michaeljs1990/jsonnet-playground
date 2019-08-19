Jsonnet Playground
==================

This is largely inspired by the work done on golang playground 
and takes advantage of the go-jsonnet program to give you the
similar features. This very much has the minimum amount of code
to get things working but lacks persistent stoage currently for
code snippets.

### Demo Site

This site will not be staying around forever as I look to move the
playground to some other place. For now you can test things out at
https://jsonnet.xrt0x.com/ though.

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
