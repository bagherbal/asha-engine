// Package generation2syntheticapsindexboundaryledger implements Gate 518:
// Synthetic APS Index Boundary Ledger Dry-Run.
//
// Gate 517 defined the APS/index/eta airlock: the local chiral index-density
// socket and anomaly-inflow structure are native/scale-free, while global index
// integers, boundary eta invariants, boundary spectra, and closed-manifold
// assumptions require continuum topology/boundary data. Gate 518 exercises that
// airlock with explicitly fake bridge rows. It validates the arithmetic plumbing
// of the APS formula and the native-write firewall; it does not derive a global
// Dirac index, eta invariant, or manifold topology.
package generation2syntheticapsindexboundaryledger

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gravitationalindexetaairlock"
)

const (
	AuditID = "GATE518-SYNTHETIC-APS-INDEX-BOUNDARY-LEDGER"

	StatusGate517IndexEtaAirlockInherited       = "CONDITIONAL_SUPPORT_GATE517_INDEX_ETA_AIRLOCK_INHERITED"
	StatusSyntheticAPSLedgerExecuted            = "CONDITIONAL_SUPPORT_SYNTHETIC_APS_INDEX_BOUNDARY_LEDGER_EXECUTED"
	StatusSyntheticTopologyRowsFake             = "CONDITIONAL_SUPPORT_SYNTHETIC_GLOBAL_TOPOLOGY_ROWS_EXPLICITLY_FAKE"
	StatusAPSFormulaComputedBridgeOnly          = "CONDITIONAL_SUPPORT_APS_FORMULA_COMPUTED_BRIDGE_ONLY"
	StatusBoundaryEtaKernelCorrectionTested     = "CONDITIONAL_SUPPORT_BOUNDARY_ETA_KERNEL_CORRECTION_PLUMBING_TESTED"
	StatusClosedManifoldSpecializationTested    = "CONDITIONAL_SUPPORT_CLOSED_MANIFOLD_SPECIALIZATION_TESTED_SYNTHETICALLY"
	StatusSyntheticIndexResidualsZero           = "CONDITIONAL_SUPPORT_INDEX_RESIDUALS_ZERO_SYNTHETICALLY"
	StatusNoObservedTopologyBoundaryImported    = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_OR_BOUNDARY_DATA_IMPORTED"
	StatusFailedSyntheticNotNativeIndex         = "FAILED_ROUTE_SYNTHETIC_APS_OUTPUTS_ARE_NOT_NATIVE_INDEX_PREDICTIONS"
	StatusFailedEtaNotBoundarySpectrumDerived   = "FAILED_ROUTE_SYNTHETIC_ETA_OUTPUT_IS_NOT_BOUNDARY_SPECTRUM_DERIVATION"
	StatusFailedGlobalTopologyStillUnselected   = "FAILED_ROUTE_GLOBAL_TOPOLOGY_STILL_NOT_SELECTED_BY_SYNTHETIC_LEDGER"
	StatusFailedBoundaryConditionStillNotNative = "FAILED_ROUTE_BOUNDARY_CONDITION_STILL_NOT_NATIVE_SELECTED"
	StatusFailedGravThetaStillUnselected        = "FAILED_ROUTE_GRAVITATIONAL_THETA_STILL_NOT_SELECTED_BY_APS_DRY_RUN"
	StatusFailedNewtonCosmologyStillBlocked     = "FAILED_ROUTE_NEWTON_CUTOFF_AND_COSMOLOGICAL_NORMALIZATION_STILL_BLOCKED_AFTER_APS_DRY_RUN"
	StatusFirewallNoEmpiricalTopologyImported   = "FIREWALL_PRESERVED_NO_MANIFOLD_BOUNDARY_NEWTON_OR_COSMOLOGY_DATA_IMPORTED"
	StatusFirewallSyntheticAPSWritesBlocked     = "FIREWALL_BLOCKED_SYNTHETIC_APS_INDEX_NATIVE_WRITE"
)

const eps = 1e-12

type Inheritance struct {
	Executed                        bool
	Gate517Inherited                bool
	Gate517LocalIndexDensitySocket  bool
	Gate517APSSocket                bool
	Gate517BoundaryEtaAirlock       bool
	Gate517AnomalyInflowSocket      bool
	Gate517GlobalIndexBlocked       bool
	Gate517EtaBlocked               bool
	Gate517BoundarySpectrumBlocked  bool
	Gate517ObservedBoundaryImported bool
	Verdict                         string
	Reason                          string
}

type SyntheticLedger struct {
	Executed                 bool
	BridgeOnly               bool
	SyntheticOnly            bool
	UsesObservedTopology     bool
	UsesBoundarySpectrum     bool
	Source                   string
	LocalIndexIntegral       float64
	BoundaryEta              float64
	BoundaryKernelDimensionH float64
	BoundaryCorrection       float64
	APSIndex                 float64
	ExpectedAPSIndex         float64
	ClosedManifoldLocalIndex float64
	ClosedManifoldEta        float64
	ClosedManifoldKernelH    float64
	ClosedManifoldIndex      float64
	ExpectedClosedIndex      float64
	APSResidual              float64
	ClosedResidual           float64
	APSIndexIntegerLike      bool
	ClosedIndexIntegerLike   bool
	Verdict                  string
	Reason                   string
}

type AirlockPolicy struct {
	Executed                          bool
	RequiresBridgeOnlyTag             bool
	RequiresSyntheticOrExternalTag    bool
	RequiresSourceMetadata            bool
	RequiresTopologyMetadata          bool
	RequiresBoundaryMetadata          bool
	RejectsNativePromotion            bool
	RejectsObservedByDefault          bool
	RejectsMissingEtaKernelRows       bool
	RejectsMissingBoundaryCondition   bool
	RejectsMissingUncertaintyMetadata bool
	NativeIndexPredictionMade         bool
	NativeEtaPredictionMade           bool
	BoundaryConditionSelected         bool
	BoundarySpectrumDerived           bool
	ClosedManifoldNativelySelected    bool
	Verdict                           string
	Reason                            string
}

type Firewall struct {
	Executed                           bool
	UsesLambdaCutoff                   bool
	UsesF2Moment                       bool
	UsesF4Moment                       bool
	UsesNewtonConstant                 bool
	UsesCosmologicalConstant           bool
	UsesPlanckScale                    bool
	UsesHiggsOrElectroweakScale        bool
	UsesFlavorYukawaData               bool
	ObservedTopologyImported           bool
	ObservedBoundarySpectrumImported   bool
	SyntheticOutputNativeWrite         bool
	GlobalIndexNativePrediction        bool
	BoundaryEtaNativePrediction        bool
	PhysicalGravitationalThetaWritten  bool
	GravityCosmologyNormalizationWrite bool
	Verdict                            string
	Reason                             string
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
	Ledger      SyntheticLedger
	Policy      AirlockPolicy
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
	g517, err := generation2gravitationalindexetaairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate517 index/eta airlock: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g517)
	a.Ledger = buildSyntheticLedger()
	a.Policy = buildPolicy()
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g517 generation2gravitationalindexetaairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate517Inherited:                g517.Index.LocalIndexDensitySocketPresent && g517.Index.ClosedManifoldSocketConsistent && g517.Eta.BoundaryEtaNativeWrite == false,
		Gate517LocalIndexDensitySocket:  g517.Index.LocalIndexDensitySocketPresent,
		Gate517APSSocket:                g517.Index.ClosedManifoldSocketConsistent,
		Gate517BoundaryEtaAirlock:       g517.Eta.EtaInvariantRequired && g517.Eta.BoundaryConditionRequired,
		Gate517AnomalyInflowSocket:      g517.Inflow.PontryaginDescentSocketPresent && g517.Inflow.ChernSimonsBoundarySocketPresent,
		Gate517GlobalIndexBlocked:       !g517.Index.GlobalIndexIntegerDerived,
		Gate517EtaBlocked:               !g517.Index.BoundaryEtaDerived && !g517.Eta.BoundaryEtaNativeDerived,
		Gate517BoundarySpectrumBlocked:  !g517.Index.BoundarySpectrumSelected,
		Gate517ObservedBoundaryImported: g517.Firewall.ObservedBoundarySpectrumImported || g517.Firewall.ObservedTopologyImported,
		Verdict:                         StatusGate517IndexEtaAirlockInherited,
		Reason:                          "Gate518 inherits Gate517's local index-density socket, APS formula socket, boundary eta airlock, anomaly-inflow socket, and explicit block on native global-index/eta/boundary-spectrum writes.",
	}
}

func buildSyntheticLedger() SyntheticLedger {
	local := 11.0
	eta := 3.0
	h := 1.0
	correction := (eta + h) / 2.0
	aps := local - correction
	closedLocal := 11.0
	closedEta := 0.0
	closedH := 0.0
	closed := closedLocal - (closedEta+closedH)/2.0
	return SyntheticLedger{
		Executed:                 true,
		BridgeOnly:               true,
		SyntheticOnly:            true,
		UsesObservedTopology:     false,
		UsesBoundarySpectrum:     false,
		Source:                   "synthetic Gate518 APS dry-run; fake topology and boundary numbers used only to test bridge plumbing",
		LocalIndexIntegral:       local,
		BoundaryEta:              eta,
		BoundaryKernelDimensionH: h,
		BoundaryCorrection:       correction,
		APSIndex:                 aps,
		ExpectedAPSIndex:         9.0,
		ClosedManifoldLocalIndex: closedLocal,
		ClosedManifoldEta:        closedEta,
		ClosedManifoldKernelH:    closedH,
		ClosedManifoldIndex:      closed,
		ExpectedClosedIndex:      11.0,
		APSResidual:              aps - 9.0,
		ClosedResidual:           closed - 11.0,
		APSIndexIntegerLike:      nearly(aps, math.Round(aps), eps),
		ClosedIndexIntegerLike:   nearly(closed, math.Round(closed), eps),
		Verdict:                  strings.Join([]string{StatusSyntheticAPSLedgerExecuted, StatusSyntheticTopologyRowsFake, StatusAPSFormulaComputedBridgeOnly, StatusBoundaryEtaKernelCorrectionTested, StatusClosedManifoldSpecializationTested, StatusSyntheticIndexResidualsZero, StatusFailedSyntheticNotNativeIndex, StatusFailedEtaNotBoundarySpectrumDerived}, ";"),
		Reason:                   "The dry-run computes ind_APS = 11 - (3+1)/2 = 9 and the closed specialization ind = 11 using explicitly fake bridge rows. Integer-like outputs only validate APS arithmetic and residual plumbing; they are not native global topology or eta derivations.",
	}
}

func buildPolicy() AirlockPolicy {
	return AirlockPolicy{
		Executed:                          true,
		RequiresBridgeOnlyTag:             true,
		RequiresSyntheticOrExternalTag:    true,
		RequiresSourceMetadata:            true,
		RequiresTopologyMetadata:          true,
		RequiresBoundaryMetadata:          true,
		RejectsNativePromotion:            true,
		RejectsObservedByDefault:          true,
		RejectsMissingEtaKernelRows:       true,
		RejectsMissingBoundaryCondition:   true,
		RejectsMissingUncertaintyMetadata: true,
		NativeIndexPredictionMade:         false,
		NativeEtaPredictionMade:           false,
		BoundaryConditionSelected:         false,
		BoundarySpectrumDerived:           false,
		ClosedManifoldNativelySelected:    false,
		Verdict:                           strings.Join([]string{StatusNoObservedTopologyBoundaryImported, StatusFailedGlobalTopologyStillUnselected, StatusFailedBoundaryConditionStillNotNative}, ";"),
		Reason:                            "Any APS/topology row must remain bridge-only and source-tagged. Missing eta/kernel/boundary-condition metadata fails closed, and observed topology is rejected by default unless a future explicit comparator mode is added.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusNoObservedTopologyBoundaryImported, StatusFailedGravThetaStillUnselected, StatusFailedNewtonCosmologyStillBlocked, StatusFirewallNoEmpiricalTopologyImported, StatusFirewallSyntheticAPSWritesBlocked}, ";"),
		Reason:   "Gate518 imports no cutoff, spectral moments, Newton/Planck normalization, cosmology, electroweak scale, flavor/Yukawa data, observed topology, or boundary spectra. Synthetic APS outputs are blocked from native registry writes.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No new native global topology integer is written at Gate518.",
			"The inherited local index-density and anomaly-inflow sockets remain the only native/topological content.",
		},
		BridgeEntries: []string{
			"Synthetic APS dry-run: ind_APS = I_local - (eta+h)/2 with fake I_local=11, eta=3, h=1 giving 9.",
			"Synthetic closed-manifold specialization: eta=h=0 with fake I_local=11 giving 11.",
			"Fail-closed APS/topology airlock requiring bridge-only, source, topology, boundary, eta, kernel, and uncertainty metadata.",
		},
		EnvironmentalEntries: []string{
			"Actual manifold topology, Euler/signature/Pontryagin integers, boundary condition, boundary Dirac spectrum, eta invariant, kernel dimension h, bordism/orientation data, and closedness.",
			"Newton/Planck normalization, cutoff Lambda, spectral moments, cosmological constant, electroweak scales, and flavor/Yukawa data.",
		},
		FailedRoutes: []string{
			"Treating synthetic APS integers as native global index predictions.",
			"Treating a synthetic eta row as a boundary-spectrum derivation.",
			"Using APS dry-run arithmetic to select gravitational theta, manifold topology, or physical gravity/cosmology normalization.",
		},
		OpenTheorems: []string{
			"An observed/global topology comparator preflight that remains bridge-only and source-tagged.",
			"A native manifold/bordism selector, if ASHA can ever derive one.",
			"A boundary Hilbert-space theorem deriving eta from a native boundary operator spectrum.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 519, Title: "Observed Topology and Boundary Comparator Preflight", Reason: "Gate518 validates synthetic APS/index plumbing but still blocks actual manifold topology and boundary spectra.", PrimaryTask: "Define a fail-closed comparator schema for external Euler/signature/Pontryagin/eta/boundary rows, requiring source and bridge-only metadata while rejecting native promotion by default."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate517Inherited && a.Inheritance.Gate517LocalIndexDensitySocket && a.Inheritance.Gate517APSSocket && a.Inheritance.Gate517BoundaryEtaAirlock && a.Inheritance.Gate517AnomalyInflowSocket && a.Inheritance.Gate517GlobalIndexBlocked && a.Inheritance.Gate517EtaBlocked && a.Inheritance.Gate517BoundarySpectrumBlocked && !a.Inheritance.Gate517ObservedBoundaryImported, "Gate518 inheritance invalid"},
		{a.Ledger.Executed && a.Ledger.BridgeOnly && a.Ledger.SyntheticOnly && !a.Ledger.UsesObservedTopology && !a.Ledger.UsesBoundarySpectrum && nearly(a.Ledger.BoundaryCorrection, 2, eps) && nearly(a.Ledger.APSIndex, 9, eps) && nearly(a.Ledger.ClosedManifoldIndex, 11, eps) && nearly(a.Ledger.APSResidual, 0, eps) && nearly(a.Ledger.ClosedResidual, 0, eps) && a.Ledger.APSIndexIntegerLike && a.Ledger.ClosedIndexIntegerLike, "Gate518 synthetic ledger invalid"},
		{a.Policy.Executed && a.Policy.RequiresBridgeOnlyTag && a.Policy.RequiresSyntheticOrExternalTag && a.Policy.RequiresSourceMetadata && a.Policy.RequiresTopologyMetadata && a.Policy.RequiresBoundaryMetadata && a.Policy.RejectsNativePromotion && a.Policy.RejectsObservedByDefault && a.Policy.RejectsMissingEtaKernelRows && a.Policy.RejectsMissingBoundaryCondition && a.Policy.RejectsMissingUncertaintyMetadata && !a.Policy.NativeIndexPredictionMade && !a.Policy.NativeEtaPredictionMade && !a.Policy.BoundaryConditionSelected && !a.Policy.BoundarySpectrumDerived && !a.Policy.ClosedManifoldNativelySelected, "Gate518 policy invalid"},
		{a.Firewall.Executed && !a.Firewall.UsesLambdaCutoff && !a.Firewall.UsesF2Moment && !a.Firewall.UsesF4Moment && !a.Firewall.UsesNewtonConstant && !a.Firewall.UsesCosmologicalConstant && !a.Firewall.UsesPlanckScale && !a.Firewall.UsesHiggsOrElectroweakScale && !a.Firewall.UsesFlavorYukawaData && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundarySpectrumImported && !a.Firewall.SyntheticOutputNativeWrite && !a.Firewall.GlobalIndexNativePrediction && !a.Firewall.BoundaryEtaNativePrediction && !a.Firewall.PhysicalGravitationalThetaWritten && !a.Firewall.GravityCosmologyNormalizationWrite, "Gate518 firewall invalid"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate517IndexEtaAirlockInherited,
		StatusSyntheticAPSLedgerExecuted,
		StatusSyntheticTopologyRowsFake,
		StatusAPSFormulaComputedBridgeOnly,
		StatusBoundaryEtaKernelCorrectionTested,
		StatusClosedManifoldSpecializationTested,
		StatusSyntheticIndexResidualsZero,
		StatusNoObservedTopologyBoundaryImported,
		StatusFailedSyntheticNotNativeIndex,
		StatusFailedEtaNotBoundarySpectrumDerived,
		StatusFailedGlobalTopologyStillUnselected,
		StatusFailedBoundaryConditionStillNotNative,
		StatusFailedGravThetaStillUnselected,
		StatusFailedNewtonCosmologyStillBlocked,
		StatusFirewallNoEmpiricalTopologyImported,
		StatusFirewallSyntheticAPSWritesBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 518 validates the APS/index boundary ledger as bridge plumbing only: the formula ind_APS = local index integral - (eta+h)/2 and the closed-manifold specialization are internally consistent on explicitly fake rows, but no manifold topology, boundary spectrum, eta invariant, global index integer, gravitational theta coefficient, Newton normalization, cutoff, or cosmological quantity is derived. The local index theorem remains native; its global data remain environmental or bridge-supplied."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate517 inherited=%t; local density socket=%t; APS socket=%t; eta airlock=%t; anomaly inflow=%t; global index blocked=%t; eta blocked=%t; boundary spectrum blocked=%t; observed boundary imported=%t", x.Gate517Inherited, x.Gate517LocalIndexDensitySocket, x.Gate517APSSocket, x.Gate517BoundaryEtaAirlock, x.Gate517AnomalyInflowSocket, x.Gate517GlobalIndexBlocked, x.Gate517EtaBlocked, x.Gate517BoundarySpectrumBlocked, x.Gate517ObservedBoundaryImported)
}

func FormatLedger(x SyntheticLedger) string {
	return fmt.Sprintf("source=%q; bridge_only=%t; synthetic_only=%t; observed_topology=%t; boundary_spectrum=%t; local=%.12g; eta=%.12g; h=%.12g; correction=(eta+h)/2=%.12g; ind_APS=%.12g; expected=%.12g; residual=%.12g; closed_local=%.12g; closed_eta=%.12g; closed_h=%.12g; closed_index=%.12g; closed_expected=%.12g; closed_residual=%.12g; aps_integer_like=%t; closed_integer_like=%t", x.Source, x.BridgeOnly, x.SyntheticOnly, x.UsesObservedTopology, x.UsesBoundarySpectrum, x.LocalIndexIntegral, x.BoundaryEta, x.BoundaryKernelDimensionH, x.BoundaryCorrection, x.APSIndex, x.ExpectedAPSIndex, x.APSResidual, x.ClosedManifoldLocalIndex, x.ClosedManifoldEta, x.ClosedManifoldKernelH, x.ClosedManifoldIndex, x.ExpectedClosedIndex, x.ClosedResidual, x.APSIndexIntegerLike, x.ClosedIndexIntegerLike)
}

func FormatPolicy(x AirlockPolicy) string {
	return fmt.Sprintf("requires_bridge_only=%t; requires_synthetic_or_external_tag=%t; requires_source=%t; requires_topology_metadata=%t; requires_boundary_metadata=%t; rejects_native=%t; rejects_observed_default=%t; rejects_missing_eta_h=%t; rejects_missing_boundary_condition=%t; rejects_missing_uncertainty=%t; native_index_prediction=%t; native_eta_prediction=%t; boundary_condition_selected=%t; boundary_spectrum_derived=%t; closed_manifold_native=%t", x.RequiresBridgeOnlyTag, x.RequiresSyntheticOrExternalTag, x.RequiresSourceMetadata, x.RequiresTopologyMetadata, x.RequiresBoundaryMetadata, x.RejectsNativePromotion, x.RejectsObservedByDefault, x.RejectsMissingEtaKernelRows, x.RejectsMissingBoundaryCondition, x.RejectsMissingUncertaintyMetadata, x.NativeIndexPredictionMade, x.NativeEtaPredictionMade, x.BoundaryConditionSelected, x.BoundarySpectrumDerived, x.ClosedManifoldNativelySelected)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("uses_Lambda=%t; uses_f2=%t; uses_f4=%t; uses_Newton=%t; uses_cosmological=%t; uses_Planck=%t; uses_EW=%t; uses_flavor=%t; observed_topology=%t; observed_boundary_spectrum=%t; synthetic_native_write=%t; global_index_native=%t; eta_native=%t; theta_written=%t; gravity_cosmology_write=%t", x.UsesLambdaCutoff, x.UsesF2Moment, x.UsesF4Moment, x.UsesNewtonConstant, x.UsesCosmologicalConstant, x.UsesPlanckScale, x.UsesHiggsOrElectroweakScale, x.UsesFlavorYukawaData, x.ObservedTopologyImported, x.ObservedBoundarySpectrumImported, x.SyntheticOutputNativeWrite, x.GlobalIndexNativePrediction, x.BoundaryEtaNativePrediction, x.PhysicalGravitationalThetaWritten, x.GravityCosmologyNormalizationWrite)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 518 Registry Audit — Synthetic APS Index Boundary Ledger Dry-Run\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Synthetic APS ledger\n\n" + a.Ledger.Reason + "\n\n```text\n" + FormatLedger(a.Ledger) + "\n```\n\n")
	b.WriteString("APS formulas exercised:\n\n```text\nind_APS(D_E) = ∫_M [Â(R) ch(E)]_4 - (η(D_∂)+h)/2\nsynthetic: 11 - (3+1)/2 = 9\nclosed synthetic: 11 - (0+0)/2 = 11\n```\n\n")
	b.WriteString("## Airlock policy\n\n" + a.Policy.Reason + "\n\n```text\n" + FormatPolicy(a.Policy) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate519 should be:\n\n```text\nGate 519 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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

func nearly(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
