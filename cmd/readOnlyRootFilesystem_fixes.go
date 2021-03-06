package cmd

import k8sRuntime "k8s.io/apimachinery/pkg/runtime"

func fixReadOnlyRootFilesystem(resource k8sRuntime.Object, occurrence Occurrence) k8sRuntime.Object {
	var containers []Container
	for _, container := range getContainers(resource) {
		if occurrence.container == container.Name {
			container.SecurityContext.ReadOnlyRootFilesystem = newTrue()
		}
		containers = append(containers, container)
	}
	return setContainers(resource, containers)
}
