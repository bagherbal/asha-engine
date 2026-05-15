// Package generation2scalarcovariantintertwiner implements Gate 492:
// Scalar Covariant Derivative and Goldstone Intertwiner Audit.
//
// Gate 491 proved a bounded scalar-edge result: the Higgs carrier is a finite
// one-form edge module and the scalar kinetic trace is positive-semidefinite,
// with a 4=1+3 Goldstone count resonance. Gate 492 asks the stronger question:
// does the current ASHA registry already contain a native covariant derivative
// DΦ and a canonical protected-to-broken gauge-eating intertwiner?
//
// The verdict is deliberately exact. Existing lower gates provide a rigorous
// abstract DΦ template and a dimensionless W/Z/photon image diagnostic: three
// broken generator images are independent and Q_em annihilates the diagnostic
// vacuum. However, the template still depends on bridge-level choices: an
// abstract SU(2) doublet representation, a diagnostic lower-component vacuum,
// Euclidean scalar kinetic metric, and unselected gauge Hessian/couplings. This
// gate therefore promotes only the diagnostic socket, not a native W/Z mass or
// full electroweak symmetry breaking theorem.
package generation2scalarcovariantintertwiner

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2scalaredgestability"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarcovariant"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarsu2"
)

const (
	AuditID = "GATE492-SCALAR-COVARIANT-DERIVATIVE-GOLDSTONE-INTERTWINER-AUDIT"

	StatusGate491Inherited                 = "CONDITIONAL_SUPPORT_GATE491_SCALAR_EDGE_STABILITY_INHERITED"
	StatusAbstractDphiTemplateFound        = "CONDITIONAL_SUPPORT_ABSTRACT_DPHI_TEMPLATE_FOUND"
	StatusGoldstoneImageDiagnosticFound    = "CONDITIONAL_SUPPORT_GOLDSTONE_IMAGE_INTERTWINER_DIAGNOSTIC_FOUND"
	StatusPhotonExemptionDiagnostic        = "CONDITIONAL_SUPPORT_PHOTON_EXEMPTION_DIAGNOSTIC_CONFIRMED"
	StatusDimensionlessWZPhotonSignature   = "CONDITIONAL_SUPPORT_DIMENSIONLESS_WZ_PHOTON_SIGNATURE_CONFIRMED"
	StatusBridgeGaugeEatingSocketPreserved = "CONDITIONAL_SUPPORT_BRIDGE_GAUGE_EATING_SOCKET_PRESERVED"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_NO_WZ_HIGGS_OR_FLAVOR_DATA_IMPORTED"
	StatusWZMassWriteBlocked               = "FIREWALL_BLOCKED_WZ_MASS_NATIVE_REGISTRY_WRITE"

	StatusFailedNativeDphiNotDerived             = "FAILED_ROUTE_NATIVE_SCALAR_COVARIANT_DERIVATIVE_NOT_DERIVED"
	StatusFailedCanonicalIntertwinerNotDerived   = "FAILED_ROUTE_CANONICAL_PROTECTED_TO_BROKEN_INTERTWINER_NOT_DERIVED"
	StatusFailedScalarSU2NativeNotSelected       = "FAILED_ROUTE_FULL_SCALAR_SU2_ACTION_NOT_SELECTED_BY_FINITE_SCALAR_DATA"
	StatusFailedVacuumOrientationNotNative       = "FAILED_ROUTE_SCALAR_VACUUM_ORIENTATION_NOT_NATIVE"
	StatusFailedKineticMetricBridge              = "FAILED_ROUTE_SCALAR_KINETIC_METRIC_STILL_BRIDGE_LEVEL"
	StatusFailedGaugeHessianCouplingsNotSelected = "FAILED_ROUTE_GAUGE_HESSIAN_AND_COUPLINGS_NOT_ACTION_SELECTED"
	StatusFailedPhysicalMassMatrixNotDerived     = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_NOT_DERIVED"
	StatusFailedWeakAngleNotDerived              = "FAILED_ROUTE_WEAK_MIXING_ANGLE_NOT_DERIVED"

	StatusGate493RedirectDefined = "CONDITIONAL_SUPPORT_GATE493_FULL_ELECTROWEAK_CURVATURE_ACTION_REDIRECT_DEFINED"
)

const eps = 1e-8

type Inheritance struct {
	Executed                          bool
	Gate491ScalarEdgeStability        bool
	HiggsOneFormEdgeSupport           bool
	ScalarKineticPositiveSemidefinite bool
	GhostRouteBlocked                 bool
	GoldstoneCountResonance           bool
	NativeGaugeEatingPreviouslyOpen   bool
	NoMassFlavorDataImported          bool
	Verdict                           string
	Reason                            string
}

type ScalarRepresentationAudit struct {
	Executed                      bool
	ActiveRealDimension           int
	AbstractDoubletRepresentation bool
	SU2ClosureResidual            float64
	SkewResidual                  float64
	PairDegenerate                bool
	PairSplit                     float64
	FullSU2SelectedByScalarData   bool
	U1PairRotationSelected        bool
	CanonicalComplexStructure     bool
	CovariantDerivativeNative     bool
	GaugeEatingTheoremNative      bool
	Verdict                       string
	Reason                        string
}

type CovariantDerivativeAudit struct {
	Executed                                bool
	AbstractTemplateAvailable               bool
	GeneratorCount                          int
	ActiveRealDimension                     int
	GeneratorSkewResidual                   float64
	VacuumOrientationDiagnostic             bool
	VacuumOrientationNative                 bool
	EMAnnihilatesVacuumNorm                 float64
	MassMatrixRank                          int
	DimensionlessWZPhotonSignature          bool
	FiniteScalarKineticNormalizationDerived bool
	GaugeCouplingsDerived                   bool
	GaugeActionHessianDerived               bool
	PhysicalMassesDerived                   bool
	NativeDphiDerived                       bool
	Verdict                                 string
	Reason                                  string
}

type IntertwinerAudit struct {
	Executed                          bool
	ActiveRealDimension               int
	GaugeGeneratorCount               int
	BrokenGeneratorCount              int
	UnbrokenGeneratorCount            int
	BrokenImageRank                   int
	BrokenImagesIndependent           bool
	BrokenImageMinEigen               float64
	BrokenImageCondition              float64
	GoldstoneImageDiagnostic          bool
	GaugeEatingCountDiagnostic        bool
	EMNullNorm                        float64
	FiniteGaugeEatingTheoremDerived   bool
	GaugeBosonMassMatrixDerived       bool
	PhysicalMassesDerived             bool
	CanonicalProtectedToBrokenDerived bool
	Verdict                           string
	Reason                            string
}

type PhotonAudit struct {
	Executed                   bool
	QEMAnnihilatesVacuum       bool
	PhotonNullResidual         float64
	PhotonMassSquaredHat       float64
	UnbrokenGeneratorCount     int
	PhotonPhysicallyNormalized bool
	WeakMixingAngleDerived     bool
	FineStructureDerived       bool
	Verdict                    string
	Reason                     string
}

type Boundary struct {
	Executed                       bool
	DiagnosticSocketPromotable     bool
	NativeIntertwinerDerived       bool
	NativeDphiDerived              bool
	NativeVacuumOrientationDerived bool
	NativeKineticMetricDerived     bool
	NativeGaugeHessianDerived      bool
	NativeCouplingsDerived         bool
	PhysicalMassMatrixDerived      bool
	Verdict                        string
	Reason                         string
}

type Firewall struct {
	Executed                  bool
	ObservedWMassImported     bool
	ObservedZMassImported     bool
	ObservedHiggsMassImported bool
	FermiConstantImported     bool
	WeakAngleImported         bool
	YukawaImported            bool
	CKMPMNSImported           bool
	NativeWZMassWritten       bool
	NativeWeakAngleWritten    bool
	NativeHiggsMassWritten    bool
	Verdict                   string
	Reason                    string
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
	Inheritance    Inheritance
	Representation ScalarRepresentationAudit
	Dphi           CovariantDerivativeAudit
	Intertwiner    IntertwinerAudit
	Photon         PhotonAudit
	Boundary       Boundary
	Firewall       Firewall
	Registry       RegistryUpdate
	Next           NextStep
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
	g491, err := generation2scalaredgestability.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate491 scalar-edge audit: %w", err)
	}
	su2, err := scalarsu2.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar SU(2) audit: %w", err)
	}
	dphi, err := scalarcovariant.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit scalar covariant derivative template: %w", err)
	}
	ge, err := gaugeeating.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit gauge-eating diagnostic: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g491)
	a.Representation = buildRepresentationAudit(su2)
	a.Dphi = buildCovariantDerivativeAudit(dphi)
	a.Intertwiner = buildIntertwinerAudit(ge)
	a.Photon = buildPhotonAudit(dphi, ge)
	a.Boundary = buildBoundary(a.Representation, a.Dphi, a.Intertwiner, a.Photon)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistryUpdate(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g491 generation2scalaredgestability.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate491ScalarEdgeStability:        g491.Boundary.EdgeSupportNative && g491.Boundary.KineticSemidefiniteNative,
		HiggsOneFormEdgeSupport:           g491.Support.HiggsIsFiniteOneForm && g491.Support.EdgeMeasureSelected,
		ScalarKineticPositiveSemidefinite: g491.Kinetic.PositiveSemidefinite,
		GhostRouteBlocked:                 g491.Kinetic.GhostRiskEliminated,
		GoldstoneCountResonance:           g491.Goldstone.CountResonance,
		NativeGaugeEatingPreviouslyOpen:   !g491.Goldstone.GaugeEatingTheoremDerived && !g491.Goldstone.CovariantDerivativeDerived,
		NoMassFlavorDataImported:          !g491.Firewall.ObservedMassesImported && !g491.Firewall.ObservedHiggsMassImported && !g491.Firewall.ObservedCKMImported && !g491.Firewall.ObservedPMNSImported,
		Verdict:                           StatusGate491Inherited,
		Reason:                            "Gate491 supplies the finite one-form scalar carrier, positive-semidefinite kinetic trace, and 4=1+3 count resonance, while explicitly leaving the covariant derivative and gauge-eating map open.",
	}
}

func buildRepresentationAudit(s scalarsu2.Analysis) ScalarRepresentationAudit {
	return ScalarRepresentationAudit{
		Executed:                      true,
		ActiveRealDimension:           s.ActiveRealDimension,
		AbstractDoubletRepresentation: s.AbstractDoubletRepresentation,
		SU2ClosureResidual:            s.SU2ClosureResidual,
		SkewResidual:                  s.SkewResidual,
		PairDegenerate:                s.PairDegenerate,
		PairSplit:                     s.PairSplit,
		FullSU2SelectedByScalarData:   s.FullSU2SelectedByScalarData,
		U1PairRotationSelected:        s.U1PairRotationSelected,
		CanonicalComplexStructure:     s.CanonicalComplexStructure,
		CovariantDerivativeNative:     s.CovariantDerivativeDerived,
		GaugeEatingTheoremNative:      s.GaugeEatingTheoremDerived,
		Verdict:                       strings.Join([]string{StatusAbstractDphiTemplateFound, StatusFailedScalarSU2NativeNotSelected}, "; "),
		Reason:                        "the four-real scalar frame supports the realification of a complex SU(2) doublet, but the full SU(2) action is not selected by the finite scalar response itself and no canonical complex structure is yet derived",
	}
}

func buildCovariantDerivativeAudit(sc scalarcovariant.Analysis) CovariantDerivativeAudit {
	return CovariantDerivativeAudit{
		Executed:                                true,
		AbstractTemplateAvailable:               sc.AbstractCovariantDerivativeTemplate,
		GeneratorCount:                          4,
		ActiveRealDimension:                     sc.ActiveRealDimension,
		GeneratorSkewResidual:                   sc.GeneratorSkewResidual,
		VacuumOrientationDiagnostic:             len(sc.VacuumVector) == 4 && sc.VacuumRadius > eps,
		VacuumOrientationNative:                 sc.VacuumOrientationChosen,
		EMAnnihilatesVacuumNorm:                 sc.EMAnnihilatesVacuumNorm,
		MassMatrixRank:                          sc.MassMatrixRank,
		DimensionlessWZPhotonSignature:          sc.DimensionlessWZPhotonSignature,
		FiniteScalarKineticNormalizationDerived: sc.FiniteScalarKineticNormalizationDerived,
		GaugeCouplingsDerived:                   sc.GaugeCouplingsDerived,
		GaugeActionHessianDerived:               sc.GaugeActionHessianDerived,
		PhysicalMassesDerived:                   sc.PhysicalMassesDerived,
		NativeDphiDerived:                       false,
		Verdict:                                 strings.Join([]string{StatusAbstractDphiTemplateFound, StatusDimensionlessWZPhotonSignature, StatusFailedNativeDphiNotDerived, StatusFailedVacuumOrientationNotNative, StatusFailedGaugeHessianCouplingsNotSelected}, "; "),
		Reason:                                  "the existing scalar-covariant package constructs a dimensionless DΦ template over {T1,T2,T3,Yφ} and obtains the W/Z/photon rank signature, but the vacuum orientation, kinetic normalization, gauge couplings, and native finite-action origin are not selected",
	}
}

func buildIntertwinerAudit(ge gaugeeating.Analysis) IntertwinerAudit {
	return IntertwinerAudit{
		Executed:                          true,
		ActiveRealDimension:               ge.ActiveRealDimension,
		GaugeGeneratorCount:               ge.GaugeGeneratorCount,
		BrokenGeneratorCount:              ge.BrokenGeneratorCount,
		UnbrokenGeneratorCount:            ge.UnbrokenGeneratorCount,
		BrokenImageRank:                   ge.BrokenImageRank,
		BrokenImagesIndependent:           ge.BrokenImagesIndependent,
		BrokenImageMinEigen:               ge.BrokenImageMinEigen,
		BrokenImageCondition:              ge.BrokenImageCondition,
		GoldstoneImageDiagnostic:          ge.GoldstoneImageTheoremDiagnostic,
		GaugeEatingCountDiagnostic:        ge.GaugeEatingCountDiagnostic,
		EMNullNorm:                        ge.EMNullNorm,
		FiniteGaugeEatingTheoremDerived:   ge.FiniteGaugeEatingTheoremDerived,
		GaugeBosonMassMatrixDerived:       ge.GaugeBosonMassMatrixDerived,
		PhysicalMassesDerived:             ge.PhysicalMassesDerived,
		CanonicalProtectedToBrokenDerived: false,
		Verdict:                           strings.Join([]string{StatusGoldstoneImageDiagnosticFound, StatusBridgeGaugeEatingSocketPreserved, StatusFailedCanonicalIntertwinerNotDerived, StatusFailedKineticMetricBridge}, "; "),
		Reason:                            "the broken generator image map has rank three and Qem annihilates the diagnostic vacuum, so the gauge-eating socket is real at bridge level; it is not yet a canonical protected-contact to broken-gauge intertwiner selected by a finite action",
	}
}

func buildPhotonAudit(sc scalarcovariant.Analysis, ge gaugeeating.Analysis) PhotonAudit {
	qNull := sc.EMAnnihilatesVacuumNorm < eps && ge.EMNullNorm < eps
	return PhotonAudit{
		Executed:                   true,
		QEMAnnihilatesVacuum:       qNull,
		PhotonNullResidual:         sc.PhotonNullResidual,
		PhotonMassSquaredHat:       sc.PhotonMassSquaredHat,
		UnbrokenGeneratorCount:     ge.UnbrokenGeneratorCount,
		PhotonPhysicallyNormalized: false,
		WeakMixingAngleDerived:     false,
		FineStructureDerived:       false,
		Verdict:                    strings.Join([]string{StatusPhotonExemptionDiagnostic, StatusFailedWeakAngleNotDerived}, "; "),
		Reason:                     "Qem=T3+Yφ annihilates the diagnostic scalar vacuum and produces a null photon direction in the dimensionless template; physical photon normalization, theta_W, and alpha_em remain sealed until the gauge Hessian/couplings are selected",
	}
}

func buildBoundary(rep ScalarRepresentationAudit, d CovariantDerivativeAudit, i IntertwinerAudit, p PhotonAudit) Boundary {
	socket := d.AbstractTemplateAvailable && d.DimensionlessWZPhotonSignature && i.GoldstoneImageDiagnostic && i.GaugeEatingCountDiagnostic && p.QEMAnnihilatesVacuum
	return Boundary{
		Executed:                       true,
		DiagnosticSocketPromotable:     socket,
		NativeIntertwinerDerived:       false,
		NativeDphiDerived:              false,
		NativeVacuumOrientationDerived: false,
		NativeKineticMetricDerived:     false,
		NativeGaugeHessianDerived:      false,
		NativeCouplingsDerived:         false,
		PhysicalMassMatrixDerived:      false,
		Verdict:                        strings.Join([]string{StatusBridgeGaugeEatingSocketPreserved, StatusFailedNativeDphiNotDerived, StatusFailedCanonicalIntertwinerNotDerived, StatusFailedPhysicalMassMatrixNotDerived}, "; "),
		Reason:                         fmt.Sprintf("bridge socket=%t: abstract SU(2) doublet=%t, rank(DΦ_broken)=%d, photon-null=%t; native promotion is blocked by missing finite-action DΦ, canonical scalar complex/vacuum selection, kinetic metric, and gauge Hessian/couplings", socket, rep.AbstractDoubletRepresentation, i.BrokenImageRank, p.QEMAnnihilatesVacuum),
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                  true,
		ObservedWMassImported:     false,
		ObservedZMassImported:     false,
		ObservedHiggsMassImported: false,
		FermiConstantImported:     false,
		WeakAngleImported:         false,
		YukawaImported:            false,
		CKMPMNSImported:           false,
		NativeWZMassWritten:       false,
		NativeWeakAngleWritten:    false,
		NativeHiggsMassWritten:    false,
		Verdict:                   strings.Join([]string{StatusFirewallPreserved, StatusWZMassWriteBlocked}, "; "),
		Reason:                    "Gate492 imports no W/Z masses, Higgs pole mass, Fermi constant, weak mixing angle, Yukawa data, CKM, or PMNS data, and writes no native mass or electroweak-angle prediction.",
	}
}

func buildRegistryUpdate(_ Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Gate491 finite scalar one-form support and positive-semidefinite kinetic trace remain native bounded results",
			"the unbroken electromagnetic generator is structurally identified as Qem=T3+Yφ inside the dimensionless scalar template",
		},
		BridgeEntries: []string{
			"an abstract finite scalar covariant-derivative template DΦ over {T1,T2,T3,Yφ} exists on the four-real scalar frame",
			"the diagnostic broken-image map has rank three and realizes the 4=1+3 Goldstone count socket",
			"the dimensionless template has two charged directions, one neutral massive direction, and one photon-null direction",
		},
		EnvironmentalEntries: []string{
			"physical W/Z masses, Higgs pole mass, Fermi constant, weak mixing angle, fine-structure constant, Yukawa amplitudes, CKM, and PMNS remain outside the native registry",
		},
		FailedRoutes: []string{
			StatusFailedNativeDphiNotDerived,
			StatusFailedCanonicalIntertwinerNotDerived,
			StatusFailedScalarSU2NativeNotSelected,
			StatusFailedVacuumOrientationNotNative,
			StatusFailedKineticMetricBridge,
			StatusFailedGaugeHessianCouplingsNotSelected,
			StatusFailedPhysicalMassMatrixNotDerived,
		},
		OpenTheorems: []string{
			"derive scalar SU(2)_L action from the finite scalar/contact module instead of installing an abstract doublet representation",
			"derive a canonical scalar complex/quaternionic structure and vacuum orientation from finite dynamics",
			"derive a finite electroweak curvature/action whose second variation selects the gauge Hessian and couplings",
			"only after those are native, revisit physical W/Z masses and electroweak mixing",
		},
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        493,
		Title:       "Full Electroweak Curvature Action and Gauge Hessian Selection Audit",
		Reason:      "Gate492 proves the DΦ/gauge-eating socket only at bridge-diagnostic level; native promotion requires a finite action for the full electroweak connection, because the broken sector alone is not Lie-closed.",
		PrimaryTask: "construct or reject a native finite field-strength/curvature action for {T1,T2,Z,Q} and test whether its second variation selects the gauge Hessian without importing W/Z masses, theta_W, or continuum couplings",
	}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate491ScalarEdgeStability || !a.Inheritance.HiggsOneFormEdgeSupport || !a.Inheritance.ScalarKineticPositiveSemidefinite || !a.Inheritance.GhostRouteBlocked || !a.Inheritance.GoldstoneCountResonance || !a.Inheritance.NativeGaugeEatingPreviouslyOpen || !a.Inheritance.NoMassFlavorDataImported {
		return fmt.Errorf("Gate492 inheritance invalid: %+v", a.Inheritance)
	}
	if !a.Representation.Executed || a.Representation.ActiveRealDimension != 4 || !a.Representation.AbstractDoubletRepresentation || a.Representation.SU2ClosureResidual > eps || a.Representation.SkewResidual > eps || !a.Representation.PairDegenerate || a.Representation.PairSplit <= eps || a.Representation.FullSU2SelectedByScalarData || !a.Representation.U1PairRotationSelected || a.Representation.CanonicalComplexStructure || a.Representation.CovariantDerivativeNative || a.Representation.GaugeEatingTheoremNative {
		return fmt.Errorf("Gate492 representation audit invalid: %+v", a.Representation)
	}
	if !a.Dphi.Executed || !a.Dphi.AbstractTemplateAvailable || a.Dphi.GeneratorCount != 4 || a.Dphi.ActiveRealDimension != 4 || a.Dphi.GeneratorSkewResidual > eps || !a.Dphi.VacuumOrientationDiagnostic || a.Dphi.VacuumOrientationNative || a.Dphi.EMAnnihilatesVacuumNorm > eps || a.Dphi.MassMatrixRank != 3 || !a.Dphi.DimensionlessWZPhotonSignature || a.Dphi.FiniteScalarKineticNormalizationDerived || a.Dphi.GaugeCouplingsDerived || a.Dphi.PhysicalMassesDerived || a.Dphi.NativeDphiDerived {
		return fmt.Errorf("Gate492 Dphi audit invalid: %+v", a.Dphi)
	}
	if !a.Intertwiner.Executed || a.Intertwiner.ActiveRealDimension != 4 || a.Intertwiner.GaugeGeneratorCount != 4 || a.Intertwiner.BrokenGeneratorCount != 3 || a.Intertwiner.UnbrokenGeneratorCount != 1 || a.Intertwiner.BrokenImageRank != 3 || !a.Intertwiner.BrokenImagesIndependent || a.Intertwiner.BrokenImageMinEigen <= eps || !a.Intertwiner.GoldstoneImageDiagnostic || !a.Intertwiner.GaugeEatingCountDiagnostic || a.Intertwiner.EMNullNorm > eps || a.Intertwiner.FiniteGaugeEatingTheoremDerived || a.Intertwiner.GaugeBosonMassMatrixDerived || a.Intertwiner.PhysicalMassesDerived || a.Intertwiner.CanonicalProtectedToBrokenDerived {
		return fmt.Errorf("Gate492 intertwiner audit invalid: %+v", a.Intertwiner)
	}
	if !a.Photon.Executed || !a.Photon.QEMAnnihilatesVacuum || math.Abs(a.Photon.PhotonMassSquaredHat) > eps || a.Photon.PhotonNullResidual > eps || a.Photon.UnbrokenGeneratorCount != 1 || a.Photon.PhotonPhysicallyNormalized || a.Photon.WeakMixingAngleDerived || a.Photon.FineStructureDerived {
		return fmt.Errorf("Gate492 photon audit invalid: %+v", a.Photon)
	}
	if !a.Boundary.Executed || !a.Boundary.DiagnosticSocketPromotable || a.Boundary.NativeIntertwinerDerived || a.Boundary.NativeDphiDerived || a.Boundary.NativeVacuumOrientationDerived || a.Boundary.NativeKineticMetricDerived || a.Boundary.NativeGaugeHessianDerived || a.Boundary.NativeCouplingsDerived || a.Boundary.PhysicalMassMatrixDerived {
		return fmt.Errorf("Gate492 boundary invalid: %+v", a.Boundary)
	}
	if !a.Firewall.Executed || a.Firewall.ObservedWMassImported || a.Firewall.ObservedZMassImported || a.Firewall.ObservedHiggsMassImported || a.Firewall.FermiConstantImported || a.Firewall.WeakAngleImported || a.Firewall.YukawaImported || a.Firewall.CKMPMNSImported || a.Firewall.NativeWZMassWritten || a.Firewall.NativeWeakAngleWritten || a.Firewall.NativeHiggsMassWritten {
		return fmt.Errorf("Gate492 firewall invalid: %+v", a.Firewall)
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate492 finds the electroweak gauge-eating mechanism as a precise bridge diagnostic, not yet as a native theorem: the abstract DΦ template on the four-real scalar frame maps three broken directions into independent Goldstone images and leaves Qem photon-null, with mass-matrix rank %d. Native promotion is blocked until ASHA derives the scalar SU(2) action, canonical complex/vacuum orientation, scalar kinetic metric, and full electroweak gauge Hessian/couplings from a finite action.", a.Dphi.MassMatrixRank)
}

func FormatInheritance(i Inheritance) string {
	return fmt.Sprintf("%s: Gate491=%t oneform=%t kinetic_semidef=%t ghost_blocked=%t count_resonance=%t Dphi_open=%t observed_data=%t; %s", i.Verdict, i.Gate491ScalarEdgeStability, i.HiggsOneFormEdgeSupport, i.ScalarKineticPositiveSemidefinite, i.GhostRouteBlocked, i.GoldstoneCountResonance, i.NativeGaugeEatingPreviouslyOpen, !i.NoMassFlavorDataImported, i.Reason)
}

func FormatRepresentation(r ScalarRepresentationAudit) string {
	return fmt.Sprintf("%s: active=%d abstract_doublet=%t su2_residual=%.3e skew=%.3e pair_degenerate=%t pair_split=%.10f full_SU2_selected=%t U1_pair_selected=%t canonical_complex=%t native_Dphi=%t native_eating=%t; %s", r.Verdict, r.ActiveRealDimension, r.AbstractDoubletRepresentation, r.SU2ClosureResidual, r.SkewResidual, r.PairDegenerate, r.PairSplit, r.FullSU2SelectedByScalarData, r.U1PairRotationSelected, r.CanonicalComplexStructure, r.CovariantDerivativeNative, r.GaugeEatingTheoremNative, r.Reason)
}

func FormatDphi(d CovariantDerivativeAudit) string {
	return fmt.Sprintf("%s: template=%t generators=%d active=%d skew=%.3e vacuum_diagnostic=%t vacuum_native=%t Qem_phi0=%.3e rank=%d WZ_photon=%t ZH_native=%t couplings=%t hessian=%t masses=%t native_Dphi=%t; %s", d.Verdict, d.AbstractTemplateAvailable, d.GeneratorCount, d.ActiveRealDimension, d.GeneratorSkewResidual, d.VacuumOrientationDiagnostic, d.VacuumOrientationNative, d.EMAnnihilatesVacuumNorm, d.MassMatrixRank, d.DimensionlessWZPhotonSignature, d.FiniteScalarKineticNormalizationDerived, d.GaugeCouplingsDerived, d.GaugeActionHessianDerived, d.PhysicalMassesDerived, d.NativeDphiDerived, d.Reason)
}

func FormatIntertwiner(i IntertwinerAudit) string {
	return fmt.Sprintf("%s: active=%d gauge=%d broken=%d unbroken=%d rank=%d independent=%t min_eigen=%.10f condition=%.10f image_diag=%t count_diag=%t EM_null=%.3e finite_theorem=%t mass_matrix=%t masses=%t canonical_map=%t; %s", i.Verdict, i.ActiveRealDimension, i.GaugeGeneratorCount, i.BrokenGeneratorCount, i.UnbrokenGeneratorCount, i.BrokenImageRank, i.BrokenImagesIndependent, i.BrokenImageMinEigen, i.BrokenImageCondition, i.GoldstoneImageDiagnostic, i.GaugeEatingCountDiagnostic, i.EMNullNorm, i.FiniteGaugeEatingTheoremDerived, i.GaugeBosonMassMatrixDerived, i.PhysicalMassesDerived, i.CanonicalProtectedToBrokenDerived, i.Reason)
}

func FormatPhoton(p PhotonAudit) string {
	return fmt.Sprintf("%s: Qem_annihilates_vacuum=%t photon_null_residual=%.3e photon_mass_hat=%.3e unbroken=%d physical_norm=%t thetaW=%t alpha=%t; %s", p.Verdict, p.QEMAnnihilatesVacuum, p.PhotonNullResidual, p.PhotonMassSquaredHat, p.UnbrokenGeneratorCount, p.PhotonPhysicallyNormalized, p.WeakMixingAngleDerived, p.FineStructureDerived, p.Reason)
}

func FormatBoundary(b Boundary) string {
	return fmt.Sprintf("%s: socket=%t native_intertwiner=%t native_Dphi=%t vacuum_native=%t kinetic_native=%t hessian_native=%t couplings_native=%t physical_mass_matrix=%t; %s", b.Verdict, b.DiagnosticSocketPromotable, b.NativeIntertwinerDerived, b.NativeDphiDerived, b.NativeVacuumOrientationDerived, b.NativeKineticMetricDerived, b.NativeGaugeHessianDerived, b.NativeCouplingsDerived, b.PhysicalMassMatrixDerived, b.Reason)
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("%s: W=%t Z=%t Higgs=%t GF=%t thetaW=%t Yukawa=%t CKMPMNS=%t native_WZ=%t native_thetaW=%t native_Higgs=%t; %s", f.Verdict, f.ObservedWMassImported, f.ObservedZMassImported, f.ObservedHiggsMassImported, f.FermiConstantImported, f.WeakAngleImported, f.YukawaImported, f.CKMPMNSImported, f.NativeWZMassWritten, f.NativeWeakAngleWritten, f.NativeHiggsMassWritten, f.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 492 Registry Audit — Scalar Covariant Derivative and Goldstone Intertwiner Audit\n\n")
	b.WriteString("## Verdict\n\n")
	for _, v := range []string{
		StatusGate491Inherited,
		StatusAbstractDphiTemplateFound,
		StatusGoldstoneImageDiagnosticFound,
		StatusPhotonExemptionDiagnostic,
		StatusDimensionlessWZPhotonSignature,
		StatusBridgeGaugeEatingSocketPreserved,
		StatusFailedNativeDphiNotDerived,
		StatusFailedCanonicalIntertwinerNotDerived,
		StatusFailedScalarSU2NativeNotSelected,
		StatusFailedVacuumOrientationNotNative,
		StatusFailedKineticMetricBridge,
		StatusFailedGaugeHessianCouplingsNotSelected,
		StatusFailedPhysicalMassMatrixNotDerived,
		StatusFailedWeakAngleNotDerived,
		StatusFirewallPreserved,
		StatusWZMassWriteBlocked,
	} {
		b.WriteString("- `" + v + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("Gate491 supplies the scalar one-form edge carrier, ghost-free positive-semidefinite kinetic trace, and 4=1+3 count resonance. Gate492 may test the DΦ/gauge-eating socket, but it may not import W/Z masses, Higgs mass, weak-angle data, Fermi constant, Yukawa amplitudes, CKM, or PMNS.\n\n")

	b.WriteString("## Algebraic DΦ sieve\n\n")
	b.WriteString(FormatRepresentation(a.Representation) + "\n\n")
	b.WriteString(FormatDphi(a.Dphi) + "\n\n")
	b.WriteString("The current object is an abstract finite scalar covariant-derivative template on the four-real scalar frame. It is mathematically useful because it types the electroweak action, but it is still not a native finite-action theorem.\n\n")

	b.WriteString("## Protected-to-broken intertwiner audit\n\n")
	b.WriteString(FormatIntertwiner(a.Intertwiner) + "\n\n")
	b.WriteString("Three broken generator images are independent, matching the three angular/Goldstone directions. This proves a bridge-level image diagnostic, not a canonical protected-contact-to-broken-gauge isometry/intertwiner.\n\n")

	b.WriteString("## Photon exemption\n\n")
	b.WriteString(FormatPhoton(a.Photon) + "\n\n")
	b.WriteString("The diagnostic photon is protected because Qem annihilates the scalar vacuum and the template contains one null gauge direction. Physical photon normalization and the weak mixing angle remain unpromoted.\n\n")

	b.WriteString("## Firewall result\n\n")
	b.WriteString(FormatBoundary(a.Boundary) + "\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No W/Z mass, Higgs pole mass, weak mixing angle, fine-structure constant, Fermi constant, Yukawa texture, CKM, or PMNS datum entered the native registry.\n\n")

	b.WriteString("## Registry update\n\n")
	b.WriteString("### Native\n\n")
	for _, x := range a.Registry.NativeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Bridge\n\n")
	for _, x := range a.Registry.BridgeEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Environmental\n\n")
	for _, x := range a.Registry.EnvironmentalEntries {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n### Failed routes\n\n")
	for _, x := range a.Registry.FailedRoutes {
		b.WriteString("- `" + x + "`\n")
	}
	b.WriteString("\n### Open theorems\n\n")
	for _, x := range a.Registry.OpenTheorems {
		b.WriteString("- " + x + "\n")
	}

	b.WriteString("\n## Next step\n\n")
	b.WriteString(fmt.Sprintf("**Gate %d — %s.** %s Primary task: %s\n\n", a.Next.Gate, a.Next.Title, a.Next.Reason, a.Next.PrimaryTask))
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
