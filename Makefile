image:
	docker build --platform linux/amd64 -t github.com/Phonevilai/assessment:v0.0.1 -f Dockerfile .

push:
	docker push github.com/Phonevilai/assessment:v0.0.1

container:
	docker run -d -p 9091:9091 --env-file ./dev.env \
    --name my-api github.com/Phonevilai/assessment:v0.0.1

test-limiter:
	echo 'GET http://localhost:2565/limitz' | vegeta attack -rate=10/s -duration=1s | vegeta report

test-ddos:
	ddosify -t https://api-cluster.ldblao.la/ldb-x-laonsw/uat-api/v1/payments -p "443" -a "LDB_SERVICES:de4122858d9545cd9f4996021f0ab67d" -m "POST" -h 'Content-Type: application/json' -b \
	"{\
        "invoiceIds": [\
            "210205000006"\
        ]\
    }" -d 2 -n 1000 -l incremental

