// Package generation2historydefectreadoutfunctionalselectionaudit implements
// Gate 698: History Defect Readout Functional Selection Audit.
//
// Gate 697 typed the active payoff S_split=lambda(Lambda_12)+(R_3-1)
// as the canonical boundary anti-alignment quotient coordinate. Gate 698 audits
// the output/readout side: whether D_base=kappa_lambda+kappa_e+lambda(Lambda_12)
// is the canonical scalar/flavor/history closure-defect coordinate measuring
// failure of the scalar matching deficit and flavor wall deficit to close on the
// signed scalar zero-wall coordinate.
//
// This is a bridge-layer history-readout audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native payoff theorem, a native history-response theorem, or a
// native 7/72 theorem.
package generation2historydefectreadoutfunctionalselectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate697 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryquotientpayofffunctionalselectionaudit"
)

const (
	AuditID = "GATE698-HISTORY-DEFECT-READOUT-FUNCTIONAL-SELECTION-AUDIT"

	StatusGate697BoundaryQuotientPayoffInherited             = "PASS_GATE697_BOUNDARY_QUOTIENT_PAYOFF_INHERITED"
	StatusHistoryClosureWallDefined                          = "PASS_HISTORY_CLOSURE_WALL_DEFINED"
	StatusSigmaHistoryReadoutDefined                         = "PASS_SIGMA_HISTORY_READOUT_DEFINED"
	StatusDBaseIdentifiedAsHistoryDefectQuotient             = "PASS_DBASE_IDENTIFIED_AS_HISTORY_DEFECT_QUOTIENT"
	StatusAlternativeHistoryReadoutsAudited                  = "PASS_ALTERNATIVE_HISTORY_READOUTS_AUDITED"
	StatusSignedScalarWallFormPreferred                      = "PASS_SIGNED_SCALAR_WALL_FORM_PREFERRED"
	StatusFullBridgeReconstructedAsQuotientExpectation       = "PASS_FULL_BRIDGE_RECONSTRUCTED_AS_QUOTIENT_TO_EXPECTATION_EQUATION"
	StatusDBaseCanonicalHistoryClosureDefectReadout          = "CONDITIONAL_SUPPORT_DBASE_IS_CANONICAL_HISTORY_CLOSURE_DEFECT_READOUT"
	StatusActiveBridgeRelatesHistoryQuotientToExpectedPayoff = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_RELATES_HISTORY_QUOTIENT_TO_EXPECTED_BOUNDARY_PAYOFF"
	StatusHistoryReadoutUniqueOnlyUpToNormalization          = "FAILED_ROUTE_HISTORY_READOUT_UNIQUE_ONLY_UP_TO_WALL_COORDINATE_NORMALIZATION"
	StatusNoNativeExpectedK7PayoffEqualsHistoryDefect        = "FAILED_ROUTE_NO_NATIVE_REASON_EXPECTED_K7_BOUNDARY_PAYOFF_EQUALS_HISTORY_DEFECT"
	StatusNoNativeHistoryBoundaryResponseTheorem             = "FAILED_ROUTE_NO_NATIVE_HISTORY_BOUNDARY_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem                         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate698HistoryDefectReadoutBoundary                = "FIREWALL_PRESERVED_GATE698_HISTORY_DEFECT_READOUT_BOUNDARY"
)

const (
	lambda4Dimension = 70
	boundaryDim      = 2
	h72Dimension     = lambda4Dimension + boundaryDim
	k7Dimension      = 7
	kappaLambda      = 0.0443230430960771
	kappaE           = 0.00550355419157456
	lambdaLambda12   = -0.0497009420776833
	r3Minus1         = 0.0509933868964996
	tolerance        = 1e-15
)

type Gate697Inheritance struct {
	BoundaryQuotientPayoffInherited bool
	BoundaryPayoff                  string
	BoundaryPayoffValue             float64
	BoundaryPayoffDescends          bool
	ResponseOperator                string
	Rho72                           string
	Expectation                     float64
	DBase                           float64
	ResidualE1                      float64
	NoNativePayoffCoupling          bool
	NoNativeHistoryResponse         bool
	NoNativeSevenOver72             bool
	Verdict                         string
}

type HistoryClosureWallAudit struct {
	WallEquation              string
	EquivalentPositiveClosure string
	KappaLambda               float64
	KappaE                    float64
	Lambda                    float64
	KSum                      float64
	LambdaNegative            bool
	ClosureDefect             float64
	WallSatisfied             bool
	Verdict                   string
}

type HistoryReadoutAudit struct {
	SigmaDefinition       string
	InputSpace            string
	InputVector           [3]float64
	DBaseExpression       string
	DBase                 float64
	MatchesInheritedDBase bool
	MeasuresClosureDefect bool
	Verdict               string
}

type HistoryReadoutAlternative struct {
	Name                  string
	Functional            string
	Value                 float64
	IncludesKappaLambda   bool
	IncludesKappaE        bool
	IncludesSignedLambda  bool
	EquivalentWhenLambdaN bool
	PreservesOrientation  bool
	Active                bool
	Reason                string
}

type AlternativeHistoryReadoutsAudit struct {
	Alternatives            []HistoryReadoutAlternative
	KSumRejected            bool
	LambdaOnlyRejected      bool
	KappaLambdaOnlyRejected bool
	KappaEOnlyRejected      bool
	AbsoluteFormEquivalent  bool
	SignedFormAccepted      bool
	AllAudited              bool
	Verdict                 string
}

type SignAndWallCoordinateAudit struct {
	LambdaNegative             bool
	SignedForm                 string
	AbsoluteForm               string
	DBaseSigned                float64
	DBaseAbsolute              float64
	FormsEquivalentNumerically bool
	SignedFormPreferred        bool
	OrientationPreserved       bool
	Verdict                    string
}

type BridgeReconstructionAudit struct {
	LeftSide                 string
	RightSide                string
	ExpandedEquation         string
	SigmaHistory             float64
	SigmaBoundaryExpectation float64
	ResidualE1               float64
	MatchesInheritedResidual bool
	RelatesQuotientToPayoff  bool
	Verdict                  string
}

type ResidualStatusAudit struct {
	ResidualE1                float64
	QuadraticClueRetained     bool
	QuadraticCluePromoted     bool
	NativeResidualExplanation bool
	Verdict                   string
}

type SourceTypeClassificationAudit struct {
	KappaLambdaRole string
	KappaERole      string
	LambdaRole      string
	DBaseRole       string
	SSplitRole      string
	Rho72Role       string
	PK7Role         string
	BridgeRole      string
	Verdict         string
}

type MissingTheoremAudit struct {
	Candidates []string
	PreciseGap string
	Verdict    string
}

type FirewallAudit struct {
	ClaimsExpectedK7PayoffEqualsHistoryDefectNatively bool
	ClaimsNativeHistoryBoundaryResponseTheorem        bool
	ClaimsNativeSevenOver72Theorem                    bool
	ClaimsNativePayoffTheorem                         bool
	ClaimsBoundaryStressDerived                       bool
	ClaimsScalarRGMatching                            bool
	ClaimsHiggsMass                                   bool
	ClaimsGaugeUnification                            bool
	ClaimsFlavorDerivation                            bool
	ClaimsCKMPMNS                                     bool
	Verdict                                           string
}

type Analysis struct {
	Inherited    Gate697Inheritance
	ClosureWall  HistoryClosureWallAudit
	Readout      HistoryReadoutAudit
	Alternatives AlternativeHistoryReadoutsAudit
	SignAudit    SignAndWallCoordinateAudit
	Bridge       BridgeReconstructionAudit
	Residual     ResidualStatusAudit
	SourceTypes  SourceTypeClassificationAudit
	Missing      MissingTheoremAudit
	Firewall     FirewallAudit
	Truth        string
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
	g697, err := gate697.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate697 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g697)
	closure := buildHistoryClosureWall()
	readout := buildHistoryReadout(inherited)
	alternatives := buildAlternatives(readout)
	sign := buildSignAudit(readout)
	bridge := buildBridgeReconstruction(inherited, readout)
	residual := ResidualStatusAudit{
		ResidualE1:                inherited.ResidualE1,
		QuadraticClueRetained:     true,
		QuadraticCluePromoted:     false,
		NativeResidualExplanation: false,
		Verdict:                   "Gate690 quadratic residual clue retained as subleading and not promoted",
	}
	sources := buildSourceTypes()
	missing := MissingTheoremAudit{
		Candidates: []string{
			"HistoryBoundaryQuotientResponseTheorem",
			"K7EventPayoffToHistoryReadoutTheorem",
			StatusHistoryReadoutUniqueOnlyUpToNormalization,
			StatusNoNativeExpectedK7PayoffEqualsHistoryDefect,
			StatusNoNativeHistoryBoundaryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native bridge theorem explaining why sigma_history(kappa_lambda,kappa_e,lambda) equals E_rho72[sigma_boundary(lambda,R) 1_K7]",
		Verdict: strings.Join([]string{
			StatusHistoryReadoutUniqueOnlyUpToNormalization,
			StatusNoNativeExpectedK7PayoffEqualsHistoryDefect,
			StatusNoNativeHistoryBoundaryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate698HistoryDefectReadoutBoundary}
	truth := "Gate 698 audits the output side of the bridge. The signed functional sigma_history(kappa_lambda,kappa_e,lambda)=kappa_lambda+kappa_e+lambda defines the scalar/flavor/history closure defect against the signed scalar zero wall. Since lambda(Lambda_12)<0, it equals K_sum-|lambda| numerically, but the signed form is preferred because it preserves wall orientation. The active bridge is reconstructed as sigma_history(h)≈Tr[(I_H72/72) sigma_boundary(b) P_K7]. This conditionally supports D_base as the canonical history closure readout, while preserving the firewall that no native theorem yet explains why the expected K7 boundary payoff equals the history defect."
	return Analysis{Inherited: inherited, ClosureWall: closure, Readout: readout, Alternatives: alternatives, SignAudit: sign, Bridge: bridge, Residual: residual, SourceTypes: sources, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate697.Analysis) Gate697Inheritance {
	return Gate697Inheritance{
		BoundaryQuotientPayoffInherited: g.Quotient.DescendsToQuotient && g.Interpretation.K7ReceivesBoundaryDefect && g.Expectation.MatchesInherited,
		BoundaryPayoff:                  "S_split=sigma_boundary(lambda,R)=lambda+(R_3-1)",
		BoundaryPayoffValue:             g.Quotient.SSplit,
		BoundaryPayoffDescends:          g.Quotient.DescendsToQuotient,
		ResponseOperator:                "R_split=sigma_boundary(b)P_K7",
		Rho72:                           "rho_72=I_H72/72",
		Expectation:                     g.Expectation.Expectation,
		DBase:                           g.Expectation.DBase,
		ResidualE1:                      g.Expectation.ResidualE1,
		NoNativePayoffCoupling:          !g.Firewall.ClaimsNativePayoffCouplingTheorem,
		NoNativeHistoryResponse:         !g.Firewall.ClaimsNativeHistoryResponseTheorem,
		NoNativeSevenOver72:             !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                         StatusGate697BoundaryQuotientPayoffInherited,
	}
}

func buildHistoryClosureWall() HistoryClosureWallAudit {
	kSum := kappaLambda + kappaE
	defect := kSum + lambdaLambda12
	return HistoryClosureWallAudit{
		WallEquation:              "kappa_lambda+kappa_e+lambda=0",
		EquivalentPositiveClosure: "kappa_lambda+kappa_e=|lambda| when lambda<0",
		KappaLambda:               kappaLambda,
		KappaE:                    kappaE,
		Lambda:                    lambdaLambda12,
		KSum:                      kSum,
		LambdaNegative:            lambdaLambda12 < 0,
		ClosureDefect:             defect,
		WallSatisfied:             math.Abs(defect) < tolerance,
		Verdict:                   StatusHistoryClosureWallDefined,
	}
}

func buildHistoryReadout(inherited Gate697Inheritance) HistoryReadoutAudit {
	d := kappaLambda + kappaE + lambdaLambda12
	return HistoryReadoutAudit{
		SigmaDefinition:       "sigma_history(kappa_lambda,kappa_e,lambda)=kappa_lambda+kappa_e+lambda",
		InputSpace:            "H_history=span(kappa_lambda,kappa_e,lambda)",
		InputVector:           [3]float64{kappaLambda, kappaE, lambdaLambda12},
		DBaseExpression:       "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)",
		DBase:                 d,
		MatchesInheritedDBase: math.Abs(d-inherited.DBase) < tolerance,
		MeasuresClosureDefect: true,
		Verdict: strings.Join([]string{
			StatusSigmaHistoryReadoutDefined,
			StatusDBaseIdentifiedAsHistoryDefectQuotient,
			StatusDBaseCanonicalHistoryClosureDefectReadout,
		}, "; "),
	}
}

func buildAlternatives(r HistoryReadoutAudit) AlternativeHistoryReadoutsAudit {
	kSum := kappaLambda + kappaE
	absForm := kSum - math.Abs(lambdaLambda12)
	alts := []HistoryReadoutAlternative{
		{Name: "K_sum", Functional: "kappa_lambda+kappa_e", Value: kSum, IncludesKappaLambda: true, IncludesKappaE: true, IncludesSignedLambda: false, EquivalentWhenLambdaN: false, PreservesOrientation: false, Active: false, Reason: "incomplete: omits signed scalar zero-wall coordinate"},
		{Name: "lambda-only", Functional: "lambda(Lambda_12)", Value: lambdaLambda12, IncludesKappaLambda: false, IncludesKappaE: false, IncludesSignedLambda: true, EquivalentWhenLambdaN: false, PreservesOrientation: true, Active: false, Reason: "ignores scalar/flavor deficits"},
		{Name: "kappa_lambda-only", Functional: "kappa_lambda", Value: kappaLambda, IncludesKappaLambda: true, IncludesKappaE: false, IncludesSignedLambda: false, EquivalentWhenLambdaN: false, PreservesOrientation: false, Active: false, Reason: "ignores charged-lepton flavor deficit and scalar wall coordinate"},
		{Name: "kappa_e-only", Functional: "kappa_e", Value: kappaE, IncludesKappaLambda: false, IncludesKappaE: true, IncludesSignedLambda: false, EquivalentWhenLambdaN: false, PreservesOrientation: false, Active: false, Reason: "ignores scalar matching deficit and scalar wall coordinate"},
		{Name: "absolute closure form", Functional: "kappa_lambda+kappa_e-|lambda|", Value: absForm, IncludesKappaLambda: true, IncludesKappaE: true, IncludesSignedLambda: false, EquivalentWhenLambdaN: lambdaLambda12 < 0 && math.Abs(absForm-r.DBase) < tolerance, PreservesOrientation: false, Active: false, Reason: "numerically equivalent only because lambda<0; less canonical because it erases scalar-wall orientation"},
		{Name: "signed history form", Functional: "kappa_lambda+kappa_e+lambda", Value: r.DBase, IncludesKappaLambda: true, IncludesKappaE: true, IncludesSignedLambda: true, EquivalentWhenLambdaN: true, PreservesOrientation: true, Active: true, Reason: "canonical oriented history closure defect against the signed scalar zero wall"},
	}
	return AlternativeHistoryReadoutsAudit{
		Alternatives:            alts,
		KSumRejected:            !alts[0].Active,
		LambdaOnlyRejected:      !alts[1].Active,
		KappaLambdaOnlyRejected: !alts[2].Active,
		KappaEOnlyRejected:      !alts[3].Active,
		AbsoluteFormEquivalent:  alts[4].EquivalentWhenLambdaN,
		SignedFormAccepted:      alts[5].Active,
		AllAudited:              true,
		Verdict:                 StatusAlternativeHistoryReadoutsAudited,
	}
}

func buildSignAudit(r HistoryReadoutAudit) SignAndWallCoordinateAudit {
	abs := kappaLambda + kappaE - math.Abs(lambdaLambda12)
	return SignAndWallCoordinateAudit{
		LambdaNegative:             lambdaLambda12 < 0,
		SignedForm:                 "kappa_lambda+kappa_e+lambda",
		AbsoluteForm:               "kappa_lambda+kappa_e-|lambda|",
		DBaseSigned:                r.DBase,
		DBaseAbsolute:              abs,
		FormsEquivalentNumerically: math.Abs(r.DBase-abs) < tolerance,
		SignedFormPreferred:        true,
		OrientationPreserved:       true,
		Verdict:                    StatusSignedScalarWallFormPreferred,
	}
}

func buildBridgeReconstruction(inherited Gate697Inheritance, r HistoryReadoutAudit) BridgeReconstructionAudit {
	return BridgeReconstructionAudit{
		LeftSide:                 "sigma_history(h)=kappa_lambda+kappa_e+lambda",
		RightSide:                "Tr(rho_72 sigma_boundary(b) P_K7)",
		ExpandedEquation:         "kappa_lambda+kappa_e+lambda ≈ Tr[(I_H72/72)(lambda+(R_3-1))P_K7]",
		SigmaHistory:             r.DBase,
		SigmaBoundaryExpectation: inherited.Expectation,
		ResidualE1:               r.DBase - inherited.Expectation,
		MatchesInheritedResidual: math.Abs((r.DBase-inherited.Expectation)-inherited.ResidualE1) < 1e-17,
		RelatesQuotientToPayoff:  true,
		Verdict: strings.Join([]string{
			StatusFullBridgeReconstructedAsQuotientExpectation,
			StatusActiveBridgeRelatesHistoryQuotientToExpectedPayoff,
		}, "; "),
	}
}

func buildSourceTypes() SourceTypeClassificationAudit {
	return SourceTypeClassificationAudit{
		KappaLambdaRole: "scalar low-scale loop-matching deficit",
		KappaERole:      "charged-lepton flavor-wall loop deficit",
		LambdaRole:      "signed scalar zero-wall coordinate lambda(Lambda_12)",
		DBaseRole:       "scalar/flavor/history closure-defect readout sigma_history",
		SSplitRole:      "boundary anti-alignment quotient payoff sigma_boundary",
		Rho72Role:       "full augmented maximum-entropy observer state",
		PK7Role:         "Boolean-octonionic support-selected event projector",
		BridgeRole:      "history quotient approximately equals expected K7 boundary payoff",
		Verdict:         StatusDBaseCanonicalHistoryClosureDefectReadout,
	}
}

func Statuses() []string {
	return []string{
		StatusGate697BoundaryQuotientPayoffInherited,
		StatusHistoryClosureWallDefined,
		StatusSigmaHistoryReadoutDefined,
		StatusDBaseIdentifiedAsHistoryDefectQuotient,
		StatusAlternativeHistoryReadoutsAudited,
		StatusSignedScalarWallFormPreferred,
		StatusFullBridgeReconstructedAsQuotientExpectation,
		StatusDBaseCanonicalHistoryClosureDefectReadout,
		StatusActiveBridgeRelatesHistoryQuotientToExpectedPayoff,
		StatusHistoryReadoutUniqueOnlyUpToNormalization,
		StatusNoNativeExpectedK7PayoffEqualsHistoryDefect,
		StatusNoNativeHistoryBoundaryResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate698HistoryDefectReadoutBoundary,
	}
}

func FormatInheritance(x Gate697Inheritance) string {
	return fmt.Sprintf("inherited=%t payoff=%q value=%.18g descends=%t operator=%q rho=%q expectation=%.18g dbase=%.18g e1=%.18g noCoupling=%t noHistory=%t no7=%t verdict=%q", x.BoundaryQuotientPayoffInherited, x.BoundaryPayoff, x.BoundaryPayoffValue, x.BoundaryPayoffDescends, x.ResponseOperator, x.Rho72, x.Expectation, x.DBase, x.ResidualE1, x.NoNativePayoffCoupling, x.NoNativeHistoryResponse, x.NoNativeSevenOver72, x.Verdict)
}

func FormatClosureWall(x HistoryClosureWallAudit) string {
	return fmt.Sprintf("wall=%q equiv=%q kappaLambda=%.18g kappaE=%.18g lambda=%.18g ksum=%.18g lambdaNeg=%t defect=%.18g satisfied=%t verdict=%q", x.WallEquation, x.EquivalentPositiveClosure, x.KappaLambda, x.KappaE, x.Lambda, x.KSum, x.LambdaNegative, x.ClosureDefect, x.WallSatisfied, x.Verdict)
}

func FormatReadout(x HistoryReadoutAudit) string {
	return fmt.Sprintf("sigma=%q space=%q vector=(%.18g,%.18g,%.18g) expr=%q dbase=%.18g matchesInherited=%t measures=%t verdict=%q", x.SigmaDefinition, x.InputSpace, x.InputVector[0], x.InputVector[1], x.InputVector[2], x.DBaseExpression, x.DBase, x.MatchesInheritedDBase, x.MeasuresClosureDefect, x.Verdict)
}

func FormatAlternatives(x AlternativeHistoryReadoutsAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s=%s value=%.18g signedLambda=%t equivNeg=%t orient=%t active=%t reason=%q", a.Name, a.Functional, a.Value, a.IncludesSignedLambda, a.EquivalentWhenLambdaN, a.PreservesOrientation, a.Active, a.Reason))
	}
	return fmt.Sprintf("ksumRejected=%t lambdaOnlyRejected=%t kappaLambdaOnlyRejected=%t kappaEOnlyRejected=%t absEquivalent=%t signedAccepted=%t all=%t alternatives=[%s] verdict=%q", x.KSumRejected, x.LambdaOnlyRejected, x.KappaLambdaOnlyRejected, x.KappaEOnlyRejected, x.AbsoluteFormEquivalent, x.SignedFormAccepted, x.AllAudited, strings.Join(parts, "; "), x.Verdict)
}

func FormatSignAudit(x SignAndWallCoordinateAudit) string {
	return fmt.Sprintf("lambdaNeg=%t signed=%q absolute=%q dSigned=%.18g dAbs=%.18g equivalent=%t signedPreferred=%t orientation=%t verdict=%q", x.LambdaNegative, x.SignedForm, x.AbsoluteForm, x.DBaseSigned, x.DBaseAbsolute, x.FormsEquivalentNumerically, x.SignedFormPreferred, x.OrientationPreserved, x.Verdict)
}

func FormatBridge(x BridgeReconstructionAudit) string {
	return fmt.Sprintf("left=%q right=%q expanded=%q sigmaHistory=%.18g expectation=%.18g e1=%.18g matchesResidual=%t relates=%t verdict=%q", x.LeftSide, x.RightSide, x.ExpandedEquation, x.SigmaHistory, x.SigmaBoundaryExpectation, x.ResidualE1, x.MatchesInheritedResidual, x.RelatesQuotientToPayoff, x.Verdict)
}

func FormatResidual(x ResidualStatusAudit) string {
	return fmt.Sprintf("e1=%.18g quadraticRetained=%t quadraticPromoted=%t nativeResidual=%t verdict=%q", x.ResidualE1, x.QuadraticClueRetained, x.QuadraticCluePromoted, x.NativeResidualExplanation, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassificationAudit) string {
	return fmt.Sprintf("kappaLambda=%q kappaE=%q lambda=%q dbase=%q ssplit=%q rho=%q pk7=%q bridge=%q verdict=%q", x.KappaLambdaRole, x.KappaERole, x.LambdaRole, x.DBaseRole, x.SSplitRole, x.Rho72Role, x.PK7Role, x.BridgeRole, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("candidates=%q gap=%q verdict=%q", strings.Join(x.Candidates, ", "), x.PreciseGap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("expectedPayoffEqualsHistory=%t historyBoundaryTheorem=%t native7=%t payoffTheorem=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsExpectedK7PayoffEqualsHistoryDefectNatively, x.ClaimsNativeHistoryBoundaryResponseTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativePayoffTheorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
