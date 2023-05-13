# Create registry
1. Build image
docker build -t server-api:v1 .
docker build -t <image-name:tag> .
docker tag server-api 65.108.131.181:5568/server-api:v1
docker tag <image-name:tag> <registry>/<image-name:tag>
**Config docker push to insecure regis**
sudo nano /etc/docker/daemon.json
{ "features": { "buildkit": true },"insecure-registries":["65.108.131.181:5568"] }
sudo service docker restart
docker push 65.108.131.181:5568/server-api:v1
docker pull 65.108.131.181:5568/server-api:v1

# K8s
minikube start --cpus 4 --memory 6000 --insecure-registry "65.108.131.181:5568"
Register secret registry, run file registrySecret.sh
kubectl apply -f deployment.yaml
kubectl delete -f deployment.yaml
kubectl apply -f hpa.yaml





# 