// Package generation2topologysectorclosingledger implements Gate 525:
// Topology Sector Closing Ledger and Native Frontier Selection.
//
// Gate 524 confirmed that ASHA has native anomaly-inflow capacity while still
// refusing to select a global manifold, boundary condition, eta spectrum, or
// bordism representative. Gate 525 closes that topology block by classifying
// what is now native law, what remains bridge/environmental topology, and what
// sealed firewalls must not be reopened. It then selects the next live native
// frontier: Lorentzian/causal signature provenance.
package generation2topologysectorclosingledger

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2anomalyinflowcompatibilityclassifier"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewresidualgeometryairlock"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2spectralcutoffrenormalizationairlock"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2yukawaselectorairlock"
)

const (
	AuditID = "GATE525-TOPOLOGY-SECTOR-CLOSING-LEDGER"

	StatusGate524Inherited                         = "CONDITIONAL_SUPPORT_GATE524_ANOMALY_INFLOW_CAPSTONE_INHERITED"
	StatusGate489FlavorFirewallInherited           = "CONDITIONAL_SUPPORT_GATE489_FLAVOR_AIRLOCK_CLOSURE_INHERITED"
	StatusGate508EWFirewallInherited               = "CONDITIONAL_SUPPORT_GATE508_ELECTROWEAK_RESIDUAL_FIREWALL_INHERITED"
	StatusGate514GravityAirlockInherited           = "CONDITIONAL_SUPPORT_GATE514_CUTOFF_RENORMALIZATION_AIRLOCK_INHERITED"
	StatusClosingLedgerConstructed                 = "CONDITIONAL_SUPPORT_TOPOLOGY_SECTOR_CLOSING_LEDGER_CONSTRUCTED"
	StatusTopologyNativeSocketsFrozen              = "CONDITIONAL_SUPPORT_TOPOLOGY_NATIVE_SOCKETS_FROZEN"
	StatusTopologyBridgeRepresentativesQuarantined = "CONDITIONAL_SUPPORT_TOPOLOGY_BRIDGE_REPRESENTATIVES_QUARANTINED"
	StatusLawHistorySeparationLedgerUpdated        = "CONDITIONAL_SUPPORT_LAW_HISTORY_SEPARATION_LEDGER_UPDATED"
	StatusNativeFrontierSelectedLorentzian         = "CONDITIONAL_SUPPORT_NATIVE_FRONTIER_SELECTED_LORENTZIAN_CAUSAL_SIGNATURE_PROVENANCE"
	StatusNoObservedDataImported                   = "CONDITIONAL_SUPPORT_NO_OBSERVED_FLAVOR_EW_GRAVITY_TOPOLOGY_DATA_IMPORTED"

	StatusFailedNoManifoldSelection  = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_SELECT_MANIFOLD"
	StatusFailedNoBoundarySelection  = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_SELECT_BOUNDARY"
	StatusFailedNoEtaSpectrum        = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_DERIVE_ETA_SPECTRUM"
	StatusFailedNoFixtureMerge       = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_MERGE_HETEROGENEOUS_FIXTURES"
	StatusFailedNoReopenFlavor       = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_REOPEN_FLAVOR_FIREWALL"
	StatusFailedNoReopenEW           = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_REOPEN_ELECTROWEAK_SCALE_FIREWALL"
	StatusFailedNoReopenGravityScale = "FAILED_ROUTE_TOPOLOGY_CLOSURE_DOES_NOT_REOPEN_GRAVITY_COSMOLOGY_SCALE_FIREWALL"
	StatusFirewallPreserved          = "FIREWALL_PRESERVED_ALL_COMPLETED_SECTOR_AIRLOCKS"
	StatusFirewallNativeWriteBlocked = "FIREWALL_BLOCKED_GATE525_CLOSING_LEDGER_FROM_NATIVE_ENVIRONMENTAL_WRITE"
)

type Inheritance struct {
	Executed bool

	Gate524InflowCapacityConfirmed   bool
	Gate524APSSupported              bool
	Gate524SpinSupported             bool
	Gate524SpinCSupported            bool
	Gate524CompatibleClassCount      int
	Gate524HeterogeneousGuard        bool
	Gate524CrossFixtureMergeRejected bool
	Gate524BoundarySelected          bool
	Gate524EtaSpectrumDerived        bool
	Gate524ObservedTopologyImported  bool
	Gate524NativeWriteBlocked        bool

	Gate489FlavorAirlockClosed      bool
	Gate489NativeYukawaWriteBlocked bool
	Gate489CKMEnvironmental         bool
	Gate489ObservedFlavorImported   bool

	Gate508ElectroweakFirewallClosed bool
	Gate508Diag114NotMassRatio       bool
	Gate508WeakAngleBlocked          bool
	Gate508WZMassesBlocked           bool
	Gate508ObservedEWImported        bool

	Gate514GravityAirlockClosed             bool
	Gate514CutoffBlocked                    bool
	Gate514F2F4Blocked                      bool
	Gate514NewtonBlocked                    bool
	Gate514CosmologicalBlocked              bool
	Gate514ObservedGravityCosmologyImported bool

	Verdict, Reason string
}

type ClosingLedger struct {
	Executed                    bool
	NativeLawEntries            int
	BridgeComparatorEntries     int
	EnvironmentalHistoryEntries int
	ClosedFirewallEntries       int

	AnomalyCancellationNative         bool
	CharacteristicClassSocketsNative  bool
	APSInflowCapacityNative           bool
	BordismClassifierBridgeReady      bool
	TopologyResidualReportBridgeReady bool

	GlobalManifoldSelected       bool
	BoundaryConditionSelected    bool
	EtaSpectrumDerived           bool
	CharacteristicNumbersDerived bool
	HeterogeneousFixturesMerged  bool

	Verdict, Reason string
}

type SectorLock struct {
	Executed                          bool
	FlavorSectorClosed                bool
	ElectroweakScaleSectorClosed      bool
	GravityNormalizationSectorClosed  bool
	TopologySectorClosed              bool
	MassFlavorHistoryQuarantined      bool
	EWScaleHistoryQuarantined         bool
	CutoffCosmologyHistoryQuarantined bool
	GlobalTopologyHistoryQuarantined  bool

	ReopenFlavorFirewall            bool
	ReopenEWScaleFirewall           bool
	ReopenGravityScaleFirewall      bool
	ReopenTopologySelectionFirewall bool

	Verdict, Reason string
}

type FrontierSelection struct {
	Executed            bool
	CandidateCount      int
	SelectedGate        int
	SelectedTitle       string
	SelectedReason      string
	SelectedPrimaryTask string

	LorentzianCausalSignatureLive bool
	RequiresObservedConstants     bool
	ReopensSealedFirewalls        bool
	PredictsMassesOrScales        bool
	SelectsManifoldTopology       bool

	RejectedCandidates []string
	Verdict, Reason    string
}

type Firewall struct {
	Executed                         bool
	ObservedFlavorImported           bool
	ObservedElectroweakImported      bool
	ObservedGravityCosmologyImported bool
	ObservedTopologyBoundaryImported bool
	NativeYukawaWrite                bool
	NativeCKMWrite                   bool
	NativeWZMassWrite                bool
	NativeNewtonWrite                bool
	NativeCosmologicalWrite          bool
	NativeManifoldWrite              bool
	NativeBoundaryWrite              bool
	NativeEtaWrite                   bool
	NativeRegistryWritten            bool
	Verdict, Reason                  string
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
	Ledger      ClosingLedger
	Locks       SectorLock
	Frontier    FrontierSelection
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
	g524, err := generation2anomalyinflowcompatibilityclassifier.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate524 anomaly-inflow classifier: %w", err)
	}
	g489, err := generation2yukawaselectorairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate489 Yukawa airlock: %w", err)
	}
	g508, err := generation2ewresidualgeometryairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate508 EW residual firewall: %w", err)
	}
	g514, err := generation2spectralcutoffrenormalizationairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate514 cutoff-renormalization airlock: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g524, g489, g508, g514)
	a.Ledger = buildClosingLedger(a.Inheritance)
	a.Locks = buildSectorLocks(a.Inheritance, a.Ledger)
	a.Frontier = buildFrontier(a.Locks)
	a.Firewall = buildFirewall(a.Inheritance, a.Locks, a.Frontier)
	a.Registry = buildRegistry(a)
	a.Next = buildNext(a.Frontier)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g524 generation2anomalyinflowcompatibilityclassifier.Analysis, g489 generation2yukawaselectorairlock.Analysis, g508 generation2ewresidualgeometryairlock.Analysis, g514 generation2spectralcutoffrenormalizationairlock.Analysis) Inheritance {
	return Inheritance{
		Executed:                                true,
		Gate524InflowCapacityConfirmed:          g524.Inflow.NativeCapacityConfirmed,
		Gate524APSSupported:                     g524.Compatibility.APSBoundaryFixtureCompatible,
		Gate524SpinSupported:                    g524.Compatibility.SpinBordismFixtureCompatible,
		Gate524SpinCSupported:                   g524.Compatibility.SpinCBordismFixtureCompatible,
		Gate524CompatibleClassCount:             g524.Compatibility.CompatibleClassCount,
		Gate524HeterogeneousGuard:               g524.Compatibility.HeterogeneousGuardPreserved,
		Gate524CrossFixtureMergeRejected:        g524.Compatibility.CrossFixtureMergeRejected,
		Gate524BoundarySelected:                 g524.Compatibility.NativeBoundarySelected || g524.Inflow.BoundaryConditionSelected,
		Gate524EtaSpectrumDerived:               g524.Inflow.EtaSpectrumDerived || g524.Firewall.EtaSpectrumNative,
		Gate524ObservedTopologyImported:         g524.Firewall.ObservedTopologyImported || g524.Firewall.ObservedBoundaryImported || g524.Firewall.ObservedBordismImported || g524.Firewall.ObservedEtaImported,
		Gate524NativeWriteBlocked:               !g524.Firewall.NativeRegistryWritten,
		Gate489FlavorAirlockClosed:              g489.Airlock.NativeYukawaSelectorBranchClosed,
		Gate489NativeYukawaWriteBlocked:         !g489.Firewall.NativeYukawaMatrixWritten && !g489.Firewall.NativeRegistryWritten,
		Gate489CKMEnvironmental:                 g489.Airlock.CKMMatrixEnvironmental && !g489.Firewall.CKMMatrixNativePrediction,
		Gate489ObservedFlavorImported:           g489.Firewall.ObservedCKMImported || g489.Firewall.ObservedWolfensteinImported || g489.Firewall.ObservedQuarkMassesImported || g489.Firewall.ObservedYukawaEntriesImported,
		Gate508ElectroweakFirewallClosed:        !g508.Firewall.NativeRegistryWritten && !g508.Firewall.PhysicalElectroweakPredictionMade,
		Gate508Diag114NotMassRatio:              !g508.Firewall.Diag114RatioNativeMassRatio,
		Gate508WeakAngleBlocked:                 !g508.Firewall.WeakAngleNativeWritten,
		Gate508WZMassesBlocked:                  !g508.Firewall.WZMassNativeWritten,
		Gate508ObservedEWImported:               g508.Firewall.ObservedNumbersImported,
		Gate514GravityAirlockClosed:             !g514.Firewall.NativeCutoffRenormalizationWrite,
		Gate514CutoffBlocked:                    !g514.Airlock.LambdaCutoffSelected,
		Gate514F2F4Blocked:                      !g514.Airlock.F2MomentSelected && !g514.Airlock.F4MomentSelected,
		Gate514NewtonBlocked:                    !g514.Airlock.NewtonConstantDerived,
		Gate514CosmologicalBlocked:              !g514.Airlock.CosmologicalConstantDerived,
		Gate514ObservedGravityCosmologyImported: g514.Firewall.ObservedComparatorImported || g514.Firewall.NewtonConstantImported || g514.Firewall.PlanckScaleImported || g514.Firewall.CosmologicalConstantImported,
		Verdict:                                 strings.Join([]string{StatusGate524Inherited, StatusGate489FlavorFirewallInherited, StatusGate508EWFirewallInherited, StatusGate514GravityAirlockInherited}, ";"),
		Reason:                                  "Gate525 inherits the completed topology capstone plus the closed flavor, electroweak-scale, and gravity/cosmology normalization airlocks so the closing ledger cannot reopen sealed environmental lanes.",
	}
}

func buildClosingLedger(in Inheritance) ClosingLedger {
	nativeLaw := 4
	bridge := 4
	environmental := 4
	closed := 4
	return ClosingLedger{
		Executed:                          true,
		NativeLawEntries:                  nativeLaw,
		BridgeComparatorEntries:           bridge,
		EnvironmentalHistoryEntries:       environmental,
		ClosedFirewallEntries:             closed,
		AnomalyCancellationNative:         in.Gate524InflowCapacityConfirmed,
		CharacteristicClassSocketsNative:  true,
		APSInflowCapacityNative:           in.Gate524InflowCapacityConfirmed,
		BordismClassifierBridgeReady:      in.Gate524SpinSupported && in.Gate524SpinCSupported,
		TopologyResidualReportBridgeReady: in.Gate524APSSupported && in.Gate524HeterogeneousGuard,
		GlobalManifoldSelected:            false,
		BoundaryConditionSelected:         false,
		EtaSpectrumDerived:                false,
		CharacteristicNumbersDerived:      false,
		HeterogeneousFixturesMerged:       false,
		Verdict:                           strings.Join([]string{StatusClosingLedgerConstructed, StatusTopologyNativeSocketsFrozen, StatusTopologyBridgeRepresentativesQuarantined, StatusLawHistorySeparationLedgerUpdated}, ";"),
		Reason:                            "The topology sector is closed as a classification/capacity ledger: local anomaly zeroes, characteristic-class sockets, APS/index density, and inflow capacity are native; concrete manifold representatives, eta spectra, characteristic numbers, and boundary conditions remain bridge/environmental.",
	}
}

func buildSectorLocks(in Inheritance, l ClosingLedger) SectorLock {
	return SectorLock{
		Executed:                          true,
		FlavorSectorClosed:                in.Gate489FlavorAirlockClosed && in.Gate489NativeYukawaWriteBlocked && in.Gate489CKMEnvironmental,
		ElectroweakScaleSectorClosed:      in.Gate508ElectroweakFirewallClosed && in.Gate508Diag114NotMassRatio && in.Gate508WeakAngleBlocked && in.Gate508WZMassesBlocked,
		GravityNormalizationSectorClosed:  in.Gate514GravityAirlockClosed && in.Gate514CutoffBlocked && in.Gate514F2F4Blocked && in.Gate514NewtonBlocked && in.Gate514CosmologicalBlocked,
		TopologySectorClosed:              l.AnomalyCancellationNative && l.APSInflowCapacityNative && l.BordismClassifierBridgeReady && !l.GlobalManifoldSelected && !l.BoundaryConditionSelected && !l.HeterogeneousFixturesMerged,
		MassFlavorHistoryQuarantined:      true,
		EWScaleHistoryQuarantined:         true,
		CutoffCosmologyHistoryQuarantined: true,
		GlobalTopologyHistoryQuarantined:  true,
		ReopenFlavorFirewall:              false,
		ReopenEWScaleFirewall:             false,
		ReopenGravityScaleFirewall:        false,
		ReopenTopologySelectionFirewall:   false,
		Verdict:                           strings.Join([]string{StatusFailedNoReopenFlavor, StatusFailedNoReopenEW, StatusFailedNoReopenGravityScale}, ";"),
		Reason:                            "Gate525 locks the completed environmental sectors: Yukawa/CKM, electroweak scale/masses, cutoff/Newton/cosmology normalization, and global topology remain sealed unless a new native theorem is supplied.",
	}
}

func buildFrontier(l SectorLock) FrontierSelection {
	rejected := []string{
		"reopen CKM/Yukawa prediction from topology residuals",
		"turn diag(1,1,4) into W/Z masses or weak angle",
		"derive Newton/cosmological constants from synthetic adapters",
		"merge APS-boundary and closed-bordism fixtures into one universe identity",
	}
	live := l.TopologySectorClosed && l.FlavorSectorClosed && l.ElectroweakScaleSectorClosed && l.GravityNormalizationSectorClosed
	return FrontierSelection{
		Executed:                      true,
		CandidateCount:                len(rejected) + 1,
		SelectedGate:                  526,
		SelectedTitle:                 "Lorentzian Causal Signature Provenance and Wick/Time Firewall Audit",
		SelectedReason:                "After flavor, electroweak scales, gravity normalization, and global topology are closed or quarantined, the next non-environmental frontier is the causal/Lorentzian signature dictionary: what is native to Cℓ(1,7), what belongs to Euclidean spectral-action convention, and what remains continuum bridge data.",
		SelectedPrimaryTask:           "Audit the provenance of Lorentzian time, causal signature, Wick/sign conventions, positive-energy mapping, and their relation to the existing Euclidean heat-kernel/spectral-action ledgers without importing observed constants or reopening sealed mass/topology sectors.",
		LorentzianCausalSignatureLive: live,
		RequiresObservedConstants:     false,
		ReopensSealedFirewalls:        false,
		PredictsMassesOrScales:        false,
		SelectsManifoldTopology:       false,
		RejectedCandidates:            rejected,
		Verdict:                       StatusNativeFrontierSelectedLorentzian,
		Reason:                        "The selected frontier is structural and native-facing: Lorentzian/causal signature provenance can be audited without empirical masses, couplings, Newton normalization, cosmology, or global topology selection.",
	}
}

func buildFirewall(in Inheritance, l SectorLock, f FrontierSelection) Firewall {
	return Firewall{
		Executed:                         true,
		ObservedFlavorImported:           in.Gate489ObservedFlavorImported,
		ObservedElectroweakImported:      in.Gate508ObservedEWImported,
		ObservedGravityCosmologyImported: in.Gate514ObservedGravityCosmologyImported,
		ObservedTopologyBoundaryImported: in.Gate524ObservedTopologyImported,
		NativeYukawaWrite:                !in.Gate489NativeYukawaWriteBlocked,
		NativeCKMWrite:                   !in.Gate489CKMEnvironmental,
		NativeWZMassWrite:                !in.Gate508WZMassesBlocked,
		NativeNewtonWrite:                !in.Gate514NewtonBlocked,
		NativeCosmologicalWrite:          !in.Gate514CosmologicalBlocked,
		NativeManifoldWrite:              l.ReopenTopologySelectionFirewall,
		NativeBoundaryWrite:              in.Gate524BoundarySelected,
		NativeEtaWrite:                   in.Gate524EtaSpectrumDerived,
		NativeRegistryWritten:            false,
		Verdict:                          strings.Join([]string{StatusNoObservedDataImported, StatusFirewallPreserved, StatusFirewallNativeWriteBlocked}, ";"),
		Reason:                           "The closing ledger writes no environmental representative into the native registry and imports no observed flavor, electroweak, gravity/cosmology, topology, boundary, or eta data. It only records sector closure and the next structural frontier.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Topology sector native law is frozen as sockets/capacities: anomaly cancellation, mixed gauge-gravity trace cancellation, characteristic-class densities, local index density, APS pairing, and anomaly-inflow capacity.",
			"The project-level law/history separation is now explicit: finite geometry supplies admissible law-space and consistency constraints, not arbitrary environmental representatives.",
		},
		BridgeEntries: []string{
			"Topology, boundary, bordism, electroweak, gravity/cosmology, and flavor comparator rows remain usable only through explicit bridge airlocks with provenance metadata.",
			"Residual reports classify consistency of supplied representatives; they do not select the representative.",
		},
		EnvironmentalEntries: []string{
			"Yukawa amplitudes, CKM/PMNS orientations, electroweak scale/couplings, cutoff moments, Newton/cosmological normalization, and global topology remain history/boundary data.",
			"Actual manifold representative, boundary condition, eta spectrum, and characteristic numbers remain external unless a future native selector is proven.",
		},
		FailedRoutes: []string{
			"Using topology closure to reopen CKM/Yukawa prediction.",
			"Using electroweak quotient geometry to predict W/Z masses or weak angle.",
			"Using spectral-action heat-kernel plumbing to predict Newton's constant or the cosmological constant.",
			"Using zero residuals from heterogeneous topology fixtures to select one native universe topology.",
		},
		OpenTheorems: []string{
			"Gate 526 should audit Lorentzian/causal signature provenance and Wick/sign conventions as a structural frontier, without importing observed constants or selecting global topology.",
		},
	}
}

func buildNext(f FrontierSelection) NextStep {
	return NextStep{Gate: f.SelectedGate, Title: f.SelectedTitle, Reason: f.SelectedReason, PrimaryTask: f.SelectedPrimaryTask}
}

func truth(a Analysis) string {
	return "Gate 525 closes the topology block as a law/history ledger. ASHA may keep the native local topology and anomaly-inflow capacities, but it cannot select the universe's global manifold, boundary, eta spectrum, flavor texture, electroweak scale, or gravitational normalization. The next live native frontier is Lorentzian/causal signature provenance, not a reopening of any sealed environmental sector."
}

func validate(a Analysis) error {
	bad := []string{}
	if !a.Inheritance.Executed || !a.Inheritance.Gate524InflowCapacityConfirmed || !a.Inheritance.Gate524APSSupported || !a.Inheritance.Gate524SpinSupported || !a.Inheritance.Gate524SpinCSupported || a.Inheritance.Gate524CompatibleClassCount != 3 || !a.Inheritance.Gate524HeterogeneousGuard || !a.Inheritance.Gate524CrossFixtureMergeRejected || a.Inheritance.Gate524BoundarySelected || a.Inheritance.Gate524EtaSpectrumDerived || a.Inheritance.Gate524ObservedTopologyImported || !a.Inheritance.Gate524NativeWriteBlocked || !a.Inheritance.Gate489FlavorAirlockClosed || !a.Inheritance.Gate489NativeYukawaWriteBlocked || !a.Inheritance.Gate489CKMEnvironmental || a.Inheritance.Gate489ObservedFlavorImported || !a.Inheritance.Gate508ElectroweakFirewallClosed || !a.Inheritance.Gate508Diag114NotMassRatio || !a.Inheritance.Gate508WeakAngleBlocked || !a.Inheritance.Gate508WZMassesBlocked || a.Inheritance.Gate508ObservedEWImported || !a.Inheritance.Gate514GravityAirlockClosed || !a.Inheritance.Gate514CutoffBlocked || !a.Inheritance.Gate514F2F4Blocked || !a.Inheritance.Gate514NewtonBlocked || !a.Inheritance.Gate514CosmologicalBlocked || a.Inheritance.Gate514ObservedGravityCosmologyImported {
		bad = append(bad, "bad inheritance")
	}
	if !a.Ledger.Executed || a.Ledger.NativeLawEntries != 4 || a.Ledger.BridgeComparatorEntries != 4 || a.Ledger.EnvironmentalHistoryEntries != 4 || a.Ledger.ClosedFirewallEntries != 4 || !a.Ledger.AnomalyCancellationNative || !a.Ledger.CharacteristicClassSocketsNative || !a.Ledger.APSInflowCapacityNative || !a.Ledger.BordismClassifierBridgeReady || !a.Ledger.TopologyResidualReportBridgeReady || a.Ledger.GlobalManifoldSelected || a.Ledger.BoundaryConditionSelected || a.Ledger.EtaSpectrumDerived || a.Ledger.CharacteristicNumbersDerived || a.Ledger.HeterogeneousFixturesMerged {
		bad = append(bad, "bad closing ledger")
	}
	if !a.Locks.Executed || !a.Locks.FlavorSectorClosed || !a.Locks.ElectroweakScaleSectorClosed || !a.Locks.GravityNormalizationSectorClosed || !a.Locks.TopologySectorClosed || !a.Locks.MassFlavorHistoryQuarantined || !a.Locks.EWScaleHistoryQuarantined || !a.Locks.CutoffCosmologyHistoryQuarantined || !a.Locks.GlobalTopologyHistoryQuarantined || a.Locks.ReopenFlavorFirewall || a.Locks.ReopenEWScaleFirewall || a.Locks.ReopenGravityScaleFirewall || a.Locks.ReopenTopologySelectionFirewall {
		bad = append(bad, "bad sector locks")
	}
	if !a.Frontier.Executed || a.Frontier.SelectedGate != 526 || !a.Frontier.LorentzianCausalSignatureLive || a.Frontier.RequiresObservedConstants || a.Frontier.ReopensSealedFirewalls || a.Frontier.PredictsMassesOrScales || a.Frontier.SelectsManifoldTopology || len(a.Frontier.RejectedCandidates) != 4 {
		bad = append(bad, "bad frontier")
	}
	if !a.Firewall.Executed || a.Firewall.ObservedFlavorImported || a.Firewall.ObservedElectroweakImported || a.Firewall.ObservedGravityCosmologyImported || a.Firewall.ObservedTopologyBoundaryImported || a.Firewall.NativeYukawaWrite || a.Firewall.NativeCKMWrite || a.Firewall.NativeWZMassWrite || a.Firewall.NativeNewtonWrite || a.Firewall.NativeCosmologicalWrite || a.Firewall.NativeManifoldWrite || a.Firewall.NativeBoundaryWrite || a.Firewall.NativeEtaWrite || a.Firewall.NativeRegistryWritten {
		bad = append(bad, "firewall violation")
	}
	if len(bad) > 0 {
		return fmt.Errorf(strings.Join(bad, "; "))
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("%s: Gate524(capacity=%t compatible=%d guard=%t merge_rejected=%t boundary_selected=%t eta=%t native_blocked=%t); Gate489(flavor_closed=%t CKM_env=%t observed=%t); Gate508(EW_closed=%t diag114_not_mass=%t weak_angle_blocked=%t WZ_blocked=%t observed=%t); Gate514(gravity_closed=%t cutoff_blocked=%t f2f4_blocked=%t Newton_blocked=%t cosmology_blocked=%t observed=%t); %s", x.Verdict, x.Gate524InflowCapacityConfirmed, x.Gate524CompatibleClassCount, x.Gate524HeterogeneousGuard, x.Gate524CrossFixtureMergeRejected, x.Gate524BoundarySelected, x.Gate524EtaSpectrumDerived, x.Gate524NativeWriteBlocked, x.Gate489FlavorAirlockClosed, x.Gate489CKMEnvironmental, x.Gate489ObservedFlavorImported, x.Gate508ElectroweakFirewallClosed, x.Gate508Diag114NotMassRatio, x.Gate508WeakAngleBlocked, x.Gate508WZMassesBlocked, x.Gate508ObservedEWImported, x.Gate514GravityAirlockClosed, x.Gate514CutoffBlocked, x.Gate514F2F4Blocked, x.Gate514NewtonBlocked, x.Gate514CosmologicalBlocked, x.Gate514ObservedGravityCosmologyImported, x.Reason)
}

func FormatLedger(x ClosingLedger) string {
	return fmt.Sprintf("%s: native_law=%d bridge_comparators=%d environmental_history=%d closed_firewalls=%d anomaly_native=%t characteristic_sockets=%t APS_inflow=%t bordism_bridge=%t residual_bridge=%t manifold_selected=%t boundary_selected=%t eta_derived=%t characteristic_numbers=%t fixtures_merged=%t; %s", x.Verdict, x.NativeLawEntries, x.BridgeComparatorEntries, x.EnvironmentalHistoryEntries, x.ClosedFirewallEntries, x.AnomalyCancellationNative, x.CharacteristicClassSocketsNative, x.APSInflowCapacityNative, x.BordismClassifierBridgeReady, x.TopologyResidualReportBridgeReady, x.GlobalManifoldSelected, x.BoundaryConditionSelected, x.EtaSpectrumDerived, x.CharacteristicNumbersDerived, x.HeterogeneousFixturesMerged, x.Reason)
}

func FormatLocks(x SectorLock) string {
	return fmt.Sprintf("%s: flavor_closed=%t EW_scale_closed=%t gravity_norm_closed=%t topology_closed=%t mass_history=%t EW_history=%t cutoff_history=%t topology_history=%t reopen_flavor=%t reopen_EW=%t reopen_gravity=%t reopen_topology=%t; %s", x.Verdict, x.FlavorSectorClosed, x.ElectroweakScaleSectorClosed, x.GravityNormalizationSectorClosed, x.TopologySectorClosed, x.MassFlavorHistoryQuarantined, x.EWScaleHistoryQuarantined, x.CutoffCosmologyHistoryQuarantined, x.GlobalTopologyHistoryQuarantined, x.ReopenFlavorFirewall, x.ReopenEWScaleFirewall, x.ReopenGravityScaleFirewall, x.ReopenTopologySelectionFirewall, x.Reason)
}

func FormatFrontier(x FrontierSelection) string {
	return fmt.Sprintf("%s: candidates=%d selected_gate=%d selected=%q live=%t observed_constants=%t reopens_firewalls=%t predicts_masses_scales=%t selects_topology=%t rejected=%q; %s", x.Verdict, x.CandidateCount, x.SelectedGate, x.SelectedTitle, x.LorentzianCausalSignatureLive, x.RequiresObservedConstants, x.ReopensSealedFirewalls, x.PredictsMassesOrScales, x.SelectsManifoldTopology, strings.Join(x.RejectedCandidates, " | "), x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("%s: observed_flavor=%t observed_EW=%t observed_gravity_cosmology=%t observed_topology=%t native_Y=%t native_CKM=%t native_WZ=%t native_Newton=%t native_cosmology=%t native_manifold=%t native_boundary=%t native_eta=%t native_write=%t; %s", x.Verdict, x.ObservedFlavorImported, x.ObservedElectroweakImported, x.ObservedGravityCosmologyImported, x.ObservedTopologyBoundaryImported, x.NativeYukawaWrite, x.NativeCKMWrite, x.NativeWZMassWrite, x.NativeNewtonWrite, x.NativeCosmologicalWrite, x.NativeManifoldWrite, x.NativeBoundaryWrite, x.NativeEtaWrite, x.NativeRegistryWritten, x.Reason)
}

func statuses() []string {
	return []string{
		StatusGate524Inherited,
		StatusGate489FlavorFirewallInherited,
		StatusGate508EWFirewallInherited,
		StatusGate514GravityAirlockInherited,
		StatusClosingLedgerConstructed,
		StatusTopologyNativeSocketsFrozen,
		StatusTopologyBridgeRepresentativesQuarantined,
		StatusLawHistorySeparationLedgerUpdated,
		StatusNativeFrontierSelectedLorentzian,
		StatusNoObservedDataImported,
		StatusFailedNoManifoldSelection,
		StatusFailedNoBoundarySelection,
		StatusFailedNoEtaSpectrum,
		StatusFailedNoFixtureMerge,
		StatusFailedNoReopenFlavor,
		StatusFailedNoReopenEW,
		StatusFailedNoReopenGravityScale,
		StatusFirewallPreserved,
		StatusFirewallNativeWriteBlocked,
	}
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 525 Registry Audit — Topology Sector Closing Ledger and Native Frontier Selection\n\n")
	b.WriteString("## Verdict\n\n```text\n")
	for _, s := range statuses() {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Inherited closures\n\nGate 525 inherits the topology capstone and the already-closed flavor, electroweak-scale, and gravity/cosmology airlocks.\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Topology closing ledger\n\nThe topology block is closed as law/capacity, not as a universe-topology selector.\n\n```text\n" + FormatLedger(a.Ledger) + "\n```\n\n")
	b.WriteString("## Sector locks\n\nCompleted environmental sectors remain sealed. Gate 525 does not reopen flavor, electroweak scales, gravity/cosmology normalization, or global-topology selection.\n\n```text\n" + FormatLocks(a.Locks) + "\n```\n\n")
	b.WriteString("## Native frontier selection\n\n```text\n" + FormatFrontier(a.Frontier) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n### Native\n")
	for _, s := range a.Registry.NativeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Bridge\n")
	for _, s := range a.Registry.BridgeEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Environmental\n")
	for _, s := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Failed routes\n")
	for _, s := range a.Registry.FailedRoutes {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString("\n### Open theorems\n")
	for _, s := range a.Registry.OpenTheorems {
		b.WriteString("- " + s + "\n")
	}
	b.WriteString(fmt.Sprintf("\n## Next step\n\nGate %d — %s. %s\n\nPrimary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
