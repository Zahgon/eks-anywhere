package handlers

import (
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func ChildObjectToClusters(log logr.Logger) handler.MapFunc {
	_ = "STUB: not implemented"
	return *new(handler.MapFunc)
}

func reconcileRequestForOwnerRef(o client.Object, owner metav1.OwnerReference) reconcile.Request {
	_ = "STUB: not implemented"
	return *new(reconcile.Request)
}
