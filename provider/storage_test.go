package provider

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/go-cmp/cmp"
)

func TestProvider_StorageAt(t *testing.T) {
	// TODO: Find a better test case
	want := "0x0"

	for _, tc := range []struct {
		address common.Hash
		key     uint64
		opts    *StorageAtOptions
	}{{
		address: common.HexToHash("0x06eeefce63bc81620e375c7501cb7b5aecdf9fb99aa5ec25886b8b854c4293cb"),
		key:     1,
		opts:    nil,
	}, {
		address: common.HexToHash("0x06eeefce63bc81620e375c7501cb7b5aecdf9fb99aa5ec25886b8b854c4293cb"),
		key:     1,
		opts:    &StorageAtOptions{BlockNumber: 582},
	}, {
		address: common.HexToHash("0x06eeefce63bc81620e375c7501cb7b5aecdf9fb99aa5ec25886b8b854c4293cb"),
		key:     1,
		opts:    &StorageAtOptions{BlockHash: "0x182d83f0ed972e97fa529be0088e20b5a7cb32e3bba0d164d668147713e79f9"},
	}} {
		ctx := context.Background()
		p := NewProvider("https://alpha-mainnet.starknet.io")
		got, err := p.StorageAt(ctx, tc.address, tc.key, tc.opts)
		if err != nil {
			t.Fatalf("getting storage at: %v", err)
		}

		if tc.opts != nil {
			if diff := cmp.Diff(want, got, nil); diff != "" {
				t.Errorf("Storage value diff mismatch (-want +got):\n%s", diff)
			}
		}
	}
}
