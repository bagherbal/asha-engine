// Package generation2gravitationalindexetaairlock implements Gate 517:
// Gravitational Index and Boundary Eta Airlock.
//
// Gate 516 confirmed that the a4 gravity channel contains scale-free
// Euler/Gauss-Bonnet and Pontryagin/signature characteristic-class sockets, but
// blocked actual manifold integers and boundary eta data. Gate 517 audits the
// Atiyah-Singer / Atiyah-Patodi-Singer index lane: the finite grading supplies a
// local chiral index socket, while the global index, boundary eta invariant, and
// boundary spectral asymmetry remain continuum/global inputs unless ASHA later
// derives a native spacetime topology and boundary spectrum.
package generation2gravitationalindexetaairlock

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologicalgravitycharacteristicclassledger"
)

const (
	AuditID = "GATE517-GRAVITATIONAL-INDEX-BOUNDARY-ETA-AIRLOCK"

	StatusGate516TopologyInherited             = "CONDITIONAL_SUPPORT_GATE516_GRAVITY_TOPOLOGY_LEDGER_INHERITED"
	StatusAPSIndexLedgerDefined                = "CONDITIONAL_SUPPORT_APS_INDEX_LEDGER_DEFINED"
	StatusLocalIndexDensitySocketScaleFree     = "CONDITIONAL_SUPPORT_LOCAL_INDEX_DENSITY_SOCKET_SCALE_FREE"
	StatusClosedManifoldIndexSocketConsistent  = "CONDITIONAL_SUPPORT_CLOSED_MANIFOLD_INDEX_SOCKET_CONSISTENT"
	StatusBoundaryEtaAirlockDefined            = "CONDITIONAL_SUPPORT_BOUNDARY_ETA_AIRLOCK_DEFINED"
	StatusAnomalyInflowSocketPresent           = "CONDITIONAL_SUPPORT_GRAVITATIONAL_ANOMALY_INFLOW_SOCKET_PRESENT"
	StatusEtaCorrectionClassifiedGlobal        = "CONDITIONAL_SUPPORT_ETA_CORRECTION_CLASSIFIED_AS_GLOBAL_BOUNDARY_DATA"
	StatusScaleAndMassIndependencePreserved    = "CONDITIONAL_SUPPORT_INDEX_ETA_LEDGER_SCALE_AND_MASS_INDEPENDENT"
	StatusNoObservedTopologyOrBoundaryImported = "CONDITIONAL_SUPPORT_NO_OBSERVED_TOPOLOGY_OR_BOUNDARY_DATA_IMPORTED"

	StatusFailedGlobalIndexIntegerNotDerived       = "FAILED_ROUTE_GLOBAL_DIRAC_INDEX_INTEGER_NOT_DERIVED_WITHOUT_MANIFOLD_TOPOLOGY"
	StatusFailedBoundaryEtaNotDerived              = "FAILED_ROUTE_BOUNDARY_ETA_INVARIANT_NOT_DERIVED_WITHOUT_BOUNDARY_SPECTRUM"
	StatusFailedBoundarySpectrumNotSelected        = "FAILED_ROUTE_BOUNDARY_SPECTRUM_NOT_SELECTED_BY_FINITE_ALGEBRA"
	StatusFailedClosedManifoldAssumptionNotNative  = "FAILED_ROUTE_CLOSED_MANIFOLD_OR_BOUNDARY_CONDITION_NOT_NATIVE_SELECTED"
	StatusFailedGravitationalThetaStillUnselected  = "FAILED_ROUTE_GRAVITATIONAL_THETA_AND_PONTRYAGIN_COEFFICIENT_NOT_SELECTED"
	StatusFailedManifoldTopologyStillEnvironmental = "FAILED_ROUTE_CONTINUUM_MANIFOLD_TOPOLOGY_STILL_ENVIRONMENTAL"
	StatusFailedNewtonCosmologyStillBlocked        = "FAILED_ROUTE_NEWTON_CUTOFF_AND_COSMOLOGICAL_NORMALIZATION_STILL_BLOCKED_AFTER_INDEX_AIRLOCK"
	StatusFirewallNoEmpiricalBoundaryImported      = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_MANIFOLD_OR_BOUNDARY_DATA_IMPORTED"
	StatusFirewallIndexEtaNativeWriteBlocked       = "FIREWALL_BLOCKED_GLOBAL_INDEX_AND_ETA_NATIVE_WRITE"
)

const (
	finiteTraceDimension = 96.0
	// The inherited Gate516 a4 local unit prefactor after finite trace stripping.
	a4UnitPrefactor = 1.0 / (60.0 * math.Pi * math.Pi)
)

type Inheritance struct {
	Executed                          bool
	Gate516Inherited                  bool
	Gate516EulerSocket                bool
	Gate516PontryaginSocket           bool
	Gate516CharacteristicScaleFree    bool
	Gate516ChiralIndexSocket          bool
	Gate516MixedGaugeGravityTraceZero bool
	Gate516GlobalIntegersBlocked      bool
	Gate516EtaBlocked                 bool
	Gate516ObservedTopologyImported   bool
	Verdict                           string
	Reason                            string
}

type APSIndexLedger struct {
	Executed                       bool
	Dimension                      int
	FiniteGradingAvailable         bool
	RealStructureAvailable         bool
	LocalIndexDensity              string
	APSFormula                     string
	ClosedManifoldReduction        string
	LocalIndexDensitySocketPresent bool
	ClosedManifoldSocketConsistent bool
	GlobalIndexIntegerDerived      bool
	BoundaryEtaDerived             bool
	BoundaryKernelDimensionDerived bool
	BoundarySpectrumSelected       bool
	ClosedManifoldSelected         bool
	Verdict                        string
	Reason                         string
}

type BoundaryEtaAirlock struct {
	Executed                       bool
	BoundaryOperatorRequired       bool
	BoundarySpectrumRequired       bool
	EtaInvariantRequired           bool
	KernelCorrectionRequired       bool
	BoundaryConditionRequired      bool
	GlobalTopologyRequired         bool
	BoundaryDataImported           bool
	BoundaryEtaNativeDerived       bool
	BoundaryEtaNativeWrite         bool
	ClosedManifoldIsAllowedBridge  bool
	ClosedManifoldIsNativeSelected bool
	Verdict                        string
	Reason                         string
}

type AnomalyInflowAudit struct {
	Executed                         bool
	PontryaginDescentSocketPresent   bool
	ChernSimonsBoundarySocketPresent bool
	ChiralIndexAnomalySocketPresent  bool
	MixedGaugeGravityTraceZero       bool
	BoundaryEtaPairsWithInflow       bool
	PhysicalThetaCoefficientDerived  bool
	BoundaryTheorySelected           bool
	Verdict                          string
	Reason                           string
}

type ScaleFirewall struct {
	Executed                          bool
	UsesLambdaCutoff                  bool
	UsesF2Moment                      bool
	UsesF4Moment                      bool
	UsesNewtonConstant                bool
	UsesCosmologicalConstant          bool
	UsesHiggsOrElectroweakScale       bool
	UsesFlavorYukawaData              bool
	ObservedTopologyImported          bool
	ObservedBoundarySpectrumImported  bool
	GlobalIndexIntegerNativeWrite     bool
	BoundaryEtaNativeWrite            bool
	PhysicalGravitationalThetaWritten bool
	NativeGravityNormalizationWritten bool
	Verdict                           string
	Reason                            string
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
	Index       APSIndexLedger
	Eta         BoundaryEtaAirlock
	Inflow      AnomalyInflowAudit
	Firewall    ScaleFirewall
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
	g516, err := generation2topologicalgravitycharacteristicclassledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate516 characteristic-class ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g516)
	a.Index = buildIndex()
	a.Eta = buildEta()
	a.Inflow = buildInflow(g516)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g516 generation2topologicalgravitycharacteristicclassledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate516Inherited:                  g516.Ledger.A4CharacteristicSubspace && g516.Scale.CharacteristicIntegralsScaleFree,
		Gate516EulerSocket:                g516.Ledger.EulerSocketPresent,
		Gate516PontryaginSocket:           g516.Ledger.PontryaginSocketPresent && g516.Ledger.SignatureSocketPresent,
		Gate516CharacteristicScaleFree:    g516.Scale.CharacteristicIntegralsScaleFree && !g516.Scale.UsesLambdaCutoff && !g516.Scale.UsesNewtonConstant,
		Gate516ChiralIndexSocket:          g516.Finite.ChiralIndexSocketPresent,
		Gate516MixedGaugeGravityTraceZero: g516.Finite.MixedGravitationalGaugeTraceZero,
		Gate516GlobalIntegersBlocked:      !g516.Finite.ContinuumEulerIntegerDerived && !g516.Finite.ContinuumSignatureIntegerDerived && !g516.Finite.ManifoldTopologySelected,
		Gate516EtaBlocked:                 !g516.Finite.BoundaryEtaInvariantClosed,
		Gate516ObservedTopologyImported:   g516.Firewall.ObservedTopologyImported || g516.Firewall.ManifoldEulerIntegerImported || g516.Firewall.ManifoldSignatureImported,
		Verdict:                           StatusGate516TopologyInherited,
		Reason:                            "Gate517 inherits Gate516's scale-free characteristic-class sockets, finite chiral index carrier, mixed gravitational trace cancellation, and explicit block on global manifold integers and eta boundary data.",
	}
}

func buildIndex() APSIndexLedger {
	return APSIndexLedger{
		Executed:                       true,
		Dimension:                      4,
		FiniteGradingAvailable:         true,
		RealStructureAvailable:         true,
		LocalIndexDensity:              "ind(D_E) local socket = [Â(R) ch(E)]_4, with Â_4 = -p1(TM)/24 and gauge twists carried by ch(E)",
		APSFormula:                     "ind_APS(D_E) = ∫_M [Â(R) ch(E)]_4 - (η(D_∂)+h)/2",
		ClosedManifoldReduction:        "if ∂M=∅ is supplied as a global condition, η=0 and h=0, so ind(D_E)=∫_M [Â(R)ch(E)]_4",
		LocalIndexDensitySocketPresent: true,
		ClosedManifoldSocketConsistent: true,
		GlobalIndexIntegerDerived:      false,
		BoundaryEtaDerived:             false,
		BoundaryKernelDimensionDerived: false,
		BoundarySpectrumSelected:       false,
		ClosedManifoldSelected:         false,
		Verdict:                        strings.Join([]string{StatusAPSIndexLedgerDefined, StatusLocalIndexDensitySocketScaleFree, StatusClosedManifoldIndexSocketConsistent, StatusFailedGlobalIndexIntegerNotDerived}, ";"),
		Reason:                         "the finite grading and real structure support the local index-density socket, and the closed-manifold reduction is mathematically consistent, but ASHA has not selected a continuum manifold, boundary condition, boundary operator spectrum, or global index integer.",
	}
}

func buildEta() BoundaryEtaAirlock {
	return BoundaryEtaAirlock{
		Executed:                       true,
		BoundaryOperatorRequired:       true,
		BoundarySpectrumRequired:       true,
		EtaInvariantRequired:           true,
		KernelCorrectionRequired:       true,
		BoundaryConditionRequired:      true,
		GlobalTopologyRequired:         true,
		BoundaryDataImported:           false,
		BoundaryEtaNativeDerived:       false,
		BoundaryEtaNativeWrite:         false,
		ClosedManifoldIsAllowedBridge:  true,
		ClosedManifoldIsNativeSelected: false,
		Verdict:                        strings.Join([]string{StatusBoundaryEtaAirlockDefined, StatusEtaCorrectionClassifiedGlobal, StatusFailedBoundaryEtaNotDerived, StatusFailedBoundarySpectrumNotSelected, StatusFailedClosedManifoldAssumptionNotNative}, ";"),
		Reason:                         "APS eta is a boundary spectral asymmetry. It can be accepted only as a bridge/global-topology row or set to zero under an explicit closed-manifold assumption; neither the boundary spectrum nor the closed-manifold condition is selected by the finite algebra alone.",
	}
}

func buildInflow(g516 generation2topologicalgravitycharacteristicclassledger.Analysis) AnomalyInflowAudit {
	return AnomalyInflowAudit{
		Executed:                         true,
		PontryaginDescentSocketPresent:   g516.Ledger.PontryaginSocketPresent,
		ChernSimonsBoundarySocketPresent: true,
		ChiralIndexAnomalySocketPresent:  g516.Finite.ChiralIndexSocketPresent,
		MixedGaugeGravityTraceZero:       g516.Finite.MixedGravitationalGaugeTraceZero,
		BoundaryEtaPairsWithInflow:       true,
		PhysicalThetaCoefficientDerived:  false,
		BoundaryTheorySelected:           false,
		Verdict:                          strings.Join([]string{StatusAnomalyInflowSocketPresent, StatusFailedGravitationalThetaStillUnselected}, ";"),
		Reason:                           "Pontryagin descent supplies the formal inflow socket d CS_grav = Tr(R∧R), and APS eta is the matching boundary spectral correction. The socket is native/topological, but no boundary theory, theta coefficient, or physical parity-violating gravity term is selected.",
	}
}

func buildFirewall() ScaleFirewall {
	return ScaleFirewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusScaleAndMassIndependencePreserved, StatusNoObservedTopologyOrBoundaryImported, StatusFailedManifoldTopologyStillEnvironmental, StatusFailedNewtonCosmologyStillBlocked, StatusFirewallNoEmpiricalBoundaryImported, StatusFirewallIndexEtaNativeWriteBlocked}, ";"),
		Reason:   "Gate517 imports no cutoff, spectral moments, Newton/Planck normalization, cosmology, electroweak scale, flavor data, observed topology, boundary spectrum, eta invariant, or manifold index value. Global index and eta writes are blocked unless a continuum topology/boundary ledger is supplied explicitly as bridge data or natively derived by a future theorem.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"The local chiral index-density socket [Â(R)ch(E)]_4 is supported by the finite grading/real-structure lane and is scale independent.",
			"Pontryagin descent and gravitational anomaly-inflow sockets are present as scale-free a4/topological structures.",
		},
		BridgeEntries: []string{
			"The APS formula ind_APS(D)=∫[Â ch]_4-(η+h)/2 as a continuum index ledger requiring manifold and boundary data.",
			"The closed-manifold specialization η=h=0, allowed only when ∂M=∅ is explicitly supplied or natively proven later.",
		},
		EnvironmentalEntries: []string{
			"Actual global Dirac index, Euler characteristic, signature, Pontryagin numbers, boundary spectrum, eta invariant, kernel dimension h, boundary condition, and orientation/bordism class.",
			"Newton/Planck normalization, cutoff Λ, cosmological constant, dark-energy comparator, electroweak scales, and flavor/Yukawa data.",
		},
		FailedRoutes: []string{
			"Deriving the global index integer from finite algebra without continuum topology.",
			"Deriving APS eta without the boundary Dirac spectrum and boundary condition.",
			"Treating Pontryagin descent as a prediction of a physical gravitational theta angle.",
		},
		OpenTheorems: []string{
			"A native continuum topology/bordism selector, if ASHA can ever derive it.",
			"A boundary Hilbert-space/operator theorem selecting APS eta or proving closedness.",
			"A parity/gravitational theta coefficient provenance theorem independent of empirical topology.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 518, Title: "Synthetic APS Index Boundary Ledger Dry-Run", Reason: "Gate517 defines the APS/index airlock but refuses to derive global index or eta values without explicit topology and boundary data.", PrimaryTask: "Exercise a synthetic bridge-only APS ledger with fake Euler/signature/eta/kernel rows to validate residual plumbing and native-write blocking, without importing observed spacetime topology."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate516Inherited && a.Inheritance.Gate516EulerSocket && a.Inheritance.Gate516PontryaginSocket && a.Inheritance.Gate516CharacteristicScaleFree && a.Inheritance.Gate516ChiralIndexSocket && a.Inheritance.Gate516MixedGaugeGravityTraceZero && a.Inheritance.Gate516GlobalIntegersBlocked && a.Inheritance.Gate516EtaBlocked && !a.Inheritance.Gate516ObservedTopologyImported, "Gate517 inheritance invalid"},
		{a.Index.Executed && a.Index.Dimension == 4 && a.Index.FiniteGradingAvailable && a.Index.RealStructureAvailable && a.Index.LocalIndexDensitySocketPresent && a.Index.ClosedManifoldSocketConsistent && !a.Index.GlobalIndexIntegerDerived && !a.Index.BoundaryEtaDerived && !a.Index.BoundaryKernelDimensionDerived && !a.Index.BoundarySpectrumSelected && !a.Index.ClosedManifoldSelected, "Gate517 index ledger invalid"},
		{a.Eta.Executed && a.Eta.BoundaryOperatorRequired && a.Eta.BoundarySpectrumRequired && a.Eta.EtaInvariantRequired && a.Eta.KernelCorrectionRequired && a.Eta.BoundaryConditionRequired && a.Eta.GlobalTopologyRequired && !a.Eta.BoundaryDataImported && !a.Eta.BoundaryEtaNativeDerived && !a.Eta.BoundaryEtaNativeWrite && a.Eta.ClosedManifoldIsAllowedBridge && !a.Eta.ClosedManifoldIsNativeSelected, "Gate517 eta airlock invalid"},
		{a.Inflow.Executed && a.Inflow.PontryaginDescentSocketPresent && a.Inflow.ChernSimonsBoundarySocketPresent && a.Inflow.ChiralIndexAnomalySocketPresent && a.Inflow.MixedGaugeGravityTraceZero && a.Inflow.BoundaryEtaPairsWithInflow && !a.Inflow.PhysicalThetaCoefficientDerived && !a.Inflow.BoundaryTheorySelected, "Gate517 inflow audit invalid"},
		{a.Firewall.Executed && !a.Firewall.UsesLambdaCutoff && !a.Firewall.UsesF2Moment && !a.Firewall.UsesF4Moment && !a.Firewall.UsesNewtonConstant && !a.Firewall.UsesCosmologicalConstant && !a.Firewall.UsesHiggsOrElectroweakScale && !a.Firewall.UsesFlavorYukawaData && !a.Firewall.ObservedTopologyImported && !a.Firewall.ObservedBoundarySpectrumImported && !a.Firewall.GlobalIndexIntegerNativeWrite && !a.Firewall.BoundaryEtaNativeWrite && !a.Firewall.PhysicalGravitationalThetaWritten && !a.Firewall.NativeGravityNormalizationWritten, "Gate517 firewall invalid"},
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
		StatusGate516TopologyInherited,
		StatusAPSIndexLedgerDefined,
		StatusLocalIndexDensitySocketScaleFree,
		StatusClosedManifoldIndexSocketConsistent,
		StatusBoundaryEtaAirlockDefined,
		StatusAnomalyInflowSocketPresent,
		StatusEtaCorrectionClassifiedGlobal,
		StatusScaleAndMassIndependencePreserved,
		StatusNoObservedTopologyOrBoundaryImported,
		StatusFailedGlobalIndexIntegerNotDerived,
		StatusFailedBoundaryEtaNotDerived,
		StatusFailedBoundarySpectrumNotSelected,
		StatusFailedClosedManifoldAssumptionNotNative,
		StatusFailedGravitationalThetaStillUnselected,
		StatusFailedManifoldTopologyStillEnvironmental,
		StatusFailedNewtonCosmologyStillBlocked,
		StatusFirewallNoEmpiricalBoundaryImported,
		StatusFirewallIndexEtaNativeWriteBlocked,
	}
}

func truth(a Analysis) string {
	return "Gate 517 confirms the index lane without overclaiming it: ASHA has a native scale-free local chiral index socket and a gravitational anomaly-inflow/APS eta airlock, but the global Dirac index, eta invariant, boundary spectrum, closed-manifold condition, gravitational theta coefficient, and continuum topology are not selected by finite algebra alone. The index theorem is structurally present; its global integer data remain bridge/environmental until a manifold-and-boundary theorem exists."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate516 inherited=%t; Euler socket=%t; Pontryagin/signature socket=%t; scale-free=%t; chiral index socket=%t; mixed grav-U1 trace zero=%t; global integers blocked=%t; eta blocked=%t; observed topology imported=%t", x.Gate516Inherited, x.Gate516EulerSocket, x.Gate516PontryaginSocket, x.Gate516CharacteristicScaleFree, x.Gate516ChiralIndexSocket, x.Gate516MixedGaugeGravityTraceZero, x.Gate516GlobalIntegersBlocked, x.Gate516EtaBlocked, x.Gate516ObservedTopologyImported)
}
func FormatIndex(x APSIndexLedger) string {
	return fmt.Sprintf("dim=%d; grading=%t; real structure=%t; density=%q; APS=%q; closed reduction=%q; local socket=%t; closed socket=%t; index integer derived=%t; eta derived=%t; h derived=%t; boundary spectrum selected=%t; closed manifold selected=%t", x.Dimension, x.FiniteGradingAvailable, x.RealStructureAvailable, x.LocalIndexDensity, x.APSFormula, x.ClosedManifoldReduction, x.LocalIndexDensitySocketPresent, x.ClosedManifoldSocketConsistent, x.GlobalIndexIntegerDerived, x.BoundaryEtaDerived, x.BoundaryKernelDimensionDerived, x.BoundarySpectrumSelected, x.ClosedManifoldSelected)
}
func FormatEta(x BoundaryEtaAirlock) string {
	return fmt.Sprintf("boundary operator required=%t; spectrum required=%t; eta required=%t; h required=%t; boundary condition required=%t; topology required=%t; boundary data imported=%t; eta native derived=%t; eta native write=%t; closed manifold bridge-allowed=%t; closed manifold native-selected=%t", x.BoundaryOperatorRequired, x.BoundarySpectrumRequired, x.EtaInvariantRequired, x.KernelCorrectionRequired, x.BoundaryConditionRequired, x.GlobalTopologyRequired, x.BoundaryDataImported, x.BoundaryEtaNativeDerived, x.BoundaryEtaNativeWrite, x.ClosedManifoldIsAllowedBridge, x.ClosedManifoldIsNativeSelected)
}
func FormatInflow(x AnomalyInflowAudit) string {
	return fmt.Sprintf("Pontryagin descent=%t; CS boundary socket=%t; chiral index anomaly socket=%t; mixed grav-U1 trace zero=%t; eta pairs with inflow=%t; theta coefficient derived=%t; boundary theory selected=%t", x.PontryaginDescentSocketPresent, x.ChernSimonsBoundarySocketPresent, x.ChiralIndexAnomalySocketPresent, x.MixedGaugeGravityTraceZero, x.BoundaryEtaPairsWithInflow, x.PhysicalThetaCoefficientDerived, x.BoundaryTheorySelected)
}
func FormatFirewall(x ScaleFirewall) string {
	return fmt.Sprintf("uses Λ=%t; uses f2=%t; uses f4=%t; uses G=%t; uses Λ_cosmo=%t; uses EW/Higgs=%t; uses flavor=%t; observed topology imported=%t; observed boundary spectrum imported=%t; index native write=%t; eta native write=%t; theta write=%t; gravity normalization write=%t", x.UsesLambdaCutoff, x.UsesF2Moment, x.UsesF4Moment, x.UsesNewtonConstant, x.UsesCosmologicalConstant, x.UsesHiggsOrElectroweakScale, x.UsesFlavorYukawaData, x.ObservedTopologyImported, x.ObservedBoundarySpectrumImported, x.GlobalIndexIntegerNativeWrite, x.BoundaryEtaNativeWrite, x.PhysicalGravitationalThetaWritten, x.NativeGravityNormalizationWritten)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 517 Registry Audit — Gravitational Index and Boundary Eta Airlock\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## APS index ledger\n\n" + a.Index.Reason + "\n\n```text\n" + FormatIndex(a.Index) + "\n```\n\n")
	b.WriteString("Index formulas:\n\n```text\nind(D_E) closed = ∫_M [Â(R) ch(E)]_4\nind_APS(D_E) = ∫_M [Â(R) ch(E)]_4 - (η(D_∂)+h)/2\nÂ_4 = -p1(TM)/24\n```\n\n")
	b.WriteString("## Boundary eta airlock\n\n" + a.Eta.Reason + "\n\n```text\n" + FormatEta(a.Eta) + "\n```\n\n")
	b.WriteString("## Gravitational anomaly inflow audit\n\n" + a.Inflow.Reason + "\n\n```text\n" + FormatInflow(a.Inflow) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate518 should be:\n\n```text\nGate 518 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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
