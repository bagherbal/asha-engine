// Package generation2cp1vacuumlineselectorfunctionalandmomentmapfirewallaudit implements
// Gate 763: CP1 Vacuum-Line Selector Functional and Moment-Map Firewall Audit.
//
// Gate 762 established that the complex Higgs vacuum line Pi_vac_C in
// K7+_J(n) ~= C^2 is the sharper missing object behind P_rad, and that no
// current ASHA structure selects a CP1 point. Gate 763 audits what kind of
// typed functional could select that CP1 point. It tests moment-map, scalar
// potential, boundary-history stress, Fano/quaternionic invariant, and
// spontaneous-orientation seal candidates. The result is firewall-preserving:
// a nonconstant CP1 functional or Hermitian SU(2) socket axis would be required,
// but none is natively certified. This is a selector-functional typing audit
// only. It does not derive electroweak symmetry breaking, radial gauge fixing,
// scalar runtime lambda, Higgs mass, pole mass, Yukawa operators, CKM/PMNS,
// flavor hierarchy, or a native HistoryLoopUnit theorem.
package generation2cp1vacuumlineselectorfunctionalandmomentmapfirewallaudit

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE763-CP1-VACUUM-LINE-SELECTOR-FUNCTIONAL-AND-MOMENT-MAP-FIREWALL-AUDIT"

	StatusGate762ComplexVacuumLineSealInherited               = "PASS_GATE762_COMPLEX_VACUUM_LINE_SEAL_INHERITED"
	StatusFunctionalSelectorRequirementDefined                = "PASS_FUNCTIONAL_SELECTOR_REQUIREMENT_DEFINED"
	StatusCP1MomentMapAudited                                 = "PASS_CP1_MOMENT_MAP_AUDITED"
	StatusScalarPotentialFunctionalAudited                    = "PASS_SCALAR_POTENTIAL_FUNCTIONAL_AUDITED"
	StatusBoundaryHistoryStressFunctionalAudited              = "PASS_BOUNDARY_HISTORY_STRESS_FUNCTIONAL_AUDITED"
	StatusFanoQuaternionicInvariantAudited                    = "PASS_FANO_QUATERNIONIC_INVARIANT_AUDITED"
	StatusOrientationSealOptionAudited                        = "PASS_ORIENTATION_SEAL_OPTION_AUDITED"
	StatusCandidateFunctionalRankingRecorded                  = "PASS_CANDIDATE_FUNCTIONAL_RANKING_RECORDED"
	StatusLineBeforeGaugeOrderPreserved                       = "PASS_LINE_BEFORE_RADIAL_GAUGE_ORDER_PRESERVED"
	StatusPhysicalFirewallsEnforced                           = "PASS_PHYSICAL_FIREWALLS_ENFORCED"
	StatusCP1SelectorRequiresNonconstantFunctionalOrAxis      = "CONDITIONAL_SUPPORT_CP1_SELECTOR_REQUIRES_NONCONSTANT_FUNCTIONAL_OR_HERMITIAN_AXIS"
	StatusMomentMapCanSelectOnlyAfterSuppliedSU2Axis          = "CONDITIONAL_SUPPORT_MOMENT_MAP_CAN_SELECT_ONLY_AFTER_SUPPLIED_SU2_AXIS"
	StatusU2InvariantScalarPotentialCP1Flat                   = "CONDITIONAL_SUPPORT_U2_INVARIANT_SCALAR_POTENTIAL_IS_CP1_FLAT"
	StatusBoundaryScalarsCannotSelectCP1WithoutVectorCoupling = "CONDITIONAL_SUPPORT_BOUNDARY_SCALARS_CANNOT_SELECT_CP1_POINT_WITHOUT_VECTOR_COUPLING"
	StatusPiVacCRemainsComplexVacuumLineSeal                  = "CONDITIONAL_SUPPORT_PI_VAC_C_REMAINS_COMPLEX_VACUUM_LINE_SEAL"
	StatusRadialGaugeFixingRemainsSecondary                   = "CONDITIONAL_SUPPORT_RADIAL_GAUGE_FIXING_REMAINS_SECONDARY_TO_CP1_LINE_SELECTION"
	StatusNoNativeCP1SelectorFunctional                       = "FAILED_ROUTE_NO_NATIVE_CP1_SELECTOR_FUNCTIONAL"
	StatusNoNativeSU2MomentMapAxis                            = "FAILED_ROUTE_NO_NATIVE_SU2_MOMENT_MAP_AXIS"
	StatusNoNativeHermitianHiggsSocketAxis                    = "FAILED_ROUTE_NO_NATIVE_HERMITIAN_HIGGS_SOCKET_AXIS"
	StatusScalarPotentialNotTypedOrientationSelector          = "FAILED_ROUTE_SCALAR_POTENTIAL_NOT_TYPED_ORIENTATION_SELECTOR"
	StatusBoundaryStressScalarNotCP1Functional                = "FAILED_ROUTE_BOUNDARY_STRESS_IS_SCALAR_NOT_CP1_FUNCTIONAL"
	StatusFanoQuaternionicDoesNotSelectCP1Point               = "FAILED_ROUTE_FANO_QUATERNIONIC_STRUCTURE_DOES_NOT_SELECT_CP1_POINT"
	StatusSpontaneousOrientationSealNotNativeTheorem          = "FAILED_ROUTE_SPONTANEOUS_ORIENTATION_SEAL_NOT_NATIVE_THEOREM"
	StatusNoNativeEWSBTheorem                                 = "FAILED_ROUTE_NO_NATIVE_ELECTROWEAK_SYMMETRY_BREAKING_THEOREM"
	StatusNoNativeHistoryLoopUnitTheorem                      = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                        = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                 = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate763CP1FunctionalSelectorBoundary                = "FIREWALL_PRESERVED_GATE763_CP1_FUNCTIONAL_SELECTOR_BOUNDARY"
)

const (
	k7PlusRealDim       = 4
	k7PlusComplexDim    = 2
	cp1ComplexDim       = 1
	cp1RealDim          = 2
	complexLineRealRank = 2
	complexLineRank     = 1
	su2AxisDim          = 3
)

type Gate762Inheritance struct {
	Inherited                    bool
	Socket                       string
	MissingObject                string
	MissingObjectRankR           int
	MissingObjectRankC           int
	PRadSecondary                bool
	NoCurrentCP1Selector         bool
	ComplexVacuumLineSealRemains bool
	Verdict                      string
}

type FunctionalRequirement struct {
	Question                           string
	Domain                             string
	Target                             string
	EquivalentHermitianAxis            string
	RequiresNonconstantFunctional      bool
	RequiresIsolatedCriticalLine       bool
	U2InvariantFunctionalSelectsLine   bool
	ConstantFunctionalLeavesCP1Flat    bool
	RadialGaugeRequiredForLineSelector bool
	Verdict                            string
}

type MomentMapAudit struct {
	MomentMapFormula           string
	HamiltonianFormula         string
	MomentMapExistsOnCP1       bool
	RequiresSU2Axis            bool
	SuppliedAxisCertified      bool
	WouldSelectEigenlineIfAxis bool
	NativeSelectorCertified    bool
	Verdict                    string
}

type ScalarPotentialAudit struct {
	InvariantPotentialFormula      string
	AnisotropicPotentialFormula    string
	U2Invariant                    bool
	FlatOnCP1                      bool
	AnisotropicTermWouldSelectLine bool
	AnisotropicTermCertified       bool
	NativeOrientationSelector      bool
	Verdict                        string
}

type BoundaryStressAudit struct {
	BoundaryObjects           []string
	CollapsedScalarLayer      bool
	ProvidesVectorInK7Plus    bool
	ProvidesHermitianAxisOnC2 bool
	CanDefineCP1Functional    bool
	CanSelectCP1Point         bool
	NeedsTypedVectorCoupling  bool
	Verdict                   string
}

type FanoQuaternionicAudit struct {
	SuppliesTwistorFamily          bool
	SuppliesJHSocket               bool
	SuppliesU2Socket               bool
	SelectsComplexStructureN       bool
	SelectsCP1Point                bool
	InvariantVectorInK7Plus        bool
	SymmetryWouldMakeLineArbitrary bool
	Verdict                        string
}

type OrientationSealAudit struct {
	SealName                string
	CanSelectCP1IfSupplied  bool
	NativeTheoremCertified  bool
	WouldPrecedeRadialGauge bool
	WouldNotDeriveEWSB      bool
	WouldNotDeriveHiggsMass bool
	Verdict                 string
}

type CandidateFunctional struct {
	Name                    string
	CanDefineFunctional     bool
	CanSelectWithExtraInput bool
	ExtraInputRequired      string
	NativeSelectorCertified bool
	FailureMode             string
	Priority                int
}

type CandidateRanking struct {
	Candidates                 []CandidateFunctional
	BestTypedFutureCandidate   string
	HighestNativeResult        string
	RankingRecorded            bool
	AnyNativeSelectorCertified bool
	Verdict                    string
}

type SealOrder struct {
	Step1           string
	Step2           string
	Step3           string
	Step4           string
	LineBeforeGauge bool
	LHopfAfterGauge bool
	Preserved       bool
	Verdict         string
}

type Firewalls struct {
	Audited                           bool
	CP1FunctionalNativeSelector       bool
	MomentMapAxisNative               bool
	HermitianAxisNative               bool
	ScalarPotentialNativeOrientation  bool
	BoundaryStressCP1Functional       bool
	PiVacCNativeEWSBTheorem           bool
	RadialGaugeFixingNativeTheorem    bool
	LHopfNativeHistoryLoopTheorem     bool
	ScalarRuntimeIndependentTheorem   bool
	HiggsMassOrPoleMassTheorem        bool
	YukawaOperatorOrEigenvalueTheorem bool
	Verdict                           string
}

type Analysis struct {
	Gate762         Gate762Inheritance
	Requirement     FunctionalRequirement
	MomentMap       MomentMapAudit
	ScalarPotential ScalarPotentialAudit
	BoundaryStress  BoundaryStressAudit
	Fano            FanoQuaternionicAudit
	Orientation     OrientationSealAudit
	Ranking         CandidateRanking
	SealOrder       SealOrder
	Firewalls       Firewalls
	Truth           string
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
	a := &Analysis{
		Gate762: Gate762Inheritance{
			Inherited:                    true,
			Socket:                       "K7+_J(n) ~= C^2",
			MissingObject:                "Pi_vac_C in CP1 = P(K7+_J(n))",
			MissingObjectRankR:           complexLineRealRank,
			MissingObjectRankC:           complexLineRank,
			PRadSecondary:                true,
			NoCurrentCP1Selector:         true,
			ComplexVacuumLineSealRemains: true,
			Verdict:                      StatusGate762ComplexVacuumLineSealInherited,
		},
		Requirement: FunctionalRequirement{
			Question:                           "Can a typed scalar/Higgs functional select Pi_vac_C in CP1?",
			Domain:                             "CP1 = P(C^2)",
			Target:                             "isolated complex rank-one projector Pi_vac_C",
			EquivalentHermitianAxis:            "nonzero traceless Hermitian axis H in su(2) on K7+_J(n)",
			RequiresNonconstantFunctional:      true,
			RequiresIsolatedCriticalLine:       true,
			U2InvariantFunctionalSelectsLine:   false,
			ConstantFunctionalLeavesCP1Flat:    true,
			RadialGaugeRequiredForLineSelector: false,
			Verdict:                            strings.Join([]string{StatusFunctionalSelectorRequirementDefined, StatusCP1SelectorRequiresNonconstantFunctionalOrAxis}, "; "),
		},
		MomentMap: MomentMapAudit{
			MomentMapFormula:           "mu([z]) = zz^dagger/<z,z> - (1/2)I in su(2)^*",
			HamiltonianFormula:         "H_h([z]) = <h, mu([z])>; h != 0 selects an eigenline/pole pair",
			MomentMapExistsOnCP1:       true,
			RequiresSU2Axis:            true,
			SuppliedAxisCertified:      false,
			WouldSelectEigenlineIfAxis: true,
			NativeSelectorCertified:    false,
			Verdict:                    strings.Join([]string{StatusCP1MomentMapAudited, StatusMomentMapCanSelectOnlyAfterSuppliedSU2Axis, StatusNoNativeSU2MomentMapAxis, StatusNoNativeHermitianHiggsSocketAxis}, "; "),
		},
		ScalarPotential: ScalarPotentialAudit{
			InvariantPotentialFormula:      "V0(|z|^2) = alpha |z|^2 + beta |z|^4",
			AnisotropicPotentialFormula:    "V_H([z]) = <z,H z>/<z,z> with nonzero Hermitian H",
			U2Invariant:                    true,
			FlatOnCP1:                      true,
			AnisotropicTermWouldSelectLine: true,
			AnisotropicTermCertified:       false,
			NativeOrientationSelector:      false,
			Verdict:                        strings.Join([]string{StatusScalarPotentialFunctionalAudited, StatusU2InvariantScalarPotentialCP1Flat, StatusScalarPotentialNotTypedOrientationSelector}, "; "),
		},
		BoundaryStress: BoundaryStressAudit{
			BoundaryObjects:           []string{"lambda(Lambda_12)", "R3-1", "xi_boundary", "F_wall_3_red", "kappa_lambda_red"},
			CollapsedScalarLayer:      true,
			ProvidesVectorInK7Plus:    false,
			ProvidesHermitianAxisOnC2: false,
			CanDefineCP1Functional:    false,
			CanSelectCP1Point:         false,
			NeedsTypedVectorCoupling:  true,
			Verdict:                   strings.Join([]string{StatusBoundaryHistoryStressFunctionalAudited, StatusBoundaryScalarsCannotSelectCP1WithoutVectorCoupling, StatusBoundaryStressScalarNotCP1Functional}, "; "),
		},
		Fano: FanoQuaternionicAudit{
			SuppliesTwistorFamily:          true,
			SuppliesJHSocket:               true,
			SuppliesU2Socket:               true,
			SelectsComplexStructureN:       false,
			SelectsCP1Point:                false,
			InvariantVectorInK7Plus:        false,
			SymmetryWouldMakeLineArbitrary: true,
			Verdict:                        strings.Join([]string{StatusFanoQuaternionicInvariantAudited, StatusFanoQuaternionicDoesNotSelectCP1Point}, "; "),
		},
		Orientation: OrientationSealAudit{
			SealName:                "SpontaneousOrientationSeal / ComplexVacuumLineSeal",
			CanSelectCP1IfSupplied:  true,
			NativeTheoremCertified:  false,
			WouldPrecedeRadialGauge: true,
			WouldNotDeriveEWSB:      true,
			WouldNotDeriveHiggsMass: true,
			Verdict:                 strings.Join([]string{StatusOrientationSealOptionAudited, StatusSpontaneousOrientationSealNotNativeTheorem, StatusPiVacCRemainsComplexVacuumLineSeal}, "; "),
		},
		Ranking: buildCandidateRanking(),
		SealOrder: SealOrder{
			Step1:           "n selects J_H(n), the complex structure on K7+",
			Step2:           "a nonconstant CP1 functional or supplied orientation seal selects Pi_vac_C",
			Step3:           "radial gauge fixing selects P_rad inside Pi_vac_C",
			Step4:           "L_Hopf uses Tr(rho_plus P_rad)=1/4 after both prior choices",
			LineBeforeGauge: true,
			LHopfAfterGauge: true,
			Preserved:       true,
			Verdict:         strings.Join([]string{StatusLineBeforeGaugeOrderPreserved, StatusRadialGaugeFixingRemainsSecondary}, "; "),
		},
		Firewalls: Firewalls{
			Audited:                           true,
			CP1FunctionalNativeSelector:       false,
			MomentMapAxisNative:               false,
			HermitianAxisNative:               false,
			ScalarPotentialNativeOrientation:  false,
			BoundaryStressCP1Functional:       false,
			PiVacCNativeEWSBTheorem:           false,
			RadialGaugeFixingNativeTheorem:    false,
			LHopfNativeHistoryLoopTheorem:     false,
			ScalarRuntimeIndependentTheorem:   false,
			HiggsMassOrPoleMassTheorem:        false,
			YukawaOperatorOrEigenvalueTheorem: false,
			Verdict:                           StatusGate763CP1FunctionalSelectorBoundary,
		},
		Truth: "Gate 763 finds that a CP1 vacuum-line selector would require a nonconstant scalar/Higgs functional, equivalently a supplied Hermitian SU(2) socket axis; no current ASHA structure certifies such a native selector, so Pi_vac_C remains ComplexVacuumLineSeal.",
	}
	cache = a
	clone := *a
	return &clone, nil
}

func buildCandidateRanking() CandidateRanking {
	candidates := []CandidateFunctional{
		{Name: "SU(2) moment-map Hamiltonian", CanDefineFunctional: true, CanSelectWithExtraInput: true, ExtraInputRequired: "nonzero h in su(2)^* / Hermitian socket axis", NativeSelectorCertified: false, FailureMode: StatusNoNativeSU2MomentMapAxis, Priority: 1},
		{Name: "anisotropic scalar potential on K7+_J(n)", CanDefineFunctional: true, CanSelectWithExtraInput: true, ExtraInputRequired: "typed non-U(2)-invariant Hermitian term", NativeSelectorCertified: false, FailureMode: StatusScalarPotentialNotTypedOrientationSelector, Priority: 2},
		{Name: "spontaneous orientation seal", CanDefineFunctional: true, CanSelectWithExtraInput: true, ExtraInputRequired: "sealed CP1 point/orientation", NativeSelectorCertified: false, FailureMode: StatusSpontaneousOrientationSealNotNativeTheorem, Priority: 3},
		{Name: "boundary-history stress", CanDefineFunctional: false, CanSelectWithExtraInput: false, ExtraInputRequired: "typed vector/Hermitian coupling from boundary scalar to K7+", NativeSelectorCertified: false, FailureMode: StatusBoundaryStressScalarNotCP1Functional, Priority: 4},
		{Name: "Fano/quaternionic invariant", CanDefineFunctional: false, CanSelectWithExtraInput: false, ExtraInputRequired: "symmetry-breaking vector beyond invariant socket", NativeSelectorCertified: false, FailureMode: StatusFanoQuaternionicDoesNotSelectCP1Point, Priority: 5},
	}
	return CandidateRanking{
		Candidates:                 candidates,
		BestTypedFutureCandidate:   "SU(2) moment-map Hamiltonian or equivalent Hermitian Higgs socket axis",
		HighestNativeResult:        "CP1 socket geometry is available, but no CP1 point selector is certified",
		RankingRecorded:            true,
		AnyNativeSelectorCertified: false,
		Verdict:                    strings.Join([]string{StatusCandidateFunctionalRankingRecorded, StatusNoNativeCP1SelectorFunctional}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate762ComplexVacuumLineSealInherited,
		StatusFunctionalSelectorRequirementDefined,
		StatusCP1MomentMapAudited,
		StatusScalarPotentialFunctionalAudited,
		StatusBoundaryHistoryStressFunctionalAudited,
		StatusFanoQuaternionicInvariantAudited,
		StatusOrientationSealOptionAudited,
		StatusCandidateFunctionalRankingRecorded,
		StatusLineBeforeGaugeOrderPreserved,
		StatusPhysicalFirewallsEnforced,
		StatusCP1SelectorRequiresNonconstantFunctionalOrAxis,
		StatusMomentMapCanSelectOnlyAfterSuppliedSU2Axis,
		StatusU2InvariantScalarPotentialCP1Flat,
		StatusBoundaryScalarsCannotSelectCP1WithoutVectorCoupling,
		StatusPiVacCRemainsComplexVacuumLineSeal,
		StatusRadialGaugeFixingRemainsSecondary,
		StatusNoNativeCP1SelectorFunctional,
		StatusNoNativeSU2MomentMapAxis,
		StatusNoNativeHermitianHiggsSocketAxis,
		StatusScalarPotentialNotTypedOrientationSelector,
		StatusBoundaryStressScalarNotCP1Functional,
		StatusFanoQuaternionicDoesNotSelectCP1Point,
		StatusSpontaneousOrientationSealNotNativeTheorem,
		StatusNoNativeEWSBTheorem,
		StatusNoNativeHistoryLoopUnitTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate763CP1FunctionalSelectorBoundary,
	}
}

func FormatGate762(v Gate762Inheritance) string {
	return fmt.Sprintf("inherited=%v socket=%s missing=%s rankR=%d rankC=%d prad_secondary=%v no_cp1_selector=%v seal_remains=%v verdict=%s", v.Inherited, v.Socket, v.MissingObject, v.MissingObjectRankR, v.MissingObjectRankC, v.PRadSecondary, v.NoCurrentCP1Selector, v.ComplexVacuumLineSealRemains, v.Verdict)
}

func FormatRequirement(v FunctionalRequirement) string {
	return fmt.Sprintf("question=%q domain=%s target=%s hermitian_axis=%s nonconstant=%v isolated=%v u2_selects=%v flat=%v radial_gauge_required=%v verdict=%s", v.Question, v.Domain, v.Target, v.EquivalentHermitianAxis, v.RequiresNonconstantFunctional, v.RequiresIsolatedCriticalLine, v.U2InvariantFunctionalSelectsLine, v.ConstantFunctionalLeavesCP1Flat, v.RadialGaugeRequiredForLineSelector, v.Verdict)
}

func FormatMomentMap(v MomentMapAudit) string {
	return fmt.Sprintf("mu=%s H=%s exists=%v requires_axis=%v axis_certified=%v would_select=%v native=%v verdict=%s", v.MomentMapFormula, v.HamiltonianFormula, v.MomentMapExistsOnCP1, v.RequiresSU2Axis, v.SuppliedAxisCertified, v.WouldSelectEigenlineIfAxis, v.NativeSelectorCertified, v.Verdict)
}

func FormatScalarPotential(v ScalarPotentialAudit) string {
	return fmt.Sprintf("V0=%s VH=%s u2=%v flat_cp1=%v anisotropic_selects=%v anisotropic_certified=%v native_orientation=%v verdict=%s", v.InvariantPotentialFormula, v.AnisotropicPotentialFormula, v.U2Invariant, v.FlatOnCP1, v.AnisotropicTermWouldSelectLine, v.AnisotropicTermCertified, v.NativeOrientationSelector, v.Verdict)
}

func FormatBoundaryStress(v BoundaryStressAudit) string {
	return fmt.Sprintf("objects=%s collapsed=%v vector_k7plus=%v hermitian_axis=%v cp1_functional=%v cp1_point=%v needs_coupling=%v verdict=%s", strings.Join(v.BoundaryObjects, "+"), v.CollapsedScalarLayer, v.ProvidesVectorInK7Plus, v.ProvidesHermitianAxisOnC2, v.CanDefineCP1Functional, v.CanSelectCP1Point, v.NeedsTypedVectorCoupling, v.Verdict)
}

func FormatFano(v FanoQuaternionicAudit) string {
	return fmt.Sprintf("twistor=%v jh_socket=%v u2_socket=%v selects_n=%v selects_cp1=%v invariant_vector=%v arbitrary_line=%v verdict=%s", v.SuppliesTwistorFamily, v.SuppliesJHSocket, v.SuppliesU2Socket, v.SelectsComplexStructureN, v.SelectsCP1Point, v.InvariantVectorInK7Plus, v.SymmetryWouldMakeLineArbitrary, v.Verdict)
}

func FormatOrientation(v OrientationSealAudit) string {
	return fmt.Sprintf("seal=%s can_select=%v native=%v precedes_gauge=%v no_ewsb=%v no_higgs=%v verdict=%s", v.SealName, v.CanSelectCP1IfSupplied, v.NativeTheoremCertified, v.WouldPrecedeRadialGauge, v.WouldNotDeriveEWSB, v.WouldNotDeriveHiggsMass, v.Verdict)
}

func FormatRanking(v CandidateRanking) string {
	parts := make([]string, 0, len(v.Candidates))
	for _, c := range v.Candidates {
		parts = append(parts, fmt.Sprintf("%d:%s(native=%v extra=%s fail=%s)", c.Priority, c.Name, c.NativeSelectorCertified, c.ExtraInputRequired, c.FailureMode))
	}
	return fmt.Sprintf("recorded=%v any_native=%v best=%s highest=%s candidates=[%s] verdict=%s", v.RankingRecorded, v.AnyNativeSelectorCertified, v.BestTypedFutureCandidate, v.HighestNativeResult, strings.Join(parts, "; "), v.Verdict)
}

func FormatSealOrder(v SealOrder) string {
	return fmt.Sprintf("1=%s 2=%s 3=%s 4=%s line_before_gauge=%v lhopf_after_gauge=%v preserved=%v verdict=%s", v.Step1, v.Step2, v.Step3, v.Step4, v.LineBeforeGauge, v.LHopfAfterGauge, v.Preserved, v.Verdict)
}

func FormatFirewalls(v Firewalls) string {
	return fmt.Sprintf("audited=%v cp1_functional=%v moment_axis=%v hermitian_axis=%v scalar_potential=%v boundary_cp1=%v ewsb=%v radial_gauge=%v lhopf=%v runtime=%v pole=%v yukawa=%v verdict=%s", v.Audited, v.CP1FunctionalNativeSelector, v.MomentMapAxisNative, v.HermitianAxisNative, v.ScalarPotentialNativeOrientation, v.BoundaryStressCP1Functional, v.PiVacCNativeEWSBTheorem, v.RadialGaugeFixingNativeTheorem, v.LHopfNativeHistoryLoopTheorem, v.ScalarRuntimeIndependentTheorem, v.HiggsMassOrPoleMassTheorem, v.YukawaOperatorOrEigenvalueTheorem, v.Verdict)
}
