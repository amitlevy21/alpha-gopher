FROM golang

ENV PROJECT_PATH /go/src/github.com/amitlevy21/alpha-gopher

RUN mkdir -p ${PROJECT_PATH}

COPY . ${PROJECT_PATH}

WORKDIR ${PROJECT_PATH}

RUN go get ${PROJECT_PATH}/...

CMD [ "go", "run" ]
