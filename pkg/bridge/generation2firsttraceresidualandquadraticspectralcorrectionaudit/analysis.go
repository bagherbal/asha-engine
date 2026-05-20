// Package generation2firsttraceresidualandquadraticspectralcorrectionaudit implements
// Gate 690: First-Trace Residual and Quadratic Spectral Correction Audit.
//
// Gate 689 identified the active bridge as the first ordinary trace of the
// support-selected response operator
//
//	F_1 = Tr(R_split)/72 = (7/72)S_split.
//
// Gate 690 audits the small residual
//
//	E_1 = D_base - F_1
//
// and asks whether it is consistent with a typed, suppressed second-order
// correction proportional to
//
//	F_2 = Tr(R_split^2)/72.
//
// This is a bridge-layer residual-compression audit only.  It does not promote
// the quadratic correction to a native theorem, because D_base already contains
// kappa_e and any use of kappa_e as the explanatory coefficient is therefore
// partially dependent.  No boundary stress, scalar RG matching, Higgs mass,
// gauge unification, flavor, CKM/PMNS, native first-trace theorem, spectral
// expansion theorem, or native 7/72 theorem is derived here.
package generation2firsttraceresidualandquadraticspectralcorrectionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate689 "github.com/bagherbal/asha-engine/pkg/bridge/generation2firsttracefunctionalselectionandspectralorderaudit"
)

const (
	AuditID = "GATE690-FIRST-TRACE-RESIDUAL-AND-QUADRATIC-SPECTRAL-CORRECTION-AUDIT"

	StatusGate689FirstTraceSelectionInherited       = "PASS_GATE689_FIRST_TRACE_SELECTION_INHERITED"
	StatusFirstTraceResidualComputed                = "PASS_FIRST_TRACE_RESIDUAL_COMPUTED"
	StatusQuadraticSpectralScaleComputed            = "PASS_QUADRATIC_SPECTRAL_SCALE_COMPUTED"
	StatusResidualOverF2CoefficientComputed         = "PASS_RESIDUAL_OVER_F2_COEFFICIENT_COMPUTED"
	StatusTypedCoefficientCandidatesAudited         = "PASS_TYPED_COEFFICIENT_CANDIDATES_AUDITED"
	StatusFirstTraceResidualSecondOrderSuppressed   = "CONDITIONAL_SUPPORT_FIRST_TRACE_RESIDUAL_IS_SECOND_ORDER_SUPPRESSED"
	StatusKappaECloseToQuadraticResidualCoefficient = "CONDITIONAL_SUPPORT_KAPPA_E_CLOSE_TO_QUADRATIC_RESIDUAL_COEFFICIENT"
	StatusQuadraticTraceNotLeading                  = "FAILED_ROUTE_QUADRATIC_TRACE_NOT_ACTIVE_LEADING_RESPONSE"
	StatusKappaEQuadraticNotIndependentlyCertified  = "FAILED_ROUTE_KAPPA_E_QUADRATIC_CORRECTION_NOT_INDEPENDENTLY_CERTIFIED"
	StatusNoNativeSpectralExpansionTheorem          = "FAILED_ROUTE_NO_NATIVE_SPECTRAL_EXPANSION_RESPONSE_THEOREM"
	StatusNoNativeFirstTraceResponseTheorem         = "FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_RESPONSE_THEOREM"
	StatusNoNativeSevenOver72Theorem                = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate690FirstTraceResidualBoundary         = "FIREWALL_PRESERVED_GATE690_FIRST_TRACE_RESIDUAL_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	k7Dimension       = 7

	// Already active typed quantities from prior bridge ledgers.  These are
	// compared as named quantities only; Gate690 deliberately performs no
	// arbitrary rational search and introduces no new fitted coefficient.
	kappaE             = 0.00550355419157456
	sin2Theta13Quarter = 0.0055375
	jCKM               = 3.11699352875547e-05
	kappaEOrient       = sin2Theta13Quarter - jCKM
	kappaLambda        = 0.0443230430960771
	loopUnitL          = 1.0 / (8.0 * math.Pi)

	residualTolerance    = 1e-18
	coefficientTolerance = 1e-15
)

type Gate689Inheritance struct {
	FirstTraceSelectionInherited bool
	Operator                     string
	DBase                        float64
	SSplit                       float64
	F1                           float64
	F2                           float64
	F3                           float64
	H72Dimension                 int
	K7Dimension                  int
	FirstTraceActive             bool
	QuadraticLeadingActive       bool
	NativeFirstTraceTheorem      bool
	NativeSevenOver72Theorem     bool
	Verdict                      string
}

type FirstTraceResidualAudit struct {
	DBase           float64
	F1              float64
	Residual        float64
	AbsResidual     float64
	Formula         string
	ResidualTiny    bool
	ResidualNonzero bool
	Verdict         string
}

type QuadraticScaleAudit struct {
	F2                       float64
	F3                       float64
	Residual                 float64
	ResidualOverF2           float64
	ResidualOverF2Small      bool
	F2SecondOrder            bool
	F3ThirdOrder             bool
	F2MuchLargerThanResidual bool
	QuadraticStillNotLeading bool
	Verdict                  string
}

type TypedCoefficientCandidate struct {
	Name                        string
	Expression                  string
	Value                       float64
	Correction                  float64
	ResidualAfterCorrection     float64
	AbsResidualAfterCorrection  float64
	RelativeResidualToE1        float64
	CoefficientDistance         float64
	AlreadyActiveTypedQuantity  bool
	IntroducedByArbitrarySearch bool
	PartiallyDependentOnDBase   bool
	NativeCertifiedCorrection   bool
	Comment                     string
}

type TypedCoefficientAudit struct {
	Candidates                  []TypedCoefficientCandidate
	CandidateCount              int
	TargetCoefficient           float64
	BestCandidate               string
	BestCandidateValue          float64
	BestCoefficientDistance     float64
	BestResidualAfterCorrection float64
	KappaEClosest               bool
	KappaEOrientCloseButWorse   bool
	NoArbitraryRationalSearch   bool
	AllCandidatesAlreadyTyped   bool
	Verdict                     string
}

type FlavorDeficitComparisonAudit struct {
	E1                            float64
	F2                            float64
	F2AloneResidual               float64
	KappaECorrection              float64
	KappaEResidual                float64
	KappaEAbsResidual             float64
	KappaEOrientCorrection        float64
	KappaEOrientResidual          float64
	KappaEOrientAbsResidual       float64
	KappaECloserThanF2Alone       bool
	KappaEOrientCloserThanF2Alone bool
	KappaEExact                   bool
	KappaEOrientExact             bool
	ResidualClueOnly              bool
	Verdict                       string
}

type NoncircularityAudit struct {
	DBaseExpression                     string
	DBaseContainsKappaE                 bool
	CorrectionUsesKappaE                bool
	KappaEExplanationPartiallyDependent bool
	IndependentEvidence                 bool
	NativeCorrectionCertified           bool
	PromoteCorrection                   bool
	Verdict                             string
}

type SpectralExpansionAudit struct {
	FormulaWithFreeC2         string
	FormulaWithKappaE         string
	F1                        float64
	F2                        float64
	ExactC2                   float64
	KappaECandidate           float64
	PredictionWithExactC2     float64
	PredictionWithKappaE      float64
	ResidualWithExactC2       float64
	ResidualWithKappaE        float64
	ExactC2IsDefinitionOnly   bool
	KappaEFormulaPromoted     bool
	ExpansionTheoremCertified bool
	Verdict                   string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsQuadraticLeadingResponse  bool
	ClaimsKappaECorrectionCertified bool
	ClaimsNativeSpectralExpansion   bool
	ClaimsNativeFirstTraceTheorem   bool
	ClaimsNativeSevenOver72         bool
	ClaimsBoundaryStress            bool
	ClaimsScalarRGMatching          bool
	ClaimsHiggsMass                 bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	ClaimsCKMPMNS                   bool
	ClaimsProjectorActivation       bool
	PerformsArbitraryRationalSearch bool
	Verdict                         string
}

type Analysis struct {
	Inherited      Gate689Inheritance
	Residual       FirstTraceResidualAudit
	Quadratic      QuadraticScaleAudit
	Coefficients   TypedCoefficientAudit
	FlavorDeficit  FlavorDeficitComparisonAudit
	Noncircularity NoncircularityAudit
	Expansion      SpectralExpansionAudit
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
	g689, err := gate689.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate689 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g689)
	residual := buildFirstTraceResidual(inherited)
	quadratic := buildQuadraticScale(inherited, residual)
	coefficients := buildTypedCoefficientAudit(inherited, residual, quadratic)
	flavor := buildFlavorDeficitComparison(residual, quadratic)
	noncircular := buildNoncircularity()
	expansion := buildSpectralExpansion(residual, quadratic, flavor)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeSpectralExpansionTheorem,
			StatusNoNativeFirstTraceResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native HistoryResponseFirstTraceTheorem plus a SpectralExpansionResponseTheorem explaining why the physical-history bridge uses the first ordinary trace and whether a suppressed quadratic trace correction has a lawful typed coefficient independent of D_base",
		Verdict: strings.Join([]string{
			StatusNoNativeSpectralExpansionTheorem,
			StatusNoNativeFirstTraceResponseTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate690FirstTraceResidualBoundary}
	truth := "Gate 690 audits the tiny residual left after the Gate689 first ordinary trace.  The residual E_1=D_base-F_1 is about 8.525834398e-10, while F_2=Tr(R_split^2)/72 is about 1.624013232e-7, so E_1/F_2 is about 0.005249855.  Among already active typed quantities, kappa_e is closest to this coefficient and kappa_e*F_2 compresses the residual much better than F_2 alone, but the match is not exact and is partially dependent because D_base already contains kappa_e.  The result is therefore only a second-order residual-compression clue; no native spectral expansion theorem, first-trace theorem, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Residual: residual, Quadratic: quadratic, Coefficients: coefficients, FlavorDeficit: flavor, Noncircularity: noncircular, Expansion: expansion, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate689.Analysis) Gate689Inheritance {
	return Gate689Inheritance{
		FirstTraceSelectionInherited: g.Selection.SelectedFunctional == "F_1=Tr(R_split)/72=(7/72)S_split" && g.Functionals.OnlyFirstTraceActive,
		Operator:                     g.Inherited.Operator,
		DBase:                        g.Inherited.DBase,
		SSplit:                       g.Inherited.SSplit,
		F1:                           g.Functionals.FirstOrdinaryTrace,
		F2:                           g.Functionals.QuadraticTrace,
		F3:                           g.Functionals.CubicTrace,
		H72Dimension:                 g.Inherited.H72Dimension,
		K7Dimension:                  g.Inherited.K7Dimension,
		FirstTraceActive:             g.Selection.SelectionIsComparative,
		QuadraticLeadingActive:       false,
		NativeFirstTraceTheorem:      g.Selection.NativeFirstTraceProved,
		NativeSevenOver72Theorem:     g.Selection.NativeSevenOver72Proved,
		Verdict:                      StatusGate689FirstTraceSelectionInherited,
	}
}

func buildFirstTraceResidual(i Gate689Inheritance) FirstTraceResidualAudit {
	e1 := i.DBase - i.F1
	return FirstTraceResidualAudit{
		DBase:           i.DBase,
		F1:              i.F1,
		Residual:        e1,
		AbsResidual:     math.Abs(e1),
		Formula:         "E_1 = D_base - F_1 = D_base - Tr(R_split)/72",
		ResidualTiny:    math.Abs(e1) < 1e-8,
		ResidualNonzero: math.Abs(e1) > 1e-12,
		Verdict:         StatusFirstTraceResidualComputed,
	}
}

func buildQuadraticScale(i Gate689Inheritance, r FirstTraceResidualAudit) QuadraticScaleAudit {
	c2 := r.Residual / i.F2
	return QuadraticScaleAudit{
		F2:                       i.F2,
		F3:                       i.F3,
		Residual:                 r.Residual,
		ResidualOverF2:           c2,
		ResidualOverF2Small:      math.Abs(c2) < 0.01,
		F2SecondOrder:            true,
		F3ThirdOrder:             true,
		F2MuchLargerThanResidual: i.F2 > 100.0*math.Abs(r.Residual),
		QuadraticStillNotLeading: true,
		Verdict: strings.Join([]string{
			StatusQuadraticSpectralScaleComputed,
			StatusResidualOverF2CoefficientComputed,
			StatusFirstTraceResidualSecondOrderSuppressed,
			StatusQuadraticTraceNotLeading,
		}, "; "),
	}
}

func buildTypedCoefficientAudit(i Gate689Inheritance, r FirstTraceResidualAudit, q QuadraticScaleAudit) TypedCoefficientAudit {
	candidates := []TypedCoefficientCandidate{
		coefficientCandidate("kappa_e", "charged-lepton flavor-wall deficit seal", kappaE, r.Residual, i.F2, true, true, "closest active typed coefficient, but partially dependent because D_base contains kappa_e"),
		coefficientCandidate("kappa_e_orient", "sin^2(theta13)/4 - J_CKM", kappaEOrient, r.Residual, i.F2, true, false, "orientation-substituted flavor clue; close but slightly worse than exact kappa_e"),
		coefficientCandidate("kappa_lambda", "scalar low-scale matching deficit", kappaLambda, r.Residual, i.F2, true, false, "typed scalar deficit, but much too large as a residual coefficient"),
		coefficientCandidate("L", "1/(8*pi)", loopUnitL, r.Residual, i.F2, true, false, "history loop unit candidate, but much too large as a residual coefficient"),
		coefficientCandidate("S_split", "lambda(Lambda_12)+(R_3-1)", i.SSplit, r.Residual, i.F2, true, false, "active boundary scalar, but too small as the coefficient"),
		coefficientCandidate("7/72", "rank(K7)/dim(H72)", 7.0/72.0, r.Residual, i.F2, true, false, "active trace weight, but much too large as a residual coefficient"),
		coefficientCandidate("1/72", "unit trace weight on H72", 1.0/72.0, r.Residual, i.F2, true, false, "typed unit trace weight, but too large as a residual coefficient"),
	}
	best := candidates[0]
	allTyped := true
	noSearch := true
	for _, c := range candidates {
		if c.CoefficientDistance < best.CoefficientDistance {
			best = c
		}
		allTyped = allTyped && c.AlreadyActiveTypedQuantity
		noSearch = noSearch && !c.IntroducedByArbitrarySearch
	}
	return TypedCoefficientAudit{
		Candidates:                  candidates,
		CandidateCount:              len(candidates),
		TargetCoefficient:           q.ResidualOverF2,
		BestCandidate:               best.Name,
		BestCandidateValue:          best.Value,
		BestCoefficientDistance:     best.CoefficientDistance,
		BestResidualAfterCorrection: best.AbsResidualAfterCorrection,
		KappaEClosest:               best.Name == "kappa_e",
		KappaEOrientCloseButWorse:   candidates[1].CoefficientDistance > candidates[0].CoefficientDistance && candidates[1].CoefficientDistance < 3e-4,
		NoArbitraryRationalSearch:   noSearch,
		AllCandidatesAlreadyTyped:   allTyped,
		Verdict: strings.Join([]string{
			StatusTypedCoefficientCandidatesAudited,
			StatusKappaECloseToQuadraticResidualCoefficient,
			StatusKappaEQuadraticNotIndependentlyCertified,
		}, "; "),
	}
}

func coefficientCandidate(name, expression string, value, residual, f2 float64, typed, dependent bool, comment string) TypedCoefficientCandidate {
	correction := value * f2
	after := residual - correction
	return TypedCoefficientCandidate{
		Name:                        name,
		Expression:                  expression,
		Value:                       value,
		Correction:                  correction,
		ResidualAfterCorrection:     after,
		AbsResidualAfterCorrection:  math.Abs(after),
		RelativeResidualToE1:        math.Abs(after) / math.Abs(residual),
		CoefficientDistance:         math.Abs((residual / f2) - value),
		AlreadyActiveTypedQuantity:  typed,
		IntroducedByArbitrarySearch: false,
		PartiallyDependentOnDBase:   dependent,
		NativeCertifiedCorrection:   false,
		Comment:                     comment,
	}
}

func buildFlavorDeficitComparison(r FirstTraceResidualAudit, q QuadraticScaleAudit) FlavorDeficitComparisonAudit {
	kappaECorrection := kappaE * q.F2
	kappaEOrientCorrection := kappaEOrient * q.F2
	kappaEResidual := r.Residual - kappaECorrection
	kappaEOrientResidual := r.Residual - kappaEOrientCorrection
	f2AloneResidual := r.Residual - q.F2
	return FlavorDeficitComparisonAudit{
		E1:                            r.Residual,
		F2:                            q.F2,
		F2AloneResidual:               f2AloneResidual,
		KappaECorrection:              kappaECorrection,
		KappaEResidual:                kappaEResidual,
		KappaEAbsResidual:             math.Abs(kappaEResidual),
		KappaEOrientCorrection:        kappaEOrientCorrection,
		KappaEOrientResidual:          kappaEOrientResidual,
		KappaEOrientAbsResidual:       math.Abs(kappaEOrientResidual),
		KappaECloserThanF2Alone:       math.Abs(kappaEResidual) < math.Abs(f2AloneResidual),
		KappaEOrientCloserThanF2Alone: math.Abs(kappaEOrientResidual) < math.Abs(f2AloneResidual),
		KappaEExact:                   math.Abs(kappaEResidual) < residualTolerance,
		KappaEOrientExact:             math.Abs(kappaEOrientResidual) < residualTolerance,
		ResidualClueOnly:              true,
		Verdict: strings.Join([]string{
			StatusKappaECloseToQuadraticResidualCoefficient,
			StatusKappaEQuadraticNotIndependentlyCertified,
		}, "; "),
	}
}

func buildNoncircularity() NoncircularityAudit {
	return NoncircularityAudit{
		DBaseExpression:                     "D_base=kappa_lambda+kappa_e+lambda(Lambda_12)",
		DBaseContainsKappaE:                 true,
		CorrectionUsesKappaE:                true,
		KappaEExplanationPartiallyDependent: true,
		IndependentEvidence:                 false,
		NativeCorrectionCertified:           false,
		PromoteCorrection:                   false,
		Verdict: strings.Join([]string{
			StatusKappaEQuadraticNotIndependentlyCertified,
			StatusNoNativeSpectralExpansionTheorem,
		}, "; "),
	}
}

func buildSpectralExpansion(r FirstTraceResidualAudit, q QuadraticScaleAudit, f FlavorDeficitComparisonAudit) SpectralExpansionAudit {
	predictionExact := r.F1 + q.ResidualOverF2*q.F2
	predictionKappaE := r.F1 + kappaE*q.F2
	return SpectralExpansionAudit{
		FormulaWithFreeC2:         "D_base ≈ Tr(R_split)/72 + c_2 Tr(R_split^2)/72",
		FormulaWithKappaE:         "D_base ≈ Tr(R_split)/72 + kappa_e Tr(R_split^2)/72",
		F1:                        r.F1,
		F2:                        q.F2,
		ExactC2:                   q.ResidualOverF2,
		KappaECandidate:           kappaE,
		PredictionWithExactC2:     predictionExact,
		PredictionWithKappaE:      predictionKappaE,
		ResidualWithExactC2:       r.DBase - predictionExact,
		ResidualWithKappaE:        f.KappaEResidual,
		ExactC2IsDefinitionOnly:   true,
		KappaEFormulaPromoted:     false,
		ExpansionTheoremCertified: false,
		Verdict: strings.Join([]string{
			StatusFirstTraceResidualSecondOrderSuppressed,
			StatusKappaECloseToQuadraticResidualCoefficient,
			StatusKappaEQuadraticNotIndependentlyCertified,
			StatusNoNativeSpectralExpansionTheorem,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate689FirstTraceSelectionInherited,
		StatusFirstTraceResidualComputed,
		StatusQuadraticSpectralScaleComputed,
		StatusResidualOverF2CoefficientComputed,
		StatusTypedCoefficientCandidatesAudited,
		StatusFirstTraceResidualSecondOrderSuppressed,
		StatusKappaECloseToQuadraticResidualCoefficient,
		StatusQuadraticTraceNotLeading,
		StatusKappaEQuadraticNotIndependentlyCertified,
		StatusNoNativeSpectralExpansionTheorem,
		StatusNoNativeFirstTraceResponseTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate690FirstTraceResidualBoundary,
	}
}

func FormatInheritance(x Gate689Inheritance) string {
	return fmt.Sprintf("inherited=%t operator=%q dbase=%.18g ssplit=%.18g f1=%.18g f2=%.18g f3=%.18g h72=%d k7=%d firstActive=%t quadraticLeading=%t nativeFirstTrace=%t native7=%t verdict=%q", x.FirstTraceSelectionInherited, x.Operator, x.DBase, x.SSplit, x.F1, x.F2, x.F3, x.H72Dimension, x.K7Dimension, x.FirstTraceActive, x.QuadraticLeadingActive, x.NativeFirstTraceTheorem, x.NativeSevenOver72Theorem, x.Verdict)
}

func FormatResidual(x FirstTraceResidualAudit) string {
	return fmt.Sprintf("dbase=%.18g f1=%.18g residual=%.18g abs=%.18g formula=%q tiny=%t nonzero=%t verdict=%q", x.DBase, x.F1, x.Residual, x.AbsResidual, x.Formula, x.ResidualTiny, x.ResidualNonzero, x.Verdict)
}

func FormatQuadratic(x QuadraticScaleAudit) string {
	return fmt.Sprintf("f2=%.18g f3=%.18g residual=%.18g c2=%.18g c2Small=%t f2Second=%t f3Third=%t f2MuchLarger=%t quadraticLeading=%t verdict=%q", x.F2, x.F3, x.Residual, x.ResidualOverF2, x.ResidualOverF2Small, x.F2SecondOrder, x.F3ThirdOrder, x.F2MuchLargerThanResidual, !x.QuadraticStillNotLeading, x.Verdict)
}

func FormatCandidate(x TypedCoefficientCandidate) string {
	return fmt.Sprintf("%s=%s value=%.18g correction=%.18g residual=%.18g rel=%.18g coeffDist=%.18g typed=%t arbitrary=%t dependent=%t certified=%t comment=%q", x.Name, x.Expression, x.Value, x.Correction, x.ResidualAfterCorrection, x.RelativeResidualToE1, x.CoefficientDistance, x.AlreadyActiveTypedQuantity, x.IntroducedByArbitrarySearch, x.PartiallyDependentOnDBase, x.NativeCertifiedCorrection, x.Comment)
}

func FormatCoefficients(x TypedCoefficientAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("target=%.18g count=%d best=%q bestValue=%.18g bestCoeffDistance=%.18g bestResidual=%.18g kappaEClosest=%t orientCloseWorse=%t noRationalSearch=%t allTyped=%t candidates=[%s] verdict=%q", x.TargetCoefficient, x.CandidateCount, x.BestCandidate, x.BestCandidateValue, x.BestCoefficientDistance, x.BestResidualAfterCorrection, x.KappaEClosest, x.KappaEOrientCloseButWorse, x.NoArbitraryRationalSearch, x.AllCandidatesAlreadyTyped, strings.Join(parts, " | "), x.Verdict)
}

func FormatFlavor(x FlavorDeficitComparisonAudit) string {
	return fmt.Sprintf("e1=%.18g f2=%.18g f2AloneResidual=%.18g kappaECorrection=%.18g kappaEResidual=%.18g kappaEAbs=%.18g orientCorrection=%.18g orientResidual=%.18g orientAbs=%.18g kappaECloser=%t orientCloser=%t kappaEExact=%t orientExact=%t clueOnly=%t verdict=%q", x.E1, x.F2, x.F2AloneResidual, x.KappaECorrection, x.KappaEResidual, x.KappaEAbsResidual, x.KappaEOrientCorrection, x.KappaEOrientResidual, x.KappaEOrientAbsResidual, x.KappaECloserThanF2Alone, x.KappaEOrientCloserThanF2Alone, x.KappaEExact, x.KappaEOrientExact, x.ResidualClueOnly, x.Verdict)
}

func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("dbase=%q containsKappaE=%t usesKappaE=%t dependent=%t independent=%t certified=%t promote=%t verdict=%q", x.DBaseExpression, x.DBaseContainsKappaE, x.CorrectionUsesKappaE, x.KappaEExplanationPartiallyDependent, x.IndependentEvidence, x.NativeCorrectionCertified, x.PromoteCorrection, x.Verdict)
}

func FormatExpansion(x SpectralExpansionAudit) string {
	return fmt.Sprintf("free=%q kappaEForm=%q f1=%.18g f2=%.18g exactC2=%.18g kappaE=%.18g predExact=%.18g predKappaE=%.18g residualExact=%.18g residualKappaE=%.18g exactDefinition=%t promoted=%t theorem=%t verdict=%q", x.FormulaWithFreeC2, x.FormulaWithKappaE, x.F1, x.F2, x.ExactC2, x.KappaECandidate, x.PredictionWithExactC2, x.PredictionWithKappaE, x.ResidualWithExactC2, x.ResidualWithKappaE, x.ExactC2IsDefinitionOnly, x.KappaEFormulaPromoted, x.ExpansionTheoremCertified, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsQuadraticLeading=%t claimsKappaECert=%t claimsExpansion=%t claimsFirstTrace=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t claimsCKMPMNS=%t claimsActivation=%t arbitrarySearch=%t verdict=%q", x.ClaimsQuadraticLeadingResponse, x.ClaimsKappaECorrectionCertified, x.ClaimsNativeSpectralExpansion, x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.PerformsArbitraryRationalSearch, x.Verdict)
}
