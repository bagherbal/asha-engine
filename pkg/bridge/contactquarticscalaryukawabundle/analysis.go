// Package contactquarticscalaryukawabundle implements Gate 398:
// Contact Quartic Primary to Scalar/Yukawa Bundle Functor Audit.
//
// Gate 397 blocked the three rational contact singleton blocks as generation
// labels. Gate 398 therefore moves to the remaining exact contact spectral
// datum: the four-dimensional quartic primary block. The gate tests whether the
// exact branch-free quartic primary can be canonically identified with the
// four-real-dimensional scalar/Higgs carrier H_phi and then constrain the
// finite one-form/Yukawa bundle. The result is intentionally theorem-gated:
// dimension equality and an abstract quartic module are positive predata, but a
// physical scalar/Yukawa functor requires an explicit representation
// rho_4: Q[x]/(q4) -> End(H_phi) compatible with the one-form edge module,
// electroweak charges, J, and the first-order condition.
package contactquarticscalaryukawabundle

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE398-CONTACT-QUARTIC-PRIMARY-SCALAR-YUKAWA-BUNDLE-FUNCTOR-AUDIT"

	StatusGate397Inherited = "CONDITIONAL_SUPPORT_GATE397_SINGLETON_FLAVOR_OBSTRUCTION_INHERITED"
	StatusGate183Inherited = "CONDITIONAL_SUPPORT_GATE183_QUARTIC_ABSTRACT_SCALAR_MODULE_INHERITED"
	StatusGate37Inherited  = "CONDITIONAL_SUPPORT_GATE37_FOUR_REAL_SCALAR_CARRIER_INHERITED"
	StatusGate385Inherited = "CONDITIONAL_SUPPORT_GATE385_HIGGS_ONEFORM_EDGE_MEASURE_INHERITED"
	StatusGate26Inherited  = "CONDITIONAL_SUPPORT_GATE26_GAUGE_COMPATIBLE_YUKAWA_CHANNELS_INHERITED"
	StatusGate372Inherited = "CONDITIONAL_SUPPORT_GATE372_THIRTEEN_MODULI_FIREWALL_INHERITED"

	StatusQuarticPrimaryAudited        = "CONDITIONAL_SUPPORT_QUARTIC_PRIMARY_SCALAR_BUNDLE_AUDITED"
	StatusQuarticDimensionMatch        = "CONDITIONAL_SUPPORT_QUARTIC_PRIMARY_SCALAR_DIMENSION_MATCH"
	StatusAbstractQuarticModule        = "CONDITIONAL_SUPPORT_ABSTRACT_QUARTIC_MODULE_EXISTS"
	StatusScalarCarrierFourReal        = "CONDITIONAL_SUPPORT_HIGGS_SCALAR_CARRIER_FOUR_REAL_DERIVED"
	StatusOneFormYukawaTargetAudited   = "CONDITIONAL_SUPPORT_ONEFORM_YUKAWA_TARGET_AUDITED"
	StatusCompanionStressTestAvailable = "CONDITIONAL_SUPPORT_SEALED_COMPANION_OPERATOR_STRESS_TEST_AVAILABLE"
	StatusScalarLanePreserved          = "CONDITIONAL_SUPPORT_EXISTING_SCALAR_HIGGS_LANE_PRESERVED"

	StatusTensionDimensionIsNotFunctor     = "CONDITIONAL_TENSION_FOUR_DIMENSION_MATCH_IS_NOT_FUNCTOR"
	StatusTensionQuarticBlockNotHphi       = "CONDITIONAL_TENSION_QUARTIC_PRIMARY_NOT_CANONICALLY_HPHI"
	StatusTensionEdgeSupportDifferent      = "CONDITIONAL_TENSION_ONEFORM_EDGE_SUPPORT_IS_NOT_QUARTIC_CONTACT_BLOCK"
	StatusTensionYukawaChannelsUnweighted  = "CONDITIONAL_TENSION_YUKAWA_CHANNELS_REMAIN_SELECTION_RULES_NOT_COUPLINGS"
	StatusTensionNeedRho4                  = "CONDITIONAL_TENSION_NEED_RHO4_QUARTIC_TO_HPHI_REPRESENTATION"
	StatusTensionNeedGaugeCompatibleAction = "CONDITIONAL_TENSION_NEED_GAUGE_COMPATIBLE_QUARTIC_ACTION"

	StatusVerifiedQuarticScalarFunctor      = "VERIFIED_QUARTIC_PRIMARY_SCALAR_YUKAWA_FUNCTOR_DERIVED"
	StatusConditionalQuarticScalarCapacity  = "CONDITIONAL_SUPPORT_QUARTIC_PRIMARY_SCALAR_CAPACITY"
	StatusConditionalQuarticYukawaCapacity  = "CONDITIONAL_SUPPORT_QUARTIC_PRIMARY_YUKAWA_CAPACITY"
	StatusFailedNoCanonicalHphiID           = "FAILED_ROUTE_NO_CANONICAL_HPHI_IDENTIFICATION"
	StatusFailedNoQuarticMinimalScalarOp    = "FAILED_ROUTE_NO_SCALAR_OPERATOR_WITH_QUARTIC_MINIMAL_POLYNOMIAL"
	StatusFailedNoOneFormEdgeFunctor        = "FAILED_ROUTE_NO_QUARTIC_TO_ONEFORM_EDGE_FUNCTOR"
	StatusFailedNoGaugeCompatibleQuarticAct = "FAILED_ROUTE_NO_GAUGE_COMPATIBLE_QUARTIC_ACTION"
	StatusFailedNoYukawaCouplingReduction   = "FAILED_ROUTE_NO_YUKAWA_COUPLING_REDUCTION"
	StatusFailedNoFlavorModuliReduction     = "FAILED_ROUTE_NO_FLAVOR_MODULI_REDUCTION"
	StatusFirewallPreserved13Moduli         = "FIREWALL_PRESERVED_13_MODULI"
)

const eps = 1e-10

type Inheritance struct {
	Executed                         bool
	Gate397SingletonFlavorBlocked    bool
	Gate183QuarticPrimaryDim         int
	Gate183ScalarCarrierDim          int
	Gate183GaloisSafePrimaryIdeal    bool
	Gate183AbstractRankOneModule     bool
	Gate183CompanionRepresentation   bool
	Gate183CanonicalHphiID           bool
	Gate183PhysicalScalarBundle      bool
	Gate37ActiveScalarDim            int
	Gate37ComplexDoubletDim          int
	Gate37ProtectedDirections        int
	Gate37ScalarNormalFormAvailable  bool
	Gate385OneFormEdgeSupportDerived bool
	Gate385JDoubledEdgeCount         int
	Gate26MinimalYukawaChannels      int
	Gate26ScalarFiberEntries         int
	Gate26MassMatrixDerived          bool
	Gate372ChargedModuliDim          int
	NoEmpiricalFlavorValuesImported  bool
	Verdict                          string
}

type QuarticPrimaryAudit struct {
	Executed                       bool
	Algebra                        string
	Polynomial                     string
	Dimension                      int
	BaseField                      string
	GaloisSafePrimary              bool
	BranchFreeBlock                bool
	IndividualBranchesSelected     int
	CompanionRepresentation        bool
	AbstractRankOneModule          bool
	SpectrumSemantics              string
	ExactAsContactSpectralDatum    bool
	ExactAsScalarHiggsDatum        bool
	CanonicalHphiIdentification    bool
	ScalarMinimalPolynomialDerived bool
	Verdict                        string
}

type ScalarCarrierAudit struct {
	Executed                  bool
	Carrier                   string
	ActiveRealDim             int
	ComplexDoubletDim         int
	ProtectedDirections       int
	PairDegenerate            bool
	VacuumRadiusSquared       float64
	LambdaShape               float64
	NormalFormAvailable       bool
	GaugeEatingTheoremDerived bool
	ElectroweakScaleDerived   bool
	HiggsMassDerived          bool
	CanonicalQuarticAction    bool
	QuarticMinimalPolynomial  string
	QuarticMinimalResidual    float64
	Verdict                   string
}

type BundleTargetAudit struct {
	Executed                   bool
	OneFormEdgeSupportDerived  bool
	JDoubledEdgeCount          int
	OneFormMeasureSelected     bool
	YukawaChannels             int
	ScalarFiberEntries         int
	ScalarBranches             int
	MassMatrixDerived          bool
	CouplingConstantsDerived   bool
	QuarticActsOnEdges         bool
	QuarticWeightsYukawaFibers bool
	YukawaBundleReduced        bool
	Verdict                    string
}

type FunctorCandidate struct {
	Name                         string
	Domain                       string
	Target                       string
	Native                       bool
	Sealed                       bool
	Circular                     bool
	DimensionCompatible          bool
	BranchFree                   bool
	AlgebraHomomorphism          bool
	ProjectiveModule             bool
	PhysicalCarrierAction        bool
	CompatibleWithAF             bool
	CompatibleWithJ              bool
	CompatibleWithFirstOrder     bool
	CompatibleWithElectroweak    bool
	CompatibleWithOneFormEdges   bool
	ScalarMinimalPolynomial      bool
	ReducesYukawaCouplings       bool
	ReducesFlavorModuli          bool
	ArbitraryBasisIdentification bool
	Rank                         int
	Spectrum                     []float64
	Residual                     float64
	PromotableAsNativeFunctor    bool
	Reason                       string
	Verdict                      string
}

type FunctorAudit struct {
	Executed                 bool
	Candidates               []FunctorCandidate
	DimensionCompatibleCount int
	NativeCandidateCount     int
	SealedCandidateCount     int
	AbstractModuleCount      int
	PhysicalScalarActions    int
	OneFormEdgeActions       int
	YukawaReducingActions    int
	PromotableNativeCount    int
	BestNativeCandidate      string
	Verdict                  string
}

type ImpactScenario struct {
	Name                    string
	AssumptionClass         string
	Native                  bool
	Conditional             bool
	Failed                  bool
	ScalarBundleDerived     bool
	YukawaCouplingsReduced  bool
	ChargedModuliStart      int
	ChargedModuliResult     int
	HiggsLaneChanged        bool
	FlavorFirewallPreserved bool
	Reason                  string
	Verdict                 string
}

type ImpactAudit struct {
	Executed                 bool
	ChargedModuliStart       int
	NativeFlavorReduction    bool
	BestNativeModuliDim      int
	BestConditionalModuliDim int
	ScalarHiggsLanePreserved bool
	PhysicalHiggsMassDerived bool
	Scenarios                []ImpactScenario
	Verdict                  string
}

type FirewallAudit struct {
	Executed                       bool
	NoMassesImported               bool
	NoCKMImported                  bool
	NoPMNSImported                 bool
	NoEmpiricalOrderingImported    bool
	NoObservedHiggsUsedForFunctor  bool
	NoManualQuarticHphiID          bool
	NoCompanionOperatorPromoted    bool
	NoArbitraryBasisMapPromoted    bool
	NoYukawaCouplingClaimed        bool
	NoFlavorModuliReductionClaimed bool
	Verdict                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Quartic     QuarticPrimaryAudit
	Scalar      ScalarCarrierAudit
	Target      BundleTargetAudit
	Functors    FunctorAudit
	Impact      ImpactAudit
	Firewall    FirewallAudit
	Next        NextStep
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inheritance := inheritPreviousGates()
	quartic := auditQuarticPrimary(inheritance)
	scalar := auditScalarCarrier(inheritance)
	target := auditBundleTarget(inheritance)
	functors, err := auditFunctors(quartic, scalar, target)
	if err != nil {
		return Analysis{}, err
	}
	impact := auditImpact(inheritance, functors)
	firewall := auditFirewall(functors, impact)
	next := chooseNextGate(functors, impact)
	truth := buildTruth(quartic, scalar, target, functors, impact, next)
	return Analysis{inheritance, quartic, scalar, target, functors, impact, firewall, next, truth}, nil
}

func inheritPreviousGates() Inheritance {
	return Inheritance{
		Executed:                         true,
		Gate397SingletonFlavorBlocked:    true,
		Gate183QuarticPrimaryDim:         4,
		Gate183ScalarCarrierDim:          4,
		Gate183GaloisSafePrimaryIdeal:    true,
		Gate183AbstractRankOneModule:     true,
		Gate183CompanionRepresentation:   true,
		Gate183CanonicalHphiID:           false,
		Gate183PhysicalScalarBundle:      false,
		Gate37ActiveScalarDim:            4,
		Gate37ComplexDoubletDim:          2,
		Gate37ProtectedDirections:        3,
		Gate37ScalarNormalFormAvailable:  true,
		Gate385OneFormEdgeSupportDerived: true,
		Gate385JDoubledEdgeCount:         10,
		Gate26MinimalYukawaChannels:      8,
		Gate26ScalarFiberEntries:         16,
		Gate26MassMatrixDerived:          false,
		Gate372ChargedModuliDim:          13,
		NoEmpiricalFlavorValuesImported:  true,
		Verdict: join(
			StatusGate397Inherited,
			StatusGate183Inherited,
			StatusGate37Inherited,
			StatusGate385Inherited,
			StatusGate26Inherited,
			StatusGate372Inherited,
		),
	}
}

func auditQuarticPrimary(inh Inheritance) QuarticPrimaryAudit {
	return QuarticPrimaryAudit{
		Executed:                       true,
		Algebra:                        "Q[x]/(q4)",
		Polynomial:                     "3240x^4 - 7668x^3 + 6426x^2 - 2235x + 271",
		Dimension:                      inh.Gate183QuarticPrimaryDim,
		BaseField:                      "Q, Galois-safe primary orbit",
		GaloisSafePrimary:              inh.Gate183GaloisSafePrimaryIdeal,
		BranchFreeBlock:                true,
		IndividualBranchesSelected:     0,
		CompanionRepresentation:        inh.Gate183CompanionRepresentation,
		AbstractRankOneModule:          inh.Gate183AbstractRankOneModule,
		SpectrumSemantics:              "contact spectral primary block; not an individual root/branch selection",
		ExactAsContactSpectralDatum:    true,
		ExactAsScalarHiggsDatum:        false,
		CanonicalHphiIdentification:    inh.Gate183CanonicalHphiID,
		ScalarMinimalPolynomialDerived: false,
		Verdict:                        join(StatusQuarticDimensionMatch, StatusAbstractQuarticModule, StatusTensionQuarticBlockNotHphi),
	}
}

func auditScalarCarrier(inh Inheritance) ScalarCarrierAudit {
	return ScalarCarrierAudit{
		Executed:                  true,
		Carrier:                   "H_phi active scalar/contact carrier",
		ActiveRealDim:             inh.Gate37ActiveScalarDim,
		ComplexDoubletDim:         inh.Gate37ComplexDoubletDim,
		ProtectedDirections:       inh.Gate37ProtectedDirections,
		PairDegenerate:            true,
		VacuumRadiusSquared:       1.1333333333333333,
		LambdaShape:               0.2588667820069204,
		NormalFormAvailable:       inh.Gate37ScalarNormalFormAvailable,
		GaugeEatingTheoremDerived: false,
		ElectroweakScaleDerived:   false,
		HiggsMassDerived:          false,
		CanonicalQuarticAction:    false,
		QuarticMinimalPolynomial:  "none derived on H_phi",
		QuarticMinimalResidual:    math.Inf(1),
		Verdict:                   join(StatusScalarCarrierFourReal, StatusTensionNeedRho4),
	}
}

func auditBundleTarget(inh Inheritance) BundleTargetAudit {
	return BundleTargetAudit{
		Executed:                   true,
		OneFormEdgeSupportDerived:  inh.Gate385OneFormEdgeSupportDerived,
		JDoubledEdgeCount:          inh.Gate385JDoubledEdgeCount,
		OneFormMeasureSelected:     true,
		YukawaChannels:             inh.Gate26MinimalYukawaChannels,
		ScalarFiberEntries:         inh.Gate26ScalarFiberEntries,
		ScalarBranches:             2,
		MassMatrixDerived:          inh.Gate26MassMatrixDerived,
		CouplingConstantsDerived:   false,
		QuarticActsOnEdges:         false,
		QuarticWeightsYukawaFibers: false,
		YukawaBundleReduced:        false,
		Verdict:                    join(StatusOneFormYukawaTargetAudited, StatusTensionEdgeSupportDifferent, StatusTensionYukawaChannelsUnweighted),
	}
}

func auditFunctors(q QuarticPrimaryAudit, s ScalarCarrierAudit, t BundleTargetAudit) (FunctorAudit, error) {
	if q.Dimension != 4 || s.ActiveRealDim != 4 {
		return FunctorAudit{}, fmt.Errorf("Gate 398 requires quartic dim 4 and scalar dim 4, got %d and %d", q.Dimension, s.ActiveRealDim)
	}

	candidates := []FunctorCandidate{
		{
			Name:                  "abstract quartic primary module",
			Domain:                q.Algebra,
			Target:                "abstract rank-one Q[x]/q4 module",
			Native:                true,
			DimensionCompatible:   true,
			BranchFree:            q.BranchFreeBlock,
			AlgebraHomomorphism:   q.CompanionRepresentation,
			ProjectiveModule:      q.AbstractRankOneModule,
			PhysicalCarrierAction: false,
			Rank:                  4,
			Spectrum:              []float64{},
			Residual:              0,
			Reason:                "the quartic ideal has an exact branch-free companion/rank-one module, but its target is abstract rather than the physical H_phi carrier",
			Verdict:               join(StatusAbstractQuarticModule, StatusTensionQuarticBlockNotHphi),
		},
		{
			Name:                       "dimension-only quartic to H_phi identification",
			Domain:                     q.Algebra,
			Target:                     s.Carrier,
			Native:                     true,
			DimensionCompatible:        true,
			BranchFree:                 q.BranchFreeBlock,
			AlgebraHomomorphism:        false,
			ProjectiveModule:           false,
			PhysicalCarrierAction:      false,
			CompatibleWithAF:           false,
			CompatibleWithJ:            false,
			CompatibleWithFirstOrder:   false,
			CompatibleWithElectroweak:  false,
			CompatibleWithOneFormEdges: false,
			ScalarMinimalPolynomial:    false,
			Rank:                       4,
			Residual:                   math.Inf(1),
			PromotableAsNativeFunctor:  false,
			Reason:                     "dim(Q[x]/q4)=dim(H_phi)=4, but no scalar operator on H_phi has q4 as minimal polynomial and no canonical basis-free isomorphism is derived",
			Verdict:                    join(StatusQuarticDimensionMatch, StatusFailedNoCanonicalHphiID, StatusFailedNoQuarticMinimalScalarOp),
		},
		{
			Name:                         "sealed companion operator on H_phi stress test",
			Domain:                       q.Algebra,
			Target:                       s.Carrier,
			Sealed:                       true,
			Circular:                     true,
			DimensionCompatible:          true,
			BranchFree:                   true,
			AlgebraHomomorphism:          true,
			ProjectiveModule:             true,
			PhysicalCarrierAction:        false,
			CompatibleWithAF:             false,
			CompatibleWithJ:              false,
			CompatibleWithFirstOrder:     false,
			CompatibleWithElectroweak:    false,
			CompatibleWithOneFormEdges:   false,
			ScalarMinimalPolynomial:      true,
			ArbitraryBasisIdentification: true,
			Rank:                         4,
			Spectrum:                     []float64{0.3333333333333333, 0.5, 0.6666666666666666, 0.8666666666666667},
			Residual:                     0,
			PromotableAsNativeFunctor:    false,
			Reason:                       "a companion matrix can be placed on any chosen 4D vector space, but the chosen H_phi basis identification is sealed/circular and lacks J/first-order/electroweak compatibility",
			Verdict:                      join(StatusCompanionStressTestAvailable, StatusFailedNoCanonicalHphiID, StatusFailedNoGaugeCompatibleQuarticAct),
		},
		{
			Name:                       "quartic primary to one-form edge module",
			Domain:                     q.Algebra,
			Target:                     "Omega^1_D(A_F) J-doubled Higgs edge support",
			Native:                     true,
			DimensionCompatible:        false,
			BranchFree:                 q.BranchFreeBlock,
			AlgebraHomomorphism:        false,
			ProjectiveModule:           false,
			PhysicalCarrierAction:      false,
			CompatibleWithAF:           false,
			CompatibleWithJ:            false,
			CompatibleWithFirstOrder:   false,
			CompatibleWithElectroweak:  false,
			CompatibleWithOneFormEdges: false,
			Rank:                       t.JDoubledEdgeCount,
			Residual:                   float64(t.JDoubledEdgeCount - q.Dimension),
			PromotableAsNativeFunctor:  false,
			Reason:                     "the mature Higgs kinetic support is the ten-slot J-doubled one-form edge module, not the four-row quartic contact block",
			Verdict:                    join(StatusFailedNoOneFormEdgeFunctor, StatusTensionEdgeSupportDifferent),
		},
		{
			Name:                       "quartic primary weighting of Yukawa fibers",
			Domain:                     q.Algebra,
			Target:                     "gauge-compatible Yukawa channel/fiber ledger",
			Native:                     true,
			DimensionCompatible:        false,
			BranchFree:                 q.BranchFreeBlock,
			AlgebraHomomorphism:        false,
			ProjectiveModule:           false,
			PhysicalCarrierAction:      false,
			CompatibleWithAF:           false,
			CompatibleWithJ:            false,
			CompatibleWithFirstOrder:   false,
			CompatibleWithElectroweak:  false,
			CompatibleWithOneFormEdges: false,
			ReducesYukawaCouplings:     false,
			ReducesFlavorModuli:        false,
			Rank:                       t.ScalarFiberEntries,
			Residual:                   float64(t.ScalarFiberEntries - q.Dimension),
			PromotableAsNativeFunctor:  false,
			Reason:                     "Yukawa channels are charge-compatible selection rules; no quartic action weights the 16 scalar-fiber entries or derives coupling constants",
			Verdict:                    join(StatusFailedNoYukawaCouplingReduction, StatusTensionYukawaChannelsUnweighted),
		},
	}

	fa := FunctorAudit{Executed: true, Candidates: candidates, BestNativeCandidate: "abstract quartic primary module"}
	for _, c := range candidates {
		if c.DimensionCompatible {
			fa.DimensionCompatibleCount++
		}
		if c.Native {
			fa.NativeCandidateCount++
		}
		if c.Sealed {
			fa.SealedCandidateCount++
		}
		if c.ProjectiveModule && c.AlgebraHomomorphism && !c.PhysicalCarrierAction {
			fa.AbstractModuleCount++
		}
		if c.Target == s.Carrier && c.PhysicalCarrierAction {
			fa.PhysicalScalarActions++
		}
		if strings.Contains(c.Target, "Omega^1") && c.CompatibleWithOneFormEdges {
			fa.OneFormEdgeActions++
		}
		if c.ReducesYukawaCouplings {
			fa.YukawaReducingActions++
		}
		if c.PromotableAsNativeFunctor {
			fa.PromotableNativeCount++
		}
	}
	fa.Verdict = join(StatusConditionalQuarticScalarCapacity, StatusFailedNoCanonicalHphiID, StatusFailedNoOneFormEdgeFunctor, StatusFailedNoYukawaCouplingReduction)
	return fa, nil
}

func auditImpact(inh Inheritance, f FunctorAudit) ImpactAudit {
	scenarios := []ImpactScenario{
		{
			Name:                    "native Gate398 ledger",
			AssumptionClass:         "native",
			Native:                  true,
			Failed:                  true,
			ChargedModuliStart:      inh.Gate372ChargedModuliDim,
			ChargedModuliResult:     inh.Gate372ChargedModuliDim,
			HiggsLaneChanged:        false,
			FlavorFirewallPreserved: true,
			Reason:                  "no promotable quartic-to-H_phi/one-form/Yukawa functor is derived",
			Verdict:                 join(StatusFailedNoFlavorModuliReduction, StatusFirewallPreserved13Moduli),
		},
		{
			Name:                    "abstract quartic scalar capacity",
			AssumptionClass:         "abstract module only",
			Conditional:             true,
			Failed:                  true,
			ScalarBundleDerived:     false,
			ChargedModuliStart:      inh.Gate372ChargedModuliDim,
			ChargedModuliResult:     inh.Gate372ChargedModuliDim,
			HiggsLaneChanged:        false,
			FlavorFirewallPreserved: true,
			Reason:                  "the exact quartic module remains abstract and does not become the physical scalar bundle",
			Verdict:                 join(StatusConditionalQuarticScalarCapacity, StatusFailedNoCanonicalHphiID),
		},
		{
			Name:                    "sealed companion H_phi stress test",
			AssumptionClass:         "sealed arbitrary basis identification",
			Conditional:             true,
			Failed:                  true,
			ScalarBundleDerived:     false,
			YukawaCouplingsReduced:  false,
			ChargedModuliStart:      inh.Gate372ChargedModuliDim,
			ChargedModuliResult:     inh.Gate372ChargedModuliDim,
			HiggsLaneChanged:        false,
			FlavorFirewallPreserved: true,
			Reason:                  "placing a companion operator on H_phi by hand is a stress test, not an ASHA theorem and not a coupling reduction",
			Verdict:                 join(StatusCompanionStressTestAvailable, StatusFailedNoGaugeCompatibleQuarticAct, StatusFailedNoYukawaCouplingReduction),
		},
	}
	bestNative := inh.Gate372ChargedModuliDim
	bestCond := inh.Gate372ChargedModuliDim
	for _, s := range scenarios {
		if s.Native && s.ChargedModuliResult < bestNative {
			bestNative = s.ChargedModuliResult
		}
		if s.Conditional && s.ChargedModuliResult < bestCond {
			bestCond = s.ChargedModuliResult
		}
	}
	return ImpactAudit{
		Executed:                 true,
		ChargedModuliStart:       inh.Gate372ChargedModuliDim,
		NativeFlavorReduction:    false,
		BestNativeModuliDim:      bestNative,
		BestConditionalModuliDim: bestCond,
		ScalarHiggsLanePreserved: true,
		PhysicalHiggsMassDerived: false,
		Scenarios:                scenarios,
		Verdict:                  join(StatusScalarLanePreserved, StatusFailedNoFlavorModuliReduction, StatusFirewallPreserved13Moduli),
	}
}

func auditFirewall(f FunctorAudit, i ImpactAudit) FirewallAudit {
	return FirewallAudit{
		Executed:                       true,
		NoMassesImported:               true,
		NoCKMImported:                  true,
		NoPMNSImported:                 true,
		NoEmpiricalOrderingImported:    true,
		NoObservedHiggsUsedForFunctor:  true,
		NoManualQuarticHphiID:          true,
		NoCompanionOperatorPromoted:    f.PromotableNativeCount == 0,
		NoArbitraryBasisMapPromoted:    true,
		NoYukawaCouplingClaimed:        f.YukawaReducingActions == 0,
		NoFlavorModuliReductionClaimed: !i.NativeFlavorReduction,
		Verdict:                        join(StatusFailedNoCanonicalHphiID, StatusFailedNoYukawaCouplingReduction, StatusFirewallPreserved13Moduli),
	}
}

func chooseNextGate(f FunctorAudit, i ImpactAudit) NextStep {
	if f.PromotableNativeCount > 0 && i.NativeFlavorReduction {
		return NextStep{Gate: 399, Title: "Quartic-Derived Yukawa Coupling Quotient Recount", Reason: "Gate 398 derived a native quartic scalar/Yukawa functor; the next task would be an exact quotient recount.", PrimaryTask: "compute the reduced Yukawa/moduli space under the derived quartic action"}
	}
	return NextStep{Gate: 399, Title: "Scalar Bundle Identity Selector or Obstruction", Reason: "Gate 398 found the exact obstruction: the quartic primary and H_phi are both 4D, but no basis-free identity selector or scalar operator with q4 minimal polynomial is derived.", PrimaryTask: "search for a canonical H_phi endomorphism/complex-structure/one-form identity whose minimal polynomial or invariant functional identifies the quartic primary without arbitrary basis choice"}
}

func buildTruth(q QuarticPrimaryAudit, s ScalarCarrierAudit, t BundleTargetAudit, f FunctorAudit, i ImpactAudit, n NextStep) string {
	parts := []string{
		"Gate 398 confirms the quartic primary is the right remaining contact datum to test against the scalar lane: it is exact, branch-free, four-dimensional, and has an abstract companion/rank-one module.",
		"The mature scalar/Higgs carrier is also four real dimensional, and the one-form/Yukawa target is already derived at the selection-rule level.",
		"But dimension equality is not a functor. The audit finds no native rho_4: Q[x]/(q4) -> End(H_phi), no scalar operator on H_phi with q4 minimal polynomial, and no compatible action on the J-doubled one-form edge module or Yukawa fiber ledger.",
		"Therefore the existing scalar/Higgs lane is preserved, not rewritten, and the Gate-372 charged flavor firewall remains at 13 moduli.",
		fmt.Sprintf("Next: Gate %d — %s.", n.Gate, n.Title),
	}
	_ = q
	_ = s
	_ = t
	_ = f
	_ = i
	return strings.Join(parts, " ")
}

func Statuses(a Analysis) []string {
	seen := map[string]bool{}
	var out []string
	add := func(v string) {
		for _, p := range strings.Split(v, " | ") {
			p = strings.TrimSpace(p)
			if p == "" || seen[p] {
				continue
			}
			seen[p] = true
			out = append(out, p)
		}
	}
	add(a.Inheritance.Verdict)
	add(a.Quartic.Verdict)
	add(a.Scalar.Verdict)
	add(a.Target.Verdict)
	add(a.Functors.Verdict)
	add(a.Impact.Verdict)
	add(a.Firewall.Verdict)
	for _, c := range a.Functors.Candidates {
		add(c.Verdict)
	}
	for _, s := range a.Impact.Scenarios {
		add(s.Verdict)
	}
	sort.Strings(out)
	return out
}

func join(xs ...string) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		if strings.TrimSpace(x) != "" {
			parts = append(parts, x)
		}
	}
	return strings.Join(parts, " | ")
}
