image:
	docker build -t github.com/phonevilai/assessment:v0.0.1 -f Dockerfile .

push:
	docker push github.com/phonevilai/assessment:v0.0.1

container:
	docker run -d -p 2565:2565 --env-file ./dev.env \
    --name my-api github.com/phonevilai/assessment:v0.0.1

test-ddos:
	ddosify -t http://localhost:2565/expenses -p "2565" -m "GET" -d 1 -n 300

sandbox:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from it_tests