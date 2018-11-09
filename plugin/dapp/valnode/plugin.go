package valnode

import (
	"gitlab.33.cn/chain33/plugin/plugin/dapp/valnode/executor"
	"gitlab.33.cn/chain33/plugin/plugin/dapp/valnode/types"
	"gitlab.33.cn/chain33/chain33/pluginmgr"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ValNodeX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      nil,
		RPC:      nil,
	})
}
