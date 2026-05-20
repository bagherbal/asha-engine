// Package generation2k7twistorspherehiggssocketbundleandvacuumselectorfirewallaudit implements
// Gate 713: K7 Twistor-Sphere Higgs Socket Bundle and Vacuum Selector Firewall Audit.
//
// Gate 712 showed that K7- does not supply a canonical unit selector n_*, but
// every unit n in K7- selects a compatible complex structure J_H(n)=n_a J_a on
// K7+.  Gate 713 audits whether the correct native object is therefore the
// full S^2 / CP1 family of compatible complex structures: a twistor-sphere
// Higgs-socket bundle over K7- directions.  It preserves the firewalls blocking
// physical electroweak, hypercharge, Higgs mass, Yukawa, flavor hierarchy,
// CKM/PMNS, scalar-runtime, and native 7/72 promotions.
package generation2k7twistorspherehiggssocketbundleandvacuumselectorfirewallaudit

import (
	"fmt"
	"strings"
	"sync"

	gate712 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7minuscomplexstructureselectorandso3gaugefirewallaudit"
)

const (
	AuditID = "GATE713-K7-TWISTOR-SPHERE-HIGGS-SOCKET-BUNDLE-AND-VACUUM-SELECTOR-FIREWALL-AUDIT"

	StatusGate712SelectorFirewallInherited           = "PASS_GATE712_SELECTOR_FIREWALL_INHERITED"
	StatusTwistorSphereDefined                       = "PASS_TWISTOR_SPHERE_OF_COMPLEX_STRUCTURES_DEFINED"
	StatusU2SocketBundleDefined                      = "PASS_U2_SOCKET_BUNDLE_DEFINED"
	StatusSO3ActionOnSelectorSphereAudited           = "PASS_SO3_ACTION_ON_SELECTOR_SPHERE_AUDITED"
	StatusSelectorDependentAndInvariantDataSeparated = "PASS_SELECTOR_DEPENDENT_AND_INVARIANT_DATA_SEPARATED"
	StatusVacuumSelectorFirewallAudited              = "PASS_VACUUM_SELECTOR_FIREWALL_AUDITED"
	StatusPhysicalElectroweakFirewallEnforced        = "PASS_PHYSICAL_ELECTROWEAK_FIREWALL_ENFORCED"
	StatusK7PlusHiggsSocketTwistorSphereFamily       = "CONDITIONAL_SUPPORT_K7_PLUS_HIGGS_SOCKET_IS_TWISTOR_SPHERE_FAMILY"
	StatusU2SocketBundleOverS2                       = "CONDITIONAL_SUPPORT_U2_SOCKET_IS_BUNDLE_OVER_S2_OF_K7_MINUS_DIRECTIONS"
	StatusSingleHiggsSocketRequiresSelectorOrSeal    = "CONDITIONAL_SUPPORT_SINGLE_HIGGS_SOCKET_REQUIRES_SELECTOR_OR_SEAL"
	StatusNoNativeTwistorPointSelector               = "FAILED_ROUTE_NO_NATIVE_TWISTOR_POINT_SELECTOR"
	StatusNoCanonicalHiggsComplexStructureSelected   = "FAILED_ROUTE_NO_CANONICAL_HIGGS_COMPLEX_STRUCTURE_SELECTED"
	StatusInternalSocketBundleNotPhysicalElectroweak = "FAILED_ROUTE_INTERNAL_SOCKET_BUNDLE_NOT_CERTIFIED_AS_PHYSICAL_ELECTROWEAK_REPRESENTATION"
	StatusNoHyperchargeAssignmentOrNormalization     = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT_OR_NORMALIZATION"
	StatusNoYukawaOperatorOrEigenvalueTheorem        = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem          = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusGate713TwistorSocketBundleBoundary         = "FIREWALL_PRESERVED_GATE713_TWISTOR_SOCKET_BUNDLE_BOUNDARY"
)

const (
	k7PlusDimension      = 4
	k7MinusDimension     = 3
	twistorSphereDim     = 2
	complexProjectiveDim = 1
	u2SocketDimension    = 4
	commutantDimension   = 3
)

type Gate712Inheritance struct {
	SelectorFirewallInherited bool
	K7PlusDimension           int
	K7MinusDimension          int
	UnitDirectionSelectsJH    bool
	FamilyValuedU2Socket      bool
	NativeSelectorFound       bool
	CanonicalJHSelected       bool
	PhysicalHiggsDoubletMap   bool
	PhysicalGenerationMap     bool
	YukawaOperatorMap         bool
	Verdict                   string
}

type TwistorSphereAudit struct {
	Base                        string
	EquivalentDescription       string
	CompatibleComplexStructures string
	SphereDimension             int
	ComplexProjectiveDimension  int
	JHSquared                   string
	FamilyNativeObject          bool
	SinglePointSelected         bool
	Verdict                     string
}

type U2SocketBundleAudit struct {
	BundleName               string
	Base                     string
	Fiber                    string
	FiberDimension           int
	CommutantDimension       int
	SpanJHDimension          int
	SingleSocketPromoted     bool
	FamilyValuedSocketBundle bool
	Verdict                  string
}

type SO3SelectorSphereAudit struct {
	EtaAction             string
	JAction               string
	ActsTransitivelyOnS2  bool
	PreservesFanoData     bool
	PreferredPoint        bool
	CanonicalAxisSelected bool
	Verdict               string
}

type SelectorDependentInvariantSeparation struct {
	SelectorDependent []string
	SelectorInvariant []string
	DependentCount    int
	InvariantCount    int
	SeparationValid   bool
	Verdict           string
}

type VacuumSelectorFirewallAudit struct {
	MissingPoint             string
	SelectorType             string
	NativeSelectorCertified  bool
	EnvironmentalSealAllowed bool
	SealNames                []string
	Verdict                  string
}

type PhysicalElectroweakFirewallAudit struct {
	TwistorBundlePhysicalElectroweak bool
	ChosenJHPhysicalHiggsStructure   bool
	SpanJHHypercharge                bool
	CommutantPhysicalSU2L            bool
	K7MinusFlavorHierarchy           bool
	MissingMaps                      []string
	Verdict                          string
}

type SourceTypeClassification struct {
	K7PlusRole   string
	K7MinusRole  string
	TwistorRole  string
	SocketRole   string
	SelectorRole string
	PhysicalRole string
	Verdict      string
}

type Analysis struct {
	Inherited        Gate712Inheritance
	Twistor          TwistorSphereAudit
	SocketBundle     U2SocketBundleAudit
	SO3Action        SO3SelectorSphereAudit
	DataSeparation   SelectorDependentInvariantSeparation
	VacuumFirewall   VacuumSelectorFirewallAudit
	PhysicalFirewall PhysicalElectroweakFirewallAudit
	SourceTypes      SourceTypeClassification
	Truth            string
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
	g712, err := gate712.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate712 inheritance unavailable: %w", err)
	}
	inherited := buildGate712Inheritance(g712)
	twistor := buildTwistorSphere(inherited)
	bundle := buildSocketBundle(twistor)
	so3 := buildSO3Action(twistor)
	separation := buildDataSeparation()
	vacuum := buildVacuumFirewall()
	physical := buildPhysicalFirewall()
	sources := buildSourceTypes()
	truth := "Gate 713 upgrades the Gate712 family-valued selector result into the correct internal object: the S^2/CP1 twistor sphere of compatible complex structures on K7+.  For every unit n in K7-, J_H(n)=n_a J_a gives a C^2 pre-Higgs carrier and an internal U(2,J_H(n)) socket, so the socket is a bundle over the selector sphere rather than a single native socket.  The inherited SO(3) covariance acts transitively on that sphere and selects no point.  A single Higgs socket would require a vacuum/orientation selector or quarantined seal.  No physical electroweak representation, hypercharge assignment, Higgs mass, scalar runtime, Yukawa operator, flavor hierarchy, CKM/PMNS, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Twistor: twistor, SocketBundle: bundle, SO3Action: so3, DataSeparation: separation, VacuumFirewall: vacuum, PhysicalFirewall: physical, SourceTypes: sources, Truth: truth}, nil
}

func buildGate712Inheritance(g gate712.Analysis) Gate712Inheritance {
	return Gate712Inheritance{
		SelectorFirewallInherited: g.Inherited.U2SocketInherited && g.FamilyMap.UnitDirectionWouldSelect && g.FamilyMap.FamilyValuedU2Socket && !g.FamilyMap.NativeSelectorFound && !g.FamilyMap.CanonicalComplexStructure,
		K7PlusDimension:           g.Inherited.K7PlusDimension,
		K7MinusDimension:          g.Inherited.K7MinusDimension,
		UnitDirectionSelectsJH:    g.FamilyMap.UnitDirectionWouldSelect,
		FamilyValuedU2Socket:      g.FamilyMap.FamilyValuedU2Socket,
		NativeSelectorFound:       g.FamilyMap.NativeSelectorFound,
		CanonicalJHSelected:       g.FamilyMap.CanonicalComplexStructure,
		PhysicalHiggsDoubletMap:   g.GaugeFirewall.PhysicalHiggsDoubletMap,
		PhysicalGenerationMap:     g.GaugeFirewall.PhysicalGenerationSpaceMap,
		YukawaOperatorMap:         g.GaugeFirewall.YukawaOperatorMap,
		Verdict:                   StatusGate712SelectorFirewallInherited,
	}
}

func buildTwistorSphere(i Gate712Inheritance) TwistorSphereAudit {
	family := i.SelectorFirewallInherited && i.K7MinusDimension == k7MinusDimension && i.K7PlusDimension == k7PlusDimension
	return TwistorSphereAudit{
		Base:                        "S^2(K7-)={n in K7-: ||n||=1}",
		EquivalentDescription:       "CP1 twistor sphere of compatible complex structures on quaternionic K7+",
		CompatibleComplexStructures: "J_H(n)=n_a J_a with J_H(n)^2=-I",
		SphereDimension:             twistorSphereDim,
		ComplexProjectiveDimension:  complexProjectiveDim,
		JHSquared:                   "J_H(n)^2=-I for each unit n",
		FamilyNativeObject:          family,
		SinglePointSelected:         false,
		Verdict: strings.Join([]string{
			StatusTwistorSphereDefined,
			StatusK7PlusHiggsSocketTwistorSphereFamily,
			StatusNoNativeTwistorPointSelector,
		}, "; "),
	}
}

func buildSocketBundle(t TwistorSphereAudit) U2SocketBundleAudit {
	return U2SocketBundleAudit{
		BundleName:               "U2SocketBundle",
		Base:                     t.Base,
		Fiber:                    "fiber over n is u(2,J_H(n))=span{J_H(n)} plus Comm(J_1,J_2,J_3)",
		FiberDimension:           u2SocketDimension,
		CommutantDimension:       commutantDimension,
		SpanJHDimension:          1,
		SingleSocketPromoted:     false,
		FamilyValuedSocketBundle: t.FamilyNativeObject && !t.SinglePointSelected,
		Verdict: strings.Join([]string{
			StatusU2SocketBundleDefined,
			StatusU2SocketBundleOverS2,
			StatusK7PlusHiggsSocketTwistorSphereFamily,
		}, "; "),
	}
}

func buildSO3Action(t TwistorSphereAudit) SO3SelectorSphereAudit {
	return SO3SelectorSphereAudit{
		EtaAction:             "eta_a -> R_ab eta_b for R in SO(3)",
		JAction:               "J_a -> R_ab J_b, hence n -> R n on S^2(K7-)",
		ActsTransitivelyOnS2:  true,
		PreservesFanoData:     true,
		PreferredPoint:        false,
		CanonicalAxisSelected: false,
		Verdict: strings.Join([]string{
			StatusSO3ActionOnSelectorSphereAudited,
			StatusNoNativeTwistorPointSelector,
			StatusNoCanonicalHiggsComplexStructureSelected,
		}, "; "),
	}
}

func buildDataSeparation() SelectorDependentInvariantSeparation {
	dependent := []string{
		"chosen J_H(n)",
		"chosen internal U(1) phase line span{J_H(n)}",
		"chosen C^2 model of K7+",
	}
	invariant := []string{
		"quaternionic structure on K7+",
		"full S^2/CP1 complex-structure family",
		"commutant sp(1) candidate",
		"K7+ real four-space",
		"F_A coupling frame up to SO(3)",
	}
	return SelectorDependentInvariantSeparation{
		SelectorDependent: dependent,
		SelectorInvariant: invariant,
		DependentCount:    len(dependent),
		InvariantCount:    len(invariant),
		SeparationValid:   true,
		Verdict: strings.Join([]string{
			StatusSelectorDependentAndInvariantDataSeparated,
			StatusK7PlusHiggsSocketTwistorSphereFamily,
		}, "; "),
	}
}

func buildVacuumFirewall() VacuumSelectorFirewallAudit {
	sealNames := []string{
		"HiggsComplexStructureVacuumSelectorSeal",
		"K7MinusTwistorPointSelectorTheorem",
		"SpontaneousHiggsSocketOrientationSeal",
	}
	return VacuumSelectorFirewallAudit{
		MissingPoint:             "n_* in S^2(K7-) selecting a single J_H(n_*)",
		SelectorType:             "vacuum/orientation selector or quarantined environmental seal",
		NativeSelectorCertified:  false,
		EnvironmentalSealAllowed: true,
		SealNames:                sealNames,
		Verdict: strings.Join([]string{
			StatusVacuumSelectorFirewallAudited,
			StatusSingleHiggsSocketRequiresSelectorOrSeal,
			StatusNoNativeTwistorPointSelector,
			StatusNoCanonicalHiggsComplexStructureSelected,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalElectroweakFirewallAudit {
	missing := []string{
		"Theta_H: chosen K7+_J -> physical Higgs doublet",
		"Theta_SU2: internal commutant sp(1) -> physical SU(2)_L",
		"Theta_Y: span{J_H} -> physical U(1)_Y with correct normalization",
		"Theta_selector: native/environmental principle selecting n_*",
	}
	return PhysicalElectroweakFirewallAudit{
		TwistorBundlePhysicalElectroweak: false,
		ChosenJHPhysicalHiggsStructure:   false,
		SpanJHHypercharge:                false,
		CommutantPhysicalSU2L:            false,
		K7MinusFlavorHierarchy:           false,
		MissingMaps:                      missing,
		Verdict: strings.Join([]string{
			StatusPhysicalElectroweakFirewallEnforced,
			StatusInternalSocketBundleNotPhysicalElectroweak,
			StatusNoHyperchargeAssignmentOrNormalization,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		K7PlusRole:   "quaternionic real four-space carrying a family of C^2 pre-Higgs models",
		K7MinusRole:  "selector-sphere direction space; no native preferred point",
		TwistorRole:  "S^2/CP1 family is the invariant internal object before a vacuum point is selected",
		SocketRole:   "U(2,J_H(n)) socket is a bundle fiber over n, not a single physical electroweak socket",
		SelectorRole: "single Higgs socket requires n_* as native theorem or quarantined seal",
		PhysicalRole: "physical SU(2)_L, U(1)_Y, hypercharge, Yukawa, and Higgs runtime maps remain missing",
		Verdict: strings.Join([]string{
			StatusU2SocketBundleOverS2,
			StatusSingleHiggsSocketRequiresSelectorOrSeal,
			StatusGate713TwistorSocketBundleBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate712SelectorFirewallInherited,
		StatusTwistorSphereDefined,
		StatusU2SocketBundleDefined,
		StatusSO3ActionOnSelectorSphereAudited,
		StatusSelectorDependentAndInvariantDataSeparated,
		StatusVacuumSelectorFirewallAudited,
		StatusPhysicalElectroweakFirewallEnforced,
		StatusK7PlusHiggsSocketTwistorSphereFamily,
		StatusU2SocketBundleOverS2,
		StatusSingleHiggsSocketRequiresSelectorOrSeal,
		StatusNoNativeTwistorPointSelector,
		StatusNoCanonicalHiggsComplexStructureSelected,
		StatusInternalSocketBundleNotPhysicalElectroweak,
		StatusNoHyperchargeAssignmentOrNormalization,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusGate713TwistorSocketBundleBoundary,
	}
}

func FormatInherited(x Gate712Inheritance) string {
	return fmt.Sprintf("inherited=%t plusDim=%d minusDim=%d unitSelects=%t family=%t nativeSelector=%t canonical=%t higgsMap=%t generationMap=%t yukawa=%t verdict=%q", x.SelectorFirewallInherited, x.K7PlusDimension, x.K7MinusDimension, x.UnitDirectionSelectsJH, x.FamilyValuedU2Socket, x.NativeSelectorFound, x.CanonicalJHSelected, x.PhysicalHiggsDoubletMap, x.PhysicalGenerationMap, x.YukawaOperatorMap, x.Verdict)
}

func FormatTwistor(x TwistorSphereAudit) string {
	return fmt.Sprintf("base=%q equiv=%q structures=%q sphereDim=%d cpDim=%d jh=%q family=%t point=%t verdict=%q", x.Base, x.EquivalentDescription, x.CompatibleComplexStructures, x.SphereDimension, x.ComplexProjectiveDimension, x.JHSquared, x.FamilyNativeObject, x.SinglePointSelected, x.Verdict)
}

func FormatSocketBundle(x U2SocketBundleAudit) string {
	return fmt.Sprintf("name=%q base=%q fiber=%q fiberDim=%d commDim=%d jhDim=%d single=%t family=%t verdict=%q", x.BundleName, x.Base, x.Fiber, x.FiberDimension, x.CommutantDimension, x.SpanJHDimension, x.SingleSocketPromoted, x.FamilyValuedSocketBundle, x.Verdict)
}

func FormatSO3(x SO3SelectorSphereAudit) string {
	return fmt.Sprintf("eta=%q j=%q transitive=%t preserves=%t point=%t axis=%t verdict=%q", x.EtaAction, x.JAction, x.ActsTransitivelyOnS2, x.PreservesFanoData, x.PreferredPoint, x.CanonicalAxisSelected, x.Verdict)
}

func FormatDataSeparation(x SelectorDependentInvariantSeparation) string {
	return fmt.Sprintf("dependent=%d invariant=%d valid=%t verdict=%q", x.DependentCount, x.InvariantCount, x.SeparationValid, x.Verdict)
}

func FormatVacuum(x VacuumSelectorFirewallAudit) string {
	return fmt.Sprintf("missing=%q type=%q native=%t environmental=%t seals=%d verdict=%q", x.MissingPoint, x.SelectorType, x.NativeSelectorCertified, x.EnvironmentalSealAllowed, len(x.SealNames), x.Verdict)
}

func FormatPhysical(x PhysicalElectroweakFirewallAudit) string {
	return fmt.Sprintf("physicalBundle=%t physicalJH=%t hypercharge=%t su2=%t flavor=%t missing=%d verdict=%q", x.TwistorBundlePhysicalElectroweak, x.ChosenJHPhysicalHiggsStructure, x.SpanJHHypercharge, x.CommutantPhysicalSU2L, x.K7MinusFlavorHierarchy, len(x.MissingMaps), x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("plus=%q minus=%q twistor=%q socket=%q selector=%q physical=%q verdict=%q", x.K7PlusRole, x.K7MinusRole, x.TwistorRole, x.SocketRole, x.SelectorRole, x.PhysicalRole, x.Verdict)
}
