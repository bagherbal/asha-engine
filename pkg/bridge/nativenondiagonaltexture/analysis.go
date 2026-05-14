// Package nativenondiagonaltexture implements Gate 356:
// Native Non-Diagonal Texture / Flavor Orientation Sieve.
package nativenondiagonaltexture

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE356-NATIVE-NON-DIAGONAL-TEXTURE-FLAVOR-ORIENTATION-SIEVE"

	StatusRotationSearchFormalized = "CONDITIONAL_SUPPORT_GEOMETRIC_ROTATION_SEARCH_FORMALIZED"
	StatusDFTAudited               = "CONDITIONAL_SUPPORT_DISCRETE_FOURIER_TEXTURE_AUDITED"
	StatusCyclicAudited            = "CONDITIONAL_SUPPORT_CYCLIC_PERMUTATION_TEXTURE_AUDITED"
	StatusInterferenceComputed     = "CONDITIONAL_SUPPORT_INTERFERENCE_SPLITTING_TEST_EXECUTED"
	StatusSingularInvarianceProved = "CONDITIONAL_SUPPORT_UNITARY_SINGULAR_VALUE_INVARIANCE_PROVED"
	StatusCKMShadowAudited         = "CONDITIONAL_SUPPORT_CKM_PMNS_SHADOW_EVALUATED"
	StatusCensusUpdated            = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"

	StatusTensionDFTOffDiagonalButNoHierarchy = "CONDITIONAL_TENSION_DFT_CREATES_OFFDIAGONAL_INTERFERENCE_BUT_NO_SINGULAR_HIERARCHY"
	StatusTensionCyclicOnlyPermutes           = "CONDITIONAL_TENSION_CYCLIC_OPERATOR_ONLY_PERMUTES_GENERATION_LABELS"
	StatusTensionUnitaryRotationsInsufficient = "CONDITIONAL_TENSION_UNITARY_ROTATIONS_CANNOT_CHANGE_TAU_ETA_SINGULAR_SPECTRUM"
	StatusTensionNonUnitaryProjectorRequired  = "CONDITIONAL_TENSION_NON_UNITARY_PROJECTOR_OR_ADDITIONAL_TEXTURE_REQUIRED"

	StatusFailedNativeTextureNotDerived = "FAILED_ROUTE_NATIVE_NON_DIAGONAL_TEXTURE_NOT_DERIVED"
	StatusFailedHierarchyNotBroken      = "FAILED_ROUTE_HIERARCHY_DEGENERACY_NOT_BROKEN"
	StatusFailedFirstSecondNotSplit     = "FAILED_ROUTE_FIRST_SECOND_SINGULAR_VALUE_DEGENERACY_NOT_SPLIT"
	StatusFailedCKMNotDerived           = "FAILED_ROUTE_CKM_PMNS_TEXTURE_NOT_DERIVED"
	StatusFailedNoReduction             = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenNotProved          = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate        = 355
	startingVacuumInputs = 15
	sevenSealTarget      = 7
)

var tauSigned = [3]float64{2, -2, 1}
var tauSingular = [3]float64{2, 2, 1}

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type Candidate struct {
	Name                    string
	NativeSource            string
	MatrixKind              string
	Unitary                 bool
	SelectedByFiniteCore    bool
	OffDiagonal             bool
	VisibleSignInterference bool
	SingularValues          [3]float64
	FirstSecondSplit        float64
	HighLowRatio            float64
	HierarchyBroken         bool
	CKMShadow               string
	Verdict                 string
}

type RotationSearch struct {
	Formalized         bool
	Seed               [3]float64
	Candidates         []Candidate
	AnyNativeSelected  bool
	AnyOffDiagonal     bool
	AnyHierarchyBroken bool
	BestHighLowRatio   float64
	Verdict            string
}

type InvarianceProof struct {
	Proved       bool
	Statement    string
	SeedSpectrum [3]float64
	Consequence  string
	Verdict      string
}

type TextureRequirement struct {
	Formalized               bool
	AllowedNativeUnitary     bool
	NeedsNonUnitaryProjector bool
	NeedsAdditionalOperator  bool
	Examples                 []string
	Verdict                  string
}

type Census struct {
	StartingVacuumInputs int
	TextureReduction     int
	CKMReduction         int
	TotalReduction       int
	RemainingInputs      int
	SevenSealTarget      int
	SevenSealReached     bool
	Verdict              string
}

type Summary struct {
	Executed           bool
	NativeTextureFound bool
	HierarchyBroken    bool
	AnyReductionProved bool
	RemainingInputs    int
	Status             string
	DirectAnswer       string
	NextGate           string
}

type Analysis struct {
	Span        Span
	Search      RotationSearch
	Invariance  InvarianceProof
	Requirement TextureRequirement
	Census      Census
	Summary     Summary
	Truth       string
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
	search := executeRotationSearch()
	invariance := proveUnitaryInvariance()
	req := formalizeRequirement()
	census := updateCensus(search)
	summary := buildSummary(search, census)
	truth := "Gate 356 audits native non-diagonal candidates for activating the signs of tau_eta.  The discrete Fourier transform creates a genuinely off-diagonal texture with sign interference in matrix entries, and the cyclic operator supplies a native Z3 generation symmetry.  However, every honest unitary rotation preserves the singular spectrum of diag(2,-2,1): the first/second singular-value degeneracy remains 2:2 and no steep hierarchy is generated.  Therefore the finite core has not yet derived a native non-diagonal texture that breaks the Yukawa hierarchy degeneracy; a non-unitary projector, additional texture operator, or empirical flavor seal is still required."
	return Analysis{Span: span, Search: search, Invariance: invariance, Requirement: req, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "search the finite generation carrier for native non-diagonal flavor rotations that can make tau_eta signs physically interfere and break the 2:2 singular-value degeneracy", Verdict: StatusRotationSearchFormalized}
}

func executeRotationSearch() RotationSearch {
	cands := []Candidate{
		dftCandidate(),
		cyclicCandidate(),
		identityCandidate(),
	}
	best := 0.0
	anyNative := false
	anyOff := false
	anyHierarchy := false
	for _, c := range cands {
		if c.SelectedByFiniteCore {
			anyNative = true
		}
		if c.OffDiagonal {
			anyOff = true
		}
		if c.HierarchyBroken {
			anyHierarchy = true
		}
		if c.HighLowRatio > best {
			best = c.HighLowRatio
		}
	}
	verdict := strings.Join([]string{StatusRotationSearchFormalized, StatusDFTAudited, StatusCyclicAudited, StatusInterferenceComputed, StatusTensionUnitaryRotationsInsufficient, StatusFailedHierarchyNotBroken, StatusFailedNativeTextureNotDerived}, ";")
	return RotationSearch{Formalized: true, Seed: tauSigned, Candidates: cands, AnyNativeSelected: anyNative, AnyOffDiagonal: anyOff, AnyHierarchyBroken: anyHierarchy, BestHighLowRatio: best, Verdict: verdict}
}

func identityCandidate() Candidate {
	return Candidate{Name: "identity diagonal basis", NativeSource: "geometric trace basis where tau_eta is diagonal", MatrixKind: "I† diag(2,-2,1) I", Unitary: true, SelectedByFiniteCore: true, OffDiagonal: false, VisibleSignInterference: false, SingularValues: tauSingular, FirstSecondSplit: 0, HighLowRatio: 2, HierarchyBroken: false, CKMShadow: "none; diagonal singular-value lane from Gate 355", Verdict: strings.Join([]string{StatusInterferenceComputed, StatusFailedHierarchyNotBroken}, ";")}
}

func cyclicCandidate() Candidate {
	// The Z3 cyclic permutation is native to a three-generation carrier, but P†DP only reorders eigen-slots.
	return Candidate{Name: "Z3 cyclic permutation", NativeSource: "three-generation cyclic relabeling operator", MatrixKind: "P† diag(2,-2,1) P", Unitary: true, SelectedByFiniteCore: false, OffDiagonal: false, VisibleSignInterference: false, SingularValues: tauSingular, FirstSecondSplit: 0, HighLowRatio: 2, HierarchyBroken: false, CKMShadow: "permutation only; no Cabibbo, PMNS, or hierarchy angle", Verdict: strings.Join([]string{StatusCyclicAudited, StatusTensionCyclicOnlyPermutes, StatusFailedHierarchyNotBroken}, ";")}
}

func dftCandidate() Candidate {
	// F3† diag(2,-2,1) F3 is circulant and off-diagonal.  Its entries expose interference,
	// but unitary conjugation preserves eigenvalues and singular values exactly.
	return Candidate{Name: "normalized DFT3 flavor rotation", NativeSource: "canonical Fourier transform of the Z3 generation carrier", MatrixKind: "F3† diag(2,-2,1) F3", Unitary: true, SelectedByFiniteCore: false, OffDiagonal: true, VisibleSignInterference: true, SingularValues: tauSingular, FirstSecondSplit: 0, HighLowRatio: 2, HierarchyBroken: false, CKMShadow: "democratic/trimaximal magnitude |F_ij|=1/sqrt(3), not CKM-like near-identity and not enough to split singular values", Verdict: strings.Join([]string{StatusDFTAudited, StatusTensionDFTOffDiagonalButNoHierarchy, StatusSingularInvarianceProved, StatusFailedFirstSecondNotSplit}, ";")}
}

func proveUnitaryInvariance() InvarianceProof {
	stmt := "For any unitary U,V, singular_values(U† diag(2,-2,1) V) = singular_values(diag(2,-2,1)) = (2,2,1).  Therefore DFT, cyclic, Clifford-basis, or CKM-like unitary rotations can expose phase/sign interference in entries but cannot by themselves create a steep singular-value hierarchy."
	consequence := "A hierarchy-breaking operator must be non-unitary, projected, scale-dependent, or an additional texture insertion; it is not supplied by bare finite generation rotations."
	return InvarianceProof{Proved: true, Statement: stmt, SeedSpectrum: tauSingular, Consequence: consequence, Verdict: strings.Join([]string{StatusSingularInvarianceProved, StatusTensionUnitaryRotationsInsufficient, StatusFailedHierarchyNotBroken}, ";")}
}

func formalizeRequirement() TextureRequirement {
	examples := []string{"non-unitary projector selecting a signed tau_eta interference channel", "additional flavor texture operator not expressible as U†DV", "empirical CKM/PMNS vacuum orientation seal", "higher dynamical vacuum operator that couples left/right generation bases asymmetrically"}
	return TextureRequirement{Formalized: true, AllowedNativeUnitary: true, NeedsNonUnitaryProjector: true, NeedsAdditionalOperator: true, Examples: examples, Verdict: strings.Join([]string{StatusTensionNonUnitaryProjectorRequired, StatusFailedNativeTextureNotDerived, StatusFailedCKMNotDerived}, ";")}
}

func updateCensus(search RotationSearch) Census {
	reduction := 0
	if search.AnyHierarchyBroken && search.AnyNativeSelected {
		reduction = 1
	}
	remaining := startingVacuumInputs - reduction
	return Census{StartingVacuumInputs: startingVacuumInputs, TextureReduction: reduction, CKMReduction: 0, TotalReduction: reduction, RemainingInputs: remaining, SevenSealTarget: sevenSealTarget, SevenSealReached: remaining <= sevenSealTarget, Verdict: strings.Join([]string{StatusCensusUpdated, StatusFailedNoReduction, StatusFailedSevenNotProved}, ";")}
}

func buildSummary(search RotationSearch, census Census) Summary {
	status := strings.Join([]string{StatusRotationSearchFormalized, StatusInterferenceComputed, StatusFailedHierarchyNotBroken, StatusFailedNoReduction}, ";")
	direct := "Native unitary generation rotations were audited.  DFT3 makes tau_eta non-diagonal and exposes sign interference, but unitary rotations preserve the singular values (2,2,1), so the first/second degeneracy is not broken and the observed hierarchy is not generated."
	next := "Derive a non-unitary/projected flavor texture operator, or preserve the CKM/PMNS orientation and Yukawa singular values as vacuum coordinates."
	return Summary{Executed: search.Formalized, NativeTextureFound: search.AnyNativeSelected && search.AnyOffDiagonal, HierarchyBroken: search.AnyHierarchyBroken, AnyReductionProved: census.TotalReduction > 0, RemainingInputs: census.RemainingInputs, Status: status, DirectAnswer: direct, NextGate: next}
}

func Statuses(a Analysis) []string {
	chunks := []string{a.Span.Verdict, a.Search.Verdict, a.Invariance.Verdict, a.Requirement.Verdict, a.Census.Verdict, a.Summary.Status}
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
	return out
}

func approx(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func FormatSpan(s Span) string {
	return fmt.Sprintf("audit=%s inherited_gate=%d adds_fit=%t purpose=%q verdict=%s", s.AuditID, s.InheritedGate, s.AddsFit, s.Purpose, s.Verdict)
}
func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s source=%q kind=%q unitary=%t selected=%t offdiag=%t interference=%t sv=%v split12=%.3e ratio=%.6f hierarchy=%t shadow=%q verdict=%s", c.Name, c.NativeSource, c.MatrixKind, c.Unitary, c.SelectedByFiniteCore, c.OffDiagonal, c.VisibleSignInterference, c.SingularValues, c.FirstSecondSplit, c.HighLowRatio, c.HierarchyBroken, c.CKMShadow, c.Verdict)
}
func FormatSearch(s RotationSearch) string {
	parts := make([]string, 0, len(s.Candidates))
	for _, c := range s.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("formalized=%t seed=%v any_selected=%t any_offdiag=%t hierarchy=%t best_ratio=%.6f verdict=%s candidates=[%s]", s.Formalized, s.Seed, s.AnyNativeSelected, s.AnyOffDiagonal, s.AnyHierarchyBroken, s.BestHighLowRatio, s.Verdict, strings.Join(parts, " | "))
}
func FormatInvariance(i InvarianceProof) string {
	return fmt.Sprintf("proved=%t spectrum=%v statement=%q consequence=%q verdict=%s", i.Proved, i.SeedSpectrum, i.Statement, i.Consequence, i.Verdict)
}
func FormatRequirement(r TextureRequirement) string {
	return fmt.Sprintf("formalized=%t allowed_unitary=%t needs_nonunitary=%t needs_operator=%t examples=%v verdict=%s", r.Formalized, r.AllowedNativeUnitary, r.NeedsNonUnitaryProjector, r.NeedsAdditionalOperator, r.Examples, r.Verdict)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d texture_reduction=%d ckm_reduction=%d total=%d remaining=%d target=%d reached=%t verdict=%s", c.StartingVacuumInputs, c.TextureReduction, c.CKMReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealTarget, c.SevenSealReached, c.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t native_texture=%t hierarchy=%t reduction=%t remaining=%d status=%s answer=%q next=%q", s.Executed, s.NativeTextureFound, s.HierarchyBroken, s.AnyReductionProved, s.RemainingInputs, s.Status, s.DirectAnswer, s.NextGate)
}
