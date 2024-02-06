crd:
	kubectl apply -f crd.yaml

ctrl:
	kubectl apply -f controller.yaml

docker:
	docker build -t kube-blitz-controller:latest .