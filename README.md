# whoa

Simple Go binary `->` tiny Docker image

```shell
$ ID=$( docker run -d -p 80:8080 rackerlabs/whoa )
$ curl $( docker port $ID 8080 )
🎉  Whoa!  🎉
```

## Building

You'll need go with cross compiling support (and `make`)!

```shell
$ make
```

## Background

The `whoa` binary that gets built is [built specifically to be small](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07), courtesy
@kelseyhightower.

```shell
$ CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
```

Assuming everything is happy in go land, that should create a `whoa` binary.

The Dockerfile relies on the super tiny [scratch image](https://docs.docker.com/articles/baseimages/#creating-a-simple-base-image-using-scratch),
adding the `whoa` server as one of its steps:

```Dockerfile
FROM scratch
EXPOSE 8080
ADD whoa /whoa
ENTRYPOINT ["/whoa"]
```

The resulting Docker image is `< 5 MB`:

```
$ docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
rackerlabs/whoa     latest              5ad725856153        About an hour ago   4.437 MB
```

How 'bout them :apple:?
