// Package ccmspectralactionsubstitution implements Gate 379:
// CCM Spectral Action Direct Substitution / Complete Coefficient Ledger.
//
// Gate 378 audited generic heat-kernel normalization factors. Gate 379 stops
// recomputing the Standard-Model almost-commutative action from isolated local
// factors and instead substitutes the ASHA finite invariants directly into the
// published Chamseddine-Connes-Marcolli coefficient ledger for M × F.
package ccmspectralactionsubstitution

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE379-CCM-SPECTRAL-ACTION-DIRECT-SUBSTITUTION"

	StatusCCMFormulaInstalled              = "CONDITIONAL_SUPPORT_CCM_SPECTRAL_ACTION_FORMULA_INSTALLED"
	StatusProductGeometryRecognized        = "CONDITIONAL_SUPPORT_PRODUCT_GEOMETRY_MATCHES_CCM_CONTEXT"
	StatusEinsteinCoefficientRecomputed    = "CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_COEFFICIENT_RECOMPUTED_FROM_CCM"
	StatusCutoffMomentCorrectionComputed   = "CONDITIONAL_SUPPORT_CUTOFF_MOMENT_CORRECTION_COMPUTED"
	StatusHiggsKineticQuarticReadOff       = "CONDITIONAL_SUPPORT_HIGGS_KINETIC_AND_QUARTIC_READ_OFF_FROM_CCM"
	StatusGaugeCoefficientLedgerReadOff    = "CONDITIONAL_SUPPORT_GAUGE_COEFFICIENT_LEDGER_READ_OFF"
	StatusYukawaTraceSymbolsPreserved      = "CONDITIONAL_SUPPORT_YUKAWA_TRACE_SYMBOLS_PRESERVED"
	StatusLagrangianAssembled              = "CONDITIONAL_SUPPORT_CCM_LAGRANGIAN_ASSEMBLED"
	StatusPreviousGenericFormulaSuperseded = "CONDITIONAL_SUPPORT_PREVIOUS_GENERIC_EH_FORMULA_SUPERSEDED"

	StatusTensionPreviousF2Mismatch        = "CONDITIONAL_TENSION_PREVIOUS_F2_PI_OVER_64_MISMATCHES_CCM_EH_BY_8PI"
	StatusTensionF0MomentSlotUnresolved    = "CONDITIONAL_TENSION_F0_EQUALS_SEVEN_STILL_REQUIRES_TEST_FUNCTION_MOMENT_PROOF"
	StatusTensionCTraceNotNumericallyFixed = "CONDITIONAL_TENSION_C_TRACE_TR_DF2_NOT_NUMERIC_WITHOUT_YUKAWA_SCALE_SEAL"
	StatusTensionHiggsQuarticChanged       = "CONDITIONAL_TENSION_HIGGS_QUARTIC_NORMALIZATION_DIFFERS_FROM_PREVIOUS_RATIO"
	StatusTensionGaugeAbsoluteOpen         = "CONDITIONAL_TENSION_ABSOLUTE_GAUGE_NORMALIZATION_NEEDS_REPRESENTATION_TRACE_LEDGER"
	StatusTensionCosmologicalMomentOpen    = "CONDITIONAL_TENSION_F4_VACUUM_SUBTRACTION_RULE_OPEN"

	StatusFailedPreviousF2NotCanonical    = "FAILED_ROUTE_PREVIOUS_F2_PI_OVER_64_NOT_CANONICAL_UNDER_CCM_EH_COEFFICIENT"
	StatusFailedFullNumericalTOENotClosed = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED"
	StatusFailedHiggsMassNotPredicted     = "FAILED_ROUTE_HIGGS_MASS_NOT_PREDICTED_WITHOUT_MU_VEV_RG_SEALS"
	StatusFailedCosmologicalConstantOpen  = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED_FROM_CCM_SUBSTITUTION"
)

const (
	TrFOne              = 96.0
	F0Candidate         = 7.0
	PreviousF2LambdaMP2 = math.Pi / 64.0
	LambdaTraceRatio    = 1197.0 / 4624.0 // e/a² in the ASHA quartic trace ledger.
	CanonicalEH         = 0.5
)

type CCMFormula struct {
	Cosmological    string
	EinsteinHilbert string
	HiggsKinetic    string
	HiggsMass       string
	HiggsQuartic    string
	Gauge           string
	Yukawa          string
	Source          string
}

type EinsteinLedger struct {
	Formula                           string
	CoefficientWithPreviousF2NoC      float64
	CanonicalCoefficient              float64
	GapFactor                         float64
	RequiredF2LambdaMP2Leading        float64
	CorrectionFromCOverMP2Coefficient float64
	RequiredFormula                   string
	Verdict                           string
}

type HiggsLedger struct {
	KineticCoefficient                      string
	CanonicalFieldRescaling                 string
	EOverA2                                 float64
	PreviousAshaRatio                       float64
	QuarticNoOuterPiConvention              float64
	QuarticCanonicalOuterPiConvention       float64
	MassParameterNoOuterPiConvention        string
	MassParameterCanonicalOuterPiConvention string
	Verdict                                 string
}

type GaugeLedger struct {
	Formula                     string
	RequiresRepresentationTrace bool
	AbsoluteClosed              bool
	Verdict                     string
}

type CosmologicalLedger struct {
	Formula                string
	NeedsF4                bool
	NeedsVacuumSubtraction bool
	Verdict                string
}

type Calculation struct {
	Executed                bool
	Formula                 CCMFormula
	Einstein                EinsteinLedger
	Higgs                   HiggsLedger
	Gauge                   GaugeLedger
	Cosmological            CosmologicalLedger
	Lagrangian              string
	Statuses                []string
	StructuralClosure       bool
	FullNumericalTOEClosure bool
	Truth                   string
}

type Analysis struct{ Calculation Calculation }

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	formula := CCMFormula{
		Cosmological:    "π⁻²[48 f₄Λ⁴ - f₂Λ² c + (f₀/4)d]",
		EinsteinHilbert: "π⁻²[(96 f₂Λ² - f₀ c)/24] R",
		HiggsKinetic:    "π⁻²[f₀ a |D_μ φ|²]",
		HiggsMass:       "π⁻²[-(f₂Λ² a/2)|φ|²]",
		HiggsQuartic:    "π⁻²[(f₀/2)e |φ|⁴]",
		Gauge:           "CCM representation-trace gauge kinetic channels; absolute coupling requires the representation trace normalization.",
		Yukawa:          "fermionic action ψ̄(D_M⊗1 + γ₅⊗D_F)ψ; the 13 charged finite-Dirac moduli remain in D_F.",
		Source:          "Chamseddine-Connes-Marcolli almost-commutative spectral action coefficient ledger.",
	}

	requiredLeading := math.Pi * math.Pi / 8.0
	coeffPreviousNoC := 4.0 * PreviousF2LambdaMP2 / (math.Pi * math.Pi)
	gap := CanonicalEH / coeffPreviousNoC
	einstein := EinsteinLedger{
		Formula:                           "C_R/M_P² = (96 F₂ - f₀ c/M_P²)/(24π²), where F₂=f₂Λ²/M_P². Canonical EH requires C_R/M_P²=1/2.",
		CoefficientWithPreviousF2NoC:      coeffPreviousNoC,
		CanonicalCoefficient:              CanonicalEH,
		GapFactor:                         gap,
		RequiredF2LambdaMP2Leading:        requiredLeading,
		CorrectionFromCOverMP2Coefficient: F0Candidate / 96.0,
		RequiredFormula:                   "F₂_required = π²/8 + (f₀/96)(c/M_P²). In the leading c≪M_P² approximation, F₂_required=π²/8.",
		Verdict:                           fmt.Sprintf("Previous F₂=π/64 gives C_R/M_P²=%.12g, short by %.12g=8π. CCM leading canonical value is π²/8=%.12g.", coeffPreviousNoC, gap, requiredLeading),
	}

	higgsNoPi := LambdaTraceRatio / (2.0 * F0Candidate)
	higgsWithPi := math.Pi * math.Pi * LambdaTraceRatio / (2.0 * F0Candidate)
	higgs := HiggsLedger{
		KineticCoefficient:                      "K_φ=f₀a/π²",
		CanonicalFieldRescaling:                 "If |D H|² is canonical, H=(√(f₀a)/π)φ under the literal outer-π² CCM convention.",
		EOverA2:                                 LambdaTraceRatio,
		PreviousAshaRatio:                       LambdaTraceRatio,
		QuarticNoOuterPiConvention:              higgsNoPi,
		QuarticCanonicalOuterPiConvention:       higgsWithPi,
		MassParameterNoOuterPiConvention:        "μ² = f₂Λ²/(2f₀) or f₂Λ²/f₀ depending on whether the Higgs potential is written with the CCM |φ|²/2 convention or the SM |H|² convention.",
		MassParameterCanonicalOuterPiConvention: "Under H=(√(f₀a)/π)φ, the quadratic coefficient is f₂Λ²/(2f₀) before SM potential convention reconciliation.",
		Verdict:                                 "The old λ_H/g_*²=e/a²=1197/4624 is a finite trace ratio, not automatically the canonically normalized quartic. CCM read-off introduces f₀ and possibly π² through field normalization.",
	}

	gauge := GaugeLedger{
		Formula:                     "1/g_i² is read from the CCM gauge kinetic representation trace after normalizing Tr(F_{μν}F^{μν}) to the canonical Yang-Mills convention.",
		RequiresRepresentationTrace: true,
		AbsoluteClosed:              false,
		Verdict:                     "Relative normalization such as sin²θ_W=3/8 can remain closed; absolute α_GUT requires the exact CCM representation trace convention and f₀ moment slot.",
	}

	cosmo := CosmologicalLedger{
		Formula:                "ρ_vac channel = π⁻²[48 f₄Λ⁴ - f₂Λ² c + (f₀/4)d].",
		NeedsF4:                true,
		NeedsVacuumSubtraction: true,
		Verdict:                "The cosmological channel is assembled symbolically but Λ_cosmo is not predicted without f₄ and a vacuum subtraction/renormalization theorem.",
	}

	lagrangian := strings.Join([]string{
		"S_CCM-ASHA[M×F] = ∫_M d⁴x√g / π² {",
		"  48 f₄Λ⁴ - f₂Λ² c + (f₀/4)d",
		"  + [(96 f₂Λ² - f₀c)/24] R",
		"  + f₀a |D_μ φ|² - (f₂Λ²a/2)|φ|² + (f₀e/2)|φ|⁴",
		"  + gauge kinetic representation-trace terms",
		"} + fermionic action ψ̄(D_M⊗1 + γ₅⊗D_F)ψ.",
		"ASHA substitutions: Tr_F(1)=96, f₀ candidate=7, e/a²=1197/4624, dim M_charged(D_F)=13.",
	}, "\n")

	statuses := []string{
		StatusCCMFormulaInstalled,
		StatusProductGeometryRecognized,
		StatusEinsteinCoefficientRecomputed,
		StatusCutoffMomentCorrectionComputed,
		StatusHiggsKineticQuarticReadOff,
		StatusGaugeCoefficientLedgerReadOff,
		StatusYukawaTraceSymbolsPreserved,
		StatusLagrangianAssembled,
		StatusPreviousGenericFormulaSuperseded,
		StatusTensionPreviousF2Mismatch,
		StatusTensionF0MomentSlotUnresolved,
		StatusTensionCTraceNotNumericallyFixed,
		StatusTensionHiggsQuarticChanged,
		StatusTensionGaugeAbsoluteOpen,
		StatusTensionCosmologicalMomentOpen,
		StatusFailedPreviousF2NotCanonical,
		StatusFailedFullNumericalTOENotClosed,
		StatusFailedHiggsMassNotPredicted,
		StatusFailedCosmologicalConstantOpen,
	}

	truth := "Gate 379 confirms the user's correction: the CCM almost-commutative coefficient ledger supersedes the Gate-378 generic Einstein-channel arithmetic. The direct substitution shifts the canonical leading cutoff moment from π/64 to π²/8, an exact 8π mismatch. It also downgrades 1197/4624 from an already-normalized Higgs quartic to a finite trace ratio that must pass through CCM kinetic normalization. Structural SM+gravity closure is strengthened; full numerical ToE closure still requires the f₀ moment theorem, c/a/e numeric seals, gauge trace normalization, f₄/vacuum subtraction, and RG/matching prescriptions."

	return Analysis{Calculation: Calculation{Executed: true, Formula: formula, Einstein: einstein, Higgs: higgs, Gauge: gauge, Cosmological: cosmo, Lagrangian: lagrangian, Statuses: statuses, StructuralClosure: true, FullNumericalTOEClosure: false, Truth: truth}}, nil
}

func NativeConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	c := a.Calculation
	return map[string]float64{
		"previous_f2_lambda_over_mp2":             PreviousF2LambdaMP2,
		"ccm_required_f2_lambda_over_mp2_leading": c.Einstein.RequiredF2LambdaMP2Leading,
		"ccm_gap_factor":                          c.Einstein.GapFactor,
		"e_over_a2":                               c.Higgs.EOverA2,
		"lambda_no_outer_pi_convention":           c.Higgs.QuarticNoOuterPiConvention,
		"lambda_canonical_outer_pi_convention":    c.Higgs.QuarticCanonicalOuterPiConvention,
	}
}

func FormatFloat(x float64) string { return fmt.Sprintf("%.15g", x) }

func StatusLine(c Calculation) string { return strings.Join(c.Statuses, ";") }
