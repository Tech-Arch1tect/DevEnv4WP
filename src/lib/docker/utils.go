package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/tech-arch1tect/DevEnv4WP/lib/utils"
)

func pullDockerImage(image string) error {
	ctx := context.Background()
	client, err := GetClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// check if image exists
	_, _, err = client.ImageInspectWithRaw(
		ctx,
		image,
	)
	if err == nil {
		utils.DebugLog("pullDockerImage (ImageInspectWithRaw) image exists: " + image)
		return nil
	}
	utils.DebugLog("pullDockerImage (ImageInspectWithRaw) image does not exist: " + image + ". Pulling image")

	// Pull image
	out, err := client.ImagePull(
		ctx,
		image,
		types.ImagePullOptions{},
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
		types.ContainerRemoveOptions{},
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

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{
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
