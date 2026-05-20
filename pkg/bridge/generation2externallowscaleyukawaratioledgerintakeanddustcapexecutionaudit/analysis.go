// Package generation2externallowscaleyukawaratioledgerintakeanddustcapexecutionaudit implements
// Gate 823: External Low-Scale Yukawa Ratio Ledger Intake and Dust-Cap Execution Audit.
//
// Gate 823 executes Gate 822's knife without adding new geometry: it searches the
// active ASHA ledger for a convention-locked low-scale Yukawa ratio ledger. If the
// ledger is absent or convention-incomplete, it returns DATA_REQUIRED and freezes
// the literal low-scale sector assignment. If a valid ledger is supplied later,
// this package provides the dust-cap execution protocol and quantitative margins.
package generation2externallowscaleyukawaratioledgerintakeanddustcapexecutionaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

const (
	AuditID = "GATE823-EXTERNAL-LOW-SCALE-YUKAWA-RATIO-LEDGER-INTAKE-DUST-CAP-EXECUTION-AUDIT"

	NEff      = 3.0023273474722147
	DeltaN    = NEff - 3.0
	SBoundary = 0.0012924448188162962
	PBoundary = 7.0 / 72.0
	CHistory  = 1.038025177923625
	CYukawa   = 0.9992248188812008
	CHiggs    = 1.0372205204048603

	StatusGate822Inherited       = "PASS_GATE822_DUST_CAP_STRESS_PROTOCOL_INHERITED"
	StatusSearchDefined          = "PASS_EXTERNAL_LOW_SCALE_YUKAWA_RATIO_LEDGER_SEARCH_DEFINED"
	StatusConventionReq          = "PASS_LEDGER_CONVENTION_REQUIREMENTS_DEFINED"
	StatusBottomExecution        = "PASS_BOTTOM_BRANCH_EXECUTION_TEST_DEFINED"
	StatusCharmExecution         = "PASS_CHARM_BRANCH_EXECUTION_TEST_DEFINED"
	StatusAbstractExecution      = "PASS_ABSTRACT_COLORED_CHAMBER_EXECUTION_TEST_DEFINED"
	StatusMarginsDefined         = "PASS_VIOLATION_MARGIN_DIAGNOSTICS_DEFINED"
	StatusDowngradeRule          = "PASS_LITERAL_SECTOR_DOWNGRADE_RULE_DEFINED"
	StatusHighScaleFirewall      = "PASS_HIGH_SCALE_ESCAPE_FIREWALL_DEFINED"
	StatusImpactFirewall         = "PASS_C_YUKAWA_AND_C_HIGGS_FIREWALL_PRESERVED"
	StatusOutcomeClassification  = "PASS_OUTCOME_CLASSIFICATION_DEFINED"
	StatusPhysicalFirewalls      = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusProtocolReady          = "PASS_DUST_CAP_PROTOCOL_READY"
	StatusDataRequiredLedger     = "DATA_REQUIRED_EXTERNAL_LOW_SCALE_YUKAWA_RATIO_LEDGER"
	StatusDataRequiredConvention = "DATA_REQUIRED_CONVENTION_LOCKED_YUKAWA_RATIO_LEDGER"

	SupportExecutionGate              = "CONDITIONAL_SUPPORT_GATE823_IS_EXECUTION_GATE_NOT_NEW_GEOMETRY"
	SupportCanDecideIfLedgerExists    = "CONDITIONAL_SUPPORT_DUST_CAP_TEST_CAN_DECIDE_LITERAL_LOW_SCALE_SECTOR_READING_IF_LEDGER_EXISTS"
	SupportMatchNeedsDustSurvival     = "CONDITIONAL_SUPPORT_BOTTOM_OR_CHARM_MATCH_REQUIRES_FULL_DUST_SURVIVAL"
	SupportFailureDowngrades          = "CONDITIONAL_SUPPORT_FAILURE_DOWNGRADES_SIMPLEX_TO_AGGREGATE_OR_HIGH_SCALE_HYPOTHESIS"
	SupportExternalR3RequiresFullPass = "CONDITIONAL_SUPPORT_EXTERNAL_R3_STATUS_REQUIRES_CONVENTION_LOCKED_LEDGER_PASS"
	SupportHighScaleEscapeSeparate    = "CONDITIONAL_SUPPORT_HIGH_SCALE_ESCAPE_REQUIRES_SEPARATE_RG_TRANSPORT_LEDGER"

	FailureNoLedger                  = "FAILED_ROUTE_NO_SECTOR_DECISION_WITHOUT_EXTERNAL_LOW_SCALE_YUKAWA_RATIO_LEDGER"
	FailureNoLedgerFound             = "FAILED_ROUTE_NO_EXTERNAL_LOW_SCALE_YUKAWA_RATIO_LEDGER_FOUND"
	FailureIncompleteLedger          = "FAILED_ROUTE_INCOMPLETE_LEDGER_CANNOT_EXECUTE_DUST_CAP_TEST"
	FailureConventionsIncomplete     = "FAILED_ROUTE_LOW_SCALE_LEDGER_CONVENTIONS_INCOMPLETE"
	FailureBottomMatchAlone          = "FAILED_ROUTE_BOTTOM_MATCH_ALONE_DOES_NOT_SURVIVE_GATE823"
	FailureCharmMatchAlone           = "FAILED_ROUTE_CHARM_MATCH_ALONE_DOES_NOT_SURVIVE_GATE823"
	FailureDustOverflow              = "FAILED_ROUTE_DUST_OVERFLOW_FALSIFIES_LITERAL_LOW_SCALE_SECTOR_READING"
	FailureExtraColoredAboveCap      = "FAILED_ROUTE_EXTRA_COLORED_SECTOR_ABOVE_ALPHA_B_CAP_CANNOT_BE_IGNORED"
	FailureUncoloredAboveCap         = "FAILED_ROUTE_UNCOLORED_SECTOR_ABOVE_SQRT3_ALPHA_B_CAP_CANNOT_BE_IGNORED_UNLESS_EXCLUDED_BY_EXPLICIT_CONVENTION"
	FailureHighScaleNeedsRG          = "FAILED_ROUTE_LOW_SCALE_FAILURE_CANNOT_BE_RESCUED_BY_HIGH_SCALE_LANGUAGE_WITHOUT_RG_LEDGER"
	FailureNoUpdateCYukawa           = "FAILED_ROUTE_GATE823_DOES_NOT_UPDATE_C_YUKAWA_WITHOUT_FULL_PASS_OR_NATIVE_MAP"
	FailureCHiggsLevelB              = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B"
	FailureExternalNotNative         = "FAILED_ROUTE_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	FailureBottomBranchDustViolated  = "FAILED_ROUTE_BOTTOM_BRANCH_DUST_CAP_VIOLATED"
	FailureCharmBranchDustViolated   = "FAILED_ROUTE_CHARM_BRANCH_DUST_CAP_VIOLATED"
	FailureMultipleColoredAboveCap   = "FAILED_ROUTE_MULTIPLE_COLORED_REST_RATIOS_EXCEED_DUST_CAP"
	FailureAbstractUncoloredAboveCap = "FAILED_ROUTE_UNCOLORED_REST_RATIO_EXCEEDS_DUST_CAP"
	StatusFirewallGate823            = "FIREWALL_PRESERVED_GATE823_EXTERNAL_LOW_SCALE_YUKAWA_DUST_CAP_EXECUTION_BOUNDARY"
)

type Ledger struct {
	NEff, DeltaN, S, P, M2                           float64
	AlphaB, BOverT, SqrtBOverT                       float64
	ExtraColoredCap, ExtraColoredTraceCap            float64
	UncoloredCap, UncoloredTraceCap                  float64
	TotalRestTraceTarget                             float64
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialCYukawa, OfficialCHiggs                  float64
	Verdicts, Supports, Failures                     []string
}

type ExternalLowScaleYukawaRatioLedger struct {
	SourceLabel, ScaleMu, Scheme, Normalization, TopSelector string
	ColorConvention, NeutrinoConvention                      string
	ColoredRatios                                            map[string]float64 // y_f/y_t for b,c,s,u,d
	UncoloredRatios                                          map[string]float64 // y_l/y_t for tau,mu,e and optional neutrinos by convention
	Uncertainties                                            map[string]float64
}

type SearchResult struct {
	Found, ConventionLocked      bool
	SourcesSearched              []string
	MissingObjects               []string
	Ledger                       *ExternalLowScaleYukawaRatioLedger
	Verdicts, Supports, Failures []string
}

type BranchDefinition struct {
	Name                         string
	Selected                     string
	Tests                        []string
	Verdicts, Supports, Failures []string
}

type Margin struct {
	Name   string
	Ratio  float64
	Cap    float64
	Margin float64
	Kind   string
	Class  string
}

type ExecutionResult struct {
	Branch                       string
	Selected                     string
	LargeTarget                  float64
	SelectedRatio                float64
	LargeTargetMargin            float64
	LargeTripletMatched          bool
	ColoredMargins               []Margin
	UncoloredMargins             []Margin
	ColoredDustOK                bool
	UncoloredDustOK              bool
	LiteralSectorSurvives        bool
	Outcome                      string
	Reasons                      []string
	Verdicts, Supports, Failures []string
}

type TestProtocol struct {
	Tests                        []string
	Forbidden                    []string
	DowngradeRule                []string
	HighScaleEscape              []string
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
	DataRequired                 bool
	LiteralSectorFrozen          bool
	ExternalR3                   bool
	CanUpdateCYukawa             bool
	Verdicts, Supports, Failures []string
}

type Impact struct {
	CandidateNEff, CandidateCYukawa, CandidateCHiggs float64
	OfficialNEff, OfficialCYukawa, OfficialCHiggs    float64
	Verdicts, Supports, Failures                     []string
}

type Firewalls struct {
	Enforced                                                                                    bool
	NoInferenceWithoutLedger, NoIncompleteLedger, NoMatchAlone, DustOverflow, NoHighScaleEscape bool
	NoCYukawaUpdate, CHiggsLevelB, ExternalNotNative                                            bool
	Verdict                                                                                     string
}

type Analysis struct {
	Ledger    Ledger
	Search    SearchResult
	Branches  []BranchDefinition
	Protocol  TestProtocol
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
func NEffSimplex(alpha float64) float64    { return 3.0 + 6.0*alpha }
func CYukawaFromNEff(nEff float64) float64 { return 3.0 / nEff }
func CHiggsFromNEff(nEff float64) float64  { return CYukawaFromNEff(nEff) * CHistory }
func UncoloredCap(alpha float64) float64   { return math.Sqrt(3.0) * alpha }

func BuildDefault() (Analysis, error) {
	m2 := M2(SBoundary)
	alpha := AlphaB(SBoundary)
	bOverT := BOverT(alpha)
	if alpha <= 0 || bOverT <= 0 {
		return Analysis{}, fmt.Errorf("invalid inherited Gate 822 scales: alpha_B=%g B/T=%g", alpha, bOverT)
	}
	candidateNEff := NEffSimplex(alpha)
	ledger := Ledger{
		NEff: NEff, DeltaN: DeltaN, S: SBoundary, P: PBoundary, M2: m2,
		AlphaB: alpha, BOverT: bOverT, SqrtBOverT: math.Sqrt(bOverT),
		ExtraColoredCap: alpha, ExtraColoredTraceCap: alpha * alpha,
		UncoloredCap: UncoloredCap(alpha), UncoloredTraceCap: 3.0 * alpha * alpha,
		TotalRestTraceTarget: 3.0 * alpha,
		CandidateNEff:        candidateNEff, CandidateCYukawa: CYukawaFromNEff(candidateNEff), CandidateCHiggs: CHiggsFromNEff(candidateNEff),
		OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs,
		Verdicts: []string{StatusGate822Inherited},
		Supports: []string{SupportExecutionGate, SupportCanDecideIfLedgerExists, SupportMatchNeedsDustSurvival},
		Failures: []string{FailureNoLedger, FailureNoUpdateCYukawa, FailureCHiggsLevelB},
	}
	search := SearchActiveProjectLedger()
	branches := []BranchDefinition{
		{Name: "bottom branch", Selected: "b", Tests: []string{"r_b ≈ sqrt(alpha_B(1-alpha_B))", "r_c,r_s,r_u,r_d <= alpha_B", "r_tau,r_mu,r_e <= sqrt(3) alpha_B unless excluded"}, Verdicts: []string{StatusBottomExecution}, Supports: []string{SupportMatchNeedsDustSurvival}, Failures: []string{FailureBottomMatchAlone, FailureBottomBranchDustViolated}},
		{Name: "charm branch", Selected: "c", Tests: []string{"r_c ≈ sqrt(alpha_B(1-alpha_B))", "r_b,r_s,r_u,r_d <= alpha_B", "r_tau,r_mu,r_e <= sqrt(3) alpha_B unless excluded"}, Verdicts: []string{StatusCharmExecution}, Supports: []string{SupportMatchNeedsDustSurvival}, Failures: []string{FailureCharmMatchAlone, FailureCharmBranchDustViolated}},
		{Name: "abstract colored chamber", Selected: "largest colored non-top", Tests: []string{"sort r_b,r_c,r_s,r_u,r_d descending", "largest ≈ sqrt(B/T)", "all remaining colored <= alpha_B", "all uncolored <= sqrt(3) alpha_B"}, Verdicts: []string{StatusAbstractExecution}, Supports: []string{SupportCanDecideIfLedgerExists}, Failures: []string{FailureMultipleColoredAboveCap, FailureAbstractUncoloredAboveCap}},
	}
	protocol := TestProtocol{
		Tests:           []string{"T1 identify whether any colored rest atom matches sqrt(B/T)", "T2 every non-selected colored atom must satisfy y_f/y_t <= alpha_B", "T3 every included uncolored atom must satisfy y_l/y_t <= sqrt(3) alpha_B", "T4 total rest trace should match 3 alpha_B", "T5 rest concentration should match q_simplex(alpha_B)", "T6 no coefficient, selector, scale, or sector convention retuned after ledger import"},
		Forbidden:       []string{"infer ratios from N_eff, C_Higgs, Higgs mass, boundary-FN closure, Koide, FN, GJ, or symbolic patterns", "bottom or charm match alone", "silent exclusion of dust-overflow sectors", "high-scale rescue without RG ledger"},
		DowngradeRule:   []string{"if more than one sizeable rest sector exceeds dust, literal low-scale sector simplex fails", "remaining statuses: aggregate concentration, hidden rest-carrier, high-scale/RG diagnostic"},
		HighScaleEscape: []string{"requires multi-scale Yukawa ledger", "requires RG transport and threshold convention", "low-scale failure does not kill aggregate simplex but blocks direct low-scale sector assignment"},
		Verdicts:        []string{StatusSearchDefined, StatusConventionReq, StatusMarginsDefined, StatusDowngradeRule, StatusHighScaleFirewall},
		Supports:        []string{SupportFailureDowngrades, SupportHighScaleEscapeSeparate},
		Failures:        []string{FailureDustOverflow, FailureHighScaleNeedsRG},
	}
	native := []NativeSourceAudit{
		{Lane: "finite spectral triple", Supplies: []string{"edge templates", "color multiplicity"}, DoesNotSupply: []string{"low-scale Yukawa ratios"}, Failures: []string{FailureNoLedger}},
		{Lane: "projective/Fock 1+3", Supplies: []string{"structural simplex shape"}, DoesNotSupply: []string{"sector assignment", "low-scale ratios"}, Failures: []string{FailureNoLedger}},
		{Lane: "K7 4|3", Supplies: []string{"carrier resonance"}, DoesNotSupply: []string{"trace atoms"}, Failures: []string{FailureNoLedger}},
		{Lane: "boundary alpha_B", Supplies: []string{"small rest-weight candidate"}, DoesNotSupply: []string{"bottom/charm/tau/dust assignment"}, Failures: []string{FailureNoLedger}},
		{Lane: "Boundary-FN", Supplies: []string{"aggregate closure candidate"}, DoesNotSupply: []string{"sector hierarchy", "literal low-scale sector map"}, Failures: []string{FailureNoLedger}},
		{Lane: "Georgi-Jarlskog", Supplies: []string{"future high-scale diagnostic after bottom-like survival or RG ledger"}, DoesNotSupply: []string{"low-scale rescue without RG"}, Failures: []string{FailureHighScaleNeedsRG}},
		{Lane: "D4/triality", Supplies: []string{"airlocked structural search branch"}, DoesNotSupply: []string{"rest atom hierarchy"}, Failures: []string{FailureNoLedger}},
	}
	status := Status{Outcome: "Outcome A — convention-locked low-scale Yukawa ratio ledger absent", Level: "DATA_REQUIRED; literal sector assignment frozen; strengthened partial R2 only; not external R3; not native R4", DataRequired: true, LiteralSectorFrozen: true, ExternalR3: false, CanUpdateCYukawa: false, Verdicts: []string{StatusOutcomeClassification, StatusProtocolReady, StatusDataRequiredLedger}, Supports: []string{SupportExecutionGate}, Failures: []string{FailureNoLedger, FailureNoLedgerFound, FailureNoUpdateCYukawa}}
	if search.Found && !search.ConventionLocked {
		status = Status{Outcome: "Outcome F — low-scale ledger present but convention-incomplete", Level: "DATA_REQUIRED_CONVENTION_LOCKED_LEDGER; dust-cap test not executed", DataRequired: true, LiteralSectorFrozen: true, ExternalR3: false, CanUpdateCYukawa: false, Verdicts: []string{StatusOutcomeClassification, StatusDataRequiredConvention}, Failures: []string{FailureConventionsIncomplete, FailureIncompleteLedger, FailureNoUpdateCYukawa}}
	}
	impact := Impact{CandidateNEff: candidateNEff, CandidateCYukawa: CYukawaFromNEff(candidateNEff), CandidateCHiggs: CHiggsFromNEff(candidateNEff), OfficialNEff: NEff, OfficialCYukawa: CYukawa, OfficialCHiggs: CHiggs, Verdicts: []string{StatusImpactFirewall}, Failures: []string{FailureNoUpdateCYukawa, FailureCHiggsLevelB}}
	firewalls := Firewalls{Enforced: true, NoInferenceWithoutLedger: true, NoIncompleteLedger: true, NoMatchAlone: true, DustOverflow: true, NoHighScaleEscape: true, NoCYukawaUpdate: true, CHiggsLevelB: true, ExternalNotNative: true, Verdict: StatusFirewallGate823}
	truth := "Gate 823 executes the Gate 822 knife: if a convention-locked low-scale Yukawa ratio ledger is absent, the literal sector assignment is frozen as DATA_REQUIRED; if supplied, bottom/charm/abstract branches must survive the full dust cap."
	final := "No active convention-locked low-scale Yukawa ratio ledger is certified in the current ASHA ledger; Gate 823 returns DATA_REQUIRED and does not update C_Yukawa or C_Higgs."
	return Analysis{Ledger: ledger, Search: search, Branches: branches, Protocol: protocol, Native: native, Status: status, Impact: impact, Firewalls: firewalls, Truth: truth, Final: final}, nil
}

// SearchActiveProjectLedger records the current active-ledger status for Gate 823.
// No convention-locked ExternalLowScaleYukawaRatioLedger is part of the active
// typed ledger in the Gate 822 package; observed PDG-like mass files are not this
// seal because they do not expose the required scale/scheme/normalization/top
// selector/color/neutrino conventions as a ratio ledger.
func SearchActiveProjectLedger() SearchResult {
	return SearchResult{
		Found:            false,
		ConventionLocked: false,
		SourcesSearched:  []string{"active aggregate a,b,N_eff ledger", "finite spectral triple trace templates", "Boundary-FN candidate ledger", "project data/data/*.json observed ledgers checked as not convention-locked ratio seal"},
		MissingObjects:   []string{"source_label", "scale_mu", "scheme", "normalization", "top_selector", "color_convention", "neutrino_convention", "r_b,r_c,r_s,r_u,r_d", "r_tau,r_mu,r_e", "uncertainties"},
		Verdicts:         []string{StatusSearchDefined, StatusProtocolReady, StatusDataRequiredLedger},
		Supports:         []string{SupportExecutionGate},
		Failures:         []string{FailureNoLedgerFound, FailureNoLedger},
	}
}

func (l ExternalLowScaleYukawaRatioLedger) MissingConventionObjects() []string {
	var missing []string
	if strings.TrimSpace(l.SourceLabel) == "" {
		missing = append(missing, "source_label")
	}
	if strings.TrimSpace(l.ScaleMu) == "" {
		missing = append(missing, "scale_mu")
	}
	if strings.TrimSpace(l.Scheme) == "" {
		missing = append(missing, "scheme")
	}
	if strings.TrimSpace(l.Normalization) == "" {
		missing = append(missing, "normalization")
	}
	if strings.TrimSpace(l.TopSelector) == "" {
		missing = append(missing, "top_selector")
	}
	if strings.TrimSpace(l.ColorConvention) == "" {
		missing = append(missing, "color_convention")
	}
	if strings.TrimSpace(l.NeutrinoConvention) == "" {
		missing = append(missing, "neutrino_convention")
	}
	for _, k := range []string{"b", "c", "s", "u", "d"} {
		if _, ok := l.ColoredRatios[k]; !ok {
			missing = append(missing, "r_"+k)
		}
	}
	for _, k := range []string{"tau", "mu", "e"} {
		if _, ok := l.UncoloredRatios[k]; !ok {
			missing = append(missing, "r_"+k)
		}
	}
	return missing
}

func (l ExternalLowScaleYukawaRatioLedger) ConventionLocked() bool {
	return len(l.MissingConventionObjects()) == 0
}

func ExecuteBranch(l ExternalLowScaleYukawaRatioLedger, selected string, relativeTolerance float64) ExecutionResult {
	alpha := AlphaB(SBoundary)
	largeTarget := math.Sqrt(BOverT(alpha))
	coloredCap := alpha
	uncoloredCap := UncoloredCap(alpha)
	if relativeTolerance <= 0 {
		relativeTolerance = 0.05
	}
	res := ExecutionResult{Branch: selected + " branch", Selected: selected, LargeTarget: largeTarget}
	selectedRatio, ok := l.ColoredRatios[selected]
	if !ok {
		res.Reasons = append(res.Reasons, "selected colored ratio missing")
		res.Outcome = "ledger incomplete for selected branch"
		res.Failures = append(res.Failures, FailureIncompleteLedger)
		return res
	}
	res.SelectedRatio = selectedRatio
	res.LargeTargetMargin = selectedRatio / largeTarget
	res.LargeTripletMatched = relativeClose(selectedRatio, largeTarget, relativeTolerance)
	res.ColoredDustOK = true
	for _, k := range sortedKeys(l.ColoredRatios) {
		if k == selected {
			continue
		}
		m := classifyMargin(k, l.ColoredRatios[k], coloredCap, "colored", l.Uncertainties)
		res.ColoredMargins = append(res.ColoredMargins, m)
		if m.Class == "hard fail" {
			res.ColoredDustOK = false
		}
	}
	res.UncoloredDustOK = true
	for _, k := range sortedKeys(l.UncoloredRatios) {
		m := classifyMargin(k, l.UncoloredRatios[k], uncoloredCap, "uncolored", l.Uncertainties)
		res.UncoloredMargins = append(res.UncoloredMargins, m)
		if m.Class == "hard fail" {
			res.UncoloredDustOK = false
		}
	}
	if !res.LargeTripletMatched {
		res.Reasons = append(res.Reasons, "selected large triplet does not match target")
	}
	if !res.ColoredDustOK {
		res.Reasons = append(res.Reasons, "one or more non-selected colored sectors exceed alpha_B cap")
	}
	if !res.UncoloredDustOK {
		res.Reasons = append(res.Reasons, "one or more uncolored sectors exceed sqrt(3) alpha_B cap")
	}
	res.LiteralSectorSurvives = res.LargeTripletMatched && res.ColoredDustOK && res.UncoloredDustOK
	if res.LiteralSectorSurvives {
		res.Outcome = "Outcome B — one large colored sector matched and dust caps passed"
		res.Verdicts = []string{branchPassVerdict(selected)}
		res.Supports = []string{SupportExternalR3RequiresFullPass}
		return res
	}
	res.Outcome = "literal low-scale sector simplex rejected for this branch"
	res.Failures = []string{matchAloneFailure(selected), FailureDustOverflow}
	return res
}

func ExecuteAbstractColoredChamber(l ExternalLowScaleYukawaRatioLedger, relativeTolerance float64) ExecutionResult {
	alpha := AlphaB(SBoundary)
	largeTarget := math.Sqrt(BOverT(alpha))
	coloredCap := alpha
	uncoloredCap := UncoloredCap(alpha)
	if relativeTolerance <= 0 {
		relativeTolerance = 0.05
	}
	type pair struct {
		name  string
		ratio float64
	}
	pairs := make([]pair, 0, len(l.ColoredRatios))
	for k, v := range l.ColoredRatios {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].ratio > pairs[j].ratio })
	res := ExecutionResult{Branch: "abstract colored chamber", LargeTarget: largeTarget, ColoredDustOK: true, UncoloredDustOK: true}
	if len(pairs) == 0 {
		res.Outcome = "no colored ratios supplied"
		res.Failures = []string{FailureIncompleteLedger}
		return res
	}
	res.Selected = pairs[0].name
	res.SelectedRatio = pairs[0].ratio
	res.LargeTargetMargin = res.SelectedRatio / largeTarget
	res.LargeTripletMatched = relativeClose(res.SelectedRatio, largeTarget, relativeTolerance)
	for _, p := range pairs[1:] {
		m := classifyMargin(p.name, p.ratio, coloredCap, "colored", l.Uncertainties)
		res.ColoredMargins = append(res.ColoredMargins, m)
		if m.Class == "hard fail" {
			res.ColoredDustOK = false
		}
	}
	for _, k := range sortedKeys(l.UncoloredRatios) {
		m := classifyMargin(k, l.UncoloredRatios[k], uncoloredCap, "uncolored", l.Uncertainties)
		res.UncoloredMargins = append(res.UncoloredMargins, m)
		if m.Class == "hard fail" {
			res.UncoloredDustOK = false
		}
	}
	res.LiteralSectorSurvives = res.LargeTripletMatched && res.ColoredDustOK && res.UncoloredDustOK
	if res.LiteralSectorSurvives {
		res.Outcome = "Outcome B — abstract one-colored chamber survives dust caps"
		res.Verdicts = []string{"PASS_ABSTRACT_ONE_COLORED_CHAMBER_SURVIVES"}
	} else {
		res.Outcome = "Outcome D/E — abstract chamber fails or downgrades"
		res.Failures = []string{FailureMultipleColoredAboveCap, FailureAbstractUncoloredAboveCap}
	}
	return res
}

func classifyMargin(name string, ratio, cap float64, kind string, uncertainties map[string]float64) Margin {
	m := Margin{Name: name, Ratio: ratio, Cap: cap, Kind: kind}
	if cap > 0 {
		m.Margin = ratio / cap
	}
	unc := uncertainties[name]
	switch {
	case ratio <= cap || (unc > 0 && ratio-unc <= cap):
		m.Class = "soft pass"
	case m.Margin > 1.0:
		m.Class = "hard fail"
	default:
		m.Class = "ambiguous"
	}
	return m
}

func branchPassVerdict(selected string) string {
	if selected == "b" {
		return "PASS_BOTTOM_LARGE_TRIPLET_MATCHED"
	}
	if selected == "c" {
		return "PASS_CHARM_LARGE_TRIPLET_MATCHED"
	}
	return "PASS_ABSTRACT_ONE_COLORED_CHAMBER_SURVIVES"
}

func matchAloneFailure(selected string) string {
	if selected == "b" {
		return FailureBottomMatchAlone
	}
	if selected == "c" {
		return FailureCharmMatchAlone
	}
	return FailureDustOverflow
}

func sortedKeys(m map[string]float64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func relativeClose(got, want, tol float64) bool {
	if want == 0 {
		return math.Abs(got) <= tol
	}
	return math.Abs(got-want)/math.Abs(want) <= tol
}

func Statuses() []string {
	return []string{StatusGate822Inherited, StatusSearchDefined, StatusConventionReq, StatusBottomExecution, StatusCharmExecution, StatusAbstractExecution, StatusMarginsDefined, StatusDowngradeRule, StatusHighScaleFirewall, StatusImpactFirewall, StatusOutcomeClassification, StatusPhysicalFirewalls, StatusProtocolReady, StatusDataRequiredLedger, StatusDataRequiredConvention, SupportExecutionGate, SupportCanDecideIfLedgerExists, SupportMatchNeedsDustSurvival, SupportFailureDowngrades, SupportExternalR3RequiresFullPass, SupportHighScaleEscapeSeparate, FailureNoLedger, FailureNoLedgerFound, FailureIncompleteLedger, FailureConventionsIncomplete, FailureBottomMatchAlone, FailureCharmMatchAlone, FailureDustOverflow, FailureExtraColoredAboveCap, FailureUncoloredAboveCap, FailureHighScaleNeedsRG, FailureNoUpdateCYukawa, FailureCHiggsLevelB, FailureExternalNotNative, StatusFirewallGate823}
}

func FormatLedger(a Ledger) string {
	return fmt.Sprintf("N_eff=%.16g Delta_N=%.16g s=%.16g p=%.16g M2=%.16g alpha_B=%.16g sqrt(B/T)=%.16g extra_colored_y/y_t_cap=%.16g uncolored_y/y_t_cap=%.16g a_rest/T=%.16g candidate_Neff=%.16g candidate_CYukawa=%.16g candidate_CHiggs=%.16g official_CYukawa=%.16g official_CHiggs=%.16g", a.NEff, a.DeltaN, a.S, a.P, a.M2, a.AlphaB, a.SqrtBOverT, a.ExtraColoredCap, a.UncoloredCap, a.TotalRestTraceTarget, a.CandidateNEff, a.CandidateCYukawa, a.CandidateCHiggs, a.OfficialCYukawa, a.OfficialCHiggs)
}

func FormatSearch(s SearchResult) string {
	return fmt.Sprintf("found=%t conventionLocked=%t sources=[%s] missing=[%s]", s.Found, s.ConventionLocked, strings.Join(s.SourcesSearched, "; "), strings.Join(s.MissingObjects, "; "))
}

func FormatBranches(rows []BranchDefinition) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s selected=%s tests=[%s]", r.Name, r.Selected, strings.Join(r.Tests, "; ")))
	}
	return strings.Join(parts, " | ")
}

func FormatProtocol(p TestProtocol) string {
	return fmt.Sprintf("tests=[%s] forbidden=[%s] downgrade=[%s] highScale=[%s]", strings.Join(p.Tests, "; "), strings.Join(p.Forbidden, "; "), strings.Join(p.DowngradeRule, "; "), strings.Join(p.HighScaleEscape, "; "))
}

func FormatMargins(rows []Margin) string {
	parts := make([]string, 0, len(rows))
	for _, m := range rows {
		parts = append(parts, fmt.Sprintf("%s(%s): ratio=%.16g cap=%.16g margin=%.16g class=%s", m.Name, m.Kind, m.Ratio, m.Cap, m.Margin, m.Class))
	}
	return strings.Join(parts, " | ")
}

func FormatExecution(r ExecutionResult) string {
	return fmt.Sprintf("branch=%s selected=%s selectedRatio=%.16g target=%.16g targetMargin=%.16g largeMatch=%t coloredOK=%t uncoloredOK=%t survives=%t coloredMargins=[%s] uncoloredMargins=[%s] reasons=[%s] outcome=%s", r.Branch, r.Selected, r.SelectedRatio, r.LargeTarget, r.LargeTargetMargin, r.LargeTripletMatched, r.ColoredDustOK, r.UncoloredDustOK, r.LiteralSectorSurvives, FormatMargins(r.ColoredMargins), FormatMargins(r.UncoloredMargins), strings.Join(r.Reasons, "; "), r.Outcome)
}

func FormatNative(rows []NativeSourceAudit) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s supplies=[%s] doesNotSupply=[%s]", r.Lane, strings.Join(r.Supplies, ","), strings.Join(r.DoesNotSupply, ",")))
	}
	return strings.Join(parts, " | ")
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("candidate NEff=%.16g CYukawa=%.16g CHiggs=%.16g official NEff=%.16g CYukawa=%.16g CHiggs=%.16g", i.CandidateNEff, i.CandidateCYukawa, i.CandidateCHiggs, i.OfficialNEff, i.OfficialCYukawa, i.OfficialCHiggs)
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
