// Package generation2coloredresttripletexclusivityanddustcapacityfalsificationaudit implements
// Gate 821: Colored RestTriplet Exclusivity and Dust-Capacity Falsification Audit.
//
// Gate 821 follows Gate 820 by auditing the stricter consequence of the literal
// 1+3 rest-simplex reading: one large non-top colored rest triplet may occupy
// the triplet chamber, and every remaining colored or uncolored rest atom must
// fit inside the tiny dust budget.
package generation2coloredresttripletexclusivityanddustcapacityfalsificationaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	AuditID = "GATE821-COLORED-REST-TRIPLET-EXCLUSIVITY-DUST-CAPACITY-FALSIFICATION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate820Inherited    = "PASS_GATE820_COLORED_REST_TRIPLET_CANDIDATE_INHERITED"
	StatusDustCapacityDerived = "PASS_DUST_CAPACITY_CONSEQUENCE_DERIVED"
	StatusSecondColoredBound  = "PASS_SECOND_COLORED_TRIPLET_BOUND_COMPUTED"
	StatusUncoloredDustBound  = "PASS_UNCOLORED_DUST_BOUND_COMPUTED"
	StatusBottomBranch        = "PASS_BOTTOM_COLOR_BRANCH_DEFINED"
	StatusCharmBranch         = "PASS_CHARM_COLOR_BRANCH_DEFINED"
	StatusAbstractBranch      = "PASS_ABSTRACT_COLORED_CHAMBER_BRANCH_DEFINED"
	StatusFailureBranch       = "PASS_FAILURE_BRANCH_DEFINED"
	StatusExternalProtocol    = "PASS_EXTERNAL_LEDGER_PROTOCOL_DEFINED"
	StatusNativeSourceAudit   = "PASS_NATIVE_SOURCE_AUDIT_DEFINED"
	StatusClassification      = "PASS_STATUS_CLASSIFICATION_DEFINED"
	StatusImpactFirewall      = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusPhysicalFirewalls   = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	SupportOneLargeTripletOnly   = "CONDITIONAL_SUPPORT_LITERAL_ONE_PLUS_THREE_SIMPLEX_PREDICTS_ONLY_ONE_LARGE_NON_TOP_COLORED_TRIPLET"
	SupportOtherColoredDustBound = "CONDITIONAL_SUPPORT_ALL_OTHER_COLORED_TRIPLETS_MUST_FIT_ALPHA_B_SQUARED_DUST_BOUND"
	SupportBottomTestable        = "CONDITIONAL_SUPPORT_BOTTOM_BRANCH_IS_TESTABLE_BUT_NOT_ACCEPTED_WITHOUT_LEDGER"
	SupportCharmTestable         = "CONDITIONAL_SUPPORT_CHARM_BRANCH_IS_TESTABLE_BUT_NOT_ACCEPTED_WITHOUT_LEDGER"
	SupportDustStronger          = "CONDITIONAL_SUPPORT_DUST_CAPACITY_IS_STRONGER_FALSIFICATION_THAN_B_OVER_T_MATCH_ALONE"
	SupportExternalFalsify       = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_FALSIFY_LITERAL_SECTOR_READING"
	SupportGJAfterBottomDust     = "CONDITIONAL_SUPPORT_GJ_LANE_ACTIVATES_ONLY_AFTER_BOTTOM_BRANCH_SURVIVES_DUST_TEST"
	SupportStrengthenedPartialR2 = "CONDITIONAL_SUPPORT_EXPECTED_STATUS_IS_STRENGTHENED_PARTIAL_R2_WITH_HARD_DUST_CAPACITY_TESTS"

	FailureMatchAloneInsufficient = "FAILED_ROUTE_ONE_LARGE_TRIPLET_MATCH_ALONE_DOES_NOT_PROVE_SIMPLEX_SECTOR_ASSIGNMENT"
	FailureBottomNeedsDust        = "FAILED_ROUTE_BOTTOM_MATCH_NOT_ALLOWED_WITHOUT_DUST_CAPACITY_CHECK"
	FailureCharmNeedsDust         = "FAILED_ROUTE_CHARM_MATCH_NOT_ALLOWED_WITHOUT_DUST_CAPACITY_CHECK"
	FailureSecondColoredAboveDust = "FAILED_ROUTE_SECOND_COLORED_TRIPLET_ABOVE_ALPHA_B_SQUARED_FALSIFIES_LITERAL_ONE_PLUS_THREE_READING"
	FailureUncoloredAboveDust     = "FAILED_ROUTE_UNCOLORED_ATOM_ABOVE_THREE_ALPHA_B_SQUARED_FALSIFIES_TINY_DUST_READING_UNLESS_CONVENTION_EXCLUDES_IT"
	FailureProjectiveNotTheorem   = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_TRACE_ATOM_THEOREM"
	FailureAlphaNotSector         = "FAILED_ROUTE_BOUNDARY_ALPHA_B_NOT_SECTOR_ASSIGNMENT_THEOREM"
	FailureBFNNotOperator         = "FAILED_ROUTE_BOUNDARY_FN_PACKAGE_NOT_YUKAWA_OPERATOR_THEOREM"
	FailureExternalNotNative      = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	FailureNoUpdateCYukawa        = "FAILED_ROUTE_GATE821_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFIED_TRACE_MAGNITUDE_MAP_OR_VALIDATED_EXTERNAL_LEDGER"
	FailureCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureTreeProxyNotPole       = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate821         = "FIREWALL_PRESERVED_GATE821_COLORED_REST_TRIPLET_DUST_CAPACITY_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2                             float64
	AlphaB, BOverT, SqrtBOverT                         float64
	DustOverT, SqrtDustOverT                           float64
	TotalRestOverT, TripletTraceOverT                  float64
	SecondColoredPerColorBound, SecondColoredSqrtBound float64
	UncoloredAtomBound, UncoloredSqrtBound             float64
	NEffSimplex, CYukawaSimplex, CHiggsSimplex         float64
	OfficialCYukawa, OfficialCHiggs                    float64
	Verdicts, Supports, Failures                       []string
}

type DustCapacity struct {
	Equations                    []string
	Interpretation               string
	Verdicts, Supports, Failures []string
}

type Branch struct {
	Name                         string
	SurvivalCondition            []string
	BlockedPromotion             []string
	Verdicts, Supports, Failures []string
}

type ExternalLedgerProtocol struct {
	RequiredObjects              []string
	ComputedQuantities           []string
	Tests                        []string
	Forbidden                    []string
	CanFalsify                   bool
	Verdicts, Supports, Failures []string
}

type NativeSourceAudit struct {
	Lane                         string
	Supplies                     []string
	DoesNotSupply                []string
	Verdicts, Supports, Failures []string
}

type Status struct {
	Outcome                      string
	Level                        string
	NativeSourceFound            bool
	ExternalLedgerSupplied       bool
	CanUpdateCYukawa             bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced                                                                                                          bool
	MatchAloneInsufficient, BottomNeedsDust, CharmNeedsDust, SecondColoredBound, UncoloredBound, ProjectiveNotTheorem bool
	AlphaNotSector, BFNNotOperator, ExternalNotNative, NoCYukawaUpdate, CHiggsLevelB, TreeProxyNotPole                bool
	Verdict                                                                                                           string
}

type Analysis struct {
	Ledger    Ledger
	Capacity  DustCapacity
	Branches  []Branch
	Protocol  ExternalLedgerProtocol
	Native    []NativeSourceAudit
	Status    Status
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func BOverT(alpha float64) float64         { return alpha * (1.0 - alpha) }
func DustOverT(alpha float64) float64      { return 3.0 * alpha * alpha }
func NEffSimplex(alpha float64) float64    { return 3.0 + 6.0*alpha }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	bOverT := BOverT(alpha)
	dOverT := DustOverT(alpha)
	totalRest := 3.0 * alpha
	tripletTrace := 3.0 * bOverT
	secondColoredBound := alpha * alpha
	uncoloredBound := dOverT
	if math.Abs((totalRest-tripletTrace)-dOverT) > 1e-18 {
		return Analysis{}, fmt.Errorf("dust capacity mismatch: total-triplet=%g dust=%g", totalRest-tripletTrace, dOverT)
	}
	if secondColoredBound <= 0 || uncoloredBound <= 0 || bOverT <= 0 {
		return Analysis{}, fmt.Errorf("invalid positive bounds: B/T=%g C/T=%g L/T=%g", bOverT, secondColoredBound, uncoloredBound)
	}
	nEffSimplex := NEffSimplex(alpha)
	ledger := Ledger{
		NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2,
		AlphaB: alpha, BOverT: bOverT, SqrtBOverT: math.Sqrt(bOverT),
		DustOverT: dOverT, SqrtDustOverT: math.Sqrt(dOverT),
		TotalRestOverT: totalRest, TripletTraceOverT: tripletTrace,
		SecondColoredPerColorBound: secondColoredBound, SecondColoredSqrtBound: math.Sqrt(secondColoredBound),
		UncoloredAtomBound: uncoloredBound, UncoloredSqrtBound: math.Sqrt(uncoloredBound),
		NEffSimplex: nEffSimplex, CYukawaSimplex: CYukawaFromNEff(nEffSimplex), CHiggsSimplex: CHiggsFromNEff(nEffSimplex),
		OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		Verdicts: []string{StatusGate820Inherited, StatusDustCapacityDerived, StatusSecondColoredBound, StatusUncoloredDustBound},
		Supports: []string{SupportOneLargeTripletOnly, SupportOtherColoredDustBound, SupportDustStronger},
		Failures: []string{FailureMatchAloneInsufficient, FailureSecondColoredAboveDust, FailureUncoloredAboveDust},
	}

	capacity := DustCapacity{
		Equations: []string{
			"a_rest/T = 3 alpha_B",
			"3B/T = 3 alpha_B(1-alpha_B)",
			"D/T = 3 alpha_B^2",
			"second colored triplet: C/T <= alpha_B^2",
			"uncolored atom: L/T <= 3 alpha_B^2 unless convention excludes it",
		},
		Interpretation: "literal 1+3 sector reading permits one large non-top colored triplet; every other rest atom must fit the dust capacity",
		Verdicts:       []string{StatusDustCapacityDerived, StatusSecondColoredBound, StatusUncoloredDustBound},
		Supports:       []string{SupportOneLargeTripletOnly, SupportOtherColoredDustBound, SupportDustStronger},
		Failures:       []string{FailureMatchAloneInsufficient, FailureSecondColoredAboveDust, FailureUncoloredAboveDust},
	}

	branches := []Branch{
		{Name: "bottom-color branch", SurvivalCondition: []string{"h_b/T near alpha_B(1-alpha_B)", "charm must fit alpha_B^2 dust bound", "tau/uncolored atoms must fit 3 alpha_B^2 unless excluded", "scale/scheme stability holds"}, BlockedPromotion: []string{"bottom match alone is not a Yukawa theorem", "bottom match alone is not GJ", "bottom match alone is not N_eff derivation"}, Verdicts: []string{StatusBottomBranch}, Supports: []string{SupportBottomTestable, SupportGJAfterBottomDust}, Failures: []string{FailureBottomNeedsDust, FailureMatchAloneInsufficient}},
		{Name: "charm-color branch", SurvivalCondition: []string{"h_c/T near alpha_B(1-alpha_B)", "bottom must fit alpha_B^2 dust bound", "tau/uncolored atoms must fit 3 alpha_B^2 unless excluded", "all remaining atoms fit dust"}, BlockedPromotion: []string{"charm match alone is not a Yukawa theorem", "charm match fails if bottom exceeds dust"}, Verdicts: []string{StatusCharmBranch}, Supports: []string{SupportCharmTestable}, Failures: []string{FailureCharmNeedsDust, FailureMatchAloneInsufficient}},
		{Name: "abstract colored chamber branch", SurvivalCondition: []string{"one color-tripled rest block", "all remaining atoms below dust capacity", "q_rest matches simplex concentration"}, BlockedPromotion: []string{"external R3-ready if ledger validated", "not native R4"}, Verdicts: []string{StatusAbstractBranch}, Supports: []string{SupportExternalFalsify}, Failures: []string{FailureExternalNotNative}},
		{Name: "failure branch", SurvivalCondition: []string{"two or more non-top colored triplets exceed alpha_B^2", "or uncolored atoms exceed 3 alpha_B^2", "or selector/convention must be forced"}, BlockedPromotion: []string{"literal 1+3 sector reading fails", "simplex may survive only as aggregate concentration"}, Verdicts: []string{StatusFailureBranch}, Failures: []string{FailureSecondColoredAboveDust, FailureUncoloredAboveDust}},
	}

	protocol := ExternalLedgerProtocol{
		RequiredObjects:    []string{"T=h_t", "colored atoms h_b,h_c,h_s,h_u,h_d", "uncolored atoms h_tau,h_mu,h_e plus explicit neutrino convention", "scale and scheme", "Yukawa normalization", "color convention"},
		ComputedQuantities: []string{"R_colored_sorted=sorted({h_b,h_c,h_s,h_u,h_d}/T)", "R_1 largest colored ratio", "R_k remaining colored ratios", "L_i/T uncolored ratios", "a_rest,b_rest,q_rest_ext"},
		Tests:              []string{"T1: R_1 ≈ alpha_B(1-alpha_B)", "T2: R_k <= alpha_B^2 for k >= 2", "T3: L_i/T <= 3 alpha_B^2 unless excluded by convention", "T4: total rest trace matches 3 alpha_B", "T5: q_rest_ext matches q_simplex(alpha_B)", "T6: no coefficient or selector is retuned after ledger import"},
		Forbidden:          []string{"choose top selector to force the result", "discard rest atoms to fit dust", "merge unrelated atoms to fake color triplet", "retune alpha_B, 9/5, or 6", "use Higgs mass or C_Higgs to tune ledger"},
		CanFalsify:         true,
		Verdicts:           []string{StatusExternalProtocol},
		Supports:           []string{SupportExternalFalsify, SupportDustStronger},
		Failures:           []string{FailureExternalNotNative, FailureMatchAloneInsufficient},
	}

	native := []NativeSourceAudit{
		{Lane: "finite spectral triple", Supplies: []string{"color multiplicity", "sector edge templates"}, DoesNotSupply: []string{"colored rest hierarchy", "B/T", "dust capacity sector atoms"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureBFNNotOperator}},
		{Lane: "projective/Fock 1+3", Supplies: []string{"one-line plus triplet shape"}, DoesNotSupply: []string{"bottom/charm/dust sector assignment", "trace atom theorem"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureProjectiveNotTheorem}},
		{Lane: "K7 4|3", Supplies: []string{"structural 4|3 resonance"}, DoesNotSupply: []string{"trace atom values", "colored rest triplet theorem"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureProjectiveNotTheorem}},
		{Lane: "boundary alpha_B", Supplies: []string{"rest weight", "dust capacity scale"}, DoesNotSupply: []string{"sector identity", "Yukawa ratio theorem"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureAlphaNotSector}},
		{Lane: "Boundary-FN package", Supplies: []string{"scalar closure", "partial positive compatibility"}, DoesNotSupply: []string{"sector atoms", "Yukawa operator theorem"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureBFNNotOperator}},
		{Lane: "Georgi-Jarlskog", Supplies: []string{"secondary diagnostic after bottom branch survives dust test"}, DoesNotSupply: []string{"low-scale dust capacity proof", "triplet identity"}, Verdicts: []string{StatusNativeSourceAudit}, Supports: []string{SupportGJAfterBottomDust}, Failures: []string{FailureMatchAloneInsufficient}},
		{Lane: "D4/triality", Supplies: []string{"airlocked structural search geometry"}, DoesNotSupply: []string{"rest atom hierarchy", "trace atom map"}, Verdicts: []string{StatusNativeSourceAudit}, Failures: []string{FailureBFNNotOperator}},
	}

	status := Status{Outcome: "Outcome C — no ledger exists", Level: "strengthened partial R2 with hard dust-capacity falsification tests; not external R3; not native R4", NativeSourceFound: false, ExternalLedgerSupplied: false, CanUpdateCYukawa: false, Verdicts: []string{StatusClassification}, Supports: []string{SupportStrengthenedPartialR2}, Failures: []string{FailureExternalNotNative, FailureNoUpdateCYukawa}}
	impact := Impact{CandidateNEff: nEffSimplex, CandidateCYukawa: CYukawaFromNEff(nEffSimplex), CandidateCHiggs: CHiggsFromNEff(nEffSimplex), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactFirewall}, Failures: []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole}}
	firewalls := Firewalls{Enforced: true, MatchAloneInsufficient: true, BottomNeedsDust: true, CharmNeedsDust: true, SecondColoredBound: true, UncoloredBound: true, ProjectiveNotTheorem: true, AlphaNotSector: true, BFNNotOperator: true, ExternalNotNative: true, NoCYukawaUpdate: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate821}

	truth := "Gate 821 strengthens Gate 820: a B/T match is insufficient; the literal 1+3 sector reading permits only one large non-top colored triplet and pushes every other rest atom into the tiny dust budget."
	final := "Without a decomposed trace ledger, the honest status is strengthened partial R2 with hard dust-capacity falsification inequalities; C_Yukawa remains unchanged."
	return Analysis{Ledger: ledger, Capacity: capacity, Branches: branches, Protocol: protocol, Native: native, Status: status, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func SortColoredRatios(ratios []float64) []float64 {
	out := append([]float64(nil), ratios...)
	sort.Sort(sort.Reverse(sort.Float64Slice(out)))
	return out
}

func CheckColoredDustCapacity(ratios []float64, alpha float64) (largest float64, remainingOK bool) {
	sorted := SortColoredRatios(ratios)
	if len(sorted) == 0 {
		return 0, true
	}
	bound := alpha * alpha
	for _, r := range sorted[1:] {
		if r > bound {
			return sorted[0], false
		}
	}
	return sorted[0], true
}

func CheckUncoloredDustCapacity(ratios []float64, alpha float64) bool {
	bound := 3.0 * alpha * alpha
	for _, r := range ratios {
		if r > bound {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{StatusGate820Inherited, StatusDustCapacityDerived, StatusSecondColoredBound, StatusUncoloredDustBound, StatusBottomBranch, StatusCharmBranch, StatusAbstractBranch, StatusFailureBranch, StatusExternalProtocol, StatusNativeSourceAudit, StatusClassification, StatusImpactFirewall, StatusPhysicalFirewalls, SupportOneLargeTripletOnly, SupportOtherColoredDustBound, SupportBottomTestable, SupportCharmTestable, SupportDustStronger, SupportExternalFalsify, SupportGJAfterBottomDust, SupportStrengthenedPartialR2, FailureMatchAloneInsufficient, FailureBottomNeedsDust, FailureCharmNeedsDust, FailureSecondColoredAboveDust, FailureUncoloredAboveDust, FailureProjectiveNotTheorem, FailureAlphaNotSector, FailureBFNNotOperator, FailureExternalNotNative, FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole, StatusFirewallGate821}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g B/T=%.16g sqrt(B/T)=%.16g a_rest/T=%.16g 3B/T=%.16g D/T=%.16g sqrt(D/T)=%.16g C/T_bound=%.16g sqrt(C/T)_bound=%.16g L/T_bound=%.16g sqrt(L/T)_bound=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.BOverT, a.SqrtBOverT, a.TotalRestOverT, a.TripletTraceOverT, a.DustOverT, a.SqrtDustOverT, a.SecondColoredPerColorBound, a.SecondColoredSqrtBound, a.UncoloredAtomBound, a.UncoloredSqrtBound)
}

func FormatCapacity(a DustCapacity) string {
	return fmt.Sprintf("%s equations=[%s]", a.Interpretation, strings.Join(a.Equations, "; "))
}

func FormatBranches(rows []Branch) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s survival=[%s] blocked=[%s]", r.Name, strings.Join(r.SurvivalCondition, "; "), strings.Join(r.BlockedPromotion, "; ")))
	}
	return strings.Join(out, " | ")
}

func FormatProtocol(a ExternalLedgerProtocol) string {
	return fmt.Sprintf("required=[%s] computed=[%s] tests=[%s] forbidden=[%s] canFalsify=%t", strings.Join(a.RequiredObjects, "; "), strings.Join(a.ComputedQuantities, "; "), strings.Join(a.Tests, "; "), strings.Join(a.Forbidden, "; "), a.CanFalsify)
}

func FormatNative(rows []NativeSourceAudit) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s supplies=[%s] doesNotSupply=[%s]", r.Lane, strings.Join(r.Supplies, ","), strings.Join(r.DoesNotSupply, ",")))
	}
	return strings.Join(out, " | ")
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.CandidateNEff, a.CandidateCYukawa, a.CandidateCHiggs, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range hay {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}
