package vm

import (
	"fmt"
	"github.com/sbinet/wasm"
	zkwasm_gas_injector "github.com/wasm0/zkwasm-gas-injector"
)

const disableGasInjection = true

var WasmOpCodeToName = map[wasm.Opcode]string{
	wasm.Op_unreachable:         "unreachable",
	wasm.Op_nop:                 "nop",
	wasm.Op_block:               "block",
	wasm.Op_loop:                "loop",
	wasm.Op_if:                  "if",
	wasm.Op_else:                "else",
	wasm.Op_end:                 "end",
	wasm.Op_br:                  "br",
	wasm.Op_br_if:               "br_if",
	wasm.Op_br_table:            "br_table",
	wasm.Op_return:              "return",
	wasm.Op_call:                "call",
	wasm.Op_call_indirect:       "call_indirect",
	wasm.Op_drop:                "drop",
	wasm.Op_select:              "select",
	wasm.Op_get_local:           "get_local",
	wasm.Op_set_local:           "set_local",
	wasm.Op_tee_local:           "tee_local",
	wasm.Op_get_global:          "get_global",
	wasm.Op_set_global:          "set_global",
	wasm.Op_i32_load:            "i32_load",
	wasm.Op_i64_load:            "i64_load",
	wasm.Op_f32_load:            "f32_load",
	wasm.Op_f64_load:            "f64_load",
	wasm.Op_i32_load8_s:         "i32_load8_s",
	wasm.Op_i32_load8_u:         "i32_load8_u",
	wasm.Op_i32_load16_s:        "i32_load16_s",
	wasm.Op_i32_load16_u:        "i32_load16_u",
	wasm.Op_i64_load8_s:         "i64_load8_s",
	wasm.Op_i64_load8_u:         "i64_load8_u",
	wasm.Op_i64_load16_s:        "i64_load16_s",
	wasm.Op_i64_load16_u:        "i64_load16_u",
	wasm.Op_i64_load32_s:        "i64_load32_s",
	wasm.Op_i64_load32_u:        "i64_load32_u",
	wasm.Op_i32_store:           "i32_store",
	wasm.Op_i64_store:           "i64_store",
	wasm.Op_f32_store:           "f32_store",
	wasm.Op_f64_store:           "f64_store",
	wasm.Op_i32_store8:          "i32_store8",
	wasm.Op_i32_store16:         "i32_store16",
	wasm.Op_i64_store8:          "i64_store8",
	wasm.Op_i64_store16:         "i64_store16",
	wasm.Op_i64_store32:         "i64_store32",
	wasm.Op_current_memory:      "current_memory",
	wasm.Op_grow_memory:         "grow_memory",
	wasm.Op_i32_const:           "i32_const",
	wasm.Op_i64_const:           "i64_const",
	wasm.Op_f32_const:           "f32_const",
	wasm.Op_f64_const:           "f64_const",
	wasm.Op_i32_eqz:             "i32_eqz",
	wasm.Op_i32_eq:              "i32_eq",
	wasm.Op_i32_ne:              "i32_ne",
	wasm.Op_i32_lt_s:            "i32_lt_s",
	wasm.Op_i32_lt_u:            "i32_lt_u",
	wasm.Op_i32_gt_s:            "i32_gt_s",
	wasm.Op_i32_gt_u:            "i32_gt_u",
	wasm.Op_i32_le_s:            "i32_le_s",
	wasm.Op_i32_le_u:            "i32_le_u",
	wasm.Op_i32_ge_s:            "i32_ge_s",
	wasm.Op_i32_ge_u:            "i32_ge_u",
	wasm.Op_i64_eqz:             "i64_eqz",
	wasm.Op_i64_eq:              "i64_eq",
	wasm.Op_i64_ne:              "i64_ne",
	wasm.Op_i64_lt_s:            "i64_lt_s",
	wasm.Op_i64_lt_u:            "i64_lt_u",
	wasm.Op_i64_gt_s:            "i64_gt_s",
	wasm.Op_i64_gt_u:            "i64_gt_u",
	wasm.Op_i64_le_s:            "i64_le_s",
	wasm.Op_i64_le_u:            "i64_le_u",
	wasm.Op_i64_ge_s:            "i64_ge_s",
	wasm.Op_i64_ge_u:            "i64_ge_u",
	wasm.Op_f32_eq:              "f32_eq",
	wasm.Op_f32_ne:              "f32_ne",
	wasm.Op_f32_lt:              "f32_lt",
	wasm.Op_f32_gt:              "f32_gt",
	wasm.Op_f32_le:              "f32_le",
	wasm.Op_f32_ge:              "f32_ge",
	wasm.Op_f64_eq:              "f64_eq",
	wasm.Op_f64_ne:              "f64_ne",
	wasm.Op_f64_lt:              "f64_lt",
	wasm.Op_f64_gt:              "f64_gt",
	wasm.Op_f64_le:              "f64_le",
	wasm.Op_f64_ge:              "f64_ge",
	wasm.Op_i32_clz:             "i32_clz",
	wasm.Op_i32_ctz:             "i32_ctz",
	wasm.Op_i32_popcnt:          "i32_popcnt",
	wasm.Op_i32_add:             "i32_add",
	wasm.Op_i32_sub:             "i32_sub",
	wasm.Op_i32_mul:             "i32_mul",
	wasm.Op_i32_div_s:           "i32_div_s",
	wasm.Op_i32_div_u:           "i32_div_u",
	wasm.Op_i32_rem_s:           "i32_rem_s",
	wasm.Op_i32_rem_u:           "i32_rem_u",
	wasm.Op_i32_and:             "i32_and",
	wasm.Op_i32_or:              "i32_or",
	wasm.Op_i32_xor:             "i32_xor",
	wasm.Op_i32_shl:             "i32_shl",
	wasm.Op_i32_shr_s:           "i32_shr_s",
	wasm.Op_i32_shr_u:           "i32_shr_u",
	wasm.Op_i32_rotl:            "i32_rotl",
	wasm.Op_i32_rotr:            "i32_rotr",
	wasm.Op_i64_clz:             "i64_clz",
	wasm.Op_i64_ctz:             "i64_ctz",
	wasm.Op_i64_popcnt:          "i64_popcnt",
	wasm.Op_i64_add:             "i64_add",
	wasm.Op_i64_sub:             "i64_sub",
	wasm.Op_i64_mul:             "i64_mul",
	wasm.Op_i64_div_s:           "i64_div_s",
	wasm.Op_i64_div_u:           "i64_div_u",
	wasm.Op_i64_rem_s:           "i64_rem_s",
	wasm.Op_i64_rem_u:           "i64_rem_u",
	wasm.Op_i64_and:             "i64_and",
	wasm.Op_i64_or:              "i64_or",
	wasm.Op_i64_xor:             "i64_xor",
	wasm.Op_i64_shl:             "i64_shl",
	wasm.Op_i64_shr_s:           "i64_shr_s",
	wasm.Op_i64_shr_u:           "i64_shr_u",
	wasm.Op_i64_rotl:            "i64_rotl",
	wasm.Op_i64_rotr:            "i64_rotr",
	wasm.Op_f32_abs:             "f32_abs",
	wasm.Op_f32_neg:             "f32_neg",
	wasm.Op_f32_ceil:            "f32_ceil",
	wasm.Op_f32_floor:           "f32_floor",
	wasm.Op_f32_trunc:           "f32_trunc",
	wasm.Op_f32_nearest:         "f32_nearest",
	wasm.Op_f32_sqrt:            "f32_sqrt",
	wasm.Op_f32_add:             "f32_add",
	wasm.Op_f32_sub:             "f32_sub",
	wasm.Op_f32_mul:             "f32_mul",
	wasm.Op_f32_div:             "f32_div",
	wasm.Op_f32_min:             "f32_min",
	wasm.Op_f32_max:             "f32_max",
	wasm.Op_f32_copysign:        "f32_copysign",
	wasm.Op_f64_abs:             "f64_abs",
	wasm.Op_f64_neg:             "f64_neg",
	wasm.Op_f64_ceil:            "f64_ceil",
	wasm.Op_f64_floor:           "f64_floor",
	wasm.Op_f64_trunc:           "f64_trunc",
	wasm.Op_f64_nearest:         "f64_nearest",
	wasm.Op_f64_sqrt:            "f64_sqrt",
	wasm.Op_f64_add:             "f64_add",
	wasm.Op_f64_sub:             "f64_sub",
	wasm.Op_f64_mul:             "f64_mul",
	wasm.Op_f64_div:             "f64_div",
	wasm.Op_f64_min:             "f64_min",
	wasm.Op_f64_max:             "f64_max",
	wasm.Op_f64_copysign:        "f64_copysign",
	wasm.Op_i32_wrap_i64:        "i32_wrap_i64",
	wasm.Op_i32_trunc_s_f32:     "i32_trunc_s_f32",
	wasm.Op_i32_trunc_u_f32:     "i32_trunc_u_f32",
	wasm.Op_i32_trunc_s_f64:     "i32_trunc_s_f64",
	wasm.Op_i32_trunc_u_f64:     "i32_trunc_u_f64",
	wasm.Op_i64_extend_s_i32:    "i64_extend_s_i32",
	wasm.Op_i64_extend_u_i32:    "i64_extend_u_i32",
	wasm.Op_i64_trunc_s_f32:     "i64_trunc_s_f32",
	wasm.Op_i64_trunc_u_f32:     "i64_trunc_u_f32",
	wasm.Op_i64_trunc_s_f64:     "i64_trunc_s_f64",
	wasm.Op_i64_trunc_u_f64:     "i64_trunc_u_f64",
	wasm.Op_f32_convert_s_i32:   "f32_convert_s_i32",
	wasm.Op_f32_convert_u_i32:   "f32_convert_u_i32",
	wasm.Op_f32_convert_s_i64:   "f32_convert_s_i64",
	wasm.Op_f32_convert_u_i64:   "f32_convert_u_i64",
	wasm.Op_f32_demote_f64:      "f32_demote_f64",
	wasm.Op_f64_convert_s_i32:   "f64_convert_s_i32",
	wasm.Op_f64_convert_u_i32:   "f64_convert_u_i32",
	wasm.Op_f64_convert_s_i64:   "f64_convert_s_i64",
	wasm.Op_f64_convert_u_i64:   "f64_convert_u_i64",
	wasm.Op_f64_promote_f32:     "f64_promote_f32",
	wasm.Op_i32_reinterpret_f32: "i32_reinterpret_f32",
	wasm.Op_i64_reinterpret_f64: "i64_reinterpret_f64",
	wasm.Op_f32_reinterpret_i32: "f32_reinterpret_i32",
	wasm.Op_f64_reinterpret_i64: "f64_reinterpret_i64",
}

const (
	GasImportedFunction = "gas"
)

func injectGasComputationAndStackProtection(
	destCode []byte,
) (ret []byte, err error) {
	if disableGasInjection {
		return destCode, nil
	}
	if destCode == nil {
		return nil, fmt.Errorf("no contract code to check")
	}
	injectRes, err := zkwasm_gas_injector.Inject(
		destCode,
		zkwasm_gas_injector.InjectTypeBoth,
		zkwasm_gas_injector.InjectGasTypeHost,
		1,
		10000,
		0,
		1024,
		zkwasm_gas_injector.ReturnFormatWasm,
	)
	if err != nil {
		return nil, err
	}
	if injectRes != nil {
		ret = injectRes
	}
	return
}