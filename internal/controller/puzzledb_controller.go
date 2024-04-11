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

	apierrors "k8s.io/apimachinery/pkg/api/errors"
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

	// // Check if prometheus rule already exists, if not create a new one
	// foundRule := &monitoringv1.PrometheusRule{}
	// err := r.Get(ctx, types.NamespacedName{Name: ruleName, Namespace: namespace}, foundRule)
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

	// Fetch the PuzzleDB instance
	// The purpose is check if the Custom Resource for the Kind puzzledb
	// is applied on the cluster if not we return nil to stop the reconciliation
	puzzledb := &apiextensionsk8siov1.PuzzleDB{}
	err := r.Get(ctx, req.NamespacedName, puzzledb)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// If the custom resource is not found then, it usually means that it was deleted or not created
			// In this way, we will stop the reconciliation
			log.Info("puzzledb resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get puzzledb")
		return ctrl.Result{}, err
	}

	// // Let's just set the status as Unknown when no status are available
	// if puzzledb.Status.Conditions == nil || len(puzzledb.Status.Conditions) == 0 {
	// 	meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeAvailableMemcached, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
	// 	if err = r.Status().Update(ctx, puzzledb); err != nil {
	// 		log.Error(err, "Failed to update puzzledb status")
	// 		return ctrl.Result{}, err
	// 	}

	// 	// Let's re-fetch the puzzledb Custom Resource after update the status
	// 	// so that we have the latest state of the resource on the cluster and we will avoid
	// 	// raise the issue "the object has been modified, please apply
	// 	// your changes to the latest version and try again" which would re-trigger the reconciliation
	// 	// if we try to update it again in the following operations
	// 	if err := r.Get(ctx, req.NamespacedName, puzzledb); err != nil {
	// 		log.Error(err, "Failed to re-fetch puzzledb")
	// 		return ctrl.Result{}, err
	// 	}
	// }

	// // Let's add a finalizer. Then, we can define some operations which should
	// // occurs before the custom resource to be deleted.
	// // More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers
	// if !controllerutil.ContainsFinalizer(puzzledb, memcachedFinalizer) {
	// 	log.Info("Adding Finalizer for puzzledb")
	// 	if ok := controllerutil.AddFinalizer(puzzledb, memcachedFinalizer); !ok {
	// 		log.Error(err, "Failed to add finalizer into the custom resource")
	// 		return ctrl.Result{Requeue: true}, nil
	// 	}

	// 	if err = r.Update(ctx, puzzledb); err != nil {
	// 		log.Error(err, "Failed to update custom resource to add finalizer")
	// 		return ctrl.Result{}, err
	// 	}
	// }

	// // Check if the puzzledb instance is marked to be deleted, which is
	// // indicated by the deletion timestamp being set.
	// isMemcachedMarkedToBeDeleted := puzzledb.GetDeletionTimestamp() != nil
	// if isMemcachedMarkedToBeDeleted {
	// 	if controllerutil.ContainsFinalizer(puzzledb, memcachedFinalizer) {
	// 		log.Info("Performing Finalizer Operations for puzzledb before delete CR")

	// 		// Let's add here an status "Downgrade" to define that this resource begin its process to be terminated.
	// 		meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeDegradedMemcached,
	// 			Status: metav1.ConditionUnknown, Reason: "Finalizing",
	// 			Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", puzzledb.Name)})

	// 		if err := r.Status().Update(ctx, puzzledb); err != nil {
	// 			log.Error(err, "Failed to update puzzledb status")
	// 			return ctrl.Result{}, err
	// 		}

	// 		// Perform all operations required before remove the finalizer and allow
	// 		// the Kubernetes API to remove the custom resource.
	// 		r.doFinalizerOperationsForMemcached(puzzledb)

	// 		// TODO(user): If you add operations to the doFinalizerOperationsForMemcached method
	// 		// then you need to ensure that all worked fine before deleting and updating the Downgrade status
	// 		// otherwise, you should requeue here.

	// 		// Re-fetch the puzzledb Custom Resource before update the status
	// 		// so that we have the latest state of the resource on the cluster and we will avoid
	// 		// raise the issue "the object has been modified, please apply
	// 		// your changes to the latest version and try again" which would re-trigger the reconciliation
	// 		if err := r.Get(ctx, req.NamespacedName, puzzledb); err != nil {
	// 			log.Error(err, "Failed to re-fetch puzzledb")
	// 			return ctrl.Result{}, err
	// 		}

	// 		meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeDegradedMemcached,
	// 			Status: metav1.ConditionTrue, Reason: "Finalizing",
	// 			Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", puzzledb.Name)})

	// 		if err := r.Status().Update(ctx, puzzledb); err != nil {
	// 			log.Error(err, "Failed to update puzzledb status")
	// 			return ctrl.Result{}, err
	// 		}

	// 		log.Info("Removing Finalizer for puzzledb after successfully perform the operations")
	// 		if ok := controllerutil.RemoveFinalizer(puzzledb, memcachedFinalizer); !ok {
	// 			log.Error(err, "Failed to remove finalizer for puzzledb")
	// 			return ctrl.Result{Requeue: true}, nil
	// 		}

	// 		if err := r.Update(ctx, puzzledb); err != nil {
	// 			log.Error(err, "Failed to remove finalizer for puzzledb")
	// 			return ctrl.Result{}, err
	// 		}
	// 	}
	// 	return ctrl.Result{}, nil
	// }

	// // Check if the deployment already exists, if not create a new one
	// found := &appsv1.Deployment{}
	// err = r.Get(ctx, types.NamespacedName{Name: puzzledb.Name, Namespace: puzzledb.Namespace}, found)
	// if err != nil && apierrors.IsNotFound(err) {
	// 	// Define a new deployment
	// 	dep, err := r.deploymentForMemcached(puzzledb)
	// 	if err != nil {
	// 		log.Error(err, "Failed to define new Deployment resource for puzzledb")

	// 		// The following implementation will update the status
	// 		meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeAvailableMemcached,
	// 			Status: metav1.ConditionFalse, Reason: "Reconciling",
	// 			Message: fmt.Sprintf("Failed to create Deployment for the custom resource (%s): (%s)", puzzledb.Name, err)})

	// 		if err := r.Status().Update(ctx, puzzledb); err != nil {
	// 			log.Error(err, "Failed to update puzzledb status")
	// 			return ctrl.Result{}, err
	// 		}

	// 		return ctrl.Result{}, err
	// 	}

	// 	log.Info("Creating a new Deployment",
	// 		"Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	// 	if err = r.Create(ctx, dep); err != nil {
	// 		log.Error(err, "Failed to create new Deployment",
	// 			"Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	// 		return ctrl.Result{}, err
	// 	}

	// 	// Deployment created successfully
	// 	// We will requeue the reconciliation so that we can ensure the state
	// 	// and move forward for the next operations
	// 	return ctrl.Result{RequeueAfter: time.Minute}, nil
	// } else if err != nil {
	// 	log.Error(err, "Failed to get Deployment")
	// 	// Let's return the error for the reconciliation be re-trigged again
	// 	return ctrl.Result{}, err
	// }

	// // The CRD API is defining that the puzzledb type, have a MemcachedSpec.Size field
	// // to set the quantity of Deployment instances is the desired state on the cluster.
	// // Therefore, the following code will ensure the Deployment size is the same as defined
	// // via the Size spec of the Custom Resource which we are reconciling.
	// size := puzzledb.Spec.Size
	// if *found.Spec.Replicas != size {
	// 	// Increment MemcachedDeploymentSizeUndesiredCountTotal metric by 1
	// 	monitoring.MemcachedDeploymentSizeUndesiredCountTotal.Inc()
	// 	found.Spec.Replicas = &size
	// 	if err = r.Update(ctx, found); err != nil {
	// 		log.Error(err, "Failed to update Deployment",
	// 			"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

	// 		// Re-fetch the puzzledb Custom Resource before update the status
	// 		// so that we have the latest state of the resource on the cluster and we will avoid
	// 		// raise the issue "the object has been modified, please apply
	// 		// your changes to the latest version and try again" which would re-trigger the reconciliation
	// 		if err := r.Get(ctx, req.NamespacedName, puzzledb); err != nil {
	// 			log.Error(err, "Failed to re-fetch puzzledb")
	// 			return ctrl.Result{}, err
	// 		}

	// 		// The following implementation will update the status
	// 		meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeAvailableMemcached,
	// 			Status: metav1.ConditionFalse, Reason: "Resizing",
	// 			Message: fmt.Sprintf("Failed to update the size for the custom resource (%s): (%s)", puzzledb.Name, err)})

	// 		if err := r.Status().Update(ctx, puzzledb); err != nil {
	// 			log.Error(err, "Failed to update puzzledb status")
	// 			return ctrl.Result{}, err
	// 		}

	// 		return ctrl.Result{}, err
	// 	}

	// 	// Now, that we update the size we want to requeue the reconciliation
	// 	// so that we can ensure that we have the latest state of the resource before
	// 	// update. Also, it will help ensure the desired state on the cluster
	// 	return ctrl.Result{Requeue: true}, nil
	// }

	// // The following implementation will update the status
	// meta.SetStatusCondition(&puzzledb.Status.Conditions, metav1.Condition{Type: typeAvailableMemcached,
	// 	Status: metav1.ConditionTrue, Reason: "Reconciling",
	// 	Message: fmt.Sprintf("Deployment for custom resource (%s) with %d replicas created successfully", puzzledb.Name, size)})

	// if err := r.Status().Update(ctx, puzzledb); err != nil {
	// 	log.Error(err, "Failed to update puzzledb status")
	// 	return ctrl.Result{}, err
	// }

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
