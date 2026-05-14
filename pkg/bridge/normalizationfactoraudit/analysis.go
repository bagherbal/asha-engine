// Package normalizationfactoraudit implements Gate 378:
// Complete Normalization Factor Audit / Product Spectral Action Convention Sieve.
//
// Gate 377 exposed that the M×F product action had coefficient arithmetic, but
// still had an Einstein-Hilbert normalization seal. Gate 378 audits the six
// proposed finite-to-continuum normalization factors and determines which are
// fixed, which are channel-specific, and whether their product actually closes
// the gravitational coefficient gap.
package normalizationfactoraudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE378-COMPLETE-NORMALIZATION-FACTOR-AUDIT"

	StatusHeatKernelVolumeComputed          = "CONDITIONAL_SUPPORT_HEAT_KERNEL_VOLUME_FACTOR_COMPUTED"
	StatusDiracA2CurvatureCorrected         = "CONDITIONAL_SUPPORT_DIRAC_A2_CURVATURE_FACTOR_CORRECTED"
	StatusRealityTraceAlternativesAudited   = "CONDITIONAL_SUPPORT_REALITY_TRACE_ALTERNATIVES_AUDITED"
	StatusF0SlotAudited                     = "CONDITIONAL_SUPPORT_F0_MOMENT_SLOT_AUDITED"
	StatusCutoffScaleAudited                = "CONDITIONAL_SUPPORT_CUTOFF_SCALE_IDENTIFICATION_AUDITED"
	StatusGaugeTraceConventionAudited       = "CONDITIONAL_SUPPORT_GAUGE_TRACE_CONVENTION_AUDITED"
	StatusChannelSeparatedLedgerConstructed = "CONDITIONAL_SUPPORT_CHANNEL_SEPARATED_NORMALIZATION_LEDGER_CONSTRUCTED"
	StatusEHGapQuantified                   = "CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_NORMALIZATION_GAP_QUANTIFIED"
	StatusSMGravityStructurePreserved       = "CONDITIONAL_SUPPORT_SM_GRAVITY_STRUCTURAL_ACTION_PRESERVED"

	StatusTensionNaiveProductInvalid       = "CONDITIONAL_TENSION_SIX_FACTOR_NAIVE_PRODUCT_IS_CHANNEL_MIXING"
	StatusTensionLichnerowiczNotStandalone = "CONDITIONAL_TENSION_LICHNEROWICZ_R_OVER_4_NOT_STANDALONE_A2_FACTOR"
	StatusTensionRealityFactorNotUniversal = "CONDITIONAL_TENSION_REALITY_HALF_FACTOR_IS_CHANNEL_CONVENTION_NOT_UNIVERSAL_ARITHMETIC"
	StatusTensionF0EqualsSevenUnresolved   = "CONDITIONAL_TENSION_F0_EQUALS_SEVEN_SPECTRAL_ACTION_MOMENT_SLOT_UNRESOLVED"
	StatusTensionCutoffPlanckMismatch      = "CONDITIONAL_TENSION_F2_CUTOFF_MOMENT_TOO_SMALL_FOR_CANONICAL_EH_WITH_CURRENT_TRACE"
	StatusTensionAbsoluteGaugeOpen         = "CONDITIONAL_TENSION_ABSOLUTE_GAUGE_COUPLING_REQUIRES_REPRESENTATION_TRACE_NORMALIZATION"

	StatusFailedNormalizationProductNotClosed        = "FAILED_ROUTE_SIX_NORMALIZATION_FACTORS_DO_NOT_CLOSE_EH_GAP"
	StatusFailedEHCanonicalCoefficientNotDerived     = "FAILED_ROUTE_CANONICAL_EINSTEIN_HILBERT_COEFFICIENT_NOT_DERIVED"
	StatusFailedF0SevenNotProvenAsTestFunctionMoment = "FAILED_ROUTE_F0_EQUALS_SEVEN_NOT_PROVEN_AS_SPECTRAL_ACTION_TEST_FUNCTION_MOMENT"
	StatusFailedAbsoluteGaugeCouplingNotClosed       = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLING_NORMALIZATION_NOT_CLOSED"
	StatusFailedFullNumericalTOENotClosed            = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_NOT_REACHED"
)

const (
	dimSpacetime       = 4
	trFDoubled         = 96.0
	trFRealityHalf     = 48.0
	f0Candidate        = 7.0
	f2LambdaOverMP2    = math.Pi / 64.0
	canonicalEHPerMP2  = 0.5
	alphaBranchInverse = 8.0 * math.Pi
)

type Factor struct {
	Name                   string
	MathematicalDefinition string
	ProposedValue          string
	CorrectedValue         string
	Numeric                float64
	Channel                string
	Status                 string
	Verdict                string
}

type EHChannel struct {
	TraceDimension          float64
	TraceLabel              string
	HeatKernelVolume        float64
	DiracA2Magnitude        float64
	F2LambdaOverMP2         float64
	CoefficientPerMP2       float64
	GapToCanonical          float64
	RequiredF2LambdaOverMP2 float64
	CurrentF2ShortBy        float64
	Formula                 string
	Status                  string
}

type GaugeChannel struct {
	F0Candidate                  float64
	TargetAlphaInverse           float64
	RequiredRepresentationTraceK float64
	Formula                      string
	Status                       string
}

type Audit struct {
	Executed                          bool
	Factors                           []Factor
	EHWithDoubledTrace                EHChannel
	EHWithRealityHalfTrace            EHChannel
	NaiveSixFactorProductUsingUserR4  float64
	NaiveSixFactorProductUsingDiracA2 float64
	SkeletonGapNormalization          float64
	Gauge                             GaugeChannel
	ChannelSeparated                  bool
	ClosesEH                          bool
	FullNumericalClosure              bool
	Verdict                           string
	Truth                             string
}

type Analysis struct{ Audit Audit }

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	heat := 1.0 / (16.0 * math.Pi * math.Pi)
	// Under the common Laplace-type convention P=-(g^{μν}∇_μ∇_ν+E),
	// Lichnerowicz gives E=-R/4 and a2_R=(E+R/6)=-R/12.  We use the
	// magnitude here because sign depends on the Euclidean action convention.
	diracA2 := 1.0 / 12.0
	realityHalf := 0.5

	doubled := buildEH("full doubled trace", trFDoubled, heat, diracA2, f2LambdaOverMP2)
	half := buildEH("J-reduced/effective half trace", trFRealityHalf, heat, diracA2, f2LambdaOverMP2)

	skeleton := 0.5 * trFDoubled * f2LambdaOverMP2
	skeletonGap := canonicalEHPerMP2 / skeleton

	naiveUser := heat * 0.25 * realityHalf * trFDoubled
	naiveDirac := heat * diracA2 * realityHalf * trFDoubled

	gaugeK := (alphaBranchInverse * 4.0 * math.Pi) / f0Candidate
	gauge := GaugeChannel{F0Candidate: f0Candidate, TargetAlphaInverse: alphaBranchInverse, RequiredRepresentationTraceK: gaugeK, Formula: "If 1/g²=f0·K/(16π²), then α⁻¹=f0·K/(4π); α⁻¹=8π requires K=32π²/f0.", Status: strings.Join([]string{StatusGaugeTraceConventionAudited, StatusTensionAbsoluteGaugeOpen, StatusFailedAbsoluteGaugeCouplingNotClosed}, ";")}

	factors := []Factor{
		{Name: "Factor 1 — heat-kernel volume", MathematicalDefinition: "For d=4, a_{2k}(P) carries (4π)^(-2).", ProposedValue: "1/(16π²)", CorrectedValue: "1/(16π²)", Numeric: heat, Channel: "all local heat-kernel channels", Status: StatusHeatKernelVolumeComputed, Verdict: "fixed"},
		{Name: "Factor 2 — Dirac curvature contribution", MathematicalDefinition: "Lichnerowicz D²=∇*∇+R/4 must be inserted into the a₂ formula, which also contains the universal R/6 term.", ProposedValue: "1/4", CorrectedValue: "|1/6 - 1/4| = 1/12 in the common Laplace-type convention; sign is convention-dependent", Numeric: diracA2, Channel: "Einstein-Hilbert a₂ channel only", Status: strings.Join([]string{StatusDiracA2CurvatureCorrected, StatusTensionLichnerowiczNotStandalone}, ";"), Verdict: "corrected; R/4 is not a standalone multiplier"},
		{Name: "Factor 3 — doubled-space reality factor", MathematicalDefinition: "Compare bosonic spectral trace over full H_F with a possible J-reduced physical trace.", ProposedValue: "1/2", CorrectedValue: "not universal; doubled=96 and half=48 are both audited", Numeric: realityHalf, Channel: "trace convention; may differ for bosonic and fermionic sectors", Status: strings.Join([]string{StatusRealityTraceAlternativesAudited, StatusTensionRealityFactorNotUniversal}, ";"), Verdict: "not a universal arithmetic closure factor"},
		{Name: "Factor 4 — f0 spectral-action moment slot", MathematicalDefinition: "In the spectral action f0 is f(0)/zeroth moment for a₄-type terms; ζ_contact(0)=7 is an eigenvalue-counting/topological value unless promoted to a test-function moment by theorem.", ProposedValue: "f0=7", CorrectedValue: "unresolved; valid as a candidate only under a cutoff-moment promotion theorem", Numeric: f0Candidate, Channel: "a₄ gauge/Higgs/curvature² channels; not Einstein-Hilbert a₂", Status: strings.Join([]string{StatusF0SlotAudited, StatusTensionF0EqualsSevenUnresolved, StatusFailedF0SevenNotProvenAsTestFunctionMoment}, ";"), Verdict: "critical slot distinction"},
		{Name: "Factor 5 — cutoff scale identification", MathematicalDefinition: "Einstein-Hilbert normalization uses f2Λ²/M_P², not f0. Current ASHA ledger gives f2(Λ/M_P)²=π/64.", ProposedValue: "Λ=M_P or Λ=Mbar_P alternatives", CorrectedValue: "current f2Λ²/M_P²=π/64; canonical EH would require π² with doubled trace or 2π² with half trace", Numeric: f2LambdaOverMP2, Channel: "Einstein-Hilbert a₂ channel", Status: strings.Join([]string{StatusCutoffScaleAudited, StatusTensionCutoffPlanckMismatch, StatusEHGapQuantified}, ";"), Verdict: "precise mismatch quantified"},
		{Name: "Factor 6 — gauge trace convention", MathematicalDefinition: "Gauge kinetic normalization uses representation trace Tr_rep(T_aT_b)=Kδ_ab, not only finite ratios.", ProposedValue: "representation trace fixes absolute coupling", CorrectedValue: fmt.Sprintf("to get α_branch⁻¹=8π with f0=7 requires K=32π²/7≈%.12g", gaugeK), Numeric: gaugeK, Channel: "a₄ gauge kinetic channel", Status: gauge.Status, Verdict: "relative ratios survive; absolute coupling still needs trace normalization theorem"},
	}

	verdict := strings.Join([]string{
		StatusChannelSeparatedLedgerConstructed,
		StatusSMGravityStructurePreserved,
		StatusTensionNaiveProductInvalid,
		StatusEHGapQuantified,
		StatusFailedNormalizationProductNotClosed,
		StatusFailedEHCanonicalCoefficientNotDerived,
		StatusFailedFullNumericalTOENotClosed,
	}, ";")

	truth := "Gate 378 validates the user's instinct that normalization bookkeeping is the next right move, but rejects the proposed single six-factor product. The six items are real audit targets, yet they belong to different heat-kernel coefficient channels. The Lichnerowicz R/4 term is not a standalone multiplier; after the a₂ formula it contributes a Dirac Einstein-Hilbert magnitude 1/12. With the current ASHA value f₂(Λ/M_P)²=π/64, the canonical Einstein-Hilbert coefficient M_P²/2 is not derived: the doubled trace is short by 64π, and the J-half trace is short by 128π. Therefore the bridge is sharper, but full numerical ToE closure still requires a spectral-action moment/Planck-normalization theorem, an f0 slot theorem, and an absolute gauge trace theorem."

	return Analysis{Audit: Audit{Executed: true, Factors: factors, EHWithDoubledTrace: doubled, EHWithRealityHalfTrace: half, NaiveSixFactorProductUsingUserR4: naiveUser, NaiveSixFactorProductUsingDiracA2: naiveDirac, SkeletonGapNormalization: skeletonGap, Gauge: gauge, ChannelSeparated: true, ClosesEH: false, FullNumericalClosure: false, Verdict: verdict, Truth: truth}}, nil
}

func buildEH(label string, tr, heat, diracA2, f2 float64) EHChannel {
	coeff := tr * heat * diracA2 * f2
	reqF2 := canonicalEHPerMP2 / (tr * heat * diracA2)
	shortBy := reqF2 / f2
	return EHChannel{TraceDimension: tr, TraceLabel: label, HeatKernelVolume: heat, DiracA2Magnitude: diracA2, F2LambdaOverMP2: f2, CoefficientPerMP2: coeff, GapToCanonical: canonicalEHPerMP2 / coeff, RequiredF2LambdaOverMP2: reqF2, CurrentF2ShortBy: shortBy, Formula: "|C_R|/M_P² = Tr_F(1) · (4π)^-2 · (1/12) · f₂(Λ/M_P)²", Status: strings.Join([]string{StatusEHGapQuantified, StatusTensionCutoffPlanckMismatch, StatusFailedEHCanonicalCoefficientNotDerived}, ";")}
}

func FormatFactor(f Factor) string {
	return fmt.Sprintf("%s: proposed=%s; corrected=%s; numeric=%.12g; channel=%s; status=%s; verdict=%s", f.Name, f.ProposedValue, f.CorrectedValue, f.Numeric, f.Channel, f.Status, f.Verdict)
}

func FormatEH(e EHChannel) string {
	return fmt.Sprintf("%s: %s = %.12g; canonical target=0.5; gap=%.12g; required f₂(Λ/M_P)²=%.12g; current short by %.12g; status=%s", e.TraceLabel, e.Formula, e.CoefficientPerMP2, e.GapToCanonical, e.RequiredF2LambdaOverMP2, e.CurrentF2ShortBy, e.Status)
}

func FormatAudit(a Audit) string {
	lines := []string{}
	for _, f := range a.Factors {
		lines = append(lines, FormatFactor(f))
	}
	lines = append(lines, FormatEH(a.EHWithDoubledTrace))
	lines = append(lines, FormatEH(a.EHWithRealityHalfTrace))
	lines = append(lines, fmt.Sprintf("naive user product without f₂ using R/4 and reality half: %.12g", a.NaiveSixFactorProductUsingUserR4))
	lines = append(lines, fmt.Sprintf("corrected Dirac-a₂ product without f₂ using reality half: %.12g", a.NaiveSixFactorProductUsingDiracA2))
	lines = append(lines, fmt.Sprintf("Gate377 prompt-skeleton normalization needed 2/(3π): %.12g", a.SkeletonGapNormalization))
	lines = append(lines, fmt.Sprintf("gauge trace audit: %s; K_required=%.12g; status=%s", a.Gauge.Formula, a.Gauge.RequiredRepresentationTraceK, a.Gauge.Status))
	return strings.Join(lines, "\n")
}

func Constants() map[string]float64 {
	a, _ := BuildDefault()
	audit := a.Audit
	return map[string]float64{
		"heat_kernel_volume_4d":                    audit.Factors[0].Numeric,
		"dirac_a2_R_magnitude":                     audit.Factors[1].Numeric,
		"reality_half_candidate":                   audit.Factors[2].Numeric,
		"f0_candidate":                             f0Candidate,
		"f2_Lambda_over_MP2":                       f2LambdaOverMP2,
		"EH_coeff_doubled_per_MP2":                 audit.EHWithDoubledTrace.CoefficientPerMP2,
		"EH_coeff_half_per_MP2":                    audit.EHWithRealityHalfTrace.CoefficientPerMP2,
		"EH_gap_doubled_to_canonical":              audit.EHWithDoubledTrace.GapToCanonical,
		"EH_gap_half_to_canonical":                 audit.EHWithRealityHalfTrace.GapToCanonical,
		"required_f2_doubled":                      audit.EHWithDoubledTrace.RequiredF2LambdaOverMP2,
		"required_f2_half":                         audit.EHWithRealityHalfTrace.RequiredF2LambdaOverMP2,
		"current_f2_short_by_doubled":              audit.EHWithDoubledTrace.CurrentF2ShortBy,
		"current_f2_short_by_half":                 audit.EHWithRealityHalfTrace.CurrentF2ShortBy,
		"naive_user_product_R4_half_trace_no_f2":   audit.NaiveSixFactorProductUsingUserR4,
		"corrected_dirac_product_half_trace_no_f2": audit.NaiveSixFactorProductUsingDiracA2,
		"skeleton_gap_normalization":               audit.SkeletonGapNormalization,
		"gauge_required_K_for_alpha_8pi_if_f0_7":   audit.Gauge.RequiredRepresentationTraceK,
	}
}
