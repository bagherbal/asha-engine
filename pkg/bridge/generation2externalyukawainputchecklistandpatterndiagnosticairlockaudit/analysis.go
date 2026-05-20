// Package generation2externalyukawainputchecklistandpatterndiagnosticairlockaudit implements
// Gate 797: External Yukawa Input Checklist and Pattern-Diagnostic Airlock Audit.
//
// Gate 797 turns Gate 796's external Yukawa intake airlock into an explicit input
// checklist and classifies Koide, Froggatt-Nielsen, and b-tau diagnostics as
// read-only pattern tests that may be run only after a convention-locked Yukawa
// ledger is supplied. It does not import, infer, or derive Yukawa atoms.
package generation2externalyukawainputchecklistandpatterndiagnosticairlockaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE797-EXTERNAL-YUKAWA-INPUT-CHECKLIST-PATTERN-DIAGNOSTIC-AIRLOCK-AUDIT"

	StatusGate796Inherited    = "PASS_GATE796_EXTERNAL_YUKAWA_INTAKE_AIRLOCK_INHERITED"
	StatusPatternsReadOnly    = "PASS_PATTERN_DIAGNOSTICS_CLASSIFIED_AS_READ_ONLY_TESTS"
	StatusChecklistDefined    = "PASS_EXTERNAL_YUKAWA_INPUT_CHECKLIST_DEFINED"
	StatusAtomProtocolDefined = "PASS_ATOM_CONSTRUCTION_PROTOCOL_DEFINED"
	StatusAggregateValidation = "PASS_AGGREGATE_VALIDATION_PROTOCOL_DEFINED"
	StatusKoideAirlock        = "PASS_KOIDE_DIAGNOSTIC_AIRLOCK_DEFINED"
	StatusFNAirlock           = "PASS_FROGGATT_NIELSEN_DIAGNOSTIC_AIRLOCK_DEFINED"
	StatusBTauAirlock         = "PASS_B_TAU_UNIFICATION_DIAGNOSTIC_AIRLOCK_DEFINED"
	StatusPatternPriority     = "PASS_PATTERN_PRIORITY_CLASSIFICATION_RECORDED"
	StatusCHiggsImpact        = "PASS_C_HIGGS_IMPACT_RECORDED"
	StatusBranchDecision      = "PASS_BRANCH_DECISION_RECORDED"
	StatusPhysicalFirewalls   = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusExternalLedgerAfterLock = "CONDITIONAL_SUPPORT_EXTERNAL_LEDGER_CAN_POPULATE_TRACE_ATOMS_AFTER_CONVENTION_LOCK"
	StatusKoideAfterLedger        = "CONDITIONAL_SUPPORT_KOIDE_CAN_BE_RUN_AFTER_CHARGED_LEPTON_LEDGER_IS_SUPPLIED"
	StatusFNAfterLedger           = "CONDITIONAL_SUPPORT_FN_POWER_PATTERN_CAN_CLASSIFY_YUKAWA_HIERARCHY_AFTER_LEDGER_INPUT"
	StatusBTauRequiresRG          = "CONDITIONAL_SUPPORT_B_TAU_DIAGNOSTIC_REQUIRES_MULTI_SCALE_OR_RG_TRANSPORTED_LEDGER"
	StatusFullAtomLedgerPrimary   = "CONDITIONAL_SUPPORT_FULL_ATOM_LEDGER_REMAINS_PRIMARY_NEED"
	StatusDiagnosticsSecondary    = "CONDITIONAL_SUPPORT_KOIDE_FN_BTAU_ARE_SECONDARY_DIAGNOSTICS_NOT_DATA_SOURCES"
	StatusValidatedUpdatesCYukawa = "CONDITIONAL_SUPPORT_VALIDATED_LEDGER_CAN_UPDATE_OR_CONFIRM_C_YUKAWA"

	StatusImplicitScaleRejected        = "FAILED_ROUTE_YUKAWA_LEDGER_REJECTED_IF_SCALE_OR_NORMALIZATION_IS_IMPLICIT"
	StatusColorNoDoubleCount           = "FAILED_ROUTE_COLOR_FACTOR_MUST_NOT_BE_DOUBLE_COUNTED"
	StatusNoSilentRenormalization      = "FAILED_ROUTE_EXTERNAL_LEDGER_MUST_NOT_BE_SILENTLY_RENORMALIZED"
	StatusFailedValidationBlocks       = "FAILED_ROUTE_FAILED_VALIDATION_BLOCKS_SECTOR_INTERPRETATION_OF_INHERITED_N_EFF"
	StatusKoideNotNative               = "FAILED_ROUTE_KOIDE_FORMULA_NOT_NATIVE_YUKAWA_THEOREM"
	StatusKoideNoBacksolve             = "FAILED_ROUTE_KOIDE_CANNOT_BE_USED_TO_SOLVE_YUKAWA_ATOMS_BACKWARDS"
	StatusKoideNotFullNEff             = "FAILED_ROUTE_KOIDE_CHARGED_LEPTON_PATTERN_NOT_FULL_N_EFF_SOURCE"
	StatusFNNotNative                  = "FAILED_ROUTE_FN_PATTERN_NOT_NATIVE_YUKAWA_OPERATOR_THEOREM"
	StatusFNPowersNoInvent             = "FAILED_ROUTE_EPSILON_POWERS_MUST_NOT_BE_USED_TO_INVENT_TRACE_ATOMS"
	StatusNoNativeFNCharges            = "FAILED_ROUTE_NO_NATIVE_FN_CHARGE_ASSIGNMENT_THEOREM"
	StatusSingleScaleNoBTau            = "FAILED_ROUTE_SINGLE_SCALE_MZ_LEDGER_CANNOT_CERTIFY_B_TAU_UNIFICATION"
	StatusBTauNotNative                = "FAILED_ROUTE_B_TAU_UNIFICATION_NOT_NATIVE_YUKAWA_THEOREM"
	StatusRGThresholdRequired          = "FAILED_ROUTE_RG_THRESHOLD_PACKAGE_REQUIRED_FOR_HIGH_SCALE_COMPARISON"
	StatusDiagnosticsDoNotModifyCHiggs = "FAILED_ROUTE_PATTERN_DIAGNOSTICS_DO_NOT_MODIFY_C_HIGGS_FORMULA_BY_THEMSELVES"
	StatusCHiggsRemainsLevelB          = "FAILED_ROUTE_C_HIGGS_REMAINS_LEVEL_B_AFTER_EXTERNAL_PATTERN_DIAGNOSTICS"
	StatusFirewallPreservedGate797     = "FIREWALL_PRESERVED_GATE797_EXTERNAL_YUKAWA_INPUT_PATTERN_AIRLOCK_BOUNDARY"
)

const (
	aInherited    = 2.8424095142339083
	bInherited    = 2.6910096440382287
	nEffInherited = 3.0023273474722147
	cHiggsLevelB  = 1.0372205204048603
)

type Gate796Inheritance struct {
	Inherited bool
	Verdict   string
}

type InputChecklist struct {
	Defined          bool
	Fields           []string
	TargetScale      string
	RequiredStatuses []string
	RejectImplicit   bool
	Verdict          string
}

type AtomConstructionProtocol struct {
	Defined                bool
	AtomFormula            string
	CoefficientColorRules  []string
	Computes               []string
	RequiresConventionLock bool
	Verdict                string
}

type AggregateValidationProtocol struct {
	Defined              bool
	InheritedA           float64
	InheritedB           float64
	InheritedNEff        float64
	ClassifiedFailures   []string
	SilentRescaleAllowed bool
	Verdict              string
}

type PatternDiagnostic struct {
	Name                  string
	Defined               bool
	Requires              []string
	DiagnosticFormulae    []string
	AllowedUse            []string
	BlockedUse            []string
	ReadOnly              bool
	CanPopulateAtomLedger bool
	Verdict               string
}

type PatternPriority struct {
	Recorded bool
	Ranking  []string
	Verdict  string
}

type CHiggsImpact struct {
	Recorded                  bool
	ValidatedLedgerCanConfirm bool
	PatternsModifyFormula     bool
	CHiggsLevel               string
	CurrentCHiggs             float64
	Verdict                   string
}

type BranchDecision struct {
	Recorded       bool
	Recommended    string
	ExternalLedger bool
	NativeYukawa   bool
	D4Package      bool
	Alternatives   []string
	Verdict        string
}

type Firewalls struct {
	Enforced                  bool
	KoideNativeYukawa         bool
	FNNativeCharge            bool
	BTauNativeUnification     bool
	PatternFitIsTraceAtomData bool
	ExternalLedgerNative      bool
	NEffGenerationTheorem     bool
	NEffD4Triality            bool
	CHiggsLevelC              bool
	TreeProxyPoleMass         bool
	Verdict                   string
}

type Analysis struct {
	Gate796        Gate796Inheritance
	Checklist      InputChecklist
	AtomProtocol   AtomConstructionProtocol
	Validation     AggregateValidationProtocol
	Koide          PatternDiagnostic
	FN             PatternDiagnostic
	BTau           PatternDiagnostic
	Priority       PatternPriority
	Impact         CHiggsImpact
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
		Gate796: Gate796Inheritance{Inherited: true, Verdict: StatusGate796Inherited},
		Checklist: InputChecklist{
			Defined:          true,
			Fields:           []string{"source_label", "scale_mu", "scheme", "Yukawa_normalization", "color_convention", "neutrino_convention", "y_u,y_c,y_t", "y_d,y_s,y_b", "y_e,y_mu,y_tau", "optional y_nu1,y_nu2,y_nu3", "uncertainties", "conversion_notes"},
			TargetScale:      "M_Z unless multi-scale ledger supplied",
			RequiredStatuses: []string{"scale_mu required", "scheme required or explicitly unknown", "normalization required", "color_convention explicit", "neutrino_convention explicit", "uncertainties required or explicitly absent"},
			RejectImplicit:   true,
			Verdict:          StatusChecklistDefined,
		},
		AtomProtocol: AtomConstructionProtocol{
			Defined:                true,
			AtomFormula:            "x_f=y_f^2; x_f^2=y_f^4",
			CoefficientColorRules:  []string{"a_u=3(sum y_u^2,y_c^2,y_t^2)", "a_d=3(sum y_d^2,y_s^2,y_b^2)", "b_u=3(sum y_u^4,y_c^4,y_t^4)", "b_d=3(sum y_d^4,y_s^4,y_b^4)"},
			Computes:               []string{"a_ext", "b_ext", "N_eff_ext", "C_Yukawa_ext"},
			RequiresConventionLock: true,
			Verdict:                StatusAtomProtocolDefined,
		},
		Validation: AggregateValidationProtocol{
			Defined:              true,
			InheritedA:           aInherited,
			InheritedB:           bInherited,
			InheritedNEff:        nEffInherited,
			ClassifiedFailures:   []string{"scale mismatch", "scheme mismatch", "normalization mismatch", "neutrino convention mismatch", "color counting mismatch", "external ledger not same as inherited ASHA ledger"},
			SilentRescaleAllowed: false,
			Verdict:              StatusAggregateValidation,
		},
		Koide: PatternDiagnostic{
			Name:                  "Koide charged-lepton diagnostic",
			Defined:               true,
			Requires:              []string{"y_e", "y_mu", "y_tau", "declared convention"},
			DiagnosticFormulae:    []string{"Q_e=(y_e+y_mu+y_tau)/(sqrt(y_e)+sqrt(y_mu)+sqrt(y_tau))^2", "compare Q_e to 2/3"},
			AllowedUse:            []string{"charged-lepton pattern diagnostic", "hidden generation geometry motivation"},
			BlockedUse:            []string{"populate Yukawa atom ledger", "derive y_e,y_mu,y_tau", "derive N_eff", "prove D4/triality", "source PMNS/CKM"},
			ReadOnly:              true,
			CanPopulateAtomLedger: false,
			Verdict:               StatusKoideAirlock,
		},
		FN: PatternDiagnostic{
			Name:                  "Froggatt-Nielsen hierarchy diagnostic",
			Defined:               true,
			Requires:              []string{"full sector Yukawa ledger", "declared epsilon", "sector normalization"},
			DiagnosticFormulae:    []string{"n_f=log(y_f/y_t_reference)/log(epsilon)", "sector-specific normalized variants allowed if declared"},
			AllowedUse:            []string{"hierarchy power classification", "charge-assignment motivation"},
			BlockedUse:            []string{"derive Yukawa values", "choose epsilon to force fit without declaration", "native charge theorem without typed FN package"},
			ReadOnly:              true,
			CanPopulateAtomLedger: false,
			Verdict:               StatusFNAirlock,
		},
		BTau: PatternDiagnostic{
			Name:                  "b-tau high-scale unification diagnostic",
			Defined:               true,
			Requires:              []string{"y_b(mu_i)", "y_tau(mu_i)", "RG scheme", "threshold convention", "scale grid or high-scale value", "uncertainty model"},
			DiagnosticFormulae:    []string{"R_btau(mu)=y_b(mu)/y_tau(mu)", "test R_btau(mu_GUT)≈1"},
			AllowedUse:            []string{"high-scale pattern diagnostic", "unification boundary motivation"},
			BlockedUse:            []string{"single-scale proof", "derive full Yukawa ledger", "derive N_eff(M_Z) without RG transport"},
			ReadOnly:              true,
			CanPopulateAtomLedger: false,
			Verdict:               StatusBTauAirlock,
		},
		Priority:       PatternPriority{Recorded: true, Ranking: []string{"full sector/atom ledger", "Froggatt-Nielsen hierarchy diagnostic", "Koide charged-lepton diagnostic", "b-tau unification diagnostic"}, Verdict: StatusPatternPriority},
		Impact:         CHiggsImpact{Recorded: true, ValidatedLedgerCanConfirm: true, PatternsModifyFormula: false, CHiggsLevel: "Level B", CurrentCHiggs: cHiggsLevelB, Verdict: StatusCHiggsImpact},
		Branch:         BranchDecision{Recorded: true, Recommended: "Gate 798 — Yukawa Pattern Diagnostics Holding Pattern and Native-Source Branch Audit", ExternalLedger: false, NativeYukawa: false, D4Package: false, Alternatives: []string{"External Yukawa Ledger Validation and Sector Contribution Audit", "Native Yukawa Operator Trace Atom Extraction Audit", "D4 Triality Trilinear Coupling and Yukawa Trace Readout Audit"}, Verdict: StatusBranchDecision},
		Firewalls:      Firewalls{Enforced: true, KoideNativeYukawa: false, FNNativeCharge: false, BTauNativeUnification: false, PatternFitIsTraceAtomData: false, ExternalLedgerNative: false, NEffGenerationTheorem: false, NEffD4Triality: false, CHiggsLevelC: false, TreeProxyPoleMass: false, Verdict: StatusFirewallPreservedGate797},
		Truth:          "External Yukawa pattern diagnostics are read-only tests on convention-locked supplied data; they do not generate trace atoms.",
		FinalStatement: "Gate 797 does not import or derive Yukawa atoms. It defines the exact external Yukawa input checklist and classifies Koide, Froggatt-Nielsen, and b-tau lanes as secondary read-only diagnostics. A full validated atom ledger remains the primary need before N_eff can be sector-audited.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate796Inherited, StatusPatternsReadOnly, StatusChecklistDefined, StatusAtomProtocolDefined, StatusAggregateValidation, StatusKoideAirlock, StatusFNAirlock, StatusBTauAirlock, StatusPatternPriority, StatusCHiggsImpact, StatusBranchDecision, StatusPhysicalFirewalls,
		StatusExternalLedgerAfterLock, StatusKoideAfterLedger, StatusFNAfterLedger, StatusBTauRequiresRG, StatusFullAtomLedgerPrimary, StatusDiagnosticsSecondary, StatusValidatedUpdatesCYukawa,
		StatusImplicitScaleRejected, StatusColorNoDoubleCount, StatusNoSilentRenormalization, StatusFailedValidationBlocks, StatusKoideNotNative, StatusKoideNoBacksolve, StatusKoideNotFullNEff, StatusFNNotNative, StatusFNPowersNoInvent, StatusNoNativeFNCharges, StatusSingleScaleNoBTau, StatusBTauNotNative, StatusRGThresholdRequired, StatusDiagnosticsDoNotModifyCHiggs, StatusCHiggsRemainsLevelB, StatusFirewallPreservedGate797,
	}
}

func FormatChecklist(c InputChecklist) string {
	return fmt.Sprintf("target=%s fields=%s statuses=%s", c.TargetScale, strings.Join(c.Fields, ","), strings.Join(c.RequiredStatuses, ";"))
}

func FormatAtomProtocol(p AtomConstructionProtocol) string {
	return fmt.Sprintf("%s computes=%s", p.AtomFormula, strings.Join(p.Computes, ","))
}

func FormatValidation(v AggregateValidationProtocol) string {
	return fmt.Sprintf("a=%.16g b=%.16g N_eff=%.16g failures=%s silent_rescale=%v", v.InheritedA, v.InheritedB, v.InheritedNEff, strings.Join(v.ClassifiedFailures, ","), v.SilentRescaleAllowed)
}

func FormatPattern(d PatternDiagnostic) string {
	return fmt.Sprintf("%s requires=%s formulae=%s read_only=%v populates_atoms=%v", d.Name, strings.Join(d.Requires, ","), strings.Join(d.DiagnosticFormulae, ";"), d.ReadOnly, d.CanPopulateAtomLedger)
}

func FormatPriority(p PatternPriority) string {
	return strings.Join(p.Ranking, " > ")
}

func FormatImpact(i CHiggsImpact) string {
	return fmt.Sprintf("C_Higgs=%.16g level=%s patterns_modify_formula=%v", i.CurrentCHiggs, i.CHiggsLevel, i.PatternsModifyFormula)
}

func FormatBranch(b BranchDecision) string {
	return fmt.Sprintf("recommended=%s alternatives=%s", b.Recommended, strings.Join(b.Alternatives, ";"))
}

func containsAll(haystack []string, needles []string) bool {
	joined := strings.ToLower(strings.Join(haystack, "\n"))
	for _, n := range needles {
		if !strings.Contains(joined, strings.ToLower(n)) {
			return false
		}
	}
	return true
}

func closeAbs(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}
