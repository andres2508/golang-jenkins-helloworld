# STEP 1 build executable binary
FROM golang:alpine as builder
# Install git
RUN apk update && apk add git 
#ENV GOPATH="~/go"
COPY ./go-service $GOPATH/src/go-service
# COPY prueba.txt /hello/prueba.txt
WORKDIR $GOPATH/src/go-service
#get dependancies
#you can also use dep
RUN go get -d -v
#build the binary
# RUN go build -o /go/bin/server

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/backoffice

# STEP 2 build a small image
# start from scratch 
# ----- NOTA ------
# Me toco cambiar scratch porque este no tiene manejador de paquetes para las credenciales de aws
# FROM traefik:alpine
FROM scratch
# FROM ubuntu
# Copy our static executable
# RUN mkdir -p /go/bin
COPY --from=builder /go/bin/backoffice /go/bin/backoffice
# RUN export AWS_CONFIG_FILE = /root/.aws
EXPOSE 3535
ENTRYPOINT ["/go/bin/backoffice"]
# WORKDIR $GOPATH/src/pekiz_tools_api
# RUN ls -la
# EXPOSE 5555
# ENTRYPOINT ["go", "run", "server.go"]
