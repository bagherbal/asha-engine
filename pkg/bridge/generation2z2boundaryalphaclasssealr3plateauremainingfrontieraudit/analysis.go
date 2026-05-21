// Package generation2z2boundaryalphaclasssealr3plateauremainingfrontieraudit
// implements Gate 910: Z2 BoundaryAlpha ClassSeal R3 Plateau and Remaining
// Frontier Audit.
//
// Gate 910 is intentionally classificatory. It follows Gate 909's result that
// BoundaryAlpha can be written on the quotient puncture class [p]_{Z2}, not on a
// chosen lambda/bar(lambda) representative. It freezes the current R3 plateau:
// the aggregate trace ledger is structurally present as a Z2-equivariant sealed
// trace ledger, while native R3 remains blocked by the native Z2 BoundaryAlpha
// functor/source and lawful status of the post-orientation stabilizer layer.
package generation2z2boundaryalphaclasssealr3plateauremainingfrontieraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE910-Z2-BOUNDARY-ALPHA-CLASSSEAL-R3-PLATEAU-REMAINING-FRONTIER-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	LinearDenom  = 10
	QuadDenom    = 72

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PunctureClass        = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	BoundaryAlphaFormula = "alpha_B^Z2=[rank([F_1/F_0]_{Z2})/10]s+[rank([F_2/F_0]_{Z2})/72]s^2"
	ReducedB2Response    = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	TraceRowMultiset     = "{(rank 3, weight 1),(rank 3, weight alpha_B(1-alpha_B)),(rank 1, weight 3 alpha_B^2)}"
	FiniteSectorLedger   = "Pi_sector^{F,Z2} with row multiset 3,3,1 and weights 1, alpha_B(1-alpha_B), 3 alpha_B^2"
	OrientedAlgebra      = "A_F^orient=C_R plus C_H plus M_3(C)"
	FullAlgebra          = "A_F=C plus H plus M_3(C)"

	Gate909Classification = "R3_Z2_EQUIVARIANT_TRACE_LEDGER_WITH_BOUNDARY_ALPHA_CLASS_SEAL_NOT_NATIVE"
	Gate909ShortStatus    = "R3_Z2_BOUNDARY_ALPHA_CLASS_SEAL_NOT_NATIVE"
	Classification        = "R3_SEALED_Z2_EQUIVARIANT_TRACE_LEDGER_WITH_BOUNDARY_ALPHA_CLASS_SEAL_NATIVE_PROMOTION_BLOCKED"
	ShortStatus           = "R3_Z2_BOUNDARY_ALPHA_CLASS_SEAL_PLATEAU_NOT_NATIVE"
	FinalTruth            = "R3_SEALED_PLATEAU_CONFIRMED_NATIVE_R3_BLOCKED_BY_Z2_BOUNDARY_ALPHA_FUNCTOR"
	StrategicConclusion   = "R3 trace ledger exists under a Z2 BoundaryAlpha class seal; the remaining native frontier is the native Z2 BoundaryAlpha functor plus lawful status of the post-orientation stabilizer layer."
	NextGate              = "NEXT_PRESSURE_GATE911_NATIVE_R3_FRONTIER_SELECTION_Z2_BOUNDARY_ALPHA_FUNCTOR_VS_FULL_A_F_DESCENT_AUDIT"

	StatusGate909Inherited      = "PASS_GATE909_Z2_BOUNDARY_ALPHA_CLASS_SEAL_INHERITED"
	StatusR3PlateauFrozen       = "PASS_R3_SEALED_PLATEAU_REACHED_AND_FROZEN"
	StatusPhaseSignClosed       = "PASS_PHASE_SIGN_AND_REPRESENTATIVE_ALPHA_LOOP_CLOSED"
	StatusBoundaryAlphaClass    = "PASS_BOUNDARY_ALPHA_CLASS_SEAL_REPRESENTATIVE_INDEPENDENT"
	StatusTraceLedgerComplete   = "PASS_Z2_TRACE_MAGNITUDE_LEDGER_COMPLETE_UNDER_SEAL"
	StatusNativeBlockersReduced = "PASS_NATIVE_R3_FRONTIER_REDUCED_TO_Z2_BOUNDARY_ALPHA_FUNCTOR_AND_FULL_A_F_DESCENT"
	StatusR4FrontierSeparated   = "PASS_GENERATION_FLAVOR_AND_INDIVIDUAL_YUKAWA_BRANCHES_CLASSIFIED_AS_R4_OR_LATER"
	StatusOfficialFreeze        = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict       = "FIREWALL_PRESERVED_GATE910_NOT_NATIVE_R3"

	SupportR3SealedPlateau           = "CONDITIONAL_SUPPORT_R3_SEALED_PLATEAU_REACHED"
	SupportPhaseSignNoLongerBlocks   = "CONDITIONAL_SUPPORT_PHASE_SIGN_NO_LONGER_BLOCKS_R3_TRACE_LEDGER"
	SupportTraceRowsZ2Invariant      = "CONDITIONAL_SUPPORT_TRACE_MAGNITUDE_ROW_MULTISET_IS_Z2_CLASS_INVARIANT"
	SupportBoundaryAlphaClassSeal    = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_CLASS_SEAL_IS_REPRESENTATIVE_INDEPENDENT"
	SupportAlphaRankPairZ2Invariant  = "CONDITIONAL_SUPPORT_ALPHA_RANK_PAIR_3_7_IS_Z2_CLASS_INVARIANT"
	SupportFiniteSectorZ2Ledger      = "CONDITIONAL_SUPPORT_FINITE_SECTOR_PROJECTOR_LEDGER_EXISTS_AS_Z2_ORIENTATION_CLASS"
	SupportPositiveReadoutZ2Class    = "CONDITIONAL_SUPPORT_POSITIVE_TRACE_MAGNITUDE_READOUT_EXISTS_ON_Z2_CLASS"
	SupportZ2LedgerCompleteUnderSeal = "CONDITIONAL_SUPPORT_Z2_EQUIVARIANT_TRACE_MAGNITUDE_LEDGER_IS_COMPLETE_UNDER_SEAL"
	SupportOperatorsReconstructed    = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_C_YUKAWA_C_HIGGS_RECONSTRUCTED_UNDER_SEAL"
	SupportNativeFrontierReduced     = "CONDITIONAL_SUPPORT_NATIVE_R3_FRONTIER_REDUCED_TO_Z2_BOUNDARY_ALPHA_FUNCTOR_AND_FULL_A_F_DESCENT"
	SupportR4Later                   = "CONDITIONAL_SUPPORT_GENERATION_FLAVOR_AND_INDIVIDUAL_YUKAWA_BRANCHES_ARE_R4_OR_LATER"

	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoNativeZ2AirlockFunctor       = "FAILED_ROUTE_NO_NATIVE_Z2_EQUIVARIANT_AIRLOCK_FUNCTOR"
	FailureNoNativeZ2BoundaryAlphaFunctor = "FAILED_ROUTE_NO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	FailureNoNativeDegreeToZ2FlagFunctor  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeZ2CrossLane            = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureReducedB2NotNativeFunctional   = "FAILED_ROUTE_REDUCED_B2_RESPONSE_NOT_NATIVE_BOUNDARY_FUNCTIONAL"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeBoundaryAlphaSource    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ALPHA_SOURCE"
	FailureNoNativeTransportS             = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_Z2_AIRLOCK_CLASS"
	FailureHiggsOrientationClassSealed    = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureAFOrientNotFullAF              = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFullAFToOrientDescent  = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoPhysicalParticleAssign       = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoIndividualYukawaSpectrum     = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawa               = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedPlateau struct {
	Gate909Classification  string
	Gate909ShortStatus     string
	BoundaryAlphaClassSeal bool
	TraceLedgerZ2Class     bool
	NativeR3               bool
	Supports, Failures     []string
}

type R3ReadyUnderSeal struct {
	PunctureClass                string
	BoundaryAlphaFormula         string
	TraceRows                    string
	FiniteSectorLedger           string
	PhaseSignBlocksTraceLedger   bool
	BoundaryAlphaRepresentative  bool
	BoundaryAlphaClassLevel      bool
	TraceRowsZ2Invariant         bool
	FiniteSectorZ2Ledger         bool
	PositiveReadoutOnZ2Class     bool
	OperatorNEffReconstructed    bool
	OperatorCYukawaReconstructed bool
	OperatorCHiggsReconstructed  bool
	Alpha                        float64
	RowWeightUnit                float64
	RowWeightRest                float64
	RowWeightPuncture            float64
	Supports, Failures           []string
}

type NativeBlockers struct {
	NativeZ2BoundaryAlphaFunctorMissing bool
	NativeReducedB2FunctionalMissing    bool
	FullAFDescentMissing                bool
	Z2AirlockFunctor                    string
	BoundaryResponse                    string
	OrientedAlgebra                     string
	FullAlgebra                         string
	CoreBlockerCount                    int
	PhaseSignStillBlocker               bool
	RepresentativeAlphaStillBlocker     bool
	IndividualYukawaStillR3Blocker      bool
	Supports, Failures                  []string
}

type LaterFrontier struct {
	GenerationCarrierR4OrLater    bool
	FlavorOrientationR4OrLater    bool
	IndividualYukawaR4OrLater     bool
	PhysicalAssignmentR4OrLater   bool
	CKMPMNSR4OrLater              bool
	ObservedMassSpectrumR4OrLater bool
	CanEnterR4FromGate910         bool
	Supports, Failures            []string
}

type OfficialFreeze struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	OperatorDiagnosticsOnly           bool
	OfficialLedgersFrozen             bool
	CanUpdateOfficialNEff             bool
	CanUpdateCYukawaCHiggs            bool
	Supports, Failures                []string
}

type PlateauClassification struct {
	Classification                 string
	ShortStatus                    string
	Verdict                        string
	NextGate                       string
	R3ReadyUnderSeal               bool
	NativeR3                       bool
	LoopBackToPhase                bool
	LoopBackToRepAlpha             bool
	NextRailA_Z2BoundaryAlphaFirst bool
	NextRailB_FullAFSecond         bool
	Supports, Failures             []string
}

type Firewalls struct {
	NativeR3                         bool
	NativeZ2AirlockFunctor           bool
	NativeZ2BoundaryAlphaFunctor     bool
	NativeDegreeToZ2FlagFunctor      bool
	NativeZ2CrossLane                bool
	ReducedB2NativeFunctional        bool
	AlphaNative                      bool
	NativeBoundaryAlphaSource        bool
	NativeTransportS                 bool
	HiggsOrientationNative           bool
	FullAFDescent                    bool
	AFOrientEqualsFullAF             bool
	FullAFToOrientDescent            bool
	OfficialLedgerUpdate             bool
	PhysicalParticleAssignment       bool
	GenerationCarrierMap             bool
	FlavorOrientationMap             bool
	IndividualYukawaValues           bool
	IndividualPhysicalYukawaSpectrum bool
	NativeYukawaOperator             bool
	R4NativeYukawa                   bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedPlateau
	Ready          R3ReadyUnderSeal
	Blockers       NativeBlockers
	Later          LaterFrontier
	Freeze         OfficialFreeze
	Plateau        PlateauClassification
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func LinearCoefficient() float64              { return float64(RankF1OverF0) / float64(LinearDenom) }
func QuadraticCoefficient() float64           { return float64(RankF2OverF0) / float64(QuadDenom) }
func LinearContribution(s float64) float64    { return LinearCoefficient() * s }
func QuadraticContribution(s float64) float64 { return QuadraticCoefficient() * s * s }
func BoundaryAlphaZ2(s float64) float64       { return LinearContribution(s) + QuadraticContribution(s) }
func RowWeightRest(alpha float64) float64     { return alpha * (1.0 - alpha) }
func RowWeightPuncture(alpha float64) float64 { return 3.0 * alpha * alpha }

func BuildDefault() (Audit, error) {
	inherited := buildInheritedPlateau()
	if !inherited.BoundaryAlphaClassSeal || !inherited.TraceLedgerZ2Class || inherited.NativeR3 {
		return Audit{}, fmt.Errorf("inherited plateau leak: %s", FormatInherited(inherited))
	}

	ready := buildR3ReadyUnderSeal()
	if ready.PhaseSignBlocksTraceLedger || ready.BoundaryAlphaRepresentative || !ready.BoundaryAlphaClassLevel || !ready.TraceRowsZ2Invariant || !ready.FiniteSectorZ2Ledger || !ready.PositiveReadoutOnZ2Class || !ready.OperatorNEffReconstructed || !ready.OperatorCYukawaReconstructed || !ready.OperatorCHiggsReconstructed || !near(ready.Alpha, AlphaB) {
		return Audit{}, fmt.Errorf("R3-ready sealed structure leak: %s", FormatReady(ready))
	}

	blockers := buildNativeBlockers()
	if !blockers.NativeZ2BoundaryAlphaFunctorMissing || !blockers.NativeReducedB2FunctionalMissing || !blockers.FullAFDescentMissing || blockers.CoreBlockerCount != 3 || blockers.PhaseSignStillBlocker || blockers.RepresentativeAlphaStillBlocker || blockers.IndividualYukawaStillR3Blocker {
		return Audit{}, fmt.Errorf("native blocker classification leak: %s", FormatBlockers(blockers))
	}

	later := buildLaterFrontier()
	if !later.GenerationCarrierR4OrLater || !later.FlavorOrientationR4OrLater || !later.IndividualYukawaR4OrLater || !later.PhysicalAssignmentR4OrLater || !later.CKMPMNSR4OrLater || !later.ObservedMassSpectrumR4OrLater || later.CanEnterR4FromGate910 {
		return Audit{}, fmt.Errorf("later frontier leak: %s", FormatLater(later))
	}

	freeze := buildOfficialFreeze()
	if !freeze.OperatorDiagnosticsOnly || !freeze.OfficialLedgersFrozen || freeze.CanUpdateOfficialNEff || freeze.CanUpdateCYukawaCHiggs || near(freeze.OperatorNEff, freeze.OfficialNEff) || near(freeze.OperatorCYukawa, freeze.OfficialCYukawa) || near(freeze.OperatorCHiggs, freeze.OfficialCHiggs) {
		return Audit{}, fmt.Errorf("official freeze leak: %s", FormatFreeze(freeze))
	}

	plateau := buildPlateauClassification()
	if !plateau.R3ReadyUnderSeal || plateau.NativeR3 || plateau.LoopBackToPhase || plateau.LoopBackToRepAlpha || !plateau.NextRailA_Z2BoundaryAlphaFirst || !plateau.NextRailB_FullAFSecond {
		return Audit{}, fmt.Errorf("plateau classification leak: %s", FormatPlateau(plateau))
	}

	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited:      inherited,
		Ready:          ready,
		Blockers:       blockers,
		Later:          later,
		Freeze:         freeze,
		Plateau:        plateau,
		Firewalls:      firewalls,
		Truth:          FinalTruth,
		Final:          "Gate 910 freezes the current rail: R3 is structurally present only as a Z2-equivariant sealed aggregate trace ledger with BoundaryAlpha defined on the class [p]_{Z2}. The phase-sign wound and representative-alpha wound are closed at class/seal level. Native R3 is not certified: the remaining native pressure is the Z2 BoundaryAlpha functor/source and lawful post-orientation stabilizer status, while generation, flavor, individual Yukawa values, CKM/PMNS, physical particle assignment, and observed spectra are R4-or-later territory. Official N_eff, C_Yukawa, and C_Higgs ledgers remain frozen.",
	}, nil
}

func buildInheritedPlateau() InheritedPlateau {
	return InheritedPlateau{
		Gate909Classification:  Gate909Classification,
		Gate909ShortStatus:     Gate909ShortStatus,
		BoundaryAlphaClassSeal: true,
		TraceLedgerZ2Class:     true,
		NativeR3:               false,
		Supports:               []string{StatusGate909Inherited, SupportBoundaryAlphaClassSeal, SupportPhaseSignNoLongerBlocks},
		Failures:               []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeZ2BoundaryAlphaFunctor},
	}
}

func buildR3ReadyUnderSeal() R3ReadyUnderSeal {
	alpha := BoundaryAlphaZ2(SBoundary)
	return R3ReadyUnderSeal{
		PunctureClass:                PunctureClass,
		BoundaryAlphaFormula:         BoundaryAlphaFormula,
		TraceRows:                    TraceRowMultiset,
		FiniteSectorLedger:           FiniteSectorLedger,
		PhaseSignBlocksTraceLedger:   false,
		BoundaryAlphaRepresentative:  false,
		BoundaryAlphaClassLevel:      true,
		TraceRowsZ2Invariant:         true,
		FiniteSectorZ2Ledger:         true,
		PositiveReadoutOnZ2Class:     true,
		OperatorNEffReconstructed:    true,
		OperatorCYukawaReconstructed: true,
		OperatorCHiggsReconstructed:  true,
		Alpha:                        alpha,
		RowWeightUnit:                1,
		RowWeightRest:                RowWeightRest(alpha),
		RowWeightPuncture:            RowWeightPuncture(alpha),
		Supports: []string{
			StatusR3PlateauFrozen,
			StatusPhaseSignClosed,
			StatusBoundaryAlphaClass,
			StatusTraceLedgerComplete,
			SupportR3SealedPlateau,
			SupportPhaseSignNoLongerBlocks,
			SupportTraceRowsZ2Invariant,
			SupportBoundaryAlphaClassSeal,
			SupportAlphaRankPairZ2Invariant,
			SupportFiniteSectorZ2Ledger,
			SupportPositiveReadoutZ2Class,
			SupportZ2LedgerCompleteUnderSeal,
			SupportOperatorsReconstructed,
		},
		Failures: []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
}

func buildNativeBlockers() NativeBlockers {
	return NativeBlockers{
		NativeZ2BoundaryAlphaFunctorMissing: true,
		NativeReducedB2FunctionalMissing:    true,
		FullAFDescentMissing:                true,
		Z2AirlockFunctor:                    "Z2EquivariantNeutralPunctureAirlockFunctor",
		BoundaryResponse:                    ReducedB2Response,
		OrientedAlgebra:                     OrientedAlgebra,
		FullAlgebra:                         FullAlgebra,
		CoreBlockerCount:                    3,
		PhaseSignStillBlocker:               false,
		RepresentativeAlphaStillBlocker:     false,
		IndividualYukawaStillR3Blocker:      false,
		Supports:                            []string{StatusNativeBlockersReduced, SupportNativeFrontierReduced},
		Failures: []string{
			FailureNoNativeZ2AirlockFunctor,
			FailureNoNativeZ2BoundaryAlphaFunctor,
			FailureNoNativeDegreeToZ2FlagFunctor,
			FailureNoNativeZ2CrossLane,
			FailureReducedB2NotNativeFunctional,
			FailureAlphaStillSealed,
			FailureNoNativeBoundaryAlphaSource,
			FailureNoNativeTransportS,
			FailureHiggsOrientationClassSealed,
			FailureFullAFDescentStillBlocked,
			FailureAFOrientNotFullAF,
			FailureNoNativeFullAFToOrientDescent,
			FailureNotNativeR3,
		},
	}
}

func buildLaterFrontier() LaterFrontier {
	return LaterFrontier{
		GenerationCarrierR4OrLater:    true,
		FlavorOrientationR4OrLater:    true,
		IndividualYukawaR4OrLater:     true,
		PhysicalAssignmentR4OrLater:   true,
		CKMPMNSR4OrLater:              true,
		ObservedMassSpectrumR4OrLater: true,
		CanEnterR4FromGate910:         false,
		Supports:                      []string{StatusR4FrontierSeparated, SupportR4Later},
		Failures: []string{
			FailureNoPhysicalParticleAssign,
			FailureNoGenerationCarrierMap,
			FailureNoFlavorOrientationMap,
			FailureNoIndividualYukawaValues,
			FailureNoIndividualYukawaSpectrum,
			FailureNoNativeYukawaOperator,
			FailureNoR4NativeYukawa,
		},
	}
}

func buildOfficialFreeze() OfficialFreeze {
	return OfficialFreeze{
		Alpha:                   AlphaB,
		OperatorNEff:            OperatorNEffDiagnostic,
		OfficialNEff:            OfficialNEffFrozen,
		OperatorCYukawa:         OperatorCYukawaDiagnostic,
		OfficialCYukawa:         OfficialCYukawaFrozen,
		OperatorCHiggs:          OperatorCHiggsDiagnostic,
		OfficialCHiggs:          OfficialCHiggsFrozen,
		OperatorDiagnosticsOnly: true,
		OfficialLedgersFrozen:   true,
		CanUpdateOfficialNEff:   false,
		CanUpdateCYukawaCHiggs:  false,
		Supports:                []string{StatusOfficialFreeze, SupportOperatorsReconstructed},
		Failures:                []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
}

func buildPlateauClassification() PlateauClassification {
	return PlateauClassification{
		Classification:                 Classification,
		ShortStatus:                    ShortStatus,
		Verdict:                        FinalTruth,
		NextGate:                       NextGate,
		R3ReadyUnderSeal:               true,
		NativeR3:                       false,
		LoopBackToPhase:                false,
		LoopBackToRepAlpha:             false,
		NextRailA_Z2BoundaryAlphaFirst: true,
		NextRailB_FullAFSecond:         true,
		Supports: []string{
			SupportR3SealedPlateau,
			SupportNativeFrontierReduced,
			SupportR4Later,
		},
		Failures: []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NativeR3:                         false,
		NativeZ2AirlockFunctor:           false,
		NativeZ2BoundaryAlphaFunctor:     false,
		NativeDegreeToZ2FlagFunctor:      false,
		NativeZ2CrossLane:                false,
		ReducedB2NativeFunctional:        false,
		AlphaNative:                      false,
		NativeBoundaryAlphaSource:        false,
		NativeTransportS:                 false,
		HiggsOrientationNative:           false,
		FullAFDescent:                    false,
		AFOrientEqualsFullAF:             false,
		FullAFToOrientDescent:            false,
		OfficialLedgerUpdate:             false,
		PhysicalParticleAssignment:       false,
		GenerationCarrierMap:             false,
		FlavorOrientationMap:             false,
		IndividualYukawaValues:           false,
		IndividualPhysicalYukawaSpectrum: false,
		NativeYukawaOperator:             false,
		R4NativeYukawa:                   false,
	}
}

func Statuses() []string {
	return []string{
		StatusGate909Inherited,
		StatusR3PlateauFrozen,
		StatusPhaseSignClosed,
		StatusBoundaryAlphaClass,
		StatusTraceLedgerComplete,
		StatusNativeBlockersReduced,
		StatusR4FrontierSeparated,
		StatusOfficialFreeze,
		StatusFirewallVerdict,
		NextGate,
		SupportR3SealedPlateau,
		SupportPhaseSignNoLongerBlocks,
		SupportTraceRowsZ2Invariant,
		SupportBoundaryAlphaClassSeal,
		SupportAlphaRankPairZ2Invariant,
		SupportFiniteSectorZ2Ledger,
		SupportPositiveReadoutZ2Class,
		SupportZ2LedgerCompleteUnderSeal,
		SupportOperatorsReconstructed,
		SupportNativeFrontierReduced,
		SupportR4Later,
		FailureNotNativeR3,
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureReducedB2NotNativeFunctional,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryAlphaSource,
		FailureNoNativeTransportS,
		FailureHiggsOrientationClassSealed,
		FailureFullAFDescentStillBlocked,
		FailureNoOfficialNEffUpdate,
		FailureNoPhysicalParticleAssign,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNotNativeR3,
		FailureNoNativeZ2AirlockFunctor,
		FailureNoNativeZ2BoundaryAlphaFunctor,
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeZ2CrossLane,
		FailureReducedB2NotNativeFunctional,
		FailureAlphaStillSealed,
		FailureNoNativeBoundaryAlphaSource,
		FailureNoNativeTransportS,
		FailureHiggsOrientationClassSealed,
		FailureFullAFDescentStillBlocked,
		FailureAFOrientNotFullAF,
		FailureNoNativeFullAFToOrientDescent,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoPhysicalParticleAssign,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoIndividualYukawaSpectrum,
		FailureNoNativeYukawaOperator,
		FailureNoR4NativeYukawa,
	}
}

func FormatInherited(a InheritedPlateau) string {
	return fmt.Sprintf("gate909_classification=%s gate909_short=%s boundary_alpha_class_seal=%t trace_ledger_Z2=%t native_R3=%t supports=%s failures=%s", a.Gate909Classification, a.Gate909ShortStatus, a.BoundaryAlphaClassSeal, a.TraceLedgerZ2Class, a.NativeR3, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatReady(a R3ReadyUnderSeal) string {
	return fmt.Sprintf("class=%s formula=%s trace_rows=%s finite_sector=%s phase_blocks=%t representative_alpha=%t class_alpha=%t rows_Z2=%t projector_Z2=%t positive_readout_Z2=%t N_eff=%t C_Y=%t C_H=%t alpha=%.17g weights=[%.16g %.16g %.16g] supports=%s failures=%s", a.PunctureClass, a.BoundaryAlphaFormula, a.TraceRows, a.FiniteSectorLedger, a.PhaseSignBlocksTraceLedger, a.BoundaryAlphaRepresentative, a.BoundaryAlphaClassLevel, a.TraceRowsZ2Invariant, a.FiniteSectorZ2Ledger, a.PositiveReadoutOnZ2Class, a.OperatorNEffReconstructed, a.OperatorCYukawaReconstructed, a.OperatorCHiggsReconstructed, a.Alpha, a.RowWeightUnit, a.RowWeightRest, a.RowWeightPuncture, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatBlockers(a NativeBlockers) string {
	return fmt.Sprintf("native_Z2_boundary_alpha_missing=%t native_reduced_B2_missing=%t full_AF_descent_missing=%t blocker_count=%d phase_still_blocker=%t rep_alpha_still_blocker=%t individual_yukawa_R3_blocker=%t functor=%s response=%s oriented=%s full=%s supports=%s failures=%s", a.NativeZ2BoundaryAlphaFunctorMissing, a.NativeReducedB2FunctionalMissing, a.FullAFDescentMissing, a.CoreBlockerCount, a.PhaseSignStillBlocker, a.RepresentativeAlphaStillBlocker, a.IndividualYukawaStillR3Blocker, a.Z2AirlockFunctor, a.BoundaryResponse, a.OrientedAlgebra, a.FullAlgebra, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatLater(a LaterFrontier) string {
	return fmt.Sprintf("generation_R4=%t flavor_R4=%t individual_yukawa_R4=%t physical_assignment_R4=%t CKM_PMNS_R4=%t observed_spectrum_R4=%t can_enter_R4_now=%t supports=%s failures=%s", a.GenerationCarrierR4OrLater, a.FlavorOrientationR4OrLater, a.IndividualYukawaR4OrLater, a.PhysicalAssignmentR4OrLater, a.CKMPMNSR4OrLater, a.ObservedMassSpectrumR4OrLater, a.CanEnterR4FromGate910, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFreeze(a OfficialFreeze) string {
	return fmt.Sprintf("alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g operator_CY=%.16g official_CY=%.16g operator_CH=%.16g official_CH=%.16g diagnostics_only=%t official_frozen=%t can_update_N_eff=%t can_update_CY_CH=%t supports=%s failures=%s", a.Alpha, a.OperatorNEff, a.OfficialNEff, a.OperatorCYukawa, a.OfficialCYukawa, a.OperatorCHiggs, a.OfficialCHiggs, a.OperatorDiagnosticsOnly, a.OfficialLedgersFrozen, a.CanUpdateOfficialNEff, a.CanUpdateCYukawaCHiggs, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatPlateau(a PlateauClassification) string {
	return fmt.Sprintf("classification=%s short=%s verdict=%s next=%s ready_under_seal=%t native_R3=%t loop_phase=%t loop_rep_alpha=%t next_A_Z2_alpha_first=%t next_B_full_AF_second=%t supports=%s failures=%s", a.Classification, a.ShortStatus, a.Verdict, a.NextGate, a.R3ReadyUnderSeal, a.NativeR3, a.LoopBackToPhase, a.LoopBackToRepAlpha, a.NextRailA_Z2BoundaryAlphaFirst, a.NextRailB_FullAFSecond, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("native_R3=%t native_Z2_airlock=%t native_Z2_boundary_alpha=%t native_degree_Z2_flag=%t native_Z2_cross_lane=%t reduced_B2_native=%t alpha_native=%t native_boundary_alpha_source=%t native_s_transport=%t higgs_native=%t full_AF_descent=%t AF_orient_equals_full_AF=%t full_AF_to_orient=%t official_update=%t physical_assignment=%t generation=%t flavor=%t individual_yukawa=%t physical_yukawa_spectrum=%t native_yukawa_operator=%t R4_native_yukawa=%t", f.NativeR3, f.NativeZ2AirlockFunctor, f.NativeZ2BoundaryAlphaFunctor, f.NativeDegreeToZ2FlagFunctor, f.NativeZ2CrossLane, f.ReducedB2NativeFunctional, f.AlphaNative, f.NativeBoundaryAlphaSource, f.NativeTransportS, f.HiggsOrientationNative, f.FullAFDescent, f.AFOrientEqualsFullAF, f.FullAFToOrientDescent, f.OfficialLedgerUpdate, f.PhysicalParticleAssignment, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.IndividualPhysicalYukawaSpectrum, f.NativeYukawaOperator, f.R4NativeYukawa)
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeR3 && !f.NativeZ2AirlockFunctor && !f.NativeZ2BoundaryAlphaFunctor && !f.NativeDegreeToZ2FlagFunctor && !f.NativeZ2CrossLane && !f.ReducedB2NativeFunctional && !f.AlphaNative && !f.NativeBoundaryAlphaSource && !f.NativeTransportS && !f.HiggsOrientationNative && !f.FullAFDescent && !f.AFOrientEqualsFullAF && !f.FullAFToOrientDescent && !f.OfficialLedgerUpdate && !f.PhysicalParticleAssignment && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.IndividualPhysicalYukawaSpectrum && !f.NativeYukawaOperator && !f.R4NativeYukawa
}

func containsAll(haystack []string, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
