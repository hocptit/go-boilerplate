# Create registry
1. Build image
docker build -f Dockerfile -t server-api:v1 .
docker build -f <Dockerfile name> -t <image-name:tag> .
docker build -t <image-name:tag> .
docker tag server-api 65.108.131.181:5568/server-api:v1
docker tag <image-name:tag> <registry>/<image-name:tag>
**Config docker push to insecure regis**
sudo nano /etc/docker/daemon.json
{ "features": { "buildkit": true },"insecure-registries":["65.108.131.181:5568"] }
sudo service docker restart
docker login 65.108.131.181:5568 -u username -p password
docker push 65.108.131.181:5568/server-api:v1
docker pull 65.108.131.181:5568/server-api:v1

# K8s
minikube start --cpus 4 --memory 6000 --insecure-registry "65.108.131.181:5568"
Register secret registry, run file registrySecret.sh
kubectl apply -f deployment.yaml
kubectl delete -f deployment.yaml
kubectl apply -f hpa.yaml
kubectl apply -f service.yaml
kubectl set image deployment/server-api-deployment server-api=65.108.131.181:5568/server-api:v2 -n demo-backend-ns
=> access minikube docker run curl 127.0.0.1:31186 => done
kubectl port-forward service/demo-backend-api-service -n demo-backend-ns 5668:80 --address='0.0.0.0'
=> using pm2 to keep context

pm2 start "kubectl port-forward service/demo-backend-api-service -n demo-backend-ns 5668:80 --address='0.0.0.0'"
kubectl port-forward service/demo-backend-api-service -n demo-backend-ns 5668:80 --address='0.0.0.0' & 
If forward 5000, not access all pods
For access all pods, access to kubemini, and curl 127.0.0.1:31186 # access nodePort public
=> Use
    In host machine install docker and minikube
    minikube ip => 192.168.49.2
    192.168.49.2:31186
    => using in host machine, not connect from client
    Using nodejs proxy in /proxy

Update with new image
 kubectl set image deployment/server-api-deployment server-api=65.108.131.181:5568/server api:performancetest -n demo-backend-ns


# 