package main

import (
	"github.com/aras/k11/apps/addons/core/executors"
	k11 "github.com/Kiisanz/k11-addon-sdk"
)

func main() {
	addon := k11.NewAddon("core", k11.Name("Core Addon"), k11.Version("1.0.0"))

	k11.Register(addon, "core.delay", k11.Node{Name: "Delay"}, &executors.DelayExecutor{})
	k11.Register(addon, "core.log", k11.Node{Name: "Log"}, &executors.LogExecutor{})
	k11.Register(addon, "core.record.build", k11.Node{Name: "Build Record"}, &executors.RecordBuilderExecutor{})
	k11.Register(addon, "core.dataset.append", k11.Node{Name: "Append Dataset"}, &executors.DatasetAppendExecutor{})
	k11.Register(addon, "core.dataset.append_many", k11.Node{Name: "Append Many to Dataset"}, &executors.DatasetAppendManyExecutor{})
	k11.Register(addon, "core.file.save", k11.Node{Name: "Save File"}, &executors.FileSaveExecutor{})

	addon.Serve()
}
