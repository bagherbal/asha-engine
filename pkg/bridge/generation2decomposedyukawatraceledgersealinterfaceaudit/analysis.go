// Package generation2decomposedyukawatraceledgersealinterfaceaudit implements
// Gate 794: DecomposedYukawaTraceLedgerSeal Specification and Data-Interface Audit.
//
// Gate 793 showed that N_eff is the highest numerical-leverage Yukawa trace
// participation input, but the active ledger only carries aggregate a,b traces.
// Gate 794 does not derive Yukawa data; it specifies the exact decomposed
// sector/atom interface needed before N_eff can be source-audited.
package generation2decomposedyukawatraceledgersealinterfaceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE794-DECOMPOSED-YUKAWA-TRACE-LEDGER-SEAL-INTERFACE-AUDIT"

	StatusGate793Inherited            = "PASS_GATE793_DECOMPOSED_YUKAWA_TRACE_AUDIT_INHERITED"
	StatusLedgerSelectedBottleneck    = "PASS_DECOMPOSED_YUKAWA_TRACE_LEDGER_SELECTED_AS_CURRENT_N_EFF_BOTTLENECK"
	StatusSealDefined                 = "PASS_DECOMPOSED_YUKAWA_TRACE_LEDGER_SEAL_DEFINED"
	StatusSectorInterface             = "PASS_SECTOR_TRACE_INTERFACE_SPECIFIED"
	StatusAtomInterface               = "PASS_TRACE_ATOM_LEDGER_INTERFACE_SPECIFIED"
	StatusColorMultiplicityRequired   = "PASS_COLOR_MULTIPLICITY_RULE_REQUIRED"
	StatusTopSelectorInterface        = "PASS_TOP_CHANNEL_SELECTOR_INTERFACE_SPECIFIED"
	StatusNeutrinoFirewall            = "PASS_NEUTRINO_SECTOR_CONVENTION_FIREWALL_DEFINED"
	StatusScaleNormalizationInterface = "PASS_SCALE_AND_NORMALIZATION_INTERFACE_SPECIFIED"
	StatusValidationRules             = "PASS_AGGREGATE_VALIDATION_RULES_DEFINED"
	StatusSourceOutputRequirements    = "PASS_SOURCE_OUTPUT_REQUIREMENTS_DEFINED"
	StatusTrialityGenerationFirewall  = "PASS_TRIALITY_AND_GENERATION_FIREWALL_PRESERVED"
	StatusCHiggsImpact                = "PASS_C_HIGGS_IMPACT_RECORDED"
	StatusBranchDecision              = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusNEffRequiresAtomLedger        = "CONDITIONAL_SUPPORT_N_EFF_SOURCE_REDUCTION_REQUIRES_EXPLICIT_TRACE_ATOM_LEDGER"
	StatusSectorLedgerAllowsAssignment  = "CONDITIONAL_SUPPORT_SECTOR_TRACE_LEDGER_WOULD_ALLOW_SOURCE_ASSIGNMENT_OF_N_EFF_MINUS_THREE"
	StatusTopRestRequiresT              = "CONDITIONAL_SUPPORT_TOP_REST_DECOMPOSITION_REQUIRES_TYPED_T_CHANNEL"
	StatusLedgerScaleLocal              = "CONDITIONAL_SUPPORT_DECOMPOSED_LEDGER_MUST_BE_SCALE_LOCAL"
	StatusValidatedLedgerAllowsAudit    = "CONDITIONAL_SUPPORT_VALIDATED_DECOMPOSED_LEDGER_WOULD_ALLOW_SECTOR_SOURCE_AUDIT_OF_N_EFF"
	StatusLedgerUpgradesNEffToAuditable = "CONDITIONAL_SUPPORT_DECOMPOSED_LEDGER_WOULD_UPGRADE_N_EFF_FROM_AGGREGATE_SEAL_TO_SECTOR_AUDITABLE_SEAL"

	StatusAggregateABDoNotIdentifyAtoms  = "FAILED_ROUTE_AGGREGATE_A_B_VALUES_ALONE_DO_NOT_IDENTIFY_TRACE_ATOMS"
	StatusNoSectorAssignmentWithoutData  = "FAILED_ROUTE_NO_SECTOR_CONTRIBUTION_ASSIGNMENT_WITHOUT_A_U_A_D_A_E_A_NU_AND_B_SECTORS"
	StatusColorMustNotDoubleCount        = "FAILED_ROUTE_COLOR_FACTOR_MUST_NOT_BE_DOUBLE_COUNTED"
	StatusAtomsUnavailableFromAggregate  = "FAILED_ROUTE_TRACE_ATOMS_NOT_AVAILABLE_FROM_AGGREGATE_A_B_ALONE"
	StatusTMustNotBeSolvedBackwards      = "FAILED_ROUTE_T_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_N_EFF"
	StatusNoAlphaBetaWithoutTopSelector  = "FAILED_ROUTE_NO_NUMERICAL_ALPHA_BETA_WITHOUT_TOP_CHANNEL_SELECTOR"
	StatusNeutrinoConventionExplicit     = "FAILED_ROUTE_NEUTRINO_TRACE_CONVENTION_MUST_NOT_BE_LEFT_IMPLICIT"
	StatusNEffScaleStabilityNotSingle    = "FAILED_ROUTE_N_EFF_SCALE_STABILITY_NOT_CERTIFIED_BY_SINGLE_SCALE_LEDGER"
	StatusRejectIfNoReproduceABNEff      = "FAILED_ROUTE_DECOMPOSED_LEDGER_REJECTED_IF_IT_DOES_NOT_REPRODUCE_A_B_N_EFF"
	StatusCannotComputeWithoutData       = "FAILED_ROUTE_GATE794_CANNOT_COMPUTE_SECTOR_CONTRIBUTIONS_WITHOUT_SUPPLIED_DECOMPOSED_DATA"
	StatusLedgerNotGenerationTheorem     = "FAILED_ROUTE_DECOMPOSED_TRACE_LEDGER_NOT_YET_GENERATION_THEOREM"
	StatusLedgerNotD4TrialityTheorem     = "FAILED_ROUTE_DECOMPOSED_TRACE_LEDGER_NOT_YET_D4_TRIALITY_THEOREM"
	StatusSectorAuditableNotNativeYukawa = "FAILED_ROUTE_SECTOR_AUDITABLE_N_EFF_NOT_NATIVE_YUKAWA_THEOREM"
	StatusCHiggsNotLevelC                = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_PREDICTION"
	StatusFirewallPreservedGate794       = "FIREWALL_PRESERVED_GATE794_DECOMPOSED_YUKAWA_TRACE_LEDGER_SEAL_BOUNDARY"
)

const (
	aSnapshot        = 2.8424095142339083
	bSnapshot        = 2.6910096440382287
	nEffSnapshot     = 3.0023273474722147
	cYukawaSnapshot  = 0.9992248188812008
	cHistorySnapshot = 1.038025177923625
	cHiggsSnapshot   = 1.0372205204048603
)

type Gate793Inheritance struct {
	Inherited      bool
	NEffBottleneck bool
	AggregateA     float64
	AggregateB     float64
	NEff           float64
	Verdict        string
}

type LedgerSealDefinition struct {
	Defined    bool
	Name       string
	Components []string
	Verdict    string
}

type SectorTraceInterface struct {
	Specified       bool
	QuadraticTraces []string
	QuarticTraces   []string
	RequiredOutputs []string
	DataAvailable   bool
	Verdict         string
}

type TraceAtomInterface struct {
	Specified         bool
	Fields            []string
	SumRules          []string
	ColorRuleRequired bool
	ColorConvention   string
	MixingForbidden   bool
	Verdict           string
}

type TopChannelSelectorInterface struct {
	Specified       bool
	TopSelectorName string
	RequiresTypedT  bool
	MayInvertNEff   bool
	Formulas        []string
	Verdict         string
}

type NeutrinoConventionFirewall struct {
	Defined         bool
	AllowedStatuses []string
	ImplicitAllowed bool
	Verdict         string
}

type ScaleNormalizationInterface struct {
	Specified        bool
	Scale            string
	Scheme           string
	Normalization    string
	Differential     string
	MultiScaleLedger bool
	Verdict          string
}

type AggregateValidationRules struct {
	Defined      bool
	Rules        []string
	Tolerance    string
	RejectOnFail bool
	Verdict      string
}

type SourceOutputRequirements struct {
	Defined           bool
	RequiresData      bool
	OutputsIfSupplied []string
	CanComputeNow     bool
	Verdict           string
}

type TrialityGenerationFirewall struct {
	Preserved                   bool
	LedgerImpliesGeneration     bool
	LedgerImpliesD4             bool
	RequiredForNativeGeneration []string
	Verdict                     string
}

type CHiggsImpact struct {
	Recorded                       bool
	CYukawa                        float64
	CHiggs                         float64
	NEffAggregateSealed            bool
	NEffSectorAuditableIfDataValid bool
	CHiggsLevelC                   bool
	Verdict                        string
}

type BranchDecision struct {
	Recorded              bool
	ValidatedLedgerExists bool
	D4PackageIntroduced   bool
	Recommended           string
	Alternatives          []string
	Verdict               string
}

type Firewalls struct {
	Enforced                       bool
	DecomposedLedgerNativeYukawa   bool
	SectorTraceGenerationTheorem   bool
	TopSelectorTopYukawaDerivation bool
	ValidatedAtomsPMNSCKM          bool
	NEffD4Triality                 bool
	ScaleLocalScaleStable          bool
	CHiggsPoleMass                 bool
	TreeProxyPoleMass              bool
	Verdict                        string
}

type Analysis struct {
	Gate793        Gate793Inheritance
	Seal           LedgerSealDefinition
	Sector         SectorTraceInterface
	Atom           TraceAtomInterface
	TopSelector    TopChannelSelectorInterface
	Neutrino       NeutrinoConventionFirewall
	Scale          ScaleNormalizationInterface
	Validation     AggregateValidationRules
	Output         SourceOutputRequirements
	Triality       TrialityGenerationFirewall
	Impact         CHiggsImpact
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

func BuildDefault() (Analysis, error) {
	nEff := (aSnapshot * aSnapshot) / bSnapshot
	cYukawa := 3.0 / nEff
	cHiggs := cYukawa * cHistorySnapshot
	if !closeAbs(nEff, nEffSnapshot, 5e-16) || !closeAbs(cYukawa, cYukawaSnapshot, 5e-16) || !closeAbs(cHiggs, cHiggsSnapshot, 5e-16) {
		return Analysis{}, fmt.Errorf("aggregate ledger mismatch: N_eff=%.17g C_Yukawa=%.17g C_Higgs=%.17g", nEff, cYukawa, cHiggs)
	}
	analysis := Analysis{
		Gate793:        Gate793Inheritance{Inherited: true, NEffBottleneck: true, AggregateA: aSnapshot, AggregateB: bSnapshot, NEff: nEff, Verdict: StatusGate793Inherited},
		Seal:           LedgerSealDefinition{Defined: true, Name: "DecomposedYukawaTraceLedgerSeal", Components: []string{"scale_convention", "Yukawa_normalization_convention", "sector_trace_ledger", "trace_atom_ledger", "top_channel_selector", "color_multiplicity_rule", "neutrino_sector_convention", "validation_rules"}, Verdict: StatusNEffRequiresAtomLedger},
		Sector:         SectorTraceInterface{Specified: true, QuadraticTraces: []string{"a_u=3 Tr(Y_u†Y_u)", "a_d=3 Tr(Y_d†Y_d)", "a_e=Tr(Y_e†Y_e)", "a_nu=Tr(Y_nu†Y_nu)"}, QuarticTraces: []string{"b_u=3 Tr((Y_u†Y_u)^2)", "b_d=3 Tr((Y_d†Y_d)^2)", "b_e=Tr((Y_e†Y_e)^2)", "b_nu=Tr((Y_nu†Y_nu)^2)"}, RequiredOutputs: []string{"a_sector/a", "b_sector/b", "sector contribution to N_eff-3"}, DataAvailable: false, Verdict: StatusNoSectorAssignmentWithoutData},
		Atom:           TraceAtomInterface{Specified: true, Fields: []string{"atom_id", "sector", "generation_label_if_available", "color_label_or_multiplicity", "squared_singular_value x_i", "quartic_atom x_i^2", "scale", "convention"}, SumRules: []string{"sum_i x_i = a", "sum_i x_i^2 = b", "N_eff=(sum_i x_i)^2/sum_i x_i^2"}, ColorRuleRequired: true, ColorConvention: "choose coefficient representation or repeated-atom representation; never both", MixingForbidden: true, Verdict: StatusAtomsUnavailableFromAggregate},
		TopSelector:    TopChannelSelectorInterface{Specified: true, TopSelectorName: "TopChannelSelector", RequiresTypedT: true, MayInvertNEff: false, Formulas: []string{"a_top=3T", "b_top=3T^2", "alpha=a_rest/(3T)", "beta=b_rest/(3T^2)", "b/a^2=(1/3)(1+beta)/(1+alpha)^2"}, Verdict: StatusNoAlphaBetaWithoutTopSelector},
		Neutrino:       NeutrinoConventionFirewall{Defined: true, AllowedStatuses: []string{"Y_nu absent", "Y_nu zero", "Y_nu Dirac sealed", "Y_nu Majorana-effective with explicit operator normalization", "Y_nu unknown means incomplete ledger"}, ImplicitAllowed: false, Verdict: StatusNeutrinoConventionExplicit},
		Scale:          ScaleNormalizationInterface{Specified: true, Scale: "M_Z", Scheme: "supplied_or_unknown", Normalization: "supplied_or_unknown", Differential: "d ln N_eff = 2 d ln a - d ln b", MultiScaleLedger: false, Verdict: StatusNEffScaleStabilityNotSingle},
		Validation:     AggregateValidationRules{Defined: true, Rules: []string{"abs((a_u+a_d+a_e+a_nu)-a_inherited)<=tolerance", "abs((b_u+b_d+b_e+b_nu)-b_inherited)<=tolerance", "abs((a^2/b)-N_eff_inherited)<=tolerance", "abs(sum_i x_i-a_inherited)<=tolerance", "abs(sum_i x_i^2-b_inherited)<=tolerance"}, Tolerance: "explicit caller-supplied numerical tolerance", RejectOnFail: true, Verdict: StatusRejectIfNoReproduceABNEff},
		Output:         SourceOutputRequirements{Defined: true, RequiresData: true, OutputsIfSupplied: []string{"sector fractions", "top/rest T alpha beta", "N_eff and epsilon_Yukawa", "top contribution to a and b", "largest non-top trace atoms"}, CanComputeNow: false, Verdict: StatusCannotComputeWithoutData},
		Triality:       TrialityGenerationFirewall{Preserved: true, LedgerImpliesGeneration: false, LedgerImpliesD4: false, RequiredForNativeGeneration: []string{"generation carrier or D4 carrier", "sector operator map", "trace-readout theorem", "breaking operator"}, Verdict: StatusLedgerNotGenerationTheorem},
		Impact:         CHiggsImpact{Recorded: true, CYukawa: cYukawa, CHiggs: cHiggs, NEffAggregateSealed: true, NEffSectorAuditableIfDataValid: true, CHiggsLevelC: false, Verdict: StatusLedgerUpgradesNEffToAuditable},
		Branch:         BranchDecision{Recorded: true, ValidatedLedgerExists: false, D4PackageIntroduced: false, Recommended: "Gate 795 — Yukawa Trace Atom Data Acquisition and Convention-Seal Audit", Alternatives: []string{"Gate 795 — Sector Contribution to N_eff Deviation and Top-Rest Dominance Audit", "Gate 795 — D4 Triality Trilinear Coupling and Yukawa Trace Readout Audit"}, Verdict: StatusBranchDecision},
		Firewalls:      Firewalls{Enforced: true, DecomposedLedgerNativeYukawa: false, SectorTraceGenerationTheorem: false, TopSelectorTopYukawaDerivation: false, ValidatedAtomsPMNSCKM: false, NEffD4Triality: false, ScaleLocalScaleStable: false, CHiggsPoleMass: false, TreeProxyPoleMass: false, Verdict: StatusFirewallPreservedGate794},
		Truth:          "Gate 794 specifies the decomposed Yukawa trace data interface required to make N_eff sector-auditable; aggregate a,b alone do not identify trace atoms.",
		FinalStatement: "Gate 794 does not decompose N_eff yet unless the sector/atom ledger is actually supplied. It specifies DecomposedYukawaTraceLedgerSeal: sector traces, atom traces, top-channel selector, color multiplicity rule, neutrino convention, scale convention, normalization convention, and validation rules. The next move is a precise data-interface problem: Gate 795 should either audit validated sector contributions if the ledger exists, or acquire/specify the Yukawa trace atom data needed to make N_eff source-reduction real.",
	}
	return analysis, nil
}

func FormatSeal(s LedgerSealDefinition) string {
	return fmt.Sprintf("%s=(%s)", s.Name, strings.Join(s.Components, ", "))
}

func FormatSector(s SectorTraceInterface) string {
	return fmt.Sprintf("quadratic=[%s] quartic=[%s] data_available=%v", strings.Join(s.QuadraticTraces, "; "), strings.Join(s.QuarticTraces, "; "), s.DataAvailable)
}

func FormatAtom(a TraceAtomInterface) string {
	return fmt.Sprintf("fields=[%s] color_rule=%s", strings.Join(a.Fields, "; "), a.ColorConvention)
}

func FormatTopSelector(t TopChannelSelectorInterface) string {
	return fmt.Sprintf("selector=%s requires_T=%v may_invert_N_eff=%v formulas=[%s]", t.TopSelectorName, t.RequiresTypedT, t.MayInvertNEff, strings.Join(t.Formulas, "; "))
}

func FormatValidation(v AggregateValidationRules) string {
	return fmt.Sprintf("rules=[%s] reject_on_fail=%v", strings.Join(v.Rules, "; "), v.RejectOnFail)
}

func FormatImpact(i CHiggsImpact) string {
	return fmt.Sprintf("C_Yukawa=%.16g C_Higgs=%.16g aggregate_sealed=%v sector_auditable_if_valid=%v level_C=%v", i.CYukawa, i.CHiggs, i.NEffAggregateSealed, i.NEffSectorAuditableIfDataValid, i.CHiggsLevelC)
}

func FormatBranch(b BranchDecision) string {
	return fmt.Sprintf("recommended=%s alternatives=%s", b.Recommended, strings.Join(b.Alternatives, " | "))
}

func Statuses() []string {
	return []string{
		StatusGate793Inherited,
		StatusLedgerSelectedBottleneck,
		StatusSealDefined,
		StatusSectorInterface,
		StatusAtomInterface,
		StatusColorMultiplicityRequired,
		StatusTopSelectorInterface,
		StatusNeutrinoFirewall,
		StatusScaleNormalizationInterface,
		StatusValidationRules,
		StatusSourceOutputRequirements,
		StatusTrialityGenerationFirewall,
		StatusCHiggsImpact,
		StatusBranchDecision,
		StatusPhysicalFirewalls,
		StatusNEffRequiresAtomLedger,
		StatusSectorLedgerAllowsAssignment,
		StatusTopRestRequiresT,
		StatusLedgerScaleLocal,
		StatusValidatedLedgerAllowsAudit,
		StatusLedgerUpgradesNEffToAuditable,
		StatusAggregateABDoNotIdentifyAtoms,
		StatusNoSectorAssignmentWithoutData,
		StatusColorMustNotDoubleCount,
		StatusAtomsUnavailableFromAggregate,
		StatusTMustNotBeSolvedBackwards,
		StatusNoAlphaBetaWithoutTopSelector,
		StatusNeutrinoConventionExplicit,
		StatusNEffScaleStabilityNotSingle,
		StatusRejectIfNoReproduceABNEff,
		StatusCannotComputeWithoutData,
		StatusLedgerNotGenerationTheorem,
		StatusLedgerNotD4TrialityTheorem,
		StatusSectorAuditableNotNativeYukawa,
		StatusCHiggsNotLevelC,
		StatusFirewallPreservedGate794,
	}
}

var once struct {
	sync.Once
	statuses []string
}

func CachedStatuses() []string {
	once.Do(func() { once.statuses = Statuses() })
	return once.statuses
}

func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }

func containsAll(haystack []string, needles []string) bool {
	joined := strings.Join(haystack, "\n")
	for _, n := range needles {
		if !strings.Contains(joined, n) {
			return false
		}
	}
	return true
}
