// Package generation2topologicaldeformation implements Gate 483:
// Finite Algebraic Deformation Operator Search.
package generation2topologicaldeformation

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID                                        = "GATE483-FINITE-ALGEBRAIC-DEFORMATION-OPERATOR-SEARCH"
	StatusAuditCompleted                           = "CONDITIONAL_SUPPORT_TOPOLOGICAL_DEFORMATION_AUDIT_COMPLETED"
	StatusInheritedNullPerturbationFrontier        = "CONDITIONAL_SUPPORT_GATE482_SECTOR_PERTURBATION_FRONTIER_INHERITED"
	StatusTopologicalSectorSeparationAudited       = "CONDITIONAL_SUPPORT_TOPOLOGICAL_SECTOR_SEPARATION_AUDITED"
	StatusQuarkLeptonSeparationOnly                = "CONDITIONAL_SUPPORT_QUARK_LEPTON_TOPOLOGICAL_SEPARATION_ONLY"
	StatusBridgeSlotPreserved                      = "CONDITIONAL_SUPPORT_TOPOLOGICAL_PERTURBATION_BRIDGE_SLOT_PRESERVED"
	StatusFirewallPreserved                        = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_TOPOLOGICAL_SEARCH"
	StatusFailedNativeTopologicalSourceAbsent      = "FAILED_ROUTE_NATIVE_TOPOLOGICAL_DEFORMATION_OPERATOR_ABSENT"
	StatusFailedColorWindingGenerationBlind        = "FAILED_ROUTE_COLOR_WINDING_GENERATION_BLIND"
	StatusFailedHolonomyNoDeltaMap                 = "FAILED_ROUTE_HOLONOMY_LABELS_NO_DELTA_ALPHA_DELTA_PHI_MAP"
	StatusFailedBettiLedgerNotPresent              = "FAILED_ROUTE_BETTI_DEFORMATION_LEDGER_NOT_NATIVE_IN_CURRENT_ATLAS"
	StatusFailedSingleElectronNotFiniteOperator    = "FAILED_ROUTE_SINGLE_ELECTRON_FLOW_NOT_IMPLEMENTED_AS_FINITE_CLIFFORD_OPERATOR"
	StatusFailedGaugeRepresentationGenerationBlind = "FAILED_ROUTE_GAUGE_REPRESENTATION_TOPOLOGY_SECTOR_AWARE_BUT_GENERATION_BLIND"
	StatusFailedYukawaEnvironmental                = "FAILED_ROUTE_TOPOLOGICAL_STRESS_COLLAPSES_TO_SEALED_YUKAWA_LEDGER"
	StatusFailedCKMPMNSRejected                    = "FAILED_ROUTE_CKM_PMNS_AS_TOPOLOGICAL_DEFORMATION_SOURCE_REJECTED"
	StatusFailedNativePromotion                    = "FAILED_ROUTE_TOPOLOGICAL_DEFORMATION_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed, Gate480NullBaseline, Gate481Cancellation, Gate482SourceAbsent bool
	AlphaVac, IKVac                                                         float64
	SectorPerturbationsUnsolved, NativeRegistryClean                        bool
	Verdict                                                                 string
}

type TopologicalCandidate struct {
	Name                                                                                                                                        string
	NativeObject, DiscreteTopological, SectorAware, QuarkLeptonSeparating, GenerationAware, CanSourceDeltaAlpha, CanSourceDeltaPhi, NeedsBridge bool
	RejectedFailure, Reason                                                                                                                     string
}

type TopologicalAudit struct {
	Executed                                                                                bool
	Candidates                                                                              []TopologicalCandidate
	SectorSeparatorsFound, QuarkLeptonOnly, NativeFullSourceFound, BridgeLikeSourcesPresent bool
	Verdict, Reason                                                                         string
}

type GenerationAwarenessTest struct {
	Executed                                                                                   bool
	RequiredDistinctions                                                                       []string
	ColorDistinguishesQuarkLepton, ColorDistinguishesGenerations, ColorDistinguishesUpCharmTop bool
	WindingDistinguishesQuarkLepton, WindingDistinguishesGenerations                           bool
	CandidatesPassingGenerationAwareness                                                       int
	Verdict, Reason                                                                            string
	Failures                                                                                   []string
}

type DeformationMap struct {
	Executed                                                                                                                         bool
	RequiredMap                                                                                                                      string
	TopologicalStressNative, DeltaAlphaMapNative, DeltaPhiMapNative, NumericCoordinateMapNative, TraceCompatible, CKMPMNSIndependent bool
	AllZeroDistance                                                                                                                  float64
	Verdict, Reason                                                                                                                  string
	Failures                                                                                                                         []string
}

type BridgeSlot struct {
	Executed                                                                                                                                                                          bool
	SlotName                                                                                                                                                                          string
	AllowedFields                                                                                                                                                                     []string
	RequiresAirlock, RequiresProvenance, RequiresBranchTags, RequiresUncertainty, AllowsTopologicalLabels, RejectsCKMPMNSAsInput, RejectsNativePromotion, CanComputeSyntheticResidual bool
	CanComputePhysicalResidual                                                                                                                                                        bool
	Verdict, Reason                                                                                                                                                                   string
}

type Firewall struct {
	Executed                                                                                                                                                                    bool
	ObservedMassImported, CKMImported, PMNSImported, VacuumIKNativeBaseline, TopologicalNativeSourceFound, SectorPerturbationsNative, PhysicalDUDComputed, PhysicalDENuComputed bool
	CKMMatrixConstructed, PMNSMatrixConstructed, NativeRegistryWritten                                                                                                          bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                      int
	Verdict, Reason                                                                                                                                                             string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance         Inheritance
	TopologicalAudit    TopologicalAudit
	GenerationAwareness GenerationAwarenessTest
	DeformationMap      DeformationMap
	BridgeSlot          BridgeSlot
	Firewall            Firewall
	Next                NextStep
	Truth               string
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
	a.TopologicalAudit = buildTopologicalAudit()
	a.GenerationAwareness = buildGenerationAwareness(a.TopologicalAudit)
	a.DeformationMap = buildDeformationMap(a.TopologicalAudit, a.GenerationAwareness)
	a.BridgeSlot = buildBridgeSlot()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	return a, validate(a)
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate480NullBaseline: true, Gate481Cancellation: true, Gate482SourceAbsent: true, AlphaVac: 1, IKVac: 0.5, SectorPerturbationsUnsolved: true, NativeRegistryClean: true, Verdict: StatusInheritedNullPerturbationFrontier}
}

func buildTopologicalAudit() TopologicalAudit {
	cs := []TopologicalCandidate{
		{Name: "SU(3) color representation / color-winding label", NativeObject: true, DiscreteTopological: true, SectorAware: true, QuarkLeptonSeparating: true, GenerationAware: false, CanSourceDeltaAlpha: false, CanSourceDeltaPhi: false, NeedsBridge: false, RejectedFailure: StatusFailedColorWindingGenerationBlind, Reason: "color distinguishes quark representations from colorless leptons, but the same SU(3) color representation is shared by u/c/t and d/s/b; it supplies no generation-indexed cylinder coordinate"},
		{Name: "gauge representation topology and holonomy class", NativeObject: true, DiscreteTopological: true, SectorAware: true, QuarkLeptonSeparating: true, GenerationAware: false, CanSourceDeltaAlpha: false, CanSourceDeltaPhi: false, NeedsBridge: false, RejectedFailure: StatusFailedGaugeRepresentationGenerationBlind, Reason: "representation topology separates gauge sectors, but it is universal across the three families and lacks a native map to delta_alpha_s or delta_phi_s"},
		{Name: "finite holonomy / winding-number stress ansatz", NativeObject: true, DiscreteTopological: true, SectorAware: true, QuarkLeptonSeparating: true, GenerationAware: false, CanSourceDeltaAlpha: false, CanSourceDeltaPhi: false, NeedsBridge: true, RejectedFailure: StatusFailedHolonomyNoDeltaMap, Reason: "a holonomy or winding label can be recorded, but the current theorem atlas contains no coefficient-free morphism from that label to the continuous family-cylinder offsets"},
		{Name: "Betti-number deformation ledger", NativeObject: false, DiscreteTopological: true, SectorAware: false, QuarkLeptonSeparating: false, GenerationAware: false, CanSourceDeltaAlpha: false, CanSourceDeltaPhi: false, NeedsBridge: true, RejectedFailure: StatusFailedBettiLedgerNotPresent, Reason: "no native finite Betti-to-family-coordinate ledger has been constructed inside the current Cℓ(1,7) board"},
		{Name: "one-electron / single light-flow worldline hypothesis", NativeObject: false, DiscreteTopological: false, SectorAware: false, QuarkLeptonSeparating: false, GenerationAware: false, CanSourceDeltaAlpha: false, CanSourceDeltaPhi: false, NeedsBridge: true, RejectedFailure: StatusFailedSingleElectronNotFiniteOperator, Reason: "the idea is a useful heuristic, but Gate483 requires a concrete finite Clifford operator with domain, codomain, trace rule, and generation action; none is present"},
		{Name: "sealed Yukawa/flavor environmental ledger", NativeObject: false, DiscreteTopological: false, SectorAware: true, QuarkLeptonSeparating: true, GenerationAware: true, CanSourceDeltaAlpha: true, CanSourceDeltaPhi: true, NeedsBridge: true, RejectedFailure: StatusFailedYukawaEnvironmental, Reason: "the only object capable of generation-aware sector offsets in the current atlas is the sealed flavor/environmental ledger, exactly the 13-moduli firewall"},
		{Name: "CKM/PMNS residual targets", NativeObject: false, DiscreteTopological: false, SectorAware: true, QuarkLeptonSeparating: true, GenerationAware: true, CanSourceDeltaAlpha: true, CanSourceDeltaPhi: true, NeedsBridge: true, RejectedFailure: StatusFailedCKMPMNSRejected, Reason: "mixing matrices may be residual targets; using them to source winding deformations would reverse the adapter and fit the answer"},
	}
	sectorSeparators, quarkLeptonOnly, fullSource, bridgeLike := false, true, false, false
	for _, c := range cs {
		if c.NativeObject && c.DiscreteTopological && c.SectorAware && c.QuarkLeptonSeparating {
			sectorSeparators = true
		}
		if c.GenerationAware && c.NativeObject && c.DiscreteTopological {
			quarkLeptonOnly = false
		}
		if c.NativeObject && c.DiscreteTopological && c.GenerationAware && c.CanSourceDeltaAlpha && c.CanSourceDeltaPhi && !c.NeedsBridge {
			fullSource = true
		}
		if c.NeedsBridge && c.CanSourceDeltaAlpha && c.CanSourceDeltaPhi {
			bridgeLike = true
		}
	}
	return TopologicalAudit{Executed: true, Candidates: cs, SectorSeparatorsFound: sectorSeparators, QuarkLeptonOnly: quarkLeptonOnly, NativeFullSourceFound: fullSource, BridgeLikeSourcesPresent: bridgeLike, Verdict: StatusTopologicalSectorSeparationAudited, Reason: "native topological/gauge labels can separate quarks from leptons, but the currently available labels are generation-blind and do not form a coefficient-free deformation operator"}
}

func buildGenerationAwareness(t TopologicalAudit) GenerationAwarenessTest {
	failures := []string{StatusFailedColorWindingGenerationBlind, StatusFailedGaugeRepresentationGenerationBlind, StatusFailedNativeTopologicalSourceAbsent}
	return GenerationAwarenessTest{Executed: true, RequiredDistinctions: []string{"u vs c vs t", "d vs s vs b", "e vs mu vs tau", "nu1 vs nu2 vs nu3", "quark vs lepton"}, ColorDistinguishesQuarkLepton: true, ColorDistinguishesGenerations: false, ColorDistinguishesUpCharmTop: false, WindingDistinguishesQuarkLepton: t.SectorSeparatorsFound, WindingDistinguishesGenerations: false, CandidatesPassingGenerationAwareness: 0, Verdict: StatusFailedColorWindingGenerationBlind, Reason: "color/winding-style labels pass the quark-versus-lepton test but fail the essential generation-awareness test: u, c, and t carry the same color representation, as do d, s, and b", Failures: failures}
}

func buildDeformationMap(t TopologicalAudit, g GenerationAwarenessTest) DeformationMap {
	failures := []string{StatusFailedNativeTopologicalSourceAbsent, StatusFailedHolonomyNoDeltaMap, StatusFailedBettiLedgerNotPresent, StatusFailedYukawaEnvironmental, StatusFailedCKMPMNSRejected}
	return DeformationMap{Executed: true, RequiredMap: "topological label w_s -> (delta_alpha_s, delta_phi_s) for each s in {u,d,e,nu} and each generation, independent of CKM/PMNS targets", TopologicalStressNative: t.SectorSeparatorsFound, DeltaAlphaMapNative: false, DeltaPhiMapNative: false, NumericCoordinateMapNative: false, TraceCompatible: true, CKMPMNSIndependent: true, AllZeroDistance: 0, Verdict: StatusFailedNativeTopologicalSourceAbsent, Reason: "Gate483 finds sector labels but no native map from those labels to generation-aware cylinder offsets; without such a map the only rigorous native perturbation remains zero, and physical residuals stay undefined", Failures: failures}
}

func buildBridgeSlot() BridgeSlot {
	return BridgeSlot{Executed: true, SlotName: "topological-sector-perturbation-ledger", AllowedFields: []string{"sector", "generation", "topological_label", "winding_number", "holonomy_class", "delta_alpha", "delta_phi", "I_spec", "I_K", "sigma_CP", "n_C3", "scale", "scheme", "source", "uncertainty", "bridge_only"}, RequiresAirlock: true, RequiresProvenance: true, RequiresBranchTags: true, RequiresUncertainty: true, AllowsTopologicalLabels: true, RejectsCKMPMNSAsInput: true, RejectsNativePromotion: true, CanComputeSyntheticResidual: true, CanComputePhysicalResidual: false, Verdict: StatusBridgeSlotPreserved, Reason: "a future topological perturbation ledger may carry explicit winding/holonomy labels, but it remains bridge-only unless a later theorem constructs a generation-aware native map to delta_alpha and delta_phi"}
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, ObservedMassImported: false, CKMImported: false, PMNSImported: false, VacuumIKNativeBaseline: true, TopologicalNativeSourceFound: false, SectorPerturbationsNative: false, PhysicalDUDComputed: false, PhysicalDENuComputed: false, CKMMatrixConstructed: false, PMNSMatrixConstructed: false, NativeRegistryWritten: false, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "topological labels may classify sectors, but no generation-aware native deformation operator is found; the 13-moduli firewall remains mathematically unbreached"}
}

func buildNext() NextStep {
	return NextStep{Gate: 484, Title: "Generation-aware finite deformation operator construction or closure", Reason: "Gate483 shows color/winding topology separates quarks from leptons but is generation-blind and lacks a native map to family-cylinder offsets.", PrimaryTask: "either construct a new generation-aware finite operator with explicit Cℓ(1,7) action on K/X/Y coordinates, or close the sector perturbation frontier as environmental bridge data"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullBaseline || !a.Inheritance.Gate481Cancellation || !a.Inheritance.Gate482SourceAbsent || a.Inheritance.AlphaVac != 1 || a.Inheritance.IKVac != 0.5 || !a.Inheritance.SectorPerturbationsUnsolved || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate483 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.TopologicalAudit.Executed || !a.TopologicalAudit.SectorSeparatorsFound || !a.TopologicalAudit.QuarkLeptonOnly || a.TopologicalAudit.NativeFullSourceFound || !a.TopologicalAudit.BridgeLikeSourcesPresent || len(a.TopologicalAudit.Candidates) < 6 {
		return fmt.Errorf("Gate483 topological audit invalid: %+v", a.TopologicalAudit)
	}
	if !a.GenerationAwareness.Executed || !a.GenerationAwareness.ColorDistinguishesQuarkLepton || a.GenerationAwareness.ColorDistinguishesGenerations || a.GenerationAwareness.ColorDistinguishesUpCharmTop || !a.GenerationAwareness.WindingDistinguishesQuarkLepton || a.GenerationAwareness.WindingDistinguishesGenerations || a.GenerationAwareness.CandidatesPassingGenerationAwareness != 0 {
		return fmt.Errorf("Gate483 generation awareness invalid: %+v", a.GenerationAwareness)
	}
	if !a.DeformationMap.Executed || !a.DeformationMap.TopologicalStressNative || a.DeformationMap.DeltaAlphaMapNative || a.DeformationMap.DeltaPhiMapNative || a.DeformationMap.NumericCoordinateMapNative || !a.DeformationMap.TraceCompatible || !a.DeformationMap.CKMPMNSIndependent || a.DeformationMap.AllZeroDistance != 0 {
		return fmt.Errorf("Gate483 deformation map invalid: %+v", a.DeformationMap)
	}
	if !a.BridgeSlot.Executed || !a.BridgeSlot.RequiresAirlock || !a.BridgeSlot.RequiresProvenance || !a.BridgeSlot.RequiresBranchTags || !a.BridgeSlot.RequiresUncertainty || !a.BridgeSlot.AllowsTopologicalLabels || !a.BridgeSlot.RejectsCKMPMNSAsInput || !a.BridgeSlot.RejectsNativePromotion || !a.BridgeSlot.CanComputeSyntheticResidual || a.BridgeSlot.CanComputePhysicalResidual {
		return fmt.Errorf("Gate483 bridge slot invalid: %+v", a.BridgeSlot)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassImported || a.Firewall.CKMImported || a.Firewall.PMNSImported || !a.Firewall.VacuumIKNativeBaseline || a.Firewall.TopologicalNativeSourceFound || a.Firewall.SectorPerturbationsNative || a.Firewall.PhysicalDUDComputed || a.Firewall.PhysicalDENuComputed || a.Firewall.CKMMatrixConstructed || a.Firewall.PMNSMatrixConstructed || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate483 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate483 result: audited %d topological deformation candidates; quark/lepton separation is present, but generation-aware native deformation source count is %d, so physical d_ud/d_eν remain undefined and the 13-moduli firewall remains closed.", len(a.TopologicalAudit.Candidates), a.GenerationAwareness.CandidatesPassingGenerationAwareness)
}

func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 483 Registry Audit — Finite Algebraic Deformation Operator Search\n\n")
	b.WriteString("## Verdict\n\n`" + StatusFailedNativeTopologicalSourceAbsent + "`\n\n")
	b.WriteString("Gate 483 tests whether algebraic winding, holonomy, color topology, or a one-light-flow style finite operator can move sectors off the Gate 480 null baseline. The answer is sharply constrained: native topology can separate quark-like from lepton-like sectors, but it does not distinguish generations and does not provide a coefficient-free map to `delta_alpha_s` and `delta_phi_s`.\n\n")
	b.WriteString("## Inherited frontier\n\n```text\n")
	b.WriteString(fmt.Sprintf("alpha_vac = %.12f\nI_K,vac = %.12f\nGate481 shared-baseline cancellation = true\nGate482 native perturbation source absent = true\n", a.Inheritance.AlphaVac, a.Inheritance.IKVac))
	b.WriteString("```\n\n")
	b.WriteString("## Topological candidate audit\n\n| candidate | native | topological | sector-aware | quark/lepton separator | generation-aware | native delta_alpha | native delta_phi | verdict |\n|---|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.TopologicalAudit.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.NativeObject, c.DiscreteTopological, c.SectorAware, c.QuarkLeptonSeparating, c.GenerationAware, c.CanSourceDeltaAlpha && !c.NeedsBridge, c.CanSourceDeltaPhi && !c.NeedsBridge, c.RejectedFailure))
	}
	b.WriteString("\n## Why the topological candidates fail\n\n")
	for _, c := range a.TopologicalAudit.Candidates {
		b.WriteString("- **" + c.Name + "** — " + c.Reason + "\n")
	}
	b.WriteString("\n## Generation-awareness test\n\n```text\n")
	b.WriteString(fmt.Sprintf("color distinguishes quark/lepton = %t\ncolor distinguishes u/c/t = %t\nwinding distinguishes quark/lepton = %t\nwinding distinguishes generations = %t\ncandidates passing generation-awareness = %d\n", a.GenerationAwareness.ColorDistinguishesQuarkLepton, a.GenerationAwareness.ColorDistinguishesUpCharmTop, a.GenerationAwareness.WindingDistinguishesQuarkLepton, a.GenerationAwareness.WindingDistinguishesGenerations, a.GenerationAwareness.CandidatesPassingGenerationAwareness))
	b.WriteString("```\n\n")
	b.WriteString("The generation-awareness test is the decisive obstruction: color charge is real sector topology, but it is shared by all three quark generations. A color/winding label may say `quark`, but it does not say `up versus charm versus top`, nor does it assign a unique family-cylinder coordinate.\n\n")
	b.WriteString("## Deformation map requirement\n\n```text\n")
	b.WriteString(a.DeformationMap.RequiredMap + "\n")
	b.WriteString(fmt.Sprintf("topological stress label native = %t\ndelta_alpha native map = %t\ndelta_phi native map = %t\nnumeric coordinate map native = %t\nall-zero native perturbation distance = %.12f\n", a.DeformationMap.TopologicalStressNative, a.DeformationMap.DeltaAlphaMapNative, a.DeformationMap.DeltaPhiMapNative, a.DeformationMap.NumericCoordinateMapNative, a.DeformationMap.AllZeroDistance))
	b.WriteString("```\n\n")
	b.WriteString("## Bridge slot preserved\n\n`" + a.BridgeSlot.SlotName + "`\n\n```text\n")
	for _, f := range a.BridgeSlot.AllowedFields {
		b.WriteString(f + "\n")
	}
	b.WriteString("```\n\nThis slot may carry explicit winding or holonomy labels in a future bridge run, but it requires airlock provenance, uncertainty, branch tags, and `bridge_only=true`.\n\n")
	b.WriteString("## Rejected routes\n\n```text\n")
	for _, s := range a.DeformationMap.Failures {
		b.WriteString(s + "\n")
	}
	for _, s := range a.GenerationAwareness.Failures {
		b.WriteString(s + "\n")
	}
	b.WriteString(StatusFailedNativePromotion + "\n```\n\n")
	b.WriteString("## Firewall state\n\n```text\n")
	b.WriteString("I_K,vac native baseline = true\n")
	b.WriteString("topological quark/lepton separator = true\n")
	b.WriteString("native generation-aware deformation source = false\n")
	b.WriteString("physical d_ud = undefined\nphysical d_eν = undefined\nCKM/PMNS = not constructed\nnative registry write = false\nnative flavor dimension = 13\ncharged K/X/Y coefficient dimension = 9\n")
	b.WriteString("```\n\n## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s: %s\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	return b.String()
}
