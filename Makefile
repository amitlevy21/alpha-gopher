build:
	docker build -t alpha-gopher .

dev: destroy
	docker run -it --name alpha_gopher -v ${GOPATH}/src/github.com/amitlevy21/alpha-gopher:/go/src/github.com/amitlevy21/alpha-gopher alpha-gopher bash

destroy:
	docker stop alpha_gopher
	docker rm alpha_gopher