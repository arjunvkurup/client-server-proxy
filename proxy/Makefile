.DEFAULT_GOAL:=run

name = "proxy"

kill:
	docker kill $(name)

run:
	docker rm $(name);
	docker build -t $(name) .
	docker run -p 8080:8080 -i --name $(name) $(name)

run_no_rm:
	docker build -t $(name) .
	docker run -p 8080:8080 -i --name $(name) $(name)
