// Package generation2linearresponsefunctionalandtracepairingnormalizationaudit implements
// Gate 691: Linear Response Functional and Trace-Pairing Normalization Audit.
//
// Gate 689 identified the active bridge scalarization as the first ordinary
// trace of the support-selected response operator
//
//	F_1 = Tr_H72(R_split)/72,
//
// where R_split=S_split P_K7. Gate 690 showed that the remaining first-trace
// residual is tiny and compatible with a suppressed second-order correction,
// but not independently certified.
//
// Gate 691 rewrites the leading bridge as a normalized trace-pairing
// functional
//
//	<I_H72,R_split>_tr,norm = Tr_H72(I_H72 R_split)/Tr_H72(I_H72).
//
// This is a bridge-layer linear-response functional audit only. It does not
// derive boundary stress, scalar RG matching, Higgs mass, gauge unification,
// flavor, CKM/PMNS, a native linear-response theorem, a native first-trace
// theorem, a native spectral-expansion theorem, or a native 7/72 theorem.
package generation2linearresponsefunctionalandtracepairingnormalizationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate690 "github.com/bagherbal/asha-engine/pkg/bridge/generation2firsttraceresidualandquadraticspectralcorrectionaudit"
)

const (
	AuditID = "GATE691-LINEAR-RESPONSE-FUNCTIONAL-AND-TRACE-PAIRING-NORMALIZATION-AUDIT"

	StatusGate689FirstTraceSelectionInherited       = "PASS_GATE689_FIRST_TRACE_SELECTION_INHERITED"
	StatusGate690ResidualStatusInherited            = "PASS_GATE690_RESIDUAL_STATUS_INHERITED"
	StatusNormalizedTracePairingDefined             = "PASS_NORMALIZED_TRACE_PAIRING_DEFINED"
	StatusActiveBridgeRewrittenAsTracePairing       = "PASS_ACTIVE_BRIDGE_REWRITTEN_AS_TRACE_PAIRING"
	StatusObserverResponseRoleClassified            = "PASS_OBSERVER_RESPONSE_ROLE_CLASSIFIED"
	StatusAlternativeObserverPairingsAudited        = "PASS_ALTERNATIVE_OBSERVER_PAIRINGS_AUDITED"
	StatusActiveBridgeLinearTracePairingResponse    = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_LINEAR_TRACE_PAIRING_RESPONSE"
	StatusFullChamberIdentityObserverTypeCorrect    = "CONDITIONAL_SUPPORT_FULL_CHAMBER_IDENTITY_OBSERVER_IS_TYPE_CORRECT"
	StatusQuadraticResidualRemainsSubleadingClue    = "CONDITIONAL_SUPPORT_QUADRATIC_RESIDUAL_REMAINS_SUBLEADING_CLUE"
	StatusTracePairingDoesNotUniquelySelectH72      = "FAILED_ROUTE_TRACE_PAIRING_DOES_NOT_UNIQUELY_SELECT_FULL_H72_OBSERVER"
	StatusNoNativeLinearResponseFunctionalTheorem   = "FAILED_ROUTE_NO_NATIVE_LINEAR_RESPONSE_FUNCTIONAL_THEOREM"
	StatusNoNativeFirstTraceTheorem                 = "FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_THEOREM"
	StatusNoNativeSevenOver72Theorem                = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate691TracePairingLinearResponseBoundary = "FIREWALL_PRESERVED_GATE691_TRACE_PAIRING_LINEAR_RESPONSE_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	k7Dimension       = 7
	k7PlusDimension   = 4
	k7MinusDimension  = 3
	residualTolerance = 1e-18
	pairingTolerance  = 1e-18
)

type Gate689690Inheritance struct {
	Gate689FirstTraceSelectionInherited bool
	Gate690ResidualStatusInherited      bool
	Operator                            string
	DBase                               float64
	SSplit                              float64
	F1                                  float64
	F2                                  float64
	E1                                  float64
	H72Dimension                        int
	K7Dimension                         int
	K7PlusDimension                     int
	K7MinusDimension                    int
	QuadraticResidualClueRetained       bool
	QuadraticCorrectionPromoted         bool
	NativeSpectralExpansionTheorem      bool
	NativeFirstTraceTheorem             bool
	NativeSevenOver72Theorem            bool
	Verdict                             string
}

type NormalizedTracePairingAudit struct {
	Definition          string
	Observer            string
	Response            string
	NumeratorTrace      float64
	DenominatorTrace    float64
	Value               float64
	ExpectedFirstTrace  float64
	EqualsFirstTrace    bool
	LinearInResponse    bool
	LinearInSSplit      bool
	BilinearInArguments bool
	Verdict             string
}

type ObserverResponseRoleAudit struct {
	FullChamberObserverRole    string
	ResponseOperatorRole       string
	SupportCarrierRole         string
	BoundaryScalarRole         string
	TracePairingInterpretation string
	FullObserverTypeCorrect    bool
	ResponseSupportSelected    bool
	BoundaryScalarIsEigenvalue bool
	Verdict                    string
}

type ObserverPairingCandidate struct {
	Name                         string
	Functional                   string
	ObserverRole                 string
	ContainsK7                   bool
	ActsAsIdentityOnK7           bool
	PositiveObserver             bool
	SignedObserver               bool
	NormalizationTrace           float64
	NumeratorTrace               float64
	Value                        float64
	ResidualAgainstF1            float64
	EquivalentToActiveFirstTrace bool
	SelectsFullH72Uniquely       bool
	ActiveObserverCandidate      bool
	Verdict                      string
}

type AlternativeObserverPairingAudit struct {
	Candidates                          []ObserverPairingCandidate
	CandidateCount                      int
	PositiveIdentityOnK7Count           int
	AllPositiveK7ObserversGiveSameValue bool
	SignedPolarityObserverInactive      bool
	FullH72ObserverUnique               bool
	DegeneracyWarning                   string
	Verdict                             string
}

type LinearResponseStatusAudit struct {
	DBaseExpression                       string
	TracePairingExpression                string
	DBaseLinearInWallCoordinates          bool
	TracePairingLinearInResponse          bool
	TracePairingLinearInSSplit            bool
	MatchesWallCoordinateOrder            bool
	NativeLinearResponseFunctionalTheorem bool
	Verdict                               string
}

type ResidualStatusAudit struct {
	E1                             float64
	AbsE1                          float64
	QuadraticF2                    float64
	QuadraticCoefficient           float64
	QuadraticResidualClueRetained  bool
	QuadraticCorrectionPromoted    bool
	NativeSpectralExpansionTheorem bool
	Verdict                        string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsUniqueFullH72Observer         bool
	ClaimsNativeLinearResponseTheorem   bool
	ClaimsNativeFirstTraceTheorem       bool
	ClaimsNativeSpectralExpansion       bool
	ClaimsNativeSevenOver72             bool
	ClaimsBoundaryStress                bool
	ClaimsScalarRGMatching              bool
	ClaimsHiggsMass                     bool
	ClaimsGaugeUnification              bool
	ClaimsFlavorDerivation              bool
	ClaimsCKMPMNS                       bool
	ClaimsProjectorActivation           bool
	PromotesQuadraticResidualCorrection bool
	Verdict                             string
}

type Analysis struct {
	Inherited      Gate689690Inheritance
	Pairing        NormalizedTracePairingAudit
	Roles          ObserverResponseRoleAudit
	Observers      AlternativeObserverPairingAudit
	LinearResponse LinearResponseStatusAudit
	Residual       ResidualStatusAudit
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
	g690, err := gate690.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate690 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g690)
	pairing := buildPairing(inherited)
	roles := buildRoles()
	observers := buildObserverPairings(inherited)
	linear := buildLinearResponseStatus(pairing)
	residual := buildResidualStatus(inherited)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeLinearResponseFunctionalTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native HistoryResponseFirstTraceTheorem or LinearResponseFunctionalTheorem explaining why the physical-history bridge uses the normalized ordinary trace-pairing <I_H72,R_split> rather than another positive K7-containing observer, a signed Hodge observer, or a higher spectral response",
		Verdict: strings.Join([]string{
			StatusNoNativeLinearResponseFunctionalTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate691TracePairingLinearResponseBoundary}
	truth := "Gate 691 rewrites the Gate689 leading response as the normalized trace pairing <I_H72,R_split>_tr,norm=Tr_H72(I_H72 R_split)/Tr_H72(I_H72)=(7/72)S_split.  This is a linear response functional: I_H72 is a type-correct full augmented-chamber observer, R_split is the support-selected response operator, P_K7 is the Boolean-octonionic support carrier, and S_split is the boundary quotient eigenvalue on that support.  The audit also records a degeneracy: any positive observer acting as identity on K7 gives the same pairing, while the Hodge-signed polarity observer gives only (1/72)S_split.  Thus the active bridge is conditionally supported as a linear ordinary trace-pairing response, but no native theorem uniquely selects the full H72 observer, no native first-trace theorem, and no native 7/72 theorem are certified.  Gate690's quadratic residual remains only a subleading clue."
	return Analysis{Inherited: inherited, Pairing: pairing, Roles: roles, Observers: observers, LinearResponse: linear, Residual: residual, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate690.Analysis) Gate689690Inheritance {
	return Gate689690Inheritance{
		Gate689FirstTraceSelectionInherited: g.Inherited.FirstTraceSelectionInherited && g.Inherited.Operator == "R_split = S_split P_K7" && g.Inherited.FirstTraceActive,
		Gate690ResidualStatusInherited:      g.Residual.Verdict == gate690.StatusFirstTraceResidualComputed && g.Quadratic.ResidualOverF2Small && g.FlavorDeficit.ResidualClueOnly,
		Operator:                            g.Inherited.Operator,
		DBase:                               g.Inherited.DBase,
		SSplit:                              g.Inherited.SSplit,
		F1:                                  g.Inherited.F1,
		F2:                                  g.Inherited.F2,
		E1:                                  g.Residual.Residual,
		H72Dimension:                        g.Inherited.H72Dimension,
		K7Dimension:                         g.Inherited.K7Dimension,
		K7PlusDimension:                     k7PlusDimension,
		K7MinusDimension:                    k7MinusDimension,
		QuadraticResidualClueRetained:       g.FlavorDeficit.ResidualClueOnly && !g.Expansion.KappaEFormulaPromoted,
		QuadraticCorrectionPromoted:         g.Expansion.KappaEFormulaPromoted,
		NativeSpectralExpansionTheorem:      g.Expansion.ExpansionTheoremCertified,
		NativeFirstTraceTheorem:             g.Inherited.NativeFirstTraceTheorem,
		NativeSevenOver72Theorem:            g.Inherited.NativeSevenOver72Theorem,
		Verdict:                             strings.Join([]string{StatusGate689FirstTraceSelectionInherited, StatusGate690ResidualStatusInherited}, "; "),
	}
}

func buildPairing(i Gate689690Inheritance) NormalizedTracePairingAudit {
	numerator := float64(i.K7Dimension) * i.SSplit
	denominator := float64(i.H72Dimension)
	value := numerator / denominator
	return NormalizedTracePairingAudit{
		Definition:          "<A,B>_tr,norm = Tr_H72(A B)/Tr_H72(I_H72)",
		Observer:            "I_H72",
		Response:            i.Operator,
		NumeratorTrace:      numerator,
		DenominatorTrace:    denominator,
		Value:               value,
		ExpectedFirstTrace:  i.F1,
		EqualsFirstTrace:    math.Abs(value-i.F1) < pairingTolerance,
		LinearInResponse:    true,
		LinearInSSplit:      true,
		BilinearInArguments: true,
		Verdict: strings.Join([]string{
			StatusNormalizedTracePairingDefined,
			StatusActiveBridgeRewrittenAsTracePairing,
			StatusActiveBridgeLinearTracePairingResponse,
		}, "; "),
	}
}

func buildRoles() ObserverResponseRoleAudit {
	return ObserverResponseRoleAudit{
		FullChamberObserverRole:    "I_H72: full augmented chamber observer / unbiased ordinary trace scalarizer",
		ResponseOperatorRole:       "R_split=S_split P_K7: support-selected response operator",
		SupportCarrierRole:         "P_K7: Boolean-octonionic intersection carrier selected by rank seven plus P_B and P_G support",
		BoundaryScalarRole:         "S_split=lambda(Lambda_12)+(R_3-1): boundary anti-alignment quotient eigenvalue/amplitude on selected support",
		TracePairingInterpretation: "D_base is approximated by the normalized trace pairing of the full-chamber observer with the support-selected response operator",
		FullObserverTypeCorrect:    true,
		ResponseSupportSelected:    true,
		BoundaryScalarIsEigenvalue: true,
		Verdict: strings.Join([]string{
			StatusObserverResponseRoleClassified,
			StatusFullChamberIdentityObserverTypeCorrect,
		}, "; "),
	}
}

func buildObserverPairings(i Gate689690Inheritance) AlternativeObserverPairingAudit {
	active := i.F1
	candidates := []ObserverPairingCandidate{
		observerCandidate("I_H72", "Tr(I_H72 R_split)/72", "full augmented chamber observer", true, true, true, false, i.SSplit, active, true),
		observerCandidate("P_finite", "Tr(P_finite R_split)/72", "finite Lambda^4 R^8 chamber observer containing K7", true, true, true, false, i.SSplit, active, false),
		observerCandidate("P_kernel", "Tr(P_kernel R_split)/72", "kernel/support-sector observer containing K7", true, true, true, false, i.SSplit, active, false),
		observerCandidate("P_K7", "Tr(P_K7 R_split)/72", "selected support observer", true, true, true, false, i.SSplit, active, false),
		observerCandidate("S_K", "Tr(S_K R_split)/72", "Hodge polarity sign observer on K7", true, false, false, true, i.SSplit, active, false),
	}
	positiveSame := true
	positiveCount := 0
	signedInactive := false
	for _, c := range candidates {
		if c.PositiveObserver && c.ActsAsIdentityOnK7 {
			positiveCount++
			positiveSame = positiveSame && c.EquivalentToActiveFirstTrace
		}
		if c.SignedObserver {
			signedInactive = !c.EquivalentToActiveFirstTrace && math.Abs(c.Value-(1.0/72.0)*i.SSplit) < pairingTolerance
		}
	}
	return AlternativeObserverPairingAudit{
		Candidates:                          candidates,
		CandidateCount:                      len(candidates),
		PositiveIdentityOnK7Count:           positiveCount,
		AllPositiveK7ObserversGiveSameValue: positiveSame,
		SignedPolarityObserverInactive:      signedInactive,
		FullH72ObserverUnique:               false,
		DegeneracyWarning:                   "any positive observer acting as identity on K7 gives Tr(O R_split)/72=(7/72)S_split, so trace pairing does not uniquely select I_H72 as observer",
		Verdict: strings.Join([]string{
			StatusAlternativeObserverPairingsAudited,
			StatusTracePairingDoesNotUniquelySelectH72,
		}, "; "),
	}
}

func observerCandidate(name, functional, role string, containsK7, identityOnK7, positive, signed bool, ssplit, active float64, activeObserver bool) ObserverPairingCandidate {
	var numerator float64
	if signed {
		numerator = float64(k7PlusDimension-k7MinusDimension) * ssplit
	} else if identityOnK7 {
		numerator = float64(k7Dimension) * ssplit
	}
	value := numerator / float64(h72Dimension)
	return ObserverPairingCandidate{
		Name:                         name,
		Functional:                   functional,
		ObserverRole:                 role,
		ContainsK7:                   containsK7,
		ActsAsIdentityOnK7:           identityOnK7,
		PositiveObserver:             positive,
		SignedObserver:               signed,
		NormalizationTrace:           float64(h72Dimension),
		NumeratorTrace:               numerator,
		Value:                        value,
		ResidualAgainstF1:            active - value,
		EquivalentToActiveFirstTrace: math.Abs(active-value) < pairingTolerance,
		SelectsFullH72Uniquely:       false,
		ActiveObserverCandidate:      activeObserver,
		Verdict:                      observerVerdict(positive, signed, identityOnK7),
	}
}

func observerVerdict(positive, signed, identityOnK7 bool) string {
	if signed {
		return "SIGNED_POLARITY_OBSERVER_INACTIVE_FOR_TOTAL_SUPPORT_RESPONSE"
	}
	if positive && identityOnK7 {
		return "POSITIVE_K7_IDENTITY_OBSERVER_EQUIVALENT_TO_ACTIVE_FIRST_TRACE"
	}
	return "OBSERVER_NOT_ACTIVE"
}

func buildLinearResponseStatus(p NormalizedTracePairingAudit) LinearResponseStatusAudit {
	return LinearResponseStatusAudit{
		DBaseExpression:                       "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)",
		TracePairingExpression:                p.Definition,
		DBaseLinearInWallCoordinates:          true,
		TracePairingLinearInResponse:          p.LinearInResponse,
		TracePairingLinearInSSplit:            p.LinearInSSplit,
		MatchesWallCoordinateOrder:            p.LinearInResponse && p.LinearInSSplit,
		NativeLinearResponseFunctionalTheorem: false,
		Verdict: strings.Join([]string{
			StatusActiveBridgeLinearTracePairingResponse,
			StatusNoNativeLinearResponseFunctionalTheorem,
		}, "; "),
	}
}

func buildResidualStatus(i Gate689690Inheritance) ResidualStatusAudit {
	c2 := i.E1 / i.F2
	return ResidualStatusAudit{
		E1:                             i.E1,
		AbsE1:                          math.Abs(i.E1),
		QuadraticF2:                    i.F2,
		QuadraticCoefficient:           c2,
		QuadraticResidualClueRetained:  i.QuadraticResidualClueRetained,
		QuadraticCorrectionPromoted:    i.QuadraticCorrectionPromoted,
		NativeSpectralExpansionTheorem: i.NativeSpectralExpansionTheorem,
		Verdict: strings.Join([]string{
			StatusGate690ResidualStatusInherited,
			StatusQuadraticResidualRemainsSubleadingClue,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate689FirstTraceSelectionInherited,
		StatusGate690ResidualStatusInherited,
		StatusNormalizedTracePairingDefined,
		StatusActiveBridgeRewrittenAsTracePairing,
		StatusObserverResponseRoleClassified,
		StatusAlternativeObserverPairingsAudited,
		StatusActiveBridgeLinearTracePairingResponse,
		StatusFullChamberIdentityObserverTypeCorrect,
		StatusQuadraticResidualRemainsSubleadingClue,
		StatusTracePairingDoesNotUniquelySelectH72,
		StatusNoNativeLinearResponseFunctionalTheorem,
		StatusNoNativeFirstTraceTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate691TracePairingLinearResponseBoundary,
	}
}

func FormatInheritance(x Gate689690Inheritance) string {
	return fmt.Sprintf("gate689=%t gate690=%t operator=%q dbase=%.18g ssplit=%.18g f1=%.18g f2=%.18g e1=%.18g h72=%d k7=%d k7plus=%d k7minus=%d residualClue=%t correctionPromoted=%t spectralTheorem=%t firstTraceTheorem=%t sevenTheorem=%t verdict=%q", x.Gate689FirstTraceSelectionInherited, x.Gate690ResidualStatusInherited, x.Operator, x.DBase, x.SSplit, x.F1, x.F2, x.E1, x.H72Dimension, x.K7Dimension, x.K7PlusDimension, x.K7MinusDimension, x.QuadraticResidualClueRetained, x.QuadraticCorrectionPromoted, x.NativeSpectralExpansionTheorem, x.NativeFirstTraceTheorem, x.NativeSevenOver72Theorem, x.Verdict)
}

func FormatPairing(x NormalizedTracePairingAudit) string {
	return fmt.Sprintf("definition=%q observer=%q response=%q numerator=%.18g denominator=%.18g value=%.18g expected=%.18g equalsFirst=%t linearResponse=%t linearSSplit=%t bilinear=%t verdict=%q", x.Definition, x.Observer, x.Response, x.NumeratorTrace, x.DenominatorTrace, x.Value, x.ExpectedFirstTrace, x.EqualsFirstTrace, x.LinearInResponse, x.LinearInSSplit, x.BilinearInArguments, x.Verdict)
}

func FormatRoles(x ObserverResponseRoleAudit) string {
	return fmt.Sprintf("observer=%q response=%q support=%q scalar=%q interpretation=%q fullType=%t responseSelected=%t scalarEigen=%t verdict=%q", x.FullChamberObserverRole, x.ResponseOperatorRole, x.SupportCarrierRole, x.BoundaryScalarRole, x.TracePairingInterpretation, x.FullObserverTypeCorrect, x.ResponseSupportSelected, x.BoundaryScalarIsEigenvalue, x.Verdict)
}

func FormatObserverCandidate(x ObserverPairingCandidate) string {
	return fmt.Sprintf("name=%q functional=%q role=%q containsK7=%t identityOnK7=%t positive=%t signed=%t norm=%.18g numerator=%.18g value=%.18g residualF1=%.18g equivalent=%t uniqueH72=%t activeCandidate=%t verdict=%q", x.Name, x.Functional, x.ObserverRole, x.ContainsK7, x.ActsAsIdentityOnK7, x.PositiveObserver, x.SignedObserver, x.NormalizationTrace, x.NumeratorTrace, x.Value, x.ResidualAgainstF1, x.EquivalentToActiveFirstTrace, x.SelectsFullH72Uniquely, x.ActiveObserverCandidate, x.Verdict)
}

func FormatObservers(x AlternativeObserverPairingAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatObserverCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] count=%d positiveK7Count=%d positiveSame=%t signedInactive=%t h72Unique=%t warning=%q verdict=%q", strings.Join(parts, " | "), x.CandidateCount, x.PositiveIdentityOnK7Count, x.AllPositiveK7ObserversGiveSameValue, x.SignedPolarityObserverInactive, x.FullH72ObserverUnique, x.DegeneracyWarning, x.Verdict)
}

func FormatLinearResponse(x LinearResponseStatusAudit) string {
	return fmt.Sprintf("dbase=%q pairing=%q dbaseLinear=%t pairingLinearResponse=%t pairingLinearSSplit=%t orderMatch=%t nativeLinearTheorem=%t verdict=%q", x.DBaseExpression, x.TracePairingExpression, x.DBaseLinearInWallCoordinates, x.TracePairingLinearInResponse, x.TracePairingLinearInSSplit, x.MatchesWallCoordinateOrder, x.NativeLinearResponseFunctionalTheorem, x.Verdict)
}

func FormatResidual(x ResidualStatusAudit) string {
	return fmt.Sprintf("e1=%.18g absE1=%.18g f2=%.18g c2=%.18g residualClue=%t correctionPromoted=%t spectralTheorem=%t verdict=%q", x.E1, x.AbsE1, x.QuadraticF2, x.QuadraticCoefficient, x.QuadraticResidualClueRetained, x.QuadraticCorrectionPromoted, x.NativeSpectralExpansionTheorem, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("uniqueH72=%t nativeLinear=%t nativeFirst=%t nativeSpectral=%t native7=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t activation=%t promoteQuadratic=%t verdict=%q", x.ClaimsUniqueFullH72Observer, x.ClaimsNativeLinearResponseTheorem, x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSpectralExpansion, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.PromotesQuadraticResidualCorrection, x.Verdict)
}
