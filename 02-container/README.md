# 02-container

https://docs.docker.com/engine/reference/builder/
https://hub.docker.com/_/golang

```bash
$ docker build -t lorem-ipsum:1.0 .
Sending build context to Docker daemon  13.82kB
Step 1/6 : FROM golang:1.18
 ---> 65375c930b21
Step 2/6 : WORKDIR /build
 ---> Using cache
 ---> 2d8df02b02d4
Step 3/6 : COPY . /build
 ---> 4bc6a6a7c2e8
Step 4/6 : RUN CGO_ENABLED=0 GOOS=linux go build -a -o lorem-ipsum .
 ---> Running in 465f0a811a04
go: downloading go.uber.org/zap v1.21.0
go: downloading gopkg.in/loremipsum.v1 v1.1.0
go: downloading go.uber.org/multierr v1.6.0
go: downloading go.uber.org/atomic v1.7.0
Removing intermediate container 465f0a811a04
 ---> f7a4e0de8404
Step 5/6 : EXPOSE 80
 ---> Running in 1bc0defdfe5b
Removing intermediate container 1bc0defdfe5b
 ---> 622fd6dae48f
Step 6/6 : ENTRYPOINT ["/build/lorem-ipsum"]
 ---> Running in 698a61d16ec3
Removing intermediate container 698a61d16ec3
 ---> 1e1f08a92b99
Successfully built 1e1f08a92b99
Successfully tagged lorem-ipsum:1.0

$ docker images
REPOSITORY            TAG       IMAGE ID       CREATED          SIZE
lorem-ipsum           1.0       1e1f08a92b99   18 seconds ago   1.03GB

$ docker run -d -p 8081:80 lorem-ipsum:1.0
ca8a4d4cf666dc634ec30b5c7b4907f69d182b455445a04a5744c855d1a5c6af

$ docker ps
CONTAINER ID   IMAGE             COMMAND                CREATED          STATUS          PORTS                                   NAMES
ca8a4d4cf666   lorem-ipsum:1.0   "/build/lorem-ipsum"   12 seconds ago   Up 11 seconds   0.0.0.0:8081->80/tcp, :::8081->80/tcp   sad_ritchie

$ curl localhost:8081
Tincidunt justo cras morbi curae fusce dui luctus varius phasellus tortor aliquet ultrices egestas mattis ligula mus mollis a lectus habitant accumsan. Potenti nam nostra faucibus lectus urna lobortis magnis leo penatibus habitasse lacinia sollicitudin fames venenatis torquent senectus auctor duis netus ipsum posuere. Eleifend turpis sapien ad mi cras cursus phasellus lorem augue facilisis erat faucibus congue facilisi. Habitasse libero elementum pulvinar maecenas quisque imperdiet hendrerit et blandit porta senectus molestie parturient aenean suscipit ac purus gravida eleifend dignissim nisl vulputate morbi netus.%
```
