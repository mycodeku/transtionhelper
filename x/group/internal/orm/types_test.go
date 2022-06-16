package orm

import (
	"reflect"
	"testing"

	"github.com/mycodeku/transtionhelper/codec"
	"github.com/mycodeku/transtionhelper/codec/types"
	"github.com/mycodeku/transtionhelper/store/prefix"
	"github.com/mycodeku/transtionhelper/testutil/testdata"
	sdk "github.com/mycodeku/transtionhelper/types"
	sdkerrors "github.com/mycodeku/transtionhelper/types/errors"
	"github.com/mycodeku/transtionhelper/x/group/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeSafeRowGetter(t *testing.T) {
	storeKey := sdk.NewKVStoreKey("test")
	ctx := NewMockContext()
	var prefixKey = [2]byte{0x2}
	store := prefix.NewStore(ctx.KVStore(storeKey), prefixKey[:])
	md := testdata.TableModel{
		Id:   1,
		Name: "some name",
	}
	bz, err := md.Marshal()
	require.NoError(t, err)
	store.Set(EncodeSequence(1), bz)

	specs := map[string]struct {
		srcRowID     RowID
		srcModelType reflect.Type
		expObj       interface{}
		expErr       *sdkerrors.Error
	}{
		"happy path": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.TableModel{}),
			expObj:       md,
		},
		"unknown rowID should return sdkerrors.ErrNotFound": {
			srcRowID:     EncodeSequence(2),
			srcModelType: reflect.TypeOf(testdata.TableModel{}),
			expErr:       sdkerrors.ErrNotFound,
		},
		"wrong type should cause sdkerrors.ErrInvalidType": {
			srcRowID:     EncodeSequence(1),
			srcModelType: reflect.TypeOf(testdata.Cat{}),
			expErr:       sdkerrors.ErrInvalidType,
		},
		"empty rowID not allowed": {
			srcRowID:     []byte{},
			srcModelType: reflect.TypeOf(testdata.TableModel{}),
			expErr:       errors.ErrORMEmptyKey,
		},
		"nil rowID not allowed": {
			srcModelType: reflect.TypeOf(testdata.TableModel{}),
			expErr:       errors.ErrORMEmptyKey,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			interfaceRegistry := types.NewInterfaceRegistry()
			cdc := codec.NewProtoCodec(interfaceRegistry)

			getter := NewTypeSafeRowGetter(prefixKey, spec.srcModelType, cdc)
			var loadedObj testdata.TableModel

			err := getter(ctx.KVStore(storeKey), spec.srcRowID, &loadedObj)
			if spec.expErr != nil {
				require.True(t, spec.expErr.Is(err), err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.expObj, loadedObj)
		})
	}
}
