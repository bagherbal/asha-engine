package familyboundarysourceaxiom

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FamilyBoundaryConditionSectorSourceAxiomMinimalitySieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Family boundary condition / sector source axiom minimality sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate415 audit", Passed: false, Detail: err.Error()}}}
		}
		var sectorOK, observedRejected, z3Under bool
		for _, c := range a.Candidates {
			if c.Name == "charge-sector source boundary" && c.CKMCapacity && c.PMNSCapacity && c.EmpiricalIndependent && !c.NativeToCurrentAsha && !c.FixesCoefficientValues {
				sectorOK = true
			}
			if c.Name == "observed Yukawa matrix source" && c.ImportsObservedYukawa && c.FixesCoefficientValues && !c.EmpiricalIndependent && !c.PromotedToTheorem {
				observedRejected = true
			}
			if c.Name == "Z3 Weyl phase source" && c.CKMCapacity && !c.FixesCoefficientValues && c.FreeRealParameters > 0 {
				z3Under = true
			}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate414 coefficient boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate414NoCoefficientSelector && a.Inheritance.Gate414CoefficientsFree && a.Inheritance.ChargedModuliDim == Gate372ChargedFlavorModuliDim, Detail: FormatInheritance(a.Inheritance)},
			{Name: "boundary/source axiom arena formalized", Passed: a.Arena.Executed && a.Arena.BaselineCoefficientRays > 0 && a.Arena.GaugeCompatibleIfFamily && !a.Arena.NativeSelectorPresent && !a.Arena.EmpiricalYukawaImported, Detail: FormatArena(a.Arena)},
			{Name: "candidate axiom ledger compiled", Passed: len(a.Candidates) >= 5 && sectorOK && observedRejected && z3Under, Detail: RenderCandidateSummary(a.Candidates)},
			{Name: "minimality ranking identifies least-cost CKM-capable axiom but keeps it quarantined", Passed: a.Ranking.Executed && a.Ranking.LeastCostName == "charge-sector source boundary" && a.Ranking.LeastCost == 2 && a.Ranking.LeastCostStillAxiom && a.Ranking.LeastCostCKMCapacity && !a.Ranking.LeastCostFixesAngles && a.Ranking.NoCandidateNative, Detail: FormatRanking(a.Ranking)},
			{Name: "CKM/PMNS capacity remains conditional", Passed: a.Capacity.Executed && a.Capacity.ConditionalCKMAvailable && a.Capacity.ConditionalPMNSAvailable && !a.Capacity.AnyCandidateFixesAngles && !a.Capacity.AnyCandidateNative && a.Capacity.AnyCandidateCurveFitting, Detail: FormatCapacity(a.Capacity)},
			{Name: "moduli firewall remains native", Passed: a.Moduli.StartDim == Gate372ChargedFlavorModuliDim && a.Moduli.BestNativeDim == Gate372ChargedFlavorModuliDim && !a.Moduli.NativeReductionBelow13 && a.Moduli.ConditionalMixingCapacity && a.Moduli.CoefficientsRemainFree && a.Moduli.FirewallPreserved, Detail: FormatModuli(a.Moduli)},
			{Name: "empirical firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaMatricesInserted && a.Firewall.AxiomsQuarantined && a.Firewall.NoNativeDerivationClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate targets minimal sector-source consistency", Passed: a.Next.Gate == 416 && a.Next.Title == "Minimal Sector-Source Axiom Consistency / Parameter-Counting Sieve", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
