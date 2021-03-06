package model

import "time"

type GitInitConfig struct {
	SshKey string `json:"sshKey,omitempty"`
	GitUrl string `json:"gitUrl,omitempty"`
}

const CRD_YAML string = "apiVersion: apiextensions.k8s.io/v1beta1\n" +
	"kind: CustomResourceDefinition\n" +
	"metadata:\n" +
	"  name: c7nhelmreleases.choerodon.io\n" +
	"spec:\n" +
	"  group: choerodon.io\n" +
	"  names:\n" +
	"    kind: C7NHelmRelease\n" +
	"    listKind: C7NHelmReleaseList\n" +
	"    plural: c7nhelmreleases\n" +
	"    singular: c7nhelmrelease\n" +
	"  scope: Namespaced\n" +
	"  version: v1alpha1\n"

type AgentInitOptions struct {
	Envs    []EnvParas `json:"envs,omitempty"`
	GitHost string     `json:"gitHost,omitempty"`
}

type AgentStatus struct {
	EnvStatuses            []EnvStatus
	HelmStatus             string
	HelmOpDuration         time.Duration
	KubeStatus             string
	LastControllerSyncTime string
}

type EnvStatus struct {
	EnvCode       string
	EnvId         int32
	GitReady      bool
	GitOpDuration time.Duration
}

type EnvParas struct {
	Namespace string `json:"namespace,omitempty"`
	EnvId     int32  `json:"envId,omitempty"`
	GitRsaKey string `json:"gitRsaKey,omitempty"`
	GitUrl    string `json:"gitUrl,omitempty"`
}
