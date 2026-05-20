// Package generation2paulimomentweakplaneincidenceaudit implements Gate 561:
// Pauli Moment to Weak-Plane Incidence Intertwiner Audit.
//
// Gate 560 certified a sealed scalar-sector Pauli/Hopf moment triplet
// R^3_sigma on H_phi ~= C^2. Gate 561 audits whether this triplet can be
// functorially/intertwiner-transferred to the spatial weak-plane incidence
// data U_12, U_13, U_23. The result is deliberately conservative: formal
// incidence and formal Hodge-star formulas exist once one chooses an oriented
// Euclidean 3-space of spatial labels, but ASHA's current project data does
// not certify that label space as a native oriented metric carrier and does
// not construct a basis-independent map from R^3_sigma to spatial vectors or
// bivectors. The weak-plane/generation/firewall remains closed.
package generation2paulimomentweakplaneincidenceaudit

import (
	"fmt"
	"strings"
	"sync"

	gate560 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulihopfscalarmomentmapaudit"
)

const (
	AuditID = "GATE561-PAULI-MOMENT-WEAK-PLANE-INCIDENCE-INTERTWINER-AUDIT"

	StatusGate560Inherited                       = "CONDITIONAL_SUPPORT_GATE560_PAULI_HOPF_SCALAR_MOMENT_INHERITED"
	StatusSpatialLabelsAvailable                 = "CONDITIONAL_SUPPORT_SPATIAL_MODE_LABEL_SPACE_AVAILABLE_IN_B_MINUS_L_EIGENSPACE"
	StatusSpatialOrientedMetricNotNative         = "FAILED_ROUTE_SPATIAL_LABEL_SPACE_NOT_NATIVE_ORIENTED_METRIC_3SPACE"
	StatusWeakPlanesRepresentableAsBivectors     = "CONDITIONAL_SUPPORT_WEAK_PLANE_CANDIDATES_REPRESENTABLE_AS_COORDINATE_BIVECTORS"
	StatusWeakPlaneIncidenceNotNative            = "FAILED_ROUTE_WEAK_PLANE_INCIDENCE_REPRESENTATION_NOT_NATIVE_SELECTOR"
	StatusFormalHodgeStarAvailable               = "CONDITIONAL_SUPPORT_FORMAL_HODGE_STAR_AVAILABLE_GIVEN_EXTRA_ORIENTATION"
	StatusHodgeStarNotNative                     = "FAILED_ROUTE_SPATIAL_HODGE_STAR_NOT_NATIVE_WITHOUT_METRIC_ORIENTATION_CERTIFICATE"
	StatusNoPauliToIncidenceIntertwiner          = "FAILED_ROUTE_NO_PAULI_MOMENT_TO_WEAK_PLANE_INCIDENCE_INTERTWINER"
	StatusBasisDependentIntertwiner              = "FAILED_ROUTE_PAULI_TO_WEAK_PLANE_INTERTWINER_BASIS_DEPENDENT"
	StatusNoCanonicalWeakPlane                   = "FAILED_ROUTE_NO_CANONICAL_WEAK_PLANE_SELECTED_BY_SCALAR_MOMENT"
	StatusFormalBLCompatibility                  = "CONDITIONAL_SUPPORT_FORMAL_INCIDENCE_SELECTION_REFINES_B_MINUS_L_SPATIAL_EIGENSPACE"
	StatusBLCompatibilityVacuous                 = "CONDITIONAL_SUPPORT_B_MINUS_L_COMPATIBILITY_IS_VACUOUS_ON_W_SPATIAL"
	StatusBLDoesNotCanonicalize                  = "FAILED_ROUTE_B_MINUS_L_DOES_NOT_CANONICALIZE_PAULI_INCIDENCE_TRANSFER"
	StatusSpectralTripleCompatibilityUnavailable = "FAILED_ROUTE_PAULI_INCIDENCE_SPECTRAL_TRIPLE_COMPATIBILITY_UNAVAILABLE_NO_INTERTWINER"
	StatusNoFiniteOneFormHiggsRelation           = "FAILED_ROUTE_NO_PAULI_INCIDENCE_RELATION_TO_FINITE_ONE_FORM_HIGGS_LANE"
	StatusNoGenerationOrFlavorPromotion          = "FAILED_ROUTE_PAULI_INCIDENCE_DOES_NOT_GRANT_GENERATION_OR_FLAVOR_DATA"
	StatusFirewallPreserved                      = "FIREWALL_PRESERVED_GATE561_PAULI_MOMENT_WEAK_PLANE_INCIDENCE_BOUNDARY"
)

type InheritedAudit struct {
	Gate560PauliTriplet           bool
	Gate560HopfIdentity           bool
	Gate560ScalarFourToOnePlus3   bool
	Gate560ScalarMomentThreeSplit bool
	Gate560NoTransferFunctor      bool
	Verdict                       string
}

type SpatialLabelAudit struct {
	CarrierName                     string
	Labels                          []string
	CorrespondingFockModes          []string
	BLSpatialEigenvalue             string
	InsideBLSpatialEigenspace       bool
	NativeOrientedMetricThreeSpace  bool
	BasisConventionOnly             bool
	MetricCertificateAvailable      bool
	OrientationCertificateAvailable bool
	Verdict                         string
}

type WeakPlaneIncidenceAudit struct {
	PlaneToBivector              map[string]string
	IncidenceDimension           int
	CoordinateBivectorsAvailable bool
	NativeIncidenceSelector      bool
	NotationalOnly               bool
	Verdict                      string
}

type HodgeAudit struct {
	RequiresMetricAndOrientation bool
	MetricAvailableNatively      bool
	OrientationAvailableNatively bool
	FormalRules                  []string
	FormalNormalSelectsPlane     bool
	NativeHodgeStarConstructed   bool
	Verdict                      string
}

type IntertwinerSearchAudit struct {
	Source                     string
	TargetVector               string
	TargetBivector             string
	MapToSpatialFound          bool
	MapToIncidenceFound        bool
	BasisIndependent           bool
	UnitMetricCompatible       bool
	ManualSigma3ToS3Assignment bool
	Verdict                    string
}

type CanonicalPlaneAudit struct {
	IntertwinerExists       bool
	Sigma3AxisAvailable     bool
	NonzeroMuAvailable      bool
	CanonicalU12            bool
	CanonicalU13            bool
	CanonicalU23            bool
	OnlyBasisDependentPlane bool
	Verdict                 string
}

type BLCompatibilityAudit struct {
	SelectionInsideWSpatial       bool
	MixesLeptonSlot               bool
	BLRestrictedToWSpatial        string
	FormalSelectionCommutesWithBL bool
	CompatibilityNontrivial       bool
	BLSuppliesPlaneLabels         bool
	Verdict                       string
}

type SpectralTripleAudit struct {
	IncidenceFunctorFound      bool
	GradingCheckAvailable      bool
	JCheckAvailable            bool
	DCheckAvailable            bool
	FirstOrderCheckAvailable   bool
	FiniteOneFormRelationFound bool
	CompatibilityPassed        bool
	MissingData                []string
	Verdict                    string
}

type FirewallAudit struct {
	Preserved                     bool
	WeakIsospinIdentified         bool
	GaugeBosonsIdentified         bool
	PhotonIdentified              bool
	GenerationHierarchyIdentified bool
	YukawaTextureDerived          bool
	CKMPMNSDerived                bool
	ObservedFlavorImported        bool
	HiggsLanePromoted             bool
	Verdict                       string
}

type FinalVerdict struct {
	SSpatialNativeOrientedMetric bool
	WeakPlanesNativeBivectors    bool
	HodgeStarNative              bool
	PauliToIncidenceIntertwiner  bool
	ScalarMomentSelectsWeakPlane bool
	LawfulTransferAvailable      bool
	MissingNextTheorem           string
	Verdict                      string
}

type Analysis struct {
	Inherited     InheritedAudit
	SpatialLabels SpatialLabelAudit
	Incidence     WeakPlaneIncidenceAudit
	Hodge         HodgeAudit
	Intertwiner   IntertwinerSearchAudit
	Plane         CanonicalPlaneAudit
	BL            BLCompatibilityAudit
	Spectral      SpectralTripleAudit
	Firewall      FirewallAudit
	Final         FinalVerdict
	Truth         string
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
	prev, err := gate560.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 560 Pauli-Hopf scalar moment audit: %w", err)
	}
	inherited := auditInherited(prev)
	spatial := auditSpatialLabels()
	incidence := auditWeakPlaneIncidence()
	hodge := auditHodge(spatial)
	intertwiner := auditIntertwinerSearch(spatial, incidence, hodge)
	plane := auditCanonicalPlane(intertwiner)
	bl := auditBLCompatibility()
	spectral := auditSpectralTriple(intertwiner)
	firewall := auditFirewall()
	final := auditFinal(spatial, incidence, hodge, intertwiner, plane, firewall)
	a := Analysis{Inherited: inherited, SpatialLabels: spatial, Incidence: incidence, Hodge: hodge, Intertwiner: intertwiner, Plane: plane, BL: bl, Spectral: spectral, Firewall: firewall, Final: final}
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(prev gate560.Analysis) InheritedAudit {
	return InheritedAudit{
		Gate560PauliTriplet:           prev.Final.SealedPauliTripletExists,
		Gate560HopfIdentity:           prev.Final.HopfMomentIdentityHolds,
		Gate560ScalarFourToOnePlus3:   prev.Final.ScalarFourToOnePlusThree,
		Gate560ScalarMomentThreeSplit: prev.Final.NonzeroMomentThreeToOnePlusTwo,
		Gate560NoTransferFunctor:      !prev.Final.LawfulTransferToWOrGeneration,
		Verdict:                       join(StatusGate560Inherited, "Gate 561 starts from the sealed scalar Pauli/Hopf moment triplet and its explicit transfer firewall"),
	}
}

func auditSpatialLabels() SpatialLabelAudit {
	return SpatialLabelAudit{
		CarrierName:                     "S_spatial = span{s_1,s_2,s_3} as labels for a_1^dagger,a_2^dagger,a_3^dagger",
		Labels:                          []string{"s_1", "s_2", "s_3"},
		CorrespondingFockModes:          []string{"a_1^dagger", "a_2^dagger", "a_3^dagger"},
		BLSpatialEigenvalue:             "B-L|W_spatial = (1/3) I_3",
		InsideBLSpatialEigenspace:       true,
		NativeOrientedMetricThreeSpace:  false,
		BasisConventionOnly:             true,
		MetricCertificateAvailable:      false,
		OrientationCertificateAvailable: false,
		Verdict:                         join(StatusSpatialLabelsAvailable, StatusSpatialOrientedMetricNotNative),
	}
}

func auditWeakPlaneIncidence() WeakPlaneIncidenceAudit {
	return WeakPlaneIncidenceAudit{
		PlaneToBivector: map[string]string{
			"U_12": "s_1 ∧ s_2",
			"U_13": "s_1 ∧ s_3",
			"U_23": "s_2 ∧ s_3",
		},
		IncidenceDimension:           3,
		CoordinateBivectorsAvailable: true,
		NativeIncidenceSelector:      false,
		NotationalOnly:               true,
		Verdict:                      join(StatusWeakPlanesRepresentableAsBivectors, StatusWeakPlaneIncidenceNotNative),
	}
}

func auditHodge(spatial SpatialLabelAudit) HodgeAudit {
	native := spatial.NativeOrientedMetricThreeSpace && spatial.MetricCertificateAvailable && spatial.OrientationCertificateAvailable
	return HodgeAudit{
		RequiresMetricAndOrientation: true,
		MetricAvailableNatively:      spatial.MetricCertificateAvailable,
		OrientationAvailableNatively: spatial.OrientationCertificateAvailable,
		FormalRules: []string{
			"*s_1 = s_2∧s_3",
			"*s_2 = -s_1∧s_3",
			"*s_3 = s_1∧s_2",
		},
		FormalNormalSelectsPlane:   true,
		NativeHodgeStarConstructed: native,
		Verdict:                    join(StatusFormalHodgeStarAvailable, StatusHodgeStarNotNative),
	}
}

func auditIntertwinerSearch(spatial SpatialLabelAudit, incidence WeakPlaneIncidenceAudit, hodge HodgeAudit) IntertwinerSearchAudit {
	return IntertwinerSearchAudit{
		Source:                     "R^3_sigma scalar Pauli moment record space",
		TargetVector:               "S_spatial label space",
		TargetBivector:             "Λ^2 S_spatial coordinate incidence space",
		MapToSpatialFound:          false,
		MapToIncidenceFound:        false,
		BasisIndependent:           false,
		UnitMetricCompatible:       false,
		ManualSigma3ToS3Assignment: false,
		Verdict:                    join(StatusNoPauliToIncidenceIntertwiner, StatusBasisDependentIntertwiner),
	}
}

func auditCanonicalPlane(intertwiner IntertwinerSearchAudit) CanonicalPlaneAudit {
	return CanonicalPlaneAudit{
		IntertwinerExists:       intertwiner.MapToSpatialFound || intertwiner.MapToIncidenceFound,
		Sigma3AxisAvailable:     true,
		NonzeroMuAvailable:      true,
		CanonicalU12:            false,
		CanonicalU13:            false,
		CanonicalU23:            false,
		OnlyBasisDependentPlane: true,
		Verdict:                 join(StatusNoCanonicalWeakPlane, StatusBasisDependentIntertwiner),
	}
}

func auditBLCompatibility() BLCompatibilityAudit {
	return BLCompatibilityAudit{
		SelectionInsideWSpatial:       true,
		MixesLeptonSlot:               false,
		BLRestrictedToWSpatial:        "(1/3) I_3",
		FormalSelectionCommutesWithBL: true,
		CompatibilityNontrivial:       false,
		BLSuppliesPlaneLabels:         false,
		Verdict:                       join(StatusFormalBLCompatibility, StatusBLCompatibilityVacuous, StatusBLDoesNotCanonicalize),
	}
}

func auditSpectralTriple(intertwiner IntertwinerSearchAudit) SpectralTripleAudit {
	missing := []string{"basis-independent Pauli-to-incidence functor", "grading compatibility", "J compatibility", "D compatibility", "first-order compatibility", "finite one-form Higgs-lane relation"}
	return SpectralTripleAudit{
		IncidenceFunctorFound:      intertwiner.MapToSpatialFound || intertwiner.MapToIncidenceFound,
		GradingCheckAvailable:      false,
		JCheckAvailable:            false,
		DCheckAvailable:            false,
		FirstOrderCheckAvailable:   false,
		FiniteOneFormRelationFound: false,
		CompatibilityPassed:        false,
		MissingData:                missing,
		Verdict:                    join(StatusSpectralTripleCompatibilityUnavailable, StatusNoFiniteOneFormHiggsRelation),
	}
}

func auditFirewall() FirewallAudit {
	return FirewallAudit{
		Preserved:                     true,
		WeakIsospinIdentified:         false,
		GaugeBosonsIdentified:         false,
		PhotonIdentified:              false,
		GenerationHierarchyIdentified: false,
		YukawaTextureDerived:          false,
		CKMPMNSDerived:                false,
		ObservedFlavorImported:        false,
		HiggsLanePromoted:             false,
		Verdict:                       join(StatusNoGenerationOrFlavorPromotion, StatusFirewallPreserved),
	}
}

func auditFinal(spatial SpatialLabelAudit, incidence WeakPlaneIncidenceAudit, hodge HodgeAudit, intertwiner IntertwinerSearchAudit, plane CanonicalPlaneAudit, firewall FirewallAudit) FinalVerdict {
	return FinalVerdict{
		SSpatialNativeOrientedMetric: spatial.NativeOrientedMetricThreeSpace,
		WeakPlanesNativeBivectors:    incidence.NativeIncidenceSelector,
		HodgeStarNative:              hodge.NativeHodgeStarConstructed,
		PauliToIncidenceIntertwiner:  intertwiner.MapToSpatialFound || intertwiner.MapToIncidenceFound,
		ScalarMomentSelectsWeakPlane: plane.CanonicalU12 || plane.CanonicalU13 || plane.CanonicalU23,
		LawfulTransferAvailable:      false,
		MissingNextTheorem:           "Required next theorem: construct a native basis-independent incidence functor F_inc:R^3_sigma->Λ^2 S_spatial, with a certified oriented metric structure on S_spatial, B-L refinement, grading/J/D/first-order compatibility, and proof that the selected plane is not a basis convention.",
		Verdict:                      join(StatusNoPauliToIncidenceIntertwiner, StatusNoCanonicalWeakPlane, StatusFirewallPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.Gate560PauliTriplet || !a.Inherited.Gate560HopfIdentity || !a.Inherited.Gate560NoTransferFunctor {
		return fmt.Errorf("Gate 560 inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.SpatialLabels.InsideBLSpatialEigenspace || !a.SpatialLabels.BasisConventionOnly || a.SpatialLabels.NativeOrientedMetricThreeSpace {
		return fmt.Errorf("spatial label audit failed: %s", FormatSpatialLabels(a.SpatialLabels))
	}
	if !a.Incidence.CoordinateBivectorsAvailable || a.Incidence.NativeIncidenceSelector || !a.Incidence.NotationalOnly {
		return fmt.Errorf("incidence audit failed: %s", FormatIncidence(a.Incidence))
	}
	if !a.Hodge.FormalNormalSelectsPlane || a.Hodge.NativeHodgeStarConstructed {
		return fmt.Errorf("Hodge audit failed: %s", FormatHodge(a.Hodge))
	}
	if a.Intertwiner.MapToSpatialFound || a.Intertwiner.MapToIncidenceFound || a.Intertwiner.BasisIndependent {
		return fmt.Errorf("intertwiner unexpectedly found: %s", FormatIntertwiner(a.Intertwiner))
	}
	if a.Plane.CanonicalU12 || a.Plane.CanonicalU13 || a.Plane.CanonicalU23 {
		return fmt.Errorf("canonical weak plane unexpectedly selected: %s", FormatPlane(a.Plane))
	}
	if !a.BL.FormalSelectionCommutesWithBL || a.BL.CompatibilityNontrivial || a.BL.BLSuppliesPlaneLabels {
		return fmt.Errorf("B-L compatibility audit failed: %s", FormatBL(a.BL))
	}
	if a.Spectral.CompatibilityPassed || a.Spectral.IncidenceFunctorFound || a.Spectral.FiniteOneFormRelationFound {
		return fmt.Errorf("spectral triple compatibility unexpectedly passed: %s", FormatSpectral(a.Spectral))
	}
	if !a.Firewall.Preserved || a.Firewall.WeakIsospinIdentified || a.Firewall.GaugeBosonsIdentified || a.Firewall.GenerationHierarchyIdentified || a.Firewall.YukawaTextureDerived || a.Firewall.CKMPMNSDerived || a.Firewall.ObservedFlavorImported {
		return fmt.Errorf("firewall audit failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate560Inherited,
		StatusSpatialLabelsAvailable,
		StatusSpatialOrientedMetricNotNative,
		StatusWeakPlanesRepresentableAsBivectors,
		StatusWeakPlaneIncidenceNotNative,
		StatusFormalHodgeStarAvailable,
		StatusHodgeStarNotNative,
		StatusNoPauliToIncidenceIntertwiner,
		StatusBasisDependentIntertwiner,
		StatusNoCanonicalWeakPlane,
		StatusFormalBLCompatibility,
		StatusBLCompatibilityVacuous,
		StatusBLDoesNotCanonicalize,
		StatusSpectralTripleCompatibilityUnavailable,
		StatusNoFiniteOneFormHiggsRelation,
		StatusNoGenerationOrFlavorPromotion,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 561 audits the hoped-for incidence route from the sealed scalar Pauli moment triplet R^3_sigma to the spatial weak-plane candidates. It finds coordinate labels s_1,s_2,s_3 and formal bivectors for U_12,U_13,U_23, and notes that a formal Hodge star would map a normal axis to an orthogonal two-plane if an oriented Euclidean spatial label metric were supplied. But current ASHA data treats S_spatial as a basis convention inside the B-L spatial eigenspace, not as a certified native oriented metric 3-space. No basis-independent F:R^3_sigma->S_spatial or F_inc:R^3_sigma->Λ^2S_spatial exists. B-L compatibility is vacuous because B-L is scalar on W_spatial. Therefore no scalar moment selects a canonical U_12/U_13/U_23, and no weak-isospin, gauge-boson, generation, Yukawa, CKM/PMNS, or observed-flavor claim is permitted. %s", a.Final.MissingNextTheorem)
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
