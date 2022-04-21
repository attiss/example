# 03-container

https://hub.docker.com/settings/security?generateToken=true

```
$ docker login -u attiss
Password:
WARNING! Your password will be stored unencrypted in /home/attiss/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

$ docker tag my-webserver:1.0 attiss/my-webserver:1.0

$ docker push attiss/my-webserver:1.0
The push refers to repository [docker.io/attiss/my-webserver]
9fc38c62701d: Pushed
ea4bc0cd4a93: Mounted from library/nginx
fac199a5a1a5: Mounted from library/nginx
5c77d760e1f4: Mounted from library/nginx
33cf1b723f65: Mounted from library/nginx
ea207a4854e7: Mounted from library/nginx
608f3a074261: Mounted from library/nginx
1.0: digest: sha256:403d488eb96b73bb54990510517595f8e4ce4bb32a2b174c4b0d367ca9374bb1 size: 1777

$ docker tag lorem-ipsum:1.0 attiss/lorem-ipsum:1.0

$ docker push attiss/lorem-ipsum:1.0
The push refers to repository [docker.io/attiss/lorem-ipsum]
af92a7c47283: Pushed
2a3196ac03fb: Pushed
0cd0e32a2277: Pushed
d9a88bc06a40: Mounted from library/golang
5214810daaa5: Mounted from library/golang
6908ca9042ea: Mounted from library/golang
08fa02ce37eb: Mounted from library/golang
a037458de4e0: Mounted from library/golang
bafdbe68e4ae: Mounted from library/golang
a13c519c6361: Mounted from library/golang
1.0: digest: sha256:b2338ed6446c38594913efdb745bc6952ce82ae790ae1ceaf9a1308506ecdc24 size: 2423
```


https://hub.docker.com/r/attiss/my-webserver
https://hub.docker.com/r/attiss/lorem-ipsum
