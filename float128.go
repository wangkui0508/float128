package float128

/*
#cgo LDFLAGS: -lsoftfloat
#define SOFTFLOAT_FAST_INT64
#include "softfloat.h"
int_fast64_t f128_to_i64_default(float128_t a) {
	return f128_to_i64(a,softfloat_round_near_even,false);
}
typedef union {
    double a;
    float64_t b;
} unionType;
float128_t double_to_f128(double a) {
	unionType v;
	v.a=a;
	return f64_to_f128(v.b);
}
double f128_to_double(float128_t x) {
	unionType v;
	v.b=f128_to_f64(x);
	return v.a;
}
*/
import "C"

type F128 struct {
	v [2]uint64
}

func f128_c_to_go(a C.float128_t) F128 {
	var res F128
	res.v[0] = uint64(a.v[0])
	res.v[1] = uint64(a.v[1])
	return res
}

func f128_go_to_c(a F128)  C.float128_t{
	var res C.float128_t
	res.v[0] = C.uint64_t(a.v[0])
	res.v[1] = C.uint64_t(a.v[1])
	return res
}

func F128FromI64(a int64) F128 {
	return f128_c_to_go(C.i64_to_f128(C.int64_t(a)))
}
func F128FromF64(a float64) F128 {
	return f128_c_to_go(C.double_to_f128(C.double(a)))
}
func (a F128) ToI64() int64 {
	return int64(C.f128_to_i64_default(f128_go_to_c(a)))
}
func (a F128) ToF64() float64 {
	return float64(C.f128_to_double(f128_go_to_c(a)))
}

func (a F128) Add(b F128) F128 {
	res := C.f128_add(f128_go_to_c(a),f128_go_to_c(b))
	return f128_c_to_go(res)
}
func (a F128) Sub(b F128) F128 {
	res := C.f128_sub(f128_go_to_c(a),f128_go_to_c(b))
	return f128_c_to_go(res)
}
func (a F128) Mul(b F128) F128 {
	res := C.f128_mul(f128_go_to_c(a),f128_go_to_c(b))
	return f128_c_to_go(res)
}
func (a F128) MulAdd(b F128, c F128) F128 {
	res := C.f128_mulAdd(f128_go_to_c(a),f128_go_to_c(b),f128_go_to_c(c))
	return f128_c_to_go(res)
}
func (a F128) Div(b F128) F128 {
	res := C.f128_div(f128_go_to_c(a),f128_go_to_c(b))
	return f128_c_to_go(res)
}
func (a F128) Rem(b F128) F128 {
	res := C.f128_rem(f128_go_to_c(a),f128_go_to_c(b))
	return f128_c_to_go(res)
}
func (a F128) Sqrt() F128 {
	res := C.f128_sqrt(f128_go_to_c(a))
	return f128_c_to_go(res)
}

func (a F128) Equal(b F128) bool {
	return bool(C.f128_eq(f128_go_to_c(a),f128_go_to_c(b)))
}
func (a F128) LTE(b F128) bool {
	return bool(C.f128_le(f128_go_to_c(a),f128_go_to_c(b)))
}
func (a F128) LT(b F128) bool {
	return bool(C.f128_lt(f128_go_to_c(a),f128_go_to_c(b)))
}
func (a F128) GTE(b F128) bool {
	return bool(C.f128_le(f128_go_to_c(b),f128_go_to_c(a)))
}
func (a F128) GT(b F128) bool {
	return bool(C.f128_lt(f128_go_to_c(b),f128_go_to_c(a)))
}
func (a F128) IsSignalingNaN() bool {
	return bool(C.f128_isSignalingNaN(f128_go_to_c(a)))
}
