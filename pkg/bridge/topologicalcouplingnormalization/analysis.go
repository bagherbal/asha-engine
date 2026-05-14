// Package topologicalcouplingnormalization implements Gate 328:
// Topological Action / Chern-Weil Coupling Normalization Factor Audit.
//
// Gate 327 discovered a striking 8π coupling witness:
//
//	S_top/π = 8π,
//	dim_R(A_F)π/N_gen = 24π/3 = 8π,
//
// yielding g_*² = 1/2 and the tree-level Higgs proxy
// m_H = v√(1197/4624) ≈ 125.274 GeV.  Gate 328 audits whether the
// denominator π is forced by spectral-action/Chern-Weil normalization, or
// whether the conventional Yang-Mills instanton normalization gives a competing
// denominator 2π.  The gate keeps both lanes explicit and refuses to promote
// the 8π witness into a native α_GUT theorem until the missing factor-of-two
// theorem is derived.
package topologicalcouplingnormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE328-TOPOLOGICAL-CHERN-WEIL-COUPLING-NORMALIZATION-FACTOR-AUDIT"

	StatusTopologicalActionLedgerFormalized = "CONDITIONAL_SUPPORT_TOPOLOGICAL_ACTION_LEDGER_FORMALIZED"
	StatusChernWeilNormalizationAudited     = "CONDITIONAL_SUPPORT_CHERN_WEIL_NORMALIZATION_AUDITED"
	StatusPiDenominatorWitnessComputed      = "CONDITIONAL_SUPPORT_PI_DENOMINATOR_EIGHT_PI_WITNESS_COMPUTED"
	StatusTwoPiDenominatorLaneComputed      = "CONDITIONAL_SUPPORT_TWO_PI_DENOMINATOR_STANDARD_LANE_COMPUTED"
	StatusFactorTwoObligationIdentified     = "CONDITIONAL_SUPPORT_FACTOR_TWO_NORMALIZATION_OBLIGATION_IDENTIFIED"
	StatusHiggsProxyBranchesComputed        = "CONDITIONAL_SUPPORT_HIGGS_PROXY_BRANCHES_COMPUTED"

	StatusTensionPiLaneMatchesHiggsButNeedsProof = "CONDITIONAL_TENSION_PI_LANE_MATCHES_HIGGS_PROXY_BUT_NEEDS_NORMALIZATION_PROOF"
	StatusTensionTwoPiLaneReturnsOldGStarOne     = "CONDITIONAL_TENSION_TWO_PI_LANE_RETURNS_OLD_GSTAR_ONE_BOUNDARY"

	StatusFailedEightPiNotPromoted        = "FAILED_ROUTE_EIGHT_PI_COUPLING_NOT_PROMOTED_TO_SPECTRAL_ACTION_THEOREM"
	StatusFailedFactorTwoNotDerived       = "FAILED_ROUTE_FACTOR_TWO_NORMALIZATION_NOT_DERIVED"
	StatusFailedTraceRepIndexStillMissing = "FAILED_ROUTE_REQUIRED_TRACE_REP_INDEX_STILL_MISSING"
	StatusFailedAlphaGUTStillSealed       = "FAILED_ROUTE_ALPHA_GUT_ABSOLUTE_VALUE_STILL_SEALED"
	StatusFailedColliderMassNotClaimed    = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 327

	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	finiteAlgebraRealDim     = 24.0
	generationCount          = 3.0
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10
)

type Inputs struct {
	HighestInheritedGate int
	STopFormula          string
	STop                 float64
	DimRealAF            float64
	GenerationCount      float64
	ContactScalarRatio   float64
	ElectroweakVEVGeV    float64
	ObservedHiggsGeV     float64
	AddsEmpiricalFit     bool
	Status               string
}

type NormalizationLane struct {
	Name             string
	AlphaInverseForm string
	AlphaInverse     float64
	GStarSquared     float64
	LambdaH          float64
	HiggsMassGeV     float64
	DifferenceGeV    float64
	RelativeErrorPct float64
	MatchesNearHiggs bool
	PromotedTheorem  bool
	Status           string
}

type ChernWeilAudit struct {
	InstantonActionFormula       string
	AlphaDefinition              string
	ConventionalAlphaInverse     float64
	PiLaneAlphaInverse           float64
	FactorOfTwo                  float64
	ConventionalDenominator      string
	PiLaneRequiresExtraHalf      bool
	DoubledSpaceCouldSupplyHalf  bool
	DoubledSpaceHalfDerivedHere  bool
	RepresentationTraceRequired  bool
	DerivedAsSpectralActionProof bool
	DiscrepancyNote              string
	Status                       string
}

type DimensionWitness struct {
	Formula                 string
	Value                   float64
	EqualsPiLane            bool
	UsesOnlyDerivedIntegers bool
	RequiresPiNormalization bool
	PromotedToActionTheorem bool
	Status                  string
}

type FirewallAudit struct {
	NoEmpiricalAlphaInserted   bool
	NoObservedHiggsFitInserted bool
	NoFactorTwoInvented        bool
	NoTraceIndexInvented       bool
	NoPoleMassClaimed          bool
	NoFinalColliderMassClaimed bool
	EightPiKeptAsWitness       bool
	FiniteCorePolluted         bool
	Status                     string
}

type Summary struct {
	PiLaneGStarHalf       bool
	TwoPiLaneGStarOne     bool
	PiLaneHiggsProxyWorks bool
	FactorTwoMissing      bool
	NativeAlphaClosed     bool
	FinalMassClaimed      bool
	DirectAnswer          string
	NextObligation        string
	Status                string
}

type Analysis struct {
	Inputs    Inputs
	PiLane    NormalizationLane
	TwoPiLane NormalizationLane
	ChernWeil ChernWeilAudit
	Dimension DimensionWitness
	Audit     FirewallAudit
	Summary   Summary
	Truth     string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inputs := compileInputs()
	piLane := computeLane("S_top / π", inputs.STop/math.Pi, StatusPiDenominatorWitnessComputed, false)
	twoPiLane := computeLane("S_top / 2π", inputs.STop/(2.0*math.Pi), StatusTwoPiDenominatorLaneComputed, false)
	cw := auditChernWeil(inputs, piLane, twoPiLane)
	dim := auditDimensionWitness(piLane)
	audit := auditFirewalls(cw, piLane, twoPiLane, dim)
	summary := compileSummary(piLane, twoPiLane, cw, audit)
	truth := "Gate 328 audits the missing normalization factor behind the Gate 327 8π witness. The π-denominator lane gives α_GUT^{-1}=8π, g_*²=1/2, and m_H≈125.274 GeV, while the conventional Chern-Weil/Yang-Mills instanton denominator 2π gives α^{-1}=4π, g_*²=1, and returns the old ≈177.164 GeV tree proxy. Therefore the 8π lane remains a powerful witness, but its promotion requires a native factor-of-two/action-normalization theorem, such as a derived half-weight from the doubled-space spectral action or a representation trace theorem."
	return Analysis{Inputs: inputs, PiLane: piLane, TwoPiLane: twoPiLane, ChernWeil: cw, Dimension: dim, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		STopFormula:          "S_top = 8π²",
		STop:                 8.0 * math.Pi * math.Pi,
		DimRealAF:            finiteAlgebraRealDim,
		GenerationCount:      generationCount,
		ContactScalarRatio:   contactScalarNumerator / contactScalarDenominator,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		ObservedHiggsGeV:     observedHiggsGeV,
		AddsEmpiricalFit:     false,
		Status:               StatusTopologicalActionLedgerFormalized,
	}
}

func computeLane(name string, alphaInv float64, status string, promoted bool) NormalizationLane {
	g2 := 4.0 * math.Pi / alphaInv
	lambda := (contactScalarNumerator / contactScalarDenominator) * g2
	mass := electroweakVEVGeV * math.Sqrt(2.0*lambda)
	diff := mass - observedHiggsGeV
	return NormalizationLane{
		Name:             name,
		AlphaInverseForm: fmt.Sprintf("α^{-1} = %s", name),
		AlphaInverse:     alphaInv,
		GStarSquared:     g2,
		LambdaH:          lambda,
		HiggsMassGeV:     mass,
		DifferenceGeV:    diff,
		RelativeErrorPct: diff / observedHiggsGeV * 100.0,
		MatchesNearHiggs: math.Abs(diff) < 0.5,
		PromotedTheorem:  promoted,
		Status:           status,
	}
}

func auditChernWeil(inputs Inputs, piLane, twoPiLane NormalizationLane) ChernWeilAudit {
	return ChernWeilAudit{
		InstantonActionFormula:       "S_YM(k=1)=8π²/g²; α=g²/(4π)",
		AlphaDefinition:              "α^{-1}=4π/g²",
		ConventionalAlphaInverse:     inputs.STop / (2.0 * math.Pi),
		PiLaneAlphaInverse:           inputs.STop / math.Pi,
		FactorOfTwo:                  piLane.AlphaInverse / twoPiLane.AlphaInverse,
		ConventionalDenominator:      "2π under the usual instanton-action conversion S=8π²/g² -> α^{-1}=S/(2π)",
		PiLaneRequiresExtraHalf:      true,
		DoubledSpaceCouldSupplyHalf:  true,
		DoubledSpaceHalfDerivedHere:  false,
		RepresentationTraceRequired:  true,
		DerivedAsSpectralActionProof: false,
		DiscrepancyNote:              "The Higgs-successful α^{-1}=S_top/π lane is exactly twice the conventional instanton conversion α^{-1}=S_top/(2π). Gate 328 therefore identifies a precise missing factor-of-two theorem rather than declaring the 8π witness proven.",
		Status:                       StatusChernWeilNormalizationAudited,
	}
}

func auditDimensionWitness(piLane NormalizationLane) DimensionWitness {
	value := finiteAlgebraRealDim * math.Pi / generationCount
	return DimensionWitness{
		Formula:                 "dim_R(A_F)π/N_gen = 24π/3",
		Value:                   value,
		EqualsPiLane:            nearlyEqual(value, piLane.AlphaInverse, 1e-12),
		UsesOnlyDerivedIntegers: true,
		RequiresPiNormalization: true,
		PromotedToActionTheorem: false,
		Status:                  StatusFactorTwoObligationIdentified,
	}
}

func auditFirewalls(cw ChernWeilAudit, piLane, twoPiLane NormalizationLane, dim DimensionWitness) FirewallAudit {
	return FirewallAudit{
		NoEmpiricalAlphaInserted:   true,
		NoObservedHiggsFitInserted: true,
		NoFactorTwoInvented:        !cw.DoubledSpaceHalfDerivedHere && !piLane.PromotedTheorem,
		NoTraceIndexInvented:       cw.RepresentationTraceRequired && !cw.DerivedAsSpectralActionProof,
		NoPoleMassClaimed:          true,
		NoFinalColliderMassClaimed: true,
		EightPiKeptAsWitness:       !piLane.PromotedTheorem && !dim.PromotedToActionTheorem,
		FiniteCorePolluted:         false,
		Status:                     StatusFailedEightPiNotPromoted,
	}
}

func compileSummary(piLane, twoPiLane NormalizationLane, cw ChernWeilAudit, audit FirewallAudit) Summary {
	factorMissing := cw.PiLaneRequiresExtraHalf && !cw.DoubledSpaceHalfDerivedHere
	return Summary{
		PiLaneGStarHalf:       nearlyEqual(piLane.GStarSquared, 0.5, 1e-12),
		TwoPiLaneGStarOne:     nearlyEqual(twoPiLane.GStarSquared, 1.0, 1e-12),
		PiLaneHiggsProxyWorks: piLane.MatchesNearHiggs,
		FactorTwoMissing:      factorMissing,
		NativeAlphaClosed:     false,
		FinalMassClaimed:      false,
		DirectAnswer:          "The 8π lane reproduces the Higgs proxy, but Gate 328 does not promote it because the standard Chern-Weil conversion prefers a 2π denominator unless a native half-weight/action-normalization theorem is derived.",
		NextObligation:        "Derive the factor-of-two normalization from the doubled spectral action, real-structure quotient, or explicit representation trace index.",
		Status:                StatusTensionPiLaneMatchesHiggsButNeedsProof,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.PiLane.Status,
		a.TwoPiLane.Status,
		a.ChernWeil.Status,
		a.Dimension.Status,
		a.Audit.Status,
		a.Summary.Status,
		StatusHiggsProxyBranchesComputed,
		StatusTensionTwoPiLaneReturnsOldGStarOne,
		StatusFailedFactorTwoNotDerived,
		StatusFailedTraceRepIndexStillMissing,
		StatusFailedAlphaGUTStillSealed,
		StatusFailedColliderMassNotClaimed,
	}
}

func FormatInputs(v Inputs) string {
	return fmt.Sprintf("gate=%d; %s; S_top=%.12f; dim_R(A_F)=%.0f; N_gen=%.0f; contact_shape=%.12f; empirical_fit=%t", v.HighestInheritedGate, v.STopFormula, v.STop, v.DimRealAF, v.GenerationCount, v.ContactScalarRatio, v.AddsEmpiricalFit)
}

func FormatLane(v NormalizationLane) string {
	return fmt.Sprintf("%s: alpha_inv=%.12f; g2=%.12f; lambda=%.12f; m=%.6f GeV; diff=%+.6f GeV (%+.6f%%); promoted=%t; status=%s", v.Name, v.AlphaInverse, v.GStarSquared, v.LambdaH, v.HiggsMassGeV, v.DifferenceGeV, v.RelativeErrorPct, v.PromotedTheorem, v.Status)
}

func FormatChernWeil(v ChernWeilAudit) string {
	return fmt.Sprintf("%s; %s; conventional_alpha_inv=%.12f; pi_lane_alpha_inv=%.12f; factor=%.6f; half_possible=%t; half_derived=%t; trace_required=%t; spectral_proof=%t; note=%s", v.InstantonActionFormula, v.ConventionalDenominator, v.ConventionalAlphaInverse, v.PiLaneAlphaInverse, v.FactorOfTwo, v.DoubledSpaceCouldSupplyHalf, v.DoubledSpaceHalfDerivedHere, v.RepresentationTraceRequired, v.DerivedAsSpectralActionProof, v.DiscrepancyNote)
}

func FormatDimension(v DimensionWitness) string {
	return fmt.Sprintf("%s=%.12f; equals_pi_lane=%t; derived_integers=%t; requires_pi_norm=%t; theorem=%t", v.Formula, v.Value, v.EqualsPiLane, v.UsesOnlyDerivedIntegers, v.RequiresPiNormalization, v.PromotedToActionTheorem)
}

func FormatAudit(v FirewallAudit) string {
	return fmt.Sprintf("no_alpha_fit=%t; no_higgs_fit=%t; no_factor_two_invented=%t; no_trace_index_invented=%t; no_pole=%t; no_final=%t; witness_only=%t; polluted=%t", v.NoEmpiricalAlphaInserted, v.NoObservedHiggsFitInserted, v.NoFactorTwoInvented, v.NoTraceIndexInvented, v.NoPoleMassClaimed, v.NoFinalColliderMassClaimed, v.EightPiKeptAsWitness, v.FiniteCorePolluted)
}

func FormatSummary(v Summary) string {
	return fmt.Sprintf("pi_g2_half=%t; twopi_g2_one=%t; pi_proxy=%t; factor_two_missing=%t; native_closed=%t; final_mass=%t; next=%s", v.PiLaneGStarHalf, v.TwoPiLaneGStarOne, v.PiLaneHiggsProxyWorks, v.FactorTwoMissing, v.NativeAlphaClosed, v.FinalMassClaimed, v.NextObligation)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 328 Registry Audit — Topological Action / Chern-Weil Coupling Normalization Factor Audit\n\n")
	b.WriteString("## Gate identity\n\n")
	b.WriteString("- **Gate:** 328\n")
	b.WriteString("- **Package:** `pkg/bridge/topologicalcouplingnormalization`\n")
	b.WriteString("- **Theorem:** `TopologicalActionChernWeilCouplingNormalizationFactorAuditTheorem`\n")
	b.WriteString("- **Audit ID:** `" + AuditID + "`\n")
	b.WriteString("- **Layer:** Bridge / Phase-II Absolute Coupling Normalization\n")
	b.WriteString("- **Purpose:** audit whether the Gate 327 `S_top/π = 8π` witness is a native spectral-action gauge-coupling theorem or still requires a missing factor-of-two/action-normalization proof.\n\n")

	b.WriteString("## Input ledger\n\n")
	b.WriteString("```text\n" + FormatInputs(a.Inputs) + "\n```\n\n")

	b.WriteString("## Two coupling-normalization lanes\n\n")
	b.WriteString("| Lane | α⁻¹ | g_*² | λ_H | Tree proxy | Verdict |\n")
	b.WriteString("| --- | ---: | ---: | ---: | ---: | --- |\n")
	b.WriteString(fmt.Sprintf("| `S_top/π` | %.12f | %.12f | %.12f | %.6f GeV | Higgs-successful witness; not yet theorem |\n", a.PiLane.AlphaInverse, a.PiLane.GStarSquared, a.PiLane.LambdaH, a.PiLane.HiggsMassGeV))
	b.WriteString(fmt.Sprintf("| `S_top/(2π)` | %.12f | %.12f | %.12f | %.6f GeV | conventional instanton conversion; returns old g_*²=1 lane |\n\n", a.TwoPiLane.AlphaInverse, a.TwoPiLane.GStarSquared, a.TwoPiLane.LambdaH, a.TwoPiLane.HiggsMassGeV))

	b.WriteString("## Chern-Weil normalization audit\n\n")
	b.WriteString("```text\n" + FormatChernWeil(a.ChernWeil) + "\n```\n\n")
	b.WriteString("The gate identifies an exact factor-of-two obstruction. The Higgs-successful lane uses `α⁻¹=S_top/π`, while the usual Yang-Mills instanton conversion from `S=8π²/g²` gives `α⁻¹=S/(2π)`. Therefore the 8π result must be supported by an explicit doubled-space, quotient, or representation-trace normalization theorem before it can be promoted.\n\n")

	b.WriteString("## Dimension/generation witness\n\n")
	b.WriteString("```text\n" + FormatDimension(a.Dimension) + "\n```\n\n")
	b.WriteString("The dimension/generation expression exactly matches the `8π` lane, but Gate 328 keeps it as a witness rather than a spectral-action theorem.\n\n")

	b.WriteString("## Final status ledger\n\n")
	b.WriteString("```text\n")
	for _, s := range Statuses(a) {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")

	b.WriteString("## Verdict\n\n")
	b.WriteString(a.Truth + "\n\n")
	b.WriteString("**Next obligation:** " + a.Summary.NextObligation + "\n")
	return b.String()
}

func nearlyEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}
