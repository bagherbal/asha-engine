package generation2stabilizerbranchfirstorderoperatorcalculationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_STABILIZER_BRANCH_FIRST_ORDER_OPERATOR_CALCULATION_AUDIT"
	theoremName = "Gate 861 — Stabilizer-Branch First-Order Operator Calculation Audit"
)

func Generation2StabilizerBranchFirstOrderOperatorCalculationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "attempt first-order operator calculation inside A_F^orient", Passed: a.FirstOrder.Attempted && a.FirstOrder.Gate860Inherited && a.FirstOrder.Algebra == "A_F^orient = C_R plus C_H plus M_3(C)" && containsAll(a.FirstOrder.Supports, []string{StatusFirstOrderAttempted, StatusGate860Inherited}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "separate nonzero D-rho commutator as allowed one-form source", Passed: a.FirstOrder.DRhoNonzeroAllowedOneForm && containsAll(a.FirstOrder.Supports, []string{StatusNonzeroDRhoAllowed}), Detail: FormatFirstOrder(a.FirstOrder)},
			{Name: "remove opposite M3 pressure by scalar color-central edges", Passed: colorEdgesCentral(a.Edges) && a.FirstOrder.ColorCentralityInstalled && a.FirstOrder.ColorObstructionRemoved && containsAll(a.FirstOrder.Supports, []string{StatusColorObstructionRemoved, SupportColorCentralitySolvesM3}), Detail: FormatEdges(a.Edges) + " | " + FormatTerms(a.FirstOrder.Terms)},
			{Name: "preserve puncture edge zero and left kernel singleton", Passed: leptonAndPunctureOK(a.Edges) && a.FirstOrder.PunctureEdgeZeroPreserved && a.FirstOrder.LeftKernelPreserved && a.Kernel.PunctureZero && a.Kernel.KernelPreserved && a.Kernel.KernelSingleton == "h_+ tensor P_1", Detail: FormatKernel(a.Kernel)},
			{Name: "classify socket character matching as remaining orientation-seal pressure", Passed: !a.FirstOrder.SocketCharacterMatchOperatorCertified && containsAll(a.FirstOrder.Supports, []string{SupportCharacterMatchNeeded, SupportOperatorFirstOrderIfChars}) && containsAll(a.FirstOrder.Failures, []string{FailureSocketCharacterSeal}), Detail: FormatFirstOrder(a.FirstOrder) + " | " + FormatTerms(a.FirstOrder.Terms)},
			{Name: "block full unbroken/native/R3/R4 promotion", Passed: !a.FirstOrder.FullUnbrokenCompatibilityCertified && !a.FirstOrder.StabilizerOperatorCompatibilityCertified && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 861 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.NotFullUnbrokenAFTheorem && a.Firewalls.SocketCharacterSeal && a.Firewalls.NoNativeFiniteTriple && a.Firewalls.AForientNotFullAF && a.Firewalls.NoYukawaMagnitudes && a.Firewalls.NoNumericalYukawa && a.Firewalls.NoAlphaSource && a.Firewalls.NoTraceReadout && a.Firewalls.NoOfficialNEffUpdate && a.Firewalls.NoCYukawaCHiggsUpdate && a.Firewalls.NoParticleAssignment && a.Firewalls.NoNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NoR3 && a.Firewalls.NoR4 && a.Firewalls.NoFullUnbrokenOperatorTheorem && a.Firewalls.NoCompleteJProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.Verdict == StatusFirewallVerdict, Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatEdges(a.Edges), FormatFirstOrder(a.FirstOrder), FormatTerms(a.FirstOrder.Terms), FormatKernel(a.Kernel), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func colorEdgesCentral(edges []EdgeOperator) bool {
	seen := map[string]bool{}
	for _, e := range edges {
		if e.ColorEdge {
			seen[e.Name] = true
			if !e.Present || !e.CoefficientSymbolic || !e.ScalarOnVisibleFactor || !e.IdentityOnVisibleFactor || e.NumericalValue || e.YukawaMagnitude {
				return false
			}
			if e.OperatorForm != "y_+3 |h_+><e_+| tensor I_{P_3}" && e.OperatorForm != "y_-3 |h_-><e_-| tensor I_{P_3}" {
				return false
			}
		}
	}
	return seen["Y_+3"] && seen["Y_-3"]
}

func leptonAndPunctureOK(edges []EdgeOperator) bool {
	lepton, puncture := false, false
	for _, e := range edges {
		if e.LeptonEdge {
			lepton = e.Name == "Y_-1" && e.Present && e.CoefficientSymbolic && e.ScalarOnVisibleFactor && e.IdentityOnVisibleFactor && !e.NumericalValue && !e.YukawaMagnitude && e.OperatorForm == "y_-1 |h_-><e_-| tensor I_{P_1}"
		}
		if e.PunctureEdge {
			puncture = e.Name == "Y_+1" && !e.Present && !e.CoefficientSymbolic && e.OperatorForm == "0" && e.ScalarOnVisibleFactor && !e.IdentityOnVisibleFactor
		}
	}
	return lepton && puncture
}
