// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package logger

import (
	"encoding/json"

	"github.com/holiman/uint256"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/common/math"
	"github.com/scroll-tech/go-ethereum/core/vm"
)

var _ = (*wasmLogMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (w WasmLog) MarshalJSON() ([]byte, error) {
	type WasmLog struct {
		Pc            uint64                      `json:"pc"`
		OpFamily      OpCodeFamily                `json:"opcodeFamily,omitempty"`
		Op            vm.OpCodeInfo               `json:"op"`
		Params        []uint64                    `json:"params,omitempty"`
		Gas           math.HexOrDecimal64         `json:"gas"`
		GasCost       math.HexOrDecimal64         `json:"gasCost"`
		MemoryChanges map[uint32]hexutil.Bytes    `json:"memoryChanges,omitempty"`
		Stack         []uint256.Int               `json:"stack"`
		ReturnData    hexutil.Bytes               `json:"returnData,omitempty"`
		Storage       map[common.Hash]common.Hash `json:"-"`
		Depth         int                         `json:"depth"`
		RefundCounter uint64                      `json:"refund"`
		Err           error                       `json:"-"`
		Keep          uint32                      `json:"keep"`
		Drop          uint32                      `json:"drop"`
		OpName        string                      `json:"opName"`
		ErrorString   string                      `json:"error,omitempty"`
	}
	var enc WasmLog
	enc.Pc = w.Pc
	enc.OpFamily = w.OpFamily
	enc.Op = w.Op
	enc.Params = w.Params
	enc.Gas = math.HexOrDecimal64(w.Gas)
	enc.GasCost = math.HexOrDecimal64(w.GasCost)
	if w.MemoryChanges != nil {
		enc.MemoryChanges = make(map[uint32]hexutil.Bytes, len(w.MemoryChanges))
		for k, v := range w.MemoryChanges {
			enc.MemoryChanges[k] = hexutil.Bytes(v)
		}
	}
	enc.Stack = w.Stack
	enc.ReturnData = w.ReturnData
	enc.Storage = w.Storage
	enc.Depth = w.Depth
	enc.RefundCounter = w.RefundCounter
	enc.Err = w.Err
	enc.Keep = w.Keep
	enc.Drop = w.Drop
	enc.OpName = w.OpName()
	enc.ErrorString = w.ErrorString()
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (w *WasmLog) UnmarshalJSON(input []byte) error {
	type WasmLog struct {
		Pc            *uint64                     `json:"pc"`
		OpFamily      *OpCodeFamily               `json:"opcodeFamily,omitempty"`
		Op            vm.OpCodeInfo               `json:"op"`
		Params        []uint64                    `json:"params,omitempty"`
		Gas           *math.HexOrDecimal64        `json:"gas"`
		GasCost       *math.HexOrDecimal64        `json:"gasCost"`
		MemoryChanges map[uint32]hexutil.Bytes    `json:"memoryChanges,omitempty"`
		Stack         []uint256.Int               `json:"stack"`
		ReturnData    *hexutil.Bytes              `json:"returnData,omitempty"`
		Storage       map[common.Hash]common.Hash `json:"-"`
		Depth         *int                        `json:"depth"`
		RefundCounter *uint64                     `json:"refund"`
		Err           error                       `json:"-"`
		Keep          *uint32                     `json:"keep"`
		Drop          *uint32                     `json:"drop"`
	}
	var dec WasmLog
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Pc != nil {
		w.Pc = *dec.Pc
	}
	if dec.OpFamily != nil {
		w.OpFamily = *dec.OpFamily
	}
	if dec.Op != nil {
		w.Op = dec.Op
	}
	if dec.Params != nil {
		w.Params = dec.Params
	}
	if dec.Gas != nil {
		w.Gas = uint64(*dec.Gas)
	}
	if dec.GasCost != nil {
		w.GasCost = uint64(*dec.GasCost)
	}
	if dec.MemoryChanges != nil {
		w.MemoryChanges = make(map[uint32]string, len(dec.MemoryChanges))
		for k, v := range dec.MemoryChanges {
			w.MemoryChanges[k] = string(v)
		}
	}
	if dec.Stack != nil {
		w.Stack = dec.Stack
	}
	if dec.ReturnData != nil {
		w.ReturnData = *dec.ReturnData
	}
	if dec.Storage != nil {
		w.Storage = dec.Storage
	}
	if dec.Depth != nil {
		w.Depth = *dec.Depth
	}
	if dec.RefundCounter != nil {
		w.RefundCounter = *dec.RefundCounter
	}
	if dec.Err != nil {
		w.Err = dec.Err
	}
	if dec.Keep != nil {
		w.Keep = *dec.Keep
	}
	if dec.Drop != nil {
		w.Drop = *dec.Drop
	}
	return nil
}
