// Package generation2scalaredgesocketoperatorrealizationaudit implements
// Gate 860: Scalar Edge-Socket Operator Realization Audit.
//
// Gate 860 follows Gate 859's stabilizer-branch first-order support / edge-
// centrality audit.  Gate 859 showed that first-order support compatibility in
// the Higgs-oriented stabilizer branch requires the color edges to be scalar on
// the P_3 color factor and the lepton edge to be scalar on the P_1 lepton
// factor.  Gate 860 now realizes those support labels as operator-valued
// symbolic socket maps:
//
//	Y = y_+3 |h_+><e_+| tensor I_{P_3}
//	  + y_-3 |h_-><e_-| tensor I_{P_3}
//	  + y_-1 |h_-><e_-| tensor I_{P_1},
//
// with y_+1=0.  This is stronger than support anatomy, but it remains a
// symbolic post-orientation seal.  The y sockets are not observed Yukawa values,
// no alpha_B source is derived, no sector trace-magnitude readout is certified,
// no full unbroken A_F theorem is claimed, and no R3/R4 promotion or official
// ledger update is allowed.
package generation2scalaredgesocketoperatorrealizationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE860-SCALAR-EDGE-SOCKET-OPERATOR-REALIZATION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank       = 1
	P3Rank       = 3
	WRank        = P1Rank + P3Rank
	HLRank       = 2 * WRank
	HRMinRank    = 7
	HPartMinRank = HLRank + HRMinRank
	YRankFull    = HRMinRank
	DSymRankFull = 2 * YRankFull
	KernelRank   = HPartMinRank - DSymRankFull

	StatusGate859Inherited     = "PASS_GATE859_EDGE_CENTRALITY_SUPPORT_INHERITED"
	StatusYOperatorDefined     = "PASS_OPERATOR_VALUED_Y_MAP_DEFINED"
	StatusColorEdgesScalar     = "PASS_Y_PLUS3_AND_Y_MINUS3_SCALAR_ON_COLOR_FACTOR"
	StatusLeptonEdgeTrivial    = "PASS_Y_MINUS1_COLOR_TRIVIAL_ON_LEPTON_FACTOR"
	StatusPunctureZero         = "PASS_PUNCTURE_EDGE_REMAINS_ZERO"
	StatusDRebuilt             = "PASS_D_F_SYM_OPERATOR_MATRIX_REBUILT_FROM_Y"
	StatusRankKernelRecomputed = "PASS_RANK_AND_KERNEL_LEDGER_RECOMPUTED_FOR_NONZERO_SYMBOLIC_SOCKETS"
	StatusFirstOrderCentrality = "PASS_FIRST_ORDER_SUPPORT_COMPATIBILITY_INHERITS_GATE859_CENTRALITY"
	StatusNoObservedDataUsed   = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen         = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict      = "FIREWALL_PRESERVED_GATE860_SCALAR_EDGE_SOCKET_OPERATOR_REALIZATION_NOT_R3"

	SupportOperatorYFirstOrderIfScalar = "CONDITIONAL_SUPPORT_OPERATOR_VALUED_EDGE_MAP_IS_FIRST_ORDER_SUPPORT_COMPATIBLE_IF_Y_COEFFICIENTS_ARE_SCALAR_SOCKETS"
	SupportColorCentralityRepairsM3    = "CONDITIONAL_SUPPORT_COLOR_CENTRALITY_REPAIRS_OPPOSITE_M3_PRESSURE_AT_SUPPORT_LEVEL"
	SupportLeftKernelPersists          = "CONDITIONAL_SUPPORT_LEFT_KERNEL_SINGLETON_PERSISTS_FOR_NONZERO_ACTIVE_SOCKETS"
	SupportNonzeroSocketsRankLedger    = "CONDITIONAL_SUPPORT_NONZERO_SYMBOLIC_SOCKETS_GIVE_RANK7_Y_AND_RANK14_D"
	SupportPunctureStillZero           = "CONDITIONAL_SUPPORT_Y_PLUS1_ZERO_PRESERVES_RIGHT_PUNCTURE_AND_LEFT_KERNEL"
	SupportStabilizerOperatorSeal      = "CONDITIONAL_SUPPORT_POST_ORIENTATION_OPERATOR_VALUED_D_F_SUPPORT_MATRIX_SEAL"

	FailureOperatorYSymbolicNotYukawa  = "FAILED_ROUTE_OPERATOR_VALUED_Y_IS_SYMBOLIC_SOCKET_MATRIX_NOT_YUKAWA_THEOREM"
	FailureNoNumericalYukawa           = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoAlphaSource               = "FAILED_ROUTE_NO_NATIVE_ALPHA_SOURCE"
	FailureNoTraceReadout              = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoFullUnbrokenAFTheorem     = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF           = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoFullOperatorFirstOrder    = "FAILED_ROUTE_NO_FULL_OPERATOR_LEVEL_FIRST_ORDER_THEOREM"
	FailureNoCompleteJOppositeProof    = "FAILED_ROUTE_NO_COMPLETE_J_OPPOSITE_OPERATOR_PROOF"
	FailureNoBimoduleCommutantProof    = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureScalarSocketSupportOnly     = "FAILED_ROUTE_SCALAR_EDGE_SOCKET_REALIZATION_IS_SUPPORT_SEAL_NOT_NATIVE_OPERATOR_THEOREM"
	FailureNoOperatorYukawaTheorem     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoOfficialNEffUpdate        = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate       = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                       = "FAILED_ROUTE_R2_OPERATOR_VALUED_SUPPORT_MATRIX_NOT_R3"
	FailureNotR4                       = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment        = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem           = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem    = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureZeroSocketBranchEnlargesKer = "FAILED_ROUTE_ZERO_ACTIVE_SOCKET_BRANCH_ENLARGES_KERNEL_AND_IS_NOT_CURRENT_MINIMAL_RANK7_SUPPORT"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type EdgeOperator struct {
	Name, Domain, Codomain, OperatorForm, VisibleFactor string
	Rank                                                int
	Present, CoefficientSymbolic                        bool
	ScalarOnVisibleFactor, IdentityOnVisibleFactor      bool
	ColorEdge, LeptonEdge, PunctureEdge                 bool
	NumericalValue, YukawaMagnitude                     bool
	Supports, Failures                                  []string
}

type YOperator struct {
	Expression, PunctureExpression                            string
	OperatorValued, SymbolicSocketMatrix, FirstOrderCandidate bool
	ColorCentrality, LeptonTriviality, PunctureZero           bool
	ActiveEdgeCount, RankIfActiveSocketsNonzero               int
	ZeroActiveSocketBranchEnlargesKernel                      bool
	Supports, Failures                                        []string
}

type DMatrix struct {
	Expression                                      string
	BuiltFromY, SelfAdjointByBlockForm              bool
	ChiralBlockForm, PostOrientationOnly            bool
	RankIfActiveSocketsNonzero, KernelRankIfNonzero int
	KernelSingleton                                 string
	Supports, Failures                              []string
}

type FirstOrderPosition struct {
	Algebra, Target                                              string
	Gate859Inherited, EdgeCentralityInstalled                    bool
	OperatorLevelFirstOrderCertified, CompleteJOppositeCertified bool
	Supports, Failures                                           []string
}

type Impact struct {
	Classification                                                                                  string
	OperatorYDefined, ScalarColorEdges, LeptonEdgeTrivial, PunctureZero, RankKernelLedgerRecomputed bool
	OperatorFirstOrderProof, NativeFiniteTripleProof                                                bool
	AlphaStillSealed, MagnitudesStillMissing                                                        bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4                bool
}

type Firewalls struct {
	Enforced                                                                                              bool
	OperatorYSymbolicNotYukawa, NoNumericalYukawa, NoAlphaSource, NoTraceReadout, NoFullUnbrokenAFTheorem bool
	AForientNotFullAF, NoFullFirstOrder, NoCompleteJOpposite, NoBimoduleProof, ScalarSocketSupportOnly    bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem    bool
	NoThreeGenerationTheorem                                                                              bool
	Verdict                                                                                               string
}

type Audit struct {
	ID         string
	Ledger     Ledger
	Edges      []EdgeOperator
	Y          YOperator
	D          DMatrix
	FirstOrder FirstOrderPosition
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Edges: []EdgeOperator{
			{Name: "Y_+3", Domain: "e_+ tensor P_3", Codomain: "h_+ tensor P_3", OperatorForm: "y_+3 |h_+><e_+| tensor I_{P_3}", VisibleFactor: "P_3 color", Rank: P3Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, ColorEdge: true, Supports: []string{StatusColorEdgesScalar, SupportColorCentralityRepairsM3}, Failures: []string{FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa}},
			{Name: "Y_-3", Domain: "e_- tensor P_3", Codomain: "h_- tensor P_3", OperatorForm: "y_-3 |h_-><e_-| tensor I_{P_3}", VisibleFactor: "P_3 color", Rank: P3Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, ColorEdge: true, Supports: []string{StatusColorEdgesScalar, SupportColorCentralityRepairsM3}, Failures: []string{FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa}},
			{Name: "Y_-1", Domain: "e_- tensor P_1", Codomain: "h_- tensor P_1", OperatorForm: "y_-1 |h_-><e_-| tensor I_{P_1}", VisibleFactor: "P_1 lepton", Rank: P1Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, LeptonEdge: true, Supports: []string{StatusLeptonEdgeTrivial}, Failures: []string{FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa}},
			{Name: "Y_+1", Domain: "e_+ tensor P_1", Codomain: "h_+ tensor P_1", OperatorForm: "0", VisibleFactor: "P_1 lepton", Rank: P1Rank, Present: false, CoefficientSymbolic: false, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: false, PunctureEdge: true, Supports: []string{StatusPunctureZero, SupportPunctureStillZero}, Failures: []string{FailureScalarSocketSupportOnly, FailureNoNeutrinoTheorem}},
		},
		Y: YOperator{
			Expression: "Y=y_+3 |h_+><e_+| tensor I_{P_3} + y_-3 |h_-><e_-| tensor I_{P_3} + y_-1 |h_-><e_-| tensor I_{P_1}", PunctureExpression: "Y_+1=0",
			OperatorValued: true, SymbolicSocketMatrix: true, FirstOrderCandidate: true, ColorCentrality: true, LeptonTriviality: true, PunctureZero: true,
			ActiveEdgeCount: 3, RankIfActiveSocketsNonzero: YRankFull, ZeroActiveSocketBranchEnlargesKernel: true,
			Supports: []string{StatusYOperatorDefined, StatusColorEdgesScalar, StatusLeptonEdgeTrivial, StatusPunctureZero, SupportOperatorYFirstOrderIfScalar, SupportNonzeroSocketsRankLedger, SupportPunctureStillZero, SupportStabilizerOperatorSeal},
			Failures: []string{FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa, FailureZeroSocketBranchEnlargesKer},
		},
		D: DMatrix{
			Expression: "D_F^sym=[[0,Y^dagger],[Y,0]]", BuiltFromY: true, SelfAdjointByBlockForm: true, ChiralBlockForm: true, PostOrientationOnly: true,
			RankIfActiveSocketsNonzero: DSymRankFull, KernelRankIfNonzero: KernelRank, KernelSingleton: "h_+ tensor P_1",
			Supports: []string{StatusDRebuilt, StatusRankKernelRecomputed, SupportLeftKernelPersists, SupportNonzeroSocketsRankLedger},
			Failures: []string{FailureScalarSocketSupportOnly, FailureNoFullUnbrokenAFTheorem, FailureAForientNotFullAF},
		},
		FirstOrder: FirstOrderPosition{
			Algebra: "A_F^orient=C_R plus C_H plus M_3(C)", Target: "[[D_F^sym,rho_F(a)],rho_F^op(b)]=0",
			Gate859Inherited: true, EdgeCentralityInstalled: true, OperatorLevelFirstOrderCertified: false, CompleteJOppositeCertified: false,
			Supports: []string{StatusGate859Inherited, StatusFirstOrderCentrality, SupportOperatorYFirstOrderIfScalar, SupportColorCentralityRepairsM3},
			Failures: []string{FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof, FailureNoBimoduleCommutantProof, FailureAForientNotFullAF},
		},
		Impact:    Impact{Classification: "R2+++++_operator_valued_scalar_edge_socket_matrix_seal", OperatorYDefined: true, ScalarColorEdges: true, LeptonEdgeTrivial: true, PunctureZero: true, RankKernelLedgerRecomputed: true, OperatorFirstOrderProof: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true},
		Firewalls: Firewalls{Enforced: true, OperatorYSymbolicNotYukawa: true, NoNumericalYukawa: true, NoAlphaSource: true, NoTraceReadout: true, NoFullUnbrokenAFTheorem: true, AForientNotFullAF: true, NoFullFirstOrder: true, NoCompleteJOpposite: true, NoBimoduleProof: true, ScalarSocketSupportOnly: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict},
		Truth:     "Gate 860 realizes the Gate 859 edge-centrality rule as an operator-valued symbolic socket matrix: the color edges are scalar identity maps on P_3, the lepton edge is scalar on P_1, and Y_+1 remains zero.",
		Final:     "The branch now has an operator-valued post-orientation finite Dirac support matrix seal. The y sockets remain symbolic edge amplitudes only; no Yukawa magnitudes, alpha source, first-order operator theorem, R3/R4 promotion, or official ledger update is certified.",
	}
	return a, a.Validate()
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger overpromoted")
	}
	if len(a.Edges) != 4 {
		return err("expected three active edges plus one puncture edge")
	}
	activeRank := 0
	activeCount := 0
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude {
			return err("edge overpromoted to numerical/magnitude data")
		}
		if e.Present {
			activeCount++
			activeRank += e.Rank
			if !e.CoefficientSymbolic || !e.ScalarOnVisibleFactor || !e.IdentityOnVisibleFactor {
				return err("active edge must be symbolic scalar identity on visible support factor")
			}
		}
		if e.PunctureEdge && (e.Present || e.IdentityOnVisibleFactor || e.OperatorForm != "0") {
			return err("puncture edge must remain zero")
		}
	}
	if activeCount != 3 || activeRank != YRankFull {
		return err("active edge rank ledger inconsistent")
	}
	if !a.Y.OperatorValued || !a.Y.SymbolicSocketMatrix || !a.Y.FirstOrderCandidate || !a.Y.ColorCentrality || !a.Y.LeptonTriviality || !a.Y.PunctureZero || a.Y.ActiveEdgeCount != 3 || a.Y.RankIfActiveSocketsNonzero != YRankFull || !a.Y.ZeroActiveSocketBranchEnlargesKernel {
		return err("Y operator flags inconsistent")
	}
	if !a.D.BuiltFromY || !a.D.SelfAdjointByBlockForm || !a.D.ChiralBlockForm || !a.D.PostOrientationOnly || a.D.RankIfActiveSocketsNonzero != DSymRankFull || a.D.KernelRankIfNonzero != KernelRank || a.D.KernelSingleton != "h_+ tensor P_1" {
		return err("D matrix ledger inconsistent")
	}
	if !a.FirstOrder.Gate859Inherited || !a.FirstOrder.EdgeCentralityInstalled || a.FirstOrder.OperatorLevelFirstOrderCertified || a.FirstOrder.CompleteJOppositeCertified {
		return err("first-order position overpromoted or inconsistent")
	}
	if !a.Impact.OperatorYDefined || !a.Impact.ScalarColorEdges || !a.Impact.LeptonEdgeTrivial || !a.Impact.PunctureZero || !a.Impact.RankKernelLedgerRecomputed || a.Impact.OperatorFirstOrderProof || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.OperatorYSymbolicNotYukawa || !a.Firewalls.NoNumericalYukawa || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoFullUnbrokenAFTheorem || !a.Firewalls.AForientNotFullAF || !a.Firewalls.NoFullFirstOrder || !a.Firewalls.NoCompleteJOpposite || !a.Firewalls.NoBimoduleProof || !a.Firewalls.ScalarSocketSupportOnly || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall flags inconsistent")
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate859Inherited, StatusYOperatorDefined, StatusColorEdgesScalar, StatusLeptonEdgeTrivial, StatusPunctureZero, StatusDRebuilt, StatusRankKernelRecomputed, StatusFirstOrderCentrality, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict, SupportOperatorYFirstOrderIfScalar, SupportColorCentralityRepairsM3, SupportLeftKernelPersists, SupportNonzeroSocketsRankLedger, SupportPunctureStillZero, SupportStabilizerOperatorSeal, FailureOperatorYSymbolicNotYukawa, FailureNoNumericalYukawa, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoFullUnbrokenAFTheorem, FailureAForientNotFullAF, FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof, FailureNoBimoduleCommutantProof, FailureScalarSocketSupportOnly, FailureNoOperatorYukawaTheorem, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureZeroSocketBranchEnlargesKer}
}

func containsAll(haystack, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if h == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g frozen=%v alpha_native=%v R3=%v R4=%v", l.AlphaB, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatEdges(edges []EdgeOperator) string {
	parts := []string{}
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s->%s op=%s present=%v symbolic=%v scalar=%v identity=%v magnitude=%v", e.Name, e.Domain, e.Codomain, e.OperatorForm, e.Present, e.CoefficientSymbolic, e.ScalarOnVisibleFactor, e.IdentityOnVisibleFactor, e.YukawaMagnitude))
	}
	return strings.Join(parts, " | ")
}
func FormatY(y YOperator) string {
	return fmt.Sprintf("%s; %s operator=%v symbolic=%v rank_if_nonzero=%d zero_branch_enlarges_kernel=%v failures=%s", y.Expression, y.PunctureExpression, y.OperatorValued, y.SymbolicSocketMatrix, y.RankIfActiveSocketsNonzero, y.ZeroActiveSocketBranchEnlargesKernel, strings.Join(y.Failures, ","))
}
func FormatD(d DMatrix) string {
	return fmt.Sprintf("%s built_from_Y=%v self_adjoint_block=%v rank_if_nonzero=%d kernel_rank=%d kernel=%s", d.Expression, d.BuiltFromY, d.SelfAdjointByBlockForm, d.RankIfActiveSocketsNonzero, d.KernelRankIfNonzero, d.KernelSingleton)
}
func FormatFirstOrder(f FirstOrderPosition) string {
	return fmt.Sprintf("%s for %s inherited859=%v centrality=%v operator_first_order=%v failures=%s", f.Target, f.Algebra, f.Gate859Inherited, f.EdgeCentralityInstalled, f.OperatorLevelFirstOrderCertified, strings.Join(f.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s Y_operator=%v scalar_color=%v puncture_zero=%v rank_kernel=%v operator_first_order=%v alpha_sealed=%v magnitudes_missing=%v R3=%v R4=%v", i.Classification, i.OperatorYDefined, i.ScalarColorEdges, i.PunctureZero, i.RankKernelLedgerRecomputed, i.OperatorFirstOrderProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%v verdict=%s symbolic_not_yukawa=%v no_yukawa=%v no_alpha=%v no_full_AF=%v no_first_order=%v not_R3=%v not_R4=%v", f.Enforced, f.Verdict, f.OperatorYSymbolicNotYukawa, f.NoNumericalYukawa, f.NoAlphaSource, f.NoFullUnbrokenAFTheorem, f.NoFullFirstOrder, f.NotR3, f.NotR4)
}
