package transforms

import (
	"time"

	machineryV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Extracts the common properties from a default k8s resource of unknown type and returns them in a map ready to be put in an Node
func commonProperties(resource machineryV1.Object) map[string]interface{} {
	ret := make(map[string]interface{})

	ret["name"] = resource.GetName()
	ret["selfLink"] = resource.GetSelfLink()
	ret["created"] = resource.GetCreationTimestamp().UTC().Format(time.RFC3339)

	if resource.GetLabels() != nil {
		ret["label"] = resource.GetLabels()
	}
	if resource.GetNamespace() != "" {
		ret["namespace"] = resource.GetNamespace()
	}
	return ret
}

// Transforms a resource of unknown type by simply pulling out the common properties.
func transformCommon(resource machineryV1.Object) Node {
	return Node{
		UID:        string(resource.GetUID()),
		Properties: commonProperties(resource),
	}
}

// Extracts the properties from a non-default k8s resource and returns them in a map ready to be put in an Node
func unstructuredProperties(resource *unstructured.Unstructured) map[string]interface{} {
	ret := make(map[string]interface{})

	ret["kind"] = resource.GetKind()
	ret["name"] = resource.GetName()
	ret["selfLink"] = resource.GetSelfLink()
	ret["created"] = resource.GetCreationTimestamp().UTC().Format(time.RFC3339)

	if resource.GetLabels() != nil {
		ret["label"] = resource.GetLabels()
	}
	if resource.GetNamespace() != "" {
		ret["namespace"] = resource.GetNamespace()
	}
	return ret

}

// Transforms an unstructured.Unstructured (which represents a non-default k8s object) into a Node
func transformUnstructured(resource *unstructured.Unstructured) Node {
	return Node{
		UID:        string(resource.GetUID()),
		Properties: unstructuredProperties(resource),
	}
}
