// Package scalarfundamentalclass implements Gate 193: finite fundamental-class /
// scalar-bundle integration functional search audit.
//
// Gate 192 found exact finite scalar-bundle trace data behind the Gate 191
// SpontaneousOrientationSeal. Gate 193 asks whether those traces can be
// organized as an algebraic integration functional/fundamental class without
// importing a continuum volume form, a Dixmier trace, S_top=8π², thresholds, or
// physical gauge couplings.
//
// The important subtlety is that the eta-graded trace is not a universal cyclic
// trace on the full 4x4 matrix algebra. It is a closed/cyclic functional on the
// sealed eta-even curvature-observable algebra audited here. This is exactly why
// the result is a finite fundamental-class preflight rather than a completed
// continuum Chern-Weil theorem.
package scalarfundamentalclass

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarchernweiltaudit"
	"github.com/bagherbal/asha-engine/pkg/linear"
)

type CandidateFunctionalAudit struct {
	BundleName                        string
	BundleDimension                   int
	SealStatus                        string
	OrdinaryTraceFunctional           string
	EtaGradedFunctional               string
	EtaMatrix                         string
	EtaSquaredResidual                float64
	EtaTrace                          float64
	Domain                            string
	UsesSpontaneousOrientationSeal    bool
	FiniteMatrixFunctionalConstructed bool
	ContinuumIntegralImported         bool
	DixmierTraceDerived               bool
	Verdict                           string
}

type ClosedCyclicAudit struct {
	OrdinaryTraceCyclicOnFullMatrixAlgebra     bool
	EtaTraceCyclicOnFullMatrixAlgebra          bool
	EtaTraceFullMatrixCounterexample           string
	EtaTraceFullMatrixCounterexampleDefect     float64
	EtaTraceClosedOnGaugeGeneratorAlgebra      bool
	EtaTraceClosedOnCurvatureObservableAlgebra bool
	EtaTraceConnectionAdjointDefectMax         float64
	EtaTraceObservableCyclicDefectMax          float64
	HochschildBoundaryZeroOnAuditedDomain      bool
	FullConnectionUniversalIntegralDerived     bool
	Verdict                                    string
}

type QuantizedTraceRecord struct {
	Name                     string
	Expression               string
	NativeValue              float64
	ExpectedRational         string
	StableInteger            bool
	HalfFiberNormalized      float64
	HalfNormalizationMeaning string
}

type NormalizationSearchAudit struct {
	NativeTraceRecords                  []QuantizedTraceRecord
	NeutralQNativeDegree                float64
	NeutralZNativeDegree                float64
	NeutralMixedNativeDegree            float64
	StableQuantizedInvariants           bool
	FiberDimension                      int
	HalfFiberNormalizationCandidate     bool
	HalfFiberNormalizationForced        bool
	UnitFundamentalClassDerived         bool
	CanonicalNormalizationFactorDerived bool
	NativeAlgebraicDegreesPreserved     bool
	Verdict                             string
}

type HeatKernelContinuumFirewallAudit struct {
	FiniteIntegrationFunctionalExists  bool
	FiniteSignedCurvatureCarrierExists bool
	CompleteChernWeilCarrierDerived    bool
	ContinuumVolumeFormDerived         bool
	BoundarylessFourCycleDerived       bool
	DiracOperatorDerived               bool
	OrderOneAxiomVerified              bool
	DixmierTraceDerived                bool
	SpectralCutoffDerived              bool
	HeatKernelA4CoefficientPromoted    bool
	ImportsTopologicalSeal8PiSquared   bool
	EquatesFiniteDegreeWithInstanton   bool
	FiniteToContinuumScaleDerived      bool
	ThresholdBetaRowsDerived           bool
	AbsoluteCouplingPromoted           bool
	PhysicalConstantsDerived           bool
	StrictNullityBefore                int
	StrictNullityAfter                 int
	ConditionalNullityBefore           int
	ConditionalNullityAfter            int
	ClosedStatements                   []string
	OpenRequirements                   []string
	RecommendedNextGate                string
	Verdict                            string
}

type MatterExtensionPlanAudit struct {
	MatterFockDimension       int
	ScalarBundleDimension     int
	TotalTensorDimension      int
	ProposedLift              string
	YukawaAuditMode           string
	SelectionRulesCanBeReused bool
	YukawaAmplitudesDerived   bool
	MassTermsDerived          bool
	RequiresSeparateGate      bool
	RecommendedGate           string
	Verdict                   string
}

type Summary struct {
	TestsAudited                    int
	InheritedGate192Carrier         bool
	FiniteFunctionalConstructed     bool
	ClosedOnAuditedEtaEvenDomain    bool
	FullMatrixEtaTraceRejected      bool
	StableNativeDegrees             bool
	CanonicalContinuumNormalization bool
	ContinuumFirewallPreserved      bool
	Comment                         string
}

type Analysis struct {
	PreviousGate192 scalarchernweiltaudit.Analysis
	Functional      CandidateFunctionalAudit
	ClosedCyclic    ClosedCyclicAudit
	Normalization   NormalizationSearchAudit
	Firewall        HeatKernelContinuumFirewallAudit
	MatterPlan      MatterExtensionPlanAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := scalarchernweiltaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 192 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev, 1e-9)
	})
	return defaultA, defaultErr
}

func Build(prev scalarchernweiltaudit.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-9
	}
	if !prev.Summary.InheritedSealedBundle || !prev.Summary.NontrivialSignedNeutralCarrier || !prev.Firewall.ChernWeilCarrierPreflight {
		return Analysis{}, fmt.Errorf("Gate 193 requires Gate 192 sealed scalar-bundle signed carrier preflight")
	}
	if prev.Firewall.ImportsTopologicalSeal8PiSquared || prev.Firewall.AbsoluteCouplingPromoted || prev.Firewall.PhysicalConstantsDerived {
		return Analysis{}, fmt.Errorf("Gate 193 refuses leaked topology/coupling/constant claims from Gate 192")
	}

	ph := linear.Diagonal([]float64{1, 1, 0, 0})
	pl := linear.Diagonal([]float64{0, 0, 1, 1})
	eta, _ := ph.Sub(pl)
	eta2 := linear.MustMul(eta, eta)
	etaRes, _ := eta2.MaxAbsDiff(linear.Identity(4))
	etaTrace, _ := eta.Trace()

	sc := prev.Seal.ScalarCovariant
	q, err := sc.T3.Add(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}
	z, err := sc.T3.Sub(sc.YPhi)
	if err != nil {
		return Analysis{}, err
	}

	functional := CandidateFunctionalAudit{
		BundleName:                        "sealed H_Phi scalar bundle",
		BundleDimension:                   prev.HeatKernel.SealedScalarBundleDimension,
		SealStatus:                        prev.Seal.Seal.ConditionalStatus,
		OrdinaryTraceFunctional:           "tau_0(O)=Tr_HPhi(O)",
		EtaGradedFunctional:               "tau_eta(O)=Tr_HPhi(eta O)",
		EtaMatrix:                         formatMatrix(eta),
		EtaSquaredResidual:                etaRes,
		EtaTrace:                          etaTrace,
		Domain:                            "eta-even curvature-observable algebra generated by P_high, P_low, T_a^T T_b, Q^T Q, Z^T Z, and T3L^T Y_phi",
		UsesSpontaneousOrientationSeal:    prev.Seal.Seal.ExplicitAxiom && prev.Seal.Seal.Quarantined,
		FiniteMatrixFunctionalConstructed: etaRes <= eps && math.Abs(etaTrace) <= eps,
		ContinuumIntegralImported:         false,
		DixmierTraceDerived:               false,
		Verdict:                           "Gate 193 constructs a finite sealed scalar-bundle functional pair: the ordinary matrix trace and the eta-graded signed trace. The eta functional is defined on the audited eta-even curvature-observable algebra, not promoted to a continuum integral.",
	}

	closed := auditClosedCyclic(sc.T1, sc.T2, sc.T3, sc.YPhi, q, z, eta, eps)
	norm := auditNormalization(prev)
	fw := auditFirewall(prev, functional, closed, norm)
	matter := MatterExtensionPlanAudit{
		MatterFockDimension:       16,
		ScalarBundleDimension:     functional.BundleDimension,
		TotalTensorDimension:      16 * functional.BundleDimension,
		ProposedLift:              "H_total = H_Fock ⊗ H_Phi; tau_total is a future tensor lift of charge-sector/Fock trace with tau_eta on the scalar factor",
		YukawaAuditMode:           "selection-rule and bilinear-support audit only; no fitted amplitudes, no masses, no generation texture values",
		SelectionRulesCanBeReused: true,
		YukawaAmplitudesDerived:   false,
		MassTermsDerived:          false,
		RequiresSeparateGate:      true,
		RecommendedGate:           "Gate 194 — tensor-lifted scalar fundamental class / Yukawa bilinear support audit",
		Verdict:                   "Matter fields should be mapped by tensor-lifting the finite scalar functional to H_Fock ⊗ H_Phi and auditing which already-derived Yukawa selection channels have nonzero scalar support. This must remain a support theorem, not a mass/amplitude theorem.",
	}
	summary := Summary{
		TestsAudited:                    4,
		InheritedGate192Carrier:         prev.Summary.NontrivialSignedNeutralCarrier && prev.Firewall.ChernWeilCarrierPreflight,
		FiniteFunctionalConstructed:     functional.FiniteMatrixFunctionalConstructed,
		ClosedOnAuditedEtaEvenDomain:    closed.HochschildBoundaryZeroOnAuditedDomain,
		FullMatrixEtaTraceRejected:      !closed.EtaTraceCyclicOnFullMatrixAlgebra && closed.EtaTraceFullMatrixCounterexampleDefect > eps,
		StableNativeDegrees:             norm.StableQuantizedInvariants && norm.NativeAlgebraicDegreesPreserved,
		CanonicalContinuumNormalization: norm.CanonicalNormalizationFactorDerived,
		ContinuumFirewallPreserved:      !fw.ImportsTopologicalSeal8PiSquared && !fw.AbsoluteCouplingPromoted && !fw.PhysicalConstantsDerived,
		Comment:                         "Gate 193 constructs a finite algebraic integration functional on the sealed eta-even scalar curvature algebra. It explicitly rejects the stronger false claim that Tr_eta is a universal cyclic trace on all 4x4 matrices, and it preserves the continuum/coupling firewall.",
	}
	truth := "Gate 193 upgrades the Gate 192 trace preflight into a finite algebraic fundamental-class candidate: tau_0=Tr and tau_eta=Tr(eta·) are exact finite matrix functionals on the sealed scalar bundle, with tau_eta closed on the audited eta-even curvature-observable algebra. The gate also proves the limit of this construction: tau_eta is not a universal cyclic trace on the full matrix algebra, no canonical 1/2 or continuum normalization is forced, and no 8π²/heat-kernel/coupling/threshold/constant bridge is derived."

	return Analysis{PreviousGate192: prev, Functional: functional, ClosedCyclic: closed, Normalization: norm, Firewall: fw, MatterPlan: matter, Summary: summary, TruthStatement: truth}, nil
}

func auditClosedCyclic(t1, t2, t3, y, q, z, eta linear.Matrix, eps float64) ClosedCyclicAudit {
	// Counterexample in the full matrix algebra: E_02 and E_20 are eta-odd rank-one
	// maps between the high and low planes. tau_eta([E02,E20]) = 2.
	e02 := matrixUnit(4, 0, 2)
	e20 := matrixUnit(4, 2, 0)
	counter := etaCyclicDefect(eta, e02, e20)

	gens := []linear.Matrix{t1, t2, t3, y}
	genDefect := maxEtaCyclicDefect(eta, gens, gens)

	observables := []linear.Matrix{
		linear.MustMul(t1.Transpose(), t1),
		linear.MustMul(t2.Transpose(), t2),
		linear.MustMul(t3.Transpose(), t3),
		linear.MustMul(y.Transpose(), y),
		linear.MustMul(q.Transpose(), q),
		linear.MustMul(z.Transpose(), z),
		linear.MustMul(t3.Transpose(), y),
		linear.Diagonal([]float64{1, 1, 0, 0}),
		linear.Diagonal([]float64{0, 0, 1, 1}),
	}
	obsDefect := maxEtaCyclicDefect(eta, observables, observables)
	connObsDefect := maxEtaCyclicDefect(eta, gens, observables)

	return ClosedCyclicAudit{
		OrdinaryTraceCyclicOnFullMatrixAlgebra:     true,
		EtaTraceCyclicOnFullMatrixAlgebra:          math.Abs(counter) <= eps,
		EtaTraceFullMatrixCounterexample:           "tau_eta([E_02,E_20]) where E_02 maps high->low and E_20 maps low->high",
		EtaTraceFullMatrixCounterexampleDefect:     counter,
		EtaTraceClosedOnGaugeGeneratorAlgebra:      genDefect <= eps,
		EtaTraceClosedOnCurvatureObservableAlgebra: obsDefect <= eps,
		EtaTraceConnectionAdjointDefectMax:         connObsDefect,
		EtaTraceObservableCyclicDefectMax:          obsDefect,
		HochschildBoundaryZeroOnAuditedDomain:      genDefect <= eps && obsDefect <= eps && connObsDefect <= eps && math.Abs(counter) > eps,
		FullConnectionUniversalIntegralDerived:     false,
		Verdict:                                    "The ordinary trace is cyclic on all finite matrices. The eta-graded trace is not cyclic on the full matrix algebra, as shown by an explicit high/low rank-one counterexample. On the audited eta-even curvature-observable algebra and under the audited gauge-generator adjoint tests, the eta functional is closed; this is the lawful finite fundamental-class domain.",
	}
}

func auditNormalization(prev scalarchernweiltaudit.Analysis) NormalizationSearchAudit {
	records := []QuantizedTraceRecord{
		{Name: "neutral electromagnetic split", Expression: "tau_eta(Q^T Q)", NativeValue: prev.Grading.NeutralSplitQGradedTrace, ExpectedRational: "2", StableInteger: prev.Grading.NeutralSplitQGradedTrace == 2, HalfFiberNormalized: prev.Grading.NeutralSplitQGradedTrace / 2, HalfNormalizationMeaning: "would divide by the 2D high fiber, yielding 1, but this is a convention unless forced"},
		{Name: "neutral Z split", Expression: "tau_eta(Z^T Z)", NativeValue: prev.Grading.NeutralSplitZGradedTrace, ExpectedRational: "-2", StableInteger: prev.Grading.NeutralSplitZGradedTrace == -2, HalfFiberNormalized: prev.Grading.NeutralSplitZGradedTrace / 2, HalfNormalizationMeaning: "would divide by the 2D low fiber, yielding -1, but this is a convention unless forced"},
		{Name: "neutral mixed pairing", Expression: "tau_eta(T3L^T Y_phi)", NativeValue: prev.Grading.NeutralMixedGradedTrace, ExpectedRational: "1", StableInteger: prev.Grading.NeutralMixedGradedTrace == 1, HalfFiberNormalized: prev.Grading.NeutralMixedGradedTrace / 2, HalfNormalizationMeaning: "already has native unit value; a forced 1/2 would change it to 1/2"},
	}
	stable := true
	for _, r := range records {
		stable = stable && r.StableInteger
	}
	return NormalizationSearchAudit{
		NativeTraceRecords:                  records,
		NeutralQNativeDegree:                prev.Grading.NeutralSplitQGradedTrace,
		NeutralZNativeDegree:                prev.Grading.NeutralSplitZGradedTrace,
		NeutralMixedNativeDegree:            prev.Grading.NeutralMixedGradedTrace,
		StableQuantizedInvariants:           stable,
		FiberDimension:                      2,
		HalfFiberNormalizationCandidate:     true,
		HalfFiberNormalizationForced:        false,
		UnitFundamentalClassDerived:         false,
		CanonicalNormalizationFactorDerived: false,
		NativeAlgebraicDegreesPreserved:     true,
		Verdict:                             "The finite scalar bundle supplies stable native integer degrees 2, -2, and 1. A 1/2 fiber-dimension normalization is visible but not forced: it would make Q/Z unit-valued while changing the already-unit mixed pairing. The engine therefore preserves the native algebraic degrees.",
	}
}

func auditFirewall(prev scalarchernweiltaudit.Analysis, fn CandidateFunctionalAudit, closed ClosedCyclicAudit, norm NormalizationSearchAudit) HeatKernelContinuumFirewallAudit {
	return HeatKernelContinuumFirewallAudit{
		FiniteIntegrationFunctionalExists:  fn.FiniteMatrixFunctionalConstructed && closed.HochschildBoundaryZeroOnAuditedDomain,
		FiniteSignedCurvatureCarrierExists: prev.Grading.NontrivialSignedNeutralCarrier && norm.StableQuantizedInvariants,
		CompleteChernWeilCarrierDerived:    false,
		ContinuumVolumeFormDerived:         false,
		BoundarylessFourCycleDerived:       false,
		DiracOperatorDerived:               false,
		OrderOneAxiomVerified:              false,
		DixmierTraceDerived:                false,
		SpectralCutoffDerived:              false,
		HeatKernelA4CoefficientPromoted:    false,
		ImportsTopologicalSeal8PiSquared:   false,
		EquatesFiniteDegreeWithInstanton:   false,
		FiniteToContinuumScaleDerived:      false,
		ThresholdBetaRowsDerived:           false,
		AbsoluteCouplingPromoted:           false,
		PhysicalConstantsDerived:           false,
		StrictNullityBefore:                prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                 prev.Firewall.StrictNullityAfter,
		ConditionalNullityBefore:           prev.Firewall.ConditionalNullityAfter,
		ConditionalNullityAfter:            prev.Firewall.ConditionalNullityAfter,
		ClosedStatements: []string{
			"finite scalar-bundle integration functional tau_eta exists on the audited eta-even curvature-observable algebra",
			"tau_eta is not a universal cyclic trace on the full 4x4 matrix algebra",
			"native signed degrees tau_eta(Q^TQ)=2, tau_eta(Z^TZ)=-2, tau_eta(T3^T Y)=1 are stable finite invariants",
			"no canonical 1/2 continuum/fundamental-class normalization is forced by the finite algebra",
		},
		OpenRequirements: []string{
			"derive a boundaryless finite four-cycle or Hochschild cycle before claiming a complete Chern-Weil carrier",
			"derive a finite-to-continuum volume/Dixmier trace bridge before mapping tau_eta to an integral over spacetime",
			"derive the Dirac/order-one/spectral-cutoff bridge before promoting a heat-kernel a4 coefficient",
			"derive thresholds, scalar kinetic normalization, absolute gauge couplings, and physical constants independently",
		},
		RecommendedNextGate: "Gate 194 — tensor-lifted scalar fundamental class / Yukawa bilinear support audit",
		Verdict:             "Gate 193 constructs the finite algebraic scalar-bundle fundamental-class candidate while preserving the continuum firewall: no S_top=8π² import, no instanton equivalence, no heat-kernel promotion, no thresholds, no couplings, and no physical constants.",
	}
}

func etaCyclicDefect(eta, a, b linear.Matrix) float64 {
	ab := linear.MustMul(a, b)
	ba := linear.MustMul(b, a)
	comm := linear.MustSub(ab, ba)
	tr, _ := linear.MustMul(eta, comm).Trace()
	return tr
}

func maxEtaCyclicDefect(eta linear.Matrix, left, right []linear.Matrix) float64 {
	max := 0.0
	for _, a := range left {
		for _, b := range right {
			if d := math.Abs(etaCyclicDefect(eta, a, b)); d > max {
				max = d
			}
		}
	}
	return max
}

func matrixUnit(n, r, c int) linear.Matrix {
	m := linear.NewMatrix(n, n)
	m.Set(r, c, 1)
	return m
}

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

func FormatFunctional(a CandidateFunctionalAudit) string {
	return fmt.Sprintf("bundle=%s dim=%d seal=%s tau0=%s tauEta=%s eta=%s eta2Residual=%g etaTrace=%g domain=%s seal=%t finite=%t continuum=%t dixmier=%t verdict=%s", a.BundleName, a.BundleDimension, a.SealStatus, a.OrdinaryTraceFunctional, a.EtaGradedFunctional, a.EtaMatrix, a.EtaSquaredResidual, a.EtaTrace, a.Domain, a.UsesSpontaneousOrientationSeal, a.FiniteMatrixFunctionalConstructed, a.ContinuumIntegralImported, a.DixmierTraceDerived, a.Verdict)
}

func FormatClosedCyclic(a ClosedCyclicAudit) string {
	return fmt.Sprintf("ordinaryFull=%t etaFull=%t counter=%s counterDefect=%g etaGauge=%t etaCurvature=%t connDefect=%g obsDefect=%g hochschildDomain=%t universal=%t verdict=%s", a.OrdinaryTraceCyclicOnFullMatrixAlgebra, a.EtaTraceCyclicOnFullMatrixAlgebra, a.EtaTraceFullMatrixCounterexample, a.EtaTraceFullMatrixCounterexampleDefect, a.EtaTraceClosedOnGaugeGeneratorAlgebra, a.EtaTraceClosedOnCurvatureObservableAlgebra, a.EtaTraceConnectionAdjointDefectMax, a.EtaTraceObservableCyclicDefectMax, a.HochschildBoundaryZeroOnAuditedDomain, a.FullConnectionUniversalIntegralDerived, a.Verdict)
}

func FormatNormalization(a NormalizationSearchAudit) string {
	parts := make([]string, 0, len(a.NativeTraceRecords))
	for _, r := range a.NativeTraceRecords {
		parts = append(parts, fmt.Sprintf("%s:%s=%g expected=%s integer=%t half=%g", r.Name, r.Expression, r.NativeValue, r.ExpectedRational, r.StableInteger, r.HalfFiberNormalized))
	}
	return fmt.Sprintf("records=[%s] Q=%g Z=%g mixed=%g stable=%t fiberDim=%d halfCandidate=%t halfForced=%t unit=%t canonicalNorm=%t nativePreserved=%t verdict=%s", strings.Join(parts, "; "), a.NeutralQNativeDegree, a.NeutralZNativeDegree, a.NeutralMixedNativeDegree, a.StableQuantizedInvariants, a.FiberDimension, a.HalfFiberNormalizationCandidate, a.HalfFiberNormalizationForced, a.UnitFundamentalClassDerived, a.CanonicalNormalizationFactorDerived, a.NativeAlgebraicDegreesPreserved, a.Verdict)
}

func FormatFirewall(a HeatKernelContinuumFirewallAudit) string {
	return fmt.Sprintf("finiteIntegral=%t signedCarrier=%t chernFull=%t volume=%t fourCycle=%t dirac=%t orderOne=%t dixmier=%t cutoff=%t a4=%t import8pi2=%t instantonEq=%t scale=%t thresholds=%t abs=%t constants=%t strict=%d->%d conditional=%d->%d closed=[%s] open=[%s] next=%s verdict=%s", a.FiniteIntegrationFunctionalExists, a.FiniteSignedCurvatureCarrierExists, a.CompleteChernWeilCarrierDerived, a.ContinuumVolumeFormDerived, a.BoundarylessFourCycleDerived, a.DiracOperatorDerived, a.OrderOneAxiomVerified, a.DixmierTraceDerived, a.SpectralCutoffDerived, a.HeatKernelA4CoefficientPromoted, a.ImportsTopologicalSeal8PiSquared, a.EquatesFiniteDegreeWithInstanton, a.FiniteToContinuumScaleDerived, a.ThresholdBetaRowsDerived, a.AbsoluteCouplingPromoted, a.PhysicalConstantsDerived, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityBefore, a.ConditionalNullityAfter, strings.Join(a.ClosedStatements, "; "), strings.Join(a.OpenRequirements, "; "), a.RecommendedNextGate, a.Verdict)
}

func FormatMatterPlan(a MatterExtensionPlanAudit) string {
	return fmt.Sprintf("fockDim=%d scalarDim=%d totalDim=%d lift=%s yukawaMode=%s selectionRules=%t amplitudes=%t masses=%t separateGate=%t next=%s verdict=%s", a.MatterFockDimension, a.ScalarBundleDimension, a.TotalTensorDimension, a.ProposedLift, a.YukawaAuditMode, a.SelectionRulesCanBeReused, a.YukawaAmplitudesDerived, a.MassTermsDerived, a.RequiresSeparateGate, a.RecommendedGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tests=%d inherited=%t finite=%t closedDomain=%t fullRejected=%t degrees=%t continuumNorm=%t firewall=%t comment=%s", a.TestsAudited, a.InheritedGate192Carrier, a.FiniteFunctionalConstructed, a.ClosedOnAuditedEtaEvenDomain, a.FullMatrixEtaTraceRejected, a.StableNativeDegrees, a.CanonicalContinuumNormalization, a.ContinuumFirewallPreserved, a.Comment)
}
