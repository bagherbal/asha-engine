package generation2phaseorientation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SignedCycleComplexPhaseOrientationSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 signed-cycle and complex phase orientation sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate446 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate445 triangle topology and sealed orientation", Passed: a.Inheritance.Executed && a.Inheritance.Gate445KGenForced && a.Inheritance.Gate445XSupportForced && a.Inheritance.Gate445AmplitudeSealed && a.Inheritance.Gate445SignedOrientationSealed && a.Inheritance.Gate445NoEmpiricalMasses, Detail: FormatInheritance(a.Inheritance)},
			{Name: "Hermitian triangular cycle arena formalized", Passed: a.Arena.Executed && a.Arena.Hermitian && a.Arena.ZeroDiagonal && a.Arena.TriangleSupportInherited && a.Arena.VertexRephasingAllowed && !a.Arena.EmpiricalDataImported, Detail: FormatArena(a.Arena)},
			{Name: "all orientation boundaries applied", Passed: len(a.Boundaries) == 5 && a.Boundaries[0].Passed && a.Boundaries[1].Passed && a.Boundaries[2].Passed && a.Boundaries[3].Passed && a.Boundaries[4].Passed, Detail: FormatBoundary(a.Boundaries[0]) + " | " + FormatBoundary(a.Boundaries[1]) + " | " + FormatBoundary(a.Boundaries[2]) + " | " + FormatBoundary(a.Boundaries[3]) + " | " + FormatBoundary(a.Boundaries[4])},
			{Name: "real signed sieve remains non-unique", Passed: a.RealSieve.Executed && !a.RealSieve.UniqueSignedCycle && a.RealSieve.PositiveCycleCount == 4 && a.RealSieve.NegativeCycleCount == 4 && a.RealSieve.Z2GaugeClasses == 2, Detail: FormatRealSignSieve(a.RealSieve)},
			{Name: "complex cycle phase continuum survives", Passed: a.PhaseSieve.Executed && a.PhaseSieve.ContinuumSurvives && a.PhaseSieve.CPConjugatePairsSurvive && !a.PhaseSieve.UniqueComplexPhase && !a.PhaseSieve.CPPhaseValuePredicted, Detail: FormatComplexPhaseSieve(a.PhaseSieve)},
			{Name: "orientation not promoted as native axiom", Passed: a.Conclusion.Executed && a.Conclusion.XSupportTopologyPreserved && !a.Conclusion.SignedCycleForced && !a.Conclusion.ComplexPhaseForced && !a.Conclusion.YGenPromotedToNative && !a.Conclusion.CPViolationPredicted, Detail: FormatConclusion(a.Conclusion)},
			{Name: "empirical flavor firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.YGenRemainsQuarantined && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimStillFree == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate audits coefficient source ledger", Passed: a.Next.Gate == 447 && a.Next.Title == "Sector-Coefficient Source Ledger / Amplitude Firewall Closure", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusFailedSignedOrientationNotUnique, StatusFailedComplexPhaseContinuum, StatusFailedYGenNotNative, a.Truth}}
	}}
}
