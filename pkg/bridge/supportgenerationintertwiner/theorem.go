package supportgenerationintertwiner

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SupportToGenerationIntertwinerTopologicalIndexMapSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-SUPPORT-TO-GENERATION-INTERTWINER-TOPOLOGICAL-INDEX-MAP-SIEVE"
	const name = "Support-to-Generation Intertwiner / Topological Index Map Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 370 audit", Passed: false, Detail: err.Error()}}}
		}
		laneA := candidateByID(a.Candidates, "A")
		laneB := candidateByID(a.Candidates, "B")
		laneC := candidateByID(a.Candidates, "C")
		laneD := candidateByID(a.Candidates, "D")
		laneE := candidateByID(a.Candidates, "E")
		laneF := candidateByID(a.Candidates, "F")
		checks := []theorem.Check{
			{Name: "Gate 369 obstruction is inherited", Passed: a.Inheritance.Executed && a.Inheritance.NativeEtaTraceCentral && a.Inheritance.TauEtaInsertionCircular && a.Inheritance.NoEmpiricalFlavorData, Detail: FormatInheritance(a.Inheritance)},
			{Name: "support-to-generation intertwiner sieve is formalized", Passed: a.Formalization.Executed && a.Formalization.TargetFormula != "" && len(a.Formalization.NativeAdmissibility) >= 3 && len(a.Formalization.ForbiddenMoves) >= 3, Detail: FormatFormalization(a.Formalization)},
			{Name: "identity broadcast factors through I3", Passed: laneA.Native && !laneA.Circular && laneA.U3Equivariant && !laneA.GenerationAddressed && laneA.Central && !laneA.Decomposition.HasNonzeroB && !laneA.Promotable, Detail: FormatCandidate(laneA)},
			{Name: "Omega_Hsigma endpoint remains support-index only", Passed: laneB.Native && laneB.U3Equivariant && !laneB.GenerationAddressed && laneB.Central && !laneB.Promotable, Detail: FormatCandidate(laneB)},
			{Name: "finite Dirac/J transport is generation-equivariant", Passed: laneC.Native && laneC.U3Equivariant && !laneC.GenerationAddressed && laneC.Central && !laneC.Promotable, Detail: FormatCandidate(laneC)},
			{Name: "Morita multiplicity broadcasts uniformly", Passed: laneD.Native && laneD.U3Equivariant && !laneD.GenerationAddressed && laneD.Central && !laneD.Promotable, Detail: FormatCandidate(laneD)},
			{Name: "B-gap scaling does not create a generation address", Passed: laneE.Native && laneE.U3Equivariant && laneE.Central && !laneE.Decomposition.HasNonzeroB && !laneE.Promotable, Detail: FormatCandidate(laneE)},
			{Name: "tau_eta map is noncentral but circular", Passed: !laneF.Native && laneF.Circular && laneF.GenerationAddressed && laneF.NonCentral && laneF.Decomposition.TargetReached && laneF.BreaksFlavorOrbit && !laneF.Promotable, Detail: FormatCandidate(laneF)},
			{Name: "equivariance no-go is exposed", Passed: a.NoGo.Executed && a.NoGo.AllNativeMapsFactorThroughI3 && a.NoGo.NativeNoncentralCount == 0 && a.NoGo.NativeGenerationAddressCount == 0 && a.NoGo.CircularNoncentralWitnessCount == 1, Detail: FormatNoGo(a.NoGo)},
			{Name: "thermal activation is refused", Passed: a.Activation.Executed && !a.Activation.NativeIntertwinerDerived && !a.Activation.InternalThermalTimeActivated && !a.Activation.TauEtaHamiltonianSelected && a.Activation.CircularCapacityWitnessed, Detail: FormatActivation(a.Activation)},
			{Name: "landscape firewalls remain preserved", Passed: a.Landscape.Executed && a.Landscape.WeakMixingPreserved && a.Landscape.QuarticRatioPreserved && a.Landscape.AlphaGUTPreserved && a.Landscape.MoritaSplitPreserved && a.Landscape.NoEmpiricalFlavorImport && !a.Landscape.FiniteCorePolluted, Detail: FormatLandscape(a.Landscape)},
			{Name: "kinetic safety is preserved", Passed: a.Kinetic.Executed && a.Kinetic.AllCandidatesSelf && a.Kinetic.NoNonunitaryPush && a.Kinetic.NoRankCollapse && a.Kinetic.NoGhostMetric && a.Kinetic.FaithfulCarrier, Detail: FormatKinetic(a.Kinetic)},
			{Name: "vacuum census remains unreduced", Passed: a.Census.StartingInputs == 15 && a.Census.Reduction == 0 && a.Census.RemainingInputs == 15 && !a.Census.SevenSealTarget, Detail: FormatCensus(a.Census)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks}
	}}
}

func candidateByID(candidates []IntertwinerCandidate, id string) IntertwinerCandidate {
	for _, c := range candidates {
		if c.Lane == id {
			return c
		}
	}
	return candidates[0]
}
