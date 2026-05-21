// Package generation2stabilizerbranchfirstordersupportedgecentralityaudit implements
// Gate 859: Stabilizer-Branch First-Order Support Calculation and Edge-Centrality Audit.
//
// Gate 859 follows Gate 858's support-level oriented bimodule/order-zero seal.
// It audits the first-order support target inside A_F^orient=C_R plus C_H plus
// M_3(C), while separating the allowed one-form commutator [D_F,rho_F(a)] from
// the true first-order obstruction [[D_F,rho_F(a)],rho_F^op(b)].  The new
// pressure point is edge centrality: the color edges must be scalar/identity on
// the P_3 color factor and the lepton edge must be scalar on the P_1 factor.
// The result remains support-level only: no operator-level J-opposite proof, no
// first-order theorem, no Yukawa magnitudes, no alpha_B source, no R3/R4, and no
// official ledger update are certified.
package generation2stabilizerbranchfirstordersupportedgecentralityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE859-STABILIZER-BRANCH-FIRST-ORDER-SUPPORT-EDGE-CENTRALITY-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank          = 1
	P3Rank          = 3
	WRank           = P1Rank + P3Rank
	WeakDoubletRank = 2
	HLRank          = WeakDoubletRank * WRank
	HRMinRank       = 7
	HPartMinRank    = HLRank + HRMinRank
	HFMinRank       = 2 * HPartMinRank
	DSymRank        = 14
	KernelRank      = HPartMinRank - DSymRank

	StatusGate858Inherited         = "PASS_GATE858_ORDER_ZERO_SUPPORT_BIMODULE_INHERITED"
	StatusFirstOrderSupportAudited = "PASS_FIRST_ORDER_SUPPORT_COMMUTATOR_AUDITED"
	StatusDRhoOneFormSeparated     = "PASS_NONZERO_D_RHO_CLASSIFIED_AS_ONE_FORM_SOURCE"
	StatusEdgeCentralityDefined    = "PASS_COLOR_EDGE_CENTRALITY_REQUIREMENT_DEFINED"
	StatusLeptonTrivialityAudited  = "PASS_LEPTON_EDGE_TRIVIALITY_AUDITED"
	StatusPunctureZeroPreserved    = "PASS_PUNCTURE_EDGE_ZERO_PRESERVED"
	StatusKernelPreserved          = "PASS_LEFT_KERNEL_SINGLETON_PRESERVED_AT_SUPPORT_LEVEL"
	StatusNoObservedDataUsed       = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen             = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE859_SUPPORT_FIRST_ORDER_EDGE_CENTRALITY_NOT_R3"

	SupportFirstOrderIfCentral       = "CONDITIONAL_SUPPORT_FIRST_ORDER_SUPPORT_COMPATIBILITY_IF_COLOR_EDGES_ARE_SCALAR_ON_P3"
	SupportColorEdgesCentral         = "CONDITIONAL_SUPPORT_COLOR_EDGES_ARE_CENTRAL_ON_P3_FACTOR"
	SupportYPlusMinusIdentityOnColor = "CONDITIONAL_SUPPORT_Y_PLUS3_AND_Y_MINUS3_MUST_BE_IDENTITY_ON_COLOR_FACTOR"
	SupportLeptonEdgeP1Compatible    = "CONDITIONAL_SUPPORT_LEPTON_EDGE_COMPATIBLE_ON_P1"
	SupportNonzeroDRhoOneForm        = "CONDITIONAL_SUPPORT_NONZERO_D_RHO_COMMUTATOR_IS_ALLOWED_ONE_FORM_SOURCE"
	SupportPunctureKernelStable      = "CONDITIONAL_SUPPORT_PUNCTURE_AND_KERNEL_STABLE_IN_ORIENTED_BRANCH_AT_SUPPORT_LEVEL"
	SupportStabilizerSupportFirst    = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_STABILIZER_FIRST_ORDER_SUPPORT_SEAL"

	FailureAForientNotFullAF             = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_UNBROKEN_A_F"
	FailureNoFullOperatorFirstOrder      = "FAILED_ROUTE_NO_FULL_OPERATOR_LEVEL_FIRST_ORDER_THEOREM"
	FailureNoCompleteJOppositeProof      = "FAILED_ROUTE_NO_COMPLETE_J_OPPOSITE_OPERATOR_PROOF"
	FailureNoBimoduleCommutantProof      = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureEdgeCentralitySupportOnly     = "FAILED_ROUTE_EDGE_CENTRALITY_IS_SUPPORT_SEAL_NOT_NATIVE_YUKAWA_THEOREM"
	FailureCharacterMatchSupportOnly     = "FAILED_ROUTE_CHARACTER_MATCH_REMAINS_SUPPORT_LABEL_NOT_OPERATOR_INTERTWINER"
	FailureSupportIntertwinerNotOperator = "FAILED_ROUTE_SUPPORT_LEVEL_INTERTWINER_NOT_OPERATOR_THEOREM"
	FailureSymbolicYNotMagnitude         = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoAlphaSource                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceReadout                = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                         = "FAILED_ROUTE_R2_SUPPORT_BIMODULE_FIRST_ORDER_SUPPORT_NOT_R3"
	FailureNotR4                         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment          = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem             = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem      = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type AlgebraLayer struct {
	Algebra                                    string
	FullUnbrokenAlgebra                        string
	PostOrientation, ContainsCR, ContainsCH    bool
	ContainsM3C, ContainsFullH                 bool
	SupportPreservesSockets, SupportPreservesP bool
	Supports, Failures                         []string
}

type FirstOrderSupport struct {
	Expression                         string
	Algebra                            string
	OrderZeroInherited                 bool
	SupportAuditable                   bool
	DRhoCommutatorAllowedOneFormSource bool
	FirstOrderSupportConditionAudited  bool
	OperatorTheoremCertified           bool
	Supports, Failures                 []string
}

type EdgeCentrality struct {
	Name, Domain, Codomain, Factor, RequiredForm  string
	Rank                                          int
	CentralOnFactor, IdentityOnFactor             bool
	ColorEdge, LeptonEdge, PunctureEdge           bool
	Present                                       bool
	OperatorIntertwinerCertified, YukawaMagnitude bool
	Supports, Failures                            []string
}

type PunctureKernel struct {
	PunctureEdge                                   string
	PunctureCoefficientZero, PunctureReintroduced  bool
	RightPunctureOutsideMinimal, LeftKernelPresent bool
	LeftKernel                                     string
	KernelRank                                     int
	Supports, Failures                             []string
}

type Impact struct {
	Classification                                                                             string
	Gate858Inherited, FirstOrderSupportAudited, EdgeCentralitySupport, PunctureKernelPreserved bool
	OperatorFirstOrderProof, OperatorJOppositeProof, NativeFiniteTripleProof                   bool
	AlphaStillSealed, MagnitudesStillMissing                                                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4           bool
}

type Firewalls struct {
	Enforced                                                                                               bool
	AForientNotFullAF, NoFullFirstOrder, NoCompleteJOpposite, NoBimoduleProof, EdgeCentralitySupportOnly   bool
	CharacterMatchSupportOnly, SupportIntertwinerNotOperator, YSymbolicOnly, NoAlphaSource, NoTraceReadout bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem     bool
	NoThreeGenerationTheorem                                                                               bool
	Verdict                                                                                                string
}

type Audit struct {
	ID             string
	Ledger         Ledger
	Algebra        AlgebraLayer
	FirstOrder     FirstOrderSupport
	Edges          []EdgeCentrality
	PunctureKernel PunctureKernel
	Impact         Impact
	Firewalls      Firewalls
	Truth, Final   string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Algebra: AlgebraLayer{
			Algebra: "A_F^orient=C_R plus C_H plus M_3(C)", FullUnbrokenAlgebra: "A_F=C plus H plus M_3(C)",
			PostOrientation: true, ContainsCR: true, ContainsCH: true, ContainsM3C: true, ContainsFullH: false,
			SupportPreservesSockets: true, SupportPreservesP: true,
			Supports: []string{StatusGate858Inherited, SupportStabilizerSupportFirst},
			Failures: []string{FailureAForientNotFullAF},
		},
		FirstOrder: FirstOrderSupport{
			Expression: "[[D_F^sym,rho_F(a)],rho_F^op(b)]=0", Algebra: "a,b in A_F^orient",
			OrderZeroInherited: true, SupportAuditable: true, DRhoCommutatorAllowedOneFormSource: true, FirstOrderSupportConditionAudited: true, OperatorTheoremCertified: false,
			Supports: []string{StatusFirstOrderSupportAudited, StatusDRhoOneFormSeparated, SupportFirstOrderIfCentral, SupportNonzeroDRhoOneForm},
			Failures: []string{FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof, FailureNoBimoduleCommutantProof, FailureSupportIntertwinerNotOperator},
		},
		Edges: []EdgeCentrality{
			{Name: "Y_+3", Domain: "e_+ tensor P_3", Codomain: "h_+ tensor P_3", Factor: "P_3 color", RequiredForm: "y_+3 I_{P_3}", Rank: P3Rank, CentralOnFactor: true, IdentityOnFactor: true, ColorEdge: true, Present: true, Supports: []string{StatusEdgeCentralityDefined, SupportColorEdgesCentral, SupportYPlusMinusIdentityOnColor}, Failures: []string{FailureEdgeCentralitySupportOnly, FailureCharacterMatchSupportOnly, FailureSymbolicYNotMagnitude}},
			{Name: "Y_-3", Domain: "e_- tensor P_3", Codomain: "h_- tensor P_3", Factor: "P_3 color", RequiredForm: "y_-3 I_{P_3}", Rank: P3Rank, CentralOnFactor: true, IdentityOnFactor: true, ColorEdge: true, Present: true, Supports: []string{StatusEdgeCentralityDefined, SupportColorEdgesCentral, SupportYPlusMinusIdentityOnColor}, Failures: []string{FailureEdgeCentralitySupportOnly, FailureCharacterMatchSupportOnly, FailureSymbolicYNotMagnitude}},
			{Name: "Y_-1", Domain: "e_- tensor P_1", Codomain: "h_- tensor P_1", Factor: "P_1 lepton", RequiredForm: "y_-1 I_{P_1}", Rank: P1Rank, CentralOnFactor: true, IdentityOnFactor: true, LeptonEdge: true, Present: true, Supports: []string{StatusLeptonTrivialityAudited, SupportLeptonEdgeP1Compatible}, Failures: []string{FailureEdgeCentralitySupportOnly, FailureCharacterMatchSupportOnly, FailureSymbolicYNotMagnitude}},
			{Name: "Y_+1", Domain: "e_+ tensor P_1", Codomain: "h_+ tensor P_1", Factor: "P_1 lepton", RequiredForm: "0", Rank: P1Rank, CentralOnFactor: true, IdentityOnFactor: false, PunctureEdge: true, Present: false, Supports: []string{StatusPunctureZeroPreserved, SupportPunctureKernelStable}, Failures: []string{FailureSupportIntertwinerNotOperator}},
		},
		PunctureKernel: PunctureKernel{
			PunctureEdge: "Y_+1:e_+ tensor P_1 -> h_+ tensor P_1", PunctureCoefficientZero: true, PunctureReintroduced: false,
			RightPunctureOutsideMinimal: true, LeftKernelPresent: true, LeftKernel: "h_+ tensor P_1", KernelRank: KernelRank,
			Supports: []string{StatusPunctureZeroPreserved, StatusKernelPreserved, SupportPunctureKernelStable},
			Failures: []string{FailureNoNeutrinoTheorem, FailureNoParticleAssignment},
		},
		Impact: Impact{
			Classification: "R2+++++_stabilizer_first_order_support_edge_centrality_seal", Gate858Inherited: true, FirstOrderSupportAudited: true, EdgeCentralitySupport: true, PunctureKernelPreserved: true,
			OperatorFirstOrderProof: false, OperatorJOppositeProof: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true,
		},
		Firewalls: Firewalls{Enforced: true, AForientNotFullAF: true, NoFullFirstOrder: true, NoCompleteJOpposite: true, NoBimoduleProof: true, EdgeCentralitySupportOnly: true, CharacterMatchSupportOnly: true, SupportIntertwinerNotOperator: true, YSymbolicOnly: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict},
		Truth:     "Gate 859 audits first-order support inside A_F^orient and identifies edge centrality as the pressure point: color edges must be scalar identity maps on P_3, the lepton edge must be scalar on P_1, and Y_+1 remains zero.",
		Final:     "The stabilizer branch is support-first-order compatible only if Y_+3 and Y_-3 are central on the color factor and Y_-1 is scalar on the lepton factor. This is still a support seal, not an operator theorem or magnitude source.",
	}
	return a, a.Validate()
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger overpromoted")
	}
	if !a.Algebra.PostOrientation || !a.Algebra.ContainsCR || !a.Algebra.ContainsCH || !a.Algebra.ContainsM3C || a.Algebra.ContainsFullH || !a.Algebra.SupportPreservesSockets || !a.Algebra.SupportPreservesP {
		return err("algebra layer inconsistent")
	}
	if !a.FirstOrder.OrderZeroInherited || !a.FirstOrder.SupportAuditable || !a.FirstOrder.DRhoCommutatorAllowedOneFormSource || !a.FirstOrder.FirstOrderSupportConditionAudited || a.FirstOrder.OperatorTheoremCertified {
		return err("first-order support overpromoted or inconsistent")
	}
	if len(a.Edges) != 4 {
		return err("expected three active edges plus puncture edge")
	}
	active := 0
	for _, e := range a.Edges {
		if e.Present {
			active++
		}
		if e.YukawaMagnitude || e.OperatorIntertwinerCertified {
			return err("edge overpromoted")
		}
		if (e.ColorEdge || e.LeptonEdge) && (!e.CentralOnFactor || !e.IdentityOnFactor || !e.Present) {
			return err("active edge centrality inconsistent")
		}
		if e.PunctureEdge && (e.Present || e.IdentityOnFactor) {
			return err("puncture edge must stay zero")
		}
	}
	if active != 3 {
		return err("expected three active support edges")
	}
	if !a.PunctureKernel.PunctureCoefficientZero || a.PunctureKernel.PunctureReintroduced || !a.PunctureKernel.RightPunctureOutsideMinimal || !a.PunctureKernel.LeftKernelPresent || a.PunctureKernel.KernelRank != KernelRank {
		return err("puncture/kernel flags inconsistent")
	}
	if !a.Impact.Gate858Inherited || !a.Impact.FirstOrderSupportAudited || !a.Impact.EdgeCentralitySupport || !a.Impact.PunctureKernelPreserved || a.Impact.OperatorFirstOrderProof || a.Impact.OperatorJOppositeProof || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AForientNotFullAF || !a.Firewalls.NoFullFirstOrder || !a.Firewalls.NoCompleteJOpposite || !a.Firewalls.NoBimoduleProof || !a.Firewalls.EdgeCentralitySupportOnly || !a.Firewalls.CharacterMatchSupportOnly || !a.Firewalls.SupportIntertwinerNotOperator || !a.Firewalls.YSymbolicOnly || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall flags inconsistent")
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate858Inherited, StatusFirstOrderSupportAudited, StatusDRhoOneFormSeparated, StatusEdgeCentralityDefined, StatusLeptonTrivialityAudited, StatusPunctureZeroPreserved, StatusKernelPreserved, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict, SupportFirstOrderIfCentral, SupportColorEdgesCentral, SupportYPlusMinusIdentityOnColor, SupportLeptonEdgeP1Compatible, SupportNonzeroDRhoOneForm, SupportPunctureKernelStable, SupportStabilizerSupportFirst, FailureAForientNotFullAF, FailureNoFullOperatorFirstOrder, FailureNoCompleteJOppositeProof, FailureNoBimoduleCommutantProof, FailureEdgeCentralitySupportOnly, FailureCharacterMatchSupportOnly, FailureSupportIntertwinerNotOperator, FailureSymbolicYNotMagnitude, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem}
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
func allNotes(parts ...[]string) []string {
	out := []string{}
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g frozen=%v alpha_native=%v R3=%v R4=%v", l.AlphaB, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatAlgebra(a AlgebraLayer) string {
	return fmt.Sprintf("%s post_orientation=%v contains_full_H=%v preserves_sockets=%v preserves_P1P3=%v failures=%s", a.Algebra, a.PostOrientation, a.ContainsFullH, a.SupportPreservesSockets, a.SupportPreservesP, strings.Join(a.Failures, ","))
}
func FormatFirstOrder(f FirstOrderSupport) string {
	return fmt.Sprintf("%s for %s order_zero_inherited=%v support_auditable=%v d_rho_one_form=%v operator_theorem=%v failures=%s", f.Expression, f.Algebra, f.OrderZeroInherited, f.SupportAuditable, f.DRhoCommutatorAllowedOneFormSource, f.OperatorTheoremCertified, strings.Join(f.Failures, ","))
}
func FormatEdges(edges []EdgeCentrality) string {
	parts := []string{}
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s->%s form=%s present=%v central=%v identity=%v operator=%v magnitude=%v", e.Name, e.Domain, e.Codomain, e.RequiredForm, e.Present, e.CentralOnFactor, e.IdentityOnFactor, e.OperatorIntertwinerCertified, e.YukawaMagnitude))
	}
	return strings.Join(parts, " | ")
}
func FormatPunctureKernel(p PunctureKernel) string {
	return fmt.Sprintf("puncture=%s zero=%v reintroduced=%v left_kernel=%s kernel_rank=%d", p.PunctureEdge, p.PunctureCoefficientZero, p.PunctureReintroduced, p.LeftKernel, p.KernelRank)
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s first_order_support=%v edge_centrality=%v operator_first_order=%v alpha_sealed=%v magnitudes_missing=%v R3=%v R4=%v", i.Classification, i.FirstOrderSupportAudited, i.EdgeCentralitySupport, i.OperatorFirstOrderProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%v verdict=%s no_full_first_order=%v no_J_op=%v edge_centrality_support_only=%v not_R3=%v not_R4=%v", f.Enforced, f.Verdict, f.NoFullFirstOrder, f.NoCompleteJOpposite, f.EdgeCentralitySupportOnly, f.NotR3, f.NotR4)
}
