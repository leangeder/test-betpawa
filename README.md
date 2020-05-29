# betpawa

betpawa code test

## Prerequsites

1. Install [Minikube](https://github.com/kubernetes/minikube/#installation)
1. Install [VirtualBOX](https://www.virtualbox.org/wiki/Downloads)
1. Install [Skaffold](https://github.com/GoogleContainerTools/skaffold#installation)
1. Install [Kustomize](https://github.com/kubernetes-sigs/kustomize)


```
$ minikube start --driver=virtualbox
$ minikube addons enable ingress
```

1. after couple of minute nginx is accessible on the ip address provide via the command line:
```
$ minikube ip
```

By default, all services are brought up with `skaffold run` which just starts
the services. In "dev" mode, code will be hot-reloading for any modification done
on files used during the build process of the container image. To do so, just go 
into the application's folder and run `skaffold dev`, e.g.:


### Deploy application

```
$ skaffold run
$ curl http://$(minikube ip)/
$ curl http://$(minikube ip)/ping
```

### Deploy in development mode

```
$ skaffold dev
$ curl http://$(minikube ip)/
$ curl http://$(minikube ip)/ping
```


### Clean up

```
$ minikube delete
$ rm -rf ~/.minikube
```