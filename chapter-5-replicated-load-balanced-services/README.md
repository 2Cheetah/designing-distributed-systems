# Description

A replicated dictionary service with a replicated varnish cache layer and behind a replicated nginx service that terminates SSL/TLS.

## Steps

- [x] Create a deployment for sample-backend with 4 replicas
- [x] Create a load balancer for the sample-backend replicas
- [x] Create config and configMap for cache layer
- [x] Create a deployment for cache with 2 replicas
- [x] Create a load balancer for the cache replicas
- [x] Create a config and configMap for ssl termination layer
- [x] Create a secret and certificate for ssl termination layer
- [x] Create a deployment for the ssl termination layer with 3 replicas
- [x] Create a load balancer for the ssl termination layer
- [ ] Use helm for varnish deployment

## How-to

1. To call cache service from host, the cache service `type` must be `NodePort`. Get its url: `minikube service cache-service --url`, use it to call the service.
2. To build and upload a new image to minikube: `minikube image build -t sample-backend:latest .`
3. To create an SSL certificate and private key: `openssl req -x509 -newkey rsa:4096 -nodes -keyout server.key -out server.crt -days 365 -sha256`
4. To create a secret in k8s cluster: `kubectl create secret tls ssl --cert=server.crt --key=server.key`
5. To create a configMap: `kubectl create configmap nginx-conf --from-file=nginx.conf`
