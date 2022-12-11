package proj3

import (
	"testing"
	"github.com/libp2p/go-libp2p/core/peer"
)

func FuzzTest(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string){
		l := localNotifee("meep meep")
		l.HandlePeerFound(peer.AddrInfo{})
	})

}

