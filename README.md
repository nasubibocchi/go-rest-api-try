0. create .env file (sample: sample.env)
1. execute `docker-compose up` in command line
2. check "mysql-container" built (`docker ps`)
3. copy CONTAINER ID of "mysql-container"
4. execute `docker inspect CONTAINER-ID | grep IPAddress`
5. copy IPAddress and paste it on "DOCKER_LOCAL_HOST" in .env

※work in process
