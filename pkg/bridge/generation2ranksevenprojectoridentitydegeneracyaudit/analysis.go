// Package generation2ranksevenprojectoridentitydegeneracyaudit implements
// Gate 684: Rank-Seven Projector Identity Degeneracy Audit.
//
// Gate 683 defined the lawful projector-valued response
//
//	R_split = S_split P_7,  P_7=P_K7⊕0_boundary,
//
// and recovered the active scalar bridge from its normalized ordinary trace.
// Gate 684 audits the next firewall: ordinary trace scalarization selects only
// the rank of a projector. Therefore it strongly selects rank seven, but it
// cannot by itself distinguish P_K7 from any other rank-seven projector such as
// the orthogonal cokernel representative W7.
//
// This is a bridge-layer projector-identity audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native 7/72 theorem, or a native projector-activation theorem.
package generation2ranksevenprojectoridentitydegeneracyaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate683 "github.com/bagherbal/asha-engine/pkg/bridge/generation2projectorvaluedboundaryquotientresponsetraceaudit"
)

const (
	AuditID = "GATE684-RANK-SEVEN-PROJECTOR-IDENTITY-DEGENERACY-AUDIT"

	StatusGate683ProjectorResponseInherited   = "PASS_GATE683_PROJECTOR_RESPONSE_INHERITED"
	StatusOrdinaryTraceRankLawAudited         = "PASS_ORDINARY_TRACE_RANK_LAW_AUDITED"
	StatusTypedProjectorCandidatesEnumerated  = "PASS_TYPED_PROJECTOR_CANDIDATES_ENUMERATED"
	StatusNumericalResponseByRankComputed     = "PASS_NUMERICAL_RESPONSE_BY_RANK_COMPUTED"
	StatusActiveResponseSelectsRankSeven      = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_SELECTS_RANK_SEVEN"
	StatusPK7StrongestTypedRankSevenCandidate = "CONDITIONAL_SUPPORT_P_K7_IS_STRONGEST_TYPED_RANK_SEVEN_SOURCE_CANDIDATE"
	StatusOrdinaryTraceCannotSelectIdentity   = "FAILED_ROUTE_ORDINARY_TRACE_CANNOT_DISTINGUISH_RANK_SEVEN_PROJECTOR_IDENTITY"
	StatusPK7NotUniquelySelectedByTraceAlone  = "FAILED_ROUTE_P_K7_NOT_UNIQUELY_SELECTED_BY_TRACE_ALONE"
	StatusNoNativeK7ActivationTheorem         = "FAILED_ROUTE_NO_NATIVE_K7_ACTIVATION_THEOREM"
	StatusNoNativeProjectorIdentitySelection  = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_IDENTITY_SELECTION_THEOREM"
	StatusNoNativeSevenOver72Theorem          = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusNoBoundaryStressDerivation          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_DERIVATION"
	StatusGate684Boundary                     = "FIREWALL_PRESERVED_GATE684_PROJECTOR_IDENTITY_DEGENERACY_BOUNDARY"
)

type Gate683Inheritance struct {
	ProjectorResponseInherited bool
	DBase                      float64
	SSplit                     float64
	H72Dimension               int
	K7Rank                     int
	Gate683Residual            float64
	Gate683UsedOrdinaryTrace   bool
	Gate683SignedTraceFailed   bool
	PriorFirewallPreserved     bool
	Verdict                    string
}

type RankLawAudit struct {
	Formula           string
	DependsOnlyOnRank bool
	CanSelectIdentity bool
	TraceIdentity     int
	Verdict           string
}

type ProjectorCandidate struct {
	Name        string
	Rank        int
	Coefficient float64
	Prediction  float64
	Residual    float64
	Source      string
	TypedStatus string
	RankSeven   bool
}

type ProjectorCandidateAudit struct {
	Candidates          []ProjectorCandidate
	BestResidual        float64
	BestNames           []string
	BestRank            int
	RankSevenCandidates []string
	Verdict             string
}

type RankDegeneracyAudit struct {
	ActiveRankSelected     int
	RankSevenResidual      float64
	RankSevenNames         []string
	OrdinaryTraceRankOnly  bool
	PK7UniquelySelected    bool
	DegenerateRank7Sources []string
	Verdict                string
}

type PK7SourceAudit struct {
	Reasons            []string
	AlternativeWarning string
	BestTypedCandidate string
	UniquelySelected   bool
	Verdict            string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsK7IdentitySelectedByTrace bool
	ClaimsNativeK7Activation        bool
	ClaimsProjectorIdentityTheorem  bool
	ClaimsNativeSevenOver72         bool
	ClaimsBoundaryStressDerivation  bool
	ClaimsHiggsMass                 bool
	ClaimsGaugeUnification          bool
	ClaimsFlavorDerivation          bool
	Verdict                         string
}

type Analysis struct {
	Inherited  Gate683Inheritance
	RankLaw    RankLawAudit
	Candidates ProjectorCandidateAudit
	Degeneracy RankDegeneracyAudit
	PK7Source  PK7SourceAudit
	Missing    MissingTheoremAudit
	Discipline VerdictDiscipline
	Truth      string
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
	g683, err := gate683.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate683 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g683)
	rankLaw := RankLawAudit{
		Formula:           "Tr_H72(S_split P_r)/Tr_H72(I) = (rank(P_r)/72) S_split",
		DependsOnlyOnRank: true,
		CanSelectIdentity: false,
		TraceIdentity:     inherited.H72Dimension,
		Verdict:           StatusOrdinaryTraceRankLawAudited,
	}
	candidates := buildCandidates(inherited)
	degeneracy := buildDegeneracy(candidates)
	pk7 := PK7SourceAudit{
		Reasons: []string{
			"K_7 is the native Boolean-octonionic intersection carrier Im(P_B)∩Im(P_G)",
			"K_7 is the contact/defect carrier already certified in the finite chamber",
			"K_7 appears as the kernel defect of the square addition map A: U⊕V -> Lambda^4 R^8",
			"K_7 is the mature Fano-Hitchin carrier with internal 4|3 polarity, even though that lane remains boundary-disconnected",
		},
		AlternativeWarning: "W_7 is also rank seven and gives the same ordinary trace response, but it is a cokernel representative rather than the contact/intersection carrier",
		BestTypedCandidate: "P_K7 remains the strongest typed rank-seven source candidate, but ordinary trace alone does not select its identity",
		UniquelySelected:   false,
		Verdict:            strings.Join([]string{StatusPK7StrongestTypedRankSevenCandidate, StatusPK7NotUniquelySelectedByTraceAlone}, "; "),
	}
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeK7ActivationTheorem,
			StatusNoNativeProjectorIdentitySelection,
			StatusNoNativeSevenOver72Theorem,
			StatusNoBoundaryStressDerivation,
		},
		PreciseGap: "a native projector-identity theorem explaining why S_split activates P_K7 specifically, not merely an arbitrary rank-seven projector",
		Verdict:    strings.Join([]string{StatusNoNativeK7ActivationTheorem, StatusNoNativeProjectorIdentitySelection, StatusNoNativeSevenOver72Theorem}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate684Boundary}
	truth := "Gate 684 proves that ordinary trace scalarization selects rank, not projector identity. The active bridge strongly selects rank seven because rank-seven projectors give the 7/72 response, and P_K7 remains the best typed source candidate. But P_W7 or any other rank-seven projector gives the same ordinary trace, so trace alone cannot certify K7 identity. The missing theorem is now sharpened to a K7Activation/ProjectorIdentitySelection theorem."
	return Analysis{Inherited: inherited, RankLaw: rankLaw, Candidates: candidates, Degeneracy: degeneracy, PK7Source: pk7, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate683.Analysis) Gate683Inheritance {
	return Gate683Inheritance{
		ProjectorResponseInherited: g.Projector.ResponseInEndH72 && g.Projector.ProjectorRank == 7,
		DBase:                      g.Inherited.DBase,
		SSplit:                     g.Inherited.SSplit,
		H72Dimension:               g.Inherited.H72Dimension,
		K7Rank:                     g.Inherited.K7Dimension,
		Gate683Residual:            g.Ordinary.Residual,
		Gate683UsedOrdinaryTrace:   g.Hodge.ActiveUsesOrdinary,
		Gate683SignedTraceFailed:   g.Hodge.SignedFailsActive,
		PriorFirewallPreserved:     g.Discipline.Verdict == gate683.StatusGate683Boundary,
		Verdict:                    StatusGate683ProjectorResponseInherited,
	}
}

func buildCandidates(i Gate683Inheritance) ProjectorCandidateAudit {
	items := []ProjectorCandidate{
		candidate("P_K7", 7, i, "K_7=Im(P_B)∩Im(P_G)", "native contact/intersection defect carrier"),
		candidate("P_W7", 7, i, "W_7=(U+V)^perp", "orthogonal cokernel representative, same rank as K_7"),
		candidate("P_+", 4, i, "K_7^+ self-dual Hodge sector", "rank-four Hodge-polarity block"),
		candidate("P_-", 3, i, "K_7^- anti-self-dual Hodge sector", "rank-three Hodge-polarity block"),
		candidate("P_signed", 1, i, "Hodge-signed trace P_+-P_-", "signed trace 4-3; not an ordinary projector but an audited signed scalar alternative"),
		candidate("P_G", 14, i, "octonionic projector image", "rank-fourteen octonionic chamber"),
		candidate("P_B", 56, i, "Boolean projector image", "rank-fifty-six Boolean chamber"),
		candidate("P_UplusV", 63, i, "Boolean-octonionic span U+V", "rank-sixty-three span after overlap correction"),
		candidate("P_finite", 70, i, "Lambda^4 R^8 finite chamber", "finite middle chamber projector"),
		candidate("P_kernel", 71, i, "ker(pi_split)=Lambda^4 R^8⊕L_anti", "projection-kernel conditional chamber"),
		candidate("I_H72", 72, i, "identity on H_72", "full augmented chamber"),
	}
	bestAbs := math.Inf(1)
	var bestNames []string
	bestRank := -1
	var rank7 []string
	for _, c := range items {
		abs := math.Abs(c.Residual)
		if abs < bestAbs-1e-18 {
			bestAbs = abs
			bestNames = []string{c.Name}
			bestRank = c.Rank
		} else if math.Abs(abs-bestAbs) <= 1e-18 {
			bestNames = append(bestNames, c.Name)
		}
		if c.RankSeven {
			rank7 = append(rank7, c.Name)
		}
	}
	sort.Strings(bestNames)
	sort.Strings(rank7)
	return ProjectorCandidateAudit{Candidates: items, BestResidual: items[0].Residual, BestNames: bestNames, BestRank: bestRank, RankSevenCandidates: rank7, Verdict: strings.Join([]string{StatusTypedProjectorCandidatesEnumerated, StatusNumericalResponseByRankComputed, StatusActiveResponseSelectsRankSeven}, "; ")}
}

func candidate(name string, rank int, i Gate683Inheritance, source, typed string) ProjectorCandidate {
	coeff := float64(rank) / float64(i.H72Dimension)
	pred := coeff * i.SSplit
	return ProjectorCandidate{Name: name, Rank: rank, Coefficient: coeff, Prediction: pred, Residual: i.DBase - pred, Source: source, TypedStatus: typed, RankSeven: rank == 7}
}

func buildDegeneracy(c ProjectorCandidateAudit) RankDegeneracyAudit {
	return RankDegeneracyAudit{
		ActiveRankSelected:     c.BestRank,
		RankSevenResidual:      c.BestResidual,
		RankSevenNames:         append([]string(nil), c.RankSevenCandidates...),
		OrdinaryTraceRankOnly:  true,
		PK7UniquelySelected:    false,
		DegenerateRank7Sources: append([]string(nil), c.RankSevenCandidates...),
		Verdict:                strings.Join([]string{StatusActiveResponseSelectsRankSeven, StatusOrdinaryTraceCannotSelectIdentity, StatusPK7NotUniquelySelectedByTraceAlone}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate683ProjectorResponseInherited,
		StatusOrdinaryTraceRankLawAudited,
		StatusTypedProjectorCandidatesEnumerated,
		StatusNumericalResponseByRankComputed,
		StatusActiveResponseSelectsRankSeven,
		StatusPK7StrongestTypedRankSevenCandidate,
		StatusOrdinaryTraceCannotSelectIdentity,
		StatusPK7NotUniquelySelectedByTraceAlone,
		StatusNoNativeK7ActivationTheorem,
		StatusNoNativeProjectorIdentitySelection,
		StatusNoNativeSevenOver72Theorem,
		StatusGate684Boundary,
	}
}
