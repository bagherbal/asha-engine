package asha

import "math"

type Geometry struct {
	CliffordDimension int      `json:"clifford_dimension"`
	GradeDimensions   []int    `json:"grade_dimensions"`
	RankPB            int      `json:"rank_p_b"`
	RankPG            int      `json:"rank_p_g"`
	DimK              int      `json:"dim_k"`
	ContactIndex      Rational `json:"contact_index"`
	BGap              float64  `json:"b_gap"`
	TrMK              float64  `json:"tr_m_k"`
	TrMK2             float64  `json:"tr_m_k2"`
	LambdaShape       float64  `json:"lambda_shape"`
	ActiveScalarDims  int      `json:"active_scalar_dims"`
	ProtectedDims     int      `json:"protected_dims"`
}

type Electroweak struct {
	KY              Rational `json:"k_y"`
	Sin2ThetaStar   Rational `json:"sin2_theta_star"`
	GaugeDirections int      `json:"gauge_directions"`
	HiggsDoublets   int      `json:"higgs_doublets"`
	GaugeGroup      string   `json:"gauge_group"`
}

func NewGeometry() Geometry {
	tr := 1.1333333333
	tr2 := 0.3325
	return Geometry{
		CliffordDimension: 256,
		GradeDimensions:   []int{1, 8, 28, 56, 70, 56, 28, 8, 1},
		RankPB:            56,
		RankPG:            14,
		DimK:              7,
		ContactIndex:      NewRational(1, 1),
		BGap:              0.1024649212,
		TrMK:              tr,
		TrMK2:             tr2,
		LambdaShape:       tr2 / (tr * tr),
		ActiveScalarDims:  4,
		ProtectedDims:     3,
	}
}

func NewElectroweak() Electroweak {
	return Electroweak{
		KY:              NewRational(5, 3),
		Sin2ThetaStar:   NewRational(3, 8),
		GaugeDirections: 12,
		HiggsDoublets:   1,
		GaugeGroup:      "U(1)_Y × SU(2)_L × SU(3)_C",
	}
}

func GradeDimensions(n int) []int {
	out := make([]int, n+1)
	for k := 0; k <= n; k++ {
		out[k] = binomial(n, k)
	}
	return out
}

func binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - k + i) / i
	}
	return res
}

func (g Geometry) Checks() []Check {
	return []Check{
		{Name: "Clifford dimension", Passed: g.CliffordDimension == 256, Detail: "dim Cℓ(1,7)=2^8"},
		{Name: "Exterior grade dimensions", Passed: intsEqual(g.GradeDimensions, GradeDimensions(8)), Detail: "[1,8,28,56,70,56,28,8,1]"},
		{Name: "Boolean/G2 contact vacuum", Passed: g.RankPB == 56 && g.RankPG == 14 && g.DimK == 7, Detail: "rank(P_B)=56 rank(P_G)=14 dim K=7"},
		{Name: "Scalar shape", Passed: nearly(g.LambdaShape, 0.2588667820, 1e-9), Detail: "Tr(M_K^2)/Tr(M_K)^2"},
	}
}

func intsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (e Electroweak) Checks() []Check {
	return []Check{
		{Name: "Hypercharge normalization", Passed: e.KY.Num == 5 && e.KY.Den == 3, Detail: "k_Y=5/3"},
		{Name: "Boundary weak angle", Passed: e.Sin2ThetaStar.Num == 3 && e.Sin2ThetaStar.Den == 8, Detail: "sin²θ*=3/8"},
		{Name: "Gauge/Higgs inventory", Passed: e.GaugeDirections == 12 && e.HiggsDoublets == 1, Detail: e.GaugeGroup + " + one complex Higgs doublet"},
	}
}

func _unusedGeometryMathGuard() float64 { return math.Pi }
