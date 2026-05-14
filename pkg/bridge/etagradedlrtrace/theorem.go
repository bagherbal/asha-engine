package etagradedlrtrace

import "github.com/bagherbal/asha-engine/pkg/theorem"

func EtaGradedLeftRightTraceNoncentralHamiltonianExtractionSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-ETA-GRADED-LEFT-RIGHT-TRACE-NONCENTRAL-HAMILTONIAN-EXTRACTION-SIEVE"
	const name = "Eta-Graded Left-Right Trace / Noncentral Hamiltonian Extraction Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 369 audit", Passed: false, Detail: err.Error()}}}
		}
		laneA := laneByID(a.Lanes, "A")
		laneB := laneByID(a.Lanes, "B")
		laneC := laneByID(a.Lanes, "C")
		laneD := laneByID(a.Lanes, "D")
		checks := []theorem.Check{
			{Name: "Gate 368 target is inherited with circularity firewall", Passed: a.Inheritance.Executed && a.Inheritance.CircularityFirewall && a.Inheritance.NoEmpiricalFlavorData, Detail: FormatInheritance(a.Inheritance)},
			{Name: "eta grading is formalized on support separately from generation tau", Passed: a.Formalization.Executed && !a.Formalization.NativeEtaActsOnGeneration && len(a.Formalization.NativeSupportEta) == 2 && len(a.Formalization.GenerationEtaCandidate) == 3, Detail: FormatFormalization(a.Formalization)},
			{Name: "native support eta trace executes but remains central", Passed: laneA.Native && !laneA.Circular && laneA.Central && !laneA.Decomposition.HasNonzeroB && !laneA.BreaksFlavorOrbit, Detail: FormatLane(laneA)},
			{Name: "balanced support trace cancellation is central/zero", Passed: laneB.Native && laneB.Central && !laneB.Decomposition.HasNonzeroB && !laneB.BreaksFlavorOrbit, Detail: FormatLane(laneB)},
			{Name: "B-gap coupling does not create generation asymmetry", Passed: laneC.Native && laneC.Central && !laneC.Decomposition.HasNonzeroB && !laneC.BreaksFlavorOrbit, Detail: FormatLane(laneC)},
			{Name: "generation eta lane hits tau target but is circular", Passed: !laneD.Native && laneD.Circular && laneD.NonCentral && laneD.Decomposition.TargetReached && laneD.BreaksFlavorOrbit && !laneD.PromotedHamiltonian, Detail: FormatLane(laneD)},
			{Name: "thermal-time activation is refused without native noncentral extraction", Passed: a.Activation.Executed && !a.Activation.NativeTargetReached && a.Activation.CircularCapacityWitnessed && !a.Activation.PromotedNative && !a.Activation.InternalTimeActivated, Detail: FormatActivation(a.Activation)},
			{Name: "landscape firewalls remain preserved", Passed: a.Landscape.Executed && a.Landscape.WeakMixingPreserved && a.Landscape.QuarticRatioPreserved && a.Landscape.AlphaGUTPreserved && a.Landscape.MoritaSplitPreserved && a.Landscape.NoEmpiricalFlavorImport && !a.Landscape.FiniteCorePolluted, Detail: FormatLandscape(a.Landscape)},
			{Name: "kinetic safety is preserved", Passed: a.Kinetic.Executed && a.Kinetic.AllCandidatesSelf && a.Kinetic.FaithfulStates && a.Kinetic.NoRankCollapse && a.Kinetic.NoGhostMetric && a.Kinetic.NoNonunitaryPush, Detail: FormatKinetic(a.Kinetic)},
			{Name: "vacuum census remains unchanged", Passed: a.Census.StartingInputs == 15 && a.Census.Reduction == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealTarget, Detail: FormatCensus(a.Census)},
		}
		passed := 0
		for _, c := range checks {
			if c.Passed {
				passed++
			}
		}
		status := theorem.FailedRoute
		if passed != len(checks) {
			status = theorem.FailedRoute
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks}
	}}
}

func laneByID(lanes []TraceLane, id string) TraceLane {
	for _, lane := range lanes {
		if lane.Lane == id {
			return lane
		}
	}
	return lanes[0]
}
