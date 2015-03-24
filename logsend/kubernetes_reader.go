package logsend

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/clientcmd"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd"
	"github.com/spf13/cobra"
	"time"
)

// KubernetesReader read data from kubernetes api each 3 seconds
func KubernetesReader() {

	var filename string
	filename = "/home/dima/.kubernetes_vagrant_kubeconfig"
	loadRules := &clientcmd.ClientConfigLoadingRules{
		ExplicitPath: filename,
	}

	apiConfig, _ := loadRules.Load()
	clientConfig := clientcmd.NewDefaultClientConfig(*apiConfig, &clientcmd.ConfigOverrides{})

	f := cmd.NewFactory(clientConfig)
	buf := bytes.NewBuffer([]byte{})
	cmd := f.NewCmdGet(buf)
	cmd.SetOutput(buf)
	for {
		time.Sleep(time.Millisecond * 1000)
		go getData(cmd, buf)
	}
	return
}

func getData(cmd *cobra.Command, buf *bytes.Buffer) {
	cmd.Run(cmd, []string{"pods"})
	fmt.Printf("RESULT: %s\n", buf)
	reader := bufio.NewReader(buf)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			break
		}
		fmt.Printf("res %s\n", line)
	}
}
