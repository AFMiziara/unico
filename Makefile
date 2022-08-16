# Build app golang
build:
	go build -o server main.go

run: build
    ./server

watch:
    reflex -s -r '\.go$$' make run

sec:
    gosec -exclude-dir=rules -exclude-dir=vendor -fmt=json -out=results.json -stdout ./...


# stop all containers:
docker-stop-all:
	sudo docker kill $(sudo docker ps -q)


# Remove all containers and docker images
docker-clean-all:
	sudo docker rm $(sudo docker ps -a -q) --force
	sudo docker rmi $(sudo docker images -q) --force
	sudo docker volume rm $(sudo docker volume ls -q) --force

docker-stats:
	sudo docker stats

docker-start:
	sudo docker-compose down
	sudo docker-compose up -d --build


# Install GoLang
install_golang:	
	wget https://dl.google.com/go/go1.16.7.linux-amd64.tar.gz --no-check-certificate
	mkdir -p ~/go/{bin,pkg,src}
	tar -C /usr/local/ -xzf go1.16.7.linux-amd64.tar.gz
	rm go1.16.7.linux-amd64.tar.gz	