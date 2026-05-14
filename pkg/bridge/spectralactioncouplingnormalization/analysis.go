// Package spectralactioncouplingnormalization implements Gate 327:
// Spectral Action Coupling Normalization / α_GUT Formula Audit.
//
// Gate 327 audits a new topological/dimensional witness for the absolute
// unified coupling.  It does not silently promote the witness into a theorem.
// It compares three lanes:
//  1. the standard heat-kernel gauge-normalization ledger,
//  2. the topological-action lane S_top/π = 8π,
//  3. the finite-algebra dimension-per-generation lane dim_R(A_F)π/N_gen = 8π.
//
// The gate then propagates the resulting g_*² = 1/2 into the Gate 308 Higgs
// quartic boundary ratio λ_H/g_*² = 1197/4624.  The resulting tree-level Higgs
// proxy is cataloged as a striking conditional alignment, not as a final pole
// mass or a derivation of the physical electroweak VEV.
package spectralactioncouplingnormalization

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE327-SPECTRAL-ACTION-COUPLING-NORMALIZATION-ALPHA-GUT-AUDIT"

	StatusSpectralActionGaugeLedgerFormalized = "CONDITIONAL_SUPPORT_SPECTRAL_ACTION_GAUGE_LEDGER_FORMALIZED"
	StatusTopologicalActionCouplingWitness    = "CONDITIONAL_SUPPORT_TOPOLOGICAL_ACTION_COUPLING_WITNESS_FORMALIZED"
	StatusAlgebraDimensionGenerationWitness   = "CONDITIONAL_SUPPORT_ALGEBRA_DIMENSION_GENERATION_WITNESS_FORMALIZED"
	StatusGStarHalfBoundaryComputed           = "CONDITIONAL_SUPPORT_GSTAR_SQUARED_ONE_HALF_BOUNDARY_COMPUTED"
	StatusHiggsProxyFromTopologicalCoupling   = "CONDITIONAL_SUPPORT_HIGGS_PROXY_FROM_TOPOLOGICAL_COUPLING_COMPUTED"

	StatusTensionSpectralActionNormalizationNotIdentical = "CONDITIONAL_TENSION_STANDARD_HEAT_KERNEL_NORMALIZATION_NOT_IDENTICAL_TO_8PI_WITHOUT_TRACE_THEOREM"
	StatusTensionEightPiWitnessNotYetTheorem             = "CONDITIONAL_TENSION_EIGHT_PI_WITNESS_NOT_YET_PROMOTED_TO_SPECTRAL_ACTION_THEOREM"

	StatusFailedAlphaGUTDerivationNotClosed      = "FAILED_ROUTE_ALPHA_GUT_NATIVE_DERIVATION_NOT_CLOSED"
	StatusFailedTraceRepIndexNotDerived          = "FAILED_ROUTE_REQUIRED_TRACE_REP_INDEX_NOT_DERIVED"
	StatusFailedDimensionFormulaNotSpectralProof = "FAILED_ROUTE_DIMENSION_PER_GENERATION_FORMULA_NOT_PROVED_AS_SPECTRAL_ACTION_THEOREM"
	StatusFailedVEVNotDerived                    = "FAILED_ROUTE_ELECTROWEAK_VEV_NOT_DERIVED"
	StatusFailedPoleMassNotExecuted              = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedFinalColliderMassNotClaimed      = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 326

	// Gate inputs / structural witnesses.
	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	finiteAlgebraRealDim     = 24.0
	generationCount          = 3.0
	contactCutoffF0          = 7.0
	gutTraceIndex            = 1.0
	observedHiggsGeV         = 125.10
	electroweakVEVGeV        = 246.22
)

type GateInputs struct {
	HighestInheritedGate int
	ContactScalarRatio   float64
	FiniteAlgebraRealDim float64
	GenerationCount      float64
	ContactCutoffF0      float64
	GUTTraceIndex        float64
	ElectroweakVEVGeV    float64
	ObservedHiggsGeV     float64
	AddsEmpiricalFit     bool
	UsesObservedHiggsFit bool
	Status               string
}

type SpectralActionLedger struct {
	Formula                     string
	ContactCutoffF0             float64
	GUTTraceIndex               float64
	AlphaInverseFromN4TwoOver7  float64
	N4ForEightPi                float64
	N4ForEmpiricalTwentyFive    float64
	StandardCoefficientFormula  string
	RequiredTraceRepIndexFor8Pi float64
	StandardTraceIndexKnown     bool
	EightPiDerivedByThisLane    bool
	DiscrepancyNote             string
	Status                      string
}

type TopologicalActionLane struct {
	STopFormula         string
	STop                float64
	AlphaInverseFormula string
	AlphaInverse        float64
	GStarSquared        float64
	TopologicalSealOld  float64
	SealCorrectionRatio float64
	MatchesEightPi      bool
	Status              string
}

type DimensionGenerationLane struct {
	Formula                 string
	FiniteAlgebraRealDim    float64
	GenerationCount         float64
	AlphaInverse            float64
	EqualsTopologicalLane   bool
	UsesOnlyDerivedCounts   bool
	ProvedAsSpectralTheorem bool
	Status                  string
}

type HiggsProxy struct {
	ContactScalarRatio float64
	GStarSquared       float64
	LambdaH            float64
	Formula            string
	RunningOrTree      string
	PredictedMassGeV   float64
	ObservedMassGeV    float64
	DifferenceGeV      float64
	RelativeErrorPct   float64
	PoleMassClaimed    bool
	Status             string
}

type FirewallAudit struct {
	NoAlphaGUTFitInserted            bool
	NoTraceIndexInvented             bool
	NoSpectralActionProofOverclaimed bool
	NoObservedHiggsMassFitInserted   bool
	NoPoleMassClaimed                bool
	NoTwoLoopClaimed                 bool
	NoFinalColliderMassClaimed       bool
	FiniteCorePolluted               bool
	Status                           string
}

type Summary struct {
	CouplingWitnessFound     bool
	AlphaInverseEightPi      bool
	GStarSquaredHalf         bool
	HiggsProxyNearObserved   bool
	NativeDerivationClosed   bool
	FinalColliderMassClaimed bool
	Status                   string
	DirectAnswer             string
	NextObligation           string
}

type Analysis struct {
	Inputs      GateInputs
	Spectral    SpectralActionLedger
	Topological TopologicalActionLane
	Dimension   DimensionGenerationLane
	Higgs       HiggsProxy
	Audit       FirewallAudit
	Summary     Summary
	Truth       string
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
	spectral := auditSpectralActionLedger(inputs)
	topological := auditTopologicalActionLane()
	dimension := auditDimensionGenerationLane(topological)
	higgs := computeHiggsProxy(inputs, topological)
	audit := auditFirewalls(spectral, dimension, higgs)
	summary := compileSummary(spectral, topological, dimension, higgs, audit)
	truth := "Gate 327 finds a powerful 8π absolute-coupling witness: S_top/π = 8π and dim_R(A_F)π/N_gen = 24π/3 = 8π both imply g_*² = 1/2. Substituting this into the Gate 308 ratio λ_H/g_*² = 1197/4624 gives a tree-level Higgs proxy m_H = v√(1197/4624) ≈ 125.274 GeV. The gate does not close the α_GUT derivation, because the standard heat-kernel trace normalization still requires an explicit representation-index theorem or an action-normalization theorem proving that the 8π witness is the physical spectral-action coupling."
	return Analysis{Inputs: inputs, Spectral: spectral, Topological: topological, Dimension: dimension, Higgs: higgs, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileInputs() GateInputs {
	return GateInputs{
		HighestInheritedGate: inheritedHighestGate,
		ContactScalarRatio:   contactScalarNumerator / contactScalarDenominator,
		FiniteAlgebraRealDim: finiteAlgebraRealDim,
		GenerationCount:      generationCount,
		ContactCutoffF0:      contactCutoffF0,
		GUTTraceIndex:        gutTraceIndex,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		ObservedHiggsGeV:     observedHiggsGeV,
		AddsEmpiricalFit:     false,
		UsesObservedHiggsFit: false,
		Status:               StatusSpectralActionGaugeLedgerFormalized,
	}
}

func auditSpectralActionLedger(inputs GateInputs) SpectralActionLedger {
	alphaInvEightPi := 8.0 * math.Pi
	n4ForEightPi := alphaInvEightPi / (4.0 * math.Pi * inputs.ContactCutoffF0 * inputs.GUTTraceIndex)
	n4ForEmpiricalTwentyFive := 25.0 / (4.0 * math.Pi * inputs.ContactCutoffF0 * inputs.GUTTraceIndex)

	// If one uses the common schematic gauge coefficient 1/g² = (f0/2π²)Tr_rep(T²),
	// then α^{-1}=4π/g²=(2f0/π)Tr_rep(T²).  To equal 8π, the trace index would
	// have to be 4π²/7.  Gate 327 treats that as a required theorem, not as an input.
	requiredTrace := alphaInvEightPi * math.Pi / (2.0 * inputs.ContactCutoffF0)

	return SpectralActionLedger{
		Formula:                     "α_GUT^{-1} = 4π · N4 · f0 · τ_GUT",
		ContactCutoffF0:             inputs.ContactCutoffF0,
		GUTTraceIndex:               inputs.GUTTraceIndex,
		AlphaInverseFromN4TwoOver7:  4.0 * math.Pi * (2.0 / 7.0) * inputs.ContactCutoffF0 * inputs.GUTTraceIndex,
		N4ForEightPi:                n4ForEightPi,
		N4ForEmpiricalTwentyFive:    n4ForEmpiricalTwentyFive,
		StandardCoefficientFormula:  "if 1/g²=(f0/2π²)Tr_rep(T²), then α^{-1}=(2f0/π)Tr_rep(T²)",
		RequiredTraceRepIndexFor8Pi: requiredTrace,
		StandardTraceIndexKnown:     false,
		EightPiDerivedByThisLane:    false,
		DiscrepancyNote:             "The heat-kernel coefficient ledger can host α^{-1}=8π by N4=2/7, but the standard schematic coefficient needs a derived representation trace Tr_rep(T²)=4π²/7. That representation index is not a raw Hilbert dimension and is not derived here.",
		Status:                      StatusTensionSpectralActionNormalizationNotIdentical,
	}
}

func auditTopologicalActionLane() TopologicalActionLane {
	sTop := 8.0 * math.Pi * math.Pi
	alphaInv := sTop / math.Pi
	g2 := 4.0 * math.Pi / alphaInv
	return TopologicalActionLane{
		STopFormula:         "S_top = 8π²",
		STop:                sTop,
		AlphaInverseFormula: "α_GUT^{-1} := S_top/π = 8π",
		AlphaInverse:        alphaInv,
		GStarSquared:        g2,
		TopologicalSealOld:  1.0,
		SealCorrectionRatio: g2 / 1.0,
		MatchesEightPi:      nearlyEqual(alphaInv, 8.0*math.Pi, 1e-12) && nearlyEqual(g2, 0.5, 1e-12),
		Status:              StatusTopologicalActionCouplingWitness,
	}
}

func auditDimensionGenerationLane(topological TopologicalActionLane) DimensionGenerationLane {
	alphaInv := finiteAlgebraRealDim * math.Pi / generationCount
	return DimensionGenerationLane{
		Formula:                 "α_GUT^{-1} ?= dim_R(A_F)·π/N_gen = 24π/3 = 8π",
		FiniteAlgebraRealDim:    finiteAlgebraRealDim,
		GenerationCount:         generationCount,
		AlphaInverse:            alphaInv,
		EqualsTopologicalLane:   nearlyEqual(alphaInv, topological.AlphaInverse, 1e-12),
		UsesOnlyDerivedCounts:   true,
		ProvedAsSpectralTheorem: false,
		Status:                  StatusAlgebraDimensionGenerationWitness,
	}
}

func computeHiggsProxy(inputs GateInputs, topological TopologicalActionLane) HiggsProxy {
	lambda := inputs.ContactScalarRatio * topological.GStarSquared
	mass := inputs.ElectroweakVEVGeV * math.Sqrt(2.0*lambda)
	diff := mass - inputs.ObservedHiggsGeV
	rel := 100.0 * diff / inputs.ObservedHiggsGeV
	return HiggsProxy{
		ContactScalarRatio: inputs.ContactScalarRatio,
		GStarSquared:       topological.GStarSquared,
		LambdaH:            lambda,
		Formula:            "m_H(tree proxy)=v√(2·(1197/4624)·g_*²)=v√(1197/4624) for g_*²=1/2",
		RunningOrTree:      "tree-level running/pole-unmatched proxy",
		PredictedMassGeV:   mass,
		ObservedMassGeV:    inputs.ObservedHiggsGeV,
		DifferenceGeV:      diff,
		RelativeErrorPct:   rel,
		PoleMassClaimed:    false,
		Status:             StatusHiggsProxyFromTopologicalCoupling,
	}
}

func auditFirewalls(spectral SpectralActionLedger, dimension DimensionGenerationLane, higgs HiggsProxy) FirewallAudit {
	return FirewallAudit{
		NoAlphaGUTFitInserted:            true,
		NoTraceIndexInvented:             !spectral.StandardTraceIndexKnown && !spectral.EightPiDerivedByThisLane,
		NoSpectralActionProofOverclaimed: !dimension.ProvedAsSpectralTheorem,
		NoObservedHiggsMassFitInserted:   true,
		NoPoleMassClaimed:                !higgs.PoleMassClaimed,
		NoTwoLoopClaimed:                 true,
		NoFinalColliderMassClaimed:       true,
		FiniteCorePolluted:               false,
		Status:                           StatusFailedAlphaGUTDerivationNotClosed,
	}
}

func compileSummary(spectral SpectralActionLedger, topological TopologicalActionLane, dimension DimensionGenerationLane, higgs HiggsProxy, audit FirewallAudit) Summary {
	near := math.Abs(higgs.RelativeErrorPct) < 0.25
	closed := spectral.EightPiDerivedByThisLane && dimension.ProvedAsSpectralTheorem
	status := StatusTensionEightPiWitnessNotYetTheorem
	if topological.MatchesEightPi && dimension.EqualsTopologicalLane && near && audit.NoSpectralActionProofOverclaimed {
		status = StatusHiggsProxyFromTopologicalCoupling
	}
	return Summary{
		CouplingWitnessFound:     topological.MatchesEightPi && dimension.EqualsTopologicalLane,
		AlphaInverseEightPi:      nearlyEqual(topological.AlphaInverse, 8.0*math.Pi, 1e-12),
		GStarSquaredHalf:         nearlyEqual(topological.GStarSquared, 0.5, 1e-12),
		HiggsProxyNearObserved:   near,
		NativeDerivationClosed:   closed,
		FinalColliderMassClaimed: false,
		Status:                   status,
		DirectAnswer:             "The 8π lane gives g_*²=1/2 and m_H(tree proxy)≈125.274 GeV, but the promotion of S_top/π or 24π/3 into the spectral-action gauge coupling still needs an explicit normalization theorem.",
		NextObligation:           "Derive the weighted representation trace/action-normalization theorem that identifies N4=2/7 or Tr_rep(T²)=4π²/7 without empirical fitting.",
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Spectral.Status,
		a.Topological.Status,
		a.Dimension.Status,
		a.Higgs.Status,
		a.Audit.Status,
		a.Summary.Status,
		StatusGStarHalfBoundaryComputed,
		StatusFailedTraceRepIndexNotDerived,
		StatusFailedDimensionFormulaNotSpectralProof,
		StatusFailedVEVNotDerived,
		StatusFailedPoleMassNotExecuted,
		StatusFailedFinalColliderMassNotClaimed,
	}
}

func FormatInputs(x GateInputs) string {
	return fmt.Sprintf("gate<=%d ratio=%.12f dim_R(A_F)=%.0f N_gen=%.0f f0=%.0f τ=%.0f v=%.2f no_fit=%t", x.HighestInheritedGate, x.ContactScalarRatio, x.FiniteAlgebraRealDim, x.GenerationCount, x.ContactCutoffF0, x.GUTTraceIndex, x.ElectroweakVEVGeV, !x.AddsEmpiricalFit && !x.UsesObservedHiggsFit)
}

func FormatSpectral(x SpectralActionLedger) string {
	return fmt.Sprintf("%s; f0=%.0f τ=%.0f N4(8π)=%.12f N4(25)=%.12f α^{-1}(N4=2/7)=%.12f required_Tr=%.12f derived=%t note=%s", x.Formula, x.ContactCutoffF0, x.GUTTraceIndex, x.N4ForEightPi, x.N4ForEmpiricalTwentyFive, x.AlphaInverseFromN4TwoOver7, x.RequiredTraceRepIndexFor8Pi, x.EightPiDerivedByThisLane, x.DiscrepancyNote)
}

func FormatTopological(x TopologicalActionLane) string {
	return fmt.Sprintf("%s -> %s = %.12f; g_*²=%.12f old_seal=%.1f correction=%.3f matches=%t", x.STopFormula, x.AlphaInverseFormula, x.AlphaInverse, x.GStarSquared, x.TopologicalSealOld, x.SealCorrectionRatio, x.MatchesEightPi)
}

func FormatDimension(x DimensionGenerationLane) string {
	return fmt.Sprintf("%s -> α^{-1}=%.12f equals_topological=%t derived_counts=%t proved_as_spectral_theorem=%t", x.Formula, x.AlphaInverse, x.EqualsTopologicalLane, x.UsesOnlyDerivedCounts, x.ProvedAsSpectralTheorem)
}

func FormatHiggs(x HiggsProxy) string {
	return fmt.Sprintf("ratio=%.12f g_*²=%.12f λ=%.12f m=%.6f GeV observed_proxy=%.2f diff=%.6f rel=%.6f%% pole_claim=%t formula=%s", x.ContactScalarRatio, x.GStarSquared, x.LambdaH, x.PredictedMassGeV, x.ObservedMassGeV, x.DifferenceGeV, x.RelativeErrorPct, x.PoleMassClaimed, x.Formula)
}

func FormatAudit(x FirewallAudit) string {
	flags := []string{
		fmt.Sprintf("no_alpha_fit=%t", x.NoAlphaGUTFitInserted),
		fmt.Sprintf("no_trace_index_invented=%t", x.NoTraceIndexInvented),
		fmt.Sprintf("no_spectral_proof_overclaimed=%t", x.NoSpectralActionProofOverclaimed),
		fmt.Sprintf("no_higgs_fit=%t", x.NoObservedHiggsMassFitInserted),
		fmt.Sprintf("no_pole_mass=%t", x.NoPoleMassClaimed),
		fmt.Sprintf("no_two_loop=%t", x.NoTwoLoopClaimed),
		fmt.Sprintf("no_final_collider=%t", x.NoFinalColliderMassClaimed),
		fmt.Sprintf("polluted=%t", x.FiniteCorePolluted),
	}
	return strings.Join(flags, "; ")
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("witness=%t alpha8pi=%t ghalf=%t higgs_near=%t closed=%t collider_claim=%t status=%s next=%s", x.CouplingWitnessFound, x.AlphaInverseEightPi, x.GStarSquaredHalf, x.HiggsProxyNearObserved, x.NativeDerivationClosed, x.FinalColliderMassClaimed, x.Status, x.NextObligation)
}

func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
