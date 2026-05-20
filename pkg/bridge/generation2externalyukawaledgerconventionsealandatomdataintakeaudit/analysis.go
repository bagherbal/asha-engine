// Package generation2externalyukawaledgerconventionsealandatomdataintakeaudit implements
// Gate 796: External Yukawa Ledger Convention Seal and Atom Data Intake Audit.
//
// Gate 795 proved that aggregate a,b,N_eff do not identify sector or atom
// Yukawa trace data. Gate 796 defines the lawful external-data airlock required
// to populate the missing DecomposedYukawaTraceLedgerSeal without contaminating
// the Level-B scalar-Higgs interface or solving backwards from Higgs data.
package generation2externalyukawaledgerconventionsealandatomdataintakeaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE796-EXTERNAL-YUKAWA-LEDGER-CONVENTION-SEAL-ATOM-DATA-INTAKE-AUDIT"

	StatusGate795Inherited      = "PASS_GATE795_NON_IDENTIFIABILITY_AUDIT_INHERITED"
	StatusExternalAirlock       = "PASS_EXTERNAL_YUKAWA_DATA_AIRLOCK_SELECTED_AS_CURRENT_REQUIRED_OBJECT"
	StatusExternalSealDefined   = "PASS_EXTERNAL_YUKAWA_LEDGER_CONVENTION_SEAL_DEFINED"
	StatusCircularFirewall      = "PASS_CIRCULAR_INTAKE_FIREWALL_DEFINED"
	StatusAtomConstructionRules = "PASS_TRACE_ATOM_CONSTRUCTION_RULES_DEFINED"
	StatusColorConvention       = "PASS_COLOR_MULTIPLICITY_CONVENTION_REQUIRED"
	StatusInputSchema           = "PASS_YUKAWA_ATOM_INPUT_SCHEMA_DEFINED"
	StatusNeutrinoAudit         = "PASS_NEUTRINO_CONVENTION_AUDITED"
	StatusValidationRules       = "PASS_AGGREGATE_VALIDATION_RULES_DEFINED"
	StatusTopChannelRules       = "PASS_TOP_CHANNEL_SELECTOR_RULES_DEFINED"
	StatusSectorOutputs         = "PASS_SECTOR_CONTRIBUTION_OUTPUTS_DEFINED"
	StatusScaleStabilityRules   = "PASS_SCALE_STABILITY_INTAKE_RULES_DEFINED"
	StatusCHiggsImpact          = "PASS_LEVEL_B_C_HIGGS_IMPACT_RECORDED"
	StatusTrialityFirewall      = "PASS_TRIALITY_FIREWALL_PRESERVED"
	StatusBranchDecision        = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls     = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusExternalLedgerExplicitConventions = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_POPULATE_TRACE_ATOMS_ONLY_WITH_EXPLICIT_CONVENTIONS"
	StatusAtomLedgerNeedsLabels             = "CONDITIONAL_SUPPORT_ATOM_LEDGER_REQUIRES_SECTOR_AND_GENERATION_LABELS"
	StatusTopRestNeedsTypedTop              = "CONDITIONAL_SUPPORT_TOP_REST_DECOMPOSITION_BECOMES_NUMERICAL_ONLY_AFTER_TYPED_TOP_INPUT"
	StatusMultiScaleWouldAllowAudit         = "CONDITIONAL_SUPPORT_MULTI_SCALE_LEDGER_WOULD_ALLOW_N_EFF_SCALE_STABILITY_AUDIT"
	StatusValidatedExternalImprovesTest     = "CONDITIONAL_SUPPORT_VALIDATED_EXTERNAL_LEDGER_IMPROVES_C_HIGGS_TESTABILITY"

	StatusExternalNotNative            = "FAILED_ROUTE_EXTERNAL_YUKAWA_VALUES_NOT_NATIVE_ASHA_DERIVATION"
	StatusNoBacksolveFromNEff          = "FAILED_ROUTE_YUKAWA_ATOMS_MUST_NOT_BE_SOLVED_BACKWARDS_FROM_N_EFF"
	StatusNoTuneToCHiggsOrHiggsMass    = "FAILED_ROUTE_YUKAWA_ATOMS_MUST_NOT_BE_TUNED_TO_C_HIGGS_OR_HIGGS_MASS"
	StatusNoDoubleCountColor           = "FAILED_ROUTE_COLOR_FACTOR_MUST_NOT_BE_DOUBLE_COUNTED"
	StatusUnlabeledCannotAuditSector   = "FAILED_ROUTE_UNLABELED_YUKAWA_VALUES_CANNOT_SOURCE_SECTOR_AUDIT"
	StatusNeutrinoMustBeExplicit       = "FAILED_ROUTE_NEUTRINO_SECTOR_MUST_NOT_REMAIN_IMPLICIT"
	StatusNoSilentRenormalization      = "FAILED_ROUTE_EXTERNAL_LEDGER_MUST_NOT_BE_SILENTLY_RENORMALIZED_TO_MATCH_A_B"
	StatusTopNotInferredFromAggregate  = "FAILED_ROUTE_TOP_CHANNEL_MUST_NOT_BE_INFERRED_FROM_AGGREGATE_N_EFF"
	StatusNoSectorOutputsWithoutLedger = "FAILED_ROUTE_NO_SECTOR_OUTPUTS_WITHOUT_SUPPLIED_EXTERNAL_LEDGER"
	StatusSingleScaleLocal             = "FAILED_ROUTE_SINGLE_SCALE_LEDGER_REMAINS_SCALE_LOCAL"
	StatusValidatedExternalNotNative   = "FAILED_ROUTE_VALIDATED_EXTERNAL_LEDGER_NOT_NATIVE_YUKAWA_THEOREM"
	StatusCHiggsNotLevelC              = "FAILED_ROUTE_C_HIGGS_NOT_LEVEL_C_PREDICTION"
	StatusExternalNotD4Triality        = "FAILED_ROUTE_EXTERNAL_TRACE_LEDGER_NOT_D4_TRIALITY_THEOREM"
	StatusExternalNotGenerationTheorem = "FAILED_ROUTE_EXTERNAL_TRACE_LEDGER_NOT_GENERATION_THEOREM"
	StatusFirewallPreservedGate796     = "FIREWALL_PRESERVED_GATE796_EXTERNAL_YUKAWA_LEDGER_INTAKE_BOUNDARY"
)

const (
	aInherited    = 2.8424095142339083
	bInherited    = 2.6910096440382287
	nEffInherited = 3.0023273474722147
	cHiggsLevelB  = 1.0372205204048603
)

type Gate795Inheritance struct {
	Inherited bool
	Verdict   string
}

type ExternalSealDefinition struct {
	Defined bool
	Fields  []string
	Verdict string
}

type CircularIntakeFirewall struct {
	Defined             bool
	ForbiddenSources    []string
	ForbiddenOperations []string
	UsesHiggsBacksolve  bool
	Verdict             string
}

type AtomConstructionRules struct {
	Defined                     bool
	AtomFormula                 string
	CoefficientColorConvention  string
	RepeatedAtomConvention      string
	RequiresExactlyOneColorRule bool
	Verdict                     string
}

type YukawaAtomInputSchema struct {
	Defined         bool
	Fields          []string
	RequiredSectors []string
	RequiresLabels  bool
	Verdict         string
}

type NeutrinoConventionAudit struct {
	Audited          bool
	AllowedStatuses  []string
	ExplicitRequired bool
	ActiveStatus     string
	Verdict          string
}

type AggregateValidationRules struct {
	Defined              bool
	InheritedA           float64
	InheritedB           float64
	InheritedNEff        float64
	MustValidateA        bool
	MustValidateB        bool
	MustValidateNEff     bool
	SilentRescaleAllowed bool
	Verdict              string
}

type TopChannelSelectorRules struct {
	Defined          bool
	RequiresTypedTop bool
	MayInferFromNEff bool
	Formulae         []string
	Verdict          string
}

type SectorContributionOutputs struct {
	Defined                bool
	ExternalLedgerSupplied bool
	Outputs                []string
	CanOutputNow           bool
	Verdict                string
}

type ScaleStabilityIntakeRules struct {
	Defined             bool
	AllowsSingleScale   bool
	AllowsMultiScale    bool
	CurrentExternalData string
	MultiScaleCertified bool
	Verdict             string
}

type CHiggsImpact struct {
	Recorded                        bool
	ValidatedExternalWouldImprove   bool
	ValidatedExternalIsNativeYukawa bool
	CHiggsLevelC                    bool
	CurrentCHiggs                   float64
	Verdict                         string
}

type TrialityFirewall struct {
	Preserved                bool
	ExternalLedgerD4Theorem  bool
	ExternalLedgerGeneration bool
	RequiresD4Package        []string
	Verdict                  string
}

type BranchDecision struct {
	Recorded               bool
	ValidatedExternalFound bool
	ValidationFailed       bool
	NativeYukawaOperators  bool
	Recommended            string
	Alternatives           []string
	Verdict                string
}

type Firewalls struct {
	Enforced                   bool
	ExternalLedgerNativeYukawa bool
	ValidatedAtomsPMNSCKM      bool
	SectorDominanceGeneration  bool
	TopDominanceTopYukawa      bool
	NEffD4Triality             bool
	SingleScaleStable          bool
	CHiggsLevelC               bool
	TreeProxyPoleMass          bool
	Verdict                    string
}

type Analysis struct {
	Gate795        Gate795Inheritance
	ExternalSeal   ExternalSealDefinition
	Circular       CircularIntakeFirewall
	AtomRules      AtomConstructionRules
	Schema         YukawaAtomInputSchema
	Neutrino       NeutrinoConventionAudit
	Validation     AggregateValidationRules
	TopChannel     TopChannelSelectorRules
	SectorOutputs  SectorContributionOutputs
	Scale          ScaleStabilityIntakeRules
	Impact         CHiggsImpact
	Triality       TrialityFirewall
	Branch         BranchDecision
	Firewalls      Firewalls
	Truth          string
	FinalStatement string
}

func BuildDefault() (Analysis, error) {
	nEff := (aInherited * aInherited) / bInherited
	if !closeAbs(nEff, nEffInherited, 5e-16) {
		return Analysis{}, fmt.Errorf("inherited aggregate mismatch: N_eff=%.17g", nEff)
	}
	return Analysis{
		Gate795: Gate795Inheritance{Inherited: true, Verdict: StatusGate795Inherited},
		ExternalSeal: ExternalSealDefinition{Defined: true, Fields: []string{
			"source_label", "scale_mu", "renormalization_scheme", "Yukawa_normalization", "VEV_or_mass_conversion_convention", "sector_singular_values", "neutrino_convention", "color_multiplicity_convention", "uncertainty_model", "validation_against_aggregate_a_b",
		}, Verdict: StatusExternalSealDefined},
		Circular:       CircularIntakeFirewall{Defined: true, ForbiddenSources: []string{"N_eff", "C_Higgs", "lambda_runtime_eff", "m_H_tree_proxy", "m_H_pole", "observed Higgs mass"}, ForbiddenOperations: []string{"choose T so N_eff matches", "choose rest atoms so C_Higgs matches", "adjust Yukawa atoms using Higgs pole mass", "infer sectors from scalar bridge closure"}, UsesHiggsBacksolve: false, Verdict: StatusNoBacksolveFromNEff},
		AtomRules:      AtomConstructionRules{Defined: true, AtomFormula: "x_i=y_i^2", CoefficientColorConvention: "quark sector traces multiplied by coefficient 3", RepeatedAtomConvention: "colored quark atoms repeated three times", RequiresExactlyOneColorRule: true, Verdict: StatusColorConvention},
		Schema:         YukawaAtomInputSchema{Defined: true, Fields: []string{"fermion_label", "sector", "generation_label", "color_multiplicity", "y_value", "y_squared_atom", "y_quartic_atom", "scale_mu", "scheme", "normalization", "uncertainty"}, RequiredSectors: []string{"up:u,c,t", "down:d,s,b", "charged_lepton:e,mu,tau", "neutrino:nu1,nu2,nu3 or explicit absent/zero/unknown"}, RequiresLabels: true, Verdict: StatusAtomLedgerNeedsLabels},
		Neutrino:       NeutrinoConventionAudit{Audited: true, AllowedStatuses: []string{"Y_nu_absent", "Y_nu_zero", "Y_nu_Dirac_sealed", "Y_nu_Majorana_effective", "Y_nu_unknown"}, ExplicitRequired: true, ActiveStatus: "requires explicit status before ledger validation", Verdict: StatusNeutrinoMustBeExplicit},
		Validation:     AggregateValidationRules{Defined: true, InheritedA: aInherited, InheritedB: bInherited, InheritedNEff: nEff, MustValidateA: true, MustValidateB: true, MustValidateNEff: true, SilentRescaleAllowed: false, Verdict: StatusNoSilentRenormalization},
		TopChannel:     TopChannelSelectorRules{Defined: true, RequiresTypedTop: true, MayInferFromNEff: false, Formulae: []string{"T=y_t^2", "a_top=3T", "b_top=3T^2", "alpha=a_rest/a_top", "beta=b_rest/b_top", "b/a^2=(1/3)(1+beta)/(1+alpha)^2"}, Verdict: StatusTopRestNeedsTypedTop},
		SectorOutputs:  SectorContributionOutputs{Defined: true, ExternalLedgerSupplied: false, Outputs: []string{"a_u/a", "a_d/a", "a_e/a", "a_nu/a", "b_u/b", "b_d/b", "b_e/b", "b_nu/b", "largest atoms", "top dominance fractions", "non-top rest pressure", "neutrino status"}, CanOutputNow: false, Verdict: StatusNoSectorOutputsWithoutLedger},
		Scale:          ScaleStabilityIntakeRules{Defined: true, AllowsSingleScale: true, AllowsMultiScale: true, CurrentExternalData: "missing external ledger; inherited aggregate is single-scale M_Z", MultiScaleCertified: false, Verdict: StatusSingleScaleLocal},
		Impact:         CHiggsImpact{Recorded: true, ValidatedExternalWouldImprove: true, ValidatedExternalIsNativeYukawa: false, CHiggsLevelC: false, CurrentCHiggs: cHiggsLevelB, Verdict: StatusValidatedExternalImprovesTest},
		Triality:       TrialityFirewall{Preserved: true, ExternalLedgerD4Theorem: false, ExternalLedgerGeneration: false, RequiresD4Package: []string{"D4TrialityCarrierPackage", "trace-readout map into a,b or N_eff", "breaking operator explaining N_eff-3"}, Verdict: StatusExternalNotD4Triality},
		Branch:         BranchDecision{Recorded: true, ValidatedExternalFound: false, ValidationFailed: false, NativeYukawaOperators: false, Recommended: "Gate 797 — External Yukawa Input Request and Convention Checklist Audit", Alternatives: []string{"Gate 797 — Sector Contribution to N_eff Deviation and Top-Rest Dominance Audit", "Gate 797 — Yukawa Ledger Convention Mismatch and Normalization Firewall Audit", "Gate 797 — Native Yukawa Operator Trace Atom Extraction Audit"}, Verdict: StatusBranchDecision},
		Firewalls:      Firewalls{Enforced: true, ExternalLedgerNativeYukawa: false, ValidatedAtomsPMNSCKM: false, SectorDominanceGeneration: false, TopDominanceTopYukawa: false, NEffD4Triality: false, SingleScaleStable: false, CHiggsLevelC: false, TreeProxyPoleMass: false, Verdict: StatusFirewallPreservedGate796},
		Truth:          "Gate 796 defines the external-data airlock required to populate Yukawa trace atoms without solving backwards from N_eff, C_Higgs, or Higgs observables.",
		FinalStatement: "Gate 796 does not import or derive Yukawa atoms by itself. It defines the lawful external intake airlock: every Yukawa atom must be sector-labeled, scale-labeled, convention-labeled, color-counted exactly once, neutrino-typed explicitly, and validated against the inherited aggregate a,b,N_eff. If a validated external ledger is supplied, ASHA can audit which sectors create N_eff-3; if not, N_eff remains an aggregate sealed participation number and the Level-B Higgs interface remains scientifically testable only at the aggregate level.",
	}, nil
}

func FormatExternalSeal(x ExternalSealDefinition) string {
	return fmt.Sprintf("defined=%t fields=[%s] verdict=%s", x.Defined, strings.Join(x.Fields, ", "), x.Verdict)
}

func FormatCircular(x CircularIntakeFirewall) string {
	return fmt.Sprintf("defined=%t forbidden_sources=[%s] forbidden_ops=[%s] backsolve=%t verdict=%s", x.Defined, strings.Join(x.ForbiddenSources, ", "), strings.Join(x.ForbiddenOperations, "; "), x.UsesHiggsBacksolve, x.Verdict)
}

func FormatAtomRules(x AtomConstructionRules) string {
	return fmt.Sprintf("defined=%t atom=%s coefficient_color=%s repeated_atom=%s exactly_one=%t verdict=%s", x.Defined, x.AtomFormula, x.CoefficientColorConvention, x.RepeatedAtomConvention, x.RequiresExactlyOneColorRule, x.Verdict)
}

func FormatSchema(x YukawaAtomInputSchema) string {
	return fmt.Sprintf("defined=%t fields=[%s] sectors=[%s] labels=%t verdict=%s", x.Defined, strings.Join(x.Fields, ", "), strings.Join(x.RequiredSectors, "; "), x.RequiresLabels, x.Verdict)
}

func FormatValidation(x AggregateValidationRules) string {
	return fmt.Sprintf("defined=%t a=%.16g b=%.16g N_eff=%.16g validate_a=%t validate_b=%t validate_N=%t silent_rescale=%t verdict=%s", x.Defined, x.InheritedA, x.InheritedB, x.InheritedNEff, x.MustValidateA, x.MustValidateB, x.MustValidateNEff, x.SilentRescaleAllowed, x.Verdict)
}

func FormatTopChannel(x TopChannelSelectorRules) string {
	return fmt.Sprintf("defined=%t typed_top=%t infer_from_N=%t formulae=[%s] verdict=%s", x.Defined, x.RequiresTypedTop, x.MayInferFromNEff, strings.Join(x.Formulae, "; "), x.Verdict)
}

func FormatBranch(x BranchDecision) string {
	return fmt.Sprintf("validated_external=%t validation_failed=%t native_operators=%t recommended=%s alternatives=[%s]", x.ValidatedExternalFound, x.ValidationFailed, x.NativeYukawaOperators, x.Recommended, strings.Join(x.Alternatives, "; "))
}

func Statuses() []string {
	return []string{
		StatusGate795Inherited, StatusExternalAirlock, StatusExternalSealDefined, StatusCircularFirewall, StatusAtomConstructionRules, StatusColorConvention,
		StatusInputSchema, StatusNeutrinoAudit, StatusValidationRules, StatusTopChannelRules, StatusSectorOutputs, StatusScaleStabilityRules,
		StatusCHiggsImpact, StatusTrialityFirewall, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusExternalLedgerExplicitConventions, StatusAtomLedgerNeedsLabels, StatusTopRestNeedsTypedTop, StatusMultiScaleWouldAllowAudit, StatusValidatedExternalImprovesTest,
		StatusExternalNotNative, StatusNoBacksolveFromNEff, StatusNoTuneToCHiggsOrHiggsMass, StatusNoDoubleCountColor, StatusUnlabeledCannotAuditSector,
		StatusNeutrinoMustBeExplicit, StatusNoSilentRenormalization, StatusTopNotInferredFromAggregate, StatusNoSectorOutputsWithoutLedger,
		StatusSingleScaleLocal, StatusValidatedExternalNotNative, StatusCHiggsNotLevelC, StatusExternalNotD4Triality, StatusExternalNotGenerationTheorem,
		StatusFirewallPreservedGate796,
	}
}

func containsAll(haystack []string, needles []string) bool {
	for _, needle := range needles {
		found := false
		for _, item := range haystack {
			if strings.Contains(item, needle) {
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

func closeAbs(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
