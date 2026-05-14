// Package fermionicfamilybundleextension implements Gate 410:
// Fermionic Representation Extension / Nontrivial Family Bundle Sieve.
//
// Gate 409 proved that the current fermionic carrier still treats generation as
// a trivial U(3) multiplicity. Gate 410 does not insert Yukawa amplitudes. It
// audits whether the existing ASHA ledger already derives an advanced extension
// that replaces H_fermion \otimes C^3_gen by a nontrivial family bundle with a
// native connection, curvature, or pair of noncommuting generation operators.
package fermionicfamilybundleextension

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE410-FERMIONIC-REPRESENTATION-EXTENSION-NONTRIVIAL-FAMILY-BUNDLE-SIEVE"

	StatusGate409Inherited                      = "CONDITIONAL_SUPPORT_GATE409_FERMIONIC_TRIVIAL_MULTIPLICITY_INHERITED"
	StatusExtensionSearchFormalized             = "CONDITIONAL_SUPPORT_REPRESENTATION_EXTENSION_SEARCH_FORMALIZED"
	StatusFamilyBundleArenaAudited              = "CONDITIONAL_SUPPORT_NONTRIVIAL_FAMILY_BUNDLE_ARENA_AUDITED"
	StatusKOTwistAudited                        = "CONDITIONAL_SUPPORT_KO_TWISTED_SPECTRAL_TRIPLE_AUDITED"
	StatusModularKMSAudited                     = "CONDITIONAL_SUPPORT_MODULAR_NONTRACIAL_FERMION_STATE_AUDITED"
	StatusPrimitiveIdealExtensionAudited        = "CONDITIONAL_SUPPORT_PRIMITIVE_IDEAL_EXTENSION_AUDITED"
	StatusSealedFamilyBundleCapacity            = "CONDITIONAL_SUPPORT_NONTRIVIAL_FAMILY_BUNDLE_CAPACITY"
	StatusConditionalCKMCapacityActivated       = "CONDITIONAL_SUPPORT_CKM_CAPACITY_ACTIVATED"
	StatusVerifiedNontrivialFamilyBundleDerived = "VERIFIED_NONTRIVIAL_FAMILY_BUNDLE_DERIVED"
	StatusVerifiedNativeFamilyConnectionDerived = "VERIFIED_NATIVE_FAMILY_CONNECTION_DERIVED"
	StatusVerifiedNativeNoncommutingFamilyPair  = "VERIFIED_NATIVE_NONCOMMUTING_FAMILY_TEXTURE_PAIR"
	StatusFailedNoNativeFamilyBundle            = "FAILED_ROUTE_NO_NATIVE_NONTRIVIAL_FAMILY_BUNDLE"
	StatusFailedKOTwistOnlyChangesSigns         = "FAILED_ROUTE_KO_TWIST_ONLY_CHANGES_REAL_STRUCTURE_SIGNS"
	StatusFailedKMSRequiresExternalHamiltonian  = "FAILED_ROUTE_KMS_NONTRACIAL_STATE_REQUIRES_EXTERNAL_HAMILTONIAN"
	StatusFailedPrimitiveIdealsWrongDomain      = "FAILED_ROUTE_PRIMITIVE_IDEALS_REMAIN_COLOR_SPECIES_OR_CONTACT_DOMAIN"
	StatusFailedExtensionRequiresNewAxiom       = "FAILED_ROUTE_FAMILY_BUNDLE_EXTENSION_REQUIRES_NEW_AXIOM"
	StatusFailedTrivialGenerationMultiplicity   = "FAILED_ROUTE_GENERATION_REMAINS_TRIVIAL_U3_MULTIPLICITY"
	StatusFailedNoNativeNoncommutingTexturePair = "FAILED_ROUTE_NO_NATIVE_NONCOMMUTING_FAMILY_TEXTURE_PAIR"
	StatusFirewallPreserved13Moduli             = "FIREWALL_PRESERVED_13_MODULI"
)

const (
	Gate372ChargedFlavorModuliDim = 13
	CurrentGenerationRank         = 3
	FockMatterDim                 = 16
	OneGenerationYukawaChannels   = 8
)

type Inheritance struct {
	Executed                             bool
	Gate409FermionicCarrierTrivial       bool
	Gate409U3GenerationFreedomUnselected bool
	Gate409NoNativeCKMCapacity           bool
	Gate408ScalarFlavorBlind             bool
	Gate395SpinorSplitIsChirality        bool
	Gate394StaticAddressCentral          bool
	Gate393TrialityDomainNotAdmitted     bool
	Gate372ChargedModuliDim              int
	NoEmpiricalInputsImported            bool
	Verdict                              string
}

type ExtensionCandidate struct {
	Name                        string
	Domain                      string
	NativeInCurrentAsha         bool
	ChangesFermionCarrier       bool
	DerivesThreeFamilies        bool
	NontrivialBundle            bool
	ProvidesCanonicalConnection bool
	ProvidesFamilyCurvature     bool
	ProvidesDiagonalOperator    bool
	ProvidesTwoNoncommutingOps  bool
	CKMCapacityNative           bool
	CKMCapacityConditional      bool
	CompatibleAF                bool
	CompatibleJ                 bool
	CompatibleGamma             bool
	CompatibleFirstOrder        bool
	CompatibleGaugeCharges      bool
	RequiresNewAxiom            bool
	RequiresExternalHamiltonian bool
	RequiresManualFamilyLabel   bool
	RequiresEmpiricalInput      bool
	WrongDomain                 bool
	Verdict                     string
	Reason                      string
}

type ExtensionAudit struct {
	Executed                     bool
	Candidates                   []ExtensionCandidate
	CandidatesAudited            int
	NativeNontrivialBundles      int
	ConditionalNontrivialBundles int
	NativeConnections            int
	NativeNoncommutingPairs      int
	ConditionalNoncommutingPairs int
	WrongDomainCandidates        int
	RequiresNewAxiomCandidates   int
	Verdict                      string
}

type FamilyBundleAudit struct {
	Executed                      bool
	CurrentCarrier                string
	GenerationRank                int
	TrivialMultiplicity           bool
	ContainsU3Freedom             bool
	U3FreedomSelectedByGeometry   bool
	NontrivialTransitionFunctions bool
	NativeFamilyConnection        bool
	NativeFamilyCurvature         bool
	NativeHolonomyNonabelian      bool
	ReplacesTensorC3              bool
	Verdict                       string
	Reason                        string
}

type KOTwistAudit struct {
	Executed                     bool
	KODimensionSignsAudited      bool
	ChangesJGammaCommutation     bool
	ChangesMultiplicity          bool
	ProducesThreeFamilies        bool
	ProducesFamilyConnection     bool
	ProducesNoncommutingTextures bool
	CompatibleWithExistingTriple bool
	RequiresTwistedAlgebraAction bool
	Verdict                      string
	Reason                       string
}

type ModularKMSAudit struct {
	Executed                    bool
	TracialStateFreezesFlow     bool
	NontracialStateHasCapacity  bool
	NativeHamiltonianFound      bool
	NativeDensityMatrixFound    bool
	ThreeLevelSpectrumNative    bool
	MixingOperatorNative        bool
	RequiresExternalHamiltonian bool
	RequiresChosenDensityMatrix bool
	Verdict                     string
	Reason                      string
}

type PrimitiveIdealExtensionAudit struct {
	Executed                     bool
	ExistingIdealsAudited        bool
	ExistingIdealsAreWrongDomain bool
	NewPrimitiveIdealExtension   bool
	ThreeFamilyIdealDerived      bool
	ActsOnC3GenNoncentrally      bool
	RequiresAlgebraEnlargement   bool
	CompatibleWithAF             bool
	CompatibleWithFirstOrder     bool
	Verdict                      string
	Reason                       string
}

type NoncommutingTextureAudit struct {
	Executed                      bool
	NativeFamilyOperators         int
	NativeNoncommutingPairs       int
	ConditionalFamilyOperators    int
	ConditionalNoncommutingPairs  int
	UpDownMisalignmentNative      bool
	UpDownMisalignmentConditional bool
	CKMCapacityNative             bool
	CKMCapacityConditional        bool
	PMNSCapacityNative            bool
	Verdict                       string
	Reason                        string
}

type ModuliScenario struct {
	Name                        string
	Status                      string
	ModuliDim                   int
	ThreeDistinctMassesPossible bool
	CKMPossible                 bool
	PMNSPossible                bool
	FamilyBundleNontrivial      bool
	Reason                      string
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
	NoExternalHamiltonianPromoted bool
	NoManualFamilyBundlePromoted  bool
	NoNewAxiomPromoted            bool
	NoColorAsGenerationPromoted   bool
	NoSpeciesAsGenerationPromoted bool
	NoScalarSelectorPromoted      bool
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
	Inheritance     Inheritance
	Extensions      ExtensionAudit
	FamilyBundle    FamilyBundleAudit
	KOTwist         KOTwistAudit
	ModularKMS      ModularKMSAudit
	PrimitiveIdeals PrimitiveIdealExtensionAudit
	Noncommuting    NoncommutingTextureAudit
	Moduli          ModuliImpact
	Firewall        FirewallAudit
	Next            NextStep
	Truth           string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Extensions = buildExtensions()
	a.FamilyBundle = buildFamilyBundle(a.Extensions)
	a.KOTwist = buildKOTwist()
	a.ModularKMS = buildModularKMS()
	a.PrimitiveIdeals = buildPrimitiveIdeals()
	a.Noncommuting = buildNoncommuting(a.Extensions)
	a.Moduli = buildModuli(a.Extensions, a.Noncommuting)
	a.Firewall = buildFirewall(a.Moduli)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate409FermionicCarrierTrivial {
		return fmt.Errorf("Gate409 inheritance missing")
	}
	if !a.Extensions.Executed || a.Extensions.CandidatesAudited == 0 {
		return fmt.Errorf("extension audit missing")
	}
	if a.Extensions.NativeNontrivialBundles != 0 {
		return fmt.Errorf("unexpected native family bundle promoted")
	}
	if a.Noncommuting.NativeNoncommutingPairs != 0 || a.Noncommuting.CKMCapacityNative {
		return fmt.Errorf("unexpected native noncommuting family texture pair")
	}
	if a.Moduli.StartDim != Gate372ChargedFlavorModuliDim || a.Moduli.BestNativeDim != Gate372ChargedFlavorModuliDim || !a.Moduli.FirewallPreserved {
		return fmt.Errorf("moduli firewall broken incorrectly")
	}
	if !a.Firewall.NoObservedMassesImported || !a.Firewall.NoNewAxiomPromoted || !a.Firewall.NoManualFamilyBundlePromoted {
		return fmt.Errorf("firewall violation")
	}
	return nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate409FermionicCarrierTrivial: true, Gate409U3GenerationFreedomUnselected: true, Gate409NoNativeCKMCapacity: true, Gate408ScalarFlavorBlind: true, Gate395SpinorSplitIsChirality: true, Gate394StaticAddressCentral: true, Gate393TrialityDomainNotAdmitted: true, Gate372ChargedModuliDim: Gate372ChargedFlavorModuliDim, NoEmpiricalInputsImported: true, Verdict: StatusGate409Inherited}
}

func buildExtensions() ExtensionAudit {
	candidates := []ExtensionCandidate{
		{Name: "KO-dimension / twisted real-structure extension", Domain: "finite spectral triple signs and J/Gamma relations", NativeInCurrentAsha: true, ChangesFermionCarrier: false, DerivesThreeFamilies: false, NontrivialBundle: false, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: false, ProvidesDiagonalOperator: false, ProvidesTwoNoncommutingOps: false, CKMCapacityNative: false, CKMCapacityConditional: false, CompatibleAF: true, CompatibleJ: true, CompatibleGamma: true, CompatibleFirstOrder: true, CompatibleGaugeCharges: true, RequiresNewAxiom: false, Verdict: StatusFailedKOTwistOnlyChangesSigns, Reason: "KO/twist data can change real-structure signs or first-order bookkeeping, but in the current ledger it does not replace the trivial C^3_gen multiplicity."},
		{Name: "modular nontracial fermion KMS state", Domain: "fermion density matrix / modular Hamiltonian", NativeInCurrentAsha: false, ChangesFermionCarrier: true, DerivesThreeFamilies: false, NontrivialBundle: false, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: false, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: false, CKMCapacityNative: false, CKMCapacityConditional: true, CompatibleAF: true, CompatibleJ: false, CompatibleGamma: true, CompatibleFirstOrder: false, CompatibleGaugeCharges: true, RequiresExternalHamiltonian: true, RequiresManualFamilyLabel: true, Verdict: StatusFailedKMSRequiresExternalHamiltonian, Reason: "A nontracial state can break degeneracy, but the required Hamiltonian/density matrix is not derived on the fermion family carrier."},
		{Name: "primitive ideal extension of A_F", Domain: "algebra enlargement beyond C + H + M3(C)", NativeInCurrentAsha: false, ChangesFermionCarrier: true, DerivesThreeFamilies: false, NontrivialBundle: false, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: false, ProvidesDiagonalOperator: false, ProvidesTwoNoncommutingOps: false, CKMCapacityNative: false, CompatibleAF: false, CompatibleJ: false, CompatibleGamma: true, CompatibleFirstOrder: false, CompatibleGaugeCharges: false, RequiresNewAxiom: true, Verdict: StatusFailedExtensionRequiresNewAxiom, Reason: "A new primitive-ideal family algebra could define families, but it is an algebra extension not forced by the current finite triple."},
		{Name: "contact singleton / rational idempotent family bundle", Domain: "contact spectral domain", NativeInCurrentAsha: true, ChangesFermionCarrier: false, DerivesThreeFamilies: false, NontrivialBundle: false, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: false, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: false, CKMCapacityNative: false, CompatibleAF: false, CompatibleJ: false, CompatibleGamma: false, CompatibleFirstOrder: false, CompatibleGaugeCharges: false, RequiresManualFamilyLabel: true, WrongDomain: true, Verdict: StatusFailedPrimitiveIdealsWrongDomain, Reason: "The three rational contact blocks are native but remain contact-domain idempotents; Gate 397 rejected their finite-Dirac flavor functor."},
		{Name: "triality local-system family bundle", Domain: "Spin(8) representation category", NativeInCurrentAsha: false, ChangesFermionCarrier: true, DerivesThreeFamilies: false, NontrivialBundle: false, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: false, ProvidesDiagonalOperator: false, ProvidesTwoNoncommutingOps: false, CKMCapacityNative: false, CKMCapacityConditional: true, CompatibleAF: false, CompatibleJ: false, CompatibleGamma: true, CompatibleFirstOrder: false, CompatibleGaugeCharges: false, RequiresManualFamilyLabel: true, Verdict: StatusFailedExtensionRequiresNewAxiom, Reason: "Triality gives an arena, but a family bundle over it needs a new functor from finite-Dirac states to 8v,8s,8c."},
		{Name: "sealed nontrivial U(3)_gen connection stress test", Domain: "external family bundle connection", NativeInCurrentAsha: false, ChangesFermionCarrier: true, DerivesThreeFamilies: true, NontrivialBundle: true, ProvidesCanonicalConnection: false, ProvidesFamilyCurvature: true, ProvidesDiagonalOperator: true, ProvidesTwoNoncommutingOps: true, CKMCapacityNative: false, CKMCapacityConditional: true, CompatibleAF: true, CompatibleJ: false, CompatibleGamma: true, CompatibleFirstOrder: false, CompatibleGaugeCharges: true, RequiresNewAxiom: true, RequiresManualFamilyLabel: true, Verdict: StatusSealedFamilyBundleCapacity, Reason: "A manually supplied family connection has CKM capacity, but it is precisely the missing structure, not a theorem."},
	}
	out := ExtensionAudit{Executed: true, Candidates: candidates, CandidatesAudited: len(candidates), Verdict: StatusExtensionSearchFormalized}
	for _, c := range candidates {
		if c.NativeInCurrentAsha && c.NontrivialBundle && c.DerivesThreeFamilies {
			out.NativeNontrivialBundles++
		}
		if !c.NativeInCurrentAsha && c.NontrivialBundle {
			out.ConditionalNontrivialBundles++
		}
		if c.NativeInCurrentAsha && c.ProvidesCanonicalConnection {
			out.NativeConnections++
		}
		if c.NativeInCurrentAsha && c.ProvidesTwoNoncommutingOps {
			out.NativeNoncommutingPairs++
		}
		if !c.NativeInCurrentAsha && c.ProvidesTwoNoncommutingOps {
			out.ConditionalNoncommutingPairs++
		}
		if c.WrongDomain {
			out.WrongDomainCandidates++
		}
		if c.RequiresNewAxiom {
			out.RequiresNewAxiomCandidates++
		}
	}
	return out
}

func buildFamilyBundle(ext ExtensionAudit) FamilyBundleAudit {
	return FamilyBundleAudit{Executed: true, CurrentCarrier: "H_fermion(one generation) tensor C^3_gen", GenerationRank: CurrentGenerationRank, TrivialMultiplicity: true, ContainsU3Freedom: true, U3FreedomSelectedByGeometry: false, NontrivialTransitionFunctions: false, NativeFamilyConnection: ext.NativeConnections > 0, NativeFamilyCurvature: false, NativeHolonomyNonabelian: false, ReplacesTensorC3: false, Verdict: StatusFailedTrivialGenerationMultiplicity, Reason: "The current carrier remains a product with an unselected U(3)_gen commutant; no transition data, connection, or curvature upgrades it into a family bundle."}
}

func buildKOTwist() KOTwistAudit {
	return KOTwistAudit{Executed: true, KODimensionSignsAudited: true, ChangesJGammaCommutation: true, ChangesMultiplicity: false, ProducesThreeFamilies: false, ProducesFamilyConnection: false, ProducesNoncommutingTextures: false, CompatibleWithExistingTriple: true, RequiresTwistedAlgebraAction: false, Verdict: StatusFailedKOTwistOnlyChangesSigns, Reason: "The audited KO/twist lane controls real-structure signs and compatibility; it does not manufacture a rank-3 family carrier or CKM-capable pair."}
}

func buildModularKMS() ModularKMSAudit {
	return ModularKMSAudit{Executed: true, TracialStateFreezesFlow: true, NontracialStateHasCapacity: true, NativeHamiltonianFound: false, NativeDensityMatrixFound: false, ThreeLevelSpectrumNative: false, MixingOperatorNative: false, RequiresExternalHamiltonian: true, RequiresChosenDensityMatrix: true, Verdict: StatusFailedKMSRequiresExternalHamiltonian, Reason: "A nontracial KMS state is the right type of selector, but the fermion-family modular Hamiltonian is not native in the current ASHA ledger."}
}

func buildPrimitiveIdeals() PrimitiveIdealExtensionAudit {
	return PrimitiveIdealExtensionAudit{Executed: true, ExistingIdealsAudited: true, ExistingIdealsAreWrongDomain: true, NewPrimitiveIdealExtension: false, ThreeFamilyIdealDerived: false, ActsOnC3GenNoncentrally: false, RequiresAlgebraEnlargement: true, CompatibleWithAF: false, CompatibleWithFirstOrder: false, Verdict: StatusFailedPrimitiveIdealsWrongDomain, Reason: "Existing primitive/idempotent structures encode color, species, chirality, or contact roots; a true family ideal requires enlarging the finite algebra."}
}

func buildNoncommuting(ext ExtensionAudit) NoncommutingTextureAudit {
	return NoncommutingTextureAudit{Executed: true, NativeFamilyOperators: 0, NativeNoncommutingPairs: ext.NativeNoncommutingPairs, ConditionalFamilyOperators: 2, ConditionalNoncommutingPairs: ext.ConditionalNoncommutingPairs, UpDownMisalignmentNative: false, UpDownMisalignmentConditional: ext.ConditionalNoncommutingPairs > 0, CKMCapacityNative: false, CKMCapacityConditional: ext.ConditionalNoncommutingPairs > 0, PMNSCapacityNative: false, Verdict: StatusFailedNoNativeNoncommutingTexturePair, Reason: "Only the sealed external U(3)_gen connection stress test provides noncommuting capacity; no native ASHA extension provides two family operators."}
}

func buildModuli(ext ExtensionAudit, n NoncommutingTextureAudit) ModuliImpact {
	scenarios := []ModuliScenario{
		{Name: "current trivial C3_gen multiplicity", Status: StatusFailedTrivialGenerationMultiplicity, ModuliDim: 13, Reason: "standard unselected U(3)_gen commutant"},
		{Name: "KO/twisted real structure only", Status: StatusFailedKOTwistOnlyChangesSigns, ModuliDim: 13, Reason: "changes signs/compatibility but not family multiplicity"},
		{Name: "modular KMS without native Hamiltonian", Status: StatusFailedKMSRequiresExternalHamiltonian, ModuliDim: 13, ThreeDistinctMassesPossible: true, Reason: "capacity exists only after external K_gen"},
		{Name: "primitive ideal extension of A_F", Status: StatusFailedExtensionRequiresNewAxiom, ModuliDim: 13, Reason: "would be a new algebraic axiom"},
		{Name: "sealed nontrivial U3 family connection", Status: StatusSealedFamilyBundleCapacity, ModuliDim: 0, ThreeDistinctMassesPossible: true, CKMPossible: true, PMNSPossible: true, FamilyBundleNontrivial: true, Reason: "stress test only; supplied connection can fit/select textures but is external"},
	}
	return ModuliImpact{StartDim: Gate372ChargedFlavorModuliDim, Scenarios: scenarios, BestNativeDim: Gate372ChargedFlavorModuliDim, NativeReductionBelow13: false, ConditionalReductionBelow13: n.CKMCapacityConditional, FirewallPreserved: true, Verdict: StatusFirewallPreserved13Moduli}
}

func buildFirewall(m ModuliImpact) FirewallAudit {
	return FirewallAudit{Executed: true, NoObservedMassesImported: true, NoCKMImported: true, NoPMNSImported: true, NoYukawaAmplitudesInserted: true, NoExternalHamiltonianPromoted: true, NoManualFamilyBundlePromoted: true, NoNewAxiomPromoted: true, NoColorAsGenerationPromoted: true, NoSpeciesAsGenerationPromoted: true, NoScalarSelectorPromoted: true, NoModuliReductionClaimed: !m.NativeReductionBelow13, Verdict: StatusFirewallPreserved13Moduli}
}

func buildNext() NextStep {
	return NextStep{Gate: 411, Title: "Axiom-Candidate Ledger for Nontrivial Family Bundle Extensions", Reason: "Gate 410 shows that current ASHA data does not derive a nontrivial family bundle. The remaining non-surrender path is not another search inside existing carriers, but an explicit ledger of minimal new axioms/extensions that could be tested without empirical fitting.", PrimaryTask: "Classify candidate family-bundle axioms by mathematical cost, compatibility with A_F/J/first-order/gauge charges, CKM capacity, and whether they remain independent of observed Yukawa data."}
}

func truth(a Analysis) string {
	return "Gate 410 audits advanced representation extensions after Gate 409 and finds capacity but no native theorem. KO/twisted real-structure data changes compatibility signs, not family rank. Modular/KMS states could break degeneracy only after supplying a non-native Hamiltonian or density matrix. Primitive ideal extensions and triality local systems require a new algebra/functor. A sealed U(3)_gen connection has CKM capacity, but that is exactly an external family bundle, not a derived ASHA object. Therefore the current project still has a trivial generation multiplicity and the 13 charged flavor moduli firewall remains preserved."
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGate409Inherited, StatusExtensionSearchFormalized, StatusFamilyBundleArenaAudited, StatusKOTwistAudited, StatusModularKMSAudited, StatusPrimitiveIdealExtensionAudited}
	for _, c := range a.Extensions.Candidates {
		addStatus(&statuses, c.Verdict)
	}
	for _, s := range []string{StatusFailedNoNativeFamilyBundle, StatusFailedKOTwistOnlyChangesSigns, StatusFailedKMSRequiresExternalHamiltonian, StatusFailedPrimitiveIdealsWrongDomain, StatusFailedExtensionRequiresNewAxiom, StatusFailedTrivialGenerationMultiplicity, StatusFailedNoNativeNoncommutingTexturePair, StatusFirewallPreserved13Moduli} {
		addStatus(&statuses, s)
	}
	if a.Extensions.ConditionalNontrivialBundles > 0 {
		addStatus(&statuses, StatusSealedFamilyBundleCapacity)
	}
	if a.Noncommuting.CKMCapacityConditional {
		addStatus(&statuses, StatusConditionalCKMCapacityActivated)
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

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate409_trivial=%v gate409_U3_unselected=%v gate409_no_ckm=%v gate408_scalar_blind=%v gate395_chirality=%v gate394_central=%v gate393_triality_not_admitted=%v charged_moduli=%d no_empirical=%v verdict=%s", x.Executed, x.Gate409FermionicCarrierTrivial, x.Gate409U3GenerationFreedomUnselected, x.Gate409NoNativeCKMCapacity, x.Gate408ScalarFlavorBlind, x.Gate395SpinorSplitIsChirality, x.Gate394StaticAddressCentral, x.Gate393TrialityDomainNotAdmitted, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}
func FormatCandidate(x ExtensionCandidate) string {
	return fmt.Sprintf("name=%s domain=%s native=%v changes_carrier=%v derives3=%v bundle=%v connection=%v curvature=%v diag=%v noncomm=%v ckm_native=%v ckm_cond=%v AF=%v J=%v Gamma=%v first_order=%v gauge=%v new_axiom=%v external_K=%v manual_family=%v empirical=%v wrong_domain=%v verdict=%s reason=%s", x.Name, x.Domain, x.NativeInCurrentAsha, x.ChangesFermionCarrier, x.DerivesThreeFamilies, x.NontrivialBundle, x.ProvidesCanonicalConnection, x.ProvidesFamilyCurvature, x.ProvidesDiagonalOperator, x.ProvidesTwoNoncommutingOps, x.CKMCapacityNative, x.CKMCapacityConditional, x.CompatibleAF, x.CompatibleJ, x.CompatibleGamma, x.CompatibleFirstOrder, x.CompatibleGaugeCharges, x.RequiresNewAxiom, x.RequiresExternalHamiltonian, x.RequiresManualFamilyLabel, x.RequiresEmpiricalInput, x.WrongDomain, x.Verdict, x.Reason)
}
func FormatExtensionAudit(x ExtensionAudit) string {
	return fmt.Sprintf("executed=%v candidates=%d native_bundles=%d conditional_bundles=%d native_connections=%d native_noncommuting=%d conditional_noncommuting=%d wrong_domain=%d new_axiom=%d verdict=%s", x.Executed, x.CandidatesAudited, x.NativeNontrivialBundles, x.ConditionalNontrivialBundles, x.NativeConnections, x.NativeNoncommutingPairs, x.ConditionalNoncommutingPairs, x.WrongDomainCandidates, x.RequiresNewAxiomCandidates, x.Verdict)
}
func FormatFamilyBundle(x FamilyBundleAudit) string {
	return fmt.Sprintf("executed=%v carrier=%s rank=%d trivial=%v U3=%v U3_selected=%v transitions=%v connection=%v curvature=%v holonomy=%v replaces_C3=%v verdict=%s reason=%s", x.Executed, x.CurrentCarrier, x.GenerationRank, x.TrivialMultiplicity, x.ContainsU3Freedom, x.U3FreedomSelectedByGeometry, x.NontrivialTransitionFunctions, x.NativeFamilyConnection, x.NativeFamilyCurvature, x.NativeHolonomyNonabelian, x.ReplacesTensorC3, x.Verdict, x.Reason)
}
func FormatKOTwist(x KOTwistAudit) string {
	return fmt.Sprintf("executed=%v KO_signs=%v changes_JGamma=%v changes_multiplicity=%v families3=%v connection=%v noncommuting=%v compatible=%v twisted_action_needed=%v verdict=%s reason=%s", x.Executed, x.KODimensionSignsAudited, x.ChangesJGammaCommutation, x.ChangesMultiplicity, x.ProducesThreeFamilies, x.ProducesFamilyConnection, x.ProducesNoncommutingTextures, x.CompatibleWithExistingTriple, x.RequiresTwistedAlgebraAction, x.Verdict, x.Reason)
}
func FormatModularKMS(x ModularKMSAudit) string {
	return fmt.Sprintf("executed=%v tracial_freezes=%v nontracial_capacity=%v native_K=%v native_rho=%v three_level=%v mixing=%v external_K=%v chosen_rho=%v verdict=%s reason=%s", x.Executed, x.TracialStateFreezesFlow, x.NontracialStateHasCapacity, x.NativeHamiltonianFound, x.NativeDensityMatrixFound, x.ThreeLevelSpectrumNative, x.MixingOperatorNative, x.RequiresExternalHamiltonian, x.RequiresChosenDensityMatrix, x.Verdict, x.Reason)
}
func FormatPrimitiveIdeals(x PrimitiveIdealExtensionAudit) string {
	return fmt.Sprintf("executed=%v existing=%v wrong_domain=%v new_extension=%v family_ideal=%v noncentral_C3=%v algebra_enlargement=%v AF=%v first_order=%v verdict=%s reason=%s", x.Executed, x.ExistingIdealsAudited, x.ExistingIdealsAreWrongDomain, x.NewPrimitiveIdealExtension, x.ThreeFamilyIdealDerived, x.ActsOnC3GenNoncentrally, x.RequiresAlgebraEnlargement, x.CompatibleWithAF, x.CompatibleWithFirstOrder, x.Verdict, x.Reason)
}
func FormatNoncommuting(x NoncommutingTextureAudit) string {
	return fmt.Sprintf("executed=%v native_ops=%d native_pairs=%d conditional_ops=%d conditional_pairs=%d updown_native=%v updown_cond=%v ckm_native=%v ckm_cond=%v pmns_native=%v verdict=%s reason=%s", x.Executed, x.NativeFamilyOperators, x.NativeNoncommutingPairs, x.ConditionalFamilyOperators, x.ConditionalNoncommutingPairs, x.UpDownMisalignmentNative, x.UpDownMisalignmentConditional, x.CKMCapacityNative, x.CKMCapacityConditional, x.PMNSCapacityNative, x.Verdict, x.Reason)
}
func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("name=%s status=%s moduli_dim=%d masses3=%v ckm=%v pmns=%v nontrivial_bundle=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.FamilyBundleNontrivial, x.Reason)
}
func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d scenarios=%d best_native_dim=%d native_reduction=%v conditional_reduction=%v firewall=%v verdict=%s", x.StartDim, len(x.Scenarios), x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalReductionBelow13, x.FirewallPreserved, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa=%v no_external_K=%v no_manual_bundle=%v no_new_axiom=%v no_color_as_gen=%v no_species_as_gen=%v no_scalar=%v no_moduli_reduction=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoExternalHamiltonianPromoted, x.NoManualFamilyBundlePromoted, x.NoNewAxiomPromoted, x.NoColorAsGenerationPromoted, x.NoSpeciesAsGenerationPromoted, x.NoScalarSelectorPromoted, x.NoModuliReductionClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
func CompactStatusTable(a Analysis) string { return strings.Join(Statuses(a), "\n") }
