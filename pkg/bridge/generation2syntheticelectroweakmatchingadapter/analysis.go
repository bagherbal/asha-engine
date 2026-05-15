// Package generation2syntheticelectroweakmatchingadapter implements Gate 505:
// Synthetic Electroweak Matching Adapter Dry-Run.
//
// Gate 504 opened a continuum-matching permission ledger for electroweak
// bridge inputs, while blocking all native writes of VEV, couplings, weak
// angle, kappa, and W/Z masses.  Gate 505 executes the adapter once with a
// deliberately fake 3-4-5 input triangle.  The purpose is not prediction; it is
// a firewall audit proving that explicit bridge inputs can be propagated
// through the tree-level electroweak formulas without being mistaken for native
// finite-geometry theorems or observed particle data.
package generation2syntheticelectroweakmatchingadapter

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2continuummatchingpermissionledger"
)

const (
	AuditID = "GATE505-SYNTHETIC-ELECTROWEAK-MATCHING-ADAPTER-DRY-RUN-AUDIT"

	StatusGate504PermissionLedgerInherited  = "CONDITIONAL_SUPPORT_GATE504_PERMISSION_LEDGER_INHERITED"
	StatusSyntheticAdapterExecuted          = "CONDITIONAL_SUPPORT_SYNTHETIC_ELECTROWEAK_MATCHING_ADAPTER_EXECUTED"
	StatusSyntheticInputsExplicitlyFake     = "CONDITIONAL_SUPPORT_SYNTHETIC_INPUTS_EXPLICITLY_FAKE_AND_TAGGED"
	StatusBridgeTreeWZComputed              = "CONDITIONAL_SUPPORT_BRIDGE_TREE_LEVEL_WZ_FORMULAS_COMPUTED_ON_FAKE_INPUTS"
	StatusPhotonZeroSyntheticPreserved      = "CONDITIONAL_SUPPORT_PHOTON_ZERO_MODE_SURVIVES_SYNTHETIC_ADAPTER"
	StatusTreeRhoIdentitySyntheticConfirmed = "CONDITIONAL_SUPPORT_TREE_RHO_IDENTITY_CONFIRMED_SYNTHETICALLY"
	StatusWeakAngleComputedAsBridgeOutput   = "CONDITIONAL_SUPPORT_WEAK_ANGLE_COMPUTED_AS_BRIDGE_OUTPUT_ONLY"
	StatusNoObservedElectroweakDataImported = "CONDITIONAL_SUPPORT_NO_OBSERVED_ELECTROWEAK_DATA_IMPORTED"
	StatusGate506ObservedComparatorRedirect = "CONDITIONAL_SUPPORT_GATE506_OBSERVED_ELECTROWEAK_COMPARATOR_AIRLOCK_REDIRECT_DEFINED"

	StatusFailedSyntheticNotNativePrediction    = "FAILED_ROUTE_SYNTHETIC_OUTPUTS_ARE_NOT_NATIVE_ELECTROWEAK_PREDICTIONS"
	StatusFailedSyntheticNotObservedMasses      = "FAILED_ROUTE_SYNTHETIC_OUTPUTS_ARE_NOT_OBSERVED_WZ_MASSES"
	StatusFailedVEVCouplingsWeakAngleNotDerived = "FAILED_ROUTE_VEV_COUPLINGS_AND_WEAK_ANGLE_STILL_NOT_DERIVED"
	StatusFailedKappaStillBridge                = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_AFTER_SYNTHETIC_ADAPTER"
	StatusFailedYukawaTraceStillSealed          = "FAILED_ROUTE_YUKAWA_TRACE_A_STILL_SEALED_AFTER_SYNTHETIC_ADAPTER"

	StatusFirewallNoObservedDataImported      = "FIREWALL_PRESERVED_NO_OBSERVED_ELECTROWEAK_DATA_IMPORTED"
	StatusFirewallSyntheticNativeWriteBlocked = "FIREWALL_BLOCKED_SYNTHETIC_ELECTROWEAK_OUTPUT_NATIVE_WRITE"
)

type Inheritance struct {
	Executed                        bool
	Gate504AuditDefined             bool
	PermissionLedgerAccepted        bool
	BridgeInputSchemaDefined        bool
	NativeRows                      int
	BridgeRows                      int
	FormulaBridgeOnly               bool
	PermissionAllowsExplicitAdapter bool
	Gate504NumericalAdapterExecuted bool
	Gate504ObservedEWDataImported   bool
	Gate504NativeWriteBlocked       bool
	Verdict                         string
	Reason                          string
}

type SyntheticInput struct {
	V                       float64
	G2                      float64
	GY                      float64
	RenormalizationScaleTag string
	Scheme                  string
	Source                  string
	Synthetic               bool
	Observed                bool
	Native                  bool
	Purpose                 string
}

type AdapterOutput struct {
	Executed             bool
	Sin2ThetaW           float64
	Cos2ThetaW           float64
	MW                   float64
	MZ                   float64
	MGamma               float64
	RhoTree              float64
	NeutralChargedRatio  float64
	UsedTreeLevelFormula bool
	Verdict              string
	Reason               string
}

type AdapterAudit struct {
	Executed                    bool
	SyntheticOnly               bool
	ObservedDataImported        bool
	NativeDataImported          bool
	InputsPositive              bool
	InputsFinite                bool
	RequiresScaleSchemeMetadata bool
	ScaleSchemeMetadataPresent  bool
	ComputedWithExplicitInputs  bool
	WeakAngleBridgeOutputOnly   bool
	WZBridgeOutputOnly          bool
	PhotonZeroPreserved         bool
	RhoTreeIdentityConfirmed    bool
	ObservedMassesClaimed       bool
	NativeWeakAngleDerived      bool
	NativeWZMassesDerived       bool
	NativeGaugeCouplingsDerived bool
	NativeVEVDerived            bool
	NativeKappaPromoted         bool
	NativeYukawaTraceDerived    bool
	Verdict                     string
	Reason                      string
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
	SyntheticOutputWrittenNative   bool
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
	Input       SyntheticInput
	Output      AdapterOutput
	Adapter     AdapterAudit
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
	g504, err := generation2continuummatchingpermissionledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate504 permission ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g504)
	a.Input = buildSyntheticInput()
	a.Output = runAdapter(a.Input)
	a.Adapter = auditAdapter(a.Inheritance, a.Input, a.Output)
	a.Firewall = buildFirewall(a.Adapter)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g504 generation2continuummatchingpermissionledger.Analysis) Inheritance {
	gate504NativeWriteBlocked := g504.Firewall.NativeVEVWritten == false &&
		g504.Firewall.NativeGaugeCouplingWritten == false &&
		g504.Firewall.NativeWeakAngleWritten == false &&
		g504.Firewall.NativeWZMassWritten == false &&
		g504.Firewall.NativeKappaWritten == false
	return Inheritance{
		Executed:                        true,
		Gate504AuditDefined:             true,
		PermissionLedgerAccepted:        g504.Boundary.PermissionLedgerAccepted,
		BridgeInputSchemaDefined:        g504.Schema.Executed && g504.Schema.NativeRows == 0 && g504.Schema.BridgeRows == len(g504.Schema.Rows),
		NativeRows:                      g504.Schema.NativeRows,
		BridgeRows:                      g504.Schema.BridgeRows,
		FormulaBridgeOnly:               g504.Formula.TreeLevelWZFormulasDefined && !g504.Formula.NativeWeakAngleDerived && !g504.Formula.NativeWZMassesDerived,
		PermissionAllowsExplicitAdapter: g504.Boundary.ContinuumAdapterMayComputeWithExplicitInputs,
		Gate504NumericalAdapterExecuted: g504.Boundary.NumericalAdapterExecuted,
		Gate504ObservedEWDataImported:   g504.Firewall.ObservedVEVImported || g504.Firewall.ObservedGaugeCouplingsImported || g504.Firewall.ObservedWeakAngleImported || g504.Firewall.ObservedWMassImported || g504.Firewall.ObservedZMassImported,
		Gate504NativeWriteBlocked:       gate504NativeWriteBlocked,
		Verdict:                         StatusGate504PermissionLedgerInherited,
		Reason:                          "Gate504 permits an explicit bridge adapter but supplies zero native electroweak matching rows and executed no numerical matching itself.",
	}
}

func buildSyntheticInput() SyntheticInput {
	return SyntheticInput{
		V:                       2,
		G2:                      3,
		GY:                      4,
		RenormalizationScaleTag: "synthetic-mu=1-no-physical-units",
		Scheme:                  "tree-level-synthetic-dry-run",
		Source:                  "fake Gate505 3-4-5 adapter fixture, not observed data",
		Synthetic:               true,
		Observed:                false,
		Native:                  false,
		Purpose:                 "exercise the Gate504 bridge formulas with explicit fake inputs and no native promotion",
	}
}

func runAdapter(in SyntheticInput) AdapterOutput {
	denom := in.G2*in.G2 + in.GY*in.GY
	sin2 := in.GY * in.GY / denom
	cos2 := in.G2 * in.G2 / denom
	mw := in.G2 * in.V / 2
	mz := math.Sqrt(denom) * in.V / 2
	mgamma := 0.0
	rho := math.NaN()
	if mz != 0 && cos2 != 0 {
		rho = mw * mw / (mz * mz * cos2)
	}
	ratio := math.NaN()
	if mw != 0 {
		ratio = mz * mz / (mw * mw)
	}
	return AdapterOutput{
		Executed:             true,
		Sin2ThetaW:           sin2,
		Cos2ThetaW:           cos2,
		MW:                   mw,
		MZ:                   mz,
		MGamma:               mgamma,
		RhoTree:              rho,
		NeutralChargedRatio:  ratio,
		UsedTreeLevelFormula: true,
		Verdict:              strings.Join([]string{StatusBridgeTreeWZComputed, StatusPhotonZeroSyntheticPreserved, StatusTreeRhoIdentitySyntheticConfirmed, StatusWeakAngleComputedAsBridgeOutput}, ";"),
		Reason:               "The fake input v=2, g2=3, gY=4 gives a 3-4-5 dry-run: mW=3, mZ=5, sin²θ=16/25, mγ=0, and ρtree=1.  These are synthetic bridge outputs only.",
	}
}

func auditAdapter(i Inheritance, in SyntheticInput, out AdapterOutput) AdapterAudit {
	finite := allFinite(in.V, in.G2, in.GY, out.Sin2ThetaW, out.Cos2ThetaW, out.MW, out.MZ, out.MGamma, out.RhoTree)
	metadata := in.RenormalizationScaleTag != "" && in.Scheme != "" && in.Source != ""
	bridgeOnly := in.Synthetic && !in.Observed && !in.Native && out.Executed && out.UsedTreeLevelFormula
	return AdapterAudit{
		Executed:                    true,
		SyntheticOnly:               in.Synthetic && !in.Observed,
		ObservedDataImported:        in.Observed,
		NativeDataImported:          in.Native,
		InputsPositive:              in.V > 0 && in.G2 > 0 && in.GY > 0,
		InputsFinite:                finite,
		RequiresScaleSchemeMetadata: true,
		ScaleSchemeMetadataPresent:  metadata,
		ComputedWithExplicitInputs:  i.PermissionAllowsExplicitAdapter && out.Executed,
		WeakAngleBridgeOutputOnly:   bridgeOnly,
		WZBridgeOutputOnly:          bridgeOnly,
		PhotonZeroPreserved:         out.MGamma == 0,
		RhoTreeIdentityConfirmed:    nearly(out.RhoTree, 1, 1e-12),
		ObservedMassesClaimed:       false,
		NativeWeakAngleDerived:      false,
		NativeWZMassesDerived:       false,
		NativeGaugeCouplingsDerived: false,
		NativeVEVDerived:            false,
		NativeKappaPromoted:         false,
		NativeYukawaTraceDerived:    false,
		Verdict: strings.Join([]string{
			StatusSyntheticAdapterExecuted,
			StatusSyntheticInputsExplicitlyFake,
			StatusBridgeTreeWZComputed,
			StatusPhotonZeroSyntheticPreserved,
			StatusTreeRhoIdentitySyntheticConfirmed,
			StatusWeakAngleComputedAsBridgeOutput,
			StatusNoObservedElectroweakDataImported,
			StatusFailedSyntheticNotNativePrediction,
			StatusFailedSyntheticNotObservedMasses,
			StatusFailedVEVCouplingsWeakAngleNotDerived,
			StatusFailedKappaStillBridge,
			StatusFailedYukawaTraceStillSealed,
		}, ";"),
		Reason: "The adapter computes only a synthetic bridge dry-run from explicit fake inputs.  The computed weak angle and W/Z values are not native finite-geometry predictions and are not observed electroweak data.",
	}
}

func buildFirewall(a AdapterAudit) Firewall {
	return Firewall{
		Executed:                     true,
		SyntheticOutputWrittenNative: false,
		Verdict:                      strings.Join([]string{StatusFirewallNoObservedDataImported, StatusFirewallSyntheticNativeWriteBlocked}, ";"),
		Reason:                       "No observed electroweak input is imported and no synthetic output is written to the native theorem registry.",
	}
}

func buildRegistry(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native electroweak scale, coupling, weak angle, kappa, VEV, or W/Z mass theorem is admitted at Gate505.",
		},
		BridgeEntries: []string{
			"A synthetic electroweak bridge adapter dry-run is admitted: fake inputs v=2, g2=3, gY=4 propagate through tree-level bridge formulas to mW=3, mZ=5, sin²θ=16/25, mγ=0, ρtree=1.",
			"The dry-run verifies the Gate504 permission ledger and metadata airlock for explicit bridge inputs.",
		},
		EnvironmentalEntries: []string{
			"Physical VEV, physical gauge couplings, physical weak angle, W/Z pole or running masses, kappa_U1, and Yukawa trace values remain bridge/environmental data, not native ASHA outputs.",
		},
		FailedRoutes: []string{
			StatusFailedSyntheticNotNativePrediction,
			StatusFailedSyntheticNotObservedMasses,
			StatusFailedVEVCouplingsWeakAngleNotDerived,
			StatusFailedKappaStillBridge,
			StatusFailedYukawaTraceStillSealed,
		},
		OpenTheorems: []string{
			"Observed electroweak comparator preflight may be implemented next, but only as an explicit bridge airlock with scale/scheme metadata and no native promotion.",
			"A separate finite-action theorem would still be required to select a nonzero Higgs ray, kappa_U1, gauge couplings, or physical mass matrix natively.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 506, Title: "Observed Electroweak Comparator Airlock Preflight", Reason: "Gate505 proves the adapter can run on fake inputs without native promotion; the next safe step is to define the observed-data preflight schema without executing observed matching by default.", PrimaryTask: "construct an observed electroweak comparator preflight that accepts VEV/coupling/mass records only with explicit bridge tags, scale/scheme metadata, and native-write rejection guards"}
}

func truth(a Analysis) string {
	if a.Adapter.SyntheticOnly && a.Adapter.PhotonZeroPreserved && a.Adapter.RhoTreeIdentityConfirmed && !a.Adapter.NativeWZMassesDerived {
		return "Gate505 proves the electroweak bridge adapter can execute safely when all inputs are explicit, fake, scale-tagged, and bridge-only: the photon zero mode and tree-level rho identity survive the dry-run, but every numerical output is synthetic adapter arithmetic, not a native ASHA electroweak prediction and not observed data."
	}
	return "Gate505 did not validate the synthetic electroweak matching adapter dry-run."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.PermissionLedgerAccepted && a.Inheritance.BridgeInputSchemaDefined && a.Inheritance.NativeRows == 0 && a.Inheritance.BridgeRows == 6 && a.Inheritance.FormulaBridgeOnly && a.Inheritance.PermissionAllowsExplicitAdapter && !a.Inheritance.Gate504NumericalAdapterExecuted && !a.Inheritance.Gate504ObservedEWDataImported && a.Inheritance.Gate504NativeWriteBlocked, "Gate504 inheritance invalid or over-promoted"},
		{a.Input.Synthetic && !a.Input.Observed && !a.Input.Native && a.Input.V == 2 && a.Input.G2 == 3 && a.Input.GY == 4 && a.Input.RenormalizationScaleTag != "" && a.Input.Scheme != "", "synthetic input fixture invalid"},
		{a.Output.Executed && a.Output.UsedTreeLevelFormula && nearly(a.Output.MW, 3, 1e-12) && nearly(a.Output.MZ, 5, 1e-12) && nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) && nearly(a.Output.Cos2ThetaW, 9.0/25.0, 1e-12) && a.Output.MGamma == 0 && nearly(a.Output.RhoTree, 1, 1e-12), "adapter output invalid"},
		{a.Adapter.Executed && a.Adapter.SyntheticOnly && !a.Adapter.ObservedDataImported && !a.Adapter.NativeDataImported && a.Adapter.InputsPositive && a.Adapter.InputsFinite && a.Adapter.ScaleSchemeMetadataPresent && a.Adapter.ComputedWithExplicitInputs && a.Adapter.WeakAngleBridgeOutputOnly && a.Adapter.WZBridgeOutputOnly && a.Adapter.PhotonZeroPreserved && a.Adapter.RhoTreeIdentityConfirmed && !a.Adapter.ObservedMassesClaimed && !a.Adapter.NativeWeakAngleDerived && !a.Adapter.NativeWZMassesDerived && !a.Adapter.NativeGaugeCouplingsDerived && !a.Adapter.NativeVEVDerived && !a.Adapter.NativeKappaPromoted && !a.Adapter.NativeYukawaTraceDerived, "adapter audit invalid or over-promoted"},
		{a.Firewall.Executed && !a.Firewall.ObservedVEVImported && !a.Firewall.ObservedGaugeCouplingsImported && !a.Firewall.ObservedWeakAngleImported && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedZMassImported && !a.Firewall.ObservedYukawaImported && !a.Firewall.NativeVEVWritten && !a.Firewall.NativeGaugeCouplingWritten && !a.Firewall.NativeWeakAngleWritten && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeKappaWritten && !a.Firewall.SyntheticOutputWrittenNative, "firewall violation"},
		{a.Next.Gate == 506, "Gate506 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func allFinite(xs ...float64) bool {
	for _, x := range xs {
		if math.IsNaN(x) || math.IsInf(x, 0) {
			return false
		}
	}
	return true
}

func nearly(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate504=%t permission=%t schema=%t native_rows=%d bridge_rows=%d formula_bridge=%t may_compute=%t gate504_computed=%t observed=%t native_blocked=%t verdict=%s reason=%s", x.Gate504AuditDefined, x.PermissionLedgerAccepted, x.BridgeInputSchemaDefined, x.NativeRows, x.BridgeRows, x.FormulaBridgeOnly, x.PermissionAllowsExplicitAdapter, x.Gate504NumericalAdapterExecuted, x.Gate504ObservedEWDataImported, x.Gate504NativeWriteBlocked, x.Verdict, x.Reason)
}
func FormatInput(x SyntheticInput) string {
	return fmt.Sprintf("v=%s g2=%s gY=%s scale=%s scheme=%s source=%s synthetic=%t observed=%t native=%t purpose=%s", fmtFloat(x.V), fmtFloat(x.G2), fmtFloat(x.GY), x.RenormalizationScaleTag, x.Scheme, x.Source, x.Synthetic, x.Observed, x.Native, x.Purpose)
}
func FormatOutput(x AdapterOutput) string {
	return fmt.Sprintf("executed=%t sin2=%s cos2=%s mW=%s mZ=%s mGamma=%s rho=%s ratio=%s tree=%t verdict=%s reason=%s", x.Executed, fmtFloat(x.Sin2ThetaW), fmtFloat(x.Cos2ThetaW), fmtFloat(x.MW), fmtFloat(x.MZ), fmtFloat(x.MGamma), fmtFloat(x.RhoTree), fmtFloat(x.NeutralChargedRatio), x.UsedTreeLevelFormula, x.Verdict, x.Reason)
}
func FormatAdapter(x AdapterAudit) string {
	return fmt.Sprintf("executed=%t synthetic_only=%t observed=%t native_input=%t finite=%t metadata=%t explicit=%t weak_bridge=%t wz_bridge=%t photon0=%t rho1=%t observed_claim=%t native_theta=%t native_WZ=%t native_gauge=%t native_v=%t native_kappa=%t native_a=%t verdict=%s reason=%s", x.Executed, x.SyntheticOnly, x.ObservedDataImported, x.NativeDataImported, x.InputsFinite, x.ScaleSchemeMetadataPresent, x.ComputedWithExplicitInputs, x.WeakAngleBridgeOutputOnly, x.WZBridgeOutputOnly, x.PhotonZeroPreserved, x.RhoTreeIdentityConfirmed, x.ObservedMassesClaimed, x.NativeWeakAngleDerived, x.NativeWZMassesDerived, x.NativeGaugeCouplingsDerived, x.NativeVEVDerived, x.NativeKappaPromoted, x.NativeYukawaTraceDerived, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("obs_v=%t obs_gauge=%t obs_theta=%t obs_W=%t obs_Z=%t obs_yukawa=%t native_v=%t native_gauge=%t native_theta=%t native_WZ=%t native_kappa=%t synthetic_native=%t verdict=%s reason=%s", x.ObservedVEVImported, x.ObservedGaugeCouplingsImported, x.ObservedWeakAngleImported, x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedYukawaImported, x.NativeVEVWritten, x.NativeGaugeCouplingWritten, x.NativeWeakAngleWritten, x.NativeWZMassWritten, x.NativeKappaWritten, x.SyntheticOutputWrittenNative, x.Verdict, x.Reason)
}
func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 505 Registry Audit — Synthetic Electroweak Matching Adapter Dry-Run\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range []string{StatusGate504PermissionLedgerInherited, StatusSyntheticAdapterExecuted, StatusSyntheticInputsExplicitlyFake, StatusBridgeTreeWZComputed, StatusPhotonZeroSyntheticPreserved, StatusTreeRhoIdentitySyntheticConfirmed, StatusWeakAngleComputedAsBridgeOutput, StatusNoObservedElectroweakDataImported, StatusFailedSyntheticNotNativePrediction, StatusFailedSyntheticNotObservedMasses, StatusFailedVEVCouplingsWeakAngleNotDerived, StatusFailedKappaStillBridge, StatusFailedYukawaTraceStillSealed, StatusFirewallNoObservedDataImported, StatusFirewallSyntheticNativeWriteBlocked, StatusGate506ObservedComparatorRedirect} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate504 opened an electroweak continuum-matching airlock: explicit `v`, `g2`, `gY`, weak-angle outputs, W/Z comparator outputs, and Yukawa trace normalization may enter only as bridge/environmental data with scale and scheme metadata.  Gate504 admitted zero native electroweak matching rows and executed no numerical adapter.\n\n")
	b.WriteString("## Synthetic adapter input\n\n")
	b.WriteString(fmt.Sprintf("```text\n%s\n```\n\n", FormatInput(a.Input)))
	b.WriteString("The fixture is deliberately fake.  The `3-4-5` pattern is chosen only to make the tree-level bridge arithmetic transparent; it is not a physical electroweak dataset.\n\n")
	b.WriteString("## Bridge computation\n\n")
	b.WriteString(fmt.Sprintf("```text\nm_W = g2 v / 2 = %s\nm_Z = sqrt(g2^2 + gY^2) v / 2 = %s\nsin^2(theta_W) = gY^2/(g2^2+gY^2) = %s\ncos^2(theta_W) = g2^2/(g2^2+gY^2) = %s\nm_gamma = %s\nrho_tree = m_W^2/(m_Z^2 cos^2 theta_W) = %s\nneutral/charged quotient ratio = m_Z^2/m_W^2 = %s\n```\n\n", fmtFloat(a.Output.MW), fmtFloat(a.Output.MZ), fmtFloat(a.Output.Sin2ThetaW), fmtFloat(a.Output.Cos2ThetaW), fmtFloat(a.Output.MGamma), fmtFloat(a.Output.RhoTree), fmtFloat(a.Output.NeutralChargedRatio)))
	b.WriteString("This computation confirms only that the bridge adapter is algebraically coherent and photon-safe when explicit fake inputs are supplied.\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("No observed VEV, gauge coupling, weak angle, W mass, Z mass, or Yukawa value is imported.  The synthetic outputs are not written to the native registry.  `kappa_U1 = 6`, the Higgs VEV, physical gauge couplings, and physical W/Z masses remain blocked as native claims.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\n")
	b.WriteString("Gate506 should be:\n\n```text\nGate 506 — Observed Electroweak Comparator Airlock Preflight\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "NaN"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.12f", x), "0"), ".")
}
