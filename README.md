# gocalc-api

**gocalc-api** is an API for calc service.
The server runs on `port 8080`.

## installation

```shell
git clone github.com/HayKor/gocalc-api
```

## Usage

```shell
go run cmd/server/main.go
```

## Examples

> [!TIP]
>
> ```shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2+2*2"
> }'
> ```


> [!WARNING]
> HTTP 422 Response (Entity unprocessable)
>
> ```shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2+++2*2"
> }'
> ```


> [!CAUTION]
> HTTP 500 Response (Internal server error)
>
> ```shell
> curl -l 'localhost:8080/api/v1/calculate' \
>     -H 'Content-Type: application/json' \
>     -d '{
>   "expression": "2+2'
> ```
