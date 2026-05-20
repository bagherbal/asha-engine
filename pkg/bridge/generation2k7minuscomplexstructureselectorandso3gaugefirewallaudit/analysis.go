// Package generation2k7minuscomplexstructureselectorandso3gaugefirewallaudit implements
// Gate 712: K7- Complex-Structure Selector and SO(3) Gauge Firewall Audit.
//
// Gate 711 showed that after choosing a compatible complex structure
// J_H=n_a J_a, the Hodge-positive K7+ carrier admits an internal U(2)-type
// socket. Gate 712 audits whether the Hodge-negative K7- sector supplies a
// canonical unit direction n selecting J_H, or whether the choice remains an
// SO(3)-gauge / vacuum-selector freedom. It preserves the firewalls blocking
// physical electroweak, Higgs, generation, Yukawa, flavor hierarchy, CKM/PMNS,
// and native 7/72 promotions.
package generation2k7minuscomplexstructureselectorandso3gaugefirewallaudit

import (
	"fmt"
	"strings"
	"sync"

	gate711 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7plusu2higgssocketandquaternioniccommutantaudit"
)

const (
	AuditID = "GATE712-K7-MINUS-COMPLEX-STRUCTURE-SELECTOR-AND-SO3-GAUGE-FIREWALL-AUDIT"

	StatusGate711U2SocketInherited                            = "PASS_GATE711_U2_SOCKET_INHERITED"
	StatusK7MinusToComplexStructureFamilyMapAudited           = "PASS_K7_MINUS_TO_COMPLEX_STRUCTURE_FAMILY_MAP_AUDITED"
	StatusSO3CovarianceOfK7MinusFrameAudited                  = "PASS_SO3_COVARIANCE_OF_K7_MINUS_FRAME_AUDITED"
	StatusSelectorCandidatesAudited                           = "PASS_SELECTOR_CANDIDATES_AUDITED"
	StatusK7MinusUnitDirectionWouldSelectJH                   = "CONDITIONAL_SUPPORT_K7_MINUS_UNIT_DIRECTION_WOULD_SELECT_JH"
	StatusU2SocketFamilyValuedOverS2                          = "CONDITIONAL_SUPPORT_U2_SOCKET_IS_FAMILY_VALUED_OVER_S2_OF_K7_MINUS_DIRECTIONS"
	StatusNoNativeK7MinusUnitVectorSelector                   = "FAILED_ROUTE_NO_NATIVE_K7_MINUS_UNIT_VECTOR_SELECTOR"
	StatusFanoVolumeOrFrameDoesNotSelectSingleAxis            = "FAILED_ROUTE_FANO_VOLUME_OR_FRAME_DOES_NOT_SELECT_SINGLE_AXIS"
	StatusBoundaryScalarAndHistoryScalarsDoNotSelectDirection = "FAILED_ROUTE_BOUNDARY_SCALAR_AND_HISTORY_SCALARS_DO_NOT_SELECT_K7_MINUS_DIRECTION"
	StatusNoCanonicalHiggsComplexStructureSelected            = "FAILED_ROUTE_NO_CANONICAL_HIGGS_COMPLEX_STRUCTURE_SELECTED"
	StatusNoTypedK7PlusToPhysicalHiggsDoubletMap              = "FAILED_ROUTE_NO_TYPED_K7_PLUS_TO_PHYSICAL_HIGGS_DOUBLET_MAP"
	StatusNoTypedK7MinusToPhysicalGenerationSpaceMap          = "FAILED_ROUTE_NO_TYPED_K7_MINUS_TO_PHYSICAL_GENERATION_SPACE_MAP"
	StatusNoYukawaOperatorOrEigenvalueTheorem                 = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate712K7MinusSelectorBoundary                      = "FIREWALL_PRESERVED_GATE712_K7_MINUS_SELECTOR_BOUNDARY"
)

const (
	k7MinusDimension = 3
	k7PlusDimension  = 4
	complexDim       = 2
	sphereDimension  = 2
)

type Gate711Inheritance struct {
	U2SocketInherited          bool
	K7PlusDimension            int
	K7MinusDimension           int
	CanSelectJH                bool
	U2SocketAfterChoice        bool
	InternalSU2SocketCandidate bool
	InternalU1SocketCandidate  bool
	CanonicalJHSelected        bool
	PhysicalElectroweakU2      bool
	PhysicalHiggsDoubletMap    bool
	HyperchargeAssignment      bool
	YukawaOperatorCertified    bool
	HiggsMassCertified         bool
	Verdict                    string
}

type ComplexStructureFamilyMapAudit struct {
	Domain                    string
	Codomain                  string
	Map                       string
	UnitSphereDimension       int
	K7MinusDimension          int
	JHSquared                 string
	UnitDirectionWouldSelect  bool
	NativeSelectorFound       bool
	FamilyValuedU2Socket      bool
	CanonicalComplexStructure bool
	Verdict                   string
}

type SO3CovarianceAudit struct {
	FrameTransform         string
	TwoFormTransform       string
	PreservesOmega         bool
	PreservesEta123        bool
	GaugeCovariantFrame    bool
	CanonicalOrderedFrame  bool
	SelectsSingleAxis      bool
	PhysicalGenerationAxis bool
	Verdict                string
}

type SelectorCandidate struct {
	Name           string
	Role           string
	HasDirection   bool
	NativeSelector bool
	Reason         string
}

type SelectorCandidatesAudit struct {
	Candidates              []SelectorCandidate
	NativeSelectorFound     bool
	BoundaryScalarsRejected bool
	HistoryScalarsRejected  bool
	ExternalSealOnly        bool
	Verdict                 string
}

type GaugeVsPhysicalSelectorFirewallAudit struct {
	FamilyValuedPreHiggsCarriers bool
	UniqueHiggsComplexStructure  bool
	PhysicalHiggsDoubletMap      bool
	PhysicalGenerationSpaceMap   bool
	YukawaOperatorMap            bool
	FlavorHierarchyTheorem       bool
	CKMPMNSThereom               bool
	Verdict                      string
}

type VacuumSelectorClassification struct {
	MissingObject            string
	SelectorType             string
	NativeTheoremCertified   bool
	EnvironmentalSealAllowed bool
	SealMustBeQuarantined    bool
	Verdict                  string
}

type FlavorRelationAudit struct {
	SO3BreakPattern      string
	ResemblesFlavorAxis  bool
	GenerationCarrierMap bool
	FlavorOrientationMap bool
	YukawaEigenvalueMap  bool
	CKMPMNSMap           bool
	Verdict              string
}

type MissingMaps struct {
	ThetaG  string
	ThetaJH string
	ThetaY  string
	Missing []string
	Verdict string
}

type SourceTypeClassification struct {
	K7MinusRole   string
	FanoFrameRole string
	BoundaryRole  string
	HistoryRole   string
	SelectorRole  string
	FirewallRole  string
	Verdict       string
}

type Analysis struct {
	Inherited      Gate711Inheritance
	FamilyMap      ComplexStructureFamilyMapAudit
	SO3Covariance  SO3CovarianceAudit
	Selectors      SelectorCandidatesAudit
	GaugeFirewall  GaugeVsPhysicalSelectorFirewallAudit
	VacuumSelector VacuumSelectorClassification
	FlavorRelation FlavorRelationAudit
	Missing        MissingMaps
	SourceTypes    SourceTypeClassification
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
	g711, err := gate711.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate711 inheritance unavailable: %w", err)
	}
	inherited := buildGate711Inheritance(g711)
	family := buildFamilyMap(inherited)
	so3 := buildSO3Covariance()
	selectors := buildSelectorCandidates()
	gaugeFirewall := buildGaugeFirewall(family)
	vacuum := buildVacuumSelector()
	flavor := buildFlavorRelation()
	missing := buildMissingMaps()
	sourceTypes := buildSourceTypes()
	truth := "Gate 712 audits the selector problem exposed by Gates 710-711.  A unit direction n in the three-dimensional K7- frame can select a compatible complex structure J_H=n_a J_a on K7+ and hence a family-valued internal U(2) socket.  The inherited Fano frame is SO(3)-covariant: it supplies a gauge-covariant three-channel frame, not a canonical unit vector or ordered physical generation frame.  Hodge polarity, Fano volume, the eta_a frame, boundary scalars, scalar-wall coordinates, history deficits, and current bridge seals do not supply a typed vector n_* in K7-.  Therefore the Higgs complex structure remains noncanonical, the K7+ U(2) socket remains family-valued over S^2, and no physical Higgs, generation, Yukawa, flavor hierarchy, CKM/PMNS, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, FamilyMap: family, SO3Covariance: so3, Selectors: selectors, GaugeFirewall: gaugeFirewall, VacuumSelector: vacuum, FlavorRelation: flavor, Missing: missing, SourceTypes: sourceTypes, Truth: truth}, nil
}

func buildGate711Inheritance(g gate711.Analysis) Gate711Inheritance {
	return Gate711Inheritance{
		U2SocketInherited:          g.Inherited.QuaternionicK7PlusInherited && g.U2Socket.SpanJHInternalU1Candidate && g.U2Socket.CommutantInternalSU2Candidate,
		K7PlusDimension:            g.Inherited.K7PlusRealDimension,
		K7MinusDimension:           g.K7Minus.K7MinusDimension,
		CanSelectJH:                g.K7Minus.CanSelectJH,
		U2SocketAfterChoice:        g.U2Socket.Dimension == 4 && g.ChosenJH.SelectedAfterChoice,
		InternalSU2SocketCandidate: g.U2Socket.CommutantInternalSU2Candidate,
		InternalU1SocketCandidate:  g.U2Socket.SpanJHInternalU1Candidate,
		CanonicalJHSelected:        g.ChosenJH.CanonicalSelected,
		PhysicalElectroweakU2:      g.U2Socket.PhysicalElectroweakU2,
		PhysicalHiggsDoubletMap:    g.Inherited.PhysicalHiggsDoubletMap,
		HyperchargeAssignment:      g.Inherited.HyperchargeCertified,
		YukawaOperatorCertified:    g.Inherited.YukawaOperatorCertified,
		HiggsMassCertified:         g.Inherited.HiggsMassCertified,
		Verdict:                    StatusGate711U2SocketInherited,
	}
}

func buildFamilyMap(i Gate711Inheritance) ComplexStructureFamilyMapAudit {
	unitSelects := i.CanSelectJH && i.K7MinusDimension == k7MinusDimension && i.U2SocketAfterChoice
	return ComplexStructureFamilyMapAudit{
		Domain:                    "S^2 of unit directions n in K7-",
		Codomain:                  "S^2 of compatible complex structures J_H(n) on K7+",
		Map:                       "n -> J_H(n)=n_a J_a",
		UnitSphereDimension:       sphereDimension,
		K7MinusDimension:          i.K7MinusDimension,
		JHSquared:                 "J_H(n)^2=-I for ||n||=1",
		UnitDirectionWouldSelect:  unitSelects,
		NativeSelectorFound:       false,
		FamilyValuedU2Socket:      unitSelects && !i.CanonicalJHSelected,
		CanonicalComplexStructure: false,
		Verdict: strings.Join([]string{
			StatusK7MinusToComplexStructureFamilyMapAudited,
			StatusK7MinusUnitDirectionWouldSelectJH,
			StatusU2SocketFamilyValuedOverS2,
			StatusNoNativeK7MinusUnitVectorSelector,
			StatusNoCanonicalHiggsComplexStructureSelected,
		}, "; "),
	}
}

func buildSO3Covariance() SO3CovarianceAudit {
	return SO3CovarianceAudit{
		FrameTransform:         "eta_a -> R_ab eta_b for R in SO(3)",
		TwoFormTransform:       "omega_a -> R_ab omega_b, preserving sum_a omega_a wedge eta_a + eta_123",
		PreservesOmega:         true,
		PreservesEta123:        true,
		GaugeCovariantFrame:    true,
		CanonicalOrderedFrame:  false,
		SelectsSingleAxis:      false,
		PhysicalGenerationAxis: false,
		Verdict: strings.Join([]string{
			StatusSO3CovarianceOfK7MinusFrameAudited,
			StatusFanoVolumeOrFrameDoesNotSelectSingleAxis,
		}, "; "),
	}
}

func buildSelectorCandidates() SelectorCandidatesAudit {
	candidates := []SelectorCandidate{
		{Name: "Hodge polarity", Role: "separates K7+ from K7-", HasDirection: false, NativeSelector: false, Reason: "no vector inside K7-"},
		{Name: "Fano volume eta_123", Role: "orients K7-", HasDirection: false, NativeSelector: false, Reason: "orientation is not a unit axis"},
		{Name: "Fano frame eta_a", Role: "three-channel frame", HasDirection: false, NativeSelector: false, Reason: "frame is defined only up to SO(3) covariance"},
		{Name: "Boundary scalar S_split", Role: "scalar boundary quotient", HasDirection: false, NativeSelector: false, Reason: "scalar carries no K7- direction"},
		{Name: "Scalar-wall airlock lambda", Role: "scalar normalization anchor", HasDirection: false, NativeSelector: false, Reason: "scalar carries no K7- direction"},
		{Name: "History deficits kappa_lambda,kappa_e", Role: "scalar bridge coordinates", HasDirection: false, NativeSelector: false, Reason: "no vector-valued map into K7-"},
		{Name: "OrientationBalanceSeal / flavor wall", Role: "external bridge candidate", HasDirection: false, NativeSelector: false, Reason: "no typed map into K7- yet"},
	}
	return SelectorCandidatesAudit{
		Candidates:              candidates,
		NativeSelectorFound:     false,
		BoundaryScalarsRejected: true,
		HistoryScalarsRejected:  true,
		ExternalSealOnly:        true,
		Verdict: strings.Join([]string{
			StatusSelectorCandidatesAudited,
			StatusNoNativeK7MinusUnitVectorSelector,
			StatusFanoVolumeOrFrameDoesNotSelectSingleAxis,
			StatusBoundaryScalarAndHistoryScalarsDoNotSelectDirection,
		}, "; "),
	}
}

func buildGaugeFirewall(f ComplexStructureFamilyMapAudit) GaugeVsPhysicalSelectorFirewallAudit {
	return GaugeVsPhysicalSelectorFirewallAudit{
		FamilyValuedPreHiggsCarriers: f.FamilyValuedU2Socket,
		UniqueHiggsComplexStructure:  false,
		PhysicalHiggsDoubletMap:      false,
		PhysicalGenerationSpaceMap:   false,
		YukawaOperatorMap:            false,
		FlavorHierarchyTheorem:       false,
		CKMPMNSThereom:               false,
		Verdict: strings.Join([]string{
			StatusU2SocketFamilyValuedOverS2,
			StatusNoCanonicalHiggsComplexStructureSelected,
			StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
			StatusNoTypedK7MinusToPhysicalGenerationSpaceMap,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate712K7MinusSelectorBoundary,
		}, "; "),
	}
}

func buildVacuumSelector() VacuumSelectorClassification {
	return VacuumSelectorClassification{
		MissingObject:            "n_* in K7-, ||n_*||=1, selecting J_H=J_{n_*}",
		SelectorType:             "vacuum/orientation selector candidate",
		NativeTheoremCertified:   false,
		EnvironmentalSealAllowed: true,
		SealMustBeQuarantined:    true,
		Verdict: strings.Join([]string{
			StatusNoNativeK7MinusUnitVectorSelector,
			StatusNoCanonicalHiggsComplexStructureSelected,
		}, "; "),
	}
}

func buildFlavorRelation() FlavorRelationAudit {
	return FlavorRelationAudit{
		SO3BreakPattern:      "choosing n in K7- would break internal SO(3) to SO(2)",
		ResemblesFlavorAxis:  true,
		GenerationCarrierMap: false,
		FlavorOrientationMap: false,
		YukawaEigenvalueMap:  false,
		CKMPMNSMap:           false,
		Verdict: strings.Join([]string{
			StatusK7MinusUnitDirectionWouldSelectJH,
			StatusNoTypedK7MinusToPhysicalGenerationSpaceMap,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
	}
}

func buildMissingMaps() MissingMaps {
	missing := []string{
		StatusNoNativeK7MinusUnitVectorSelector,
		StatusNoCanonicalHiggsComplexStructureSelected,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoTypedK7MinusToPhysicalGenerationSpaceMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
	}
	return MissingMaps{
		ThetaG:  "Theta_G: K7- -> physical generation carrier or typed family-label space",
		ThetaJH: "Theta_JH: K7- unit direction -> physical Higgs complex structure",
		ThetaY:  "Theta_Y: F_A/Omega -> Yukawa operator family and singular-value data",
		Missing: missing,
		Verdict: strings.Join(missing, "; "),
	}
}

func buildSourceTypes() SourceTypeClassification {
	return SourceTypeClassification{
		K7MinusRole:   "real three-dimensional internal frame whose unit directions can label complex-structure choices",
		FanoFrameRole: "SO(3)-covariant frame, not a canonical ordered physical generation axis",
		BoundaryRole:  "boundary/scalar-wall data are scalar coordinates and do not carry K7- direction information",
		HistoryRole:   "history deficits are scalar bridge readouts, not vector selectors in K7-",
		SelectorRole:  "missing n_* is a vacuum/orientation selector candidate, native theorem not certified",
		FirewallRole:  "family-valued internal U(2) socket remains nonphysical until representation and Yukawa maps are typed",
		Verdict: strings.Join([]string{
			StatusK7MinusUnitDirectionWouldSelectJH,
			StatusU2SocketFamilyValuedOverS2,
			StatusNoNativeK7MinusUnitVectorSelector,
			StatusGate712K7MinusSelectorBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate711U2SocketInherited,
		StatusK7MinusToComplexStructureFamilyMapAudited,
		StatusSO3CovarianceOfK7MinusFrameAudited,
		StatusSelectorCandidatesAudited,
		StatusK7MinusUnitDirectionWouldSelectJH,
		StatusU2SocketFamilyValuedOverS2,
		StatusNoNativeK7MinusUnitVectorSelector,
		StatusFanoVolumeOrFrameDoesNotSelectSingleAxis,
		StatusBoundaryScalarAndHistoryScalarsDoNotSelectDirection,
		StatusNoCanonicalHiggsComplexStructureSelected,
		StatusNoTypedK7PlusToPhysicalHiggsDoubletMap,
		StatusNoTypedK7MinusToPhysicalGenerationSpaceMap,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate712K7MinusSelectorBoundary,
	}
}

func FormatInherited(x Gate711Inheritance) string {
	return fmt.Sprintf("u2Inherited=%t plusDim=%d minusDim=%d canSelect=%t u2AfterChoice=%t su2=%t u1=%t canonical=%t physicalU2=%t higgsMap=%t hypercharge=%t yukawa=%t higgsMass=%t verdict=%q", x.U2SocketInherited, x.K7PlusDimension, x.K7MinusDimension, x.CanSelectJH, x.U2SocketAfterChoice, x.InternalSU2SocketCandidate, x.InternalU1SocketCandidate, x.CanonicalJHSelected, x.PhysicalElectroweakU2, x.PhysicalHiggsDoubletMap, x.HyperchargeAssignment, x.YukawaOperatorCertified, x.HiggsMassCertified, x.Verdict)
}

func FormatFamilyMap(x ComplexStructureFamilyMapAudit) string {
	return fmt.Sprintf("domain=%q codomain=%q map=%q sphereDim=%d minusDim=%d jh=%q unitSelects=%t nativeSelector=%t family=%t canonical=%t verdict=%q", x.Domain, x.Codomain, x.Map, x.UnitSphereDimension, x.K7MinusDimension, x.JHSquared, x.UnitDirectionWouldSelect, x.NativeSelectorFound, x.FamilyValuedU2Socket, x.CanonicalComplexStructure, x.Verdict)
}

func FormatSO3(x SO3CovarianceAudit) string {
	return fmt.Sprintf("eta=%q omega=%q omegaPreserved=%t eta123=%t gaugeFrame=%t ordered=%t axis=%t physicalGenerationAxis=%t verdict=%q", x.FrameTransform, x.TwoFormTransform, x.PreservesOmega, x.PreservesEta123, x.GaugeCovariantFrame, x.CanonicalOrderedFrame, x.SelectsSingleAxis, x.PhysicalGenerationAxis, x.Verdict)
}

func FormatSelectors(x SelectorCandidatesAudit) string {
	return fmt.Sprintf("candidates=%d native=%t boundaryRejected=%t historyRejected=%t externalSeal=%t verdict=%q", len(x.Candidates), x.NativeSelectorFound, x.BoundaryScalarsRejected, x.HistoryScalarsRejected, x.ExternalSealOnly, x.Verdict)
}

func FormatGaugeFirewall(x GaugeVsPhysicalSelectorFirewallAudit) string {
	return fmt.Sprintf("family=%t unique=%t higgsMap=%t generationMap=%t yukawa=%t hierarchy=%t ckm=%t verdict=%q", x.FamilyValuedPreHiggsCarriers, x.UniqueHiggsComplexStructure, x.PhysicalHiggsDoubletMap, x.PhysicalGenerationSpaceMap, x.YukawaOperatorMap, x.FlavorHierarchyTheorem, x.CKMPMNSThereom, x.Verdict)
}

func FormatVacuumSelector(x VacuumSelectorClassification) string {
	return fmt.Sprintf("missing=%q type=%q native=%t environmental=%t quarantine=%t verdict=%q", x.MissingObject, x.SelectorType, x.NativeTheoremCertified, x.EnvironmentalSealAllowed, x.SealMustBeQuarantined, x.Verdict)
}

func FormatFlavorRelation(x FlavorRelationAudit) string {
	return fmt.Sprintf("break=%q resembles=%t generation=%t flavor=%t yukawa=%t ckm=%t verdict=%q", x.SO3BreakPattern, x.ResemblesFlavorAxis, x.GenerationCarrierMap, x.FlavorOrientationMap, x.YukawaEigenvalueMap, x.CKMPMNSMap, x.Verdict)
}

func FormatMissing(x MissingMaps) string {
	return fmt.Sprintf("thetaG=%q thetaJH=%q thetaY=%q missing=%d verdict=%q", x.ThetaG, x.ThetaJH, x.ThetaY, len(x.Missing), x.Verdict)
}

func FormatSourceTypes(x SourceTypeClassification) string {
	return fmt.Sprintf("k7minus=%q fano=%q boundary=%q history=%q selector=%q firewall=%q verdict=%q", x.K7MinusRole, x.FanoFrameRole, x.BoundaryRole, x.HistoryRole, x.SelectorRole, x.FirewallRole, x.Verdict)
}
