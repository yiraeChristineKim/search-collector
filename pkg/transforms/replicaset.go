package transforms

import (
	v1 "k8s.io/api/apps/v1"
)

// Takes a *v1.ReplicaSet and yields a Node
func TransformReplicaSet(resource *v1.ReplicaSet) Node {

	replicaSet := TransformCommon(resource) // Start off with the common properties

	// Extract the properties specific to this type
	replicaSet.Properties["kind"] = "ReplicaSet"
	replicaSet.Properties["current"] = resource.Status.Replicas
	replicaSet.Properties["desired"] = resource.Spec.Replicas

	return replicaSet
}