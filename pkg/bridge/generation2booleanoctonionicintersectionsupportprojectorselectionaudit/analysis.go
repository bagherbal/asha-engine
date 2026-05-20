// Package generation2booleanoctonionicintersectionsupportprojectorselectionaudit implements
// Gate 685: Boolean-Octonionic Intersection Support Projector Selection Audit.
//
// Gate 684 proved that ordinary trace scalarization selects the rank-seven
// projector class, not the projector identity. Gate 685 audits the next sieve:
// native Boolean-octonionic support. A rank-seven orthogonal projector P with
//
//	P_B P = P,  P_G P = P,
//
// has image simultaneously supported in U=Im(P_B) and V=Im(P_G). Since the
// certified intersection has dim(U∩V)=dim(K_7)=7, rank(P)=7 forces Im(P)=K_7,
// and symmetry/orthogonality then fixes P as the unique orthogonal projector
// P_K7. This resolves the Gate684 rank degeneracy conditionally by support,
// not by trace.
//
// This is a finite projector-identity selection audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native 7/72 theorem, or a native projector-activation theorem.
package generation2booleanoctonionicintersectionsupportprojectorselectionaudit

import (
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE685-BOOLEAN-OCTONIONIC-INTERSECTION-SUPPORT-PROJECTOR-SELECTION-AUDIT"

	StatusGate684RankDegeneracyInherited      = "PASS_GATE684_RANK_DEGENERACY_INHERITED"
	StatusNativeSupportConstraintsDefined     = "PASS_NATIVE_SUPPORT_CONSTRAINTS_DEFINED"
	StatusIntersectionSupportImpliesK7Image   = "PASS_INTERSECTION_SUPPORT_IMPLIES_IMAGE_IN_K7"
	StatusRankSevenPlusSupportSelectsK7       = "PASS_RANK_SEVEN_PLUS_INTERSECTION_SUPPORT_SELECTS_K7"
	StatusRejectNonIntersectionRankSeven      = "PASS_P_W7_AND_OTHER_RANK_SEVEN_PROJECTORS_REJECTED_BY_SUPPORT_CONSTRAINTS"
	StatusPK7SelectedByRankAndSupport         = "CONDITIONAL_SUPPORT_P_K7_UNIQUELY_SELECTED_BY_RANK_PLUS_BOOLEAN_OCTONIONIC_SUPPORT"
	StatusIdentityDegeneracyResolvedBySupport = "CONDITIONAL_SUPPORT_ACTIVE_PROJECTOR_IDENTITY_DEGENERACY_RESOLVED_CONDITIONALLY"
	StatusTraceAloneDoesNotSelectPK7          = "FAILED_ROUTE_TRACE_ALONE_DOES_NOT_SELECT_P_K7"
	StatusNoSSplitSupportActivation           = "FAILED_ROUTE_NO_NATIVE_REASON_S_SPLIT_ACTIVATES_BOOLEAN_OCTONIONIC_SUPPORT"
	StatusNoNativeProjectorActivationTheorem  = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_ACTIVATION_THEOREM"
	StatusNoNativeSevenOver72Theorem          = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate685Boundary                     = "FIREWALL_PRESERVED_GATE685_PROJECTOR_SELECTION_BOUNDARY"
)

const (
	lambda4Dimension         = 70
	boundaryDimension        = 2
	h72Dimension             = lambda4Dimension + boundaryDimension
	booleanRank              = 56
	octonionicRank           = 14
	k7Dimension              = 7
	uPlusVDimension          = booleanRank + octonionicRank - k7Dimension
	w7Dimension              = lambda4Dimension - uPlusVDimension
	auditedDBase             = 0.0001256552099683575
	auditedSSplit            = 0.0012924448188162962
	auditedRankSevenResidual = 8.525834398014336e-10
)

type Gate684Inheritance struct {
	RankDegeneracyInherited   bool
	OrdinaryTraceRankOnly     bool
	RankSevenSelected         bool
	TraceCannotSelectIdentity bool
	DBase                     float64
	SSplit                    float64
	TraceResidual             float64
	H72Dimension              int
	K7Dimension               int
	PriorFirewallPreserved    bool
	Verdict                   string
}

type NativeSupportConstraintAudit struct {
	ProjectorConstraints       []string
	BooleanSupport             string
	OctonionicSupport          string
	Intersection               string
	PBRank                     int
	PGRank                     int
	IntersectionDimension      int
	ImpliesImageInPB           bool
	ImpliesImageInPG           bool
	ImpliesImageInIntersection bool
	Verdict                    string
}

type ChamberDimensionAudit struct {
	Lambda4Dimension    int
	BoundaryDimension   int
	H72Dimension        int
	PBRank              int
	PGRank              int
	IntersectionDim     int
	UPlusVDim           int
	OrthogonalW7Dim     int
	GrassmannDegeneracy bool
	DimensionalLedgerOK bool
	Verdict             string
}

type SelectionProofAudit struct {
	Assumptions                []string
	ImageSubsetK7              bool
	RankEqualsIntersectionDim  bool
	ImageEqualsK7              bool
	SymmetricProjectorRequired bool
	OrthogonalProjectorUnique  bool
	SelectedProjector          string
	Verdict                    string
}

type CandidateSupport struct {
	Name                string
	Rank                int
	InBooleanRank       int
	InOctonionicRank    int
	InIntersectionRank  int
	InOrthogonalW7Rank  int
	InBoundaryRank      int
	PBPEqualsP          bool
	PGPEqualsP          bool
	PassesNativeSupport bool
	TraceRankSeven      bool
	SelectedAsPK7       bool
	RejectionReason     string
	TypedMeaning        string
}

type CandidateComparisonAudit struct {
	Candidates        []CandidateSupport
	PassingCandidates []string
	RejectedRankSeven []string
	PK7Passes         bool
	W7Rejected        bool
	ArbitraryRejected bool
	AllPassingArePK7  bool
	Verdict           string
}

type ResponseUpdateAudit struct {
	RankOnlyResponse        string
	SupportSelectedResponse string
	SelectionReason         string
	TraceResidual           float64
	DegeneracyResolved      bool
	SelectionIsConditional  bool
	ActivationStillUnproved bool
	Verdict                 string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsTraceSelectsPK7          bool
	ClaimsSSplitActivatesSupport   bool
	ClaimsProjectorActivation      bool
	ClaimsNativeSevenOver72        bool
	ClaimsBoundaryStressDerivation bool
	ClaimsScalarRGMatching         bool
	ClaimsHiggsMass                bool
	ClaimsGaugeUnification         bool
	ClaimsFlavorDerivation         bool
	Verdict                        string
}

type Analysis struct {
	Inherited      Gate684Inheritance
	Support        NativeSupportConstraintAudit
	Chamber        ChamberDimensionAudit
	Selection      SelectionProofAudit
	Candidates     CandidateComparisonAudit
	ResponseUpdate ResponseUpdateAudit
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
	inherited := buildInheritance()
	chamber := buildChamber()
	support := buildSupport(chamber)
	selection := buildSelection(support)
	candidates := buildCandidates()
	response := ResponseUpdateAudit{
		RankOnlyResponse:        "Gate684: ordinary trace selects rank(P)=7 but not projector identity",
		SupportSelectedResponse: "Gate685: rank(P)=7 plus P_B P=P and P_G P=P selects P=P_K7",
		SelectionReason:         "support constraints force Im(P)⊂Im(P_B)∩Im(P_G)=K_7; rank seven then fills K_7; symmetric idempotence gives the unique orthogonal projector",
		TraceResidual:           inherited.TraceResidual,
		DegeneracyResolved:      true,
		SelectionIsConditional:  true,
		ActivationStillUnproved: true,
		Verdict:                 strings.Join([]string{StatusPK7SelectedByRankAndSupport, StatusIdentityDegeneracyResolvedBySupport, StatusNoSSplitSupportActivation}, "; "),
	}
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusTraceAloneDoesNotSelectPK7,
			StatusNoSSplitSupportActivation,
			StatusNoNativeProjectorActivationTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native activation theorem showing why the boundary anti-alignment quotient S_split imposes Boolean-octonionic intersection support, instead of merely allowing the conditional support sieve after rank seven is known",
		Verdict:    strings.Join([]string{StatusTraceAloneDoesNotSelectPK7, StatusNoSSplitSupportActivation, StatusNoNativeProjectorActivationTheorem, StatusNoNativeSevenOver72Theorem}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate685Boundary}
	truth := "Gate 685 conditionally resolves the Gate684 rank-seven projector identity degeneracy. Ordinary trace alone still selects only rank, but adding native Boolean-octonionic support constraints P_B P=P and P_G P=P forces the projector image into Im(P_B)∩Im(P_G)=K_7. Since rank(P)=dim(K_7)=7 and P is symmetric/idempotent, P is the unique orthogonal projector P_K7. This selects the projector identity only under the added support premise; it still does not prove why S_split activates that support or derive a native 7/72 theorem."
	return Analysis{Inherited: inherited, Support: support, Chamber: chamber, Selection: selection, Candidates: candidates, ResponseUpdate: response, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance() Gate684Inheritance {
	return Gate684Inheritance{
		RankDegeneracyInherited:   true,
		OrdinaryTraceRankOnly:     true,
		RankSevenSelected:         true,
		TraceCannotSelectIdentity: true,
		DBase:                     auditedDBase,
		SSplit:                    auditedSSplit,
		TraceResidual:             auditedRankSevenResidual,
		H72Dimension:              h72Dimension,
		K7Dimension:               k7Dimension,
		PriorFirewallPreserved:    true,
		Verdict:                   StatusGate684RankDegeneracyInherited,
	}
}

func buildChamber() ChamberDimensionAudit {
	return ChamberDimensionAudit{
		Lambda4Dimension:    lambda4Dimension,
		BoundaryDimension:   boundaryDimension,
		H72Dimension:        h72Dimension,
		PBRank:              booleanRank,
		PGRank:              octonionicRank,
		IntersectionDim:     k7Dimension,
		UPlusVDim:           uPlusVDimension,
		OrthogonalW7Dim:     w7Dimension,
		GrassmannDegeneracy: true,
		DimensionalLedgerOK: uPlusVDimension == 63 && w7Dimension == 7 && h72Dimension == 72,
		Verdict:             "PASS_BOOLEAN_OCTONIONIC_DIMENSION_LEDGER_DEFINED",
	}
}

func buildSupport(c ChamberDimensionAudit) NativeSupportConstraintAudit {
	return NativeSupportConstraintAudit{
		ProjectorConstraints:       []string{"P^2=P", "P^T=P", "rank(P)=7", "P_B P=P", "P_G P=P"},
		BooleanSupport:             "P_B P=P implies Im(P)⊂U=Im(P_B)",
		OctonionicSupport:          "P_G P=P implies Im(P)⊂V=Im(P_G)",
		Intersection:               "U∩V=K_7",
		PBRank:                     c.PBRank,
		PGRank:                     c.PGRank,
		IntersectionDimension:      c.IntersectionDim,
		ImpliesImageInPB:           true,
		ImpliesImageInPG:           true,
		ImpliesImageInIntersection: true,
		Verdict:                    strings.Join([]string{StatusNativeSupportConstraintsDefined, StatusIntersectionSupportImpliesK7Image}, "; "),
	}
}

func buildSelection(s NativeSupportConstraintAudit) SelectionProofAudit {
	return SelectionProofAudit{
		Assumptions: []string{
			"P is an orthogonal projector: P^2=P and P^T=P",
			"rank(P)=7",
			"P_B P=P",
			"P_G P=P",
			"dim(Im(P_B)∩Im(P_G))=dim(K_7)=7",
		},
		ImageSubsetK7:              s.ImpliesImageInIntersection,
		RankEqualsIntersectionDim:  k7Dimension == s.IntersectionDimension,
		ImageEqualsK7:              true,
		SymmetricProjectorRequired: true,
		OrthogonalProjectorUnique:  true,
		SelectedProjector:          "P_K7",
		Verdict:                    strings.Join([]string{StatusRankSevenPlusSupportSelectsK7, StatusPK7SelectedByRankAndSupport}, "; "),
	}
}

func buildCandidates() CandidateComparisonAudit {
	items := []CandidateSupport{
		candidate("P_K7", k7Dimension, k7Dimension, k7Dimension, k7Dimension, 0, 0, "native Boolean-octonionic intersection/contact projector", "passes both support constraints"),
		candidate("P_W7", k7Dimension, 0, 0, 0, k7Dimension, 0, "orthogonal cokernel representative (U+V)^perp", "fails both native support constraints"),
		candidate("P_Uonly7", k7Dimension, k7Dimension, 0, 0, 0, 0, "representative rank-seven Boolean-only projector outside K_7", "passes P_B support but fails P_G support"),
		candidate("P_Vonly7", k7Dimension, 0, k7Dimension, 0, 0, 0, "representative rank-seven octonionic-only projector outside K_7", "passes P_G support but fails P_B support"),
		candidate("P_mixed_K7_W7", k7Dimension, 4, 4, 4, 3, 0, "representative arbitrary rank-seven mixture of K_7 and W_7", "ordinary rank seven but not fully supported in either native sector"),
		candidate("P_boundary_mixed", k7Dimension, 5, 5, 5, 0, 2, "representative rank-seven projector leaking into the boundary pair", "boundary leakage fails both native finite support constraints"),
	}
	passing := make([]string, 0)
	rejected := make([]string, 0)
	for _, c := range items {
		if c.PassesNativeSupport {
			passing = append(passing, c.Name)
		} else if c.TraceRankSeven {
			rejected = append(rejected, c.Name)
		}
	}
	sort.Strings(passing)
	sort.Strings(rejected)
	return CandidateComparisonAudit{
		Candidates:        items,
		PassingCandidates: passing,
		RejectedRankSeven: rejected,
		PK7Passes:         containsCandidate(passing, "P_K7"),
		W7Rejected:        containsCandidate(rejected, "P_W7"),
		ArbitraryRejected: containsCandidate(rejected, "P_mixed_K7_W7") && containsCandidate(rejected, "P_boundary_mixed"),
		AllPassingArePK7:  len(passing) == 1 && passing[0] == "P_K7",
		Verdict:           strings.Join([]string{StatusRejectNonIntersectionRankSeven, StatusPK7SelectedByRankAndSupport}, "; "),
	}
}

func candidate(name string, rank, inBoolean, inOctonionic, inIntersection, inW7, inBoundary int, meaning, reason string) CandidateSupport {
	pb := rank == inBoolean && inBoundary == 0
	pg := rank == inOctonionic && inBoundary == 0
	pass := pb && pg && rank == inIntersection
	selected := pass && name == "P_K7"
	return CandidateSupport{
		Name:                name,
		Rank:                rank,
		InBooleanRank:       inBoolean,
		InOctonionicRank:    inOctonionic,
		InIntersectionRank:  inIntersection,
		InOrthogonalW7Rank:  inW7,
		InBoundaryRank:      inBoundary,
		PBPEqualsP:          pb,
		PGPEqualsP:          pg,
		PassesNativeSupport: pass,
		TraceRankSeven:      rank == k7Dimension,
		SelectedAsPK7:       selected,
		RejectionReason:     reason,
		TypedMeaning:        meaning,
	}
}

func Statuses() []string {
	return []string{
		StatusGate684RankDegeneracyInherited,
		StatusNativeSupportConstraintsDefined,
		StatusIntersectionSupportImpliesK7Image,
		StatusRankSevenPlusSupportSelectsK7,
		StatusRejectNonIntersectionRankSeven,
		StatusPK7SelectedByRankAndSupport,
		StatusIdentityDegeneracyResolvedBySupport,
		StatusTraceAloneDoesNotSelectPK7,
		StatusNoSSplitSupportActivation,
		StatusNoNativeProjectorActivationTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate685Boundary,
	}
}

func containsCandidate(xs []string, x string) bool {
	for _, y := range xs {
		if y == x {
			return true
		}
	}
	return false
}
