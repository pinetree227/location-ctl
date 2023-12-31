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

package ctl

import (
	"context"
        "fmt"
        "strconv"
	"regexp"
	"math"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        appsv1 "k8s.io/api/apps/v1"
        corev1 "k8s.io/api/core/v1"
        "k8s.io/apimachinery/pkg/api/equality"
        "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
        "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
        "k8s.io/apimachinery/pkg/runtime"
        //"k8s.io/apimachinery/pkg/util/intstr"
        appsv1apply "k8s.io/client-go/applyconfigurations/apps/v1"
        corev1apply "k8s.io/client-go/applyconfigurations/core/v1"
        metav1apply "k8s.io/client-go/applyconfigurations/meta/v1"
        "k8s.io/client-go/tools/record"
        "k8s.io/utils/pointer"
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/client"
        "sigs.k8s.io/controller-runtime/pkg/client/apiutil"
        //"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
        "sigs.k8s.io/controller-runtime/pkg/log"

	pinev1 "github.com/pinetree227/location-ctl/api/ctl/v1"
)

// LocationCtlReconciler reconciles a LocationCtl object
type LocationCtlReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Recorder record.EventRecorder
}



//+kubebuilder:rbac:groups=ctl.pinetree227.github.io,resources=locationctls,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ctl.pinetree227.github.io,resources=locationctls/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ctl.pinetree227.github.io,resources=locationctls/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the LocationCtl object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
/*
func (r *LocationCtlReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}
*/
func (r *LocationCtlReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
        logger := log.FromContext(ctx)

        //! [call-remove-metrics]
        var mdView pinev1.LocationCtl
        err := r.Get(ctx, req.NamespacedName, &mdView)
        if errors.IsNotFound(err) {
                return ctrl.Result{}, nil
        }
        //! [call-remove-metrics]
        if err != nil {
                logger.Error(err, "unable to get MarkdownView", "name", req.NamespacedName)
                return ctrl.Result{}, err
        }

        if !mdView.ObjectMeta.DeletionTimestamp.IsZero() {
                return ctrl.Result{}, nil
        }
	if mdView.Spec.Update == 1{
		var pod appsv1.Deployment
		err := r.Get(ctx, client.ObjectKey{Namespace: "default", Name: "viewer-" + mdView.Name}, &pod)
		        if err != nil {
				return ctrl.Result{}, err
			}
		nodeName := pod.Spec.Template.Spec.NodeName
                        value1 := mdView.Spec.PodX
                        if value1 != "" {
                                if _, err := strconv.Atoi(value1); err != nil {
                                        return ctrl.Result{}, err

                                }
                        }
                        value2 := mdView.Spec.PodY
                        if value2 != "" {
                                if _, err := strconv.Atoi(value2); err != nil {
                                  return ctrl.Result{}, err

                         }
                 }
                        podx, _ := strconv.ParseFloat(value1,64)
                        pody, _ := strconv.ParseFloat(value2,64)
                        y,x,err := extractNumbers(nodeName)
                        if err != nil {
                                        fmt.Printf("Error getting last digit from node name %s: %v\n", nodeName, err)
                                         return ctrl.Result{}, err

                                }
                                dest := (x - podx) * (x - podx) + (y - pody) * (y - pody)
                                dest = math.Sqrt(float64(dest))
			if mdView.Spec.Apptype == "RealTime"{
				if dest > 250 {
    uid := pod.GetUID()
    resourceVersion := pod.GetResourceVersion()
    cond := metav1.Preconditions{
        UID:             &uid,
        ResourceVersion: &resourceVersion,
    }
    err = r.Delete(ctx, &pod, &client.DeleteOptions{
        Preconditions: &cond,
    })

				}
			} else {
				if dest > 500 {
    uid := pod.GetUID()
    resourceVersion := pod.GetResourceVersion()
    cond := metav1.Preconditions{
        UID:             &uid,
        ResourceVersion: &resourceVersion,
    }
    err = r.Delete(ctx, &pod, &client.DeleteOptions{
        Preconditions: &cond,
    })

			}
	}
}
/*
        err = r.reconcileConfigMap(ctx, mdView)
        if err != nil {
                return ctrl.Result{}, err
        }
*/
        err = r.reconcileDeployment(ctx, mdView)
        if err != nil {
                return ctrl.Result{}, err
        }
/*
        err = r.reconcileService(ctx, mdView)
        if err != nil {
                return ctrl.Result{}, err
        }
*/

        return  ctrl.Result{}, nil
}

func (r *LocationCtlReconciler) reconcileDeployment(ctx context.Context, mdView pinev1.LocationCtl) error {
        logger := log.FromContext(ctx)

        depName := "viewer-" + mdView.Name

        /*      viewerImage := "peaceiris/mdbook:latest"
        if len(mdView.Spec.ViewerImage) != 0 {
                viewerImage = mdView.Spec.ViewerImage
        }
        */

        owner, err := controllerReference(mdView, r.Scheme)
        if err != nil {
                return err
        }
        dep := appsv1apply.Deployment(depName, mdView.Namespace).
	WithName(depName).
                WithOwnerReferences(owner).
               WithSpec(appsv1apply.DeploymentSpec().
            WithReplicas(1).
            WithSelector(metav1apply.LabelSelector().WithMatchLabels(map[string]string{"app": "nginx"})).
            WithTemplate(corev1apply.PodTemplateSpec().
                WithLabels(map[string]string{
                        "app": "nginx",
                        "podx": mdView.Spec.PodX,
                        "pody": mdView.Spec.PodY,
			"apptype": mdView.Spec.Apptype,
                }).
                WithSpec(corev1apply.PodSpec().
                    WithContainers(corev1apply.Container().
                        WithName("nginx").
                        WithImage("nginx:latest").
			WithResources(corev1apply.ResourceRequirements().
			WithRequests(corev1.ResourceList{
				corev1.ResourceMemory: resource.MustParse("1024Mi"),
				corev1.ResourceCPU : resource.MustParse("500m"),
			}),
			),
                    ),
                ),
            ),
        )
        obj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(dep)
        if err != nil {
                return err
        }
        patch := &unstructured.Unstructured{
                Object: obj,
        }

        var current appsv1.Deployment
        err = r.Get(ctx, client.ObjectKey{Namespace: mdView.Namespace, Name: depName}, &current)
        if err != nil && !errors.IsNotFound(err) {
                return err
        }

        currApplyConfig, err := appsv1apply.ExtractDeployment(&current, "location-ctl-controller")
        if err != nil {
                return err
        }

        if equality.Semantic.DeepEqual(dep, currApplyConfig) {
                return nil
        }

        err = r.Patch(ctx, patch, client.Apply, &client.PatchOptions{
                FieldManager: "location-ctl-controller",
                Force:        pointer.Bool(true),
        })

        if err != nil {
                logger.Error(err, "unable to create or update Deployment")
                return err
        }
        logger.Info("reconcile Deployment successfully", "name", mdView.Name)
        return nil
}
func controllerReference(mdView pinev1.LocationCtl, scheme *runtime.Scheme) (*metav1apply.OwnerReferenceApplyConfiguration, error) {
        gvk, err := apiutil.GVKForObject(&mdView, scheme)
        if err != nil {
                return nil, err
        }
        ref := metav1apply.OwnerReference().
                WithAPIVersion(gvk.GroupVersion().String()).
                WithKind(gvk.Kind).
                WithName(mdView.Name).
                WithName(mdView.Name).
                WithUID(mdView.GetUID()).
                WithBlockOwnerDeletion(true).
                WithController(true)
        return ref, nil
}

func extractNumbers(inputString string) (float64, float64, error) {
        // 正規表現のパターン
        pattern := `(\d{2})(\d{2})$`

        // 正規表現に一致するか確認
        re := regexp.MustCompile(pattern)
        matches := re.FindStringSubmatch(inputString)

        if len(matches) != 3 {
                return 0, 0, fmt.Errorf("数値に変換できません")

        }

        // 文字列から数値に変換
        firstTwo, err := strconv.ParseFloat(matches[1],64)
        if err != nil {
                return 0, 0, fmt.Errorf("数値に変換できません: %v", err)
        }

        lastTwo, err := strconv.ParseFloat(matches[2],64)
        if err != nil {
                return 0, 0, fmt.Errorf("数値に変換できません: %v", err)
        }

        return firstTwo*8.66, lastTwo*5, nil
}



// SetupWithManager sets up the controller with the Manager.
func (r *LocationCtlReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pinev1.LocationCtl{}).
		Complete(r)
	}
