package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func pullDockerImage(i string) error {
	ctx := context.Background()
	client, err := GetClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// check if image exists
	_, _, err = client.ImageInspectWithRaw(
		ctx,
		i,
	)
	if err == nil {
		utils.DebugLog("pullDockerImage (ImageInspectWithRaw) image exists: " + i)
		return nil
	}
	utils.DebugLog("pullDockerImage (ImageInspectWithRaw) image does not exist: " + i + ". Pulling image")

	// Pull image
	out, err := client.ImagePull(
		ctx,
		i,
		image.PullOptions{},
	)
	if err != nil {
		utils.DebugLog("pullDockerImage (ImagePull) error" + err.Error())
		return err
	}
	defer out.Close()

	output, err := io.ReadAll(out)
	if err != nil {
		utils.DebugLog("pullDockerImage (ReadAll) error" + err.Error())
		return err
	}
	utils.DebugLog("pullDockerImage (ReadAll) output: " + string(output))
	utils.DebugLog("pullDockerImage (ImagePull) success")

	return nil
}

func removeContainerIfExists(containerName string) error {
	conatinerID, err := getContainerIDbyName(containerName)
	if err != nil {
		utils.DebugLog("Remove container error (get container ID): " + err.Error())
		return err
	}

	ctx := context.Background()
	client, err := GetClient()
	if err != nil {
		utils.DebugLog("Remove container error (get docker client): " + err.Error())
		return err
	}
	defer client.Close()

	err = client.ContainerRemove(
		ctx,
		conatinerID,
		container.RemoveOptions{},
	)
	if err != nil {
		utils.DebugLog("Remove container error (remove container): " + err.Error())
		return err
	}

	return nil
}

func getContainerIDbyName(containerName string) (string, error) {
	ctx := context.Background()
	client, err := GetClient()
	if err != nil {
		utils.DebugLog("Get container ID error (get docker client): " + err.Error())
		return "", err
	}
	defer client.Close()

	containers, err := client.ContainerList(ctx, container.ListOptions{
		All: true,
	})
	if err != nil {
		utils.DebugLog("Get container ID error (list containers): " + err.Error())
		return "", err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == "/"+containerName {
				return container.ID, nil
			}
		}
	}

	return "", fmt.Errorf("container not found")
}
