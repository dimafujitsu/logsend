package logsend

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/clientcmd"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd"
	"github.com/spf13/cobra"
	"time"
)

type Record struct {
	Kind              string      `json:"kind"`
	Creationtimestamp interface{} `json:"creationTimestamp"`
	Resourceversion   int         `json:"resourceVersion"`
	Apiversion        string      `json:"apiVersion"`
	Items             []struct {
		Kind              string `json:"kind"`
		ID                string `json:"id"`
		UID               string `json:"uid"`
		Creationtimestamp string `json:"creationTimestamp"`
		Selflink          string `json:"selfLink"`
		Resourceversion   int    `json:"resourceVersion"`
		Apiversion        string `json:"apiVersion"`
		Namespace         string `json:"namespace"`
		Generatename      string `json:"generateName"`
		Labels            struct {
			RunContainer string `json:"run-container"`
		} `json:"labels"`
		Desiredstate struct {
			Manifest struct {
				Version    string      `json:"version"`
				ID         string      `json:"id"`
				Volumes    interface{} `json:"volumes"`
				Containers []struct {
					Name  string `json:"name"`
					Image string `json:"image"`
					Ports []struct {
						Containerport int    `json:"containerPort"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
					Resources              struct{} `json:"resources"`
					Terminationmessagepath string   `json:"terminationMessagePath"`
					Imagepullpolicy        string   `json:"imagePullPolicy"`
					Capabilities           struct{} `json:"capabilities"`
				} `json:"containers"`
				Restartpolicy struct {
					Always struct{} `json:"always"`
				} `json:"restartPolicy"`
				Dnspolicy string `json:"dnsPolicy"`
			} `json:"manifest"`
			Host string `json:"host"`
		} `json:"desiredState"`
		Currentstate struct {
			Manifest struct {
				Version       string      `json:"version"`
				ID            string      `json:"id"`
				Volumes       interface{} `json:"volumes"`
				Containers    interface{} `json:"containers"`
				Restartpolicy struct{}    `json:"restartPolicy"`
			} `json:"manifest"`
			Status string `json:"status"`
			Host   string `json:"host"`
		} `json:"currentState"`
	} `json:"items"`
}

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
	cmd.Flags().Set("output", "json")
	rules := RuleLoad()
	for {
		time.Sleep(time.Millisecond * 1000)
		go getData(cmd, buf, rules)
	}
}

func getData(cmd *cobra.Command, buf *bytes.Buffer, rules []*Rule) {
	cmd.Run(cmd, []string{"pods"})
	reader := bufio.NewReader(buf)
	x := new(Record)
	err := json.NewDecoder(reader).Decode(x)
	if err != nil {
		fmt.Printf("Error json: %+v\n", err)
	} else {
		fmt.Printf("JSON: %+v\n", x)
		Sender.Send(x)
	}
}
