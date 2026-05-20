package generation2covariantphasespacecliffordsourcerouterandyukawagapaccelerationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-824-COVARIANT-PHASE-SPACE-CLIFFORD-SOURCE-ROUTER-YUKAWA-GAP-ACCELERATION"
	theoremName = "Gate 824 — Covariant Phase-Space Clifford Source-Router and Yukawa-Gap Acceleration Audit"
)

func Generation2CovariantPhaseSpaceCliffordSourceRouterAndYukawaGapAccelerationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Notes: []string{err.Error()}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 823 DATA_REQUIRED and boundary-FN pressure", Passed: strings.Contains(a.Inherited.Gate823Status, "DATA_REQUIRED") && math.Abs(a.Inherited.AlphaB-0.0003878958469680527) < 1e-18 && !a.Impact.CanUpdate, Detail: FormatInherited(a.Inherited)},
			{Name: "execute real-form chirality precheck", Passed: a.RealForm.OmegaSquared == "-1" && !a.RealForm.NaiveProjectorsIdempotent && a.RealForm.RequiresAirlock && containsAll(a.RealForm.Failures, []string{FailureNaiveRealChirality}), Detail: FormatRealForm(a.RealForm)},
			{Name: "route master claim inventory", Passed: len(a.Router.Claims) == 12 && strings.Contains(FormatClaims(a.Router.Claims), "SO(8)/D4 triality") && strings.Contains(FormatClaims(a.Router.Claims), "nu_R dark-sector candidate"), Detail: FormatClaims(a.Router.Claims)},
			{Name: "reject state/gauge/edge container as trace-magnitude source", Passed: containsAll(a.Router.Failures, []string{FailureContainerNotYukawa, FailureStateInventoryNoTrace, FailureMassBridgeNotHierarchy}) && containsAll(a.Router.Supports, []string{SupportContainerOnly}), Detail: strings.Join(a.Router.Failures, "; ")},
			{Name: "preserve triality and chirality firewalls", Passed: containsAll(a.Router.Failures, []string{FailureTrialityNoReadout, FailureChiralityNotYukawa}) && strings.Contains(FormatClaims(a.Router.Claims), "airlocked triality candidate"), Detail: FormatClaims(a.Router.Claims)},
			{Name: "audit boundary-FN source routing", Passed: !a.BoundaryFN.CanMoveBeyondPartialR2 && containsAll(a.BoundaryFN.Failures, []string{FailureBoundaryNoPositiveSpectra, FailureNoBoundaryMapFound}) && strings.Contains(FormatBoundaryFN(a.BoundaryFN), "BoundaryToTraceMagnitudeRestMap"), Detail: FormatBoundaryFN(a.BoundaryFN)},
			{Name: "select partial source-container outcome", Passed: a.Outcome.Outcome == OutcomePartialContainer && strings.Contains(a.Outcome.NextGate, "BoundaryToTraceMagnitudeRestMap") && containsAll(a.Outcome.Failures, []string{FailureNoPositiveTraceReadoutFound}), Detail: FormatOutcome(a.Outcome)},
			{Name: "preserve C_Yukawa and C_Higgs", Passed: !a.Impact.CanUpdate && math.Abs(a.Impact.OfficialCYukawa-CYukawa) < 1e-18 && math.Abs(a.Impact.OfficialCHiggs-CHiggs) < 1e-18 && containsAll(a.Impact.Failures, []string{FailureNoCYukawaUpdate, FailureCHiggsLevelB}), Detail: FormatImpact(a.Impact)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.ContainerNotSpectrum && a.Firewalls.StatesNotValues && a.Firewalls.TrialityNotGeneration && a.Firewalls.ChiralityNotHierarchy && a.Firewalls.MassBridgeNotValues && a.Firewalls.LabelsNotTraceAtoms && a.Firewalls.LagrangianNotNEff && a.Firewalls.DarkNotNuLedger && a.Firewalls.NoCYukawaUpdate && a.Firewalls.Verdict == StatusFirewallGate824, Detail: a.Firewalls.Verdict},
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
		notes := []string{a.Truth, FormatInherited(a.Inherited), FormatRealForm(a.RealForm), FormatClaims(a.Router.Claims), FormatBoundaryFN(a.BoundaryFN), FormatOutcome(a.Outcome), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func containsAll(haystack []string, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}
