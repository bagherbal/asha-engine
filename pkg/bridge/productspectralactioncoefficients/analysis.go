// Package productspectralactioncoefficients implements Gate 377:
// Product Spectral Action Coefficient Calculator / Almost-Commutative Closure Audit.
//
// Gate 376 assembled the M×F product geometry and identified the SM+gravity
// sectors. Gate 377 is stricter: it performs the coefficient arithmetic that
// Gate 376 did not perform. It substitutes the ASHA finite invariants into a
// declared heat-kernel convention, reads off normalized action coefficients,
// and reports which sectors are actually fixed and which still require an
// empirical/conventional seal.
package productspectralactioncoefficients

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE377-PRODUCT-SPECTRAL-ACTION-COEFFICIENT-CALCULATOR"

	StatusGate376Inspected                       = "CONDITIONAL_SUPPORT_GATE376_PRODUCT_TRIPLE_INHERITED"
	StatusCoefficientArithmeticExecuted          = "CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_ACTION_COEFFICIENT_ARITHMETIC_EXECUTED"
	StatusHeatKernelConventionDeclared           = "CONDITIONAL_SUPPORT_HEAT_KERNEL_CONVENTION_DECLARED"
	StatusA0CoefficientComputed                  = "CONDITIONAL_SUPPORT_A0_COSMOLOGICAL_CHANNEL_COEFFICIENT_COMPUTED"
	StatusA2CoefficientComputed                  = "CONDITIONAL_SUPPORT_A2_EINSTEIN_HILBERT_CHANNEL_COEFFICIENT_COMPUTED"
	StatusA4GaugeCoefficientComputed             = "CONDITIONAL_SUPPORT_A4_GAUGE_KINETIC_CHANNEL_COMPUTED"
	StatusA4HiggsCoefficientComputed             = "CONDITIONAL_SUPPORT_A4_HIGGS_CHANNEL_COMPUTED"
	StatusFermionicYukawaTermComputed            = "CONDITIONAL_SUPPORT_FERMIONIC_YUKAWA_TERM_READ_OFF"
	StatusSMGravityLagrangianStructural          = "CONDITIONAL_SUPPORT_SM_GRAVITY_LAGRANGIAN_STRUCTURALLY_RECOVERED"
	StatusFiniteCoefficientsSubstituted          = "CONDITIONAL_SUPPORT_ASHA_FINITE_COEFFICIENTS_SUBSTITUTED"
	StatusRelativeGaugeRatiosFixed               = "CONDITIONAL_SUPPORT_RELATIVE_GAUGE_RATIOS_FIXED"
	StatusQuarticRatioFixed                      = "CONDITIONAL_SUPPORT_HIGGS_QUARTIC_RATIO_FIXED"
	StatusYukawaModuliPreserved                  = "CONDITIONAL_SUPPORT_13_YUKAWA_CKM_MODULI_PRESERVED"
	StatusTensionGate376WasFormalNotArithmetic   = "CONDITIONAL_TENSION_GATE376_WAS_FORMAL_ASSEMBLY_NOT_FULL_COEFFICIENT_CALCULATION"
	StatusTensionF4MomentMissing                 = "CONDITIONAL_TENSION_F4_MOMENT_OR_VACUUM_SUBTRACTION_MISSING"
	StatusTensionEHNormalizationNotClosed        = "CONDITIONAL_TENSION_EINSTEIN_HILBERT_NORMALIZATION_REQUIRES_CONVENTION_OR_TRACE_RENORMALIZATION"
	StatusTensionAbsoluteGaugeCouplingNotClosed  = "CONDITIONAL_TENSION_ABSOLUTE_GAUGE_COUPLING_NORMALIZATION_REQUIRES_RUNNING_AND_CONVENTION"
	StatusTensionYukawaTextureNotSelected        = "CONDITIONAL_TENSION_YUKAWA_TEXTURE_NOT_SELECTED"
	StatusTensionMassParameterNotFixed           = "CONDITIONAL_TENSION_HIGGS_MASS_PARAMETER_REQUIRES_TR_DF2_AND_VACUUM_CHOICE"
	StatusFailedFullTOEClosureNotReached         = "FAILED_ROUTE_FULL_THEORY_OF_EVERYTHING_NUMERICAL_CLOSURE_NOT_REACHED"
	StatusFailedAllCoefficientsNotDetermined     = "FAILED_ROUTE_ALL_LAGRANGIAN_COEFFICIENTS_NOT_DETERMINED_BY_ASHA_FINITE_DATA"
	StatusFailedCosmologicalConstantNotPredicted = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_PREDICTED"
	StatusFailedYukawaSectorNotPredicted         = "FAILED_ROUTE_YUKAWA_SECTOR_REMAINS_13_MODULI"
)

const (
	finiteHilbertDimDoubled = 96.0
	f0Contact               = 7.0
	f2LambdaOverMP2         = math.Pi / 64.0
	lambdaHOverGStarSq      = 1197.0 / 4624.0
	sin2ThetaWBoundary      = 3.0 / 8.0
	alphaBranchInverse      = 8.0 * math.Pi
	chargedModuli           = 13
	thresholdDeltaLambda    = -0.0978
)

type ProductTriple struct {
	Algebra, HilbertSpace, Dirac, Real, Grading string
	Valid                                       bool
}
type HeatKernelConvention struct {
	Name, Expansion, A0Density, A2DiracRDensity, A4Channels, Verdict string
	Dimension                                                        int
	IncludesRaw16Pi2, DiracA2RSignKnown                              bool
	Notes                                                            []string
}
type FiniteData struct {
	TrOne, F0, F2LambdaOverMP2, LambdaHOverGStarSq, Sin2ThetaW, AlphaBranchInverse, ThresholdDeltaLambda float64
	ChargedModuli                                                                                        int
}
type Coefficient struct {
	Sector, Formula, Value, Missing, Status string
	Numeric                                 float64
	DeterminedByASHA, FullyPhysical         bool
}
type Calculation struct {
	Executed                                                                             bool
	Product                                                                              ProductTriple
	Convention                                                                           HeatKernelConvention
	Finite                                                                               FiniteData
	A0CosmologicalPrefactorPerF4Lambda4                                                  Coefficient
	A2RawEinsteinCoefficientPerMP2                                                       Coefficient
	A2SkeletonEinsteinCoefficientPerMP2                                                  Coefficient
	A2NormalizationNeededToMatchMP                                                       float64
	GaugeSin2ThetaW, GaugeAlphaBranch, HiggsQuarticRatio, HiggsThresholdJump, YukawaTerm Coefficient
	Lagrangian                                                                           string
	AllCoefficientsDetermined, StandardModelGravityStructural, HardTOEClosure            bool
	Verdict                                                                              string
}
type Analysis struct {
	Calculation Calculation
	Truth       string
}

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	product := ProductTriple{Algebra: "C∞(M) ⊗ (C ⊕ H ⊕ M₃(C))", HilbertSpace: "L²(M,S) ⊗ H_F", Dirac: "D_total = D_M ⊗ 1_F + γ₅ ⊗ D_F", Real: "J_total = J_M ⊗ J_F", Grading: "γ_total = γ_M ⊗ γ_F", Valid: true}
	conv := HeatKernelConvention{Name: "declared 4D product heat-kernel audit convention", Dimension: 4, Expansion: "Tr f(D²/Λ²) ≃ f₄Λ⁴a₀(D²)+f₂Λ²a₂(D²)+f₀a₄(D²)+…", A0Density: "a₀=(4π)⁻²∫√g Tr(1)", A2DiracRDensity: "raw Dirac channel magnitude |a₂_R|=(4π)⁻²∫√g Tr(1)·R/12; sign depends on Laplace convention", A4Channels: "a₄ contains gauge curvature, Higgs kinetic/potential, and curvature² channels", IncludesRaw16Pi2: true, DiracA2RSignKnown: false, Notes: []string{"This gate computes coefficient channels, not just names sectors.", "The raw heat-kernel coefficient and the prompt's skeleton coefficient are both reported because the project has not fixed one universal gravitational trace-renormalization convention."}, Verdict: strings.Join([]string{StatusHeatKernelConventionDeclared, StatusTensionEHNormalizationNotClosed}, ";")}
	finite := FiniteData{TrOne: finiteHilbertDimDoubled, F0: f0Contact, F2LambdaOverMP2: f2LambdaOverMP2, LambdaHOverGStarSq: lambdaHOverGStarSq, Sin2ThetaW: sin2ThetaWBoundary, AlphaBranchInverse: alphaBranchInverse, ChargedModuli: chargedModuli, ThresholdDeltaLambda: thresholdDeltaLambda}
	a0 := finite.TrOne / (16 * math.Pi * math.Pi)
	a2raw := finite.TrOne * finite.F2LambdaOverMP2 / (192 * math.Pi * math.Pi)
	a2skel := 0.5 * finite.TrOne * finite.F2LambdaOverMP2
	norm := 0.5 / a2skel
	calc := Calculation{Executed: true, Product: product, Convention: conv, Finite: finite,
		A0CosmologicalPrefactorPerF4Lambda4: Coefficient{Sector: "cosmological/vacuum", Formula: "C_Λ/(f₄Λ⁴)=Tr_F(1)/(16π²)", Value: fmt.Sprintf("%.12g", a0), Numeric: a0, DeterminedByASHA: true, FullyPhysical: false, Missing: "f₄Λ⁴ moment and vacuum subtraction/renormalization condition", Status: strings.Join([]string{StatusA0CoefficientComputed, StatusTensionF4MomentMissing, StatusFailedCosmologicalConstantNotPredicted}, ";")},
		A2RawEinsteinCoefficientPerMP2:      Coefficient{Sector: "Einstein-Hilbert raw heat-kernel channel", Formula: "|C_R|/M_P²=Tr_F(1)·[f₂(Λ/M_P)²]/(192π²)", Value: fmt.Sprintf("%.12g", a2raw), Numeric: a2raw, DeterminedByASHA: true, FullyPhysical: false, Missing: "sign convention, trace normalization, and matching to M_P²/2", Status: strings.Join([]string{StatusA2CoefficientComputed, StatusTensionEHNormalizationNotClosed}, ";")},
		A2SkeletonEinsteinCoefficientPerMP2: Coefficient{Sector: "Einstein-Hilbert prompt-skeleton channel", Formula: "C_R/M_P²=(1/2)·Tr_F(1)·f₂(Λ/M_P)²", Value: fmt.Sprintf("%.12g", a2skel), Numeric: a2skel, DeterminedByASHA: true, FullyPhysical: false, Missing: "normalization factor required to equal canonical M_P²/2", Status: strings.Join([]string{StatusA2CoefficientComputed, StatusTensionEHNormalizationNotClosed}, ";")},
		A2NormalizationNeededToMatchMP:      norm,
		GaugeSin2ThetaW:                     Coefficient{Sector: "relative gauge normalization", Formula: "sin²θ_W(Λ)=3/8", Value: "3/8", Numeric: finite.Sin2ThetaW, DeterminedByASHA: true, FullyPhysical: true, Missing: "IR RG transport for low-energy comparison", Status: strings.Join([]string{StatusA4GaugeCoefficientComputed, StatusRelativeGaugeRatiosFixed}, ";")},
		GaugeAlphaBranch:                    Coefficient{Sector: "absolute gauge branch", Formula: "α_branch⁻¹=8π", Value: fmt.Sprintf("%.12g", finite.AlphaBranchInverse), Numeric: finite.AlphaBranchInverse, DeterminedByASHA: true, FullyPhysical: false, Missing: "threshold transport and absolute convention audit", Status: strings.Join([]string{StatusA4GaugeCoefficientComputed, StatusTensionAbsoluteGaugeCouplingNotClosed}, ";")},
		HiggsQuarticRatio:                   Coefficient{Sector: "Higgs quartic", Formula: "λ_H/g_*²=1197/4624", Value: "1197/4624", Numeric: finite.LambdaHOverGStarSq, DeterminedByASHA: true, FullyPhysical: false, Missing: "RG/matching to pole Higgs mass and μ²/vacuum choice", Status: strings.Join([]string{StatusA4HiggsCoefficientComputed, StatusQuarticRatioFixed, StatusTensionMassParameterNotFixed}, ";")},
		HiggsThresholdJump:                  Coefficient{Sector: "heavy-sector threshold", Formula: "Δλ≈-0.0978", Value: "-0.0978", Numeric: finite.ThresholdDeltaLambda, DeterminedByASHA: true, FullyPhysical: false, Missing: "continuum matching prescription", Status: StatusA4HiggsCoefficientComputed},
		YukawaTerm:                          Coefficient{Sector: "fermionic/Yukawa", Formula: "∫√g ψ̄(D_M⊗1+γ₅⊗D_F)ψ", Value: "13 charged finite-Dirac moduli", Numeric: float64(finite.ChargedModuli), DeterminedByASHA: true, FullyPhysical: false, Missing: "9 charged masses + 4 CKM coordinates", Status: strings.Join([]string{StatusFermionicYukawaTermComputed, StatusYukawaModuliPreserved, StatusTensionYukawaTextureNotSelected, StatusFailedYukawaSectorNotPredicted}, ";")},
		StandardModelGravityStructural:      true, AllCoefficientsDetermined: false, HardTOEClosure: false}
	calc.Lagrangian = buildLagrangian(calc)
	calc.Verdict = strings.Join([]string{StatusGate376Inspected, StatusTensionGate376WasFormalNotArithmetic, StatusCoefficientArithmeticExecuted, StatusFiniteCoefficientsSubstituted, StatusSMGravityLagrangianStructural, StatusFailedAllCoefficientsNotDetermined, StatusFailedFullTOEClosureNotReached}, ";")
	truth := "Gate 377 confirms the criticism of Gate 376: the previous gate assembled the product-action ledger but did not perform a complete coefficient calculation. This gate substitutes ASHA finite constants into an explicit heat-kernel convention and reads off coefficient channels. The SM+Einstein-gravity structure is recovered, sin²θW and the Higgs quartic ratio are fixed as finite ratios, and the 13 Yukawa/CKM moduli are preserved. Full Theory-of-Everything numerical closure is not reached: the cosmological f₄/vacuum subtraction channel, gravitational normalization convention, absolute low-energy running/matching, Higgs μ²/vacuum choice, and the 13 flavor coordinates remain open or sealed inputs."
	return Analysis{Calculation: calc, Truth: truth}, nil
}

func buildLagrangian(c Calculation) string {
	return strings.Join([]string{"S[M×F] = ∫_M d⁴x √g {", fmt.Sprintf("  C_R R                                  with raw |C_R|/M_P² = %.12g; skeleton C_R/M_P² = %.12g", c.A2RawEinsteinCoefficientPerMP2.Numeric, c.A2SkeletonEinsteinCoefficientPerMP2.Numeric), fmt.Sprintf("  - ρ_vac                                with C_Λ/(f₄Λ⁴) = %.12g, but f₄/subtraction open", c.A0CosmologicalPrefactorPerF4Lambda4.Numeric), "  + Σ_i (1/4g_i²) Tr(F_i μν F_i^μν)      with sin²θ_W(Λ)=3/8 and α_branch⁻¹=8π channel", "  + Z_H |∇H|²", "  - μ_H² |H|² + λ_H |H|⁴                with λ_H/g_*²=1197/4624 and Δλ≈-0.0978 threshold channel", "  + ψ̄ iD_M ψ + ψ̄ γ₅D_F ψ              with 13 charged Yukawa/CKM moduli", "  + curvature² spectral-action terms", "}"}, "\n")
}

func NativeCoefficientConstants() map[string]float64 {
	a0 := finiteHilbertDimDoubled / (16 * math.Pi * math.Pi)
	a2raw := finiteHilbertDimDoubled * f2LambdaOverMP2 / (192 * math.Pi * math.Pi)
	a2skel := 0.5 * finiteHilbertDimDoubled * f2LambdaOverMP2
	return map[string]float64{"Tr_F_1": finiteHilbertDimDoubled, "f0": f0Contact, "f2_Lambda_over_MP2": f2LambdaOverMP2, "a0_prefactor_per_f4_Lambda4": a0, "a2_raw_EH_coeff_per_MP2": a2raw, "a2_skeleton_EH_coeff_per_MP2": a2skel, "EH_skeleton_norm_needed_to_match_MP2_over_2": 0.5 / a2skel, "sin2_thetaW": sin2ThetaWBoundary, "alpha_branch_inverse": alphaBranchInverse, "lambdaH_over_gstar2": lambdaHOverGStarSq, "threshold_delta_lambda": thresholdDeltaLambda, "charged_moduli": chargedModuli}
}
func FormatCoefficient(x Coefficient) string {
	return fmt.Sprintf("%s: %s = %s; ASHA=%t; physical=%t; missing=%s; status=%s", x.Sector, x.Formula, x.Value, x.DeterminedByASHA, x.FullyPhysical, x.Missing, x.Status)
}
func FormatCalculation(c Calculation) string {
	parts := []string{FormatCoefficient(c.A0CosmologicalPrefactorPerF4Lambda4), FormatCoefficient(c.A2RawEinsteinCoefficientPerMP2), FormatCoefficient(c.A2SkeletonEinsteinCoefficientPerMP2), fmt.Sprintf("EH skeleton normalization needed for canonical M_P²/2: %.12g", c.A2NormalizationNeededToMatchMP), FormatCoefficient(c.GaugeSin2ThetaW), FormatCoefficient(c.GaugeAlphaBranch), FormatCoefficient(c.HiggsQuarticRatio), FormatCoefficient(c.HiggsThresholdJump), FormatCoefficient(c.YukawaTerm)}
	return strings.Join(parts, "\n")
}
