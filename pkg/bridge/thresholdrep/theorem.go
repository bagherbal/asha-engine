package thresholdrep

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ThresholdRepresentationAssignmentAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-THRESHOLD-REPRESENTATION-ASSIGNMENT-AUDIT"
	const name = "finite threshold representation assignment audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct threshold representation audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "threshold candidate inventory", Passed: a.CandidateCount > 0, Detail: fmt.Sprintf("candidates=%d", a.CandidateCount)},
			{Name: "scalar doublet sector assignment", Passed: a.ScalarDoubletSectorDerived, Detail: "active scalar/contact sector assigned at sector level as (1,2)_{±1/2}; individual real-mode thresholds are not selected"},
			{Name: "individual threshold-mode assignments", Passed: a.IndividualAssignedCount > 0, Detail: fmt.Sprintf("individual assigned=%d; current assignments are sector-level only", a.IndividualAssignedCount)},
			{Name: "B-sector representation assignment", Passed: a.BGapRepresentationDerived, Detail: "not derived; B-sector gap cannot enter beta thresholds yet"},
			{Name: "contact partial-overlap representation assignment", Passed: a.ContactOverlapRepresentationsDerived, Detail: "not derived; seven partial-overlap modes remain unclassified"},
			{Name: "leakage is not a threshold", Passed: !a.LeakageModeRepresentationDerived, Detail: "bare contact leakage remains a frustration invariant, not a field representation"},
			{Name: "continuum-active threshold modes", Passed: a.ContinuumActiveCount > 0, Detail: fmt.Sprintf("continuum-active derived=%d; activation/decoupling rule is still absent", a.ContinuumActiveCount)},
			{Name: "threshold-corrected beta coefficients", Passed: a.ThresholdCorrectedBetaDerived, Detail: "not derived; representation assignment is incomplete and no activation rule exists"},
			{Name: "assignment summary", Passed: true, Detail: fmt.Sprintf("sector/bridge assigned=%d, individual assigned=%d, unassigned=%d, non-threshold=%d; %s", a.SectorAssignedCount, a.IndividualAssignedCount, a.UnassignedCount, a.NonThresholdCount, FormatAssignments(a.Assignments, 8))},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("minimum missing data: %v", a.MinimumMissingData)}}
	}}
}
