// Package generation2k7representationairlockcomplexhiggsandgenerationcarrieraudit implements
// Gate 709: K7 Representation Airlock: Complex-Higgs and Generation-Carrier Audit.
//
// Gate 708 recorded the native K7 Hodge split
//
//	K7 = K7+ ⊕ K7-, dim(K7+)=4, dim(K7-)=3
//
// as a Higgs/flavor shadow candidate only. Gate 709 audits the representation
// airlock that must be crossed before this internal 4|3 carrier can be promoted
// to a physical Higgs doublet or generation carrier.  It preserves the firewall:
// real dimension matches and Fano/quaternionic structure are not typed physical
// representation maps, Yukawa operators, eigenvalue theorems, or flavor
// hierarchy theorems.
package generation2k7representationairlockcomplexhiggsandgenerationcarrieraudit

import (
	"fmt"
	"strings"
	"sync"

	gate708 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7hodge43higgsflavorshadowfirewallaudit"
	gate654 "github.com/bagherbal/asha-engine/pkg/bridge/generation2pgtofanonormalformsourcetheoremaudit"
)

const (
	AuditID = "GATE709-K7-REPRESENTATION-AIRLOCK-COMPLEX-HIGGS-AND-GENERATION-CARRIER-AUDIT"

	StatusGate708HiggsFlavorShadowInherited      = "PASS_GATE708_HIGGS_FLAVOR_SHADOW_INHERITED"
	StatusK7PlusRealFourSpaceAudited             = "PASS_K7_PLUS_REAL_FOUR_SPACE_AUDITED"
	StatusK7MinusRealThreeChannelFrameAudited    = "PASS_K7_MINUS_REAL_THREE_CHANNEL_FRAME_AUDITED"
	StatusFanoCouplingFrameMapAudited            = "PASS_FANO_COUPLING_FRAME_MAP_AUDITED"
	StatusComplexificationFirewallAudited        = "PASS_COMPLEXIFICATION_FIREWALL_AUDITED"
	StatusPhysicalRepresentationFirewallEnforced = "PASS_PHYSICAL_REPRESENTATION_FIREWALL_ENFORCED"
	StatusK7PlusHiggsRealSpaceCandidate          = "CONDITIONAL_SUPPORT_K7_PLUS_IS_HIGGS_REAL_SPACE_CANDIDATE"
	StatusK7MinusFlavorChannelCandidate          = "CONDITIONAL_SUPPORT_K7_MINUS_IS_FLAVOR_CHANNEL_CANDIDATE"
	StatusFanoNormalFormCouplingFrameCandidate   = "CONDITIONAL_SUPPORT_FANO_NORMAL_FORM_IS_COUPLING_FRAME_CANDIDATE"
	StatusNoTypedK7PlusToSU2HiggsDoubletMap      = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_SU2_HIGGS_DOUBLET_MAP"
	StatusNoTypedK7MinusToComplexGenerationSpace = "FAILED_ROUTE_NO_TYPED_K7_MINUS_TO_COMPLEX_GENERATION_SPACE_MAP"
	StatusNoTypedFanoToYukawaOperatorMap         = "FAILED_ROUTE_NO_TYPED_FANO_TO_YUKAWA_OPERATOR_MAP"
	StatusNoYukawaEigenvalueOrFlavorHierarchy    = "FAILED_ROUTE_NO_YUKAWA_EIGENVALUE_OR_FLAVOR_HIERARCHY_THEOREM"
	StatusNoHiggsMassOrScalarRuntimeTheorem      = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusGate709RepresentationAirlockBoundary   = "FIREWALL_PRESERVED_GATE709_REPRESENTATION_AIRLOCK_BOUNDARY"
)

const (
	k7PlusRealDimension       = 4
	k7MinusRealDimension      = 3
	higgsDoubletRealDimension = 4
	complexGenerationRealDim  = 6
	couplingFrameSize         = k7PlusRealDimension * k7MinusRealDimension
)

type Gate708Inheritance struct {
	ShadowInherited             bool
	K7PlusDimension             int
	K7MinusDimension            int
	FanoCouplingFrameCandidate  bool
	PhysicalHiggsMapCertified   bool
	GenerationMapCertified      bool
	YukawaTheoremCertified      bool
	FlavorHierarchyCertified    bool
	CKMPMNSTheoremCertified     bool
	HiggsFlavorMapCertified     bool
	SevenOver72TheoremCertified bool
	Verdict                     string
}

type HiggsRealSpaceCompatibilityAudit struct {
	Domain                    string
	RealDimension             int
	HiggsDoubletRealDimension int
	QuaternionicTripleVisible bool
	CandidateComplexStructure bool
	SU2LikeInternalAction     bool
	SU2LMapCertified          bool
	HyperchargeCertified      bool
	ScalarPotentialCertified  bool
	QuarticLaneCertified      bool
	PhysicalHiggsDoubletMap   bool
	Verdict                   string
}

type GenerationCarrierCompatibilityAudit struct {
	Domain                         string
	RealDimension                  int
	ChannelCount                   int
	SO3CovariantInternalFrame      bool
	DiscreteFamilyLabelsCertified  bool
	ComplexGenerationSpaceRequired string
	ComplexGenerationRealDimension int
	C3GenerationMapCertified       bool
	FlavorHilbertFactorCertified   bool
	YukawaOperatorsCertified       bool
	Verdict                        string
}

type CouplingFrameAudit struct {
	NormalForm                  string
	MapName                     string
	Domain                      string
	Codomain                    string
	Rank                        int
	FrameSize                   int
	QuaternionicFanoCalibration bool
	SO3GaugeCovariant           bool
	ProtoYukawaOrientation      bool
	YukawaOperatorCertified     bool
	SingularValuesCertified     bool
	MixingMatricesCertified     bool
	Verdict                     string
}

type ComplexificationFirewallAudit struct {
	K7MinusRealDimension         int
	C3GenerationComplexDimension int
	C3GenerationRealDimension    int
	Real3EqualsComplex3          bool
	LabelSpacePossibleFuture     bool
	ComplexificationCertified    bool
	Verdict                      string
}

type PhysicalRepresentationFirewallAudit struct {
	ClaimsK7PlusPhysicalHiggsDoublet bool
	ClaimsK7MinusPhysicalGeneration  bool
	ClaimsOmegaYukawaMatrix          bool
	ClaimsFanoObservedFlavorTheorem  bool
	ClaimsFourPlusThreeDerivation    bool
	ClaimsHiggsMass                  bool
	ClaimsYukawaEigenvalues          bool
	ClaimsFlavorHierarchy            bool
	ClaimsCKMPMNS                    bool
	Verdict                          string
}

type MissingRepresentationMaps struct {
	ThetaH  string
	ThetaG  string
	ThetaY  string
	Missing []string
	Verdict string
}

type SourceTypeClassification struct {
	K7PlusRole       string
	K7MinusRole      string
	FanoRole         string
	QuaternionicRole string
	AirlockBoundary  string
	Verdict          string
}

type Analysis struct {
	Inherited        Gate708Inheritance
	Higgs            HiggsRealSpaceCompatibilityAudit
	Generation       GenerationCarrierCompatibilityAudit
	CouplingFrame    CouplingFrameAudit
	Complexification ComplexificationFirewallAudit
	Firewalls        PhysicalRepresentationFirewallAudit
	Missing          MissingRepresentationMaps
	SourceTypes      SourceTypeClassification
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
	g708, err := gate708.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate708 inheritance unavailable: %w", err)
	}
	g654, err := gate654.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate654 Fano/quaternionic source inheritance unavailable: %w", err)
	}
	inherited := buildGate708Inheritance(g708)
	higgs := buildHiggsRealSpace(g654)
	generation := buildGenerationCarrier(g654)
	coupling := buildCouplingFrame(g654)
	complexification := buildComplexification()
	firewalls := PhysicalRepresentationFirewallAudit{Verdict: StatusGate709RepresentationAirlockBoundary}
	missing := buildMissing()
	sourceTypes := buildSourceTypes()
	truth := "Gate 709 keeps the Gate708 4|3 Higgs/flavor shadow behind a representation airlock.  K7+ is a real four-dimensional Hodge-positive sector with an inherited quaternionic/Fano two-form triple candidate, but no typed SU(2)_L, hypercharge, scalar-potential, or physical Higgs-doublet map is certified.  K7- is a real three-channel SO(3)-covariant internal frame, not a C^3 generation space.  The Fano normal form supplies F_A:K7-→Λ^2(K7+)^* as an internal coupling-frame candidate only; it is not a Yukawa operator, singular-value theorem, flavor hierarchy theorem, CKM/PMNS theorem, Higgs mass theorem, or native 7/72 theorem."
	return Analysis{Inherited: inherited, Higgs: higgs, Generation: generation, CouplingFrame: coupling, Complexification: complexification, Firewalls: firewalls, Missing: missing, SourceTypes: sourceTypes, Truth: truth}, nil
}

func buildGate708Inheritance(g gate708.Analysis) Gate708Inheritance {
	return Gate708Inheritance{
		ShadowInherited:             g.Shadow.HiggsDimensionMatches && g.Shadow.FlavorTripletMatches && g.Shadow.OnlyDimensionShadow,
		K7PlusDimension:             g.Shadow.PlusDimension,
		K7MinusDimension:            g.Shadow.MinusDimension,
		FanoCouplingFrameCandidate:  g.FanoHitchin.CouplingFrameCandidate,
		PhysicalHiggsMapCertified:   g.Firewalls.ClaimsK7PlusIsPhysicalHiggsDoublet,
		GenerationMapCertified:      g.Firewalls.ClaimsK7MinusIsGenerationSpace,
		YukawaTheoremCertified:      g.Firewalls.ClaimsYukawaEigenvalueTheorem,
		FlavorHierarchyCertified:    g.Firewalls.ClaimsFlavorHierarchyTheorem,
		CKMPMNSTheoremCertified:     g.Firewalls.ClaimsCKMPMNSTheorem,
		HiggsFlavorMapCertified:     !strings.Contains(g.Missing.Verdict, gate708.StatusNoNativeHiggsFlavorRepresentationMap),
		SevenOver72TheoremCertified: !strings.Contains(g.Missing.Verdict, gate708.StatusNoNativeSevenOver72Theorem),
		Verdict:                     StatusGate708HiggsFlavorShadowInherited,
	}
}

func buildHiggsRealSpace(g gate654.Analysis) HiggsRealSpaceCompatibilityAudit {
	quatVisible := g.Quaternionic.FormsDefineEndomorphisms && g.Quaternionic.QuaternionicTriple && g.Quaternionic.JIdentityResidual == 0
	return HiggsRealSpaceCompatibilityAudit{
		Domain:                    "K7+",
		RealDimension:             k7PlusRealDimension,
		HiggsDoubletRealDimension: higgsDoubletRealDimension,
		QuaternionicTripleVisible: quatVisible,
		CandidateComplexStructure: quatVisible,
		SU2LikeInternalAction:     quatVisible,
		SU2LMapCertified:          false,
		HyperchargeCertified:      false,
		ScalarPotentialCertified:  false,
		QuarticLaneCertified:      false,
		PhysicalHiggsDoubletMap:   false,
		Verdict: strings.Join([]string{
			StatusK7PlusRealFourSpaceAudited,
			StatusK7PlusHiggsRealSpaceCandidate,
			StatusNoTypedK7PlusToSU2HiggsDoubletMap,
		}, "; "),
	}
}

func buildGenerationCarrier(g gate654.Analysis) GenerationCarrierCompatibilityAudit {
	return GenerationCarrierCompatibilityAudit{
		Domain:                         "K7-",
		RealDimension:                  k7MinusRealDimension,
		ChannelCount:                   k7MinusRealDimension,
		SO3CovariantInternalFrame:      g.Gauge.FMapEquivariant && g.Gauge.NormalFormGaugeCovariant,
		DiscreteFamilyLabelsCertified:  false,
		ComplexGenerationSpaceRequired: "C^3_generation has complex dimension 3 and real dimension 6",
		ComplexGenerationRealDimension: complexGenerationRealDim,
		C3GenerationMapCertified:       false,
		FlavorHilbertFactorCertified:   false,
		YukawaOperatorsCertified:       false,
		Verdict: strings.Join([]string{
			StatusK7MinusRealThreeChannelFrameAudited,
			StatusK7MinusFlavorChannelCandidate,
			StatusNoTypedK7MinusToComplexGenerationSpace,
		}, "; "),
	}
}

func buildCouplingFrame(g gate654.Analysis) CouplingFrameAudit {
	return CouplingFrameAudit{
		NormalForm:                  "Omega=sum_a omega_a∧eta_a + eta_123",
		MapName:                     "F_A:K7- -> Lambda^2(K7+)^*",
		Domain:                      g.AMap.Domain,
		Codomain:                    g.AMap.Codomain,
		Rank:                        g.AMap.Rank,
		FrameSize:                   couplingFrameSize,
		QuaternionicFanoCalibration: g.Quaternionic.QuaternionicTriple && g.AMap.WedgeOrthonormal,
		SO3GaugeCovariant:           g.Gauge.FMapEquivariant,
		ProtoYukawaOrientation:      true,
		YukawaOperatorCertified:     false,
		SingularValuesCertified:     false,
		MixingMatricesCertified:     false,
		Verdict: strings.Join([]string{
			StatusFanoCouplingFrameMapAudited,
			StatusFanoNormalFormCouplingFrameCandidate,
			StatusNoTypedFanoToYukawaOperatorMap,
		}, "; "),
	}
}

func buildComplexification() ComplexificationFirewallAudit {
	return ComplexificationFirewallAudit{
		K7MinusRealDimension:         k7MinusRealDimension,
		C3GenerationComplexDimension: 3,
		C3GenerationRealDimension:    complexGenerationRealDim,
		Real3EqualsComplex3:          false,
		LabelSpacePossibleFuture:     true,
		ComplexificationCertified:    false,
		Verdict: strings.Join([]string{
			StatusComplexificationFirewallAudited,
			StatusNoTypedK7MinusToComplexGenerationSpace,
		}, "; "),
	}
}

func buildMissing() MissingRepresentationMaps {
	missing := []string{
		StatusNoTypedK7PlusToSU2HiggsDoubletMap,
		StatusNoTypedK7MinusToComplexGenerationSpace,
		StatusNoTypedFanoToYukawaOperatorMap,
		StatusNoYukawaEigenvalueOrFlavorHierarchy,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
	}
	return MissingRepresentationMaps{
		ThetaH:  "Theta_H:K7+ -> SU(2)_L Higgs doublet representation with hypercharge",
		ThetaG:  "Theta_G:K7- -> complex generation carrier or typed family-label space",
		ThetaY:  "Theta_Y:F_A or Omega -> Yukawa operator family",
		Missing: missing,
		Verdict: strings.Join(missing, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		K7PlusRole:       "real four-dimensional Hodge-positive sector; Higgs-real-space candidate only",
		K7MinusRole:      "real three-channel Hodge-negative internal frame; flavor-channel candidate only, not C^3_generation",
		FanoRole:         "F_A maps the three-channel frame into two-forms on K7+; coupling-frame candidate only",
		QuaternionicRole: "quaternionic/Fano triple can supply a candidate complex structure/SU2-like internal action, not physical SU(2)_L×U(1)_Y",
		AirlockBoundary:  "representation airlock remains closed until Theta_H, Theta_G, and Theta_Y are typed",
		Verdict: strings.Join([]string{
			StatusK7PlusHiggsRealSpaceCandidate,
			StatusK7MinusFlavorChannelCandidate,
			StatusFanoNormalFormCouplingFrameCandidate,
			StatusPhysicalRepresentationFirewallEnforced,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate708HiggsFlavorShadowInherited,
		StatusK7PlusRealFourSpaceAudited,
		StatusK7MinusRealThreeChannelFrameAudited,
		StatusFanoCouplingFrameMapAudited,
		StatusComplexificationFirewallAudited,
		StatusPhysicalRepresentationFirewallEnforced,
		StatusK7PlusHiggsRealSpaceCandidate,
		StatusK7MinusFlavorChannelCandidate,
		StatusFanoNormalFormCouplingFrameCandidate,
		StatusNoTypedK7PlusToSU2HiggsDoubletMap,
		StatusNoTypedK7MinusToComplexGenerationSpace,
		StatusNoTypedFanoToYukawaOperatorMap,
		StatusNoYukawaEigenvalueOrFlavorHierarchy,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusGate709RepresentationAirlockBoundary,
	}
}

func FormatInherited(x Gate708Inheritance) string {
	return fmt.Sprintf("shadow=%t plus=%d minus=%d fano=%t higgsMap=%t genMap=%t yukawa=%t hierarchy=%t ckm=%t hfMap=%t seven72=%t verdict=%q", x.ShadowInherited, x.K7PlusDimension, x.K7MinusDimension, x.FanoCouplingFrameCandidate, x.PhysicalHiggsMapCertified, x.GenerationMapCertified, x.YukawaTheoremCertified, x.FlavorHierarchyCertified, x.CKMPMNSTheoremCertified, x.HiggsFlavorMapCertified, x.SevenOver72TheoremCertified, x.Verdict)
}

func FormatHiggs(x HiggsRealSpaceCompatibilityAudit) string {
	return fmt.Sprintf("domain=%q real=%d higgsReal=%d quat=%t complexCandidate=%t su2Like=%t su2L=%t hypercharge=%t potential=%t quartic=%t physicalMap=%t verdict=%q", x.Domain, x.RealDimension, x.HiggsDoubletRealDimension, x.QuaternionicTripleVisible, x.CandidateComplexStructure, x.SU2LikeInternalAction, x.SU2LMapCertified, x.HyperchargeCertified, x.ScalarPotentialCertified, x.QuarticLaneCertified, x.PhysicalHiggsDoubletMap, x.Verdict)
}

func FormatGeneration(x GenerationCarrierCompatibilityAudit) string {
	return fmt.Sprintf("domain=%q real=%d channels=%d so3=%t labels=%t required=%q c3real=%d c3map=%t hilbert=%t yukawaOps=%t verdict=%q", x.Domain, x.RealDimension, x.ChannelCount, x.SO3CovariantInternalFrame, x.DiscreteFamilyLabelsCertified, x.ComplexGenerationSpaceRequired, x.ComplexGenerationRealDimension, x.C3GenerationMapCertified, x.FlavorHilbertFactorCertified, x.YukawaOperatorsCertified, x.Verdict)
}

func FormatCouplingFrame(x CouplingFrameAudit) string {
	return fmt.Sprintf("normal=%q map=%q domain=%q codomain=%q rank=%d frameSize=%d fano=%t so3=%t proto=%t yukawaOp=%t singular=%t mixing=%t verdict=%q", x.NormalForm, x.MapName, x.Domain, x.Codomain, x.Rank, x.FrameSize, x.QuaternionicFanoCalibration, x.SO3GaugeCovariant, x.ProtoYukawaOrientation, x.YukawaOperatorCertified, x.SingularValuesCertified, x.MixingMatricesCertified, x.Verdict)
}

func FormatComplexification(x ComplexificationFirewallAudit) string {
	return fmt.Sprintf("k7minusReal=%d c3complex=%d c3real=%d real3eqC3=%t futureLabel=%t complexified=%t verdict=%q", x.K7MinusRealDimension, x.C3GenerationComplexDimension, x.C3GenerationRealDimension, x.Real3EqualsComplex3, x.LabelSpacePossibleFuture, x.ComplexificationCertified, x.Verdict)
}

func FormatFirewalls(x PhysicalRepresentationFirewallAudit) string {
	return fmt.Sprintf("higgs=%t gen=%t omegaYukawa=%t fanoFlavor=%t fourThree=%t higgsMass=%t yukawa=%t hierarchy=%t ckm=%t verdict=%q", x.ClaimsK7PlusPhysicalHiggsDoublet, x.ClaimsK7MinusPhysicalGeneration, x.ClaimsOmegaYukawaMatrix, x.ClaimsFanoObservedFlavorTheorem, x.ClaimsFourPlusThreeDerivation, x.ClaimsHiggsMass, x.ClaimsYukawaEigenvalues, x.ClaimsFlavorHierarchy, x.ClaimsCKMPMNS, x.Verdict)
}

func FormatMissing(x MissingRepresentationMaps) string {
	return fmt.Sprintf("thetaH=%q thetaG=%q thetaY=%q missing=%s verdict=%q", x.ThetaH, x.ThetaG, x.ThetaY, strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("plus=%q minus=%q fano=%q quat=%q boundary=%q verdict=%q", x.K7PlusRole, x.K7MinusRole, x.FanoRole, x.QuaternionicRole, x.AirlockBoundary, x.Verdict)
}
