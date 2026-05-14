// Package topologicalboundaryviability implements Gate 200: topological
// boundary viability / bottom-up convergence comparison audit.
//
// Gate 199 built a symbolic threshold-corrected RG expression after explicit
// UV boundary seals M* and u*=1/g_*^2.  Gate 200 deliberately crosses into a
// phenomenological comparison layer: it quarantines observed Z-pole couplings,
// solves the one-loop pairwise UV intersection equations, measures the mismatch
// triangle, and compares the inferred boundary intercept to the optional
// topological branch u*=1.
//
// This package never promotes observed values to finite theorems.  It does not
// evaluate a threshold-corrected physical prediction because numerical threshold
// ordering, W/Z thresholds, finite matching corrections, and continuum scheme
// data remain sealed.
package topologicalboundaryviability

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugecouplingboundaryseal"
)

type BetaVector = gaugecouplingboundaryseal.BetaVector

type EmpiricalComparisonLedger struct {
	Name                          string
	ScaleSymbol                   string
	ScaleGeV                      float64
	AlphaEMInverse                float64
	Sin2ThetaMSbar                float64
	AlphaS                        float64
	Alpha1GUTInverse              float64
	AlphaYInverse                 float64
	Alpha2Inverse                 float64
	Alpha3Inverse                 float64
	Source                        string
	ExplicitPhenomenologicalInput bool
	Quarantined                   bool
	UsedForFiniteDerivation       bool
	UsedForBoundaryDerivation     bool
	Verdict                       string
}

type PairwiseIntersection struct {
	Pair                    string
	GaugeI                  string
	GaugeJ                  string
	LogIntervalL            float64
	ScaleGeV                float64
	Log10ScaleGeV           float64
	InferredAlphaInverse    float64
	InferredU               float64
	FiniteDerived           bool
	UsesObservedLedger      bool
	UsesThresholdLedger     bool
	ClosedFormSolved        bool
	PhysicalPredictionClaim bool
	Verdict                 string
}

type ConvergenceTriangle struct {
	Pairwise                   []PairwiseIntersection
	PairwiseCount              int
	AllClosedFormSolved        bool
	SingleIntersectionFound    bool
	LogSpread                  float64
	ScaleRatioMaxOverMin       float64
	TriangleAreaInLogUPlane    float64
	MismatchNonzero            bool
	UsesObservedLedger         bool
	ThresholdCorrected         bool
	PhysicalUnificationDerived bool
	Verdict                    string
}

type TopologicalBenchmark struct {
	TopologicalBranchAvailableForComparison bool
	TopologicalBranchAssumedAsTruth         bool
	TopologicalU                            float64
	CentroidLogInterval                     float64
	CentroidScaleGeV                        float64
	InferredU1                              float64
	InferredU2                              float64
	InferredU3                              float64
	AverageInferredU                        float64
	DeltaUFromTopologicalUnit               float64
	CloseToUnitAtLooseTolerance             bool
	FiniteDerivationClaim                   bool
	Verdict                                 string
}

type ThresholdComparisonAudit struct {
	Gate199SymbolicThresholdTreeInherited  bool
	ExactThresholdBetaRowsAvailable        bool
	EmpiricalThresholdLedgerSupplied       bool
	ThresholdOrderingKnown                 bool
	NumericalThresholdCorrectedRunAllowed  bool
	NumericalThresholdCorrectedRunExecuted bool
	SharpStepSchemeStillConditional        bool
	FiniteMatchingCorrectionsAvailable     bool
	WZThresholdsAvailable                  bool
	LowEnergyZPoleStrictlyInsideDomain     bool
	FormalThresholdReadyExpression         string
	BlockReason                            string
	Verdict                                string
}

type MatchingAndSolverAudit struct {
	SolverKind                        string
	UsesExactClosedFormPairwiseLogs   bool
	RequiresNumericalOptimization     bool
	NumericalOptimizationUsed         bool
	TreeLevelContinuityEnforced       bool
	FiniteThresholdCorrectionsAudited bool
	FiniteThresholdCorrectionsDerived bool
	ContinuityCondition               string
	FiniteCorrectionTerm              string
	Verdict                           string
}

type FirewallAudit struct {
	Gate199Inherited                       bool
	EmpiricalLedgerQuarantined             bool
	ObservedInputsUsedForFiniteDerivation  bool
	ObservedInputsUsedForPhenomenologyOnly bool
	BoundaryScaleDerived                   bool
	AbsoluteCouplingDerived                bool
	TopologicalUOneDerived                 bool
	TopologicalUOneAssumed                 bool
	EightPiSquaredImported                 bool
	ThresholdCorrectedPhysicalFitClaimed   bool
	PhysicalUnificationClaimed             bool
	PhysicalGaugeCouplingsDerived          bool
	WZThresholdsDerived                    bool
	FiniteMatchingCorrectionsDerived       bool
	FiniteToContinuumNormalizationDerived  bool
	StrictNullityBefore                    int
	StrictNullityAfter                     int
	PhenomenologyNullityBefore             int
	PhenomenologyNullityAfter              int
	PhysicalPredictionNullityBefore        int
	PhysicalPredictionNullityAfter         int
	OpenRequirements                       []string
	RecommendedNextGate                    string
	Verdict                                string
}

type Summary struct {
	TestsAudited                        int
	EmpiricalLedgerQuarantined          bool
	PairwiseIntersectionsSolved         bool
	MismatchTriangleNonzero             bool
	TopologicalBenchmarkComputed        bool
	ThresholdCorrectedEvaluationBlocked bool
	NoFiniteDerivationClaim             bool
	Comment                             string
}

type Analysis struct {
	PreviousGate199 gaugecouplingboundaryseal.Analysis
	Ledger          EmpiricalComparisonLedger
	Triangle        ConvergenceTriangle
	Benchmark       TopologicalBenchmark
	ThresholdAudit  ThresholdComparisonAudit
	SolverAudit     MatchingAndSolverAudit
	Firewall        FirewallAudit
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
		prev, err := gaugecouplingboundaryseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 199 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(prev)
	})
	return defaultA, defaultErr
}

func Build(prev gaugecouplingboundaryseal.Analysis) (Analysis, error) {
	if !prev.Summary.SymbolicTopDownTrajectoryBuilt || !prev.Summary.BottomUpViabilityAuditSeparated {
		return Analysis{}, fmt.Errorf("Gate 200 requires Gate 199 symbolic RG expression and separated bottom-up audit")
	}
	if prev.Firewall.ObservedInputsImported || prev.Firewall.PhysicalRGPredictionMade || prev.Firewall.NumericalTrajectoryEvaluated || prev.Firewall.AbsoluteCouplingDerivedStrict || prev.Firewall.BoundaryScaleDerivedStrict {
		return Analysis{}, fmt.Errorf("Gate 200 refuses inherited observed-input, physical-prediction, numerical-trajectory, or boundary-derivation leakage")
	}

	ledger := buildEmpiricalLedger()
	triangle := buildConvergenceTriangle(ledger)
	benchmark := buildTopologicalBenchmark(ledger, triangle)
	threshold := auditThresholdComparison(prev)
	solver := auditSolverAndMatching(prev)
	fw := auditFirewall(prev, ledger, triangle, benchmark, threshold, solver)
	summary := Summary{
		TestsAudited:                        7,
		EmpiricalLedgerQuarantined:          ledger.Quarantined && ledger.ExplicitPhenomenologicalInput && !ledger.UsedForFiniteDerivation && !ledger.UsedForBoundaryDerivation,
		PairwiseIntersectionsSolved:         triangle.PairwiseCount == 3 && triangle.AllClosedFormSolved,
		MismatchTriangleNonzero:             triangle.MismatchNonzero && !triangle.SingleIntersectionFound,
		TopologicalBenchmarkComputed:        benchmark.TopologicalBranchAvailableForComparison && !benchmark.TopologicalBranchAssumedAsTruth && benchmark.DeltaUFromTopologicalUnit > 0,
		ThresholdCorrectedEvaluationBlocked: !threshold.NumericalThresholdCorrectedRunAllowed && !threshold.NumericalThresholdCorrectedRunExecuted && !threshold.EmpiricalThresholdLedgerSupplied,
		NoFiniteDerivationClaim:             !fw.ObservedInputsUsedForFiniteDerivation && !fw.BoundaryScaleDerived && !fw.AbsoluteCouplingDerived && !fw.PhysicalUnificationClaimed && fw.StrictNullityBefore == fw.StrictNullityAfter,
		Comment:                             "Gate 200 quarantines a Z-pole comparison ledger, solves the closed-form pairwise one-loop UV intersections, finds a nonzero mismatch triangle, and compares the inferred intercept with the optional u*=1 topological branch. Threshold-corrected numerical evaluation remains blocked until an empirical threshold ledger and scheme matching data are supplied.",
	}
	truth := "Gate 200 is a phenomenological diagnostic, not a finite theorem. With the quarantined Z-pole comparison ledger, the three pairwise one-loop UV intersections are computable in closed form, but they do not coincide. The inferred boundary intercept is far from the optional topological u*=1 branch under the default non-thresholded Z-pole comparison. This does not refute the finite algebra; it proves that physical viability requires additional sealed threshold data, W/Z treatment, scheme matching, and possibly new finite/BSM structure."

	return Analysis{
		PreviousGate199: prev,
		Ledger:          ledger,
		Triangle:        triangle,
		Benchmark:       benchmark,
		ThresholdAudit:  threshold,
		SolverAudit:     solver,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  truth,
	}, nil
}

func buildEmpiricalLedger() EmpiricalComparisonLedger {
	const (
		mz       = 91.1876
		alphaInv = 127.955
		sin2     = 0.23122
		alphaS   = 0.1179
	)
	alpha2Inv := sin2 * alphaInv
	alphaYInv := (1 - sin2) * alphaInv
	alpha1Inv := (3.0 / 5.0) * alphaYInv
	alpha3Inv := 1 / alphaS
	return EmpiricalComparisonLedger{
		Name:                          "Z-pole empirical comparison ledger",
		ScaleSymbol:                   "M_Z",
		ScaleGeV:                      mz,
		AlphaEMInverse:                alphaInv,
		Sin2ThetaMSbar:                sin2,
		AlphaS:                        alphaS,
		Alpha1GUTInverse:              alpha1Inv,
		AlphaYInverse:                 alphaYInv,
		Alpha2Inverse:                 alpha2Inv,
		Alpha3Inverse:                 alpha3Inv,
		Source:                        "PDG-style Z-pole electroweak/QCD comparison values; phenomenology ledger only",
		ExplicitPhenomenologicalInput: true,
		Quarantined:                   true,
		UsedForFiniteDerivation:       false,
		UsedForBoundaryDerivation:     false,
		Verdict:                       "observed couplings admitted only as dirty comparison input",
	}
}

func buildConvergenceTriangle(obs EmpiricalComparisonLedger) ConvergenceTriangle {
	const (
		b1 = 41.0 / 10.0
		b2 = -19.0 / 6.0
		b3 = -7.0
	)
	pairs := []PairwiseIntersection{
		buildPair("12", "U1_GUT", "SU2L", obs.Alpha1GUTInverse, obs.Alpha2Inverse, b1, b2, obs),
		buildPair("13", "U1_GUT", "SU3C", obs.Alpha1GUTInverse, obs.Alpha3Inverse, b1, b3, obs),
		buildPair("23", "SU2L", "SU3C", obs.Alpha2Inverse, obs.Alpha3Inverse, b2, b3, obs),
	}
	minL, maxL := pairs[0].LogIntervalL, pairs[0].LogIntervalL
	for _, p := range pairs[1:] {
		if p.LogIntervalL < minL {
			minL = p.LogIntervalL
		}
		if p.LogIntervalL > maxL {
			maxL = p.LogIntervalL
		}
	}
	area := triangleAreaLogU(pairs)
	spread := maxL - minL
	return ConvergenceTriangle{
		Pairwise:                   pairs,
		PairwiseCount:              len(pairs),
		AllClosedFormSolved:        true,
		SingleIntersectionFound:    spread < 1e-9 && area < 1e-9,
		LogSpread:                  spread,
		ScaleRatioMaxOverMin:       math.Exp(spread),
		TriangleAreaInLogUPlane:    area,
		MismatchNonzero:            spread > 1e-6 && area > 1e-6,
		UsesObservedLedger:         true,
		ThresholdCorrected:         false,
		PhysicalUnificationDerived: false,
		Verdict:                    "nonzero mismatch triangle under the quarantined default Z-pole comparison",
	}
}

func buildPair(pair, gi, gj string, ai, aj, bi, bj float64, obs EmpiricalComparisonLedger) PairwiseIntersection {
	L := 2 * math.Pi * (ai - aj) / (bi - bj)
	M := obs.ScaleGeV * math.Exp(L)
	alphaStar := ai - (bi/(2*math.Pi))*L
	u := alphaStar / (4 * math.Pi)
	return PairwiseIntersection{
		Pair:                    pair,
		GaugeI:                  gi,
		GaugeJ:                  gj,
		LogIntervalL:            L,
		ScaleGeV:                M,
		Log10ScaleGeV:           math.Log10(M),
		InferredAlphaInverse:    alphaStar,
		InferredU:               u,
		FiniteDerived:           false,
		UsesObservedLedger:      true,
		UsesThresholdLedger:     false,
		ClosedFormSolved:        finite(L) && finite(M) && finite(u),
		PhysicalPredictionClaim: false,
		Verdict:                 "pairwise closed-form intersection; comparison-only",
	}
}

func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }

func triangleAreaLogU(ps []PairwiseIntersection) float64 {
	if len(ps) != 3 {
		return math.NaN()
	}
	x1, y1 := ps[0].LogIntervalL, ps[0].InferredU
	x2, y2 := ps[1].LogIntervalL, ps[1].InferredU
	x3, y3 := ps[2].LogIntervalL, ps[2].InferredU
	return math.Abs(x1*(y2-y3)+x2*(y3-y1)+x3*(y1-y2)) / 2
}

func buildTopologicalBenchmark(obs EmpiricalComparisonLedger, tri ConvergenceTriangle) TopologicalBenchmark {
	L := 0.0
	for _, p := range tri.Pairwise {
		L += p.LogIntervalL
	}
	L /= float64(len(tri.Pairwise))
	u1 := (obs.Alpha1GUTInverse - (41.0/10.0)/(2*math.Pi)*L) / (4 * math.Pi)
	u2 := (obs.Alpha2Inverse - (-19.0/6.0)/(2*math.Pi)*L) / (4 * math.Pi)
	u3 := (obs.Alpha3Inverse - (-7.0)/(2*math.Pi)*L) / (4 * math.Pi)
	avg := (u1 + u2 + u3) / 3
	M := obs.ScaleGeV * math.Exp(L)
	delta := math.Abs(avg - 1)
	return TopologicalBenchmark{
		TopologicalBranchAvailableForComparison: true,
		TopologicalBranchAssumedAsTruth:         false,
		TopologicalU:                            1,
		CentroidLogInterval:                     L,
		CentroidScaleGeV:                        M,
		InferredU1:                              u1,
		InferredU2:                              u2,
		InferredU3:                              u3,
		AverageInferredU:                        avg,
		DeltaUFromTopologicalUnit:               delta,
		CloseToUnitAtLooseTolerance:             delta < 0.15,
		FiniteDerivationClaim:                   false,
		Verdict:                                 "default comparison is not close to u*=1; this is phenomenological tension, not a finite-theorem failure",
	}
}

func auditThresholdComparison(prev gaugecouplingboundaryseal.Analysis) ThresholdComparisonAudit {
	return ThresholdComparisonAudit{
		Gate199SymbolicThresholdTreeInherited:  prev.Summary.SymbolicTopDownTrajectoryBuilt,
		ExactThresholdBetaRowsAvailable:        prev.Trajectory.ThresholdRows == 12 && prev.Trajectory.FermionContributionVector.Equal(BetaVector{U1Y: gaugecouplingboundaryseal.R(4, 1), SU2L: gaugecouplingboundaryseal.R(4, 1), SU3C: gaugecouplingboundaryseal.R(4, 1)}),
		EmpiricalThresholdLedgerSupplied:       false,
		ThresholdOrderingKnown:                 false,
		NumericalThresholdCorrectedRunAllowed:  false,
		NumericalThresholdCorrectedRunExecuted: false,
		SharpStepSchemeStillConditional:        prev.Trajectory.TreeLevelContinuityInherited && prev.Trajectory.FiniteMatchingCorrectionsSealed,
		FiniteMatchingCorrectionsAvailable:     false,
		WZThresholdsAvailable:                  false,
		LowEnergyZPoleStrictlyInsideDomain:     false,
		FormalThresholdReadyExpression:         prev.Trajectory.Expression,
		BlockReason:                            "Gate 200 has observed Z-pole couplings but no empirical threshold mass ledger, no threshold ordering, no W/Z thresholds, and no finite matching corrections; threshold-corrected numerical running remains forbidden.",
		Verdict:                                "threshold-corrected convergence is prepared symbolically but blocked numerically",
	}
}

func auditSolverAndMatching(prev gaugecouplingboundaryseal.Analysis) MatchingAndSolverAudit {
	return MatchingAndSolverAudit{
		SolverKind:                        "closed-form pairwise logarithmic intersections inside a fixed beta-region",
		UsesExactClosedFormPairwiseLogs:   true,
		RequiresNumericalOptimization:     false,
		NumericalOptimizationUsed:         false,
		TreeLevelContinuityEnforced:       prev.Trajectory.TreeLevelContinuityInherited,
		FiniteThresholdCorrectionsAudited: true,
		FiniteThresholdCorrectionsDerived: false,
		ContinuityCondition:               "A_i(M_f^-)=A_i(M_f^+) under the sharp-step comparison convention",
		FiniteCorrectionTerm:              "δ_i^match(M_f): scheme-dependent and not finite-derived",
		Verdict:                           "exact pairwise algebra is used for the current fixed-region audit; finite matching corrections remain sealed",
	}
}

func auditFirewall(prev gaugecouplingboundaryseal.Analysis, ledger EmpiricalComparisonLedger, tri ConvergenceTriangle, bench TopologicalBenchmark, threshold ThresholdComparisonAudit, solver MatchingAndSolverAudit) FirewallAudit {
	return FirewallAudit{
		Gate199Inherited:                       prev.Summary.SymbolicTopDownTrajectoryBuilt && prev.Summary.BottomUpViabilityAuditSeparated,
		EmpiricalLedgerQuarantined:             ledger.Quarantined && ledger.ExplicitPhenomenologicalInput,
		ObservedInputsUsedForFiniteDerivation:  ledger.UsedForFiniteDerivation,
		ObservedInputsUsedForPhenomenologyOnly: ledger.Quarantined && !ledger.UsedForFiniteDerivation,
		BoundaryScaleDerived:                   false,
		AbsoluteCouplingDerived:                false,
		TopologicalUOneDerived:                 false,
		TopologicalUOneAssumed:                 bench.TopologicalBranchAssumedAsTruth,
		EightPiSquaredImported:                 false,
		ThresholdCorrectedPhysicalFitClaimed:   threshold.NumericalThresholdCorrectedRunExecuted,
		PhysicalUnificationClaimed:             tri.PhysicalUnificationDerived,
		PhysicalGaugeCouplingsDerived:          false,
		WZThresholdsDerived:                    threshold.WZThresholdsAvailable,
		FiniteMatchingCorrectionsDerived:       solver.FiniteThresholdCorrectionsDerived,
		FiniteToContinuumNormalizationDerived:  false,
		StrictNullityBefore:                    3,
		StrictNullityAfter:                     3,
		PhenomenologyNullityBefore:             1,
		PhenomenologyNullityAfter:              0,
		PhysicalPredictionNullityBefore:        1,
		PhysicalPredictionNullityAfter:         1,
		OpenRequirements: []string{
			"empirical threshold mass ledger and ordering, if a threshold-corrected comparison is requested",
			"W/Z threshold treatment and electroweak broken-phase domain convention",
			"finite threshold matching corrections or explicit continuum scheme convention",
			"absolute finite-to-continuum normalization if u*=1 is to be promoted beyond comparison",
			"possible extra BSM/finite sector if the mismatch triangle is to be closed",
		},
		RecommendedNextGate: "Gate 201 — empirical threshold ledger / B-sector deformation viability search audit",
		Verdict:             "phenomenology comparison is computable and quarantined; finite theorem and physical prediction firewalls remain sealed",
	}
}

func FormatLedger(x EmpiricalComparisonLedger) string {
	return fmt.Sprintf("%s: scale=%s=%.6g GeV αem⁻¹=%.6g sin²=%.6g αs=%.6g α1GUT⁻¹=%.6g α2⁻¹=%.6g α3⁻¹=%.6g phenom=%t quarantine=%t finiteDerivation=%t boundaryDerivation=%t source=%s", x.Name, x.ScaleSymbol, x.ScaleGeV, x.AlphaEMInverse, x.Sin2ThetaMSbar, x.AlphaS, x.Alpha1GUTInverse, x.Alpha2Inverse, x.Alpha3Inverse, x.ExplicitPhenomenologicalInput, x.Quarantined, x.UsedForFiniteDerivation, x.UsedForBoundaryDerivation, x.Source)
}

func FormatPair(p PairwiseIntersection) string {
	return fmt.Sprintf("%s %s/%s: L=%.9g M*=%.9g GeV log10M=%.6g α*⁻¹=%.6g u*=%.6g observed=%t threshold=%t closed=%t physicalClaim=%t", p.Pair, p.GaugeI, p.GaugeJ, p.LogIntervalL, p.ScaleGeV, p.Log10ScaleGeV, p.InferredAlphaInverse, p.InferredU, p.UsesObservedLedger, p.UsesThresholdLedger, p.ClosedFormSolved, p.PhysicalPredictionClaim)
}

func FormatPairs(xs []PairwiseIntersection) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatPair(x))
	}
	return strings.Join(parts, " | ")
}

func FormatTriangle(t ConvergenceTriangle) string {
	return fmt.Sprintf("pairs=%d single=%t spreadL=%.9g scaleRatio=%.9g area=%.9g mismatch=%t observed=%t threshold=%t physicalUnification=%t :: %s", t.PairwiseCount, t.SingleIntersectionFound, t.LogSpread, t.ScaleRatioMaxOverMin, t.TriangleAreaInLogUPlane, t.MismatchNonzero, t.UsesObservedLedger, t.ThresholdCorrected, t.PhysicalUnificationDerived, FormatPairs(t.Pairwise))
}

func FormatBenchmark(b TopologicalBenchmark) string {
	return fmt.Sprintf("u_top=%.6g comparison=%t assumed=%t Lcentroid=%.9g Mcentroid=%.9g GeV inferred=(%.6g,%.6g,%.6g) avg=%.6g Δu=%.6g closeUnit=%t finiteClaim=%t verdict=%s", b.TopologicalU, b.TopologicalBranchAvailableForComparison, b.TopologicalBranchAssumedAsTruth, b.CentroidLogInterval, b.CentroidScaleGeV, b.InferredU1, b.InferredU2, b.InferredU3, b.AverageInferredU, b.DeltaUFromTopologicalUnit, b.CloseToUnitAtLooseTolerance, b.FiniteDerivationClaim, b.Verdict)
}

func FormatThresholdAudit(t ThresholdComparisonAudit) string {
	return fmt.Sprintf("g199=%t rows=%t empiricalThresholds=%t ordering=%t allowed=%t executed=%t sharpConditional=%t finiteCorrections=%t WZ=%t ZDomain=%t block=%s", t.Gate199SymbolicThresholdTreeInherited, t.ExactThresholdBetaRowsAvailable, t.EmpiricalThresholdLedgerSupplied, t.ThresholdOrderingKnown, t.NumericalThresholdCorrectedRunAllowed, t.NumericalThresholdCorrectedRunExecuted, t.SharpStepSchemeStillConditional, t.FiniteMatchingCorrectionsAvailable, t.WZThresholdsAvailable, t.LowEnergyZPoleStrictlyInsideDomain, t.BlockReason)
}

func FormatSolver(s MatchingAndSolverAudit) string {
	return fmt.Sprintf("solver=%s closedForm=%t needsOpt=%t optUsed=%t continuity=%t finiteCorrectionsAudited=%t finiteCorrectionsDerived=%t condition=%s correction=%s", s.SolverKind, s.UsesExactClosedFormPairwiseLogs, s.RequiresNumericalOptimization, s.NumericalOptimizationUsed, s.TreeLevelContinuityEnforced, s.FiniteThresholdCorrectionsAudited, s.FiniteThresholdCorrectionsDerived, s.ContinuityCondition, s.FiniteCorrectionTerm)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("g199=%t ledger=%t observedFinite=%t observedPheno=%t Mstar=%t ustar=%t u1derived=%t u1assumed=%t 8pi2=%t thresholdFit=%t unification=%t gauge=%t WZ=%t finiteMatch=%t contNorm=%t strict=%d->%d pheno=%d->%d pred=%d->%d next=%s", f.Gate199Inherited, f.EmpiricalLedgerQuarantined, f.ObservedInputsUsedForFiniteDerivation, f.ObservedInputsUsedForPhenomenologyOnly, f.BoundaryScaleDerived, f.AbsoluteCouplingDerived, f.TopologicalUOneDerived, f.TopologicalUOneAssumed, f.EightPiSquaredImported, f.ThresholdCorrectedPhysicalFitClaimed, f.PhysicalUnificationClaimed, f.PhysicalGaugeCouplingsDerived, f.WZThresholdsDerived, f.FiniteMatchingCorrectionsDerived, f.FiniteToContinuumNormalizationDerived, f.StrictNullityBefore, f.StrictNullityAfter, f.PhenomenologyNullityBefore, f.PhenomenologyNullityAfter, f.PhysicalPredictionNullityBefore, f.PhysicalPredictionNullityAfter, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d ledger=%t intersections=%t mismatch=%t topo=%t thresholdsBlocked=%t noFinite=%t :: %s", s.TestsAudited, s.EmpiricalLedgerQuarantined, s.PairwiseIntersectionsSolved, s.MismatchTriangleNonzero, s.TopologicalBenchmarkComputed, s.ThresholdCorrectedEvaluationBlocked, s.NoFiniteDerivationClaim, s.Comment)
}
