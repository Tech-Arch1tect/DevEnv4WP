package docker

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func RunContainerWithCommand(command string, uid int, gid int, image string, name string, binds []string) error {
	removeContainerIfExists(name)

	ctx := context.Background()
	utils.DebugLog("Pulling docker image: " + image)
	err := pullDockerImage(image)
	if err != nil {
		utils.DebugLog("Error pulling docker image: " + image)
		return err
	}
	utils.DebugLog("Pulled docker image" + image)

	client, err := GetClient()
	if err != nil {
		utils.DebugLog("Error getting docker client")
		return err
	}
	defer client.Close()

	utils.DebugLog("Creating container: " + name)
	runContainer, err := client.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			Cmd:   strings.Fields(command),
			Tty:   true,
			User:  fmt.Sprintf("%d:%d", uid, gid),
		},
		&container.HostConfig{
			AutoRemove: false,
			Binds:      binds,
		},
		nil,
		nil,
		name,
	)
	if err != nil {
		utils.DebugLog("Error creating container: " + name)
		return err
	}

	err = client.ContainerStart(
		ctx,
		runContainer.ID,
		container.StartOptions{},
	)
	if err != nil {
		utils.DebugLog("Error starting container: " + name)
		return err
	}
	utils.DebugLog("Started container: " + name)

	utils.DebugLog("Waiting for container to finish: " + name)
	statusCh, errCh := client.ContainerWait(
		ctx,
		runContainer.ID,
		container.WaitConditionNotRunning,
	)
	select {
	case err := <-errCh:
		if err != nil {
			utils.DebugLog("Error waiting for container to finish: " + name)
			return err
		}
	case <-statusCh:
	}

	utils.DebugLog("Container finished, getting logs: " + name)
	out, err := client.ContainerLogs(
		ctx,
		runContainer.ID,
		container.LogsOptions{ShowStdout: true},
	)
	if err != nil {
		utils.DebugLog("Error getting container logs: " + name)
		return err
	}
	utils.DebugLog("Got container logs: " + name)
	defer out.Close()
	output, err := io.ReadAll(out)
	if err != nil {
		utils.DebugLog("Error reading container logs: " + name)
		return err
	}
	utils.DebugLog("DEBUG: Container (" + name + ") logs: " + string(output))

	err = removeContainerIfExists(name)
	if err != nil {
		utils.DebugLog("Error removing container: " + name)
		return err
	}

	return err
}
