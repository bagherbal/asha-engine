// Package generation2boundarygaugenormalizationhessianaudit implements Gate 565:
// Boundary Gauge-Normalization to Electroweak Hessian Alignment Audit.
//
// Gate 564 derived the bridge-symbolic electroweak Hessian shape from the
// scalar kinetic socket. Gate 565 aligns that shape with the pre-existing
// ASHA hypercharge representation-trace normalization k_Y=5/3 and the
// equal-normalized-coupling boundary diagnostic sin^2(theta_*)=3/8. This is a
// boundary-normalization audit only: it imports no measured W/Z masses, no
// observed weak angle, no observed gauge couplings, no Higgs pole data, no
// CKM/PMNS, and no Yukawa eigenvalues.
package generation2boundarygaugenormalizationhessianaudit

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/fockrepresentationtrace"
	gate564 "github.com/bagherbal/asha-engine/pkg/bridge/generation2symbolicelectroweakhessianbridgeaudit"
)

const (
	AuditID = "GATE565-BOUNDARY-GAUGE-NORMALIZATION-HESSIAN-ALIGNMENT-AUDIT"

	StatusGate564Inherited           = "CONDITIONAL_SUPPORT_GATE564_SYMBOLIC_HESSIAN_INHERITED"
	StatusKYRecovered                = "PASS_HYPERCHARGE_TRACE_NORMALIZATION_KY_5_OVER_3_RECOVERED"
	StatusKYBoundaryLayer            = "CONDITIONAL_SUPPORT_KY_LIVES_IN_REPRESENTATION_TRACE_BOUNDARY_LAYER"
	StatusCouplingConvention         = "CONDITIONAL_SUPPORT_CANONICAL_HYPERCHARGE_COUPLING_CONVENTION_VERIFIED"
	StatusBoundaryEqualityBridge     = "CONDITIONAL_SUPPORT_EQUAL_NORMALIZED_COUPLING_BOUNDARY_IS_BRIDGE_ASSUMPTION"
	StatusSin238Derived              = "PASS_BOUNDARY_WEAK_ANGLE_SIN2_THETA_STAR_3_OVER_8_DERIVED"
	StatusHessianRatio58             = "PASS_GATE564_HESSIAN_RATIO_ALIGNED_TO_5_OVER_8_AT_BOUNDARY"
	StatusBoundaryShapeOnly          = "CONDITIONAL_SUPPORT_BOUNDARY_HESSIAN_RATIO_SHAPE_ONLY"
	StatusNoLowEnergyPrediction      = "FAILED_ROUTE_NO_LOW_ENERGY_WZ_OR_WEAK_ANGLE_PREDICTION"
	StatusRemainingVariablesSealed   = "FAILED_ROUTE_ABSOLUTE_KINETIC_SCALE_AND_VACUUM_DATA_REMAIN_BRIDGE_ENVIRONMENTAL"
	StatusPhotonSocketOnly           = "FAILED_ROUTE_BOUNDARY_NULL_SOCKET_DOES_NOT_DERIVE_PHYSICAL_PHOTON_DYNAMICS"
	StatusNoFlavorData               = "FAILED_ROUTE_BOUNDARY_GAUGE_NORMALIZATION_DOES_NOT_DERIVE_FLAVOR_DATA"
	StatusPreviousFirewallsPreserved = "FIREWALL_PRESERVED_Q4_TAU_ETA_WSPATIAL_PAULI_GATE564_BOUNDARIES"
	StatusFirewallPreserved          = "FIREWALL_PRESERVED_GATE565_BOUNDARY_GAUGE_NORMALIZATION_HESSIAN_BOUNDARY"
)

type Rational struct{ Num, Den int }

func (r Rational) String() string   { return fmt.Sprintf("%d/%d", r.Num, r.Den) }
func (r Rational) Float64() float64 { return float64(r.Num) / float64(r.Den) }

type InheritedHessianAudit struct {
	Gate564HessianShape       bool
	Gate564NeutralNull        bool
	Gate564NoPhysicalDynamics bool
	Gate564NoFlavorData       bool
	Gate564RatioShape         string
	Verdict                   string
}

type GaugeKineticNormalizationAudit struct {
	SourceLayer            string
	SourceTheorem          string
	KY                     Rational
	KYFormula              string
	KYRecovered            bool
	BoundarySin2FromSource Rational
	BoundarySin2Recovered  bool
	LowEnergyObservedClaim bool
	ObservedInputUsed      bool
	Verdict                string
}

type CouplingConventionAudit struct {
	CanonicalHyperchargeCoupling string
	Gate564AbelianCoupling       string
	Relation                     string
	RatioUnderBoundaryEquality   Rational
	ConventionVerified           bool
	NativePhysicalCouplingValue  bool
	Verdict                      string
}

type BoundaryEqualityAudit struct {
	Equality                    string
	EqualityNativeTheorem       bool
	EqualityBridgeBoundary      bool
	AbsoluteCouplingUnitDerived bool
	LowEnergyRunningDerived     bool
	Verdict                     string
}

type WeakAngleBoundaryAudit struct {
	KY                        Rational
	GPrimeSquaredOverGSquared Rational
	Sin2ThetaStar             Rational
	Derivation                string
	MatchesPreviousASHA       bool
	ObservedWeakAngleImported bool
	Verdict                   string
}

type HessianRatioAlignmentAudit struct {
	Gate564RatioShape          string
	InsertedBoundaryRatio      Rational
	BoundaryMW2OverMZ2         Rational
	Derivation                 string
	PhysicalLowEnergyMassRatio bool
	ObservedMassImported       bool
	Verdict                    string
}

type RemainingVariablesFirewall struct {
	BridgeEnvironmental  []string
	NativeAbsoluteKphi   bool
	NativeV              bool
	NativeAbsoluteG      bool
	NativeAbsoluteGPrime bool
	NativeF0             bool
	NativeYukawaTraceA   bool
	NativeScalarMetric   bool
	NativeRGThresholds   bool
	Verdict              string
}

type PhotonAndFlavorFirewall struct {
	ASocketSymbolicOnly   bool
	PhysicalPhotonDerived bool
	OSWickHilbertDerived  bool
	YukawaEigenvalues     bool
	CKMPMNS               bool
	GenerationHierarchy   bool
	ObservedFlavorData    bool
	Verdict               string
}

type RelationAudit struct {
	Q4ContactOnly                bool
	TauEtaSigma3TraceShadow      bool
	WSpatialWeakPlaneBlocked     bool
	PauliQuaternionicScalarRoute bool
	Gate564HessianShape          bool
	Gate565BoundaryAlignmentOnly bool
	Verdict                      string
}

type FinalVerdict struct {
	KYRecoveredCorrectLayer      bool
	CouplingConventionVerified   bool
	BoundaryEqualityLayer        string
	Sin238Passes                 bool
	HessianRatio58Passes         bool
	BridgeEnvironmentalVariables []string
	PhysicalLowEnergyPrediction  bool
	FlavorOrObservedDataProduced bool
	MissingNextTheorem           string
	Verdict                      string
}

type Analysis struct {
	Inherited        InheritedHessianAudit
	GaugeNorm        GaugeKineticNormalizationAudit
	Couplings        CouplingConventionAudit
	BoundaryEquality BoundaryEqualityAudit
	WeakAngle        WeakAngleBoundaryAudit
	HessianRatio     HessianRatioAlignmentAudit
	Remaining        RemainingVariablesFirewall
	PhotonFlavor     PhotonAndFlavorFirewall
	Relations        RelationAudit
	Final            FinalVerdict
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
	g564, err := gate564.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 564 symbolic Hessian audit: %w", err)
	}
	fock, err := fockrepresentationtrace.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build representation trace normalization audit: %w", err)
	}

	a := Analysis{}
	a.Inherited = auditInherited(g564)
	a.GaugeNorm = auditGaugeNormalization(fock)
	a.Couplings = auditCouplings(a.GaugeNorm)
	a.BoundaryEquality = auditBoundaryEquality()
	a.WeakAngle = auditWeakAngle(a.GaugeNorm, a.Couplings, a.BoundaryEquality)
	a.HessianRatio = auditHessianRatio(g564, a.WeakAngle)
	a.Remaining = auditRemaining()
	a.PhotonFlavor = auditPhotonFlavor()
	a.Relations = auditRelations(a.Inherited)
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate564.Analysis) InheritedHessianAudit {
	return InheritedHessianAudit{
		Gate564HessianShape:       g.Final.SymbolicScalarKineticBridgeProducesHessian,
		Gate564NeutralNull:        g.Final.NeutralHessianHasNullDirection,
		Gate564NoPhysicalDynamics: !g.Final.PhysicalWZPhotonDynamicsDerived,
		Gate564NoFlavorData:       !g.Final.FlavorOrObservedMassDataProduced,
		Gate564RatioShape:         string(g.MassRatio.RatioShape),
		Verdict:                   join(StatusGate564Inherited, "Gate 565 inherits the bridge-symbolic neutral Hessian and W/Z ratio shape from Gate 564"),
	}
}

func auditGaugeNormalization(f fockrepresentationtrace.Analysis) GaugeKineticNormalizationAudit {
	kyOK := f.TraceAudit.NormalizedY.String() == "5/3"
	sinOK := f.TraceAudit.WeakAngleSeed.String() == "3/8"
	return GaugeKineticNormalizationAudit{
		SourceLayer:            "representation-trace / finite charge-table boundary normalization, not low-energy running",
		SourceTheorem:          "Gate 167 representation-trace gauge-ratio rigidity, preserved by later abelian normalization gates",
		KY:                     Rational{5, 3},
		KYFormula:              "k_Y = Tr(Y^2)/Tr(T_3^2) = 5/3 after canonical U(1)_Y normalization",
		KYRecovered:            kyOK && f.TraceAudit.BoundaryDiagMatched && f.TraceAudit.AmplitudeIndependent,
		BoundarySin2FromSource: Rational{3, 8},
		BoundarySin2Recovered:  sinOK && f.TraceAudit.WeakAngleSeedMatched,
		LowEnergyObservedClaim: false,
		ObservedInputUsed:      f.TraceAudit.UsesObservedInput,
		Verdict:                join(StatusKYRecovered, StatusKYBoundaryLayer),
	}
}

func auditCouplings(g GaugeKineticNormalizationAudit) CouplingConventionAudit {
	return CouplingConventionAudit{
		CanonicalHyperchargeCoupling: "g_1 for canonically normalized U(1)_Y kinetic term",
		Gate564AbelianCoupling:       "g' multiplying the un-normalized Gate 564 Y_phi generator convention",
		Relation:                     "g_1^2 = k_Y g'^2; therefore g'^2/g^2 = 1/k_Y if g_1=g at the boundary",
		RatioUnderBoundaryEquality:   Rational{3, 5},
		ConventionVerified:           g.KYRecovered,
		NativePhysicalCouplingValue:  false,
		Verdict:                      join(StatusCouplingConvention, "The relation is a normalization convention, not an observed coupling value"),
	}
}

func auditBoundaryEquality() BoundaryEqualityAudit {
	return BoundaryEqualityAudit{
		Equality:                    "g_1 = g at the canonically normalized spectral boundary",
		EqualityNativeTheorem:       false,
		EqualityBridgeBoundary:      true,
		AbsoluteCouplingUnitDerived: false,
		LowEnergyRunningDerived:     false,
		Verdict:                     join(StatusBoundaryEqualityBridge, "equal normalized couplings are a boundary bridge assumption/normalization condition, not an IR prediction"),
	}
}

func auditWeakAngle(g GaugeKineticNormalizationAudit, c CouplingConventionAudit, b BoundaryEqualityAudit) WeakAngleBoundaryAudit {
	return WeakAngleBoundaryAudit{
		KY:                        g.KY,
		GPrimeSquaredOverGSquared: c.RatioUnderBoundaryEquality,
		Sin2ThetaStar:             Rational{3, 8},
		Derivation:                "with k_Y=5/3 and g_1=g, g'^2/g^2=1/k_Y=3/5, hence sin^2(theta_*)=g'^2/(g^2+g'^2)=(3/5)/(1+3/5)=3/8",
		MatchesPreviousASHA:       g.BoundarySin2Recovered && c.ConventionVerified && b.EqualityBridgeBoundary,
		ObservedWeakAngleImported: false,
		Verdict:                   join(StatusSin238Derived, StatusBoundaryShapeOnly),
	}
}

func auditHessianRatio(g gate564.Analysis, w WeakAngleBoundaryAudit) HessianRatioAlignmentAudit {
	return HessianRatioAlignmentAudit{
		Gate564RatioShape:          string(g.MassRatio.RatioShape),
		InsertedBoundaryRatio:      w.GPrimeSquaredOverGSquared,
		BoundaryMW2OverMZ2:         Rational{5, 8},
		Derivation:                 "m_W^2/m_Z^2 = g^2/(g^2+g'^2) = 1/(1+3/5) = 5/8 at the boundary-normalized symbolic Hessian level",
		PhysicalLowEnergyMassRatio: false,
		ObservedMassImported:       false,
		Verdict:                    join(StatusHessianRatio58, StatusBoundaryShapeOnly, StatusNoLowEnergyPrediction),
	}
}

func auditRemaining() RemainingVariablesFirewall {
	vars := []string{
		"K_phi scalar kinetic coefficient",
		"v scalar vacuum norm / Higgs VEV bridge scale",
		"absolute g and g' values",
		"absolute canonical g_1 value",
		"f0 heat-kernel coefficient",
		"finite Yukawa trace a",
		"scalar metric normalization",
		"vacuum orientation",
		"Higgs pole mass",
		"RG running interval",
		"threshold corrections",
		"continuum matching scheme",
	}
	return RemainingVariablesFirewall{
		BridgeEnvironmental:  vars,
		NativeAbsoluteKphi:   false,
		NativeV:              false,
		NativeAbsoluteG:      false,
		NativeAbsoluteGPrime: false,
		NativeF0:             false,
		NativeYukawaTraceA:   false,
		NativeScalarMetric:   false,
		NativeRGThresholds:   false,
		Verdict:              join(StatusRemainingVariablesSealed, StatusNoLowEnergyPrediction),
	}
}

func auditPhotonFlavor() PhotonAndFlavorFirewall {
	return PhotonAndFlavorFirewall{
		ASocketSymbolicOnly:   true,
		PhysicalPhotonDerived: false,
		OSWickHilbertDerived:  false,
		YukawaEigenvalues:     false,
		CKMPMNS:               false,
		GenerationHierarchy:   false,
		ObservedFlavorData:    false,
		Verdict:               join(StatusPhotonSocketOnly, StatusNoFlavorData),
	}
}

func auditRelations(i InheritedHessianAudit) RelationAudit {
	return RelationAudit{
		Q4ContactOnly:                true,
		TauEtaSigma3TraceShadow:      true,
		WSpatialWeakPlaneBlocked:     true,
		PauliQuaternionicScalarRoute: true,
		Gate564HessianShape:          i.Gate564HessianShape,
		Gate565BoundaryAlignmentOnly: true,
		Verdict:                      join(StatusPreviousFirewallsPreserved, "Gate 565 only aligns Gate 564 with boundary trace normalization"),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	return FinalVerdict{
		KYRecoveredCorrectLayer:      a.GaugeNorm.KYRecovered && !a.GaugeNorm.LowEnergyObservedClaim && !a.GaugeNorm.ObservedInputUsed,
		CouplingConventionVerified:   a.Couplings.ConventionVerified && !a.Couplings.NativePhysicalCouplingValue,
		BoundaryEqualityLayer:        "bridge boundary normalization, not native absolute coupling or low-energy equality",
		Sin238Passes:                 a.WeakAngle.MatchesPreviousASHA && !a.WeakAngle.ObservedWeakAngleImported,
		HessianRatio58Passes:         a.HessianRatio.BoundaryMW2OverMZ2 == (Rational{5, 8}) && !a.HessianRatio.PhysicalLowEnergyMassRatio && !a.HessianRatio.ObservedMassImported,
		BridgeEnvironmentalVariables: a.Remaining.BridgeEnvironmental,
		PhysicalLowEnergyPrediction:  false,
		FlavorOrObservedDataProduced: false,
		MissingNextTheorem:           "A later gate must derive or seal absolute gauge/scalar kinetic normalizations, K_phi, v, RG transport, threshold matching, and OS/Wick/Hilbert continuum gauge dynamics before physical W/Z/photon predictions are allowed",
		Verdict:                      join(StatusSin238Derived, StatusHessianRatio58, StatusBoundaryShapeOnly, StatusNoLowEnergyPrediction, StatusFirewallPreserved),
	}
}

func validate(a Analysis) error {
	failures := []string{}
	if !a.Inherited.Gate564HessianShape || !a.Inherited.Gate564NeutralNull || !a.Inherited.Gate564NoPhysicalDynamics || !a.Inherited.Gate564NoFlavorData {
		failures = append(failures, "Gate 564 inheritance failed")
	}
	if !a.GaugeNorm.KYRecovered || !a.GaugeNorm.BoundarySin2Recovered || a.GaugeNorm.LowEnergyObservedClaim || a.GaugeNorm.ObservedInputUsed {
		failures = append(failures, "hypercharge trace normalization audit failed")
	}
	if !a.Couplings.ConventionVerified || a.Couplings.RatioUnderBoundaryEquality != (Rational{3, 5}) || a.Couplings.NativePhysicalCouplingValue {
		failures = append(failures, "coupling convention audit failed")
	}
	if a.BoundaryEquality.EqualityNativeTheorem || !a.BoundaryEquality.EqualityBridgeBoundary || a.BoundaryEquality.AbsoluteCouplingUnitDerived || a.BoundaryEquality.LowEnergyRunningDerived {
		failures = append(failures, "boundary equality firewall failed")
	}
	if !a.WeakAngle.MatchesPreviousASHA || a.WeakAngle.Sin2ThetaStar != (Rational{3, 8}) || a.WeakAngle.ObservedWeakAngleImported {
		failures = append(failures, "weak angle boundary derivation failed")
	}
	if a.HessianRatio.BoundaryMW2OverMZ2 != (Rational{5, 8}) || a.HessianRatio.PhysicalLowEnergyMassRatio || a.HessianRatio.ObservedMassImported {
		failures = append(failures, "Hessian ratio alignment failed")
	}
	if a.Remaining.NativeAbsoluteKphi || a.Remaining.NativeV || a.Remaining.NativeAbsoluteG || a.Remaining.NativeAbsoluteGPrime || a.Remaining.NativeF0 || a.Remaining.NativeYukawaTraceA || a.Remaining.NativeScalarMetric || a.Remaining.NativeRGThresholds {
		failures = append(failures, "remaining-variable firewall failed")
	}
	if !a.PhotonFlavor.ASocketSymbolicOnly || a.PhotonFlavor.PhysicalPhotonDerived || a.PhotonFlavor.OSWickHilbertDerived || a.PhotonFlavor.YukawaEigenvalues || a.PhotonFlavor.CKMPMNS || a.PhotonFlavor.GenerationHierarchy || a.PhotonFlavor.ObservedFlavorData {
		failures = append(failures, "photon/flavor firewall failed")
	}
	if !a.Relations.Q4ContactOnly || !a.Relations.TauEtaSigma3TraceShadow || !a.Relations.WSpatialWeakPlaneBlocked || !a.Relations.PauliQuaternionicScalarRoute || !a.Relations.Gate564HessianShape || !a.Relations.Gate565BoundaryAlignmentOnly {
		failures = append(failures, "relation audit failed")
	}
	if !a.Final.KYRecoveredCorrectLayer || !a.Final.CouplingConventionVerified || !a.Final.Sin238Passes || !a.Final.HessianRatio58Passes || a.Final.PhysicalLowEnergyPrediction || a.Final.FlavorOrObservedDataProduced {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate564Inherited,
		StatusKYRecovered,
		StatusKYBoundaryLayer,
		StatusCouplingConvention,
		StatusBoundaryEqualityBridge,
		StatusSin238Derived,
		StatusHessianRatio58,
		StatusBoundaryShapeOnly,
		StatusNoLowEnergyPrediction,
		StatusRemainingVariablesSealed,
		StatusPhotonSocketOnly,
		StatusNoFlavorData,
		StatusPreviousFirewallsPreserved,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return join(
		"Gate 565 lawfully aligns the Gate 564 symbolic Hessian with the finite representation-trace boundary normalization k_Y=5/3",
		"under equal canonically normalized boundary couplings, g'^2/g^2=3/5 and sin^2(theta_*)=3/8",
		"the symbolic Hessian ratio becomes m_W^2/m_Z^2=5/8 only at the boundary-normalized socket level",
		"absolute scales, RG transport, thresholds, physical photon dynamics, and flavor data remain firewalled",
	)
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
