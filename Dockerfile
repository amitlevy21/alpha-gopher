FROM golang

ENV PROJECT_PATH /go/src/github.com/amitlevy21/alpha-gopher

EXPOSE 8080

RUN mkdir -p ${PROJECT_PATH} /backup

WORKDIR ${PROJECT_PATH}

RUN apt-get update && apt-get install vim sudo rsync aha htop libxml2-utils -y

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/stretchr/testify/assert

COPY . ${PROJECT_PATH}

CMD [ "./web/alpha-gopher-frontend/vue_htop.sh" ]
