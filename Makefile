PROJECT_NAME=telegram-bot

init:
	./ngrok http 8080 > /dev/null &

build:
	go build -a -x -race

test:
	go test -v -cover . ./bot

clean:
	rm telegram-bot
