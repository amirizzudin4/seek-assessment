package utils

type ApiVersion string

const (
	ApiVersionV1                ApiVersion = "v1"
	ApiVersionAppsV1            ApiVersion = "apps/v1"
	ApiVersionExtensionsV1Beta1 ApiVersion = "extensions/v1beta1"
	ApiVersionNetworkingV1      ApiVersion = "networking.k8s.io/v1"
	ApiVersionRbacV1            ApiVersion = "rbac.authorization.k8s.io/v1"
	ApiVersionStorageV1         ApiVersion = "storage.k8s.io/v1"
	ApiVersionBatchV1           ApiVersion = "batch/v1"
	ApiVersionBatchV1Beta1      ApiVersion = "batch/v1beta1"
	ApiVersionAutoscalingV1     ApiVersion = "autoscaling/v1"
	ApiVersionAutoscalingV2     ApiVersion = "autoscaling/v2"
)

// Kind enum type
type Kind string

const (
	KindPod                   Kind = "Pod"
	KindService               Kind = "Service"
	KindDeployment            Kind = "Deployment"
	KindReplicaSet            Kind = "ReplicaSet"
	KindConfigMap             Kind = "ConfigMap"
	KindSecret                Kind = "Secret"
	KindIngress               Kind = "Ingress"
	KindNamespace             Kind = "Namespace"
	KindPersistentVolume      Kind = "PersistentVolume"
	KindPersistentVolumeClaim Kind = "PersistentVolumeClaim"
	KindJob                   Kind = "Job"
	KindCronJob               Kind = "CronJob"
	KindDaemonSet             Kind = "DaemonSet"
	KindStatefulSet           Kind = "StatefulSet"
)

type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)
