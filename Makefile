cassie-run:
	docker-compose up -d
	go build -o twitz
	./migrate-up.sh

cassie-stop:
	docker-compose down
