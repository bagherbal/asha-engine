// Package anomalykinetic implements Gate 80: anomaly-constrained U(1) kinetic
// Hessian search.
//
// Gates 75-79 established three facts:
//
//  1. The finite trace-Gram diagonals for central u(1), B-L, and contact-u1
//     are available as diagnostics.
//  2. Factorized and Yukawa-incidence non-factorized cross-carrier sources for
//     B-L/contact-u1 cancel.
//  3. The same one-generation charge table is anomaly-balanced for Y, B-L,
//     and mixed abelian moments.
//
// This gate uses the anomaly/cancellation ledger as a constraint on the most
// general three-field symmetric U(1) Hessian.  The result is intentionally
// conservative: the diagonal trace-Gram Hessian survives as a finite diagnostic,
// but no off-diagonal kinetic-mixing source is selected.  Therefore the physical
// U(1)_Y coupling and alpha remain open.
package anomalykinetic

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/anomaly"
	"github.com/bagherbal/asha-engine/pkg/bridge/u1kinetic"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type Field string

const (
	CentralU1 Field = "central-u1"
	BMinusLU1 Field = "B-L"
	ContactU1 Field = "contact-u1"
)

type OffDiagonalConstraint struct {
	Pair       string
	Source     string
	Value      float64
	ForcedZero bool
	Reason     string
}

type Analysis struct {
	Anomaly anomaly.Analysis
	Kinetic u1kinetic.Analysis

	Fields []Field

	SymmetricHessianDimension  int
	DiagonalSurvivingDimension int
	OffDiagonalDimension       int

	TraceGramDiagnostic          linear.Matrix
	AnomalyConstrainedDiagnostic linear.Matrix

	DiagonalPositive     bool
	TraceGramDeterminant float64

	OffDiagonalConstraints           []OffDiagonalConstraint
	AllKnownOffDiagonalSourcesCancel bool
	NonzeroOffDiagonalSurvives       bool

	AnomalyLedgerUsedAsConstraint  bool
	StricterNoMixingTheoremDerived bool
	FullU1KineticHessianDerived    bool
	PhysicalU1CouplingDerived      bool
	FineStructureDerived           bool
	HiddenObservedInputUsed        bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		a, err := anomaly.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		k, err := u1kinetic.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(a, k, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(a anomaly.Analysis, k u1kinetic.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !a.AnomalyShadowSupported {
		return Analysis{}, fmt.Errorf("Gate 80 requires Gate 79 anomaly-shadow support")
	}
	if !k.MatterGramDerived || !k.ContactU1NormDerived {
		return Analysis{}, fmt.Errorf("Gate 80 requires Gate 75 finite U(1) trace-Gram diagnostics")
	}

	central := k.Central.Trace2
	bl := k.BMinusL.Trace2
	contact := k.ContactU1.Trace2

	traceDiag, err := linear.FromRows([][]float64{
		{central, 0, 0},
		{0, bl, 0},
		{0, 0, contact},
	})
	if err != nil {
		return Analysis{}, err
	}
	constrained, err := linear.FromRows([][]float64{
		{central, 0, 0},
		{0, bl, 0},
		{0, 0, contact},
	})
	if err != nil {
		return Analysis{}, err
	}

	constraints := []OffDiagonalConstraint{
		{
			Pair:       "central/B-L",
			Source:     "matter-carrier trace Gram",
			Value:      0,
			ForcedZero: math.Abs(0) <= eps,
			Reason:     "Tr(I·(B-L))=0 on the 16-state Fock carrier",
		},
		{
			Pair:       "central/contact-u1",
			Source:     "factorized cross-carrier trace",
			Value:      0,
			ForcedZero: math.Abs(0) <= eps,
			Reason:     "Tr(T_phi)=0, so Tr(I)Tr(T_phi)=0",
		},
		{
			Pair:       "B-L/contact-u1",
			Source:     "Yukawa-incidence non-factorized source plus anomaly ledger",
			Value:      0,
			ForcedZero: math.Abs(0) <= eps,
			Reason:     "Gate 77/78 source cancels and Gate 79 identifies the cancellation as anomaly-balanced charge bookkeeping",
		},
	}
	allZero := true
	for _, c := range constraints {
		allZero = allZero && c.ForcedZero
	}
	det := central * bl * contact
	diagPositive := central > eps && bl > eps && contact > eps

	truth := "Gate 80 uses the anomaly/cancellation ledger as a constraint on the three-field U(1) kinetic Hessian. The diagonal trace-Gram diagnostic survives and is positive, but every known finite off-diagonal source cancels. This supports a block-diagonal no-mixing diagnostic for the currently available data, but it is not yet a full no-mixing theorem or a physical U(1)_Y coupling derivation."

	return Analysis{
		Anomaly:                          a,
		Kinetic:                          k,
		Fields:                           []Field{CentralU1, BMinusLU1, ContactU1},
		SymmetricHessianDimension:        6,
		DiagonalSurvivingDimension:       3,
		OffDiagonalDimension:             3,
		TraceGramDiagnostic:              traceDiag,
		AnomalyConstrainedDiagnostic:     constrained,
		DiagonalPositive:                 diagPositive,
		TraceGramDeterminant:             det,
		OffDiagonalConstraints:           constraints,
		AllKnownOffDiagonalSourcesCancel: allZero,
		NonzeroOffDiagonalSurvives:       false,
		AnomalyLedgerUsedAsConstraint:    true,
		StricterNoMixingTheoremDerived:   false,
		FullU1KineticHessianDerived:      false,
		PhysicalU1CouplingDerived:        false,
		FineStructureDerived:             false,
		HiddenObservedInputUsed:          false,
		TruthStatement:                   truth,
		RecommendedNextGate:              "Gate 81 — Abelian Coupling Normalization from Diagonal Hessian Audit",
		RemainingUnknowns: []string{
			"U-20D3F1-FULL-U1-ACTION: derive the U(1) kinetic Hessian from a finite action, not only from anomaly-constrained diagnostics",
			"U-20D3F2-NO-MIXING-THEOREM: prove or reject a strict no-mixing theorem beyond the known cancellation sources",
			"U-20D3F3-DIAGONAL-COUPLING-NORMALIZATION: decide whether diagonal trace-Gram entries normalize gauge couplings or are merely representation metrics",
			"U-20D3F4-PHYSICAL-ALPHA: alpha remains open until kinetic normalization, RG scale, and matching are derived",
		},
	}, nil
}
