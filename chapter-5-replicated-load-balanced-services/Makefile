deploy:
	kubectl apply -f sample-backend-deployment.yaml
	kubectl apply -f sample-backend-service.yaml
	kubectl create configmap varnish-config --from-file=varnish-cache-config.vcl
	kubectl apply -f cache-deployment.yaml
	kubectl apply -f cache-service.yaml
	kubectl create configmap nginx-conf --from-file=nginx.conf
	kubectl apply -f nginx-deployment.yaml
	kubectl apply -f nginx-service.yaml

delete-deployment:
	kubectl delete -f nginx-service.yaml
	kubectl delete -f nginx-deployment.yaml
	kubectl delete configmaps nginx-conf
	kubectl delete -f cache-service.yaml
	kubectl delete -f cache-deployment.yaml
	kubectl delete configmaps varnish-config
	kubectl delete -f sample-backend-service.yaml
	kubectl delete -f sample-backend-deployment.yaml

time-call:
	curl -k -w "@curl-timing.txt" -o /dev/null -s https://192.168.64.2:31457/term
