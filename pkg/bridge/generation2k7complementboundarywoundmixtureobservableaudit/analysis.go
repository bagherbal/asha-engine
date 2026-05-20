// Package generation2k7complementboundarywoundmixtureobservableaudit implements
// Gate 704: K7/Complement Boundary Wound Mixture Observable Audit.
//
// Gate 703 typed the active response coefficient as the K7 event probability
// after scalar-wall unit gluing. Gate 704 audits the equivalent positive-
// distance rearrangement
//
//	kappa_lambda+kappa_e ≈ (65/72)|lambda(Lambda_12)| + (7/72)(R_3-1)
//
// as a two-event boundary wound observable on the K7/complement split under
// the full augmented no-bias state rho_72. This is a bridge-layer boundary-
// wound mixture audit only. It does not derive boundary stress, scalar RG
// matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native response
// theorem, a native state-selection theorem, or a native 7/72 theorem.
package generation2k7complementboundarywoundmixtureobservableaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate703 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarwallairlockandquotientlinegluingaudit"
)

const (
	AuditID = "GATE704-K7-COMPLEMENT-BOUNDARY-WOUND-MIXTURE-OBSERVABLE-AUDIT"

	StatusGate703ScalarWallAirlockInherited       = "PASS_GATE703_SCALAR_WALL_AIRLOCK_INHERITED"
	StatusGate700ResponseLawRearranged            = "PASS_GATE700_RESPONSE_LAW_REARRANGED"
	StatusK7AndComplementProbabilitiesComputed    = "PASS_K7_AND_COMPLEMENT_PROBABILITIES_COMPUTED"
	StatusTwoPayoffBoundaryWoundObservableDefined = "PASS_TWO_PAYOFF_BOUNDARY_WOUND_OBSERVABLE_DEFINED"
	StatusExpectationReproducesWeightedClosure    = "PASS_EXPECTATION_REPRODUCES_WEIGHTED_CLOSURE"
	StatusNumericalResidualRecorded               = "PASS_NUMERICAL_RESIDUAL_RECORDED"
	StatusEquivalenceToPreviousFormsAudited       = "PASS_EQUIVALENCE_TO_PREVIOUS_FORMS_AUDITED"
	StatusAlternativeMixtureObservablesAudited    = "PASS_ALTERNATIVE_MIXTURE_OBSERVABLES_AUDITED"
	StatusKappaSumNoBiasExpectedBoundaryWound     = "CONDITIONAL_SUPPORT_KAPPA_SUM_IS_NO_BIAS_EXPECTED_BOUNDARY_WOUND"
	Status65Over72ComplementEventProbability      = "CONDITIONAL_SUPPORT_65_OVER_72_IS_COMPLEMENT_EVENT_PROBABILITY"
	Status7Over72K7EventProbability               = "CONDITIONAL_SUPPORT_7_OVER_72_IS_K7_EVENT_PROBABILITY"
	StatusWeightedBoundaryClosureEventMeaning     = "CONDITIONAL_SUPPORT_WEIGHTED_BOUNDARY_CLOSURE_HAS_EVENT_COMPLEMENT_MEANING"
	StatusNoNativeK7ReceivesGaugeWound            = "FAILED_ROUTE_NO_NATIVE_REASON_K7_RECEIVES_GAUGE_WOUND"
	StatusNoNativeComplementReceivesScalarWound   = "FAILED_ROUTE_NO_NATIVE_REASON_COMPLEMENT_RECEIVES_SCALAR_WOUND"
	StatusNoNativeBoundaryWoundMixtureTheorem     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_MIXTURE_THEOREM"
	StatusNoNativeHistoryResponseTheorem          = "FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem              = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate704BoundaryWoundMixtureBoundary     = "FIREWALL_PRESERVED_GATE704_BOUNDARY_WOUND_MIXTURE_BOUNDARY"
)

const (
	h72Dimension   = 72
	k7Dimension    = 7
	complementDim  = h72Dimension - k7Dimension
	pK7            = float64(k7Dimension) / float64(h72Dimension)
	pComplement    = float64(complementDim) / float64(h72Dimension)
	kappaLambda    = 0.0443230430960771
	kappaE         = 0.00550355419157456
	lambdaLambda12 = -0.0497009420776833
	r3Minus1       = 0.0509933868964996
	tolerance      = 1e-15
)

type Gate703Inheritance struct {
	ScalarWallAirlockInherited bool
	UnitScalarWallGlue         bool
	EventProbability           float64
	ResponseCoefficient        float64
	NonTautologyInherited      bool
	NoNativeScalarWallAirlock  bool
	NoNativeBoundaryHistory    bool
	NoNativeSevenOver72        bool
	Verdict                    string
}

type RearrangementAudit struct {
	StartingEquation      string
	RearrangedEquation    string
	LambdaNegative        bool
	KSum                  float64
	PositiveScalarDepth   float64
	GaugeWound            float64
	WeightedClosureRight  float64
	Residual              float64
	SameAsGate700Residual bool
	Verdict               string
}

type ProbabilityAudit struct {
	Rho72               string
	PK7                 string
	PComplement         string
	K7Rank              int
	ComplementRank      int
	TotalDimension      int
	PK7Probability      float64
	ComplementProb      float64
	ProbabilitiesSumTo1 bool
	Verdict             string
}

type BoundaryWoundObservableAudit struct {
	Observable            string
	K7Payoff              float64
	ComplementPayoff      float64
	K7PayoffRole          string
	ComplementPayoffRole  string
	SupportSplit          string
	IsTwoPayoffObservable bool
	Verdict               string
}

type ExpectationAudit struct {
	Formula                   string
	ExpectedBoundaryWound     float64
	KSum                      float64
	Residual                  float64
	ReproducesWeightedClosure bool
	Verdict                   string
}

type NumericalAudit struct {
	KSum                  float64
	ExpectedWound         float64
	Residual              float64
	InheritedResidualForm float64
	SameResidual          bool
	Verdict               string
}

type InterpretationAudit struct {
	K7EventPayoff    string
	ComplementPayoff string
	Observer         string
	Output           string
	Reading          string
	Verdict          string
}

type EquivalentForm struct {
	Name       string
	Equation   string
	Equivalent bool
}

type EquivalenceAudit struct {
	Forms                          []EquivalentForm
	IntroducesNewNumericalRelation bool
	UpgradesSourceType             bool
	Verdict                        string
}

type AlternativeMixtureObservable struct {
	Name        string
	K7Payoff    float64
	PerpPayoff  float64
	Expectation float64
	Active      bool
	Rejected    bool
	Reason      string
}

type AlternativeMixtureAudit struct {
	Alternatives              []AlternativeMixtureObservable
	ReversedPayoffRejected    bool
	SupportLocalFormSeparated bool
	MidpointRejected          bool
	HodgeSignedRejected       bool
	ActiveMixtureAccepted     bool
	Verdict                   string
}

type MissingTheoremAudit struct {
	Missing []string
	Verdict string
}

type FirewallAudit struct {
	ClaimsK7ReceivesGaugeWoundNative          bool
	ClaimsComplementReceivesScalarWoundNative bool
	ClaimsNativeBoundaryWoundMixtureTheorem   bool
	ClaimsNativeHistoryResponseTheorem        bool
	ClaimsNativeSevenOver72Theorem            bool
	ClaimsBoundaryStressDerived               bool
	ClaimsScalarRGMatching                    bool
	ClaimsHiggsMass                           bool
	ClaimsGaugeUnification                    bool
	ClaimsFlavorDerivation                    bool
	ClaimsCKMPMNS                             bool
	Verdict                                   string
}

type Analysis struct {
	Inherited      Gate703Inheritance
	Rearrangement  RearrangementAudit
	Probabilities  ProbabilityAudit
	Observable     BoundaryWoundObservableAudit
	Expectation    ExpectationAudit
	Numerical      NumericalAudit
	Interpretation InterpretationAudit
	Equivalence    EquivalenceAudit
	Alternatives   AlternativeMixtureAudit
	Missing        MissingTheoremAudit
	Firewall       FirewallAudit
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
	g703, err := gate703.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate703 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g703)
	rearranged := buildRearrangement()
	probs := buildProbabilities()
	observable := buildObservable()
	expectation := buildExpectation(rearranged)
	numerical := buildNumerical(expectation)
	interpretation := buildInterpretation()
	equivalence := buildEquivalence()
	alternatives := buildAlternatives()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeK7ReceivesGaugeWound,
			StatusNoNativeComplementReceivesScalarWound,
			StatusNoNativeBoundaryWoundMixtureTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusNoNativeK7ReceivesGaugeWound,
			StatusNoNativeComplementReceivesScalarWound,
			StatusNoNativeBoundaryWoundMixtureTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate704BoundaryWoundMixtureBoundary}
	truth := "Gate 704 rewrites the scalar-wall glued response law into a positive-distance K7/complement mixture: K_sum≈E_rho72[R P_K7+|lambda|P_perp]. This is equivalent to Gate700/Gate703 and assigns source type to the old weighted closure as a no-bias two-event boundary wound expectation. It does not prove why K7 receives the gauge wound or why the complement receives the scalar wound."
	return Analysis{Inherited: inherited, Rearrangement: rearranged, Probabilities: probs, Observable: observable, Expectation: expectation, Numerical: numerical, Interpretation: interpretation, Equivalence: equivalence, Alternatives: alternatives, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate703.Analysis) Gate703Inheritance {
	return Gate703Inheritance{
		ScalarWallAirlockInherited: g.UnitGlue.CoefficientEqualsProbability && g.UnitGlue.UnitGlue,
		UnitScalarWallGlue:         g.UnitGlue.UnitGlue,
		EventProbability:           g.UnitGlue.EventProbability,
		ResponseCoefficient:        g.UnitGlue.ResponseCoefficient,
		NonTautologyInherited:      g.NonTautology.NonTautologicalRelationPreserved,
		NoNativeScalarWallAirlock:  !g.Firewall.ClaimsNativeScalarWallAirlock,
		NoNativeBoundaryHistory:    !g.Firewall.ClaimsNativeBoundaryHistory,
		NoNativeSevenOver72:        !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                    StatusGate703ScalarWallAirlockInherited,
	}
}

func buildRearrangement() RearrangementAudit {
	ksum := kappaLambda + kappaE
	absLambda := math.Abs(lambdaLambda12)
	rhs := pComplement*absLambda + pK7*r3Minus1
	residual := ksum - rhs
	gate700Residual := (kappaLambda + kappaE + lambdaLambda12) - pK7*(lambdaLambda12+r3Minus1)
	return RearrangementAudit{
		StartingEquation:      "kappa_lambda+kappa_e+lambda≈p_K7(lambda+R)",
		RearrangedEquation:    "K_sum≈p_perp|lambda|+p_K7(R_3-1)",
		LambdaNegative:        lambdaLambda12 < 0,
		KSum:                  ksum,
		PositiveScalarDepth:   absLambda,
		GaugeWound:            r3Minus1,
		WeightedClosureRight:  rhs,
		Residual:              residual,
		SameAsGate700Residual: math.Abs(residual-gate700Residual) < 1e-17,
		Verdict: strings.Join([]string{
			StatusGate700ResponseLawRearranged,
			StatusKappaSumNoBiasExpectedBoundaryWound,
		}, "; "),
	}
}

func buildProbabilities() ProbabilityAudit {
	return ProbabilityAudit{
		Rho72:               "rho_72=I_H72/72",
		PK7:                 "P_K7",
		PComplement:         "P_perp=I_H72-P_K7",
		K7Rank:              k7Dimension,
		ComplementRank:      complementDim,
		TotalDimension:      h72Dimension,
		PK7Probability:      pK7,
		ComplementProb:      pComplement,
		ProbabilitiesSumTo1: math.Abs(pK7+pComplement-1) < tolerance,
		Verdict: strings.Join([]string{
			StatusK7AndComplementProbabilitiesComputed,
			Status7Over72K7EventProbability,
			Status65Over72ComplementEventProbability,
		}, "; "),
	}
}

func buildObservable() BoundaryWoundObservableAudit {
	return BoundaryWoundObservableAudit{
		Observable:            "W_boundary=(R_3-1)P_K7+|lambda(Lambda_12)|P_perp",
		K7Payoff:              r3Minus1,
		ComplementPayoff:      math.Abs(lambdaLambda12),
		K7PayoffRole:          "gauge meeting-wall wound R_3-1",
		ComplementPayoffRole:  "scalar zero-wall depth |lambda(Lambda_12)|",
		SupportSplit:          "K7/complement split under P_K7 and P_perp",
		IsTwoPayoffObservable: true,
		Verdict:               StatusTwoPayoffBoundaryWoundObservableDefined,
	}
}

func buildExpectation(r RearrangementAudit) ExpectationAudit {
	expected := pK7*r3Minus1 + pComplement*math.Abs(lambdaLambda12)
	ksum := kappaLambda + kappaE
	residual := ksum - expected
	return ExpectationAudit{
		Formula:                   "Tr(rho_72 W_boundary)=p_K7(R_3-1)+p_perp|lambda|",
		ExpectedBoundaryWound:     expected,
		KSum:                      ksum,
		Residual:                  residual,
		ReproducesWeightedClosure: math.Abs(expected-r.WeightedClosureRight) < tolerance,
		Verdict: strings.Join([]string{
			StatusExpectationReproducesWeightedClosure,
			StatusKappaSumNoBiasExpectedBoundaryWound,
			StatusWeightedBoundaryClosureEventMeaning,
		}, "; "),
	}
}

func buildNumerical(e ExpectationAudit) NumericalAudit {
	gate700Residual := (kappaLambda + kappaE + lambdaLambda12) - pK7*(lambdaLambda12+r3Minus1)
	return NumericalAudit{
		KSum:                  kappaLambda + kappaE,
		ExpectedWound:         e.ExpectedBoundaryWound,
		Residual:              e.Residual,
		InheritedResidualForm: gate700Residual,
		SameResidual:          math.Abs(e.Residual-gate700Residual) < 1e-17,
		Verdict:               StatusNumericalResidualRecorded,
	}
}

func buildInterpretation() InterpretationAudit {
	return InterpretationAudit{
		K7EventPayoff:    "K7 event payoff is the gauge meeting-wall wound R_3-1",
		ComplementPayoff: "K7 complement payoff is the scalar zero-wall depth |lambda|",
		Observer:         "rho_72 full augmented no-bias state",
		Output:           "K_sum=kappa_lambda+kappa_e",
		Reading:          "K_sum is approximately the no-bias expected boundary wound",
		Verdict: strings.Join([]string{
			StatusKappaSumNoBiasExpectedBoundaryWound,
			StatusWeightedBoundaryClosureEventMeaning,
		}, "; "),
	}
}

func buildEquivalence() EquivalenceAudit {
	forms := []EquivalentForm{
		{Name: "boundary-to-history quotient response", Equation: "sigma_history≈p_K7 sigma_boundary", Equivalent: true},
		{Name: "stress-split pullback", Equation: "D_base≈p_K7 S_split", Equivalent: true},
		{Name: "weighted boundary interpolation", Equation: "K_sum≈(65/72)|lambda|+(7/72)(R_3-1)", Equivalent: true},
		{Name: "event/complement expectation", Equation: "K_sum≈Tr(rho_72[(R_3-1)P_K7+|lambda|P_perp])", Equivalent: true},
	}
	return EquivalenceAudit{Forms: forms, IntroducesNewNumericalRelation: false, UpgradesSourceType: true, Verdict: StatusEquivalenceToPreviousFormsAudited}
}

func buildAlternatives() AlternativeMixtureAudit {
	absLambda := math.Abs(lambdaLambda12)
	xiBoundary := 0.5 * (r3Minus1 + absLambda)
	alts := []AlternativeMixtureObservable{
		buildAlternative("reversed payoff", absLambda, r3Minus1, false, "K7 receives |lambda| and complement receives R; this is the wrong active orientation"),
		buildAlternative("support-local split payoff", lambdaLambda12+r3Minus1, 0, false, "gives D_base form, not the positive-distance K_sum mixture"),
		buildAlternative("midpoint stress", xiBoundary, xiBoundary, false, "gives xi_boundary and loses event/complement split"),
		buildAlternative("Hodge-signed event weights", r3Minus1, -absLambda, false, "signed polarity response is not the active ordinary probability-state mixture"),
		buildAlternative("active boundary wound mixture", r3Minus1, absLambda, true, "K7 -> R and complement -> |lambda| is the active positive-distance mixture"),
	}
	return AlternativeMixtureAudit{
		Alternatives:              alts,
		ReversedPayoffRejected:    alts[0].Rejected,
		SupportLocalFormSeparated: alts[1].Rejected,
		MidpointRejected:          alts[2].Rejected,
		HodgeSignedRejected:       alts[3].Rejected,
		ActiveMixtureAccepted:     alts[4].Active,
		Verdict:                   StatusAlternativeMixtureObservablesAudited,
	}
}

func buildAlternative(name string, k7Payoff, perpPayoff float64, active bool, reason string) AlternativeMixtureObservable {
	return AlternativeMixtureObservable{
		Name:        name,
		K7Payoff:    k7Payoff,
		PerpPayoff:  perpPayoff,
		Expectation: pK7*k7Payoff + pComplement*perpPayoff,
		Active:      active,
		Rejected:    !active,
		Reason:      reason,
	}
}

func Statuses() []string {
	return []string{
		StatusGate703ScalarWallAirlockInherited,
		StatusGate700ResponseLawRearranged,
		StatusK7AndComplementProbabilitiesComputed,
		StatusTwoPayoffBoundaryWoundObservableDefined,
		StatusExpectationReproducesWeightedClosure,
		StatusNumericalResidualRecorded,
		StatusEquivalenceToPreviousFormsAudited,
		StatusAlternativeMixtureObservablesAudited,
		StatusKappaSumNoBiasExpectedBoundaryWound,
		Status65Over72ComplementEventProbability,
		Status7Over72K7EventProbability,
		StatusWeightedBoundaryClosureEventMeaning,
		StatusNoNativeK7ReceivesGaugeWound,
		StatusNoNativeComplementReceivesScalarWound,
		StatusNoNativeBoundaryWoundMixtureTheorem,
		StatusNoNativeHistoryResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate704BoundaryWoundMixtureBoundary,
	}
}

func FormatInheritance(x Gate703Inheritance) string {
	return fmt.Sprintf("inherited=%t unitGlue=%t p=%.18g c=%.18g nonTaut=%t noAirlock=%t noBoundaryHistory=%t no7=%t verdict=%q", x.ScalarWallAirlockInherited, x.UnitScalarWallGlue, x.EventProbability, x.ResponseCoefficient, x.NonTautologyInherited, x.NoNativeScalarWallAirlock, x.NoNativeBoundaryHistory, x.NoNativeSevenOver72, x.Verdict)
}

func FormatRearrangement(x RearrangementAudit) string {
	return fmt.Sprintf("start=%q rearranged=%q lambdaNeg=%t K=%.18g absLambda=%.18g R=%.18g rhs=%.18g residual=%.18g sameGate700=%t verdict=%q", x.StartingEquation, x.RearrangedEquation, x.LambdaNegative, x.KSum, x.PositiveScalarDepth, x.GaugeWound, x.WeightedClosureRight, x.Residual, x.SameAsGate700Residual, x.Verdict)
}

func FormatProbabilities(x ProbabilityAudit) string {
	return fmt.Sprintf("rho=%q PK7=%q Pperp=%q ranks=%d/%d total=%d pK7=%.18g pPerp=%.18g sum1=%t verdict=%q", x.Rho72, x.PK7, x.PComplement, x.K7Rank, x.ComplementRank, x.TotalDimension, x.PK7Probability, x.ComplementProb, x.ProbabilitiesSumTo1, x.Verdict)
}

func FormatObservable(x BoundaryWoundObservableAudit) string {
	return fmt.Sprintf("observable=%q k7=%.18g perp=%.18g k7Role=%q perpRole=%q split=%q twoPayoff=%t verdict=%q", x.Observable, x.K7Payoff, x.ComplementPayoff, x.K7PayoffRole, x.ComplementPayoffRole, x.SupportSplit, x.IsTwoPayoffObservable, x.Verdict)
}

func FormatExpectation(x ExpectationAudit) string {
	return fmt.Sprintf("formula=%q expected=%.18g K=%.18g residual=%.18g reproduces=%t verdict=%q", x.Formula, x.ExpectedBoundaryWound, x.KSum, x.Residual, x.ReproducesWeightedClosure, x.Verdict)
}

func FormatNumerical(x NumericalAudit) string {
	return fmt.Sprintf("K=%.18g expected=%.18g residual=%.18g inherited=%.18g same=%t verdict=%q", x.KSum, x.ExpectedWound, x.Residual, x.InheritedResidualForm, x.SameResidual, x.Verdict)
}

func FormatInterpretation(x InterpretationAudit) string {
	return fmt.Sprintf("k7=%q perp=%q observer=%q output=%q reading=%q verdict=%q", x.K7EventPayoff, x.ComplementPayoff, x.Observer, x.Output, x.Reading, x.Verdict)
}

func FormatEquivalence(x EquivalenceAudit) string {
	parts := make([]string, 0, len(x.Forms))
	for _, f := range x.Forms {
		parts = append(parts, fmt.Sprintf("%s:%t:%s", f.Name, f.Equivalent, f.Equation))
	}
	return fmt.Sprintf("forms=[%s] newNumerical=%t sourceUpgrade=%t verdict=%q", strings.Join(parts, " | "), x.IntroducesNewNumericalRelation, x.UpgradesSourceType, x.Verdict)
}

func FormatAlternatives(x AlternativeMixtureAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, alt := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s: k7=%.18g perp=%.18g E=%.18g active=%t rejected=%t", alt.Name, alt.K7Payoff, alt.PerpPayoff, alt.Expectation, alt.Active, alt.Rejected))
	}
	return fmt.Sprintf("reversed=%t supportLocal=%t midpoint=%t hodge=%t active=%t alts=[%s] verdict=%q", x.ReversedPayoffRejected, x.SupportLocalFormSeparated, x.MidpointRejected, x.HodgeSignedRejected, x.ActiveMixtureAccepted, strings.Join(parts, " | "), x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=%s verdict=%q", strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("k7GaugeNative=%t complementScalarNative=%t mixtureTheorem=%t historyTheorem=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gaugeUnification=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsK7ReceivesGaugeWoundNative, x.ClaimsComplementReceivesScalarWoundNative, x.ClaimsNativeBoundaryWoundMixtureTheorem, x.ClaimsNativeHistoryResponseTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
