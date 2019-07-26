FROM golang

ENV PROJECT_PATH /go/src/github.com/amitlevy21/alpha-gopher

RUN mkdir -p ${PROJECT_PATH} /backup

WORKDIR ${PROJECT_PATH}

RUN apt-get update && apt-get install vim sudo rsync tree -y
RUN wget https://sourceforge.net/projects/videlibri/files/Xidel/Xidel%200.9.8/xidel-0.9.8.linux64.tar.gz
RUN tar xvzf xidel-0.9.8.linux64.tar.gz && cp xidel /usr/bin/xidel

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/stretchr/testify/assert
RUN go get -u github.com/gin-contrib/cors

COPY . ${PROJECT_PATH}

CMD [ "go", "run", "./cmd/main.go" ]
