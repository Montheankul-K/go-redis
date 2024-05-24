## Redis
In-memory data structure store (key-value), use as database cache and message broker
Storage base on instance memory

## Redis CLI:
Open Redis: redis-cli 
default host is localhost use -h host to specific host
1. Ping Redis: ping
2. Insert data to Redis: set key value value
Ex. set name value oat
- option ex time: set storage time
Ex. set name value oat set 5 -> store in Redis 5 sec
3. Get data from Redis: get key
Ex. get name

## K6
Run load test:
- docker compose run --rm k6 run /scripts/test.js (k6 in container)
- k6 run ./scripts/test.js (not in container)
options:
- -u: number of virtual users
- -d: test duration
- -i: total iteration
- -o: send metrics to influxdb: -o influxdb=host
Ex. docker compose run --rm k6 run /scripts/test.js -u 5 -d 5s

K6 can store data in influx (time-series database)
if use influx in container, recommend use k6 in container

