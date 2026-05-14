// Package abelianmixing implements Gate 74: the abelian mixing / hypercharge
// coupling-normalization search.
//
// Gate 73 reduced the raw dual-carrier coupling tensor to two symmetry-allowed
// abelian bridge slots:
//
//	central u(1) ↔ contact-u1
//	B-L          ↔ contact-u1
//
// This gate audits whether those two slots can be identified with the finite
// hypercharge bridge already reconstructed in the matter/electroweak gates.
// The result is deliberately strict:
//
//   - B-L enters hypercharge with coefficient 1/2 at the charge-table level.
//   - central u(1) is rejected as a hypercharge component by tracelessness.
//   - the full gauge coupling normalization is still not derived.
//
// In other words, this gate selects the charge-level abelian direction, not the
// physical U(1)_Y kinetic coupling.
package abelianmixing

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/dualcoupling"
	"github.com/bagherbal/asha-engine/pkg/bridge/ewprojection"
	"github.com/bagherbal/asha-engine/pkg/matter/hyperaudit"
)

type Coefficient struct {
	Name        string
	Value       float64
	Selected    bool
	Role        string
	Reason      string
	Obstruction string
}

type Analysis struct {
	DualCoupling dualcoupling.Analysis
	EWProjection ewprojection.Analysis
	HyperAudit   hyperaudit.Analysis

	AbelianBridgeDimension int

	Central Coefficient
	BMinusL Coefficient

	RightSingletTrace             float64
	CentralTraceShiftPerUnit      float64
	CentralRejectedForHypercharge bool
	BMinusLCoefficientSelected    bool
	ChargeLevelHyperchargeBridge  bool

	KY                          float64
	BoundarySin2                float64
	NormalizedHyperchargeFactor float64

	GaugeKineticNormalizationDerived bool
	PhysicalU1CouplingDerived        bool
	ContactU1KineticMixingDerived    bool
	RGBoundaryScaleDerived           bool
	FineStructureDerived             bool
	HiddenObservedInputUsed          bool

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
		dc, err := dualcoupling.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ew, err := ewprojection.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ha, err := hyperaudit.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(dc, ew, ha, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(dc dualcoupling.Analysis, ew ewprojection.Analysis, ha hyperaudit.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !dc.AbelianBridgeDomainExposed || dc.AbelianBridgeDimension != 2 {
		return Analysis{}, fmt.Errorf("Gate 73 did not expose the expected 2D abelian bridge domain")
	}
	if !ha.ChiralOrientationSelected || ha.PreferredBranchName == "" {
		return Analysis{}, fmt.Errorf("hypercharge orientation must be selected before abelian mixing audit")
	}
	if math.Abs(ew.HyperchargeNormalizationKY-5.0/3.0) > 1e-8 {
		return Analysis{}, fmt.Errorf("unexpected kY %.12f; expected the previously derived 5/3 normalization", ew.HyperchargeNormalizationKY)
	}

	trace := traceChargeCounts(ha.Odd.HyperchargeCounts)
	// The odd branch is the selected right-singlet/conjugate table.  Adding a
	// central component c·I to a 16-state generation shifts the trace by 16c.
	centralTraceShift := 16.0
	centralRejected := math.Abs(trace) <= eps && centralTraceShift > 0

	central := Coefficient{
		Name:        "kappa_central",
		Value:       0,
		Selected:    centralRejected,
		Role:        "central u(1) coefficient inside the hypercharge bridge",
		Reason:      "a nonzero central component shifts every one-generation hypercharge equally and destroys the traceless finite table",
		Obstruction: "central u(1) may still exist as a separate universal current, but it is not selected as Standard-Model hypercharge",
	}
	bminusl := Coefficient{
		Name:        "kappa_B-L",
		Value:       0.5,
		Selected:    ha.RightSingletConjugateTableDerived,
		Role:        "B-L coefficient in Y = T3_R + (B-L)/2",
		Reason:      "the odd chiral branch matches the right-singlet/conjugate hypercharge table only with the half-weight B-L coefficient already used by the finite T3_R audit",
		Obstruction: "this is charge normalization, not the physical U(1)_Y gauge coupling",
	}

	truth := "Gate 74 selects the charge-level abelian hypercharge direction inside the two-slot abelian bridge exposed by Gate 73. The finite table rejects central u(1) as a hypercharge component and retains B-L with coefficient 1/2, consistent with Y=T3_R+(B-L)/2 and k_Y=5/3. This still does not derive the U(1)_Y kinetic coupling, contact-u1 kinetic mixing, RG boundary scale, or alpha_em."

	return Analysis{
		DualCoupling: dc,
		EWProjection: ew,
		HyperAudit:   ha,

		AbelianBridgeDimension: dc.AbelianBridgeDimension,
		Central:                central,
		BMinusL:                bminusl,

		RightSingletTrace:                trace,
		CentralTraceShiftPerUnit:         centralTraceShift,
		CentralRejectedForHypercharge:    centralRejected,
		BMinusLCoefficientSelected:       bminusl.Selected,
		ChargeLevelHyperchargeBridge:     centralRejected && bminusl.Selected,
		KY:                               ew.HyperchargeNormalizationKY,
		BoundarySin2:                     ew.EqualNormalizedCouplingBoundarySin2,
		NormalizedHyperchargeFactor:      ew.NormalizedHyperchargeFactor,
		GaugeKineticNormalizationDerived: false,
		PhysicalU1CouplingDerived:        false,
		ContactU1KineticMixingDerived:    false,
		RGBoundaryScaleDerived:           false,
		FineStructureDerived:             false,
		HiddenObservedInputUsed:          false,
		TruthStatement:                   truth,
		RecommendedNextGate:              "Gate 75 — U(1) Kinetic Mixing / Gauge Coupling Hessian Search",
		RemainingUnknowns: []string{
			"U-20D3A-U1-KINETIC-HESSIAN: derive the kinetic matrix for central, B-L, and contact-u1 fields",
			"U-20D3B-CONTACT-U1-NORMALIZATION: connect contact-u1 generator normalization to U(1)_Y kinetic normalization",
			"U-20D3C-CENTRAL-U1-FATE: decide whether central u(1) is projected out, massive, confined to a source sector, or separately gauged",
			"U-20D3D-RG-BOUNDARY: no physical alpha or weak angle follows until the U(1) kinetic matrix and RG scale are derived",
		},
	}, nil
}

func traceChargeCounts(xs []hyperaudit.ChargeCount) float64 {
	tr := 0.0
	for _, x := range xs {
		tr += x.Charge * float64(x.Count)
	}
	return tr
}
