package executor

import (
    "log"
    "os/exec"
)

// DeployFunction builds and deploys a Docker image for a function
func DeployFunction(imageName string, dockerfilePath string) error {
    log.Printf("Deploying function: %s from %s", imageName, dockerfilePath)
    cmd := exec.Command("docker", "build", "-t", imageName, dockerfilePath)
    cmd.Stdout = log.Writer()
    cmd.Stderr = log.Writer()
    return cmd.Run()
}
