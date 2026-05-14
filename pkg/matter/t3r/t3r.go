// Package t3r searches for the matter-side weak-isospin bridge that Gate 20
// proved was still missing.
//
// The scalar/contact side already supplies a canonical 2+2 scalar charge
// T_Φ=diag(+1/2,+1/2,-1/2,-1/2).  But the matter side still needs an operator
// that plays the role of T3_R in the Pati-Salam relation
//
//	Y = T3_R + (B-L)/2.
//
// This package tests the simplest finite operators available in the current
// typed Fock basis.  The temporal-occupation polarization
//
//	T0 |n0,n1,n2,n3⟩ = (1/2 - n0)|n0,n1,n2,n3⟩
//
// is canonical, traceless, color-blind, and commutes with B-L.  By itself it is
// vectorlike: it still does not make occupation-parity Γ_F into physical
// electroweak chirality.  The next honest possibility is a chiral restriction
// of this temporal polarization to one Γ_F sector.  Both even and odd choices
// work algebraically, so the engine records a mirror ambiguity rather than
// pretending that the physical orientation has already been derived.
package t3r

import (
	"sync"

	"fmt"
	"math"
	"sort"

	"github.com/bagherbal/asha-engine/pkg/linear"
	"github.com/bagherbal/asha-engine/pkg/matter/hypercharge"
	"github.com/bagherbal/asha-engine/pkg/spinor"
)

type Sector struct {
	Charge   float64
	EvenDim  int
	OddDim   int
	TotalDim int
}

type Candidate struct {
	Name string

	MatterT3R         linear.Matrix
	MatterHypercharge linear.Matrix
	TotalCharge       linear.Matrix

	TraceT3R                float64
	TraceT3RSquared         float64
	CommutesWithBMinusLNorm float64
	CommutesWithGradingNorm float64

	Sectors           []Sector
	PreservingDim     int
	FlippingDim       int
	FlippingAvailable bool
}

type Analysis struct {
	Hypercharge hypercharge.Analysis

	MatterDimension int
	ScalarDimension int
	TensorDimension int

	TemporalPolarization            linear.Matrix
	TemporalSpectrum                []float64
	TemporalTrace                   float64
	TemporalTraceSquared            float64
	TemporalCommutesWithBMinusLNorm float64
	TemporalCommutesWithGradingNorm float64

	Vectorlike Candidate
	ChiralEven Candidate
	ChiralOdd  Candidate

	MatterSideOperatorFound               bool
	VectorlikeTemporalIsPhysicalChirality bool
	ChiralRestrictedBridgeAvailable       bool
	PhysicalOrientationSelected           bool
	MirrorAmbiguity                       bool
	HyperchargeCandidateConstructed       bool
	ElectroweakYukawaDerived              bool
	RemainingUnknowns                     []string
}

var (
	t3rDefaultOnce  sync.Once
	t3rDefaultValue Analysis
	t3rDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	t3rDefaultOnce.Do(func() {
		t3rDefaultValue, t3rDefaultErr = buildT3rDefaultUncached()
	})
	return t3rDefaultValue, t3rDefaultErr
}

func buildT3rDefaultUncached() (Analysis, error) {
	h, err := hypercharge.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(h, 1e-10)
}

func Build(h hypercharge.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	t := h.Electroweak.Yukawa.Tensor
	states := t.Charge.Action.Bridge.Fock.States
	if len(states) != t.MatterDimension {
		return Analysis{}, fmt.Errorf("state count %d does not match matter dimension %d", len(states), t.MatterDimension)
	}
	if t.ScalarDimension != 4 {
		return Analysis{}, fmt.Errorf("T3_R search expects scalar dimension 4, got %d", t.ScalarDimension)
	}

	temporal := temporalPolarization(states)
	traceT, _ := temporal.Trace()
	traceT2 := diagonalTraceSquared(temporal)

	commB, err := linear.Commutator(temporal, t.Charge.FockChargeOperator)
	if err != nil {
		return Analysis{}, err
	}
	commG, err := linear.Commutator(temporal, h.Electroweak.FockGrading)
	if err != nil {
		return Analysis{}, err
	}

	evenProjector, oddProjector := parityProjectors(states)
	evenTemporal, err := evenProjector.Mul(temporal)
	if err != nil {
		return Analysis{}, err
	}
	oddTemporal, err := oddProjector.Mul(temporal)
	if err != nil {
		return Analysis{}, err
	}

	vectorlike, err := buildCandidate("vectorlike temporal T0", temporal, h, eps)
	if err != nil {
		return Analysis{}, err
	}
	chiralEven, err := buildCandidate("even-sector temporal T3_R candidate", evenTemporal, h, eps)
	if err != nil {
		return Analysis{}, err
	}
	chiralOdd, err := buildCandidate("odd-sector temporal T3_R candidate", oddTemporal, h, eps)
	if err != nil {
		return Analysis{}, err
	}

	spectrum := make([]float64, len(states))
	for i := range states {
		spectrum[i] = temporal.At(i, i)
	}

	return Analysis{
		Hypercharge:                           h,
		MatterDimension:                       t.MatterDimension,
		ScalarDimension:                       t.ScalarDimension,
		TensorDimension:                       t.TensorDimension,
		TemporalPolarization:                  temporal,
		TemporalSpectrum:                      spectrum,
		TemporalTrace:                         traceT,
		TemporalTraceSquared:                  traceT2,
		TemporalCommutesWithBMinusLNorm:       commB.FrobeniusNorm(),
		TemporalCommutesWithGradingNorm:       commG.FrobeniusNorm(),
		Vectorlike:                            vectorlike,
		ChiralEven:                            chiralEven,
		ChiralOdd:                             chiralOdd,
		MatterSideOperatorFound:               true,
		VectorlikeTemporalIsPhysicalChirality: false,
		ChiralRestrictedBridgeAvailable:       chiralEven.FlippingAvailable && chiralOdd.FlippingAvailable,
		PhysicalOrientationSelected:           false,
		MirrorAmbiguity:                       true,
		HyperchargeCandidateConstructed:       true,
		ElectroweakYukawaDerived:              false,
		RemainingUnknowns: []string{
			"U-11A-CHIRAL-ORIENTATION: decide whether the physical right-handed sector is the even or odd Fock parity branch, or replace Γ_F by a better chirality operator",
			"U-06B-HYPERCHARGE-NORMALIZATION: connect the finite T3_R candidate to a full Standard Model hypercharge table, including conjugate-state conventions",
			"U-07-YUKAWA: construct the actual gauge-compatible chirality-changing intertwiner after the orientation is fixed",
		},
	}, nil
}

func temporalPolarization(states []spinor.FockState) linear.Matrix {
	m := linear.NewMatrix(len(states), len(states))
	for i, s := range states {
		value := 0.5
		if len(s.Occupation) > 0 && s.Occupation[0] {
			value = -0.5
		}
		m.Set(i, i, value)
	}
	return m
}

func parityProjectors(states []spinor.FockState) (linear.Matrix, linear.Matrix) {
	even := linear.NewMatrix(len(states), len(states))
	odd := linear.NewMatrix(len(states), len(states))
	for i, s := range states {
		if s.ExcitationNumber()%2 == 0 {
			even.Set(i, i, 1)
		} else {
			odd.Set(i, i, 1)
		}
	}
	return even, odd
}

func buildCandidate(name string, t3 linear.Matrix, h hypercharge.Analysis, eps float64) (Candidate, error) {
	tensor := h.Electroweak.Yukawa.Tensor
	matterY, err := t3.Add(tensor.Charge.FockChargeOperator.Scale(0.5))
	if err != nil {
		return Candidate{}, err
	}
	total, err := linear.Kronecker(matterY, linear.Identity(tensor.ScalarDimension)).Add(linear.Kronecker(linear.Identity(tensor.MatterDimension), h.ScalarWeakChargeOperator))
	if err != nil {
		return Candidate{}, err
	}

	commB, err := linear.Commutator(t3, tensor.Charge.FockChargeOperator)
	if err != nil {
		return Candidate{}, err
	}
	commG, err := linear.Commutator(t3, h.Electroweak.FockGrading)
	if err != nil {
		return Candidate{}, err
	}

	sectors, preserving, flipping := sectorsForCharge(total, h.Electroweak.TensorGrading, eps)
	traceT, _ := t3.Trace()
	return Candidate{
		Name:                    name,
		MatterT3R:               t3,
		MatterHypercharge:       matterY,
		TotalCharge:             total,
		TraceT3R:                traceT,
		TraceT3RSquared:         diagonalTraceSquared(t3),
		CommutesWithBMinusLNorm: commB.FrobeniusNorm(),
		CommutesWithGradingNorm: commG.FrobeniusNorm(),
		Sectors:                 sectors,
		PreservingDim:           preserving,
		FlippingDim:             flipping,
		FlippingAvailable:       flipping > 0,
	}, nil
}

func diagonalTraceSquared(m linear.Matrix) float64 {
	n := m.Rows()
	if m.Cols() < n {
		n = m.Cols()
	}
	sum := 0.0
	for i := 0; i < n; i++ {
		v := m.At(i, i)
		sum += v * v
	}
	return sum
}

func sectorsForCharge(charge linear.Matrix, grading linear.Matrix, eps float64) ([]Sector, int, int) {
	sectors := make([]Sector, 0)
	for i := 0; i < charge.Rows(); i++ {
		q := charge.At(i, i)
		if math.Abs(q) < eps {
			q = 0
		}
		idx := -1
		for j := range sectors {
			if math.Abs(sectors[j].Charge-q) <= eps {
				idx = j
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
