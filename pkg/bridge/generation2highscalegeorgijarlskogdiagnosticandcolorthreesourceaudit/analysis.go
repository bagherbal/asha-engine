// Package generation2highscalegeorgijarlskogdiagnosticandcolorthreesourceaudit implements
// Gate 798: High-Scale Georgi-Jarlskog Diagnostic and Color-Three Source Audit.
//
// Gate 798 defines a read-only diagnostic branch for comparing the low-scale
// color-three participation shadow N_eff≈3 with possible high-scale
// Georgi-Jarlskog Clebsch-three structure, while preserving every Yukawa,
// triality, and Higgs firewall.
package generation2highscalegeorgijarlskogdiagnosticandcolorthreesourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE798-HIGH-SCALE-GEORGI-JARLSKOG-COLOR-THREE-DIAGNOSTIC-AUDIT"

	StatusGate797Inherited        = "PASS_GATE797_PATTERN_AIRLOCK_INHERITED"
	StatusNEffThreeInherited      = "PASS_N_EFF_THREE_SOURCE_STATUS_INHERITED"
	StatusHypothesisDefined       = "PASS_GEORGI_JARLSKOG_DIAGNOSTIC_HYPOTHESIS_DEFINED"
	StatusMultiScaleRequirement   = "PASS_MULTISCALE_YUKAWA_LEDGER_REQUIREMENT_DEFINED"
	StatusGJRatiosDefined         = "PASS_GEORGI_JARLSKOG_RATIO_DIAGNOSTICS_DEFINED"
	StatusThreeReadoutsDistinct   = "PASS_N_EFF_AND_GJ_THREE_TYPED_AS_DISTINCT_READOUTS"
	StatusFNCompatibilityDefined  = "PASS_FN_COMPATIBILITY_CHECK_DEFINED"
	StatusKoideScaleDefined       = "PASS_KOIDE_SCALE_COMPATIBILITY_CHECK_DEFINED"
	StatusHexagramFirewall        = "PASS_HEXAGRAM_MOTIF_FIREWALL_AUDITED"
	StatusOutcomeTableDefined     = "PASS_DIAGNOSTIC_OUTCOME_TABLE_DEFINED"
	StatusCHiggsFirewallPreserved = "PASS_C_HIGGS_FORMULA_FIREWALL_PRESERVED"
	StatusBranchDecision          = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusCompareColorClebsch       = "CONDITIONAL_SUPPORT_COLOR_THREE_AND_CLEBSCH_THREE_MAY_BE_COMPARED_AS_DISTINCT_READOUTS"
	StatusGJHighScaleClebsch        = "CONDITIONAL_SUPPORT_GJ_RATIOS_TEST_HIGH_SCALE_CLEBSCH_THREE_STRUCTURE"
	StatusLowHighComparisonLawful   = "CONDITIONAL_SUPPORT_COMPARING_LOW_SCALE_PARTICIPATION_AND_HIGH_SCALE_CLEBSCH_THREE_IS_LAWFUL"
	StatusFNAfterLedger             = "CONDITIONAL_SUPPORT_FN_DIAGNOSTIC_CAN_CLASSIFY_HIERARCHY_AFTER_LEDGER_INPUT"
	StatusKoideAfterLedger          = "CONDITIONAL_SUPPORT_KOIDE_CAN_BE_USED_AS_CHARGED_LEPTON_DIAGNOSTIC_AFTER_LEDGER_INPUT"
	StatusHexagonalSearchMotivation = "CONDITIONAL_SUPPORT_HEXAGONAL_GEOMETRY_CAN_MOTIVATE_SU3_A2_WEIGHT_SEARCH"
	StatusOutcomeCanDecideThree     = "CONDITIONAL_SUPPORT_GATE798_CAN_DECIDE_WHETHER_THREE_IS_COLOR_ONLY_OR_HIGH_SCALE_CLEBSCH_RESONANT"
	StatusValidatedLedgerUpdatesC   = "CONDITIONAL_SUPPORT_VALIDATED_LEDGER_CAN_UPDATE_OR_CONFIRM_C_YUKAWA"

	StatusHypothesisNotNative      = "FAILED_ROUTE_HYPOTHESIS_NOT_NATIVE_THEOREM"
	StatusSingleScaleNoGJ          = "FAILED_ROUTE_SINGLE_SCALE_MZ_LEDGER_CANNOT_TEST_GEORGI_JARLSKOG"
	StatusRGThresholdRequired      = "FAILED_ROUTE_RG_THRESHOLD_PACKAGE_REQUIRED_FOR_HIGH_SCALE_DIAGNOSTIC"
	StatusGJNotNative              = "FAILED_ROUTE_GJ_RATIOS_NOT_NATIVE_YUKAWA_DERIVATION"
	StatusThreeReadoutsNotSame     = "FAILED_ROUTE_N_EFF_THREE_AND_GJ_THREE_NOT_IDENTICAL_THEOREMS"
	StatusFNNotNative              = "FAILED_ROUTE_FN_POWERS_NOT_NATIVE_CHARGE_THEOREM"
	StatusEpsilonNoSilentFit       = "FAILED_ROUTE_EPSILON_MUST_NOT_BE_FITTED_SILENTLY"
	StatusKoideNotNative           = "FAILED_ROUTE_KOIDE_NOT_NATIVE_YUKAWA_THEOREM"
	StatusVisualMotifNotEvidence   = "FAILED_ROUTE_SYMBOLIC_VISUAL_MOTIF_NOT_TYPED_EVIDENCE"
	StatusHexagramNotYukawa        = "FAILED_ROUTE_HEXAGRAM_NOT_YUKAWA_THEOREM"
	StatusPatternsDoNotModifyCH    = "FAILED_ROUTE_GJ_FN_KOIDE_DO_NOT_MODIFY_C_HIGGS_FORMULA"
	StatusCHiggsNotLevelC          = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_AFTER_PATTERN_DIAGNOSTICS"
	StatusFirewallPreservedGate798 = "FIREWALL_PRESERVED_GATE798_HIGH_SCALE_GJ_COLOR_THREE_DIAGNOSTIC_BOUNDARY"
)

const (
	nEffSnapshot    = 3.0023273474722147
	cYukawaSnapshot = 0.9992248188812008
	cHiggsSnapshot  = 1.0372205204048603
)

type Gate797Inheritance struct {
	Inherited              bool
	NEffThreeStatusKnown   bool
	CurrentThreeSource     string
	NotGenerationTriality  bool
	NotD4Triality          bool
	NotNativeYukawaTheorem bool
	NEff                   float64
	CYukawa                float64
	Verdict                string
}

type DiagnosticHypothesis struct {
	Defined   bool
	Name      string
	LowScale  string
	HighScale string
	Blocked   []string
	Verdict   string
}

type MultiScaleRequirement struct {
	Defined       bool
	SealName      string
	Fields        []string
	MinimumValues []string
	Optional      []string
	SingleScaleOK bool
	Verdict       string
}

type GJDiagnostic struct {
	Defined     bool
	Ratios      []string
	LogResidual []string
	ClosureNorm string
	HighScale   bool
	Verdict     string
}

type ThreeComparison struct {
	Recorded         bool
	NEffThreeType    string
	GJThreeType      string
	SameTypedObject  bool
	LawfulComparison bool
	BlockedShortcuts []string
	Verdict          string
}

type CompatibilityDiagnostic struct {
	Name     string
	Defined  bool
	Requires []string
	Formulae []string
	Allowed  []string
	Blocked  []string
	Verdict  string
}

type HexagramFirewall struct {
	Audited         bool
	LawfulReadings  []string
	AllowedUse      string
	ForbiddenUse    []string
	RequiredTheorem []string
	Verdict         string
}

type OutcomeTable struct {
	Defined  bool
	Outcomes []string
	Verdict  string
}

type CHiggsImpact struct {
	Recorded               bool
	Formula                string
	ValidatedLedgerCanSetC bool
	PatternsModifyFormula  bool
	CHiggsLevel            string
	CurrentCHiggs          float64
	Verdict                string
}

type BranchDecision struct {
	Recorded          bool
	Recommended       string
	MultiScaleLedger  bool
	SingleScaleLedger bool
	AnyLedger         bool
	Alternatives      []string
	Verdict           string
}

type Firewalls struct {
	Enforced              bool
	GJNativeYukawa        bool
	GUTUnificationTheorem bool
	FNChargeTheorem       bool
	KoideNativeYukawa     bool
	VisualMotifProof      bool
	NEffGenerationTheorem bool
	NEffD4Triality        bool
	CHiggsLevelC          bool
	TreeProxyPoleMass     bool
	Verdict               string
}

type Analysis struct {
	Gate797        Gate797Inheritance
	Hypothesis     DiagnosticHypothesis
	Requirement    MultiScaleRequirement
	GJ             GJDiagnostic
	Comparison     ThreeComparison
	FN             CompatibilityDiagnostic
	Koide          CompatibilityDiagnostic
	Hexagram       HexagramFirewall
	Outcomes       OutcomeTable
	Impact         CHiggsImpact
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

func BuildDefault() (Analysis, error) {
	if nEffSnapshot <= 3 || cYukawaSnapshot <= 0 || cHiggsSnapshot <= 0 || math.IsNaN(nEffSnapshot) {
		return Analysis{}, fmt.Errorf("invalid inherited Gate 798 snapshots")
	}
	return Analysis{
		Gate797: Gate797Inheritance{
			Inherited: true, NEffThreeStatusKnown: true,
			CurrentThreeSource:    "color-tripled top dominance",
			NotGenerationTriality: true, NotD4Triality: true, NotNativeYukawaTheorem: true,
			NEff: nEffSnapshot, CYukawa: cYukawaSnapshot, Verdict: StatusGate797Inherited,
		},
		Hypothesis: DiagnosticHypothesis{
			Defined: true, Name: "H_GJ",
			LowScale:  "N_eff≈3 from color-tripled top dominance",
			HighScale: "Georgi-Jarlskog Clebsch factors in down/lepton Yukawa ratios",
			Blocked:   []string{"native Yukawa theorem", "generation theorem", "D4 triality theorem"},
			Verdict:   StatusHypothesisDefined,
		},
		Requirement: MultiScaleRequirement{
			Defined: true, SealName: "MultiScaleYukawaLedgerSeal",
			Fields:        []string{"low_scale_mu", "high_scale_mu", "RG_scheme", "threshold_convention", "sector_singular_values_at_mu_i", "uncertainty_model", "normalization_convention"},
			MinimumValues: []string{"y_d(mu_high)", "y_s(mu_high)", "y_b(mu_high)", "y_e(mu_high)", "y_mu(mu_high)", "y_tau(mu_high)"},
			Optional:      []string{"y_u,y_c,y_t", "neutrino convention", "full scale trajectory y_f(mu_i)"},
			SingleScaleOK: false,
			Verdict:       StatusMultiScaleRequirement,
		},
		GJ: GJDiagnostic{
			Defined: true, HighScale: true,
			Ratios:      []string{"R_GJ_3=y_b/y_tau", "R_GJ_2=y_mu/(3y_s)", "R_GJ_1=(3y_e)/y_d"},
			LogResidual: []string{"Delta_GJ_3=log(y_b/y_tau)", "Delta_GJ_2=log(y_mu/(3y_s))", "Delta_GJ_1=log(3y_e/y_d)"},
			ClosureNorm: "||Delta_GJ||^2=Delta_GJ_1^2+Delta_GJ_2^2+Delta_GJ_3^2",
			Verdict:     StatusGJRatiosDefined,
		},
		Comparison: ThreeComparison{
			Recorded:         true,
			NEffThreeType:    "inverse participation count with certified low-scale color multiplicity source",
			GJThreeType:      "high-scale down/lepton Clebsch factor diagnostic",
			SameTypedObject:  false,
			LawfulComparison: true,
			BlockedShortcuts: []string{"N_eff≈3 implies GJ", "GJ implies N_eff≈3", "both prove native triality"},
			Verdict:          StatusThreeReadoutsDistinct,
		},
		FN: CompatibilityDiagnostic{
			Name: "Froggatt-Nielsen compatibility check", Defined: true,
			Requires: []string{"full Yukawa ledger", "declared epsilon", "fit declaration if epsilon fitted"},
			Formulae: []string{"n_f(epsilon)=log(y_f/y_t_reference)/log(epsilon)"},
			Allowed:  []string{"classify hierarchy powers", "check alignment/conflict with GJ pattern"},
			Blocked:  []string{"native charge theorem", "silent epsilon fit", "invent trace atoms"},
			Verdict:  StatusFNCompatibilityDefined,
		},
		Koide: CompatibilityDiagnostic{
			Name: "Koide scale compatibility check", Defined: true,
			Requires: []string{"y_e,y_mu,y_tau", "declared convention", "optional multi-scale values"},
			Formulae: []string{"Q_e=(y_e+y_mu+y_tau)/(sqrt(y_e)+sqrt(y_mu)+sqrt(y_tau))^2"},
			Allowed:  []string{"charged-lepton diagnostic", "low/high scale drift diagnostic"},
			Blocked:  []string{"derive charged-lepton Yukawas", "prove generation geometry", "source N_eff"},
			Verdict:  StatusKoideScaleDefined,
		},
		Hexagram: HexagramFirewall{
			Audited:         true,
			LawfulReadings:  []string{"A2 / SU(3) root or weight hexagon", "color-anticolor triangular duality", "six outer weights around a center", "two interlaced triangles as mnemonic"},
			AllowedUse:      "motivate checking color-three, Clebsch-three, and SU(3)-type weight geometry",
			ForbiddenUse:    []string{"visual proof", "typed ASHA evidence", "D4 triality theorem", "Yukawa theorem"},
			RequiredTheorem: []string{"typed carrier", "root/weight system", "representation map into Yukawa sectors", "trace-readout map into a,b,N_eff or GJ ratios"},
			Verdict:         StatusHexagramFirewall,
		},
		Outcomes: OutcomeTable{Defined: true, Outcomes: []string{
			"A: GJ ratios close and N_eff near 3 => shared three-pressure research branch, not native theorem",
			"B: N_eff near 3 but GJ fails => current three likely top-color participation only",
			"C: GJ succeeds but N_eff decomposition unexpected => distinct mechanisms",
			"D: both fail => inherited aggregate ledger or convention may be mismatched",
		}, Verdict: StatusOutcomeTableDefined},
		Impact:         CHiggsImpact{Recorded: true, Formula: "C_Higgs=(3/N_eff)C_History", ValidatedLedgerCanSetC: true, PatternsModifyFormula: false, CHiggsLevel: "Level B", CurrentCHiggs: cHiggsSnapshot, Verdict: StatusCHiggsFirewallPreserved},
		Branch:         BranchDecision{Recorded: true, Recommended: "Gate 799 — Native Three-Source Candidate Ranking and D4/SU3 Carrier Firewall Audit", MultiScaleLedger: false, SingleScaleLedger: false, AnyLedger: false, Alternatives: []string{"Georgi-Jarlskog Ratio Evaluation and N_eff Scale-Comparison Audit", "Sector Contribution to N_eff Deviation and Top-Rest Dominance Audit"}, Verdict: StatusBranchDecision},
		Firewalls:      Firewalls{Enforced: true, Verdict: StatusFirewallPreservedGate798},
		Truth:          "Gate 798 defines GJ as a high-scale read-only diagnostic and keeps low-scale N_eff three and high-scale Clebsch three as distinct typed readouts.",
		FinalStatement: "Gate 798 tests the right hypothesis without cheating. It does not claim that N_eff≈3 is triality; it asks whether a convention-locked Yukawa ledger shows both low-scale color-three participation and high-scale Georgi-Jarlskog Clebsch-three structure. Without a validated multi-scale ledger, the certified source remains top-color dominance plus small non-top rest pressure.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate797Inherited, StatusNEffThreeInherited, StatusHypothesisDefined, StatusMultiScaleRequirement, StatusGJRatiosDefined, StatusThreeReadoutsDistinct, StatusFNCompatibilityDefined, StatusKoideScaleDefined, StatusHexagramFirewall, StatusOutcomeTableDefined, StatusCHiggsFirewallPreserved, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusCompareColorClebsch, StatusGJHighScaleClebsch, StatusLowHighComparisonLawful, StatusFNAfterLedger, StatusKoideAfterLedger, StatusHexagonalSearchMotivation, StatusOutcomeCanDecideThree, StatusValidatedLedgerUpdatesC,
		StatusHypothesisNotNative, StatusSingleScaleNoGJ, StatusRGThresholdRequired, StatusGJNotNative, StatusThreeReadoutsNotSame, StatusFNNotNative, StatusEpsilonNoSilentFit, StatusKoideNotNative, StatusVisualMotifNotEvidence, StatusHexagramNotYukawa, StatusPatternsDoNotModifyCH, StatusCHiggsNotLevelC, StatusFirewallPreservedGate798,
	}
}

func FormatHypothesis(h DiagnosticHypothesis) string {
	return h.Name + ": low-scale=" + h.LowScale + "; high-scale=" + h.HighScale + "; blocked=" + strings.Join(h.Blocked, ", ")
}
func FormatRequirement(r MultiScaleRequirement) string {
	return r.SealName + ": fields=" + strings.Join(r.Fields, ", ") + "; minimum=" + strings.Join(r.MinimumValues, ", ")
}
func FormatGJ(g GJDiagnostic) string {
	return strings.Join(g.Ratios, "; ") + "; " + strings.Join(g.LogResidual, "; ") + "; " + g.ClosureNorm
}
func FormatComparison(c ThreeComparison) string {
	return "N_eff three=" + c.NEffThreeType + "; GJ three=" + c.GJThreeType
}
func FormatDiagnostic(d CompatibilityDiagnostic) string {
	return d.Name + ": requires=" + strings.Join(d.Requires, ", ") + "; formulae=" + strings.Join(d.Formulae, ", ")
}
func FormatHexagram(h HexagramFirewall) string {
	return "lawful=" + strings.Join(h.LawfulReadings, ", ") + "; forbidden=" + strings.Join(h.ForbiddenUse, ", ")
}
func FormatOutcomes(o OutcomeTable) string { return strings.Join(o.Outcomes, " | ") }
func FormatImpact(i CHiggsImpact) string {
	return fmt.Sprintf("%s; patterns_modify=%v; C_Higgs=%.16g", i.Formula, i.PatternsModifyFormula, i.CurrentCHiggs)
}
func FormatBranch(b BranchDecision) string {
	return b.Recommended + "; alternatives=" + strings.Join(b.Alternatives, "; ")
}

func closeAbs(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}
func containsAll(hay []string, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range hay {
			if strings.Contains(h, n) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
