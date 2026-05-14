// Package ewprojection audits the electroweak neutral projection after the
// finite SU(2)_L and hypercharge table have been constructed.
//
// The package deliberately separates three objects that are often conflated:
//
//  1. the charge identity Q = T3_L + Y, which is a representation theorem;
//  2. trace-metric directions inside the finite charge table, which are
//     dimensionless representation diagnostics;
//  3. the physical weak mixing angle, which also requires gauge kinetic
//     normalizations g and g' plus an RG scale.
//
// The finite charge table does contain a notable normalization result: when the
// full one-generation table is used, Tr(Y^2)/Tr(T3_L^2)=5/3. This is the usual
// hypercharge normalization factor. Under the additional bridge hypothesis of
// equal normalized gauge couplings at a boundary scale, it gives sin^2(theta)=3/8.
// That is a boundary-candidate diagnostic, not a measured low-energy prediction.
package ewprojection

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/matter/su2lgauge"
)

type Analysis struct {
	SU2L su2lgauge.Analysis

	LeftDimension int

	T3Trace2Left float64
	YTrace2Left  float64
	QTrace2Left  float64
	T3YTraceLeft float64

	ChargeIdentityResidual          float64
	HyperchargeCommutesWithSU2LNorm float64

	// TraceDirectionSin2 is the squared projection of the finite electric charge
	// generator Q=T3+Y onto the hypercharge axis using the unnormalized left-
	// doublet trace metric. It is a representation-direction diagnostic, not the
	// physical weak mixing angle.
	TraceDirectionSin2 float64
	TraceDirectionCos2 float64

	RightYTrace2Standard        float64
	FullYTrace2OneGeneration    float64
	FullT3Trace2OneGeneration   float64
	HyperchargeNormalizationKY  float64
	NormalizedHyperchargeFactor float64
	NormalizedYTrace2           float64

	EqualNormalizedCouplingBoundarySin2      float64
	EqualNormalizedCouplingBoundaryCandidate bool

	ElectromagneticGeneratorDerived  bool
	GaugeKineticNormalizationDerived bool
	WeakMixingAngleDerived           bool
	RGBoundaryScaleDerived           bool
	FineStructureDerived             bool
	HiddenObservedInputUsed          bool

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		s, err := su2lgauge.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(s, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(s su2lgauge.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !s.NonabelianSU2LGeneratorsDerived {
		return Analysis{}, fmt.Errorf("SU(2)_L finite generator audit must pass first")
	}
	if len(s.States) == 0 {
		return Analysis{}, fmt.Errorf("SU(2)_L state table is empty")
	}

	t3y := 0.0
	t3sq := 0.0
	ysq := 0.0
	qsq := 0.0
	for _, st := range s.States {
		t3sq += st.T3 * st.T3
		ysq += st.Hypercharge * st.Hypercharge
		q := st.T3 + st.Hypercharge
		qsq += q * q
		t3y += st.T3 * st.Hypercharge
	}
	if math.Abs(qsq-(t3sq+ysq+2*t3y)) > 1e-8 {
		return Analysis{}, fmt.Errorf("charge trace decomposition failed")
	}
	traceSin2 := 0.0
	traceCos2 := 0.0
	if qsq > eps {
		traceSin2 = ysq / qsq
		traceCos2 = t3sq / qsq
	}

	o := s.Audit.Standard
	rightY2 := 0.0
	rightY2 += float64(o.Up.Multiplicity) * o.Up.RightCharge * o.Up.RightCharge
	rightY2 += float64(o.Down.Multiplicity) * o.Down.RightCharge * o.Down.RightCharge
	rightY2 += float64(o.Neutrino.Multiplicity) * o.Neutrino.RightCharge * o.Neutrino.RightCharge
	rightY2 += float64(o.Electron.Multiplicity) * o.Electron.RightCharge * o.Electron.RightCharge
	fullY2 := ysq + rightY2
	fullT3 := t3sq // right singlets are SU(2)_L singlets.
	kY := 0.0
	normFactor := 0.0
	normY2 := 0.0
	equalBoundarySin2 := 0.0
	if fullT3 > eps && fullY2 > eps {
		kY = fullY2 / fullT3
		normFactor = math.Sqrt(1 / kY)
		normY2 = normFactor * normFactor * fullY2
		// Equal normalized couplings: g_2=g_1 with g_1^2=kY*g_Y^2.
		// Therefore sin^2(theta)=g_Y^2/(g_2^2+g_Y^2)=1/(1+kY).
		equalBoundarySin2 = 1 / (1 + kY)
	}

	return Analysis{
		SU2L:                                     s,
		LeftDimension:                            len(s.States),
		T3Trace2Left:                             t3sq,
		YTrace2Left:                              ysq,
		QTrace2Left:                              qsq,
		T3YTraceLeft:                             t3y,
		ChargeIdentityResidual:                   s.GellMannNishijimaNorm,
		HyperchargeCommutesWithSU2LNorm:          s.CommutesWithHyperchargeNorm,
		TraceDirectionSin2:                       traceSin2,
		TraceDirectionCos2:                       traceCos2,
		RightYTrace2Standard:                     rightY2,
		FullYTrace2OneGeneration:                 fullY2,
		FullT3Trace2OneGeneration:                fullT3,
		HyperchargeNormalizationKY:               kY,
		NormalizedHyperchargeFactor:              normFactor,
		NormalizedYTrace2:                        normY2,
		EqualNormalizedCouplingBoundarySin2:      equalBoundarySin2,
		EqualNormalizedCouplingBoundaryCandidate: kY > eps,
		ElectromagneticGeneratorDerived:          s.GellMannNishijimaNorm <= eps,
		GaugeKineticNormalizationDerived:         false,
		WeakMixingAngleDerived:                   false,
		RGBoundaryScaleDerived:                   false,
		FineStructureDerived:                     false,
		HiddenObservedInputUsed:                  false,
		TruthStatement:                           "The finite charge table derives Q=T3_L+Y and the full-generation hypercharge normalization k_Y=5/3. It also gives trace-direction diagnostics such as sin²=1/4 on the left doublet sector and a boundary-candidate sin²=3/8 under equal normalized couplings. None of these is yet the physical weak mixing angle because the gauge kinetic normalization, RG boundary scale, and running are still open.",
		RemainingUnknowns: []string{
			"U-20B-TRACE-KINETIC-NORMALIZATION: derive the kinetic normalization of SU(2)_L and U(1)_Y from the finite action, not only from charge traces",
			"U-20C-RG-BOUNDARY: identify the scale at which any equal-normalized-coupling boundary condition applies",
			"U-20D-ELECTROWEAK-MIXING: derive the physical weak mixing angle after coupling normalization and RG flow",
			"U-20E-FINE-STRUCTURE: compute alpha_em only after electromagnetic projection, normalization, and scale are fixed",
		},
	}, nil
}
