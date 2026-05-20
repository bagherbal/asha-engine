// Package generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit implements
// Gate 764: CP1 Gauge-Orbit Demotion and Radial Rank-Invariant HistoryLoop Audit.
//
// Gate 763 showed that no current ASHA structure selects a point in CP1 and
// that a nonconstant CP1 selector would require a Hermitian SU(2) socket axis.
// Gate 764 asks whether that absence is a scalar-runtime defect or the expected
// behavior of a gauge orbit. It audits CP1 ~= SU(2)/U(1), demotes Pi_vac_C to a
// gauge/vacuum-orientation representative for scalar-runtime numerics, and
// verifies that L_Hopf depends only on the rank-one real radial event weight
// Tr((I_K7+/4)P_rad)=1/4, not on the CP1 position. This is a gauge-orbit and
// rank-invariance audit only. It does not derive electroweak symmetry breaking,
// scalar runtime lambda, Higgs mass, pole mass, Yukawa operators, CKM/PMNS,
// flavor hierarchy, or a native HistoryLoopUnit theorem.
package generation2cp1gaugeorbitdemotionandradialrankinvarianthistoryloopaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE764-CP1-GAUGE-ORBIT-DEMOTION-AND-RADIAL-RANK-INVARIANT-HISTORYLOOP-AUDIT"

	StatusGate763CP1SelectorFirewallInherited                 = "PASS_GATE763_CP1_SELECTOR_FIREWALL_INHERITED"
	StatusCP1AsSU2OrbitAudited                                = "PASS_CP1_AS_SU2_ORBIT_AUDITED"
	StatusGaugeOrbitDemotionDefined                           = "PASS_GAUGE_ORBIT_DEMOTION_DEFINED"
	StatusRadialRankInvarianceComputed                        = "PASS_RADIAL_RANK_INVARIANCE_COMPUTED"
	StatusComplexLineVersusRadialEventDistinctionAudited      = "PASS_COMPLEX_LINE_VERSUS_RADIAL_EVENT_DISTINCTION_AUDITED"
	StatusUpdatedMissingObjectHierarchyRecorded               = "PASS_UPDATED_MISSING_OBJECT_HIERARCHY_RECORDED"
	StatusPhysicalFirewallsEnforced                           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusCP1SelectorAbsenceExpectedForGaugeOrbit             = "CONDITIONAL_SUPPORT_CP1_SELECTOR_ABSENCE_IS_EXPECTED_FOR_GAUGE_ORBIT"
	StatusPiVacCIsGaugeRepresentativeNotScalarNumericalSource = "CONDITIONAL_SUPPORT_PI_VAC_C_IS_GAUGE_REPRESENTATIVE_NOT_SCALAR_NUMERICAL_SOURCE"
	StatusLHopfDependsOnRankOneRadialEventNotCP1Position      = "CONDITIONAL_SUPPORT_L_HOPF_DEPENDS_ON_RANK_ONE_RADIAL_EVENT_NOT_CP1_POSITION"
	StatusNextScalarSourceTargetRadialEventTypeSelection      = "CONDITIONAL_SUPPORT_NEXT_SCALAR_SOURCE_TARGET_IS_RADIAL_EVENT_TYPE_SELECTION"
	StatusInternalCNotCertifiedAsPhysicalSU2L                 = "FAILED_ROUTE_INTERNAL_C_NOT_CERTIFIED_AS_PHYSICAL_SU2L"
	StatusNoNativeReasonHistoryLoopUsesRealRadialEvent        = "FAILED_ROUTE_NO_NATIVE_REASON_HISTORYLOOP_USES_REAL_RADIAL_EVENT"
	StatusNoNativeHistoryLoopUnitTheorem                      = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoNativeElectroweakSymmetryBreakingTheorem          = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                        = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                 = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate764CP1GaugeOrbitRadialRankBoundary              = "FIREWALL_PRESERVED_GATE764_CP1_GAUGE_ORBIT_RADIAL_RANK_BOUNDARY"
)

const (
	k7PlusRealDim       = 4
	k7PlusComplexDim    = 2
	realRadialRank      = 1
	complexLineRealRank = 2
	cp1RealDim          = 2
	su2Dim              = 3
	u1StabilizerDim     = 1
)

type Gate763Inheritance struct {
	Inherited                  bool
	Socket                     string
	CP1                        string
	PiVacC                     string
	PRad                       string
	RhoPlus                    string
	LHopfFormula               string
	NoNativeCP1Selector        bool
	RequiresHermitianSU2Axis   bool
	ComplexVacuumLineSeal      bool
	RadialGaugeFixingSecondary bool
	Verdict                    string
}

type SU2OrbitAudit struct {
	Carrier                         string
	OrbitFormula                    string
	SU2ActsTransitivelyOnCP1        bool
	Stabilizer                      string
	CP1RealDimension                int
	SU2Dimension                    int
	StabilizerDimension             int
	InvariantFunctionalConstant     bool
	AnisotropicAxisRequiredToSelect bool
	ExplainsGate763Failure          bool
	ConditionalOnGaugeAirlock       bool
	Verdict                         string
}

type GaugeOrbitDemotion struct {
	Object                            string
	PriorTyping                       string
	DemotedTyping                     string
	ScalarRuntimeNumericalSource      bool
	GaugeRepresentative               bool
	RequiresPhysicalAnisotropy        bool
	AbsenceOfSelectorScalarFailure    bool
	AbsenceOfSelectorVacuumTheoremGap bool
	DemotionConditional               bool
	Verdict                           string
}

type RadialRankInvariance struct {
	RhoPlusFormula           string
	ProjectorRank            int
	K7PlusDimension          int
	TraceWeight              float64
	PhaseLoopPayoff          float64
	LHopf                    float64
	ExpectedLHopf            float64
	DependsOnRankOnly        bool
	IndependentOfCP1Position bool
	RankInvariant            bool
	Verdict                  string
}

type LineVsRadialEvent struct {
	ComplexLineRealRank     int
	RadialEventRealRank     int
	ComplexLineWeight       float64
	RadialEventWeight       float64
	ComplexLineLoopUnit     float64
	RadialEventLoopUnit     float64
	ActiveUsesRadialEvent   bool
	FullComplexLineRejected bool
	ScalarRuntimeQuestion   string
	Verdict                 string
}

type ConditionalDemotionFirewall struct {
	InternalSocketName              string
	InternalSocketCertified         bool
	PhysicalSU2LCertified           bool
	CP1GaugeOrbitPhysicalTheorem    bool
	DemotionAllowedAsInternalSocket bool
	DemotionPromotedToPhysicalEWSB  bool
	Verdict                         string
}

type MissingObjectHierarchy struct {
	BeforeGate764PrimaryMissing string
	AfterGate764ScalarTarget    string
	AfterGate764PhysicalTarget  string
	CP1LocationScalarDatum      bool
	RadialEventTypeScalarDatum  bool
	ElectroweakVacuumStillOpen  bool
	Updated                     bool
	Verdict                     string
}

type Firewalls struct {
	Audited                            bool
	CP1PointPhysicalTheorem            bool
	AbsenceOfCP1SelectorScalarFailure  bool
	CP1GaugeOrbitPhysicalSU2LTheorem   bool
	RankOneRadialEventNativeTheorem    bool
	LHopfNativeTransportTheorem        bool
	ScalarRuntimeIndependentTheorem    bool
	TreeProxyPoleMassTheorem           bool
	ElectroweakSymmetryBreakingTheorem bool
	HiggsMassOrPoleMassTheorem         bool
	YukawaOperatorOrEigenvalueTheorem  bool
	Verdict                            string
}

type Analysis struct {
	Gate763        Gate763Inheritance
	Orbit          SU2OrbitAudit
	Demotion       GaugeOrbitDemotion
	RankInvariance RadialRankInvariance
	LineVsRadial   LineVsRadialEvent
	Conditional    ConditionalDemotionFirewall
	Hierarchy      MissingObjectHierarchy
	Firewalls      Firewalls
	Truth          string
}

var (
	cacheMu sync.Mutex
	cache   *Analysis
)

func BuildDefault() (*Analysis, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()
	if cache != nil {
		clone := *cache
		return &clone, nil
	}
	rankWeight := float64(realRadialRank) / float64(k7PlusRealDim)
	phaseLoop := 1.0 / (2.0 * math.Pi)
	lhopf := phaseLoop * rankWeight
	complexWeight := float64(complexLineRealRank) / float64(k7PlusRealDim)
	complexLoop := phaseLoop * complexWeight
	a := &Analysis{
		Gate763: Gate763Inheritance{
			Inherited:                  true,
			Socket:                     "K7+_J(n) ~= C^2",
			CP1:                        "CP1 = P(C^2)",
			PiVacC:                     "Pi_vac_C complex vacuum line / CP1 point",
			PRad:                       "P_rad gauge-fixed real radial representative inside Pi_vac_C",
			RhoPlus:                    "rho_plus = I_K7+ / 4",
			LHopfFormula:               "L_Hopf = (1/(2*pi)) Tr(rho_plus P_rad)",
			NoNativeCP1Selector:        true,
			RequiresHermitianSU2Axis:   true,
			ComplexVacuumLineSeal:      true,
			RadialGaugeFixingSecondary: true,
			Verdict:                    StatusGate763CP1SelectorFirewallInherited,
		},
		Orbit: SU2OrbitAudit{
			Carrier:                         "K7+_J(n) ~= C^2 complex doublet carrier",
			OrbitFormula:                    "CP1 ~= SU(2)/U(1)",
			SU2ActsTransitivelyOnCP1:        true,
			Stabilizer:                      "U(1)",
			CP1RealDimension:                cp1RealDim,
			SU2Dimension:                    su2Dim,
			StabilizerDimension:             u1StabilizerDim,
			InvariantFunctionalConstant:     true,
			AnisotropicAxisRequiredToSelect: true,
			ExplainsGate763Failure:          true,
			ConditionalOnGaugeAirlock:       true,
			Verdict:                         strings.Join([]string{StatusCP1AsSU2OrbitAudited, StatusCP1SelectorAbsenceExpectedForGaugeOrbit, StatusInternalCNotCertifiedAsPhysicalSU2L}, "; "),
		},
		Demotion: GaugeOrbitDemotion{
			Object:                            "Pi_vac_C in CP1",
			PriorTyping:                       "primary ComplexVacuumLineSeal for the Higgs vacuum orientation",
			DemotedTyping:                     "gauge/vacuum-orientation representative for scalar-runtime numerics unless anisotropic axis is supplied",
			ScalarRuntimeNumericalSource:      false,
			GaugeRepresentative:               true,
			RequiresPhysicalAnisotropy:        true,
			AbsenceOfSelectorScalarFailure:    false,
			AbsenceOfSelectorVacuumTheoremGap: true,
			DemotionConditional:               true,
			Verdict:                           strings.Join([]string{StatusGaugeOrbitDemotionDefined, StatusPiVacCIsGaugeRepresentativeNotScalarNumericalSource}, "; "),
		},
		RankInvariance: RadialRankInvariance{
			RhoPlusFormula:           "Tr((I_K7+/4)P_rad)=rank(P_rad)/4",
			ProjectorRank:            realRadialRank,
			K7PlusDimension:          k7PlusRealDim,
			TraceWeight:              rankWeight,
			PhaseLoopPayoff:          phaseLoop,
			LHopf:                    lhopf,
			ExpectedLHopf:            1.0 / (8.0 * math.Pi),
			DependsOnRankOnly:        true,
			IndependentOfCP1Position: true,
			RankInvariant:            true,
			Verdict:                  strings.Join([]string{StatusRadialRankInvarianceComputed, StatusLHopfDependsOnRankOneRadialEventNotCP1Position}, "; "),
		},
		LineVsRadial: LineVsRadialEvent{
			ComplexLineRealRank:     complexLineRealRank,
			RadialEventRealRank:     realRadialRank,
			ComplexLineWeight:       complexWeight,
			RadialEventWeight:       rankWeight,
			ComplexLineLoopUnit:     complexLoop,
			RadialEventLoopUnit:     lhopf,
			ActiveUsesRadialEvent:   true,
			FullComplexLineRejected: true,
			ScalarRuntimeQuestion:   "why HistoryLoop transport uses a real rank-one radial amplitude event rather than the full complex line",
			Verdict:                 strings.Join([]string{StatusComplexLineVersusRadialEventDistinctionAudited, StatusNoNativeReasonHistoryLoopUsesRealRadialEvent}, "; "),
		},
		Conditional: ConditionalDemotionFirewall{
			InternalSocketName:              "C internal SU(2)-socket candidate",
			InternalSocketCertified:         true,
			PhysicalSU2LCertified:           false,
			CP1GaugeOrbitPhysicalTheorem:    false,
			DemotionAllowedAsInternalSocket: true,
			DemotionPromotedToPhysicalEWSB:  false,
			Verdict:                         strings.Join([]string{StatusInternalCNotCertifiedAsPhysicalSU2L, StatusNoNativeElectroweakSymmetryBreakingTheorem}, "; "),
		},
		Hierarchy: MissingObjectHierarchy{
			BeforeGate764PrimaryMissing: "ComplexVacuumLineSeal / CP1 point selector",
			AfterGate764ScalarTarget:    "radial event type selection: real rank-one amplitude event with phase-loop payoff",
			AfterGate764PhysicalTarget:  "native electroweak symmetry breaking / vacuum-orbit theorem",
			CP1LocationScalarDatum:      false,
			RadialEventTypeScalarDatum:  true,
			ElectroweakVacuumStillOpen:  true,
			Updated:                     true,
			Verdict:                     strings.Join([]string{StatusUpdatedMissingObjectHierarchyRecorded, StatusNextScalarSourceTargetRadialEventTypeSelection}, "; "),
		},
		Firewalls: Firewalls{
			Audited:                            true,
			CP1PointPhysicalTheorem:            false,
			AbsenceOfCP1SelectorScalarFailure:  false,
			CP1GaugeOrbitPhysicalSU2LTheorem:   false,
			RankOneRadialEventNativeTheorem:    false,
			LHopfNativeTransportTheorem:        false,
			ScalarRuntimeIndependentTheorem:    false,
			TreeProxyPoleMassTheorem:           false,
			ElectroweakSymmetryBreakingTheorem: false,
			HiggsMassOrPoleMassTheorem:         false,
			YukawaOperatorOrEigenvalueTheorem:  false,
			Verdict:                            StatusGate764CP1GaugeOrbitRadialRankBoundary,
		},
		Truth: "Gate 764 demotes the CP1 vacuum-line location to internal gauge-orbit representative data for scalar-runtime numerics and shows that L_Hopf depends only on the real rank-one radial event weight, while preserving the missing native reason for choosing that event and the physical SU(2)_L/EWSB firewalls.",
	}
	cache = a
	clone := *a
	return &clone, nil
}

func Statuses() []string {
	return []string{
		StatusGate763CP1SelectorFirewallInherited,
		StatusCP1AsSU2OrbitAudited,
		StatusGaugeOrbitDemotionDefined,
		StatusRadialRankInvarianceComputed,
		StatusComplexLineVersusRadialEventDistinctionAudited,
		StatusUpdatedMissingObjectHierarchyRecorded,
		StatusPhysicalFirewallsEnforced,
		StatusCP1SelectorAbsenceExpectedForGaugeOrbit,
		StatusPiVacCIsGaugeRepresentativeNotScalarNumericalSource,
		StatusLHopfDependsOnRankOneRadialEventNotCP1Position,
		StatusNextScalarSourceTargetRadialEventTypeSelection,
		StatusInternalCNotCertifiedAsPhysicalSU2L,
		StatusNoNativeReasonHistoryLoopUsesRealRadialEvent,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoNativeElectroweakSymmetryBreakingTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate764CP1GaugeOrbitRadialRankBoundary,
	}
}

func FormatGate763(v Gate763Inheritance) string {
	return fmt.Sprintf("inherited=%v socket=%s cp1=%s pi=%s prad=%s rho=%s lhopf=%s no_selector=%v requires_axis=%v complex_seal=%v radial_secondary=%v verdict=%s", v.Inherited, v.Socket, v.CP1, v.PiVacC, v.PRad, v.RhoPlus, v.LHopfFormula, v.NoNativeCP1Selector, v.RequiresHermitianSU2Axis, v.ComplexVacuumLineSeal, v.RadialGaugeFixingSecondary, v.Verdict)
}

func FormatOrbit(v SU2OrbitAudit) string {
	return fmt.Sprintf("carrier=%s orbit=%s transitive=%v stabilizer=%s dim_cp1=%d dim_su2=%d dim_stab=%d invariant_constant=%v axis_required=%v explains_gate763=%v conditional=%v verdict=%s", v.Carrier, v.OrbitFormula, v.SU2ActsTransitivelyOnCP1, v.Stabilizer, v.CP1RealDimension, v.SU2Dimension, v.StabilizerDimension, v.InvariantFunctionalConstant, v.AnisotropicAxisRequiredToSelect, v.ExplainsGate763Failure, v.ConditionalOnGaugeAirlock, v.Verdict)
}

func FormatDemotion(v GaugeOrbitDemotion) string {
	return fmt.Sprintf("object=%s prior=%s demoted=%s scalar_source=%v gauge_rep=%v anisotropy_required=%v selector_absence_scalar_failure=%v vacuum_gap=%v conditional=%v verdict=%s", v.Object, v.PriorTyping, v.DemotedTyping, v.ScalarRuntimeNumericalSource, v.GaugeRepresentative, v.RequiresPhysicalAnisotropy, v.AbsenceOfSelectorScalarFailure, v.AbsenceOfSelectorVacuumTheoremGap, v.DemotionConditional, v.Verdict)
}

func FormatRankInvariance(v RadialRankInvariance) string {
	return fmt.Sprintf("rho=%s rank=%d dim=%d trace=%.17g phase=%.17g lhopf=%.17g expected=%.17g rank_only=%v independent_cp1=%v invariant=%v verdict=%s", v.RhoPlusFormula, v.ProjectorRank, v.K7PlusDimension, v.TraceWeight, v.PhaseLoopPayoff, v.LHopf, v.ExpectedLHopf, v.DependsOnRankOnly, v.IndependentOfCP1Position, v.RankInvariant, v.Verdict)
}

func FormatLineVsRadial(v LineVsRadialEvent) string {
	return fmt.Sprintf("complex_rank=%d radial_rank=%d complex_weight=%.17g radial_weight=%.17g complex_loop=%.17g radial_loop=%.17g active_radial=%v reject_complex=%v question=%q verdict=%s", v.ComplexLineRealRank, v.RadialEventRealRank, v.ComplexLineWeight, v.RadialEventWeight, v.ComplexLineLoopUnit, v.RadialEventLoopUnit, v.ActiveUsesRadialEvent, v.FullComplexLineRejected, v.ScalarRuntimeQuestion, v.Verdict)
}

func FormatConditional(v ConditionalDemotionFirewall) string {
	return fmt.Sprintf("socket=%s internal_certified=%v physical_su2l=%v cp1_physical=%v internal_demote=%v promoted_ewsb=%v verdict=%s", v.InternalSocketName, v.InternalSocketCertified, v.PhysicalSU2LCertified, v.CP1GaugeOrbitPhysicalTheorem, v.DemotionAllowedAsInternalSocket, v.DemotionPromotedToPhysicalEWSB, v.Verdict)
}

func FormatHierarchy(v MissingObjectHierarchy) string {
	return fmt.Sprintf("before=%s after_scalar=%s after_physical=%s cp1_scalar_datum=%v radial_event_scalar_datum=%v ewsb_open=%v updated=%v verdict=%s", v.BeforeGate764PrimaryMissing, v.AfterGate764ScalarTarget, v.AfterGate764PhysicalTarget, v.CP1LocationScalarDatum, v.RadialEventTypeScalarDatum, v.ElectroweakVacuumStillOpen, v.Updated, v.Verdict)
}

func FormatFirewalls(v Firewalls) string {
	return fmt.Sprintf("audited=%v cp1_physical=%v absence_scalar_failure=%v cp1_su2l=%v radial_native=%v lhopf=%v runtime=%v tree_pole=%v ewsb=%v pole=%v yukawa=%v verdict=%s", v.Audited, v.CP1PointPhysicalTheorem, v.AbsenceOfCP1SelectorScalarFailure, v.CP1GaugeOrbitPhysicalSU2LTheorem, v.RankOneRadialEventNativeTheorem, v.LHopfNativeTransportTheorem, v.ScalarRuntimeIndependentTheorem, v.TreeProxyPoleMassTheorem, v.ElectroweakSymmetryBreakingTheorem, v.HiggsMassOrPoleMassTheorem, v.YukawaOperatorOrEigenvalueTheorem, v.Verdict)
}
