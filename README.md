# go-simple-todo
A simple todo API written in Golang

### PreInstall
1. Docker https://docs.docker.com/desktop/install/windows-install/

### How to use with Docker
1. Clone this Repository using `git clone`
2. Run Docker Desktop
3. Check config/db.go for the commented out code and adjust accordingly
4. use CMD Command on this project root
    1. `docker-compose build`
    2. `docker-compose up`
5. access http://localhost:8080/swagger/index.html#/
6. The app is ready to use.
---
### How to use with Go Build
1. Clone this Repository using `git clone`
2. Check config/db.go for the commented out code and adjust accordingly
3. use CMD Command on this project root
    1. `go build -tags netgo -ldflags '-s -w' -o app`
    2. `./app`
4. access http://localhost:8080/swagger/index.html#/
5. The app is ready to use.

The live demo is available at https://micbun-golang-simple-todo.onrender.com/swagger/index.html

For more information, please contact me LinkedIn: https://www.linkedin.com/in/MicBun