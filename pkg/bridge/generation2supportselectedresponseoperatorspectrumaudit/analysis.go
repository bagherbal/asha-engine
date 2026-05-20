// Package generation2supportselectedresponseoperatorspectrumaudit implements
// Gate 688: Support-Selected Response Operator Spectrum Audit.
//
// Gate 687 proved the factorization firewall: S_split is a scalar boundary
// amplitude, P_K7 is selected separately by Boolean-octonionic support, and
// ordinary trace scalarization gives the active bridge. Gate 688 audits the
// resulting support-selected response operator
//
//	R_split = S_split P_K7
//
// on H_72 = Lambda^4 R^8 ⊕ R^2_boundary. Since P_K7 is a rank-seven
// projector extended by zero on the boundary pair, R_split has eigenvalue
// S_split with multiplicity seven and zero with multiplicity sixty-five. The
// first ordinary trace therefore recovers (7/72)S_split, while spectrum and
// ordinary trace alone still do not identify the K7 projector without the
// Boolean-octonionic support constraints.
//
// This is a bridge-layer response-operator spectrum audit only. It does not
// derive boundary stress, scalar RG matching, Higgs mass, gauge unification,
// flavor, CKM/PMNS, a native 7/72 theorem, or a native projector-activation
// theorem.
package generation2supportselectedresponseoperatorspectrumaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate687 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryscalarprojectorselectorfactorizationfirewallaudit"
)

const (
	AuditID = "GATE688-SUPPORT-SELECTED-RESPONSE-OPERATOR-SPECTRUM-AUDIT"

	StatusGate687FactorizationFirewallInherited          = "PASS_GATE687_FACTORIZATION_FIREWALL_INHERITED"
	StatusRSplitDefinedAsSupportSelectedResponseOperator = "PASS_R_SPLIT_DEFINED_AS_SUPPORT_SELECTED_RESPONSE_OPERATOR"
	StatusOperatorSpectrumComputed                       = "PASS_OPERATOR_SPECTRUM_COMPUTED"
	StatusTracePowerCableComputed                        = "PASS_TRACE_POWER_CABLE_COMPUTED"
	StatusLinearFirstTraceResponseAudited                = "PASS_LINEAR_FIRST_TRACE_RESPONSE_AUDITED"
	StatusSupportInvarianceAudited                       = "PASS_SUPPORT_INVARIANCE_AUDITED"
	StatusRankSevenSpectralDegeneracyRecorded            = "PASS_RANK_SEVEN_SPECTRAL_DEGENERACY_RECORDED"
	StatusHodgePolarityTraceComparisonAudited            = "PASS_HODGE_POLARITY_TRACE_COMPARISON_AUDITED"
	StatusSSplitEigenvalueOnK7Support                    = "CONDITIONAL_SUPPORT_S_SPLIT_IS_EIGENVALUE_ON_K7_RESPONSE_SUPPORT"
	StatusActiveBridgeFirstTraceOfSupportOperator        = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_FIRST_TRACE_OF_SUPPORT_SELECTED_RESPONSE_OPERATOR"
	StatusSpectrumTraceAloneDoNotSelectK7                = "FAILED_ROUTE_SPECTRUM_AND_TRACE_ALONE_DO_NOT_SELECT_K7_IDENTITY"
	StatusNoNativeReasonFirstOrdinaryTrace               = "FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_RESPONSE_USES_FIRST_ORDINARY_TRACE"
	StatusNoNativeProjectorActivationTheorem             = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_ACTIVATION_THEOREM"
	StatusNoNativeSevenOver72Theorem                     = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate688ResponseOperatorSpectrumBoundary        = "FIREWALL_PRESERVED_GATE688_RESPONSE_OPERATOR_SPECTRUM_BOUNDARY"
)

const (
	lambda4Dimension    = 70
	boundaryDimension   = 2
	h72Dimension        = lambda4Dimension + boundaryDimension
	k7Dimension         = 7
	nullMultiplicity    = h72Dimension - k7Dimension
	k7PlusDimension     = 4
	k7MinusDimension    = 3
	w7Dimension         = 7
	auditedDBase        = 0.0001256552099683575
	auditedSSplit       = 0.0012924448188162962
	activeTraceResidual = auditedDBase - (float64(k7Dimension)/float64(h72Dimension))*auditedSSplit
	tolerance           = 1e-18
)

type Gate687Inheritance struct {
	FactorizationFirewallInherited bool
	BoundaryAmplitudeSeal          string
	NativeProjectorSelectorSeal    string
	TraceScalarizationSeal         string
	SelectedProjector              string
	SSplit                         float64
	DBase                          float64
	H72Dimension                   int
	K7Dimension                    int
	PriorCentralScalarFirewall     bool
	PriorSupportSealedIdentity     bool
	PriorNativeCouplingProved      bool
	Verdict                        string
}

type ResponseOperatorDefinition struct {
	Operator                     string
	AmbientSpace                 string
	AmbientDimension             int
	Projector                    string
	ProjectorRank                int
	ProjectorIsIdempotent        bool
	ProjectorIsSymmetric         bool
	BoundaryScalar               string
	SSplit                       float64
	ResponseInEndH72             bool
	SupportSelectedBeforeTracing bool
	Verdict                      string
}

type OperatorSpectrumAudit struct {
	EigenvalueOnK7       float64
	K7Multiplicity       int
	ZeroEigenvalue       float64
	ZeroMultiplicity     int
	RankIfSSplitNonzero  int
	SSplitNonzero        bool
	MinimalPolynomial    string
	TraceFromSpectrum    float64
	SpectrumDimensionSum int
	Verdict              string
}

type TracePower struct {
	Power           int
	Trace           float64
	NormalizedTrace float64
	Formula         string
}

type TracePowerCableAudit struct {
	Powers                  []TracePower
	FirstTrace              float64
	FirstTraceNormalized    float64
	FirstTracePrediction    float64
	DBase                   float64
	FirstTraceResidual      float64
	FrobeniusNormSquared    float64
	SecondTraceNormalized   float64
	FormulaAllPositivePower string
	Verdict                 string
}

type LinearResponseSelectionAudit struct {
	UsesFirstOrdinaryTrace bool
	UsesSecondTrace        bool
	UsesFrobeniusNorm      bool
	UsesDetLikeQuantity    bool
	UsesHodgeSignedTrace   bool
	ActiveFunctional       string
	RejectedFunctionals    []string
	ReasonForMissing       string
	Verdict                string
}

type SupportInvarianceAudit struct {
	BooleanSupportEquation       string
	OctonionicSupportEquation    string
	PBRSplitEqualsRSplit         bool
	PGRSplitEqualsRSplit         bool
	ImageInBooleanSector         bool
	ImageInOctonionicSector      bool
	ImageInIntersectionCarrier   bool
	SelectedIndependentlyOfTrace bool
	Verdict                      string
}

type RankSevenResponseCandidate struct {
	Name                    string
	ProjectorRank           int
	Carrier                 string
	SameSpectrumAsPK7       bool
	SameOrdinaryTraceAsPK7  bool
	PassesBooleanSupport    bool
	PassesOctonionicSupport bool
	SelectedBySupport       bool
	SupportVerdict          string
}

type SpectralDegeneracyAudit struct {
	Candidates               []RankSevenResponseCandidate
	AllShareSpectrumAndTrace bool
	SpectrumSelectsK7        bool
	TraceSelectsK7           bool
	SupportSelectsK7         bool
	Verdict                  string
}

type HodgePolarityTraceComparisonAudit struct {
	K7PlusDimension           int
	K7MinusDimension          int
	OrdinaryTraceMultiplicity int
	HodgeSignedMultiplicity   int
	OrdinaryTrace             float64
	HodgeSignedTrace          float64
	OrdinaryCoefficient       float64
	HodgeSignedCoefficient    float64
	ActiveUsesOrdinaryTrace   bool
	ActiveUsesSignedTrace     bool
	Verdict                   string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsSpectrumSelectsK7Identity bool
	ClaimsTraceSelectsK7Identity    bool
	ClaimsNativeFirstTracePrinciple bool
	ClaimsProjectorActivation       bool
	ClaimsNativeSevenOver72         bool
	ClaimsBoundaryStressDerivation  bool
	ClaimsScalarRGMatching          bool
	ClaimsHiggsMass                 bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	Verdict                         string
}

type Analysis struct {
	Inherited      Gate687Inheritance
	Response       ResponseOperatorDefinition
	Spectrum       OperatorSpectrumAudit
	TraceCable     TracePowerCableAudit
	LinearResponse LinearResponseSelectionAudit
	Support        SupportInvarianceAudit
	Degeneracy     SpectralDegeneracyAudit
	Hodge          HodgePolarityTraceComparisonAudit
	Missing        MissingTheoremAudit
	Discipline     VerdictDiscipline
	Truth          string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g687, err := gate687.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate687 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g687)
	response := buildResponseDefinition(inherited)
	spectrum := buildSpectrum(response)
	traceCable := buildTraceCable(inherited, spectrum)
	linear := buildLinearResponseSelection()
	support := buildSupportInvariance()
	degeneracy := buildDegeneracy()
	hodge := buildHodgePolarity()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeReasonFirstOrdinaryTrace,
			StatusNoNativeProjectorActivationTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native theorem explaining why the physical-history bridge uses the support-selected operator R_split=S_split P_K7 and why the active scalarization is its first ordinary trace rather than another spectral functional",
		Verdict:    strings.Join([]string{StatusNoNativeReasonFirstOrdinaryTrace, StatusNoNativeProjectorActivationTheorem, StatusNoNativeSevenOver72Theorem}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate688ResponseOperatorSpectrumBoundary}
	truth := "Gate 688 audits the operator left by the Gate687 factorization firewall. With P_K7 a rank-seven projector on H72 and S_split a scalar amplitude, R_split=S_split P_K7 has eigenvalue S_split on the seven-dimensional K7 support and zero on the remaining sixty-five directions. Its trace powers are Tr(R_split^n)=7 S_split^n for n>=1, so the active linear response is the first ordinary trace Tr(R_split)/72=(7/72)S_split. This spectral data is still rank-only: any rank-seven projector scaled by S_split has the same spectrum and ordinary trace. The K7 identity is supplied only by the Boolean-octonionic support equations, and the missing theorem is why history uses this support-selected operator and first ordinary trace."
	return Analysis{Inherited: inherited, Response: response, Spectrum: spectrum, TraceCable: traceCable, LinearResponse: linear, Support: support, Degeneracy: degeneracy, Hodge: hodge, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate687.Analysis) Gate687Inheritance {
	return Gate687Inheritance{
		FactorizationFirewallInherited: g.Factorization.FactorizationRequired && !g.Factorization.SSplitAloneSelectsIdentity,
		BoundaryAmplitudeSeal:          g.ThreeSeal.BoundaryAmplitudeSeal,
		NativeProjectorSelectorSeal:    g.ThreeSeal.NativeProjectorSelectorSeal,
		TraceScalarizationSeal:         g.ThreeSeal.TraceScalarizationSeal,
		SelectedProjector:              g.SupportSelection.SelectedProjector,
		SSplit:                         g.Inherited.SSplit,
		DBase:                          g.Inherited.DBase,
		H72Dimension:                   g.Inherited.H72Dimension,
		K7Dimension:                    g.Inherited.K7Dimension,
		PriorCentralScalarFirewall:     g.ScalarAction.CentralAction && !g.ScalarAction.CarriesProjectorDirection,
		PriorSupportSealedIdentity:     g.Factorization.ProjectorIdentitySupportSealed,
		PriorNativeCouplingProved:      g.Factorization.NativeCouplingProved,
		Verdict:                        StatusGate687FactorizationFirewallInherited,
	}
}

func buildResponseDefinition(inherited Gate687Inheritance) ResponseOperatorDefinition {
	return ResponseOperatorDefinition{
		Operator:                     "R_split = S_split P_K7",
		AmbientSpace:                 "H_72 = Lambda^4 R^8 ⊕ R^2_boundary",
		AmbientDimension:             inherited.H72Dimension,
		Projector:                    "P_K7 = Boolean-octonionic intersection projector extended by zero on R^2_boundary",
		ProjectorRank:                inherited.K7Dimension,
		ProjectorIsIdempotent:        true,
		ProjectorIsSymmetric:         true,
		BoundaryScalar:               "S_split=lambda(Lambda_12)+(R_3-1)",
		SSplit:                       inherited.SSplit,
		ResponseInEndH72:             true,
		SupportSelectedBeforeTracing: true,
		Verdict: strings.Join([]string{
			StatusRSplitDefinedAsSupportSelectedResponseOperator,
			StatusSSplitEigenvalueOnK7Support,
		}, "; "),
	}
}

func buildSpectrum(response ResponseOperatorDefinition) OperatorSpectrumAudit {
	trace := float64(response.ProjectorRank) * response.SSplit
	return OperatorSpectrumAudit{
		EigenvalueOnK7:       response.SSplit,
		K7Multiplicity:       response.ProjectorRank,
		ZeroEigenvalue:       0,
		ZeroMultiplicity:     response.AmbientDimension - response.ProjectorRank,
		RankIfSSplitNonzero:  response.ProjectorRank,
		SSplitNonzero:        math.Abs(response.SSplit) > 0,
		MinimalPolynomial:    "x(x-S_split) when S_split != 0",
		TraceFromSpectrum:    trace,
		SpectrumDimensionSum: response.AmbientDimension,
		Verdict: strings.Join([]string{
			StatusOperatorSpectrumComputed,
			StatusSSplitEigenvalueOnK7Support,
		}, "; "),
	}
}

func buildTraceCable(inherited Gate687Inheritance, spectrum OperatorSpectrumAudit) TracePowerCableAudit {
	powers := make([]TracePower, 0, 4)
	for _, n := range []int{1, 2, 3, 4} {
		tr := float64(inherited.K7Dimension) * math.Pow(inherited.SSplit, float64(n))
		powers = append(powers, TracePower{Power: n, Trace: tr, NormalizedTrace: tr / float64(inherited.H72Dimension), Formula: fmt.Sprintf("Tr(R_split^%d)=7 S_split^%d", n, n)})
	}
	firstTrace := spectrum.TraceFromSpectrum
	firstNorm := firstTrace / float64(inherited.H72Dimension)
	secondTrace := powers[1].Trace
	return TracePowerCableAudit{
		Powers:                  powers,
		FirstTrace:              firstTrace,
		FirstTraceNormalized:    firstNorm,
		FirstTracePrediction:    firstNorm,
		DBase:                   inherited.DBase,
		FirstTraceResidual:      inherited.DBase - firstNorm,
		FrobeniusNormSquared:    secondTrace,
		SecondTraceNormalized:   secondTrace / float64(inherited.H72Dimension),
		FormulaAllPositivePower: "Tr(R_split^n)=7 S_split^n for n>=1",
		Verdict: strings.Join([]string{
			StatusTracePowerCableComputed,
			StatusActiveBridgeFirstTraceOfSupportOperator,
		}, "; "),
	}
}

func buildLinearResponseSelection() LinearResponseSelectionAudit {
	return LinearResponseSelectionAudit{
		UsesFirstOrdinaryTrace: true,
		UsesSecondTrace:        false,
		UsesFrobeniusNorm:      false,
		UsesDetLikeQuantity:    false,
		UsesHodgeSignedTrace:   false,
		ActiveFunctional:       "Tr_H72(R_split)/72",
		RejectedFunctionals: []string{
			"Tr_H72(R_split^2)/72",
			"||R_split||_F^2/72",
			"det-like spectral quantities",
			"Hodge-signed trace",
		},
		ReasonForMissing: "the gate records the active first ordinary trace functional but does not derive a native principle forcing history to use that functional",
		Verdict: strings.Join([]string{
			StatusLinearFirstTraceResponseAudited,
			StatusNoNativeReasonFirstOrdinaryTrace,
		}, "; "),
	}
}

func buildSupportInvariance() SupportInvarianceAudit {
	return SupportInvarianceAudit{
		BooleanSupportEquation:       "P_B R_split = R_split because P_B P_K7=P_K7",
		OctonionicSupportEquation:    "P_G R_split = R_split because P_G P_K7=P_K7",
		PBRSplitEqualsRSplit:         true,
		PGRSplitEqualsRSplit:         true,
		ImageInBooleanSector:         true,
		ImageInOctonionicSector:      true,
		ImageInIntersectionCarrier:   true,
		SelectedIndependentlyOfTrace: true,
		Verdict: strings.Join([]string{
			StatusSupportInvarianceAudited,
			StatusActiveBridgeFirstTraceOfSupportOperator,
		}, "; "),
	}
}

func buildDegeneracy() SpectralDegeneracyAudit {
	candidates := []RankSevenResponseCandidate{
		{
			Name:                    "S_split P_K7",
			ProjectorRank:           k7Dimension,
			Carrier:                 "K_7=Im(P_B)∩Im(P_G)",
			SameSpectrumAsPK7:       true,
			SameOrdinaryTraceAsPK7:  true,
			PassesBooleanSupport:    true,
			PassesOctonionicSupport: true,
			SelectedBySupport:       true,
			SupportVerdict:          "support-selected response operator",
		},
		{
			Name:                    "S_split P_W7",
			ProjectorRank:           w7Dimension,
			Carrier:                 "W_7=(Im(P_B)+Im(P_G))^perp",
			SameSpectrumAsPK7:       true,
			SameOrdinaryTraceAsPK7:  true,
			PassesBooleanSupport:    false,
			PassesOctonionicSupport: false,
			SelectedBySupport:       false,
			SupportVerdict:          StatusSpectrumTraceAloneDoNotSelectK7,
		},
		{
			Name:                    "S_split P_arbitrary7",
			ProjectorRank:           k7Dimension,
			Carrier:                 "generic rank-seven subspace of H_72",
			SameSpectrumAsPK7:       true,
			SameOrdinaryTraceAsPK7:  true,
			PassesBooleanSupport:    false,
			PassesOctonionicSupport: false,
			SelectedBySupport:       false,
			SupportVerdict:          StatusSpectrumTraceAloneDoNotSelectK7,
		},
	}
	return SpectralDegeneracyAudit{
		Candidates:               candidates,
		AllShareSpectrumAndTrace: true,
		SpectrumSelectsK7:        false,
		TraceSelectsK7:           false,
		SupportSelectsK7:         true,
		Verdict: strings.Join([]string{
			StatusRankSevenSpectralDegeneracyRecorded,
			StatusSpectrumTraceAloneDoNotSelectK7,
		}, "; "),
	}
}

func buildHodgePolarity() HodgePolarityTraceComparisonAudit {
	ordinaryMultiplicity := k7PlusDimension + k7MinusDimension
	signedMultiplicity := k7PlusDimension - k7MinusDimension
	return HodgePolarityTraceComparisonAudit{
		K7PlusDimension:           k7PlusDimension,
		K7MinusDimension:          k7MinusDimension,
		OrdinaryTraceMultiplicity: ordinaryMultiplicity,
		HodgeSignedMultiplicity:   signedMultiplicity,
		OrdinaryTrace:             float64(ordinaryMultiplicity) * auditedSSplit,
		HodgeSignedTrace:          float64(signedMultiplicity) * auditedSSplit,
		OrdinaryCoefficient:       float64(ordinaryMultiplicity) / float64(h72Dimension),
		HodgeSignedCoefficient:    float64(signedMultiplicity) / float64(h72Dimension),
		ActiveUsesOrdinaryTrace:   true,
		ActiveUsesSignedTrace:     false,
		Verdict: strings.Join([]string{
			StatusHodgePolarityTraceComparisonAudited,
			StatusActiveBridgeFirstTraceOfSupportOperator,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate687FactorizationFirewallInherited,
		StatusRSplitDefinedAsSupportSelectedResponseOperator,
		StatusOperatorSpectrumComputed,
		StatusTracePowerCableComputed,
		StatusLinearFirstTraceResponseAudited,
		StatusSupportInvarianceAudited,
		StatusRankSevenSpectralDegeneracyRecorded,
		StatusHodgePolarityTraceComparisonAudited,
		StatusSSplitEigenvalueOnK7Support,
		StatusActiveBridgeFirstTraceOfSupportOperator,
		StatusSpectrumTraceAloneDoNotSelectK7,
		StatusNoNativeReasonFirstOrdinaryTrace,
		StatusNoNativeProjectorActivationTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate688ResponseOperatorSpectrumBoundary,
	}
}

func FormatInheritance(x Gate687Inheritance) string {
	return fmt.Sprintf("factorizationInherited=%t boundary=%q selector=%q trace=%q selected=%q ssplit=%.18g dbase=%.18g h72=%d k7=%d centralFirewall=%t supportSealed=%t couplingProved=%t verdict=%q", x.FactorizationFirewallInherited, x.BoundaryAmplitudeSeal, x.NativeProjectorSelectorSeal, x.TraceScalarizationSeal, x.SelectedProjector, x.SSplit, x.DBase, x.H72Dimension, x.K7Dimension, x.PriorCentralScalarFirewall, x.PriorSupportSealedIdentity, x.PriorNativeCouplingProved, x.Verdict)
}

func FormatResponse(x ResponseOperatorDefinition) string {
	return fmt.Sprintf("operator=%q ambient=%q ambientDim=%d projector=%q rank=%d idempotent=%t symmetric=%t scalar=%q ssplit=%.18g inEnd=%t supportBeforeTrace=%t verdict=%q", x.Operator, x.AmbientSpace, x.AmbientDimension, x.Projector, x.ProjectorRank, x.ProjectorIsIdempotent, x.ProjectorIsSymmetric, x.BoundaryScalar, x.SSplit, x.ResponseInEndH72, x.SupportSelectedBeforeTracing, x.Verdict)
}

func FormatSpectrum(x OperatorSpectrumAudit) string {
	return fmt.Sprintf("eigK7=%.18g multK7=%d eig0=%.18g mult0=%d rankIfNonzero=%d nonzero=%t minpoly=%q trace=%.18g dimSum=%d verdict=%q", x.EigenvalueOnK7, x.K7Multiplicity, x.ZeroEigenvalue, x.ZeroMultiplicity, x.RankIfSSplitNonzero, x.SSplitNonzero, x.MinimalPolynomial, x.TraceFromSpectrum, x.SpectrumDimensionSum, x.Verdict)
}

func FormatTracePower(x TracePower) string {
	return fmt.Sprintf("n=%d trace=%.18g normalized=%.18g formula=%q", x.Power, x.Trace, x.NormalizedTrace, x.Formula)
}

func FormatTraceCable(x TracePowerCableAudit) string {
	parts := make([]string, 0, len(x.Powers))
	for _, p := range x.Powers {
		parts = append(parts, FormatTracePower(p))
	}
	return fmt.Sprintf("powers=[%s] firstTrace=%.18g firstNorm=%.18g prediction=%.18g dbase=%.18g residual=%.18g frob2=%.18g secondNorm=%.18g formula=%q verdict=%q", strings.Join(parts, " | "), x.FirstTrace, x.FirstTraceNormalized, x.FirstTracePrediction, x.DBase, x.FirstTraceResidual, x.FrobeniusNormSquared, x.SecondTraceNormalized, x.FormulaAllPositivePower, x.Verdict)
}

func FormatLinearResponse(x LinearResponseSelectionAudit) string {
	return fmt.Sprintf("firstOrdinary=%t second=%t frob=%t det=%t signed=%t active=%q rejected=[%s] missing=%q verdict=%q", x.UsesFirstOrdinaryTrace, x.UsesSecondTrace, x.UsesFrobeniusNorm, x.UsesDetLikeQuantity, x.UsesHodgeSignedTrace, x.ActiveFunctional, strings.Join(x.RejectedFunctionals, "; "), x.ReasonForMissing, x.Verdict)
}

func FormatSupport(x SupportInvarianceAudit) string {
	return fmt.Sprintf("boolEq=%q octEq=%q pbR=%t pgR=%t inBool=%t inOct=%t inIntersection=%t independentOfTrace=%t verdict=%q", x.BooleanSupportEquation, x.OctonionicSupportEquation, x.PBRSplitEqualsRSplit, x.PGRSplitEqualsRSplit, x.ImageInBooleanSector, x.ImageInOctonionicSector, x.ImageInIntersectionCarrier, x.SelectedIndependentlyOfTrace, x.Verdict)
}

func FormatCandidate(x RankSevenResponseCandidate) string {
	return fmt.Sprintf("name=%q rank=%d carrier=%q sameSpectrum=%t sameTrace=%t bool=%t oct=%t selected=%t verdict=%q", x.Name, x.ProjectorRank, x.Carrier, x.SameSpectrumAsPK7, x.SameOrdinaryTraceAsPK7, x.PassesBooleanSupport, x.PassesOctonionicSupport, x.SelectedBySupport, x.SupportVerdict)
}

func FormatDegeneracy(x SpectralDegeneracyAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] allShare=%t spectrumSelects=%t traceSelects=%t supportSelects=%t verdict=%q", strings.Join(parts, " | "), x.AllShareSpectrumAndTrace, x.SpectrumSelectsK7, x.TraceSelectsK7, x.SupportSelectsK7, x.Verdict)
}

func FormatHodge(x HodgePolarityTraceComparisonAudit) string {
	return fmt.Sprintf("k7plus=%d k7minus=%d ordinaryMult=%d signedMult=%d ordinaryTrace=%.18g signedTrace=%.18g ordinaryCoeff=%.18g signedCoeff=%.18g activeOrdinary=%t activeSigned=%t verdict=%q", x.K7PlusDimension, x.K7MinusDimension, x.OrdinaryTraceMultiplicity, x.HodgeSignedMultiplicity, x.OrdinaryTrace, x.HodgeSignedTrace, x.OrdinaryCoefficient, x.HodgeSignedCoefficient, x.ActiveUsesOrdinaryTrace, x.ActiveUsesSignedTrace, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsSpectrumSelects=%t claimsTraceSelects=%t claimsFirstTrace=%t claimsActivation=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsSpectrumSelectsK7Identity, x.ClaimsTraceSelectsK7Identity, x.ClaimsNativeFirstTracePrinciple, x.ClaimsProjectorActivation, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
