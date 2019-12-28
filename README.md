# Chaos Mesh 
Chaos Mesh is a powerful chaos engineering platform for kubernetes. Now it has two parts Chaos-Operator and Chaos-Dashboard
Chaos-Operator is in use now. You can only install Chaos-operator to do chaos operations. Chaos-Dashboard is under development. It now can shows the impact of chaos to TiDB Cluster(other target such as etcd need some configure)


## Chaos Operator

It is used to inject chaos into the applications and Kubernetes infrastructure in a managed fashion. Which provides easy definitions for chaos experiments and automates the execution of chaos experiments.

![Chaos Operator](./static/chaos-mesh-overview.png)

### Deploy 

#### Prerequisites 

Before deploying Chaos Mesh, make sure the following items are installed on your machine: 

* Kubernetes >= v1.12
* [RBAC](https://kubernetes.io/docs/admin/authorization/rbac) enabled (optional)
* Helm version >= v2.8.2 and < v3.0.0
* Protobuf version >= v3.6.1 and < v3.7.0 and go support

#### Install Chaos Operator 

##### Get the Helm files

```bash
$ git clone https://github.com/pingcap/chaos-mesh.git
$ cd chaos-mesh/
```

##### Create custom resource type

Chaos Mesh uses [CRD (Custom Resource Definition)](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/) 
to extend Kubernetes. Therefore, to use Chaos Mesh, you must first create the related custom resource type.

```bash
$ kubectl apply -f manifests/
$ kubectl get crd podchaos.pingcap.com
```

##### Install Chaos Operator 

```bash
$ helm install helm/chaos-mesh --name=chaos-mesh --namespace=chaos-testing
$ kubectl get pods --namespace chaos-testing -l app.kubernetes.io/instance=chaos-mesh
```

### Usage

#### Define chaos experiment config file 

eg: define a chaos experiment to kill one tikv pod randomly

create a chaos experiment file and name it pod-kill-example.yaml

```yaml
apiVersion: pingcap.com/v1alpha1
kind: PodChaos
metadata:
  name: pod-kill-example
  namespace: chaos-testing
spec:
  action: pod-kill
  mode: one
  selector:
    namespaces:
      - tidb-cluster-demo
    labelSelectors:
      "app.kubernetes.io/component": "tikv"
  scheduler:
    cron: "@every 1m"
```

##### PodChaos

PodChaos designs for the chaos experiments about pods.

* **action** defines the specific pod chaos action, supported action: pod-kill / pod-failure
* **mode** defines the mode to run chaos action, supported mode: one / all / fixed / fixed-percent / random-max-percent
* **selector** is used to select pods that are used to inject chaos action.
* **scheduler** defines some scheduler rules to the running time of the chaos experiment about pods. 
More cron rule info: https://godoc.org/github.com/robfig/cron


more examples: [https://github.com/pingcap/chaos-mesh/tree/master/examples](https://github.com/pingcap/chaos-mesh/tree/master/examples) 

#### Create a chaos experiment

```bash
$ kubectl apply -f pod-kill-example.yaml
$ kubectl get podchaos --namespace=chaos-testing
```

#### Update a chaos experiment

```bash
$ vim pod-kill-example.yaml
$ kubectl apply -f pod-kill-example.yaml
```

#### Delete a chaos experiment

```bash
$ kubectl delete -f pod-kill-example.yaml
```

### additional

There are now support such kind of chaos:
* pod-kill, pod-kill only works on pods which created by deployment / statefulset, when this kind of pod kill, it will be restarted by kubenetes.
* pod-fail, pod-fail means pod will fail but not death. Only the process will be kill, the pod is still there
* netem chaos, netem chaos contains some kind of chaos. such as delay, duplicate etc. you can find more in example
* network partition, network partition can partition pod into differenet networks. If helpful for corss IDC's application
* IO chaos, now IO chaos contains IO delay and IO errno. io delay means you can specify how slow the io will return. IO errno means you read/write IO operation will return error

## Chaos Dashboard

Chaos dashboard can be used to visualize chaos events. However, it **only** supports tidb now (so it isn't installed by default).

### Deploy

#### Prerequisites 

Before deploying Chaos Mesh, make sure the following items are installed on your machine: 

* Kubernetes >= v1.12
* [RBAC](https://kubernetes.io/docs/admin/authorization/rbac) enabled (optional)
* Helm version >= v2.8.2 and < v3.0.0
* Protobuf version >= v3.6.1 and < v3.7.0 and go support

#### Install Chaos Dashboard
To install dashboard with `chaos-mesh`:

```
helm install helm/chaos-mesh --name=chaos-mesh --namespace=chaos-testing --set dashboard.create=true
```

Then `svc/chaos-dashboard` will be created under `chaos-testing` namespace and you can access `chaos-dashboard` throught it.

#### Demo
[![Watch the video](https://raw.github.com/GabLeRoux/WebMole/master/ressources/WebMole_Youtube_Video.png)](./static/dashboard.mov)