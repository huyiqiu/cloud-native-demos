# Demo-device-plugin

## 1. Create test device node
```bash
sudo mknod /dev/coffee0 c 100 0
sudo chmod 666 /dev/coffee0

sudo mknod /dev/coffee1 c 100 1
sudo chmod 666 /dev/coffee1
```

## 2. Build
```bash
make docker-build
make docker-push
```

## 3. Deploy
```bash
kubectl apply -f deploy/coffee-ds.yaml
```

## 4. Create test pod
```bash
kubectl apply -f e2e/nginx-deploy.yaml
```

## Reference
[Kubernetes开发知识–device-plugin的实现](https://www.myway5.com/index.php/2020/03/24/kubernetes-device-plugin/)