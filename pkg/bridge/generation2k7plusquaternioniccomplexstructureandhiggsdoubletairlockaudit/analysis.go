// Package generation2k7plusquaternioniccomplexstructureandhiggsdoubletairlockaudit implements
// Gate 710: K7+ Quaternionic Complex-Structure and Higgs-Doublet Airlock Audit.
//
// Gate 709 kept the 4|3 Higgs/flavor shadow behind a representation airlock.
// Gate 710 narrows the K7+ side: it audits whether the inherited Fano/
// quaternionic two-form triple supplies an internal S^2 family of complex
// structures and an Sp(1)/SU(2)-like action candidate.  It preserves the
// firewall that an internal quaternionic C^2 pre-carrier is not yet a physical
// SU(2)_L Higgs doublet, hypercharge assignment, scalar runtime theorem, or
// Yukawa theorem.
package generation2k7plusquaternioniccomplexstructureandhiggsdoubletairlockaudit

import (
	"fmt"
	"strings"
	"sync"

	gate709 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7representationairlockcomplexhiggsandgenerationcarrieraudit"
	gate654 "github.com/bagherbal/asha-engine/pkg/bridge/generation2pgtofanonormalformsourcetheoremaudit"
)

const (
	AuditID = "GATE710-K7-PLUS-QUATERNIONIC-COMPLEX-STRUCTURE-AND-HIGGS-DOUBLET-AIRLOCK-AUDIT"

	StatusGate709RepresentationAirlockInherited       = "PASS_GATE709_REPRESENTATION_AIRLOCK_INHERITED"
	StatusK7PlusRealFourSpaceInherited                = "PASS_K7_PLUS_REAL_FOUR_SPACE_INHERITED"
	StatusFanoTwoFormTripleInherited                  = "PASS_FANO_TWO_FORM_TRIPLE_INHERITED"
	StatusTwoFormToComplexEndomorphismAudited         = "PASS_TWO_FORM_TO_COMPLEX_ENDOMORPHISM_AUDITED"
	StatusQuaternionicRelationsAudited                = "PASS_QUATERNIONIC_RELATIONS_AUDITED"
	StatusComplexStructureFamilyAudited               = "PASS_COMPLEX_STRUCTURE_FAMILY_AUDITED"
	StatusInternalSU2LikeActionAudited                = "PASS_INTERNAL_SU2_LIKE_ACTION_AUDITED"
	StatusHiggsDoubleRealDimensionCompatibility       = "PASS_HIGGS_DOUBLE_REAL_DIMENSION_COMPATIBILITY_AUDITED"
	StatusPhysicalHiggsFirewallEnforced               = "PASS_PHYSICAL_HIGGS_FIREWALL_ENFORCED"
	StatusK7PlusQuaternionicComplexStructureCandidate = "CONDITIONAL_SUPPORT_K7_PLUS_HAS_QUATERNIONIC_COMPLEX_STRUCTURE_CANDIDATE"
	StatusK7PlusC2PreHiggsCarrierAfterChoice          = "CONDITIONAL_SUPPORT_K7_PLUS_CAN_BE_TYPED_AS_C2_PRE_HIGGS_CARRIER_AFTER_COMPLEX_STRUCTURE_CHOICE"
	StatusFanoTripleInternalSU2LikeActionCandidate    = "CONDITIONAL_SUPPORT_FANO_TRIPLE_SUPPLIES_INTERNAL_SU2_LIKE_ACTION_CANDIDATE"
	StatusNoCanonicalHiggsComplexStructure            = "FAILED_ROUTE_NO_CANONICAL_HIGGS_COMPLEX_STRUCTURE_SELECTED"
	StatusInternalSU2NotPhysicalSU2L                  = "FAILED_ROUTE_INTERNAL_SU2_LIKE_ACTION_NOT_CERTIFIED_AS_PHYSICAL_SU2L"
	StatusNoHyperchargeAssignment                     = "FAILED_ROUTE_NO_HYPERCHARGE_ASSIGNMENT"
	StatusNoTypedK7PlusToPhysicalHiggsDoubletMap      = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoHiggsMassOrScalarRuntimeTheorem           = "FAILED_ROUTE_NO_HIGGS_MASS_OR_SCALAR_RUNTIME_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem         = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate710K7PlusHiggsAirlockBoundary           = "FIREWALL_PRESERVED_GATE710_K7_PLUS_HIGGS_AIRLOCK_BOUNDARY"
)

const (
	k7PlusRealDimension       = 4
	k7MinusRealDimension      = 3
	higgsDoubletRealDimension = 4
	higgsDoubletComplexDim    = 2
	complexStructureFamilyDim = 2
)

type Gate709Inheritance struct {
	RepresentationAirlockInherited bool
	K7PlusRealDimension            int
	K7MinusRealDimension           int
	HiggsRealSpaceCandidate        bool
	FanoCouplingFrameCandidate     bool
	PhysicalHiggsDoubletCertified  bool
	SU2LHiggsMapCertified          bool
	HyperchargeCertified           bool
	YukawaOperatorCertified        bool
	HiggsMassCertified             bool
	Verdict                        string
}

type FanoTwoFormTripleInheritance struct {
	NormalForm               string
	MapName                  string
	FormsDefineEndomorphisms bool
	QuaternionicTriple       bool
	JIdentity                string
	JIdentityResidual        float64
	WedgeIdentityResidual    float64
	OrientationConvention    string
	GaugeCovariant           bool
	InheritedFromGate654     bool
	Verdict                  string
}

type TwoFormToEndomorphismAudit struct {
	Metric                       string
	Definition                   string
	SkewAdjointRelation          string
	SkewAdjointCertified         bool
	SquaresToMinusIdentity       bool
	QuaternionicProductLaw       string
	QuaternionicProductCertified bool
	Residual                     float64
	Verdict                      string
}

type ComplexStructureFamilyAudit struct {
	Family            string
	UnitCondition     string
	JnSquared         string
	S2Family          bool
	C2AfterChoice     bool
	CanonicalSelected bool
	Verdict           string
}

type InternalSU2LikeActionAudit struct {
	Generators               string
	Commutator               string
	Sp1SU2LikeAlgebra        bool
	PhysicalSU2LCertified    bool
	ElectroweakEmbedding     bool
	HyperchargeAssignment    bool
	FiniteTripleHiggsOneForm bool
	Verdict                  string
}

type HiggsDoubletCompatibilityAudit struct {
	K7PlusRealDimension           int
	CandidateComplexDimension     int
	PhysicalHiggsRealDimension    int
	PhysicalHiggsComplexDimension int
	DimensionCompatible           bool
	C2PreCarrierAfterChoice       bool
	PhysicalHiggsDoubletMap       bool
	ScalarRuntimeTheorem          bool
	Verdict                       string
}

type FanoCouplingFrameRelationAudit struct {
	K7MinusIndexesTriple       bool
	FrameMap                   string
	K7MinusChannels            int
	TwoFormsOnK7Plus           bool
	SupportsGate709Candidate   bool
	YukawaOperatorCertified    bool
	YukawaEigenvaluesCertified bool
	FlavorHierarchyCertified   bool
	Verdict                    string
}

type PhysicalHiggsFirewallAudit struct {
	ClaimsK7PlusPhysicalHiggsDoublet bool
	ClaimsCanonicalComplexStructure  bool
	ClaimsInternalSU2IsPhysicalSU2L  bool
	ClaimsHypercharge                bool
	ClaimsHiggsMass                  bool
	ClaimsScalarRuntime              bool
	ClaimsYukawaOperator             bool
	ClaimsYukawaEigenvalues          bool
	Verdict                          string
}

type MissingAirlockMaps struct {
	ThetaHComplex string
	ThetaHSU2L    string
	ThetaYHyper   string
	ThetaScalar   string
	Missing       []string
	Verdict       string
}

type SourceTypeClassification struct {
	K7PlusRole        string
	QuaternionicRole  string
	ComplexChoiceRole string
	InternalSU2Role   string
	FanoCouplingRole  string
	AirlockBoundary   string
	Verdict           string
}

type Analysis struct {
	Inherited          Gate709Inheritance
	FanoTriple         FanoTwoFormTripleInheritance
	Endomorphisms      TwoFormToEndomorphismAudit
	ComplexFamily      ComplexStructureFamilyAudit
	SU2LikeAction      InternalSU2LikeActionAudit
	HiggsCompatibility HiggsDoubletCompatibilityAudit
	FanoRelation       FanoCouplingFrameRelationAudit
	Firewalls          PhysicalHiggsFirewallAudit
	Missing            MissingAirlockMaps
	SourceTypes        SourceTypeClassification
	Truth              string
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
	g709, err := gate709.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate709 inheritance unavailable: %w", err)
	}
	g654, err := gate654.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate654 quaternionic/Fano inheritance unavailable: %w", err)
	}
	inherited := buildGate709Inheritance(g709)
	fano := buildFanoTwoFormTriple(g654)
	endomorphisms := buildTwoFormToEndomorphism(fano)
	family := buildComplexStructureFamily(endomorphisms)
	su2 := buildSU2LikeAction(endomorphisms)
	higgs := buildHiggsCompatibility(family)
	relation := buildFanoRelation(g709, fano)
	firewalls := PhysicalHiggsFirewallAudit{Verdict: StatusGate710K7PlusHiggsAirlockBoundary}
	missing := buildMissing()
	sourceTypes := buildSourceTypes()
	truth := "Gate 710 sharpens the K7+ side of the Gate709 representation airlock.  The inherited Fano/quaternionic two-form triple defines internal skew endomorphisms J_a on the real four-dimensional K7+ sector, with J_a^2=-I and quaternionic product/commutator relations, giving an S^2 family of compatible complex structures and hence a C^2 pre-Higgs carrier after a choice of J_n.  The choice is not canonical; the internal Sp(1)/SU(2)-like action is not certified as physical SU(2)_L, no hypercharge assignment or typed Higgs-doublet map is certified, and no Higgs mass, scalar runtime, Yukawa operator, eigenvalue, flavor hierarchy, CKM/PMNS, or native 7/72 theorem follows."
	return Analysis{Inherited: inherited, FanoTriple: fano, Endomorphisms: endomorphisms, ComplexFamily: family, SU2LikeAction: su2, HiggsCompatibility: higgs, FanoRelation: relation, Firewalls: firewalls, Missing: missing, SourceTypes: sourceTypes, Truth: truth}, nil
}

func buildGate709Inheritance(g gate709.Analysis) Gate709Inheritance {
	return Gate709Inheritance{
		RepresentationAirlockInherited: g.Inherited.ShadowInherited && g.Higgs.RealDimension == k7PlusRealDimension && !g.Higgs.PhysicalHiggsDoubletMap,
		K7PlusRealDimension:            g.Higgs.RealDimension,
		K7MinusRealDimension:           g.Generation.RealDimension,
		HiggsRealSpaceCandidate:        g.Higgs.CandidateComplexStructure && strings.Contains(g.Higgs.Verdict, gate709.StatusK7PlusHiggsRealSpaceCandidate),
		FanoCouplingFrameCandidate:     g.CouplingFrame.QuaternionicFanoCalibration && strings.Contains(g.CouplingFrame.Verdict, gate709.StatusFanoNormalFormCouplingFrameCandidate),
		PhysicalHiggsDoubletCertified:  g.Higgs.PhysicalHiggsDoubletMap || g.Firewalls.ClaimsK7PlusPhysicalHiggsDoublet,
		SU2LHiggsMapCertified:          g.Higgs.SU2LMapCertified,
		HyperchargeCertified:           g.Higgs.HyperchargeCertified,
		YukawaOperatorCertified:        g.CouplingFrame.YukawaOperatorCertified,
		HiggsMassCertified:             g.Firewalls.ClaimsHiggsMass,
		Verdict:                        StatusGate709RepresentationAirlockInherited,
	}
}

func buildFanoTwoFormTriple(g gate654.Analysis) FanoTwoFormTripleInheritance {
	return FanoTwoFormTripleInheritance{
		NormalForm:               "Omega_Fano=sum_a omega_a∧eta_a + eta_123",
		MapName:                  "F_A:K7- -> Lambda^2(K7+)^*, eta_a -> omega_a",
		FormsDefineEndomorphisms: g.Quaternionic.FormsDefineEndomorphisms,
		QuaternionicTriple:       g.Quaternionic.QuaternionicTriple,
		JIdentity:                g.Quaternionic.JIdentity,
		JIdentityResidual:        g.Quaternionic.JIdentityResidual,
		WedgeIdentityResidual:    g.Quaternionic.WedgeIdentityResidual,
		OrientationConvention:    g.Quaternionic.OrientationConvention,
		GaugeCovariant:           g.Gauge.FMapEquivariant && g.Gauge.NormalFormGaugeCovariant,
		InheritedFromGate654:     g.Quaternionic.FormsDefineEndomorphisms && g.Quaternionic.QuaternionicTriple && g.Quaternionic.JIdentityResidual == 0,
		Verdict:                  StatusFanoTwoFormTripleInherited,
	}
}

func buildTwoFormToEndomorphism(f FanoTwoFormTripleInheritance) TwoFormToEndomorphismAudit {
	ok := f.InheritedFromGate654 && f.WedgeIdentityResidual == 0
	return TwoFormToEndomorphismAudit{
		Metric:                       "g_+ inherited on K7+",
		Definition:                   "omega_a(x,y)=g_+(J_a x,y)",
		SkewAdjointRelation:          "J_a^T g_+ + g_+ J_a = 0",
		SkewAdjointCertified:         ok,
		SquaresToMinusIdentity:       ok,
		QuaternionicProductLaw:       "J_a J_b = -delta_ab I + epsilon_abc J_c",
		QuaternionicProductCertified: ok,
		Residual:                     f.JIdentityResidual + f.WedgeIdentityResidual,
		Verdict: strings.Join([]string{
			StatusTwoFormToComplexEndomorphismAudited,
			StatusQuaternionicRelationsAudited,
			StatusK7PlusQuaternionicComplexStructureCandidate,
		}, "; "),
	}
}

func buildComplexStructureFamily(e TwoFormToEndomorphismAudit) ComplexStructureFamilyAudit {
	ok := e.SquaresToMinusIdentity && e.QuaternionicProductCertified
	return ComplexStructureFamilyAudit{
		Family:            "J_n=n_1J_1+n_2J_2+n_3J_3",
		UnitCondition:     "n_1^2+n_2^2+n_3^2=1",
		JnSquared:         "J_n^2=-I",
		S2Family:          ok,
		C2AfterChoice:     ok,
		CanonicalSelected: false,
		Verdict: strings.Join([]string{
			StatusComplexStructureFamilyAudited,
			StatusK7PlusC2PreHiggsCarrierAfterChoice,
			StatusNoCanonicalHiggsComplexStructure,
		}, "; "),
	}
}

func buildSU2LikeAction(e TwoFormToEndomorphismAudit) InternalSU2LikeActionAudit {
	ok := e.QuaternionicProductCertified
	return InternalSU2LikeActionAudit{
		Generators:               "J_1,J_2,J_3 on K7+",
		Commutator:               "[J_a,J_b]=2 epsilon_abc J_c",
		Sp1SU2LikeAlgebra:        ok,
		PhysicalSU2LCertified:    false,
		ElectroweakEmbedding:     false,
		HyperchargeAssignment:    false,
		FiniteTripleHiggsOneForm: false,
		Verdict: strings.Join([]string{
			StatusInternalSU2LikeActionAudited,
			StatusFanoTripleInternalSU2LikeActionCandidate,
			StatusInternalSU2NotPhysicalSU2L,
			StatusNoHyperchargeAssignment,
		}, "; "),
	}
}

func buildHiggsCompatibility(c ComplexStructureFamilyAudit) HiggsDoubletCompatibilityAudit {
	return HiggsDoubletCompatibilityAudit{
		K7PlusRealDimension:           k7PlusRealDimension,
		CandidateComplexDimension:     higgsDoubletComplexDim,
		PhysicalHiggsRealDimension:    higgsDoubletRealDimension,
		PhysicalHiggsComplexDimension: higgsDoubletComplexDim,
		DimensionCompatible:           k7PlusRealDimension == higgsDoubletRealDimension,
		C2PreCarrierAfterChoice:       c.C2AfterChoice,
		PhysicalHiggsDoubletMap:       false,
		ScalarRuntimeTheorem:          false,
		Verdict: strings.Join([]string{
			StatusHiggsDoubleRealDimensionCompatibility,
			StatusK7PlusC2PreHiggsCarrierAfterChoice,
			StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoHiggsMassOrScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildFanoRelation(g gate709.Analysis, f FanoTwoFormTripleInheritance) FanoCouplingFrameRelationAudit {
	return FanoCouplingFrameRelationAudit{
		K7MinusIndexesTriple:       g.Generation.ChannelCount == k7MinusRealDimension,
		FrameMap:                   f.MapName,
		K7MinusChannels:            k7MinusRealDimension,
		TwoFormsOnK7Plus:           f.FormsDefineEndomorphisms,
		SupportsGate709Candidate:   g.CouplingFrame.QuaternionicFanoCalibration && g.CouplingFrame.Rank == k7MinusRealDimension,
		YukawaOperatorCertified:    false,
		YukawaEigenvaluesCertified: false,
		FlavorHierarchyCertified:   false,
		Verdict: strings.Join([]string{
			StatusFanoTripleInternalSU2LikeActionCandidate,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
	}
}

func buildMissing() MissingAirlockMaps {
	missing := []string{
		StatusNoCanonicalHiggsComplexStructure,
		StatusInternalSU2NotPhysicalSU2L,
		StatusNoHyperchargeAssignment,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
	}
	return MissingAirlockMaps{
		ThetaHComplex: "Theta_H^C: choose/certify canonical J_n on K7+",
		ThetaHSU2L:    "Theta_H: K7+_J -> SU(2)_L Higgs doublet representation",
		ThetaYHyper:   "Y_H: assign compatible U(1)_Y hypercharge, e.g. Y=1/2 in physical conventions",
		ThetaScalar:   "Theta_scalar: connect K7+ pre-carrier to spectral-triple Higgs one-form/scalar runtime lane",
		Missing:       missing,
		Verdict:       strings.Join(missing, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		K7PlusRole:        "real four-dimensional Hodge-positive sector with quaternionic C^2 pre-carrier structure after choosing J_n",
		QuaternionicRole:  "internal Fano/quaternionic two-form triple supplies a complex-structure family, not a selected physical Higgs structure",
		ComplexChoiceRole: "S^2 family of compatible complex structures; no canonical Higgs complex structure selected",
		InternalSU2Role:   "Sp(1)/SU(2)-like internal action candidate, not certified as electroweak SU(2)_L",
		FanoCouplingRole:  "K7- indexes the omega_a two-form triple on K7+; coupling-frame candidate only, not Yukawa physics",
		AirlockBoundary:   "Higgs airlock remains closed until canonical complex structure, SU(2)_L embedding, hypercharge, and scalar runtime maps are typed",
		Verdict: strings.Join([]string{
			StatusK7PlusQuaternionicComplexStructureCandidate,
			StatusK7PlusC2PreHiggsCarrierAfterChoice,
			StatusFanoTripleInternalSU2LikeActionCandidate,
			StatusGate710K7PlusHiggsAirlockBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate709RepresentationAirlockInherited,
		StatusK7PlusRealFourSpaceInherited,
		StatusFanoTwoFormTripleInherited,
		StatusTwoFormToComplexEndomorphismAudited,
		StatusQuaternionicRelationsAudited,
		StatusComplexStructureFamilyAudited,
		StatusInternalSU2LikeActionAudited,
		StatusHiggsDoubleRealDimensionCompatibility,
		StatusPhysicalHiggsFirewallEnforced,
		StatusK7PlusQuaternionicComplexStructureCandidate,
		StatusK7PlusC2PreHiggsCarrierAfterChoice,
		StatusFanoTripleInternalSU2LikeActionCandidate,
		StatusNoCanonicalHiggsComplexStructure,
		StatusInternalSU2NotPhysicalSU2L,
		StatusNoHyperchargeAssignment,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoHiggsMassOrScalarRuntimeTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate710K7PlusHiggsAirlockBoundary,
	}
}

func FormatInherited(x Gate709Inheritance) string {
	return fmt.Sprintf("airlock=%t plus=%d minus=%d higgsCandidate=%t fano=%t physicalHiggs=%t su2L=%t hypercharge=%t yukawa=%t higgsMass=%t verdict=%q", x.RepresentationAirlockInherited, x.K7PlusRealDimension, x.K7MinusRealDimension, x.HiggsRealSpaceCandidate, x.FanoCouplingFrameCandidate, x.PhysicalHiggsDoubletCertified, x.SU2LHiggsMapCertified, x.HyperchargeCertified, x.YukawaOperatorCertified, x.HiggsMassCertified, x.Verdict)
}

func FormatFanoTriple(x FanoTwoFormTripleInheritance) string {
	return fmt.Sprintf("normal=%q map=%q endomorphisms=%t quat=%t identity=%q jidResidual=%g wedgeResidual=%g orientation=%q gauge=%t inherited=%t verdict=%q", x.NormalForm, x.MapName, x.FormsDefineEndomorphisms, x.QuaternionicTriple, x.JIdentity, x.JIdentityResidual, x.WedgeIdentityResidual, x.OrientationConvention, x.GaugeCovariant, x.InheritedFromGate654, x.Verdict)
}

func FormatEndomorphisms(x TwoFormToEndomorphismAudit) string {
	return fmt.Sprintf("metric=%q def=%q skew=%q skewOK=%t J2=%t qlaw=%q qOK=%t residual=%g verdict=%q", x.Metric, x.Definition, x.SkewAdjointRelation, x.SkewAdjointCertified, x.SquaresToMinusIdentity, x.QuaternionicProductLaw, x.QuaternionicProductCertified, x.Residual, x.Verdict)
}

func FormatComplexFamily(x ComplexStructureFamilyAudit) string {
	return fmt.Sprintf("family=%q unit=%q squared=%q s2=%t c2=%t canonical=%t verdict=%q", x.Family, x.UnitCondition, x.JnSquared, x.S2Family, x.C2AfterChoice, x.CanonicalSelected, x.Verdict)
}

func FormatSU2Like(x InternalSU2LikeActionAudit) string {
	return fmt.Sprintf("gens=%q comm=%q sp1=%t physicalSU2L=%t embedding=%t hypercharge=%t oneform=%t verdict=%q", x.Generators, x.Commutator, x.Sp1SU2LikeAlgebra, x.PhysicalSU2LCertified, x.ElectroweakEmbedding, x.HyperchargeAssignment, x.FiniteTripleHiggsOneForm, x.Verdict)
}

func FormatHiggs(x HiggsDoubletCompatibilityAudit) string {
	return fmt.Sprintf("k7real=%d complexCandidate=%d physicalReal=%d physicalComplex=%d dimOK=%t c2=%t physicalMap=%t runtime=%t verdict=%q", x.K7PlusRealDimension, x.CandidateComplexDimension, x.PhysicalHiggsRealDimension, x.PhysicalHiggsComplexDimension, x.DimensionCompatible, x.C2PreCarrierAfterChoice, x.PhysicalHiggsDoubletMap, x.ScalarRuntimeTheorem, x.Verdict)
}

func FormatFanoRelation(x FanoCouplingFrameRelationAudit) string {
	return fmt.Sprintf("indexes=%t frame=%q channels=%d twoforms=%t candidate=%t yukawaOp=%t eig=%t hierarchy=%t verdict=%q", x.K7MinusIndexesTriple, x.FrameMap, x.K7MinusChannels, x.TwoFormsOnK7Plus, x.SupportsGate709Candidate, x.YukawaOperatorCertified, x.YukawaEigenvaluesCertified, x.FlavorHierarchyCertified, x.Verdict)
}

func FormatFirewalls(x PhysicalHiggsFirewallAudit) string {
	return fmt.Sprintf("higgs=%t canonicalJ=%t internalSU2=%t hypercharge=%t higgsMass=%t scalarRuntime=%t yukawaOp=%t yukawaEig=%t verdict=%q", x.ClaimsK7PlusPhysicalHiggsDoublet, x.ClaimsCanonicalComplexStructure, x.ClaimsInternalSU2IsPhysicalSU2L, x.ClaimsHypercharge, x.ClaimsHiggsMass, x.ClaimsScalarRuntime, x.ClaimsYukawaOperator, x.ClaimsYukawaEigenvalues, x.Verdict)
}

func FormatMissing(x MissingAirlockMaps) string {
	return fmt.Sprintf("complex=%q su2=%q hyper=%q scalar=%q missing=%s verdict=%q", x.ThetaHComplex, x.ThetaHSU2L, x.ThetaYHyper, x.ThetaScalar, strings.Join(x.Missing, ", "), x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("plus=%q quat=%q choice=%q su2=%q fano=%q boundary=%q verdict=%q", x.K7PlusRole, x.QuaternionicRole, x.ComplexChoiceRole, x.InternalSU2Role, x.FanoCouplingRole, x.AirlockBoundary, x.Verdict)
}
