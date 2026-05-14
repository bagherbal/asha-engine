// Package normalizationthresholdaudit implements Gate 177: normalization-prefactor
// or threshold-deformation branch audit after the quarantined u=1 M_Z rejection.
//
// Gate 176 showed that the conditional instanton branch u=1/g_*^2=1 is
// internally computable but does not land near the M_Z electroweak/QCD comparison
// ledger under the unthresholded one-loop flow. Gate 177 asks which missing
// ingredient could mathematically repair the mismatch:
//
//	(1) an absolute normalization prefactor u only,
//	(2) a universal threshold shift, or
//	(3) non-universal sector threshold deformations Δb_i.
//
// The observed M_Z ledger inherited from Gate 176 remains quarantined comparison
// data. This package is a solvability/firewall audit, not a derivation of physical
// constants and not a finite-core fit.
package normalizationthresholdaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/conditionalrgbranch"
)

type InputAudit struct {
	Gate176ConditionalUOneRejected    bool
	Gate176StrictUStillOpen           bool
	Gate176RatioOnlyCheckFailed       bool
	ObservedComparisonQuarantined     bool
	UsesObservedInputForFiniteTheorem bool
	B1                                float64
	B2                                float64
	B3                                float64
	Alpha1GUTInverseObserved          float64
	Alpha2InverseObserved             float64
	Alpha3InverseObserved             float64
	A1GUTCoefficientObserved          float64
	A2CoefficientObserved             float64
	A3CoefficientObserved             float64
	Verdict                           string
}

type PairFit struct {
	PairName                     string
	FitSectors                   string
	PredictedSector              string
	UInverseGStar                float64
	LogIntervalL                 float64
	MStarGeV                     float64
	PredictedMissingAlphaInverse float64
	ObservedMissingAlphaInverse  float64
	MissingRelativeResidual      float64
	PositiveU                    bool
	PositiveL                    bool
	ExactPairFit                 bool
	TripleConsistent             bool
	Verdict                      string
}

type NormalizationOnlyAudit struct {
	Formula                    string
	Unknowns                   int
	Equations                  int
	ExactTripleFit             bool
	BestFitUInverseGStar       float64
	BestFitGStar               float64
	BestFitLogIntervalL        float64
	BestFitMStarGeV            float64
	BestFitAlpha1GUTInverse    float64
	BestFitAlpha2Inverse       float64
	BestFitAlpha3Inverse       float64
	Alpha1RelativeResidual     float64
	Alpha2RelativeResidual     float64
	Alpha3RelativeResidual     float64
	MaxRelativeResidual        float64
	PositiveU                  bool
	PositiveL                  bool
	PairFits                   []PairFit
	PairLogIntervalsConsistent bool
	Verdict                    string
}

type UniversalThresholdAudit struct {
	Formula                           string
	AddsSectorRatioFreedom            bool
	EquivalentToInterceptShift        bool
	PairLogIntervalsStillInconsistent bool
	RatioOnlyMismatchStillPresent     bool
	CanRepairGate176Failure           bool
	Verdict                           string
}

type ThresholdVector struct {
	Name                 string
	UInverseGStar        float64
	LogIntervalL         float64
	MStarGeV             float64
	DeltaB1              float64
	DeltaB2              float64
	DeltaB3              float64
	EffectiveB1          float64
	EffectiveB2          float64
	EffectiveB3          float64
	EuclideanNorm        float64
	SignPatternPreserved bool
	FiniteDerived        bool
	Verdict              string
}

type NonUniversalThresholdAudit struct {
	Formula                          string
	SectorSpecificThresholds         int
	FitsExactlyForAnyChosenPositiveL bool
	UnderdeterminedWithoutFiniteRule bool
	MinimumNormForUOne               ThresholdVector
	Alpha3FitForUOne                 ThresholdVector
	UniversalShiftSufficient         bool
	FiniteThresholdOperatorDerived   bool
	CanRepairPhenomenologyByFit      bool
	CanReduceStrictNullity           bool
	Verdict                          string
}

type FirewallAudit struct {
	NormalizationPrefactorAloneSufficient     bool
	UniversalThresholdAloneSufficient         bool
	NonUniversalThresholdCanFitByConstruction bool
	NonUniversalThresholdDerived              bool
	UsesObservedInputOnlyForComparison        bool
	HiddenObservedInputUsedForDerivation      bool
	StrictNullityBefore                       int
	StrictNullityAfter                        int
	ConditionalNullityBefore                  int
	ConditionalNullityAfter                   int
	PhysicalConstantsDerived                  bool
	BoundaryScaleDerivedStrict                bool
	ThresholdCorrectionsDerived               bool
	RemainingStrictUnknowns                   []string
	RecommendedNextGate                       string
	Verdict                                   string
}

type Analysis struct {
	Previous       conditionalrgbranch.Analysis
	Input          InputAudit
	Normalization  NormalizationOnlyAudit
	Universal      UniversalThresholdAudit
	Thresholds     NonUniversalThresholdAudit
	Firewall       FirewallAudit
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := conditionalrgbranch.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev conditionalrgbranch.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.Firewall.ConditionalBranchRejectedByMZ {
		return Analysis{}, fmt.Errorf("Gate 177 requires Gate 176 to reject the unthresholded u=1 branch")
	}
	if !prev.Firewall.StrictUStillOpen || prev.Firewall.StrictNullityAfter != 3 {
		return Analysis{}, fmt.Errorf("Gate 177 requires strict absolute normalization to remain open")
	}
	if prev.Observed.UsedForDerivation || !prev.Observed.QuarantinedComparison || prev.Input.UsesObservedInputForDerivation {
		return Analysis{}, fmt.Errorf("Gate 177 refuses non-quarantined observed input")
	}
	if prev.Ratio.LIntervalsAgree || prev.Firewall.RatioCheckPasses {
		return Analysis{}, fmt.Errorf("Gate 177 expects the Gate 176 ratio-only check to fail")
	}

	b := []float64{prev.Input.B1, prev.Input.B2, prev.Input.B3}
	obsInv := []float64{prev.Observed.Alpha1GUTInverse, prev.Observed.Alpha2Inverse, prev.Observed.Alpha3Inverse}
	A := []float64{obsInv[0] / (4 * math.Pi), obsInv[1] / (4 * math.Pi), obsInv[2] / (4 * math.Pi)}
	input := InputAudit{
		Gate176ConditionalUOneRejected:    prev.Firewall.ConditionalBranchRejectedByMZ,
		Gate176StrictUStillOpen:           prev.Firewall.StrictUStillOpen,
		Gate176RatioOnlyCheckFailed:       !prev.Ratio.LIntervalsAgree && !prev.Firewall.RatioCheckPasses,
		ObservedComparisonQuarantined:     prev.Observed.QuarantinedComparison && !prev.Observed.UsedForDerivation,
		UsesObservedInputForFiniteTheorem: false,
		B1:                                b[0],
		B2:                                b[1],
		B3:                                b[2],
		Alpha1GUTInverseObserved:          obsInv[0],
		Alpha2InverseObserved:             obsInv[1],
		Alpha3InverseObserved:             obsInv[2],
		A1GUTCoefficientObserved:          A[0],
		A2CoefficientObserved:             A[1],
		A3CoefficientObserved:             A[2],
		Verdict:                           "Gate 176 rejection may be analyzed only as a quarantined comparison problem, not as finite-core input",
	}

	norm := buildNormalizationOnlyAudit(b, A, obsInv, prev.Observed.ScaleGeV, eps)
	univ := UniversalThresholdAudit{
		Formula:                           "A_i = u + ((b_i+δ)/(8π²))L; subtracting sectors cancels δ, so all ratio/L-consistency obstructions are unchanged",
		AddsSectorRatioFreedom:            false,
		EquivalentToInterceptShift:        true,
		PairLogIntervalsStillInconsistent: !norm.PairLogIntervalsConsistent,
		RatioOnlyMismatchStillPresent:     !prev.Ratio.LIntervalsAgree,
		CanRepairGate176Failure:           false,
		Verdict:                           "a universal threshold shift is mathematically indistinguishable from an intercept/prefactor change for relative running and cannot repair the failed ratio audit",
	}
	thresholds := buildNonUniversalThresholdAudit(b, A, obsInv, prev.Observed.ScaleGeV, eps)
	firewall := FirewallAudit{
		NormalizationPrefactorAloneSufficient:     norm.ExactTripleFit,
		UniversalThresholdAloneSufficient:         univ.CanRepairGate176Failure,
		NonUniversalThresholdCanFitByConstruction: thresholds.CanRepairPhenomenologyByFit,
		NonUniversalThresholdDerived:              thresholds.FiniteThresholdOperatorDerived,
		UsesObservedInputOnlyForComparison:        input.ObservedComparisonQuarantined && !input.UsesObservedInputForFiniteTheorem,
		HiddenObservedInputUsedForDerivation:      false,
		StrictNullityBefore:                       3,
		StrictNullityAfter:                        3,
		ConditionalNullityBefore:                  2,
		ConditionalNullityAfter:                   2,
		PhysicalConstantsDerived:                  false,
		BoundaryScaleDerivedStrict:                false,
		ThresholdCorrectionsDerived:               false,
		RemainingStrictUnknowns: []string{
			"u=1/g_*²: absolute normalization still lacks a finite-to-continuum trace bridge",
			"L=ln(M*/M_Z): no finite boundary-scale selector is derived",
			"Δb_i: non-universal threshold deformation can fit by construction but lacks a finite activation/operator derivation",
		},
		RecommendedNextGate: "Gate 178 — finite threshold operator / decoupling spectrum search",
		Verdict:             "normalization-only and universal-threshold repairs are overconstrained; arbitrary non-universal thresholds can fit but are underived, so no nullity reduction is allowed",
	}

	return Analysis{
		Previous:       prev,
		Input:          input,
		Normalization:  norm,
		Universal:      univ,
		Thresholds:     thresholds,
		Firewall:       firewall,
		TruthStatement: "Gate 177 separates the Gate 176 failure into three repair classes. A free absolute normalization prefactor improves the comparison but cannot solve all three one-loop equations with the closed beta vector: the pairwise log intervals remain inconsistent. A universal threshold shift adds no relative-running freedom. Non-universal sector thresholds can exactly reproduce the comparison ledger for infinitely many choices of L, but that is a fit family, not a finite theorem; the minimum-norm u=1 threshold vector is itself selected by an external Euclidean criterion. Strict nullity therefore remains 3.",
	}, nil
}

func buildNormalizationOnlyAudit(b, A, obsInv []float64, muGeV, eps float64) NormalizationOnlyAudit {
	beta := make([]float64, 3)
	for i := range b {
		beta[i] = b[i] / (8 * math.Pi * math.Pi)
	}
	u, L := leastSquaresInterceptSlope(beta, A)
	predA := []float64{u + beta[0]*L, u + beta[1]*L, u + beta[2]*L}
	predInv := []float64{4 * math.Pi * predA[0], 4 * math.Pi * predA[1], 4 * math.Pi * predA[2]}
	res := []float64{relResidual(predInv[0], obsInv[0]), relResidual(predInv[1], obsInv[1]), relResidual(predInv[2], obsInv[2])}
	maxRes := maxAbs(res)
	pairs := []PairFit{
		buildPairFit("α1-α2", 0, 1, 2, b, beta, A, obsInv, muGeV, eps),
		buildPairFit("α1-α3", 0, 2, 1, b, beta, A, obsInv, muGeV, eps),
		buildPairFit("α2-α3", 1, 2, 0, b, beta, A, obsInv, muGeV, eps),
	}
	consistent := true
	for i := 1; i < len(pairs); i++ {
		if math.Abs(pairs[i].LogIntervalL-pairs[0].LogIntervalL)/math.Max(1, math.Abs(pairs[0].LogIntervalL)) > 0.05 {
			consistent = false
		}
	}
	exact := maxRes < 1e-6
	verdict := "normalization prefactor alone does not give an exact three-sector M_Z solution with the closed one-loop beta vector"
	if exact {
		verdict = "normalization prefactor alone closes the three-sector comparison system"
	}
	return NormalizationOnlyAudit{
		Formula:                    "A_i(M_Z)=α_i^{-1}/4π = u + (b_i/8π²)L, with unknown u and L but fixed b_i",
		Unknowns:                   2,
		Equations:                  3,
		ExactTripleFit:             exact,
		BestFitUInverseGStar:       u,
		BestFitGStar:               1 / math.Sqrt(u),
		BestFitLogIntervalL:        L,
		BestFitMStarGeV:            muGeV * math.Exp(L),
		BestFitAlpha1GUTInverse:    predInv[0],
		BestFitAlpha2Inverse:       predInv[1],
		BestFitAlpha3Inverse:       predInv[2],
		Alpha1RelativeResidual:     res[0],
		Alpha2RelativeResidual:     res[1],
		Alpha3RelativeResidual:     res[2],
		MaxRelativeResidual:        maxRes,
		PositiveU:                  u > 0,
		PositiveL:                  L > 0,
		PairFits:                   pairs,
		PairLogIntervalsConsistent: consistent,
		Verdict:                    verdict,
	}
}

func buildPairFit(name string, i, j, missing int, b, beta, A, obsInv []float64, muGeV, eps float64) PairFit {
	L := (A[i] - A[j]) / (beta[i] - beta[j])
	u := A[i] - beta[i]*L
	predMissingA := u + beta[missing]*L
	predMissingInv := 4 * math.Pi * predMissingA
	resid := relResidual(predMissingInv, obsInv[missing])
	triple := math.Abs(resid) < 1e-6
	verdict := "pair fits exactly but misses the third sector"
	if triple {
		verdict = "pair fit is triple-consistent"
	}
	return PairFit{
		PairName:                     name,
		FitSectors:                   sectorName(i) + "," + sectorName(j),
		PredictedSector:              sectorName(missing),
		UInverseGStar:                u,
		LogIntervalL:                 L,
		MStarGeV:                     muGeV * math.Exp(L),
		PredictedMissingAlphaInverse: predMissingInv,
		ObservedMissingAlphaInverse:  obsInv[missing],
		MissingRelativeResidual:      resid,
		PositiveU:                    u > 0,
		PositiveL:                    L > 0,
		ExactPairFit:                 math.Abs((u+beta[i]*L)-A[i]) < eps && math.Abs((u+beta[j]*L)-A[j]) < eps,
		TripleConsistent:             triple,
		Verdict:                      verdict,
	}
}

func buildNonUniversalThresholdAudit(b, A, obsInv []float64, muGeV, eps float64) NonUniversalThresholdAudit {
	minNorm := minimumNormThresholdForFixedU("minimum-norm u=1 threshold vector", 1, b, A, muGeV)
	alpha3L := (A[2] - 1) / (b[2] / (8 * math.Pi * math.Pi))
	alpha3Fit := thresholdVectorAt("u=1, α3-fit threshold witness", 1, alpha3L, b, A, muGeV)
	return NonUniversalThresholdAudit{
		Formula:                          "A_i = u + ((b_i+Δb_i)/(8π²))L; hence Δb_i(L,u)=8π²(A_i-u)/L - b_i",
		SectorSpecificThresholds:         3,
		FitsExactlyForAnyChosenPositiveL: true,
		UnderdeterminedWithoutFiniteRule: true,
		MinimumNormForUOne:               minNorm,
		Alpha3FitForUOne:                 alpha3Fit,
		UniversalShiftSufficient:         false,
		FiniteThresholdOperatorDerived:   false,
		CanRepairPhenomenologyByFit:      true,
		CanReduceStrictNullity:           false,
		Verdict:                          "sector-specific threshold vectors can fit the comparison ledger by construction, but no finite threshold operator or decoupling spectrum selects Δb_i or L",
	}
}

func minimumNormThresholdForFixedU(name string, u float64, b, A []float64, muGeV float64) ThresholdVector {
	c := make([]float64, 3)
	for i := range A {
		c[i] = 8 * math.Pi * math.Pi * (A[i] - u)
	}
	dotCB := dot(c, b)
	dotCC := dot(c, c)
	L := math.Inf(1)
	if dotCB > 0 && dotCC > 0 {
		L = dotCC / dotCB
	}
	return thresholdVectorAt(name, u, L, b, A, muGeV)
}

func thresholdVectorAt(name string, u, L float64, b, A []float64, muGeV float64) ThresholdVector {
	d := []float64{math.NaN(), math.NaN(), math.NaN()}
	eff := []float64{math.NaN(), math.NaN(), math.NaN()}
	if L != 0 && !math.IsInf(L, 0) && !math.IsNaN(L) {
		for i := range A {
			d[i] = 8*math.Pi*math.Pi*(A[i]-u)/L - b[i]
			eff[i] = b[i] + d[i]
		}
	}
	norm := math.Sqrt(d[0]*d[0] + d[1]*d[1] + d[2]*d[2])
	signs := eff[0] > 0 && eff[1] < 0 && eff[2] < 0
	return ThresholdVector{
		Name:                 name,
		UInverseGStar:        u,
		LogIntervalL:         L,
		MStarGeV:             muGeV * math.Exp(L),
		DeltaB1:              d[0],
		DeltaB2:              d[1],
		DeltaB3:              d[2],
		EffectiveB1:          eff[0],
		EffectiveB2:          eff[1],
		EffectiveB3:          eff[2],
		EuclideanNorm:        norm,
		SignPatternPreserved: signs,
		FiniteDerived:        false,
		Verdict:              "comparison witness only; selected by external numerical criterion, not by finite algebra",
	}
}

func leastSquaresInterceptSlope(beta, A []float64) (float64, float64) {
	n := float64(len(A))
	var sb, sbb, sa, sba float64
	for i := range A {
		sb += beta[i]
		sbb += beta[i] * beta[i]
		sa += A[i]
		sba += beta[i] * A[i]
	}
	det := n*sbb - sb*sb
	if math.Abs(det) == 0 {
		return math.NaN(), math.NaN()
	}
	u := (sa*sbb - sb*sba) / det
	L := (n*sba - sb*sa) / det
	return u, L
}

func sectorName(i int) string {
	switch i {
	case 0:
		return "α1_GUT"
	case 1:
		return "α2"
	case 2:
		return "α3"
	default:
		return fmt.Sprintf("sector%d", i)
	}
}

func dot(a, b []float64) float64 {
	var s float64
	for i := range a {
		s += a[i] * b[i]
	}
	return s
}

func relResidual(predicted, observed float64) float64 {
	if observed == 0 || math.IsNaN(predicted) || math.IsInf(predicted, 0) {
		return math.Inf(1)
	}
	return (predicted - observed) / observed
}

func maxAbs(xs []float64) float64 {
	m := 0.0
	for _, x := range xs {
		if math.Abs(x) > m {
			m = math.Abs(x)
		}
	}
	return m
}

func FormatInput(x InputAudit) string {
	return fmt.Sprintf("gate176-rejected=%t strict-u-open=%t ratio-failed=%t comparison-quarantined=%t finite-observed-input=%t b=(%.12g,%.12g,%.12g) αinv=(%.9g,%.9g,%.9g) A=(%.9g,%.9g,%.9g) verdict=%s", x.Gate176ConditionalUOneRejected, x.Gate176StrictUStillOpen, x.Gate176RatioOnlyCheckFailed, x.ObservedComparisonQuarantined, x.UsesObservedInputForFiniteTheorem, x.B1, x.B2, x.B3, x.Alpha1GUTInverseObserved, x.Alpha2InverseObserved, x.Alpha3InverseObserved, x.A1GUTCoefficientObserved, x.A2CoefficientObserved, x.A3CoefficientObserved, x.Verdict)
}

func FormatPairFit(x PairFit) string {
	return fmt.Sprintf("%s fit(%s) predicts %s: u=%.9g L=%.9g M*=%.9g GeV pred=%.9g obs=%.9g rel=%.3g positive=(u:%t,L:%t) triple=%t verdict=%s", x.PairName, x.FitSectors, x.PredictedSector, x.UInverseGStar, x.LogIntervalL, x.MStarGeV, x.PredictedMissingAlphaInverse, x.ObservedMissingAlphaInverse, x.MissingRelativeResidual, x.PositiveU, x.PositiveL, x.TripleConsistent, x.Verdict)
}

func FormatPairFits(xs []PairFit) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatPairFit(x))
	}
	return strings.Join(parts, " | ")
}

func FormatNormalization(x NormalizationOnlyAudit) string {
	return fmt.Sprintf("%s: unknowns=%d equations=%d exact=%t best(u=%.9g,g=%.9g,L=%.9g,M*=%.9g GeV) predαinv=(%.9g,%.9g,%.9g) rel=(%.3g,%.3g,%.3g) max=%.3g positive=(u:%t,L:%t) pair-L-consistent=%t pairs=[%s] verdict=%s", x.Formula, x.Unknowns, x.Equations, x.ExactTripleFit, x.BestFitUInverseGStar, x.BestFitGStar, x.BestFitLogIntervalL, x.BestFitMStarGeV, x.BestFitAlpha1GUTInverse, x.BestFitAlpha2Inverse, x.BestFitAlpha3Inverse, x.Alpha1RelativeResidual, x.Alpha2RelativeResidual, x.Alpha3RelativeResidual, x.MaxRelativeResidual, x.PositiveU, x.PositiveL, x.PairLogIntervalsConsistent, FormatPairFits(x.PairFits), x.Verdict)
}

func FormatUniversal(x UniversalThresholdAudit) string {
	return fmt.Sprintf("%s: adds-ratio-freedom=%t equivalent-intercept=%t pair-L-inconsistent=%t ratio-mismatch=%t repair=%t verdict=%s", x.Formula, x.AddsSectorRatioFreedom, x.EquivalentToInterceptShift, x.PairLogIntervalsStillInconsistent, x.RatioOnlyMismatchStillPresent, x.CanRepairGate176Failure, x.Verdict)
}

func FormatThresholdVector(x ThresholdVector) string {
	return fmt.Sprintf("%s: u=%.9g L=%.9g M*=%.9g GeV Δb=(%.9g,%.9g,%.9g) b_eff=(%.9g,%.9g,%.9g) ||Δb||=%.9g sign-preserved=%t finite-derived=%t verdict=%s", x.Name, x.UInverseGStar, x.LogIntervalL, x.MStarGeV, x.DeltaB1, x.DeltaB2, x.DeltaB3, x.EffectiveB1, x.EffectiveB2, x.EffectiveB3, x.EuclideanNorm, x.SignPatternPreserved, x.FiniteDerived, x.Verdict)
}

func FormatThresholds(x NonUniversalThresholdAudit) string {
	return fmt.Sprintf("%s: sectors=%d exact-any-L=%t underdetermined=%t universal-sufficient=%t finite-operator=%t fit-by-construction=%t reduce-nullity=%t min=[%s] alpha3=[%s] verdict=%s", x.Formula, x.SectorSpecificThresholds, x.FitsExactlyForAnyChosenPositiveL, x.UnderdeterminedWithoutFiniteRule, x.UniversalShiftSufficient, x.FiniteThresholdOperatorDerived, x.CanRepairPhenomenologyByFit, x.CanReduceStrictNullity, FormatThresholdVector(x.MinimumNormForUOne), FormatThresholdVector(x.Alpha3FitForUOne), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("norm-only=%t universal-threshold=%t nonuniversal-fit=%t nonuniversal-derived=%t comparison-only=%t hidden-observed=%t strict-nullity=%d->%d conditional-nullity=%d->%d constants=%t L-strict=%t thresholds-derived=%t remaining=%s next=%s verdict=%s", x.NormalizationPrefactorAloneSufficient, x.UniversalThresholdAloneSufficient, x.NonUniversalThresholdCanFitByConstruction, x.NonUniversalThresholdDerived, x.UsesObservedInputOnlyForComparison, x.HiddenObservedInputUsedForDerivation, x.StrictNullityBefore, x.StrictNullityAfter, x.ConditionalNullityBefore, x.ConditionalNullityAfter, x.PhysicalConstantsDerived, x.BoundaryScaleDerivedStrict, x.ThresholdCorrectionsDerived, strings.Join(x.RemainingStrictUnknowns, "; "), x.RecommendedNextGate, x.Verdict)
}
