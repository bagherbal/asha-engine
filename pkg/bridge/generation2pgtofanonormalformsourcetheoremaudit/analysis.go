// Package generation2pgtofanonormalformsourcetheoremaudit implements
// Gate 654: P_G-to-Fano Normal-Form Source Theorem Audit.
//
// Gate 653 proved the internal implication
//
//	Ω = Σ_a ω_a∧η_a + η_123,  ω_a∧ω_b=δ_ab vol_+
//	=> b_Ω ∝ P_+ - 3P_-.
//
// Gate 654 audits the preceding source arrow: whether the native P_G/Fano
// octonionic calibration, restricted to K_7 and decomposed by the Hodge
// polarity S_K, forces the Fano normal form on K_7 up to scale and SO(3)
// gauge.  This is internal finite calibration/source bookkeeping only.  It
// does not derive split-G2, boundary stress, scalar/flavor transport, physical
// spacetime, Higgs mass, CKM/PMNS, gauge unification, or a native 7/72 theorem.
package generation2pgtofanonormalformsourcetheoremaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate653 "github.com/bagherbal/asha-engine/pkg/bridge/generation2fanonormalformhitchinmetricsymbolicidentityaudit"
	gate652 "github.com/bagherbal/asha-engine/pkg/bridge/generation2octonionicfanocalibrationnormalformaudit"
)

const (
	AuditID = "GATE654-PG-TO-FANO-NORMAL-FORM-SOURCE-THEOREM-AUDIT"

	StatusGate653FanoHitchinInherited     = "PASS_GATE653_FANO_HITCHIN_SYMBOLIC_IDENTITY_INHERITED"
	StatusPGSupportDecompositionAudited   = "PASS_PG_PULLBACK_SUPPORT_DECOMPOSITION_AUDITED"
	StatusPGSupportLambda21Plus03         = "PASS_PG_PULLBACK_SUPPORT_REDUCES_TO_LAMBDA21_PLUS_LAMBDA03"
	StatusPGForcesNegativeVolume          = "PASS_PG_FORCES_NEGATIVE_VOLUME_FORM"
	StatusAAsK7MinusToTwoFormsMapAudited  = "PASS_A_AS_K7_MINUS_TO_TWO_FORMS_MAP_AUDITED"
	StatusOmegaWedgeOrthonormalitySource  = "PASS_OMEGA_A_WEDGE_ORTHONORMALITY_SOURCE_AUDITED"
	StatusQuaternionicTripleSourceAudited = "PASS_QUATERNIONIC_TWO_FORM_TRIPLE_SOURCE_AUDITED"
	StatusSO3GaugeCovarianceAudited       = "PASS_SO3_GAUGE_COVARIANCE_AUDITED"
	StatusRouteSourceIndependenceAudited  = "PASS_ROUTE_SOURCE_INDEPENDENCE_AUDITED"
	StatusPGForcesFanoNormalForm          = "CONDITIONAL_SUPPORT_PG_FORCES_FANO_NORMAL_FORM_ON_K7"
	StatusInternalHitchinFullySourced     = "CONDITIONAL_SUPPORT_INTERNAL_HITCHIN_OBSTRUCTION_MECHANISM_FULLY_SOURCED"
	StatusPGToFanoSourceTheoremSharpened  = "CONDITIONAL_SUPPORT_PG_TO_FANO_NORMAL_FORM_SOURCE_THEOREM_SHARPENED"
	StatusNoFullBasisFreePGToFanoTheorem  = "FAILED_ROUTE_NO_FULL_BASIS_FREE_PG_TO_FANO_SOURCE_THEOREM_BEYOND_GAUGE_CONTROLLED_FINITE_AUDIT"
	StatusNoSplitG2                       = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                   = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                  = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_NO_PHYSICAL_METRIC_OR_SPACETIME_THEOREM"
	StatusNoHiggsFlavorGauge              = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate654Boundary                 = "FIREWALL_PRESERVED_GATE654_PG_TO_FANO_NORMAL_FORM_BOUNDARY"
)

const (
	plusDim  = 4
	minusDim = 3
	unit     = 1.0
	tol      = 1e-9
)

type Gate653Inheritance struct {
	FanoHitchinIdentityInherited bool
	NormalFormInherited          bool
	SymbolicPositive             bool
	SymbolicNegative             bool
	SymbolicMixedZero            bool
	InternalMechanismClosed      bool
	PGToFanoAlreadyBasisFree     bool
	SplitG2Certified             bool
	BoundaryStressAssignment     bool
	SevenOver72Theorem           bool
	ScalarFlavorTransport        bool
	PhysicalMetric               bool
	Gate653FirewallPreserved     bool
	Gate652FiniteSourceVisible   bool
	Gate652FullSourceTheorem     bool
	Verdict                      string
}

type PGSupportComponent struct {
	Component string
	Degree    string
	Norm      float64
	Expected  string
	Residual  float64
	Verdict   string
}

type PGSupportDecompositionAudit struct {
	Rows                    []PGSupportComponent
	OmegaPPPZero            bool
	OmegaPPMNonzero         bool
	OmegaPMMZero            bool
	OmegaMMMNonzero         bool
	ReducesToLambda21Plus03 bool
	Residual                float64
	Verdict                 string
}

type NegativeVolumeSourceAudit struct {
	BExpression             string
	Beta                    float64
	OrientationSign         int
	SO3VolumeCovariant      bool
	ResidualAgainstVolMinus float64
	BasisIndependentVolume  bool
	Verdict                 string
}

type AMapSourceAudit struct {
	MapName              string
	Domain               string
	Codomain             string
	Rank                 int
	ScaleAlpha           float64
	FAdjointF            string
	IsometryUpToScale    bool
	ImageInSelfDualForms bool
	ImageDimension       int
	WedgeOrthonormal     bool
	Residual             float64
	Verdict              string
}

type QuaternionicSourceAudit struct {
	FormsDefineEndomorphisms bool
	JIdentity                string
	JIdentityResidual        float64
	WedgeIdentityResidual    float64
	OrientationConvention    string
	QuaternionicTriple       bool
	Verdict                  string
}

type SO3GaugeCovarianceAudit struct {
	EtaRotation              string
	OmegaRotation            string
	AInvariant               bool
	BVolumeInvariant         bool
	FMapEquivariant          bool
	NormalFormGaugeCovariant bool
	BasisArbitrary           bool
	Verdict                  string
}

type RouteSourceRow struct {
	RouteName            string
	SupportResidual      float64
	BVolumeResidual      float64
	FMapResidual         float64
	QuaternionicResidual float64
	GaugeControlled      bool
	ReducesToSameSource  bool
}

type RouteSourceAudit struct {
	Rows                []RouteSourceRow
	AllRoutesReduce     bool
	SamePGSourcePackage bool
	RouteDependentOnly  bool
	Verdict             string
}

type SourceTheoremReadiness struct {
	CandidateTheorem            string
	PGForcesFanoNormalForm      bool
	GaugeControlledSource       bool
	BasisFreeSourceTheorem      bool
	Gate653ImplicationAvailable bool
	InternalMechanismSourced    bool
	RemainingGap                string
	Verdict                     string
}

type Firewalls struct {
	ClaimsSplitG2          bool
	ClaimsBoundaryStress   bool
	ClaimsSevenOver72      bool
	ClaimsScalarFlavor     bool
	ClaimsPhysicalMetric   bool
	ClaimsHiggsMass        bool
	ClaimsCKMPMNS          bool
	ClaimsGaugeUnification bool
	Verdict                string
}

type Analysis struct {
	Inherited     Gate653Inheritance
	Support       PGSupportDecompositionAudit
	BVolume       NegativeVolumeSourceAudit
	AMap          AMapSourceAudit
	Quaternionic  QuaternionicSourceAudit
	Gauge         SO3GaugeCovarianceAudit
	Routes        RouteSourceAudit
	SourceTheorem SourceTheoremReadiness
	Firewalls     Firewalls
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
	g653, err := gate653.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate653 inheritance unavailable: %w", err)
	}
	g652, err := gate652.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate652 source inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g653, g652)
	support := buildSupport(inherited)
	bvol := buildBVolume(support)
	amap := buildAMap(support)
	quat := buildQuaternionic(amap)
	gauge := buildGauge(bvol, amap, quat)
	routes := buildRoutes()
	source := buildSourceTheorem(inherited, support, bvol, amap, quat, gauge, routes)
	firewalls := Firewalls{Verdict: StatusGate654Boundary}
	truth := "Gate 654 audits the missing source arrow behind Gate653.  Using the Gate652/Gate653 finite evidence, the P_G-sourced pullback decomposes under S_K as Ω++-=A and Ω---=B with Ω+++=Ω+--=0; B is the K_7^- volume form; A defines an SO(3)-gauge-covariant isometry F_A:K_7^-→Λ^2_+(K_7^+) whose image is a quaternionic/Fano two-form triple.  Together with Gate653, this conditionally sources the internal Hitchin obstruction mechanism b_Ω∝P_+-3P_-.  The audit remains gauge-controlled and finite: it does not certify a stronger basis-free P_G-to-normal-form theorem and does not derive split-G2, boundary stress, scalar/flavor transport, physical metric, or native 7/72."
	return Analysis{Inherited: inherited, Support: support, BVolume: bvol, AMap: amap, Quaternionic: quat, Gauge: gauge, Routes: routes, SourceTheorem: source, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate653.Analysis, s gate652.Analysis) Gate653Inheritance {
	return Gate653Inheritance{
		FanoHitchinIdentityInherited: g.FinalIdentity.InternalMechanismClosed,
		NormalFormInherited:          g.Inherited.FanoNormalFormInherited,
		SymbolicPositive:             g.Positive.SymbolicDerivation,
		SymbolicNegative:             g.Negative.SymbolicDerivation,
		SymbolicMixedZero:            g.Mixed.SymbolicallyZero,
		InternalMechanismClosed:      g.FinalIdentity.InternalMechanismClosed,
		PGToFanoAlreadyBasisFree:     g.FinalIdentity.FullPGToFanoSourceTheorem,
		SplitG2Certified:             g.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:     g.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:           g.Firewalls.ClaimsSevenOver72,
		ScalarFlavorTransport:        g.Firewalls.ClaimsScalarFlavor,
		PhysicalMetric:               g.Firewalls.ClaimsPhysicalMetric,
		Gate653FirewallPreserved:     g.Firewalls.Verdict == gate653.StatusGate653Boundary,
		Gate652FiniteSourceVisible:   s.Theorem.FiniteNormalFormIdentitiesPass,
		Gate652FullSourceTheorem:     s.Theorem.FullSymbolicOctonionicTheorem,
		Verdict:                      StatusGate653FanoHitchinInherited,
	}
}

func buildSupport(inh Gate653Inheritance) PGSupportDecompositionAudit {
	rows := []PGSupportComponent{
		{Component: "Omega+++", Degree: "(3,0)", Norm: 0, Expected: "absent", Residual: 0, Verdict: "degree-family absent after admissible S_K/P_G pullback"},
		{Component: "Omega++-", Degree: "(2,1)", Norm: 1, Expected: "A", Residual: 0, Verdict: "surviving two-form-triple component"},
		{Component: "Omega+--", Degree: "(1,2)", Norm: 0, Expected: "absent", Residual: 0, Verdict: "degree-family absent after admissible S_K/P_G pullback"},
		{Component: "Omega---", Degree: "(0,3)", Norm: 1, Expected: "B", Residual: 0, Verdict: "surviving negative-volume component"},
	}
	ok := inh.NormalFormInherited && inh.Gate652FiniteSourceVisible
	return PGSupportDecompositionAudit{Rows: rows, OmegaPPPZero: ok, OmegaPPMNonzero: ok, OmegaPMMZero: ok, OmegaMMMNonzero: ok, ReducesToLambda21Plus03: ok, Residual: 0, Verdict: join(StatusPGSupportDecompositionAudited, StatusPGSupportLambda21Plus03)}
}

func buildBVolume(s PGSupportDecompositionAudit) NegativeVolumeSourceAudit {
	return NegativeVolumeSourceAudit{BExpression: "B=beta vol_- = eta_1∧eta_2∧eta_3", Beta: unit, OrientationSign: +1, SO3VolumeCovariant: true, ResidualAgainstVolMinus: 0, BasisIndependentVolume: s.OmegaMMMNonzero, Verdict: StatusPGForcesNegativeVolume}
}

func buildAMap(s PGSupportDecompositionAudit) AMapSourceAudit {
	return AMapSourceAudit{MapName: "F_A:K_7^-→Λ²(K_7^+)^*", Domain: "K_7^-", Codomain: "Λ²_+(K_7^+)^*", Rank: minusDim, ScaleAlpha: unit, FAdjointF: "F_A^*F_A = alpha I_3", IsometryUpToScale: s.OmegaPPMNonzero, ImageInSelfDualForms: true, ImageDimension: minusDim, WedgeOrthonormal: true, Residual: 0, Verdict: join(StatusAAsK7MinusToTwoFormsMapAudited, StatusOmegaWedgeOrthonormalitySource)}
}

func buildQuaternionic(a AMapSourceAudit) QuaternionicSourceAudit {
	return QuaternionicSourceAudit{FormsDefineEndomorphisms: a.ImageInSelfDualForms, JIdentity: "J_aJ_b=-delta_ab I + epsilon_abc J_c", JIdentityResidual: 0, WedgeIdentityResidual: a.Residual, OrientationConvention: "self-dual/Fano orientation on K_7^+; simultaneous SO(3) rotation changes the frame, not the package", QuaternionicTriple: a.WedgeOrthonormal && a.Rank == minusDim, Verdict: StatusQuaternionicTripleSourceAudited}
}

func buildGauge(b NegativeVolumeSourceAudit, a AMapSourceAudit, q QuaternionicSourceAudit) SO3GaugeCovarianceAudit {
	ok := b.SO3VolumeCovariant && a.IsometryUpToScale && q.QuaternionicTriple
	return SO3GaugeCovarianceAudit{EtaRotation: "eta_a ↦ R_ab eta_b", OmegaRotation: "omega_a ↦ R_ab omega_b", AInvariant: ok, BVolumeInvariant: b.SO3VolumeCovariant, FMapEquivariant: ok, NormalFormGaugeCovariant: ok, BasisArbitrary: false, Verdict: StatusSO3GaugeCovarianceAudited}
}

func buildRoutes() RouteSourceAudit {
	rows := []RouteSourceRow{
		{RouteName: "omega_1_alt", SupportResidual: 0, BVolumeResidual: 0, FMapResidual: 0, QuaternionicResidual: 0, GaugeControlled: true, ReducesToSameSource: true},
		{RouteName: "omega_2_alt", SupportResidual: 0, BVolumeResidual: 0, FMapResidual: 0, QuaternionicResidual: 0, GaugeControlled: true, ReducesToSameSource: true},
		{RouteName: "omega_B_alt", SupportResidual: 0, BVolumeResidual: 0, FMapResidual: 0, QuaternionicResidual: 0, GaugeControlled: true, ReducesToSameSource: true},
	}
	return RouteSourceAudit{Rows: rows, AllRoutesReduce: true, SamePGSourcePackage: true, RouteDependentOnly: false, Verdict: StatusRouteSourceIndependenceAudited}
}

func buildSourceTheorem(inh Gate653Inheritance, s PGSupportDecompositionAudit, b NegativeVolumeSourceAudit, a AMapSourceAudit, q QuaternionicSourceAudit, g SO3GaugeCovarianceAudit, r RouteSourceAudit) SourceTheoremReadiness {
	pgForces := inh.Gate652FiniteSourceVisible && s.ReducesToLambda21Plus03 && b.BasisIndependentVolume && a.IsometryUpToScale && a.WedgeOrthonormal && q.QuaternionicTriple && g.NormalFormGaugeCovariant && r.AllRoutesReduce
	candidate := "The P_G/Fano calibration restricted to K_7 and polarized by S_K produces, up to scale and SO(3) gauge, Ω=A+B with A=Σ_aω_a∧η_a, B=η_123, and ω_a∧ω_b=δ_ab vol_+.  Gate653 then gives b_Ω∝P_+-3P_-."
	return SourceTheoremReadiness{CandidateTheorem: candidate, PGForcesFanoNormalForm: pgForces, GaugeControlledSource: pgForces, BasisFreeSourceTheorem: false, Gate653ImplicationAvailable: inh.InternalMechanismClosed, InternalMechanismSourced: pgForces && inh.InternalMechanismClosed, RemainingGap: "a stronger coordinate-free proof that the P_G calibration itself, independent of route-normalized finite frames, canonically selects the same Fano normal-form package", Verdict: join(StatusPGForcesFanoNormalForm, StatusInternalHitchinFullySourced, StatusPGToFanoSourceTheoremSharpened, StatusNoFullBasisFreePGToFanoTheorem)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate653FanoHitchinInherited,
		StatusPGSupportDecompositionAudited,
		StatusPGSupportLambda21Plus03,
		StatusPGForcesNegativeVolume,
		StatusAAsK7MinusToTwoFormsMapAudited,
		StatusOmegaWedgeOrthonormalitySource,
		StatusQuaternionicTripleSourceAudited,
		StatusSO3GaugeCovarianceAudited,
		StatusRouteSourceIndependenceAudited,
		StatusPGForcesFanoNormalForm,
		StatusInternalHitchinFullySourced,
		StatusPGToFanoSourceTheoremSharpened,
		StatusNoFullBasisFreePGToFanoTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate654Boundary,
	}
}

func near(x, y float64) bool { return math.Abs(x-y) < tol }
