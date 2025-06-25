package hasher

import "github.com/ethereum/go-ethereum/common"

func AddrHash(key string) uint32 {
	addr := common.HexToAddress(key)
	return uint32(addr[0])<<24 | uint32(addr[1])<<16 | uint32(addr[2])<<8 | uint32(addr[3])
}
