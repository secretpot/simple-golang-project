#! /bin/bash -ex

echo "Start test"
./simple-gin &
sleep 1

curl http://localhost:8000/api/v1/test/ping
curl http://localhost:8000/api/v1/test/hello

curl http://localhost:8000/api/v1/text/echo?message=hi

kill -9 $(jobs -p)

echo "End test"
