// Package hilbertspacetracecapacity implements Gate 317:
// Hilbert Space Dimension / Trace Capacity Ledger Audit.
//
// Gate 316 reduced the native absolute unified-coupling problem to a very sharp
// target: derive a finite trace capacity C_trace=25, or equivalently the missing
// normalization N4=25/(28*pi), from the completed finite representation rather
// than importing alpha_GUT=1/25 phenomenologically.
//
// Gate 317 therefore performs the literal Hilbert-space count.  It counts the
// Standard-Model-like finite carrier H_F per generation, includes the right
// handed neutrino required by the Majorana/B-gap channel, extends the result to
// the doubled carrier H_F \oplus H_F^*, and audits whether any canonical count
// produces exactly 25.  The result is intentionally strict: the physical carrier
// counts are 16 per generation, 32 when doubled, 48 for three generations, and
// 96 for three doubled generations.  Several 25-shaped coincidences can be
// manufactured only by mixing incompatible categories such as spinor slots,
// adjoint dimensions, contact vacuum dimensions, or empirical target values.
// Those are cataloged and rejected.  Consequently Gate 317 does not derive the
// absolute unified coupling; it localizes the missing theorem more precisely.
package hilbertspacetracecapacity

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE317-HILBERT-SPACE-TRACE-CAPACITY-LEDGER"

	StatusPhysicalStateLedgerFormalized       = "CONDITIONAL_SUPPORT_PHYSICAL_STATE_LEDGER_FORMALIZED"
	StatusDoubledSpaceCapacitySieveFormalized = "CONDITIONAL_SUPPORT_DOUBLED_SPACE_CAPACITY_SIEVE_FORMALIZED"
	StatusTraceTarget25Audited                = "CONDITIONAL_SUPPORT_TRACE_TARGET_25_AUDITED"
	StatusNonCanonical25CoincidencesCataloged = "CONDITIONAL_SUPPORT_25_SHAPED_NONCANONICAL_CANDIDATES_CATALOGED"
	StatusCanonicalCountsCataloged            = "CONDITIONAL_SUPPORT_CANONICAL_HILBERT_COUNTS_CATALOGED"
	StatusGate317FirewallsPreserved           = "CONDITIONAL_SUPPORT_GATE317_FIREWALLS_PRESERVED"

	StatusTensionNoCanonical25HilbertCount    = "CONDITIONAL_TENSION_NO_CANONICAL_HILBERT_SPACE_COUNT_EQUALS_25"
	StatusTensionTraceCapacityNotRawDimension = "CONDITIONAL_TENSION_TRACE_CAPACITY_IS_NOT_A_RAW_HILBERT_DIMENSION"

	StatusFailedNativeTraceCapacity25NotDerived = "FAILED_ROUTE_NATIVE_TRACE_CAPACITY_25_NOT_DERIVED"
	StatusFailedCanonicalCountDoesNotEqual25    = "FAILED_ROUTE_CANONICAL_HILBERT_SPACE_COUNT_DOES_NOT_EQUAL_25"
	StatusFailedMixedCategoryCandidatesRejected = "FAILED_ROUTE_MIXED_CATEGORY_25_CANDIDATES_REJECTED"
	StatusFailedAlphaGUTStillSealed             = "FAILED_ROUTE_ALPHA_GUT_ABSOLUTE_VALUE_STILL_SEALED"
	StatusFailedContinuumPrefactorStillMissing  = "FAILED_ROUTE_CONTINUUM_PREFACTOR_SELECTION_STILL_MISSING"
	StatusFailedHiggsProxyNotUpgraded           = "FAILED_ROUTE_HIGGS_PROXY_NOT_UPGRADED_TO_NATIVE_DERIVATION"
)

const (
	generations      = 3
	targetCapacity25 = 25
)

type StateBlock struct {
	Name                  string
	Chirality             string
	WeakSlots             int
	ColorSlots            int
	GenerationSlots       int
	IncludedInCompletedHF bool
	Comment               string
}

func (b StateBlock) Slots() int { return b.WeakSlots * b.ColorSlots * b.GenerationSlots }

type PhysicalStateLedger struct {
	Blocks                []StateBlock
	LeptonSlotsPerGen     int
	QuarkSlotsPerGen      int
	SlotsPerGeneration    int
	SlotsWithoutNuRPerGen int
	ThreeGenerationSlots  int
	IncludesRightNeutrino bool
	KappaC                int
	KappaQ                int
	Verdict               string
}

type DoubledSpaceCapacity struct {
	ParticleSlotsPerGen       int
	AntiparticleSlotsPerGen   int
	DoubledSlotsPerGen        int
	ParticleSlotsThreeGen     int
	AntiparticleSlotsThreeGen int
	DoubledSlotsThreeGen      int
	DoubledSpaceMandated      bool
	Equals25                  bool
	Verdict                   string
}

type CountCandidate struct {
	Name                  string
	Formula               string
	Value                 int
	CanonicalHilbertCount bool
	Equals25              bool
	CanonicallySelected   bool
	RejectedReason        string
	Status                string
}

type TraceTargetVerification struct {
	TargetCapacity             int
	Candidates                 []CountCandidate
	CanonicalValues            []int
	HasCanonical25             bool
	HasAny25Coincidence        bool
	NativeTraceCapacityDerived bool
	RequiredNextTheorem        string
	Verdict                    string
}

type FirewallAudit struct {
	NoAlphaGUTDerivationClaimed  bool
	NoForced25Selection          bool
	NoMixedCategoryPromotion     bool
	NoContinuumPrefactorInvented bool
	NoHiggsProxyUpgradeClaimed   bool
	FiniteCorePolluted           bool
	Obligations                  []string
	Verdict                      string
}

type Summary struct {
	PhysicalLedgerBuilt    bool
	DoubledSpaceAudited    bool
	CanonicalCountsAudited bool
	Target25Audited        bool
	NativeTrace25Derived   bool
	AlphaGUTDerived        bool
	FirewallsPreserved     bool
	Status                 string
	DirectAnswer           string
	NextGate               string
}

type Analysis struct {
	Physical  PhysicalStateLedger
	Doubled   DoubledSpaceCapacity
	Target    TraceTargetVerification
	Firewalls FirewallAudit
	Summary   Summary
	Truth     string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	physical := buildPhysicalStateLedger()
	doubled := buildDoubledSpaceCapacity(physical)
	target := auditTarget25(physical, doubled)
	firewalls := auditFirewalls(target)
	summary := buildSummary(physical, doubled, target, firewalls)
	truth := "Gate 317 performs the literal finite-Hilbert trace-slot count requested by Gate 316.  The completed one-generation carrier has 16 particle slots including nu_R; Gate-293 doubling gives 32; three generations give 48, or 96 when doubled.  Canonical particle, doubled, and generation-completed counts therefore do not equal 25.  The audit finds 25-shaped coincidences, such as 16+8+1 or 15+7+3, but each mixes incompatible categories and is rejected.  Therefore alpha_GUT^{-1}=25 is not derived as a raw Hilbert-space dimension; the missing theorem must be a weighted trace-capacity / heat-kernel normalization theorem, not a simple degree-of-freedom count."
	return Analysis{Physical: physical, Doubled: doubled, Target: target, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func buildPhysicalStateLedger() PhysicalStateLedger {
	blocks := []StateBlock{
		{Name: "L_L", Chirality: "left", WeakSlots: 2, ColorSlots: 1, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "lepton weak doublet"},
		{Name: "e_R", Chirality: "right", WeakSlots: 1, ColorSlots: 1, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "charged lepton singlet"},
		{Name: "nu_R", Chirality: "right", WeakSlots: 1, ColorSlots: 1, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "right-handed neutrino required by the Majorana/B-gap carrier"},
		{Name: "Q_L", Chirality: "left", WeakSlots: 2, ColorSlots: 3, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "quark weak doublet with color multiplicity kappa_Q=3"},
		{Name: "u_R", Chirality: "right", WeakSlots: 1, ColorSlots: 3, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "up-type quark singlet with color"},
		{Name: "d_R", Chirality: "right", WeakSlots: 1, ColorSlots: 3, GenerationSlots: 1, IncludedInCompletedHF: true, Comment: "down-type quark singlet with color"},
	}
	lepton := blocks[0].Slots() + blocks[1].Slots() + blocks[2].Slots()
	quark := blocks[3].Slots() + blocks[4].Slots() + blocks[5].Slots()
	slots := lepton + quark
	return PhysicalStateLedger{
		Blocks:                blocks,
		LeptonSlotsPerGen:     lepton,
		QuarkSlotsPerGen:      quark,
		SlotsPerGeneration:    slots,
		SlotsWithoutNuRPerGen: slots - blocks[2].Slots(),
		ThreeGenerationSlots:  slots * generations,
		IncludesRightNeutrino: true,
		KappaC:                1,
		KappaQ:                3,
		Verdict:               StatusPhysicalStateLedgerFormalized,
	}
}

func buildDoubledSpaceCapacity(p PhysicalStateLedger) DoubledSpaceCapacity {
	return DoubledSpaceCapacity{
		ParticleSlotsPerGen:       p.SlotsPerGeneration,
		AntiparticleSlotsPerGen:   p.SlotsPerGeneration,
		DoubledSlotsPerGen:        2 * p.SlotsPerGeneration,
		ParticleSlotsThreeGen:     p.ThreeGenerationSlots,
		AntiparticleSlotsThreeGen: p.ThreeGenerationSlots,
		DoubledSlotsThreeGen:      2 * p.ThreeGenerationSlots,
		DoubledSpaceMandated:      true,
		Equals25:                  2*p.SlotsPerGeneration == targetCapacity25 || 2*p.ThreeGenerationSlots == targetCapacity25,
		Verdict:                   strings.Join([]string{StatusDoubledSpaceCapacitySieveFormalized, StatusTensionNoCanonical25HilbertCount}, ";"),
	}
}

func auditTarget25(p PhysicalStateLedger, d DoubledSpaceCapacity) TraceTargetVerification {
	candidates := []CountCandidate{
		{Name: "completed one-generation H_F", Formula: "L_L+e_R+nu_R+Q_L+u_R+d_R", Value: p.SlotsPerGeneration, CanonicalHilbertCount: true, Equals25: p.SlotsPerGeneration == targetCapacity25, CanonicallySelected: true, RejectedReason: "canonical completed particle carrier equals 16, not 25", Status: StatusCanonicalCountsCataloged},
		{Name: "one-generation H_F without nu_R", Formula: "16-1", Value: p.SlotsWithoutNuRPerGen, CanonicalHilbertCount: false, Equals25: p.SlotsWithoutNuRPerGen == targetCapacity25, CanonicallySelected: false, RejectedReason: "excludes the Majorana carrier required by the completed ASHA finite triple", Status: StatusFailedCanonicalCountDoesNotEqual25},
		{Name: "Gate-293 doubled one-generation space", Formula: "H_F plus H_F* = 2*16", Value: d.DoubledSlotsPerGen, CanonicalHilbertCount: true, Equals25: d.DoubledSlotsPerGen == targetCapacity25, CanonicallySelected: true, RejectedReason: "canonical doubled carrier equals 32, not 25", Status: StatusCanonicalCountsCataloged},
		{Name: "three-generation particle carrier", Formula: "3*16", Value: p.ThreeGenerationSlots, CanonicalHilbertCount: true, Equals25: p.ThreeGenerationSlots == targetCapacity25, CanonicallySelected: true, RejectedReason: "generation-completed particle carrier equals 48, not 25", Status: StatusCanonicalCountsCataloged},
		{Name: "three-generation doubled carrier", Formula: "2*3*16", Value: d.DoubledSlotsThreeGen, CanonicalHilbertCount: true, Equals25: d.DoubledSlotsThreeGen == targetCapacity25, CanonicallySelected: true, RejectedReason: "fully doubled generation carrier equals 96, not 25", Status: StatusCanonicalCountsCataloged},
		{Name: "gauge-charged doubled one-generation carrier without nu_R", Formula: "2*15", Value: 2 * p.SlotsWithoutNuRPerGen, CanonicalHilbertCount: false, Equals25: 2*p.SlotsWithoutNuRPerGen == targetCapacity25, CanonicallySelected: false, RejectedReason: "non-completed projection and still equals 30, not 25", Status: StatusFailedCanonicalCountDoesNotEqual25},
		{Name: "spinor plus color adjoint plus abelian singlet", Formula: "16+8+1", Value: 25, CanonicalHilbertCount: false, Equals25: true, CanonicallySelected: false, RejectedReason: "mixes fermion Hilbert slots with gauge-adjoint dimensions; not a single trace carrier count", Status: StatusFailedMixedCategoryCandidatesRejected},
		{Name: "SM-without-nuR plus contact vacuum plus generation count", Formula: "15+7+3", Value: 25, CanonicalHilbertCount: false, Equals25: true, CanonicallySelected: false, RejectedReason: "mixes particle slots, contact cutoff moment, and generation count; category error", Status: StatusFailedMixedCategoryCandidatesRejected},
		{Name: "Gate-316 empirical target", Formula: "alpha_GUT^{-1}", Value: 25, CanonicalHilbertCount: false, Equals25: true, CanonicallySelected: false, RejectedReason: "this is the target reconstructed from phenomenology, not a finite Hilbert count", Status: StatusFailedAlphaGUTStillSealed},
	}

	canonicalValues := make([]int, 0)
	hasCanonical25 := false
	hasAny25 := false
	for _, c := range candidates {
		if c.CanonicalHilbertCount {
			canonicalValues = append(canonicalValues, c.Value)
			if c.Equals25 && c.CanonicallySelected {
				hasCanonical25 = true
			}
		}
		hasAny25 = hasAny25 || c.Equals25
	}
	return TraceTargetVerification{
		TargetCapacity:             targetCapacity25,
		Candidates:                 candidates,
		CanonicalValues:            canonicalValues,
		HasCanonical25:             hasCanonical25,
		HasAny25Coincidence:        hasAny25,
		NativeTraceCapacityDerived: hasCanonical25,
		RequiredNextTheorem:        "derive a weighted trace-capacity / heat-kernel normalization invariant C_trace=25; raw Hilbert-space dimension counting is insufficient",
		Verdict: strings.Join([]string{
			StatusTraceTarget25Audited,
			StatusNonCanonical25CoincidencesCataloged,
			StatusFailedNativeTraceCapacity25NotDerived,
			StatusFailedCanonicalCountDoesNotEqual25,
			StatusFailedMixedCategoryCandidatesRejected,
		}, ";"),
	}
}

func auditFirewalls(t TraceTargetVerification) FirewallAudit {
	obligations := []string{
		StatusFailedNativeTraceCapacity25NotDerived + ": no canonical H_F or H_F⊕H_F* count equals 25",
		StatusFailedMixedCategoryCandidatesRejected + ": 25-shaped sums require category mixing and are not promoted",
		StatusFailedContinuumPrefactorStillMissing + ": alpha_GUT requires a weighted heat-kernel normalization, not only state counting",
		StatusFailedAlphaGUTStillSealed + ": alpha_GUT=1/25 remains empirical until a native capacity theorem is found",
		StatusFailedHiggsProxyNotUpgraded + ": Gate-315 Higgs proxy remains conditional on empirical alpha_GUT",
	}
	return FirewallAudit{
		NoAlphaGUTDerivationClaimed:  !t.NativeTraceCapacityDerived,
		NoForced25Selection:          !t.HasCanonical25,
		NoMixedCategoryPromotion:     true,
		NoContinuumPrefactorInvented: true,
		NoHiggsProxyUpgradeClaimed:   true,
		FiniteCorePolluted:           false,
		Obligations:                  obligations,
		Verdict: strings.Join([]string{
			StatusGate317FirewallsPreserved,
			StatusFailedAlphaGUTStillSealed,
			StatusFailedHiggsProxyNotUpgraded,
		}, ";"),
	}
}

func buildSummary(p PhysicalStateLedger, d DoubledSpaceCapacity, t TraceTargetVerification, f FirewallAudit) Summary {
	native := t.NativeTraceCapacityDerived
	status := StatusFailedNativeTraceCapacity25NotDerived
	if native {
		status = "CONDITIONAL_SUPPORT_NATIVE_TRACE_CAPACITY_25_DERIVED"
	}
	return Summary{
		PhysicalLedgerBuilt:    p.SlotsPerGeneration == 16 && p.ThreeGenerationSlots == 48,
		DoubledSpaceAudited:    d.DoubledSlotsPerGen == 32 && d.DoubledSlotsThreeGen == 96,
		CanonicalCountsAudited: len(t.CanonicalValues) >= 4,
		Target25Audited:        t.TargetCapacity == 25 && t.HasAny25Coincidence,
		NativeTrace25Derived:   native,
		AlphaGUTDerived:        native,
		FirewallsPreserved:     f.NoAlphaGUTDerivationClaimed && f.NoForced25Selection && f.NoMixedCategoryPromotion && f.NoContinuumPrefactorInvented && f.NoHiggsProxyUpgradeClaimed && !f.FiniteCorePolluted,
		Status:                 status,
		DirectAnswer:           "No canonical Hilbert-space dimension count yields C_trace=25: completed H_F=16, doubled H_F⊕H_F*=32, three-generation H_F=48, and three-generation doubled space=96.  Therefore alpha_GUT^{-1}=25 is not a raw slot count; the missing Phase-II theorem must be a weighted trace-capacity or continuum-normalization invariant.",
		NextGate:               "Search for a weighted heat-kernel trace-capacity invariant rather than a raw Hilbert dimension, or return to the threshold portal tensor with alpha_GUT quarantined.",
	}
}

func FormatPhysicalLedger(x PhysicalStateLedger) string {
	parts := make([]string, 0, len(x.Blocks))
	for _, b := range x.Blocks {
		parts = append(parts, fmt.Sprintf("%s=%d", b.Name, b.Slots()))
	}
	return fmt.Sprintf("blocks=[%s]; lepton=%d; quark=%d; per_gen=%d; without_nuR=%d; three_gen=%d; kappaC=%d; kappaQ=%d; verdict=%s", strings.Join(parts, ","), x.LeptonSlotsPerGen, x.QuarkSlotsPerGen, x.SlotsPerGeneration, x.SlotsWithoutNuRPerGen, x.ThreeGenerationSlots, x.KappaC, x.KappaQ, x.Verdict)
}

func FormatDoubled(x DoubledSpaceCapacity) string {
	return fmt.Sprintf("particle_per_gen=%d; antiparticle_per_gen=%d; doubled_per_gen=%d; particle_3gen=%d; anti_3gen=%d; doubled_3gen=%d; mandated=%t; equals25=%t; verdict=%s", x.ParticleSlotsPerGen, x.AntiparticleSlotsPerGen, x.DoubledSlotsPerGen, x.ParticleSlotsThreeGen, x.AntiparticleSlotsThreeGen, x.DoubledSlotsThreeGen, x.DoubledSpaceMandated, x.Equals25, x.Verdict)
}

func FormatTarget(x TraceTargetVerification) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s:%s=%d canonical=%t selected=%t equals25=%t", c.Name, c.Formula, c.Value, c.CanonicalHilbertCount, c.CanonicallySelected, c.Equals25))
	}
	return fmt.Sprintf("target=%d; canonical_values=%v; has_canonical25=%t; has_any25=%t; native=%t; candidates=[%s]; next=%s; verdict=%s", x.TargetCapacity, x.CanonicalValues, x.HasCanonical25, x.HasAny25Coincidence, x.NativeTraceCapacityDerived, strings.Join(parts, "; "), x.RequiredNextTheorem, x.Verdict)
}

func FormatFirewalls(x FirewallAudit) string {
	return fmt.Sprintf("no_alpha_claim=%t; no_forced25=%t; no_mixed_promotion=%t; no_prefactor_invented=%t; no_higgs_upgrade=%t; polluted=%t; obligations=[%s]; verdict=%s", x.NoAlphaGUTDerivationClaimed, x.NoForced25Selection, x.NoMixedCategoryPromotion, x.NoContinuumPrefactorInvented, x.NoHiggsProxyUpgradeClaimed, x.FiniteCorePolluted, strings.Join(x.Obligations, "; "), x.Verdict)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("physical=%t; doubled=%t; canonical_counts=%t; target25=%t; native25=%t; alpha=%t; firewalls=%t; status=%s; answer=%s; next=%s", x.PhysicalLedgerBuilt, x.DoubledSpaceAudited, x.CanonicalCountsAudited, x.Target25Audited, x.NativeTrace25Derived, x.AlphaGUTDerived, x.FirewallsPreserved, x.Status, x.DirectAnswer, x.NextGate)
}
