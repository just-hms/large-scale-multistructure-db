# large-scale-multi-structure-DB

## Requirements

- `docker`
- `docker compose` plugin

## deploy

```shell
# rename .env.template to .env
# add secrets
make down
make deploy
```

## dev

### docker

```shell
# rename .env.template to .env
# add secrets
make down
make dev

# get the doc in json and html :)
curl http://localhost:8080/api/swagger/doc.json
curl http://localhost:7000/api/swagger/index.html
```

### locally

```shell
# rename .env.template to .env
# add secrets
make down
make infrastructure
go run ./...
```

## test

### docker

```shell
# rename .env.template to .env
# add secrets
make down
make test
```

### locally

```shell
# rename .env.template to .env
# add secrets
make down
make infrastructure
go test ./...
```


## benchmark

```shell
make infrastructure
# run the importer
cd be/internal/usecase/repo
go test -run=^$ -bench=.
```






## google docs

- https://docs.google.com/document/d/18_51BtU6xor5OmC4hQTUU8mCdDXr_RVSLhZaOjpiBwo

## presentations

- https://docs.google.com/presentation/d/1h7jxOkOvwBJSB7aGfDJR2bifR1MJf_eIA3p5B8Vb1QY/edit?usp=sharing