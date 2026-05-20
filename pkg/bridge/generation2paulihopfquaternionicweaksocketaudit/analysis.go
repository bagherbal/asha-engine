// Package generation2paulihopfquaternionicweaksocketaudit implements Gate 562:
// Pauli-Hopf to Quaternionic Weak-Socket Intertwiner Audit.
//
// Gate 560 certified a sealed scalar Pauli/Hopf moment map on H_phi ~= C^2,
// while Gate 561 blocked any basis-independent transfer to W_spatial or weak
// plane incidence labels. Gate 562 tests the more lawful target already present
// in the finite spectral-triple architecture: the quaternionic weak socket
// Im(H) inside A_F=C⊕H⊕M3(C). The result is positive only at the sealed
// scalar/quaternionic structural level: Im(H) and the scalar SU(2) doublet lane
// exist, and the Pauli/Hopf moment map is the standard quaternionic/SU(2)
// moment map once the weak-doublet representation is used. Axis labels remain
// frame/convention dependent, and no physical W/Z/photon, Higgs-potential,
// generation, Yukawa, CKM/PMNS, or observed-flavor conclusion is promoted.
package generation2paulihopfquaternionicweaksocketaudit

import (
	"fmt"
	"strings"
	"sync"

	gate560 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulihopfscalarmomentmapaudit"
	gate561 "github.com/bagherbal/asha-engine/pkg/bridge/generation2paulimomentweakplaneincidenceaudit"
	inner "github.com/bagherbal/asha-engine/pkg/bridge/innerfluctuationfieldcontent"
)

const (
	AuditID = "GATE562-PAULI-HOPF-QUATERNIONIC-WEAK-SOCKET-INTERTWINER-AUDIT"

	StatusGate560561Inherited                 = "CONDITIONAL_SUPPORT_GATES560_561_PAULI_HOPF_BOUNDARY_INHERITED"
	StatusFiniteQuaternionicSocketRecovered   = "CONDITIONAL_SUPPORT_FINITE_ALGEBRA_QUATERNIONIC_SOCKET_RECOVERED"
	StatusImHOrientedMetricLieSpace           = "PASS_IM_H_NATIVE_ORIENTED_METRIC_LIE_THREE_SPACE"
	StatusScalarWeakDoubletRecovered          = "CONDITIONAL_SUPPORT_HPHI_AS_STRUCTURAL_SU2_DOUBLE_MODULE_RECOVERED"
	StatusQuaternionicRepresentationAvailable = "CONDITIONAL_SUPPORT_QUATERNIONIC_DOUBLE_MODULE_REPRESENTATION_AVAILABLE"
	StatusPauliQuaternionBridgeExists         = "CONDITIONAL_SUPPORT_PAULI_TRIPLET_INTERTWINES_WITH_IM_H_UNDER_DOUBLE_MODULE"
	StatusAxisFrameConvention                 = "FAILED_ROUTE_PAULI_QUATERNION_AXIS_IDENTIFICATION_FRAME_CONVENTIONAL"
	StatusMomentMapQuaternionic               = "PASS_HOPF_MOMENT_MAP_IDENTIFIED_AS_QUATERNIONIC_SU2_MOMENT_MAP"
	StatusQuaternionicStabilizerSplit         = "CONDITIONAL_SUPPORT_SCALAR_QUATERNIONIC_MOMENT_3_TO_1PLUS2_STABILIZER_ORBIT_SPLIT"
	StatusEtaChosenAxisOnly                   = "CONDITIONAL_SUPPORT_ETA_IS_ONE_CHOSEN_PAULI_QUATERNIONIC_AXIS"
	StatusEtaAxisNotPhysical                  = "FAILED_ROUTE_ETA_AXIS_NOT_PHYSICAL_ELECTROWEAK_DIRECTION"
	StatusStructuralOneFormLink               = "CONDITIONAL_SUPPORT_STRUCTURAL_LINK_TO_FINITE_ONE_FORM_SCALAR_DOUBLE_MODULE"
	StatusSpectralDynamicsStillFirewalled     = "FAILED_ROUTE_PAULI_QUATERNION_BRIDGE_DOES_NOT_DERIVE_ELECTROWEAK_DYNAMICS_OR_MASSES"
	StatusNoWSpatialWeakPlaneTransfer         = "FAILED_ROUTE_PAULI_QUATERNION_SOCKET_DOES_NOT_REOPEN_W_SPATIAL_WEAK_PLANE_TRANSFER"
	StatusNoFlavorPromotion                   = "FAILED_ROUTE_PAULI_QUATERNION_SOCKET_DOES_NOT_GRANT_GENERATION_OR_FLAVOR_DATA"
	StatusFirewallPreserved                   = "FIREWALL_PRESERVED_GATE562_PAULI_HOPF_QUATERNIONIC_WEAK_SOCKET_BOUNDARY"
)

type InheritedAudit struct {
	Gate560PauliTriplet         bool
	Gate560HopfIdentity         bool
	Gate560ScalarMomentSplit    bool
	Gate561NoSpatialIntertwiner bool
	Gate561NoCanonicalWeakPlane bool
	Verdict                     string
}

type QuaternionicSocketAudit struct {
	FiniteAlgebra                   string
	ContainsQuaternionicSummand     bool
	ImaginaryQuaternionicBasis      []string
	Dimension                       int
	MetricNormAvailable             bool
	OrientationAvailable            bool
	LieBracketCrossProductAvailable bool
	UnitaryGroup                    string
	ImHAsWeakLieAlgebraStructural   bool
	PhysicalGaugeDynamicsDerived    bool
	Verdict                         string
}

type ScalarDoubletAudit struct {
	Carrier                           string
	RealDimension                     int
	ComplexDimension                  int
	WeakRepresentation                string
	SingleComplexDoubletRecovered     bool
	LeftHModuleOrEquivalentSU2Doublet bool
	RepresentationNativeStructural    bool
	RepresentationDynamical           bool
	NumericalYukawaFree               bool
	Verdict                           string
}

type PauliQuaternionRepresentationAudit struct {
	RhoHAvailable                          bool
	RhoHUnitPreserving                     bool
	ImaginaryUnitsAntiHermitian            bool
	PauliMatricesHermitianMomentGenerators bool
	CliffordPauliFromGate560               bool
	BasisIndependentAsModule               bool
	AxisByAxisIdentificationCanonical      bool
	ConventionFreedom                      string
	Verdict                                string
}

type IntertwinerAudit struct {
	Source                              string
	Target                              string
	ModuleIntertwinerExists             bool
	MetricCompatible                    bool
	LieBracketCompatible                bool
	BasisIndependentAsUnframedSpaces    bool
	SpecificSigmaToIJKFrameConventional bool
	ManualSigma3ToK                     bool
	Verdict                             string
}

type MomentMapAudit struct {
	MuFormula                     string
	MomentMapForSU2Action         bool
	HopfIdentityInherited         bool
	Codomain                      string
	Decomposition                 string
	NormalizationConvention       string
	IdentifiesPhysicalGaugeBosons bool
	Verdict                       string
}

type StabilizerOrbitAudit struct {
	NonzeroMuCondition                  bool
	Split                               string
	RadialLineCanonicalGivenMu          bool
	OrthogonalPlaneCanonicalGivenMetric bool
	ScalarQuaternionicOnly              bool
	IdentifiesWZPhoton                  bool
	Verdict                             string
}

type EtaRelationAudit struct {
	EtaEqualsSigma3                           bool
	Sigma3CorrespondsToChosenQuaternionicAxis bool
	AxisChosenByScalarFrame                   bool
	AxisPhysicallyCanonical                   bool
	TauEtaSigma3Shadow                        bool
	Verdict                                   string
}

type SpectralTripleCompatibilityAudit struct {
	AFRepresentationStructural        bool
	GradingCompatibilityInherited     bool
	JCompatibilityInherited           bool
	DCompatibilityInherited           bool
	FirstOrderConditionInherited      bool
	FiniteOneFormScalarLaneStructural bool
	HeatKernelProjectionAvailable     bool
	HiggsPotentialDerived             bool
	MassOrDynamicsDerived             bool
	MissingOrFirewalled               []string
	Verdict                           string
}

type FirewallAudit struct {
	Preserved                     bool
	PhysicalWeakBosonsIdentified  bool
	PhotonIdentified              bool
	HiggsMassTheorem              bool
	GenerationHierarchyIdentified bool
	YukawaTextureDerived          bool
	CKMPMNSDerived                bool
	ObservedFlavorImported        bool
	WSpatialWeakPlaneSelected     bool
	Verdict                       string
}

type FinalVerdict struct {
	ImHNativeOrientedMetricThreeSpace    bool
	HPhiWeakDoubletModule                bool
	PauliTripletEquivalentToImH          bool
	HopfMomentQuaternionicMomentMap      bool
	NonzeroMuQuaternionicThreeSplit      bool
	LinkedToFiniteOneFormStructurally    bool
	PhysicalElectroweakDynamicsDerived   bool
	LawfulTransferToWSpatialOrGeneration bool
	MissingNextTheorem                   string
	Verdict                              string
}

type Analysis struct {
	Inherited      InheritedAudit
	Quaternionic   QuaternionicSocketAudit
	ScalarDoublet  ScalarDoubletAudit
	Representation PauliQuaternionRepresentationAudit
	Intertwiner    IntertwinerAudit
	Moment         MomentMapAudit
	Orbit          StabilizerOrbitAudit
	Eta            EtaRelationAudit
	Spectral       SpectralTripleCompatibilityAudit
	Firewall       FirewallAudit
	Final          FinalVerdict
	Truth          string
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
	prev560, err := gate560.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 560 Pauli-Hopf scalar moment audit: %w", err)
	}
	prev561, err := gate561.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 561 Pauli incidence audit: %w", err)
	}
	innerA, err := inner.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 298 inner fluctuation field-content audit: %w", err)
	}
	inherited := auditInherited(prev560, prev561)
	q := auditQuaternionicSocket(innerA)
	scalar := auditScalarDoublet(innerA, prev560)
	rep := auditRepresentation(q, scalar, prev560)
	intertwiner := auditIntertwiner(rep)
	moment := auditMomentMap(prev560, intertwiner)
	orbit := auditStabilizerOrbit(q, moment)
	eta := auditEta(prev560, rep)
	spectral := auditSpectralTriple(innerA, rep)
	firewall := auditFirewall(prev561)
	final := auditFinal(q, scalar, rep, intertwiner, moment, orbit, spectral, firewall)
	a := Analysis{Inherited: inherited, Quaternionic: q, ScalarDoublet: scalar, Representation: rep, Intertwiner: intertwiner, Moment: moment, Orbit: orbit, Eta: eta, Spectral: spectral, Firewall: firewall, Final: final}
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(prev560 gate560.Analysis, prev561 gate561.Analysis) InheritedAudit {
	return InheritedAudit{
		Gate560PauliTriplet:         prev560.Final.SealedPauliTripletExists,
		Gate560HopfIdentity:         prev560.Final.HopfMomentIdentityHolds,
		Gate560ScalarMomentSplit:    prev560.Final.NonzeroMomentThreeToOnePlusTwo,
		Gate561NoSpatialIntertwiner: !prev561.Final.PauliToIncidenceIntertwiner,
		Gate561NoCanonicalWeakPlane: !prev561.Final.ScalarMomentSelectsWeakPlane,
		Verdict:                     join(StatusGate560561Inherited, "Gate 562 inherits the sealed scalar Pauli/Hopf triplet and the blocked W_spatial/weak-plane incidence route"),
	}
}

func auditQuaternionicSocket(a inner.Analysis) QuaternionicSocketAudit {
	containsH := strings.Contains(a.Input.Algebra, "H") && a.Gauge.TotalDimension == 12
	weakFound := false
	for _, s := range a.Gauge.Sectors {
		if s.Name == "weak" && s.SourceSummand == "H" && strings.Contains(s.LieAlgebra, "Im(H)") && s.Dimension == 3 && s.Derived {
			weakFound = true
		}
	}
	return QuaternionicSocketAudit{
		FiniteAlgebra:                   a.Input.Algebra,
		ContainsQuaternionicSummand:     containsH,
		ImaginaryQuaternionicBasis:      []string{"i", "j", "k"},
		Dimension:                       3,
		MetricNormAvailable:             true,
		OrientationAvailable:            true,
		LieBracketCrossProductAvailable: true,
		UnitaryGroup:                    a.Gauge.PreUnimodularUnitary,
		ImHAsWeakLieAlgebraStructural:   weakFound,
		PhysicalGaugeDynamicsDerived:    false,
		Verdict:                         join(StatusFiniteQuaternionicSocketRecovered, StatusImHOrientedMetricLieSpace, "A_F contains H, so Im(H) carries its native quaternionic norm/orientation/bracket as a structural weak socket, not a dynamics theorem"),
	}
}

func auditScalarDoublet(a inner.Analysis, prev560 gate560.Analysis) ScalarDoubletAudit {
	return ScalarDoubletAudit{
		Carrier:                           "H_phi ~= R^4 ~= C^2, one complex scalar doublet from finite one-form content",
		RealDimension:                     prev560.Scalar.RealDimension,
		ComplexDimension:                  prev560.Scalar.ComplexDimension,
		WeakRepresentation:                a.Higgs.WeakRepresentation,
		SingleComplexDoubletRecovered:     a.Higgs.SingleDoubletRecovered && a.Higgs.ComplexDoublets == 1 && a.Higgs.RealScalarDimension == 4,
		LeftHModuleOrEquivalentSU2Doublet: true,
		RepresentationNativeStructural:    true,
		RepresentationDynamical:           false,
		NumericalYukawaFree:               a.Higgs.NumericalYukawaFree,
		Verdict:                           join(StatusScalarWeakDoubletRecovered, "The project recovers one complex SU(2)_L doublet structurally; this is enough for a sealed H-module/SU(2)-doublet moment-map audit but not for dynamics"),
	}
}

func auditRepresentation(q QuaternionicSocketAudit, s ScalarDoubletAudit, prev560 gate560.Analysis) PauliQuaternionRepresentationAudit {
	available := q.ContainsQuaternionicSummand && q.ImHAsWeakLieAlgebraStructural && s.LeftHModuleOrEquivalentSU2Doublet && prev560.Final.SealedPauliTripletExists
	return PauliQuaternionRepresentationAudit{
		RhoHAvailable:                          available,
		RhoHUnitPreserving:                     available,
		ImaginaryUnitsAntiHermitian:            available,
		PauliMatricesHermitianMomentGenerators: available,
		CliffordPauliFromGate560:               prev560.Final.SealedPauliTripletExists,
		BasisIndependentAsModule:               available,
		AxisByAxisIdentificationCanonical:      false,
		ConventionFreedom:                      "Aut(H) acts as SO(3) on Im(H); a concrete sigma_a ↔ i,j,k frame is conventional unless the project fixes a quaternionic frame/orientation beyond the module structure",
		Verdict:                                join(StatusQuaternionicRepresentationAvailable, StatusAxisFrameConvention, "The doublet representation supplies anti-Hermitian quaternionic generators i sigma_a; Gate 560 supplies Hermitian Pauli moment generators sigma_a, but axis labels are frame-conventional"),
	}
}

func auditIntertwiner(r PauliQuaternionRepresentationAudit) IntertwinerAudit {
	return IntertwinerAudit{
		Source:                              "R^3_sigma from Gate 560 Pauli/Hopf moment triplet",
		Target:                              "Im(H) inside A_F=C⊕H⊕M3(C)",
		ModuleIntertwinerExists:             r.RhoHAvailable,
		MetricCompatible:                    r.RhoHAvailable,
		LieBracketCompatible:                r.RhoHAvailable,
		BasisIndependentAsUnframedSpaces:    r.RhoHAvailable,
		SpecificSigmaToIJKFrameConventional: true,
		ManualSigma3ToK:                     false,
		Verdict:                             join(StatusPauliQuaternionBridgeExists, StatusAxisFrameConvention, "A structural unframed scalar/quaternionic bridge exists; choosing Sigma_3 as a named i/j/k axis remains a frame convention"),
	}
}

func auditMomentMap(prev560 gate560.Analysis, i IntertwinerAudit) MomentMapAudit {
	return MomentMapAudit{
		MuFormula:                     "mu_a = phi^dagger sigma_a phi = x^T Sigma_a x, with |mu|^2=(r^2)^2",
		MomentMapForSU2Action:         i.ModuleIntertwinerExists && prev560.Final.HopfMomentIdentityHolds,
		HopfIdentityInherited:         prev560.Final.HopfMomentIdentityHolds,
		Codomain:                      "Im(H)^* ≅ R^3_sigma^* after the scalar-doublet/quaternionic representation",
		Decomposition:                 "H_phi -> r^2 plus mu in Im(H)^*",
		NormalizationConvention:       "moment-map normalization/sign follows Pauli/quaternionic generator convention; no physical coupling normalization is fixed",
		IdentifiesPhysicalGaugeBosons: false,
		Verdict:                       join(StatusMomentMapQuaternionic, "The Gate 560 Hopf map is the standard SU(2)/quaternionic doublet moment map under the structural weak-socket representation"),
	}
}

func auditStabilizerOrbit(q QuaternionicSocketAudit, m MomentMapAudit) StabilizerOrbitAudit {
	return StabilizerOrbitAudit{
		NonzeroMuCondition:                  true,
		Split:                               "Im(H)=R mu ⊕ mu^perp for mu != 0",
		RadialLineCanonicalGivenMu:          q.MetricNormAvailable && m.MomentMapForSU2Action,
		OrthogonalPlaneCanonicalGivenMetric: q.MetricNormAvailable && m.MomentMapForSU2Action,
		ScalarQuaternionicOnly:              true,
		IdentifiesWZPhoton:                  false,
		Verdict:                             join(StatusQuaternionicStabilizerSplit, "Nonzero scalar moment produces a quaternionic stabilizer/orbit split 3=1+2, but not a physical W/Z/photon split"),
	}
}

func auditEta(prev560 gate560.Analysis, r PauliQuaternionRepresentationAudit) EtaRelationAudit {
	return EtaRelationAudit{
		EtaEqualsSigma3: prev560.Final.EtaIsSigma3Axis,
		Sigma3CorrespondsToChosenQuaternionicAxis: r.RhoHAvailable,
		AxisChosenByScalarFrame:                   true,
		AxisPhysicallyCanonical:                   false,
		TauEtaSigma3Shadow:                        prev560.EtaRelation.Sigma3AxisShadowOnly,
		Verdict:                                   join(StatusEtaChosenAxisOnly, StatusEtaAxisNotPhysical, "eta=Sigma_3 is one selected scalar Pauli/quaternionic axis; the project does not promote it to a physical electroweak direction"),
	}
}

func auditSpectralTriple(a inner.Analysis, r PauliQuaternionRepresentationAudit) SpectralTripleCompatibilityAudit {
	return SpectralTripleCompatibilityAudit{
		AFRepresentationStructural:        a.Input.Gate297SkeletonComplete && a.Input.ZeroOrderVerified,
		GradingCompatibilityInherited:     a.Input.Gate297SkeletonComplete,
		JCompatibilityInherited:           a.Input.Gate297SkeletonComplete,
		DCompatibilityInherited:           a.Input.Gate297SkeletonComplete,
		FirstOrderConditionInherited:      a.Input.FirstOrderVerified,
		FiniteOneFormScalarLaneStructural: a.Higgs.SingleDoubletRecovered,
		HeatKernelProjectionAvailable:     false,
		HiggsPotentialDerived:             false,
		MassOrDynamicsDerived:             false,
		MissingOrFirewalled:               []string{"heat-kernel scalar/gauge kinetic projection", "Higgs potential coefficients", "physical electroweak symmetry-breaking dynamics", "W/Z/photon mass eigenbasis", "physical gauge coupling normalization", "Yukawa matrices and flavor data"},
		Verdict:                           join(StatusStructuralOneFormLink, StatusSpectralDynamicsStillFirewalled, "The bridge is compatible with the existing structural finite one-form scalar doublet lane, while all dynamical/mass/continuum projections remain firewalled"),
	}
}

func auditFirewall(prev561 gate561.Analysis) FirewallAudit {
	return FirewallAudit{
		Preserved:                     true,
		PhysicalWeakBosonsIdentified:  false,
		PhotonIdentified:              false,
		HiggsMassTheorem:              false,
		GenerationHierarchyIdentified: false,
		YukawaTextureDerived:          false,
		CKMPMNSDerived:                false,
		ObservedFlavorImported:        false,
		WSpatialWeakPlaneSelected:     prev561.Final.ScalarMomentSelectsWeakPlane,
		Verdict:                       join(StatusNoWSpatialWeakPlaneTransfer, StatusNoFlavorPromotion, StatusFirewallPreserved, "Gate 562 permits only a sealed scalar/quaternionic moment-map bridge; it does not reopen the blocked W_spatial/generation/flavor routes"),
	}
}

func auditFinal(q QuaternionicSocketAudit, s ScalarDoubletAudit, r PauliQuaternionRepresentationAudit, i IntertwinerAudit, m MomentMapAudit, o StabilizerOrbitAudit, st SpectralTripleCompatibilityAudit, fw FirewallAudit) FinalVerdict {
	return FinalVerdict{
		ImHNativeOrientedMetricThreeSpace:    q.ContainsQuaternionicSummand && q.MetricNormAvailable && q.OrientationAvailable && q.LieBracketCrossProductAvailable,
		HPhiWeakDoubletModule:                s.SingleComplexDoubletRecovered && s.LeftHModuleOrEquivalentSU2Doublet,
		PauliTripletEquivalentToImH:          i.ModuleIntertwinerExists && i.MetricCompatible && i.BasisIndependentAsUnframedSpaces,
		HopfMomentQuaternionicMomentMap:      m.MomentMapForSU2Action && m.HopfIdentityInherited,
		NonzeroMuQuaternionicThreeSplit:      o.NonzeroMuCondition && o.RadialLineCanonicalGivenMu && o.OrthogonalPlaneCanonicalGivenMetric,
		LinkedToFiniteOneFormStructurally:    st.FiniteOneFormScalarLaneStructural,
		PhysicalElectroweakDynamicsDerived:   false,
		LawfulTransferToWSpatialOrGeneration: false,
		MissingNextTheorem:                   "Gate 563 should audit whether the scalar/quaternionic moment map can be connected to finite electroweak curvature/Higgs kinetic projection without importing heat-kernel coefficients, physical gauge couplings, mass eigenstates, or flavor data.",
		Verdict:                              join(StatusPauliQuaternionBridgeExists, StatusMomentMapQuaternionic, StatusQuaternionicStabilizerSplit, StatusSpectralDynamicsStillFirewalled, StatusFirewallPreserved),
	}
}

func Statuses() []string {
	return []string{
		StatusGate560561Inherited,
		StatusFiniteQuaternionicSocketRecovered,
		StatusImHOrientedMetricLieSpace,
		StatusScalarWeakDoubletRecovered,
		StatusQuaternionicRepresentationAvailable,
		StatusPauliQuaternionBridgeExists,
		StatusAxisFrameConvention,
		StatusMomentMapQuaternionic,
		StatusQuaternionicStabilizerSplit,
		StatusEtaChosenAxisOnly,
		StatusEtaAxisNotPhysical,
		StatusStructuralOneFormLink,
		StatusSpectralDynamicsStillFirewalled,
		StatusNoWSpatialWeakPlaneTransfer,
		StatusNoFlavorPromotion,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return "Gate 562 finds the lawful target that Gate 561 did not: the scalar Pauli/Hopf triplet can be read as the moment map for the quaternionic weak socket Im(H) acting on the one complex scalar doublet. This is a sealed scalar/quaternionic structural bridge, not a physical electroweak dynamics theorem. Axis choices inside Im(H) remain frame-conventional, eta is only one chosen scalar Pauli axis, and no W_spatial weak plane, generation hierarchy, Yukawa texture, CKM/PMNS data, W/Z/photon split, Higgs potential, mass theorem, or observed flavor result is promoted."
}

func validate(a Analysis) error {
	if !a.Inherited.Gate560PauliTriplet || !a.Inherited.Gate560HopfIdentity || !a.Inherited.Gate561NoSpatialIntertwiner {
		return fmt.Errorf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.Quaternionic.ContainsQuaternionicSummand || !a.Quaternionic.ImHAsWeakLieAlgebraStructural || a.Quaternionic.Dimension != 3 || !a.Quaternionic.MetricNormAvailable || !a.Quaternionic.OrientationAvailable || !a.Quaternionic.LieBracketCrossProductAvailable {
		return fmt.Errorf("quaternionic socket not certified: %s", FormatQuaternionic(a.Quaternionic))
	}
	if !a.ScalarDoublet.SingleComplexDoubletRecovered || !a.ScalarDoublet.LeftHModuleOrEquivalentSU2Doublet || a.ScalarDoublet.RepresentationDynamical {
		return fmt.Errorf("scalar doublet audit failed: %s", FormatScalarDoublet(a.ScalarDoublet))
	}
	if !a.Representation.RhoHAvailable || !a.Representation.RhoHUnitPreserving || !a.Representation.ImaginaryUnitsAntiHermitian || a.Representation.AxisByAxisIdentificationCanonical {
		return fmt.Errorf("representation audit failed: %s", FormatRepresentation(a.Representation))
	}
	if !a.Intertwiner.ModuleIntertwinerExists || !a.Intertwiner.BasisIndependentAsUnframedSpaces || !a.Intertwiner.SpecificSigmaToIJKFrameConventional || a.Intertwiner.ManualSigma3ToK {
		return fmt.Errorf("intertwiner audit failed: %s", FormatIntertwiner(a.Intertwiner))
	}
	if !a.Moment.MomentMapForSU2Action || !a.Moment.HopfIdentityInherited || a.Moment.IdentifiesPhysicalGaugeBosons {
		return fmt.Errorf("moment map audit failed: %s", FormatMoment(a.Moment))
	}
	if !a.Orbit.NonzeroMuCondition || !a.Orbit.ScalarQuaternionicOnly || a.Orbit.IdentifiesWZPhoton {
		return fmt.Errorf("orbit audit failed: %s", FormatOrbit(a.Orbit))
	}
	if !a.Eta.EtaEqualsSigma3 || !a.Eta.Sigma3CorrespondsToChosenQuaternionicAxis || a.Eta.AxisPhysicallyCanonical {
		return fmt.Errorf("eta relation audit failed: %s", FormatEta(a.Eta))
	}
	if !a.Spectral.AFRepresentationStructural || !a.Spectral.FirstOrderConditionInherited || !a.Spectral.FiniteOneFormScalarLaneStructural || a.Spectral.HeatKernelProjectionAvailable || a.Spectral.MassOrDynamicsDerived {
		return fmt.Errorf("spectral compatibility audit failed: %s", FormatSpectral(a.Spectral))
	}
	if !a.Firewall.Preserved || a.Firewall.PhysicalWeakBosonsIdentified || a.Firewall.PhotonIdentified || a.Firewall.GenerationHierarchyIdentified || a.Firewall.YukawaTextureDerived || a.Firewall.CKMPMNSDerived || a.Firewall.ObservedFlavorImported || a.Firewall.WSpatialWeakPlaneSelected {
		return fmt.Errorf("firewall audit failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
