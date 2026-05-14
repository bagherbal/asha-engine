package modularhamiltonianorigin

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ModularHamiltonianOriginTrialityEnergyConstraintDerivationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-MODULAR-HAMILTONIAN-ORIGIN-TRIALITY-ENERGY-CONSTRAINT"
	const name = "Modular Hamiltonian Origin / Triality Energy Constraint Derivation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 366 audit", Passed: false, Detail: err.Error()}}}
		}
		identity := candidateByName(a.Candidates, "identity")
		mag := candidateByName(a.Candidates, "tau magnitude")
		tau := candidateByName(a.Candidates, "triality signed tau_eta")
		checks := []theorem.Check{
			{Name: "modular Hamiltonian origin criteria are formalized", Passed: a.Criteria.Formalized && a.Criteria.MustBeNative && a.Criteria.MustBeNonCircular, Detail: FormatCriteria(a.Criteria)},
			{Name: "identity Hamiltonian is native but freezes modular time", Passed: identity.Central && !identity.BreaksAllDegeneracy && !identity.PromotedNative, Detail: FormatCandidate(identity)},
			{Name: "magnitude Hamiltonian is noncentral but retains 1-2 degeneracy", Passed: !mag.Central && !mag.BreaksAllDegeneracy && !mag.PromotedNative, Detail: FormatCandidate(mag)},
			{Name: "signed tau_eta activates all frequencies but is not selected as energy", Passed: tau.NativeOperator && !tau.Central && tau.BreaksAllDegeneracy && !tau.EnergyRoleDerived && !tau.PromotedNative, Detail: FormatCandidate(tau)},
			{Name: "energy constraint inversion is circular without native expectation value", Passed: a.Energy.Formalized && a.Energy.Circular && !a.Energy.ConstraintNative && !a.Energy.PromotesKMSNative, Detail: FormatEnergy(a.Energy)},
			{Name: "flow capacity exists but vacuum remains unselected", Passed: a.Flow.Executed && a.Flow.NontrivialCapacity && a.Flow.PreservesLandscape && a.Flow.KineticSafe && !a.Flow.SelectsVacuum, Detail: FormatFlow(a.Flow)},
			{Name: "vacuum census remains unchanged", Passed: a.Census.StartingInputs == 15 && a.Census.Reduction == 0 && a.Census.RemainingInputs == 15, Detail: FormatCensus(a.Census)},
		}
		passed := 0
		for _, c := range checks {
			if c.Passed {
				passed++
			}
		}
		status := theorem.BridgeRequired
		if passed != len(checks) {
			status = theorem.FailedRoute
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks}
	}}
}
