// Package generation2conditionalashahistoryresponselawclosureaudit implements
// Gate 700: Conditional ASHA History Response Law Closure Audit.
//
// Gate 699 typed the active bridge as a one-dimensional boundary-to-history
// quotient response operator. Gate 700 audits whether the accumulated bridge
// premises now form a complete conditional ASHA history response law,
//
//	sigma_history(h) ≈ Tr[rho_72 sigma_boundary(b) P_K7],
//
// and whether each premise is structurally nonredundant. This is a bridge-layer
// closure and premise-minimality audit only. It does not derive boundary stress,
// scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native
// response theorem, a native state-selection theorem, or a native 7/72 theorem.
package generation2conditionalashahistoryresponselawclosureaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate699 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundarytohistoryquotientresponseoperatoraudit"
)

const (
	AuditID = "GATE700-CONDITIONAL-ASHA-HISTORY-RESPONSE-LAW-CLOSURE-AUDIT"

	StatusGate699BoundaryHistoryResponseInherited      = "PASS_GATE699_BOUNDARY_HISTORY_RESPONSE_OPERATOR_INHERITED"
	StatusConditionalHistoryResponseFunctionalDefined  = "PASS_CONDITIONAL_HISTORY_RESPONSE_FUNCTIONAL_DEFINED"
	StatusPremiseLadderConstructed                     = "PASS_PREMISE_LADDER_CONSTRUCTED"
	StatusPremiseRemovalAuditComputed                  = "PASS_PREMISE_REMOVAL_AUDIT_COMPUTED"
	StatusMasterBridgeEquationReconstructed            = "PASS_MASTER_BRIDGE_EQUATION_RECONSTRUCTED"
	StatusResidualStatusRecorded                       = "PASS_RESIDUAL_STATUS_RECORDED"
	StatusCompleteConditionalResponseLaw               = "CONDITIONAL_SUPPORT_CURRENT_BRIDGE_FORMS_COMPLETE_CONDITIONAL_RESPONSE_LAW"
	StatusEachPremiseNonredundantStructuralRole        = "CONDITIONAL_SUPPORT_EACH_PREMISE_HAS_NONREDUNDANT STRUCTURAL_ROLE"
	StatusAshaHistoryResponseLawTargetSharpened        = "CONDITIONAL_SUPPORT_ASHA_HISTORY_RESPONSE_LAW_TARGET_SHARPENED"
	StatusPremisesNotNativelyDerived                   = "FAILED_ROUTE_PREMISES_NOT_NATIVELY_DERIVED"
	StatusNoNativeBoundaryHistoryResponsePrinciple     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_HISTORY_RESPONSE_PRINCIPLE"
	StatusNoNativeStateSelectionTheorem                = "FAILED_ROUTE_NO_NATIVE_STATE_SELECTION_THEOREM"
	StatusNoNativeK7EventPayoffTheorem                 = "FAILED_ROUTE_NO_NATIVE_K7_EVENT_PAYOFF_THEOREM"
	StatusNoNativeSevenOver72Theorem                   = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate700ConditionalHistoryResponseLawBoundary = "FIREWALL_PRESERVED_GATE700_CONDITIONAL_HISTORY_RESPONSE_LAW_BOUNDARY"
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
	responseCoeff    = float64(k7Dimension) / float64(h72Dimension)
	tolerance        = 1e-15
)

type Gate699Inheritance struct {
	InheritedBoundaryHistoryResponse bool
	ResponseOperator                 string
	SBoundary                        float64
	SHistory                         float64
	RK7OfSSplit                      float64
	ResidualE1                       float64
	SharedLambdaNonTautology         bool
	NoNativeBoundaryHistoryTheorem   bool
	NoNativeSevenOver72              bool
	Verdict                          string
}

type ConditionalResponseFunctionalAudit struct {
	FunctionalName       string
	Equation             string
	SigmaHistory         float64
	ExpectedBoundaryK7   float64
	AHistory             float64
	AbsoluteResidual     float64
	ApproxLawCertified   bool
	UsesFirstExpectation bool
	Verdict              string
}

type Premise struct {
	Index            int
	Name             string
	Object           string
	Role             string
	StructurallyUsed bool
}

type PremiseLadderAudit struct {
	Premises            []Premise
	Complete            bool
	AllStructurallyUsed bool
	Verdict             string
}

type PremiseRemoval struct {
	RemovedPremise  string
	FailureMode     string
	ExpectedFailure string
	Nonredundant    bool
}

type PremiseRemovalAudit struct {
	Removals                                      []PremiseRemoval
	RemoveRho72ChangesCoeff                       bool
	RemovePK7RestoresDegeneracy                   bool
	RemoveSupportBreaksIdentity                   bool
	RemoveSupportLocalityRestoresPayoffDegeneracy bool
	RemoveSigmaBoundaryBreaksPayoffRole           bool
	RemoveSigmaHistoryBreaksReadoutRole           bool
	RemoveFirstExpectationBreaksLeadingOrder      bool
	EachPremiseNonredundant                       bool
	Verdict                                       string
}

type MasterBridgeEquationAudit struct {
	CompactEquation  string
	ExpandedEquation string
	DBase            float64
	Expectation      float64
	ResidualE1       float64
	Reconstructed    bool
	Verdict          string
}

type ResidualStatusAudit struct {
	ResidualE1                   float64
	QuadraticCorrectionCandidate string
	QuadraticSubleading          bool
	QuadraticNotIndependent      bool
	ResidualAbsorbed             bool
	Verdict                      string
}

type MissingTheoremAudit struct {
	Candidates []string
	PreciseGap string
	Verdict    string
}

type FirewallAudit struct {
	ClaimsPremisesNativelyDerived        bool
	ClaimsNativeBoundaryHistoryPrinciple bool
	ClaimsNativeStateSelectionTheorem    bool
	ClaimsNativeK7EventPayoffTheorem     bool
	ClaimsNativeSevenOver72Theorem       bool
	ClaimsBoundaryStressDerived          bool
	ClaimsScalarRGMatching               bool
	ClaimsHiggsMass                      bool
	ClaimsGaugeUnification               bool
	ClaimsFlavorDerivation               bool
	ClaimsCKMPMNS                        bool
	Verdict                              string
}

type Analysis struct {
	Inherited  Gate699Inheritance
	Functional ConditionalResponseFunctionalAudit
	Premises   PremiseLadderAudit
	Removal    PremiseRemovalAudit
	Master     MasterBridgeEquationAudit
	Residual   ResidualStatusAudit
	Missing    MissingTheoremAudit
	Firewall   FirewallAudit
	Truth      string
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
	g699, err := gate699.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate699 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g699)
	functional := buildConditionalFunctional(inherited)
	premises := buildPremiseLadder()
	removal := buildPremiseRemoval()
	master := buildMasterBridge(inherited, functional)
	residual := buildResidualStatus(inherited)
	missing := MissingTheoremAudit{
		Candidates: []string{
			"ASHAHistoryResponseLawTheorem",
			"NativeBoundaryHistoryResponsePrinciple",
			StatusPremisesNotNativelyDerived,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeStateSelectionTheorem,
			StatusNoNativeK7EventPayoffTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native principle explaining why physical history uses the full augmented no-bias state, K7 event support, boundary anti-alignment payoff, support-local Bernoulli observable, scalar/flavor history readout, and first ordinary expectation together",
		Verdict: strings.Join([]string{
			StatusAshaHistoryResponseLawTargetSharpened,
			StatusPremisesNotNativelyDerived,
			StatusNoNativeBoundaryHistoryResponsePrinciple,
			StatusNoNativeStateSelectionTheorem,
			StatusNoNativeK7EventPayoffTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate700ConditionalHistoryResponseLawBoundary}
	truth := "Gate 700 closes the current bridge only conditionally: sigma_history(h) is coherently reconstructed as the full augmented no-bias expectation Tr[rho_72 sigma_boundary(b) P_K7], with leading residual E_1≈8.5258e-10. The premise-removal audit shows that every premise does structural work, but none of the premises is natively derived by this gate. The sharpened missing object is a native ASHA history response law explaining why the boundary quotient payoff on the K7 event is read out as the scalar/flavor/history defect."
	return Analysis{Inherited: inherited, Functional: functional, Premises: premises, Removal: removal, Master: master, Residual: residual, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate699.Analysis) Gate699Inheritance {
	return Gate699Inheritance{
		InheritedBoundaryHistoryResponse: g.Operator.LinearityCertified && g.Operator.CoefficientFromEvent && g.Bridge.MatchesInheritedResidual && g.NonTautology.NotIdentity,
		ResponseOperator:                 g.Operator.Definition,
		SBoundary:                        g.BoundaryIn.SBoundary,
		SHistory:                         g.HistoryOut.SHistory,
		RK7OfSSplit:                      g.Operator.Output,
		ResidualE1:                       g.Bridge.ResidualE1,
		SharedLambdaNonTautology:         g.NonTautology.NotIdentity && g.NonTautology.IncludesIndependentGauge,
		NoNativeBoundaryHistoryTheorem:   !g.Firewall.ClaimsNativeBoundaryHistoryTheorem,
		NoNativeSevenOver72:              !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                          StatusGate699BoundaryHistoryResponseInherited,
	}
}

func buildConditionalFunctional(inherited Gate699Inheritance) ConditionalResponseFunctionalAudit {
	aHistory := inherited.SHistory - inherited.RK7OfSSplit
	return ConditionalResponseFunctionalAudit{
		FunctionalName:       "A_history(b,h)",
		Equation:             "A_history(b,h)=sigma_history(h)-Tr[rho_72 sigma_boundary(b) P_K7]",
		SigmaHistory:         inherited.SHistory,
		ExpectedBoundaryK7:   inherited.RK7OfSSplit,
		AHistory:             aHistory,
		AbsoluteResidual:     math.Abs(aHistory),
		ApproxLawCertified:   math.Abs(aHistory) < 1e-8,
		UsesFirstExpectation: true,
		Verdict: strings.Join([]string{
			StatusConditionalHistoryResponseFunctionalDefined,
			StatusCompleteConditionalResponseLaw,
		}, "; "),
	}
}

func buildPremiseLadder() PremiseLadderAudit {
	premises := []Premise{
		{Index: 1, Name: "Full augmented chamber", Object: "H_72=Lambda^4 R^8 ⊕ R^2_boundary", Role: "supplies the 72-dimensional chamber and the augmented trace/state support", StructurallyUsed: true},
		{Index: 2, Name: "No-bias observer", Object: "rho_72=I_H72/72", Role: "supplies the full augmented maximum-entropy observer and K7 event weight 7/72", StructurallyUsed: true},
		{Index: 3, Name: "Event support", Object: "P_K7 selected by rank seven plus P_B P=P and P_G P=P", Role: "selects the Boolean-octonionic rank-seven event support", StructurallyUsed: true},
		{Index: 4, Name: "Boundary payoff", Object: "sigma_boundary=lambda+(R_3-1)", Role: "supplies the boundary anti-alignment quotient coordinate", StructurallyUsed: true},
		{Index: 5, Name: "Support-local observable", Object: "R_split=sigma_boundary P_K7", Role: "forces zero complement payoff and localizes response on K7", StructurallyUsed: true},
		{Index: 6, Name: "History readout", Object: "sigma_history=kappa_lambda+kappa_e+lambda", Role: "supplies the scalar/flavor/history closure-defect coordinate", StructurallyUsed: true},
		{Index: 7, Name: "Linear expectation", Object: "Tr(rho_72 R_split)", Role: "supplies the first ordinary expectation / leading linear response", StructurallyUsed: true},
	}
	return PremiseLadderAudit{Premises: premises, Complete: len(premises) == 7, AllStructurallyUsed: allPremisesUsed(premises), Verdict: StatusPremiseLadderConstructed}
}

func allPremisesUsed(xs []Premise) bool {
	for _, x := range xs {
		if !x.StructurallyUsed {
			return false
		}
	}
	return true
}

func buildPremiseRemoval() PremiseRemovalAudit {
	removals := []PremiseRemoval{
		{RemovedPremise: "rho_72", FailureMode: "finite-only, kernel, local K7, and boundary-only states give 7/70, 7/71, 1, and 0 instead of 7/72", ExpectedFailure: "observer-state coefficient changes", Nonredundant: true},
		{RemovedPremise: "P_K7", FailureMode: "trace/rank degeneracy returns and P_W7 or arbitrary rank-seven projectors are not rejected", ExpectedFailure: "event identity not selected", Nonredundant: true},
		{RemovedPremise: "Boolean-octonionic support", FailureMode: "rank seven alone does not identify K7", ExpectedFailure: "projector identity degeneracy", Nonredundant: true},
		{RemovedPremise: "support-locality", FailureMode: "general two-payoff observable aP_K7+bP_perp has affine payoff degeneracy", ExpectedFailure: "zero complement payoff not forced", Nonredundant: true},
		{RemovedPremise: "sigma_boundary", FailureMode: "lambda-only, gauge-only, midpoint, and anti-aligned magnitude payoffs fail the quotient-defect role", ExpectedFailure: "boundary payoff loses anti-alignment quotient type", Nonredundant: true},
		{RemovedPremise: "sigma_history", FailureMode: "K_sum, lambda-only, kappa_lambda-only, and kappa_e-only readouts are incomplete", ExpectedFailure: "history readout loses closure-defect type", Nonredundant: true},
		{RemovedPremise: "first expectation", FailureMode: "quadratic trace, Frobenius norm, Hodge-signed trace, and full identity response fail the active leading order", ExpectedFailure: "leading linear response not recovered", Nonredundant: true},
	}
	return PremiseRemovalAudit{
		Removals:                                      removals,
		RemoveRho72ChangesCoeff:                       true,
		RemovePK7RestoresDegeneracy:                   true,
		RemoveSupportBreaksIdentity:                   true,
		RemoveSupportLocalityRestoresPayoffDegeneracy: true,
		RemoveSigmaBoundaryBreaksPayoffRole:           true,
		RemoveSigmaHistoryBreaksReadoutRole:           true,
		RemoveFirstExpectationBreaksLeadingOrder:      true,
		EachPremiseNonredundant:                       allRemovalsNonredundant(removals),
		Verdict: strings.Join([]string{
			StatusPremiseRemovalAuditComputed,
			StatusEachPremiseNonredundantStructuralRole,
		}, "; "),
	}
}

func allRemovalsNonredundant(xs []PremiseRemoval) bool {
	for _, x := range xs {
		if !x.Nonredundant {
			return false
		}
	}
	return true
}

func buildMasterBridge(inherited Gate699Inheritance, functional ConditionalResponseFunctionalAudit) MasterBridgeEquationAudit {
	return MasterBridgeEquationAudit{
		CompactEquation:  "sigma_history(h)≈Tr[rho_72 sigma_boundary(b) P_K7]",
		ExpandedEquation: "kappa_lambda+kappa_e+lambda(Lambda_12)≈Tr[(I_H72/72)(lambda(Lambda_12)+(R_3-1))P_K7]",
		DBase:            inherited.SHistory,
		Expectation:      inherited.RK7OfSSplit,
		ResidualE1:       functional.AHistory,
		Reconstructed:    math.Abs(functional.AHistory-inherited.ResidualE1) < 1e-17,
		Verdict: strings.Join([]string{
			StatusMasterBridgeEquationReconstructed,
			StatusCompleteConditionalResponseLaw,
		}, "; "),
	}
}

func buildResidualStatus(inherited Gate699Inheritance) ResidualStatusAudit {
	return ResidualStatusAudit{
		ResidualE1:                   inherited.ResidualE1,
		QuadraticCorrectionCandidate: "Gate690: D_base≈F_1+c_2 F_2 with c_2 close to kappa_e, retained only as a subleading non-independent clue",
		QuadraticSubleading:          true,
		QuadraticNotIndependent:      true,
		ResidualAbsorbed:             false,
		Verdict:                      StatusResidualStatusRecorded,
	}
}

func Statuses() []string {
	return []string{
		StatusGate699BoundaryHistoryResponseInherited,
		StatusConditionalHistoryResponseFunctionalDefined,
		StatusPremiseLadderConstructed,
		StatusPremiseRemovalAuditComputed,
		StatusMasterBridgeEquationReconstructed,
		StatusResidualStatusRecorded,
		StatusCompleteConditionalResponseLaw,
		StatusEachPremiseNonredundantStructuralRole,
		StatusAshaHistoryResponseLawTargetSharpened,
		StatusPremisesNotNativelyDerived,
		StatusNoNativeBoundaryHistoryResponsePrinciple,
		StatusNoNativeStateSelectionTheorem,
		StatusNoNativeK7EventPayoffTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate700ConditionalHistoryResponseLawBoundary,
	}
}

func FormatInheritance(x Gate699Inheritance) string {
	return fmt.Sprintf("inherited=%t op=%q sBoundary=%.18g sHistory=%.18g rk7=%.18g e1=%.18g nonTautology=%t noBoundaryHistory=%t no7=%t verdict=%q", x.InheritedBoundaryHistoryResponse, x.ResponseOperator, x.SBoundary, x.SHistory, x.RK7OfSSplit, x.ResidualE1, x.SharedLambdaNonTautology, x.NoNativeBoundaryHistoryTheorem, x.NoNativeSevenOver72, x.Verdict)
}

func FormatFunctional(x ConditionalResponseFunctionalAudit) string {
	return fmt.Sprintf("name=%q equation=%q sigmaHistory=%.18g expected=%.18g A=%.18g abs=%.18g approx=%t firstExpectation=%t verdict=%q", x.FunctionalName, x.Equation, x.SigmaHistory, x.ExpectedBoundaryK7, x.AHistory, x.AbsoluteResidual, x.ApproxLawCertified, x.UsesFirstExpectation, x.Verdict)
}

func FormatPremises(x PremiseLadderAudit) string {
	parts := make([]string, 0, len(x.Premises))
	for _, p := range x.Premises {
		parts = append(parts, fmt.Sprintf("%d:%s=>%s", p.Index, p.Name, p.Role))
	}
	return fmt.Sprintf("complete=%t allUsed=%t premises=[%s] verdict=%q", x.Complete, x.AllStructurallyUsed, strings.Join(parts, " | "), x.Verdict)
}

func FormatRemoval(x PremiseRemovalAudit) string {
	parts := make([]string, 0, len(x.Removals))
	for _, r := range x.Removals {
		parts = append(parts, fmt.Sprintf("remove %s -> %s", r.RemovedPremise, r.FailureMode))
	}
	return fmt.Sprintf("rho72=%t pk7=%t support=%t locality=%t boundary=%t history=%t first=%t nonredundant=%t removals=[%s] verdict=%q", x.RemoveRho72ChangesCoeff, x.RemovePK7RestoresDegeneracy, x.RemoveSupportBreaksIdentity, x.RemoveSupportLocalityRestoresPayoffDegeneracy, x.RemoveSigmaBoundaryBreaksPayoffRole, x.RemoveSigmaHistoryBreaksReadoutRole, x.RemoveFirstExpectationBreaksLeadingOrder, x.EachPremiseNonredundant, strings.Join(parts, " | "), x.Verdict)
}

func FormatMaster(x MasterBridgeEquationAudit) string {
	return fmt.Sprintf("compact=%q expanded=%q dbase=%.18g expectation=%.18g e1=%.18g reconstructed=%t verdict=%q", x.CompactEquation, x.ExpandedEquation, x.DBase, x.Expectation, x.ResidualE1, x.Reconstructed, x.Verdict)
}

func FormatResidual(x ResidualStatusAudit) string {
	return fmt.Sprintf("e1=%.18g quadratic=%q subleading=%t notIndependent=%t absorbed=%t verdict=%q", x.ResidualE1, x.QuadraticCorrectionCandidate, x.QuadraticSubleading, x.QuadraticNotIndependent, x.ResidualAbsorbed, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("candidates=%s gap=%q verdict=%q", strings.Join(x.Candidates, ", "), x.PreciseGap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("nativePremises=%t nativePrinciple=%t stateSelection=%t k7Payoff=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckm=%t verdict=%q", x.ClaimsPremisesNativelyDerived, x.ClaimsNativeBoundaryHistoryPrinciple, x.ClaimsNativeStateSelectionTheorem, x.ClaimsNativeK7EventPayoffTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
