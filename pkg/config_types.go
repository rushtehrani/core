package v1

import (
	"fmt"
	"github.com/onepanelio/core/pkg/util/ptr"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

// SystemConfig is configuration loaded from kubernetes config and secrets that includes information about the
// database, server, etc.
type SystemConfig map[string]string

type NodePoolOption struct {
	ParameterOption
	Resources corev1.ResourceRequirements
}

// GetValue returns the value in the underlying map if it exists, otherwise nil is returned
// If the value does not exist, it is also logged.
func (s SystemConfig) GetValue(name string) *string {
	value, ok := s[name]
	if !ok {
		log.WithFields(log.Fields{
			"Method": "SystemConfig.GetValue",
			"Name":   name,
			"Error":  "does not exist",
		})

		return nil
	}

	return &value
}

// Domain gets the ONEPANEL_DOMAIN value, or nil.
func (s SystemConfig) Domain() *string {
	return s.GetValue("ONEPANEL_DOMAIN")
}

// APIURL gets the ONEPANEL_API_URL, or nil.
func (s SystemConfig) APIURL() *string {
	return s.GetValue("ONEPANEL_API_URL")
}

// APIProtocol returns either http:// or https:// or nil.
// It is based on the ONEPANEL_API_URL config value and checks if it has https or http
func (s SystemConfig) APIProtocol() *string {
	url := s.APIURL()
	if url == nil {
		return nil
	}

	if strings.HasPrefix(*url, "https://") {
		return ptr.String("https://")
	}

	return ptr.String("http://")
}

// FQDN gets the ONEPANEL_FQDN value or nil.
func (s SystemConfig) FQDN() *string {
	return s.GetValue("ONEPANEL_FQDN")
}

// NodePoolLabel gets the applicationNodePoolLabel from the config or returns nil.
func (s SystemConfig) NodePoolLabel() (label *string) {
	return s.GetValue("applicationNodePoolLabel")
}

// NodePoolOptions loads and parses the applicationNodePoolOptions from the config.
// If there is no data, an error is returned.
func (s SystemConfig) NodePoolOptions() (options []*NodePoolOption, err error) {
	data := s.GetValue("applicationNodePoolOptions")
	if data == nil {
		return nil, fmt.Errorf("no nodePoolOptions in config")
	}

	if err = yaml.Unmarshal([]byte(*data), &options); err != nil {
		return
	}

	return
}

func (s SystemConfig) NodePoolOptionByValue(value string) (option *NodePoolOption, err error) {
	options, err := s.NodePoolOptions()
	if err != nil {
		return
	}
	for _, opt := range options {
		if opt.Value == value {
			option = opt
			return
		}
	}
	return
}

// DatabaseDriverName gets the databaseDriverName value, or nil.
func (s SystemConfig) DatabaseDriverName() *string {
	return s.GetValue("databaseDriverName")
}

type ArtifactRepositoryS3Config struct {
	KeyFormat       string
	Bucket          string
	Endpoint        string
	Insecure        bool
	Region          string
	AccessKeySecret corev1.SecretKeySelector
	SecretKeySecret corev1.SecretKeySelector
	AccessKey       string
	Secretkey       string
}

// FormatKey replaces placeholder values with their actual values and returns this string.
// {{workflow.namespace}} -> namespace
// {{workflow.name}} -> workflowName
// {{pod.name}} -> podName
func (a *ArtifactRepositoryS3Config) FormatKey(namespace, workflowName, podName string) string {
	keyFormat := a.KeyFormat

	keyFormat = strings.Replace(keyFormat, "{{workflow.namespace}}", namespace, -1)
	keyFormat = strings.Replace(keyFormat, "{{workflow.name}}", workflowName, -1)
	keyFormat = strings.Replace(keyFormat, "{{pod.name}}", podName, -1)

	return keyFormat
}

type ArtifactRepositoryConfig struct {
	S3 *ArtifactRepositoryS3Config
}

type NamespaceConfig struct {
	ArtifactRepository ArtifactRepositoryConfig
}
