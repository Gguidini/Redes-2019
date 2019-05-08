FROM golang:1.11.10-alpine3.9

RUN mkdir -p $GOPATH/src/github.com/Redes-2019

WORKDIR $GOPATH/src/github.com/Redes-2019

COPY ircclient/* ircclient/
COPY connection/* connection/
COPY userinterface/* userinterface/

RUN go build github.com/Redes-2019/connection/
RUN go build github.com/Redes-2019/userinterface/
RUN go install github.com/Redes-2019/ircclient

CMD ["ircclient"]