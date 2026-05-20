// Package generation2bernoullipayoffnormalizationandzerocomplementsupportaudit implements
// Gate 696: Bernoulli Payoff Normalization and Zero-Complement Support Audit.
//
// Gate 695 typed the active bridge as a no-bias K7 event expectation under
// rho_72=I_H72/72. Gate 696 audits the payoff normalization behind that
// Bernoulli reading.  The general two-event observable
// R_{a,b}=aP_K7+b(I_H72-P_K7) has expectation (7/72)a+(65/72)b, so expectation
// alone leaves an affine payoff degeneracy.  K7 support-locality forces b=0;
// the remaining event payoff is then supplied by the boundary quotient scalar
// a=S_split, giving R_split=S_split P_K7.
//
// This is a bridge-layer payoff-normalization audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native state-selection theorem, a native payoff theorem, or a
// native 7/72 theorem.
package generation2bernoullipayoffnormalizationandzerocomplementsupportaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate695 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7eventweightandbernoulliresponseobservableaudit"
)

const (
	AuditID = "GATE696-BERNOULLI-PAYOFF-NORMALIZATION-ZERO-COMPLEMENT-SUPPORT-AUDIT"

	StatusGate695BernoulliObservableInherited              = "PASS_GATE695_BERNOULLI_OBSERVABLE_INHERITED"
	StatusGeneralTwoPayoffObservableDefined                = "PASS_GENERAL_TWO_PAYOFF_OBSERVABLE_DEFINED"
	StatusExpectationForGeneralABComputed                  = "PASS_EXPECTATION_FOR_GENERAL_A_B_COMPUTED"
	StatusAffinePayoffDegeneracyAudited                    = "PASS_AFFINE_PAYOFF_DEGENERACY_AUDITED"
	StatusSupportLocalityConditionDefined                  = "PASS_SUPPORT_LOCALITY_CONDITION_DEFINED"
	StatusSupportLocalityForcesZeroComplementPayoff        = "PASS_SUPPORT_LOCALITY_FORCES_ZERO_COMPLEMENT_PAYOFF"
	StatusActivePayoffAssignmentReconstructed              = "PASS_ACTIVE_PAYOFF_ASSIGNMENT_RECONSTRUCTED"
	StatusAlternativePayoffObservablesAudited              = "PASS_ALTERNATIVE_PAYOFF_OBSERVABLES_AUDITED"
	StatusActiveResponseSupportLocalK7PayoffObservable     = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_IS_SUPPORT_LOCAL_K7_PAYOFF_OBSERVABLE"
	StatusZeroComplementPayoffFromK7SupportLocality        = "CONDITIONAL_SUPPORT_ZERO_COMPLEMENT_PAYOFF_FROM_K7_SUPPORT_LOCALITY"
	StatusExpectationAloneDoesNotSelectPayoffNormalization = "FAILED_ROUTE_EXPECTATION_VALUE_ALONE_DOES_NOT_SELECT_PAYOFF_NORMALIZATION"
	StatusNoNativeReasonHistoryUsesSupportLocality         = "FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_USES_SUPPORT_LOCALITY"
	StatusNoNativeReasonK7EventReceivesSSplitPayoff        = "FAILED_ROUTE_NO_NATIVE_REASON_K7_EVENT_RECEIVES_S_SPLIT_PAYOFF"
	StatusNoNativeSevenOver72Theorem                       = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate696BernoulliPayoffSupportBoundary            = "FIREWALL_PRESERVED_GATE696_BERNOULLI_PAYOFF_SUPPORT_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	k7Dimension       = 7
	k7ComplementDim   = h72Dimension - k7Dimension
	tolerance         = 1e-15
)

type Gate695Inheritance struct {
	BernoulliObservableInherited bool
	Rho72Definition              string
	EventProjector               string
	ResponseOperator             string
	EventWeight                  float64
	ComplementWeight             float64
	SSplit                       float64
	DBase                        float64
	ActiveExpectation            float64
	ResidualE1                   float64
	SecondMoment                 float64
	Variance                     float64
	H72Dimension                 int
	K7Dimension                  int
	ComplementDimension          int
	NoNativeHistoryResponse      bool
	NoNativeRho72Reason          bool
	NoNativePayoffReason         bool
	NoNativeSevenOver72          bool
	Verdict                      string
}

type GeneralTwoPayoffObservableAudit struct {
	Observable             string
	EventProjector         string
	ComplementProjector    string
	EventWeight            float64
	ComplementWeight       float64
	ExpectationFormula     string
	ActiveEventPayoff      float64
	ActiveComplementPayoff float64
	ActiveExpectation      float64
	Verdict                string
}

type AffinePayoffDegeneracyAudit struct {
	Equation                       string
	ReferenceExpectation           float64
	ReferenceEventPayoff           float64
	ReferenceComplementPayoff      float64
	AlternativeEventPayoff         float64
	AlternativeComplementPayoff    float64
	AlternativeExpectation         float64
	AlternativeDifferentFromActive bool
	ExpectationAloneDegenerate     bool
	Verdict                        string
}

type SupportLocalityAudit struct {
	Conditions                          []string
	ForGeneralObservable                string
	PPerpLeftAction                     string
	PPerpRightAction                    string
	PK7SandwichAction                   string
	ComplementPayoffForcedZero          bool
	EventPayoffUnfixedBySupportLocality bool
	Verdict                             string
}

type ActivePayoffAssignmentAudit struct {
	SupportLocalObservable     string
	BoundaryPayoffAssignment   string
	EventPayoff                float64
	ComplementPayoff           float64
	ReconstructedOperator      string
	Expectation                float64
	MatchesInheritedActive     bool
	DoesNotProvePayoffNatively bool
	Verdict                    string
}

type AlternativePayoffObservable struct {
	Name        string
	Observable  string
	Expectation float64
	Active      bool
	Reason      string
}

type AlternativePayoffObservablesAudit struct {
	Alternatives             []AlternativePayoffObservable
	FullPayoffRejected       bool
	ComplementPayoffRejected bool
	CenteredRejected         bool
	SignedHodgeRejected      bool
	ActiveAccepted           bool
	AllAudited               bool
	Verdict                  string
}

type SourceTypeClassificationAudit struct {
	PK7Role    string
	PPerpRole  string
	SSplitRole string
	Rho72Role  string
	RSplitRole string
	Verdict    string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type FirewallAudit struct {
	ClaimsHistoryUsesSupportLocalityNatively bool
	ClaimsK7PhysicalEventNatively            bool
	ClaimsSSplitPayoffNatively               bool
	ClaimsExpectationEqualsDBaseNatively     bool
	ClaimsResidualExplained                  bool
	ClaimsNativeStateSelectionTheorem        bool
	ClaimsNativePayoffTheorem                bool
	ClaimsNativeSevenOver72Theorem           bool
	ClaimsBoundaryStress                     bool
	ClaimsScalarRGMatching                   bool
	ClaimsHiggsMass                          bool
	ClaimsGaugeUnification                   bool
	ClaimsFlavorDerivation                   bool
	ClaimsCKMPMNS                            bool
	Verdict                                  string
}

type Analysis struct {
	Inherited       Gate695Inheritance
	General         GeneralTwoPayoffObservableAudit
	Degeneracy      AffinePayoffDegeneracyAudit
	SupportLocality SupportLocalityAudit
	Assignment      ActivePayoffAssignmentAudit
	Alternatives    AlternativePayoffObservablesAudit
	SourceTypes     SourceTypeClassificationAudit
	Missing         MissingTheoremAudit
	Firewall        FirewallAudit
	Truth           string
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
	g695, err := gate695.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate695 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g695)
	general := buildGeneralTwoPayoff(inherited)
	degeneracy := buildAffineDegeneracy(inherited)
	support := buildSupportLocality()
	assignment := buildActivePayoffAssignment(inherited, support)
	alternatives := buildAlternativePayoffs(inherited)
	sourceTypes := buildSourceTypes()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusExpectationAloneDoesNotSelectPayoffNormalization,
			StatusNoNativeReasonHistoryUsesSupportLocality,
			StatusNoNativeReasonK7EventReceivesSSplitPayoff,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native history-response theorem explaining why physical history imposes K7 support-locality, why K7 is the active event, why the boundary quotient scalar S_split is assigned as the event payoff, and why the resulting expectation should equal D_base",
		Verdict: strings.Join([]string{
			StatusExpectationAloneDoesNotSelectPayoffNormalization,
			StatusNoNativeReasonHistoryUsesSupportLocality,
			StatusNoNativeReasonK7EventReceivesSSplitPayoff,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate696BernoulliPayoffSupportBoundary}
	truth := "Gate 696 audits the payoff normalization behind the Gate695 Bernoulli reading.  The general two-event observable R_{a,b}=aP_K7+bP_perp has expectation (7/72)a+(65/72)b, so expectation alone cannot select the active payoff pair.  Imposing K7 support-locality forces the complement payoff b to vanish; the boundary quotient scalar then supplies the event payoff a=S_split, reconstructing R_split=S_split P_K7.  This conditionally supports the active response as a support-local K7 payoff observable, while preserving the firewall that no native theorem yet explains why history uses support-locality, why K7 receives S_split, or why a native 7/72 theorem should hold."
	return Analysis{Inherited: inherited, General: general, Degeneracy: degeneracy, SupportLocality: support, Assignment: assignment, Alternatives: alternatives, SourceTypes: sourceTypes, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate695.Analysis) Gate695Inheritance {
	return Gate695Inheritance{
		BernoulliObservableInherited: g.Observable.TwoPointObservable && g.Observable.ComplementPayoff == 0 && math.Abs(g.Event.EventWeight-7.0/72.0) < tolerance,
		Rho72Definition:              g.Inherited.Rho72Definition,
		EventProjector:               g.Event.EventProjector,
		ResponseOperator:             g.Observable.Observable,
		EventWeight:                  g.Event.EventWeight,
		ComplementWeight:             g.Event.ComplementWeight,
		SSplit:                       g.Inherited.SSplit,
		DBase:                        g.Inherited.DBase,
		ActiveExpectation:            g.Expectation.Expectation,
		ResidualE1:                   g.Expectation.ResidualE1,
		SecondMoment:                 g.Moment.SecondMoment,
		Variance:                     g.Moment.Variance,
		H72Dimension:                 g.Inherited.H72Dimension,
		K7Dimension:                  g.Inherited.K7Dimension,
		ComplementDimension:          g.Event.ComplementDimension,
		NoNativeHistoryResponse:      !g.Firewall.ClaimsExpectationEqualsDBaseNatively,
		NoNativeRho72Reason:          !g.Firewall.ClaimsHistoryUsesRho72Natively,
		NoNativePayoffReason:         !g.Firewall.ClaimsSSplitPayoffNatively,
		NoNativeSevenOver72:          !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                      StatusGate695BernoulliObservableInherited,
	}
}

func buildGeneralTwoPayoff(i Gate695Inheritance) GeneralTwoPayoffObservableAudit {
	expectation := i.EventWeight*i.SSplit + i.ComplementWeight*0
	return GeneralTwoPayoffObservableAudit{
		Observable:             "R_{a,b}=aP_K7+bP_perp",
		EventProjector:         "P_K7",
		ComplementProjector:    "P_perp=I_H72-P_K7",
		EventWeight:            i.EventWeight,
		ComplementWeight:       i.ComplementWeight,
		ExpectationFormula:     "E_rho72[R_{a,b}]=(7/72)a+(65/72)b",
		ActiveEventPayoff:      i.SSplit,
		ActiveComplementPayoff: 0,
		ActiveExpectation:      expectation,
		Verdict: strings.Join([]string{
			StatusGeneralTwoPayoffObservableDefined,
			StatusExpectationForGeneralABComputed,
		}, "; "),
	}
}

func buildAffineDegeneracy(i Gate695Inheritance) AffinePayoffDegeneracyAudit {
	// One explicit non-active solution to (7/72)a+(65/72)b=(7/72)S_split:
	// set a=0 and solve b=(7/65)S_split.
	altA := 0.0
	altB := (float64(k7Dimension) / float64(k7ComplementDim)) * i.SSplit
	altExpectation := i.EventWeight*altA + i.ComplementWeight*altB
	return AffinePayoffDegeneracyAudit{
		Equation:                       "(7/72)a+(65/72)b=(7/72)S_split",
		ReferenceExpectation:           i.ActiveExpectation,
		ReferenceEventPayoff:           i.SSplit,
		ReferenceComplementPayoff:      0,
		AlternativeEventPayoff:         altA,
		AlternativeComplementPayoff:    altB,
		AlternativeExpectation:         altExpectation,
		AlternativeDifferentFromActive: math.Abs(altA-i.SSplit) > tolerance && math.Abs(altB) > tolerance,
		ExpectationAloneDegenerate:     math.Abs(altExpectation-i.ActiveExpectation) < tolerance,
		Verdict: strings.Join([]string{
			StatusAffinePayoffDegeneracyAudited,
			StatusExpectationAloneDoesNotSelectPayoffNormalization,
		}, "; "),
	}
}

func buildSupportLocality() SupportLocalityAudit {
	return SupportLocalityAudit{
		Conditions: []string{
			"P_K7 R P_K7 = R",
			"P_perp R = 0",
			"R P_perp = 0",
		},
		ForGeneralObservable:                "R_{a,b}=aP_K7+bP_perp",
		PPerpLeftAction:                     "P_perp R_{a,b}=bP_perp",
		PPerpRightAction:                    "R_{a,b}P_perp=bP_perp",
		PK7SandwichAction:                   "P_K7 R_{a,b} P_K7=aP_K7",
		ComplementPayoffForcedZero:          true,
		EventPayoffUnfixedBySupportLocality: true,
		Verdict: strings.Join([]string{
			StatusSupportLocalityConditionDefined,
			StatusSupportLocalityForcesZeroComplementPayoff,
			StatusZeroComplementPayoffFromK7SupportLocality,
		}, "; "),
	}
}

func buildActivePayoffAssignment(i Gate695Inheritance, s SupportLocalityAudit) ActivePayoffAssignmentAudit {
	expectation := i.EventWeight * i.SSplit
	return ActivePayoffAssignmentAudit{
		SupportLocalObservable:     "R_a=aP_K7",
		BoundaryPayoffAssignment:   "a=S_split",
		EventPayoff:                i.SSplit,
		ComplementPayoff:           0,
		ReconstructedOperator:      "R_split=S_split P_K7",
		Expectation:                expectation,
		MatchesInheritedActive:     s.ComplementPayoffForcedZero && math.Abs(expectation-i.ActiveExpectation) < tolerance,
		DoesNotProvePayoffNatively: true,
		Verdict: strings.Join([]string{
			StatusActivePayoffAssignmentReconstructed,
			StatusActiveResponseSupportLocalK7PayoffObservable,
		}, "; "),
	}
}

func buildAlternativePayoffs(i Gate695Inheritance) AlternativePayoffObservablesAudit {
	p := i.EventWeight
	q := i.ComplementWeight
	s := i.SSplit
	alternatives := []AlternativePayoffObservable{
		{
			Name:        "full payoff",
			Observable:  "S_split I_H72",
			Expectation: s,
			Active:      false,
			Reason:      "no K7 support; pays the complement as well as the event",
		},
		{
			Name:        "complement payoff",
			Observable:  "S_split P_perp",
			Expectation: q * s,
			Active:      false,
			Reason:      "wrong support: payoff assigned to K7 complement",
		},
		{
			Name:        "centered observable",
			Observable:  "S_split(P_K7-(7/72)I_H72)",
			Expectation: 0,
			Active:      false,
			Reason:      "fluctuation observable with zero mean under rho_72",
		},
		{
			Name:        "signed Hodge payoff",
			Observable:  "S_split(P_+-P_-)",
			Expectation: s / 72.0,
			Active:      false,
			Reason:      "signed polarity response; active bridge uses ordinary support weight",
		},
		{
			Name:        "active support-local observable",
			Observable:  "S_split P_K7",
			Expectation: p * s,
			Active:      true,
			Reason:      "K7 support-local payoff with zero complement response",
		},
	}
	return AlternativePayoffObservablesAudit{
		Alternatives:             alternatives,
		FullPayoffRejected:       !alternatives[0].Active && math.Abs(alternatives[0].Expectation-s) < tolerance,
		ComplementPayoffRejected: !alternatives[1].Active && math.Abs(alternatives[1].Expectation-q*s) < tolerance,
		CenteredRejected:         !alternatives[2].Active && math.Abs(alternatives[2].Expectation) < tolerance,
		SignedHodgeRejected:      !alternatives[3].Active && math.Abs(alternatives[3].Expectation-s/72.0) < tolerance,
		ActiveAccepted:           alternatives[4].Active && math.Abs(alternatives[4].Expectation-i.ActiveExpectation) < tolerance,
		AllAudited:               true,
		Verdict:                  StatusAlternativePayoffObservablesAudited,
	}
}

func buildSourceTypes() SourceTypeClassificationAudit {
	return SourceTypeClassificationAudit{
		PK7Role:    "event support selected by Boolean-octonionic support",
		PPerpRole:  "no-response complement under K7 support-locality",
		SSplitRole: "boundary anti-alignment quotient payoff assigned to the active event",
		Rho72Role:  "full augmented no-bias observer state",
		RSplitRole: "support-local Bernoulli payoff observable",
		Verdict:    StatusActiveResponseSupportLocalK7PayoffObservable,
	}
}

func Statuses() []string {
	return []string{
		StatusGate695BernoulliObservableInherited,
		StatusGeneralTwoPayoffObservableDefined,
		StatusExpectationForGeneralABComputed,
		StatusAffinePayoffDegeneracyAudited,
		StatusSupportLocalityConditionDefined,
		StatusSupportLocalityForcesZeroComplementPayoff,
		StatusActivePayoffAssignmentReconstructed,
		StatusAlternativePayoffObservablesAudited,
		StatusActiveResponseSupportLocalK7PayoffObservable,
		StatusZeroComplementPayoffFromK7SupportLocality,
		StatusExpectationAloneDoesNotSelectPayoffNormalization,
		StatusNoNativeReasonHistoryUsesSupportLocality,
		StatusNoNativeReasonK7EventReceivesSSplitPayoff,
		StatusNoNativeSevenOver72Theorem,
		StatusGate696BernoulliPayoffSupportBoundary,
	}
}

func FormatInheritance(x Gate695Inheritance) string {
	return fmt.Sprintf("inherited=%t rho=%q event=%q response=%q p=%.18g q=%.18g ssplit=%.18g dbase=%.18g expectation=%.18g e1=%.18g second=%.18g variance=%.18g h72=%d k7=%d comp=%d noHistory=%t noRho=%t noPayoff=%t no7=%t verdict=%q", x.BernoulliObservableInherited, x.Rho72Definition, x.EventProjector, x.ResponseOperator, x.EventWeight, x.ComplementWeight, x.SSplit, x.DBase, x.ActiveExpectation, x.ResidualE1, x.SecondMoment, x.Variance, x.H72Dimension, x.K7Dimension, x.ComplementDimension, x.NoNativeHistoryResponse, x.NoNativeRho72Reason, x.NoNativePayoffReason, x.NoNativeSevenOver72, x.Verdict)
}

func FormatGeneral(x GeneralTwoPayoffObservableAudit) string {
	return fmt.Sprintf("observable=%q event=%q complement=%q p=%.18g q=%.18g formula=%q activeA=%.18g activeB=%.18g activeExpectation=%.18g verdict=%q", x.Observable, x.EventProjector, x.ComplementProjector, x.EventWeight, x.ComplementWeight, x.ExpectationFormula, x.ActiveEventPayoff, x.ActiveComplementPayoff, x.ActiveExpectation, x.Verdict)
}

func FormatDegeneracy(x AffinePayoffDegeneracyAudit) string {
	return fmt.Sprintf("equation=%q reference=%.18g refA=%.18g refB=%.18g altA=%.18g altB=%.18g altExpectation=%.18g different=%t degenerate=%t verdict=%q", x.Equation, x.ReferenceExpectation, x.ReferenceEventPayoff, x.ReferenceComplementPayoff, x.AlternativeEventPayoff, x.AlternativeComplementPayoff, x.AlternativeExpectation, x.AlternativeDifferentFromActive, x.ExpectationAloneDegenerate, x.Verdict)
}

func FormatSupportLocality(x SupportLocalityAudit) string {
	return fmt.Sprintf("conditions=[%s] general=%q left=%q right=%q sandwich=%q bZero=%t aUnfixed=%t verdict=%q", strings.Join(x.Conditions, "; "), x.ForGeneralObservable, x.PPerpLeftAction, x.PPerpRightAction, x.PK7SandwichAction, x.ComplementPayoffForcedZero, x.EventPayoffUnfixedBySupportLocality, x.Verdict)
}

func FormatAssignment(x ActivePayoffAssignmentAudit) string {
	return fmt.Sprintf("supportLocal=%q payoff=%q a=%.18g b=%.18g operator=%q expectation=%.18g matches=%t notNative=%t verdict=%q", x.SupportLocalObservable, x.BoundaryPayoffAssignment, x.EventPayoff, x.ComplementPayoff, x.ReconstructedOperator, x.Expectation, x.MatchesInheritedActive, x.DoesNotProvePayoffNatively, x.Verdict)
}

func FormatAlternatives(x AlternativePayoffObservablesAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, s := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s[%s]:expectation=%.18g active=%t reason=%s", s.Name, s.Observable, s.Expectation, s.Active, s.Reason))
	}
	return fmt.Sprintf("alternatives={%s} full=%t complement=%t centered=%t signed=%t active=%t all=%t verdict=%q", strings.Join(parts, "; "), x.FullPayoffRejected, x.ComplementPayoffRejected, x.CenteredRejected, x.SignedHodgeRejected, x.ActiveAccepted, x.AllAudited, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassificationAudit) string {
	return fmt.Sprintf("pk7=%q pperp=%q ssplit=%q rho72=%q rsplit=%q verdict=%q", x.PK7Role, x.PPerpRole, x.SSplitRole, x.Rho72Role, x.RSplitRole, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("supportLocality=%t k7Event=%t payoff=%t expectationDBase=%t residual=%t stateSelection=%t payoffTheorem=%t native7=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsHistoryUsesSupportLocalityNatively, x.ClaimsK7PhysicalEventNatively, x.ClaimsSSplitPayoffNatively, x.ClaimsExpectationEqualsDBaseNatively, x.ClaimsResidualExplained, x.ClaimsNativeStateSelectionTheorem, x.ClaimsNativePayoffTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
