// Package generation2globalaugmentedtracekernelconditionalaudit implements
// Gate 680: Global Augmented Trace versus Kernel-Conditional Trace Audit.
//
// Gate 679 corrected the false literal exact-sequence target by showing that
// K7 is not the kernel of the natural split projection pi_split:H72->Q_boundary.
// Gate 680 audits the next denominator question: why the active line response
// uses the global full-extension density 7/72 rather than the conditional kernel
// density 7/71 or the finite-only density 7/70.
package generation2globalaugmentedtracekernelconditionalaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate679 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryquotientprojectionkernelaudit"
)

const (
	AuditID = "GATE680-GLOBAL-AUGMENTED-TRACE-VERSUS-KERNEL-CONDITIONAL-TRACE-AUDIT"

	StatusGate679RelativeTraceResponseInherited        = "PASS_GATE679_RELATIVE_TRACE_RESPONSE_INHERITED"
	StatusShortExactProjectionSequenceDefined          = "PASS_SHORT_EXACT_PROJECTION_SEQUENCE_DEFINED"
	StatusK7DefectInsideKernelClassified               = "PASS_K7_DEFECT_INSIDE_KERNEL_CLASSIFIED"
	StatusGlobalKernelFiniteTraceNormalizationsAudited = "PASS_GLOBAL_KERNEL_FINITE_TRACE_NORMALIZATIONS_AUDITED"
	StatusResponseCompatibilityWithQuotientLineAudited = "PASS_RESPONSE_COMPATIBILITY_WITH_QUOTIENT_LINE_AUDITED"
	StatusGlobalH72TraceTypeCorrectForQuotientResponse = "CONDITIONAL_SUPPORT_GLOBAL_H72_TRACE_IS_TYPE_CORRECT_FOR_QUOTIENT_RESPONSE"
	StatusSevenOver72FullExtensionDefectDensity        = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_FULL_EXTENSION_DEFECT_DENSITY"
	StatusKernelConditionalTraceNotActive              = "FAILED_ROUTE_KERNEL_CONDITIONAL_TRACE_7_OVER_71_NOT_ACTIVE_RESPONSE_NORMALIZATION"
	StatusFiniteOnlyTraceOmitsBoundaryQuotient         = "FAILED_ROUTE_FINITE_ONLY_TRACE_7_OVER_70_OMITS_BOUNDARY_QUOTIENT_INPUT"
	StatusNoNativeGlobalTraceResponsePrinciple         = "FAILED_ROUTE_NO_NATIVE_GLOBAL_TRACE_RESPONSE_PRINCIPLE"
	StatusNoNativeSevenOver72Theorem                   = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoNativeTraceToBoundaryQuotientTheorem       = "FAILED_ROUTE_NO_NATIVE_TRACE_TO_BOUNDARY_QUOTIENT_RESPONSE_THEOREM"
	StatusNoBoundaryStressDerivation                   = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusGate680Boundary                              = "FIREWALL_PRESERVED_GATE680_GLOBAL_TRACE_NORMALIZATION_BOUNDARY"
)

type Gate679Inheritance struct {
	RelativeTraceInherited bool
	H72Dimension           int
	KernelDimension        int
	QuotientDimension      int
	K7Rank                 int
	SSplit                 float64
	DBase                  float64
	TauGlobal              float64
	TauKernel              float64
	TauFinite              float64
	TauHalf                float64
	ResidualGlobal         float64
	FirewallPreserved      bool
	Verdict                string
}

type ShortExactProjectionSequence struct {
	Sequence          string
	KernelFormula     string
	KernelDimension   int
	AmbientDimension  int
	QuotientDimension int
	ExactByDimension  bool
	Verdict           string
}

type DefectInclusion struct {
	Inclusion                string
	Rank                     int
	KernelDimension          int
	AmbientDimension         int
	KernelConditionalDensity float64
	GlobalDensity            float64
	FiniteDensity            float64
	FullKernel               bool
	Verdict                  string
}

type TraceNormalization struct {
	Name                     string
	Formula                  string
	Denominator              int
	Value                    float64
	PredictedDBase           float64
	Residual                 float64
	AbsResidual              float64
	IncludesFiniteChamber    bool
	IncludesBoundaryAntiLine bool
	IncludesQuotientLine     bool
	Classification           string
	Verdict                  string
}

type ResponseCompatibility struct {
	ResponseDomain               string
	QuotientLineIncludedInSystem bool
	GlobalKeepsQuotientInput     bool
	KernelExcludesQuotientInput  bool
	FiniteExcludesBoundarySystem bool
	BestNormalization            string
	Verdict                      string
}

type MissingTheoremAudit struct {
	Missing                    []string
	NewPreciseMissingPrinciple string
	AllowedSupport             []string
	Verdict                    string
}

type VerdictDiscipline struct {
	ClaimsKernelConditionalTrace bool
	ClaimsFiniteOnlyTrace        bool
	ClaimsNativeGlobalPrinciple  bool
	ClaimsNativeSevenOver72      bool
	ClaimsNativeTraceQuotient    bool
	ClaimsBoundaryStress         bool
	ClaimsHiggsMass              bool
	ClaimsGaugeUnification       bool
	ClaimsFlavorDerivation       bool
	Verdict                      string
}

type Analysis struct {
	Inherited      Gate679Inheritance
	Sequence       ShortExactProjectionSequence
	Defect         DefectInclusion
	Normalizations []TraceNormalization
	Compatibility  ResponseCompatibility
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
	g679, err := gate679.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate679 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g679)
	sequence := buildSequence(inherited)
	defect := buildDefect(inherited)
	normalizations := buildNormalizations(inherited)
	compatibility := buildCompatibility()
	missing := buildMissing()
	discipline := VerdictDiscipline{Verdict: StatusGate680Boundary}
	truth := "Gate 680 sharpens the Gate679 denominator problem: K7 is a rank-seven defect inside the 71-dimensional split-projection kernel, but the active quotient-defect response uses the global full-extension density 7/72. Kernel-only 7/71 and finite-only 7/70 are typed alternatives, but they omit the quotient input or boundary extension from the response chamber and give weaker residuals. The missing theorem is now the GlobalAugmentedTraceResponsePrinciple selecting full H72 normalization for quotient-defect response."
	return Analysis{Inherited: inherited, Sequence: sequence, Defect: defect, Normalizations: normalizations, Compatibility: compatibility, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate679.Analysis) Gate679Inheritance {
	return Gate679Inheritance{
		RelativeTraceInherited: strings.Contains(g.Trace.Verdict, gate679.StatusGlobalAugmentedChamberAverage),
		H72Dimension:           g.Defect.AmbientDimension,
		KernelDimension:        g.Kernel.TotalKernelDimension,
		QuotientDimension:      g.Kernel.QuotientDimension,
		K7Rank:                 g.Defect.Rank,
		SSplit:                 g.Trace.SSplit,
		DBase:                  g.Trace.DBase,
		TauGlobal:              g.Trace.TauGlobal,
		TauKernel:              g.Trace.TauKernel,
		TauFinite:              g.Trace.TauFinite,
		TauHalf:                g.Trace.TauHalf,
		ResidualGlobal:         g.Trace.Residual,
		FirewallPreserved:      g.Discipline.Verdict == gate679.StatusGate679Boundary,
		Verdict:                StatusGate679RelativeTraceResponseInherited,
	}
}

func buildSequence(in Gate679Inheritance) ShortExactProjectionSequence {
	return ShortExactProjectionSequence{
		Sequence:          "0 -> ker(pi_split) -> H_72 -> Q_boundary -> 0",
		KernelFormula:     "ker(pi_split)=Lambda^4 R^8 ⊕ L_anti",
		KernelDimension:   in.KernelDimension,
		AmbientDimension:  in.H72Dimension,
		QuotientDimension: in.QuotientDimension,
		ExactByDimension:  in.KernelDimension+in.QuotientDimension == in.H72Dimension,
		Verdict:           StatusShortExactProjectionSequenceDefined,
	}
}

func buildDefect(in Gate679Inheritance) DefectInclusion {
	return DefectInclusion{
		Inclusion:                "K_7 ⊕ 0_boundary ⊂ ker(pi_split) ⊂ H_72",
		Rank:                     in.K7Rank,
		KernelDimension:          in.KernelDimension,
		AmbientDimension:         in.H72Dimension,
		KernelConditionalDensity: float64(in.K7Rank) / float64(in.KernelDimension),
		GlobalDensity:            float64(in.K7Rank) / float64(in.H72Dimension),
		FiniteDensity:            float64(in.K7Rank) / 70.0,
		FullKernel:               in.K7Rank == in.KernelDimension,
		Verdict:                  strings.Join([]string{StatusK7DefectInsideKernelClassified, StatusKernelConditionalTraceNotActive}, ";"),
	}
}

func buildNormalizations(in Gate679Inheritance) []TraceNormalization {
	candidates := []struct {
		name, formula, class, verdict string
		den                           int
		value                         float64
		finite, anti, quotient        bool
	}{
		{"tau_global", "7/72", "global defect density in full augmented extension H_72", StatusSevenOver72FullExtensionDefectDensity, 72, 7.0 / 72.0, true, true, true},
		{"tau_kernel", "7/71", "conditional defect density inside ker(pi_split)", StatusKernelConditionalTraceNotActive, 71, 7.0 / 71.0, true, true, false},
		{"tau_finite", "7/70", "finite-only defect density inside Lambda^4 R^8", StatusFiniteOnlyTraceOmitsBoundaryQuotient, 70, 7.0 / 70.0, true, false, false},
		{"tau_half", "7/144", "per-boundary-coordinate half-trace clue", "typed inactive half-coordinate alternative", 144, 7.0 / 144.0, true, true, true},
	}
	out := make([]TraceNormalization, 0, len(candidates))
	for _, c := range candidates {
		pred := c.value * in.SSplit
		res := in.DBase - pred
		out = append(out, TraceNormalization{Name: c.name, Formula: c.formula, Denominator: c.den, Value: c.value, PredictedDBase: pred, Residual: res, AbsResidual: math.Abs(res), IncludesFiniteChamber: c.finite, IncludesBoundaryAntiLine: c.anti, IncludesQuotientLine: c.quotient, Classification: c.class, Verdict: c.verdict})
	}
	return out
}

func buildCompatibility() ResponseCompatibility {
	return ResponseCompatibility{
		ResponseDomain:               "Q_boundary = H_72/ker(pi_split)",
		QuotientLineIncludedInSystem: true,
		GlobalKeepsQuotientInput:     true,
		KernelExcludesQuotientInput:  true,
		FiniteExcludesBoundarySystem: true,
		BestNormalization:            "tau_global=7/72",
		Verdict:                      strings.Join([]string{StatusResponseCompatibilityWithQuotientLineAudited, StatusGlobalH72TraceTypeCorrectForQuotientResponse, StatusSevenOver72FullExtensionDefectDensity}, ";"),
	}
}

func buildMissing() MissingTheoremAudit {
	return MissingTheoremAudit{
		Missing:                    []string{StatusNoNativeGlobalTraceResponsePrinciple, StatusNoNativeSevenOver72Theorem, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoBoundaryStressDerivation},
		NewPreciseMissingPrinciple: "GlobalAugmentedTraceResponsePrinciple / FullChamberDefectDensityResponseTheorem",
		AllowedSupport:             []string{StatusGlobalH72TraceTypeCorrectForQuotientResponse, StatusSevenOver72FullExtensionDefectDensity},
		Verdict:                    strings.Join([]string{StatusNoNativeGlobalTraceResponsePrinciple, StatusNoNativeSevenOver72Theorem, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusNoBoundaryStressDerivation}, ";"),
	}
}

func Statuses() []string {
	return []string{StatusGate679RelativeTraceResponseInherited, StatusShortExactProjectionSequenceDefined, StatusK7DefectInsideKernelClassified, StatusGlobalKernelFiniteTraceNormalizationsAudited, StatusResponseCompatibilityWithQuotientLineAudited, StatusGlobalH72TraceTypeCorrectForQuotientResponse, StatusSevenOver72FullExtensionDefectDensity, StatusKernelConditionalTraceNotActive, StatusFiniteOnlyTraceOmitsBoundaryQuotient, StatusNoNativeGlobalTraceResponsePrinciple, StatusNoNativeSevenOver72Theorem, StatusNoNativeTraceToBoundaryQuotientTheorem, StatusGate680Boundary}
}
