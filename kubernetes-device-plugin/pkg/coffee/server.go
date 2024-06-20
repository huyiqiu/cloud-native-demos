package coffee

import (
	"coffee-device-plugin/pkg/utils"
	"context"
	"net"
	"os"
	"strings"
	"time"
	"path"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

var (
	coffeePath = "/home/hyq/coffee"
	resourceName = "hyq.com/coffee"
	coffeeSocket = "coffee.sock"
	kubeletSocket = "kubelet.sock"
	devicePluginPath = "/var/lib/kubelet/device-plugins/"
)

type CoffeeDevicePlugin struct {
	devices []*pluginapi.Device
	server *grpc.Server
}

func NewCoffeePulgin() *CoffeeDevicePlugin {
	devices := listDevice()
	cdp := &CoffeeDevicePlugin{
		devices: devices,
	}
	return cdp
}

func RegisterWithKubelet() error {
	conn, err := grpc.Dial("unix://"+devicePluginPath+kubeletSocket, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	client := pluginapi.NewRegistrationClient(conn)
	_, err = client.Register(context.Background(), &pluginapi.RegisterRequest{
		Version: pluginapi.Version,
		Endpoint: path.Base(devicePluginPath + coffeePath),
		ResourceName: resourceName,
	})
	return err
}

func (dp *CoffeeDevicePlugin) ListAndWatch(e *pluginapi.Empty, srv pluginapi.DevicePlugin_ListAndWatchServer) error {
	if err := srv.Send(&pluginapi.ListAndWatchResponse{Devices: dp.devices}); err != nil {
        return err
    }
	for {
		time.Sleep(time.Second * 10)
		devices := listDevice()
		if utils.Diff(devices, dp.devices) {
			if err := srv.Send(&pluginapi.ListAndWatchResponse{Devices: dp.devices}); err != nil {
				return err
			}
		}
	}
}

func (dp *CoffeeDevicePlugin) Allocate(ctx context.Context, reqs *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
	allocateResp := &pluginapi.AllocateResponse{}
	for _, req := range reqs.ContainerRequests {
		resp := &pluginapi.ContainerAllocateResponse{
			Devices: []*pluginapi.DeviceSpec{
				{
					ContainerPath: coffeePath,
					HostPath: coffeePath,
					Permissions: "rw",
				},
			},
			Envs: map[string]string{
				"COFFEE_DEVICES": strings.Join(req.DevicesIDs, ","),
			},
		}
		allocateResp.ContainerResponses = append(allocateResp.ContainerResponses, resp)
	}
	return allocateResp, nil
}

func (dp *CoffeeDevicePlugin) Serve() error {
	listener, err := net.Listen("unix", devicePluginPath+coffeeSocket)
	if err != nil {
		return err
	}
	dp.server = grpc.NewServer(grpc.EmptyServerOption{})
	pluginapi.RegisterDevicePluginServer(dp.server, dp)
	return dp.server.Serve(listener)
}

func (dp *CoffeeDevicePlugin) GetDevicePluginOptions(ctx context.Context, e *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
	return &pluginapi.DevicePluginOptions{}, nil
}

func (dp *CoffeeDevicePlugin) PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
	return &pluginapi.PreStartContainerResponse{}, nil
}

func (dp *CoffeeDevicePlugin) GetPreferredAllocation(context.Context, *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
	return &pluginapi.PreferredAllocationResponse{}, nil
}


func listDevice() []*pluginapi.Device {
	fs, err := os.ReadDir(coffeePath)
	if err != nil {
		log.Fatal("fail to get device info of coffee...")
	}
	devices := make([]*pluginapi.Device, 0)
	for i := range fs {
		devices = append(devices, &pluginapi.Device{
			ID: fs[i].Name(),
			Health: pluginapi.Healthy,
		})
	}
	return devices
}