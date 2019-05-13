package porcelain

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-filecoin/protocol/storage/storagedeal"
)

type strgdlsPlumbing interface {
	DealsLs() ([]*storagedeal.Deal, error)
}

// DealGet returns a single deal matching a given cid or nil
func DealGet(plumbing strgdlsPlumbing, dealCid cid.Cid) *storagedeal.Deal {
	deals, err := plumbing.DealsLs()
	if err != nil {
		return nil
	}
	for _, storageDeal := range deals {
		if storageDeal.Response.ProposalCid == dealCid {
			return storageDeal
		}
	}
	return nil
}
