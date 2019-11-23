// +build linux

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var pauseCommand = cli.Command{
	Name:  "pause",
	Usage: "pause suspends all processes inside the container",
	ArgsUsage: `<container-id>

Where "<container-id>" is the name for the instance of the container to be
paused. `,
	Description: `The pause command suspends all processes in the instance of the container.

Use runc list to identify instances of containers and their current status.`,
	Action: func(context *cli.Context) error {
		// 检查参数
		if err := checkArgs(context, 1, exactArgs); err != nil {
			return err
		}
		// 判断是否是rootless的cgroup, rootless无法进行pause
		rootlessCg, err := shouldUseRootlessCgroupManager(context)
		if err != nil {
			return err
		}
		if rootlessCg {
			logrus.Warnf("runc pause may fail if you don't have the full access to cgroups")
		}
		// 得到Container
		container, err := getContainer(context)
		if err != nil {
			return err
		}
		// 调用Container.Pause停止
		return container.Pause()
	},
}

var resumeCommand = cli.Command{
	Name:  "resume",
	Usage: "resumes all processes that have been previously paused",
	ArgsUsage: `<container-id>

Where "<container-id>" is the name for the instance of the container to be
resumed.`,
	Description: `The resume command resumes all processes in the instance of the container.

Use runc list to identify instances of containers and their current status.`,
	Action: func(context *cli.Context) error {
		// 检查参数
		if err := checkArgs(context, 1, exactArgs); err != nil {
			return err
		}
		// 判断是否是rootless的cgroup, rootless无法进行resume
		rootlessCg, err := shouldUseRootlessCgroupManager(context)
		if err != nil {
			return err
		}
		if rootlessCg {
			logrus.Warn("runc resume may fail if you don't have the full access to cgroups")
		}
		// 得到Container，并调用Resume
		container, err := getContainer(context)
		if err != nil {
			return err
		}
		return container.Resume()
	},
}
