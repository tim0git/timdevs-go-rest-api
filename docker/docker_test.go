//go:build !docker

package docker

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

const tag = "vehicles-api"

func removeContainer(t *testing.T, id string) {
	cmd := shell.Command{
		Command: "docker",
		Args:    []string{"container", "rm", "--force", id},
	}

	shell.RunCommand(t, cmd)
}
func deleteImage(t *testing.T, tag string) {
	cmd := shell.Command{
		Command: "docker",
		Args:    []string{"image", "rm", "--force", tag},
	}

	shell.RunCommand(t, cmd)
}

func TestBuildsWithoutErrors(t *testing.T) {
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	err := docker.BuildE(t, "..", buildOptions)

	assert.Equal(t, nil, err)
}

func TestStartsWithoutErrors(t *testing.T) {
	name := "inspect-test-" + random.UniqueId()

	options := &docker.RunOptions{
		Detach: true,
		Name:   name,
	}

	id := docker.RunAndGetID(t, tag, options)
	defer removeContainer(t, id)

	c := docker.Inspect(t, id)

	exitCode0 := uint8(0x0)

	assert.Equal(t, exitCode0, c.ExitCode)
}

func TestHealthCheckReturnsStatus200(t *testing.T) {
	name := "inspect-test-" + random.UniqueId()

	options := &docker.RunOptions{
		Detach:               true,
		Name:                 name,
		EnvironmentVariables: []string{"PORT=8443"},
		OtherOptions:         []string{"-p", "8443:8443"},
	}

	id := docker.RunAndGetID(t, tag, options)
	defer removeContainer(t, id)
	defer deleteImage(t, tag)

	time.Sleep(1 * time.Second)

	response, err := http.Get("http://localhost:8443/health")

	assert.Equal(t, nil, err)
	assert.Equal(t, response.StatusCode, 200)
}
