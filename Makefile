PROJECT_NAME=telegram-bot

rundev:
	./ngrok http 8080 > /dev/null &

init:
	. ./setvars.sh
	echo $$WEBHOOK_URL