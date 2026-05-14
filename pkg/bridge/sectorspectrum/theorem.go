package sectorspectrum

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentSectorSpectralAssignmentSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-SECTOR-SPECTRAL-ASSIGNMENT"
	const name = "current-sector spectral assignment search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct current-sector spectral assignment audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "current-sector decomposition", Passed: len(a.Sectors) == 4, Detail: "central(1), color-su3(8), B-L(1), leptoquark(6)"},
			{Name: "finite spectral carrier inventory", Passed: len(a.Carriers) > 0, Detail: FormatCarriers(a.Carriers)},
			{Name: "multiplicity audit", Passed: true, Detail: fmt.Sprintf("exact count matches=%d, singleton ambiguities=%d; %s", a.ExactCountMatches, a.PotentialScalarSingletonMatches, FormatAttempts(a.Attempts))},
			{Name: "representation-level maps", Passed: a.RepresentationMapsDerived == a.RequiredSectorAssignments, Detail: fmt.Sprintf("derived=%d/%d; no spectral carrier currently carries central/color/B-L/leptoquark action", a.RepresentationMapsDerived, a.RequiredSectorAssignments)},
			{Name: "color-sector assignment", Passed: false, Detail: "requires adjoint 8 spectral carrier; the 7 contact partial modes are a near-miss, not SU(3) adjoint data"},
			{Name: "leptoquark-sector assignment", Passed: false, Detail: "requires 6 off-diagonal carrier; no six-mode finite spectral carrier is derived"},
			{Name: "abelian-sector separation", Passed: false, Detail: "central and B-L cannot both be assigned from ambiguous singleton scalar invariants"},
			{Name: "propagator denominators", Passed: a.PropagatorDenominatorsDerived, Detail: "open; rho_A remains unassigned for every current sector"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; Gate 64 diagnostic denominator families cannot update G_hat without sector maps"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarChannelDerived, Detail: "open; representation assignment does not derive exchange sign, propagator weights, or criticality"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; no current-sector spectral carrier selects top-like up over bottom-like down"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no NJL condensation, Higgs VEV, or fermion masses are claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed mass, coupling, threshold, y_t, v, or Higgs mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
