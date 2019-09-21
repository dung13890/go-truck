# go-truck
Tool deployer

### With docker & docker-compose

```sh

## Outsite docker
docker-compose up -d
docker-compose exec truck bash


## Inside docker
dep ensure -update

go run main.go
```


