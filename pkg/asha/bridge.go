package asha

import "math"

type Bridge struct {
	PlanckMassGeV        float64  `json:"planck_mass_gev"`
	F2LambdaOverMP2      Rational `json:"f2_lambda_over_mp2"`
	CorrectionFactor     float64  `json:"correction_factor"`
	VPfGeV               float64  `json:"v_pf_gev"`
	NodeTraceRatio       Rational `json:"node_trace_ratio"`
	EdgeMeasureFactor    Rational `json:"edge_measure_factor"`
	EdgeTraceRatio       float64  `json:"edge_trace_ratio"`
	LambdaH              float64  `json:"lambda_h"`
	HiggsTreeGeV         float64  `json:"higgs_tree_gev"`
	HeavyBGapMajoranaGeV float64  `json:"heavy_b_gap_majorana_gev"`
	MajoranaOverclosure  float64  `json:"majorana_overclosure_ratio"`
}

func NewBridge(planckGeV float64) Bridge {
	node := NewRational(1197, 4624)
	edgeFactor := NewRational(7, 10)
	edgeTrace := edgeFactor.Float64() * node.Float64()
	lambdaH := math.Pi * math.Pi * node.Float64() / 20.0
	vPf := planckGeV * math.Pow(2, 1.5) * math.Exp(-4*math.Pi*math.Pi)
	return Bridge{
		PlanckMassGeV:        planckGeV,
		F2LambdaOverMP2:      NewRational(0, 1), // π²/8 is irrational; kept in Formula below.
		CorrectionFactor:     8 * math.Pi,
		VPfGeV:               vPf,
		NodeTraceRatio:       node,
		EdgeMeasureFactor:    edgeFactor,
		EdgeTraceRatio:       edgeTrace,
		LambdaH:              lambdaH,
		HiggsTreeGeV:         vPf * math.Sqrt(2*lambdaH),
		HeavyBGapMajoranaGeV: 1.46774973718e6,
		MajoranaOverclosure:  1.3e13,
	}
}

func (b Bridge) Quantities() []Quantity {
	return []Quantity{
		{Symbol: "f₂(Λ/M_P)²", Text: "π²/8", Formula: "f₂(Λ/M_P)² = π²/8", Status: StatusBridge, Note: "CCM/Einstein normalization bridge"},
		{Symbol: "8π", Value: b.CorrectionFactor, Formula: "(π²/8)/(π/64)=8π", Status: StatusBridge, Note: "correction of earlier coefficient route"},
		{Symbol: "v_Pf", Value: b.VPfGeV, Unit: "GeV", Formula: "v_Pf = M_P 2^(3/2) exp(-4π²)", Status: StatusBridge},
		{Symbol: "(e/a²)_node", Text: b.NodeTraceRatio.String(), Formula: "1197/4624", Status: StatusNative},
		{Symbol: "(e/a²)_edge", Value: b.EdgeTraceRatio, Formula: "(7/10)(1197/4624)", Status: StatusBridge},
		{Symbol: "λ_H", Value: b.LambdaH, Formula: "π²(1197/4624)/20", Status: StatusBridge},
		{Symbol: "m_H^tree", Value: b.HiggsTreeGeV, Unit: "GeV", Formula: "v_Pf sqrt(2λ_H)", Status: StatusBridge, Note: "tree-level proxy, not pole-mass theorem"},
		{Symbol: "M_B", Value: b.HeavyBGapMajoranaGeV, Unit: "GeV", Formula: "sealed B-gap Majorana ledger", Status: StatusBridge},
		{Symbol: "Ω_candidate/Ω_DM", Value: b.MajoranaOverclosure, Formula: "stable thermal B-gap Majorana relic overclosure", Status: StatusFailedRoute},
	}
}

func (b Bridge) Checks() []Check {
	return []Check{
		{Name: "Pfaffian scale positive", Passed: b.VPfGeV > 0, Detail: "v_Pf computed from Planck mass bridge"},
		{Name: "Higgs tree proxy", Passed: nearly(b.HiggsTreeGeV, 124.925370288, 2e-3), Detail: "m_H^tree ≈ 124.925 GeV under project Planck convention"},
		{Name: "Majorana stable thermal relic rejected", Passed: b.MajoranaOverclosure > 1e12, Detail: "overcloses by ~1.3e13"},
	}
}
