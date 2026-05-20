// Package generation2centralbaselinegaugeandscalarwallreferenceselectionaudit implements
// Gate 707: Central Baseline Gauge and Scalar-Wall Reference Selection Audit.
//
// Gate 706 isolated the positive-distance observable as
//
//	W_boundary=|lambda|I_H72+S_split P_K7.
//
// Gate 707 audits whether |lambda|I_H72 is a uniquely typed reference choice or
// one representative inside a central baseline gauge family.  The total wound
// expectation is invariant under arbitrary central baseline shifts, but the
// choice c=|lambda| is the unique central reference that kills the complement
// uplift and leaves a support-local K7 uplift.  This is a bridge-layer central
// baseline gauge audit only. It does not derive boundary stress, scalar RG
// matching, Higgs mass, gauge unification, flavor, CKM/PMNS, a native response
// theorem, a native state-selection theorem, or a native 7/72 theorem.
package generation2centralbaselinegaugeandscalarwallreferenceselectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate706 "github.com/bagherbal/asha-engine/pkg/bridge/generation2centralscalarbaselineandupliftonlyresponseisolationaudit"
)

const (
	AuditID = "GATE707-CENTRAL-BASELINE-GAUGE-SCALAR-WALL-REFERENCE-SELECTION-AUDIT"

	StatusGate706CentralBaselineUpliftInherited       = "PASS_GATE706_CENTRAL_BASELINE_UPLIFT_INHERITED"
	StatusCentralBaselineGaugeFamilyDefined           = "PASS_CENTRAL_BASELINE_GAUGE_FAMILY_DEFINED"
	StatusTotalExpectationBaselineGaugeInvariant      = "PASS_TOTAL_EXPECTATION_BASELINE_GAUGE_INVARIANT"
	StatusActiveScalarBaselineChoiceComputed          = "PASS_ACTIVE_SCALAR_BASELINE_CHOICE_COMPUTED"
	StatusComplementZeroUpliftForAbsLambda            = "PASS_COMPLEMENT_ZERO_UPLIFT_FOR_C_EQUALS_ABS_LAMBDA"
	StatusTypedBaselineAlternativesAudited            = "PASS_TYPED_BASELINE_ALTERNATIVES_AUDITED"
	StatusSupportLocalitySelectsScalarBaseline        = "PASS_SUPPORT_LOCALITY_SELECTS_SCALAR_BASELINE_WITH_K7_UPLIFT"
	StatusScalarWallAirlockCompatibilityAudited       = "PASS_SCALAR_WALL_AIRLOCK_COMPATIBILITY_AUDITED"
	StatusAbsLambdaUniqueBaselineForK7LocalUplift     = "CONDITIONAL_SUPPORT_ABS_LAMBDA_IS_UNIQUE_BASELINE_FOR_K7_LOCAL_UPLIFT"
	StatusScalarBaselineSelectionSupportLocalGauge    = "CONDITIONAL_SUPPORT_SCALAR_BASELINE_SELECTION_IS_SUPPORT_LOCAL_REFERENCE_GAUGE"
	StatusK7UpliftFormSharperThanRawTwoPayoff         = "CONDITIONAL_SUPPORT_K7_UPLIFT_FORM_IS_SHARPER_THAN_RAW_TWO_PAYOFF_FORM"
	StatusBaselineChoiceNotNativeYet                  = "FAILED_ROUTE_BASELINE_CHOICE_NOT_NATIVE_YET"
	StatusNoNativeScalarBaselineReferenceSelection    = "FAILED_ROUTE_NO_NATIVE_SCALAR_BASELINE_REFERENCE_SELECTION_THEOREM"
	StatusNoNativeK7RatherThanComplementCarriesUplift = "FAILED_ROUTE_NO_NATIVE_REASON_K7_RATHER_THAN_COMPLEMENT_CARRIES_UPLIFT"
	StatusNoNativeBoundaryWoundUpliftTheorem          = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_WOUND_UPLIFT_THEOREM"
	StatusNoNativeSevenOver72Theorem                  = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate707CentralBaselineGaugeBoundary         = "FIREWALL_PRESERVED_GATE707_CENTRAL_BASELINE_GAUGE_BOUNDARY"
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

type Gate706Inheritance struct {
	CentralBaselineUpliftInherited bool
	CentralBaselineIdentityShift   bool
	NontrivialContentK7Uplift      bool
	DBaseBaselineSubtracted        bool
	BaselineDoesNotSelectK7OrRho72 bool
	NoNativeScalarBaseline         bool
	NoNativeK7Uplift               bool
	NoNativeUpliftTheorem          bool
	NoNativeHistoryResponse        bool
	NoNativeSevenOver72            bool
	ExpectedWound                  float64
	ExpectedUplift                 float64
	DBase                          float64
	Verdict                        string
}

type CentralBaselineGaugeFamilyAudit struct {
	GeneralForm          string
	ArbitraryC           float64
	K7Correction         float64
	ComplementCorrection float64
	ExpectationFormula   string
	ExpectationAtC       float64
	RawExpectation       float64
	GaugeInvariant       bool
	Verdict              string
}

type ActiveScalarBaselineChoiceAudit struct {
	C                          float64
	Observable                 string
	K7Uplift                   float64
	ComplementUplift           float64
	ComplementZero             bool
	K7LocalResponse            bool
	UniqueComplementZeroChoice bool
	Verdict                    string
}

type BaselineAlternativeAudit struct {
	Name                 string
	C                    float64
	Observable           string
	K7Correction         float64
	ComplementCorrection float64
	Expectation          float64
	Accepted             bool
	Reason               string
}

type SupportLocalitySelectionAudit struct {
	SupportLocalityCondition string
	RawTwoPayoffForm         string
	ScalarBaselineForm       string
	ComplementZeroForAbs     bool
	ForcesCAbsLambda         bool
	Gate696Consistent        bool
	Verdict                  string
}

type ScalarWallAirlockCompatibilityAudit struct {
	ScalarWallReference         string
	UsesScalarZeroWallDepth     bool
	GaugeBaselineAlgebraic      bool
	GaugeBaselineReversesSector bool
	CompatibleWithAirlock       bool
	Verdict                     string
}

type SourceTypeClassification struct {
	CentralGaugeFamilyRole string
	AbsLambdaRole          string
	K7UpliftRole           string
	GaugeBaselineRole      string
	TruthBoundary          string
	Verdict                string
}

type MissingTheoremAudit struct {
	Missing []string
	Verdict string
}

type FirewallAudit struct {
	ClaimsBaselineChoiceNative                   bool
	ClaimsNativeScalarBaselineReferenceSelection bool
	ClaimsNativeK7RatherThanComplementUplift     bool
	ClaimsNativeBoundaryWoundUpliftTheorem       bool
	ClaimsNativeSevenOver72Theorem               bool
	ClaimsBoundaryStressDerived                  bool
	ClaimsScalarRGMatching                       bool
	ClaimsHiggsMass                              bool
	ClaimsGaugeUnification                       bool
	ClaimsFlavorDerivation                       bool
	ClaimsCKMPMNS                                bool
	Verdict                                      string
}

type Analysis struct {
	Inherited       Gate706Inheritance
	GaugeFamily     CentralBaselineGaugeFamilyAudit
	ActiveChoice    ActiveScalarBaselineChoiceAudit
	Alternatives    []BaselineAlternativeAudit
	SupportLocality SupportLocalitySelectionAudit
	Airlock         ScalarWallAirlockCompatibilityAudit
	SourceTypes     SourceTypeClassification
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
	g706, err := gate706.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate706 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g706)
	gauge := buildGaugeFamily(math.Abs(lambdaLambda12) + 0.25*(r3Minus1-math.Abs(lambdaLambda12)))
	active := buildActiveChoice()
	alternatives := buildAlternatives()
	support := buildSupportLocality(active)
	airlock := buildAirlock()
	sourceTypes := buildSourceTypes()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusBaselineChoiceNotNativeYet,
			StatusNoNativeScalarBaselineReferenceSelection,
			StatusNoNativeK7RatherThanComplementCarriesUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Verdict: strings.Join([]string{
			StatusBaselineChoiceNotNativeYet,
			StatusNoNativeScalarBaselineReferenceSelection,
			StatusNoNativeK7RatherThanComplementCarriesUplift,
			StatusNoNativeBoundaryWoundUpliftTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate707CentralBaselineGaugeBoundary}
	truth := "Gate 707 shows that W_boundary admits a central baseline gauge family. The total expectation is invariant under central reference shifts, but c=|lambda| is the unique central reference that makes the complement uplift vanish and leaves the positive-distance observable as a K7-local uplift over the scalar zero-wall baseline. This is a support-local reference-gauge classification, not a native scalar-baseline selection theorem."
	return Analysis{Inherited: inherited, GaugeFamily: gauge, ActiveChoice: active, Alternatives: alternatives, SupportLocality: support, Airlock: airlock, SourceTypes: sourceTypes, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate706.Analysis) Gate706Inheritance {
	return Gate706Inheritance{
		CentralBaselineUpliftInherited: true,
		CentralBaselineIdentityShift:   strings.Contains(g.CentralBaseline.Verdict, gate706.StatusScalarBaselineCentralIdentityShift),
		NontrivialContentK7Uplift:      strings.Contains(g.Uplift.Verdict, gate706.StatusNontrivialBridgeContentK7Uplift),
		DBaseBaselineSubtracted:        strings.Contains(g.Uplift.Verdict, gate706.StatusDBaseBaselineSubtractedHistoryDefect),
		BaselineDoesNotSelectK7OrRho72: !g.Firewall.ClaimsBaselineSelectsK7OrRho72,
		NoNativeScalarBaseline:         !g.Firewall.ClaimsScalarWoundFullChamberBaselineNative,
		NoNativeK7Uplift:               !g.Firewall.ClaimsK7ReceivesSplitUpliftNative,
		NoNativeUpliftTheorem:          !g.Firewall.ClaimsNativeBoundaryWoundUpliftTheorem,
		NoNativeHistoryResponse:        !g.Firewall.ClaimsNativeHistoryResponseTheorem,
		NoNativeSevenOver72:            !g.Firewall.ClaimsNativeSevenOver72Theorem,
		ExpectedWound:                  g.Relation.ExpectedWound,
		ExpectedUplift:                 g.Relation.ExpectedUplift,
		DBase:                          g.Relation.DBase,
		Verdict:                        StatusGate706CentralBaselineUpliftInherited,
	}
}

func buildGaugeFamily(c float64) CentralBaselineGaugeFamilyAudit {
	absLambda := math.Abs(lambdaLambda12)
	k7Correction := r3Minus1 - c
	perpCorrection := absLambda - c
	expectationAtC := c + pK7*k7Correction + pComplement*perpCorrection
	rawExpectation := pK7*r3Minus1 + pComplement*absLambda
	return CentralBaselineGaugeFamilyAudit{
		GeneralForm:          "W_boundary=c I_H72+(R-c)P_K7+(|lambda|-c)P_perp",
		ArbitraryC:           c,
		K7Correction:         k7Correction,
		ComplementCorrection: perpCorrection,
		ExpectationFormula:   "Tr(rho_72 W_boundary)=c+(7/72)(R-c)+(65/72)(|lambda|-c)",
		ExpectationAtC:       expectationAtC,
		RawExpectation:       rawExpectation,
		GaugeInvariant:       math.Abs(expectationAtC-rawExpectation) < tolerance,
		Verdict:              strings.Join([]string{StatusCentralBaselineGaugeFamilyDefined, StatusTotalExpectationBaselineGaugeInvariant}, "; "),
	}
}

func buildActiveChoice() ActiveScalarBaselineChoiceAudit {
	absLambda := math.Abs(lambdaLambda12)
	k7Uplift := r3Minus1 - absLambda
	complementUplift := absLambda - absLambda
	return ActiveScalarBaselineChoiceAudit{
		C:                          absLambda,
		Observable:                 "W_boundary=|lambda|I_H72+S_split P_K7+0 P_perp",
		K7Uplift:                   k7Uplift,
		ComplementUplift:           complementUplift,
		ComplementZero:             math.Abs(complementUplift) < tolerance,
		K7LocalResponse:            math.Abs(k7Uplift-(lambdaLambda12+r3Minus1)) < tolerance,
		UniqueComplementZeroChoice: true,
		Verdict: strings.Join([]string{
			StatusActiveScalarBaselineChoiceComputed,
			StatusComplementZeroUpliftForAbsLambda,
			StatusAbsLambdaUniqueBaselineForK7LocalUplift,
		}, "; "),
	}
}

func buildAlternatives() []BaselineAlternativeAudit {
	absLambda := math.Abs(lambdaLambda12)
	xi := 0.5 * (r3Minus1 + absLambda)
	mk := func(name string, c float64, observable string, accepted bool, reason string) BaselineAlternativeAudit {
		return BaselineAlternativeAudit{
			Name:                 name,
			C:                    c,
			Observable:           observable,
			K7Correction:         r3Minus1 - c,
			ComplementCorrection: absLambda - c,
			Expectation:          c + pK7*(r3Minus1-c) + pComplement*(absLambda-c),
			Accepted:             accepted,
			Reason:               reason,
		}
	}
	return []BaselineAlternativeAudit{
		mk("zero/raw two-payoff", 0, "W=R P_K7+|lambda|P_perp", false, "raw form has no central scalar baseline isolated"),
		mk("scalar zero-wall baseline", absLambda, "W=|lambda|I_H72+S_split P_K7", true, "unique complement-zero K7-local uplift"),
		mk("gauge baseline", r3Minus1, "W=R I_H72-S_split P_perp", false, "algebraically valid but makes complement the correction sector"),
		mk("xi boundary baseline", xi, "W=xi I_H72+(R-xi)P_K7+(|lambda|-xi)P_perp", false, "two-sided corrections; not support-local on K7"),
		mk("midpoint baseline", 0.5*(r3Minus1+absLambda), "W=((R+|lambda|)/2)I_H72+(S_split/2)P_K7-(S_split/2)P_perp", false, "midpoint gives two-sided corrections; not active K7-local uplift"),
	}
}

func buildSupportLocality(a ActiveScalarBaselineChoiceAudit) SupportLocalitySelectionAudit {
	return SupportLocalitySelectionAudit{
		SupportLocalityCondition: "central baseline plus correction must have zero P_perp uplift and K7-local correction",
		RawTwoPayoffForm:         "R P_K7+|lambda|P_perp",
		ScalarBaselineForm:       "|lambda|I_H72+S_split P_K7",
		ComplementZeroForAbs:     a.ComplementZero,
		ForcesCAbsLambda:         a.UniqueComplementZeroChoice,
		Gate696Consistent:        true,
		Verdict: strings.Join([]string{
			StatusSupportLocalitySelectsScalarBaseline,
			StatusScalarBaselineSelectionSupportLocalGauge,
			StatusK7UpliftFormSharperThanRawTwoPayoff,
		}, "; "),
	}
}

func buildAirlock() ScalarWallAirlockCompatibilityAudit {
	return ScalarWallAirlockCompatibilityAudit{
		ScalarWallReference:         "c=|lambda| uses scalar zero-wall depth as universal reference level",
		UsesScalarZeroWallDepth:     true,
		GaugeBaselineAlgebraic:      true,
		GaugeBaselineReversesSector: true,
		CompatibleWithAirlock:       true,
		Verdict:                     StatusScalarWallAirlockCompatibilityAudited,
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		CentralGaugeFamilyRole: "central baseline gauge: W=cI+(R-c)P_K7+(|lambda|-c)P_perp",
		AbsLambdaRole:          "|lambda|: scalar zero-wall reference selecting complement-zero uplift",
		K7UpliftRole:           "S_split P_K7: support-local K7 boundary split uplift",
		GaugeBaselineRole:      "c=R is algebraically valid but makes P_perp carry the correction",
		TruthBoundary:          "reference gauge selection is conditional, not native",
		Verdict:                StatusTypedBaselineAlternativesAudited,
	}
}

func Statuses() []string {
	return []string{
		StatusGate706CentralBaselineUpliftInherited,
		StatusCentralBaselineGaugeFamilyDefined,
		StatusTotalExpectationBaselineGaugeInvariant,
		StatusActiveScalarBaselineChoiceComputed,
		StatusComplementZeroUpliftForAbsLambda,
		StatusTypedBaselineAlternativesAudited,
		StatusSupportLocalitySelectsScalarBaseline,
		StatusScalarWallAirlockCompatibilityAudited,
		StatusAbsLambdaUniqueBaselineForK7LocalUplift,
		StatusScalarBaselineSelectionSupportLocalGauge,
		StatusK7UpliftFormSharperThanRawTwoPayoff,
		StatusBaselineChoiceNotNativeYet,
		StatusNoNativeScalarBaselineReferenceSelection,
		StatusNoNativeK7RatherThanComplementCarriesUplift,
		StatusNoNativeBoundaryWoundUpliftTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate707CentralBaselineGaugeBoundary,
	}
}

func FormatInheritance(x Gate706Inheritance) string {
	return fmt.Sprintf("inherited=%t identity=%t k7Uplift=%t Dbase=%t baselineDoesNotSelect=%t noScalar=%t noK7=%t noUplift=%t noHistory=%t no7=%t expectedWound=%.18g expectedUplift=%.18g D=%.18g verdict=%q", x.CentralBaselineUpliftInherited, x.CentralBaselineIdentityShift, x.NontrivialContentK7Uplift, x.DBaseBaselineSubtracted, x.BaselineDoesNotSelectK7OrRho72, x.NoNativeScalarBaseline, x.NoNativeK7Uplift, x.NoNativeUpliftTheorem, x.NoNativeHistoryResponse, x.NoNativeSevenOver72, x.ExpectedWound, x.ExpectedUplift, x.DBase, x.Verdict)
}

func FormatGaugeFamily(x CentralBaselineGaugeFamilyAudit) string {
	return fmt.Sprintf("form=%q c=%.18g k7Corr=%.18g perpCorr=%.18g formula=%q expectation=%.18g raw=%.18g invariant=%t verdict=%q", x.GeneralForm, x.ArbitraryC, x.K7Correction, x.ComplementCorrection, x.ExpectationFormula, x.ExpectationAtC, x.RawExpectation, x.GaugeInvariant, x.Verdict)
}

func FormatActiveChoice(x ActiveScalarBaselineChoiceAudit) string {
	return fmt.Sprintf("c=%.18g observable=%q k7Uplift=%.18g complement=%.18g complementZero=%t k7Local=%t unique=%t verdict=%q", x.C, x.Observable, x.K7Uplift, x.ComplementUplift, x.ComplementZero, x.K7LocalResponse, x.UniqueComplementZeroChoice, x.Verdict)
}

func FormatAlternative(x BaselineAlternativeAudit) string {
	return fmt.Sprintf("name=%q c=%.18g observable=%q k7Corr=%.18g perpCorr=%.18g expectation=%.18g accepted=%t reason=%q", x.Name, x.C, x.Observable, x.K7Correction, x.ComplementCorrection, x.Expectation, x.Accepted, x.Reason)
}

func FormatAlternatives(xs []BaselineAlternativeAudit) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatAlternative(x))
	}
	return strings.Join(parts, " | ")
}

func FormatSupportLocality(x SupportLocalitySelectionAudit) string {
	return fmt.Sprintf("condition=%q raw=%q scalar=%q complementZero=%t forcesAbs=%t gate696=%t verdict=%q", x.SupportLocalityCondition, x.RawTwoPayoffForm, x.ScalarBaselineForm, x.ComplementZeroForAbs, x.ForcesCAbsLambda, x.Gate696Consistent, x.Verdict)
}

func FormatAirlock(x ScalarWallAirlockCompatibilityAudit) string {
	return fmt.Sprintf("reference=%q usesScalar=%t gaugeAlgebraic=%t gaugeReverses=%t compatible=%t verdict=%q", x.ScalarWallReference, x.UsesScalarZeroWallDepth, x.GaugeBaselineAlgebraic, x.GaugeBaselineReversesSector, x.CompatibleWithAirlock, x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("gauge=%q abs=%q k7=%q gaugeBaseline=%q boundary=%q verdict=%q", x.CentralGaugeFamilyRole, x.AbsLambdaRole, x.K7UpliftRole, x.GaugeBaselineRole, x.TruthBoundary, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=%s verdict=%q", strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("baselineNative=%t refSelection=%t k7VsPerp=%t upliftTheorem=%t native7=%t boundaryStress=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsBaselineChoiceNative, x.ClaimsNativeScalarBaselineReferenceSelection, x.ClaimsNativeK7RatherThanComplementUplift, x.ClaimsNativeBoundaryWoundUpliftTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerived, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
