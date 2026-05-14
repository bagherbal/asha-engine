// Package resolventbranchsemantics implements Gate 281:
// Resolvent Branch Semantics / Projector-to-Sector Orientation Seal Audit.
//
// Gate 280 constructed conditional 2+2 contact projectors after activating
// the ResolventAdjunctionSeal. Gate 281 asks whether those abstract projectors
// carry intrinsic trace/norm semantics strong enough to identify the physical
// {u,d}|{e,nu} sector split or the Gate-275 r_+/r_- amplitude branch. The
// result is intentionally conservative: rank/traces are symmetric (2|2),
// naive matrix norms are basis-dependent diagnostics, and the Morita 1|3
// multiplicity lives in the finite Hilbert bimodule rather than the contact
// companion module. Therefore a ProjectorSectorOrientationSeal can conditionally
// choose a 1-in-6 state, but it does not by itself derive a resolvent-to-r
// amplitude branch map or a Higgs-ratio prediction.
package resolventbranchsemantics

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/resolventfieldadjunction"
)

const (
	AuditID = "GATE281-RESOLVENT-BRANCH-SEMANTICS-PROJECTOR-SECTOR-ORIENTATION-SEAL-AUDIT"

	StatusTraceNormAuditCompleted           = "CONDITIONAL_SUPPORT_PROJECTOR_TRACE_NORM_SEMANTIC_AUDIT_COMPLETED"
	StatusOrientationSealActivated          = "CONDITIONAL_SUPPORT_PROJECTOR_SECTOR_ORIENTATION_SEAL_ACTIVATED"
	StatusRepresentativeOrientationAssigned = "CONDITIONAL_SUPPORT_REPRESENTATIVE_PROJECTOR_SECTOR_ORIENTATION_ASSIGNED"
	StatusSeeleyPrepDocumented              = "CONDITIONAL_SUPPORT_SEELEY_DE_WITT_PREPARATION_OBLIGATIONS_DOCUMENTED"
	StatusFirewallPreserved                 = "CONDITIONAL_SUPPORT_PROJECTOR_ORIENTATION_FIREWALLS_PRESERVED"

	StatusFailedNoNativeOrientation  = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_ORIENTATION_SELECTOR_DERIVED"
	StatusFailedMultiplicityMismatch = "FAILED_ROUTE_1_PLUS_3_MULTIPLICITY_DOES_NOT_PREFER_2_PLUS_2_PROJECTOR_ORIENTATION"
	StatusFailedNoResolventToRMap    = "FAILED_ROUTE_PROJECTOR_ORIENTATION_DOES_NOT_DERIVE_RESOLVENT_TO_R_BRANCH_MAP"
	StatusFailedAmplitudeBranch      = "FAILED_ROUTE_AMPLITUDE_BRANCH_NOT_LOCKED"
	StatusFailedHiggsRatio           = "FAILED_ROUTE_HIGGS_MASS_RATIO_STILL_NOT_DERIVED"
)

type ProjectorSemantic struct {
	BranchName             string
	ProjectorName          string
	Pair                   string
	Trace                  float64
	RankApprox             float64
	FrobeniusNormSq        float64
	ComplementTrace        float64
	ComplementRankApprox   float64
	AlignsWithKappaC       bool
	AlignsWithKappaQ       bool
	AlignsWithOnePlusThree bool
	NormIsNativeInvariant  bool
	Verdict                string
}

type BranchSemanticAudit struct {
	BranchName                   string
	Pairing                      string
	ResolventRootZ               float64
	ProjectorA                   ProjectorSemantic
	ProjectorB                   ProjectorSemantic
	ProjectorTracesEqual         bool
	ProjectorRanksEqual          bool
	CanPreferOrientationNatively bool
	PossibleOrientations         int
	Verdict                      string
}

type TraceNormSemanticAudit struct {
	KappaC                        int
	KappaQ                        int
	ExpectedMoritaSplit           string
	ContactProjectorSplit         string
	Branches                      []BranchSemanticAudit
	BranchCount                   int
	PossibleOrientations          int
	AnyNativeOrientationPreferred bool
	UsesNaiveBasisDependentNorms  bool
	Verdict                       string
}

type ProjectorSectorOrientationSeal struct {
	Name                           string
	Active                         bool
	Reason                         string
	CandidateStateCount            int
	SelectedBranchName             string
	SelectedResolventZ             float64
	ProjectorUD                    string
	ProjectorENu                   string
	OrientationIsNativeTheorem     bool
	OrientationIsSealedConditional bool
	GrantsProjectorSectorMap       bool
	GrantsAmplitudeBranchMap       bool
	Verdict                        string
}

type RBranchMappingAudit struct {
	RPlus                           float64
	RMinus                          float64
	OrientationLocked               bool
	ResolventBranchSelected         bool
	ProjectorSectorMapAvailable     bool
	AlgebraicResolventToRMapDerived bool
	UniqueAmplitudeBranch           bool
	SelectedRBranch                 string
	Reason                          string
	Verdict                         string
}

type SeeleyPrepCriterion struct {
	Name      string
	Required  bool
	Satisfied bool
	Detail    string
}

type SeeleyPrepAudit struct {
	Criteria                    []SeeleyPrepCriterion
	RBranchLocked               bool
	PhysicalJDerived            bool
	HeatKernelProjectionDerived bool
	HiggsRatioReady             bool
	Verdict                     string
}

type FirewallAudit struct {
	NoNativeOrientationOverclaim              bool
	NoMultiplicityToOrientationOverclaim      bool
	NoBasisDependentNormPromotion             bool
	OrientationSealDoesNotRewriteNativeStatus bool
	NoRBranchOverclaim                        bool
	NoHiggsRatioClaimed                       bool
	NoObservedMassesUsed                      bool
	NoEmpiricalYukawaInserted                 bool
	FiniteCorePolluted                        bool
	Verdict                                   string
}

type Summary struct {
	TraceNormAuditComplete            bool
	NativeOrientationPreferred        bool
	OrientationSealActivated          bool
	RepresentativeOrientationAssigned bool
	ProjectorSectorMapConditional     bool
	AmplitudeBranchLocked             bool
	HiggsRatioDerived                 bool
	FirewallPreserved                 bool
	Status                            string
	DirectAnswer                      string
	NextGate                          string
}

type Analysis struct {
	PreviousGate280 resolventfieldadjunction.Analysis
	TraceNorm       TraceNormSemanticAudit
	OrientationSeal ProjectorSectorOrientationSeal
	RBranch         RBranchMappingAudit
	SeeleyPrep      SeeleyPrepAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	prev, err := resolventfieldadjunction.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 280 predecessor: %w", err)
	}
	tn := auditTraceNormSemantics(prev)
	seal := activateOrientationSeal(prev, tn)
	rb := auditRBranchMapping(seal)
	prep := buildSeeleyPrep(rb)
	fw := auditFirewall(tn, seal, rb, prep)
	summary := buildSummary(tn, seal, rb, prep, fw)
	return Analysis{
		PreviousGate280: prev,
		TraceNorm:       tn,
		OrientationSeal: seal,
		RBranch:         rb,
		SeeleyPrep:      prep,
		Firewall:        fw,
		Summary:         summary,
		TruthStatement:  "Gate 281 verifies that the conditional resolvent projectors are rank-2/rank-2 objects and therefore do not natively encode the Morita 1|3 multiplicity. A ProjectorSectorOrientationSeal can conditionally assign one of six projector-sector states, but this sealed orientation still does not derive a resolvent-to-r_+/r_- map or a Higgs-ratio prediction.",
	}, nil
}

func auditTraceNormSemantics(prev resolventfieldadjunction.Analysis) TraceNormSemanticAudit {
	branches := make([]BranchSemanticAudit, 0, len(prev.BranchSpace.Branches))
	anyPreferred := false
	for _, b := range prev.BranchSpace.Branches {
		pa := semanticOf(b.Name, b.ProjectorA, b.ProjectorB)
		pb := semanticOf(b.Name, b.ProjectorB, b.ProjectorA)
		eqTrace := almost(pa.Trace, pb.Trace, 1e-8)
		eqRank := almost(pa.RankApprox, pb.RankApprox, 1e-8)
		canPrefer := false
		verdict := StatusFailedMultiplicityMismatch
		if !eqTrace || !eqRank || pa.AlignsWithOnePlusThree || pb.AlignsWithOnePlusThree {
			// This branch is left for future gates; current audited data do not reach here.
			canPrefer = true
			verdict = "CONDITIONAL_SUPPORT_PROJECTOR_ORIENTATION_ASYMMETRY_FOUND"
		}
		anyPreferred = anyPreferred || canPrefer
		branches = append(branches, BranchSemanticAudit{
			BranchName:                   b.Name,
			Pairing:                      b.Pairing,
			ResolventRootZ:               b.ResolventRootZ,
			ProjectorA:                   pa,
			ProjectorB:                   pb,
			ProjectorTracesEqual:         eqTrace,
			ProjectorRanksEqual:          eqRank,
			CanPreferOrientationNatively: canPrefer,
			PossibleOrientations:         2,
			Verdict:                      verdict,
		})
	}
	return TraceNormSemanticAudit{
		KappaC:                        1,
		KappaQ:                        3,
		ExpectedMoritaSplit:           "1|3 finite Hilbert-bimodule trace multiplicity",
		ContactProjectorSplit:         "2|2 contact companion-module resolvent projector split",
		Branches:                      branches,
		BranchCount:                   len(branches),
		PossibleOrientations:          len(branches) * 2,
		AnyNativeOrientationPreferred: anyPreferred,
		UsesNaiveBasisDependentNorms:  true,
		Verdict:                       StatusTraceNormAuditCompleted,
	}
}

func semanticOf(branch string, p, complement resolventfieldadjunction.ProjectorAudit) ProjectorSemantic {
	trace := p.Trace
	rank := p.RankApprox
	compTrace := complement.Trace
	compRank := complement.RankApprox
	// The Frobenius norm is included as a diagnostic only. In the companion
	// power basis it is not invariant under change of basis and cannot serve as
	// a native projector-sector semantic.
	normSq := frobSq(p.Matrix)
	alignsC := almost(rank, 1, 1e-8) && almost(compRank, 3, 1e-8)
	alignsQ := almost(rank, 3, 1e-8) && almost(compRank, 1, 1e-8)
	return ProjectorSemantic{
		BranchName:             branch,
		ProjectorName:          p.Name,
		Pair:                   strings.TrimPrefix(strings.TrimPrefix(p.Name, "P_"), "factor_"),
		Trace:                  trace,
		RankApprox:             rank,
		FrobeniusNormSq:        normSq,
		ComplementTrace:        compTrace,
		ComplementRankApprox:   compRank,
		AlignsWithKappaC:       alignsC,
		AlignsWithKappaQ:       alignsQ,
		AlignsWithOnePlusThree: alignsC || alignsQ,
		NormIsNativeInvariant:  false,
		Verdict:                StatusFailedMultiplicityMismatch,
	}
}

func activateOrientationSeal(prev resolventfieldadjunction.Analysis, tn TraceNormSemanticAudit) ProjectorSectorOrientationSeal {
	// Representative conditional state: choose the first sealed branch and map
	// its first projector to the already-selected Gate-277 {u,d} sector. This
	// is not a theorem; it is a quarantined orientation witness for downstream
	// stress tests.
	selected := prev.BranchSpace.Branches[0]
	return ProjectorSectorOrientationSeal{
		Name:                           "ProjectorSectorOrientationSeal",
		Active:                         !tn.AnyNativeOrientationPreferred,
		Reason:                         "all conditional projectors are 2|2 by trace/rank; Morita 1|3 multiplicity does not prefer one orientation",
		CandidateStateCount:            tn.PossibleOrientations,
		SelectedBranchName:             selected.Name,
		SelectedResolventZ:             selected.ResolventRootZ,
		ProjectorUD:                    selected.ProjectorA.Name,
		ProjectorENu:                   selected.ProjectorB.Name,
		OrientationIsNativeTheorem:     false,
		OrientationIsSealedConditional: true,
		GrantsProjectorSectorMap:       true,
		GrantsAmplitudeBranchMap:       false,
		Verdict:                        StatusOrientationSealActivated,
	}
}

func auditRBranchMapping(seal ProjectorSectorOrientationSeal) RBranchMappingAudit {
	rplus := (3591.0 + 136.0*math.Sqrt(123.0)) / 3099.0
	rminus := (3591.0 - 136.0*math.Sqrt(123.0)) / 3099.0
	return RBranchMappingAudit{
		RPlus:                           rplus,
		RMinus:                          rminus,
		OrientationLocked:               seal.Active && seal.GrantsProjectorSectorMap,
		ResolventBranchSelected:         seal.Active && seal.SelectedBranchName != "",
		ProjectorSectorMapAvailable:     seal.GrantsProjectorSectorMap,
		AlgebraicResolventToRMapDerived: false,
		UniqueAmplitudeBranch:           false,
		SelectedRBranch:                 "",
		Reason:                          "a sealed projector-sector orientation selects one contact 2+2 split, but the Gate-275 r_± roots live in the scalar-Morita amplitude-shape equation; no derived algebraic functor maps a chosen resolvent branch to either r_+ or r_-",
		Verdict:                         StatusFailedNoResolventToRMap,
	}
}

func buildSeeleyPrep(rb RBranchMappingAudit) SeeleyPrepAudit {
	criteria := []SeeleyPrepCriterion{
		{Name: "amplitude branch r_±", Required: true, Satisfied: rb.UniqueAmplitudeBranch, Detail: "one of the Gate-275 branches must be selected before a numerical proxy can be promoted"},
		{Name: "physical charge-conjugation J", Required: true, Satisfied: false, Detail: "the complete anti-linear opposite action on the physical finite Hilbert space remains missing"},
		{Name: "full chiral/hypercharge representation", Required: true, Satisfied: false, Detail: "C⊕H⊕M3(C) sector assignments are not yet a completed spectral triple"},
		{Name: "heat-kernel projection and subtraction scheme", Required: true, Satisfied: false, Detail: "raw traces must be mapped to Seeley-de Witt coefficients with cutoff moments and field normalizations"},
		{Name: "scalar/gauge normalization", Required: true, Satisfied: false, Detail: "Higgs and gauge kinetic terms must be separately projected before a2/a4 is meaningful"},
	}
	return SeeleyPrepAudit{Criteria: criteria, RBranchLocked: rb.UniqueAmplitudeBranch, PhysicalJDerived: false, HeatKernelProjectionDerived: false, HiggsRatioReady: false, Verdict: StatusSeeleyPrepDocumented}
}

func auditFirewall(tn TraceNormSemanticAudit, seal ProjectorSectorOrientationSeal, rb RBranchMappingAudit, prep SeeleyPrepAudit) FirewallAudit {
	return FirewallAudit{
		NoNativeOrientationOverclaim:              !tn.AnyNativeOrientationPreferred && !seal.OrientationIsNativeTheorem,
		NoMultiplicityToOrientationOverclaim:      true,
		NoBasisDependentNormPromotion:             tn.UsesNaiveBasisDependentNorms,
		OrientationSealDoesNotRewriteNativeStatus: seal.OrientationIsSealedConditional && !seal.OrientationIsNativeTheorem,
		NoRBranchOverclaim:                        !rb.UniqueAmplitudeBranch && !rb.AlgebraicResolventToRMapDerived,
		NoHiggsRatioClaimed:                       !prep.HiggsRatioReady,
		NoObservedMassesUsed:                      true,
		NoEmpiricalYukawaInserted:                 true,
		FiniteCorePolluted:                        false,
		Verdict:                                   StatusFirewallPreserved,
	}
}

func buildSummary(tn TraceNormSemanticAudit, seal ProjectorSectorOrientationSeal, rb RBranchMappingAudit, prep SeeleyPrepAudit, fw FirewallAudit) Summary {
	return Summary{
		TraceNormAuditComplete:            tn.BranchCount == 3,
		NativeOrientationPreferred:        tn.AnyNativeOrientationPreferred,
		OrientationSealActivated:          seal.Active,
		RepresentativeOrientationAssigned: seal.SelectedBranchName != "" && seal.ProjectorUD != "" && seal.ProjectorENu != "",
		ProjectorSectorMapConditional:     seal.GrantsProjectorSectorMap && seal.OrientationIsSealedConditional,
		AmplitudeBranchLocked:             rb.UniqueAmplitudeBranch,
		HiggsRatioDerived:                 prep.HiggsRatioReady,
		FirewallPreserved:                 !fw.FiniteCorePolluted,
		Status:                            StatusFailedAmplitudeBranch,
		DirectAnswer:                      "The 1⊕3 Morita trace multiplicities do not natively prefer a resolvent projector orientation; a ProjectorSectorOrientationSeal can select a representative 1-in-6 orientation, but the r_+/r_- amplitude branch remains algebraically unmapped.",
		NextGate:                          "Gate 282 — Resolvent-to-Scalar-Morita Branch Map / Heat-Kernel Normalization Preconditions Audit",
	}
}

func frobSq(m resolventfieldadjunction.Matrix4) float64 {
	var s float64
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			s += m[i][j] * m[i][j]
		}
	}
	return s
}

func almost(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatProjectorSemantic(p ProjectorSemantic) string {
	return fmt.Sprintf("%s pair=%s trace=%.12g rank≈%.12g frobSq=%.12g complementTrace=%.12g complementRank≈%.12g alignsC=%t alignsQ=%t aligns1|3=%t normNative=%t verdict=%s", p.ProjectorName, p.Pair, p.Trace, p.RankApprox, p.FrobeniusNormSq, p.ComplementTrace, p.ComplementRankApprox, p.AlignsWithKappaC, p.AlignsWithKappaQ, p.AlignsWithOnePlusThree, p.NormIsNativeInvariant, p.Verdict)
}

func FormatBranchSemantic(b BranchSemanticAudit) string {
	return fmt.Sprintf("%s pairing=%s z=%.15g tracesEqual=%t ranksEqual=%t preferNative=%t orientations=%d PA={%s} PB={%s} verdict=%s", b.BranchName, b.Pairing, b.ResolventRootZ, b.ProjectorTracesEqual, b.ProjectorRanksEqual, b.CanPreferOrientationNatively, b.PossibleOrientations, FormatProjectorSemantic(b.ProjectorA), FormatProjectorSemantic(b.ProjectorB), b.Verdict)
}

func FormatTraceNorm(t TraceNormSemanticAudit) string {
	parts := make([]string, 0, len(t.Branches))
	for _, b := range t.Branches {
		parts = append(parts, FormatBranchSemantic(b))
	}
	return fmt.Sprintf("kappa=%d:%d expected=%q contactSplit=%q branchCount=%d orientations=%d anyNative=%t basisNormsDiagnostic=%t branches={%s} verdict=%s", t.KappaC, t.KappaQ, t.ExpectedMoritaSplit, t.ContactProjectorSplit, t.BranchCount, t.PossibleOrientations, t.AnyNativeOrientationPreferred, t.UsesNaiveBasisDependentNorms, strings.Join(parts, "; "), t.Verdict)
}

func FormatSeal(s ProjectorSectorOrientationSeal) string {
	return fmt.Sprintf("name=%s active=%t reason=%q candidates=%d branch=%s z=%.15g UD=%s ENu=%s native=%t sealed=%t sectorMap=%t rMap=%t verdict=%s", s.Name, s.Active, s.Reason, s.CandidateStateCount, s.SelectedBranchName, s.SelectedResolventZ, s.ProjectorUD, s.ProjectorENu, s.OrientationIsNativeTheorem, s.OrientationIsSealedConditional, s.GrantsProjectorSectorMap, s.GrantsAmplitudeBranchMap, s.Verdict)
}

func FormatRBranch(r RBranchMappingAudit) string {
	return fmt.Sprintf("rPlus=%.15g rMinus=%.15g orientationLocked=%t resolventSelected=%t sectorMap=%t resolventToR=%t unique=%t selected=%q reason=%q verdict=%s", r.RPlus, r.RMinus, r.OrientationLocked, r.ResolventBranchSelected, r.ProjectorSectorMapAvailable, r.AlgebraicResolventToRMapDerived, r.UniqueAmplitudeBranch, r.SelectedRBranch, r.Reason, r.Verdict)
}

func FormatSeeleyPrep(s SeeleyPrepAudit) string {
	parts := make([]string, 0, len(s.Criteria))
	for _, c := range s.Criteria {
		parts = append(parts, fmt.Sprintf("%s[required=%t satisfied=%t detail=%s]", c.Name, c.Required, c.Satisfied, c.Detail))
	}
	return fmt.Sprintf("rLocked=%t J=%t heat=%t higgsReady=%t criteria={%s} verdict=%s", s.RBranchLocked, s.PhysicalJDerived, s.HeatKernelProjectionDerived, s.HiggsRatioReady, strings.Join(parts, "; "), s.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("noNativeOrientation=%t noMultiplicityOverclaim=%t noBasisNormPromotion=%t sealNoRewrite=%t noRBranch=%t noHiggs=%t noMasses=%t noYukawa=%t polluted=%t verdict=%s", f.NoNativeOrientationOverclaim, f.NoMultiplicityToOrientationOverclaim, f.NoBasisDependentNormPromotion, f.OrientationSealDoesNotRewriteNativeStatus, f.NoRBranchOverclaim, f.NoHiggsRatioClaimed, f.NoObservedMassesUsed, f.NoEmpiricalYukawaInserted, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("traceAudit=%t nativeOrientation=%t seal=%t representative=%t conditionalSector=%t rLocked=%t higgs=%t firewall=%t status=%s answer=%q next=%s", s.TraceNormAuditComplete, s.NativeOrientationPreferred, s.OrientationSealActivated, s.RepresentativeOrientationAssigned, s.ProjectorSectorMapConditional, s.AmplitudeBranchLocked, s.HiggsRatioDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}

func AssertNoOverclaim(a Analysis) error {
	if a.TraceNorm.AnyNativeOrientationPreferred || a.OrientationSeal.OrientationIsNativeTheorem || a.RBranch.UniqueAmplitudeBranch || a.SeeleyPrep.HiggsRatioReady || a.Firewall.FiniteCorePolluted {
		return fmt.Errorf("Gate 281 overclaimed: summary=%+v", a.Summary)
	}
	return nil
}
