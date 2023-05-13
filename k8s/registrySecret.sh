kubectl create secret docker-registry regcred \
  --docker-server=http://<ip>:<port> \
  --docker-username=<username>  \
  --docker-password=<password>  \
  --docker-email=hoc.nguyen@sotatek.com
  -n demo-backend-ns