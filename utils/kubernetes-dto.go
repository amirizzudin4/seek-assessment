package utils

type KubernetesManifest struct {
	ApiVersion ApiVersion          `yaml:"apiVersion" json:"apiVersion"`
	Kind       Kind                `yaml:"kind" json:"kind"`
	Metadata   *KubernetesMetadata `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec       *KubernetesSpec     `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type KubernetesMetadata struct {
	Name      *string           `yaml:"name,omitempty" json:"name,omitempty"`
	Namespace *string           `yaml:"namespace,omitempty" json:"namespace,omitempty"`
	Labels    map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
}

type KubernetesSpec struct {
	// Deployment/ReplicaSet fields
	Replicas *int32           `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	Selector *LabelSelector   `yaml:"selector,omitempty" json:"selector,omitempty"`
	Template *PodTemplateSpec `yaml:"template,omitempty" json:"template,omitempty"`

	// Service fields
	Type  *string       `yaml:"type,omitempty" json:"type,omitempty"`
	Ports []ServicePort `yaml:"ports,omitempty" json:"ports,omitempty"`

	// Pod fields
	Containers    []Container       `yaml:"containers,omitempty" json:"containers,omitempty"`
	RestartPolicy *string           `yaml:"restartPolicy,omitempty" json:"restartPolicy,omitempty"`
	NodeSelector  map[string]string `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
}

type LabelSelector struct {
	MatchLabels map[string]string `yaml:"matchLabels,omitempty" json:"matchLabels,omitempty"`
}

type PodTemplateSpec struct {
	Metadata *KubernetesMetadata `yaml:"metadata,omitempty" json:"metadata,omitempty"`
	Spec     *PodSpec            `yaml:"spec,omitempty" json:"spec,omitempty"`
}

type PodSpec struct {
	Containers    []Container       `yaml:"containers,omitempty" json:"containers,omitempty"`
	RestartPolicy *RestartPolicy    `yaml:"restartPolicy,omitempty" json:"restartPolicy,omitempty"`
	NodeSelector  map[string]string `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
}

type Container struct {
	Name      string          `yaml:"name" json:"name"`
	Image     string          `yaml:"image" json:"image"`
	Ports     []ContainerPort `yaml:"ports,omitempty" json:"ports,omitempty"`
	Resources *ResourceSpec   `yaml:"resources,omitempty" json:"resources,omitempty"`
}

type ContainerPort struct {
	ContainerPort int32   `yaml:"containerPort" json:"containerPort"`
	Protocol      *string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
}

type ServicePort struct {
	Port       int32   `yaml:"port" json:"port"`
	TargetPort *int32  `yaml:"targetPort,omitempty" json:"targetPort,omitempty"`
	Protocol   *string `yaml:"protocol,omitempty" json:"protocol,omitempty"`
}

type ResourceSpec struct {
	Requests map[string]string `yaml:"requests,omitempty" json:"requests,omitempty"`
	Limits   map[string]string `yaml:"limits,omitempty" json:"limits,omitempty"`
}
