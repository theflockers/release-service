/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/redhat-appstudio/release-service/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// ReleaseReconciler reconciles a Release object
type ReleaseReconciler struct {
	client.Client
	Log                   logr.Logger
	OrganizationNamespace string
	Scheme                *runtime.Scheme
}

//+kubebuilder:rbac:groups=appstudio.redhat.com,resources=releases,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=appstudio.redhat.com,resources=releases/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=appstudio.redhat.com,resources=releases/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *ReleaseReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	release := &v1alpha1.Release{}
	err := r.Get(ctx, req.NamespacedName, release)
	if err != nil {
		if errors.IsNotFound(err) {
			l.Info("Release resource not found")

			return ctrl.Result{}, nil
		}

		l.Error(err, "Failed to get Release")

		return ctrl.Result{}, err
	}

	l.Info("Release status after being loaded", "Release", release.Status)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager which monitors PipelineRuns and filter them
// using a predicate to filter out everything but just succeeded Build PipelineRuns.
func (r *ReleaseReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.Release{}).
		WithEventFilter(predicate.GenerationChangedPredicate{}).
		Complete(r)
}
