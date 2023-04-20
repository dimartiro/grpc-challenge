FROM golang:1.20  

WORKDIR /usr/src/app

#Copiamos los archivos de la carpeta actual hacia docker
COPY . .
#Descargamos las dependencias
RUN go mod download && go mod verify

#Copiamos el resto de los archivos y compilamos la aplicación
WORKDIR /usr/src/app/server
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 50051
#Ejecutamos nuestra app en docker
ENTRYPOINT ["app"]