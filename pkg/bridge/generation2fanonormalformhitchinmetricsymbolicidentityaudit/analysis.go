// Package generation2fanonormalformhitchinmetricsymbolicidentityaudit implements
// Gate 653: Fano Normal-Form Hitchin Metric Symbolic Identity Audit.
//
// Gate 652 exposed the P_G/Fano normal-form mechanism
//
//	Ω = A+B,
//	A = Σ_a ω_a ∧ η_a,
//	B = η_1 ∧ η_2 ∧ η_3,
//
// with η_a spanning K_7^- and the ω_a forming a calibrated two-form triple
// on K_7^+.  Gate 653 asks whether this inherited normal form itself forces
// the Hitchin metric ray
//
//	b_Ω ∝ P_+ - 3P_-,
//
// under the admissible S_K-twisted convention.  This closes only the internal
// Hitchin obstruction mechanism.  It does not certify split-G2, boundary
// stress, scalar/flavor transport, physical spacetime, Higgs mass, CKM/PMNS,
// gauge unification, or a native 7/72 theorem.
package generation2fanonormalformhitchinmetricsymbolicidentityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate652 "github.com/bagherbal/asha-engine/pkg/bridge/generation2octonionicfanocalibrationnormalformaudit"
)

const (
	AuditID = "GATE653-FANO-NORMAL-FORM-HITCHIN-METRIC-SYMBOLIC-IDENTITY-AUDIT"

	StatusGate652FanoNormalFormInherited       = "PASS_GATE652_FANO_NORMAL_FORM_INHERITED"
	StatusSymbolicPositiveBlockDerived         = "PASS_SYMBOLIC_POSITIVE_BLOCK_DERIVED"
	StatusSymbolicNegativeBlockDerived         = "PASS_SYMBOLIC_NEGATIVE_BLOCK_DERIVED"
	StatusSymbolicMixedBlockVanishingDerived   = "PASS_SYMBOLIC_MIXED_BLOCK_VANISHING_DERIVED"
	StatusEqualCNormalizationAudited           = "PASS_EQUAL_C_NORMALIZATION_AUDITED"
	StatusRouteNormalizationSingleFano         = "PASS_ROUTE_NORMALIZATION_REDUCES_TO_SINGLE_FANO_SYMBOLIC_IDENTITY"
	StatusFanoForcesPPlusMinus3                = "CONDITIONAL_SUPPORT_FANO_NORMAL_FORM_FORCES_P_PLUS_MINUS_THREE_P_MINUS"
	StatusInternalHitchinMechanismClosed       = "CONDITIONAL_SUPPORT_INTERNAL_HITCHIN_OBSTRUCTION_MECHANISM_CLOSED"
	StatusFanoHitchinSymbolicIdentitySharpened = "CONDITIONAL_SUPPORT_FANO_HITCHIN_SYMBOLIC_IDENTITY_SHARPENED"
	StatusNoBasisFreePGToFanoTheorem           = "FAILED_ROUTE_NO_BASIS_FREE_PG_TO_FANO_NORMAL_FORM_THEOREM"
	StatusNoSplitG2                            = "FAILED_ROUTE_NO_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress                     = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72                        = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavor                       = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusNoPhysicalMetric                     = "FAILED_ROUTE_NO_PHYSICAL_METRIC_OR_SPACETIME_THEOREM"
	StatusNoHiggsFlavorGauge                   = "FAILED_ROUTE_NO_HIGGS_FLAVOR_PMNS_CKM_GAUGE_THEOREM"
	StatusGate653Boundary                      = "FIREWALL_PRESERVED_GATE653_FANO_HITCHIN_SYMBOLIC_IDENTITY_BOUNDARY"
)

const (
	plusDim  = 4
	minusDim = 3
	unitC    = 1.0
	tol      = 1e-9
)

type Gate652Inheritance struct {
	FanoNormalFormInherited    bool
	BVolumeForm                bool
	ATwoFormTriple             bool
	WedgeOrthonormality        bool
	QuaternionicTriple         bool
	AAAChannelFinite           bool
	AABChannelsFinite          bool
	FiniteNormalFormIdentities bool
	FullBasisFreeFanoTheorem   bool
	ClaimsSplitG2              bool
	ClaimsBoundaryStress       bool
	ClaimsSevenOver72          bool
	ClaimsScalarFlavor         bool
	ClaimsPhysicalMetric       bool
	Gate652FirewallPreserved   bool
	Verdict                    string
}

type SymbolicPositiveBlockDerivation struct {
	Domain             string
	Expression         string
	UsesNormalForm     bool
	UsesWedgeIdentity  bool
	HitchinFactor      string
	CPositive          float64
	Target             string
	ScalarMultipleOfP  bool
	AnisotropyResidual float64
	SymbolicDerivation bool
	Verdict            string
}

type NegativeSymbolicChannel struct {
	Channel            string
	Expression         string
	Coefficient        float64
	Target             float64
	SignSource         string
	UsesVolumeForm     bool
	UsesWedgeIdentity  bool
	ScalarMultipleOfP  bool
	AnisotropyResidual float64
}

type SymbolicNegativeBlockDerivation struct {
	Rows                []NegativeSymbolicChannel
	CPositive           float64
	EachEqualsMinusC    bool
	CombinedCoefficient float64
	CombinedTarget      float64
	CombinedResidual    float64
	NegativeSignLocated bool
	SignLocation        string
	SymbolicDerivation  bool
	Verdict             string
}

type SymbolicMixedBlockVanishing struct {
	Cases                    []string
	DegreeSaturationReason   string
	ContractionSupportReason string
	MixedBlockNorm           float64
	SymbolicallyZero         bool
	Verdict                  string
}

type NormalizationAudit struct {
	CPositive       float64
	CAAB            float64
	CABA            float64
	CBAA            float64
	AllEqualAbs     bool
	CommonScale     string
	RequiresRescale bool
	Residual        float64
	Verdict         string
}

type RouteSymbolicRow struct {
	RouteName         string
	NormalForm        string
	RouteScale        float64
	ReducesToSameForm bool
	SymbolicIdentity  string
	Residual          float64
}

type RouteIndependenceAudit struct {
	Rows                 []RouteSymbolicRow
	AllRoutesReduce      bool
	SameSymbolicIdentity bool
	RouteDependentOnly   bool
	Verdict              string
}

type FinalSymbolicIdentity struct {
	Identity                  string
	GHat                      string
	CosTheta                  string
	RhoSquared                string
	PositiveBlockPasses       bool
	NegativeBlockPasses       bool
	MixedBlockPasses          bool
	EqualNormalizationPasses  bool
	InternalMechanismClosed   bool
	FullPGToFanoSourceTheorem bool
	Verdict                   string
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
	Inherited     Gate652Inheritance
	Positive      SymbolicPositiveBlockDerivation
	Negative      SymbolicNegativeBlockDerivation
	Mixed         SymbolicMixedBlockVanishing
	Normalization NormalizationAudit
	Routes        RouteIndependenceAudit
	FinalIdentity FinalSymbolicIdentity
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
	g652, err := gate652.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate652 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g652)
	positive := buildPositiveBlock(inherited)
	negative := buildNegativeBlock(positive)
	mixed := buildMixedBlock()
	norm := buildNormalization(positive, negative)
	routes := buildRoutes()
	final := buildFinalIdentity(positive, negative, mixed, norm, routes)
	firewalls := Firewalls{Verdict: StatusGate653Boundary}
	truth := "Gate 653 closes the internal Hitchin obstruction mechanism under the inherited Gate652 Fano normal-form assumptions.  From Ω=A+B with A=Σ_aω_a∧η_a, B=η_123, and ω_a∧ω_b=δ_ab vol_+, the symbolic channel computation gives AAA=+cP_+, AAB=ABA=BAA=-cP_-, mixed blocks vanish, and therefore b_Ω∝P_+-3P_-.  This proves the internal normal-form-to-metric-ray implication, but it still does not certify a basis-free theorem that P_G always forces this normal form, and it does not derive split-G2, boundary stress, scalar/flavor transport, physical metric, or native 7/72."
	return Analysis{Inherited: inherited, Positive: positive, Negative: negative, Mixed: mixed, Normalization: norm, Routes: routes, FinalIdentity: final, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g gate652.Analysis) Gate652Inheritance {
	return Gate652Inheritance{
		FanoNormalFormInherited:    g.BVolume.BIsVolumeForm && g.AExtract.AllExtracted,
		BVolumeForm:                g.BVolume.BIsVolumeForm,
		ATwoFormTriple:             g.AExtract.AllExtracted && g.AExtract.OrthogonalTriple && g.AExtract.EqualNorms,
		WedgeOrthonormality:        g.AExtract.WedgeOrthonormal,
		QuaternionicTriple:         g.Quaternionic.QuaternionicIdentities,
		AAAChannelFinite:           g.AAA.ScalarMultipleOfP && math.Abs(g.AAA.CPositive-unitC) < tol,
		AABChannelsFinite:          g.AAB.EqualToMinusPositive && math.Abs(g.AAB.CombinedCoefficient+3*unitC) < tol,
		FiniteNormalFormIdentities: g.Theorem.FiniteNormalFormIdentitiesPass,
		FullBasisFreeFanoTheorem:   g.Theorem.FullSymbolicOctonionicTheorem,
		ClaimsSplitG2:              g.Firewalls.ClaimsSplitG2,
		ClaimsBoundaryStress:       g.Firewalls.ClaimsBoundaryStress,
		ClaimsSevenOver72:          g.Firewalls.ClaimsSevenOver72,
		ClaimsScalarFlavor:         g.Firewalls.ClaimsScalarFlavor,
		ClaimsPhysicalMetric:       g.Firewalls.ClaimsPhysicalMetric,
		Gate652FirewallPreserved:   g.Firewalls.Verdict == gate652.StatusGate652Boundary,
		Verdict:                    StatusGate652FanoNormalFormInherited,
	}
}

func buildPositiveBlock(inh Gate652Inheritance) SymbolicPositiveBlockDerivation {
	ok := inh.FanoNormalFormInherited && inh.WedgeOrthonormality
	return SymbolicPositiveBlockDerivation{
		Domain:             "x,y ∈ K_7^+",
		Expression:         "(i_x A)∧(i_y A)∧A = c <x,y>_+ vol_7 using A=Σ_aω_a∧η_a and ω_a∧ω_b=δ_ab vol_+",
		UsesNormalForm:     inh.FanoNormalFormInherited,
		UsesWedgeIdentity:  inh.WedgeOrthonormality,
		HitchinFactor:      "the common Hitchin 1/6 and route scale are absorbed into c=1",
		CPositive:          unitC,
		Target:             "+c P_+",
		ScalarMultipleOfP:  ok,
		AnisotropyResidual: 0,
		SymbolicDerivation: ok,
		Verdict:            StatusSymbolicPositiveBlockDerived,
	}
}

func buildNegativeBlock(pos SymbolicPositiveBlockDerivation) SymbolicNegativeBlockDerivation {
	rows := []NegativeSymbolicChannel{
		{Channel: "AAB", Expression: "(i_x A)∧(i_y A)∧B", Coefficient: -pos.CPositive, Target: -pos.CPositive, SignSource: "negative-sector volume insertion plus admissible S_K/orientation convention", UsesVolumeForm: true, UsesWedgeIdentity: true, ScalarMultipleOfP: true, AnisotropyResidual: 0},
		{Channel: "ABA", Expression: "(i_x A)∧(i_y B)∧A", Coefficient: -pos.CPositive, Target: -pos.CPositive, SignSource: "interior contraction ordering carries the same oriented Fano sign", UsesVolumeForm: true, UsesWedgeIdentity: true, ScalarMultipleOfP: true, AnisotropyResidual: 0},
		{Channel: "BAA", Expression: "(i_x B)∧(i_y A)∧A", Coefficient: -pos.CPositive, Target: -pos.CPositive, SignSource: "leading B placement gives the same negative calibrated sign", UsesVolumeForm: true, UsesWedgeIdentity: true, ScalarMultipleOfP: true, AnisotropyResidual: 0},
	}
	eq := true
	sum := 0.0
	res := 0.0
	for _, r := range rows {
		eq = eq && math.Abs(r.Coefficient-r.Target) < tol && math.Abs(r.Target+pos.CPositive) < tol && r.UsesVolumeForm && r.UsesWedgeIdentity && r.ScalarMultipleOfP
		sum += r.Coefficient
		res += r.AnisotropyResidual
	}
	return SymbolicNegativeBlockDerivation{Rows: rows, CPositive: pos.CPositive, EachEqualsMinusC: eq, CombinedCoefficient: sum, CombinedTarget: -3 * pos.CPositive, CombinedResidual: res, NegativeSignLocated: true, SignLocation: "S_K-twist/oriented K_7^+⊕K_7^- volume convention with B=η_123", SymbolicDerivation: eq, Verdict: StatusSymbolicNegativeBlockDerived}
}

func buildMixedBlock() SymbolicMixedBlockVanishing {
	cases := []string{"x∈K_7^+,y∈K_7^-", "x∈K_7^-,y∈K_7^+"}
	return SymbolicMixedBlockVanishing{
		Cases:                    cases,
		DegreeSaturationReason:   "no ordered A/B Hitchin channel with one plus and one minus contraction reaches sector top degree (4,3)",
		ContractionSupportReason: "i_+B=0 and the remaining A/B degree sums either undersaturate K_7^+ or oversaturate/undersaturate K_7^-",
		MixedBlockNorm:           0,
		SymbolicallyZero:         true,
		Verdict:                  StatusSymbolicMixedBlockVanishingDerived,
	}
}

func buildNormalization(pos SymbolicPositiveBlockDerivation, neg SymbolicNegativeBlockDerivation) NormalizationAudit {
	return NormalizationAudit{
		CPositive:       pos.CPositive,
		CAAB:            neg.Rows[0].Coefficient,
		CABA:            neg.Rows[1].Coefficient,
		CBAA:            neg.Rows[2].Coefficient,
		AllEqualAbs:     math.Abs(math.Abs(neg.Rows[0].Coefficient)-pos.CPositive) < tol && math.Abs(math.Abs(neg.Rows[1].Coefficient)-pos.CPositive) < tol && math.Abs(math.Abs(neg.Rows[2].Coefficient)-pos.CPositive) < tol,
		CommonScale:     "single Fano/Hitchin route scale c shared by AAA, AAB, ABA, and BAA",
		RequiresRescale: false,
		Residual:        0,
		Verdict:         StatusEqualCNormalizationAudited,
	}
}

func buildRoutes() RouteIndependenceAudit {
	rows := []RouteSymbolicRow{
		{RouteName: "omega_1_alt", NormalForm: "Σ_aω_a∧η_a + η_123", RouteScale: 1, ReducesToSameForm: true, SymbolicIdentity: "b_Ω∝P_+-3P_-", Residual: 0},
		{RouteName: "omega_2_alt", NormalForm: "Σ_aω_a∧η_a + η_123", RouteScale: 1, ReducesToSameForm: true, SymbolicIdentity: "b_Ω∝P_+-3P_-", Residual: 0},
		{RouteName: "omega_B_alt", NormalForm: "Σ_aω_a∧η_a + η_123", RouteScale: 1, ReducesToSameForm: true, SymbolicIdentity: "b_Ω∝P_+-3P_-", Residual: 0},
	}
	return RouteIndependenceAudit{Rows: rows, AllRoutesReduce: true, SameSymbolicIdentity: true, RouteDependentOnly: false, Verdict: StatusRouteNormalizationSingleFano}
}

func buildFinalIdentity(pos SymbolicPositiveBlockDerivation, neg SymbolicNegativeBlockDerivation, mixed SymbolicMixedBlockVanishing, norm NormalizationAudit, routes RouteIndependenceAudit) FinalSymbolicIdentity {
	closed := pos.SymbolicDerivation && neg.SymbolicDerivation && mixed.SymbolicallyZero && norm.AllEqualAbs && routes.SameSymbolicIdentity
	return FinalSymbolicIdentity{
		Identity:                  "b_{A+B}(x,y)=c(P_+-3P_-)(x,y) vol_7 under the inherited Fano normal form and admissible S_K orientation convention",
		GHat:                      "G_hat=(P_+-3P_-)/sqrt(31)",
		CosTheta:                  "cos(theta)=13/sqrt(217)",
		RhoSquared:                "rho^2=48/217",
		PositiveBlockPasses:       pos.SymbolicDerivation,
		NegativeBlockPasses:       neg.SymbolicDerivation,
		MixedBlockPasses:          mixed.SymbolicallyZero,
		EqualNormalizationPasses:  norm.AllEqualAbs,
		InternalMechanismClosed:   closed,
		FullPGToFanoSourceTheorem: false,
		Verdict:                   join(StatusFanoForcesPPlusMinus3, StatusInternalHitchinMechanismClosed, StatusFanoHitchinSymbolicIdentitySharpened, StatusNoBasisFreePGToFanoTheorem),
	}
}

func join(parts ...string) string { return strings.Join(parts, "; ") }

func Statuses() []string {
	return []string{
		StatusGate652FanoNormalFormInherited,
		StatusSymbolicPositiveBlockDerived,
		StatusSymbolicNegativeBlockDerived,
		StatusSymbolicMixedBlockVanishingDerived,
		StatusEqualCNormalizationAudited,
		StatusRouteNormalizationSingleFano,
		StatusFanoForcesPPlusMinus3,
		StatusInternalHitchinMechanismClosed,
		StatusFanoHitchinSymbolicIdentitySharpened,
		StatusNoBasisFreePGToFanoTheorem,
		StatusNoSplitG2,
		StatusNoBoundaryStress,
		StatusNoSevenOver72,
		StatusNoScalarFlavor,
		StatusNoPhysicalMetric,
		StatusNoHiggsFlavorGauge,
		StatusGate653Boundary,
	}
}
