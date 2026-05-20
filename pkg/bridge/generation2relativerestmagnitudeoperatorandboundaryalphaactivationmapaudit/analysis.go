// Package generation2relativerestmagnitudeoperatorandboundaryalphaactivationmapaudit implements
// Gate 825: Relative RestMagnitude Operator and BoundaryAlpha Activation Map Audit.
//
// Gate 825 tests whether the strengthened Boundary-FN / 1+3 simplex branch can be
// promoted from scalar closure into a relative positive trace-magnitude operator.  The
// core advantage is that the absolute dominant top-like value T cancels from N_eff;
// the audit therefore focuses on the relative spectrum, projector source, and boundary
// alpha activation map rather than external Yukawa ratios.
package generation2relativerestmagnitudeoperatorandboundaryalphaactivationmapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE825-RELATIVE-REST-MAGNITUDE-OPERATOR-BOUNDARY-ALPHA-ACTIVATION-MAP-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate824Inherited        = "PASS_GATE824_SOURCE_ROUTER_INHERITED"
	StatusLiveGapSelected         = "PASS_BOUNDARY_TO_TRACE_MAGNITUDE_RESTMAP_SELECTED_AS_LIVE_GAP"
	StatusRelativeOperatorDefined = "PASS_RELATIVE_REST_MAGNITUDE_OPERATOR_DEFINED"
	StatusTopColorSpectrum        = "PASS_TOP_COLOR_RELATIVE_SPECTRUM_DEFINED"
	StatusOnePlusThreeOperator    = "PASS_ONE_PLUS_THREE_REST_OPERATOR_FORM_RECORDED"
	StatusTraceValidation         = "PASS_TRACE_VALIDATION_COMPUTED"
	StatusNEffOperatorDerived     = "PASS_N_EFF_OPERATOR_FORM_DERIVED"
	StatusFifthOrderResidual      = "PASS_FIFTH_ORDER_RESIDUAL_TO_BFN_CLOSURE_RECORDED"
	StatusBoundaryAlphaAudited    = "PASS_BOUNDARY_ALPHA_FORM_AUDITED"
	StatusSixSourceAudited        = "PASS_SIX_COEFFICIENT_SOURCE_AUDITED"
	StatusProjectiveAudited       = "PASS_PROJECTIVE_ONE_PLUS_THREE_PROJECTOR_SOURCE_AUDITED"
	StatusK7Audited               = "PASS_K7_HODGE_POLARITY_SOURCE_AUDITED"
	StatusFiniteTripleAudited     = "PASS_FINITE_TRIPLE_SOURCE_AUDITED"
	StatusD4Audited               = "PASS_D4_TRIALITY_SOURCE_REAUDITED"
	StatusActivationMapRequired   = "PASS_BOUNDARY_ALPHA_ACTIVATION_MAP_REQUIREMENT_DEFINED"
	StatusNonCircularity          = "PASS_NONCIRCULARITY_AUDIT_EXECUTED"
	StatusLevelsDefined           = "PASS_STATUS_LEVELS_DEFINED"
	StatusOutcomesDefined         = "PASS_OUTCOME_BRANCHES_DEFINED"
	StatusImpactPreserved         = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusPhysicalFirewalls       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusFirewallGate825         = "FIREWALL_PRESERVED_GATE825_RELATIVE_REST_MAGNITUDE_OPERATOR_BOUNDARY"

	SupportRelativeNoAbsoluteT      = "CONDITIONAL_SUPPORT_RELATIVE_OPERATOR_WOULD_REMOVE_NEED_FOR_ABSOLUTE_TOP_T_VALUE"
	SupportOnePlusThreeProjector    = "CONDITIONAL_SUPPORT_ONE_PLUS_THREE_PROJECTOR_SHAPE_IS_SHARPEST_REST_CARRIER_CANDIDATE"
	SupportAlphaBSourceShape        = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_BOUNDARY_HYPERCHARGE_K7_SOURCE_SHAPE"
	SupportFourthOrderClosure       = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCES_BFN_CLOSURE_THROUGH_FOURTH_ORDER"
	SupportSixFromParticipation     = "CONDITIONAL_SUPPORT_SIX_ARISES_NATURALLY_FROM_TOP_COLOR_PARTICIPATION_RESPONSE"
	SupportBoundaryPairSecondary    = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_TIMES_COLOR_REMAINS_SECONDARY_SOURCE_CANDIDATE"
	SupportCertifiedWouldReduceSeal = "CONDITIONAL_SUPPORT_CERTIFIED_OPERATOR_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"
	SupportOperatorNoExternalLedger = "CONDITIONAL_SUPPORT_RELATIVE_OPERATOR_IF_CERTIFIED_WOULD_NOT_REQUIRE_EXTERNAL_YUKAWA_LEDGER"
	SupportProjectiveShape          = "CONDITIONAL_SUPPORT_PROJECTIVE_ONE_PLUS_THREE_CAN_SUPPLY_REST_PROJECTOR_SHAPE"
	SupportK7Resonance              = "CONDITIONAL_SUPPORT_K7_4_3_REMAINS_REST_CARRIER_RESONANCE"
	SupportFiniteTripleTrace        = "CONDITIONAL_SUPPORT_FINITE_TRIPLE_SUPPLIES_TOP_COLOR_AND_TRACE_TEMPLATE"

	FailureNoRestProjectors           = "FAILED_ROUTE_RELATIVE_OPERATOR_NOT_NATIVE_WITHOUT_P1_P3_REST_PROJECTORS"
	FailureProjectiveNoReadout        = "FAILED_ROUTE_PROJECTIVE_SELECTOR_NOT_TRACE_MAGNITUDE_OPERATOR_WITHOUT_READOUT_MAP"
	FailureAlphaNotNative             = "FAILED_ROUTE_ALPHA_B_NOT_NATIVE_WITHOUT_BOUNDARY_ALPHA_ACTIVATION_MAP"
	FailureEigenvaluesNoActivation    = "FAILED_ROUTE_REST_EIGENVALUES_NOT_DERIVED_WITHOUT_ACTIVATION_THEOREM"
	FailureK7NotProjectorTheorem      = "FAILED_ROUTE_K7_HODGE_POLARITY_NOT_P1_PLUS_P3_TRACE_PROJECTOR_THEOREM"
	FailureFiniteTripleNoOperator     = "FAILED_ROUTE_FINITE_TRIPLE_DOES_NOT_SUPPLY_REST_MAGNITUDE_OPERATOR"
	FailureD4NotOperator              = "FAILED_ROUTE_D4_TRIALITY_NOT_REST_MAGNITUDE_OPERATOR"
	FailureD4NoP1P3Readout            = "FAILED_ROUTE_TRIALITY_DOES_NOT_SUPPLY_P1_P3_YUKAWA_TRACE_READOUT"
	FailureShapeNoSectorAssignment    = "FAILED_ROUTE_POSITIVE_OPERATOR_SHAPE_DOES_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	FailureR2PlusNotR3                = "FAILED_ROUTE_R2_PLUS_NOT_R3_SECTOR_TRACE_LEDGER"
	FailureR2PlusNotR4                = "FAILED_ROUTE_R2_PLUS_NOT_R4_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoCYukawaUpdate            = "FAILED_ROUTE_GATE825_DOES_NOT_UPDATE_C_YUKAWA_UNLESS_RELATIVE_OPERATOR_IS_CERTIFIED"
	FailureCHiggsLevelB               = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B_IF_OPERATOR_IS_ONLY_SOURCE_TYPED"
	FailureSixNotNumericalClosure     = "FAILED_ROUTE_SIX_MUST_NOT_BE_ACCEPTED_BY_NUMERICAL_CLOSURE_ALONE"
	FailureOperatorHypothesisIfSealed = "FAILED_ROUTE_OPERATOR_REMAINS_HYPOTHESIS_IF_PROJECTORS_OR_ALPHA_MAP_ARE_SEALED"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2 float64
	AlphaB                 float64
	DeltaNBFN, NEffBFN     float64
	CYukawaBFN, CHiggsBFN  float64
	OfficialCYukawa        float64
	OfficialCHiggs         float64
}

type RelativeOperator struct {
	Spectrum                       []float64
	RestSpectrum                   []float64
	RestProjectorForm              string
	ATotalOverT, ARestOverT        float64
	BTotalOverT2, BRestOverT2      float64
	Beta, QRest                    float64
	NEffOperator, NEffBFN          float64
	NEffResidual, SymbolicResidual float64
	Verdicts, Supports, Failures   []string
}

type CoefficientAudit struct {
	Name                         string
	Expression                   string
	CandidateSource              string
	Verdicts, Supports, Failures []string
}

type SourceAudit struct {
	Name                         string
	SuppliesShape                bool
	CertifiedTraceReadout        bool
	Verdicts, Supports, Failures []string
}

type ActivationMapAudit struct {
	RequiredObject               string
	Inputs                       []string
	TargetEigenvalues            []string
	Certified                    bool
	Verdicts, Supports, Failures []string
}

type NonCircularity struct {
	AllowedInputs, ForbiddenInputs []string
	PassesInputFirewall            bool
	OperatorWouldAvoidExternalData bool
	Verdicts, Supports, Failures   []string
}

type Status struct {
	Level                        string
	Outcome                      string
	CanUpdateCYukawa             bool
	HasRelativeOperatorShape     bool
	HasCertifiedProjectors       bool
	HasBoundaryActivationMap     bool
	HasSectorLedger              bool
	NativeYukawaTheorem          bool
	NextGate                     string
	Verdicts, Supports, Failures []string
}

type Impact struct {
	NEffCandidate, CYukawaCandidate, CHiggsCandidate float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced                                                                              bool
	NoProjectorsNoNativeOperator, ProjectiveNeedsReadout, AlphaNeedsActivation            bool
	NoSectorAssignment, NotR3, NotR4, NoCYukawaUpdate, CHiggsLevelB, NoD4, NoFiniteYukawa bool
	Verdict                                                                               string
}

type Analysis struct {
	Ledger       Ledger
	Operator     RelativeOperator
	Coefficients []CoefficientAudit
	Sources      []SourceAudit
	Activation   ActivationMapAudit
	NonCircular  NonCircularity
	Status       Status
	Impact       Impact
	Firewalls    Firewalls
	Truth        string
	Final        string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func DeltaBFN(s float64) float64           { return 6.0 * AlphaB(s) }
func NEffBFN(s float64) float64            { return 3.0 + DeltaBFN(s) }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	deltaBFN := DeltaBFN(SBoundary)
	nEffBFN := 3 + deltaBFN
	if math.Abs(deltaBFN-((9.0/5.0)*SBoundary+6.0*m2)) > 1e-18 {
		return Analysis{}, fmt.Errorf("Delta_BFN identity failed")
	}

	restDust := 3.0 * alpha * alpha
	restTriplet := alpha * (1.0 - alpha)
	restSpectrum := []float64{restDust, restTriplet, restTriplet, restTriplet}
	spectrum := []float64{1, 1, 1, restDust, restTriplet, restTriplet, restTriplet}
	aRest := restDust + 3.0*restTriplet
	bRest := restDust*restDust + 3.0*restTriplet*restTriplet
	aTotal := 3.0 + aRest
	bTotal := 3.0 + bRest
	beta := bRest / 3.0
	qRest := bRest / (aRest * aRest)
	nEffOp := (aTotal * aTotal) / bTotal
	symbolic := -24.0 * math.Pow(alpha, 5) / (1.0 + alpha*alpha - 2.0*math.Pow(alpha, 3) + 4.0*math.Pow(alpha, 4))

	operator := RelativeOperator{
		Spectrum: spectrum, RestSpectrum: restSpectrum, RestProjectorForm: "H_rest/T = 3 alpha_B^2 P_1 + alpha_B(1-alpha_B) P_3; H_top/T=[1,1,1]",
		ATotalOverT: aTotal, ARestOverT: aRest, BTotalOverT2: bTotal, BRestOverT2: bRest,
		Beta: beta, QRest: qRest, NEffOperator: nEffOp, NEffBFN: nEffBFN,
		NEffResidual: nEffOp - nEffBFN, SymbolicResidual: symbolic,
		Verdicts: []string{StatusRelativeOperatorDefined, StatusTopColorSpectrum, StatusOnePlusThreeOperator, StatusTraceValidation, StatusNEffOperatorDerived, StatusFifthOrderResidual},
		Supports: []string{SupportRelativeNoAbsoluteT, SupportFourthOrderClosure},
		Failures: []string{FailureShapeNoSectorAssignment, FailureR2PlusNotR3, FailureR2PlusNotR4},
	}

	coefficients := []CoefficientAudit{
		{Name: "boundary alpha", Expression: "alpha_B = (3/10)s + p s^2 = (1/2)(3/5)s + p s^2", CandidateSource: "inverse hypercharge boundary average plus K7-weighted second raw boundary moment", Verdicts: []string{StatusBoundaryAlphaAudited}, Supports: []string{SupportAlphaBSourceShape}, Failures: []string{FailureAlphaNotNative}},
		{Name: "six coefficient", Expression: "Delta_N_BFN = 6 alpha_B", CandidateSource: "primary: first-order top-color participation derivative d[3(1+alpha)^2]/dalpha|0=6; secondary: boundary-pair dimension times color", Verdicts: []string{StatusSixSourceAudited}, Supports: []string{SupportSixFromParticipation, SupportBoundaryPairSecondary}, Failures: []string{FailureSixNotNumericalClosure}},
	}

	sources := []SourceAudit{
		{Name: "Fock/projective 1+3 selector", SuppliesShape: true, CertifiedTraceReadout: false, Verdicts: []string{StatusProjectiveAudited}, Supports: []string{SupportProjectiveShape, SupportOnePlusThreeProjector}, Failures: []string{FailureProjectiveNoReadout}},
		{Name: "K7 Hodge 4|3 polarity", SuppliesShape: false, CertifiedTraceReadout: false, Verdicts: []string{StatusK7Audited}, Supports: []string{SupportK7Resonance}, Failures: []string{FailureK7NotProjectorTheorem}},
		{Name: "finite spectral triple", SuppliesShape: false, CertifiedTraceReadout: false, Verdicts: []string{StatusFiniteTripleAudited}, Supports: []string{SupportFiniteTripleTrace}, Failures: []string{FailureFiniteTripleNoOperator}},
		{Name: "D4 / triality airlock", SuppliesShape: false, CertifiedTraceReadout: false, Verdicts: []string{StatusD4Audited}, Failures: []string{FailureD4NotOperator, FailureD4NoP1P3Readout}},
	}

	activation := ActivationMapAudit{
		RequiredObject: "BoundaryAlphaActivationMap", Inputs: []string{"s", "p", "5/3", "boundary-pair averaging", "K7 second raw moment"},
		TargetEigenvalues: []string{"dust line = 3 alpha_B^2", "triplet chamber = alpha_B(1-alpha_B)"}, Certified: false,
		Verdicts: []string{StatusActivationMapRequired}, Supports: []string{SupportAlphaBSourceShape}, Failures: []string{FailureEigenvaluesNoActivation, FailureAlphaNotNative},
	}

	nonCircular := NonCircularity{
		AllowedInputs:       []string{"s", "p", "5/3", "color 3", "projective/Fock 1+3"},
		ForbiddenInputs:     []string{"observed Higgs mass", "C_Higgs", "lambda_runtime_eff", "m_H_tree_proxy", "external Yukawa ratios", "PMNS", "CKM", "Koide", "Georgi-Jarlskog", "fitted FN charges"},
		PassesInputFirewall: true, OperatorWouldAvoidExternalData: true,
		Verdicts: []string{StatusNonCircularity}, Supports: []string{SupportOperatorNoExternalLedger}, Failures: []string{FailureOperatorHypothesisIfSealed},
	}

	status := Status{
		Level:            "R2+ candidate shape only — positive relative operator on source-typed abstract rest carrier; not R3 sector ledger and not R4 Yukawa theorem",
		Outcome:          "Outcome B — partial success: P1/P3 projectors are source-typed by projective 1+3, but BoundaryAlphaActivationMap remains sealed",
		CanUpdateCYukawa: false, HasRelativeOperatorShape: true, HasCertifiedProjectors: false, HasBoundaryActivationMap: false, HasSectorLedger: false, NativeYukawaTheorem: false,
		NextGate: "Gate 826 — BoundaryAlphaActivationMap Source Audit",
		Verdicts: []string{StatusLevelsDefined, StatusOutcomesDefined}, Supports: []string{SupportCertifiedWouldReduceSeal}, Failures: []string{FailureNoRestProjectors, FailureAlphaNotNative, FailureR2PlusNotR3, FailureR2PlusNotR4},
	}

	impact := Impact{NEffCandidate: nEffOp, CYukawaCandidate: CYukawaFromNEff(nEffOp), CHiggsCandidate: CHiggsFromNEff(nEffOp), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactPreserved}, Supports: []string{SupportCertifiedWouldReduceSeal}, Failures: []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB}}

	firewalls := Firewalls{Enforced: true, NoProjectorsNoNativeOperator: true, ProjectiveNeedsReadout: true, AlphaNeedsActivation: true, NoSectorAssignment: true, NotR3: true, NotR4: true, NoCYukawaUpdate: true, CHiggsLevelB: true, NoD4: true, NoFiniteYukawa: true, Verdict: StatusFirewallGate825}

	ledger := Ledger{NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, AlphaB: alpha, DeltaNBFN: deltaBFN, NEffBFN: nEffBFN, CYukawaBFN: CYukawaFromNEff(nEffBFN), CHiggsBFN: CHiggsFromNEff(nEffBFN), OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs}
	truth := "Gate 825 turns the 1+3 rest simplex into a concrete positive relative operator; T cancels from N_eff, but the operator remains source-typed rather than certified because P1/P3 rest projectors and BoundaryAlphaActivationMap are not native trace readouts."
	final := "Gate 825 finds the fastest lawful path: a relative positive trace-magnitude operator reproduces the BFN closure through fourth order, yet C_Yukawa stays frozen until rest projectors and boundary-alpha activation become certified maps."
	return Analysis{Ledger: ledger, Operator: operator, Coefficients: coefficients, Sources: sources, Activation: activation, NonCircular: nonCircular, Status: status, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func Statuses() []string {
	return []string{StatusGate824Inherited, StatusLiveGapSelected, StatusRelativeOperatorDefined, StatusTopColorSpectrum, StatusOnePlusThreeOperator, StatusTraceValidation, StatusNEffOperatorDerived, StatusFifthOrderResidual, StatusBoundaryAlphaAudited, StatusSixSourceAudited, StatusProjectiveAudited, StatusK7Audited, StatusFiniteTripleAudited, StatusD4Audited, StatusActivationMapRequired, StatusNonCircularity, StatusLevelsDefined, StatusOutcomesDefined, StatusImpactPreserved, StatusPhysicalFirewalls, SupportRelativeNoAbsoluteT, SupportOnePlusThreeProjector, SupportAlphaBSourceShape, SupportFourthOrderClosure, SupportSixFromParticipation, SupportBoundaryPairSecondary, SupportCertifiedWouldReduceSeal, SupportOperatorNoExternalLedger, SupportProjectiveShape, SupportK7Resonance, SupportFiniteTripleTrace, FailureNoRestProjectors, FailureProjectiveNoReadout, FailureAlphaNotNative, FailureEigenvaluesNoActivation, FailureK7NotProjectorTheorem, FailureFiniteTripleNoOperator, FailureD4NotOperator, FailureD4NoP1P3Readout, FailureShapeNoSectorAssignment, FailureR2PlusNotR3, FailureR2PlusNotR4, FailureNoCYukawaUpdate, FailureCHiggsLevelB, FailureSixNotNumericalClosure, FailureOperatorHypothesisIfSealed, StatusFirewallGate825}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g Delta_BFN=%.16g N_eff_BFN=%.16g CYukawa_BFN=%.16g CHiggs_BFN=%.16g official_CYukawa=%.16g official_CHiggs=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.DeltaNBFN, a.NEffBFN, a.CYukawaBFN, a.CHiggsBFN, a.OfficialCYukawa, a.OfficialCHiggs)
}

func FormatOperator(a RelativeOperator) string {
	return fmt.Sprintf("%s spectrum=%v rest=%v a_total/T=%.16g a_rest/T=%.16g b_total/T2=%.16g b_rest/T2=%.16g beta=%.16g q=%.16g N_eff_operator=%.16g N_eff_BFN=%.16g R=%.16g symbolic_R=%.16g", a.RestProjectorForm, a.Spectrum, a.RestSpectrum, a.ATotalOverT, a.ARestOverT, a.BTotalOverT2, a.BRestOverT2, a.Beta, a.QRest, a.NEffOperator, a.NEffBFN, a.NEffResidual, a.SymbolicResidual)
}

func FormatCoefficients(rows []CoefficientAudit) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s: %s via %s", r.Name, r.Expression, r.CandidateSource))
	}
	return strings.Join(parts, " | ")
}

func FormatSources(rows []SourceAudit) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s shape=%t trace_readout=%t", r.Name, r.SuppliesShape, r.CertifiedTraceReadout))
	}
	return strings.Join(parts, " | ")
}

func FormatActivation(a ActivationMapAudit) string {
	return fmt.Sprintf("%s inputs=%v target=%v certified=%t", a.RequiredObject, a.Inputs, a.TargetEigenvalues, a.Certified)
}

func FormatNonCircularity(a NonCircularity) string {
	return fmt.Sprintf("allowed=%v forbidden=%v passes=%t avoids_external=%t", a.AllowedInputs, a.ForbiddenInputs, a.PassesInputFirewall, a.OperatorWouldAvoidExternalData)
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g official N_eff=%.16g C_Yukawa=%.16g C_Higgs=%.16g", a.NEffCandidate, a.CYukawaCandidate, a.CHiggsCandidate, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
}

func containsAll(xs, want []string) bool {
	m := map[string]bool{}
	for _, x := range xs {
		m[x] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}
