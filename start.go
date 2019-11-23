package main

import (
	"errors"
	"fmt"

	"github.com/opencontainers/runc/libcontainer"
	"github.com/urfave/cli"
)

var startCommand = cli.Command{
	Name:  "start",
	Usage: "executes the user defined process in a created container",
	ArgsUsage: `<container-id>

Where "<container-id>" is your name for the instance of the container that you
are starting. The name you provide for the container instance must be unique on
your host.`,
	Description: `The start command executes the user defined process in a created container.`,
	Action: func(context *cli.Context) error {
		// 依旧检查参数
		if err := checkArgs(context, 1, exactArgs); err != nil {
			return err
		}

		// 通过"id"获取容器
		container, err := getContainer(context)
		if err != nil {
			return err
		}

		// 查看容器status
		status, err := container.Status()
		if err != nil {
			return err
		}

		// 根据容器状态决定start的行为
		switch status {
		case libcontainer.Created:
			return container.Exec()
		case libcontainer.Stopped:
			return errors.New("cannot start a container that has stopped")
		case libcontainer.Running:
			return errors.New("cannot start an already running container")
		default:
			return fmt.Errorf("cannot start a container in the %s state\n", status)
		}
	},
}
