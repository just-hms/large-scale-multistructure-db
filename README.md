# large-scale-multi-structure-DB

## Requirements

- `docker`
- `docker compose` plugin

## deploy

```shell
make deploy
```

## dev

```shell
sudo make dev

# get the doc in json and html :)
curl http://localhost:8080/api/swagger/doc.json
curl http://localhost:7000/api/swagger/index.html
```

## test

### docker

```shell
sudo make test
```

### locally

```shell
# edit the env file
sudo make down
sudo make infrastructure
go test ./...
```




## google docs

- https://docs.google.com/document/d/18_51BtU6xor5OmC4hQTUU8mCdDXr_RVSLhZaOjpiBwo

## presentations

- https://docs.google.com/presentation/d/1h7jxOkOvwBJSB7aGfDJR2bifR1MJf_eIA3p5B8Vb1QY/edit?usp=sharing