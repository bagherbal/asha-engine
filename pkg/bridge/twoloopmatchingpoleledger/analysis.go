// Package twoloopmatchingpoleledger implements Gate 310:
// Two-Loop / Matching / Pole-Mass Conversion Ledger Audit.
//
// Gate 309 transported the Gate-308 UV quartic boundary through a conditional
// one-loop RG lane and obtained a high Higgs-mass diagnostic. Gate 310 does not
// repair that number by tuning. It builds the missing precision ledger: two-loop
// RG terms, finite threshold jumps, and MS-bar-to-pole conversion. The result is
// a capacity audit: which correction classes can plausibly move the diagnostic,
// which cannot, and which data must be sealed before a final collider-scale mass
// can be claimed.
package twoloopmatchingpoleledger

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE310-TWO-LOOP-MATCHING-POLE-MASS-CONVERSION-LEDGER"

	StatusGate309DiagnosticInherited           = "CONDITIONAL_SUPPORT_GATE309_ONE_LOOP_DIAGNOSTIC_INHERITED"
	StatusTwoLoopLedgerFormalized              = "CONDITIONAL_SUPPORT_TWO_LOOP_RG_LEDGER_FORMALIZED"
	StatusThresholdMatchingLedgerFormalized    = "CONDITIONAL_SUPPORT_THRESHOLD_MATCHING_LEDGER_FORMALIZED"
	StatusPoleMassConversionLedgerFormalized   = "CONDITIONAL_SUPPORT_POLE_MASS_CONVERSION_LEDGER_FORMALIZED"
	StatusHigherOrderTransportLedgerFormalized = "CONDITIONAL_SUPPORT_HIGHER_ORDER_TRANSPORT_LEDGER_FORMALIZED"
	StatusTensionCapacityAssessed              = "CONDITIONAL_SUPPORT_ONE_LOOP_TENSION_CAPACITY_ASSESSED"
	StatusStructuralTopSectorQuestionOpened    = "CONDITIONAL_TENSION_MODIFIED_TOP_SECTOR_OR_THRESHOLD_LEDGER_REQUIRED"
	StatusGate310FirewallsPreserved            = "CONDITIONAL_SUPPORT_GATE310_FIREWALLS_PRESERVED"

	StatusFailedTwoLoopNotExecuted             = "FAILED_ROUTE_TWO_LOOP_RGE_NOT_EXECUTED"
	StatusFailedFullTwoLoopCoefficientsMissing = "FAILED_ROUTE_FULL_TWO_LOOP_COEFFICIENT_TABLE_NOT_INSTALLED"
	StatusFailedThresholdValuesNotDerived      = "FAILED_ROUTE_THRESHOLD_MATCHING_VALUES_NOT_DERIVED"
	StatusFailedPoleSelfEnergiesNotComputed    = "FAILED_ROUTE_POLE_SELF_ENERGIES_NOT_COMPUTED"
	StatusFailedNoResolvedHiggsMass            = "FAILED_ROUTE_FINAL_HIGGS_MASS_NOT_RESOLVED"
	StatusFailedNoColliderPrediction           = "FAILED_ROUTE_FINAL_COLLIDER_SCALE_PREDICTION_NOT_CLAIMED"
)

const (
	vEVGeV                  = 246.22
	uvQuarticBoundary       = 1197.0 / 4624.0
	gate309PrimaryLambdaAtV = 0.907051722647
	gate309PrimaryMassGeV   = 331.630412
	observedReferenceGeV    = 125.10
	loopSuppressionOne      = 1.0 / (16.0 * math.Pi * math.Pi)
	loopSuppressionTwo      = loopSuppressionOne * loopSuppressionOne
)

type Gate309Inheritance struct {
	BoundaryEquation         string
	UVLambda                 float64
	GStarSquaredSeal         float64
	PrimaryLane              string
	PrimaryLambdaAtV         float64
	PrimaryMassGeV           float64
	GaugeOnlyMassGeV         float64
	PureSMHighScaleRejected  bool
	OneLoopOnly              bool
	ThresholdMatchingOmitted bool
	PoleMassMatchingOmitted  bool
	FinalColliderMassClaimed bool
	InheritedAsDiagnostic    bool
	Verdict                  string
}

type TwoLoopTerm struct {
	Name                 string
	StructuralForm       string
	RepresentativeSign   string
	DownwardFlowEffect   string
	CanSoftenTension     bool
	CanExacerbateTension bool
	RequiresCoefficient  bool
	ExactCoefficientUsed bool
}

type TwoLoopLedger struct {
	Formalized                    bool
	BetaConvention                string
	OneLoopSuppression            float64
	TwoLoopSuppression            float64
	RepresentativeTerms           []TwoLoopTerm
	PositiveBetaSoftensDownward   bool
	NegativeBetaAmplifiesDownward bool
	ExactFullSystemInstalled      bool
	TwoLoopIntegrationExecuted    bool
	ExpectedCapacity              string
	CanResolveAlone               bool
	Verdict                       string
}

type ThresholdSource struct {
	Name                    string
	MatchingEquation        string
	TypicalSignFreedom      string
	CanGenerateNegativeJump bool
	TreeLevelPossible       bool
	OneLoopPossible         bool
	ValueDerived            bool
	RequiresMassLedger      bool
	RequiresCouplingLedger  bool
	Verdict                 string
}

type ThresholdLedger struct {
	Formalized            bool
	MatchingRule          string
	RequiredIRLambdaShift float64
	RequiredMassShiftGeV  float64
	RequiredShiftClass    string
	Sources               []ThresholdSource
	HasCapacityToResolve  bool
	ValuesDerived         bool
	ThresholdsExecuted    bool
	Verdict               string
}

type PoleMassLedger struct {
	Formalized             bool
	RunningMassFormula     string
	PoleMassFormula        string
	RunningMassInputGeV    float64
	ReferencePoleMassGeV   float64
	DirectMassGapGeV       float64
	RequiredLambdaAtV      float64
	RequiredLambdaShiftAtV float64
	SelfEnergySources      []string
	ExpectedCapacity       string
	CanResolveAlone        bool
	SelfEnergiesComputed   bool
	UsesMeasuredMassForFit bool
	Verdict                string
}

type TensionResolutionAudit struct {
	Formalized                      bool
	Gate309LambdaAtV                float64
	ReferenceLambdaAtV              float64
	LambdaExcessAtV                 float64
	MassExcessGeV                   float64
	TwoLoopCanResolveAlone          bool
	PoleMassCanResolveAlone         bool
	ThresholdsCanResolveInPrinciple bool
	ModifiedTopSectorMayBeRequired  bool
	NeedsFullPrecisionRun           bool
	FinalMassResolved               bool
	Verdict                         string
}

type FirewallAudit struct {
	NoTwoLoopNumericalTransportRun bool
	NoThresholdJumpInserted        bool
	NoPoleSelfEnergyInserted       bool
	NoObservedHiggsUsedAsFit       bool
	NoObservedTopUsedAsFit         bool
	NoFinalMassClaimed             bool
	NoFiniteCorePolluted           bool
	Obligations                    []RemainingObligation
	Verdict                        string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksFinalPrediction     bool
}

type Summary struct {
	Gate309Inherited               bool
	TwoLoopLedgerReady             bool
	ThresholdLedgerReady           bool
	PoleMassLedgerReady            bool
	CapacityAssessed               bool
	FinalMassResolved              bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Inheritance Gate309Inheritance
	TwoLoop     TwoLoopLedger
	Thresholds  ThresholdLedger
	PoleMass    PoleMassLedger
	Tension     TensionResolutionAudit
	Firewalls   FirewallAudit
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
	inheritance := inheritGate309()
	twoLoop := formalizeTwoLoopLedger()
	pole := formalizePoleMassLedger(inheritance)
	thresholds := formalizeThresholdLedger(pole)
	tension := assessTension(inheritance, twoLoop, thresholds, pole)
	firewalls := auditFirewalls()
	summary := buildSummary(inheritance, twoLoop, thresholds, pole, tension, firewalls)
	truth := "Gate 310 formalizes the precision-continuum ledger missing from Gate 309. Two-loop RG terms can soften or exacerbate the downward quartic flow depending on sign, pole-mass conversion is too small to erase a ~200 GeV diagnostic by itself, and threshold matching has the mathematical capacity to move lambda only if finite jumps and heavy-sector couplings are derived. The 331 GeV one-loop number is therefore a tension diagnostic, not a final prediction."
	return Analysis{Inheritance: inheritance, TwoLoop: twoLoop, Thresholds: thresholds, PoleMass: pole, Tension: tension, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritGate309() Gate309Inheritance {
	return Gate309Inheritance{
		BoundaryEquation:         "λ_H(Λ_GUT) = (1197/4624) · g_*² with g_*²=1",
		UVLambda:                 uvQuarticBoundary,
		GStarSquaredSeal:         1,
		PrimaryLane:              "Gate206 Dirac vectorlike quark doublet PeV lane + r_plus_top_yukawa_boundary_seal",
		PrimaryLambdaAtV:         gate309PrimaryLambdaAtV,
		PrimaryMassGeV:           gate309PrimaryMassGeV,
		GaugeOnlyMassGeV:         157.0,
		PureSMHighScaleRejected:  true,
		OneLoopOnly:              true,
		ThresholdMatchingOmitted: true,
		PoleMassMatchingOmitted:  true,
		FinalColliderMassClaimed: false,
		InheritedAsDiagnostic:    true,
		Verdict:                  StatusGate309DiagnosticInherited,
	}
}

func formalizeTwoLoopLedger() TwoLoopLedger {
	terms := []TwoLoopTerm{
		{Name: "top sextic", StructuralForm: "+ y_t^6 representative term", RepresentativeSign: "positive in βλ", DownwardFlowEffect: "positive β contribution lowers λ during UV-to-IR transport, softening the Gate 309 upward drift", CanSoftenTension: true, RequiresCoefficient: true},
		{Name: "quartic cubic", StructuralForm: "- λ^3 representative term", RepresentativeSign: "negative in βλ", DownwardFlowEffect: "negative β contribution raises λ during UV-to-IR transport, potentially exacerbating the high-mass diagnostic", CanExacerbateTension: true, RequiresCoefficient: true},
		{Name: "mixed gauge-Yukawa", StructuralForm: "g_i² y_t⁴, g_i⁴ y_t², λ g_i² y_t²", RepresentativeSign: "mixed", DownwardFlowEffect: "direction depends on the full coefficient table and current coupling values", CanSoftenTension: true, CanExacerbateTension: true, RequiresCoefficient: true},
		{Name: "pure gauge", StructuralForm: "g_i^6 and mixed gauge monomials", RepresentativeSign: "mixed", DownwardFlowEffect: "can shift the quartic flow but is generally loop-suppressed relative to the one-loop top diagnostic", CanSoftenTension: true, CanExacerbateTension: true, RequiresCoefficient: true},
	}
	return TwoLoopLedger{
		Formalized:                    true,
		BetaConvention:                "β_x = β_x^(1)/(16π²) + β_x^(2)/(16π²)²; transport is top-down with dt=dlnμ<0",
		OneLoopSuppression:            loopSuppressionOne,
		TwoLoopSuppression:            loopSuppressionTwo,
		RepresentativeTerms:           terms,
		PositiveBetaSoftensDownward:   true,
		NegativeBetaAmplifiesDownward: true,
		ExactFullSystemInstalled:      false,
		TwoLoopIntegrationExecuted:    false,
		ExpectedCapacity:              "perturbative correction class; important for precision, but not authorized as a standalone 200 GeV repair without coefficients and integration",
		CanResolveAlone:               false,
		Verdict:                       strings.Join([]string{StatusTwoLoopLedgerFormalized, StatusFailedFullTwoLoopCoefficientsMissing, StatusFailedTwoLoopNotExecuted}, ";"),
	}
}

func formalizePoleMassLedger(in Gate309Inheritance) PoleMassLedger {
	requiredLambda := observedReferenceGeV * observedReferenceGeV / (2.0 * vEVGeV * vEVGeV)
	return PoleMassLedger{
		Formalized:             true,
		RunningMassFormula:     "m_run(v)=v·sqrt(2λ(v))",
		PoleMassFormula:        "m_pole²=m_run²+Π_HH(m_pole²)-counterterms; requires W/Z/top/Higgs self-energy ledger",
		RunningMassInputGeV:    in.PrimaryMassGeV,
		ReferencePoleMassGeV:   observedReferenceGeV,
		DirectMassGapGeV:       in.PrimaryMassGeV - observedReferenceGeV,
		RequiredLambdaAtV:      requiredLambda,
		RequiredLambdaShiftAtV: requiredLambda - in.PrimaryLambdaAtV,
		SelfEnergySources:      []string{"top-quark self-energy", "W-boson loop", "Z-boson loop", "Higgs/Goldstone loops", "renormalization-scheme counterterms"},
		ExpectedCapacity:       "precision conversion; essential for final comparison, but not naturally a ~200 GeV correction in a controlled perturbative ledger",
		CanResolveAlone:        false,
		SelfEnergiesComputed:   false,
		UsesMeasuredMassForFit: false,
		Verdict:                strings.Join([]string{StatusPoleMassConversionLedgerFormalized, StatusFailedPoleSelfEnergiesNotComputed}, ";"),
	}
}

func formalizeThresholdLedger(p PoleMassLedger) ThresholdLedger {
	sources := []ThresholdSource{
		{Name: "PeV vectorlike/adjoint states", MatchingEquation: "λ(μ_-)=λ(μ_+)+Δλ_PeV", TypicalSignFreedom: "model-dependent", CanGenerateNegativeJump: true, TreeLevelPossible: false, OneLoopPossible: true, RequiresMassLedger: true, RequiresCouplingLedger: true, Verdict: StatusFailedThresholdValuesNotDerived},
		{Name: "B-gap / right-handed-neutrino sector", MatchingEquation: "λ(μ_-)=λ(μ_+)+Δλ_Bgap", TypicalSignFreedom: "depends on Majorana/seesaw activation theorem", CanGenerateNegativeJump: true, TreeLevelPossible: true, OneLoopPossible: true, RequiresMassLedger: true, RequiresCouplingLedger: true, Verdict: StatusFailedThresholdValuesNotDerived},
		{Name: "scalar portal / heavy scalar residue", MatchingEquation: "λ_eff=λ_full-κ²/M² + loop residues", TypicalSignFreedom: "negative tree-level shifts possible if a heavy scalar is integrated out", CanGenerateNegativeJump: true, TreeLevelPossible: true, OneLoopPossible: true, RequiresMassLedger: true, RequiresCouplingLedger: true, Verdict: StatusFailedThresholdValuesNotDerived},
	}
	return ThresholdLedger{
		Formalized:            true,
		MatchingRule:          "at each heavy threshold M: λ_below(M)=λ_above(M)+Δλ_threshold, with β-functions changed below M",
		RequiredIRLambdaShift: p.RequiredLambdaShiftAtV,
		RequiredMassShiftGeV:  -p.DirectMassGapGeV,
		RequiredShiftClass:    "large negative effective shift at IR-equivalent scale; exact UV threshold size requires running sensitivity matrix",
		Sources:               sources,
		HasCapacityToResolve:  true,
		ValuesDerived:         false,
		ThresholdsExecuted:    false,
		Verdict:               strings.Join([]string{StatusThresholdMatchingLedgerFormalized, StatusFailedThresholdValuesNotDerived}, ";"),
	}
}

func assessTension(in Gate309Inheritance, two TwoLoopLedger, th ThresholdLedger, pole PoleMassLedger) TensionResolutionAudit {
	lambdaExcess := in.PrimaryLambdaAtV - pole.RequiredLambdaAtV
	massExcess := in.PrimaryMassGeV - pole.ReferencePoleMassGeV
	verdicts := []string{StatusTensionCapacityAssessed, StatusStructuralTopSectorQuestionOpened, StatusFailedNoResolvedHiggsMass}
	return TensionResolutionAudit{
		Formalized:                      true,
		Gate309LambdaAtV:                in.PrimaryLambdaAtV,
		ReferenceLambdaAtV:              pole.RequiredLambdaAtV,
		LambdaExcessAtV:                 lambdaExcess,
		MassExcessGeV:                   massExcess,
		TwoLoopCanResolveAlone:          two.CanResolveAlone,
		PoleMassCanResolveAlone:         pole.CanResolveAlone,
		ThresholdsCanResolveInPrinciple: th.HasCapacityToResolve,
		ModifiedTopSectorMayBeRequired:  true,
		NeedsFullPrecisionRun:           true,
		FinalMassResolved:               false,
		Verdict:                         strings.Join(verdicts, ";"),
	}
}

func auditFirewalls() FirewallAudit {
	obligations := []RemainingObligation{
		{Name: "full two-loop SM/BSM coefficient table", WhyRequired: "needed before the direction and size of βλ^(2) can be numerically integrated", Status: StatusFailedFullTwoLoopCoefficientsMissing, BlocksFinalPrediction: true},
		{Name: "finite threshold matching values", WhyRequired: "thresholds are the only formal class in this ledger with enough capacity to bridge the full λ gap", Status: StatusFailedThresholdValuesNotDerived, BlocksFinalPrediction: true},
		{Name: "pole-mass self-energy calculation", WhyRequired: "required to compare an MS-bar running mass to the collider pole mass", Status: StatusFailedPoleSelfEnergiesNotComputed, BlocksFinalPrediction: true},
		{Name: "top-sector amplitude origin", WhyRequired: "the r_+ top seal drives the Gate 309 tension and may need a tensor-sector correction", Status: "FAILED_ROUTE_TOP_YUKAWA_ORIGIN_STILL_SEALED", BlocksFinalPrediction: true},
	}
	return FirewallAudit{
		NoTwoLoopNumericalTransportRun: true,
		NoThresholdJumpInserted:        true,
		NoPoleSelfEnergyInserted:       true,
		NoObservedHiggsUsedAsFit:       true,
		NoObservedTopUsedAsFit:         true,
		NoFinalMassClaimed:             true,
		NoFiniteCorePolluted:           true,
		Obligations:                    obligations,
		Verdict:                        strings.Join([]string{StatusGate310FirewallsPreserved, StatusFailedNoColliderPrediction}, ";"),
	}
}

func buildSummary(in Gate309Inheritance, two TwoLoopLedger, th ThresholdLedger, pole PoleMassLedger, tension TensionResolutionAudit, f FirewallAudit) Summary {
	return Summary{
		Gate309Inherited:     in.InheritedAsDiagnostic,
		TwoLoopLedgerReady:   two.Formalized,
		ThresholdLedgerReady: th.Formalized,
		PoleMassLedgerReady:  pole.Formalized,
		CapacityAssessed:     tension.Formalized,
		FinalMassResolved:    tension.FinalMassResolved,
		FirewallPreserved:    f.NoFinalMassClaimed && f.NoFiniteCorePolluted,
		Status:               strings.Join([]string{StatusHigherOrderTransportLedgerFormalized, StatusFailedNoResolvedHiggsMass}, ";"),
		DirectAnswer:         "Gate 310 authorizes the precision ledger, not a corrected mass: two-loop and pole corrections are necessary but not enough by themselves; finite threshold matching or a modified top-sector tensor is the likely resolution class.",
		NextGate:             "Gate 311 should instantiate a controlled threshold-sensitivity matrix Δλ_i -> λ(v), then test whether sealed PeV/B-gap thresholds can reduce the Gate 309 tension without empirical tuning.",
	}
}

func FormatInheritance(x Gate309Inheritance) string {
	return fmt.Sprintf("boundary=%s λUV=%.12f lane=%s λ(v)=%.12f m=%.6f GeV oneLoop=%v thresholdMatchingOmitted=%v poleOmitted=%v finalClaim=%v verdict=%s", x.BoundaryEquation, x.UVLambda, x.PrimaryLane, x.PrimaryLambdaAtV, x.PrimaryMassGeV, x.OneLoopOnly, x.ThresholdMatchingOmitted, x.PoleMassMatchingOmitted, x.FinalColliderMassClaimed, x.Verdict)
}

func FormatTwoLoop(x TwoLoopLedger) string {
	names := make([]string, 0, len(x.RepresentativeTerms))
	for _, t := range x.RepresentativeTerms {
		names = append(names, t.Name+":"+t.DownwardFlowEffect)
	}
	return fmt.Sprintf("formalized=%v convention=%s eps1=%.6g eps2=%.6g positiveBetaSoftens=%v negativeBetaAmplifies=%v exactSystem=%v executed=%v canResolveAlone=%v terms=[%s] verdict=%s", x.Formalized, x.BetaConvention, x.OneLoopSuppression, x.TwoLoopSuppression, x.PositiveBetaSoftensDownward, x.NegativeBetaAmplifiesDownward, x.ExactFullSystemInstalled, x.TwoLoopIntegrationExecuted, x.CanResolveAlone, strings.Join(names, " | "), x.Verdict)
}

func FormatThresholds(x ThresholdLedger) string {
	names := make([]string, 0, len(x.Sources))
	for _, s := range x.Sources {
		names = append(names, s.Name+":"+s.MatchingEquation)
	}
	return fmt.Sprintf("formalized=%v rule=%s requiredIRΔλ=%.12f requiredΔm=%.6f GeV class=%s hasCapacity=%v valuesDerived=%v executed=%v sources=[%s] verdict=%s", x.Formalized, x.MatchingRule, x.RequiredIRLambdaShift, x.RequiredMassShiftGeV, x.RequiredShiftClass, x.HasCapacityToResolve, x.ValuesDerived, x.ThresholdsExecuted, strings.Join(names, " | "), x.Verdict)
}

func FormatPole(x PoleMassLedger) string {
	return fmt.Sprintf("formalized=%v runFormula=%s poleFormula=%s mrun=%.6f reference=%.6f gap=%.6f λtarget=%.12f Δλtarget=%.12f canResolveAlone=%v selfEnergiesComputed=%v fit=%v verdict=%s", x.Formalized, x.RunningMassFormula, x.PoleMassFormula, x.RunningMassInputGeV, x.ReferencePoleMassGeV, x.DirectMassGapGeV, x.RequiredLambdaAtV, x.RequiredLambdaShiftAtV, x.CanResolveAlone, x.SelfEnergiesComputed, x.UsesMeasuredMassForFit, x.Verdict)
}

func FormatTension(x TensionResolutionAudit) string {
	return fmt.Sprintf("formalized=%v λGate309=%.12f λref=%.12f λexcess=%.12f massExcess=%.6f twoLoopAlone=%v poleAlone=%v thresholdsCapacity=%v modifiedTopSectorMayBeRequired=%v needsFullPrecision=%v resolved=%v verdict=%s", x.Formalized, x.Gate309LambdaAtV, x.ReferenceLambdaAtV, x.LambdaExcessAtV, x.MassExcessGeV, x.TwoLoopCanResolveAlone, x.PoleMassCanResolveAlone, x.ThresholdsCanResolveInPrinciple, x.ModifiedTopSectorMayBeRequired, x.NeedsFullPrecisionRun, x.FinalMassResolved, x.Verdict)
}

func FormatFirewalls(x FirewallAudit) string {
	obligations := make([]string, 0, len(x.Obligations))
	for _, o := range x.Obligations {
		obligations = append(obligations, o.Name+":"+o.Status)
	}
	return fmt.Sprintf("noTwoLoopRun=%v noThresholdInserted=%v noPoleInserted=%v noHiggsFit=%v noTopFit=%v noFinalMass=%v noPollution=%v obligations=[%s] verdict=%s", x.NoTwoLoopNumericalTransportRun, x.NoThresholdJumpInserted, x.NoPoleSelfEnergyInserted, x.NoObservedHiggsUsedAsFit, x.NoObservedTopUsedAsFit, x.NoFinalMassClaimed, x.NoFiniteCorePolluted, strings.Join(obligations, " | "), x.Verdict)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("gate309=%v twoLoop=%v thresholds=%v pole=%v capacity=%v resolved=%v firewall=%v status=%s answer=%s next=%s", x.Gate309Inherited, x.TwoLoopLedgerReady, x.ThresholdLedgerReady, x.PoleMassLedgerReady, x.CapacityAssessed, x.FinalMassResolved, x.FirewallPreserved, x.Status, x.DirectAnswer, x.NextGate)
}
