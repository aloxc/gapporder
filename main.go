package main

import (
	"context"
	"encoding/json"
	"github.com/aloxc/gapporder/config"
	"github.com/aloxc/gapporder/io"
	"github.com/aloxc/gapporder/module"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/log"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"os"
	"strconv"
	"time"
)

type Gapporder struct {
}

func init() {
	module.InsertTestOrder()
}
func (*Gapporder) getOrder(ctx context.Context, request *io.Request, response *io.Response) error {
	var order = module.Order{
		Id: request.Params["id"].(int32),
	}
	err := module.GetOrder(&order)
	if err != nil {
		log.Error(err)
		return err
	}
	response.Code = 0
	response.Message = "正常请求"
	response.Data = order
	return nil
}
func (this *Gapporder) Execute(ctx context.Context, request *io.Request, response *io.Response) error {
	bytes, _ := json.Marshal(request)
	log.Info("请求", string(bytes[0:]))
	switch request.Method {
	case "getOrder":
		return this.getOrder(ctx, request, response)
	}
	return nil
}

func main() {
	port := os.Getenv(config.SERVER_PORT)
	if port == "" {
		port = strconv.Itoa(config.SERVER_PORT_DEFAULT)
	}
	srv := server.NewServer(server.WithReadTimeout(time.Duration(2)*time.Second), server.WithWriteTimeout(time.Duration(2)*time.Second))
	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	srv.Plugins.Add(p)
	srv.RegisterName("gapporder", new(Gapporder), "")
	err := srv.Serve("tcp", ":"+port)
	if err != nil {
		log.Error("rpcx服务无法启动", err)
		os.Exit(1)
	}
	log.Info("服务已启动，监听端口[", port, "]")
}
