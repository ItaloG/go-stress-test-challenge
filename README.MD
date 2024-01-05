# Aplicação para realizar teste de estresse em Golang
### O objetivo da aplicação é realizar request em um servidor web e retornar um relatório com as informações das requests.

``Docker``

Para realizar o teste com o docker você deve possuir o [docker](https://www.docker.com/) na sua máquina e executar o seguinte comando

```sh
  docker run --rm italog09/go-stress-cli:latest -u "http://google.com" -c 2 -r 10
```

``Localmente``

Para realizar o teste local basta executar o seguinte comando

```sh
  go run main.go -u "http://google.com" -c 2 -r 10
```

🚧Atenção: Caso você encontre o status code 0 no seu relatório saiba que ele representa as request que não foram resolvidas internamente, ou seja não foi possível realizar a request.🚧