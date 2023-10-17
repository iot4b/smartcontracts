## EVER-IOT: NODE service

### Dependencies
This package requires `ever-client-go` to be installed. 
It is required to install and setup -lton_client: [guide](https://github.com/markgenuine/ever-client-go) in related lib

### Install

```shell
git clone https://github.com/ever-iot/node
cd ./node
go mod tidy
```

### Run

```shell
go run main.go <config-name>
```
example
```shell
go run main.go dev
```
will be used ./config/dev.yml
