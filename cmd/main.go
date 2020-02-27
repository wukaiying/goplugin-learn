package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/wukaiying/goplugin-learn/pkg/shared"

	"github.com/hashicorp/go-plugin"
)

func main()  {
	log.SetOutput(ioutil.Discard)
	//we are host, launching the plugin process
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: shared.PluginMap,
		Cmd: exec.Command("sh", "-c", "/Users/wukaiying/go/src/goplugin-learn/pluginclient/pluginclient"),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC,
		},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("kv_grpc")
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	kv := raw.(shared.KV)
	os.Args = os.Args[1:]
	//os.Args = []string{"put","wky", "111"}
	switch os.Args[0] {
	case "get":
		result,err := kv.Get(os.Args[1])
		if err != nil {
			fmt.Println("Error", err.Error())
			os.Exit(1)
		}
		fmt.Println(string(result))
	case "put":
		err := kv.Put(os.Args[1], []byte(os.Args[2]))
		if err != nil {
			fmt.Println("Error", err.Error())
			os.Exit(1)
		}
	default:
		fmt.Printf("Please only use 'get' or 'put', given: %q", os.Args[0])
		os.Exit(1)
	}
	os.Exit(0)
}
