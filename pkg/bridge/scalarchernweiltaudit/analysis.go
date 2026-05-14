// Package scalarchernweiltaudit implements Gate 192: sealed scalar-bundle
// Chern-Weil carrier / heat-kernel preflight audit.
//
// Gate 191 deliberately inserted the eta orientation as an explicit
// SpontaneousOrientationSeal and used it to trivialize the H_Phi scalar bundle.
// Gate 192 asks a narrower question: once that sealed bundle exists, can it
// carry finite matrix traces that look like the local algebraic ingredients of
// Chern-Weil and heat-kernel gauge terms, without importing the 8π² topological
// seal, continuum integration, spectral cutoffs, thresholds, or physical
// couplings?
package scalarchernweiltaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarorientationseal"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type CurvatureTraceRecord struct {
	Name                  string
	Basis                 string
	GeneratorMatrix       string
	KineticOperator       string
	TotalTrace            float64
	HighFiberTrace        float64
	LowFiberTrace         float64
	EtaGradedTrace        float64
	ExpectedTotalRational string
	ExpectedHighRational  string
	ExpectedLowRational   string
	ExpectedEtaRational   string
	StableRational        bool
	PreservesFibers       bool
	MixesFibers           bool
	PhysicalMeaning       string
}

type FiniteCurvatureTraceAudit struct {
	SealStatus                        string
	TraceFunctional                   string
	GeneratorsAudited                 []CurvatureTraceRecord
	PrimitiveGaugeKineticTracesStable bool
	T1TraceStable                     bool
	T2TraceStable                     bool
	T3TraceStable                     bool
	YPhiTraceStable                   bool
	NeutralQTraceStable               bool
	NeutralZTraceStable               bool
	HighLowFiberTraceAvailable        bool
	GaugeKineticCarrierPreflight      bool
	PhysicalGaugeCouplingsDerived     bool
	Verdict                           string
}

type OrientationGradingAudit struct {
	EtaMatrix                          string
	EtaSquaredResidual                 float64
	EtaTrace                           float64
	EtaHighEigenvalue                  float64
	EtaLowEigenvalue                   float64
	EtaDerivedFromSeal                 bool
	EtaIsPhysicalBoundaryData          bool
	PrimitiveDiagonalGradedTracesZero  bool
	ChargedGradedTracesZero            bool
	NeutralMixedPairName               string
	NeutralMixedGradedTrace            float64
	NeutralMixedExpectedRational       string
	NeutralSplitQGradedTrace           float64
	NeutralSplitZGradedTrace           float64
	NontrivialSignedNeutralCarrier     bool
	IntegerTopologicalChargeMapDerived bool
	ContinuumOrientationDerived        bool
	Verdict                            string
}

type HeatKernelPreflightAudit struct {
	FiniteMatrixTraceAvailable         bool
	SealedScalarBundleDimension        int
	GaugeFluctuationSquareTraceNonzero bool
	A4LocalAlgebraicIngredientPresent  bool
	DiracOperatorDerived               bool
	OrderOneAxiomVerified              bool
	DixmierTraceDerived                bool
	ContinuumVolumeFormDerived         bool
	SpectralCutoffDerived              bool
	SpectralActionEvaluated            bool
	PhysicalYangMillsActionEvaluated   bool
	HeatKernelCoefficientPromoted      bool
	Verdict                            string
}

type TopologicalCouplingFirewallAudit struct {
	UsesSpontaneousOrientationSeal    bool
	ImportsTopologicalSeal8PiSquared  bool
	EquatesFiniteTraceWithInstanton   bool
	ChernWeilCarrierPreflight         bool
	CompleteChernWeilCarrierDerived   bool
	HeatKernelPreflightPassed         bool
	HeatKernelMatchingDerived         bool
	ThresholdBetaRowsDerived          bool
	AbsoluteCouplingPromoted          bool
	PhysicalConstantsDerived          bool
	ScalarKineticNormalizationDerived bool
	PhysicalGaugeCouplingsDerived     bool
	PhysicalMassesDerived             bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalNullityBefore          int
	ConditionalNullityAfter           int
	ClosedStatements                  []string
	OpenRequirements                  []string
	RecommendedNextGate               string
	Verdict                           string
}

type Summary struct {
	TestsAudited                      int
	InheritedSealedBundle             bool
	FiniteCurvatureTracesStable       bool
	EtaGradingValid                   bool
	NontrivialSignedNeutralCarrier    bool
	HeatKernelPreflightPassed         bool
	ChernWeilOnlyPreflight            bool
	CouplingsAndThresholdsStillSealed bool
	Comment                           string
}

type Analysis struct {
	Seal           scalarorientationseal.Analysis
	Curvature      FiniteCurvatureTraceAudit
	Grading        OrientationGradingAudit
	HeatKernel     HeatKernelPreflightAudit
	Firewall       TopologicalCouplingFirewallAudit
	Summary        Summary
	TruthStatement string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		seal, err := scalarorientationseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 191 seal input: %w", err)
			return
		}
		defaultA, defaultErr = Build(seal, 1e-9)
	})
	return defaultA, defaultErr
}

func Build(seal scalarorientationseal.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !seal.Firewall.ConditionalPhysicalBundleDerived || !seal.Trivialization.ProjectorIntertwiningVerified || !seal.Seal.Quarantined {
		return Analysis{}, fmt.Errorf("Gate 192 requires the Gate 191 sealed H_Phi trivialization")
	}
	if seal.Firewall.ChernWeilCarrierDerived || seal.Firewall.HeatKernelMatchingDerived || seal.Firewall.AbsoluteCouplingPromoted || seal.Firewall.PhysicalConstantsDerived {
		return Analysis{}, fmt.Errorf("Gate 192 refuses a pre-leaked Chern-Weil/heat-kernel/coupling claim")
	}

	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	eta, _ := ph.Sub(pl)

	sc := seal.ScalarCovariant
	z, err := sc.T3.Sub(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	q, err := sc.T3.Add(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}

	records := []CurvatureTraceRecord{
		traceRecord("T1", "primitive SU(2)_L", sc.T1, ph, pl, eta, "1", "1/2", "1/2", "0", "charged weak generator; off-diagonal on the sealed A/B fibers", eps),
		traceRecord("T2", "primitive SU(2)_L", sc.T2, ph, pl, eta, "1", "1/2", "1/2", "0", "charged weak generator; off-diagonal on the sealed A/B fibers", eps),
		traceRecord("T3L", "primitive SU(2)_L", sc.T3, ph, pl, eta, "1", "1/2", "1/2", "0", "diagonal weak-isospin generator preserving the sealed fibers", eps),
		traceRecord("Y_phi", "primitive U(1)_Y", sc.YPhi, ph, pl, eta, "1", "1/2", "1/2", "0", "scalar hypercharge generator preserving the sealed fibers", eps),
		traceRecord("Q=T3+Y_phi", "neutral split", q, ph, pl, eta, "2", "2", "0", "2", "electromagnetic neutral split in the sealed scalar frame", eps),
		traceRecord("Z=T3-Y_phi", "neutral split", z, ph, pl, eta, "2", "0", "2", "-2", "orthogonal neutral broken split in the sealed scalar frame", eps),
	}

	curv := auditCurvature(seal, records, eps)
	grading := auditGrading(seal, eta, sc.T1, sc.T2, sc.T3, sc.YPhi, q, z, eps)
	heat := auditHeatKernel(seal, curv, grading)
	fw := auditFirewall(seal, curv, grading, heat)
	summary := Summary{
		TestsAudited:                      4,
		InheritedSealedBundle:             seal.Firewall.ConditionalPhysicalBundleDerived && seal.Seal.Quarantined,
		FiniteCurvatureTracesStable:       curv.PrimitiveGaugeKineticTracesStable && curv.NeutralQTraceStable && curv.NeutralZTraceStable,
		EtaGradingValid:                   grading.EtaDerivedFromSeal && grading.EtaSquaredResidual <= eps && math.Abs(grading.EtaTrace) <= eps,
		NontrivialSignedNeutralCarrier:    grading.NontrivialSignedNeutralCarrier,
		HeatKernelPreflightPassed:         heat.A4LocalAlgebraicIngredientPresent && !heat.SpectralActionEvaluated,
		ChernWeilOnlyPreflight:            fw.ChernWeilCarrierPreflight && !fw.CompleteChernWeilCarrierDerived,
		CouplingsAndThresholdsStillSealed: !fw.AbsoluteCouplingPromoted && !fw.ThresholdBetaRowsDerived && !fw.PhysicalConstantsDerived,
		Comment:                           "Gate 192 finds exact finite scalar-bundle trace data: primitive gauge kinetic traces are stable, simple eta-graded primitive squares vanish, and the sealed neutral split carries a nontrivial signed trace. This is a Chern-Weil/heat-kernel preflight, not a continuum normalization or coupling theorem.",
	}
	truth := "Gate 192 shows that the sealed H_Phi scalar bundle can carry finite matrix-trace curvature data. The positive kinetic traces Tr(T_a^T T_a) are exact rational diagnostics on the high/low fibers; the seal grading eta is a valid involution; the primitive eta-graded diagonal traces vanish; and the neutral split Q=T3+Y_phi, Z=T3-Y_phi produces a nontrivial signed carrier. This opens a local Chern-Weil/heat-kernel preflight, but it does not derive a boundaryless four-cycle, continuum integration, spectral action, thresholds, absolute couplings, or physical constants."

	return Analysis{Seal: seal, Curvature: curv, Grading: grading, HeatKernel: heat, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func auditCurvature(seal scalarorientationseal.Analysis, records []CurvatureTraceRecord, eps float64) FiniteCurvatureTraceAudit {
	by := func(name string) CurvatureTraceRecord {
		for _, r := range records {
			if r.Name == name {
				return r
			}
		}
		return CurvatureTraceRecord{}
	}
	t1, t2, t3, y := by("T1"), by("T2"), by("T3L"), by("Y_phi")
	q, z := by("Q=T3+Y_phi"), by("Z=T3-Y_phi")
	primitive := t1.StableRational && t2.StableRational && t3.StableRational && y.StableRational
	return FiniteCurvatureTraceAudit{
		SealStatus:                        seal.Seal.ConditionalStatus,
		TraceFunctional:                   "finite bundle matrix trace Tr_HPhi(K), with K=T^T T for skew real gauge generators",
		GeneratorsAudited:                 records,
		PrimitiveGaugeKineticTracesStable: primitive,
		T1TraceStable:                     t1.StableRational,
		T2TraceStable:                     t2.StableRational,
		T3TraceStable:                     t3.StableRational,
		YPhiTraceStable:                   y.StableRational,
		NeutralQTraceStable:               q.StableRational,
		NeutralZTraceStable:               z.StableRational,
		HighLowFiberTraceAvailable:        primitive && approx(t1.HighFiberTrace, 0.5, eps) && approx(t1.LowFiberTrace, 0.5, eps),
		GaugeKineticCarrierPreflight:      primitive && q.StableRational && z.StableRational,
		PhysicalGaugeCouplingsDerived:     false,
		Verdict:                           "The sealed scalar bundle supports exact finite positive trace diagnostics for the primitive weak/hypercharge generators and for the neutral Q/Z split. These are representation traces only; no physical gauge coupling is derived.",
	}
}

func auditGrading(seal scalarorientationseal.Analysis, eta, t1, t2, t3, y, q, z linear.Matrix, eps float64) OrientationGradingAudit {
	eta2, _ := eta.Mul(eta)
	etaRes, _ := eta2.MaxAbsDiff(linear.Identity(4))
	etaTrace, _ := eta.Trace()
	gt1 := gradedTrace(eta, positiveSquare(t1))
	gt2 := gradedTrace(eta, positiveSquare(t2))
	gt3 := gradedTrace(eta, positiveSquare(t3))
	gy := gradedTrace(eta, positiveSquare(y))
	gq := gradedTrace(eta, positiveSquare(q))
	gz := gradedTrace(eta, positiveSquare(z))
	mixed := gradedTrace(eta, positiveProduct(t3, y))
	primitiveZero := approx(gt1, 0, eps) && approx(gt2, 0, eps) && approx(gt3, 0, eps) && approx(gy, 0, eps)
	return OrientationGradingAudit{
		EtaMatrix:                          formatMatrix(eta),
		EtaSquaredResidual:                 etaRes,
		EtaTrace:                           etaTrace,
		EtaHighEigenvalue:                  1,
		EtaLowEigenvalue:                   -1,
		EtaDerivedFromSeal:                 seal.Seal.ExplicitAxiom && seal.Seal.Quarantined,
		EtaIsPhysicalBoundaryData:          seal.Seal.BreaksEtaInvolutionAsBoundaryData,
		PrimitiveDiagonalGradedTracesZero:  primitiveZero,
		ChargedGradedTracesZero:            approx(gt1, 0, eps) && approx(gt2, 0, eps),
		NeutralMixedPairName:               "Tr_eta(T3L^T Y_phi)",
		NeutralMixedGradedTrace:            mixed,
		NeutralMixedExpectedRational:       "1",
		NeutralSplitQGradedTrace:           gq,
		NeutralSplitZGradedTrace:           gz,
		NontrivialSignedNeutralCarrier:     approx(mixed, 1, eps) && approx(gq, 2, eps) && approx(gz, -2, eps),
		IntegerTopologicalChargeMapDerived: false,
		ContinuumOrientationDerived:        false,
		Verdict:                            "The seal supplies a valid eta grading on H_Phi. The primitive graded square traces vanish, so no instanton number is produced directly. The neutral Q/Z split and the T3-Y mixed pairing carry a nontrivial signed finite trace, which is only a local algebraic carrier until an integration/fundamental-class map is derived.",
	}
}

func auditHeatKernel(seal scalarorientationseal.Analysis, curv FiniteCurvatureTraceAudit, grading OrientationGradingAudit) HeatKernelPreflightAudit {
	nonzero := curv.GaugeKineticCarrierPreflight && grading.NontrivialSignedNeutralCarrier
	return HeatKernelPreflightAudit{
		FiniteMatrixTraceAvailable:         true,
		SealedScalarBundleDimension:        seal.ScalarCovariant.ActiveRealDimension,
		GaugeFluctuationSquareTraceNonzero: nonzero,
		A4LocalAlgebraicIngredientPresent:  nonzero && seal.Firewall.ConditionalPhysicalBundleDerived,
		DiracOperatorDerived:               false,
		OrderOneAxiomVerified:              false,
		DixmierTraceDerived:                false,
		ContinuumVolumeFormDerived:         false,
		SpectralCutoffDerived:              false,
		SpectralActionEvaluated:            false,
		PhysicalYangMillsActionEvaluated:   false,
		HeatKernelCoefficientPromoted:      false,
		Verdict:                            "The sealed scalar bundle has the finite trace and nonzero fluctuation-square data needed for an a4-style local preflight. A full heat-kernel coefficient is not promoted because no canonical finite Dirac operator, order-one spectral triple, Dixmier/continuum trace, or cutoff scale is derived here.",
	}
}

func auditFirewall(seal scalarorientationseal.Analysis, curv FiniteCurvatureTraceAudit, grading OrientationGradingAudit, heat HeatKernelPreflightAudit) TopologicalCouplingFirewallAudit {
	return TopologicalCouplingFirewallAudit{
		UsesSpontaneousOrientationSeal:    seal.Seal.ExplicitAxiom,
		ImportsTopologicalSeal8PiSquared:  false,
		EquatesFiniteTraceWithInstanton:   false,
		ChernWeilCarrierPreflight:         curv.GaugeKineticCarrierPreflight && grading.NontrivialSignedNeutralCarrier,
		CompleteChernWeilCarrierDerived:   false,
		HeatKernelPreflightPassed:         heat.A4LocalAlgebraicIngredientPresent,
		HeatKernelMatchingDerived:         false,
		ThresholdBetaRowsDerived:          false,
		AbsoluteCouplingPromoted:          false,
		PhysicalConstantsDerived:          false,
		ScalarKineticNormalizationDerived: seal.Firewall.ScalarKineticNormalizationDerived,
		PhysicalGaugeCouplingsDerived:     false,
		PhysicalMassesDerived:             false,
		StrictNullityBefore:               seal.Firewall.StrictNullityAfter,
		StrictNullityAfter:                seal.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:          seal.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:           seal.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"finite positive curvature traces exist on the sealed scalar bundle",
			"eta is a valid seal-grading involution with trace zero and eta^2=1",
			"primitive eta-graded square traces vanish, preventing a fake instanton-number claim",
			"the neutral Q/Z split supplies a nontrivial signed finite trace carrier",
			"heat-kernel a4 support is only local/preflight because no Dirac/order-one/Dixmier/volume bridge is derived",
		},
		OpenRequirements: []string{
			"derive a boundaryless finite four-cycle or Hochschild/fundamental class before claiming a complete Chern-Weil carrier",
			"derive the finite-to-continuum integration and trace normalization bridge before using S_top=8π² as an absolute normalization",
			"derive a canonical finite Dirac operator/order-one spectral triple before promoting an a4 heat-kernel coefficient",
			"derive scalar kinetic normalization, physical gauge couplings, thresholds, and evaluation scale independently",
		},
		RecommendedNextGate: "Gate 193 — finite fundamental-class / scalar-bundle integration functional search audit",
		Verdict:             "Gate 192 opens a sealed scalar-bundle Chern-Weil/heat-kernel preflight, but keeps the topological-coupling firewall intact: no 8π² import, no instanton identification, no thresholds, no absolute couplings, and no physical constants.",
	}
}

func traceRecord(name, basis string, g, ph, pl, eta linear.Matrix, expectedTotal, expectedHigh, expectedLow, expectedEta, meaning string, eps float64) CurvatureTraceRecord {
	k := positiveSquare(g)
	total, _ := k.Trace()
	high := projectedTrace(ph, k)
	low := projectedTrace(pl, k)
	etaTrace := gradedTrace(eta, k)
	stable := approxToExpected(expectedTotal, total, eps) && approxToExpected(expectedHigh, high, eps) && approxToExpected(expectedLow, low, eps) && approxToExpected(expectedEta, etaTrace, eps)
	return CurvatureTraceRecord{
		Name:                  name,
		Basis:                 basis,
		GeneratorMatrix:       formatMatrix(g),
		KineticOperator:       formatMatrix(k),
		TotalTrace:            total,
		HighFiberTrace:        high,
		LowFiberTrace:         low,
		EtaGradedTrace:        etaTrace,
		ExpectedTotalRational: expectedTotal,
		ExpectedHighRational:  expectedHigh,
		ExpectedLowRational:   expectedLow,
		ExpectedEtaRational:   expectedEta,
		StableRational:        stable,
		PreservesFibers:       offDiagonalNorm(g, ph, pl) <= eps,
		MixesFibers:           offDiagonalNorm(g, ph, pl) > eps,
		PhysicalMeaning:       meaning,
	}
}

func positiveSquare(g linear.Matrix) linear.Matrix {
	return linear.MustMul(g.Transpose(), g)
}

func positiveProduct(a, b linear.Matrix) linear.Matrix {
	return linear.MustMul(a.Transpose(), b)
}

func projectedTrace(p, k linear.Matrix) float64 {
	pkp := linear.MustMul(linear.MustMul(p, k), p)
	tr, _ := pkp.Trace()
	return tr
}

func gradedTrace(eta, k linear.Matrix) float64 {
	tr, _ := linear.MustMul(eta, k).Trace()
	return tr
}

func offDiagonalNorm(g, ph, pl linear.Matrix) float64 {
	hl := linear.MustMul(linear.MustMul(ph, g), pl)
	lh := linear.MustMul(linear.MustMul(pl, g), ph)
	off, _ := hl.Add(lh)
	return off.FrobeniusNorm()
}

func approxToExpected(expected string, got, eps float64) bool {
	switch expected {
	case "0":
		return approx(got, 0, eps)
	case "1/2":
		return approx(got, 0.5, eps)
	case "1":
		return approx(got, 1, eps)
	case "2":
		return approx(got, 2, eps)
	case "-2":
		return approx(got, -2, eps)
	default:
		return false
	}
}

func approx(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func formatMatrix(m linear.Matrix) string {
	rows := make([]string, 0, m.Rows())
	for r := 0; r < m.Rows(); r++ {
		vals := make([]string, 0, m.Cols())
		for c := 0; c < m.Cols(); c++ {
			vals = append(vals, fmt.Sprintf("%.6g", m.At(r, c)))
		}
		rows = append(rows, "["+strings.Join(vals, ", ")+"]")
	}
	return "[" + strings.Join(rows, ", ") + "]"
}

func FormatCurvature(a FiniteCurvatureTraceAudit) string {
	parts := make([]string, 0, len(a.GeneratorsAudited))
	for _, r := range a.GeneratorsAudited {
		parts = append(parts, fmt.Sprintf("%s basis=%s Tr=%.6g high=%.6g low=%.6g eta=%.6g expected=(%s,%s,%s,%s) stable=%t preserve=%t mix=%t meaning=%q", r.Name, r.Basis, r.TotalTrace, r.HighFiberTrace, r.LowFiberTrace, r.EtaGradedTrace, r.ExpectedTotalRational, r.ExpectedHighRational, r.ExpectedLowRational, r.ExpectedEtaRational, r.StableRational, r.PreservesFibers, r.MixesFibers, r.PhysicalMeaning))
	}
	return fmt.Sprintf("seal=%s trace=%q primitive=%t T1=%t T2=%t T3=%t Y=%t Q=%t Z=%t fibers=%t preflight=%t couplings=%t records=[%s] verdict=%s", a.SealStatus, a.TraceFunctional, a.PrimitiveGaugeKineticTracesStable, a.T1TraceStable, a.T2TraceStable, a.T3TraceStable, a.YPhiTraceStable, a.NeutralQTraceStable, a.NeutralZTraceStable, a.HighLowFiberTraceAvailable, a.GaugeKineticCarrierPreflight, a.PhysicalGaugeCouplingsDerived, strings.Join(parts, "; "), a.Verdict)
}

func FormatGrading(a OrientationGradingAudit) string {
	return fmt.Sprintf("eta=%s eta2=%.3g trace=%.3g high=%.3g low=%.3g fromSeal=%t boundary=%t primitiveZero=%t chargedZero=%t mixed=%s:%.6g expected=%s Qeta=%.6g Zeta=%.6g signed=%t integerMap=%t continuum=%t verdict=%s", a.EtaMatrix, a.EtaSquaredResidual, a.EtaTrace, a.EtaHighEigenvalue, a.EtaLowEigenvalue, a.EtaDerivedFromSeal, a.EtaIsPhysicalBoundaryData, a.PrimitiveDiagonalGradedTracesZero, a.ChargedGradedTracesZero, a.NeutralMixedPairName, a.NeutralMixedGradedTrace, a.NeutralMixedExpectedRational, a.NeutralSplitQGradedTrace, a.NeutralSplitZGradedTrace, a.NontrivialSignedNeutralCarrier, a.IntegerTopologicalChargeMapDerived, a.ContinuumOrientationDerived, a.Verdict)
}

func FormatHeatKernel(a HeatKernelPreflightAudit) string {
	return fmt.Sprintf("matrixTrace=%t dim=%d fluctuationTrace=%t a4Local=%t dirac=%t orderOne=%t dixmier=%t volume=%t cutoff=%t action=%t YM=%t promoted=%t verdict=%s", a.FiniteMatrixTraceAvailable, a.SealedScalarBundleDimension, a.GaugeFluctuationSquareTraceNonzero, a.A4LocalAlgebraicIngredientPresent, a.DiracOperatorDerived, a.OrderOneAxiomVerified, a.DixmierTraceDerived, a.ContinuumVolumeFormDerived, a.SpectralCutoffDerived, a.SpectralActionEvaluated, a.PhysicalYangMillsActionEvaluated, a.HeatKernelCoefficientPromoted, a.Verdict)
}

func FormatFirewall(a TopologicalCouplingFirewallAudit) string {
	return fmt.Sprintf("seal=%t import8pi2=%t instantonEq=%t chernPre=%t chernFull=%t heatPre=%t heatMatch=%t thresholds=%t abs=%t constants=%t scalarKinetic=%t gaugeCouplings=%t masses=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.UsesSpontaneousOrientationSeal, a.ImportsTopologicalSeal8PiSquared, a.EquatesFiniteTraceWithInstanton, a.ChernWeilCarrierPreflight, a.CompleteChernWeilCarrierDerived, a.HeatKernelPreflightPassed, a.HeatKernelMatchingDerived, a.ThresholdBetaRowsDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.ScalarKineticNormalizationDerived, a.PhysicalGaugeCouplingsDerived, a.PhysicalMassesDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d sealed=%t traces=%t eta=%t signed=%t heat=%t chernPreOnly=%t firewall=%t comment=%s", a.TestsAudited, a.InheritedSealedBundle, a.FiniteCurvatureTracesStable, a.EtaGradingValid, a.NontrivialSignedNeutralCarrier, a.HeatKernelPreflightPassed, a.ChernWeilOnlyPreflight, a.CouplingsAndThresholdsStillSealed, a.Comment)
}
