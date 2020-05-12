package patch

import (
	"bufio"
	"bytes"
	"text/template"
)

const secretSidecarTemplate = `
{
  "name": "secret-sidecar",
  "image": "{{ .ContainerImage }}",
  "securityContext": {
    "runAsUser": 1337
  },
  "ports": [
    {
      "containerPort": 9901,
      "name": "stats",
      "protocol": "TCP"
    }
  ],
  "env": [
    {
      "name": "SECRET_NAME",
      "value": "{{ .SecretName }}"
    },
    {
      "name": "AWS_REGION",
      "value": "{{ .Region }}"
    }
  ],
  "resources": {
    "requests": {
      "cpu": "{{ .CpuRequests }}",
      "memory": "{{ .MemoryRequests }}"
    }
  }
}
`

type SidecarMeta struct {
	ContainerImage string
	SecretName     string
	Region         string
	CpuRequests    string
	MemoryRequests string
}

func renderSidecars(meta SidecarMeta) ([]string, error) {
	var sidecars []string

	secretSidecar, err := renderTemplate("secret-sidecar", envoyContainerTemplate, meta)
	if err != nil {
		return sidecars, err
	}

	sidecars = append(sidecars, secretSidecar)
	return sidecars, nil
}

func renderTemplate(name string, t string, meta SidecarMeta) (string, error) {
	tmpl, err := template.New(name).Parse(t)
	if err != nil {
		return "", err
	}
	var data bytes.Buffer
	b := bufio.NewWriter(&data)

	if err := tmpl.Execute(b, meta); err != nil {
		return "", err
	}
	err = b.Flush()
	if err != nil {
		return "", err
	}
	return data.String(), nil
}
