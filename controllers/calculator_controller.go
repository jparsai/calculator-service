/*
Copyright 2023.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mathcomv1alpha1 "github.com/math-com/calculator-service/api/v1alpha1"
)

// CalculatorReconciler reconciles a Calculator object
type CalculatorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=math.com,resources=calculators,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=math.com,resources=calculators/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=math.com,resources=calculators/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Calculator object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *CalculatorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// Retrieve the latest version of the resource from ze cluster
	calculator := mathcomv1alpha1.Calculator{}
	err := r.Client.Get(ctx, req.NamespacedName, &calculator)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Sum up the numbers from spec
	sum := calculator.Spec.NumberOne + calculator.Spec.NumberTwo

	// Store in status
	calculator.Status.Sum = sum
	err = r.Client.Status().Update(ctx, &calculator)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CalculatorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mathcomv1alpha1.Calculator{}).
		Complete(r)
}
