# 04-container

https://kind.sigs.k8s.io/docs/user/quick-start/#installation
https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/

```
$ mkdir -p ~/.bin

$ curl -Lo ~/.bin/kind https://kind.sigs.k8s.io/dl/v0.12.0/kind-linux-amd64
$ chmod +x ~/.bin/kind

$ curl -Lo ~/.bin/kubectl "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
$ chmod +x ~/.bin/kubectl

$ export PATH="~/.bin:$PATH"
```

```
$ kind create cluster --config example-cluster.yaml --name example-cluster
Creating cluster "example-cluster" ...
 ✓ Ensuring node image (kindest/node:v1.23.4) 🖼
 ✓ Preparing nodes 📦 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
 ✓ Joining worker nodes 🚜
Set kubectl context to "kind-example-cluster"
You can now use your cluster with:

kubectl cluster-info --context kind-example-cluster

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/
```

```
$ kubectl get nodes --watch
NAME                            STATUS   ROLES                  AGE   VERSION
example-cluster-control-plane   Ready    control-plane,master   67s   v1.23.4
example-cluster-worker          Ready    <none>                 34s   v1.23.4
```

```
$ kubectl create -f deployment.yaml
deployment.apps/my-webserver created
service/my-webserver created
deployment.apps/lorem-ipsum created
service/lorem-ipsum created
```

```
$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
lorem-ipsum-85556b6b59-dnhd7    1/1     Running   0          37s
lorem-ipsum-85556b6b59-vq925    1/1     Running   0          37s
my-webserver-77f864867f-gdv9t   1/1     Running   0          37s
my-webserver-77f864867f-tz8vp   1/1     Running   0          37s
```

```
$ kubectl get services
NAME           TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
kubernetes     ClusterIP   10.96.0.1      <none>        443/TCP        2m29s
lorem-ipsum    NodePort    10.96.29.228   <none>        80:30001/TCP   53s
my-webserver   NodePort    10.96.130.42   <none>        80:30000/TCP   53s
```

```
$ curl localhost:30000
```
