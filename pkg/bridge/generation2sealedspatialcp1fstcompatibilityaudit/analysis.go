// Package generation2sealedspatialcp1fstcompatibilityaudit implements Gate 575:
// Sealed Spatial CP1 Compatibility with Finite Spectral Triple Audit.
//
// Gate 574 introduced the SpatialProjectiveOrientationSeal: a sealed projective
// point [u] in CP^2_sp, equivalently a rank-one projector P_u, sufficient and
// minimal for the projective split CP^2_sp -> CP^1|CP^0. Gate 575 works under
// that seal and asks a narrower compatibility question: does the sealed CP^1
// complement u^perp lawfully coincide with, or carry, the finite spectral-triple
// quaternionic weak-doublet structure? The result is an obstruction audit. The
// sealed CP^1 split exists algebraically and commutes with B-L because B-L is
// scalar on W_spatial, but current ASHA data supplies no Im(H)->su(u^perp)
// intertwiner, no H->End(u^perp) module structure, and no D/J/grading/first-
// order compatible finite spectral-triple carrier action on u^perp. The sealed
// CP^1 remains projective orientation data only.
package generation2sealedspatialcp1fstcompatibilityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate574 "github.com/bagherbal/asha-engine/pkg/bridge/generation2spatialprojectiveorientationsealminimalityconsequenceaudit"
)

const (
	AuditID = "GATE575-SEALED-SPATIAL-CP1-COMPATIBILITY-WITH-FINITE-SPECTRAL-TRIPLE-AUDIT"

	StatusGate574Inherited                      = "CONDITIONAL_SUPPORT_GATE574_SPATIAL_ORIENTATION_SEAL_INHERITED"
	StatusSealedDecompositionVerified           = "PASS_SEALED_SPATIAL_CP1_SPLIT_ALGEBRAICALLY_EXISTS"
	StatusProjectorPropertiesVerified           = "PASS_SEALED_RANK_ONE_PROJECTOR_AND_COMPLEMENT_PROPERTIES_VERIFIED"
	StatusBMinusLCompatibilityVacuous           = "CONDITIONAL_SUPPORT_SEALED_CP1_COMMUTES_WITH_B_MINUS_L_ONLY_BECAUSE_B_MINUS_L_IS_SCALAR"
	StatusBMinusLCommutatorZero                 = "PASS_B_MINUS_L_COMMUTES_WITH_SEALED_SPATIAL_PROJECTOR"
	StatusCommutantReconfirmed                  = "PASS_SEALED_SELECTOR_COMMUTANT_U2_PLUS_U1_DIMENSION_5_RECONFIRMED"
	StatusGate555SelectorFormulaReused          = "PASS_GATE555_SELECTOR_COMMUTANT_FORMULA_REUSED_FOR_SEALED_CP1"
	StatusQuaternionicSocketExistsElsewhere     = "CONDITIONAL_SUPPORT_IM_H_SOCKET_EXISTS_ON_SCALAR_HPHI_LANE"
	StatusNoImHToSealedCP1Intertwiner           = "FAILED_ROUTE_NO_IMH_TO_SEALED_SPATIAL_CP1_INTERTWINER"
	StatusNoHToEndUperpModule                   = "FAILED_ROUTE_NO_H_TO_END_U_PERP_MODULE_COMPATIBLE_WITH_SPATIAL_SEAL"
	StatusNoFiniteWeakDoubletCarrier            = "FAILED_ROUTE_SEALED_CP1_NOT_FINITE_WEAK_DOUBLET_CARRIER"
	StatusNoDJGradingFirstOrderCompatibility    = "FAILED_ROUTE_NO_D_J_GRADING_FIRST_ORDER_COMPATIBILITY_FOR_SEALED_CP1"
	StatusNoOneFormHiggsLaneCarrier             = "FAILED_ROUTE_SEALED_CP1_NOT_FINITE_ONE_FORM_HIGGS_LANE_CARRIER"
	StatusGate562ImHNotWSpatialPreserved        = "FIREWALL_PRESERVED_GATE562_IM_H_SOCKET_DOES_NOT_REOPEN_W_SPATIAL_TRANSFER"
	StatusRepresentativeU12NotPhysicalWeakPlane = "FAILED_ROUTE_REPRESENTATIVE_U12_NOT_PHYSICAL_WEAK_PLANE"
	StatusWeakPlaneRequiresFSTCompatibility     = "FAILED_ROUTE_PHYSICAL_WEAK_PLANE_REQUIRES_FST_QUATERNIONIC_D_J_GRADING_FIRST_ORDER_COMPATIBILITY"
	StatusNoFlavorEWObservedData                = "FAILED_ROUTE_SEALED_CP1_DOES_NOT_DERIVE_FLAVOR_OR_ELECTROWEAK_OBSERVED_DATA"
	StatusGate564565BoundaryPreserved           = "FIREWALL_PRESERVED_GATE564_GATE565_ELECTROWEAK_BRIDGE_SYMBOLIC_BOUNDARY"
	StatusK7TimeBoundaryPreserved               = "FIREWALL_PRESERVED_K7_TIME_OS_HILBERT_RG_BOUNDARY"
	StatusGate575BoundaryPreserved              = "FIREWALL_PRESERVED_GATE575_SEALED_SPATIAL_CP1_FST_COMPATIBILITY_BOUNDARY"
)

type InheritedGate574Audit struct {
	SealName                       string
	SealSufficient                 bool
	SealMinimal                    bool
	CP1CP0Split                    bool
	CommutantU2U1                  bool
	SealNative                     bool
	PhysicalWeakPlaneDerived       bool
	FlavorElectroweakDataDerived   bool
	K7OrProductTimeOpened          bool
	AdditionalTheoremRequiredPrior string
	Verdict                        string
}

type SealedDecompositionAudit struct {
	SealName                  string
	RepresentativeRay         string
	Projector                 [][]float64
	ComplementProjector       [][]float64
	ProjectorRank             int
	ProjectorTrace            float64
	ProjectorIdempotentError  float64
	ComplementRank            int
	ComplementTrace           float64
	ComplementIdempotentError float64
	OrthogonalityResidual     float64
	WSpatialDecomposition     string
	DimCUperp                 int
	DimCCu                    int
	ProjectiveSplit           string
	CP1CP0SplitExists         bool
	NativeWithoutSeal         bool
	Verdict                   string
}

type BMinusLCompatibilityAudit struct {
	BLRestriction           string
	BLSpatialEigenvalue     float64
	CommutatorResidual      float64
	CommutesWithPU          bool
	CommutesWithComplement  bool
	CompatibilityVacuous    bool
	SuppliesFurtherSelector bool
	Verdict                 string
}

type CommutantAudit struct {
	SelectorFormula       string
	MultiplicityPattern   string
	Commutant             string
	DimensionFormula      string
	Dimension             int
	MatchesGate555Formula bool
	SealedSupportOnly     bool
	Verdict               string
}

type QuaternionicSocketComparisonAudit struct {
	ImHSocketAvailableElsewhere     bool
	HPhiDoubletModuleAvailable      bool
	PauliHopfMomentQuaternionic     bool
	FiniteOneFormLinkedStructurally bool
	SourceCarrier                   string
	TestedTarget                    string
	ImHToSuUperpIntertwinerSupplied bool
	HToEndUperpModuleSupplied       bool
	CompatibleWithPU                bool
	WSpatialTransferBlocked         bool
	Verdict                         string
}

type FiniteSpectralTripleCarrierAudit struct {
	AFRepresentationStructural              bool
	FiniteWeakDoubletCarrierExistsElsewhere bool
	UperpUsedAsFiniteWeakDoubletCarrier     bool
	DCompatibilityForUperp                  bool
	JCompatibilityForUperp                  bool
	GradingCompatibilityForUperp            bool
	FirstOrderCompatibilityForUperp         bool
	OrderOneCarrierActionProven             bool
	Verdict                                 string
}

type OneFormHiggsLaneCompatibilityAudit struct {
	FiniteOneFormContainsScalarDoublet bool
	ImHActsOnHPhiStructurally          bool
	SealedCP1AppearsInOneFormLane      bool
	SealedCP1AppearsInHiggsLane        bool
	PauliRouteSeparateFromWSpatial     bool
	Gate562BoundaryPreserved           bool
	Verdict                            string
}

type PhysicalWeakPlaneFirewallAudit struct {
	RepresentativeGauge            string
	Complement                     string
	ConventionalName               string
	CanCallPhysicalWeakPlane       bool
	FiniteSpectralTripleCompatible bool
	QuaternionicCompatible         bool
	DCompatible                    bool
	JCompatible                    bool
	GradingCompatible              bool
	FirstOrderCompatible           bool
	Verdict                        string
}

type FlavorElectroweakFirewallAudit struct {
	GenerationHierarchyDerived bool
	YukawaTextureDerived       bool
	CKMPMNSDerived             bool
	ObservedFlavorImported     bool
	PhysicalEWDynamicsDerived  bool
	PhotonDynamicsDerived      bool
	WZMassesDerived            bool
	WeakIsospinDerived         bool
	Verdict                    string
}

type PreviousGateBoundaryAudit struct {
	TauEtaTraceShadowOnly              bool
	Q4ContactOnly                      bool
	PauliQuaternionicSocketNotWSpatial bool
	Gate564565BridgeSymbolic           bool
	K7TimeRoutesSealed                 bool
	OrientationSealProjectiveOnly      bool
	Verdict                            string
}

type FinalVerdict struct {
	SealedCP1SplitExistsAlgebraically bool
	CommutesWithBMinusL               bool
	CarriesNativeOrSealedImHAction    bool
	PartOfFiniteWeakDoubletCarrier    bool
	CanBeCalledPhysicalWeakPlane      bool
	DerivesFlavorOrEWObservedData     bool
	AdditionalTheoremRequired         string
	Verdict                           string
}

type Analysis struct {
	Inherited     InheritedGate574Audit
	Decomposition SealedDecompositionAudit
	BMinusL       BMinusLCompatibilityAudit
	Commutant     CommutantAudit
	Quaternionic  QuaternionicSocketComparisonAudit
	FiniteTriple  FiniteSpectralTripleCarrierAudit
	OneForm       OneFormHiggsLaneCompatibilityAudit
	WeakPlane     PhysicalWeakPlaneFirewallAudit
	FlavorEW      FlavorElectroweakFirewallAudit
	Boundaries    PreviousGateBoundaryAudit
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
	g574, err := gate574.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate574 spatial orientation seal predecessor: %w", err)
	}
	a := Analysis{}
	a.Inherited = auditInherited(g574)
	a.Decomposition = auditSealedDecomposition(g574)
	a.BMinusL = auditBMinusLCompatibility(a.Decomposition)
	a.Commutant = auditCommutant(g574)
	a.Quaternionic = auditQuaternionicSocket()
	a.FiniteTriple = auditFiniteSpectralTripleCarrier()
	a.OneForm = auditOneFormHiggsLane()
	a.WeakPlane = auditPhysicalWeakPlaneFirewall()
	a.FlavorEW = auditFlavorElectroweakFirewall()
	a.Boundaries = auditPreviousGateBoundaries(g574)
	a.Final = auditFinal(a)
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func auditInherited(g gate574.Analysis) InheritedGate574Audit {
	return InheritedGate574Audit{
		SealName:                       g.Seal.SealName,
		SealSufficient:                 g.Final.SealSufficient,
		SealMinimal:                    g.Final.SealMinimal,
		CP1CP0Split:                    g.Selector.ConstructsCP1CP0Split,
		CommutantU2U1:                  g.Commutant.Commutant == "u(2)+u(1)" && g.Commutant.Dimension == 5,
		SealNative:                     g.Seal.NativeDerived,
		PhysicalWeakPlaneDerived:       g.WeakPlane.ComplementaryCP1CanBeCalledPhysicalWeakPlane,
		FlavorElectroweakDataDerived:   g.Final.DerivesPhysicalWeakFlavorElectroweakData,
		K7OrProductTimeOpened:          g.Final.K7OrProductTimeOpened,
		AdditionalTheoremRequiredPrior: g.Final.AdditionalTheoremRequired,
		Verdict:                        StatusGate574Inherited,
	}
}

func auditSealedDecomposition(g gate574.Analysis) SealedDecompositionAudit {
	p := clone3(g.Seal.Projector)
	q := subtract3(identity3(), p)
	return SealedDecompositionAudit{
		SealName:                  g.Seal.SealName,
		RepresentativeRay:         "[u]=[a_3^dagger] representative gauge only",
		Projector:                 p,
		ComplementProjector:       q,
		ProjectorRank:             rankFromTrace(p),
		ProjectorTrace:            trace3(p),
		ProjectorIdempotentError:  projectorIdempotentResidual(p),
		ComplementRank:            rankFromTrace(q),
		ComplementTrace:           trace3(q),
		ComplementIdempotentError: projectorIdempotentResidual(q),
		OrthogonalityResidual:     productMaxAbs(p, q),
		WSpatialDecomposition:     "W_spatial = u^perp \u2295 C u",
		DimCUperp:                 2,
		DimCCu:                    1,
		ProjectiveSplit:           "CP^2_sp -> CP^1=P(u^perp) | CP^0=[u]",
		CP1CP0SplitExists:         g.Selector.ConstructsCP1CP0Split,
		NativeWithoutSeal:         false,
		Verdict:                   join(StatusSealedDecompositionVerified, StatusProjectorPropertiesVerified),
	}
}

func auditBMinusLCompatibility(d SealedDecompositionAudit) BMinusLCompatibilityAudit {
	bl := scale3(identity3(), 1.0/3.0)
	commP := commutatorMaxAbs(bl, d.Projector)
	commQ := commutatorMaxAbs(bl, d.ComplementProjector)
	return BMinusLCompatibilityAudit{
		BLRestriction:           "(B-L)|W_spatial=(1/3)I_3",
		BLSpatialEigenvalue:     1.0 / 3.0,
		CommutatorResidual:      math.Max(commP, commQ),
		CommutesWithPU:          commP < 1e-12,
		CommutesWithComplement:  commQ < 1e-12,
		CompatibilityVacuous:    true,
		SuppliesFurtherSelector: false,
		Verdict:                 join(StatusBMinusLCommutatorZero, StatusBMinusLCompatibilityVacuous),
	}
}

func auditCommutant(g gate574.Analysis) CommutantAudit {
	return CommutantAudit{
		SelectorFormula:       "S_sp=lambda_2(I-P_u)+lambda_1P_u, lambda_1 != lambda_2",
		MultiplicityPattern:   "2+1",
		Commutant:             g.Commutant.Commutant,
		DimensionFormula:      g.Commutant.DimensionFormula,
		Dimension:             g.Commutant.Dimension,
		MatchesGate555Formula: g.Commutant.MatchesGate555Formula && g.Commutant.Dimension == 5,
		SealedSupportOnly:     g.Commutant.SealedSupportOnly,
		Verdict:               join(StatusCommutantReconfirmed, StatusGate555SelectorFormulaReused),
	}
}

func auditQuaternionicSocket() QuaternionicSocketComparisonAudit {
	return QuaternionicSocketComparisonAudit{
		ImHSocketAvailableElsewhere:     true,
		HPhiDoubletModuleAvailable:      true,
		PauliHopfMomentQuaternionic:     true,
		FiniteOneFormLinkedStructurally: true,
		SourceCarrier:                   "Im(H) acting structurally on the H_phi scalar/finite-one-form socket",
		TestedTarget:                    "su(u^perp) or End(u^perp) for the sealed spatial CP^1 complement",
		ImHToSuUperpIntertwinerSupplied: false,
		HToEndUperpModuleSupplied:       false,
		CompatibleWithPU:                false,
		WSpatialTransferBlocked:         true,
		Verdict:                         join(StatusQuaternionicSocketExistsElsewhere, StatusNoImHToSealedCP1Intertwiner, StatusNoHToEndUperpModule, StatusGate562ImHNotWSpatialPreserved),
	}
}

func auditFiniteSpectralTripleCarrier() FiniteSpectralTripleCarrierAudit {
	return FiniteSpectralTripleCarrierAudit{
		AFRepresentationStructural:              true,
		FiniteWeakDoubletCarrierExistsElsewhere: true,
		UperpUsedAsFiniteWeakDoubletCarrier:     false,
		DCompatibilityForUperp:                  false,
		JCompatibilityForUperp:                  false,
		GradingCompatibilityForUperp:            false,
		FirstOrderCompatibilityForUperp:         false,
		OrderOneCarrierActionProven:             false,
		Verdict:                                 join(StatusNoFiniteWeakDoubletCarrier, StatusNoDJGradingFirstOrderCompatibility),
	}
}

func auditOneFormHiggsLane() OneFormHiggsLaneCompatibilityAudit {
	return OneFormHiggsLaneCompatibilityAudit{
		FiniteOneFormContainsScalarDoublet: true,
		ImHActsOnHPhiStructurally:          true,
		SealedCP1AppearsInOneFormLane:      false,
		SealedCP1AppearsInHiggsLane:        false,
		PauliRouteSeparateFromWSpatial:     true,
		Gate562BoundaryPreserved:           true,
		Verdict:                            join(StatusNoOneFormHiggsLaneCarrier, StatusGate562ImHNotWSpatialPreserved),
	}
}

func auditPhysicalWeakPlaneFirewall() PhysicalWeakPlaneFirewallAudit {
	return PhysicalWeakPlaneFirewallAudit{
		RepresentativeGauge:            "[u]=[a_3^dagger] representative gauge",
		Complement:                     "span_C{a_1^dagger,a_2^dagger}, projectively CP^1, conventionally U_12",
		ConventionalName:               "U_12",
		CanCallPhysicalWeakPlane:       false,
		FiniteSpectralTripleCompatible: false,
		QuaternionicCompatible:         false,
		DCompatible:                    false,
		JCompatible:                    false,
		GradingCompatible:              false,
		FirstOrderCompatible:           false,
		Verdict:                        join(StatusRepresentativeU12NotPhysicalWeakPlane, StatusWeakPlaneRequiresFSTCompatibility),
	}
}

func auditFlavorElectroweakFirewall() FlavorElectroweakFirewallAudit {
	return FlavorElectroweakFirewallAudit{
		GenerationHierarchyDerived: false,
		YukawaTextureDerived:       false,
		CKMPMNSDerived:             false,
		ObservedFlavorImported:     false,
		PhysicalEWDynamicsDerived:  false,
		PhotonDynamicsDerived:      false,
		WZMassesDerived:            false,
		WeakIsospinDerived:         false,
		Verdict:                    StatusNoFlavorEWObservedData,
	}
}

func auditPreviousGateBoundaries(g574 gate574.Analysis) PreviousGateBoundaryAudit {
	return PreviousGateBoundaryAudit{
		TauEtaTraceShadowOnly:              true,
		Q4ContactOnly:                      true,
		PauliQuaternionicSocketNotWSpatial: true,
		Gate564565BridgeSymbolic:           g574.Boundaries.Gate564565BridgeSymbolic,
		K7TimeRoutesSealed:                 g574.Boundaries.K7TimeRoutesSealed,
		OrientationSealProjectiveOnly:      g574.Final.SealSufficient && !g574.Final.DerivesPhysicalWeakFlavorElectroweakData,
		Verdict:                            join(StatusGate562ImHNotWSpatialPreserved, StatusGate564565BoundaryPreserved, StatusK7TimeBoundaryPreserved),
	}
}

func auditFinal(a Analysis) FinalVerdict {
	physical := a.WeakPlane.CanCallPhysicalWeakPlane || a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.ObservedFlavorImported || a.FlavorEW.PhysicalEWDynamicsDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived || a.FlavorEW.WeakIsospinDerived
	return FinalVerdict{
		SealedCP1SplitExistsAlgebraically: a.Decomposition.CP1CP0SplitExists,
		CommutesWithBMinusL:               a.BMinusL.CommutesWithPU && a.BMinusL.CommutesWithComplement,
		CarriesNativeOrSealedImHAction:    a.Quaternionic.ImHToSuUperpIntertwinerSupplied || a.Quaternionic.HToEndUperpModuleSupplied,
		PartOfFiniteWeakDoubletCarrier:    a.FiniteTriple.UperpUsedAsFiniteWeakDoubletCarrier,
		CanBeCalledPhysicalWeakPlane:      a.WeakPlane.CanCallPhysicalWeakPlane,
		DerivesFlavorOrEWObservedData:     physical,
		AdditionalTheoremRequired:         "To promote the sealed CP^1 complement beyond projective orientation, ASHA would need a native theorem constructing an Im(H)->su(u^perp) or H->End(u^perp) intertwiner compatible with the chosen P_u and proving that u^perp is the finite spectral-triple weak-doublet carrier with D, J, grading, first-order, finite one-form/Higgs-lane, B-L, K7/time, and flavor/electroweak firewalls satisfied.",
		Verdict:                           join(StatusSealedDecompositionVerified, StatusBMinusLCommutatorZero, StatusNoImHToSealedCP1Intertwiner, StatusNoFiniteWeakDoubletCarrier, StatusRepresentativeU12NotPhysicalWeakPlane, StatusNoFlavorEWObservedData, StatusGate575BoundaryPreserved),
	}
}

func validate(a Analysis) error {
	if !a.Inherited.SealSufficient || !a.Inherited.SealMinimal || !a.Inherited.CP1CP0Split || !a.Inherited.CommutantU2U1 || a.Inherited.SealNative || a.Inherited.PhysicalWeakPlaneDerived || a.Inherited.FlavorElectroweakDataDerived || a.Inherited.K7OrProductTimeOpened {
		return fmt.Errorf("Gate574 inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.Decomposition.CP1CP0SplitExists || a.Decomposition.ProjectorRank != 1 || a.Decomposition.ComplementRank != 2 || a.Decomposition.DimCUperp != 2 || a.Decomposition.ProjectorIdempotentError > 1e-12 || a.Decomposition.ComplementIdempotentError > 1e-12 || a.Decomposition.OrthogonalityResidual > 1e-12 || a.Decomposition.NativeWithoutSeal {
		return fmt.Errorf("sealed decomposition failed: %s", FormatSealedDecomposition(a.Decomposition))
	}
	if !a.BMinusL.CommutesWithPU || !a.BMinusL.CommutesWithComplement || a.BMinusL.CommutatorResidual > 1e-12 || !a.BMinusL.CompatibilityVacuous || a.BMinusL.SuppliesFurtherSelector {
		return fmt.Errorf("B-L compatibility failed: %s", FormatBMinusL(a.BMinusL))
	}
	if a.Commutant.Commutant != "u(2)+u(1)" || a.Commutant.Dimension != 5 || !a.Commutant.MatchesGate555Formula || !a.Commutant.SealedSupportOnly {
		return fmt.Errorf("commutant audit failed: %s", FormatCommutant(a.Commutant))
	}
	if !a.Quaternionic.ImHSocketAvailableElsewhere || !a.Quaternionic.HPhiDoubletModuleAvailable || !a.Quaternionic.WSpatialTransferBlocked || a.Quaternionic.ImHToSuUperpIntertwinerSupplied || a.Quaternionic.HToEndUperpModuleSupplied || a.Quaternionic.CompatibleWithPU {
		return fmt.Errorf("quaternionic socket audit failed: %s", FormatQuaternionic(a.Quaternionic))
	}
	if !a.FiniteTriple.AFRepresentationStructural || !a.FiniteTriple.FiniteWeakDoubletCarrierExistsElsewhere || a.FiniteTriple.UperpUsedAsFiniteWeakDoubletCarrier || a.FiniteTriple.DCompatibilityForUperp || a.FiniteTriple.JCompatibilityForUperp || a.FiniteTriple.GradingCompatibilityForUperp || a.FiniteTriple.FirstOrderCompatibilityForUperp || a.FiniteTriple.OrderOneCarrierActionProven {
		return fmt.Errorf("finite spectral-triple carrier audit failed: %s", FormatFiniteTriple(a.FiniteTriple))
	}
	if !a.OneForm.FiniteOneFormContainsScalarDoublet || !a.OneForm.ImHActsOnHPhiStructurally || a.OneForm.SealedCP1AppearsInOneFormLane || a.OneForm.SealedCP1AppearsInHiggsLane || !a.OneForm.PauliRouteSeparateFromWSpatial || !a.OneForm.Gate562BoundaryPreserved {
		return fmt.Errorf("one-form/Higgs lane audit failed: %s", FormatOneForm(a.OneForm))
	}
	if a.WeakPlane.CanCallPhysicalWeakPlane || a.WeakPlane.FiniteSpectralTripleCompatible || a.WeakPlane.QuaternionicCompatible || a.WeakPlane.DCompatible || a.WeakPlane.JCompatible || a.WeakPlane.GradingCompatible || a.WeakPlane.FirstOrderCompatible {
		return fmt.Errorf("physical weak-plane firewall failed: %s", FormatWeakPlane(a.WeakPlane))
	}
	if a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.ObservedFlavorImported || a.FlavorEW.PhysicalEWDynamicsDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived || a.FlavorEW.WeakIsospinDerived {
		return fmt.Errorf("flavor/electroweak firewall failed: %s", FormatFlavorEW(a.FlavorEW))
	}
	if !a.Boundaries.TauEtaTraceShadowOnly || !a.Boundaries.Q4ContactOnly || !a.Boundaries.PauliQuaternionicSocketNotWSpatial || !a.Boundaries.Gate564565BridgeSymbolic || !a.Boundaries.K7TimeRoutesSealed || !a.Boundaries.OrientationSealProjectiveOnly {
		return fmt.Errorf("previous boundary audit failed: %s", FormatBoundaries(a.Boundaries))
	}
	if !a.Final.SealedCP1SplitExistsAlgebraically || !a.Final.CommutesWithBMinusL || a.Final.CarriesNativeOrSealedImHAction || a.Final.PartOfFiniteWeakDoubletCarrier || a.Final.CanBeCalledPhysicalWeakPlane || a.Final.DerivesFlavorOrEWObservedData || a.Final.AdditionalTheoremRequired == "" {
		return fmt.Errorf("final verdict failed: %s", FormatFinal(a.Final))
	}
	return nil
}

func truth(a Analysis) string {
	return fmt.Sprintf("Gate 575 verdict: the sealed CP^1 complement exists algebraically under %s and commutes with B-L because (B-L)|W_spatial=(1/3)I. Current ASHA data supplies no Im(H)->su(u^perp) intertwiner, no H->End(u^perp) module, and no finite spectral-triple D/J/grading/first-order carrier action on u^perp. The representative U_12 plane remains sealed projective geometry only, not physical weak isospin, flavor, electroweak dynamics, K7, time, OS/Hilbert, RG, or observed data.", a.Inherited.SealName)
}

func Statuses() []string {
	return []string{
		StatusGate574Inherited,
		StatusSealedDecompositionVerified,
		StatusProjectorPropertiesVerified,
		StatusBMinusLCommutatorZero,
		StatusBMinusLCompatibilityVacuous,
		StatusCommutantReconfirmed,
		StatusGate555SelectorFormulaReused,
		StatusQuaternionicSocketExistsElsewhere,
		StatusNoImHToSealedCP1Intertwiner,
		StatusNoHToEndUperpModule,
		StatusNoFiniteWeakDoubletCarrier,
		StatusNoDJGradingFirstOrderCompatibility,
		StatusNoOneFormHiggsLaneCarrier,
		StatusGate562ImHNotWSpatialPreserved,
		StatusRepresentativeU12NotPhysicalWeakPlane,
		StatusWeakPlaneRequiresFSTCompatibility,
		StatusNoFlavorEWObservedData,
		StatusGate564565BoundaryPreserved,
		StatusK7TimeBoundaryPreserved,
		StatusGate575BoundaryPreserved,
	}
}

func join(parts ...string) string { return strings.Join(parts, ";") }

func identity3() [][]float64 { return [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} }

func clone3(m [][]float64) [][]float64 {
	out := make([][]float64, 3)
	for i := range out {
		out[i] = make([]float64, 3)
		copy(out[i], m[i])
	}
	return out
}

func scale3(m [][]float64, c float64) [][]float64 {
	out := make([][]float64, 3)
	for i := range out {
		out[i] = make([]float64, 3)
		for j := range out[i] {
			out[i][j] = c * m[i][j]
		}
	}
	return out
}

func subtract3(a, b [][]float64) [][]float64 {
	out := make([][]float64, 3)
	for i := range out {
		out[i] = make([]float64, 3)
		for j := range out[i] {
			out[i][j] = a[i][j] - b[i][j]
		}
	}
	return out
}

func trace3(m [][]float64) float64 { return m[0][0] + m[1][1] + m[2][2] }

func rankFromTrace(m [][]float64) int { return int(math.Round(trace3(m))) }

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

func productMaxAbs(a, b [][]float64) float64 {
	maxAbs := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s := 0.0
			for k := 0; k < 3; k++ {
				s += a[i][k] * b[k][j]
			}
			if math.Abs(s) > maxAbs {
				maxAbs = math.Abs(s)
			}
		}
	}
	return maxAbs
}

func commutatorMaxAbs(a, b [][]float64) float64 {
	maxAbs := 0.0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ab, ba := 0.0, 0.0
			for k := 0; k < 3; k++ {
				ab += a[i][k] * b[k][j]
				ba += b[i][k] * a[k][j]
			}
			d := math.Abs(ab - ba)
			if d > maxAbs {
				maxAbs = d
			}
		}
	}
	return maxAbs
}
