image:
	docker build --platform linux/amd64 -t github.com/Phonevilai/assessment:v0.0.1 -f Dockerfile .

push:
	docker push github.com/Phonevilai/assessment:v0.0.1

container:
	docker run -d -p 9091:9091 --env-file ./dev.env \
    --name my-api github.com/Phonevilai/assessment:v0.0.1

test-ddos:
	ddosify -t http://localhost:2565/healthz -p "2565" -m "GET" -d 2 -n 1000

run-sandbox:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from it_tests