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
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiextensionsk8siov1 "github.com/cybergarage/puzzledb-operator/api/v1"
)

// PuzzleDBReconciler reconciles a PuzzleDB object
type PuzzleDBReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
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

	log.Info("Reconciling", "Request.Namespace", req.Namespace, "Request.Name", req.Name)

	// Check if prometheus rule already exists, if not create a new one
	// foundRule := &monitoringv1.PrometheusRule{}
	// err := r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, foundRule)
	// if err != nil && apierrors.IsNotFound(err) {
	// 	// Define a new prometheus rule
	// 	prometheusRule := monitoring.NewPrometheusRule(namespace)
	// 	if err := r.Create(ctx, prometheusRule); err != nil {
	// 		log.Error(err, "Failed to create prometheus rule")
	// 		return ctrl.Result{}, nil
	// 	}
	// }

	// if err == nil {
	// 	// Check if prometheus rule spec was changed, if so set as desired
	// 	desiredRuleSpec := monitoring.NewPrometheusRuleSpec()
	// 	if !reflect.DeepEqual(foundRule.Spec.DeepCopy(), desiredRuleSpec) {
	// 		desiredRuleSpec.DeepCopyInto(&foundRule.Spec)
	// 		if r.Update(ctx, foundRule); err != nil {
	// 			log.Error(err, "Failed to update prometheus rule")
	// 			return ctrl.Result{}, nil
	// 		}
	// 	}
	// }

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

// finalizeMemcached will perform the required operations before delete the CR.
func (r *PuzzleDBReconciler) doFinalizerOperationsForMemcached(cr *apiextensionsk8siov1.PuzzleDB) {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted. Examples
	// of finalizers include performing backups and deleting
	// resources that are not owned by this CR, like a PVC.

	// Note: It is not recommended to use finalizers with the purpose of delete resources which are
	// created and managed in the reconciliation. These ones, such as the Deployment created on this reconcile,
	// are defined as depended of the custom resource. See that we use the method ctrl.SetControllerReference.
	// to set the ownerRef which means that the Deployment will be deleted by the Kubernetes API.
	// More info: https://kubernetes.io/docs/tasks/administer-cluster/use-cascading-deletion/

	// The following implementation will raise an event
	r.Recorder.Event(cr, "Warning", "Deleting",
		fmt.Sprintf("Custom Resource %s is being deleted from the namespace %s",
			cr.Name,
			cr.Namespace))
}

// SetupWithManager sets up the controller with the Manager.
func (r *PuzzleDBReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiextensionsk8siov1.PuzzleDB{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 1}).
		Complete(r)
}
