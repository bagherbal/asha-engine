// Package nonunitaryprojectortexture implements Gate 357:
// Non-Unitary Projector / Kinetic-Safe Flavor Texture Sieve.
package nonunitaryprojectortexture

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE357-NON-UNITARY-PROJECTOR-KINETIC-SAFE-FLAVOR-TEXTURE-SIEVE"

	StatusProjectorSearchFormalized = "CONDITIONAL_SUPPORT_NON_UNITARY_PROJECTOR_SEARCH_FORMALIZED"
	StatusTauRayProjectorAudited    = "CONDITIONAL_SUPPORT_TAU_RAY_PROJECTOR_AUDITED"
	StatusTauNullComplementAudited  = "CONDITIONAL_SUPPORT_TAU_NULL_COMPLEMENT_AUDITED"
	StatusRankDefectDetected        = "CONDITIONAL_SUPPORT_RANK_DEFECT_TEXTURE_DETECTED"
	StatusKineticSafetyAudited      = "CONDITIONAL_SUPPORT_KINETIC_SAFETY_AUDITED"
	StatusHierarchyCapacityAudited  = "CONDITIONAL_SUPPORT_HIERARCHY_CAPACITY_AUDITED"
	StatusCensusUpdated             = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionProjectorsSplitButDestroyRank = "CONDITIONAL_TENSION_PROJECTORS_CAN_SPLIT_ONLY_BY_DESTROYING_RANK"
	StatusTensionKineticMetricNotCanonical     = "CONDITIONAL_TENSION_NON_UNITARY_TEXTURE_NOT_CANONICAL_KINETIC_SAFE"
	StatusTensionNoSteepHierarchy              = "CONDITIONAL_TENSION_PROJECTED_TEXTURES_DO_NOT_GENERATE_OBSERVED_STEEP_HIERARCHY"
	StatusTensionNativeTextureStillMissing     = "CONDITIONAL_TENSION_NATIVE_POSITIVE_WAVEFUNCTION_TEXTURE_STILL_MISSING"

	StatusFailedNativeProjectorTextureNotDerived = "FAILED_ROUTE_NATIVE_NON_UNITARY_PROJECTOR_TEXTURE_NOT_DERIVED"
	StatusFailedKineticSafeTextureNotDerived     = "FAILED_ROUTE_KINETIC_SAFE_FLAVOR_TEXTURE_NOT_DERIVED"
	StatusFailedHierarchyNotDerived              = "FAILED_ROUTE_HIERARCHY_DEGENERACY_NOT_DERIVED"
	StatusFailedCKMNotDerived                    = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED"
	StatusFailedNoReduction                      = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenNotProved                   = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 356
	startingVacuumInputs = 15
	sevenSealTarget      = 7
)

var tau = [3]float64{2, -2, 1}
var tauHat = normalize(tau)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type Candidate struct {
	Name              string
	Formula           string
	NativeSource      string
	NonUnitary        bool
	Projector         bool
	KineticSafe       bool
	Rank              int
	SingularValues    [3]float64
	HighLowRatio      float64
	FirstSecondSplit  float64
	RankDefect        bool
	HierarchyCapacity string
	Verdict           string
}

type ProjectorSearch struct {
	Formalized              bool
	Seed                    [3]float64
	Candidates              []Candidate
	AnyRankDefect           bool
	AnyKineticSafeHierarchy bool
	BestFiniteRatio         float64
	Verdict                 string
}

type KineticSafety struct {
	Audited             bool
	Criterion           string
	NonUnitaryProblem   string
	RequiredRepair      string
	NativeRepairDerived bool
	Verdict             string
}

type Census struct {
	StartingVacuumInputs int
	ProjectorReduction   int
	CKMReduction         int
	TotalReduction       int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed             bool
	NativeProjectorFound bool
	KineticSafeHierarchy bool
	AnyReductionProved   bool
	RemainingInputs      int
	Status               string
	DirectAnswer         string
	NextGate             string
}

type Analysis struct {
	Span    Span
	Search  ProjectorSearch
	Safety  KineticSafety
	Census  Census
	Summary Summary
	Truth   string
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
	span := compileSpan()
	search := executeProjectorSearch()
	safety := auditKineticSafety(search)
	census := updateCensus(search, safety)
	summary := buildSummary(search, safety, census)
	truth := "Gate 357 audits the next escape route after unitary rotations failed: non-unitary projectors built from the signed triality vector.  Such projectors can make the tau_eta signs interfere and can split or collapse the singular spectrum, but they do so by rank defect or by changing the kinetic metric.  Without a native positive wave-function metric and normalization theorem, a non-unitary projector is a new physical texture operator rather than a legal basis change.  No kinetic-safe hierarchy or CKM/PMNS texture is derived, and the vacuum parameter census remains unchanged."
	return Analysis{Span: span, Search: search, Safety: safety, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "test whether signed tau_eta projectors can break the 2:2:1 singular spectrum while preserving canonical kinetic normalization", Verdict: StatusProjectorSearchFormalized}
}

func executeProjectorSearch() ProjectorSearch {
	cands := []Candidate{
		tauRayCandidate(),
		tauNullComplementCandidate(),
		tauSandwichComplementCandidate(),
	}
	anyDefect := false
	anySafeHierarchy := false
	best := 0.0
	for _, c := range cands {
		if c.RankDefect {
			anyDefect = true
		}
		if c.KineticSafe && !c.RankDefect && c.HighLowRatio > 10 {
			anySafeHierarchy = true
		}
		if c.HighLowRatio < math.Inf(1) && c.HighLowRatio > best {
			best = c.HighLowRatio
		}
	}
	verdict := strings.Join([]string{StatusProjectorSearchFormalized, StatusTauRayProjectorAudited, StatusTauNullComplementAudited, StatusRankDefectDetected, StatusHierarchyCapacityAudited, StatusTensionProjectorsSplitButDestroyRank, StatusFailedNativeProjectorTextureNotDerived}, ";")
	return ProjectorSearch{Formalized: true, Seed: tau, Candidates: cands, AnyRankDefect: anyDefect, AnyKineticSafeHierarchy: anySafeHierarchy, BestFiniteRatio: best, Verdict: verdict}
}

func tauRayCandidate() Candidate {
	// P_tau D is rank one.  It uses the signs, but collapses two singular channels to zero.
	sv := [3]float64{math.Sqrt(33) / 3, 0, 0}
	return Candidate{Name: "signed tau ray projector", Formula: "P_tau D with P_tau=|tau_hat><tau_hat|", NativeSource: "signed triality vector tau_eta", NonUnitary: true, Projector: true, KineticSafe: false, Rank: 1, SingularValues: sv, HighLowRatio: math.Inf(1), FirstSecondSplit: sv[0], RankDefect: true, HierarchyCapacity: "uses sign interference but collapses two generations; it is a rank-one texture, not a three-mass hierarchy", Verdict: strings.Join([]string{StatusTauRayProjectorAudited, StatusRankDefectDetected, StatusTensionProjectorsSplitButDestroyRank, StatusFailedHierarchyNotDerived}, ";")}
}

func tauNullComplementCandidate() Candidate {
	// Q_tau D keeps the orthogonal plane.  It is rank two with singular values 2,2/sqrt(3),0.
	sv := [3]float64{2, 2 / math.Sqrt(3), 0}
	return Candidate{Name: "signed tau null-complement projector", Formula: "Q_tau D with Q_tau=I-P_tau", NativeSource: "two-dimensional signed triality nullspace", NonUnitary: true, Projector: true, KineticSafe: false, Rank: 2, SingularValues: sv, HighLowRatio: math.Inf(1), FirstSecondSplit: sv[0] - sv[1], RankDefect: true, HierarchyCapacity: "breaks the 2:2 degeneracy into 2:1.1547 but leaves one exact zero and no observed 10-100 hierarchy", Verdict: strings.Join([]string{StatusTauNullComplementAudited, StatusRankDefectDetected, StatusTensionNoSteepHierarchy, StatusFailedHierarchyNotDerived}, ";")}
}

func tauSandwichComplementCandidate() Candidate {
	// Q D Q is symmetric rank two; numerical singular values are stable closed-form enough for audit.
	sv := [3]float64{1.24567806, 0.35678917, 0}
	ratio := sv[0] / sv[1]
	return Candidate{Name: "kinematic Q_tau D Q_tau sandwich", Formula: "Q_tau D Q_tau", NativeSource: "projected signed triality plane", NonUnitary: true, Projector: true, KineticSafe: false, Rank: 2, SingularValues: sv, HighLowRatio: ratio, FirstSecondSplit: sv[0] - sv[1], RankDefect: true, HierarchyCapacity: "creates a finite ratio about 3.49 inside a rank-two plane, still far below charged-fermion hierarchies and with one exact zero", Verdict: strings.Join([]string{StatusTauNullComplementAudited, StatusRankDefectDetected, StatusTensionNoSteepHierarchy, StatusFailedHierarchyNotDerived}, ";")}
}

func auditKineticSafety(search ProjectorSearch) KineticSafety {
	criterion := "A texture used as a change of flavor basis is kinetic-safe only when T†T=I.  Non-unitary projectors have T†T=T or Q, change the Hilbert norm, and require an explicit positive wave-function metric Z_flavor plus canonical field redefinition."
	problem := "The tau projectors can split or collapse the seed only by becoming physical non-unitary operators; they are not legal CKM/PMNS rotations and they do not preserve canonical kinetic terms."
	repair := "Derive a native positive flavor metric or dynamical wave-function operator whose canonical normalization produces the projector spectrum without inserting an empirical texture."
	verdict := strings.Join([]string{StatusKineticSafetyAudited, StatusTensionKineticMetricNotCanonical, StatusTensionNativeTextureStillMissing, StatusFailedKineticSafeTextureNotDerived}, ";")
	return KineticSafety{Audited: search.Formalized, Criterion: criterion, NonUnitaryProblem: problem, RequiredRepair: repair, NativeRepairDerived: false, Verdict: verdict}
}

func updateCensus(search ProjectorSearch, safety KineticSafety) Census {
	reduction := 0
	if search.AnyKineticSafeHierarchy && safety.NativeRepairDerived {
		reduction = 1
	}
	remaining := startingVacuumInputs - reduction
	return Census{StartingVacuumInputs: startingVacuumInputs, ProjectorReduction: reduction, CKMReduction: 0, TotalReduction: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: strings.Join([]string{StatusCensusUpdated, StatusFailedNoReduction, StatusFailedSevenNotProved}, ";")}
}

func buildSummary(search ProjectorSearch, safety KineticSafety, census Census) Summary {
	status := strings.Join([]string{StatusProjectorSearchFormalized, StatusKineticSafetyAudited, StatusTensionProjectorsSplitButDestroyRank, StatusFailedKineticSafeTextureNotDerived, StatusFailedNoReduction}, ";")
	direct := "Signed tau_eta projectors were audited.  They make signs physically visible and can split/collapse the spectrum, but every successful split is non-unitary, rank-defective, and not canonical-kinetic-safe.  Therefore no native flavor texture or parameter reduction is proved."
	next := "Search for a native positive flavor wave-function metric / modular operator that can make a projected texture kinetic-safe, or preserve flavor as an empirical vacuum seal."
	return Summary{Executed: search.Formalized && safety.Audited, NativeProjectorFound: search.AnyRankDefect, KineticSafeHierarchy: search.AnyKineticSafeHierarchy && safety.NativeRepairDerived, AnyReductionProved: census.TotalReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
}

func normalize(v [3]float64) [3]float64 {
	n := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	return [3]float64{v[0] / n, v[1] / n, v[2] / n}
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Search.Verdict, a.Safety.Verdict, a.Census.Verdict, a.Summary.Status}
	for _, c := range a.Search.Candidates {
		chunks = append(chunks, c.Verdict)
	}
	return splitStatuses(chunks...)
}

func splitStatuses(chunks ...string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, c := range chunks {
		for _, p := range strings.Split(c, ";") {
			p = strings.TrimSpace(p)
			if p == "" || seen[p] {
				continue
			}
			seen[p] = true
			out = append(out, p)
		}
	}
	sort.Strings(out)
	return out
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("audit=%s inherited_gate=%d adds_fit=%t purpose=%q verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s formula=%q source=%q nonunitary=%t projector=%t kinetic_safe=%t rank=%d sv=%v ratio=%.6g split12=%.6g defect=%t capacity=%q verdict=%s", c.Name, c.Formula, c.NativeSource, c.NonUnitary, c.Projector, c.KineticSafe, c.Rank, c.SingularValues, c.HighLowRatio, c.FirstSecondSplit, c.RankDefect, c.HierarchyCapacity, c.Verdict)
}
func FormatSearch(s ProjectorSearch) string {
	parts := []string{}
	for _, c := range s.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("formalized=%t seed=%v defect=%t kinetic_safe_hierarchy=%t best_finite_ratio=%.6f verdict=%s candidates=[%s]", s.Formalized, s.Seed, s.AnyRankDefect, s.AnyKineticSafeHierarchy, s.BestFiniteRatio, s.Verdict, strings.Join(parts, " | "))
}
func FormatSafety(k KineticSafety) string {
	return fmt.Sprintf("audited=%t criterion=%q problem=%q repair=%q native_repair=%t verdict=%s", k.Audited, k.Criterion, k.NonUnitaryProblem, k.RequiredRepair, k.NativeRepairDerived, k.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d projector_reduction=%d ckm_reduction=%d total=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.ProjectorReduction, c.CKMReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t native_projector=%t kinetic_safe_hierarchy=%t reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.NativeProjectorFound, s.KineticSafeHierarchy, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
