package sectoroperators

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CurrentSectorOperatorConstructionSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CURRENT-SECTOR-OPERATOR-CONSTRUCTION"
	const name = "current-sector operator construction search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct current-sector operators", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "operator-construction domain", Passed: a.FlavorDimension == 4, Detail: "one-generation flavor space is 1 lepton + 3 color seeds"},
			{Name: "current generator inventory", Passed: len(a.Generators) == 16, Detail: "u(4)-shaped inventory: central 1 + color-su3 8 + B-L 1 + leptoquark 6"},
			{Name: "sector Casimir operators", Passed: a.SectorOperatorsConstructed && a.AllExpectedSectorsConstructed, Detail: FormatOperators(a.Operators)},
			{Name: "representation-level maps", Passed: a.RepresentationLevelMapsDerived, Detail: "constructed C_A=sum_a T_a^T T_a for every current sector, replacing raw spectral-list matching"},
			{Name: "operator positivity", Passed: a.OperatorsPositive, Detail: "all sector Casimir operators are positive semidefinite in this finite representation"},
			{Name: "flavor-basis diagonal audit", Passed: a.OperatorsDiagonalInFlavorBasis, Detail: "sector Casimirs are diagonal on the lepton/color seed basis; this exposes their charge support"},
			{Name: "color Casimir carrier", Passed: a.ColorCasimirValue > 0, Detail: fmt.Sprintf("color-su3 C has spectrum [0, %.10f, %.10f, %.10f] on [lepton,color1,color2,color3]", a.ColorCasimirValue, a.ColorCasimirValue, a.ColorCasimirValue)},
			{Name: "leptoquark Casimir carrier", Passed: a.LeptoquarkLeptonValue > a.LeptoquarkColorValue, Detail: fmt.Sprintf("leptoquark C spectrum [%.10f, %.10f, %.10f, %.10f]", a.LeptoquarkLeptonValue, a.LeptoquarkColorValue, a.LeptoquarkColorValue, a.LeptoquarkColorValue)},
			{Name: "abelian separation carrier", Passed: a.BLLeptonValue > a.BLColorValue && a.CentralValue > 0, Detail: fmt.Sprintf("central spectrum is uniform %.10f; B-L square spectrum [%.10f, %.10f, %.10f, %.10f]", a.CentralValue, a.BLLeptonValue, a.BLColorValue, a.BLColorValue, a.BLColorValue)},
			{Name: "propagator denominators", Passed: a.PropagatorDenominatorsDerived, Detail: "open; sector Casimirs are representation/kinetic diagnostics, not exchange propagator masses"},
			{Name: "exchange kernel update", Passed: a.ExchangeKernelUpdated, Detail: "open; no rho_A or sector propagator rule has been selected"},
			{Name: "attractive scalar-channel theorem", Passed: a.AttractiveScalarChannelDerived, Detail: "open; operator construction does not derive exchange sign, NJL criticality, or condensation"},
			{Name: "up/down splitting", Passed: a.UpDownSplittingDerived, Detail: "open; current-sector Casimirs act on lepton/color flavor and still do not split up-type from down-type quarks"},
			{Name: "condensation claim", Passed: a.CondensationClaimAllowed, Detail: "false by design; no top condensation, Higgs VEV, or fermion mass is claimed"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed mass, coupling, threshold, y_t, v, or Higgs mass was inserted"},
		}, Notes: []string{
			a.TruthStatement,
			fmt.Sprintf("recommended next gate: %s", a.RecommendedNextGate),
			fmt.Sprintf("remaining unknowns: %s", FormatUnknowns(a.RemainingUnknowns)),
		}}
	}}
}
