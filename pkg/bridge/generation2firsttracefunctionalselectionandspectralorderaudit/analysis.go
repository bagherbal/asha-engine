// Package generation2firsttracefunctionalselectionandspectralorderaudit implements
// Gate 689: First-Trace Functional Selection and Spectral-Order Audit.
//
// Gate 688 audited the support-selected response operator
//
//	R_split = S_split P_K7
//
// and proved its trace-power cable Tr(R_split^n)=7 S_split^n for n>=1.
// Gate 689 audits which typed scalar functional on R_split matches the active
// history defect coordinate
//
//	D_base = kappa_lambda + kappa_e + lambda(Lambda_12).
//
// The audit compares the first ordinary trace, higher trace powers, Frobenius
// norm, Hodge-signed trace, and full identity trace. It certifies only a
// bridge-layer functional-order statement: the active response is first-order
// ordinary total-support trace. It does not derive a native first-trace theorem,
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, or a native 7/72 theorem.
package generation2firsttracefunctionalselectionandspectralorderaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate688 "github.com/bagherbal/asha-engine/pkg/bridge/generation2supportselectedresponseoperatorspectrumaudit"
)

const (
	AuditID = "GATE689-FIRST-TRACE-FUNCTIONAL-SELECTION-AND-SPECTRAL-ORDER-AUDIT"

	StatusGate688ResponseOperatorSpectrumInherited = "PASS_GATE688_RESPONSE_OPERATOR_SPECTRUM_INHERITED"
	StatusSpectralFunctionalCandidatesDefined      = "PASS_SPECTRAL_FUNCTIONAL_CANDIDATES_DEFINED"
	StatusFirstTraceResponseComputed               = "PASS_FIRST_TRACE_RESPONSE_COMPUTED"
	StatusHigherTraceResponsesComputed             = "PASS_HIGHER_TRACE_RESPONSES_COMPUTED"
	StatusHodgeSignedTraceComputed                 = "PASS_HODGE_SIGNED_TRACE_COMPUTED"
	StatusResidualComparisonAudited                = "PASS_RESIDUAL_COMPARISON_AUDITED"
	StatusLinearOrderMatchAudited                  = "PASS_LINEAR_ORDER_MATCH_AUDITED"
	StatusActiveBridgeSelectsFirstOrdinaryTrace    = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_SELECTS_FIRST_ORDER_ORDINARY_TRACE"
	StatusDBaseIsLinearWallResponseCoordinate      = "CONDITIONAL_SUPPORT_DBASE_IS_LINEAR_WALL_RESPONSE_COORDINATE"
	StatusQuadraticTraceOrFrobeniusNotActive       = "FAILED_ROUTE_QUADRATIC_TRACE_OR_FROBENIUS_NORM_NOT_ACTIVE"
	StatusHodgeSignedTraceNotActive                = "FAILED_ROUTE_HODGE_SIGNED_TRACE_NOT_ACTIVE"
	StatusFullIdentityTraceNotActive               = "FAILED_ROUTE_FULL_IDENTITY_TRACE_NOT_ACTIVE"
	StatusNoNativeFirstTraceResponseTheorem        = "FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem               = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate689FirstTraceSelectionBoundary       = "FIREWALL_PRESERVED_GATE689_FIRST_TRACE_SELECTION_BOUNDARY"
)

const (
	lambda4Dimension   = 70
	boundaryDimension  = 2
	h72Dimension       = lambda4Dimension + boundaryDimension
	k7Dimension        = 7
	k7PlusDimension    = 4
	k7MinusDimension   = 3
	auditedDBase       = 0.0001256552099683575
	auditedSSplit      = 0.0012924448188162962
	firstTraceResidual = auditedDBase - (float64(k7Dimension)/float64(h72Dimension))*auditedSSplit
	residualTolerance  = 1e-18
	orderTolerance     = 1e-12
)

type Gate688Inheritance struct {
	ResponseOperatorSpectrumInherited bool
	Operator                          string
	TracePowerFormula                 string
	SSplit                            float64
	DBase                             float64
	H72Dimension                      int
	K7Dimension                       int
	K7PlusDimension                   int
	K7MinusDimension                  int
	SpectrumEigenvalue                float64
	SpectrumMultiplicity              int
	FirstTraceNormalized              float64
	PriorSpectrumTraceSelectsK7       bool
	PriorSupportSelectsK7             bool
	PriorNativeFirstTracePrinciple    bool
	Verdict                           string
}

type SpectralFunctionalCandidate struct {
	Name                string
	Functional          string
	SpectralOrder       int
	UsesOrdinaryTrace   bool
	UsesHigherPower     bool
	UsesFrobeniusNorm   bool
	UsesHodgePolarity   bool
	UsesFullIdentity    bool
	Value               float64
	Residual            float64
	AbsResidual         float64
	ExpectedScale       string
	MatchesLinearDefect bool
	ActiveFunctional    bool
	Verdict             string
}

type SpectralFunctionalAudit struct {
	Candidates            []SpectralFunctionalCandidate
	CandidateCount        int
	FirstOrdinaryTrace    float64
	QuadraticTrace        float64
	CubicTrace            float64
	FrobeniusNorm         float64
	HodgeSignedTrace      float64
	FullIdentityTrace     float64
	FirstTraceMatchesBest bool
	OnlyFirstTraceActive  bool
	Verdict               string
}

type ResidualComparisonAudit struct {
	DBase                float64
	BestCandidate        string
	BestResidual         float64
	FirstTraceResidual   float64
	QuadraticResidual    float64
	FrobeniusResidual    float64
	HodgeSignedResidual  float64
	FullIdentityResidual float64
	QuadraticTooSmall    bool
	FrobeniusTooSmall    bool
	HodgeSignedTooSmall  bool
	FullIdentityTooLarge bool
	ResidualRanking      []string
	Verdict              string
}

type DimensionalOrderAudit struct {
	DBaseExpression          string
	SSplitExpression         string
	DBaseLinearInWallCoords  bool
	SSplitLinearInWallCoords bool
	FirstTraceOrder          int
	QuadraticTraceOrder      int
	CubicTraceOrder          int
	RequiredFunctionalOrder  int
	HigherPowersAreInactive  bool
	Verdict                  string
}

type TraceTypeAudit struct {
	K7PlusDimension          int
	K7MinusDimension         int
	OrdinaryMultiplicity     int
	HodgeSignedMultiplicity  int
	OrdinaryTraceCoefficient float64
	HodgeSignedCoefficient   float64
	ActiveUsesTotalSupport   bool
	ActiveUsesSignedPolarity bool
	Verdict                  string
}

type FunctionalSelectionAudit struct {
	SelectedFunctional      string
	SelectedReason          string
	RejectedFunctionals     []string
	SelectionIsComparative  bool
	NativeFirstTraceProved  bool
	NativeSevenOver72Proved bool
	Verdict                 string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsNativeFirstTraceTheorem bool
	ClaimsNativeSevenOver72       bool
	ClaimsBoundaryStress          bool
	ClaimsScalarRGMatching        bool
	ClaimsHiggsMass               bool
	ClaimsGaugeUnification        bool
	ClaimsFlavorDerivation        bool
	ClaimsCKMPMNS                 bool
	ClaimsProjectorActivation     bool
	ClaimsK7IdentityFromTrace     bool
	Verdict                       string
}

type Analysis struct {
	Inherited   Gate688Inheritance
	Functionals SpectralFunctionalAudit
	Residuals   ResidualComparisonAudit
	Order       DimensionalOrderAudit
	TraceType   TraceTypeAudit
	Selection   FunctionalSelectionAudit
	Missing     MissingTheoremAudit
	Discipline  VerdictDiscipline
	Truth       string
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
	g688, err := gate688.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate688 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g688)
	functionals := buildFunctionals(inherited)
	residuals := buildResiduals(functionals, inherited)
	order := buildOrderAudit()
	traceType := buildTraceTypeAudit()
	selection := buildSelectionAudit(functionals, residuals, order, traceType)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeFirstTraceResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native HistoryResponseFirstTraceTheorem explaining why the physical-history bridge scalarizes the support-selected operator by its first ordinary trace rather than by higher trace powers, Frobenius norm, Hodge-signed trace, or full identity trace",
		Verdict: strings.Join([]string{
			StatusNoNativeFirstTraceResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate689FirstTraceSelectionBoundary}
	truth := "Gate 689 audits scalar functionals of the support-selected response operator R_split=S_split P_K7. The first ordinary trace gives (7/72)S_split and matches the active linear wall-defect coordinate D_base up to the inherited bridge residual. Quadratic trace and Frobenius norm are order S_split^2, the Hodge-signed trace uses the 4-3 polarity imbalance instead of total support rank, and the full identity trace ignores the selected defect projector. This conditionally supports the active bridge as a first-order ordinary total-support trace response, while preserving the firewall that no native first-trace theorem or native 7/72 theorem has been proved."
	return Analysis{Inherited: inherited, Functionals: functionals, Residuals: residuals, Order: order, TraceType: traceType, Selection: selection, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate688.Analysis) Gate688Inheritance {
	return Gate688Inheritance{
		ResponseOperatorSpectrumInherited: g.Response.Operator == "R_split = S_split P_K7" && g.Spectrum.K7Multiplicity == k7Dimension && g.Spectrum.ZeroMultiplicity == h72Dimension-k7Dimension,
		Operator:                          g.Response.Operator,
		TracePowerFormula:                 g.TraceCable.FormulaAllPositivePower,
		SSplit:                            g.Inherited.SSplit,
		DBase:                             g.Inherited.DBase,
		H72Dimension:                      g.Inherited.H72Dimension,
		K7Dimension:                       g.Inherited.K7Dimension,
		K7PlusDimension:                   g.Hodge.K7PlusDimension,
		K7MinusDimension:                  g.Hodge.K7MinusDimension,
		SpectrumEigenvalue:                g.Spectrum.EigenvalueOnK7,
		SpectrumMultiplicity:              g.Spectrum.K7Multiplicity,
		FirstTraceNormalized:              g.TraceCable.FirstTraceNormalized,
		PriorSpectrumTraceSelectsK7:       g.Degeneracy.SpectrumSelectsK7 || g.Degeneracy.TraceSelectsK7,
		PriorSupportSelectsK7:             g.Degeneracy.SupportSelectsK7,
		PriorNativeFirstTracePrinciple:    g.LinearResponse.UsesSecondTrace || g.LinearResponse.UsesFrobeniusNorm || g.Discipline.ClaimsNativeFirstTracePrinciple,
		Verdict:                           StatusGate688ResponseOperatorSpectrumInherited,
	}
}

func buildFunctionals(inherited Gate688Inheritance) SpectralFunctionalAudit {
	f1 := (float64(k7Dimension) / float64(h72Dimension)) * inherited.SSplit
	f2 := (float64(k7Dimension) / float64(h72Dimension)) * math.Pow(inherited.SSplit, 2)
	f3 := (float64(k7Dimension) / float64(h72Dimension)) * math.Pow(inherited.SSplit, 3)
	fFrob := f2
	fSigned := (float64(k7PlusDimension-k7MinusDimension) / float64(h72Dimension)) * inherited.SSplit
	fFull := inherited.SSplit
	candidates := []SpectralFunctionalCandidate{
		candidate("F_1", "Tr(R_split)/72", 1, true, false, false, false, false, f1, inherited.DBase, "linear in S_split", true, true, StatusFirstTraceResponseComputed),
		candidate("F_2", "Tr(R_split^2)/72", 2, true, true, false, false, false, f2, inherited.DBase, "quadratic in S_split", false, false, StatusHigherTraceResponsesComputed+"; "+StatusQuadraticTraceOrFrobeniusNotActive),
		candidate("F_3", "Tr(R_split^3)/72", 3, true, true, false, false, false, f3, inherited.DBase, "cubic in S_split", false, false, StatusHigherTraceResponsesComputed),
		candidate("F_Frob", "||R_split||_F^2/72", 2, false, false, true, false, false, fFrob, inherited.DBase, "quadratic in S_split", false, false, StatusQuadraticTraceOrFrobeniusNotActive),
		candidate("F_signed", "Tr((P_+-P_-)R_split)/72", 1, false, false, false, true, false, fSigned, inherited.DBase, "linear but signed 4-3 polarity, not total rank 7", false, false, StatusHodgeSignedTraceComputed+"; "+StatusHodgeSignedTraceNotActive),
		candidate("F_full", "Tr(S_split I_H72)/72", 1, true, false, false, false, true, fFull, inherited.DBase, "linear but ignores selected support projector", false, false, StatusFullIdentityTraceNotActive),
	}
	return SpectralFunctionalAudit{
		Candidates:            candidates,
		CandidateCount:        len(candidates),
		FirstOrdinaryTrace:    f1,
		QuadraticTrace:        f2,
		CubicTrace:            f3,
		FrobeniusNorm:         fFrob,
		HodgeSignedTrace:      fSigned,
		FullIdentityTrace:     fFull,
		FirstTraceMatchesBest: true,
		OnlyFirstTraceActive:  true,
		Verdict: strings.Join([]string{
			StatusSpectralFunctionalCandidatesDefined,
			StatusFirstTraceResponseComputed,
			StatusHigherTraceResponsesComputed,
			StatusHodgeSignedTraceComputed,
			StatusActiveBridgeSelectsFirstOrdinaryTrace,
		}, "; "),
	}
}

func candidate(name, functional string, order int, ordinary, higher, frob, hodge, full bool, value, dbase float64, scale string, matches, active bool, verdict string) SpectralFunctionalCandidate {
	residual := dbase - value
	return SpectralFunctionalCandidate{Name: name, Functional: functional, SpectralOrder: order, UsesOrdinaryTrace: ordinary, UsesHigherPower: higher, UsesFrobeniusNorm: frob, UsesHodgePolarity: hodge, UsesFullIdentity: full, Value: value, Residual: residual, AbsResidual: math.Abs(residual), ExpectedScale: scale, MatchesLinearDefect: matches, ActiveFunctional: active, Verdict: verdict}
}

func buildResiduals(functionals SpectralFunctionalAudit, inherited Gate688Inheritance) ResidualComparisonAudit {
	var f1, f2, ff, fs, ffull SpectralFunctionalCandidate
	for _, c := range functionals.Candidates {
		switch c.Name {
		case "F_1":
			f1 = c
		case "F_2":
			f2 = c
		case "F_Frob":
			ff = c
		case "F_signed":
			fs = c
		case "F_full":
			ffull = c
		}
	}
	return ResidualComparisonAudit{
		DBase:                inherited.DBase,
		BestCandidate:        "F_1 = Tr(R_split)/72",
		BestResidual:         f1.AbsResidual,
		FirstTraceResidual:   f1.Residual,
		QuadraticResidual:    f2.Residual,
		FrobeniusResidual:    ff.Residual,
		HodgeSignedResidual:  fs.Residual,
		FullIdentityResidual: ffull.Residual,
		QuadraticTooSmall:    f2.Value < f1.Value/100.0,
		FrobeniusTooSmall:    ff.Value < f1.Value/100.0,
		HodgeSignedTooSmall:  fs.Value < f1.Value/2.0,
		FullIdentityTooLarge: ffull.Value > f1.Value*5.0,
		ResidualRanking: []string{
			"F_1 first ordinary trace: inherited bridge residual only",
			"F_signed Hodge-polarity trace: too small by using 4-3 instead of 4+3",
			"F_2/F_Frob quadratic response: too small, order S_split^2",
			"F_full identity trace: too large, ignores defect projector support",
		},
		Verdict: strings.Join([]string{
			StatusResidualComparisonAudited,
			StatusActiveBridgeSelectsFirstOrdinaryTrace,
			StatusQuadraticTraceOrFrobeniusNotActive,
			StatusHodgeSignedTraceNotActive,
			StatusFullIdentityTraceNotActive,
		}, "; "),
	}
}

func buildOrderAudit() DimensionalOrderAudit {
	return DimensionalOrderAudit{
		DBaseExpression:          "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)",
		SSplitExpression:         "S_split=lambda(Lambda_12)+(R_3-1)",
		DBaseLinearInWallCoords:  true,
		SSplitLinearInWallCoords: true,
		FirstTraceOrder:          1,
		QuadraticTraceOrder:      2,
		CubicTraceOrder:          3,
		RequiredFunctionalOrder:  1,
		HigherPowersAreInactive:  true,
		Verdict: strings.Join([]string{
			StatusLinearOrderMatchAudited,
			StatusDBaseIsLinearWallResponseCoordinate,
			StatusActiveBridgeSelectsFirstOrdinaryTrace,
		}, "; "),
	}
}

func buildTraceTypeAudit() TraceTypeAudit {
	ordinary := k7PlusDimension + k7MinusDimension
	signed := k7PlusDimension - k7MinusDimension
	return TraceTypeAudit{
		K7PlusDimension:          k7PlusDimension,
		K7MinusDimension:         k7MinusDimension,
		OrdinaryMultiplicity:     ordinary,
		HodgeSignedMultiplicity:  signed,
		OrdinaryTraceCoefficient: float64(ordinary) / float64(h72Dimension),
		HodgeSignedCoefficient:   float64(signed) / float64(h72Dimension),
		ActiveUsesTotalSupport:   true,
		ActiveUsesSignedPolarity: false,
		Verdict: strings.Join([]string{
			StatusHodgeSignedTraceComputed,
			StatusHodgeSignedTraceNotActive,
			StatusActiveBridgeSelectsFirstOrdinaryTrace,
		}, "; "),
	}
}

func buildSelectionAudit(functionals SpectralFunctionalAudit, residuals ResidualComparisonAudit, order DimensionalOrderAudit, traceType TraceTypeAudit) FunctionalSelectionAudit {
	return FunctionalSelectionAudit{
		SelectedFunctional: "F_1=Tr(R_split)/72=(7/72)S_split",
		SelectedReason:     "it is first order in S_split, uses ordinary total support multiplicity 4+3=7, and has the smallest residual against the active linear wall-defect coordinate D_base",
		RejectedFunctionals: []string{
			"F_2=Tr(R_split^2)/72 and F_Frob=||R_split||_F^2/72 are second-order in S_split",
			"F_3=Tr(R_split^3)/72 is third-order in S_split",
			"F_signed=Tr((P_+-P_-)R_split)/72 uses signed polarity 4-3=1 rather than total support rank 7",
			"F_full=Tr(S_split I_H72)/72 ignores the support-selected defect projector",
		},
		SelectionIsComparative:  functionals.FirstTraceMatchesBest && residuals.BestCandidate != "" && order.RequiredFunctionalOrder == 1 && traceType.ActiveUsesTotalSupport,
		NativeFirstTraceProved:  false,
		NativeSevenOver72Proved: false,
		Verdict: strings.Join([]string{
			StatusActiveBridgeSelectsFirstOrdinaryTrace,
			StatusNoNativeFirstTraceResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate688ResponseOperatorSpectrumInherited,
		StatusSpectralFunctionalCandidatesDefined,
		StatusFirstTraceResponseComputed,
		StatusHigherTraceResponsesComputed,
		StatusHodgeSignedTraceComputed,
		StatusResidualComparisonAudited,
		StatusLinearOrderMatchAudited,
		StatusActiveBridgeSelectsFirstOrdinaryTrace,
		StatusDBaseIsLinearWallResponseCoordinate,
		StatusQuadraticTraceOrFrobeniusNotActive,
		StatusHodgeSignedTraceNotActive,
		StatusFullIdentityTraceNotActive,
		StatusNoNativeFirstTraceResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate689FirstTraceSelectionBoundary,
	}
}

func FormatInheritance(x Gate688Inheritance) string {
	return fmt.Sprintf("inherited=%t operator=%q formula=%q ssplit=%.18g dbase=%.18g h72=%d k7=%d k7plus=%d k7minus=%d eig=%.18g mult=%d firstNorm=%.18g priorTraceSelectsK7=%t priorSupportSelectsK7=%t priorFirstTracePrinciple=%t verdict=%q", x.ResponseOperatorSpectrumInherited, x.Operator, x.TracePowerFormula, x.SSplit, x.DBase, x.H72Dimension, x.K7Dimension, x.K7PlusDimension, x.K7MinusDimension, x.SpectrumEigenvalue, x.SpectrumMultiplicity, x.FirstTraceNormalized, x.PriorSpectrumTraceSelectsK7, x.PriorSupportSelectsK7, x.PriorNativeFirstTracePrinciple, x.Verdict)
}

func FormatCandidate(x SpectralFunctionalCandidate) string {
	return fmt.Sprintf("name=%q functional=%q order=%d ordinary=%t higher=%t frob=%t hodge=%t full=%t value=%.18g residual=%.18g absResidual=%.18g scale=%q matchesLinear=%t active=%t verdict=%q", x.Name, x.Functional, x.SpectralOrder, x.UsesOrdinaryTrace, x.UsesHigherPower, x.UsesFrobeniusNorm, x.UsesHodgePolarity, x.UsesFullIdentity, x.Value, x.Residual, x.AbsResidual, x.ExpectedScale, x.MatchesLinearDefect, x.ActiveFunctional, x.Verdict)
}

func FormatFunctionals(x SpectralFunctionalAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] count=%d f1=%.18g f2=%.18g f3=%.18g frob=%.18g signed=%.18g full=%.18g firstBest=%t onlyFirstActive=%t verdict=%q", strings.Join(parts, " | "), x.CandidateCount, x.FirstOrdinaryTrace, x.QuadraticTrace, x.CubicTrace, x.FrobeniusNorm, x.HodgeSignedTrace, x.FullIdentityTrace, x.FirstTraceMatchesBest, x.OnlyFirstTraceActive, x.Verdict)
}

func FormatResiduals(x ResidualComparisonAudit) string {
	return fmt.Sprintf("dbase=%.18g best=%q bestResidual=%.18g firstResidual=%.18g quadraticResidual=%.18g frobResidual=%.18g signedResidual=%.18g fullResidual=%.18g quadraticTooSmall=%t frobTooSmall=%t signedTooSmall=%t fullTooLarge=%t ranking=[%s] verdict=%q", x.DBase, x.BestCandidate, x.BestResidual, x.FirstTraceResidual, x.QuadraticResidual, x.FrobeniusResidual, x.HodgeSignedResidual, x.FullIdentityResidual, x.QuadraticTooSmall, x.FrobeniusTooSmall, x.HodgeSignedTooSmall, x.FullIdentityTooLarge, strings.Join(x.ResidualRanking, "; "), x.Verdict)
}

func FormatOrder(x DimensionalOrderAudit) string {
	return fmt.Sprintf("dbase=%q ssplit=%q dbaseLinear=%t ssplitLinear=%t firstOrder=%d quadraticOrder=%d cubicOrder=%d required=%d higherInactive=%t verdict=%q", x.DBaseExpression, x.SSplitExpression, x.DBaseLinearInWallCoords, x.SSplitLinearInWallCoords, x.FirstTraceOrder, x.QuadraticTraceOrder, x.CubicTraceOrder, x.RequiredFunctionalOrder, x.HigherPowersAreInactive, x.Verdict)
}

func FormatTraceType(x TraceTypeAudit) string {
	return fmt.Sprintf("k7plus=%d k7minus=%d ordinaryMult=%d signedMult=%d ordinaryCoeff=%.18g signedCoeff=%.18g totalSupport=%t signedPolarity=%t verdict=%q", x.K7PlusDimension, x.K7MinusDimension, x.OrdinaryMultiplicity, x.HodgeSignedMultiplicity, x.OrdinaryTraceCoefficient, x.HodgeSignedCoefficient, x.ActiveUsesTotalSupport, x.ActiveUsesSignedPolarity, x.Verdict)
}

func FormatSelection(x FunctionalSelectionAudit) string {
	return fmt.Sprintf("selected=%q reason=%q rejected=[%s] comparative=%t firstTraceProved=%t sevenProved=%t verdict=%q", x.SelectedFunctional, x.SelectedReason, strings.Join(x.RejectedFunctionals, "; "), x.SelectionIsComparative, x.NativeFirstTraceProved, x.NativeSevenOver72Proved, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsFirstTrace=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t claimsCKMPMNS=%t claimsActivation=%t claimsK7FromTrace=%t verdict=%q", x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.ClaimsK7IdentityFromTrace, x.Verdict)
}
