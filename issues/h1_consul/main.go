package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"github.com/leonelquinteros/gorand"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	consulCheck()
	//curlCheck()

}

// GenerateUUID generates v4 uuid in text form
func GenerateUUID() (string, error) {
	uuid, err := gorand.UUIDv4()
	if err != nil {
		return "", err
	}

	uuidStr, err := gorand.MarshalUUID(uuid)
	if err != nil {
		return "", err
	}
	return uuidStr, nil
}

// VerifyUUID verifies uuid in text form
func VerifyUUID(uuid string) bool {
	_, err := gorand.UnmarshalUUID(uuid)
	if err != nil {
		return false
	}
	return true
}

func consulCheck() {
	fmt.Println("start consul")
	var consulRegistrar sd.Registrar
	consulClient, err := MakeClient("localhost", "8500")
	if err != nil {
		panic(err)
	}
	serviceUUID, err := GenerateUUID()
	if err != nil {
		panic(err)
	}

	portInt := 8500

	agent := api.AgentServiceRegistration{
		ID:      serviceUUID,
		Name:    "some_name_to_debug",
		Tags:    []string{},
		Port:    portInt,
		Address: "localhost",
		Check: &api.AgentServiceCheck{
			Name:     fmt.Sprintf("%s health check on %s", "some_name_to_debug", net.JoinHostPort("localhost", "8500")),
			HTTP:     "http://" + net.JoinHostPort("localhost", "8500") + "/healthcheck",
			Interval: "10s",
			Timeout:  "10s",
			Method:   http.MethodGet,
		},
	}
	w := log.NewSyncWriter(os.Stderr)
	logger := log.NewLogfmtLogger(w)
	consulRegistrar = consulsd.NewRegistrar(consulClient, &agent, logger)
	consulRegistrar.Register()
	fmt.Println("finished consul")
}

func curlCheck() {
	fmt.Println("Starting to make a curl req")
	command := "curl https://127.0.0.1:8500/v1/agent/members"
	cmcc := exec.Command("bash", "-c", command)
	data, err := cmcc.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("command out")
	fmt.Println(string(data))
	fmt.Println("Finish")
}

const defaultIndex = 0

type config struct {
	// consul=======
	client      consul.Client
	service     string
	tags        []string
	passingOnly bool
}

//MakeClient is used to make wrapper from std consul client to gokit
func MakeClient(consulHost, consulPort string) (consulsd.Client, error) {
	consulConfig := api.DefaultConfig()

	consulConfig.Address = net.JoinHostPort(consulHost, consulPort)
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new api client error")
	}
	consulClient := consulsd.NewClient(client)
	return consulClient, nil
}

//func (c *config) getInstances() error {
//	tag := ""
//	if len(c.tags) > 0 {
//		// Consul doesn't support more than one tag in its service query method.
//		// https://github.com/hashicorp/consul/issues/294
//		// Hashi suggest prepared queries, but they don't support blocking.
//		// https://www.consul.io/docs/agent/http/query.html#execute
//		// If we want blocking for efficiency, we must filter tags manually.
//		// So for that we are using `filterEntries` function
//		tag = c.tags[0]
//	}
//
//	entries, _, err := c.client.Service(c.service, tag, c.passingOnly, nil)
//	if err != nil {
//		return err
//	}
//	if len(c.tags) > 1 {
//		entries = filterEntries(entries, c.tags[1:]...)
//	}
//
//	instances := makeInstances(entries)
//
//	if len(instances) > 0 {
//		// in case of debug, print all instances
//		for i := 0; i < len(instances); i++ {
//			fmt.Sprintf("instance is %s", instances[i])
//			//c.logger.Debugw("instances", "instance index", i, "value: ", instances[i])
//		}
//		// instances by index 0 is the most recent service
//	} else {
//		return errors.New("error in getInstances")
//	}
//
//	return nil
//}

// IMPORTANT NOTE
// If we have more than 1 tag, make sure, that we have at least 1 service registered in consul with this tag. For example
// tag: dev
//func filterEntries(entries []*api.ServiceEntry, tags ...string) []*api.ServiceEntry {
//	var es []*api.ServiceEntry
//	// range over the all ServiceEntries
//	for _, entry := range entries {
//		// make a temporary struct to hold all tags from ServiceEntries stored in consul
//		ts := make(map[string]struct{}, len(entry.Service.Tags))
//		// add tags to map
//		for _, tag := range entry.Service.Tags {
//			ts[tag] = struct{}{}
//		}
//
//		// range over the user defined tags for search
//		for _, tag := range tags {
//			// if we found this tag in previously added entry, we appending this entry to the slice, otherwise we
//			// ignore it
//			if _, ok := ts[tag]; ok {
//				es = append(es, entry)
//			}
//		}
//	}
//
//	return es
//}
//
//func makeInstances(entries []*api.ServiceEntry) []string {
//	instances := make([]string, len(entries))
//	for i, entry := range entries {
//		// we also need to check the status of the service
//		if status := entry.Checks.AggregatedStatus(); status != api.HealthPassing {
//			// skip this service as not passing the healthchecks
//			continue
//		}
//
//		addr := entry.Node.Address
//		if entry.Service.Address != "" {
//			addr = entry.Service.Address
//		}
//		// add only host ip address, because in case of GRPC we already know the port
//		instances[i] = fmt.Sprintf("%s", addr)
//	}
//	return instances
//}
