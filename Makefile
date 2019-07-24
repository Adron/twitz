cassie-run:
	docker-compose up -d
	go build -o twitz

cassie-stop:
	docker-compose down
