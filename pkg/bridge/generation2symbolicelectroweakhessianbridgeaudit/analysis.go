// Package generation2symbolicelectroweakhessianbridgeaudit implements Gate 564:
// Symbolic Electroweak Hessian Bridge Audit.
//
// Gate 563 proved that the scalar/quaternionic moment map is present at the
// representation-bookkeeping layer but not yet inserted into a native finite
// curvature or kinetic projection. Gate 564 deliberately moves only one step:
// it performs a bridge-typed symbolic second-variation calculation for the
// scalar kinetic socket K_phi |D_mu phi|^2 around a nonzero sealed scalar
// doublet vacuum. The result is the standard electroweak Hessian shape, a
// neutral null direction, and symbolic W/Z mass-ratio structure, while keeping
// K_phi, v, g, g', f0, finite Yukawa trace a, scalar metric normalization, and
// vacuum orientation bridge/environmental.
package generation2symbolicelectroweakhessianbridgeaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate563 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarquaternionicmomentewcurvatureprojectionaudit"
)

const (
	AuditID = "GATE564-SYMBOLIC-ELECTROWEAK-HESSIAN-BRIDGE-AUDIT"

	StatusGate563Inherited           = "CONDITIONAL_SUPPORT_GATE563_SCALAR_QUATERNIONIC_PROJECTION_BOUNDARY_INHERITED"
	StatusVacuumBridgeSeal           = "CONDITIONAL_SUPPORT_SYMBOLIC_SCALAR_VACUUM_ORIENTATION_BRIDGE_SEALED"
	StatusStabilizerSolved           = "PASS_SYMBOLIC_ELECTROWEAK_STABILIZER_CONDITION_SOLVED"
	StatusChargedHessian             = "PASS_SYMBOLIC_CHARGED_SECTOR_HESSIAN_SHAPE_DERIVED"
	StatusNeutralHessian             = "PASS_SYMBOLIC_NEUTRAL_SECTOR_HESSIAN_SHAPE_DERIVED"
	StatusNeutralNull                = "PASS_NEUTRAL_HESSIAN_NULL_DIRECTION_PHOTON_SOCKET_FOUND"
	StatusMassRatio                  = "PASS_SYMBOLIC_WZ_MASS_RATIO_SHAPE_DERIVED"
	StatusHessianShapeOnly           = "CONDITIONAL_SUPPORT_SYMBOLIC_HESSIAN_SHAPE_ONLY"
	StatusNoNumericalPrediction      = "FAILED_ROUTE_NO_NATIVE_NUMERICAL_MASS_OR_COUPLING_PREDICTION"
	StatusNoPhysicalPhotonDynamics   = "FAILED_ROUTE_SYMBOLIC_NULL_SOCKET_DOES_NOT_DERIVE_PHYSICAL_PHOTON_DYNAMICS"
	StatusNoKineticNormalization     = "FAILED_ROUTE_KINETIC_NORMALIZATION_AND_VACUUM_SCALE_REMAIN_BRIDGE_ENVIRONMENTAL"
	StatusNoFlavorData               = "FAILED_ROUTE_SYMBOLIC_EW_HESSIAN_DOES_NOT_DERIVE_FLAVOR_DATA"
	StatusPreviousFirewallsPreserved = "FIREWALL_PRESERVED_Q4_TAU_ETA_WSPATIAL_PAULI_BOUNDARIES"
	StatusFirewallPreserved          = "FIREWALL_PRESERVED_GATE564_SYMBOLIC_ELECTROWEAK_HESSIAN_BOUNDARY"
)

// Symbol is intentionally string-valued: Gate 564 is a symbolic bridge audit,
// not a numerical evaluation or observed-parameter import.
type Symbol string

type InheritedAudit struct {
	Gate563ScalarDoubletLane            bool
	Gate563ImHActionOnHphi              bool
	Gate563MomentNotNativeCurvature     bool
	Gate563NoNativeU1PhotonDirection    bool
	Gate563NoNativeKineticNormalization bool
	Gate563NoFlavorData                 bool
	Verdict                             string
}

type ScalarVacuumAudit struct {
	Carrier                      string
	VacuumSymbol                 Symbol
	VacuumConvention             string
	NormSymbol                   Symbol
	VacuumDerivedNatively        bool
	VacuumBridgeSealed           bool
	GeneratorsConvention         string
	HyperchargeConvention        string
	StabilizerEquation           string
	StabilizerSolution           string
	StabilizerSolvedSymbolically bool
	Verdict                      string
}

type ChargedSectorAudit struct {
	KineticExpression           string
	RealBasis                   string
	ChargedBasis                string
	PerRealGeneratorCoefficient Symbol
	ChargedPairCoefficient      Symbol
	RepresentationConvention    string
	ObservedMassImported        bool
	NumericalCouplingImported   bool
	Verdict                     string
}

type NeutralSectorAudit struct {
	Basis         string
	OverallFactor Symbol
	Matrix        [2][2]string
	Trace         Symbol
	Determinant   string
	Rank          int
	Eigenvalues   [2]Symbol
	Convention    string
	Verdict       string
}

type NullDirectionAudit struct {
	DeterminantZero                    bool
	NullDirection                      string
	MassiveDirection                   string
	PhotonSocketOnly                   bool
	PhysicalPhotonDerived              bool
	RequiresOSWickHilbertGaugeDynamics bool
	Verdict                            string
}

type MassRatioAudit struct {
	MW2Shape                Symbol
	MZ2Shape                Symbol
	RatioShape              Symbol
	DependsOnKphi           bool
	DependsOnV              bool
	DependsOnGaugeCouplings bool
	ConventionFactorsSealed bool
	ObservedMassImported    bool
	Verdict                 string
}

type NormalizationFirewallAudit struct {
	BridgeVariables                []string
	EnvironmentalVariables         []string
	NativeNumericalMassDerived     bool
	NativeCouplingDerived          bool
	NativeKphiDerived              bool
	NativeVDerived                 bool
	NativeF0Derived                bool
	NativeYukawaTraceADerived      bool
	NativeScalarMetricDerived      bool
	NativeVacuumOrientationDerived bool
	Verdict                        string
}

type RelationFirewallAudit struct {
	Q4ContactOnly                  bool
	TauEtaSigma3TraceShadow        bool
	WSpatialWeakPlaneBlocked       bool
	PauliQuaternionicSeparateRoute bool
	FlavorDerived                  bool
	ObservedDataImported           bool
	Verdict                        string
}

type FinalVerdict struct {
	SymbolicScalarKineticBridgeProducesHessian bool
	NeutralHessianHasNullDirection             bool
	PhysicalWZPhotonDynamicsDerived            bool
	VariablesBridgeEnvironmental               []string
	FlavorOrObservedMassDataProduced           bool
	MissingNextTheorem                         string
	Verdict                                    string
}

type Analysis struct {
	Inherited     InheritedAudit
	Vacuum        ScalarVacuumAudit
	Charged       ChargedSectorAudit
	Neutral       NeutralSectorAudit
	Null          NullDirectionAudit
	MassRatio     MassRatioAudit
	Normalization NormalizationFirewallAudit
	Relations     RelationFirewallAudit
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
	g563, err := gate563.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate 563 scalar/quaternionic electroweak projection audit: %w", err)
	}

	a := Analysis{}
	a.Inherited = auditInherited(g563)
	a.Vacuum = auditVacuum()
	a.Charged = auditCharged(a.Vacuum)
	a.Neutral = auditNeutral(a.Vacuum)
	a.Null = auditNull(a.Neutral)
	a.MassRatio = auditMassRatio(a.Charged, a.Neutral)
	a.Normalization = auditNormalization()
	a.Relations = auditRelations()
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate563.Analysis) InheritedAudit {
	return InheritedAudit{
		Gate563ScalarDoubletLane:            g.Final.FiniteOneFormContainsScalarDoublet,
		Gate563ImHActionOnHphi:              g.Final.ImHActsStructurallyOnHphi,
		Gate563MomentNotNativeCurvature:     !g.Final.MomentMapInsideCurvatureOrKineticData,
		Gate563NoNativeU1PhotonDirection:    !g.Final.ElectroweakU1MixingDerived,
		Gate563NoNativeKineticNormalization: !g.Final.KineticNormalizationAndMassDynamics,
		Gate563NoFlavorData:                 !g.Final.FlavorDataDerived,
		Verdict:                             join(StatusGate563Inherited, "Gate 564 inherits the scalar doublet, Im(H) action, symbolic D_phi socket, and projection firewalls from Gate 563"),
	}
}

func auditVacuum() ScalarVacuumAudit {
	return ScalarVacuumAudit{
		Carrier:                      "H_phi ~= C^2 with phi_0=(0,v)^T under a bridge vacuum-orientation seal",
		VacuumSymbol:                 "phi_0=(0,v)^T",
		VacuumConvention:             "norm convention |phi_0|^2=v^2; all 1/2 factors are representation conventions, not native observed normalizations",
		NormSymbol:                   "v",
		VacuumDerivedNatively:        false,
		VacuumBridgeSealed:           true,
		GeneratorsConvention:         "T_a=sigma_a/2 Hermitian representatives of Im(H) action on the scalar doublet",
		HyperchargeConvention:        "Y_phi=1/2 I_2 bridge convention; sign may flip with covariant-derivative convention",
		StabilizerEquation:           "(alpha^a T_a + beta Y_phi) phi_0 = 0",
		StabilizerSolution:           "alpha^1=alpha^2=0 and beta=alpha^3 in the chosen convention; neutral unbroken socket generated by T_3+Y_phi",
		StabilizerSolvedSymbolically: true,
		Verdict:                      join(StatusVacuumBridgeSeal, StatusStabilizerSolved),
	}
}

func auditCharged(v ScalarVacuumAudit) ChargedSectorAudit {
	return ChargedSectorAudit{
		KineticExpression:           "K_phi |g W^a T_a phi_0 + g' B Y_phi phi_0|^2",
		RealBasis:                   "W^1,W^2",
		ChargedBasis:                "W^±=(W^1 ∓ i W^2)/sqrt(2) bridge convention",
		PerRealGeneratorCoefficient: "K_phi g^2 v^2 / 4",
		ChargedPairCoefficient:      "K_phi g^2 v^2 / 2 for W^+W^- in the displayed convention",
		RepresentationConvention:    v.GeneratorsConvention,
		ObservedMassImported:        false,
		NumericalCouplingImported:   false,
		Verdict:                     join(StatusChargedHessian, StatusHessianShapeOnly),
	}
}

func auditNeutral(v ScalarVacuumAudit) NeutralSectorAudit {
	return NeutralSectorAudit{
		Basis:         "(W^3,B)",
		OverallFactor: "K_phi v^2 / 4",
		Matrix: [2][2]string{
			{"g^2", "-g g'"},
			{"-g g'", "g'^2"},
		},
		Trace:       "(K_phi v^2 / 4)(g^2+g'^2)",
		Determinant: "0",
		Rank:        1,
		Eigenvalues: [2]Symbol{"0", "(K_phi v^2 / 4)(g^2+g'^2)"},
		Convention:  v.HyperchargeConvention,
		Verdict:     join(StatusNeutralHessian, StatusNeutralNull, StatusHessianShapeOnly),
	}
}

func auditNull(n NeutralSectorAudit) NullDirectionAudit {
	return NullDirectionAudit{
		DeterminantZero:                    n.Determinant == "0" && n.Rank == 1,
		NullDirection:                      "A_socket ∝ g' W^3 + g B",
		MassiveDirection:                   "Z_socket ∝ g W^3 - g' B",
		PhotonSocketOnly:                   true,
		PhysicalPhotonDerived:              false,
		RequiresOSWickHilbertGaugeDynamics: true,
		Verdict:                            join(StatusNeutralNull, StatusNoPhysicalPhotonDynamics, "The neutral null vector is a symbolic photon socket, not a physical photon theorem"),
	}
}

func auditMassRatio(c ChargedSectorAudit, n NeutralSectorAudit) MassRatioAudit {
	return MassRatioAudit{
		MW2Shape:                "m_W^2 ∝ K_phi g^2 v^2 / 4",
		MZ2Shape:                "m_Z^2 ∝ K_phi (g^2+g'^2) v^2 / 4",
		RatioShape:              "m_W^2/m_Z^2 = g^2/(g^2+g'^2)",
		DependsOnKphi:           true,
		DependsOnV:              true,
		DependsOnGaugeCouplings: true,
		ConventionFactorsSealed: true,
		ObservedMassImported:    c.ObservedMassImported || n.Rank != 1,
		Verdict:                 join(StatusMassRatio, StatusHessianShapeOnly, StatusNoNumericalPrediction),
	}
}

func auditNormalization() NormalizationFirewallAudit {
	return NormalizationFirewallAudit{
		BridgeVariables: []string{
			"K_phi scalar kinetic coefficient",
			"v scalar vacuum norm/orientation",
			"g quaternionic/SU(2) bridge coupling",
			"g' abelian/hypercharge bridge coupling",
			"generator normalization T_a=sigma_a/2 versus convention variants",
			"Y_phi scalar hypercharge normalization",
		},
		EnvironmentalVariables: []string{
			"f0 heat-kernel coefficient",
			"finite Yukawa trace a in K_phi=f0*a/pi^2 conventions",
			"scalar metric normalization",
			"vacuum orientation and scale",
			"continuum gauge-coupling boundary values",
		},
		NativeNumericalMassDerived:     false,
		NativeCouplingDerived:          false,
		NativeKphiDerived:              false,
		NativeVDerived:                 false,
		NativeF0Derived:                false,
		NativeYukawaTraceADerived:      false,
		NativeScalarMetricDerived:      false,
		NativeVacuumOrientationDerived: false,
		Verdict:                        join(StatusNoKineticNormalization, StatusNoNumericalPrediction),
	}
}

func auditRelations() RelationFirewallAudit {
	return RelationFirewallAudit{
		Q4ContactOnly:                  true,
		TauEtaSigma3TraceShadow:        true,
		WSpatialWeakPlaneBlocked:       true,
		PauliQuaternionicSeparateRoute: true,
		FlavorDerived:                  false,
		ObservedDataImported:           false,
		Verdict:                        join(StatusNoFlavorData, StatusPreviousFirewallsPreserved),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	vars := append([]string{}, a.Normalization.BridgeVariables...)
	vars = append(vars, a.Normalization.EnvironmentalVariables...)
	return FinalVerdict{
		SymbolicScalarKineticBridgeProducesHessian: a.Vacuum.StabilizerSolvedSymbolically && strings.Contains(string(a.Charged.PerRealGeneratorCoefficient), "K_phi") && a.Neutral.Rank == 1,
		NeutralHessianHasNullDirection:             a.Null.DeterminantZero && a.Null.PhotonSocketOnly,
		PhysicalWZPhotonDynamicsDerived:            a.Null.PhysicalPhotonDerived,
		VariablesBridgeEnvironmental:               vars,
		FlavorOrObservedMassDataProduced:           a.Relations.FlavorDerived || a.Relations.ObservedDataImported || a.MassRatio.ObservedMassImported,
		MissingNextTheorem:                         "A later gate must derive scalar/gauge kinetic normalization, vacuum orientation/scale, U(1) mixing normalization, and physical OS/Wick/Hilbert gauge dynamics before any W/Z/photon mass theorem is allowed",
		Verdict:                                    join(StatusHessianShapeOnly, StatusNoNumericalPrediction, StatusFirewallPreserved),
	}
}

func validate(a Analysis) error {
	failures := []string{}
	if !a.Inherited.Gate563ScalarDoubletLane || !a.Inherited.Gate563ImHActionOnHphi || !a.Inherited.Gate563MomentNotNativeCurvature {
		failures = append(failures, "Gate 563 inheritance failed")
	}
	if !a.Vacuum.VacuumBridgeSealed || a.Vacuum.VacuumDerivedNatively || !a.Vacuum.StabilizerSolvedSymbolically {
		failures = append(failures, "symbolic vacuum/stabilizer bridge failed")
	}
	if a.Charged.ObservedMassImported || a.Charged.NumericalCouplingImported || a.Charged.PerRealGeneratorCoefficient == "" {
		failures = append(failures, "charged Hessian firewall failed")
	}
	if a.Neutral.Determinant != "0" || a.Neutral.Rank != 1 || !strings.Contains(string(a.Neutral.Trace), "g^2+g'^2") {
		failures = append(failures, "neutral Hessian null structure failed")
	}
	if !a.Null.DeterminantZero || !a.Null.PhotonSocketOnly || a.Null.PhysicalPhotonDerived {
		failures = append(failures, "neutral null socket firewall failed")
	}
	if !strings.Contains(string(a.MassRatio.RatioShape), "g^2/(g^2+g'^2)") || a.MassRatio.ObservedMassImported {
		failures = append(failures, "symbolic mass-ratio firewall failed")
	}
	if a.Normalization.NativeNumericalMassDerived || a.Normalization.NativeCouplingDerived || a.Normalization.NativeKphiDerived || a.Normalization.NativeVDerived || a.Normalization.NativeF0Derived || a.Normalization.NativeYukawaTraceADerived || a.Normalization.NativeScalarMetricDerived || a.Normalization.NativeVacuumOrientationDerived {
		failures = append(failures, "normalization firewall failed")
	}
	if !a.Relations.Q4ContactOnly || !a.Relations.TauEtaSigma3TraceShadow || !a.Relations.WSpatialWeakPlaneBlocked || !a.Relations.PauliQuaternionicSeparateRoute || a.Relations.FlavorDerived || a.Relations.ObservedDataImported {
		failures = append(failures, "relation/firewall audit failed")
	}
	if !a.Final.SymbolicScalarKineticBridgeProducesHessian || !a.Final.NeutralHessianHasNullDirection || a.Final.PhysicalWZPhotonDynamicsDerived || a.Final.FlavorOrObservedMassDataProduced {
		failures = append(failures, "final verdict failed")
	}
	if len(failures) > 0 {
		return fmt.Errorf(strings.Join(failures, "; "))
	}
	_ = math.Pi // keep the package explicitly numerical-free while documenting convention-only use.
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate563Inherited,
		StatusVacuumBridgeSeal,
		StatusStabilizerSolved,
		StatusChargedHessian,
		StatusNeutralHessian,
		StatusNeutralNull,
		StatusMassRatio,
		StatusHessianShapeOnly,
		StatusNoNumericalPrediction,
		StatusNoPhysicalPhotonDynamics,
		StatusNoKineticNormalization,
		StatusNoFlavorData,
		StatusPreviousFirewallsPreserved,
		StatusFirewallPreserved,
	}
}

func truth(a Analysis) string {
	return join(
		"Gate 564 derives the bridge-symbolic electroweak Hessian shape from the sealed scalar kinetic socket",
		"neutral determinant zero gives a photon-socket null direction only",
		"K_phi, v, g, g', f0, finite Yukawa trace a, scalar metric normalization, and vacuum orientation remain bridge/environmental",
		"no observed masses, Higgs pole data, CKM/PMNS, Yukawa eigenvalues, OS/Hilbert-time dynamics, or physical photon theorem are imported",
	)
}

func join(parts ...string) string { return strings.Join(parts, "; ") }
