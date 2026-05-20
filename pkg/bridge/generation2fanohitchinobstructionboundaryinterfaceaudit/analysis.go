// Package generation2fanohitchinobstructionboundaryinterfaceaudit implements
// Gate 655: Fano-Hitchin Obstruction Boundary-Interface Audit.
//
// Gate 654 conditionally closed the internal finite Fano/Hitchin mechanism
//
//	P_G + S_K => Ω_Fano => b_Ω ∝ P_+ - 3P_-
//	=> cos(theta)=13/sqrt(217), rho^2=48/217.
//
// Gate 655 asks whether that mature internal obstruction package supplies a
// lawful boundary-facing invariant for the Gate626/Gate628 7/72 trace-weight
// clue, the Gate613 boundary-stress seal, the Gate623 history-loop unit, or
// the flavor orientation seal.  It is a boundary-interface audit only: it does
// not derive boundary stress, scalar/flavor transport, physical spacetime,
// Higgs mass, CKM/PMNS, gauge unification, split-G2, or a native 7/72 theorem.
package generation2fanohitchinobstructionboundaryinterfaceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate654 "github.com/bagherbal/asha-engine/pkg/bridge/generation2pgtofanonormalformsourcetheoremaudit"
)

const (
	AuditID = "GATE655-FANO-HITCHIN-OBSTRUCTION-BOUNDARY-INTERFACE-AUDIT"

	StatusGate654InternalHitchinInherited = "PASS_GATE654_INTERNAL_HITCHIN_MECHANISM_INHERITED"
	StatusInternalInvariantLedgerBuilt    = "PASS_INTERNAL_INVARIANT_LEDGER_CONSTRUCTED"
	StatusSevenOver72InterfaceAudited     = "PASS_7_OVER_72_INTERFACE_AUDITED"
	StatusBoundaryStressInterfaceAudited  = "PASS_BOUNDARY_STRESS_INTERFACE_AUDITED"
	StatusHistoryLoopInterfaceAudited     = "PASS_HISTORY_LOOP_UNIT_INTERFACE_AUDITED"
	StatusFlavorOrientationAudited        = "PASS_FLAVOR_ORIENTATION_INTERFACE_AUDITED"
	StatusBoundaryMapObstructionAudited   = "PASS_BOUNDARY_MAP_OBSTRUCTION_AUDITED"
	StatusFanoStructuresNumerator7        = "CONDITIONAL_SUPPORT_FANO_HITCHIN_PACKAGE_STRUCTURES_NUMERATOR_7"
	StatusFanoHitchinSealDefined          = "CONDITIONAL_SUPPORT_FANO_HITCHIN_OBSTRUCTION_SEAL_DEFINED"
	StatusNoBoundaryInterface             = "FAILED_ROUTE_NO_BOUNDARY_INTERFACE_FROM_FANO_HITCHIN_PACKAGE"
	StatusNoSevenOver72Theorem            = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoBoundaryStress                = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoScalarFlavorMap               = "FAILED_ROUTE_NO_SCALAR_FLAVOR_TRANSPORT_MAP"
	StatusNoHistoryLoopSource             = "FAILED_ROUTE_NO_HISTORY_LOOP_UNIT_SOURCE_FROM_FANO_HITCHIN"
	StatusNoSplitG2                       = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoPhysicalMetric                = "FAILED_ROUTE_NO_PHYSICAL_METRIC_OR_SPACETIME_THEOREM"
	StatusNoHiggsCKMGauge                 = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate655Boundary                 = "FIREWALL_PRESERVED_GATE655_FANO_HITCHIN_BOUNDARY_INTERFACE_BOUNDARY"
)

const (
	plusDim  = 4
	minusDim = 3
	k7Dim    = plusDim + minusDim
	lambda4  = 70
	boundary = 2
	chamber  = lambda4 + boundary
	tol      = 1e-9

	xiBoundary = 0.0503471644870914
	r3Minus1   = 0.0509933868964996
	absLambda  = 0.0497009420776833
	LHistory   = 0.0397887357729738
)

type Gate654Inheritance struct {
	InternalMechanismSourced bool
	PGForcesFanoNormalForm   bool
	GaugeControlledSource    bool
	BasisFreeSourceTheorem   bool
	HitchinRay               string
	CosTheta                 float64
	RhoSquared               float64
	ClaimsSplitG2            bool
	ClaimsBoundaryStress     bool
	ClaimsSevenOver72        bool
	ClaimsScalarFlavor       bool
	ClaimsPhysicalMetric     bool
	Gate654FirewallPreserved bool
	Verdict                  string
}

type InternalInvariant struct {
	Name             string
	Value            float64
	Expression       string
	Classification   string
	GaugeInvariant   bool
	OrientationBound bool
}

type InternalInvariantLedger struct {
	Rows                []InternalInvariant
	TraceSK             float64
	TraceGUn            float64
	Norm2SK             float64
	Norm2GUn            float64
	DetGUn              float64
	ProjectiveInner     float64
	ObstructionSquare   float64
	RankK7              int
	RankPlus            int
	RankMinus           int
	SO3GaugeDim         int
	FanoTripleCount     int
	ChannelCount        string
	AllNativeFinite     bool
	AllGaugeClassified  bool
	BoundaryDataPresent bool
	Verdict             string
}

type SevenOver72InterfaceAudit struct {
	CandidateWeight             float64
	NumeratorSources            []string
	DenominatorCandidate        string
	FanoAddsBeyondNumerator     bool
	BoundaryPairSupplied        bool
	TraceMapSupplied            bool
	StructuresNumerator7        bool
	CertifiedSevenOver72Theorem bool
	Verdict                     string
}

type BoundaryStressCandidate struct {
	Candidate      string
	Value          float64
	ClosestSeal    string
	ClosestValue   float64
	AbsResidual    float64
	RelResidual    float64
	Classification string
}

type BoundaryStressInterfaceAudit struct {
	Rows                          []BoundaryStressCandidate
	CertifiedBoundaryStressSource bool
	NearBridgeClueOnly            bool
	NoArbitrarySearch             bool
	Verdict                       string
}

type HistoryLoopUnitInterfaceAudit struct {
	TargetL                float64
	SuppliesPiOrS1         bool
	SuppliesHeatKernel     bool
	SuppliesAngularMeasure bool
	FiniteAlgebraicOnly    bool
	CertifiedSource        bool
	Verdict                string
}

type FlavorOrientationInterfaceAudit struct {
	Targets                             []string
	UsesFlavorData                      bool
	TypedIntertwinerSupplied            bool
	ObstructionAngleMappedToFlavor      bool
	RejectsNumericalProximityWithoutMap bool
	CertifiedFlavorMap                  bool
	Verdict                             string
}

type BoundaryMapObstructionAudit struct {
	MissingPsi            string
	MissingTau            string
	HasPsi                bool
	HasTau                bool
	CanAssignBoundaryPair bool
	CanAssignSevenOver72  bool
	CanAssignScalarFlavor bool
	Verdict               string
}

type FanoHitchinObstructionSeal struct {
	Name             string
	Carrier          string
	Split            string
	Source           string
	NormalForm       string
	HitchinMetricRay string
	CosTheta         string
	ResidualSquare   string
	BoundaryStatus   string
	InternalOnly     bool
	Verdict          string
}

type Firewalls struct {
	ClaimsSplitG2          bool
	ClaimsBoundaryStress   bool
	ClaimsSevenOver72      bool
	ClaimsScalarFlavor     bool
	ClaimsHistoryLoopUnit  bool
	ClaimsPhysicalMetric   bool
	ClaimsHiggsMass        bool
	ClaimsCKMPMNS          bool
	ClaimsGaugeUnification bool
	Verdict                string
}

type Analysis struct {
	Inherited      Gate654Inheritance
	Invariants     InternalInvariantLedger
	SevenOver72    SevenOver72InterfaceAudit
	BoundaryStress BoundaryStressInterfaceAudit
	HistoryLoop    HistoryLoopUnitInterfaceAudit
	Flavor         FlavorOrientationInterfaceAudit
	BoundaryMap    BoundaryMapObstructionAudit
	Seal           FanoHitchinObstructionSeal
	Firewalls      Firewalls
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
	g654, err := gate654.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate654 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g654)
	invariants := buildInvariants(inherited)
	seven := buildSevenOver72(invariants)
	boundaryStress := buildBoundaryStress()
	history := buildHistoryLoop()
	flavor := buildFlavor()
	boundaryMap := buildBoundaryMap()
	seal := buildSeal(inherited, invariants, seven, boundaryMap)
	firewalls := Firewalls{Verdict: StatusGate655Boundary}
	truth := "Gate 655 seals the Fano-Hitchin lane as a mature internal obstruction package unless a new boundary map is constructed.  The package structures the numerator 7 through the full K_7 Fano/Hitchin carrier and records native invariants such as trace(S_K)=1, ||P_+-3P_-||^2=31, cos(theta)=13/sqrt(217), and rho^2=48/217.  It does not supply the R^2_boundary assignment, a normalized trace 7/72 theorem, a HistoryLoopUnit source, or a scalar/flavor transport map."
	return Analysis{Inherited: inherited, Invariants: invariants, SevenOver72: seven, BoundaryStress: boundaryStress, HistoryLoop: history, Flavor: flavor, BoundaryMap: boundaryMap, Seal: seal, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate654.Analysis) Gate654Inheritance {
	return Gate654Inheritance{
		InternalMechanismSourced: g.SourceTheorem.InternalMechanismSourced,
		PGForcesFanoNormalForm:   g.SourceTheorem.PGForcesFanoNormalForm,
		GaugeControlledSource:    g.SourceTheorem.GaugeControlledSource,
		BasisFreeSourceTheorem:   g.SourceTheorem.BasisFreeSourceTheorem,
		HitchinRay:               "P_+ - 3P_-",
		CosTheta:                 13 / math.Sqrt(217),
		RhoSquared:               48.0 / 217.0,
		ClaimsSplitG2:            g.Firewalls.ClaimsSplitG2,
		ClaimsBoundaryStress:     g.Firewalls.ClaimsBoundaryStress,
		ClaimsSevenOver72:        g.Firewalls.ClaimsSevenOver72,
		ClaimsScalarFlavor:       g.Firewalls.ClaimsScalarFlavor,
		ClaimsPhysicalMetric:     g.Firewalls.ClaimsPhysicalMetric,
		Gate654FirewallPreserved: g.Firewalls.Verdict == gate654.StatusGate654Boundary,
		Verdict:                  StatusGate654InternalHitchinInherited,
	}
}

func buildInvariants(inh Gate654Inheritance) InternalInvariantLedger {
	traceSK := float64(plusDim - minusDim)
	traceG := float64(plusDim - 3*minusDim)
	norm2SK := float64(k7Dim)
	norm2G := float64(plusDim + 9*minusDim)
	detG := -27.0
	inner := inh.CosTheta
	rho2 := inh.RhoSquared
	rows := []InternalInvariant{
		{Name: "trace(S_K)", Value: traceSK, Expression: "4-3", Classification: "native finite Hodge imbalance", GaugeInvariant: true},
		{Name: "trace(G_un)", Value: traceG, Expression: "tr(P_+-3P_-)=4-9", Classification: "native finite Hitchin obstruction trace", GaugeInvariant: true},
		{Name: "||S_K||_F^2", Value: norm2SK, Expression: "rank(K_7)=7", Classification: "native finite norm", GaugeInvariant: true},
		{Name: "||G_un||_F^2", Value: norm2G, Expression: "4+9*3=31", Classification: "native finite Hitchin metric-ray norm", GaugeInvariant: true},
		{Name: "det(G_un)", Value: detG, Expression: "1^4*(-3)^3", Classification: "orientation-sensitive finite determinant", GaugeInvariant: true, OrientationBound: true},
		{Name: "<G_hat,B_hat>", Value: inner, Expression: "13/sqrt(217)", Classification: "internal projective obstruction alignment", GaugeInvariant: true},
		{Name: "rho^2", Value: rho2, Expression: "48/217", Classification: "internal compact/split obstruction square", GaugeInvariant: true},
		{Name: "rank(K_7)", Value: k7Dim, Expression: "4+3", Classification: "native finite Fano-Hitchin carrier dimension", GaugeInvariant: true},
		{Name: "SO(3) gauge dimension", Value: 3, Expression: "dim so(3)", Classification: "normal-form gauge freedom", GaugeInvariant: true},
		{Name: "Fano triple count", Value: 3, Expression: "omega_1,omega_2,omega_3", Classification: "quaternionic/Fano triple count", GaugeInvariant: true},
	}
	return InternalInvariantLedger{Rows: rows, TraceSK: traceSK, TraceGUn: traceG, Norm2SK: norm2SK, Norm2GUn: norm2G, DetGUn: detG, ProjectiveInner: inner, ObstructionSquare: rho2, RankK7: k7Dim, RankPlus: plusDim, RankMinus: minusDim, SO3GaugeDim: 3, FanoTripleCount: 3, ChannelCount: "1 positive + 3 negative", AllNativeFinite: true, AllGaugeClassified: true, BoundaryDataPresent: false, Verdict: StatusInternalInvariantLedgerBuilt}
}

func buildSevenOver72(inv InternalInvariantLedger) SevenOver72InterfaceAudit {
	return SevenOver72InterfaceAudit{
		CandidateWeight:             float64(k7Dim) / float64(chamber),
		NumeratorSources:            []string{"dim(K_7)=7", "rank(P_+)+rank(P_-)=7", "Fano-Hitchin carrier dimension=7"},
		DenominatorCandidate:        "72 = dim(Lambda^4 R^8)+dim(R^2_boundary)=70+2 inherited from Gates 628-630",
		FanoAddsBeyondNumerator:     inv.RankK7 == k7Dim,
		BoundaryPairSupplied:        false,
		TraceMapSupplied:            false,
		StructuresNumerator7:        true,
		CertifiedSevenOver72Theorem: false,
		Verdict:                     join(StatusSevenOver72InterfaceAudited, StatusFanoStructuresNumerator7, StatusNoSevenOver72Theorem),
	}
}

func buildBoundaryStress() BoundaryStressInterfaceAudit {
	candidates := []struct {
		name  string
		value float64
	}{
		{"7/72", 7.0 / 72.0},
		{"7/144", 7.0 / 144.0},
		{"1/sqrt(217)", 1 / math.Sqrt(217)},
		{"13/217", 13.0 / 217.0},
		{"48/217", 48.0 / 217.0},
	}
	rows := make([]BoundaryStressCandidate, 0, len(candidates))
	for _, c := range candidates {
		seal, val, absResidual := closestBoundarySeal(c.value)
		classification := "failed numeric coincidence; typed internal invariant but no boundary map"
		if c.name == "7/144" {
			classification = "near bridge clue only; closest to |lambda(Lambda_12)| but no certified source"
		}
		rows = append(rows, BoundaryStressCandidate{Candidate: c.name, Value: c.value, ClosestSeal: seal, ClosestValue: val, AbsResidual: absResidual, RelResidual: absResidual / math.Abs(val), Classification: classification})
	}
	return BoundaryStressInterfaceAudit{Rows: rows, CertifiedBoundaryStressSource: false, NearBridgeClueOnly: true, NoArbitrarySearch: true, Verdict: join(StatusBoundaryStressInterfaceAudited, StatusNoBoundaryStress)}
}

func closestBoundarySeal(x float64) (string, float64, float64) {
	seals := []struct {
		name string
		val  float64
	}{
		{"xi_boundary", xiBoundary},
		{"R_3-1", r3Minus1},
		{"|lambda(Lambda_12)|", absLambda},
	}
	bestName, bestVal, bestAbs := seals[0].name, seals[0].val, math.Abs(x-seals[0].val)
	for _, s := range seals[1:] {
		if d := math.Abs(x - s.val); d < bestAbs {
			bestName, bestVal, bestAbs = s.name, s.val, d
		}
	}
	return bestName, bestVal, bestAbs
}

func buildHistoryLoop() HistoryLoopUnitInterfaceAudit {
	return HistoryLoopUnitInterfaceAudit{TargetL: LHistory, SuppliesPiOrS1: false, SuppliesHeatKernel: false, SuppliesAngularMeasure: false, FiniteAlgebraicOnly: true, CertifiedSource: false, Verdict: join(StatusHistoryLoopInterfaceAudited, StatusNoHistoryLoopSource)}
}

func buildFlavor() FlavorOrientationInterfaceAudit {
	return FlavorOrientationInterfaceAudit{Targets: []string{"epsilon_e", "kappa_e", "sin^2(theta13)/4", "J_CKM", "B_flav"}, UsesFlavorData: false, TypedIntertwinerSupplied: false, ObstructionAngleMappedToFlavor: false, RejectsNumericalProximityWithoutMap: true, CertifiedFlavorMap: false, Verdict: join(StatusFlavorOrientationAudited, StatusNoScalarFlavorMap)}
}

func buildBoundaryMap() BoundaryMapObstructionAudit {
	return BoundaryMapObstructionAudit{MissingPsi: "Psi: K_7 or FanoHitchinPackage -> R^2_boundary", MissingTau: "tau_defect: FanoHitchinPackage -> scalar trace weight with normalized trace 7/72", HasPsi: false, HasTau: false, CanAssignBoundaryPair: false, CanAssignSevenOver72: false, CanAssignScalarFlavor: false, Verdict: join(StatusBoundaryMapObstructionAudited, StatusNoBoundaryInterface)}
}

func buildSeal(inh Gate654Inheritance, inv InternalInvariantLedger, seven SevenOver72InterfaceAudit, b BoundaryMapObstructionAudit) FanoHitchinObstructionSeal {
	internalOnly := inh.InternalMechanismSourced && inv.AllNativeFinite && seven.StructuresNumerator7 && !seven.CertifiedSevenOver72Theorem && !b.HasPsi && !b.HasTau
	return FanoHitchinObstructionSeal{Name: "FanoHitchinObstructionSeal", Carrier: "K_7", Split: "K_7^+|K_7^- = 4|3", Source: "P_G + S_K", NormalForm: "Omega=sum_a omega_a wedge eta_a + eta_123", HitchinMetricRay: "P_+ - 3P_-", CosTheta: "13/sqrt(217)", ResidualSquare: "48/217", BoundaryStatus: "internal only; no boundary interface", InternalOnly: internalOnly, Verdict: join(StatusFanoHitchinSealDefined, StatusNoBoundaryInterface)}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate654InternalHitchinInherited,
		StatusInternalInvariantLedgerBuilt,
		StatusSevenOver72InterfaceAudited,
		StatusBoundaryStressInterfaceAudited,
		StatusHistoryLoopInterfaceAudited,
		StatusFlavorOrientationAudited,
		StatusBoundaryMapObstructionAudited,
		StatusFanoStructuresNumerator7,
		StatusFanoHitchinSealDefined,
		StatusNoBoundaryInterface,
		StatusNoSevenOver72Theorem,
		StatusNoBoundaryStress,
		StatusNoScalarFlavorMap,
		StatusNoHistoryLoopSource,
		StatusNoSplitG2,
		StatusNoPhysicalMetric,
		StatusNoHiggsCKMGauge,
		StatusGate655Boundary,
	}
}

func near(x, y float64) bool { return math.Abs(x-y) < tol }
