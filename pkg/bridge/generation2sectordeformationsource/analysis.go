// Package generation2sectordeformationsource implements Gate 482:
// Null-Baseline Sector Deformation Source Search.
package generation2sectordeformationsource

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID                                 = "GATE482-NULL-BASELINE-SECTOR-DEFORMATION-SOURCE-SEARCH"
	StatusAuditCompleted                    = "CONDITIONAL_SUPPORT_SECTOR_DEFORMATION_SOURCE_AUDIT_COMPLETED"
	StatusBaselineInherited                 = "CONDITIONAL_SUPPORT_GATE481_NULL_BASELINE_CANCELLATION_INHERITED"
	StatusCandidateLedgerAudited            = "CONDITIONAL_SUPPORT_NATIVE_DEFORMATION_CANDIDATE_LEDGER_AUDITED"
	StatusEmpiricalSlotPreserved            = "CONDITIONAL_SUPPORT_BRIDGE_ONLY_SECTOR_PERTURBATION_SLOT_PRESERVED"
	StatusFirewallPreserved                 = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED_WITH_GATE482_SOURCE_SEARCH"
	StatusFailedNativeSourceAbsent          = "FAILED_ROUTE_NATIVE_SECTOR_PERTURBATION_SOURCE_ABSENT"
	StatusFailedOrientationGenerationOnly   = "FAILED_ROUTE_FINITE_ORIENTATION_FIXES_GENERATION_ADDRESS_NOT_SECTOR_PERTURBATIONS"
	StatusFailedChiralityGenerationBlind    = "FAILED_ROUTE_CHIRALITY_REAL_STRUCTURE_GENERATION_BLIND"
	StatusFailedHiggsEdgeScaleOnly          = "FAILED_ROUTE_HIGGS_EDGE_OPERATOR_SCALE_NORMALIZES_BUT_DOES_NOT_SELECT_FAMILY_RAY"
	StatusFailedGaugeChargesGenerationBlind = "FAILED_ROUTE_ELECTROWEAK_CHARGES_DISTINGUISH_SECTORS_BUT_ARE_GENERATION_UNIVERSAL"
	StatusFailedYukawaSealed                = "FAILED_ROUTE_YUKAWA_FLAVOR_LEDGER_IS_SEALED_ENVIRONMENTAL_DATA"
	StatusFailedPMNSCKMRejected             = "FAILED_ROUTE_CKM_PMNS_AS_DEFORMATION_SOURCE_REJECTED"
	StatusFailedNativePromotion             = "FAILED_ROUTE_SECTOR_DEFORMATION_NATIVE_PROMOTION_REJECTED"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed, Gate480NullBaseline, Gate481BaselineCancellation         bool
	AlphaVac, IKVac                                                    float64
	PhysicalDUDUnresolved, PhysicalDENuUnresolved, NativeRegistryClean bool
	Verdict                                                            string
}
type Candidate struct {
	Name                                                                                                                   string
	NativeObject, Universal, GenerationAware, SectorAware, CanSourceDeltaAlpha, CanSourceDeltaPhi, RequiresEmpiricalLedger bool
	RejectedFailure, Reason                                                                                                string
}
type CandidateAudit struct {
	Executed                                bool
	Candidates                              []Candidate
	NativeSourceFound, BridgeSourcesPresent bool
	Verdict, Reason                         string
}
type SourceSieve struct {
	Executed                                              bool
	RequiredOutput                                        string
	RequiredSectors, RequiredCoordinatesPerSector         []string
	NativeCandidatesTested, CandidatesPassingNativeSource int
	DeltaAlphaNative, DeltaPhiNative                      bool
	AllZeroPerturbationDistance                           float64
	AllZeroWouldPredictNoMixing                           bool
	Verdict, Reason                                       string
	Failures                                              []string
}
type BridgeSlot struct {
	Executed                                                                                                                                                                             bool
	SlotName                                                                                                                                                                             string
	AllowedFields                                                                                                                                                                        []string
	RequiresAirlock, RequiresProvenance, RequiresBranchTags, RequiresUncertainty, RejectsCKMPMNSAsInput, RejectsNativePromotion, CanComputeSyntheticResidual, CanComputePhysicalResidual bool
	Verdict, Reason                                                                                                                                                                      string
}
type Firewall struct {
	Executed                                                                                                                                                                                                                                     bool
	ObservedMassImported, CKMImported, PMNSImported, VacuumIKNativeBaseline, SectorPerturbationsNative, SectorPerturbationsSolved, PhysicalDUDComputed, PhysicalDENuComputed, CKMMatrixConstructed, PMNSMatrixConstructed, NativeRegistryWritten bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                                                                                       int
	Verdict, Reason                                                                                                                                                                                                                              string
}
type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}
type Analysis struct {
	Inheritance Inheritance
	Candidates  CandidateAudit
	Sieve       SourceSieve
	BridgeSlot  BridgeSlot
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
	a.Candidates = buildCandidates()
	a.Sieve = buildSieve(a.Candidates)
	a.BridgeSlot = buildBridgeSlot()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	return a, validate(a)
}
func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate480NullBaseline: true, Gate481BaselineCancellation: true, AlphaVac: 1, IKVac: 0.5, PhysicalDUDUnresolved: true, PhysicalDENuUnresolved: true, NativeRegistryClean: true, Verdict: StatusBaselineInherited}
}
func buildCandidates() CandidateAudit {
	cs := []Candidate{{"finite orientation / triality family address", true, true, true, false, false, false, false, StatusFailedOrientationGenerationOnly, "the finite orientation and triality-style family address can organize the three-generation board, but it does not distinguish u, d, e, and nu sector perturbation rays"}, {"chirality grading gamma_F and real structure J", true, true, false, false, false, false, false, StatusFailedChiralityGenerationBlind, "gamma_F and J enforce chirality/reality/Hermiticity structure; they are generation-blind and supply no K-axis overlap perturbation per sector"}, {"Higgs one-form edge / electroweak VEV socket", true, true, false, true, false, false, false, StatusFailedHiggsEdgeScaleOnly, "the Higgs-edge lane can normalize electroweak scale and chiral edges, but it does not select the family-cylinder coordinates delta_alpha_s or delta_phi_s"}, {"electroweak gauge charges and W/Z couplings", true, true, false, true, false, false, false, StatusFailedGaugeChargesGenerationBlind, "gauge charges distinguish representation sectors, but they are generation-universal and cannot rank-complete the family perturbation ray"}, {"Yukawa/flavor coefficient ledger", false, false, true, true, true, true, true, StatusFailedYukawaSealed, "the sealed Yukawa/flavor ledger is precisely where sector perturbations may live, but it is environmental/bridge data, not native law-space"}, {"CKM/PMNS residual targets", false, false, true, true, true, true, true, StatusFailedPMNSCKMRejected, "CKM/PMNS may only be residual targets; using them as perturbation sources would invert the comparator and contaminate the theorem registry"}}
	found, bridge := false, false
	for _, c := range cs {
		if c.NativeObject && c.CanSourceDeltaAlpha && c.CanSourceDeltaPhi && !c.RequiresEmpiricalLedger {
			found = true
		}
		if c.RequiresEmpiricalLedger && c.CanSourceDeltaAlpha && c.CanSourceDeltaPhi {
			bridge = true
		}
	}
	return CandidateAudit{true, cs, found, bridge, StatusCandidateLedgerAudited, "native candidates were audited as operators on the sector-perturbation slot; none supplies both delta_alpha and delta_phi without entering the bridge/environmental ledger"}
}
func buildSieve(c CandidateAudit) SourceSieve {
	fs := []string{StatusFailedNativeSourceAbsent}
	for _, x := range c.Candidates {
		fs = append(fs, x.RejectedFailure)
	}
	return SourceSieve{Executed: true, RequiredOutput: "for each sector s in {u,d,e,nu}: delta_alpha_s and delta_phi_s", RequiredSectors: []string{"u", "d", "e", "nu"}, RequiredCoordinatesPerSector: []string{"delta_alpha", "delta_phi"}, NativeCandidatesTested: len(c.Candidates), CandidatesPassingNativeSource: 0, DeltaAlphaNative: false, DeltaPhiNative: false, AllZeroPerturbationDistance: 0, AllZeroWouldPredictNoMixing: true, Verdict: StatusFailedNativeSourceAbsent, Reason: "if no sector perturbation source is supplied, all sectors remain on the shared null baseline and all relative cylinder distances vanish; nonzero CKM/PMNS residuals therefore require sector perturbation data or a later native source theorem", Failures: fs}
}
func buildBridgeSlot() BridgeSlot {
	return BridgeSlot{Executed: true, SlotName: "sector-perturbation-source-ledger", AllowedFields: []string{"sector", "delta_alpha", "delta_phi", "I_spec", "I_K", "sigma_CP", "n_C3", "scale", "scheme", "source", "uncertainty", "bridge_only"}, RequiresAirlock: true, RequiresProvenance: true, RequiresBranchTags: true, RequiresUncertainty: true, RejectsCKMPMNSAsInput: true, RejectsNativePromotion: true, CanComputeSyntheticResidual: true, CanComputePhysicalResidual: false, Verdict: StatusEmpiricalSlotPreserved, Reason: "Gate482 preserves a legal bridge slot for future sector perturbations while forbidding CKM/PMNS targets, observed masses, or fitted rays from becoming native source theorems"}
}
func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, ObservedMassImported: false, CKMImported: false, PMNSImported: false, VacuumIKNativeBaseline: true, SectorPerturbationsNative: false, SectorPerturbationsSolved: false, PhysicalDUDComputed: false, PhysicalDENuComputed: false, CKMMatrixConstructed: false, PMNSMatrixConstructed: false, NativeRegistryWritten: false, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate482 finds no native sector-deformation source; the null baseline remains native, but all sector perturbations and physical mixing residuals remain bridge/firewalled"}
}
func buildNext() NextStep {
	return NextStep{Gate: 483, Title: "Finite deformation-source theorem search", Reason: "Gate482 found no existing native operator that sources sector perturbations after null-baseline cancellation.", PrimaryTask: "construct or rule out a new finite algebraic deformation operator that is sector-aware, generation-aware, trace-compatible, and independent of observed CKM/PMNS data"}
}
func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate480NullBaseline || !a.Inheritance.Gate481BaselineCancellation || a.Inheritance.AlphaVac != 1 || a.Inheritance.IKVac != 0.5 || !a.Inheritance.PhysicalDUDUnresolved || !a.Inheritance.PhysicalDENuUnresolved || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate482 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Candidates.Executed || a.Candidates.NativeSourceFound || !a.Candidates.BridgeSourcesPresent || len(a.Candidates.Candidates) < 5 {
		return fmt.Errorf("Gate482 candidate audit invalid: %+v", a.Candidates)
	}
	if !a.Sieve.Executed || a.Sieve.CandidatesPassingNativeSource != 0 || a.Sieve.DeltaAlphaNative || a.Sieve.DeltaPhiNative || a.Sieve.AllZeroPerturbationDistance != 0 || !a.Sieve.AllZeroWouldPredictNoMixing {
		return fmt.Errorf("Gate482 source sieve invalid: %+v", a.Sieve)
	}
	if !a.BridgeSlot.Executed || !a.BridgeSlot.RequiresAirlock || !a.BridgeSlot.RequiresProvenance || !a.BridgeSlot.RequiresBranchTags || !a.BridgeSlot.RequiresUncertainty || !a.BridgeSlot.RejectsCKMPMNSAsInput || !a.BridgeSlot.RejectsNativePromotion || !a.BridgeSlot.CanComputeSyntheticResidual || a.BridgeSlot.CanComputePhysicalResidual {
		return fmt.Errorf("Gate482 bridge slot invalid: %+v", a.BridgeSlot)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedMassImported || a.Firewall.CKMImported || a.Firewall.PMNSImported || !a.Firewall.VacuumIKNativeBaseline || a.Firewall.SectorPerturbationsNative || a.Firewall.SectorPerturbationsSolved || a.Firewall.PhysicalDUDComputed || a.Firewall.PhysicalDENuComputed || a.Firewall.CKMMatrixConstructed || a.Firewall.PMNSMatrixConstructed || a.Firewall.NativeRegistryWritten || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("Gate482 firewall invalid: %+v", a.Firewall)
	}
	return nil
}
func truth(a Analysis) string {
	return fmt.Sprintf("Gate482 result: audited %d deformation candidates and found no native source for sector perturbations; with all perturbations zero the relative distance is %.12g, so CKM/PMNS remain undefined until a bridge ledger or new native source theorem appears.", a.Sieve.NativeCandidatesTested, a.Sieve.AllZeroPerturbationDistance)
}
func fmtFloat(x float64) string {
	if math.IsNaN(x) {
		return "undefined"
	}
	return fmt.Sprintf("%.12g", x)
}
func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 482 Registry Audit — Null-Baseline Sector Deformation Source Search\n\n## Verdict\n\n`" + StatusFailedNativeSourceAbsent + "`\n\nGate 482 inherits Gate 480/481: the null baseline gives `alpha_vac=1`, `I_K,vac=1/2`, but common baseline terms cancel from relative distances. The remaining question is whether native finite geometry already supplies the sector perturbations `delta_alpha_s` and `delta_phi_s`. The answer is no.\n\n## Required source object\n\n```text\nfor each s in {u,d,e,nu}:\n  delta_alpha_s = native sector deformation along K-axis\n  delta_phi_s   = native sector deformation around X/Y phase circle\nrelative distances depend only on differences of these perturbations\n```\n\nIf all sector perturbations are zero, all sectors sit on the same null baseline and the relative cylinder distance is exactly zero. This does not match nonzero CKM/PMNS residual targets, but it also does not license a fit.\n\n## Candidate source audit\n\n| candidate | native | generation-aware | sector-aware | sources delta_alpha | sources delta_phi | verdict |\n|---|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Candidates.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.NativeObject, c.GenerationAware, c.SectorAware, c.CanSourceDeltaAlpha && !c.RequiresEmpiricalLedger, c.CanSourceDeltaPhi && !c.RequiresEmpiricalLedger, c.RejectedFailure))
	}
	b.WriteString("\n## Why the candidates fail\n\n")
	for _, c := range a.Candidates.Candidates {
		b.WriteString("- **" + c.Name + "** — " + c.Reason + "\n")
	}
	b.WriteString("\n## Bridge slot preserved\n\nGate 482 preserves a legal future slot, but it is explicitly bridge-only:\n\n`" + a.BridgeSlot.SlotName + "`\n\n```text\n")
	for _, f := range a.BridgeSlot.AllowedFields {
		b.WriteString(f + "\n")
	}
	b.WriteString("```\n\nThe slot requires airlock provenance, uncertainty, branch tags, and `bridge_only=true`. CKM/PMNS targets cannot be used as deformation sources.\n\n## Firewall state\n\n```text\nI_K,vac native baseline = true\nsector perturbations native = false\nphysical d_ud = undefined\nphysical d_eν = undefined\nCKM/PMNS matrix export = rejected\nnative registry write = false\nnative flavor dimension = 13\ncharged K/X/Y coefficient dimension = 9\n```\n\n## Rejected routes\n\n```text\n")
	for _, s := range a.Sieve.Failures {
		b.WriteString(s + "\n")
	}
	b.WriteString(StatusFailedNativePromotion + "\n```\n\n## Numerical output\n\n```text\n")
	b.WriteString(fmt.Sprintf("alpha_vac = %.12f\nI_K,vac = %.12f\nall-zero perturbation distance = %.12f\n", a.Inheritance.AlphaVac, a.Inheritance.IKVac, a.Sieve.AllZeroPerturbationDistance))
	b.WriteString("physical d_ud = undefined\nphysical d_eν = undefined\nCKM/PMNS = not constructed\n```\n\n## Next step\n\n")
	b.WriteString(fmt.Sprintf("Gate %d — %s: %s\n", a.Next.Gate, a.Next.Title, a.Next.PrimaryTask))
	return b.String()
}
