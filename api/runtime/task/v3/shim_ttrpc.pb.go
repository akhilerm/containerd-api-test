// Code generated by protoc-gen-go-ttrpc. DO NOT EDIT.
// source: github.com/containerd/containerd/api/runtime/task/v3/shim.proto
package task

import (
	context "context"
	ttrpc "github.com/containerd/ttrpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type TTRPCTaskService interface {
	State(context.Context, *StateRequest) (*StateResponse, error)
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Pids(context.Context, *PidsRequest) (*PidsResponse, error)
	Pause(context.Context, *PauseRequest) (*emptypb.Empty, error)
	Resume(context.Context, *ResumeRequest) (*emptypb.Empty, error)
	Checkpoint(context.Context, *CheckpointTaskRequest) (*emptypb.Empty, error)
	Kill(context.Context, *KillRequest) (*emptypb.Empty, error)
	Exec(context.Context, *ExecProcessRequest) (*emptypb.Empty, error)
	ResizePty(context.Context, *ResizePtyRequest) (*emptypb.Empty, error)
	CloseIO(context.Context, *CloseIORequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateTaskRequest) (*emptypb.Empty, error)
	Wait(context.Context, *WaitRequest) (*WaitResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Shutdown(context.Context, *ShutdownRequest) (*emptypb.Empty, error)
}

func RegisterTTRPCTaskService(srv *ttrpc.Server, svc TTRPCTaskService) {
	srv.RegisterService("containerd.task.v3.Task", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"State": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.State(ctx, &req)
			},
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req DeleteRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Delete(ctx, &req)
			},
			"Pids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pids(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req PauseRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResumeRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Resume(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Checkpoint(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.CloseIO(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Update(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Wait(ctx, &req)
			},
			"Stats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req StatsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Stats(ctx, &req)
			},
			"Connect": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ConnectRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Connect(ctx, &req)
			},
			"Shutdown": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req ShutdownRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}
				return svc.Shutdown(ctx, &req)
			},
		},
	})
}

type ttrpctaskClient struct {
	client *ttrpc.Client
}

func NewTTRPCTaskClient(client *ttrpc.Client) TTRPCTaskService {
	return &ttrpctaskClient{
		client: client,
	}
}

func (c *ttrpctaskClient) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	var resp StateResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "State", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Create(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	var resp CreateTaskResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Create", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	var resp StartResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Start", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	var resp DeleteResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Delete", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Pids(ctx context.Context, req *PidsRequest) (*PidsResponse, error) {
	var resp PidsResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Pids", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Pause(ctx context.Context, req *PauseRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Pause", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Resume(ctx context.Context, req *ResumeRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Resume", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Checkpoint(ctx context.Context, req *CheckpointTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Checkpoint", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Kill(ctx context.Context, req *KillRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Kill", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Exec(ctx context.Context, req *ExecProcessRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Exec", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) ResizePty(ctx context.Context, req *ResizePtyRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "ResizePty", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) CloseIO(ctx context.Context, req *CloseIORequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "CloseIO", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Update(ctx context.Context, req *UpdateTaskRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Update", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Wait(ctx context.Context, req *WaitRequest) (*WaitResponse, error) {
	var resp WaitResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Wait", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Stats(ctx context.Context, req *StatsRequest) (*StatsResponse, error) {
	var resp StatsResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Stats", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Connect(ctx context.Context, req *ConnectRequest) (*ConnectResponse, error) {
	var resp ConnectResponse
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Connect", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *ttrpctaskClient) Shutdown(ctx context.Context, req *ShutdownRequest) (*emptypb.Empty, error) {
	var resp emptypb.Empty
	if err := c.client.Call(ctx, "containerd.task.v3.Task", "Shutdown", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
