/*
Copyright 2024.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiextensionsk8siov1 "github.com/cybergarage/puzzledb-operator/api/v1"
)

// PuzzleDBReconciler reconciles a PuzzleDB object
type PuzzleDBReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apiextensions.k8s.io.cybergarage.org,resources=puzzledbs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apiextensions.k8s.io.cybergarage.org,resources=puzzledbs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apiextensions.k8s.io.cybergarage.org,resources=puzzledbs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PuzzleDB object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *PuzzleDBReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Reconciling VisitorsApp", "Request.Namespace", req.Namespace, "Request.Name", req.Name)

	// Fetch the VisitorsApp instance
	v := &apiextensionsk8siov1.PuzzleDB{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, v)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after ctrl req.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the req.
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PuzzleDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensionsk8siov1.PuzzleDB{}).
		Complete(r)
}
