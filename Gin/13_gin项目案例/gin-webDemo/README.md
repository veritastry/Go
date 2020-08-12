# Web-Server-Framework-Golang

## Script Help

```sh
./bin/dev help
```

## Start Server

```sh
./bin/dev start server
```

## Show Server Logs

```sh
./bin/dev logs server
```

## Stop Server

```sh
./bin/dev stop server
```

## Build Server

```sh
./bin/dev build
```

## Build Docker Image

<!-- TODO: move to dev script -->

```sh
docker build -t $imageName:tag .
```

## Start Databases

<!-- TODO: move to dev script -->

```sh
docker swarm init --advertise-addr 127.0.0.1
docker stack deploy -c docker-stack.yml mydbs
```
