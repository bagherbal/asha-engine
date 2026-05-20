// Package generation2oneplusthreerestsimplexsourceminimalityandexternalledgerfalsificationaudit implements
// Gate 819: OnePlusThree RestSimplex Source Minimality and External Ledger Falsification Audit.
//
// Gate 819 asks whether Gate 818's boundary-alpha 1+3 positive rest simplex is
// typed by a current ASHA source object, or whether it must be frozen as a
// falsifiable external-ledger hypothesis.
package generation2oneplusthreerestsimplexsourceminimalityandexternalledgerfalsificationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE819-ONE-PLUS-THREE-REST-SIMPLEX-SOURCE-MINIMALITY-EXTERNAL-LEDGER-FALSIFICATION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	C1NineFive = 9.0 / 5.0
	C2Six      = 6.0

	StatusGate818Inherited     = "PASS_GATE818_BOUNDARY_ALPHA_SIMPLEX_INHERITED"
	StatusSourceSealDefined    = "PASS_ONE_PLUS_THREE_REST_SIMPLEX_SOURCE_SEAL_DEFINED"
	StatusProjectiveAudited    = "PASS_PROJECTIVE_FOCK_ONE_PLUS_THREE_SOURCE_AUDITED"
	StatusK7Audited            = "PASS_K7_HODGE_4_3_SOURCE_AUDITED"
	StatusBoundaryAlphaAudited = "PASS_BOUNDARY_ALPHA_DISTINGUISHED_WEIGHT_AUDITED"
	StatusBoundaryColorAudited = "PASS_BOUNDARY_COLOR_COEFFICIENT_SOURCE_AUDITED"
	StatusExternalProtocol     = "PASS_EXTERNAL_LEDGER_FALSIFICATION_PROTOCOL_DEFINED"
	StatusNonCircularity       = "PASS_NONCIRCULARITY_FIREWALL_DEFINED"
	StatusRStatusDefined       = "PASS_R_STATUS_CLASSIFICATION_DEFINED"
	StatusImpactRecorded       = "PASS_C_YUKAWA_AND_C_HIGGS_IMPACT_RECORDED"
	StatusOutcomeRecorded      = "PASS_OUTCOME_CLASSIFICATION_RECORDED"
	StatusBranchRecorded       = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls    = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	SupportSharpestSimplex       = "CONDITIONAL_SUPPORT_ONE_PLUS_THREE_SIMPLEX_IS_CURRENT_SHARPEST_REST_CONCENTRATION_CANDIDATE"
	SupportProjectiveRelevant    = "CONDITIONAL_SUPPORT_PROJECTIVE_ONE_PLUS_THREE_IS_RELEVANT_SOURCE_CANDIDATE"
	SupportBoundaryAlphaNatural  = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_B_IS_NATURAL_DISTINGUISHED_REST_WEIGHT_CANDIDATE"
	SupportExternalFalsifies     = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_FALSIFY_SIMPLEX_REST_SHAPE"
	SupportLedgerUpgradesR3      = "CONDITIONAL_SUPPORT_SUCCESSFUL_LEDGER_TEST_WOULD_UPGRADE_BRANCH_TO_EXTERNAL_R3"
	SupportStrengthenedPartialR2 = "CONDITIONAL_SUPPORT_SIMPLEX_BRANCH_REMAINS_STRENGTHENED_PARTIAL_R2_WITH_EXTERNAL_R3_READY_PROTOCOL"
	SupportCYukawaIfCertified    = "CONDITIONAL_SUPPORT_CERTIFIED_SIMPLEX_SOURCE_WOULD_REDUCE_N_EFF_SEAL_DEPENDENCE"

	FailureNoNativeSimplex        = "FAILED_ROUTE_ONE_PLUS_THREE_SIMPLEX_NOT_NATIVE_WITHOUT_TRACE_READOUT_MAP"
	FailureProjectiveNotTheorem   = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_YUKAWA_TRACE_MAGNITUDE_THEOREM"
	FailureNoProjectiveReadout    = "FAILED_ROUTE_NO_PROJECTIVE_TO_TRACE_ATOM_READOUT_MAP"
	FailureK7NotTheorem           = "FAILED_ROUTE_K7_4_3_NOT_REST_TRACE_MAGNITUDE_THEOREM"
	FailureNoK7Readout            = "FAILED_ROUTE_NO_K7_POLARITY_TO_YUKAWA_REST_ATOM_MAP"
	FailureBoundaryAlphaNotAtom   = "FAILED_ROUTE_BOUNDARY_ALPHA_B_NOT_YUKAWA_REST_ATOM_THEOREM"
	FailureBoundaryColorNoSimplex = "FAILED_ROUTE_BOUNDARY_COLOR_COEFFICIENTS_DO_NOT_CONSTRUCT_REST_SIMPLEX_BY_THEMSELVES"
	FailureAbstractNoSector       = "FAILED_ROUTE_ABSTRACT_SIMPLEX_DOES_NOT_ASSIGN_STANDARD_MODEL_SECTORS"
	FailureExternalNotNative      = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	FailureShapeMustNotBeForced   = "FAILED_ROUTE_SIMPLEX_SHAPE_MUST_NOT_BE_FORCED_BY_ATOM_SELECTION"
	FailureCoefficientsFrozen     = "FAILED_ROUTE_COEFFICIENTS_MUST_NOT_BE_RETUNED_AFTER_LEDGER_INPUT"
	FailureNoUpdateCYukawa        = "FAILED_ROUTE_GATE819_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_CERTIFIED_TRACE_MAGNITUDE_MAP"
	FailureCHiggsLevelB           = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureTreeProxyNotPole       = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate819         = "FIREWALL_PRESERVED_GATE819_ONE_PLUS_THREE_REST_SIMPLEX_SOURCE_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2       float64
	AlphaB                       float64
	DeltaNBFN, NEffBFN, QRestBFN float64
	QSimplex, QResidual          float64
	NEffSimplex, NEffResidual    float64
	Verdicts, Supports, Failures []string
}

type SourceSeal struct {
	Name                         string
	Components                   []string
	TargetChain                  []string
	CurrentSupplied              bool
	Verdicts, Supports, Failures []string
}

type CandidateAudit struct {
	Name                         string
	SourceShape                  string
	SuppliesCarrier              bool
	SuppliesPositiveMeasure      bool
	SuppliesAlpha                bool
	SuppliesTraceReadout         bool
	SuppliesSectorAtoms          bool
	Verdicts, Supports, Failures []string
}

type ExternalFalsificationProtocol struct {
	RequiredObjects              []string
	PrimaryTests                 []string
	FailureCases                 []string
	CanUpgradeExternalR3         bool
	Verdicts, Supports, Failures []string
}

type NonCircularity struct {
	Forbidden                    []string
	FrozenCoefficients           []string
	Enforced                     bool
	Verdicts, Supports, Failures []string
}

type RStatus struct {
	Outcome                      string
	Level                        string
	NativeSourceFound            bool
	ExternalR3Ready              bool
	CanUpdateCYukawa             bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	NEffCandidate, CYukawaCandidate, CHiggsCandidate float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type BranchDecision struct {
	NextGate                     string
	Reason                       string
	Verdicts, Supports, Failures []string
}

type Firewalls struct {
	Enforced                                                                                                         bool
	SimplexNotNative, ProjectiveNotYukawa, K7NotYukawa, BoundaryAlphaNotYukawa, AbstractNotSector, ExternalNotNative bool
	NoCYukawaUpdate, CHiggsLevelB, TreeProxyNotPole                                                                  bool
	Verdict                                                                                                          string
}

type Analysis struct {
	Ledger      Ledger
	SourceSeal  SourceSeal
	Candidates  []CandidateAudit
	Protocol    ExternalFalsificationProtocol
	NonCircular NonCircularity
	Status      RStatus
	Impact      Impact
	Branch      BranchDecision
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func DeltaBFN(s float64) float64           { return C1NineFive*s + C2Six*M2(s) }
func NEffBFN(s float64) float64            { return 3.0 + DeltaBFN(s) }
func QSimplex(t float64) float64           { return t*t + math.Pow(1.0-t, 2)/3.0 }
func Beta(alpha, q float64) float64        { return 3.0 * alpha * alpha * q }
func NEffFrom(alpha, q float64) float64    { return 3.0 * math.Pow(1.0+alpha, 2) / (1.0 + Beta(alpha, q)) }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	delta := DeltaBFN(SBoundary)
	nEffBFN := 3.0 + delta
	qRest := 1.0 / nEffBFN
	qSimplex := QSimplex(alpha)
	nEffSimplex := NEffFrom(alpha, qSimplex)
	if math.Abs(delta-6.0*alpha) > 1e-18 {
		return Analysis{}, fmt.Errorf("Delta_N_BFN != 6 alpha_B: %.18g vs %.18g", delta, 6*alpha)
	}
	ledger := Ledger{NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2, AlphaB: alpha, DeltaNBFN: delta, NEffBFN: nEffBFN, QRestBFN: qRest, QSimplex: qSimplex, QResidual: qSimplex - qRest, NEffSimplex: nEffSimplex, NEffResidual: nEffSimplex - nEffBFN, Verdicts: []string{StatusGate818Inherited}, Supports: []string{SupportSharpestSimplex, SupportBoundaryAlphaNatural}, Failures: []string{FailureNoNativeSimplex}}

	seal := SourceSeal{Name: "OnePlusThreeRestSimplexSourceSeal", Components: []string{"rest carrier R_rest", "distinguished rest line L_1", "triplet rest chamber R_3", "boundary dust weight alpha_B", "normalized simplex map", "trace-magnitude readout", "positive rest spectrum construction", "sector/atom validation rule", "scale/scheme convention", "noncircularity proof"}, TargetChain: []string{"boundary data s,p -> alpha_B", "one distinguished rest line + triplet chamber", "w_rest(alpha_B)", "q_rest", "beta", "N_eff", "C_Yukawa"}, CurrentSupplied: false, Verdicts: []string{StatusSourceSealDefined}, Supports: []string{SupportSharpestSimplex}, Failures: []string{FailureNoNativeSimplex}}

	candidates := []CandidateAudit{
		{Name: "Candidate A — Fock/projective 1+3 selector", SourceShape: "4 = 1 + 3; native projective/Fock resonance", SuppliesCarrier: true, SuppliesPositiveMeasure: false, SuppliesAlpha: false, SuppliesTraceReadout: false, SuppliesSectorAtoms: false, Verdicts: []string{StatusProjectiveAudited}, Supports: []string{SupportProjectiveRelevant}, Failures: []string{FailureProjectiveNotTheorem, FailureNoProjectiveReadout}},
		{Name: "Candidate B — K7 Hodge 4|3 polarity", SourceShape: "K7 = K7+ ⊕ K7-, dim K7+ = 4, dim K7- = 3", SuppliesCarrier: true, SuppliesPositiveMeasure: false, SuppliesAlpha: false, SuppliesTraceReadout: false, SuppliesSectorAtoms: false, Verdicts: []string{StatusK7Audited}, Failures: []string{FailureK7NotTheorem, FailureNoK7Readout}},
		{Name: "Candidate C — boundary alpha plus boundary-pair/color coefficients", SourceShape: "alpha_B = (3/10)s + p s^2; 6 = boundary-pair 2 × color 3", SuppliesCarrier: false, SuppliesPositiveMeasure: false, SuppliesAlpha: true, SuppliesTraceReadout: false, SuppliesSectorAtoms: false, Verdicts: []string{StatusBoundaryAlphaAudited, StatusBoundaryColorAudited}, Supports: []string{SupportBoundaryAlphaNatural}, Failures: []string{FailureBoundaryAlphaNotAtom, FailureBoundaryColorNoSimplex}},
		{Name: "Candidate D — external Yukawa trace ledger", SourceShape: "decomposed sector/atom data can test the simplex shape", SuppliesCarrier: false, SuppliesPositiveMeasure: true, SuppliesAlpha: true, SuppliesTraceReadout: true, SuppliesSectorAtoms: true, Verdicts: []string{StatusExternalProtocol}, Supports: []string{SupportExternalFalsifies, SupportLedgerUpgradesR3}, Failures: []string{FailureExternalNotNative}},
	}

	protocol := ExternalFalsificationProtocol{RequiredObjects: []string{"typed top-like Hermitian eigenvalue T", "a_top = 3T and b_top = 3T^2", "positive rest atoms r_j", "scale/scheme/color/neutrino conventions", "a_ext,b_ext,N_eff_ext validation"}, PrimaryTests: []string{"T1: alpha_ext ≈ alpha_B", "T2: q_rest_ext ≈ q_simplex(alpha_B)", "T3: sorted rest weights resemble [alpha_B,(1-alpha_B)/3,(1-alpha_B)/3,(1-alpha_B)/3]", "T4: c2_ext ≈ 6", "T5: N_eff_ext ≈ 3 + 6 alpha_B"}, FailureCases: []string{"F1: alpha_ext not near alpha_B", "F2: q_rest_ext outside simplex concentration band", "F3: rest weights do not show one tiny atom plus three comparable atoms", "F4: c2_ext not near 6", "F5: top selector chosen unnaturally", "F6: scale/scheme destroys relation"}, CanUpgradeExternalR3: true, Verdicts: []string{StatusExternalProtocol}, Supports: []string{SupportExternalFalsifies, SupportLedgerUpgradesR3}, Failures: []string{FailureExternalNotNative}}

	noncircular := NonCircularity{Forbidden: []string{"choose top channel to force alpha_ext ≈ alpha_B", "discard rest atoms to manufacture 1+3 shape", "retune c1=9/5 or c2=6", "solve alpha_B from N_eff", "use observed Higgs mass to source Yukawa atoms", "use Koide/FN/GJ patterns to invent missing atoms"}, FrozenCoefficients: []string{"c1=9/5", "c2=6", "alpha_B=(3/10)s+p s^2"}, Enforced: true, Verdicts: []string{StatusNonCircularity}, Failures: []string{FailureShapeMustNotBeForced, FailureCoefficientsFrozen}}

	status := RStatus{Outcome: "Outcome B/C — source-typed but not native; Level-B+ external R3-ready hypothesis", Level: "strengthened partial R2 with external R3 falsification protocol; not native R4", NativeSourceFound: false, ExternalR3Ready: true, CanUpdateCYukawa: false, Verdicts: []string{StatusRStatusDefined, StatusOutcomeRecorded}, Supports: []string{SupportStrengthenedPartialR2, SupportExternalFalsifies}, Failures: []string{FailureNoNativeSimplex, FailureExternalNotNative}}

	impact := Impact{NEffCandidate: nEffBFN, CYukawaCandidate: CYukawaFromNEff(nEffBFN), CHiggsCandidate: CHiggsFromNEff(nEffBFN), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactRecorded}, Supports: []string{SupportCYukawaIfCertified}, Failures: []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole}}

	branch := BranchDecision{NextGate: "Gate 820 — OnePlusThree RestSimplex External Ledger Checklist and Native Trace-Readout Obstruction Audit", Reason: "Freeze the simplex as a Level-B+ falsifiable pattern and specify the exact trace-ledger intake plus the remaining native readout obstruction.", Verdicts: []string{StatusBranchRecorded}, Supports: []string{SupportExternalFalsifies}, Failures: []string{FailureNoNativeSimplex}}

	firewalls := Firewalls{Enforced: true, SimplexNotNative: true, ProjectiveNotYukawa: true, K7NotYukawa: true, BoundaryAlphaNotYukawa: true, AbstractNotSector: true, ExternalNotNative: true, NoCYukawaUpdate: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate819}

	truth := "Gate 819 finds no native trace-readout source for the 1+3 rest simplex; it remains the sharpest rest-concentration candidate and becomes an external-ledger falsifiable Level-B+ hypothesis."
	final := "The 1+3 simplex is source-aligned with projective 1+3 and boundary alpha_B, but no current ASHA carrier maps it to Yukawa trace atoms. C_Yukawa stays frozen until an external ledger validates it or a native trace-readout map is constructed."

	return Analysis{Ledger: ledger, SourceSeal: seal, Candidates: candidates, Protocol: protocol, NonCircular: noncircular, Status: status, Impact: impact, Branch: branch, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func Statuses() []string {
	return []string{StatusGate818Inherited, StatusSourceSealDefined, StatusProjectiveAudited, StatusK7Audited, StatusBoundaryAlphaAudited, StatusBoundaryColorAudited, StatusExternalProtocol, StatusNonCircularity, StatusRStatusDefined, StatusImpactRecorded, StatusOutcomeRecorded, StatusBranchRecorded, StatusPhysicalFirewalls, SupportSharpestSimplex, SupportProjectiveRelevant, SupportBoundaryAlphaNatural, SupportExternalFalsifies, SupportLedgerUpgradesR3, SupportStrengthenedPartialR2, SupportCYukawaIfCertified, FailureNoNativeSimplex, FailureProjectiveNotTheorem, FailureNoProjectiveReadout, FailureK7NotTheorem, FailureNoK7Readout, FailureBoundaryAlphaNotAtom, FailureBoundaryColorNoSimplex, FailureAbstractNoSector, FailureExternalNotNative, FailureShapeMustNotBeForced, FailureCoefficientsFrozen, FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole, StatusFirewallGate819}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g Delta_BFN=%.16g N_eff_BFN=%.16g q_rest_B=%.16g q_simplex=%.16g q_R=%.16g N_eff_simplex=%.16g N_R=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.DeltaNBFN, a.NEffBFN, a.QRestBFN, a.QSimplex, a.QResidual, a.NEffSimplex, a.NEffResidual)
}

func FormatSourceSeal(a SourceSeal) string {
	return fmt.Sprintf("%s supplied=%t components=[%s] target=[%s]", a.Name, a.CurrentSupplied, strings.Join(a.Components, "; "), strings.Join(a.TargetChain, " -> "))
}

func FormatCandidates(rows []CandidateAudit) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s source=%s carrier=%t measure=%t alpha=%t traceReadout=%t sectors=%t", r.Name, r.SourceShape, r.SuppliesCarrier, r.SuppliesPositiveMeasure, r.SuppliesAlpha, r.SuppliesTraceReadout, r.SuppliesSectorAtoms))
	}
	return strings.Join(out, " | ")
}

func FormatProtocol(a ExternalFalsificationProtocol) string {
	return fmt.Sprintf("required=[%s] tests=[%s] failures=[%s] externalR3=%t", strings.Join(a.RequiredObjects, "; "), strings.Join(a.PrimaryTests, "; "), strings.Join(a.FailureCases, "; "), a.CanUpgradeExternalR3)
}

func FormatImpact(a Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", a.NEffCandidate, a.CYukawaCandidate, a.CHiggsCandidate, a.OfficialNEff, a.OfficialCYukawa, a.OfficialCHiggs)
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
