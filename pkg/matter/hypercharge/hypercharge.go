// Package hypercharge searches for the first finite scalar-charge bridge after
// the electroweak operator search.
//
// Gate 19 established two facts:
//   - the current Fock bookkeeping has a useful grading Γ_F=(-1)^N;
//   - B-L plus a neutral scalar factor cannot produce chirality-changing
//     electroweak-style intertwiners.
//
// The finite Higgs/contact sector, however, has a pair-degenerate 2+2 active
// spectrum. This package treats that 2+2 structure as a realification of a
// two-component complex scalar doublet and constructs the canonical trace-zero
// scalar weak-charge candidate on H_Φ:
//
//	T_Φ = diag(+w,+w,-w,-w),  w = 1 / pairMultiplicity.
//
// Since the pair multiplicity is 2, w=1/2 is not fitted from particle data; it
// is the fundamental SU(2) weight forced by the 2+2 scalar-doublet structure.
//
// The package is careful about what this solves. It gives a scalar-charge
// bridge. It does not derive the matter-side T3_R operator, Standard Model
// hypercharge, or a physical chirality operator. It also checks that the usual
// Pati-Salam-shaped template (B-L)/2 + T_Φ is still insufficient if Γ_F is used
// as physical chirality, which means Γ_F is only a grading candidate.
package hypercharge

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/electroweak"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type Sector struct {
	Charge   float64
	EvenDim  int
	OddDim   int
	TotalDim int
}

type Analysis struct {
	Electroweak electroweak.Analysis

	MatterDimension int
	ScalarDimension int
	TensorDimension int

	PairMultiplicity  int
	FundamentalWeight float64

	ScalarWeakChargeOperator linear.Matrix // T_Φ on the 4D active scalar/contact factor.
	TensorScalarCharge       linear.Matrix // I_Fock ⊗ T_Φ.

	ScalarChargeSpectrum []float64
	ScalarChargeTrace    float64
	ScalarChargeTrace2   float64
	ScalarChargeClusters []int

	CommutesWithScalarResponseNorm float64
	CommutesWithBMinusLNorm        float64
	CommutesWithFockGradingNorm    float64

	// Standard Pati-Salam-shaped diagnostic: Q_PS = (B-L)/2 + T_Φ.
	PatiSalamCoefficient   float64
	PatiSalamCharge        linear.Matrix
	PatiSalamSectors       []Sector
	PatiSalamFlippingDim   int
	PatiSalamPreservingDim int

	// Raw compensator diagnostic: Q_raw = (B-L) + T_Φ. This is not Standard
	// Model hypercharge. It is included only to show that scalar charge can
	// create grading-flipping channels once the correct matter charge operator
	// exists.
	RawCoefficient   float64
	RawCharge        linear.Matrix
	RawSectors       []Sector
	RawFlippingDim   int
	RawPreservingDim int

	ScalarDoubletCandidate          bool
	ScalarChargeBridgeConstructed   bool
	MatterT3ROperatorConstructed    bool
	StandardModelHyperchargeDerived bool
	PhysicalChiralityDerived        bool
	ElectroweakYukawaDerived        bool
	RemainingUnknowns               []string
}

var (
	hyperchargeDefaultOnce  sync.Once
	hyperchargeDefaultValue Analysis
	hyperchargeDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	hyperchargeDefaultOnce.Do(func() {
		hyperchargeDefaultValue, hyperchargeDefaultErr = buildHyperchargeDefaultUncached()
	})
	return hyperchargeDefaultValue, hyperchargeDefaultErr
}

func buildHyperchargeDefaultUncached() (Analysis, error) {
	ew, err := electroweak.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(ew, 1e-10)
}

func Build(ew electroweak.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	t := ew.Yukawa.Tensor
	if t.ScalarDimension != 4 {
		return Analysis{}, fmt.Errorf("scalar-charge bridge expects a 4D active scalar/contact factor, got %d", t.ScalarDimension)
	}

	weights := append([]float64(nil), t.ScalarOneParticleWeights...)
	if len(weights) != 4 {
		return Analysis{}, fmt.Errorf("expected four scalar/contact weights, got %d", len(weights))
	}
	clusters := multiplicityClusters(weights, eps)
	pairMultiplicity := commonPairMultiplicity(clusters)
	if pairMultiplicity <= 0 {
		return Analysis{}, fmt.Errorf("active scalar spectrum is not a 2+2 pair-degenerate doublet candidate: clusters=%v", clusters)
	}
	fundamentalWeight := 1.0 / float64(pairMultiplicity)

	scalarCharges := []float64{fundamentalWeight, fundamentalWeight, -fundamentalWeight, -fundamentalWeight}
	tPhi := linear.Diagonal(scalarCharges)
	tensorTPhi := linear.Kronecker(linear.Identity(t.MatterDimension), tPhi)

	scalarResponseOne := linear.Diagonal(weights)
	commScalar, err := linear.Commutator(tPhi, scalarResponseOne)
	if err != nil {
		return Analysis{}, err
	}
	commB, err := linear.Commutator(t.MatterChargeOperator, tensorTPhi)
	if err != nil {
		return Analysis{}, err
	}
	commG, err := linear.Commutator(ew.TensorGrading, tensorTPhi)
	if err != nil {
		return Analysis{}, err
	}

	traceT, _ := tPhi.Trace()
	traceT2 := 0.0
	for _, q := range scalarCharges {
		traceT2 += q * q
	}

	// The coefficient 1/pairMultiplicity is the same representation-theoretic
	// half-weight. It is not a measured constant.
	psCoeff := fundamentalWeight
	psMatter := t.MatterChargeOperator.Scale(psCoeff)
	psCharge, err := psMatter.Add(tensorTPhi)
	if err != nil {
		return Analysis{}, err
	}
	psSectors, psPres, psFlip := sectorsForCharge(psCharge, ew.TensorGrading, eps)

	rawCoeff := 1.0
	rawMatter := t.MatterChargeOperator.Scale(rawCoeff)
	rawCharge, err := rawMatter.Add(tensorTPhi)
	if err != nil {
		return Analysis{}, err
	}
	rawSectors, rawPres, rawFlip := sectorsForCharge(rawCharge, ew.TensorGrading, eps)

	return Analysis{
		Electroweak:                     ew,
		MatterDimension:                 t.MatterDimension,
		ScalarDimension:                 t.ScalarDimension,
		TensorDimension:                 t.TensorDimension,
		PairMultiplicity:                pairMultiplicity,
		FundamentalWeight:               fundamentalWeight,
		ScalarWeakChargeOperator:        tPhi,
		TensorScalarCharge:              tensorTPhi,
		ScalarChargeSpectrum:            scalarCharges,
		ScalarChargeTrace:               traceT,
		ScalarChargeTrace2:              traceT2,
		ScalarChargeClusters:            clusters,
		CommutesWithScalarResponseNorm:  commScalar.FrobeniusNorm(),
		CommutesWithBMinusLNorm:         commB.FrobeniusNorm(),
		CommutesWithFockGradingNorm:     commG.FrobeniusNorm(),
		PatiSalamCoefficient:            psCoeff,
		PatiSalamCharge:                 psCharge,
		PatiSalamSectors:                psSectors,
		PatiSalamPreservingDim:          psPres,
		PatiSalamFlippingDim:            psFlip,
		RawCoefficient:                  rawCoeff,
		RawCharge:                       rawCharge,
		RawSectors:                      rawSectors,
		RawPreservingDim:                rawPres,
		RawFlippingDim:                  rawFlip,
		ScalarDoubletCandidate:          pairMultiplicity == 2,
		ScalarChargeBridgeConstructed:   true,
		MatterT3ROperatorConstructed:    false,
		StandardModelHyperchargeDerived: false,
		PhysicalChiralityDerived:        false,
		ElectroweakYukawaDerived:        false,
		RemainingUnknowns: []string{
			"U-06A-MATTER-T3R: construct a matter-side finite T3_R operator, not only a scalar doublet charge",
			"U-11-CHIRALITY: replace occupation parity Γ_F with the physical left/right chirality projector",
			"U-07-YUKAWA: build a gauge-compatible chirality-changing intertwiner after matter T3_R and chirality are known",
		},
	}, nil
}

func multiplicityClusters(values []float64, eps float64) []int {
	v := append([]float64(nil), values...)
	sort.Sort(sort.Reverse(sort.Float64Slice(v)))
	out := make([]int, 0)
	for i := 0; i < len(v); {
		count := 1
		base := v[i]
		j := i + 1
		for ; j < len(v); j++ {
			if math.Abs(v[j]-base) > eps {
				break
			}
			count++
		}
		out = append(out, count)
		i = j
	}
	return out
}

func commonPairMultiplicity(clusters []int) int {
	if len(clusters) != 2 {
		return 0
	}
	if clusters[0] != clusters[1] {
		return 0
	}
	return clusters[0]
}

func sectorsForCharge(charge linear.Matrix, grading linear.Matrix, eps float64) ([]Sector, int, int) {
	sectors := make([]Sector, 0)
	n := charge.Rows()
	for i := 0; i < n; i++ {
		q := charge.At(i, i)
		if math.Abs(q) < eps {
			q = 0
		}
		idx := -1
		for s := range sectors {
			if math.Abs(sectors[s].Charge-q) <= eps {
				idx = s
				break
			}
		}
		if idx < 0 {
			sectors = append(sectors, Sector{Charge: q})
			idx = len(sectors) - 1
		}
		if grading.At(i, i) >= 0 {
			sectors[idx].EvenDim++
		} else {
			sectors[idx].OddDim++
		}
		sectors[idx].TotalDim++
	}
	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Charge < sectors[j].Charge })
	preserving, flipping := 0, 0
	for _, s := range sectors {
		preserving += s.EvenDim*s.EvenDim + s.OddDim*s.OddDim
		flipping += 2 * s.EvenDim * s.OddDim
	}
	return sectors, preserving, flipping
}

func FormatFloatSlice(values []float64) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10g", v)
	}
	return out + "]"
}

func FormatIntSlice(values []int) string {
	out := "["
	for i, v := range values {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%d", v)
	}
	return out + "]"
}

func FormatSectors(sectors []Sector) string {
	out := "["
	for i, s := range sectors {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("q=%.6g: even=%d odd=%d total=%d", s.Charge, s.EvenDim, s.OddDim, s.TotalDim)
	}
	return out + "]"
}

func FormatUnknowns(unknowns []string) string {
	out := "["
	for i, u := range unknowns {
		if i > 0 {
			out += "; "
		}
		out += u
	}
	return out + "]"
}

// Avoid an otherwise-unused import in documentation-only paths when future
// refactors no longer mention spinor directly.
var _ = spinor.FockState{}
