package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yvz5/kube-blitz/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Printf("Error getting kubeconfig: %v", err)
		os.Exit(1)
	}

	mgr, err := manager.New(cfg, manager.Options{})
	if err != nil {
		fmt.Printf("Error creating manager: %v", err)
		os.Exit(1)
	}

	ctrl, err := controller.New("kube-blitz-controller", mgr, controller.Options{
		Reconciler: &reconcileHandler{},
	})
	if err != nil {
		fmt.Printf("Error creating controller: %v", err)
		os.Exit(1)
	}

	if err := ctrl.Watch(source.Kind(mgr.GetCache(), &corev1.Pod{}), &handler.EnqueueRequestForObject{}); err != nil {
		fmt.Printf("Error watching Pod: %v", err)
		os.Exit(1)
	}

	if err := mgr.Start(context.Background()); err != nil {
		fmt.Printf("Error starting manager: %v", err)
		os.Exit(1)
	}
}

type reconcileHandler struct {
	client client.Client
}

// Implement reconcile method
func (r *reconcileHandler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	// Fetch LoadTest instance
	instance := new(v1alpha1.LoadTest)

	if err := r.client.Get(ctx, req.NamespacedName, instance); err != nil {
		return reconcile.Result{}, err
	}

	// Create a Pod object for Locust
	locustPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "locust-master",
			Namespace: instance.Namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "locust",
					Image: "locustio/locust",
					// Add more container configuration as needed
				},
			},
		},
	}

	// Set owner reference to the LoadTest instance
	if err := controllerutil.SetControllerReference(instance, locustPod, &runtime.Scheme{}); err != nil {
		return reconcile.Result{}, err
	}

	// Create the Locust pod
	if err := r.client.Create(ctx, locustPod); err != nil {
		return reconcile.Result{}, err
	}

	// Pod created successfully
	fmt.Println("Locust pod created successfully")

	// Run Locust load test in a Kubernetes pod
	// You'll need to implement the logic to create the pod and execute the load test

	// Once load test completes, collect results and publish them
	// You'll need to implement the logic to collect and publish load test results
	return reconcile.Result{}, nil
}
