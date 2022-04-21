# 01-container

https://docs.docker.com/engine/reference/builder/
https://hub.docker.com/_/nginx

```bash
$ docker build -t my-webserver:1.0 .
Sending build context to Docker daemon  6.144kB
Step 1/2 : FROM nginx:1.21.6
 ---> fa5269854a5e
Step 2/2 : COPY index.html /usr/share/nginx/html/index.html
 ---> f2fbbc62c156
Successfully built f2fbbc62c156
Successfully tagged my-webserver:1.0

$ docker images
REPOSITORY           TAG       IMAGE ID       CREATED          SIZE
my-webserver         1.0       f2fbbc62c156   12 seconds ago   142MB
nginx                1.21.6    fa5269854a5e   22 hours ago     142MB

$ docker run -d -p 8080:80 my-webserver:1.0
fc9473babaf0b27962e180f77e2c79876ac8d2c47bfcfa4707b6f62a3c01657a

$ docker ps
CONTAINER ID   IMAGE              COMMAND                  CREATED         STATUS         PORTS                                   NAMES
fc9473babaf0   my-webserver:1.0   "/docker-entrypoint.…"   6 seconds ago   Up 5 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   dazzling_moser

$ curl http://localhost:8080
<html>
    <body>
        <h1>example page</h1>

        <img src="https://picsum.photos/200/300?random=1" />
    </body>
</html>%
```
