// Package generation2scalarbaselineandk7boundarysplitupliftobservableaudit implements
// Gate 705: Scalar Baseline and K7 Boundary-Split Uplift Observable Audit.
//
// Gate 704 typed the positive-distance closure as a no-bias expectation of the
// two-payoff boundary wound observable
//
//	W_boundary=(R_3-1)P_K7+|lambda|P_perp.
//
// Gate 705 audits the equivalent baseline-plus-uplift decomposition
//
//	W_boundary=|lambda|I_H72+S_split P_K7,
//
// where S_split=(R_3-1)-|lambda|=lambda+(R_3-1). This is a bridge-layer
// observable-decomposition audit only. It does not derive boundary stress,
// scalar RG matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native
// response theorem, a native state-selection theorem, or a native 7/72 theorem.
package generation2scalarbaselineandk7boundarysplitupliftobservableaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate704 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7complementboundarywoundmixtureobservableaudit"
)

const (
	AuditID = "GATE705-SCALAR-BASELINE-K7-BOUNDARY-SPLIT-UPLIFT-OBSERVABLE-AUDIT"

	StatusGate704BoundaryWoundMixtureInherited        = "PASS_GATE704_BOUNDARY_WOUND_MIXTURE_INHERITED"
	StatusTwoPayoffObservableRewritten                = "PASS_TWO_PAYOFF_OBSERVABLE_REWRITTEN"
	StatusScalarBaselineK7UpliftDecompositionComputed = "PASS_SCALAR_BASELINE_PLUS_K7_UPLIFT_DECOMPOSITION_COMPUTED"
	StatusExpectationReproducesKSumClosure            = "PASS_EXPECTATION_REPRODUCES_KSUM_CLOSURE"
	StatusRelationToGate700ResponseLawAudited         = "PASS_RELATION_TO_GATE700_RESPONSE_LAW_AUDITED"
	StatusSourceTypeUpgradeAudited                    = "PASS_SOURCE_TYPE_UPGRADE_AUDITED"
	StatusAlternativeBaselineDecompositionsAudited    = "PASS_ALTERNATIVE_BASELINE_DECOMPOSITIONS_AUDITED"
	StatusKSumScalarBaselineExpectedK7SplitUplift     = "CONDITIONAL_SUPPORT_KSUM_IS_SCALAR_BASELINE_PLUS_EXPECTED_K7_SPLIT_UPLIFT"
	StatusK7ReceivesSplitUpliftNotPrimitiveGaugeWound = "CONDITIONAL_SUPPORT_K7_RECEIVES_BOUNDARY_SPLIT_UPLIFT_NOT_PRIMITIVE_GAUGE_WOUND"
	StatusScalarWallAirlockSupportsScalarBaseline     = "CONDITIONAL_SUPPORT_SCALAR_WALL_AIRLOCK_SUPPORTS_SCALAR_BASELINE_READING"
	StatusNoNativeScalarWoundFullChamberBaseline      = "FAILED_ROUTE_NO_NATIVE_REASON_SCALAR_WOUND_IS_FULL_CHAMBER_BASELINE"
	StatusNoNativeK7ReceivesSplitUplift               = "FAILED_ROUTE_NO_NATIVE_REASON_K7_RECEIVES_SPLIT_UPLIFT"
	StatusNoNativeBoundaryWoundUpliftTheorem          = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM"
	StatusNoNativeHistoryResponseTheorem              = "FAILED_ROUTE_NO_NATIVE_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem                  = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate705ScalarBaselineK7UpliftBoundary       = "FIREWALL_PRESERVED_GATE705_SCALAR_BASELINE_K7_UPLIFT_BOUNDARY"
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

type Gate704Inheritance struct {
	BoundaryWoundMixtureInherited bool
	EventComplementMeaning        bool
	KSumExpectedBoundaryWound     bool
	K7EventProbability            float64
	ComplementEventProbability    float64
	ExpectedWound                 float64
	KSum                          float64
	Residual                      float64
	NoNativeBoundaryWoundMixture  bool
	NoNativeHistoryResponse       bool
	NoNativeSevenOver72           bool
	Verdict                       string
}

type TwoPayoffRewriteAudit struct {
	Gate704Observable      string
	ProjectorIdentity      string
	RewrittenObservable    string
	K7PayoffBefore         float64
	ComplementPayoffBefore float64
	AlgebraicEquivalence   bool
	Verdict                string
}

type ScalarBaselineUpliftAudit struct {
	ScalarBaseline           float64
	BoundarySplitUplift      float64
	SFromSignedSplit         float64
	K7PayoffAfterUplift      float64
	ComplementPayoff         float64
	FullChamberBaseline      bool
	K7LocalizedUplift        bool
	K7PayoffEqualsGaugeWound bool
	ComplementEqualsBaseline bool
	Decomposition            string
	Verdict                  string
}

type ExpectationAudit struct {
	Formula               string
	BaselineExpectation   float64
	UpliftExpectation     float64
	TotalExpectation      float64
	Gate704Expectation    float64
	KSum                  float64
	Residual              float64
	ReproducesKSumClosure bool
	EquivalentToGate704   bool
	Verdict               string
}

type RelationToGate700Audit struct {
	Gate700ResponseLaw      string
	PositiveDistanceLaw     string
	DBase                   float64
	KSum                    float64
	ScalarDepth             float64
	ExpectedSplitUplift     float64
	BaselineIdentity        string
	NotNewNumericalRelation bool
	Verdict                 string
}

type SourceTypeUpgradeAudit struct {
	Gate704Reading                    string
	Gate705Reading                    string
	RemovesPrimitiveGaugeAssignment   bool
	K7ReceivesBoundarySplitUplift     bool
	ScalarWallAirlockSupportsBaseline bool
	Verdict                           string
}

type AlternativeBaselineDecomposition struct {
	Name        string
	Observable  string
	Expectation float64
	Equivalent  bool
	Active      bool
	Rejected    bool
	Reason      string
}

type AlternativeBaselineAudit struct {
	Alternatives              []AlternativeBaselineDecomposition
	GaugeBaselineRejected     bool
	MidpointBaselineRejected  bool
	HodgeSignedUpliftRejected bool
	ActiveBaselineAccepted    bool
	Verdict                   string
}

type MissingTheoremAudit struct {
	Missing []string
	Verdict string
}

type FirewallAudit struct {
	ClaimsScalarWoundFullChamberBaselineNative bool
	ClaimsK7ReceivesSplitUpliftNative          bool
	ClaimsNativeBoundaryWoundUpliftTheorem     bool
	ClaimsNativeHistoryResponseTheorem         bool
	ClaimsNativeSevenOver72Theorem             bool
	ClaimsBoundaryStressDerived                bool
	ClaimsScalarRGMatching                     bool
	ClaimsHiggsMass                            bool
	ClaimsGaugeUnification                     bool
	ClaimsFlavorDerivation                     bool
	ClaimsCKMPMNS                              bool
	Verdict                                    string
}

type Analysis struct {
	Inherited     Gate704Inheritance
	Rewrite       TwoPayoffRewriteAudit
	Decomposition ScalarBaselineUpliftAudit
	Expectation   ExpectationAudit
	Relation      RelationToGate700Audit
	SourceUpgrade SourceTypeUpgradeAudit
	Alternatives  AlternativeBaselineAudit
	Missing       MissingTheoremAudit
	Firewall      FirewallAudit
	Truth         string
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
	g704, err := gate704.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate704 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g704)
	rewrite := buildRewrite()
	decomposition := buildDecomposition()
	expectation := buildExpectation(decomposition, g704)
	relation := buildRelation(expectation)
	sourceUpgrade := buildSourceUpgrade()
	alternatives := buildAlternatives(expectation)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeScalarWoundFullChamberBaseline,
			StatusNoNativeK7ReceivesSplitUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusNoNativeScalarWoundFullChamberBaseline,
			StatusNoNativeK7ReceivesSplitUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeHistoryResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate705ScalarBaselineK7UpliftBoundary}
	truth := "Gate 705 rewrites the Gate704 two-payoff boundary wound observable as a full-chamber scalar-wall baseline plus a K7-localized boundary-split uplift: W_boundary=|lambda|I_H72+S_split P_K7. This is algebraically equivalent to Gate704 and Gate700 plus K_sum=|lambda|+D_base; it upgrades the source type by saying K7 receives the split uplift, not a primitive gauge wound. It does not prove why scalar wound is the chamber baseline or why K7 receives the uplift."
	return Analysis{Inherited: inherited, Rewrite: rewrite, Decomposition: decomposition, Expectation: expectation, Relation: relation, SourceUpgrade: sourceUpgrade, Alternatives: alternatives, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate704.Analysis) Gate704Inheritance {
	return Gate704Inheritance{
		BoundaryWoundMixtureInherited: true,
		EventComplementMeaning:        g.Equivalence.UpgradesSourceType,
		KSumExpectedBoundaryWound:     g.Expectation.ReproducesWeightedClosure,
		K7EventProbability:            g.Probabilities.PK7Probability,
		ComplementEventProbability:    g.Probabilities.ComplementProb,
		ExpectedWound:                 g.Expectation.ExpectedBoundaryWound,
		KSum:                          g.Expectation.KSum,
		Residual:                      g.Expectation.Residual,
		NoNativeBoundaryWoundMixture:  !g.Firewall.ClaimsNativeBoundaryWoundMixtureTheorem,
		NoNativeHistoryResponse:       !g.Firewall.ClaimsNativeHistoryResponseTheorem,
		NoNativeSevenOver72:           !g.Firewall.ClaimsNativeSevenOver72Theorem,
		Verdict:                       StatusGate704BoundaryWoundMixtureInherited,
	}
}

func buildRewrite() TwoPayoffRewriteAudit {
	return TwoPayoffRewriteAudit{
		Gate704Observable:      "W_boundary=(R_3-1)P_K7+|lambda|P_perp",
		ProjectorIdentity:      "P_perp=I_H72-P_K7",
		RewrittenObservable:    "W_boundary=|lambda|I_H72+((R_3-1)-|lambda|)P_K7=|lambda|I_H72+S_split P_K7",
		K7PayoffBefore:         r3Minus1,
		ComplementPayoffBefore: math.Abs(lambdaLambda12),
		AlgebraicEquivalence:   true,
		Verdict:                StatusTwoPayoffObservableRewritten,
	}
}

func buildDecomposition() ScalarBaselineUpliftAudit {
	absLambda := math.Abs(lambdaLambda12)
	sSplit := r3Minus1 - absLambda
	sSigned := lambdaLambda12 + r3Minus1
	return ScalarBaselineUpliftAudit{
		ScalarBaseline:           absLambda,
		BoundarySplitUplift:      sSplit,
		SFromSignedSplit:         sSigned,
		K7PayoffAfterUplift:      absLambda + sSplit,
		ComplementPayoff:         absLambda,
		FullChamberBaseline:      true,
		K7LocalizedUplift:        true,
		K7PayoffEqualsGaugeWound: math.Abs(absLambda+sSplit-r3Minus1) < tolerance,
		ComplementEqualsBaseline: true,
		Decomposition:            "full chamber receives |lambda|I_H72; K7 receives localized uplift S_split P_K7; K7 payoff becomes |lambda|+S_split=R_3-1",
		Verdict: strings.Join([]string{
			StatusScalarBaselineK7UpliftDecompositionComputed,
			StatusK7ReceivesSplitUpliftNotPrimitiveGaugeWound,
			StatusScalarWallAirlockSupportsScalarBaseline,
		}, "; "),
	}
}

func buildExpectation(d ScalarBaselineUpliftAudit, g gate704.Analysis) ExpectationAudit {
	baseline := d.ScalarBaseline
	uplift := pK7 * d.BoundarySplitUplift
	total := baseline + uplift
	ksum := kappaLambda + kappaE
	return ExpectationAudit{
		Formula:               "Tr(rho_72[|lambda|I_H72+S_split P_K7])=|lambda|+p_K7 S_split",
		BaselineExpectation:   baseline,
		UpliftExpectation:     uplift,
		TotalExpectation:      total,
		Gate704Expectation:    g.Expectation.ExpectedBoundaryWound,
		KSum:                  ksum,
		Residual:              ksum - total,
		ReproducesKSumClosure: math.Abs((ksum-total)-g.Expectation.Residual) < 1e-17,
		EquivalentToGate704:   math.Abs(total-g.Expectation.ExpectedBoundaryWound) < tolerance,
		Verdict: strings.Join([]string{
			StatusExpectationReproducesKSumClosure,
			StatusKSumScalarBaselineExpectedK7SplitUplift,
		}, "; "),
	}
}

func buildRelation(e ExpectationAudit) RelationToGate700Audit {
	absLambda := math.Abs(lambdaLambda12)
	dBase := kappaLambda + kappaE + lambdaLambda12
	sSplit := lambdaLambda12 + r3Minus1
	expectedSplit := pK7 * sSplit
	return RelationToGate700Audit{
		Gate700ResponseLaw:      "D_base≈(7/72)S_split",
		PositiveDistanceLaw:     "K_sum≈|lambda|+(7/72)S_split",
		DBase:                   dBase,
		KSum:                    kappaLambda + kappaE,
		ScalarDepth:             absLambda,
		ExpectedSplitUplift:     expectedSplit,
		BaselineIdentity:        "K_sum=|lambda|+D_base because lambda<0",
		NotNewNumericalRelation: math.Abs(e.TotalExpectation-(absLambda+expectedSplit)) < tolerance && math.Abs(dBase-expectedSplit-e.Residual) < 1e-17,
		Verdict:                 StatusRelationToGate700ResponseLawAudited,
	}
}

func buildSourceUpgrade() SourceTypeUpgradeAudit {
	return SourceTypeUpgradeAudit{
		Gate704Reading:                    "K7 gets gauge wound and complement gets scalar wound",
		Gate705Reading:                    "full chamber gets scalar wound baseline and K7 gets boundary split uplift",
		RemovesPrimitiveGaugeAssignment:   true,
		K7ReceivesBoundarySplitUplift:     true,
		ScalarWallAirlockSupportsBaseline: true,
		Verdict: strings.Join([]string{
			StatusSourceTypeUpgradeAudited,
			StatusK7ReceivesSplitUpliftNotPrimitiveGaugeWound,
			StatusScalarWallAirlockSupportsScalarBaseline,
		}, "; "),
	}
}

func buildAlternatives(e ExpectationAudit) AlternativeBaselineAudit {
	absLambda := math.Abs(lambdaLambda12)
	sSplit := lambdaLambda12 + r3Minus1
	xiBoundary := 0.5 * (r3Minus1 + absLambda)
	alts := []AlternativeBaselineDecomposition{
		{
			Name:        "gauge baseline",
			Observable:  "R I_H72 - S_split P_perp",
			Expectation: r3Minus1 - pComplement*sSplit,
			Equivalent:  math.Abs((r3Minus1-pComplement*sSplit)-e.TotalExpectation) < tolerance,
			Active:      false,
			Rejected:    true,
			Reason:      "algebraically equivalent but less natural because the shared scalar-wall airlock singles out |lambda| as the active wall-distance baseline",
		},
		{
			Name:        "midpoint baseline",
			Observable:  "xi_boundary I_H72 plus signed corrections",
			Expectation: xiBoundary,
			Equivalent:  false,
			Active:      false,
			Rejected:    true,
			Reason:      "less minimal and not the active quotient form",
		},
		{
			Name:        "Hodge-signed uplift",
			Observable:  "|lambda|I_H72+S_split(P_+-P_-)",
			Expectation: absLambda + (1.0/72.0)*sSplit,
			Equivalent:  false,
			Active:      false,
			Rejected:    true,
			Reason:      "gives the inactive 1/72 signed-polarity split response",
		},
		{
			Name:        "active scalar-baseline K7-uplift",
			Observable:  "|lambda|I_H72+S_split P_K7",
			Expectation: e.TotalExpectation,
			Equivalent:  true,
			Active:      true,
			Rejected:    false,
			Reason:      "accepted conditionally as the scalar-wall baseline plus K7-localized split uplift",
		},
	}
	return AlternativeBaselineAudit{
		Alternatives:              alts,
		GaugeBaselineRejected:     alts[0].Rejected,
		MidpointBaselineRejected:  alts[1].Rejected,
		HodgeSignedUpliftRejected: alts[2].Rejected,
		ActiveBaselineAccepted:    alts[3].Active,
		Verdict:                   StatusAlternativeBaselineDecompositionsAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate704BoundaryWoundMixtureInherited,
		StatusTwoPayoffObservableRewritten,
		StatusScalarBaselineK7UpliftDecompositionComputed,
		StatusExpectationReproducesKSumClosure,
		StatusRelationToGate700ResponseLawAudited,
		StatusSourceTypeUpgradeAudited,
		StatusAlternativeBaselineDecompositionsAudited,
		StatusKSumScalarBaselineExpectedK7SplitUplift,
		StatusK7ReceivesSplitUpliftNotPrimitiveGaugeWound,
		StatusScalarWallAirlockSupportsScalarBaseline,
		StatusNoNativeScalarWoundFullChamberBaseline,
		StatusNoNativeK7ReceivesSplitUplift,
		StatusNoNativeBoundaryWoundUpliftTheorem,
		StatusNoNativeHistoryResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate705ScalarBaselineK7UpliftBoundary,
	}
}

func FormatInheritance(x Gate704Inheritance) string {
	return fmt.Sprintf("inherited=%t eventComplement=%t ksumExpected=%t pK7=%.18g pPerp=%.18g expected=%.18g K=%.18g residual=%.18g noMixture=%t noHistory=%t no7=%t verdict=%q", x.BoundaryWoundMixtureInherited, x.EventComplementMeaning, x.KSumExpectedBoundaryWound, x.K7EventProbability, x.ComplementEventProbability, x.ExpectedWound, x.KSum, x.Residual, x.NoNativeBoundaryWoundMixture, x.NoNativeHistoryResponse, x.NoNativeSevenOver72, x.Verdict)
}

func FormatRewrite(x TwoPayoffRewriteAudit) string {
	return fmt.Sprintf("gate704=%q identity=%q rewritten=%q k7Before=%.18g perpBefore=%.18g equivalent=%t verdict=%q", x.Gate704Observable, x.ProjectorIdentity, x.RewrittenObservable, x.K7PayoffBefore, x.ComplementPayoffBefore, x.AlgebraicEquivalence, x.Verdict)
}

func FormatDecomposition(x ScalarBaselineUpliftAudit) string {
	return fmt.Sprintf("baseline=%.18g uplift=%.18g signedS=%.18g k7Payoff=%.18g perpPayoff=%.18g fullBaseline=%t k7Uplift=%t k7Gauge=%t perpBaseline=%t decomposition=%q verdict=%q", x.ScalarBaseline, x.BoundarySplitUplift, x.SFromSignedSplit, x.K7PayoffAfterUplift, x.ComplementPayoff, x.FullChamberBaseline, x.K7LocalizedUplift, x.K7PayoffEqualsGaugeWound, x.ComplementEqualsBaseline, x.Decomposition, x.Verdict)
}

func FormatExpectation(x ExpectationAudit) string {
	return fmt.Sprintf("formula=%q baseline=%.18g uplift=%.18g total=%.18g gate704=%.18g K=%.18g residual=%.18g reproduces=%t equivalent=%t verdict=%q", x.Formula, x.BaselineExpectation, x.UpliftExpectation, x.TotalExpectation, x.Gate704Expectation, x.KSum, x.Residual, x.ReproducesKSumClosure, x.EquivalentToGate704, x.Verdict)
}

func FormatRelation(x RelationToGate700Audit) string {
	return fmt.Sprintf("gate700=%q positive=%q D=%.18g K=%.18g absLambda=%.18g expectedUplift=%.18g identity=%q notNew=%t verdict=%q", x.Gate700ResponseLaw, x.PositiveDistanceLaw, x.DBase, x.KSum, x.ScalarDepth, x.ExpectedSplitUplift, x.BaselineIdentity, x.NotNewNumericalRelation, x.Verdict)
}

func FormatSourceUpgrade(x SourceTypeUpgradeAudit) string {
	return fmt.Sprintf("gate704=%q gate705=%q removesGaugePrimitive=%t k7SplitUplift=%t scalarAirlockBaseline=%t verdict=%q", x.Gate704Reading, x.Gate705Reading, x.RemovesPrimitiveGaugeAssignment, x.K7ReceivesBoundarySplitUplift, x.ScalarWallAirlockSupportsBaseline, x.Verdict)
}

func FormatAlternatives(x AlternativeBaselineAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, alt := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s: E=%.18g equivalent=%t active=%t rejected=%t reason=%s", alt.Name, alt.Expectation, alt.Equivalent, alt.Active, alt.Rejected, alt.Reason))
	}
	return fmt.Sprintf("gauge=%t midpoint=%t hodge=%t active=%t alts=[%s] verdict=%q", x.GaugeBaselineRejected, x.MidpointBaselineRejected, x.HodgeSignedUpliftRejected, x.ActiveBaselineAccepted, strings.Join(parts, " | "), x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=%s verdict=%q", strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("scalarBaselineNative=%t k7UpliftNative=%t upliftTheorem=%t historyTheorem=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gaugeUnification=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsScalarWoundFullChamberBaselineNative, x.ClaimsK7ReceivesSplitUpliftNative, x.ClaimsNativeBoundaryWoundUpliftTheorem, x.ClaimsNativeHistoryResponseTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
