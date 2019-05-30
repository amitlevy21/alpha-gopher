run:
	docker build -t alpha-gopher .
	docker run -it --name alpha_gopher alpha-gopher

stop:
	docker stop alpha_gopher
	docker rm alpha_gopher