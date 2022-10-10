package queryfrontend

import (
	"testing"
)

func TestBuildQueryFrontend(t *testing.T) {
	//objects := BuildQueryFrontend(v1alpha1.Microservices{
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name:      "test",
	//		Namespace: "project1",
	//	},
	//})
	//
	//// labels := manifestutils.ComponentLabels("query-frontend", "test")
	//assert.Equal(t, 3, len(objects))
	//
	//// Test the services
	//frontendService := objects[1].(*corev1.Service)
	//assert.Equal(t, "tempo-test-query-frontend", frontendService.Name)
	//assert.Len(t, frontendService.Spec.Ports, 4)
	//// TODO check port values
	//
	//frontEndDiscoveryService := objects[2].(*corev1.Service)
	//assert.Equal(t, "tempo-test-query-frontend-discovery", frontEndDiscoveryService.Name)
	//assert.Len(t, frontEndDiscoveryService.Spec.Ports, 5)
	//// TODO check port values
	//
	//deployment := objects[0].(*v1.Deployment)
	//assert.Equal(t, "tempo-test-query-frontend", deployment.Name)
	//assert.Len(t, deployment.Spec.Template.Spec.Containers, 2)

}
