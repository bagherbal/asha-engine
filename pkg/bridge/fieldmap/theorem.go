package fieldmap

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteToContinuumFieldMapAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-TO-CONTINUUM-FIELD-MAP-AUDIT"
	const name = "finite-to-continuum scalar/contact field-map audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct finite-to-continuum field-map audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "scalar/contact active sector", Passed: a.ActiveRealDirections == 4 && a.ComplexDoubletComponents == 2, Detail: fmt.Sprintf("%d real active directions = %d complex doublet components", a.ActiveRealDirections, a.ComplexDoubletComponents)},
			{Name: "sector-level doublet evidence", Passed: a.SectorLevelDoubletDerived, Detail: "classified as a continuum-field candidate at sector level, not a beta-correcting threshold"},
			{Name: "finite scalar potential evidence", Passed: a.ScalarPotentialDerived, Detail: fmt.Sprintf("r0²=%.10f, lambda_shape=%.10f", a.Scalar.VacuumRadiusSquared, a.Scalar.LambdaShape)},
			{Name: "protected-direction resonance", Passed: a.ProtectedDirections == 3, Detail: "three protected directions exist; gauge-eating theorem remains open"},
			{Name: "primary classification", Passed: a.PrimaryClassification == ContinuumScalarCandidate, Detail: string(a.PrimaryClassification)},
			{Name: "observed low-energy Higgs identity", Passed: a.LowEnergyScaleDerived, Detail: "not derived; physical unit, v, and Higgs mass bridge remain open"},
			{Name: "high-scale source-seed identity", Passed: a.BoundaryScaleDerived, Detail: "not derived; boundary scale M* and finite-to-continuum matching remain open"},
			{Name: "finite regulator classification", Passed: a.RegulatorClassificationDerived, Detail: "rejected for now; no decoupling/regulator cancellation theorem is derived"},
			{Name: "threshold correction permission", Passed: a.ThresholdCorrectionAllowed, Detail: "not allowed; scalar/contact doublet is not a derived heavy threshold correction"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInput, Detail: "no observed Higgs mass, electroweak vev, or threshold scale was inserted"},
			{Name: "role audit", Passed: true, Detail: FormatRoleRequirements(a.RoleRequirements)},
		}, Notes: []string{a.TruthStatement, fmt.Sprintf("evidence: %s", FormatEvidence(a.Evidence)), fmt.Sprintf("minimum missing data: %v", a.MinimumMissingData)}}
	}}
}
