#!/usr/bin/env bash

CONCURRENCY=${1:-10}
REQUESTS=${2:-100000}
PORT=1234
URL=http://127.0.0.1:"$PORT"/api/v1/difuntos/1

SERVICE_PID=0

service_start () {
    router=$1

    ./skel -router "$router" -port "$PORT" -loglevel panic & # do not log when benchmarking
    SERVICE_PID=$!
    sleep 1
}

service_stop () {
    kill -9 "$SERVICE_PID"
    wait "$SERVICE_PID" 2>/dev/null
}

benchmark_router () {
    router=$1

    r=$(echo "$router" | cut -c 1)

    service_start "$r"

    ab -q -k -c "$CONCURRENCY" -n "$REQUESTS" -g plot/concurrency-"$r-$CONCURRENCY".dat "$URL" | \
        grep -E '(Concurrency|Failed|Requests|Time).*:'

    service_stop
}

benchmark_simple () {

    echo "Benchmarking router adapters using $URL"

    service_start nethttp
    curl "$URL"
    service_stop

    echo ""

    for router in nethttp iris httprouter echo fasthttp gin
    do
        echo "Benchmark $router"
        benchmark_router "$router"
        echo ""
        echo "-----"
    done
}

benchmark_degradation () {

    echo "Benchmarking router adapters with increasing concurrency using $URL"

    service_start nethttp
    curl "$URL"
    service_stop

    echo ""

    for router in nethttp iris httprouter echo fasthttp gin
    do
        echo "Benchmark $router"
        for concurrency in 1 5 10 20 100
        do
            CONCURRENCY=$concurrency
            benchmark_router "$router" | grep Concurrency
        done
    done

}

case "$1" in
    degradation)
        benchmark_degradation
        ;;
    *)
        benchmark_simple
        ;;
esac
