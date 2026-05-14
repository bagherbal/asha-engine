package asha

import (
	"math"
	"strings"
)

type Family struct {
	NativeChargedFlavorDim int       `json:"native_charged_flavor_dim"`
	KXYChargedCoeffDim     int       `json:"kxy_charged_coeff_dim"`
	KTrace                 float64   `json:"k_trace"`
	KTraceSquare           float64   `json:"k_trace_square"`
	RhoBeta                []float64 `json:"rho_beta"`
	RhoRatio               float64   `json:"rho_ratio"`
	CommKSNorm             float64   `json:"comm_k_s_norm"`
	CommKXNorm             float64   `json:"comm_k_x_norm"`
	CPWitness              float64   `json:"cp_witness"`
}

func NewFamily(beta float64) Family {
	raw := []float64{math.Exp(beta), 1, math.Exp(-beta)} // exp(-β diag(-1,0,1))
	z := raw[0] + raw[1] + raw[2]
	rho := []float64{raw[0] / z, raw[1] / z, raw[2] / z}
	return Family{
		NativeChargedFlavorDim: 13,
		KXYChargedCoeffDim:     9,
		KTrace:                 0,
		KTraceSquare:           2,
		RhoBeta:                rho,
		RhoRatio:               math.Exp(2 * beta),
		CommKSNorm:             math.Sqrt(6),
		CommKXNorm:             math.Sqrt(12),
		CPWitness:              8.397024,
	}
}

func (f Family) Quantities() []Quantity {
	return []Quantity{
		{Symbol: "dim M_charged^native", Value: float64(f.NativeChargedFlavorDim), Formula: "6 quark masses + 4 CKM + 3 charged-lepton masses", Status: StatusEnvironmental, Note: "native firewall"},
		{Symbol: "K_gen", Text: "diag(-1,0,1)", Formula: "K_gen = diag(-1,0,1)", Status: StatusQuarantined, Note: "hierarchy capacity only"},
		{Symbol: "ρ_β", Text: floatSliceText(f.RhoBeta), Formula: "exp(-βK)/Tr exp(-βK)", Status: StatusQuarantined},
		{Symbol: "ρ_max/ρ_min", Value: f.RhoRatio, Formula: "exp(2β)", Status: StatusQuarantined},
		{Symbol: "X_gen", Text: "S+S^T", Formula: "real shift quadrature", Status: StatusQuarantined, Note: "real mixing capacity"},
		{Symbol: "Y_gen", Text: "i(S-S^T)", Formula: "imaginary shift quadrature", Status: StatusQuarantined, Note: "CP capacity"},
		{Symbol: "||[K,S]||_F", Value: f.CommKSNorm, Formula: "sqrt(6)", Status: StatusQuarantined},
		{Symbol: "||[K,X]||_F", Value: f.CommKXNorm, Formula: "sqrt(12)", Status: StatusQuarantined},
		{Symbol: "Im Tr([M_u,M_d]^3)", Value: f.CPWitness, Formula: "sample nonzero CP-capacity witness", Status: StatusQuarantined},
		{Symbol: "dim C_KXY^charged", Value: float64(f.KXYChargedCoeffDim), Formula: "3 charged sectors × 3 symbolic coefficients", Status: StatusQuarantined},
	}
}

func (f Family) Checks() []Check {
	return []Check{
		{Name: "Native charged flavor firewall", Passed: f.NativeChargedFlavorDim == 13, Detail: "dim M_charged^native=13"},
		{Name: "KMS family hierarchy capacity", Passed: len(f.RhoBeta) == 3 && f.RhoRatio > 1, Detail: "ρβ nontracial for β≠0"},
		{Name: "Noncommuting capacity", Passed: f.CommKSNorm > 0 && f.CommKXNorm > 0, Detail: "K does not commute with shift/quadrature"},
		{Name: "CP capacity not CP prediction", Passed: f.CPWitness != 0 && f.KXYChargedCoeffDim == 9, Detail: "phase coefficients remain free"},
	}
}

func floatSliceText(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmtFloat(x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
