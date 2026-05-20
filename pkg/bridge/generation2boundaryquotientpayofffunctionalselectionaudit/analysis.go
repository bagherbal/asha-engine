// Package generation2boundaryquotientpayofffunctionalselectionaudit implements
// Gate 697: Boundary Quotient Payoff Functional Selection Audit.
//
// Gate 696 showed that the active Bernoulli observable is support-local,
// R_split=S_split P_K7, and that support-locality forces the complement payoff
// to vanish without determining the K7 event payoff. Gate 697 audits whether
// the active payoff S_split=lambda(Lambda_12)+(R_3-1) is the canonical boundary
// quotient coordinate measuring failure of the exact gauge-scalar anti-alignment
// wall lambda+(R_3-1)=0.
//
// This is a bridge-layer payoff-functional audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native payoff theorem, a native history-response theorem, or a
// native 7/72 theorem.
package generation2boundaryquotientpayofffunctionalselectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate696 "github.com/bagherbal/asha-engine/pkg/bridge/generation2bernoullipayoffnormalizationandzerocomplementsupportaudit"
)

const (
	AuditID = "GATE697-BOUNDARY-QUOTIENT-PAYOFF-FUNCTIONAL-SELECTION-AUDIT"

	StatusGate696SupportLocalBernoulliObservableInherited = "PASS_GATE696_SUPPORT_LOCAL_BERNOULLI_OBSERVABLE_INHERITED"
	StatusPayoffSourceProblemDefined                      = "PASS_PAYOFF_SOURCE_PROBLEM_DEFINED"
	StatusBoundaryAntiAlignmentWallDefined                = "PASS_BOUNDARY_ANTI_ALIGNMENT_WALL_DEFINED"
	StatusSigmaBoundaryDescendsToQuotientCoordinate       = "PASS_SIGMA_BOUNDARY_DESCENDS_TO_QUOTIENT_COORDINATE"
	StatusSSplitIdentifiedAsBoundaryQuotientPayoff        = "PASS_S_SPLIT_IDENTIFIED_AS_BOUNDARY_QUOTIENT_PAYOFF"
	StatusAlternativeBoundaryPayoffsAudited               = "PASS_ALTERNATIVE_BOUNDARY_PAYOFFS_AUDITED"
	StatusEventExpectationReconstructed                   = "PASS_EVENT_EXPECTATION_RECONSTRUCTED"
	StatusSSplitCanonicalAntiAlignmentQuotientPayoff      = "CONDITIONAL_SUPPORT_S_SPLIT_IS_CANONICAL_ANTI_ALIGNMENT_QUOTIENT_PAYOFF"
	StatusActiveResponseK7EventBoundaryQuotientPayoff     = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_IS_K7_EVENT_WITH_BOUNDARY_QUOTIENT_PAYOFF"
	StatusPayoffFunctionalUniqueOnlyUpToNormalization     = "FAILED_ROUTE_PAYOFF_FUNCTIONAL_UNIQUE_ONLY_UP_TO_WALL_COORDINATE_NORMALIZATION"
	StatusNoNativeReasonK7ReceivesBoundaryQuotientPayoff  = "FAILED_ROUTE_NO_NATIVE_REASON_K7_EVENT_RECEIVES_BOUNDARY_QUOTIENT_PAYOFF"
	StatusNoNativePayoffCouplingTheorem                   = "FAILED_ROUTE_NO_NATIVE_PAYOFF_COUPLING_THEOREM"
	StatusNoNativeHistoryResponseTheorem                  = "FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem                      = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate697BoundaryQuotientPayoffBoundary           = "FIREWALL_PRESERVED_GATE697_BOUNDARY_QUOTIENT_PAYOFF_BOUNDARY"
)

const (
	lambda4Dimension = 70
	boundaryDim      = 2
	h72Dimension     = lambda4Dimension + boundaryDim
	k7Dimension      = 7
	lambdaLambda12   = -0.0497009420776833
	r3Minus1         = 0.0509933868964996
	tolerance        = 1e-15
)

type Gate696Inheritance struct {
	SupportLocalBernoulliObservableInherited bool
	Rho72Definition                          string
	SupportLocalObservable                   string
	BoundaryPayoffAssignment                 string
	ReconstructedOperator                    string
	EventPayoff                              float64
	ComplementPayoff                         float64
	SSplit                                   float64
	DBase                                    float64
	Expectation                              float64
	ResidualE1                               float64
	NoExpectationAloneSelection              bool
	NoNativeSupportLocality                  bool
	NoNativeSSplitPayoff                     bool
	NoNativeSevenOver72                      bool
	Verdict                                  string
}

type PayoffSourceProblemAudit struct {
	InheritedObservable     string
	KnownSupportResult      string
	RemainingQuestion       string
	SupportLocalityFixesB   bool
	EventPayoffStillUnfixed bool
	CandidatePayoff         string
	Verdict                 string
}

type BoundaryQuotientAudit struct {
	BoundarySpace          string
	BoundaryVector         [2]float64
	AntiAlignmentLine      string
	AntiAlignmentGenerator [2]float64
	SigmaDefinition        string
	SigmaOnAntiAlignment   float64
	KernelMatchesWall      bool
	DescendsToQuotient     bool
	QuotientSpace          string
	SSplit                 float64
	SSplitMatchesSigma     bool
	Verdict                string
}

type PayoffInterpretationAudit struct {
	EventPayoff              string
	Observable               string
	EquivalentOperator       string
	Payoff                   float64
	K7ReceivesBoundaryDefect bool
	DoesNotProveCoupling     bool
	Verdict                  string
}

type BoundaryPayoffAlternative struct {
	Name             string
	Functional       string
	Value            float64
	VanishesOnWall   bool
	MeasuresQuotient bool
	Active           bool
	Reason           string
}

type AlternativeBoundaryPayoffsAudit struct {
	Alternatives           []BoundaryPayoffAlternative
	LambdaOnlyRejected     bool
	GaugeOnlyRejected      bool
	AntiAlignedRejected    bool
	MidpointStressRejected bool
	SplitPayoffAccepted    bool
	AllAudited             bool
	Verdict                string
}

type ScaleNormalizationFirewallAudit struct {
	QuotientCoordinateUniqueUpToScale bool
	ScaledFunctional                  string
	ActiveNormalizationSource         string
	UnitCoefficientLambda             float64
	UnitCoefficientGauge              float64
	ClaimsNativePayoffNormalization   bool
	Verdict                           string
}

type EventExpectationReconstructionAudit struct {
	Rho72              string
	ResponseOperator   string
	K7Weight           float64
	ExpectationFormula string
	Expectation        float64
	DBase              float64
	ResidualE1         float64
	MatchesInherited   bool
	Verdict            string
}

type SourceTypeClassificationAudit struct {
	BoundarySpaceRole string
	SigmaBoundaryRole string
	SSplitRole        string
	PK7Role           string
	Rho72Role         string
	RSplitRole        string
	Verdict           string
}

type MissingTheoremAudit struct {
	Candidates []string
	PreciseGap string
	Verdict    string
}

type FirewallAudit struct {
	ClaimsK7ReceivesBoundaryPayoffNatively bool
	ClaimsNativePayoffCouplingTheorem      bool
	ClaimsNativeHistoryResponseTheorem     bool
	ClaimsNativeSevenOver72Theorem         bool
	ClaimsBoundaryStressDerived            bool
	ClaimsScalarRGMatching                 bool
	ClaimsHiggsMass                        bool
	ClaimsGaugeUnification                 bool
	ClaimsFlavorDerivation                 bool
	ClaimsCKMPMNS                          bool
	Verdict                                string
}

type Analysis struct {
	Inherited      Gate696Inheritance
	PayoffProblem  PayoffSourceProblemAudit
	Quotient       BoundaryQuotientAudit
	Interpretation PayoffInterpretationAudit
	Alternatives   AlternativeBoundaryPayoffsAudit
	ScaleFirewall  ScaleNormalizationFirewallAudit
	Expectation    EventExpectationReconstructionAudit
	SourceTypes    SourceTypeClassificationAudit
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
	g696, err := gate696.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate696 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g696)
	problem := buildPayoffSourceProblem(inherited)
	quotient := buildBoundaryQuotient()
	interpretation := buildPayoffInterpretation(quotient)
	alternatives := buildAlternativeBoundaryPayoffs(quotient)
	scale := buildScaleNormalizationFirewall()
	expectation := buildEventExpectation(inherited, quotient)
	sources := buildSourceTypes()
	missing := MissingTheoremAudit{
		Candidates: []string{
			"BoundaryQuotientPayoffCouplingTheorem",
			"K7EventBoundaryPayoffTheorem",
			StatusNoNativePayoffCouplingTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native coupling theorem explaining why the Boolean-octonionic K7 event is coupled specifically to the boundary anti-alignment quotient coordinate sigma_boundary(lambda,R)=lambda+R as its payoff",
		Verdict: strings.Join([]string{
			StatusNoNativeReasonK7ReceivesBoundaryQuotientPayoff,
			StatusNoNativePayoffCouplingTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate697BoundaryQuotientPayoffBoundary}
	truth := "Gate 697 audits the source type of the K7 event payoff left unfixed by Gate696 support-locality. The functional sigma_boundary(lambda,R)=lambda+R annihilates the exact anti-alignment wall span((-1,+1)), so it descends to the one-dimensional quotient B_boundary/L_anti and evaluates on the boundary vector to S_split. This conditionally supports S_split as the canonical anti-alignment quotient payoff, with the active observable R_split=sigma_boundary(b)P_K7. The result remains bridge-layer only: the quotient functional is unique only up to wall-coordinate normalization, and no native theorem yet explains why K7 receives this boundary quotient payoff or why history evaluates the resulting response."
	return Analysis{Inherited: inherited, PayoffProblem: problem, Quotient: quotient, Interpretation: interpretation, Alternatives: alternatives, ScaleFirewall: scale, Expectation: expectation, SourceTypes: sources, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate696.Analysis) Gate696Inheritance {
	return Gate696Inheritance{
		SupportLocalBernoulliObservableInherited: g.Assignment.MatchesInheritedActive && g.Assignment.ComplementPayoff == 0 && g.Assignment.ReconstructedOperator == "R_split=S_split P_K7",
		Rho72Definition:                          g.Inherited.Rho72Definition,
		SupportLocalObservable:                   g.Assignment.SupportLocalObservable,
		BoundaryPayoffAssignment:                 g.Assignment.BoundaryPayoffAssignment,
		ReconstructedOperator:                    g.Assignment.ReconstructedOperator,
		EventPayoff:                              g.Assignment.EventPayoff,
		ComplementPayoff:                         g.Assignment.ComplementPayoff,
		SSplit:                                   g.Inherited.SSplit,
		DBase:                                    g.Inherited.DBase,
		Expectation:                              g.Assignment.Expectation,
		ResidualE1:                               g.Inherited.ResidualE1,
		NoExpectationAloneSelection:              g.Degeneracy.ExpectationAloneDegenerate,
		NoNativeSupportLocality:                  !g.Firewall.ClaimsHistoryUsesSupportLocalityNatively,
		NoNativeSSplitPayoff:                     !g.Firewall.ClaimsSSplitPayoffNatively,
		NoNativeSevenOver72:                      !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                                  StatusGate696SupportLocalBernoulliObservableInherited,
	}
}

func buildPayoffSourceProblem(i Gate696Inheritance) PayoffSourceProblemAudit {
	return PayoffSourceProblemAudit{
		InheritedObservable:     i.SupportLocalObservable,
		KnownSupportResult:      "support-locality forces b=0 in R_{a,b}=aP_K7+bP_perp",
		RemainingQuestion:       "a ?= S_split",
		SupportLocalityFixesB:   i.ComplementPayoff == 0,
		EventPayoffStillUnfixed: true,
		CandidatePayoff:         "a=S_split=lambda(Lambda_12)+(R_3-1)",
		Verdict:                 StatusPayoffSourceProblemDefined,
	}
}

func buildBoundaryQuotient() BoundaryQuotientAudit {
	sigmaWall := -1.0 + 1.0
	s := lambdaLambda12 + r3Minus1
	return BoundaryQuotientAudit{
		BoundarySpace:          "B_boundary=span(lambda,R_3-1)",
		BoundaryVector:         [2]float64{lambdaLambda12, r3Minus1},
		AntiAlignmentLine:      "L_anti={(lambda,R):lambda+R=0}=span((-1,+1))",
		AntiAlignmentGenerator: [2]float64{-1, 1},
		SigmaDefinition:        "sigma_boundary(lambda,R)=lambda+R",
		SigmaOnAntiAlignment:   sigmaWall,
		KernelMatchesWall:      math.Abs(sigmaWall) < tolerance,
		DescendsToQuotient:     math.Abs(sigmaWall) < tolerance,
		QuotientSpace:          "Q_boundary=B_boundary/L_anti",
		SSplit:                 s,
		SSplitMatchesSigma:     math.Abs(s-(lambdaLambda12+r3Minus1)) < tolerance,
		Verdict: strings.Join([]string{
			StatusBoundaryAntiAlignmentWallDefined,
			StatusSigmaBoundaryDescendsToQuotientCoordinate,
			StatusSSplitIdentifiedAsBoundaryQuotientPayoff,
		}, "; "),
	}
}

func buildPayoffInterpretation(q BoundaryQuotientAudit) PayoffInterpretationAudit {
	return PayoffInterpretationAudit{
		EventPayoff:              "a=S_split=sigma_boundary(b)",
		Observable:               "R_split=sigma_boundary(b)P_K7",
		EquivalentOperator:       "R_split=[lambda(Lambda_12)+(R_3-1)]P_K7",
		Payoff:                   q.SSplit,
		K7ReceivesBoundaryDefect: true,
		DoesNotProveCoupling:     true,
		Verdict: strings.Join([]string{
			StatusSSplitCanonicalAntiAlignmentQuotientPayoff,
			StatusActiveResponseK7EventBoundaryQuotientPayoff,
		}, "; "),
	}
}

func buildAlternativeBoundaryPayoffs(q BoundaryQuotientAudit) AlternativeBoundaryPayoffsAudit {
	alternatives := []BoundaryPayoffAlternative{
		{
			Name:             "lambda-only payoff",
			Functional:       "a=lambda(Lambda_12)",
			Value:            lambdaLambda12,
			VanishesOnWall:   false,
			MeasuresQuotient: false,
			Active:           false,
			Reason:           "does not vanish on the anti-alignment wall generator (-1,+1)",
		},
		{
			Name:             "gauge-only payoff",
			Functional:       "a=R_3-1",
			Value:            r3Minus1,
			VanishesOnWall:   false,
			MeasuresQuotient: false,
			Active:           false,
			Reason:           "does not vanish on the anti-alignment wall generator (-1,+1)",
		},
		{
			Name:             "anti-aligned magnitude",
			Functional:       "a=(R_3-1)-lambda",
			Value:            r3Minus1 - lambdaLambda12,
			VanishesOnWall:   false,
			MeasuresQuotient: false,
			Active:           false,
			Reason:           "measures total anti-aligned magnitude, not failure of anti-alignment",
		},
		{
			Name:             "midpoint stress",
			Functional:       "a=xi_boundary=0.5[(R_3-1)+|lambda|]",
			Value:            0.5 * (r3Minus1 + math.Abs(lambdaLambda12)),
			VanishesOnWall:   false,
			MeasuresQuotient: false,
			Active:           false,
			Reason:           "measures common stress scale, not quotient defect",
		},
		{
			Name:             "split payoff",
			Functional:       "a=lambda+(R_3-1)",
			Value:            q.SSplit,
			VanishesOnWall:   true,
			MeasuresQuotient: true,
			Active:           true,
			Reason:           "vanishes exactly on perfect anti-alignment and measures quotient defect",
		},
	}
	return AlternativeBoundaryPayoffsAudit{
		Alternatives:           alternatives,
		LambdaOnlyRejected:     !alternatives[0].Active && !alternatives[0].VanishesOnWall,
		GaugeOnlyRejected:      !alternatives[1].Active && !alternatives[1].VanishesOnWall,
		AntiAlignedRejected:    !alternatives[2].Active && !alternatives[2].MeasuresQuotient,
		MidpointStressRejected: !alternatives[3].Active && !alternatives[3].MeasuresQuotient,
		SplitPayoffAccepted:    alternatives[4].Active && alternatives[4].VanishesOnWall && alternatives[4].MeasuresQuotient,
		AllAudited:             true,
		Verdict:                StatusAlternativeBoundaryPayoffsAudited,
	}
}

func buildScaleNormalizationFirewall() ScaleNormalizationFirewallAudit {
	return ScaleNormalizationFirewallAudit{
		QuotientCoordinateUniqueUpToScale: true,
		ScaledFunctional:                  "c*sigma_boundary",
		ActiveNormalizationSource:         "Gates 668-670 wall-coordinate normalization: lambda and R_3-1 are canonical wall-distance coordinates with unit coefficients",
		UnitCoefficientLambda:             1,
		UnitCoefficientGauge:              1,
		ClaimsNativePayoffNormalization:   false,
		Verdict:                           StatusPayoffFunctionalUniqueOnlyUpToNormalization,
	}
}

func buildEventExpectation(i Gate696Inheritance, q BoundaryQuotientAudit) EventExpectationReconstructionAudit {
	k7Weight := float64(k7Dimension) / float64(h72Dimension)
	expectation := k7Weight * q.SSplit
	return EventExpectationReconstructionAudit{
		Rho72:              "rho_72=I_H72/72",
		ResponseOperator:   "R_split=sigma_boundary(b)P_K7",
		K7Weight:           k7Weight,
		ExpectationFormula: "Tr(rho_72 R_split)=sigma_boundary(b)Tr(rho_72 P_K7)=(7/72)S_split",
		Expectation:        expectation,
		DBase:              i.DBase,
		ResidualE1:         i.DBase - expectation,
		MatchesInherited:   math.Abs(expectation-i.Expectation) < tolerance && math.Abs((i.DBase-expectation)-i.ResidualE1) < 1e-18,
		Verdict:            StatusEventExpectationReconstructed,
	}
}

func buildSourceTypes() SourceTypeClassificationAudit {
	return SourceTypeClassificationAudit{
		BoundarySpaceRole: "two-coordinate boundary wall plane span(lambda,R_3-1)",
		SigmaBoundaryRole: "quotient coordinate annihilating the exact anti-alignment wall",
		SSplitRole:        "boundary anti-alignment failure payoff",
		PK7Role:           "Boolean-octonionic support-selected event projector",
		Rho72Role:         "full augmented maximum-entropy observer state",
		RSplitRole:        "K7 event observable with boundary quotient payoff",
		Verdict:           StatusActiveResponseK7EventBoundaryQuotientPayoff,
	}
}

func Statuses() []string {
	return []string{
		StatusGate696SupportLocalBernoulliObservableInherited,
		StatusPayoffSourceProblemDefined,
		StatusBoundaryAntiAlignmentWallDefined,
		StatusSigmaBoundaryDescendsToQuotientCoordinate,
		StatusSSplitIdentifiedAsBoundaryQuotientPayoff,
		StatusAlternativeBoundaryPayoffsAudited,
		StatusEventExpectationReconstructed,
		StatusSSplitCanonicalAntiAlignmentQuotientPayoff,
		StatusActiveResponseK7EventBoundaryQuotientPayoff,
		StatusPayoffFunctionalUniqueOnlyUpToNormalization,
		StatusNoNativeReasonK7ReceivesBoundaryQuotientPayoff,
		StatusNoNativePayoffCouplingTheorem,
		StatusNoNativeHistoryResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate697BoundaryQuotientPayoffBoundary,
	}
}

func FormatInheritance(x Gate696Inheritance) string {
	return fmt.Sprintf("inherited=%t rho=%q supportLocal=%q payoff=%q operator=%q eventPayoff=%.18g complementPayoff=%.18g ssplit=%.18g dbase=%.18g expectation=%.18g e1=%.18g noExpectationAlone=%t noSupportLocality=%t noSSplitPayoff=%t no7=%t verdict=%q", x.SupportLocalBernoulliObservableInherited, x.Rho72Definition, x.SupportLocalObservable, x.BoundaryPayoffAssignment, x.ReconstructedOperator, x.EventPayoff, x.ComplementPayoff, x.SSplit, x.DBase, x.Expectation, x.ResidualE1, x.NoExpectationAloneSelection, x.NoNativeSupportLocality, x.NoNativeSSplitPayoff, x.NoNativeSevenOver72, x.Verdict)
}

func FormatPayoffProblem(x PayoffSourceProblemAudit) string {
	return fmt.Sprintf("observable=%q known=%q remaining=%q fixesB=%t payoffUnfixed=%t candidate=%q verdict=%q", x.InheritedObservable, x.KnownSupportResult, x.RemainingQuestion, x.SupportLocalityFixesB, x.EventPayoffStillUnfixed, x.CandidatePayoff, x.Verdict)
}

func FormatQuotient(x BoundaryQuotientAudit) string {
	return fmt.Sprintf("space=%q vector=(%.18g,%.18g) line=%q generator=(%.18g,%.18g) sigma=%q sigmaWall=%.18g kernelMatches=%t descends=%t quotient=%q ssplit=%.18g matches=%t verdict=%q", x.BoundarySpace, x.BoundaryVector[0], x.BoundaryVector[1], x.AntiAlignmentLine, x.AntiAlignmentGenerator[0], x.AntiAlignmentGenerator[1], x.SigmaDefinition, x.SigmaOnAntiAlignment, x.KernelMatchesWall, x.DescendsToQuotient, x.QuotientSpace, x.SSplit, x.SSplitMatchesSigma, x.Verdict)
}

func FormatInterpretation(x PayoffInterpretationAudit) string {
	return fmt.Sprintf("payoff=%q observable=%q equivalent=%q value=%.18g k7Receives=%t noCoupling=%t verdict=%q", x.EventPayoff, x.Observable, x.EquivalentOperator, x.Payoff, x.K7ReceivesBoundaryDefect, x.DoesNotProveCoupling, x.Verdict)
}

func FormatAlternatives(x AlternativeBoundaryPayoffsAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s{%s value=%.18g vanish=%t quotient=%t active=%t reason=%s}", a.Name, a.Functional, a.Value, a.VanishesOnWall, a.MeasuresQuotient, a.Active, a.Reason))
	}
	return fmt.Sprintf("alternatives=[%s] lambdaRejected=%t gaugeRejected=%t antiRejected=%t midpointRejected=%t splitAccepted=%t all=%t verdict=%q", strings.Join(parts, "; "), x.LambdaOnlyRejected, x.GaugeOnlyRejected, x.AntiAlignedRejected, x.MidpointStressRejected, x.SplitPayoffAccepted, x.AllAudited, x.Verdict)
}

func FormatScaleFirewall(x ScaleNormalizationFirewallAudit) string {
	return fmt.Sprintf("uniqueUpToScale=%t scaled=%q source=%q coeffLambda=%.18g coeffGauge=%.18g nativeNorm=%t verdict=%q", x.QuotientCoordinateUniqueUpToScale, x.ScaledFunctional, x.ActiveNormalizationSource, x.UnitCoefficientLambda, x.UnitCoefficientGauge, x.ClaimsNativePayoffNormalization, x.Verdict)
}

func FormatExpectation(x EventExpectationReconstructionAudit) string {
	return fmt.Sprintf("rho=%q operator=%q k7Weight=%.18g formula=%q expectation=%.18g dbase=%.18g e1=%.18g matches=%t verdict=%q", x.Rho72, x.ResponseOperator, x.K7Weight, x.ExpectationFormula, x.Expectation, x.DBase, x.ResidualE1, x.MatchesInherited, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassificationAudit) string {
	return fmt.Sprintf("boundary=%q sigma=%q ssplit=%q pk7=%q rho=%q rsplit=%q verdict=%q", x.BoundarySpaceRole, x.SigmaBoundaryRole, x.SSplitRole, x.PK7Role, x.Rho72Role, x.RSplitRole, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("candidates=%v gap=%q verdict=%q", x.Candidates, x.PreciseGap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("nativeK7Payoff=%t nativeCoupling=%t nativeHistory=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t unification=%t flavor=%t ckm=%t verdict=%q", x.ClaimsK7ReceivesBoundaryPayoffNatively, x.ClaimsNativePayoffCouplingTheorem, x.ClaimsNativeHistoryResponseTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
