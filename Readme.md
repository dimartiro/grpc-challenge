# Instrucciones

## Requisitos previos

### Instalar golang

Puedes descargar golang desde su [web oficial](https://go.dev/dl/)

### Instalar protoc

Para instalar el compilador de protoc podes ver las instrucciones en su [web](https://grpc.io/docs/protoc-installation/)

Para linux con paquetes apt
```
apt install -y protobuf-compiler
```

Para mac usando brew
```
brew install protobuf
```

En windows
```
Mejor usa linux :D
```

Para windows se puede descargar el ejecutable de protocol buffers desde su [github](https://github.com/protocolbuffers/protobuf/releases) (revisar la seccion de assets por si no aparece entera) y siguiendo [esta guia](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/) para la instalación

## Generar los stubs
```
make gen-stubs
```
o ejecutando
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/demo.proto
```

## Ejecutar el servidor
```
go run server/main.go
```

En caso de querer ejecutarlo en un puerto distinto

```
go run server/main.go --port PUERTO
```

Ejemplo:

```
go run server/main.go --port 50052
```