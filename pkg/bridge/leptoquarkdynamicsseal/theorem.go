package leptoquarkdynamicsseal

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func LeptoquarkDynamicsSealBaryonConservationTheorem() theorem.Theorem {
	const id = "BRIDGE-LEPTOQUARK-DYNAMICS-SEAL-BARYON-CONSERVATION"
	const name = "Pati-Salam leptoquark current dynamics seal and B-L-preserving proton-decay obstruction"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 209 audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: []theorem.Check{
			{Name: "Gate 208 threat inherited", Passed: a.Gate208.CurrentConnectionProtonStable && a.Gate208.LeptoquarkCurrentSlotsPresent && !a.Gate208.AbsoluteBaryonConservationProved, Detail: a.Gate208.TruthStatement},
			{Name: "native leptoquark dynamics audit", Passed: a.Dynamics.AllSlotsKinematicOnly && !a.Dynamics.AnyDynamicMediator, Detail: FormatDynamics(a.Dynamics)},
			{Name: "curvature/action/local-field/propagator absent", Passed: !a.Dynamics.AnyGaugeCurvatureDerived && !a.Dynamics.AnyFiniteActionDerived && !a.Dynamics.AnyLocalFieldMapDerived && !a.Dynamics.AnyPropagatorDerived, Detail: a.Dynamics.Verdict},
			{Name: "mass and coefficient absent", Passed: !a.Dynamics.AnyMassScaleDerived && !a.Dynamics.AnyCouplingCoefficientKnown, Detail: "no suppression scale, exchange denominator, or operator coefficient is derived"},
			{Name: "LeptoquarkDynamicsSeal active", Passed: a.Seal.Active && a.Seal.ForbidsGaugeActivation && a.Seal.ForbidsPropagatorUse && a.Seal.ForbidsOperatorCoefficient, Detail: FormatSeal(a.Seal)},
			{Name: "B-L preserving templates still audited honestly", Passed: a.Operators.AllTemplatesBMinusLPreserving && a.Gate208.BMinusLDoesNotForbidStandardTemplates, Detail: fmt.Sprintf("templates=%d; B-L is not used as a fake firewall", a.Operators.TemplatesAudited)},
			{Name: "sealed operator obstruction", Passed: a.Operators.SealedObstruction && !a.Operators.AnyTemplateConstructible && !a.Operators.AnySuppressionScaleComputed, Detail: a.Operators.Verdict},
			{Name: "sealed-connection baryon conservation", Passed: a.Conservation.SealedConnectionBaryonConservation && a.Conservation.ConditionalOnLeptoquarkSeal, Detail: a.Conservation.Verdict},
			{Name: "proton lifetime computation obstructed", Passed: a.Conservation.ProtonLifetimeStrictlyObstructed && !a.Operators.ProtonLifetimeComputationLegal, Detail: "no proton lifetime formula is legal while mediator dynamics and operator coefficients are sealed absent"},
			{Name: "firewalls", Passed: a.Firewall.NoSU5Imported && a.Firewall.NoSO10Imported && a.Firewall.NoPatiSalamGaugeDynamicsImported && a.Firewall.NoProtonLifetimeComputed && a.Firewall.SealDoesNotRewriteNativeFailure, Detail: "no unified gauge dynamics, no Pati-Salam dynamics, no leptoquark mass/propagator/coefficient, and no lifetime calculation were imported"},
		}, Notes: []string{
			StatusNativeLeptoquarkDynamicsObstructed,
			StatusConditionalOnLeptoquarkDynamicsSeal,
			StatusSealedConnectionBaryonConservation,
			a.TruthStatement,
			"remaining unknowns: " + FormatUnknowns(a.Firewall.RemainingUnknowns),
		}}
	}}
}
