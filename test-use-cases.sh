#!/usr/bin/env bash

CONCURRENCY=${1:-10}
REQUESTS=${2:-100000}
PORT=1234
URL=http://127.0.0.1:"$PORT"/api/v1/difuntos/1

SERVICE_PID=0

service_start () {
    router=$1

    ./skel -router "$router" -port "$PORT" &
    SERVICE_PID=$!
    sleep 1
}

service_stop () {
    kill -9 "$SERVICE_PID"
    wait "$SERVICE_PID" 2>/dev/null
}

test_use_cases () {
    router=${1:-echo}

    echo "Testing nested routes with $router"

    service_start "$router"

    echo "=========="
    echo "POST shrine"
    echo "=========="
    curl -X POST "http://localhost:$PORT/api/v1/altares" -H "Content-type: application/json" -d '{"MexicanID":1,"Levels":3}'
    echo ""
    echo "=========="
    echo "PUT gift"
    echo "=========="
    curl -X PUT "http://localhost:$PORT/api/v1/altares/1/niveles/1" -H "Content-type: application/json" -d '{"ID":10,"Name":"Cempasúchil","Type":"flower"}'
    echo ""
    echo "=========="
    echo "GET shrine"
    echo "=========="
    curl -X GET "http://localhost:$PORT/api/v1/altares/1"
    echo ""
    echo "=========="
    echo "GET status"
    echo "=========="
    curl -X GET "http://localhost:$PORT/api/v1/status"
    echo ""

    service_stop
}

case "$1" in
    *)
        test_use_cases
        ;;
esac
