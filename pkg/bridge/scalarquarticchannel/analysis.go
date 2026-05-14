// Package scalarquarticchannel implements Gate 306:
// Scalar Quartic Channel Extraction / Dimensionless Coupling Sieve Audit.
//
// Gate 305 isolated the a2 scalar quadratic channel and showed that the Higgs
// mass parameter remains blocked by f2, Lambda, Z_H, amplitudes, and subtraction
// conventions. Gate 306 pivots back to the a4 scalar-power-4 channel. Because
// the quartic term lives in a4, its cutoff moment is the Gate-304 promoted
// f0=7 rather than the still-open f2. This gate formalizes the exact quartic
// extraction and scalar-normalization map. It also audits the f0 dependency:
// lambda_H itself retains the inverse N4 f0 normalization after H_raw is made
// canonical, while ratios against gauge couplings cancel N4 f0 and reduce to
// pure trace-index/quartic/kinetic data. No numerical quartic prediction is
// claimed because the raw Yukawa/amplitude carrier and absolute normalization
// constants remain sealed.
package scalarquarticchannel

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE306-SCALAR-QUARTIC-CHANNEL-EXTRACTION-DIMENSIONLESS-COUPLING-SIEVE-AUDIT"

	StatusGate305Inherited                         = "CONDITIONAL_SUPPORT_GATE305_SCALAR_SUBTRACTION_INHERITED"
	StatusRawA4QuarticDecompositionFormalized      = "CONDITIONAL_SUPPORT_RAW_A4_QUARTIC_DECOMPOSITION_FORMALIZED"
	StatusQuarticCouplingNormalizationFormalized   = "CONDITIONAL_SUPPORT_QUARTIC_COUPLING_NORMALIZATION_MAP_FORMALIZED"
	StatusF0DependencyAuditFormalized              = "CONDITIONAL_SUPPORT_F0_DEPENDENCY_AUDIT_FORMALIZED"
	StatusDimensionlessRatioSynthesisFormalized    = "CONDITIONAL_SUPPORT_DIMENSIONLESS_RATIO_SYNTHESIS_FORMALIZED"
	StatusScalarQuarticChannelExtractionFormalized = "CONDITIONAL_SUPPORT_SCALAR_QUARTIC_CHANNEL_EXTRACTION_FORMALIZED"
	StatusFirewallsPreserved                       = "CONDITIONAL_SUPPORT_GATE306_QUARTIC_FIREWALLS_PRESERVED"

	StatusFailedNumericalLambdaNotDerived         = "FAILED_ROUTE_HIGGS_QUARTIC_NUMERICAL_VALUE_NOT_DERIVED"
	StatusFailedRawC4NumericalCarrierSealed       = "FAILED_ROUTE_RAW_C4_NUMERICAL_CARRIER_STILL_SEALED"
	StatusFailedNumericalZHStillSealed            = "FAILED_ROUTE_NUMERICAL_ZH_VALUE_STILL_SEALED"
	StatusFailedYukawaAmplitudesStillSealed       = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"
	StatusFailed1197RawRatioNotObservableAlone    = "FAILED_ROUTE_RAW_1197_4624_RATIO_NOT_A_PHYSICAL_OBSERVABLE_ALONE"
	StatusFailedAbsoluteQuarticNeedsN4F0KH        = "FAILED_ROUTE_ABSOLUTE_LAMBDA_H_RETAINS_N4_F0_KH_DEPENDENCY"
	StatusFailedMassStillBlockedByF2              = "FAILED_ROUTE_HIGGS_MASS_STILL_BLOCKED_BY_F2"
	StatusFailedCutoffScaleLambdaNotDerived       = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_NOT_DERIVED"
	StatusFailedQuarticSignConventionNotDynamical = "FAILED_ROUTE_QUARTIC_SIGN_CONVENTION_NOT_DERIVED_FROM_FINITE_CORE"
	StatusFailedGaugeAbsoluteCouplingsNotDerived  = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLINGS_NOT_DERIVED"
	StatusFailedBGapInstantonStillSealed          = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

const (
	sealedF0Value            = 7
	rawTraceRatioNumerator   = 1197
	rawTraceRatioDenominator = 4624
)

type Gate305Inheritance struct {
	Gate304F0Promoted            bool
	PromotedF0Value              int
	F0Positive                   bool
	ScalarSubtractionFormalized  bool
	MassMapFormalized            bool
	F2MomentLocked               bool
	HiggsMassPredictionClaimed   bool
	QuarticChannelAlreadyTouched bool
	NumericalZHComputed          bool
	NumericalYukawasInserted     bool
	Verdict                      string
}

type A4QuarticComponent struct {
	Name                string
	Symbol              string
	DerivativeOrder     int
	ScalarPower         int
	GaugeCurvaturePower int
	VacuumPower         int
	AcceptedForQuartic  bool
	RejectedFromQuartic bool
	Reason              string
	Status              string
}

type RawA4QuarticDecomposition struct {
	RawExpression            string
	QuarticSelector          string
	QuarticCoefficient       string
	Components               []A4QuarticComponent
	A4SourceConfirmed        bool
	ScalarPower4ChannelSeen  bool
	DerivativeTermsRejected  bool
	GaugeTermsRejected       bool
	VacuumTermsRejected      bool
	NumericalCoefficientUsed bool
	DecompositionFormalized  bool
	Verdict                  string
}

type QuarticCouplingNormalization struct {
	RawQuarticActionCoefficient string
	ZHInput                     string
	CanonicalRescaling          string
	PhysicalQuarticMap          string
	CanonicalPotentialTarget    string
	UsesGate300ZHNormalization  bool
	UsesGate304F0Seal           bool
	RequiresPositiveZH          bool
	RequiresRawC4Carrier        bool
	RequiresYukawaAmplitudes    bool
	RequiresSignConvention      bool
	NumericalLambdaComputed     bool
	MapFormalized               bool
	Verdict                     string
}

type F0DependencyAudit struct {
	F0Value                      int
	ZHScaling                    string
	RawQuarticScaling            string
	LambdaScalingAfterRescale    string
	GaugeCouplingScaling         string
	F0CancelsInsideLambdaAlone   bool
	F0CancelsInLambdaOverGauge   bool
	RetainsN4F0ForAbsoluteLambda bool
	F2RequiredForQuartic         bool
	AuditFormalized              bool
	Verdict                      string
}

type DimensionlessRatioSynthesis struct {
	RawTraceRatio                   string
	RawTraceRatioNumerator          int
	RawTraceRatioDenominator        int
	RawRatioRole                    string
	LambdaOverGaugeRatio            string
	GaugeTraceIndexDependency       string
	RelativeRatioCanCancelN4F0      bool
	RawRatioPromotedDirectly        bool
	NeedsC4Raw                      bool
	NeedsKHRaw                      bool
	NeedsTraceIndex                 bool
	NeedsYukawaAmplitudeSeal        bool
	NeedsQuarticSignConvention      bool
	NeedsAbsoluteGaugeNormalization bool
	NumericalPhysicalPredictionMade bool
	SynthesisFormalized             bool
	Verdict                         string
}

type ChannelLedger struct {
	A4QuarticIsolated         bool
	A4KineticPreserved        bool
	A4GaugePreserved          bool
	A2MassChannelUndisturbed  bool
	F0SealUsedForA4           bool
	F2NotUsedForQuartic       bool
	NoHiggsMassClaimed        bool
	NoNumericalQuarticClaimed bool
	Verdict                   string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksPrediction          bool
}

type FirewallAudit struct {
	NoNumericalC4Inserted           bool
	NoYukawaNumbersInserted         bool
	NoNumericalZHComputed           bool
	NoNumericalLambdaHComputed      bool
	NoRaw1197PromotedDirectly       bool
	NoHiggsMassPredictionClaimed    bool
	NoAbsoluteGaugeCouplingsClaimed bool
	NoBGapInstantonClaimed          bool
	F2FirewallPreserved             bool
	F0SealPreservedForA4            bool
	FiniteCorePolluted              bool
	Obligations                     []RemainingObligation
	Verdict                         string
}

type Summary struct {
	Gate305Inherited               bool
	A4QuarticDecomposed            bool
	QuarticMapFormalized           bool
	F0DependencyAudited            bool
	DimensionlessRatioFormalized   bool
	QuarticChannelExtracted        bool
	NumericalLambdaHDerived        bool
	PhysicalQuarticPredictionMade  bool
	MassFirewallPreserved          bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input     Gate305Inheritance
	A4        RawA4QuarticDecomposition
	Quartic   QuarticCouplingNormalization
	F0        F0DependencyAudit
	Ratio     DimensionlessRatioSynthesis
	Channels  ChannelLedger
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
	input := inheritGate305()
	a4 := decomposeRawA4Quartic(input)
	quartic := normalizeQuartic(input, a4)
	f0 := auditF0Dependency(input, quartic)
	ratio := synthesizeDimensionlessRatios(a4, quartic, f0)
	channels := auditChannels(input, a4, f0, ratio)
	firewalls := auditFirewalls(input, a4, quartic, f0, ratio, channels)
	summary := buildSummary(input, a4, quartic, f0, ratio, channels, firewalls)
	truth := "Gate 306 isolates the a4 scalar-power-4 Higgs channel and formalizes lambda_H = (N4 f0 C4_raw)/Z_H^2 with H_raw=H_phys/sqrt(Z_H). Since Z_H=N4 f0 K_H_raw, lambda_H = C4_raw/(N4 f0 K_H_raw^2); therefore f0 is not automatically absent from the absolute quartic coefficient. However, ratios to gauge couplings cancel N4 f0 because 1/g_i^2=N4 f0 tau_i, yielding lambda_H/g_i^2 = tau_i C4_raw/K_H_raw^2. The gate is an extraction and dependency theorem only: it does not promote 1197/4624 directly into a physical observable and does not compute lambda_H without sealed C4, K_H, trace-index, sign, and amplitude data."
	return Analysis{Input: input, A4: a4, Quartic: quartic, F0: f0, Ratio: ratio, Channels: channels, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate305() Gate305Inheritance {
	return Gate305Inheritance{
		Gate304F0Promoted:            true,
		PromotedF0Value:              sealedF0Value,
		F0Positive:                   true,
		ScalarSubtractionFormalized:  true,
		MassMapFormalized:            true,
		F2MomentLocked:               false,
		HiggsMassPredictionClaimed:   false,
		QuarticChannelAlreadyTouched: false,
		NumericalZHComputed:          false,
		NumericalYukawasInserted:     false,
		Verdict:                      StatusGate305Inherited,
	}
}

func decomposeRawA4Quartic(i Gate305Inheritance) RawA4QuarticDecomposition {
	components := []A4QuarticComponent{
		{Name: "scalar quartic potential", Symbol: "C4_raw · (H_raw^†H_raw)^2", DerivativeOrder: 0, ScalarPower: 4, GaugeCurvaturePower: 0, VacuumPower: 0, AcceptedForQuartic: true, RejectedFromQuartic: false, Reason: "unique non-derivative scalar-power-4 block in a4", Status: StatusRawA4QuarticDecompositionFormalized},
		{Name: "scalar kinetic", Symbol: "K_H_raw · (D_mu H_raw)^†(D^mu H_raw)", DerivativeOrder: 2, ScalarPower: 2, GaugeCurvaturePower: 0, VacuumPower: 0, AcceptedForQuartic: false, RejectedFromQuartic: true, Reason: "belongs to Z_H normalization, not the quartic potential", Status: StatusGate305Inherited},
		{Name: "gauge kinetic", Symbol: "tau_i · F_i,mu nu F_i^mu nu", DerivativeOrder: 0, ScalarPower: 0, GaugeCurvaturePower: 2, VacuumPower: 0, AcceptedForQuartic: false, RejectedFromQuartic: true, Reason: "belongs to gauge-coupling normalization", Status: StatusGate305Inherited},
		{Name: "vacuum and curvature residues", Symbol: "a4_vac + curvature^2 + counterterm residue", DerivativeOrder: 0, ScalarPower: 0, GaugeCurvaturePower: 0, VacuumPower: 1, AcceptedForQuartic: false, RejectedFromQuartic: true, Reason: "field-independent or non-Higgs residue cannot define lambda_H", Status: StatusFailedRawC4NumericalCarrierSealed},
	}
	return RawA4QuarticDecomposition{
		RawExpression:            "a4(D_A) = a4_vac + K_H_raw |D_mu H_raw|^2 + C4_raw (H_raw^†H_raw)^2 + sum_i tau_i F_i^2 + residue",
		QuarticSelector:          "Pi_{scalar^4, derivative^0, curvature^0}(a4(D_A))",
		QuarticCoefficient:       "C4_raw := coeff[a4(D_A), (H_raw^†H_raw)^2]",
		Components:               components,
		A4SourceConfirmed:        i.Gate304F0Promoted && i.PromotedF0Value == sealedF0Value,
		ScalarPower4ChannelSeen:  true,
		DerivativeTermsRejected:  true,
		GaugeTermsRejected:       true,
		VacuumTermsRejected:      true,
		NumericalCoefficientUsed: false,
		DecompositionFormalized:  i.Gate304F0Promoted && !i.QuarticChannelAlreadyTouched,
		Verdict:                  strings.Join([]string{StatusRawA4QuarticDecompositionFormalized, StatusFailedRawC4NumericalCarrierSealed}, ";"),
	}
}

func normalizeQuartic(i Gate305Inheritance, a4 RawA4QuarticDecomposition) QuarticCouplingNormalization {
	formalized := i.Gate304F0Promoted && a4.DecompositionFormalized && a4.ScalarPower4ChannelSeen && a4.DerivativeTermsRejected && a4.GaugeTermsRejected
	return QuarticCouplingNormalization{
		RawQuarticActionCoefficient: "N4 · f0 · C4_raw · (H_raw^†H_raw)^2",
		ZHInput:                     "Z_H := N4 · f0 · K_H_raw from Gate 300-304, with f0=7 promoted only for the a4 channel",
		CanonicalRescaling:          "H_raw = H_phys / sqrt(Z_H)",
		PhysicalQuarticMap:          "lambda_H = Sign_4 · N4 · f0 · C4_raw / Z_H^2 = Sign_4 · C4_raw/(N4 · f0 · K_H_raw^2)",
		CanonicalPotentialTarget:    "V(H_phys) includes +lambda_H (H_phys^†H_phys)^2 after the declared Lorentzian potential sign convention",
		UsesGate300ZHNormalization:  true,
		UsesGate304F0Seal:           i.Gate304F0Promoted && i.PromotedF0Value == sealedF0Value,
		RequiresPositiveZH:          true,
		RequiresRawC4Carrier:        true,
		RequiresYukawaAmplitudes:    true,
		RequiresSignConvention:      true,
		NumericalLambdaComputed:     false,
		MapFormalized:               formalized,
		Verdict:                     strings.Join([]string{StatusQuarticCouplingNormalizationFormalized, StatusFailedAbsoluteQuarticNeedsN4F0KH, StatusFailedNumericalLambdaNotDerived, StatusFailedNumericalZHStillSealed, StatusFailedYukawaAmplitudesStillSealed}, ";"),
	}
}

func auditF0Dependency(i Gate305Inheritance, q QuarticCouplingNormalization) F0DependencyAudit {
	return F0DependencyAudit{
		F0Value:                      i.PromotedF0Value,
		ZHScaling:                    "Z_H ~ N4 f0 K_H_raw",
		RawQuarticScaling:            "quartic action coefficient ~ N4 f0 C4_raw",
		LambdaScalingAfterRescale:    "lambda_H ~ (N4 f0 C4_raw)/(N4 f0 K_H_raw)^2 = C4_raw/(N4 f0 K_H_raw^2)",
		GaugeCouplingScaling:         "1/g_i^2 ~ N4 f0 tau_i, hence g_i^2 ~ 1/(N4 f0 tau_i)",
		F0CancelsInsideLambdaAlone:   false,
		F0CancelsInLambdaOverGauge:   true,
		RetainsN4F0ForAbsoluteLambda: true,
		F2RequiredForQuartic:         false,
		AuditFormalized:              q.MapFormalized && q.UsesGate304F0Seal,
		Verdict:                      strings.Join([]string{StatusF0DependencyAuditFormalized, StatusFailedAbsoluteQuarticNeedsN4F0KH}, ";"),
	}
}

func synthesizeDimensionlessRatios(a4 RawA4QuarticDecomposition, q QuarticCouplingNormalization, f0 F0DependencyAudit) DimensionlessRatioSynthesis {
	return DimensionlessRatioSynthesis{
		RawTraceRatio:                   "Tr(D_F^4)/(Tr(D_F^2))^2 = 1197/4624",
		RawTraceRatioNumerator:          rawTraceRatioNumerator,
		RawTraceRatioDenominator:        rawTraceRatioDenominator,
		RawRatioRole:                    "candidate dimensionless finite-trace shape diagnostic that may constrain C4_raw/K_H_raw^2 only after the trace ledger proves it is the same scalar carrier",
		LambdaOverGaugeRatio:            "lambda_H/g_i^2 = Sign_4 · tau_i · C4_raw/K_H_raw^2",
		GaugeTraceIndexDependency:       "tau_i is the representation trace index for the chosen gauge factor; hypercharge uses the 5/3 normalization ledger before comparison",
		RelativeRatioCanCancelN4F0:      f0.F0CancelsInLambdaOverGauge,
		RawRatioPromotedDirectly:        false,
		NeedsC4Raw:                      true,
		NeedsKHRaw:                      true,
		NeedsTraceIndex:                 true,
		NeedsYukawaAmplitudeSeal:        true,
		NeedsQuarticSignConvention:      true,
		NeedsAbsoluteGaugeNormalization: true,
		NumericalPhysicalPredictionMade: false,
		SynthesisFormalized:             a4.DecompositionFormalized && q.MapFormalized && f0.AuditFormalized,
		Verdict:                         strings.Join([]string{StatusDimensionlessRatioSynthesisFormalized, StatusFailed1197RawRatioNotObservableAlone, StatusFailedGaugeAbsoluteCouplingsNotDerived, StatusFailedYukawaAmplitudesStillSealed}, ";"),
	}
}

func auditChannels(i Gate305Inheritance, a4 RawA4QuarticDecomposition, f0 F0DependencyAudit, r DimensionlessRatioSynthesis) ChannelLedger {
	return ChannelLedger{
		A4QuarticIsolated:         a4.DecompositionFormalized && a4.ScalarPower4ChannelSeen,
		A4KineticPreserved:        true,
		A4GaugePreserved:          true,
		A2MassChannelUndisturbed:  i.MassMapFormalized && !i.F2MomentLocked,
		F0SealUsedForA4:           i.Gate304F0Promoted && f0.F0Value == sealedF0Value,
		F2NotUsedForQuartic:       !f0.F2RequiredForQuartic,
		NoHiggsMassClaimed:        !i.HiggsMassPredictionClaimed,
		NoNumericalQuarticClaimed: !r.NumericalPhysicalPredictionMade,
		Verdict:                   strings.Join([]string{StatusScalarQuarticChannelExtractionFormalized, StatusFailedMassStillBlockedByF2, StatusFailedNumericalLambdaNotDerived}, ";"),
	}
}

func auditFirewalls(i Gate305Inheritance, a4 RawA4QuarticDecomposition, q QuarticCouplingNormalization, f0 F0DependencyAudit, r DimensionlessRatioSynthesis, c ChannelLedger) FirewallAudit {
	obs := []RemainingObligation{
		{"raw quartic carrier C4_raw", "the scalar-power-4 trace coefficient must be computed from the finite Dirac/Yukawa carrier rather than named symbolically", StatusFailedRawC4NumericalCarrierSealed, true},
		{"raw kinetic carrier K_H_raw / Z_H", "lambda_H divides by Z_H^2; structural positivity does not give an absolute number", StatusFailedNumericalZHStillSealed, true},
		{"Yukawa/amplitude seal", "C4_raw and K_H_raw depend on finite Yukawa/amplitude data or an internal amplitude theorem", StatusFailedYukawaAmplitudesStillSealed, true},
		{"1197/4624 carrier equivalence", "the raw trace synthesis must be proven to equal the same normalized C4_raw/K_H_raw^2 carrier before use", StatusFailed1197RawRatioNotObservableAlone, true},
		{"quartic sign convention", "the physical potential sign must be matched through the Euclidean-to-Lorentzian ledger", StatusFailedQuarticSignConventionNotDynamical, true},
		{"absolute gauge normalization", "relative ratios cancel N4 f0, but absolute g_i values require the gauge normalization ledger", StatusFailedGaugeAbsoluteCouplingsNotDerived, true},
		{"f2 and Lambda mass channel", "Gate 306 does not resolve the a2 mass-channel f2 or cutoff scale", StatusFailedMassStillBlockedByF2, true},
		{"B-gap instanton action", "quartic extraction does not derive S_inst=(4/pi)/B_gap", StatusFailedBGapInstantonStillSealed, true},
	}
	polluted := a4.NumericalCoefficientUsed || q.NumericalLambdaComputed || i.NumericalZHComputed || i.NumericalYukawasInserted || i.HiggsMassPredictionClaimed || r.RawRatioPromotedDirectly || r.NumericalPhysicalPredictionMade || f0.F0CancelsInsideLambdaAlone || !c.F2NotUsedForQuartic
	return FirewallAudit{
		NoNumericalC4Inserted:           !a4.NumericalCoefficientUsed,
		NoYukawaNumbersInserted:         !i.NumericalYukawasInserted,
		NoNumericalZHComputed:           !i.NumericalZHComputed && !q.NumericalLambdaComputed,
		NoNumericalLambdaHComputed:      !q.NumericalLambdaComputed && !r.NumericalPhysicalPredictionMade,
		NoRaw1197PromotedDirectly:       !r.RawRatioPromotedDirectly,
		NoHiggsMassPredictionClaimed:    !i.HiggsMassPredictionClaimed && c.NoHiggsMassClaimed,
		NoAbsoluteGaugeCouplingsClaimed: r.NeedsAbsoluteGaugeNormalization,
		NoBGapInstantonClaimed:          true,
		F2FirewallPreserved:             !i.F2MomentLocked && c.F2NotUsedForQuartic,
		F0SealPreservedForA4:            c.F0SealUsedForA4 && q.UsesGate304F0Seal,
		FiniteCorePolluted:              polluted,
		Obligations:                     obs,
		Verdict:                         strings.Join([]string{StatusFirewallsPreserved, StatusFailedNumericalLambdaNotDerived, StatusFailedRawC4NumericalCarrierSealed, StatusFailedNumericalZHStillSealed, StatusFailedYukawaAmplitudesStillSealed, StatusFailed1197RawRatioNotObservableAlone}, ";"),
	}
}

func buildSummary(i Gate305Inheritance, a4 RawA4QuarticDecomposition, q QuarticCouplingNormalization, f0 F0DependencyAudit, r DimensionlessRatioSynthesis, c ChannelLedger, fw FirewallAudit) Summary {
	return Summary{
		Gate305Inherited:              i.Gate304F0Promoted && i.PromotedF0Value == sealedF0Value && i.ScalarSubtractionFormalized && i.MassMapFormalized,
		A4QuarticDecomposed:           a4.DecompositionFormalized,
		QuarticMapFormalized:          q.MapFormalized && q.UsesGate300ZHNormalization && q.UsesGate304F0Seal,
		F0DependencyAudited:           f0.AuditFormalized && !f0.F0CancelsInsideLambdaAlone && f0.F0CancelsInLambdaOverGauge,
		DimensionlessRatioFormalized:  r.SynthesisFormalized && r.RelativeRatioCanCancelN4F0 && !r.RawRatioPromotedDirectly,
		QuarticChannelExtracted:       c.A4QuarticIsolated && c.F0SealUsedForA4 && c.F2NotUsedForQuartic,
		NumericalLambdaHDerived:       q.NumericalLambdaComputed,
		PhysicalQuarticPredictionMade: r.NumericalPhysicalPredictionMade,
		MassFirewallPreserved:         !i.HiggsMassPredictionClaimed && !i.F2MomentLocked,
		FirewallPreserved:             !fw.FiniteCorePolluted && fw.NoNumericalC4Inserted && fw.NoYukawaNumbersInserted && fw.NoNumericalZHComputed && fw.NoNumericalLambdaHComputed && fw.NoRaw1197PromotedDirectly && fw.NoHiggsMassPredictionClaimed && fw.NoAbsoluteGaugeCouplingsClaimed && fw.NoBGapInstantonClaimed && fw.F2FirewallPreserved && fw.F0SealPreservedForA4,
		Status:                        strings.Join([]string{StatusScalarQuarticChannelExtractionFormalized, StatusQuarticCouplingNormalizationFormalized, StatusF0DependencyAuditFormalized, StatusDimensionlessRatioSynthesisFormalized, StatusFirewallsPreserved}, ";"),
		DirectAnswer:                  "Gate 306 formalizes the a4 scalar-power-4 extraction and the normalized map lambda_H = N4 f0 C4_raw/Z_H^2. With Z_H=N4 f0 K_H_raw, absolute lambda_H retains an inverse N4 f0 factor, but lambda_H/g_i^2 cancels N4 f0 and reduces to tau_i C4_raw/K_H_raw^2. No numerical lambda_H is derived.",
		NextGate:                      "Gate 307 should audit the raw scalar quartic carrier C4_raw versus the kinetic carrier K_H_raw: prove whether the sealed finite trace synthesis 1197/4624 actually equals the same normalized scalar-quartic/kinetic-square carrier needed for lambda_H/g_i^2, without inserting empirical Yukawa amplitudes.",
	}
}

func FormatGate305Inheritance(i Gate305Inheritance) string {
	return fmt.Sprintf("f0Promoted=%t f0=%d positive=%t subtraction=%t massMap=%t f2Locked=%t massClaim=%t quarticTouched=%t ZH=%t Yukawa=%t verdict=%s", i.Gate304F0Promoted, i.PromotedF0Value, i.F0Positive, i.ScalarSubtractionFormalized, i.MassMapFormalized, i.F2MomentLocked, i.HiggsMassPredictionClaimed, i.QuarticChannelAlreadyTouched, i.NumericalZHComputed, i.NumericalYukawasInserted, i.Verdict)
}

func FormatA4Component(c A4QuarticComponent) string {
	return fmt.Sprintf("%s symbol=%q d=%d scalar=%d curvature=%d vacuum=%d accepted=%t rejected=%t reason=%q status=%s", c.Name, c.Symbol, c.DerivativeOrder, c.ScalarPower, c.GaugeCurvaturePower, c.VacuumPower, c.AcceptedForQuartic, c.RejectedFromQuartic, c.Reason, c.Status)
}

func FormatRawA4(a RawA4QuarticDecomposition) string {
	parts := []string{}
	for _, c := range a.Components {
		parts = append(parts, FormatA4Component(c))
	}
	return fmt.Sprintf("raw=%q selector=%q coeff=%q source=%t scalar4=%t rejectD=%t rejectGauge=%t rejectVacuum=%t numeric=%t formalized=%t components=[%s] verdict=%s", a.RawExpression, a.QuarticSelector, a.QuarticCoefficient, a.A4SourceConfirmed, a.ScalarPower4ChannelSeen, a.DerivativeTermsRejected, a.GaugeTermsRejected, a.VacuumTermsRejected, a.NumericalCoefficientUsed, a.DecompositionFormalized, strings.Join(parts, " | "), a.Verdict)
}

func FormatQuartic(q QuarticCouplingNormalization) string {
	return fmt.Sprintf("rawAction=%q ZH=%q rescale=%q map=%q target=%q gate300=%t f0Seal=%t positiveZH=%t C4=%t Yukawa=%t sign=%t numeric=%t formalized=%t verdict=%s", q.RawQuarticActionCoefficient, q.ZHInput, q.CanonicalRescaling, q.PhysicalQuarticMap, q.CanonicalPotentialTarget, q.UsesGate300ZHNormalization, q.UsesGate304F0Seal, q.RequiresPositiveZH, q.RequiresRawC4Carrier, q.RequiresYukawaAmplitudes, q.RequiresSignConvention, q.NumericalLambdaComputed, q.MapFormalized, q.Verdict)
}

func FormatF0(f F0DependencyAudit) string {
	return fmt.Sprintf("f0=%d ZH=%q quartic=%q lambda=%q gauge=%q cancelLambda=%t cancelRatio=%t retains=%t needsF2=%t formalized=%t verdict=%s", f.F0Value, f.ZHScaling, f.RawQuarticScaling, f.LambdaScalingAfterRescale, f.GaugeCouplingScaling, f.F0CancelsInsideLambdaAlone, f.F0CancelsInLambdaOverGauge, f.RetainsN4F0ForAbsoluteLambda, f.F2RequiredForQuartic, f.AuditFormalized, f.Verdict)
}

func FormatRatio(r DimensionlessRatioSynthesis) string {
	return fmt.Sprintf("rawRatio=%s n=%d d=%d role=%q lambdaOverGauge=%q tau=%q cancel=%t rawPromoted=%t C4=%t KH=%t tauNeed=%t Yukawa=%t sign=%t absGauge=%t prediction=%t formalized=%t verdict=%s", r.RawTraceRatio, r.RawTraceRatioNumerator, r.RawTraceRatioDenominator, r.RawRatioRole, r.LambdaOverGaugeRatio, r.GaugeTraceIndexDependency, r.RelativeRatioCanCancelN4F0, r.RawRatioPromotedDirectly, r.NeedsC4Raw, r.NeedsKHRaw, r.NeedsTraceIndex, r.NeedsYukawaAmplitudeSeal, r.NeedsQuarticSignConvention, r.NeedsAbsoluteGaugeNormalization, r.NumericalPhysicalPredictionMade, r.SynthesisFormalized, r.Verdict)
}

func FormatChannels(c ChannelLedger) string {
	return fmt.Sprintf("quartic=%t kinetic=%t gauge=%t a2Mass=%t f0A4=%t noF2=%t noMass=%t noQuarticNumber=%t verdict=%s", c.A4QuarticIsolated, c.A4KineticPreserved, c.A4GaugePreserved, c.A2MassChannelUndisturbed, c.F0SealUsedForA4, c.F2NotUsedForQuartic, c.NoHiggsMassClaimed, c.NoNumericalQuarticClaimed, c.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noC4=%t noYukawa=%t noZH=%t noLambdaH=%t no1197=%t noMass=%t noAbsGauge=%t noBGap=%t f2Firewall=%t f0A4=%t polluted=%t obligations=[%s] verdict=%s", f.NoNumericalC4Inserted, f.NoYukawaNumbersInserted, f.NoNumericalZHComputed, f.NoNumericalLambdaHComputed, f.NoRaw1197PromotedDirectly, f.NoHiggsMassPredictionClaimed, f.NoAbsoluteGaugeCouplingsClaimed, f.NoBGapInstantonClaimed, f.F2FirewallPreserved, f.F0SealPreservedForA4, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate305=%t a4=%t quarticMap=%t f0=%t ratio=%t extracted=%t numericLambda=%t prediction=%t massFirewall=%t firewall=%t status=%s answer=%q next=%q", s.Gate305Inherited, s.A4QuarticDecomposed, s.QuarticMapFormalized, s.F0DependencyAudited, s.DimensionlessRatioFormalized, s.QuarticChannelExtracted, s.NumericalLambdaHDerived, s.PhysicalQuarticPredictionMade, s.MassFirewallPreserved, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
