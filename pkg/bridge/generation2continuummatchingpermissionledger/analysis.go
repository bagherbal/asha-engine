// Package generation2continuummatchingpermissionledger implements Gate 504:
// Continuum Matching Permission Ledger for Electroweak Scales.
//
// Gate 503 promoted the photon kernel and rank-three broken electroweak orbit
// to a conditional representation-index theorem, provided a nonzero Higgs ray
// is present.  It deliberately did not derive the nonzero ray, Higgs VEV,
// gauge couplings, physical weak angle, or W/Z masses.  Gate 504 therefore
// constructs the permission ledger that allows those quantities to enter only
// through an explicit continuum/environmental bridge adapter, never as native
// finite-geometry registry writes.
package generation2continuummatchingpermissionledger

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewkernelindexclosure"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2yukawatracescalarnormalization"
)

const (
	AuditID = "GATE504-CONTINUUM-MATCHING-PERMISSION-LEDGER-AUDIT"

	StatusGate503IndexInherited              = "CONDITIONAL_SUPPORT_GATE503_REPRESENTATION_INDEX_INHERITED"
	StatusGate501YukawaTraceAirlockInherited = "CONDITIONAL_SUPPORT_GATE501_YUKAWA_TRACE_AIRLOCK_INHERITED"
	StatusPermissionLedgerConstructed        = "CONDITIONAL_SUPPORT_CONTINUUM_MATCHING_PERMISSION_LEDGER_CONSTRUCTED"
	StatusBridgeInputSchemaDefined           = "CONDITIONAL_SUPPORT_ELECTROWEAK_BRIDGE_INPUT_SCHEMA_DEFINED"
	StatusTreeLevelWZFormulaBridgeOnly       = "CONDITIONAL_SUPPORT_TREE_LEVEL_WZ_FORMULA_ALLOWED_AS_BRIDGE_ONLY"
	StatusPhotonZeroSymbolicPreserved        = "CONDITIONAL_SUPPORT_PHOTON_ZERO_MODE_PRESERVED_SYMBOLICALLY"
	StatusNoNumericAdapterExecuted           = "CONDITIONAL_SUPPORT_NO_NUMERICAL_ELECTROWEAK_ADAPTER_EXECUTED"
	StatusSyntheticNextGateDefined           = "CONDITIONAL_SUPPORT_GATE505_SYNTHETIC_ELECTROWEAK_MATCHING_ADAPTER_REDIRECT_DEFINED"

	StatusFailedNoNativeVEVSelection          = "FAILED_ROUTE_NO_NATIVE_HIGGS_VEV_SELECTION"
	StatusFailedNoNativeGaugeCouplings        = "FAILED_ROUTE_NO_NATIVE_GAUGE_COUPLING_SELECTION"
	StatusFailedWeakAngleNotNative            = "FAILED_ROUTE_PHYSICAL_WEAK_MIXING_ANGLE_NOT_NATIVE"
	StatusFailedWZMassesNotNative             = "FAILED_ROUTE_WZ_MASSES_NOT_NATIVE_DERIVED"
	StatusFailedKappaStillBridge              = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_AFTER_PERMISSION_LEDGER"
	StatusFailedYukawaTraceStillEnvironmental = "FAILED_ROUTE_YUKAWA_TRACE_A_STILL_ENVIRONMENTAL_FOR_SCALAR_NORMALIZATION"

	StatusFirewallNoNumericalDataImported = "FIREWALL_PRESERVED_NO_NUMERICAL_ELECTROWEAK_DATA_IMPORTED"
	StatusFirewallNativeWriteBlocked      = "FIREWALL_BLOCKED_ELECTROWEAK_MATCHING_DATA_NATIVE_WRITE"
)

type Inheritance struct {
	Executed                         bool
	Gate503AuditDefined              bool
	Gate503ConditionalIndexAccepted  bool
	Gate503NonzeroRayAssumed         bool
	Gate503UnconditionalVacuumProven bool
	Gate503WZMassDerived             bool
	Gate503NativeWriteBlocked        bool
	Gate501AuditDefined              bool
	YukawaTraceBridgeScalarNorm      bool
	YukawaTraceNativeNumeric         bool
	ScalarNormalizationEnvironmental bool
	Verdict                          string
	Reason                           string
}

type PermissionRow struct {
	Name                               string
	Category                           string
	BridgePermitted                    bool
	NativePermitted                    bool
	RequiresExplicitValue              bool
	RequiresRenormalizationScale       bool
	RequiresSchemeConvention           bool
	CanBeComputedFromOtherBridgeInputs bool
	FormulaOrUse                       string
	Reason                             string
}

type InputSchema struct {
	Executed                    bool
	Rows                        []PermissionRow
	RequiredMetadata            []string
	BridgeRows                  int
	NativeRows                  int
	RowsRequiringExplicitValues int
	RowsRequiringSchemeScale    int
	Verdict                     string
	Reason                      string
}

type FormulaLedger struct {
	Executed                       bool
	TreeLevelWZFormulasDefined     bool
	RequiresExplicitVEV            bool
	RequiresExplicitGaugeCouplings bool
	ComputesNow                    bool
	PhotonZeroSymbolic             bool
	RhoTreeSymbolic                bool
	WeakAngleBridgeFormula         string
	WMassBridgeFormula             string
	ZMassBridgeFormula             string
	NativeWeakAngleDerived         bool
	NativeWZMassesDerived          bool
	NativeKappaPromoted            bool
	Verdict                        string
	Reason                         string
}

type Boundary struct {
	Executed                                     bool
	PermissionLedgerAccepted                     bool
	ContinuumAdapterMayComputeWithExplicitInputs bool
	NumericalAdapterExecuted                     bool
	NativeVEVSelected                            bool
	NativeGaugeCouplingsSelected                 bool
	NativeWeakAngleDerived                       bool
	NativeWZMassesDerived                        bool
	NativeKappaSelected                          bool
	YukawaTraceStillEnvironmental                bool
	Verdict                                      string
	Reason                                       string
}

type Firewall struct {
	Executed                       bool
	ObservedVEVImported            bool
	ObservedGaugeCouplingsImported bool
	ObservedWeakAngleImported      bool
	ObservedWMassImported          bool
	ObservedZMassImported          bool
	ObservedYukawaImported         bool
	NativeVEVWritten               bool
	NativeGaugeCouplingWritten     bool
	NativeWeakAngleWritten         bool
	NativeWZMassWritten            bool
	NativeKappaWritten             bool
	Verdict                        string
	Reason                         string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Schema      InputSchema
	Formula     FormulaLedger
	Boundary    Boundary
	Firewall    Firewall
	Registry    RegistryUpdate
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
	g503, err := generation2ewkernelindexclosure.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate503 kernel index closure: %w", err)
	}
	g501, err := generation2yukawatracescalarnormalization.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate501 Yukawa trace airlock: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g503, g501)
	a.Schema = buildSchema()
	a.Formula = buildFormulaLedger()
	a.Boundary = buildBoundary(a.Schema, a.Formula, a.Inheritance)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g503 generation2ewkernelindexclosure.Analysis, g501 generation2yukawatracescalarnormalization.Analysis) Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate503AuditDefined:              true,
		Gate503ConditionalIndexAccepted:  g503.Boundary.ConditionalRepresentationIndexAccepted,
		Gate503NonzeroRayAssumed:         g503.Kernel.ConditionalOnNonzeroRay,
		Gate503UnconditionalVacuumProven: g503.Kernel.UnconditionalNativeVacuumProvenance,
		Gate503WZMassDerived:             g503.Boundary.NativeWZMassMatrixDerived,
		Gate503NativeWriteBlocked:        g503.Firewall.NativeWZMassWritten == false && g503.Boundary.NativeWZMassMatrixDerived == false,
		Gate501AuditDefined:              true,
		YukawaTraceBridgeScalarNorm:      g501.Decision.TraceABridgeScalarNormAccepted,
		YukawaTraceNativeNumeric:         g501.Decision.TraceANativeNumericAccepted,
		ScalarNormalizationEnvironmental: !g501.Decision.ScalarKineticCoefficientNative,
		Verdict:                          strings.Join([]string{StatusGate503IndexInherited, StatusGate501YukawaTraceAirlockInherited}, ";"),
		Reason:                           "Gate503 supplies only a conditional Higgs-doublet representation index; Gate501 confirms the scalar normalization trace a is bridge/environmental, not a native numeric coefficient.",
	}
}

func buildSchema() InputSchema {
	rows := []PermissionRow{
		{Name: "Higgs vacuum expectation value v", Category: "continuum/environmental scale", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: true, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, FormulaOrUse: "sets the electroweak scale in bridge computations", Reason: "the finite core has not selected a nonzero vacuum ray or its magnitude"},
		{Name: "SU(2)_L gauge coupling g2", Category: "continuum running coupling", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: true, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, FormulaOrUse: "m_W = g2 v / 2 in the bridge adapter", Reason: "finite representation indices do not determine continuum coupling units"},
		{Name: "U(1)_Y gauge coupling gY", Category: "continuum running coupling", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: true, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, FormulaOrUse: "m_Z = sqrt(g2^2+gY^2) v / 2 in the bridge adapter", Reason: "hypercharge trace normalization is not a physical low-scale coupling value"},
		{Name: "physical weak mixing angle sin^2(theta_W)", Category: "derived continuum bridge quantity", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: false, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, CanBeComputedFromOtherBridgeInputs: true, FormulaOrUse: "sin^2(theta_W)=gY^2/(g2^2+gY^2)", Reason: "the finite sin^2=3/8 boundary diagnostic is not the physical renormalized low-scale angle"},
		{Name: "W and Z pole or running masses", Category: "continuum comparator/output", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: false, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, CanBeComputedFromOtherBridgeInputs: true, FormulaOrUse: "tree-level bridge formulas only after explicit v,g2,gY inputs", Reason: "Gate503 gives rank and kernel, not mass eigenvalues in GeV"},
		{Name: "Yukawa trace a = Tr(Y†Y)", Category: "scalar-normalization bridge coefficient", BridgePermitted: true, NativePermitted: false, RequiresExplicitValue: true, RequiresRenormalizationScale: true, RequiresSchemeConvention: true, FormulaOrUse: "K_phi = f0 a / pi^2", Reason: "a is invariant but depends on sealed Yukawa amplitude history"},
	}
	bridge, native, explicit, scheme := 0, 0, 0, 0
	for _, r := range rows {
		if r.BridgePermitted {
			bridge++
		}
		if r.NativePermitted {
			native++
		}
		if r.RequiresExplicitValue {
			explicit++
		}
		if r.RequiresRenormalizationScale || r.RequiresSchemeConvention {
			scheme++
		}
	}
	return InputSchema{Executed: true, Rows: rows, RequiredMetadata: []string{"renormalization scale μ", "renormalization scheme/convention", "tree/running/pole interpretation", "source tag marking data as bridge/environmental"}, BridgeRows: bridge, NativeRows: native, RowsRequiringExplicitValues: explicit, RowsRequiringSchemeScale: scheme, Verdict: strings.Join([]string{StatusPermissionLedgerConstructed, StatusBridgeInputSchemaDefined}, ";"), Reason: "Every electroweak scale/coupling/mass quantity is permitted only as an explicit bridge datum or bridge output with scale and scheme metadata; zero rows are native-permitted."}
}

func buildFormulaLedger() FormulaLedger {
	return FormulaLedger{
		Executed:                       true,
		TreeLevelWZFormulasDefined:     true,
		RequiresExplicitVEV:            true,
		RequiresExplicitGaugeCouplings: true,
		ComputesNow:                    false,
		PhotonZeroSymbolic:             true,
		RhoTreeSymbolic:                true,
		WeakAngleBridgeFormula:         "sin^2(theta_W)=gY^2/(g2^2+gY^2)",
		WMassBridgeFormula:             "m_W = g2 v / 2",
		ZMassBridgeFormula:             "m_Z = sqrt(g2^2 + gY^2) v / 2",
		NativeWeakAngleDerived:         false,
		NativeWZMassesDerived:          false,
		NativeKappaPromoted:            false,
		Verdict:                        strings.Join([]string{StatusTreeLevelWZFormulaBridgeOnly, StatusPhotonZeroSymbolicPreserved, StatusNoNumericAdapterExecuted}, ";"),
		Reason:                         "The formulas are the permitted continuum bridge map compatible with the Gate503 kernel/rank index.  No numeric adapter is executed because no explicit scale, scheme, VEV, or couplings were supplied by this gate.",
	}
}

func buildBoundary(s InputSchema, f FormulaLedger, i Inheritance) Boundary {
	accepted := s.Executed && s.BridgeRows == len(s.Rows) && s.NativeRows == 0 && f.TreeLevelWZFormulasDefined && !f.ComputesNow && i.Gate503ConditionalIndexAccepted
	return Boundary{
		Executed:                 true,
		PermissionLedgerAccepted: accepted,
		ContinuumAdapterMayComputeWithExplicitInputs: accepted,
		NumericalAdapterExecuted:                     f.ComputesNow,
		NativeVEVSelected:                            false,
		NativeGaugeCouplingsSelected:                 false,
		NativeWeakAngleDerived:                       false,
		NativeWZMassesDerived:                        false,
		NativeKappaSelected:                          false,
		YukawaTraceStillEnvironmental:                i.ScalarNormalizationEnvironmental && !i.YukawaTraceNativeNumeric,
		Verdict:                                      strings.Join([]string{StatusPermissionLedgerConstructed, StatusFailedNoNativeVEVSelection, StatusFailedNoNativeGaugeCouplings, StatusFailedWeakAngleNotNative, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge, StatusFailedYukawaTraceStillEnvironmental}, ";"),
		Reason:                                       "Gate504 opens a controlled bridge permission door but keeps the native theorem registry closed for VEV, couplings, weak angle, kappa, Yukawa trace value, and W/Z masses.",
	}
}

func buildFirewall() Firewall {
	return Firewall{Executed: true, Verdict: strings.Join([]string{StatusFirewallNoNumericalDataImported, StatusFirewallNativeWriteBlocked}, ";"), Reason: "No numerical VEV, gauge coupling, weak angle, W/Z mass, or Yukawa value is imported; none is written as native."}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries:        []string{"No new electroweak scale, coupling, weak-angle, kappa, VEV, or W/Z mass theorem is admitted at Gate504."},
		BridgeEntries:        []string{"A continuum matching permission ledger is admitted: explicit v, g2, gY, sin^2(theta_W), W/Z comparator masses, and Yukawa trace a may be used only with bridge/environmental tags and scale/scheme metadata.", "The symbolic tree-level bridge map m_W=g2 v/2, m_Z=sqrt(g2^2+gY^2)v/2, sin^2(theta_W)=gY^2/(g2^2+gY^2), and m_gamma=0 is permitted for explicit bridge adapters only."},
		EnvironmentalEntries: []string{"Higgs VEV, physical gauge couplings, low-scale weak angle, W/Z pole or running masses, Yukawa amplitudes, CKM, and PMNS remain environmental/continuum matching data."},
		FailedRoutes:         []string{StatusFailedNoNativeVEVSelection, StatusFailedNoNativeGaugeCouplings, StatusFailedWeakAngleNotNative, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge, StatusFailedYukawaTraceStillEnvironmental},
		OpenTheorems:         []string{"Implement a synthetic electroweak matching adapter that accepts explicit non-observed test inputs and computes bridge-only W/Z outputs without native promotion.", "A separate finite-action theorem would still be required to select a nonzero Higgs ray, kappa_U1, or gauge Hessian natively."},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 505, Title: "Synthetic Electroweak Matching Adapter Dry-Run", Reason: "Gate504 defines the permission schema; the next safe step is to test the adapter with synthetic bridge inputs, not observed data.", PrimaryTask: "execute a synthetic-only electroweak bridge adapter that computes m_W, m_Z, m_gamma, and sin^2(theta_W) from explicit fake v,g2,gY inputs while proving no native registry promotion occurs"}
}

func truth(a Analysis) string {
	if a.Boundary.PermissionLedgerAccepted && !a.Boundary.NativeWZMassesDerived {
		return "Gate504 establishes the electroweak continuum-matching airlock: ASHA may use VEV, gauge couplings, physical weak angle, W/Z masses, and Yukawa trace normalization only as explicitly tagged bridge/environmental data with scale and scheme metadata.  The native core keeps only representation topology, photon kernel, and rank-three broken orbit; it does not derive electroweak scales, couplings, kappa_U1, or physical masses."
	}
	return "Gate504 did not construct the electroweak continuum-matching permission ledger."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate503ConditionalIndexAccepted && !a.Inheritance.Gate503WZMassDerived && a.Inheritance.Gate501AuditDefined, "required Gate503/Gate501 inheritance missing or over-promoted"},
		{a.Schema.Executed && a.Schema.BridgeRows == len(a.Schema.Rows) && a.Schema.NativeRows == 0 && a.Schema.RowsRequiringExplicitValues >= 3 && a.Schema.RowsRequiringSchemeScale == len(a.Schema.Rows), "permission schema invalid"},
		{a.Formula.Executed && a.Formula.TreeLevelWZFormulasDefined && a.Formula.RequiresExplicitVEV && a.Formula.RequiresExplicitGaugeCouplings && !a.Formula.ComputesNow && a.Formula.PhotonZeroSymbolic && !a.Formula.NativeWeakAngleDerived && !a.Formula.NativeWZMassesDerived, "formula ledger invalid or over-promoted"},
		{a.Boundary.Executed && a.Boundary.PermissionLedgerAccepted && a.Boundary.ContinuumAdapterMayComputeWithExplicitInputs && !a.Boundary.NumericalAdapterExecuted && !a.Boundary.NativeVEVSelected && !a.Boundary.NativeGaugeCouplingsSelected && !a.Boundary.NativeWeakAngleDerived && !a.Boundary.NativeWZMassesDerived && !a.Boundary.NativeKappaSelected, "boundary invalid or over-promoted"},
		{a.Firewall.Executed && !a.Firewall.ObservedVEVImported && !a.Firewall.ObservedGaugeCouplingsImported && !a.Firewall.ObservedWeakAngleImported && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedZMassImported && !a.Firewall.ObservedYukawaImported && !a.Firewall.NativeVEVWritten && !a.Firewall.NativeGaugeCouplingWritten && !a.Firewall.NativeWeakAngleWritten && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeKappaWritten, "firewall violation"},
		{a.Next.Gate == 505, "Gate505 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate503=%t index=%t nonzero_assumed=%t unconditional_vacuum=%t WZ=%t gate501=%t a_bridge=%t a_native=%t scalar_env=%t verdict=%s reason=%s", x.Gate503AuditDefined, x.Gate503ConditionalIndexAccepted, x.Gate503NonzeroRayAssumed, x.Gate503UnconditionalVacuumProven, x.Gate503WZMassDerived, x.Gate501AuditDefined, x.YukawaTraceBridgeScalarNorm, x.YukawaTraceNativeNumeric, x.ScalarNormalizationEnvironmental, x.Verdict, x.Reason)
}
func FormatSchema(x InputSchema) string {
	return fmt.Sprintf("rows=%d bridge=%d native=%d explicit=%d scheme_scale=%d metadata=[%s] verdict=%s reason=%s", len(x.Rows), x.BridgeRows, x.NativeRows, x.RowsRequiringExplicitValues, x.RowsRequiringSchemeScale, strings.Join(x.RequiredMetadata, "; "), x.Verdict, x.Reason)
}
func FormatFormula(x FormulaLedger) string {
	return fmt.Sprintf("tree=%t v=%t couplings=%t computed=%t photon0=%t rho=%t weak=%s W=%s Z=%s native_theta=%t native_WZ=%t native_kappa=%t verdict=%s reason=%s", x.TreeLevelWZFormulasDefined, x.RequiresExplicitVEV, x.RequiresExplicitGaugeCouplings, x.ComputesNow, x.PhotonZeroSymbolic, x.RhoTreeSymbolic, x.WeakAngleBridgeFormula, x.WMassBridgeFormula, x.ZMassBridgeFormula, x.NativeWeakAngleDerived, x.NativeWZMassesDerived, x.NativeKappaPromoted, x.Verdict, x.Reason)
}
func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("accepted=%t may_compute=%t computed=%t native_v=%t native_gauge=%t native_theta=%t native_WZ=%t native_kappa=%t a_env=%t verdict=%s reason=%s", x.PermissionLedgerAccepted, x.ContinuumAdapterMayComputeWithExplicitInputs, x.NumericalAdapterExecuted, x.NativeVEVSelected, x.NativeGaugeCouplingsSelected, x.NativeWeakAngleDerived, x.NativeWZMassesDerived, x.NativeKappaSelected, x.YukawaTraceStillEnvironmental, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("obs_v=%t obs_gauge=%t obs_theta=%t obs_W=%t obs_Z=%t obs_yukawa=%t native_v=%t native_gauge=%t native_theta=%t native_WZ=%t native_kappa=%t verdict=%s reason=%s", x.ObservedVEVImported, x.ObservedGaugeCouplingsImported, x.ObservedWeakAngleImported, x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedYukawaImported, x.NativeVEVWritten, x.NativeGaugeCouplingWritten, x.NativeWeakAngleWritten, x.NativeWZMassWritten, x.NativeKappaWritten, x.Verdict, x.Reason)
}
func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 504 Registry Audit — Continuum Matching Permission Ledger for Electroweak Scales\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{StatusGate503IndexInherited, StatusGate501YukawaTraceAirlockInherited, StatusPermissionLedgerConstructed, StatusBridgeInputSchemaDefined, StatusTreeLevelWZFormulaBridgeOnly, StatusPhotonZeroSymbolicPreserved, StatusNoNumericAdapterExecuted, StatusFailedNoNativeVEVSelection, StatusFailedNoNativeGaugeCouplings, StatusFailedWeakAngleNotNative, StatusFailedWZMassesNotNative, StatusFailedKappaStillBridge, StatusFailedYukawaTraceStillEnvironmental, StatusFirewallNoNumericalDataImported, StatusFirewallNativeWriteBlocked, StatusSyntheticNextGateDefined} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate503 gives a conditional representation-index theorem: for a finite one-form Higgs doublet and a nonzero Higgs ray, `U(1)_em` is the one-dimensional stabilizer, the broken electroweak orbit has dimension three, and one radial scalar quotient remains.  Gate503 does not select the nonzero ray, its VEV, `kappa_U1`, gauge couplings, weak angle, or W/Z masses.\n\n")
	b.WriteString("Gate501 confirms that `a = Tr(Y†Y)` is a basis/rephasing-invariant scalar normalization trace, but its numeric value depends on sealed Yukawa amplitudes.\n\n")
	b.WriteString("## Permission ledger\n\n")
	b.WriteString(fmt.Sprintf("```text\nbridge-permitted rows = %d\nnative-permitted rows = %d\nrows requiring explicit values = %d\nrows requiring scale/scheme metadata = %d\nrequired metadata = %s\n```\n\n", a.Schema.BridgeRows, a.Schema.NativeRows, a.Schema.RowsRequiringExplicitValues, a.Schema.RowsRequiringSchemeScale, strings.Join(a.Schema.RequiredMetadata, "; ")))
	b.WriteString("| Quantity | Category | Bridge? | Native? | Requirement |\n")
	b.WriteString("|---|---:|---:|---:|---|\n")
	for _, r := range a.Schema.Rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %s |\n", r.Name, r.Category, r.BridgePermitted, r.NativePermitted, r.Reason))
	}
	b.WriteString("\n## Bridge formula ledger\n\n")
	b.WriteString(fmt.Sprintf("```text\n%s\n%s\n%s\nm_gamma = 0  symbolic photon kernel\nrho_tree = 1 symbolic doublet bridge identity\ncomputed now = %t\nnative weak angle derived = %t\nnative W/Z masses derived = %t\nnative kappa_U1 promoted = %t\n```\n\n", a.Formula.WMassBridgeFormula, a.Formula.ZMassBridgeFormula, a.Formula.WeakAngleBridgeFormula, a.Formula.ComputesNow, a.Formula.NativeWeakAngleDerived, a.Formula.NativeWZMassesDerived, a.Formula.NativeKappaPromoted))
	b.WriteString("These formulas are allowed only inside an explicit continuum adapter with bridge/environmental tags.  They are not ASHA-native mass or coupling predictions.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("No numerical VEV, gauge coupling, weak angle, W/Z mass, or Yukawa value is imported.  No native registry write is made for VEV, gauge couplings, weak angle, W/Z masses, `kappa_U1`, or Yukawa trace value.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate505 should be:\n\n```text\nGate 505 — Synthetic Electroweak Matching Adapter Dry-Run\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
