package familycoefficientselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FamilyCoefficientSelectorConstrainedConnectionCurvatureSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Family coefficient selector / constrained connection curvature sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate414 audit", Passed: false, Detail: err.Error()}}}
		}
		functionalsOK := len(a.Functionals) >= 4
		traceNoSelector := false
		curvatureNoSelector := false
		for _, f := range a.Functionals {
			if f.Name == "quadratic trace/norm" && f.Executed && f.GaugeCompatible && f.EmpiricalIndependent && !f.UniqueCoefficientRay && !f.SelectorNative {
				traceNoSelector = true
			}
			if f.Name == "adjoint-curvature relative to K" && f.Executed && f.SelectsNoncommutingTexture && !f.SelectsPhysicalSectorWeights && !f.SelectorNative {
				curvatureNoSelector = true
			}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate413 coefficient-free noncommuting boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate413PairAxiomCompatible && a.Inheritance.Gate413PairNotNative && a.Inheritance.Gate413CKMCapacity && a.Inheritance.Gate413CoefficientsFree && a.Inheritance.ChargedModuliDim == Gate372ChargedFlavorModuliDim, Detail: FormatInheritance(a.Inheritance)},
			{Name: "family coefficient selector arena is formalized", Passed: a.Arena.Executed && a.Arena.NoncommutingCapacity && a.Arena.GeneratedAlgebraDimension == 9 && !a.Arena.CoefficientsNative, Detail: FormatArena(a.Arena)},
			{Name: "trace and curvature functionals audited without selector promotion", Passed: functionalsOK && traceNoSelector && curvatureNoSelector, Detail: RenderFunctionalSummary(a.Functionals)},
			{Name: "constrained connection curvature does not fix coefficients", Passed: a.Connection.Executed && a.Connection.FamilyCurvatureSampleNorm > 0 && a.Connection.YangMillsMinimizerFlat && a.Connection.FlatMinimizerCommutes && a.Connection.NonzeroCurvatureRequiresSource && a.Connection.GaugeCompatibilityIfFamilyOnly && !a.Connection.ConnectionNativeInCurrentAsha && !a.Connection.CoefficientsFixedByCurvature && a.Connection.CKMCapacityConditional && !a.Connection.CKMAnglePredicted, Detail: FormatConnection(a.Connection)},
			{Name: "sector coefficients remain free", Passed: a.Coefficients.Executed && a.Coefficients.TotalFreeTextureCoefficients > 0 && a.Coefficients.TopologicalCoefficientValuesFound == 0 && !a.Coefficients.RootsOfUnityFixCoefficients && !a.Coefficients.TraceFixesCoefficients && !a.Coefficients.CurvatureFixesCoefficients && !a.Coefficients.SectorSplittingNative && !a.Coefficients.YukawaDataImported, Detail: FormatCoefficients(a.Coefficients)},
			{Name: "moduli firewall remains native", Passed: a.Moduli.StartDim == Gate372ChargedFlavorModuliDim && a.Moduli.BestNativeDim == Gate372ChargedFlavorModuliDim && !a.Moduli.NativeReductionBelow13 && a.Moduli.ConditionalMixingCapacity && a.Moduli.CoefficientsRemainFree && a.Moduli.FirewallPreserved, Detail: FormatModuli(a.Moduli)},
			{Name: "empirical firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaMatricesInserted && a.Firewall.AxiomStatusPreserved && a.Firewall.NoNativeDerivationClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate targets boundary/source axiom minimality", Passed: a.Next.Gate == 415 && a.Next.Title == "Family Boundary Condition / Sector Source Axiom Minimality Sieve", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}

func RenderFunctionalSummary(fs []FunctionalAudit) string {
	out := ""
	for i, f := range fs {
		if i > 0 {
			out += "\n"
		}
		out += FormatFunctional(f)
	}
	return out
}
