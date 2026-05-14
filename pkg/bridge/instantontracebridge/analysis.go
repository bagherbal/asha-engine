// Package instantontracebridge implements Gate 175: finite-to-continuum
// instanton trace-normalization bridge.
//
// Gate 174 exposed the clean conditional branch S_top=8*pi^2 -> u=1/g_*^2=1,
// but refused to reduce strict nullity because two identifications were still
// missing: the finite contact index must be carried to the continuum
// Yang--Mills topological charge, and the finite trace/Hessian normalization
// must be carried to the continuum kinetic trace normalization. Gate 175 audits
// those missing bridges directly.
//
// The result is intentionally conservative. The topological seal and the Fock
// representation trace remain exact finite data, and the conditional u=1 branch
// remains available. But no currently derived object supplies the continuum
// four-manifold, principal bundle/connection, Chern--Weil curvature map, Hodge
// normalization, or absolute Lie-algebra trace normalization required to turn
// that branch into a strict physical coupling theorem.
package instantontracebridge

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/topologicalnormalization"
)

type InputAudit struct {
	Gate174ConditionalBranchAvailable bool
	Gate174StrictAbsoluteUDerived     bool
	RelativeGaugeRatioClosed          bool
	WeakAngleSeedClosed               bool
	ContactIndex                      float64
	TopologicalActionSeal             float64
	ConditionalUInverseGStar          float64
	ConditionalNullity                int
	StrictNullityBefore               int
	UsesObservedInput                 bool
	Verdict                           string
}

type BridgeRequirement struct {
	Name      string
	NeededFor string
	Available bool
	Canonical bool
	Detail    string
}

type CandidateRoute struct {
	Name                          string
	FiniteDataUsed                string
	Claim                         string
	ContinuumIndexBridgeDerived   bool
	TraceKineticBridgeDerived     bool
	FixesAbsoluteKappa            bool
	UsesConvention                bool
	UsesObservedInput             bool
	AdmissibleAsConditionalBranch bool
	AdmissibleAsStrictTheorem     bool
	Verdict                       string
}

type TraceNormalizationAudit struct {
	RepresentationTraceRatioClosed bool
	GeneratorNormalizationRelative bool
	AbsoluteTraceScaleDerived      bool
	FieldRescalingFreedomOpen      bool
	F0ConventionDependenceOpen     bool
	KineticIntegralNormalization   bool
	CandidateAbsoluteKappas        []string
	RejectedShortcutCount          int
	Verdict                        string
}

type BridgeFirewall struct {
	ContinuumIndexRequirements        int
	ContinuumIndexRequirementsMet     int
	TraceRequirements                 int
	TraceRequirementsMet              int
	CandidateRoutesAudited            int
	StrictContinuumIndexBridgeDerived bool
	StrictTraceKineticBridgeDerived   bool
	StrictAbsoluteUDerived            bool
	ConditionalAbsoluteUPreserved     bool
	StrictNullityBefore               int
	StrictNullityAfter                int
	ConditionalNullityAfter           int
	PhysicalCouplingsDerived          bool
	FineStructureDerived              bool
	BoundaryScaleDerived              bool
	ThresholdCorrectionsDerived       bool
	HiddenObservedInputUsed           bool
	RemainingStrictUnknowns           []string
	RecommendedNextGate               string
	Verdict                           string
}

type Analysis struct {
	Previous       topologicalnormalization.Analysis
	Input          InputAudit
	IndexNeeds     []BridgeRequirement
	TraceNeeds     []BridgeRequirement
	Routes         []CandidateRoute
	TraceAudit     TraceNormalizationAudit
	Firewall       BridgeFirewall
	TruthStatement string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		prev, err := topologicalnormalization.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(prev, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(prev topologicalnormalization.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	if !prev.Firewall.ConditionalAbsoluteUAvailable {
		return Analysis{}, fmt.Errorf("Gate 175 requires Gate 174 conditional topological-normalization branch")
	}
	if prev.Input.UsesObservedInput || prev.Firewall.HiddenObservedInputUsed {
		return Analysis{}, fmt.Errorf("Gate 175 refuses observed coupling input")
	}
	if prev.Input.ContactIndex <= eps || prev.Input.TopologicalActionSeal <= eps {
		return Analysis{}, fmt.Errorf("Gate 175 requires positive topological seal data")
	}
	if math.Abs(prev.Input.TopologicalActionSeal-8*math.Pi*math.Pi) > 1e-8 {
		return Analysis{}, fmt.Errorf("Gate 175 expected the 8π² topological action seal, got %.12g", prev.Input.TopologicalActionSeal)
	}

	input := InputAudit{
		Gate174ConditionalBranchAvailable: prev.Firewall.ConditionalAbsoluteUAvailable,
		Gate174StrictAbsoluteUDerived:     prev.Firewall.StrictAbsoluteUDerived,
		RelativeGaugeRatioClosed:          prev.Input.GaugeRatioClosed,
		WeakAngleSeedClosed:               prev.Input.WeakAngleSeedClosed,
		ContactIndex:                      prev.Input.ContactIndex,
		TopologicalActionSeal:             prev.Input.TopologicalActionSeal,
		ConditionalUInverseGStar:          prev.Matching.ConditionalUInverseGStar,
		ConditionalNullity:                prev.Firewall.ConditionalNullityAfter,
		StrictNullityBefore:               prev.Firewall.StrictNullityAfter,
		UsesObservedInput:                 false,
		Verdict:                           "Gate 174 supplies exact finite topological data and a conditional u=1 branch; Gate 175 audits whether the missing continuum bridges are actually derived",
	}

	indexNeeds := buildIndexNeeds()
	traceNeeds := buildTraceNeeds(prev)
	routes := buildRoutes()
	traceAudit := buildTraceAudit(prev)
	indexMet := countMet(indexNeeds)
	traceMet := countMet(traceNeeds)

	firewall := BridgeFirewall{
		ContinuumIndexRequirements:        len(indexNeeds),
		ContinuumIndexRequirementsMet:     indexMet,
		TraceRequirements:                 len(traceNeeds),
		TraceRequirementsMet:              traceMet,
		CandidateRoutesAudited:            len(routes),
		StrictContinuumIndexBridgeDerived: indexMet == len(indexNeeds),
		StrictTraceKineticBridgeDerived:   traceMet == len(traceNeeds),
		StrictAbsoluteUDerived:            false,
		ConditionalAbsoluteUPreserved:     prev.Firewall.ConditionalAbsoluteUAvailable && close(prev.Matching.ConditionalUInverseGStar, 1, 1e-10),
		StrictNullityBefore:               prev.Firewall.StrictNullityAfter,
		StrictNullityAfter:                prev.Firewall.StrictNullityAfter,
		ConditionalNullityAfter:           prev.Firewall.ConditionalNullityAfter,
		PhysicalCouplingsDerived:          false,
		FineStructureDerived:              false,
		BoundaryScaleDerived:              false,
		ThresholdCorrectionsDerived:       false,
		HiddenObservedInputUsed:           false,
		RemainingStrictUnknowns: []string{
			"u=1/g_*^2: strict-open because finite index and finite kinetic trace have not been canonically mapped to the continuum Yang--Mills normalization",
			"L=ln(M*/mu): boundary scale/evaluation scale remains open",
			"Δb_i: threshold activation/decoupling remains open",
		},
		RecommendedNextGate: "Gate 176 — conditional RG boundary-scale solvability audit under the quarantined u=1 branch",
		Verdict:             "the finite-to-continuum instanton trace-normalization bridge is not derived; the u=1 branch remains conditional and quarantined",
	}

	return Analysis{
		Previous:       prev,
		Input:          input,
		IndexNeeds:     indexNeeds,
		TraceNeeds:     traceNeeds,
		Routes:         routes,
		TraceAudit:     traceAudit,
		Firewall:       firewall,
		TruthStatement: "Gate 175 rejects a strict finite-to-continuum instanton normalization theorem at the current stage. The exact seal S_top=8π² and the representation trace ratio are real finite achievements, but they do not by themselves construct a continuum Chern--Weil topological charge or fix the absolute kinetic trace normalization. Therefore u=1/g_*²=1 remains a useful conditional branch, not a derived physical coupling, and strict nullity remains 3.",
	}, nil
}

func buildIndexNeeds() []BridgeRequirement {
	return []BridgeRequirement{
		{Name: "oriented continuum four-cycle", NeededFor: "define ∫ tr(F∧F)", Available: false, Canonical: false, Detail: "the finite engine has Lorentzian/signature data and contact carriers, not a canonical continuum four-manifold or compact four-cycle"},
		{Name: "principal gauge bundle", NeededFor: "define instanton number as a bundle characteristic class", Available: false, Canonical: false, Detail: "the Fock representation gives gauge generators, not a continuum principal bundle with transition functions"},
		{Name: "continuum connection curvature F", NeededFor: "evaluate Chern--Weil density", Available: false, Canonical: false, Detail: "finite curvature/Hessian diagnostics do not yet supply a local continuum field strength"},
		{Name: "Chern--Weil normalization", NeededFor: "identify finite I_BG with k=(1/8π²)∫tr(F∧F)", Available: false, Canonical: false, Detail: "no finite theorem fixes the trace convention inside the Chern--Weil integer"},
		{Name: "integer charge orientation", NeededFor: "fix sign and unit k=+1 rather than a convention", Available: true, Canonical: false, Detail: "the finite index is positive and unit-sized, but its continuum orientation/unit interpretation is still a matching rule"},
	}
}

func buildTraceNeeds(prev topologicalnormalization.Analysis) []BridgeRequirement {
	return []BridgeRequirement{
		{Name: "relative representation trace", NeededFor: "fix SU(2):U(1) ratio", Available: prev.Input.GaugeRatioClosed, Canonical: prev.Input.GaugeRatioClosed, Detail: "Gate 167 fixes K=(2,2,2,10/3), hence diag(1,1,1,5/3)"},
		{Name: "absolute finite Hilbert trace scale", NeededFor: "fix f0 rather than only ratios", Available: false, Canonical: false, Detail: "multiplying the full finite action by a constant preserves all relative traces"},
		{Name: "continuum kinetic inner product", NeededFor: "identify finite trace with ∫ tr(FμνFμν)", Available: false, Canonical: false, Detail: "no local field map/Hodge-star normalization is derived"},
		{Name: "generator trace convention", NeededFor: "match tr(TaTb)=c δab to the continuum convention", Available: true, Canonical: false, Detail: "commutators and ratios are fixed, but the continuum normalization constant c remains convention-sensitive"},
		{Name: "coupling placement convention", NeededFor: "decide whether 1/g² multiplies action or is absorbed into F", Available: false, Canonical: false, Detail: "field rescaling can move g between kinetic prefactor and covariant derivative until a continuum convention is derived"},
	}
}

func buildRoutes() []CandidateRoute {
	return []CandidateRoute{
		{Name: "direct seal identification", FiniteDataUsed: "S_top=8π² and I_BG=1", Claim: "identify S_top with a unit Yang--Mills instanton action", ContinuumIndexBridgeDerived: false, TraceKineticBridgeDerived: false, FixesAbsoluteKappa: false, UsesConvention: true, AdmissibleAsConditionalBranch: true, AdmissibleAsStrictTheorem: false, Verdict: "preserves the Gate-174 u=1 conditional branch, but assumes the missing bridges"},
		{Name: "representation trace normalization", FiniteDataUsed: "K_rep=(2,2,2,10/3)", Claim: "use fermion trace to fix absolute f0", ContinuumIndexBridgeDerived: false, TraceKineticBridgeDerived: false, FixesAbsoluteKappa: false, UsesConvention: true, AdmissibleAsConditionalBranch: true, AdmissibleAsStrictTheorem: false, Verdict: "fixes ratios only; multiplying all K_a by f0 remains free"},
		{Name: "canonical action Hessian", FiniteDataUsed: "Gate 100/102 Hessian ratio", Claim: "use scalar-orbit second variation to fix the absolute gauge kinetic scale", ContinuumIndexBridgeDerived: false, TraceKineticBridgeDerived: false, FixesAbsoluteKappa: false, UsesConvention: true, AdmissibleAsConditionalBranch: true, AdmissibleAsStrictTheorem: false, Verdict: "selects κ_U1 and the embedded ratio, not the continuum action unit"},
		{Name: "SU(2) generator algebra", FiniteDataUsed: "[T_i,T_j]=ε_ijk T_k and Tr_rep(T_i²)=2", Claim: "use Lie algebra closure to fix continuum normalization", ContinuumIndexBridgeDerived: false, TraceKineticBridgeDerived: false, FixesAbsoluteKappa: false, UsesConvention: true, AdmissibleAsConditionalBranch: false, AdmissibleAsStrictTheorem: false, Verdict: "closure fixes relative generator structure; continuum trace normalization is still a convention"},
		{Name: "observed coupling fit", FiniteDataUsed: "external α or measured g", Claim: "choose κ to match experiment", ContinuumIndexBridgeDerived: false, TraceKineticBridgeDerived: false, FixesAbsoluteKappa: true, UsesConvention: false, UsesObservedInput: true, AdmissibleAsConditionalBranch: false, AdmissibleAsStrictTheorem: false, Verdict: "forbidden by the no-observed-constants rule"},
	}
}

func buildTraceAudit(prev topologicalnormalization.Analysis) TraceNormalizationAudit {
	return TraceNormalizationAudit{
		RepresentationTraceRatioClosed: prev.Input.GaugeRatioClosed && prev.Input.WeakAngleSeedClosed,
		GeneratorNormalizationRelative: true,
		AbsoluteTraceScaleDerived:      false,
		FieldRescalingFreedomOpen:      true,
		F0ConventionDependenceOpen:     len(prev.Conventions) == 2 && prev.Conventions[0].ConditionalF0 != prev.Conventions[1].ConditionalF0,
		KineticIntegralNormalization:   false,
		CandidateAbsoluteKappas: []string{
			"κ=1 under direct unit-trace instanton convention",
			"f0=1/2 under 1/g_a²=f0 Tr_rep(T_a²)",
			"f0=1/4 under 1/g_a²=2f0 Tr_rep(T_a²)",
		},
		RejectedShortcutCount: 4,
		Verdict:               "relative trace normalization is closed, but absolute kinetic normalization remains convention-dependent",
	}
}

func countMet(xs []BridgeRequirement) int {
	n := 0
	for _, x := range xs {
		if x.Available && x.Canonical {
			n++
		}
	}
	return n
}

func close(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatInput(a InputAudit) string {
	return fmt.Sprintf("gate174Conditional=%t gate174StrictU=%t ratioClosed=%t weakSeed=%t I_BG=%.12g S_top=%.12g conditionalU=%.12g nullityStrictBefore=%d conditionalNullity=%d observed=%t",
		a.Gate174ConditionalBranchAvailable, a.Gate174StrictAbsoluteUDerived, a.RelativeGaugeRatioClosed, a.WeakAngleSeedClosed, a.ContactIndex, a.TopologicalActionSeal, a.ConditionalUInverseGStar, a.StrictNullityBefore, a.ConditionalNullity, a.UsesObservedInput)
}

func FormatRequirements(xs []BridgeRequirement) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s[needed=%s available=%t canonical=%t detail=%s]", x.Name, x.NeededFor, x.Available, x.Canonical, x.Detail)
	}
	return strings.Join(parts, "; ")
}

func FormatRoutes(xs []CandidateRoute) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%s: claim=%s strict=%t conditional=%t observed=%t verdict=%s", x.Name, x.Claim, x.AdmissibleAsStrictTheorem, x.AdmissibleAsConditionalBranch, x.UsesObservedInput, x.Verdict)
	}
	return strings.Join(parts, "; ")
}

func FormatTraceAudit(a TraceNormalizationAudit) string {
	return fmt.Sprintf("ratioClosed=%t relativeGeneratorNorm=%t absoluteTrace=%t fieldRescaleOpen=%t f0ConventionOpen=%t kineticIntegralNorm=%t kappas=%v rejectedShortcuts=%d verdict=%s",
		a.RepresentationTraceRatioClosed, a.GeneratorNormalizationRelative, a.AbsoluteTraceScaleDerived, a.FieldRescalingFreedomOpen, a.F0ConventionDependenceOpen, a.KineticIntegralNormalization, a.CandidateAbsoluteKappas, a.RejectedShortcutCount, a.Verdict)
}

func FormatFirewall(a BridgeFirewall) string {
	return fmt.Sprintf("indexReq=%d/%d traceReq=%d/%d routes=%d strictIndex=%t strictTrace=%t strictU=%t conditionalU=%t nullityStrict=%d->%d conditional->%d physicalCouplings=%t alpha=%t scale=%t thresholds=%t observed=%t",
		a.ContinuumIndexRequirementsMet, a.ContinuumIndexRequirements, a.TraceRequirementsMet, a.TraceRequirements, a.CandidateRoutesAudited, a.StrictContinuumIndexBridgeDerived, a.StrictTraceKineticBridgeDerived, a.StrictAbsoluteUDerived, a.ConditionalAbsoluteUPreserved, a.StrictNullityBefore, a.StrictNullityAfter, a.ConditionalNullityAfter, a.PhysicalCouplingsDerived, a.FineStructureDerived, a.BoundaryScaleDerived, a.ThresholdCorrectionsDerived, a.HiddenObservedInputUsed)
}
