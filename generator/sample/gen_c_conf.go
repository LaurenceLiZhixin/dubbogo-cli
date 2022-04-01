package sample

const (
	clientConfigFile = `dubbo:
  registries:
    demoZK:
      protocol: zookeeper
      timeout: 3s
      address: 127.0.0.1:2181
  consumer:
    references:
      GreeterClientImpl:
        protocol: tri
        interface: com.apache.dubbo.sample.basic.IGreeter # must be compatible with grpc or dubbo-java`
)

func init() {
	fileMap["clientConfigFile"] = &fileGenerator{
		path:    "./go-client/conf",
		file:    "dubbogo.yml",
		context: clientConfigFile,
	}
}