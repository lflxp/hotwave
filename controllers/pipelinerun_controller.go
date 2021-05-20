package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	hotwaveiov1alpha1 "hotwave/api/v1alpha1"
)

// PipelineRunReconciler reconciles a PipelineRun object
type PipelineRunReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=hotwave.io.hotwave.io,resources=pipelineruns,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=hotwave.io.hotwave.io,resources=pipelineruns/status,verbs=get;update;patch

func (r *PipelineRunReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("pipelinerun", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *PipelineRunReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&hotwaveiov1alpha1.PipelineRun{}).
		Complete(r)
}
