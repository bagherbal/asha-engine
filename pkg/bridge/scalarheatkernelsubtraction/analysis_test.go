package scalarheatkernelsubtraction

import "testing"

func TestGate304Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.ContactF0Promoted || a.Input.PromotedF0Value != 7 || !a.Input.F0Positive || !a.Input.KineticNormalizationAnchored {
		t.Fatalf("bad Gate 304 inheritance: %s", FormatGate304Inheritance(a.Input))
	}
	if a.Input.HigherMomentsLocked || a.Input.UniqueProfileShapeDerived || a.Input.HeatKernelSubtractionClaimed || a.Input.NumericalZHComputed || a.Input.HiggsMassPredictionClaimed || a.Input.NumericalYukawasInserted {
		t.Fatalf("Gate 305 inherited overclaimed state: %s", FormatGate304Inheritance(a.Input))
	}
}

func TestRawA2Decomposition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.A2.DecompositionFormalized || !a.A2.FieldIndependentVacuumSeen || !a.A2.ScalarPower2ChannelSeen || !a.A2.MixedTermsFirewalled || a.A2.NumericalCoefficientUsed || len(a.A2.Components) < 3 {
		t.Fatalf("bad raw a2 decomposition: %s", FormatRawA2(a.A2))
	}
	var vacuum, scalar bool
	for _, c := range a.A2.Components {
		if c.VacuumArtifact && c.Subtracted {
			vacuum = true
		}
		if c.DynamicalScalar && c.PhysicalCandidate && !c.Subtracted && c.FieldPower == 2 {
			scalar = true
		}
	}
	if !vacuum || !scalar {
		t.Fatalf("a2 components did not mark vacuum/scalar channels correctly: %s", FormatRawA2(a.A2))
	}
}

func TestSubtractionScheme(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Subtraction.Formalized || !a.Subtraction.LinearityRequired || !a.Subtraction.GaugeCovariant || len(a.Subtraction.SubtractedVacuumPieces) == 0 || len(a.Subtraction.RetainedDynamicalPieces) == 0 {
		t.Fatalf("subtraction scheme not formalized: %s", FormatSubtraction(a.Subtraction))
	}
	if a.Subtraction.BackgroundIndependent || a.Subtraction.SchemeUnique || a.Subtraction.CountertermPhysicallyFixed || a.Subtraction.NumericalCountertermUsed {
		t.Fatalf("subtraction scheme overclaimed uniqueness/counterterms: %s", FormatSubtraction(a.Subtraction))
	}
}

func TestHiggsMassExtractionMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.MassMap.MapFormalized || !a.MassMap.UsesGate300Normalization || !a.MassMap.UsesSubtractedA2 || !a.MassMap.RequiresPositiveZH || !a.MassMap.RequiresF2 || !a.MassMap.RequiresCutoffScale || !a.MassMap.RequiresYukawaAmplitudes {
		t.Fatalf("mass map missing required obligations: %s", FormatMassMap(a.MassMap))
	}
	if a.MassMap.NumericalMassComputed {
		t.Fatalf("Gate 305 must not compute a numerical Higgs mass: %s", FormatMassMap(a.MassMap))
	}
}

func TestF2DependencySieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.F2.DependencyFormalized || a.F2.F2LockedByGate304 || !a.F2.SameProfileCouldVaryF2 || !a.F2.RequiresProfileShapeRule || !a.F2.RequiresCutoffScaleLambda || a.F2.CanClaimMassWithoutF2 || !a.F2.PredictivePowerLostIfFreeF2 {
		t.Fatalf("bad f2 dependency sieve: %s", FormatF2(a.F2))
	}
}

func TestChannelSeparationLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Channels.QuadraticChannelIsolated || !a.Channels.VacuumChannelSubtracted || !a.Channels.A4F0SealPreserved || !a.Channels.A2F2StillOpen || !a.Channels.NoDynamicsOverclaimed {
		t.Fatalf("bad channel separation: %s", FormatChannels(a.Channels))
	}
	if a.Channels.QuarticChannelTouched || a.Channels.GaugeKineticDisturbed {
		t.Fatalf("Gate 305 disturbed unrelated channels: %s", FormatChannels(a.Channels))
	}
}

func TestFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoNumericalF2Inserted || !a.Firewalls.NoCutoffScaleInserted || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoNumericalZHComputed || !a.Firewalls.NoHiggsMassPredictionClaimed || !a.Firewalls.NoHiggsQuarticPredictionClaimed || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.NoUniqueSubtractionSchemeClaimed || !a.Firewalls.NoProfileHigherMomentLockClaimed || !a.Firewalls.F0SealPreservedOnlyForA4 || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := ScalarHeatKernelSubtractionHiggsPotentialChannelSeparationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
