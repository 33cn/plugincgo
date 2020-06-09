package plugin

import (
	_ "github.com/33cn/plugincgo/plugin/consensus/init" //consensus init
	_ "github.com/33cn/plugincgo/plugin/crypto/init"    //crypto init
	_ "github.com/33cn/plugincgo/plugin/dapp/init"      //dapp init
	_ "github.com/33cn/plugincgo/plugin/mempool/init"   //mempool init
	_ "github.com/33cn/plugincgo/plugin/p2p/init"       //p2p init
	_ "github.com/33cn/plugincgo/plugin/store/init"     //store init
)
