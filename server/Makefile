.DEFAULT_GOAL:=run

name = "dslab1"

kill:
	docker kill $(name)

run:
	docker rm $(name);
	docker build -t $(name) .
	docker run -p 8000:8000 -i --name $(name) $(name)

run_no_rm:
	docker build -t $(name) .
	docker run -p 8000:8000 -i --name $(name) $(name)
