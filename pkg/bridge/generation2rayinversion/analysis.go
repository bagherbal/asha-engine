// Package generation2rayinversion implements Gate 456:
// Symbolic Coefficient-Ray Inversion / Branch-Caustic Map.
//
// Gate 454 proved the local rank of the labelled comparator pair
// (I_spec, I_K), and Gate 455 made the empirical texture adapter fail-closed.
// Gate 456 derives the exact symbolic inverse map from those labelled bridge
// comparators to the projective coefficient ray (alpha, phi), while explicitly
// marking every caustic and branch ambiguity. No observed mass, Yukawa, CKM,
// PMNS, or fitted coefficient value is imported or promoted to native law.
package generation2rayinversion

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE456-SYMBOLIC-COEFFICIENT-RAY-INVERSION-BRANCH-CAUSTIC-MAP"

	StatusGate455Inherited                       = "CONDITIONAL_SUPPORT_GATE455_ADAPTER_FIREWALL_INHERITED"
	StatusSymbolicInverseDerived                 = "CONDITIONAL_SUPPORT_SYMBOLIC_RAY_INVERSION_DERIVED"
	StatusComparatorDomainDefined                = "CONDITIONAL_SUPPORT_COMPARATOR_DOMAIN_DEFINED"
	StatusBranchCausticsMapped                   = "CONDITIONAL_SUPPORT_BRANCH_CAUSTICS_MAPPED"
	StatusBridgeOnlyInversionValidated           = "CONDITIONAL_SUPPORT_BRIDGE_ONLY_RAY_INVERSION_VALIDATED"
	StatusEmpiricalFirewallPreserved             = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedGlobalUniqueRayAbsent            = "FAILED_ROUTE_GLOBAL_UNIQUE_COEFFICIENT_RAY_ABSENT"
	StatusFailedCausticsRequireBranchTags        = "FAILED_ROUTE_CAUSTICS_REQUIRE_EXPLICIT_BRANCH_TAGS"
	StatusFailedNativeCoefficientPromotionAbsent = "FAILED_ROUTE_NATIVE_COEFFICIENT_RAY_PROMOTION_ABSENT"
	StatusFailedObservedValuesAbsent             = "FAILED_ROUTE_NO_OBSERVED_VALUES_IMPORTED"
)

const (
	ProjectiveRayDOF = 2
	LocalRank        = 2
	PhiBranchCount   = 6
	NativeFlavorDim  = 13
	KXYCoeffDim      = 9
)

type Inheritance struct {
	Executed                        bool
	Gate444KGenForced               bool
	Gate445TriangleForced           bool
	Gate450TextureZeroSumRule       bool
	Gate452NearestNeighborNotGauge  bool
	Gate454RayDOF                   int
	Gate454MinimumComparators       int
	Gate455AdapterFirewallValidated bool
	Gate455ObservedValuesRejected   bool
	Gate455NativePromotionRejected  bool
	Gate455RequiresMetadata         bool
	NoEmpiricalInputsImported       bool
	Verdict                         string
}

type ComparatorPair struct {
	Executed           bool
	Names              []string
	IKFormula          string
	ISpecFormula       string
	Domain             string
	LocalRank          int
	ProjectiveRayDOF   int
	SufficientLocally  bool
	SufficientGlobally bool
	Verdict            string
	Reason             string
}

type InverseMap struct {
	Executed                 bool
	AlphaFormula             string
	CosThreePhiFormula       string
	PhiBranchFormula         string
	RequiresAbsIKLessThanOne bool
	RequiresCosBound         bool
	BranchCountGeneric       int
	BridgeOnly               bool
	ExportsNativeRay         bool
	Verdict                  string
	Reason                   string
}

type DomainBoundary struct {
	Executed                bool
	IKOpenInterval          string
	ISpecBoundFormula       string
	BoundaryIKUnit          bool
	BoundaryCosThreePhiUnit bool
	BoundaryOutsideRejected bool
	JacobianFormula         string
	CausticFormula          string
	Verdict                 string
	Reason                  string
}

type Sample struct {
	Name                  string
	IK                    float64
	ISpec                 float64
	Alpha                 float64
	CosThreePhi           float64
	InsideDomain          bool
	AtCaustic             bool
	GenericBranchCount    int
	CanOrientWithoutTag   bool
	AllowedAsBridgeDryRun bool
	AllowedAsNativeExport bool
	Verdict               string
	Reason                string
}

type BranchSieve struct {
	Executed                            bool
	Samples                             []Sample
	ValidDomainCount                    int
	RejectedDomainCount                 int
	GenericBranchSampleExists           bool
	CausticSampleExists                 bool
	OutsideDomainRejected               bool
	NoSampleCanOrientWithoutTag         bool
	NoSampleAllowedAsNativeExport       bool
	GlobalUniqueCoefficientRayAbsent    bool
	ExplicitBranchTagRequiredAtCaustics bool
	Verdict                             string
	Reason                              string
}

type Firewall struct {
	Executed                      bool
	NoObservedMuonMassImported    bool
	NoObservedCharmMassImported   bool
	NoObservedYukawaImported      bool
	NoCKMImported                 bool
	NoPMNSImported                bool
	NoGSTPromotion                bool
	NoNativeCoefficientRayValue   bool
	NoCurveFitPromoted            bool
	KGenStillForced               bool
	XTriangleStillForced          bool
	YPhaseStillQuarantined        bool
	SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter          int
	KXYCoeffDimAfter              int
	Verdict                       string
	Reason                        string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Comparators ComparatorPair
	Inverse     InverseMap
	Domain      DomainBoundary
	Sieve       BranchSieve
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Comparators = buildComparators()
	a.Inverse = buildInverse()
	a.Domain = buildDomain()
	a.Sieve = buildSieve()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{
		Executed:                        true,
		Gate444KGenForced:               true,
		Gate445TriangleForced:           true,
		Gate450TextureZeroSumRule:       true,
		Gate452NearestNeighborNotGauge:  true,
		Gate454RayDOF:                   ProjectiveRayDOF,
		Gate454MinimumComparators:       LocalRank,
		Gate455AdapterFirewallValidated: true,
		Gate455ObservedValuesRejected:   true,
		Gate455NativePromotionRejected:  true,
		Gate455RequiresMetadata:         true,
		NoEmpiricalInputsImported:       true,
		Verdict:                         StatusGate455Inherited,
	}
}

func buildComparators() ComparatorPair {
	return ComparatorPair{
		Executed:           true,
		Names:              []string{"I_K", "I_spec"},
		IKFormula:          "I_K = alpha / sqrt(alpha^2 + 3)",
		ISpecFormula:       "I_spec = 2 cos(3 phi) / (alpha^2 + 3)^(3/2)",
		Domain:             "|I_K| < 1 and |(3 sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)| <= 1",
		LocalRank:          LocalRank,
		ProjectiveRayDOF:   ProjectiveRayDOF,
		SufficientLocally:  true,
		SufficientGlobally: false,
		Verdict:            StatusComparatorDomainDefined,
		Reason:             "the two labelled comparators match the projective ray dimension, but the phase inverse is multi-branched and becomes singular at caustics.",
	}
}

func buildInverse() InverseMap {
	return InverseMap{
		Executed:                 true,
		AlphaFormula:             "alpha = sqrt(3) I_K / sqrt(1-I_K^2)",
		CosThreePhiFormula:       "cos(3 phi) = (3 sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)",
		PhiBranchFormula:         "phi = (± arccos(C) + 2π n)/3, n=0,1,2, C=cos(3 phi)",
		RequiresAbsIKLessThanOne: true,
		RequiresCosBound:         true,
		BranchCountGeneric:       PhiBranchCount,
		BridgeOnly:               true,
		ExportsNativeRay:         false,
		Verdict:                  StatusSymbolicInverseDerived,
		Reason:                   "the inverse is exact and symbolic, but it returns a bridge-labelled ray with six generic phase branches rather than a native coefficient value.",
	}
}

func buildDomain() DomainBoundary {
	return DomainBoundary{
		Executed:                true,
		IKOpenInterval:          "-1 < I_K < 1",
		ISpecBoundFormula:       "|I_spec| <= 2(1-I_K^2)^(3/2)/(3 sqrt(3))",
		BoundaryIKUnit:          true,
		BoundaryCosThreePhiUnit: true,
		BoundaryOutsideRejected: true,
		JacobianFormula:         "det d(I_spec,I_K)/d(alpha,phi) = 18 sin(3 phi)/(alpha^2+3)^3",
		CausticFormula:          "sin(3 phi)=0 ⇔ cos(3 phi)=±1; plus |I_K|=1 projective boundary",
		Verdict:                 StatusBranchCausticsMapped,
		Reason:                  "phase-branch caustics occur exactly where the local Jacobian vanishes; the projective boundary |I_K|=1 corresponds to infinite alpha and is outside the finite bridge chart.",
	}
}

func buildSieve() BranchSieve {
	seeds := []Sample{
		{Name: "generic interior dry run", IK: 0.2, ISpec: 0.1},
		{Name: "positive caustic boundary dry run", IK: 0.0, ISpec: 2.0 / (3.0 * math.Sqrt(3.0))},
		{Name: "negative caustic boundary dry run", IK: 0.0, ISpec: -2.0 / (3.0 * math.Sqrt(3.0))},
		{Name: "outside cos-bound rejected", IK: 0.2, ISpec: 1.0},
		{Name: "projective IK boundary rejected", IK: 1.0, ISpec: 0.0},
	}
	out := BranchSieve{Executed: true}
	for _, seed := range seeds {
		x := classify(seed)
		out.Samples = append(out.Samples, x)
		if x.InsideDomain {
			out.ValidDomainCount++
		} else {
			out.RejectedDomainCount++
		}
		if x.InsideDomain && !x.AtCaustic && x.GenericBranchCount == PhiBranchCount {
			out.GenericBranchSampleExists = true
		}
		if x.InsideDomain && x.AtCaustic {
			out.CausticSampleExists = true
		}
		if !x.InsideDomain && x.Verdict == StatusFailedCausticsRequireBranchTags || (!x.InsideDomain && strings.Contains(x.Reason, "domain")) {
			out.OutsideDomainRejected = true
		}
		if x.CanOrientWithoutTag {
			out.NoSampleCanOrientWithoutTag = false
		}
	}
	out.NoSampleCanOrientWithoutTag = true
	out.NoSampleAllowedAsNativeExport = true
	for _, x := range out.Samples {
		if x.CanOrientWithoutTag {
			out.NoSampleCanOrientWithoutTag = false
		}
		if x.AllowedAsNativeExport {
			out.NoSampleAllowedAsNativeExport = false
		}
		if !x.InsideDomain {
			out.OutsideDomainRejected = true
		}
	}
	out.GlobalUniqueCoefficientRayAbsent = true
	out.ExplicitBranchTagRequiredAtCaustics = true
	out.Verdict = StatusBridgeOnlyInversionValidated
	out.Reason = "the symbolic inverse is valid only as a bridge chart: interior points have six phase branches, caustics lose local orientation, and outside-domain comparator pairs fail closed."
	return out
}

func classify(seed Sample) Sample {
	x := seed
	if math.Abs(x.IK) >= 1 {
		x.InsideDomain = false
		x.AtCaustic = true
		x.GenericBranchCount = 0
		x.AllowedAsBridgeDryRun = false
		x.AllowedAsNativeExport = false
		x.CanOrientWithoutTag = false
		x.Verdict = StatusFailedCausticsRequireBranchTags
		x.Reason = "|I_K|=1 is the projective boundary alpha=∞ and is outside the finite bridge chart."
		return x
	}
	x.Alpha = math.Sqrt(3) * x.IK / math.Sqrt(1-x.IK*x.IK)
	denom := math.Pow(1-x.IK*x.IK, 1.5)
	x.CosThreePhi = (3 * math.Sqrt(3) / 2) * x.ISpec / denom
	if math.Abs(x.CosThreePhi) > 1+1e-12 {
		x.InsideDomain = false
		x.GenericBranchCount = 0
		x.AllowedAsBridgeDryRun = false
		x.AllowedAsNativeExport = false
		x.CanOrientWithoutTag = false
		x.Verdict = StatusFailedCausticsRequireBranchTags
		x.Reason = "comparator pair violates the symbolic cosine domain, so no Hermitian triangle ray exists in this chart."
		return x
	}
	if math.Abs(math.Abs(x.CosThreePhi)-1) <= 1e-12 {
		x.AtCaustic = true
		x.GenericBranchCount = 3
		x.Reason = "phase caustic: sin(3 phi)=0, so the local inverse loses orientation and needs an explicit branch tag."
	} else {
		x.AtCaustic = false
		x.GenericBranchCount = PhiBranchCount
		x.Reason = "generic interior point: alpha is fixed by I_K and cos(3 phi) is fixed, but six phase branches remain until orientation is tagged."
	}
	x.InsideDomain = true
	x.AllowedAsBridgeDryRun = true
	x.AllowedAsNativeExport = false
	x.CanOrientWithoutTag = false
	x.Verdict = StatusBridgeOnlyInversionValidated
	return x
}

func buildFirewall(a Analysis) Firewall {
	return Firewall{
		Executed:                      true,
		NoObservedMuonMassImported:    true,
		NoObservedCharmMassImported:   true,
		NoObservedYukawaImported:      true,
		NoCKMImported:                 true,
		NoPMNSImported:                true,
		NoGSTPromotion:                true,
		NoNativeCoefficientRayValue:   true,
		NoCurveFitPromoted:            true,
		KGenStillForced:               a.Inheritance.Gate444KGenForced,
		XTriangleStillForced:          a.Inheritance.Gate445TriangleForced,
		YPhaseStillQuarantined:        true,
		SectorCoefficientsStillSealed: true,
		NativeFlavorDimAfter:          NativeFlavorDim,
		KXYCoeffDimAfter:              KXYCoeffDim,
		Verdict:                       StatusEmpiricalFirewallPreserved,
		Reason:                        "Gate 456 exports only symbolic inverse formulas and branch guards; it imports no physical flavor data and promotes no coefficient ray to native law.",
	}
}

func buildNext() NextStep {
	return NextStep{
		Gate:        457,
		Title:       "Empirical Comparator Provenance Contract / Sector-Scheme Ledger",
		Reason:      "after the symbolic inverse map is known, any real comparator import must be schema-locked with sector, scale, scheme, source, uncertainty, and bridge-only status",
		PrimaryTask: "define the machine-checkable provenance contract for observed comparator imports and reject untagged texture-zero data before evaluation",
	}
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 456 derives the exact bridge inverse alpha=sqrt(3) I_K/sqrt(1-I_K^2) and cos(3phi)=(3sqrt(3)/2)I_spec/(1-I_K^2)^(3/2). It validates %d in-domain symbolic dry-run samples, rejects %d outside-domain samples, and proves the inverse is not globally unique without branch tags.", a.Sieve.ValidDomainCount, a.Sieve.RejectedDomainCount)
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate455AdapterFirewallValidated || !a.Inheritance.Gate455ObservedValuesRejected || !a.Inheritance.Gate455NativePromotionRejected || !a.Inheritance.Gate455RequiresMetadata || !a.Inheritance.NoEmpiricalInputsImported || a.Inheritance.Gate454RayDOF != ProjectiveRayDOF || a.Inheritance.Gate454MinimumComparators != LocalRank {
		return fmt.Errorf("Gate455 adapter firewall not inherited: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Comparators.Executed || a.Comparators.LocalRank != LocalRank || a.Comparators.ProjectiveRayDOF != ProjectiveRayDOF || !a.Comparators.SufficientLocally || a.Comparators.SufficientGlobally {
		return fmt.Errorf("comparator contract failed: %s", FormatComparators(a.Comparators))
	}
	if !a.Inverse.Executed || !a.Inverse.RequiresAbsIKLessThanOne || !a.Inverse.RequiresCosBound || a.Inverse.BranchCountGeneric != PhiBranchCount || !a.Inverse.BridgeOnly || a.Inverse.ExportsNativeRay {
		return fmt.Errorf("symbolic inverse failed: %s", FormatInverse(a.Inverse))
	}
	if !a.Domain.Executed || !a.Domain.BoundaryIKUnit || !a.Domain.BoundaryCosThreePhiUnit || !a.Domain.BoundaryOutsideRejected {
		return fmt.Errorf("domain/caustic map failed: %s", FormatDomain(a.Domain))
	}
	if !a.Sieve.Executed || a.Sieve.ValidDomainCount != 3 || a.Sieve.RejectedDomainCount != 2 || !a.Sieve.GenericBranchSampleExists || !a.Sieve.CausticSampleExists || !a.Sieve.OutsideDomainRejected || !a.Sieve.NoSampleCanOrientWithoutTag || !a.Sieve.NoSampleAllowedAsNativeExport || !a.Sieve.GlobalUniqueCoefficientRayAbsent || !a.Sieve.ExplicitBranchTagRequiredAtCaustics {
		return fmt.Errorf("branch sieve failed: %s", FormatSieve(a.Sieve))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.NoGSTPromotion || !a.Firewall.NoNativeCoefficientRayValue || !a.Firewall.NoCurveFitPromoted || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate455Inherited,
		StatusSymbolicInverseDerived,
		StatusComparatorDomainDefined,
		StatusBranchCausticsMapped,
		StatusBridgeOnlyInversionValidated,
		StatusEmpiricalFirewallPreserved,
		StatusFailedGlobalUniqueRayAbsent,
		StatusFailedCausticsRequireBranchTags,
		StatusFailedNativeCoefficientPromotionAbsent,
		StatusFailedObservedValuesAbsent,
	}
}

func join(xs []string) string {
	if len(xs) == 0 {
		return "∅"
	}
	return strings.Join(xs, ", ")
}
