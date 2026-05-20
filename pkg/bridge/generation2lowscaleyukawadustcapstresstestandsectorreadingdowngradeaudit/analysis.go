// Package generation2lowscaleyukawadustcapstresstestandsectorreadingdowngradeaudit implements
// Gate 822: Low-Scale Yukawa Dust-Cap Stress Test and Sector-Reading Downgrade Audit.
//
// Gate 822 turns Gate 821's dust-capacity consequence into a kill-test: a literal
// low-scale sector reading of the 1+3 rest simplex permits exactly one sizeable
// non-top colored triplet, while every other colored or uncolored rest atom must
// fit inside the tiny dust capacity.
package generation2lowscaleyukawadustcapstresstestandsectorreadingdowngradeaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	AuditID = "GATE822-LOW-SCALE-YUKAWA-DUST-CAP-STRESS-TEST-SECTOR-READING-DOWNGRADE-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate821Inherited  = "PASS_GATE821_DUST_CAPACITY_AUDIT_INHERITED"
	StatusStressDefined     = "PASS_LOW_SCALE_YUKAWA_DUST_CAP_STRESS_TEST_DEFINED"
	StatusBottomStress      = "PASS_BOTTOM_BRANCH_STRESS_TEST_DEFINED"
	StatusCharmStress       = "PASS_CHARM_BRANCH_STRESS_TEST_DEFINED"
	StatusAbstractStress    = "PASS_ABSTRACT_COLORED_CHAMBER_STRESS_TEST_DEFINED"
	StatusKillSwitch        = "PASS_LOW_SCALE_KILL_SWITCH_DEFINED"
	StatusExternalLedgerReq = "PASS_EXTERNAL_RATIO_LEDGER_REQUIREMENTS_DEFINED"
	StatusTestProtocol      = "PASS_TEST_PROTOCOL_DEFINED"
	StatusClassification    = "PASS_OUTCOME_CLASSIFICATION_DEFINED"
	StatusNativeAudit       = "PASS_NATIVE_SOURCE_AUDIT_DEFINED"
	StatusImpactFirewall    = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusPhysicalFirewall  = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	SupportDustCapSharpest        = "CONDITIONAL_SUPPORT_DUST_CAPACITY_IS_THE_SHARPEST_FALSIFICATION_OF_LITERAL_SECTOR_SIMPLEX"
	SupportMatchInsufficient      = "CONDITIONAL_SUPPORT_BOTTOM_OR_CHARM_MATCH_IS_INSUFFICIENT_WITHOUT_DUST_SURVIVAL"
	SupportOnlyOneSizeableRest    = "CONDITIONAL_SUPPORT_LITERAL_ONE_PLUS_THREE_READING_PREDICTS_ONLY_ONE_SIZEABLE_REST_SECTOR"
	SupportLowScaleCanFalsify     = "CONDITIONAL_SUPPORT_LOW_SCALE_LEDGER_CAN_FALSIFY_DIRECT_SECTOR_ASSIGNMENT"
	SupportFailureDowngrades      = "CONDITIONAL_SUPPORT_FAILURE_DOWNGRADES_SIMPLEX_TO_AGGREGATE_CONCENTRATION_MODEL"
	SupportHighScaleEscape        = "CONDITIONAL_SUPPORT_HIGH_SCALE_LEDGER_REMAINS_SEPARATE_ESCAPE_BRANCH_IF_LOW_SCALE_FAILS"
	SupportStrengthenedPartialR2  = "CONDITIONAL_SUPPORT_EXPECTED_STATUS_IS_STRENGTHENED_PARTIAL_R2_WITH_DUST_CAP_STRESS_PROTOCOL"
	SupportExternalProtocolFrozen = "CONDITIONAL_SUPPORT_EXTERNAL_RATIO_LEDGER_CAN_RUN_BOTTOM_CHARM_ABSTRACT_OR_FAILURE_BRANCHES"

	FailureBottomMatchAlone     = "FAILED_ROUTE_BOTTOM_MATCH_ALONE_DOES_NOT_PROVE_SECTOR_READING"
	FailureCharmMatchAlone      = "FAILED_ROUTE_CHARM_MATCH_ALONE_DOES_NOT_PROVE_SECTOR_READING"
	FailureExtraColoredAboveCap = "FAILED_ROUTE_ANY_EXTRA_COLORED_SECTOR_ABOVE_ALPHA_B_DUST_CAP_FALSIFIES_LITERAL_SIMPLEX_AT_THAT_SCALE"
	FailureUncoloredAboveCap    = "FAILED_ROUTE_ANY_UNCOLORED_SECTOR_ABOVE_SQRT3_ALPHA_B_DUST_CAP_FALSIFIES_TINY_DUST_READING_UNLESS_EXCLUDED"
	FailureExternalNotNative    = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	FailureLowScaleNotAggregate = "FAILED_ROUTE_LOW_SCALE_FAILURE_DOES_NOT_KILL_AGGREGATE_SIMPLEX_CLOSURE"
	FailureHighScaleNeedsRG     = "FAILED_ROUTE_HIGH_SCALE_ESCAPE_REQUIRES_RG_AND_MULTISCALE_LEDGER"
	FailureNoUpdateCYukawa      = "FAILED_ROUTE_GATE822_DOES_NOT_UPDATE_C_YUKAWA"
	FailureCHiggsLevelB         = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureProjectiveNotTheorem = "FAILED_ROUTE_PROJECTIVE_ONE_PLUS_THREE_NOT_TRACE_ATOM_THEOREM"
	FailureK7NotTraceAtom       = "FAILED_ROUTE_K7_4_3_NOT_TRACE_ATOM_THEOREM"
	FailureAlphaNotSector       = "FAILED_ROUTE_BOUNDARY_ALPHA_B_NOT_SECTOR_ASSIGNMENT_THEOREM"
	FailureBFNNotOperator       = "FAILED_ROUTE_BOUNDARY_FN_PACKAGE_NOT_YUKAWA_OPERATOR_THEOREM"
	FailureD4NotHierarchy       = "FAILED_ROUTE_D4_TRIALITY_NOT_REST_ATOM_HIERARCHY_THEOREM"
	FailureTreeProxyNotPole     = "FAILED_ROUTE_TREE_PROXY_NOT_POLE_MASS"
	StatusFirewallGate822       = "FIREWALL_PRESERVED_GATE822_LOW_SCALE_YUKAWA_DUST_CAP_STRESS_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2                      float64
	AlphaB, BOverT, SqrtBOverT                  float64
	DustOverT, SqrtDustOverT                    float64
	ExtraColoredTraceCap, ExtraColoredYukawaCap float64
	UncoloredTraceCap, UncoloredYukawaCap       float64
	TotalRestOverT, TripletTraceOverT           float64
	NEffSimplex, CYukawaSimplex, CHiggsSimplex  float64
	OfficialCYukawa, OfficialCHiggs             float64
	Verdicts, Supports, Failures                []string
}

type StressRule struct {
	Name                         string
	Logic                        []string
	FailureCondition             []string
	Verdicts, Supports, Failures []string
}

type ExternalRatioLedgerRequirement struct {
	Objects                      []string
	Forbidden                    []string
	Verdicts, Supports, Failures []string
}

type TestProtocol struct {
	Tests                        []string
	KillSwitch                   []string
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
	Enforced                                                                                                  bool
	BottomMatchAlone, CharmMatchAlone, ExtraColoredCap, UncoloredCap, ExternalNotNative, LowScaleNotAggregate bool
	HighScaleNeedsRG, NoCYukawaUpdate, CHiggsLevelB, TreeProxyNotPole                                         bool
	Verdict                                                                                                   string
}

type Analysis struct {
	Ledger      Ledger
	StressRules []StressRule
	Requirement ExternalRatioLedgerRequirement
	Protocol    TestProtocol
	Native      []NativeSourceAudit
	Status      Status
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

type RatioLedger struct {
	Scale, Scheme, Normalization, TopSelector, NeutrinoConvention string
	ColoredYukawaRatios                                           map[string]float64 // y_f / y_t for colored non-top sectors
	UncoloredYukawaRatios                                         map[string]float64 // y_l / y_t for uncolored sectors
}

type StressResult struct {
	Selected                     string
	SelectedRatio                float64
	LargeTripletMatch            bool
	ColoredDustOK                bool
	UncoloredDustOK              bool
	TotalRestTraceOverT          float64
	TotalRestTraceMatch          bool
	RestConcentration            float64
	RestConcentrationMatch       bool
	LiteralSectorReadingSurvives bool
	FailureReasons               []string
}

func M2(s float64) float64                 { return PBoundary * s * s }
func AlphaB(s float64) float64             { return (3.0/10.0)*s + M2(s) }
func BOverT(alpha float64) float64         { return alpha * (1.0 - alpha) }
func DustOverT(alpha float64) float64      { return 3.0 * alpha * alpha }
func NEffSimplex(alpha float64) float64    { return 3.0 + 6.0*alpha }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }
func QSimplex(alpha float64) float64       { return alpha*alpha + math.Pow(1.0-alpha, 2.0)/3.0 }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	bOverT := BOverT(alpha)
	dustOverT := DustOverT(alpha)
	if alpha <= 0 || bOverT <= 0 || dustOverT <= 0 {
		return Analysis{}, fmt.Errorf("invalid inherited positive scales: alpha=%g B/T=%g D/T=%g", alpha, bOverT, dustOverT)
	}
	nEffSimplex := NEffSimplex(alpha)
	ledger := Ledger{
		NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2,
		AlphaB: alpha, BOverT: bOverT, SqrtBOverT: math.Sqrt(bOverT),
		DustOverT: dustOverT, SqrtDustOverT: math.Sqrt(dustOverT),
		ExtraColoredTraceCap: alpha * alpha, ExtraColoredYukawaCap: alpha,
		UncoloredTraceCap: dustOverT, UncoloredYukawaCap: math.Sqrt(dustOverT),
		TotalRestOverT: 3.0 * alpha, TripletTraceOverT: 3.0 * bOverT,
		NEffSimplex: nEffSimplex, CYukawaSimplex: CYukawaFromNEff(nEffSimplex), CHiggsSimplex: CHiggsFromNEff(nEffSimplex),
		OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		Verdicts: []string{StatusGate821Inherited, StatusStressDefined, StatusKillSwitch},
		Supports: []string{SupportDustCapSharpest, SupportOnlyOneSizeableRest, SupportLowScaleCanFalsify},
		Failures: []string{FailureBottomMatchAlone, FailureCharmMatchAlone, FailureExtraColoredAboveCap, FailureUncoloredAboveCap},
	}

	stressRules := []StressRule{
		{Name: "bottom-branch stress test", Logic: []string{"if h_b/T ≈ alpha_B(1-alpha_B), bottom only survives first test", "charm, strange, up, down must satisfy y_f/y_t <= alpha_B", "tau, muon, electron must satisfy y_l/y_t <= sqrt(3) alpha_B unless excluded"}, FailureCondition: []string{"if charm, tau, or any other rest atom exceeds dust, bottom-sector reading fails even if bottom matches"}, Verdicts: []string{StatusBottomStress}, Supports: []string{SupportMatchInsufficient}, Failures: []string{FailureBottomMatchAlone, FailureExtraColoredAboveCap, FailureUncoloredAboveCap}},
		{Name: "charm-branch stress test", Logic: []string{"if h_c/T ≈ alpha_B(1-alpha_B), charm only survives first test", "bottom, strange, up, down must satisfy y_f/y_t <= alpha_B", "uncolored atoms must satisfy y_l/y_t <= sqrt(3) alpha_B unless excluded"}, FailureCondition: []string{"if bottom exceeds colored dust, charm-sector reading fails"}, Verdicts: []string{StatusCharmStress}, Supports: []string{SupportMatchInsufficient}, Failures: []string{FailureCharmMatchAlone, FailureExtraColoredAboveCap, FailureUncoloredAboveCap}},
		{Name: "abstract colored chamber stress test", Logic: []string{"one and only one colored triplet may be large", "all remaining colored and uncolored atoms must fit dust", "multiple sizeable rest sectors downgrade literal sector reading"}, FailureCondition: []string{"if ledger contains multiple sizeable rest sectors, literal 1+3 chamber reading fails"}, Verdicts: []string{StatusAbstractStress}, Supports: []string{SupportLowScaleCanFalsify}, Failures: []string{FailureExtraColoredAboveCap, FailureUncoloredAboveCap}},
	}

	requirement := ExternalRatioLedgerRequirement{
		Objects:   []string{"scale_mu", "scheme", "normalization", "top selector T", "y_b/y_t", "y_c/y_t", "y_s/y_t", "y_u/y_t", "y_d/y_t", "y_tau/y_t", "y_mu/y_t", "y_e/y_t", "neutrino convention", "uncertainties"},
		Forbidden: []string{"infer ratios from N_eff", "infer ratios from C_Higgs", "infer ratios from lambda_runtime_eff", "infer ratios from m_H_tree_proxy", "infer ratios from observed Higgs mass", "infer ratios from boundary-FN closure"},
		Verdicts:  []string{StatusExternalLedgerReq},
		Supports:  []string{SupportExternalProtocolFrozen},
		Failures:  []string{FailureExternalNotNative},
	}

	protocol := TestProtocol{
		Tests:      []string{"T1: identify whether any colored rest atom matches sqrt(B/T) ≈ 0.01969125", "T2: verify every non-selected colored atom satisfies y_f/y_t <= alpha_B", "T3: verify every uncolored atom satisfies y_l/y_t <= sqrt(3) alpha_B unless excluded", "T4: verify total rest trace a_rest/T ≈ 3 alpha_B", "T5: verify rest concentration q_rest ≈ q_simplex(alpha_B)", "T6: verify no coefficient, top selector, scale, or sector convention is retuned after import"},
		KillSwitch: []string{"if any non-selected colored ratio y_f/y_t >> alpha_B, literal low-scale sector simplex is falsified", "if any uncolored ratio y_l/y_t >> sqrt(3) alpha_B, tiny-dust reading is falsified unless excluded", "low-scale failure downgrades sector map but not aggregate concentration closure", "high-scale escape requires RG and multi-scale ledger"},
		CanFalsify: true,
		Verdicts:   []string{StatusTestProtocol, StatusKillSwitch},
		Supports:   []string{SupportDustCapSharpest, SupportLowScaleCanFalsify, SupportFailureDowngrades, SupportHighScaleEscape},
		Failures:   []string{FailureExtraColoredAboveCap, FailureUncoloredAboveCap, FailureLowScaleNotAggregate, FailureHighScaleNeedsRG},
	}

	native := []NativeSourceAudit{
		{Lane: "finite spectral triple", Supplies: []string{"edge templates", "color multiplicity"}, DoesNotSupply: []string{"low-scale Yukawa ratios", "colored rest hierarchy"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureExternalNotNative}},
		{Lane: "projective/Fock 1+3", Supplies: []string{"structural simplex shape"}, DoesNotSupply: []string{"bottom/charm/tau identity", "trace atoms"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureProjectiveNotTheorem}},
		{Lane: "K7 4|3", Supplies: []string{"carrier resonance"}, DoesNotSupply: []string{"trace atom values", "sector hierarchy"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureK7NotTraceAtom}},
		{Lane: "boundary alpha_B", Supplies: []string{"small rest-weight candidate"}, DoesNotSupply: []string{"bottom, charm, tau, or dust assignment"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureAlphaNotSector}},
		{Lane: "Boundary-FN", Supplies: []string{"aggregate closure candidate"}, DoesNotSupply: []string{"sector hierarchy", "Yukawa operator theorem"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureBFNNotOperator}},
		{Lane: "Georgi-Jarlskog", Supplies: []string{"secondary high-scale branch after bottom-like branch survives dust test"}, DoesNotSupply: []string{"low-scale bottom-sector proof", "dust-cap survival"}, Verdicts: []string{StatusNativeAudit}, Supports: []string{SupportHighScaleEscape}, Failures: []string{FailureHighScaleNeedsRG}},
		{Lane: "D4/triality", Supplies: []string{"airlocked structural search geometry"}, DoesNotSupply: []string{"rest atom hierarchy", "low-scale ratios"}, Verdicts: []string{StatusNativeAudit}, Failures: []string{FailureD4NotHierarchy}},
	}

	status := Status{Outcome: "Outcome C — no external ratio ledger active", Level: "strengthened partial R2; no sector-level decision; not external R3; not native R4", NativeSourceFound: false, ExternalLedgerSupplied: false, CanUpdateCYukawa: false, Verdicts: []string{StatusClassification}, Supports: []string{SupportStrengthenedPartialR2}, Failures: []string{FailureExternalNotNative, FailureNoUpdateCYukawa}}
	impact := Impact{CandidateNEff: nEffSimplex, CandidateCYukawa: CYukawaFromNEff(nEffSimplex), CandidateCHiggs: CHiggsFromNEff(nEffSimplex), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactFirewall}, Failures: []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureTreeProxyNotPole}}
	firewalls := Firewalls{Enforced: true, BottomMatchAlone: true, CharmMatchAlone: true, ExtraColoredCap: true, UncoloredCap: true, ExternalNotNative: true, LowScaleNotAggregate: true, HighScaleNeedsRG: true, NoCYukawaUpdate: true, CHiggsLevelB: true, TreeProxyNotPole: true, Verdict: StatusFirewallGate822}

	truth := "Gate 822 makes Gate 821 a kill-test: a low-scale sector-literal 1+3 simplex permits one sizeable non-top colored triplet and almost no other rest mass."
	final := "Without a convention-locked low-scale Yukawa ratio ledger, no bottom/charm/failure decision is made; the branch remains strengthened partial R2 with strict dust-cap falsification conditions and C_Yukawa frozen."
	return Analysis{Ledger: ledger, StressRules: stressRules, Requirement: requirement, Protocol: protocol, Native: native, Status: status, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

func StressTestLiteralSector(l RatioLedger, selected string, relativeTolerance float64) StressResult {
	alpha := AlphaB(SBoundary)
	largeTarget := math.Sqrt(BOverT(alpha))
	coloredCap := alpha
	uncoloredCap := math.Sqrt(DustOverT(alpha))
	if relativeTolerance <= 0 {
		relativeTolerance = 0.05
	}
	res := StressResult{Selected: selected}
	selectedRatio, ok := l.ColoredYukawaRatios[selected]
	if !ok {
		res.FailureReasons = append(res.FailureReasons, "selected colored sector missing")
		return res
	}
	res.SelectedRatio = selectedRatio
	res.LargeTripletMatch = relativeClose(selectedRatio, largeTarget, relativeTolerance)
	if !res.LargeTripletMatch {
		res.FailureReasons = append(res.FailureReasons, "selected colored sector does not match large triplet target")
	}
	res.ColoredDustOK = true
	for name, r := range l.ColoredYukawaRatios {
		if name == selected {
			continue
		}
		if r > coloredCap {
			res.ColoredDustOK = false
			res.FailureReasons = append(res.FailureReasons, fmt.Sprintf("colored sector %s exceeds alpha_B dust cap", name))
		}
	}
	res.UncoloredDustOK = true
	for name, r := range l.UncoloredYukawaRatios {
		if r > uncoloredCap {
			res.UncoloredDustOK = false
			res.FailureReasons = append(res.FailureReasons, fmt.Sprintf("uncolored sector %s exceeds sqrt(3) alpha_B dust cap", name))
		}
	}
	res.TotalRestTraceOverT = totalRestTraceOverT(l)
	res.TotalRestTraceMatch = relativeClose(res.TotalRestTraceOverT, 3.0*alpha, relativeTolerance)
	if !res.TotalRestTraceMatch {
		res.FailureReasons = append(res.FailureReasons, "total rest trace does not match 3 alpha_B")
	}
	res.RestConcentration = restConcentration(l)
	res.RestConcentrationMatch = relativeClose(res.RestConcentration, QSimplex(alpha), relativeTolerance)
	if !res.RestConcentrationMatch {
		res.FailureReasons = append(res.FailureReasons, "rest concentration does not match q_simplex(alpha_B)")
	}
	res.LiteralSectorReadingSurvives = res.LargeTripletMatch && res.ColoredDustOK && res.UncoloredDustOK && res.TotalRestTraceMatch && res.RestConcentrationMatch
	return res
}

func ratioKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func totalRestTraceOverT(l RatioLedger) float64 {
	total := 0.0
	for _, r := range l.ColoredYukawaRatios {
		total += 3.0 * r * r
	}
	for _, r := range l.UncoloredYukawaRatios {
		total += r * r
	}
	return total
}

func restConcentration(l RatioLedger) float64 {
	a := totalRestTraceOverT(l)
	if a == 0 {
		return 0
	}
	b := 0.0
	for _, r := range l.ColoredYukawaRatios {
		h := r * r
		b += 3.0 * h * h
	}
	for _, r := range l.UncoloredYukawaRatios {
		h := r * r
		b += h * h
	}
	return b / (a * a)
}

func relativeClose(got, want, tol float64) bool {
	if want == 0 {
		return math.Abs(got) <= tol
	}
	return math.Abs(got-want)/math.Abs(want) <= tol
}

func Statuses() []string {
	return []string{StatusGate821Inherited, StatusStressDefined, StatusBottomStress, StatusCharmStress, StatusAbstractStress, StatusKillSwitch, StatusExternalLedgerReq, StatusTestProtocol, StatusClassification, StatusNativeAudit, StatusImpactFirewall, StatusPhysicalFirewall, SupportDustCapSharpest, SupportMatchInsufficient, SupportOnlyOneSizeableRest, SupportLowScaleCanFalsify, SupportFailureDowngrades, SupportHighScaleEscape, SupportStrengthenedPartialR2, SupportExternalProtocolFrozen, FailureBottomMatchAlone, FailureCharmMatchAlone, FailureExtraColoredAboveCap, FailureUncoloredAboveCap, FailureExternalNotNative, FailureLowScaleNotAggregate, FailureHighScaleNeedsRG, FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureProjectiveNotTheorem, FailureK7NotTraceAtom, FailureAlphaNotSector, FailureBFNNotOperator, FailureD4NotHierarchy, FailureTreeProxyNotPole, StatusFirewallGate822}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g B/T=%.16g sqrt(B/T)=%.16g D/T=%.16g sqrt(D/T)=%.16g extra_colored_C/T_cap=%.16g extra_colored_y/y_t_cap=%.16g uncolored_L/T_cap=%.16g uncolored_y/y_t_cap=%.16g a_rest/T=%.16g 3B/T=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.BOverT, a.SqrtBOverT, a.DustOverT, a.SqrtDustOverT, a.ExtraColoredTraceCap, a.ExtraColoredYukawaCap, a.UncoloredTraceCap, a.UncoloredYukawaCap, a.TotalRestOverT, a.TripletTraceOverT)
}

func FormatStressRules(rows []StressRule) string {
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s logic=[%s] fail=[%s]", r.Name, strings.Join(r.Logic, "; "), strings.Join(r.FailureCondition, "; ")))
	}
	return strings.Join(out, " | ")
}

func FormatRequirement(r ExternalRatioLedgerRequirement) string {
	return fmt.Sprintf("objects=[%s] forbidden=[%s]", strings.Join(r.Objects, "; "), strings.Join(r.Forbidden, "; "))
}

func FormatProtocol(a TestProtocol) string {
	return fmt.Sprintf("tests=[%s] killSwitch=[%s] canFalsify=%t", strings.Join(a.Tests, "; "), strings.Join(a.KillSwitch, "; "), a.CanFalsify)
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
