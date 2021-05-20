package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hotwaveiov1alpha1 "hotwave/api/v1alpha1"
)

// PlaybookRunReconciler reconciles a PlaybookRun object
type PlaybookRunReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=hotwave.io.hotwave.io,resources=playbookruns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hotwave.io.hotwave.io,resources=playbookruns/status,verbs=get;update;patch

func (r *PlaybookRunReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("playbookrun", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PlaybookRunReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hotwaveiov1alpha1.PlaybookRun{}).
		Complete(r)
}
