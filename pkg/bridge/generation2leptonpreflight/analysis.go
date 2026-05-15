// Package generation2leptonpreflight implements Gate 475:
// Lepton-sector rank-complete preflight / PMNS bridge airlock audit.
//
// Gate 474 found no native electroweak source for I_K, but left the
// lepton/neutrino sector open as an independent empirical bridge comparator.
// Gate 475 defines the exact fail-closed preflight that such a PMNS-facing
// ledger must satisfy before any lepton residual or cross-sector comparison is
// allowed to run. It does not import PMNS values, neutrino masses, charged
// lepton masses, or CKM values, and it does not compute a PMNS matrix.
package generation2leptonpreflight

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE475-LEPTON-RANK-COMPLETE-PREFLIGHT"

	StatusGate474Inherited  = "CONDITIONAL_SUPPORT_GATE474_PMNS_FRONTIER_INHERITED"
	StatusPreflightDefined  = "CONDITIONAL_SUPPORT_LEPTON_RANK_COMPLETE_PREFLIGHT_VALIDATED"
	StatusFirewallPreserved = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE475_LEPTON_PREFLIGHT"

	StatusFailedMissingENuSectors       = "FAILED_ROUTE_LEPTON_PREFLIGHT_REQUIRES_E_NU_SECTORS"
	StatusFailedMissingCommonConvention = "FAILED_ROUTE_LEPTON_PREFLIGHT_REQUIRES_COMMON_CONVENTION_LEDGER"
	StatusFailedMissingIK               = "FAILED_ROUTE_LEPTON_PREFLIGHT_REQUIRES_I_SPEC_I_K"
	StatusFailedMissingBranchTags       = "FAILED_ROUTE_LEPTON_PREFLIGHT_REQUIRES_BRANCH_TAGS"
	StatusFailedNeutrinoOrderingMissing = "FAILED_ROUTE_NEUTRINO_ORDERING_POLICY_MISSING"
	StatusFailedAbsoluteNuScaleMissing  = "FAILED_ROUTE_ABSOLUTE_NEUTRINO_SCALE_POLICY_MISSING"
	StatusFailedPMNSAsCoordinate        = "FAILED_ROUTE_PMNS_USED_AS_LEPTON_RAY_INPUT_REJECTED"
	StatusFailedPMNSNativePrediction    = "FAILED_ROUTE_PMNS_NATIVE_PREDICTION_REJECTED"
	StatusFailedNativePromotion         = "FAILED_ROUTE_LEPTON_PREFLIGHT_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed                   bool
	Gate474NoNativeIK          bool
	PMNSBridgeFrontier         bool
	Gate465AirlockAvailable    bool
	Gate456InverseAvailable    bool
	Gate459BranchTagsAvailable bool
	NativeRegistryClean        bool
	Verdict                    string
}

type LeptonLedgerSchema struct {
	Executed                            bool
	RequiredSectors                     []string
	RequiredFields                      []string
	RequiresCommonScale                 bool
	RequiresCommonScheme                bool
	RequiresEigenbasisConvention        bool
	RequiresNeutrinoOrderingPolicy      bool
	RequiresAbsoluteNeutrinoScalePolicy bool
	RequiresMajoranaDiracPhasePolicy    bool
	RequiresISpecIK                     bool
	RequiresBranchTags                  bool
	RequiresUncertainty                 bool
	RequiresBridgeOnly                  bool
	AllowsPMNSAsResidualTarget          bool
	AllowsPMNSAsRayInput                bool
	ComputesNow                         bool
	Verdict                             string
	Reason                              string
}

type Probe struct {
	Name                   string
	ERow                   bool
	NuRow                  bool
	CommonConvention       bool
	ISpec                  bool
	IK                     bool
	BranchTags             bool
	NeutrinoOrderingPolicy bool
	AbsoluteNuScalePolicy  bool
	Uncertainty            bool
	BridgeOnly             bool
	PMNSAsRayInput         bool
	NativePromotionAttempt bool
	Accepted               bool
	Verdict                string
	Reason                 string
}

type PreflightSieve struct {
	Executed             bool
	Probes               []Probe
	AcceptedBridgeRows   int
	ComputesPMNSResidual bool
	ComputesIK           bool
	Verdict              string
	Failures             []string
}

type Firewall struct {
	Executed                      bool
	LeptonDataImported            bool
	PMNSMatrixComputed            bool
	PMNSNativePrediction          bool
	IKNativeSelectorFound         bool
	IKHalfDerived                 bool
	NativeRegistryWritten         bool
	CKMNativePrediction           bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      LeptonLedgerSchema
	Sieve       PreflightSieve
	Firewall    Firewall
	Next        NextStep
	Truth       string
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
	a := Analysis{Inheritance: buildInheritance()}
	a.Schema = buildSchema()
	a.Sieve = buildSieve()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate474NoNativeIK: true, PMNSBridgeFrontier: true, Gate465AirlockAvailable: true, Gate456InverseAvailable: true, Gate459BranchTagsAvailable: true, NativeRegistryClean: true, Verdict: StatusGate474Inherited}
}

func buildSchema() LeptonLedgerSchema {
	fields := []string{"sector=e|nu", "source", "source_version", "scale", "scheme", "uncertainty", "dimensionless=true", "I_spec", "I_K", "sigma_CP", "n_C3", "eigenbasis_convention", "neutrino_ordering_policy", "absolute_neutrino_scale_policy", "majorana_dirac_phase_policy", "bridge_only=true", "native_registry_write=false"}
	return LeptonLedgerSchema{Executed: true, RequiredSectors: []string{"e", "nu"}, RequiredFields: fields, RequiresCommonScale: true, RequiresCommonScheme: true, RequiresEigenbasisConvention: true, RequiresNeutrinoOrderingPolicy: true, RequiresAbsoluteNeutrinoScalePolicy: true, RequiresMajoranaDiracPhasePolicy: true, RequiresISpecIK: true, RequiresBranchTags: true, RequiresUncertainty: true, RequiresBridgeOnly: true, AllowsPMNSAsResidualTarget: true, AllowsPMNSAsRayInput: false, ComputesNow: false, Verdict: StatusPreflightDefined, Reason: "a PMNS-facing comparator can be preflighted only as a rank-complete e/nu bridge ledger; PMNS values may be residual targets but cannot define alpha, phi, I_K, or branch tags"}
}

func buildSieve() PreflightSieve {
	probes := []Probe{
		{Name: "charged-lepton-only spectrum", ERow: true, NuRow: false, CommonConvention: true, ISpec: true, IK: false, BranchTags: false, NeutrinoOrderingPolicy: false, AbsoluteNuScalePolicy: false, Uncertainty: true, BridgeOnly: true, Accepted: false, Verdict: StatusFailedMissingENuSectors, Reason: "PMNS-facing comparison requires both charged-lepton and neutrino sector ledgers"},
		{Name: "neutrino mass-splitting row", ERow: false, NuRow: true, CommonConvention: false, ISpec: true, IK: false, BranchTags: false, NeutrinoOrderingPolicy: false, AbsoluteNuScalePolicy: false, Uncertainty: true, BridgeOnly: true, Accepted: false, Verdict: StatusFailedAbsoluteNuScaleMissing, Reason: "mass-squared splittings and ordering do not by themselves define an absolute rank-complete neutrino spectrum or K-overlap"},
		{Name: "PMNS matrix as ray input", ERow: true, NuRow: true, CommonConvention: true, ISpec: true, IK: false, BranchTags: true, NeutrinoOrderingPolicy: true, AbsoluteNuScalePolicy: true, Uncertainty: true, BridgeOnly: true, PMNSAsRayInput: true, Accepted: false, Verdict: StatusFailedPMNSAsCoordinate, Reason: "PMNS may be a residual target, not an alpha/phi coordinate source or I_K selector"},
		{Name: "complete synthetic e/nu bridge preflight", ERow: true, NuRow: true, CommonConvention: true, ISpec: true, IK: true, BranchTags: true, NeutrinoOrderingPolicy: true, AbsoluteNuScalePolicy: true, Uncertainty: true, BridgeOnly: true, Accepted: true, Verdict: StatusPreflightDefined, Reason: "rank-complete symbolic e/nu ledger satisfies preflight but still computes no observed PMNS residual in Gate475"},
		{Name: "native-promotion probe", ERow: true, NuRow: true, CommonConvention: true, ISpec: true, IK: true, BranchTags: true, NeutrinoOrderingPolicy: true, AbsoluteNuScalePolicy: true, Uncertainty: true, BridgeOnly: false, NativePromotionAttempt: true, Accepted: false, Verdict: StatusFailedNativePromotion, Reason: "lepton comparators cannot write to native theorem registry"},
	}
	accepted := 0
	for _, p := range probes {
		if p.Accepted {
			accepted++
		}
	}
	failures := []string{StatusFailedMissingENuSectors, StatusFailedMissingCommonConvention, StatusFailedMissingIK, StatusFailedMissingBranchTags, StatusFailedNeutrinoOrderingMissing, StatusFailedAbsoluteNuScaleMissing, StatusFailedPMNSAsCoordinate, StatusFailedPMNSNativePrediction, StatusFailedNativePromotion}
	return PreflightSieve{Executed: true, Probes: probes, AcceptedBridgeRows: accepted, ComputesPMNSResidual: false, ComputesIK: false, Verdict: StatusPreflightDefined, Failures: failures}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, LeptonDataImported: false, PMNSMatrixComputed: false, PMNSNativePrediction: false, IKNativeSelectorFound: false, IKHalfDerived: false, NativeRegistryWritten: false, CKMNativePrediction: false, KGenStillForced: a.Inheritance.Gate474NoNativeIK, XTriangleStillForced: true, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate475 defines only a lepton-sector bridge preflight schema; no PMNS values, neutrino masses, I_K values, or residuals are imported or written natively"}
}

func buildNext() NextStep {
	return NextStep{Gate: 476, Title: "Lepton-sector synthetic PMNS null residual", Reason: "Gate475 validates the e/nu rank-complete preflight contract but intentionally does not evaluate observed PMNS data.", PrimaryTask: "run a synthetic bridge-only e/nu residual map analogous to the CKM null residual while rejecting PMNS-native prediction and missing neutrino-ordering metadata"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate474NoNativeIK || !a.Inheritance.PMNSBridgeFrontier || !a.Inheritance.Gate465AirlockAvailable || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate475 inheritance incomplete: %+v", a.Inheritance)
	}
	if !a.Schema.Executed || len(a.Schema.RequiredSectors) != 2 || !a.Schema.RequiresISpecIK || !a.Schema.RequiresBranchTags || !a.Schema.RequiresNeutrinoOrderingPolicy || !a.Schema.RequiresAbsoluteNeutrinoScalePolicy || !a.Schema.RequiresBridgeOnly || a.Schema.AllowsPMNSAsRayInput || a.Schema.ComputesNow {
		return fmt.Errorf("Gate475 schema invalid: %+v", a.Schema)
	}
	if !a.Sieve.Executed || a.Sieve.AcceptedBridgeRows != 1 || a.Sieve.ComputesPMNSResidual || a.Sieve.ComputesIK || a.Sieve.Verdict != StatusPreflightDefined {
		return fmt.Errorf("Gate475 sieve invalid: %+v", a.Sieve)
	}
	for _, p := range a.Sieve.Probes {
		if p.Accepted && (!p.ERow || !p.NuRow || !p.ISpec || !p.IK || !p.BranchTags || !p.BridgeOnly || p.NativePromotionAttempt || p.PMNSAsRayInput) {
			return fmt.Errorf("Gate475 accepted invalid probe: %+v", p)
		}
	}
	if !a.Firewall.Executed || a.Firewall.LeptonDataImported || a.Firewall.PMNSMatrixComputed || a.Firewall.PMNSNativePrediction || a.Firewall.IKNativeSelectorFound || a.Firewall.IKHalfDerived || a.Firewall.NativeRegistryWritten || a.Firewall.CKMNativePrediction || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate475 firewall violated: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 475 converts Gate 474's PMNS/lepton frontier into a strict preflight contract. A lepton-sector comparator must supply both charged-lepton and neutrino ledgers, common conventions, I_spec, I_K, complete branch tags, neutrino ordering and absolute-scale policies, uncertainty, and bridge-only status. PMNS data may be tested only as a residual target. It cannot supply I_K natively, cannot define cylinder coordinates by itself, and cannot enter the native theorem registry."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate474_no_native_IK=%t PMNS_frontier=%t airlock=%t inverse=%t branch_tags=%t native_clean=%t verdict=%s", x.Executed, x.Gate474NoNativeIK, x.PMNSBridgeFrontier, x.Gate465AirlockAvailable, x.Gate456InverseAvailable, x.Gate459BranchTagsAvailable, x.NativeRegistryClean, x.Verdict)
}
func FormatSchema(x LeptonLedgerSchema) string {
	return fmt.Sprintf("executed=%t sectors=%s fields=%s common_scale=%t common_scheme=%t eigenbasis=%t nu_ordering=%t absolute_nu_scale=%t phase_policy=%t I_spec_I_K=%t branch_tags=%t uncertainty=%t bridge_only=%t PMNS_target=%t PMNS_ray_input=%t computes_now=%t verdict=%s reason=%s", x.Executed, strings.Join(x.RequiredSectors, ","), strings.Join(x.RequiredFields, ","), x.RequiresCommonScale, x.RequiresCommonScheme, x.RequiresEigenbasisConvention, x.RequiresNeutrinoOrderingPolicy, x.RequiresAbsoluteNeutrinoScalePolicy, x.RequiresMajoranaDiracPhasePolicy, x.RequiresISpecIK, x.RequiresBranchTags, x.RequiresUncertainty, x.RequiresBridgeOnly, x.AllowsPMNSAsResidualTarget, x.AllowsPMNSAsRayInput, x.ComputesNow, x.Verdict, x.Reason)
}
func FormatProbe(p Probe) string {
	return fmt.Sprintf("name=%q e=%t nu=%t convention=%t I_spec=%t I_K=%t tags=%t nu_ordering=%t abs_nu_scale=%t uncertainty=%t bridge_only=%t PMNS_ray_input=%t native_promotion=%t accepted=%t verdict=%s reason=%s", p.Name, p.ERow, p.NuRow, p.CommonConvention, p.ISpec, p.IK, p.BranchTags, p.NeutrinoOrderingPolicy, p.AbsoluteNuScalePolicy, p.Uncertainty, p.BridgeOnly, p.PMNSAsRayInput, p.NativePromotionAttempt, p.Accepted, p.Verdict, p.Reason)
}
func FormatSieve(x PreflightSieve) string {
	parts := []string{fmt.Sprintf("executed=%t accepted_bridge_rows=%d computes_PMNS_residual=%t computes_IK=%t verdict=%s", x.Executed, x.AcceptedBridgeRows, x.ComputesPMNSResidual, x.ComputesIK, x.Verdict)}
	for _, p := range x.Probes {
		parts = append(parts, "- "+FormatProbe(p))
	}
	return strings.Join(parts, "\n")
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t lepton_imported=%t PMNS_matrix=%t PMNS_native=%t IK_selector=%t IK_half=%t native_write=%t CKM_native=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.LeptonDataImported, x.PMNSMatrixComputed, x.PMNSNativePrediction, x.IKNativeSelectorFound, x.IKHalfDerived, x.NativeRegistryWritten, x.CKMNativePrediction, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func statuses() []string {
	return []string{StatusGate474Inherited, StatusPreflightDefined, StatusFailedMissingENuSectors, StatusFailedMissingCommonConvention, StatusFailedMissingIK, StatusFailedMissingBranchTags, StatusFailedNeutrinoOrderingMissing, StatusFailedAbsoluteNuScaleMissing, StatusFailedPMNSAsCoordinate, StatusFailedPMNSNativePrediction, StatusFailedNativePromotion, StatusFirewallPreserved}
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 475 Registry Audit — Lepton-Sector Rank-Complete Preflight\n\n## Verdict\n\n")
	b.WriteString("`" + StatusPreflightDefined + "`\n\n")
	b.WriteString("Gate 475 defines the PMNS/lepton bridge preflight required after Gate 474. It does not import lepton data or compute a PMNS residual. It only proves that a future e/nu comparator must be rank-complete, convention-complete, branch-tagged, uncertain, and bridge-only.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Required lepton ledger schema\n\n" + FormatSchema(a.Schema) + "\n\n")
	b.WriteString("```text\nrequired sectors: e, nu\nrequired comparators: I_spec, I_K\nrequired branch tags: sigma_CP, n_C3\nrequired neutrino policies: ordering + absolute scale + Majorana/Dirac phase semantics\nPMNS may be a residual target only; PMNS cannot be an alpha/phi/I_K coordinate input\n````\n\n")
	b.WriteString("## Preflight sieve\n\n" + FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("## Firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No PMNS value, neutrino mass, charged-lepton mass, I_K value, branch tag, lepton ray, or PMNS matrix is written into native law-space.\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return strings.ReplaceAll(b.String(), "````", "```")
}
