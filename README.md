# merge

<h1> How to start the application </h1>
Pre-requisites: Docker

1) Go inside the directory and run `docker compose up` command
2) Once the images are pulled and built, run the command `docker ps` to know about all the running containers.
3) Get the container id for `merge_assignment` container and go inside it's terminal using `docker exec -it ${container_id} /bin/sh`
4) Inside the terminal of container `merge_assignment` run the migrations: `migrate -database postgres://postgres:postgres@postgres:5432/merge?sslmode=disable -path db/migrations up`
This will create all the required db tables

<h1> API Collections </h1>
https://api.postman.com/collections/23172320-bfc4c7e6-49e0-4d9a-8ccc-264ab14dd2b4?access_key=PMAT-01GQ6ZGXEQXQ9MKD9D5W5M6TR8

Design document can be found here: https://www.notion.so/WIP-Merge-Design-Doc-0d9bad04685444ef8425f7150b94f560
Please note that the doc is still work in progress mode, some formatting is going on


<h1> Tests </h1>
Please note that tests for some of the modules aren't written. Right now only accounts/service and utils tests are written
