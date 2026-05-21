package generation2socketsectortofinitesectorliftcandidateaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SOCKET_SECTOR_TO_FINITE_SECTOR_LIFT_CANDIDATE_AUDIT"
	theoremName = "Gate 887 — SocketSectorToFiniteSector Lift Candidate Audit"
)

func Generation2SocketSectorToFiniteSectorLiftCandidateAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit strong socket trace atom domain", Passed: len(a.Domain.Atoms) == 3 && a.Domain.ActiveRank == RankHRMin && a.Domain.CompleteOnHRMin && containsAll(a.Domain.Supports, []string{SupportSocketSectorDomainStrong, SupportEdgeSupportFiniteSectorCandidate}), Detail: FormatDomain(a.Domain)},
			{Name: "audit stabilizer-sector lift candidate", Passed: routeOK(a.Lift.Routes, RouteStabilizerSectorLift, true, false, true) && a.Lift.BestRoute == RouteStabilizerSectorLift && containsAll(a.Lift.Supports, []string{SupportPostOrientationLiftCandidate, SupportR3PreparationUnderSeals}), Detail: FormatLift(a.Lift)},
			{Name: "block full unbroken A_F lift", Passed: routeOK(a.Lift.Routes, RouteFullAFLift, false, false, false) && containsAll(a.Lift.Failures, []string{FailureNoFullAFLift, FailureNoNativeSocketToFiniteSectorMap}) && containsAll(a.Domain.Failures, []string{FailureSocketAtomsNotStableFullH}), Detail: FormatLift(a.Lift)},
			{Name: "audit edge-support lift candidate without physical assignment", Passed: routeOK(a.Lift.Routes, RouteEdgeSupportLift, true, false, true) && containsAll(routeFailures(a.Lift.Routes, RouteEdgeSupportLift), []string{FailureEdgeAtomNotPhysicalSector, FailureEdgeAtomNotYukawaValue}) && !hasPhysicalLeak(a), Detail: FormatLift(a.Lift)},
			{Name: "preserve trace diagnostics under candidate lift", Passed: near(a.Domain.TraceTotal, 3+3*AlphaB) && near(a.Domain.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Domain.OperatorNEff, OperatorNEffDiagnostic) && near(a.Domain.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatDomain(a.Domain)},
			{Name: "verify classification and next frontier", Passed: a.Lift.LiftCandidate && !a.Lift.LiftCertified && !a.Lift.NativeR3 && a.Lift.NextRequired == NextBranch && containsAll(a.Lift.Supports, []string{SupportExactNextObject}), Detail: FormatLift(a.Lift)},
			{Name: "preserve official ledger freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 887 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatDomain(a.Domain), FormatLift(a.Lift), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func routeOK(routes []LiftRoute, name string, candidate, certified, postOrient bool) bool {
	for _, r := range routes {
		if r.Name == name {
			return r.LiftCandidate == candidate && r.LiftCertified == certified && r.PostOrientationOnly == postOrient && !r.PhysicalAssignment && !r.GenerationCarrierPresent && !r.FlavorOrientationPresent && !r.IndividualYukawaValues
		}
	}
	return false
}

func routeFailures(routes []LiftRoute, name string) []string {
	for _, r := range routes {
		if r.Name == name {
			return r.Failures
		}
	}
	return nil
}

func hasPhysicalLeak(a Audit) bool {
	for _, atom := range a.Domain.Atoms {
		if atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return true
		}
	}
	for _, route := range a.Lift.Routes {
		if route.PhysicalAssignment || route.GenerationCarrierPresent || route.FlavorOrientationPresent || route.IndividualYukawaValues {
			return true
		}
	}
	return false
}
