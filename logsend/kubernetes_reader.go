package logsend

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/clientcmd"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd"
)

// KubernetesReader read data from kubernetes api each 3 seconds
func KubernetesReader() {

	var filename string
	filename = "/home/dima/.kubernetes_vagrant_kubeconfig"
	loadRules := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: filename,
	}

	apiConfig, _ := loadRules.Load()
	fmt.Printf("Res: %+v\n", apiConfig)
	clientConfig := clientcmd.NewDefaultClientConfig(*apiConfig, &clientcmd.ConfigOverrides{})
	fmt.Printf("TEST1: %+v\n", clientConfig)

	f := cmd.NewFactory(clientConfig)
	buf := bytes.NewBuffer([]byte{})
	cmd := f.NewCmdGet(buf)
	cmd.SetOutput(buf)
	cmd.Run(cmd, []string{"pods"})
	fmt.Printf("RESULT: %s\n", buf)
	reader := bufio.NewReader(buf)
	// TODO read data each few seconds
	// TODO read data in sep. gorutine
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		fmt.Printf("res %s\n", line)
	}
	return
}
