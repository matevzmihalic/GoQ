SetUp
-----

1. Install [docker compose](https://docs.docker.com/compose/install/)
2. Run `docker-compose up` which will build images and start containers
3. Run client which will create random request every time you press enter: `docker-compose run --rm goq ./client -a goq_goq_1`
4. Use `docker-compose scale arithmetics=2 bcrypt=4 reversetext=2 fibonacci=3` to add more workers


Slow workers
------------

For testing purposes workers can have simulated slow output with `-s` flag (which just adds 1 second delay).
If you want fast workers remove `-s` flag from workers in `docker-compose.yml`.