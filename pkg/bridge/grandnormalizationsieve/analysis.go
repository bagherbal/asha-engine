// Package grandnormalizationsieve implements Gate 300:
// Grand Normalization Sieve / Wave-Function Renormalization Extraction Audit.
//
// Gate 299 formalized the Seeley-de Witt heat-kernel coefficient channels and
// deliberately stopped before physical dynamics. Gate 300 advances exactly one
// epistemic layer: it formalizes the algebraic normalization algorithm that must
// be applied to raw a2/a4 trace channels before any Higgs mass, quartic, or
// gauge-coupling statement can be physical. It does not insert empirical
// Yukawa data, cutoff moments, subtraction constants, or B-gap instanton input.
package grandnormalizationsieve

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE300-GRAND-NORMALIZATION-SIEVE-WAVE-FUNCTION-RENORMALIZATION-EXTRACTION-AUDIT"

	StatusGate299Inherited           = "CONDITIONAL_SUPPORT_GATE299_HEAT_KERNEL_CHANNELS_INHERITED"
	StatusKineticIsolationFormalized = "CONDITIONAL_SUPPORT_KINETIC_ISOLATION_ALGORITHM_FORMALIZED"
	StatusZHFormalized               = "CONDITIONAL_SUPPORT_WAVE_FUNCTION_RENORMALIZATION_ZH_FORMALIZED"
	StatusRescalingMapFormalized     = "CONDITIONAL_SUPPORT_MASS_QUARTIC_RESCALING_MAP_FORMALIZED"
	StatusGaugeNormFormalized        = "CONDITIONAL_SUPPORT_GAUGE_KINETIC_NORMALIZATION_MAP_FORMALIZED"
	StatusAlgorithmFormalized        = "CONDITIONAL_SUPPORT_KINETIC_NORMALIZATION_ALGORITHM_FORMALIZED"
	StatusFirewallsPreserved         = "CONDITIONAL_SUPPORT_GATE300_EMPIRICAL_FIREWALLS_PRESERVED"

	StatusFailedCutoffMomentsUnfixed     = "FAILED_ROUTE_CUTOFF_MOMENTS_STILL_UNFIXED"
	StatusFailedSubtractionScheme        = "FAILED_ROUTE_HEAT_KERNEL_SUBTRACTION_SCHEME_STILL_MISSING"
	StatusFailedYukawaAmplitudesFree     = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_FREE"
	StatusFailedPositiveZHNotProved      = "FAILED_ROUTE_POSITIVE_ZH_NOT_NUMERICALLY_PROVED"
	StatusFailedAbsoluteGaugeCouplings   = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLINGS_NOT_DERIVED"
	StatusFailedHiggsMassNotDerived      = "FAILED_ROUTE_HIGGS_MASS_PARAMETER_NOT_DERIVED"
	StatusFailedHiggsQuarticNotDerived   = "FAILED_ROUTE_HIGGS_QUARTIC_NOT_DERIVED"
	StatusFailedRawRatioNotObservable    = "FAILED_ROUTE_RAW_1197_4624_RATIO_NOT_A_PHYSICAL_OBSERVABLE"
	StatusFailedBGapInstantonStillSealed = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

type InheritedGate299 struct {
	HeatKernelExpansionFormalized bool
	A2ScalarQuadraticChannel      bool
	A4ScalarKineticChannel        bool
	A4GaugeKineticChannel         bool
	A4ScalarQuarticChannel        bool
	RawTraceRatio                 string
	RawTraceRatioNumerator        int
	RawTraceRatioDenominator      int
	PhysicalDynamicsDerived       bool
	Verdict                       string
}

type MonomialClassifier struct {
	Name                 string
	HeatKernelSource     string
	DerivativeOrder      int
	ScalarPower          int
	GaugeCurvaturePower  int
	VacuumPower          int
	AcceptedForZH        bool
	AcceptedForGaugeNorm bool
	AcceptedForPotential bool
	Reason               string
}

type KineticIsolation struct {
	ClassifierRules           []MonomialClassifier
	ScalarKineticSelector     string
	GaugeKineticSelector      string
	ScalarPotentialSelector   string
	SeparatesKineticPotential bool
	RejectsVacuumTerms        bool
	RejectsBGapMassInsertion  bool
	Algorithm                 string
	Verdict                   string
}

type WaveFunctionRenormalization struct {
	RawScalarKineticCoefficient string
	ZHDefinition                string
	RequiredCondition           string
	Rescaling                   string
	CanonicalTarget             string
	PositiveZHProved            bool
	NumericalZHComputed         bool
	AlgorithmValid              bool
	Verdict                     string
}

type PotentialCoefficient struct {
	Name                string
	RawChannel          string
	RawCoefficient      string
	PhysicalCoefficient string
	Requires            string
	DerivedNumerically  bool
}

type MassQuarticRescaling struct {
	Coefficients           []PotentialCoefficient
	RawRatio               string
	RawToPhysicalMap       string
	RatioInterpretation    string
	RawRatioPromoted       bool
	PhysicalMassDerived    bool
	PhysicalQuarticDerived bool
	AlgorithmFormalized    bool
	Verdict                string
}

type GaugeKineticNormalization struct {
	GaugeGroups                  []string
	TraceIndexDefinition         string
	RawCoefficientMap            string
	PhysicalCouplingMap          string
	HyperchargeNormalization     string
	AbsoluteCouplingsDerived     bool
	RelativeNormalizationAudited bool
	AlgorithmFormalized          bool
	Verdict                      string
}

type RemainingObligation struct {
	Name             string
	WhyRequired      string
	Status           string
	BlocksPrediction bool
}

type FirewallAudit struct {
	NoCutoffMomentsInserted     bool
	NoSubtractionSchemeInvented bool
	NoYukawaNumbersInserted     bool
	NoObservedMassesInserted    bool
	NoBGapInstantonClaimed      bool
	NoRawRatioPromotion         bool
	FiniteCorePolluted          bool
	Obligations                 []RemainingObligation
	Verdict                     string
}

type Summary struct {
	Gate299Inherited        bool
	KineticIsolation        bool
	ZHAlgorithm             bool
	MassQuarticMap          bool
	GaugeNormalizationMap   bool
	PhysicalDynamicsDerived bool
	FirewallPreserved       bool
	Status                  string
	DirectAnswer            string
	NextGate                string
}

type Analysis struct {
	Input     InheritedGate299
	Kinetic   KineticIsolation
	ZH        WaveFunctionRenormalization
	Rescaling MassQuarticRescaling
	Gauge     GaugeKineticNormalization
	Firewalls FirewallAudit
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
	input := inheritGate299()
	kinetic := formalizeKineticIsolation()
	zh := formalizeZH(kinetic)
	rescaling := formalizeMassQuarticRescaling(zh)
	gauge := formalizeGaugeNormalization()
	firewalls := auditFirewalls(zh, rescaling, gauge)
	summary := buildSummary(input, kinetic, zh, rescaling, gauge, firewalls)
	truth := "Gate 300 formalizes the canonical normalization sieve: isolate derivative-quadratic scalar terms inside a4, define Z_H from their raw coefficient, rescale H_raw=H_phys/sqrt(Z_H), divide the a2 scalar channel by Z_H, divide the a4 quartic channel by Z_H², and separately normalize gauge curvature traces into 1/g_i² coefficients. It is an algorithmic bridge only: no cutoff moments, subtraction constants, numerical Yukawa amplitudes, absolute couplings, Higgs mass, Higgs quartic, or B-gap instanton action are derived."
	return Analysis{Input: input, Kinetic: kinetic, ZH: zh, Rescaling: rescaling, Gauge: gauge, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate299() InheritedGate299 {
	return InheritedGate299{
		HeatKernelExpansionFormalized: true,
		A2ScalarQuadraticChannel:      true,
		A4ScalarKineticChannel:        true,
		A4GaugeKineticChannel:         true,
		A4ScalarQuarticChannel:        true,
		RawTraceRatio:                 "Tr(D_F^4)/(Tr(D_F^2))^2 = 1197/4624",
		RawTraceRatioNumerator:        1197,
		RawTraceRatioDenominator:      4624,
		PhysicalDynamicsDerived:       false,
		Verdict:                       StatusGate299Inherited,
	}
}

func formalizeKineticIsolation() KineticIsolation {
	classifiers := []MonomialClassifier{
		{Name: "scalar kinetic", HeatKernelSource: "a_4(D_A)", DerivativeOrder: 2, ScalarPower: 2, GaugeCurvaturePower: 0, VacuumPower: 0, AcceptedForZH: true, AcceptedForGaugeNorm: false, AcceptedForPotential: false, Reason: "unique derivative-quadratic Higgs channel: (D_mu H_raw)^†(D^mu H_raw)"},
		{Name: "gauge kinetic", HeatKernelSource: "a_4(D_A)", DerivativeOrder: 0, ScalarPower: 0, GaugeCurvaturePower: 2, VacuumPower: 0, AcceptedForZH: false, AcceptedForGaugeNorm: true, AcceptedForPotential: false, Reason: "curvature-quadratic channel: Tr_i(F_i,mu nu F_i^mu nu)"},
		{Name: "scalar quadratic potential", HeatKernelSource: "a_2(D_A)", DerivativeOrder: 0, ScalarPower: 2, GaugeCurvaturePower: 0, VacuumPower: 0, AcceptedForZH: false, AcceptedForGaugeNorm: false, AcceptedForPotential: true, Reason: "mass-parameter location after subtraction and scalar normalization"},
		{Name: "scalar quartic potential", HeatKernelSource: "a_4(D_A)", DerivativeOrder: 0, ScalarPower: 4, GaugeCurvaturePower: 0, VacuumPower: 0, AcceptedForZH: false, AcceptedForGaugeNorm: false, AcceptedForPotential: true, Reason: "quartic Higgs channel after scalar normalization"},
		{Name: "vacuum/cosmological", HeatKernelSource: "a_0(D_A), a_2(D_A), a_4(D_A)", DerivativeOrder: 0, ScalarPower: 0, GaugeCurvaturePower: 0, VacuumPower: 1, AcceptedForZH: false, AcceptedForGaugeNorm: false, AcceptedForPotential: false, Reason: "must be removed or fixed by a subtraction/renormalization scheme before scalar dynamics"},
	}
	return KineticIsolation{
		ClassifierRules:           classifiers,
		ScalarKineticSelector:     "select terms in a_4 with derivativeOrder=2, scalarPower=2, curvaturePower=0; project onto the Higgs-doublet one-form block",
		GaugeKineticSelector:      "select terms in a_4 with derivativeOrder=0, scalarPower=0, curvaturePower=2; group by U(1)_Y, SU(2)_L, SU(3)_C representation trace index",
		ScalarPotentialSelector:   "select non-derivative scalarPower=2 terms from a_2 after subtraction and scalarPower=4 terms from a_4",
		SeparatesKineticPotential: true,
		RejectsVacuumTerms:        true,
		RejectsBGapMassInsertion:  true,
		Algorithm:                 "decompose the heat-kernel polynomial by tensor degree: spacetime derivative degree, gauge curvature degree, scalar one-form degree, and vacuum degree; only the derivative-quadratic scalar block defines Z_H",
		Verdict:                   StatusKineticIsolationFormalized,
	}
}

func formalizeZH(k KineticIsolation) WaveFunctionRenormalization {
	return WaveFunctionRenormalization{
		RawScalarKineticCoefficient: "K_H^raw := coeff[a_4(D_A), (D_mu H_raw)^†(D^mu H_raw)]",
		ZHDefinition:                "Z_H := N_4 f_0 K_H^raw, where N_4 includes the chosen Seeley-de Witt numerical convention such as (4π)^-2 and trace normalization",
		RequiredCondition:           "Z_H must be finite and positive for a canonical propagating Higgs field; Gate 300 defines this test but does not prove numerical positivity without amplitudes",
		Rescaling:                   "H_raw = H_phys / sqrt(Z_H)",
		CanonicalTarget:             "Z_H |D_mu H_raw|^2 -> |D_mu H_phys|^2",
		PositiveZHProved:            false,
		NumericalZHComputed:         false,
		AlgorithmValid:              k.SeparatesKineticPotential && k.RejectsVacuumTerms,
		Verdict:                     strings.Join([]string{StatusZHFormalized, StatusFailedPositiveZHNotProved}, ";"),
	}
}

func formalizeMassQuarticRescaling(z WaveFunctionRenormalization) MassQuarticRescaling {
	coeffs := []PotentialCoefficient{
		{Name: "quadratic mass channel", RawChannel: "a_2(D_A)|_|H_raw|^2 after subtraction", RawCoefficient: "C_2^raw := N_2 f_2 Λ^2 [T_2 - S_2]", PhysicalCoefficient: "C_2^phys = C_2^raw / Z_H; in the common sign convention μ_H^2 = -C_2^phys", Requires: "subtraction S_2, cutoff moment f_2, scale Λ, sign convention, and positive Z_H", DerivedNumerically: false},
		{Name: "quartic scalar channel", RawChannel: "a_4(D_A)|_|H_raw|^4", RawCoefficient: "C_4^raw := N_4 f_0 T_4", PhysicalCoefficient: "λ_H = C_4^raw / Z_H^2", Requires: "scalar projection, trace convention, numerical Yukawa/Dirac amplitudes or a sealed amplitude theorem, and positive Z_H", DerivedNumerically: false},
	}
	return MassQuarticRescaling{
		Coefficients:           coeffs,
		RawRatio:               "R_raw := T_4 / T_2^2 = 1197/4624",
		RawToPhysicalMap:       "given T_2,T_4,K_H: Z_H=N_4 f_0 K_H; C_2^phys=N_2 f_2 Λ^2(T_2-S_2)/Z_H; λ_H=N_4 f_0 T_4/Z_H^2",
		RatioInterpretation:    "1197/4624 can enter only as a scale-free trace-shape factor inside T_4/T_2^2; it is not itself μ_H², λ_H, or a Higgs mass prediction until Z_H, f_0, f_2, Λ, S_2, and trace conventions are fixed",
		RawRatioPromoted:       false,
		PhysicalMassDerived:    false,
		PhysicalQuarticDerived: false,
		AlgorithmFormalized:    z.AlgorithmValid,
		Verdict:                strings.Join([]string{StatusRescalingMapFormalized, StatusFailedHiggsMassNotDerived, StatusFailedHiggsQuarticNotDerived, StatusFailedRawRatioNotObservable}, ";"),
	}
}

func formalizeGaugeNormalization() GaugeKineticNormalization {
	return GaugeKineticNormalization{
		GaugeGroups:                  []string{"U(1)_Y", "SU(2)_L", "SU(3)_C"},
		TraceIndexDefinition:         "τ_i := Tr_F(ρ(T_i)ρ(T_i)) over the completed finite Hilbert representation, using the selected generator normalization",
		RawCoefficientMap:            "K_i^raw := coeff[a_4(D_A), Tr(F_i,mu nu F_i^mu nu)] = N_4 f_0 τ_i",
		PhysicalCouplingMap:          "match K_i^raw F_i^2 to the canonical convention (1/4g_i^2)F_i^2; equivalently g_i^{-2}=4K_i^raw in that convention, with convention factors explicit",
		HyperchargeNormalization:     "the inherited k_Y=5/3 trace normalization supports the relative boundary sin²θ_W=3/8 when the common f_0 convention is shared, but not an absolute value of g_i without f_0",
		AbsoluteCouplingsDerived:     false,
		RelativeNormalizationAudited: true,
		AlgorithmFormalized:          true,
		Verdict:                      strings.Join([]string{StatusGaugeNormFormalized, StatusFailedAbsoluteGaugeCouplings}, ";"),
	}
}

func auditFirewalls(z WaveFunctionRenormalization, r MassQuarticRescaling, g GaugeKineticNormalization) FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "cutoff moments f_0,f_2,f_4", WhyRequired: "multiply the raw heat-kernel channels and set absolute normalization", Status: StatusFailedCutoffMomentsUnfixed, BlocksPrediction: true},
		{Name: "heat-kernel subtraction S_2", WhyRequired: "separates scalar mass term from vacuum/regulator-dependent contributions", Status: StatusFailedSubtractionScheme, BlocksPrediction: true},
		{Name: "numerical Yukawa/Dirac amplitudes", WhyRequired: "determine T_2,T_4,K_H after the structural edge graph is known", Status: StatusFailedYukawaAmplitudesFree, BlocksPrediction: true},
		{Name: "positive Z_H theorem", WhyRequired: "canonical scalar propagation requires finite positive wave-function coefficient", Status: StatusFailedPositiveZHNotProved, BlocksPrediction: true},
		{Name: "absolute gauge coupling scale", WhyRequired: "relative trace ratios do not determine f_0 or absolute g_i values", Status: StatusFailedAbsoluteGaugeCouplings, BlocksPrediction: true},
		{Name: "B-gap instanton action", WhyRequired: "normalization of polynomial heat-kernel traces does not derive S_inst=(4/pi)/B_gap", Status: StatusFailedBGapInstantonStillSealed, BlocksPrediction: true},
	}
	return FirewallAudit{
		NoCutoffMomentsInserted:     true,
		NoSubtractionSchemeInvented: true,
		NoYukawaNumbersInserted:     true,
		NoObservedMassesInserted:    true,
		NoBGapInstantonClaimed:      true,
		NoRawRatioPromotion:         !r.RawRatioPromoted,
		FiniteCorePolluted:          false,
		Obligations:                 obligations,
		Verdict:                     strings.Join([]string{StatusFirewallsPreserved, StatusFailedCutoffMomentsUnfixed, StatusFailedSubtractionScheme, StatusFailedYukawaAmplitudesFree, StatusFailedBGapInstantonStillSealed}, ";"),
	}
}

func buildSummary(i InheritedGate299, k KineticIsolation, z WaveFunctionRenormalization, r MassQuarticRescaling, g GaugeKineticNormalization, f FirewallAudit) Summary {
	statuses := []string{
		StatusGate299Inherited,
		StatusKineticIsolationFormalized,
		StatusZHFormalized,
		StatusRescalingMapFormalized,
		StatusGaugeNormFormalized,
		StatusAlgorithmFormalized,
		StatusFirewallsPreserved,
		StatusFailedCutoffMomentsUnfixed,
		StatusFailedSubtractionScheme,
		StatusFailedYukawaAmplitudesFree,
		StatusFailedPositiveZHNotProved,
		StatusFailedAbsoluteGaugeCouplings,
		StatusFailedHiggsMassNotDerived,
		StatusFailedHiggsQuarticNotDerived,
		StatusFailedRawRatioNotObservable,
		StatusFailedBGapInstantonStillSealed,
	}
	return Summary{
		Gate299Inherited:        i.HeatKernelExpansionFormalized && !i.PhysicalDynamicsDerived,
		KineticIsolation:        k.SeparatesKineticPotential,
		ZHAlgorithm:             z.AlgorithmValid,
		MassQuarticMap:          r.AlgorithmFormalized && !r.RawRatioPromoted,
		GaugeNormalizationMap:   g.AlgorithmFormalized && g.RelativeNormalizationAudited,
		PhysicalDynamicsDerived: false,
		FirewallPreserved:       !f.FiniteCorePolluted && f.NoCutoffMomentsInserted && f.NoSubtractionSchemeInvented && f.NoYukawaNumbersInserted && f.NoBGapInstantonClaimed,
		Status:                  strings.Join(statuses, ";"),
		DirectAnswer:            "Gate 300 converts Gate 299's missing-normalization obstruction into a precise algebraic instruction manual: isolate scalar kinetic and gauge curvature monomials, define Z_H and gauge K_i coefficients, and map raw a2/a4 potential traces into canonical parameters only after the required schemes are supplied.",
		NextGate:                "Gate 301 should audit the minimal finite data needed to make Z_H evaluable: a scalar kinetic trace functional K_H tied to the completed physical Hilbert representation, while preserving the empirical Yukawa-amplitude seal.",
	}
}

func FormatInput(i InheritedGate299) string {
	return fmt.Sprintf("heat=%t a2Scalar=%t a4ScalarKin=%t a4Gauge=%t a4Quartic=%t ratio=%s num=%d den=%d dynamics=%t verdict=%s", i.HeatKernelExpansionFormalized, i.A2ScalarQuadraticChannel, i.A4ScalarKineticChannel, i.A4GaugeKineticChannel, i.A4ScalarQuarticChannel, i.RawTraceRatio, i.RawTraceRatioNumerator, i.RawTraceRatioDenominator, i.PhysicalDynamicsDerived, i.Verdict)
}

func FormatClassifier(c MonomialClassifier) string {
	return fmt.Sprintf("%s source=%s d=%d scalar=%d curvature=%d vacuum=%d ZH=%t gauge=%t potential=%t reason=%q", c.Name, c.HeatKernelSource, c.DerivativeOrder, c.ScalarPower, c.GaugeCurvaturePower, c.VacuumPower, c.AcceptedForZH, c.AcceptedForGaugeNorm, c.AcceptedForPotential, c.Reason)
}

func FormatKinetic(k KineticIsolation) string {
	parts := []string{}
	for _, c := range k.ClassifierRules {
		parts = append(parts, FormatClassifier(c))
	}
	return fmt.Sprintf("rules=[%s] scalarSelector=%q gaugeSelector=%q potentialSelector=%q separates=%t vacuumRejected=%t bgapRejected=%t algorithm=%q verdict=%s", strings.Join(parts, " | "), k.ScalarKineticSelector, k.GaugeKineticSelector, k.ScalarPotentialSelector, k.SeparatesKineticPotential, k.RejectsVacuumTerms, k.RejectsBGapMassInsertion, k.Algorithm, k.Verdict)
}

func FormatZH(z WaveFunctionRenormalization) string {
	return fmt.Sprintf("raw=%q ZH=%q condition=%q rescale=%q target=%q positive=%t numeric=%t valid=%t verdict=%s", z.RawScalarKineticCoefficient, z.ZHDefinition, z.RequiredCondition, z.Rescaling, z.CanonicalTarget, z.PositiveZHProved, z.NumericalZHComputed, z.AlgorithmValid, z.Verdict)
}

func FormatPotentialCoefficient(c PotentialCoefficient) string {
	return fmt.Sprintf("%s rawChannel=%q rawCoeff=%q physical=%q requires=%q numeric=%t", c.Name, c.RawChannel, c.RawCoefficient, c.PhysicalCoefficient, c.Requires, c.DerivedNumerically)
}

func FormatRescaling(r MassQuarticRescaling) string {
	parts := []string{}
	for _, c := range r.Coefficients {
		parts = append(parts, FormatPotentialCoefficient(c))
	}
	return fmt.Sprintf("coefficients=[%s] rawRatio=%q map=%q interpretation=%q promoted=%t mass=%t quartic=%t formalized=%t verdict=%s", strings.Join(parts, " | "), r.RawRatio, r.RawToPhysicalMap, r.RatioInterpretation, r.RawRatioPromoted, r.PhysicalMassDerived, r.PhysicalQuarticDerived, r.AlgorithmFormalized, r.Verdict)
}

func FormatGauge(g GaugeKineticNormalization) string {
	return fmt.Sprintf("groups=%s trace=%q raw=%q physical=%q hypercharge=%q absolute=%t relative=%t formalized=%t verdict=%s", strings.Join(g.GaugeGroups, ","), g.TraceIndexDefinition, g.RawCoefficientMap, g.PhysicalCouplingMap, g.HyperchargeNormalization, g.AbsoluteCouplingsDerived, g.RelativeNormalizationAudited, g.AlgorithmFormalized, g.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	parts := []string{}
	for _, o := range f.Obligations {
		parts = append(parts, FormatObligation(o))
	}
	return fmt.Sprintf("noCutoff=%t noSubtraction=%t noYukawa=%t noObservedMass=%t noBGap=%t noRatioPromotion=%t polluted=%t obligations=[%s] verdict=%s", f.NoCutoffMomentsInserted, f.NoSubtractionSchemeInvented, f.NoYukawaNumbersInserted, f.NoObservedMassesInserted, f.NoBGapInstantonClaimed, f.NoRawRatioPromotion, f.FiniteCorePolluted, strings.Join(parts, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("inherit=%t kinetic=%t ZH=%t massQuartic=%t gauge=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate299Inherited, s.KineticIsolation, s.ZHAlgorithm, s.MassQuarticMap, s.GaugeNormalizationMap, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
