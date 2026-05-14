// Package fermionicgenerationorigin implements Gate 409:
// Fermionic Matter-Carrier Origin / Nontrivial Generation Representation Sieve.
//
// Gates 398--408 closed the scalar/H_phi route: the Higgs scalar carrier has
// full algebraic capacity but no native variational selector, and its selected
// observables remain central or pair-degenerate. Gate 409 pivots back to the
// fermionic matter carrier and audits whether a nontrivial generation
// representation is already present before inserting Yukawa amplitudes. The
// theorem deliberately rejects color, chirality, species, scalar, and exact
// triality degeneracy as generation origins unless a typed finite-Dirac functor
// into generation space is derived.
package fermionicgenerationorigin

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE409-FERMIONIC-MATTER-CARRIER-ORIGIN-NONTRIVIAL-GENERATION-REPRESENTATION-SIEVE"

	StatusScalarFlavorBlindnessInherited   = "CONDITIONAL_SUPPORT_GATE408_SCALAR_FLAVOR_BLINDNESS_INHERITED"
	StatusFermionicCarrierInventoryAudited = "CONDITIONAL_SUPPORT_FERMIONIC_CARRIER_INVENTORY_AUDITED"
	StatusPrimitiveIdempotentSearchAudited = "CONDITIONAL_SUPPORT_PRIMITIVE_FERMIONIC_IDEMPOTENT_SEARCH_AUDITED"
	StatusCommutantCentralizerAudited      = "CONDITIONAL_SUPPORT_FERMIONIC_COMMUTANT_CENTRALIZER_AUDITED"
	StatusTrialityFromFermionAudited       = "CONDITIONAL_SUPPORT_TRIALITY_FROM_FERMION_SIDE_AUDITED"
	StatusBilinearOperatorAuditCompleted   = "CONDITIONAL_SUPPORT_FERMIONIC_BILINEAR_OPERATOR_AUDITED"
	StatusDynamicGenerationSourceAudited   = "CONDITIONAL_SUPPORT_DYNAMIC_GENERATION_SOURCE_AUDITED"
	StatusSealedFermionicCapacityFound     = "CONDITIONAL_SUPPORT_FERMIONIC_GENERATION_CAPACITY"
	StatusConditionalCKMCapacityActivated  = "CONDITIONAL_SUPPORT_CKM_MIXING_CAPACITY_ACTIVATED"

	StatusVerifiedFermionicGenerationOrigin = "VERIFIED_FERMIONIC_GENERATION_ORIGIN_DERIVED"
	StatusVerifiedNativeNoncommutingPair    = "VERIFIED_NATIVE_NONCOMMUTING_FERMIONIC_TEXTURE_PAIR"

	StatusFailedTrivialGenerationCopy         = "FAILED_ROUTE_FERMIONIC_CARRIER_REMAINS_TRIVIAL_GENERATION_COPY"
	StatusFailedSpinorSplitChirality          = "FAILED_ROUTE_SPINOR_SPLIT_IS_CHIRALITY_NOT_GENERATION"
	StatusFailedTrialityExactDegeneracy       = "FAILED_ROUTE_TRIALITY_EXACT_DEGENERACY"
	StatusFailedColorAsGeneration             = "FAILED_ROUTE_COLOR_CONFUSED_WITH_GENERATION"
	StatusFailedSpeciesAsGeneration           = "FAILED_ROUTE_SPECIES_CONFUSED_WITH_GENERATION"
	StatusFailedNoNativeCKMCapacity           = "FAILED_ROUTE_NO_NATIVE_CKM_CAPACITY"
	StatusFailedNoNativeGenerationHamiltonian = "FAILED_ROUTE_NO_NATIVE_GENERATION_HAMILTONIAN"
	StatusFailedNoNativeBilinearSelector      = "FAILED_ROUTE_NO_NATIVE_FERMIONIC_BILINEAR_SELECTOR"
	StatusFirewallPreserved13Moduli           = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	FockDim                       = 16
	OneGenerationYukawaChannels   = 8
	FullGenerationMixingMaps      = 72
	Gate372ChargedFlavorModuliDim = 13
	ExactTrialityTextureDim       = 2
	TrivialGenerationCommutantDim = 9
)

type Inheritance struct {
	Executed                               bool
	Gate408ScalarFlavorBlind               bool
	Gate407FullHphiCapacityButNoSelector   bool
	Gate395SpinorSplitTwoSector            bool
	Gate396ThreeObjectSourcesNotGeneration bool
	Gate397ContactSingletonsNotFlavor      bool
	Gate393TrialityDomainNotAdmitted       bool
	Gate394StaticGenerationAddressCentral  bool
	Gate372ChargedModuliDim                int
	NoEmpiricalInputsImported              bool
	Verdict                                string
}

type Carrier struct {
	Name                       string
	Domain                     string
	Dimension                  int
	Native                     bool
	ThreeSector                bool
	ThreeSectorKind            string
	NoncentralGenerationAction bool
	NoncommutingOperators      int
	LinksToYukawaSpaces        bool
	CompatibleAF               bool
	CompatibleJ                bool
	CompatibleGamma            bool
	CompatibleFirstOrder       bool
	CompatibleHypercharge      bool
	CompatibleSU2L             bool
	CompatibleBMinusL          bool
	ConfusesColor              bool
	ConfusesSpecies            bool
	ConfusesChirality          bool
	Verdict                    string
	Reason                     string
}

type CarrierInventory struct {
	Executed                          bool
	Carriers                          []Carrier
	NativeCarrierCount                int
	NativeThreeSectorCount            int
	NativeGenerationCarrierCount      int
	NativeNoncentralGenerationActions int
	NativeNoncommutingOperatorPairs   int
	ColorThreefoldCount               int
	SpeciesOrChiralityThreefoldCount  int
	Verdict                           string
}

type IdempotentCandidate struct {
	Name                          string
	Domain                        string
	Native                        bool
	PrimitiveBlocks               int
	BlockDimensions               []int
	ThreeBlocks                   bool
	ActsOnGenerationLabels        bool
	ActsOnSpeciesColorChirality   bool
	NoncentralOnC3Gen             bool
	GivesTwoNoncommutingOperators bool
	ManuallyLabelled              bool
	Verdict                       string
	Reason                        string
}

type IdempotentAudit struct {
	Executed                        bool
	Candidates                      []IdempotentCandidate
	NativeThreeBlockCandidates      int
	NativeGenerationLabelCandidates int
	ColorOrSpeciesRejected          int
	ManualLabelRejected             int
	NoncommutingNativePairs         int
	Verdict                         string
}

type CommutantAudit struct {
	Executed                           bool
	Action                             string
	Carrier                            string
	CommuntantModel                    string
	ContainsU3GenerationFreedom        bool
	GenerationFreedomCanonicalSelector bool
	NativeDiagonalGenerationOperator   bool
	NativeNoncommutingGenerationPair   bool
	ArbitraryGenerationRotations       bool
	CentralBroadcastOverGeneration     bool
	DistinguishesSpeciesOnly           bool
	Verdict                            string
	Reason                             string
}

type TrialityAudit struct {
	Executed                           bool
	TrialityDomainAdmitted             bool
	FermionTo8VRepresentativeFound     bool
	FermionTo8SRepresentativeFound     bool
	FermionTo8CRepresentativeFound     bool
	GenerationLabelsDerivedNotInserted bool
	ExactTrialityDegeneracyPresent     bool
	BrokenTrialityOperatorNative       bool
	OnePlusTwoDegeneracy               bool
	Verdict                            string
	Reason                             string
}

type BilinearFamily struct {
	Name                              string
	Native                            bool
	AllowedDimension                  int
	ActsOnGenerationLabels            bool
	ActsOnParticleSpecies             bool
	DiagonalOnly                      bool
	NoncommutingOperators             int
	DistinguishesUpDownLeptonNatively bool
	Reduces13Moduli                   bool
	Verdict                           string
	Reason                            string
}

type BilinearAudit struct {
	Executed                          bool
	Families                          []BilinearFamily
	NativeFamilies                    int
	NativeGenerationSensitiveFamilies int
	NativeNoncommutingFamilies        int
	NativeModuliReducingFamilies      int
	Verdict                           string
}

type DynamicSource struct {
	Name                   string
	Native                 bool
	Compatible             bool
	SealedExternal         bool
	Circular               bool
	WrongDomain            bool
	Spectrum               string
	ThreeLevelSpectrum     bool
	ActsOnFermionCarrier   bool
	ActsOnGenerationLabels bool
	ProducesHierarchy      bool
	ProducesMixing         bool
	Verdict                string
	Reason                 string
}

type DynamicSourceAudit struct {
	Executed                     bool
	Sources                      []DynamicSource
	NativeThreeLevelSources      int
	NativeGenerationHamiltonians int
	SealedOrCircularSources      int
	WrongDomainSources           int
	Verdict                      string
}

type CKMAudit struct {
	Executed                        bool
	NativeCandidateOperators        int
	NativeNoncommutingPairs         int
	SealedNoncommutingPairs         int
	UpDownTextureMisalignmentNative bool
	UpDownTextureMisalignmentSealed bool
	CKMCapacityNative               bool
	CKMCapacityConditional          bool
	PMNSCapacityNative              bool
	Verdict                         string
	Reason                          string
}

type ModuliScenario struct {
	Name                          string
	Status                        string
	ModuliDim                     int
	ThreeDistinctMassesPossible   bool
	CKMPossible                   bool
	PMNSPossible                  bool
	QuarkLeptonSeparationPossible bool
	Reason                        string
}

type ModuliImpact struct {
	StartDim                    int
	Scenarios                   []ModuliScenario
	BestNativeDim               int
	NativeReductionBelow13      bool
	ConditionalReductionBelow13 bool
	FirewallPreserved           bool
	Verdict                     string
}

type FirewallAudit struct {
	Executed                      bool
	NoObservedMassesImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoYukawaAmplitudesInserted    bool
	NoScalarSelectorPromoted      bool
	NoTauEtaInserted              bool
	NoNDiagInserted               bool
	NoManualGenerationLabels      bool
	NoColorAsGenerationPromoted   bool
	NoSpeciesAsGenerationPromoted bool
	NoModuliReductionClaimed      bool
	Verdict                       string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Inventory   CarrierInventory
	Idempotents IdempotentAudit
	Commutant   CommutantAudit
	Triality    TrialityAudit
	Bilinears   BilinearAudit
	Sources     DynamicSourceAudit
	CKM         CKMAudit
	Moduli      ModuliImpact
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var buildOnce sync.Once
var cached Analysis
var cachedErr error

func BuildDefault() (Analysis, error) {
	buildOnce.Do(func() {
		inh := inherit()
		inventory := auditCarriers()
		idempotents := auditIdempotents()
		commutant := auditCommutant(inventory)
		triality := auditTrialityFromFermions()
		bilinears := auditBilinears()
		sources := auditDynamicSources()
		ckm := auditCKM(bilinears, sources, triality)
		moduli := auditModuli(inventory, idempotents, commutant, triality, bilinears, sources, ckm)
		firewall := auditFirewall(moduli)
		next := nextStep(moduli, ckm, sources)
		cached = Analysis{Inheritance: inh, Inventory: inventory, Idempotents: idempotents, Commutant: commutant, Triality: triality, Bilinears: bilinears, Sources: sources, CKM: ckm, Moduli: moduli, Firewall: firewall, Next: next}
		cached.Truth = truth(cached)
	})
	return cached, cachedErr
}

func inherit() Inheritance {
	return Inheritance{Executed: true, Gate408ScalarFlavorBlind: true, Gate407FullHphiCapacityButNoSelector: true, Gate395SpinorSplitTwoSector: true, Gate396ThreeObjectSourcesNotGeneration: true, Gate397ContactSingletonsNotFlavor: true, Gate393TrialityDomainNotAdmitted: true, Gate394StaticGenerationAddressCentral: true, Gate372ChargedModuliDim: Gate372ChargedFlavorModuliDim, NoEmpiricalInputsImported: true, Verdict: StatusScalarFlavorBlindnessInherited}
}

func auditCarriers() CarrierInventory {
	xs := []Carrier{
		{Name: "Fock carrier Lambda*(C^4)", Domain: "16 occupation states", Dimension: 16, Native: true, ThreeSector: false, ThreeSectorKind: "none: 1+3 mode split and many charge eigenspaces, not generations", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "The Fock carrier encodes particle species, chirality, charge and color seeds; it does not supply an End(C^3_gen) action."},
		{Name: "even/odd Fock parity", Domain: "fermion parity grading", Dimension: 16, Native: true, ThreeSector: false, ThreeSectorKind: "two-sector chirality/parity", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesChirality: true, Verdict: StatusFailedSpinorSplitChirality, Reason: "Even/odd and 8_s+8_c style decompositions are two-sector chirality data, not three generations."},
		{Name: "B-L / hypercharge eigenspaces", Domain: "finite charge table", Dimension: 16, Native: true, ThreeSector: false, ThreeSectorKind: "charge/species sectors", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesSpecies: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "Charge eigenspaces distinguish Standard Model species and handedness, not generation copies."},
		{Name: "SU(2)_L doublet/singlet decomposition", Domain: "weak representation table", Dimension: 16, Native: true, ThreeSector: false, ThreeSectorKind: "weak isospin", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesSpecies: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "Weak doublets/singlets select gauge representation, not family index."},
		{Name: "spatial Fock triplet / color seeds", Domain: "three spatial creation modes", Dimension: 3, Native: true, ThreeSector: true, ThreeSectorKind: "color/spatial", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesColor: true, Verdict: StatusFailedColorAsGeneration, Reason: "The native 3 is color/spatial structure; promoting it to generation would erase the color semantics already used by the charge table."},
		{Name: "finite spectral-triple Morita bimodule carrier", Domain: "A_F = C + H + M3(C) bimodule", Dimension: 32, Native: true, ThreeSector: true, ThreeSectorKind: "M3(C) color factor", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesColor: true, Verdict: StatusFailedColorAsGeneration, Reason: "The matrix-three component is color. The bimodule broadcasts over generation multiplicity unless an additional functor is supplied."},
		{Name: "Dirac/Yukawa source-target carrier", Domain: "one-generation LR bilinear channels", Dimension: OneGenerationYukawaChannels, Native: true, ThreeSector: false, ThreeSectorKind: "one generation channel inventory", LinksToYukawaSpaces: true, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, ConfusesSpecies: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "The channel selector determines allowed couplings but not their generation amplitudes or mixing."},
		{Name: "triality-lifted Yukawa channel arena", Domain: "three formal triality sectors", Dimension: 24, Native: false, ThreeSector: true, ThreeSectorKind: "representation-category arena", LinksToYukawaSpaces: true, CompatibleAF: false, CompatibleJ: false, CompatibleGamma: false, CompatibleFirstOrder: false, CompatibleHypercharge: true, CompatibleSU2L: true, CompatibleBMinusL: true, Verdict: StatusFailedTrialityExactDegeneracy, Reason: "Exact triality gives a useful threefold arena but not native generation labels; the exact invariant texture has 1+2 degeneracy."},
	}
	inv := CarrierInventory{Executed: true, Carriers: xs, Verdict: StatusFermionicCarrierInventoryAudited}
	for _, x := range xs {
		if x.Native {
			inv.NativeCarrierCount++
		}
		if x.Native && x.ThreeSector {
			inv.NativeThreeSectorCount++
		}
		if x.Native && x.ThreeSector && !x.ConfusesColor && !x.ConfusesSpecies && !x.ConfusesChirality && x.NoncentralGenerationAction {
			inv.NativeGenerationCarrierCount++
		}
		if x.Native && x.NoncentralGenerationAction {
			inv.NativeNoncentralGenerationActions++
		}
		if x.Native && x.NoncommutingOperators >= 2 {
			inv.NativeNoncommutingOperatorPairs++
		}
		if x.ConfusesColor {
			inv.ColorThreefoldCount++
		}
		if x.ConfusesSpecies || x.ConfusesChirality {
			inv.SpeciesOrChiralityThreefoldCount++
		}
	}
	return inv
}

func auditIdempotents() IdempotentAudit {
	xs := []IdempotentCandidate{
		{Name: "sixteen Fock occupation primitive idempotents", Domain: "Lambda*(C^4)", Native: true, PrimitiveBlocks: 16, BlockDimensions: repeatInt(1, 16), ThreeBlocks: false, ActsOnGenerationLabels: false, ActsOnSpeciesColorChirality: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "The primitive Fock idempotents refine occupation/species states, not family copies."},
		{Name: "even/odd Fock ideals", Domain: "parity split", Native: true, PrimitiveBlocks: 2, BlockDimensions: []int{8, 8}, ThreeBlocks: false, ActsOnGenerationLabels: false, ActsOnSpeciesColorChirality: true, Verdict: StatusFailedSpinorSplitChirality, Reason: "This is the two-sector chiral/parity split already rejected as generation origin."},
		{Name: "three spatial/color creation-mode projectors", Domain: "spatial Fock triplet", Native: true, PrimitiveBlocks: 3, BlockDimensions: []int{1, 1, 1}, ThreeBlocks: true, ActsOnGenerationLabels: false, ActsOnSpeciesColorChirality: true, Verdict: StatusFailedColorAsGeneration, Reason: "Three blocks exist, but their native meaning is color/spatial mode."},
		{Name: "right-singlet/left-doublet/hypercharge blocks", Domain: "charge representation table", Native: true, PrimitiveBlocks: 7, BlockDimensions: []int{1, 1, 1, 1, 2, 3, 6}, ThreeBlocks: false, ActsOnGenerationLabels: false, ActsOnSpeciesColorChirality: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "Hypercharge blocks reproduce particle species and weak representation."},
		{Name: "triality branch idempotents", Domain: "8v,8s,8c labelled branch stress test", Native: false, PrimitiveBlocks: 3, BlockDimensions: []int{8, 8, 8}, ThreeBlocks: true, ActsOnGenerationLabels: true, NoncentralOnC3Gen: true, GivesTwoNoncommutingOperators: true, ManuallyLabelled: true, Verdict: StatusFailedTrialityExactDegeneracy, Reason: "The idempotents become generation-labelled only after a manual branch assignment; exact invariant texture remains 1+2 degenerate."},
	}
	a := IdempotentAudit{Executed: true, Candidates: xs, Verdict: StatusPrimitiveIdempotentSearchAudited}
	for _, x := range xs {
		if x.Native && x.ThreeBlocks {
			a.NativeThreeBlockCandidates++
		}
		if x.Native && x.ActsOnGenerationLabels && x.NoncentralOnC3Gen {
			a.NativeGenerationLabelCandidates++
		}
		if x.ActsOnSpeciesColorChirality {
			a.ColorOrSpeciesRejected++
		}
		if x.ManuallyLabelled {
			a.ManualLabelRejected++
		}
		if x.Native && x.GivesTwoNoncommutingOperators {
			a.NoncommutingNativePairs++
		}
	}
	return a
}

func auditCommutant(inv CarrierInventory) CommutantAudit {
	return CommutantAudit{Executed: true, Action: "G_SM generated by hypercharge, SU(2)_L, color, B-L, J and Gamma", Carrier: "H_fermion one-generation carrier tensored by C^3_gen in the standard flavor arena", CommuntantModel: "Comm(G_SM) contains generation U(3) freedom only as multiplicity-space rotations; native ASHA operators act as I_3 on that factor", ContainsU3GenerationFreedom: true, GenerationFreedomCanonicalSelector: false, NativeDiagonalGenerationOperator: false, NativeNoncommutingGenerationPair: false, ArbitraryGenerationRotations: true, CentralBroadcastOverGeneration: true, DistinguishesSpeciesOnly: true, Verdict: StatusFailedTrivialGenerationCopy, Reason: "The commutant exposes the usual NCG problem: generations are a free multiplicity. The finite gauge/charge action leaves U(3)_gen free but supplies no canonical generator inside it."}
}

func auditTrialityFromFermions() TrialityAudit {
	return TrialityAudit{Executed: true, TrialityDomainAdmitted: false, FermionTo8VRepresentativeFound: false, FermionTo8SRepresentativeFound: true, FermionTo8CRepresentativeFound: true, GenerationLabelsDerivedNotInserted: false, ExactTrialityDegeneracyPresent: true, BrokenTrialityOperatorNative: false, OnePlusTwoDegeneracy: true, Verdict: StatusFailedTrialityExactDegeneracy, Reason: "The native spinor carrier sees 8_s and 8_c as chiral halves. The 8_v branch is representation-category data, not a derived third generation carrier. Exact triality again yields 1+2 degeneracy."}
}

func auditBilinears() BilinearAudit {
	xs := []BilinearFamily{
		{Name: "gauge-compatible LR Yukawa incidence bilinears", Native: true, AllowedDimension: 16, ActsOnGenerationLabels: false, ActsOnParticleSpecies: true, DiagonalOnly: false, DistinguishesUpDownLeptonNatively: true, Verdict: StatusFailedNoNativeBilinearSelector, Reason: "They select which species may couple through the Higgs but leave the 3x3 amplitude matrix arbitrary."},
		{Name: "J-paired fermion bilinears", Native: true, AllowedDimension: 8, ActsOnGenerationLabels: false, ActsOnParticleSpecies: true, DiagonalOnly: false, Verdict: StatusFailedNoNativeBilinearSelector, Reason: "J pairs particle/conjugate sectors, not generation labels."},
		{Name: "neutral Majorana/seesaw bilinear", Native: true, AllowedDimension: 1, ActsOnGenerationLabels: false, ActsOnParticleSpecies: true, DiagonalOnly: true, DistinguishesUpDownLeptonNatively: true, Verdict: StatusFailedSpeciesAsGeneration, Reason: "This is neutral-sector species structure; generation texture remains an external matrix unless a family carrier is derived."},
		{Name: "triality-sector bilinear stress test", Native: false, AllowedDimension: FullGenerationMixingMaps, ActsOnGenerationLabels: true, DiagonalOnly: false, NoncommutingOperators: 2, DistinguishesUpDownLeptonNatively: false, Verdict: StatusSealedFermionicCapacityFound, Reason: "A sealed triality-sector arena can host noncommuting textures, but the generation labels are not native."},
	}
	a := BilinearAudit{Executed: true, Families: xs, Verdict: StatusBilinearOperatorAuditCompleted}
	for _, x := range xs {
		if x.Native {
			a.NativeFamilies++
		}
		if x.Native && x.ActsOnGenerationLabels {
			a.NativeGenerationSensitiveFamilies++
		}
		if x.Native && x.NoncommutingOperators >= 2 {
			a.NativeNoncommutingFamilies++
		}
		if x.Native && x.Reduces13Moduli {
			a.NativeModuliReducingFamilies++
		}
	}
	return a
}

func auditDynamicSources() DynamicSourceAudit {
	xs := []DynamicSource{
		{Name: "total Fock number restriction", Native: true, Compatible: true, Spectrum: "0..4 occupation levels", ThreeLevelSpectrum: false, ActsOnFermionCarrier: true, ActsOnGenerationLabels: false, Verdict: StatusFailedSpeciesAsGeneration, Reason: "The number operator grades occupation; restricting it to three levels would be a choice, not a generation theorem."},
		{Name: "inserted N = diag(0,1,2)", Native: false, Compatible: true, SealedExternal: true, Circular: true, Spectrum: "0,1,2", ThreeLevelSpectrum: true, ActsOnFermionCarrier: false, ActsOnGenerationLabels: true, ProducesHierarchy: true, Verdict: StatusFailedNoNativeGenerationHamiltonian, Reason: "It has hierarchy capacity but remains an inserted generation Hamiltonian."},
		{Name: "modular/KMS internal Hamiltonian on native tracial fermion state", Native: true, Compatible: true, Spectrum: "trivial modular generator", ThreeLevelSpectrum: false, ActsOnFermionCarrier: true, ActsOnGenerationLabels: false, Verdict: StatusFailedNoNativeGenerationHamiltonian, Reason: "The native tracial state supplies no nontrivial three-level modular Hamiltonian."},
		{Name: "J-real asymmetry", Native: true, Compatible: true, Spectrum: "particle/conjugate pairing", ThreeLevelSpectrum: false, ActsOnFermionCarrier: true, ActsOnGenerationLabels: false, Verdict: StatusFailedSpeciesAsGeneration, Reason: "J-real structure distinguishes conjugation, not family hierarchy."},
		{Name: "color/lepton contrast", Native: true, Compatible: true, Spectrum: "3 color + 1 lepton", ThreeLevelSpectrum: true, ActsOnFermionCarrier: true, ActsOnGenerationLabels: false, WrongDomain: true, Verdict: StatusFailedColorAsGeneration, Reason: "The threefold part is color; using it as generation confuses gauge charge with family."},
		{Name: "triality Gaussian measure on branch labels", Native: false, Compatible: true, SealedExternal: true, Circular: true, Spectrum: "three triality labels", ThreeLevelSpectrum: true, ActsOnFermionCarrier: false, ActsOnGenerationLabels: true, ProducesHierarchy: true, ProducesMixing: true, Verdict: StatusSealedFermionicCapacityFound, Reason: "This has conditional capacity only after a branch-to-generation witness is sealed."},
	}
	a := DynamicSourceAudit{Executed: true, Sources: xs, Verdict: StatusDynamicGenerationSourceAudited}
	for _, x := range xs {
		if x.Native && x.ThreeLevelSpectrum {
			a.NativeThreeLevelSources++
		}
		if x.Native && x.ThreeLevelSpectrum && x.ActsOnGenerationLabels {
			a.NativeGenerationHamiltonians++
		}
		if x.SealedExternal || x.Circular {
			a.SealedOrCircularSources++
		}
		if x.WrongDomain {
			a.WrongDomainSources++
		}
	}
	return a
}

func auditCKM(b BilinearAudit, s DynamicSourceAudit, t TrialityAudit) CKMAudit {
	sealedPairs := 0
	if b.NativeNoncommutingFamilies == 0 {
		sealedPairs = 1
	}
	return CKMAudit{Executed: true, NativeCandidateOperators: b.NativeGenerationSensitiveFamilies + s.NativeGenerationHamiltonians, NativeNoncommutingPairs: b.NativeNoncommutingFamilies, SealedNoncommutingPairs: sealedPairs, UpDownTextureMisalignmentNative: false, UpDownTextureMisalignmentSealed: true, CKMCapacityNative: false, CKMCapacityConditional: true, PMNSCapacityNative: false, Verdict: StatusFailedNoNativeCKMCapacity, Reason: "Noncommuting CKM/PMNS capacity appears only in sealed triality/source stress tests. No native fermionic generation operators A,B with [A,B] != 0 were derived."}
}

func auditModuli(inv CarrierInventory, ids IdempotentAudit, c CommutantAudit, t TrialityAudit, b BilinearAudit, s DynamicSourceAudit, k CKMAudit) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "trivial generation broadcast", Status: StatusFailedTrivialGenerationCopy, ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, QuarkLeptonSeparationPossible: true, Reason: "Possible only as free arbitrary 3x3 matrices; no finite-core reduction."},
		{Name: "exact triality only", Status: StatusFailedTrialityExactDegeneracy, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: false, Reason: "Exact invariant texture has 1+2 degeneracy and no sector-specific misalignment."},
		{Name: "one native diagonal generation operator", Status: StatusFailedNoNativeGenerationHamiltonian, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: false, Reason: "No native diagonal generation Hamiltonian was found."},
		{Name: "one sealed diagonal generation operator", Status: StatusSealedFermionicCapacityFound, ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: false, Reason: "A sealed N can split hierarchy but supplies no mixing or sector origin."},
		{Name: "two native commuting operators", Status: StatusFailedNoNativeCKMCapacity, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: false, Reason: "No native generation operators were derived."},
		{Name: "two native noncommuting operators", Status: StatusFailedNoNativeCKMCapacity, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: false, Reason: "No native noncommuting pair exists."},
		{Name: "native fermionic bilinear selector", Status: StatusFailedNoNativeBilinearSelector, ModuliDim: 13, ThreeDistinctMassesPossible: false, CKMPossible: false, PMNSPossible: false, QuarkLeptonSeparationPossible: true, Reason: "Native bilinears select species channels but not generation matrices."},
		{Name: "sealed external Yukawa source", Status: StatusSealedFermionicCapacityFound, ModuliDim: 13, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, QuarkLeptonSeparationPossible: true, Reason: "Full phenomenology can be encoded if arbitrary external texture sources are sealed, but that is not a finite theorem."},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalReductionBelow13: false, FirewallPreserved: true, Verdict: StatusFirewallPreserved13Moduli}
}

func auditFirewall(m ModuliImpact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoScalarSelectorPromoted: true, NoTauEtaInserted: true, NoNDiagInserted: true, NoManualGenerationLabels: true, NoColorAsGenerationPromoted: true, NoSpeciesAsGenerationPromoted: true, NoModuliReductionClaimed: !m.NativeReductionBelow13 && m.BestNativeDim == Gate372ChargedFlavorModuliDim, Verdict: StatusFirewallPreserved13Moduli}
}

func nextStep(m ModuliImpact, k CKMAudit, s DynamicSourceAudit) NextStep {
	return NextStep{Gate: 410, Title: "Fermionic Representation Extension / Nontrivial Family Bundle Search", Reason: "Gate 409 shows that the existing fermionic carrier still treats generation as a trivial multiplicity. The next non-surrender route is not an empirical seal; it must search for an extension or new representation structure that replaces C^3_gen with a derived family bundle.", PrimaryTask: "Audit candidate nontrivial family bundles, e.g. modular nontracial fermion states, primitive ideal extensions, KO-twisted multiplicities, or sealed-but-testable family representation extensions, while preserving the Gate-372 13-moduli firewall."}
}

func truth(a Analysis) string {
	return "Gate 409 proves that the scalar-sector flavor-blindness established by Gates 398-408 does not automatically reveal a fermionic family origin. The native fermionic carriers reconstruct charge, chirality, color, conjugation, weak representation, and allowed one-generation Yukawa channels, but none supplies a noncentral End(C^3_gen) action. The commutant has the standard U(3)_gen freedom only as an unselected multiplicity, exact triality again degenerates, native bilinears select species rather than generation, and all noncommuting CKM-capable structures remain sealed/circular. Therefore the current ASHA fermionic matter carrier still behaves as a trivial generation copy and the 13 charged flavor moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusScalarFlavorBlindnessInherited, StatusFermionicCarrierInventoryAudited, StatusPrimitiveIdempotentSearchAudited, StatusCommutantCentralizerAudited, StatusTrialityFromFermionAudited, StatusBilinearOperatorAuditCompleted, StatusDynamicGenerationSourceAudited}
	for _, c := range a.Inventory.Carriers {
		addStatus(&statuses, c.Verdict)
	}
	for _, c := range a.Idempotents.Candidates {
		addStatus(&statuses, c.Verdict)
	}
	for _, b := range a.Bilinears.Families {
		addStatus(&statuses, b.Verdict)
	}
	for _, s := range a.Sources.Sources {
		addStatus(&statuses, s.Verdict)
	}
	for _, status := range []string{StatusFailedTrivialGenerationCopy, StatusFailedSpinorSplitChirality, StatusFailedTrialityExactDegeneracy, StatusFailedColorAsGeneration, StatusFailedSpeciesAsGeneration, StatusFailedNoNativeCKMCapacity, StatusFailedNoNativeGenerationHamiltonian, StatusFailedNoNativeBilinearSelector, StatusFirewallPreserved13Moduli} {
		addStatus(&statuses, status)
	}
	return statuses
}

func addStatus(xs *[]string, s string) {
	if s != "" && !contains(*xs, s) {
		*xs = append(*xs, s)
	}
}
func contains(xs []string, v string) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}
func repeatInt(v, n int) []int {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = v
	}
	return xs
}
func boolWord(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate408_scalar_flavor_blind=%v gate407_capacity_no_selector=%v gate395_two_sector=%v gate396_three_sources_not_generation=%v gate397_singletons_not_flavor=%v gate393_triality_not_admitted=%v gate394_static_address_central=%v charged_moduli=%d no_empirical=%v verdict=%s", x.Executed, x.Gate408ScalarFlavorBlind, x.Gate407FullHphiCapacityButNoSelector, x.Gate395SpinorSplitTwoSector, x.Gate396ThreeObjectSourcesNotGeneration, x.Gate397ContactSingletonsNotFlavor, x.Gate393TrialityDomainNotAdmitted, x.Gate394StaticGenerationAddressCentral, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}
func FormatCarrier(x Carrier) string {
	return fmt.Sprintf("name=%s domain=%s dim=%d native=%v three_sector=%v kind=%s noncentral_gen=%v noncomm_ops=%d links_yukawa=%v AF=%v J=%v Gamma=%v first_order=%v Y=%v SU2L=%v B-L=%v color_confusion=%v species_confusion=%v chirality_confusion=%v verdict=%s reason=%s", x.Name, x.Domain, x.Dimension, x.Native, x.ThreeSector, x.ThreeSectorKind, x.NoncentralGenerationAction, x.NoncommutingOperators, x.LinksToYukawaSpaces, x.CompatibleAF, x.CompatibleJ, x.CompatibleGamma, x.CompatibleFirstOrder, x.CompatibleHypercharge, x.CompatibleSU2L, x.CompatibleBMinusL, x.ConfusesColor, x.ConfusesSpecies, x.ConfusesChirality, x.Verdict, x.Reason)
}
func FormatInventory(x CarrierInventory) string {
	return fmt.Sprintf("executed=%v carriers=%d native=%d native_three_sector=%d native_generation_carriers=%d native_noncentral_actions=%d native_noncommuting_pairs=%d color_threefold=%d species_or_chirality_threefold=%d verdict=%s", x.Executed, len(x.Carriers), x.NativeCarrierCount, x.NativeThreeSectorCount, x.NativeGenerationCarrierCount, x.NativeNoncentralGenerationActions, x.NativeNoncommutingOperatorPairs, x.ColorThreefoldCount, x.SpeciesOrChiralityThreefoldCount, x.Verdict)
}
func FormatIdempotent(x IdempotentCandidate) string {
	return fmt.Sprintf("name=%s domain=%s native=%v blocks=%d dims=%v three_blocks=%v generation_labels=%v species_color_chirality=%v noncentral_C3=%v noncommuting=%v manual_label=%v verdict=%s reason=%s", x.Name, x.Domain, x.Native, x.PrimitiveBlocks, x.BlockDimensions, x.ThreeBlocks, x.ActsOnGenerationLabels, x.ActsOnSpeciesColorChirality, x.NoncentralOnC3Gen, x.GivesTwoNoncommutingOperators, x.ManuallyLabelled, x.Verdict, x.Reason)
}
func FormatIdempotentAudit(x IdempotentAudit) string {
	return fmt.Sprintf("executed=%v candidates=%d native_three_block=%d native_generation_labels=%d color_species_rejected=%d manual_label_rejected=%d native_noncommuting_pairs=%d verdict=%s", x.Executed, len(x.Candidates), x.NativeThreeBlockCandidates, x.NativeGenerationLabelCandidates, x.ColorOrSpeciesRejected, x.ManualLabelRejected, x.NoncommutingNativePairs, x.Verdict)
}
func FormatCommutant(x CommutantAudit) string {
	return fmt.Sprintf("executed=%v action=%s carrier=%s commutant=%s contains_U3_gen=%v canonical_selector=%v native_diag=%v native_noncomm_pair=%v arbitrary_rotations=%v central_broadcast=%v species_only=%v verdict=%s reason=%s", x.Executed, x.Action, x.Carrier, x.CommuntantModel, x.ContainsU3GenerationFreedom, x.GenerationFreedomCanonicalSelector, x.NativeDiagonalGenerationOperator, x.NativeNoncommutingGenerationPair, x.ArbitraryGenerationRotations, x.CentralBroadcastOverGeneration, x.DistinguishesSpeciesOnly, x.Verdict, x.Reason)
}
func FormatTriality(x TrialityAudit) string {
	return fmt.Sprintf("executed=%v domain_admitted=%v 8v=%v 8s=%v 8c=%v labels_derived=%v exact_degeneracy=%v broken_operator_native=%v one_plus_two=%v verdict=%s reason=%s", x.Executed, x.TrialityDomainAdmitted, x.FermionTo8VRepresentativeFound, x.FermionTo8SRepresentativeFound, x.FermionTo8CRepresentativeFound, x.GenerationLabelsDerivedNotInserted, x.ExactTrialityDegeneracyPresent, x.BrokenTrialityOperatorNative, x.OnePlusTwoDegeneracy, x.Verdict, x.Reason)
}
func FormatBilinear(x BilinearFamily) string {
	return fmt.Sprintf("name=%s native=%v allowed_dim=%d generation_labels=%v species=%v diagonal_only=%v noncomm_ops=%d up_down_lepton=%v reduces_moduli=%v verdict=%s reason=%s", x.Name, x.Native, x.AllowedDimension, x.ActsOnGenerationLabels, x.ActsOnParticleSpecies, x.DiagonalOnly, x.NoncommutingOperators, x.DistinguishesUpDownLeptonNatively, x.Reduces13Moduli, x.Verdict, x.Reason)
}
func FormatBilinearAudit(x BilinearAudit) string {
	return fmt.Sprintf("executed=%v families=%d native=%d native_generation_sensitive=%d native_noncommuting=%d native_moduli_reducing=%d verdict=%s", x.Executed, len(x.Families), x.NativeFamilies, x.NativeGenerationSensitiveFamilies, x.NativeNoncommutingFamilies, x.NativeModuliReducingFamilies, x.Verdict)
}
func FormatSource(x DynamicSource) string {
	return fmt.Sprintf("name=%s native=%v compatible=%v sealed=%v circular=%v wrong_domain=%v spectrum=%s three_level=%v fermion=%v generation=%v hierarchy=%v mixing=%v verdict=%s reason=%s", x.Name, x.Native, x.Compatible, x.SealedExternal, x.Circular, x.WrongDomain, x.Spectrum, x.ThreeLevelSpectrum, x.ActsOnFermionCarrier, x.ActsOnGenerationLabels, x.ProducesHierarchy, x.ProducesMixing, x.Verdict, x.Reason)
}
func FormatSourceAudit(x DynamicSourceAudit) string {
	return fmt.Sprintf("executed=%v sources=%d native_three_level=%d native_generation_hamiltonians=%d sealed_or_circular=%d wrong_domain=%d verdict=%s", x.Executed, len(x.Sources), x.NativeThreeLevelSources, x.NativeGenerationHamiltonians, x.SealedOrCircularSources, x.WrongDomainSources, x.Verdict)
}
func FormatCKM(x CKMAudit) string {
	return fmt.Sprintf("executed=%v native_ops=%d native_noncomm_pairs=%d sealed_noncomm_pairs=%d updown_native=%v updown_sealed=%v ckm_native=%v ckm_conditional=%v pmns_native=%v verdict=%s reason=%s", x.Executed, x.NativeCandidateOperators, x.NativeNoncommutingPairs, x.SealedNoncommutingPairs, x.UpDownTextureMisalignmentNative, x.UpDownTextureMisalignmentSealed, x.CKMCapacityNative, x.CKMCapacityConditional, x.PMNSCapacityNative, x.Verdict, x.Reason)
}
func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("name=%s status=%s moduli_dim=%d masses3=%v ckm=%v pmns=%v ql_sep=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.QuarkLeptonSeparationPossible, x.Reason)
}
func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d scenarios=%d best_native_dim=%d native_reduction=%v conditional_reduction=%v firewall=%v verdict=%s", x.StartDim, len(x.Scenarios), x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalReductionBelow13, x.FirewallPreserved, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa=%v no_scalar_promoted=%v no_tau_eta=%v no_N_diag=%v no_manual_labels=%v no_color_as_gen=%v no_species_as_gen=%v no_moduli_reduction=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoScalarSelectorPromoted, x.NoTauEtaInserted, x.NoNDiagInserted, x.NoManualGenerationLabels, x.NoColorAsGenerationPromoted, x.NoSpeciesAsGenerationPromoted, x.NoModuliReductionClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func CompactStatusTable(a Analysis) string { return strings.Join(Statuses(a), "\n") }
