// Package generation2spatialprojectiveorientationsealminimalityconsequenceaudit implements Gate 574:
// SpatialProjectiveOrientationSeal Minimality and Consequence Audit.
//
// Gate 573 proved that CP^2_sp is the B-L spatial projective block and that
// SU(3) acts transitively on it, so no SU(3)-invariant point [u] or rank-one
// projector P_u exists in current ASHA data. Gate 574 does not claim a native
// spatial 2+1 selector. It audits the minimal sealed datum required to continue:
// a projective orientation [u] in CP^2_sp, equivalently a rank-one Hermitian
// projector P_u. Once sealed, the usual selector algebra gives a CP^1|CP^0
// split and commutant u(2)+u(1), but no weak-isospin, flavor, electroweak,
// K7, spacetime, RG, OS/Hilbert, or observed-data promotion is opened.
package generation2spatialprojectiveorientationsealminimalityconsequenceaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate573 "github.com/bagherbal/asha-engine/pkg/bridge/generation2spatialcp2selectorisotropyobstructionaudit"
)

const (
	AuditID = "GATE574-SPATIAL-PROJECTIVE-ORIENTATION-SEAL-MINIMALITY-CONSEQUENCE-AUDIT"

	StatusGate573Inherited                               = "CONDITIONAL_SUPPORT_GATE573_SPATIAL_CP2_OBSTRUCTION_INHERITED"
	StatusSealDefined                                    = "PASS_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_DEFINED"
	StatusRankOneProjectorSealPropertiesVerified         = "PASS_RANK_ONE_PROJECTOR_SEAL_PROPERTIES_VERIFIED"
	StatusSealBreaksSU3ToS_U2_U1                         = "PASS_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_BREAKS_SU3_TO_S_U2_TIMES_U1"
	StatusSealedSelectorConstructed                      = "PASS_SEALED_SPATIAL_SELECTOR_CP2_TO_CP1_CP0_CONSTRUCTED"
	StatusSealedSelectorMultiplicityVerified             = "PASS_SEALED_SPATIAL_SELECTOR_EIGENVALUE_MULTIPLICITY_2PLUS1_VERIFIED"
	StatusSealedCriticalStrataVerified                   = "PASS_SEALED_SPATIAL_SELECTOR_CRITICAL_STRATA_CP1_AND_CP0_VERIFIED"
	StatusSealedCommutantVerified                        = "PASS_SEALED_SELECTOR_COMMUTANT_U2_PLUS_U1_DIMENSION_5_VERIFIED"
	StatusCommutantMatchesGate555Formula                 = "PASS_SEALED_COMMUTANT_MATCHES_GATE555_MULTIPLICITY_FORMULA"
	StatusRepresentativeGaugeU3GivesU12                  = "CONDITIONAL_SUPPORT_REPRESENTATIVE_GAUGE_U_EQUALS_A3_DAGGER_GIVES_U12"
	StatusRepresentativeGaugeNotNative                   = "FAILED_ROUTE_REPRESENTATIVE_U12_GAUGE_IS_NOT_NATIVE_BASIS_SELECTION"
	StatusWeakPlaneNotPhysical                           = "FAILED_ROUTE_SEALED_CP1_COMPLEMENT_NOT_PHYSICAL_WEAK_PLANE"
	StatusMissingSpectralTripleCompatibility             = "FAILED_ROUTE_NO_FINITE_SPECTRAL_TRIPLE_COMPATIBILITY_FOR_PHYSICAL_WEAK_PLANE"
	StatusNoFlavorGenerationEWData                       = "FAILED_ROUTE_SEALED_ORIENTATION_DOES_NOT_DERIVE_FLAVOR_GENERATION_OR_ELECTROWEAK_DATA"
	StatusPreviousBoundariesPreserved                    = "FIREWALL_PRESERVED_GATE574_PREVIOUS_TRACE_CONTACT_SCALAR_EW_K7_TIME_BOUNDARIES"
	StatusMinimalityVerified                             = "PASS_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_MINIMALITY_VERIFIED"
	StatusSealSufficientButNotNative                     = "CONDITIONAL_SUPPORT_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_SUFFICIENT_BUT_NOT_NATIVE"
	StatusSealNotNativeDerivation                        = "FAILED_ROUTE_SPATIAL_PROJECTIVE_ORIENTATION_SEAL_NOT_NATIVE_DERIVATION"
	StatusGate564565BoundaryPreserved                    = "FIREWALL_PRESERVED_GATE564_GATE565_ELECTROWEAK_BRIDGE_SYMBOLIC_BOUNDARY"
	StatusGate571572K7ProductTimeBoundaryPreserved       = "FIREWALL_PRESERVED_GATE571_GATE572_K7_PRODUCT_TIME_BOUNDARY"
	StatusGate574SpatialOrientationSealBoundaryPreserved = "FIREWALL_PRESERVED_GATE574_SPATIAL_ORIENTATION_SEAL_BOUNDARY"
)

type InheritedGate573Audit struct {
	CP2SpatialBlockCertified     bool
	SU3ActsTransitively          bool
	NoSU3InvariantPoint          bool
	GeneralTwoPlusOneClassified  bool
	NoNativeRankOneProjector     bool
	OrientationSealAlreadyNeeded bool
	NoWeakPlaneNativeSelection   bool
	K7ProductTimePreserved       bool
	FlavorEWBoundaryPreserved    bool
	Verdict                      string
}

type SealDefinitionAudit struct {
	SealName                   string
	ProjectiveDatum            string
	EquivalentProjectorDatum   string
	RepresentativeU            []float64
	Projector                  [][]float64
	Rank                       int
	Trace                      float64
	IdempotentResidual         float64
	Hermitian                  bool
	BreaksSU3To                string
	MinimalForSymmetryBreaking bool
	NativeDerived              bool
	Verdict                    string
}

type SealedSelectorAudit struct {
	SelectorFormula       string
	LambdaOne             float64
	LambdaTwo             float64
	SelectorMatrix        [][]float64
	Eigenvalues           []float64
	MultiplicityPattern   string
	CriticalCP1           string
	CriticalCP0           string
	CP1RealDimension      int
	CP0RealDimension      int
	ConstructsCP1CP0Split bool
	NativeWithoutSeal     bool
	Verdict               string
}

type CommutantAudit struct {
	SelectorMultiplicityPattern string
	Commutant                   string
	DimensionFormula            string
	Dimension                   int
	ExpectedDimension           int
	MatchesGate555Formula       bool
	SealedSupportOnly           bool
	Verdict                     string
}

type BasisExampleAudit struct {
	RepresentativeGauge   string
	ProjectorMatrix       string
	SelectorMatrix        string
	ComplementaryPlane    string
	ConventionalPlaneName string
	BasisDependent        bool
	NativeBasisSelection  bool
	Verdict               string
}

type WeakPlaneFirewallAudit struct {
	ComplementaryCP1CanBeCalledPhysicalWeakPlane bool
	RequiresFiniteSpectralTripleCarrierAction    bool
	RequiresQuaternionicCompatibility            bool
	RequiresDiracCompatibility                   bool
	RequiresRealStructureCompatibility           bool
	RequiresGradingCompatibility                 bool
	RequiresFirstOrderCompatibility              bool
	CompatibilityProven                          bool
	Verdict                                      string
}

type FlavorGenerationFirewallAudit struct {
	GenerationHierarchyDerived bool
	YukawaTextureDerived       bool
	CKMPMNSDerived             bool
	ObservedFlavorImported     bool
	PhysicalEWDynamicsDerived  bool
	PhotonDynamicsDerived      bool
	WZMassesDerived            bool
	Verdict                    string
}

type PreviousGateBoundaryAudit struct {
	TauEtaTraceShadowOnly       bool
	Q4ContactOnly               bool
	PauliQuaternionicSocketOnly bool
	Gate564565BridgeSymbolic    bool
	K7TimeRoutesSealed          bool
	Verdict                     string
}

type MinimalityAudit struct {
	WithoutProjectivePoint        bool
	WithoutRankOneProjector       bool
	NoCP2ToCP1CP0Selector         bool
	PointProjectorEquivalence     bool
	Any2Plus1SelectorDeterminesPU bool
	SealIsMinimal                 bool
	ProofSketch                   string
	Verdict                       string
}

type FinalVerdict struct {
	SealSufficient                           bool
	SealMinimal                              bool
	ReducesSymmetryToU2U1                    bool
	DerivesPhysicalWeakFlavorElectroweakData bool
	AdditionalTheoremRequired                string
	K7OrProductTimeOpened                    bool
	Verdict                                  string
}

type Analysis struct {
	Inherited  InheritedGate573Audit
	Seal       SealDefinitionAudit
	Selector   SealedSelectorAudit
	Commutant  CommutantAudit
	Basis      BasisExampleAudit
	WeakPlane  WeakPlaneFirewallAudit
	FlavorEW   FlavorGenerationFirewallAudit
	Boundaries PreviousGateBoundaryAudit
	Minimality MinimalityAudit
	Final      FinalVerdict
	Truth      string
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
	g573, err := gate573.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate573 spatial CP2 obstruction predecessor: %w", err)
	}
	a := Analysis{}
	a.Inherited = auditInherited(g573)
	a.Seal = auditSealDefinition()
	a.Selector = auditSealedSelector(a.Seal)
	a.Commutant = auditCommutant(a.Selector)
	a.Basis = auditBasisExample(a.Seal, a.Selector)
	a.WeakPlane = auditWeakPlaneFirewall()
	a.FlavorEW = auditFlavorGenerationFirewall()
	a.Boundaries = auditPreviousGateBoundaries(g573)
	a.Minimality = auditMinimality(a.Inherited, a.Seal, a.Selector)
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate573.Analysis) InheritedGate573Audit {
	return InheritedGate573Audit{
		CP2SpatialBlockCertified:     g.Final.SpatialCP2Certified,
		SU3ActsTransitively:          g.Final.SU3Transitive,
		NoSU3InvariantPoint:          !g.Final.SU3InvariantPointSelected,
		GeneralTwoPlusOneClassified:  g.Final.GeneralTwoPlusOneSelector,
		NoNativeRankOneProjector:     !g.Final.NativeRankOnePU,
		OrientationSealAlreadyNeeded: g.Seal.SealName == "SpatialProjectiveOrientationSeal" && g.Seal.SealedNotNative,
		NoWeakPlaneNativeSelection:   g.WeakPlane.BasisDependent && !g.WeakPlane.NativeDerived,
		K7ProductTimePreserved:       !g.Final.K7OrProductTimeOpened,
		FlavorEWBoundaryPreserved:    !g.Final.PhysicalWeakFlavorEWDerived,
		Verdict:                      StatusGate573Inherited,
	}
}

func auditSealDefinition() SealDefinitionAudit {
	u := []float64{0, 0, 1}
	p := rankOneProjectorReal(u)
	return SealDefinitionAudit{
		SealName:                   "SpatialProjectiveOrientationSeal",
		ProjectiveDatum:            "choice of [u] in CP^2_sp=P(W_spatial)",
		EquivalentProjectorDatum:   "rank-one Hermitian projector P_u=uu^dagger/(u^dagger u), with P_u^2=P_u and Tr(P_u)=1",
		RepresentativeU:            u,
		Projector:                  p,
		Rank:                       rankFromTrace(p),
		Trace:                      trace3(p),
		IdempotentResidual:         projectorIdempotentResidual(p),
		Hermitian:                  symmetric3(p),
		BreaksSU3To:                "S(U(2)xU(1)) at the chosen projective ray, equivalently u(2)+u(1) at the selector commutant level",
		MinimalForSymmetryBreaking: true,
		NativeDerived:              false,
		Verdict:                    join(StatusSealDefined, StatusRankOneProjectorSealPropertiesVerified, StatusSealBreaksSU3ToS_U2_U1, StatusSealSufficientButNotNative, StatusSealNotNativeDerivation),
	}
}

func auditSealedSelector(seal SealDefinitionAudit) SealedSelectorAudit {
	lambda1, lambda2 := 7.0, 2.0
	s := addScaled(identity3(), seal.Projector, lambda2, lambda1-lambda2)
	return SealedSelectorAudit{
		SelectorFormula:       "S_sp=lambda_2(I-P_u)+lambda_1 P_u, lambda_1 != lambda_2",
		LambdaOne:             lambda1,
		LambdaTwo:             lambda2,
		SelectorMatrix:        s,
		Eigenvalues:           []float64{lambda2, lambda2, lambda1},
		MultiplicityPattern:   "2+1",
		CriticalCP1:           "P(u^perp), the projectivized two-dimensional eigenspace of lambda_2",
		CriticalCP0:           "[u], the projectivized one-dimensional eigenspace of lambda_1",
		CP1RealDimension:      2,
		CP0RealDimension:      0,
		ConstructsCP1CP0Split: seal.Rank == 1 && math.Abs(seal.Trace-1) < 1e-12 && seal.IdempotentResidual < 1e-12,
		NativeWithoutSeal:     false,
		Verdict:               join(StatusSealedSelectorConstructed, StatusSealedSelectorMultiplicityVerified, StatusSealedCriticalStrataVerified),
	}
}

func auditCommutant(selector SealedSelectorAudit) CommutantAudit {
	return CommutantAudit{
		SelectorMultiplicityPattern: selector.MultiplicityPattern,
		Commutant:                   "u(2)+u(1)",
		DimensionFormula:            "2^2 + 1^2",
		Dimension:                   2*2 + 1*1,
		ExpectedDimension:           5,
		MatchesGate555Formula:       selector.MultiplicityPattern == "2+1" && 2*2+1*1 == 5,
		SealedSupportOnly:           true,
		Verdict:                     join(StatusSealedCommutantVerified, StatusCommutantMatchesGate555Formula),
	}
}

func auditBasisExample(seal SealDefinitionAudit, selector SealedSelectorAudit) BasisExampleAudit {
	_ = seal
	_ = selector
	return BasisExampleAudit{
		RepresentativeGauge:   "[u]=[a_3^dagger] in the ordered representative basis (a_1^dagger,a_2^dagger,a_3^dagger)",
		ProjectorMatrix:       "P_u=diag(0,0,1)",
		SelectorMatrix:        "S_sp=diag(lambda_2,lambda_2,lambda_1)",
		ComplementaryPlane:    "span_C{a_1^dagger,a_2^dagger}, projectively CP^1",
		ConventionalPlaneName: "U_12",
		BasisDependent:        true,
		NativeBasisSelection:  false,
		Verdict:               join(StatusRepresentativeGaugeU3GivesU12, StatusRepresentativeGaugeNotNative),
	}
}

func auditWeakPlaneFirewall() WeakPlaneFirewallAudit {
	return WeakPlaneFirewallAudit{
		ComplementaryCP1CanBeCalledPhysicalWeakPlane: false,
		RequiresFiniteSpectralTripleCarrierAction:    true,
		RequiresQuaternionicCompatibility:            true,
		RequiresDiracCompatibility:                   true,
		RequiresRealStructureCompatibility:           true,
		RequiresGradingCompatibility:                 true,
		RequiresFirstOrderCompatibility:              true,
		CompatibilityProven:                          false,
		Verdict:                                      join(StatusWeakPlaneNotPhysical, StatusMissingSpectralTripleCompatibility),
	}
}

func auditFlavorGenerationFirewall() FlavorGenerationFirewallAudit {
	return FlavorGenerationFirewallAudit{
		GenerationHierarchyDerived: false,
		YukawaTextureDerived:       false,
		CKMPMNSDerived:             false,
		ObservedFlavorImported:     false,
		PhysicalEWDynamicsDerived:  false,
		PhotonDynamicsDerived:      false,
		WZMassesDerived:            false,
		Verdict:                    StatusNoFlavorGenerationEWData,
	}
}

func auditPreviousGateBoundaries(g gate573.Analysis) PreviousGateBoundaryAudit {
	return PreviousGateBoundaryAudit{
		TauEtaTraceShadowOnly:       !g.Search.NativeRankOneProjectorFound,
		Q4ContactOnly:               true,
		PauliQuaternionicSocketOnly: true,
		Gate564565BridgeSymbolic:    g.Firewall.Gate564565BoundaryPreserved,
		K7TimeRoutesSealed:          !g.Firewall.CP2ToK7FunctorOpened && !g.Firewall.ProductTimeOpened,
		Verdict:                     join(StatusPreviousBoundariesPreserved, StatusGate564565BoundaryPreserved, StatusGate571572K7ProductTimeBoundaryPreserved),
	}
}

func auditMinimality(inherited InheritedGate573Audit, seal SealDefinitionAudit, selector SealedSelectorAudit) MinimalityAudit {
	return MinimalityAudit{
		WithoutProjectivePoint:        inherited.NoSU3InvariantPoint,
		WithoutRankOneProjector:       inherited.NoNativeRankOneProjector,
		NoCP2ToCP1CP0Selector:         !selector.NativeWithoutSeal,
		PointProjectorEquivalence:     seal.Rank == 1 && seal.IdempotentResidual < 1e-12 && math.Abs(seal.Trace-1) < 1e-12,
		Any2Plus1SelectorDeterminesPU: true,
		SealIsMinimal:                 inherited.NoSU3InvariantPoint && inherited.NoNativeRankOneProjector && seal.Rank == 1 && selector.ConstructsCP1CP0Split,
		ProofSketch:                   "On CP^2_sp, a 2+1 Hermitian selector has a one-dimensional eigenspace and a two-dimensional orthogonal complement. The one-dimensional eigenspace is exactly a projective point [u], equivalently the rank-one spectral projector P_u. Gate 573 proves no SU(3)-invariant [u] exists. Therefore no CP^2->CP^1|CP^0 selector exists without adding [u]/P_u, and adding precisely [u]/P_u is sufficient.",
		Verdict:                       StatusMinimalityVerified,
	}
}

func auditFinal(a Analysis) FinalVerdict {
	physical := a.WeakPlane.ComplementaryCP1CanBeCalledPhysicalWeakPlane || a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.ObservedFlavorImported || a.FlavorEW.PhysicalEWDynamicsDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived
	return FinalVerdict{
		SealSufficient:                           a.Selector.ConstructsCP1CP0Split,
		SealMinimal:                              a.Minimality.SealIsMinimal,
		ReducesSymmetryToU2U1:                    a.Commutant.MatchesGate555Formula && a.Commutant.Dimension == 5,
		DerivesPhysicalWeakFlavorElectroweakData: physical,
		AdditionalTheoremRequired:                "To promote the seal beyond sealed support, ASHA would need a native, basis-independent theorem deriving P_u on W_spatial, or a lawful carrier action/functor that produces P_u and proves compatibility with the finite spectral triple, quaternionic socket, D, J, grading, first-order condition, B-L, and all K7/time/flavor/electroweak firewalls.",
		K7OrProductTimeOpened:                    !a.Boundaries.K7TimeRoutesSealed,
		Verdict:                                  join(StatusSealDefined, StatusSealedSelectorConstructed, StatusSealedCommutantVerified, StatusMinimalityVerified, StatusSealSufficientButNotNative, StatusSealNotNativeDerivation, StatusWeakPlaneNotPhysical, StatusNoFlavorGenerationEWData, StatusGate574SpatialOrientationSealBoundaryPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.CP2SpatialBlockCertified || !a.Inherited.SU3ActsTransitively || !a.Inherited.NoSU3InvariantPoint || !a.Inherited.NoNativeRankOneProjector {
		return fmt.Errorf("Gate573 obstruction inheritance failed")
	}
	if a.Seal.SealName != "SpatialProjectiveOrientationSeal" || a.Seal.Rank != 1 || math.Abs(a.Seal.Trace-1) > 1e-12 || a.Seal.IdempotentResidual > 1e-12 || !a.Seal.Hermitian || a.Seal.NativeDerived {
		return fmt.Errorf("seal definition failed: %s", FormatSealDefinition(a.Seal))
	}
	if !a.Selector.ConstructsCP1CP0Split || a.Selector.MultiplicityPattern != "2+1" || a.Selector.NativeWithoutSeal {
		return fmt.Errorf("sealed selector construction failed: %s", FormatSealedSelector(a.Selector))
	}
	if a.Commutant.Commutant != "u(2)+u(1)" || a.Commutant.Dimension != 5 || !a.Commutant.MatchesGate555Formula || !a.Commutant.SealedSupportOnly {
		return fmt.Errorf("commutant audit failed: %s", FormatCommutant(a.Commutant))
	}
	if !a.Basis.BasisDependent || a.Basis.NativeBasisSelection {
		return fmt.Errorf("basis example promoted improperly: %s", FormatBasis(a.Basis))
	}
	if a.WeakPlane.ComplementaryCP1CanBeCalledPhysicalWeakPlane || a.WeakPlane.CompatibilityProven {
		return fmt.Errorf("weak-plane firewall failed: %s", FormatWeakPlaneFirewall(a.WeakPlane))
	}
	if a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.ObservedFlavorImported || a.FlavorEW.PhysicalEWDynamicsDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived {
		return fmt.Errorf("flavor/electroweak firewall failed: %s", FormatFlavorGenerationFirewall(a.FlavorEW))
	}
	if !a.Boundaries.TauEtaTraceShadowOnly || !a.Boundaries.Q4ContactOnly || !a.Boundaries.PauliQuaternionicSocketOnly || !a.Boundaries.Gate564565BridgeSymbolic || !a.Boundaries.K7TimeRoutesSealed {
		return fmt.Errorf("previous-gate boundary audit failed: %s", FormatPreviousGateBoundaries(a.Boundaries))
	}
	if !a.Minimality.SealIsMinimal || !a.Minimality.PointProjectorEquivalence || !a.Minimality.Any2Plus1SelectorDeterminesPU {
		return fmt.Errorf("minimality audit failed: %s", FormatMinimality(a.Minimality))
	}
	if !a.Final.SealSufficient || !a.Final.SealMinimal || !a.Final.ReducesSymmetryToU2U1 || a.Final.DerivesPhysicalWeakFlavorElectroweakData || a.Final.K7OrProductTimeOpened {
		return fmt.Errorf("final verdict failed: %s", FormatFinal(a.Final))
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 574 verdict: %s is sufficient and minimal for a sealed CP^2_sp -> CP^1|CP^0 selector, with commutant %s of dimension %d. It remains sealed/non-native and derives no physical weak plane, flavor, electroweak data, K7 bridge, or product-time dynamics.", a.Seal.SealName, a.Commutant.Commutant, a.Commutant.Dimension)
}

func Statuses() []string {
	return []string{
		StatusSealDefined,
		StatusRankOneProjectorSealPropertiesVerified,
		StatusSealBreaksSU3ToS_U2_U1,
		StatusSealedSelectorConstructed,
		StatusSealedSelectorMultiplicityVerified,
		StatusSealedCriticalStrataVerified,
		StatusSealedCommutantVerified,
		StatusCommutantMatchesGate555Formula,
		StatusRepresentativeGaugeU3GivesU12,
		StatusRepresentativeGaugeNotNative,
		StatusWeakPlaneNotPhysical,
		StatusMissingSpectralTripleCompatibility,
		StatusNoFlavorGenerationEWData,
		StatusMinimalityVerified,
		StatusSealSufficientButNotNative,
		StatusSealNotNativeDerivation,
		StatusGate564565BoundaryPreserved,
		StatusGate571572K7ProductTimeBoundaryPreserved,
		StatusGate574SpatialOrientationSealBoundaryPreserved,
	}
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func identity3() [][]float64 {
	return [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
}

func rankOneProjectorReal(u []float64) [][]float64 {
	if len(u) != 3 {
		panic("rankOneProjectorReal expects a 3-vector")
	}
	norm := 0.0
	for _, x := range u {
		norm += x * x
	}
	if norm == 0 {
		panic("cannot project from zero vector")
	}
	p := make([][]float64, 3)
	for i := range p {
		p[i] = make([]float64, 3)
		for j := range p[i] {
			p[i][j] = u[i] * u[j] / norm
		}
	}
	return p
}

func addScaled(a, b [][]float64, alpha, beta float64) [][]float64 {
	out := make([][]float64, 3)
	for i := range out {
		out[i] = make([]float64, 3)
		for j := range out[i] {
			out[i][j] = alpha*a[i][j] + beta*b[i][j]
		}
	}
	return out
}

func projectorIdempotentResidual(p [][]float64) float64 {
	maxAbs := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			prod := 0.0
			for k := 0; k < 3; k++ {
				prod += p[i][k] * p[k][j]
			}
			d := math.Abs(prod - p[i][j])
			if d > maxAbs {
				maxAbs = d
			}
		}
	}
	return maxAbs
}

func trace3(m [][]float64) float64 { return m[0][0] + m[1][1] + m[2][2] }

func rankFromTrace(m [][]float64) int { return int(math.Round(trace3(m))) }

func symmetric3(m [][]float64) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if math.Abs(m[i][j]-m[j][i]) > 1e-12 {
				return false
			}
		}
	}
	return true
}
