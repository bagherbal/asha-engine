// Package electroweak searches for the next missing operator after the neutral
// Yukawa/intertwiner selection rule.
//
// Gate 18 established that, with the scalar/contact factor currently neutral
// under B-L, neutral intertwiners must preserve B-L charge sectors. The next
// Standard-Model-shaped ingredient is a finite left/right grading and a
// hypercharge bridge. This package constructs the honest part that is already
// canonical in the current Fock bookkeeping: the occupation-parity grading
//
//	Γ_F |n⟩ = (-1)^{N(n)} |n⟩,
//
// lifted to H_Fock ⊗ H_Φ. It behaves like a finite chirality/grading candidate:
// Γ²=I, Tr Γ=0, and it splits the tensor space into equal positive/negative
// sectors. But the package deliberately does not call this the completed
// Standard Model chirality operator, because the true hypercharge bridge still
// needs an independent T3_R / scalar-charge operator.
package electroweak

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/yukawa"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type ChargeParitySector struct {
	Charge   float64
	EvenDim  int
	OddDim   int
	TotalDim int
}

type Analysis struct {
	Yukawa yukawa.Analysis

	MatterDimension int
	ScalarDimension int
	TensorDimension int

	FockGrading   linear.Matrix
	TensorGrading linear.Matrix

	GradingSquareResidual    float64
	GradingTrace             float64
	PositiveGradingDimension int
	NegativeGradingDimension int
	GradingBalanceResidual   int

	CommutesWithBMinusLNorm        float64
	CommutesWithScalarResponseNorm float64
	CommutesWithMatterResponseNorm float64

	ChargeParitySectors                 []ChargeParitySector
	NeutralChiralityPreservingDimension int
	NeutralChiralityFlippingDimension   int
	NeutralChiralityFlippingAvailable   bool

	BMinusLPresent           bool
	FockGradingPresent       bool
	T3ROperatorPresent       bool
	ScalarHyperchargePresent bool
	HyperchargeDerived       bool
	HyperchargeFormula       string
	PhysicalChiralityDerived bool
	ElectroweakYukawaDerived bool
	RemainingUnknowns        []string
}

var (
	electroweakDefaultOnce  sync.Once
	electroweakDefaultValue Analysis
	electroweakDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	electroweakDefaultOnce.Do(func() {
		electroweakDefaultValue, electroweakDefaultErr = buildElectroweakDefaultUncached()
	})
	return electroweakDefaultValue, electroweakDefaultErr
}

func buildElectroweakDefaultUncached() (Analysis, error) {
	y, err := yukawa.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(y, 1e-10)
}

func Build(y yukawa.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	t := y.Tensor
	if t.MatterDimension <= 0 || t.ScalarDimension <= 0 || t.TensorDimension != t.MatterDimension*t.ScalarDimension {
		return Analysis{}, fmt.Errorf("invalid tensor bridge dimensions matter=%d scalar=%d tensor=%d", t.MatterDimension, t.ScalarDimension, t.TensorDimension)
	}

	gammaF, _, _ := fockParityGrading(t.Charge.Action.Bridge.Fock.States)
	gamma := linear.Kronecker(gammaF, linear.Identity(t.ScalarDimension))

	gamma2, err := gamma.Mul(gamma)
	if err != nil {
		return Analysis{}, err
	}
	gamma2Residual, err := gamma2.MaxAbsDiff(linear.Identity(t.TensorDimension))
	if err != nil {
		return Analysis{}, err
	}
	traceGamma, err := gamma.Trace()
	if err != nil {
		return Analysis{}, err
	}
	posDim, negDim := gradingDimensions(gamma, eps)

	commQ, err := linear.Commutator(t.MatterChargeOperator, gamma)
	if err != nil {
		return Analysis{}, err
	}
	commS, err := linear.Commutator(t.ScalarResponse, gamma)
	if err != nil {
		return Analysis{}, err
	}
	commM, err := linear.Commutator(t.MatterResponse, gamma)
	if err != nil {
		return Analysis{}, err
	}

	sectors := chargeParitySectors(t.Charge.Action.Bridge.Fock.States, t.ScalarDimension, eps)
	preserving := 0
	flipping := 0
	for _, s := range sectors {
		preserving += s.EvenDim*s.EvenDim + s.OddDim*s.OddDim
		flipping += 2 * s.EvenDim * s.OddDim
	}

	return Analysis{
		Yukawa:                              y,
		MatterDimension:                     t.MatterDimension,
		ScalarDimension:                     t.ScalarDimension,
		TensorDimension:                     t.TensorDimension,
		FockGrading:                         gammaF,
		TensorGrading:                       gamma,
		GradingSquareResidual:               gamma2Residual,
		GradingTrace:                        traceGamma,
		PositiveGradingDimension:            posDim,
		NegativeGradingDimension:            negDim,
		GradingBalanceResidual:              int(math.Abs(float64(posDim - negDim))),
		CommutesWithBMinusLNorm:             commQ.FrobeniusNorm(),
		CommutesWithScalarResponseNorm:      commS.FrobeniusNorm(),
		CommutesWithMatterResponseNorm:      commM.FrobeniusNorm(),
		ChargeParitySectors:                 sectors,
		NeutralChiralityPreservingDimension: preserving,
		NeutralChiralityFlippingDimension:   flipping,
		NeutralChiralityFlippingAvailable:   flipping > 0,
		BMinusLPresent:                      true,
		FockGradingPresent:                  true,
		T3ROperatorPresent:                  false,
		ScalarHyperchargePresent:            false,
		HyperchargeDerived:                  false,
		HyperchargeFormula:                  "Y = T3_R + (B-L)/2 requires an independent finite T3_R/scalar-charge bridge",
		PhysicalChiralityDerived:            false,
		ElectroweakYukawaDerived:            false,
		RemainingUnknowns: []string{
			"U-06-HYPERCHARGE: construct finite T3_R or equivalent weak-isospin/hypercharge operator",
			"U-11-CHIRALITY: upgrade occupation parity Γ_F into a physical left/right chirality projector",
			"U-12-SCALAR-CHARGE: assign a non-neutral electroweak charge to the scalar/contact factor if charge-changing Yukawa maps are required",
			"U-07-YUKAWA: build a gauge-compatible chirality-flipping intertwiner after hypercharge is known",
		},
	}, nil
}

func fockParityGrading(states []spinor.FockState) (linear.Matrix, int, int) {
	g := linear.NewMatrix(len(states), len(states))
	even := 0
	odd := 0
	for i, state := range states {
		value := -1.0
		if state.ExcitationNumber()%2 == 0 {
			value = 1.0
			even++
		} else {
			odd++
		}
		g.Set(i, i, value)
	}
	return g, even, odd
}

func gradingDimensions(g linear.Matrix, eps float64) (positive int, negative int) {
	for i := 0; i < g.Rows() && i < g.Cols(); i++ {
		v := g.At(i, i)
		if v > eps {
			positive++
		}
		if v < -eps {
			negative++
		}
	}
	return positive, negative
}

func chargeParitySectors(states []spinor.FockState, scalarDim int, eps float64) []ChargeParitySector {
	sectors := make([]ChargeParitySector, 0)
	for _, state := range states {
		q := state.BMinusL()
		if math.Abs(q) < eps {
			q = 0
		}
		idx := -1
		for i := range sectors {
			if math.Abs(sectors[i].Charge-q) <= eps {
				idx = i
				break
			}
		}
		if idx < 0 {
			sectors = append(sectors, ChargeParitySector{Charge: q})
			idx = len(sectors) - 1
		}
		if state.ExcitationNumber()%2 == 0 {
			sectors[idx].EvenDim += scalarDim
		} else {
			sectors[idx].OddDim += scalarDim
		}
		sectors[idx].TotalDim += scalarDim
	}
	sort.Slice(sectors, func(i, j int) bool { return sectors[i].Charge < sectors[j].Charge })
	return sectors
}

func FormatChargeParitySectors(sectors []ChargeParitySector) string {
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
