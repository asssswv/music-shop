# How to run and testing this project
* ## Run db (postgresql).
      ### We will use docker to run the database
      docker run -it --name some-postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -e POSTGRES_DB=db -p 5432:5432 --rm postgres
* ## Run server
   + ### You can change port on which the server will run in the folder ___configs___ in the file ___configs.yml___
   + ### Run server
         go run cmd/main.go
* ## Run unit tests
      cd tests
      go test -v
