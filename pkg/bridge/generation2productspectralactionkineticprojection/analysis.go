// Package generation2productspectralactionkineticprojection implements Gate 500:
// Product Spectral-Action Scalar Kinetic Projection Audit.
//
// Gate 499 promoted the structural Higgs-doublet/DΦ transformation socket from
// finite inner fluctuations, but it still blocked the action-level kinetic
// projection.  Gate 500 therefore asks whether the product almost-commutative
// spectral action itself supplies the scalar kinetic term and its normalization.
//
// The answer is deliberately surgical.  The CCM product-action ledger does read
// off the scalar channel f0 a |D_μ φ|² / π², so the symbolic product kinetic
// projection exists.  But the coefficient contains the finite Dirac/Yukawa trace
// a=Tr(Y†Y).  Since Gate 489 sealed Yukawa textures and Gate 499 did not derive
// the normalized scalar action, the exact scalar kinetic coefficient, canonical
// field rescaling, Higgs VEV, kappa promotion, and W/Z mass matrix remain
// bridge/environmental.  Gate 500 promotes only the symbolic action form, not a
// physical electroweak scale.
package generation2productspectralactionkineticprojection

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ccmspectralactionsubstitution"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2innerfluctuationdphiprovenance"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2yukawaselectorairlock"
)

const (
	AuditID = "GATE500-PRODUCT-SPECTRAL-ACTION-SCALAR-KINETIC-PROJECTION-AUDIT"

	StatusGate499Inherited                 = "CONDITIONAL_SUPPORT_GATE499_STRUCTURAL_DPHI_SOCKET_INHERITED"
	StatusCCMProductActionInherited        = "CONDITIONAL_SUPPORT_CCM_PRODUCT_SPECTRAL_ACTION_LEDGER_INHERITED"
	StatusSymbolicKineticProjectionReadOff = "CONDITIONAL_SUPPORT_SYMBOLIC_SCALAR_KINETIC_PROJECTION_READ_OFF"
	StatusDphiSquaredChannelIdentified     = "CONDITIONAL_SUPPORT_DPHI_DAGGER_DPHI_ACTION_FORM_IDENTIFIED"
	StatusYukawaTraceDependenceExposed     = "CONDITIONAL_SUPPORT_SCALAR_KINETIC_COEFFICIENT_DEPENDS_ON_YUKAWA_TRACE_A"
	StatusCanonicalRescalingFormulaReadOff = "CONDITIONAL_SUPPORT_CANONICAL_SCALAR_RESCALING_FORMULA_READ_OFF"
	StatusRepresentationActionCompatible   = "CONDITIONAL_SUPPORT_INNER_FLUCTUATION_REPRESENTATION_AND_PRODUCT_ACTION_COMPATIBLE"
	StatusSymbolicProjectionBridgeAccepted = "CONDITIONAL_SUPPORT_SYMBOLIC_PRODUCT_ACTION_KINETIC_PROJECTION_BRIDGE_ACCEPTED"
	StatusFirewallPreserved                = "FIREWALL_PRESERVED_NO_ELECTROWEAK_MASS_OR_FLAVOR_DATA_IMPORTED"
	StatusRegistryWriteBlocked             = "FIREWALL_BLOCKED_NATIVE_SCALAR_KINETIC_AND_WZ_REGISTRY_WRITE"

	StatusFailedYukawaTraceASealed            = "FAILED_ROUTE_YUKAWA_TRACE_A_NOT_NATIVE_NUMERIC_DUE_TO_13_MODULI_FIREWALL"
	StatusFailedHeatKernelCoefficientNotFixed = "FAILED_ROUTE_HEAT_KERNEL_SCALAR_KINETIC_COEFFICIENT_NOT_NUMERICALLY_DERIVED"
	StatusFailedCanonicalI4MetricNotSelected  = "FAILED_ROUTE_CANONICAL_I4_SCALAR_METRIC_NOT_SELECTED_BY_PRODUCT_ACTION_ALONE"
	StatusFailedVacuumOrientationStillOpen    = "FAILED_ROUTE_SCALAR_VACUUM_ORIENTATION_STILL_NOT_NATIVE_SELECTED"
	StatusFailedKappaStillBridge              = "FAILED_ROUTE_KAPPA_U1_SIX_REMAINS_BRIDGE_CANDIDATE"
	StatusFailedWZMassStillBlocked            = "FAILED_ROUTE_PHYSICAL_WZ_MASS_MATRIX_STILL_BLOCKED"
	StatusGate501RedirectDefined              = "CONDITIONAL_SUPPORT_GATE501_YUKAWA_TRACE_SCALAR_NORMALIZATION_AIRLOCK_REDIRECT_DEFINED"
)

type Inheritance struct {
	Executed                               bool
	Gate499AuditDefined                    bool
	StructuralDphiSocketFound              bool
	StructuralScalarSU2RepresentationFound bool
	ProductKineticProjectionWasOpen        bool
	NativeDphiActionWasOpen                bool
	HeatKernelScalarCoefficientWasOpen     bool
	KappaStillBridge                       bool
	NoElectroweakFlavorDataImported        bool
	Verdict                                string
	Reason                                 string
}

type ProductActionLedger struct {
	Executed                  bool
	CCMFormulaInstalled       bool
	ProductGeometryRecognized bool
	HiggsKineticFormula       string
	CanonicalRescalingFormula string
	HiggsQuarticTraceRatio    float64
	GaugeLedgerRequiresTrace  bool
	CosmologicalOpen          bool
	YukawaTraceSymbolsKept    bool
	StructuralClosure         bool
	FullNumericalClosure      bool
	Verdict                   string
	Reason                    string
}

type KineticProjection struct {
	Executed                         bool
	ScalarOneFormRepresentationKnown bool
	ProductActionContainsDphiSquared bool
	SymbolicKineticProjectionReadOff bool
	CoefficientSymbol                string
	CoefficientDependsOnF0           bool
	CoefficientDependsOnYukawaTraceA bool
	YukawaTraceANativelyNumeric      bool
	HeatKernelCoefficientNumeric     bool
	CanonicalScalarRescalingReadOff  bool
	CanonicalI4MetricSelected        bool
	NativeKineticProjectionClosed    bool
	Verdict                          string
	Reason                           string
}

type YukawaAirlock struct {
	Executed                    bool
	Gate489AirlockAvailable     bool
	YukawaNativeSelectors       int
	YukawaEntriesEnvironmental  bool
	YukawaRankThreeDerived      bool
	CKMOrientationEnvironmental bool
	TraceAUsesYukawaSpectrum    bool
	TraceASealedByFirewall      bool
	Verdict                     string
	Reason                      string
}

type Boundary struct {
	Executed                               bool
	SymbolicProductKineticProjectionAccept bool
	DphiActionFormAccepted                 bool
	NativeScalarKineticCoefficientDerived  bool
	NativeCanonicalScalarMetricDerived     bool
	NativeVacuumOrientationDerived         bool
	NativeKappaSelected                    bool
	NativeGaugeHessianSelected             bool
	NativeWZMassMatrixDerived              bool
	Verdict                                string
	Reason                                 string
}

type Firewall struct {
	Executed                  bool
	ObservedWMassImported     bool
	ObservedZMassImported     bool
	ObservedHiggsMassImported bool
	FermiConstantImported     bool
	WeakAngleImported         bool
	FineStructureImported     bool
	GaugeCouplingImported     bool
	HiggsVEVImported          bool
	YukawaImported            bool
	CKMPMNSImported           bool
	NativeKineticWritten      bool
	NativeMetricWritten       bool
	NativeVacuumWritten       bool
	NativeWZMassWritten       bool
	Verdict                   string
	Reason                    string
}

type RegistryUpdate struct {
	NativeEntries        []string
	BridgeEntries        []string
	EnvironmentalEntries []string
	FailedRoutes         []string
	OpenTheorems         []string
}

type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}

type Analysis struct {
	Inheritance       Inheritance
	ProductAction     ProductActionLedger
	KineticProjection KineticProjection
	YukawaAirlock     YukawaAirlock
	Boundary          Boundary
	Firewall          Firewall
	Registry          RegistryUpdate
	Next              NextStep
	Truth             string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g499, err := generation2innerfluctuationdphiprovenance.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate499 inner-fluctuation Dphi provenance audit: %w", err)
	}
	ccm, err := ccmspectralactionsubstitution.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit CCM product spectral-action ledger: %w", err)
	}
	g489, err := generation2yukawaselectorairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate489 Yukawa selector airlock: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g499)
	a.ProductAction = buildProductAction(ccm)
	a.KineticProjection = buildKineticProjection(a.Inheritance, a.ProductAction)
	a.YukawaAirlock = buildYukawaAirlock(g489)
	a.Boundary = buildBoundary(a.KineticProjection, a.YukawaAirlock)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g generation2innerfluctuationdphiprovenance.Analysis) Inheritance {
	return Inheritance{
		Executed:                               true,
		Gate499AuditDefined:                    true,
		StructuralDphiSocketFound:              g.Dphi.StructuralDphiSocketFound,
		StructuralScalarSU2RepresentationFound: g.Dphi.ScalarSU2RepresentationProvenanceClosed,
		ProductKineticProjectionWasOpen:        !g.Dphi.ProductGeometryKineticProjectionDerived,
		NativeDphiActionWasOpen:                !g.Dphi.NativeDphiActionDerived,
		HeatKernelScalarCoefficientWasOpen:     g.InnerFluctuation.HeatKernelProjectionMissing,
		KappaStillBridge:                       !g.Boundary.NativeKappaSelected,
		NoElectroweakFlavorDataImported:        g.Firewall.Executed && !g.Firewall.ObservedWMassImported && !g.Firewall.ObservedZMassImported && !g.Firewall.WeakAngleImported && !g.Firewall.HiggsVEVImported && !g.Firewall.YukawaImported && !g.Firewall.CKMPMNSImported,
		Verdict:                                StatusGate499Inherited,
		Reason:                                 "Gate499 closes structural representation provenance but leaves the product-action kinetic projection, scalar normalization, kappa, and W/Z mass theorem open.",
	}
}

func buildProductAction(c ccmspectralactionsubstitution.Analysis) ProductActionLedger {
	calc := c.Calculation
	return ProductActionLedger{
		Executed:                  true,
		CCMFormulaInstalled:       strings.Contains(calc.Formula.HiggsKinetic, "f₀ a"),
		ProductGeometryRecognized: calc.StructuralClosure && strings.Contains(calc.Lagrangian, "D_M⊗1") && strings.Contains(calc.Lagrangian, "γ₅⊗D_F"),
		HiggsKineticFormula:       calc.Formula.HiggsKinetic,
		CanonicalRescalingFormula: calc.Higgs.CanonicalFieldRescaling,
		HiggsQuarticTraceRatio:    calc.Higgs.EOverA2,
		GaugeLedgerRequiresTrace:  calc.Gauge.RequiresRepresentationTrace,
		CosmologicalOpen:          calc.Cosmological.NeedsF4 && calc.Cosmological.NeedsVacuumSubtraction,
		YukawaTraceSymbolsKept:    strings.Contains(calc.Formula.Yukawa, "13") || strings.Contains(calc.Lagrangian, "dim M_charged") || strings.Contains(calc.Formula.Yukawa, "moduli"),
		StructuralClosure:         calc.StructuralClosure,
		FullNumericalClosure:      calc.FullNumericalTOEClosure,
		Verdict:                   strings.Join([]string{StatusCCMProductActionInherited, StatusSymbolicKineticProjectionReadOff}, ";"),
		Reason:                    "The CCM almost-commutative product spectral-action ledger explicitly contains the scalar kinetic channel f0 a |D_mu phi|^2 / pi^2, but it leaves the finite trace a symbolic.",
	}
}

func buildKineticProjection(in Inheritance, p ProductActionLedger) KineticProjection {
	symbolic := in.StructuralDphiSocketFound && in.StructuralScalarSU2RepresentationFound && p.CCMFormulaInstalled && p.ProductGeometryRecognized
	return KineticProjection{
		Executed:                         true,
		ScalarOneFormRepresentationKnown: in.StructuralScalarSU2RepresentationFound,
		ProductActionContainsDphiSquared: strings.Contains(p.HiggsKineticFormula, "D") && strings.Contains(p.HiggsKineticFormula, "φ"),
		SymbolicKineticProjectionReadOff: symbolic,
		CoefficientSymbol:                "K_phi = f0 * a / pi^2, with a = Tr(Y†Y) in the finite Dirac/Yukawa sector",
		CoefficientDependsOnF0:           true,
		CoefficientDependsOnYukawaTraceA: true,
		YukawaTraceANativelyNumeric:      false,
		HeatKernelCoefficientNumeric:     false,
		CanonicalScalarRescalingReadOff:  strings.Contains(p.CanonicalRescalingFormula, "√(f₀a)") || strings.Contains(p.CanonicalRescalingFormula, "sqrt"),
		CanonicalI4MetricSelected:        false,
		NativeKineticProjectionClosed:    false,
		Verdict:                          strings.Join([]string{StatusSymbolicKineticProjectionReadOff, StatusDphiSquaredChannelIdentified, StatusYukawaTraceDependenceExposed, StatusFailedHeatKernelCoefficientNotFixed}, ";"),
		Reason:                           "The action form is read off, but the coefficient is not native-numeric because it contains the sealed Yukawa trace a; therefore the canonical scalar metric and normalized kinetic coefficient cannot be written natively.",
	}
}

func buildYukawaAirlock(g generation2yukawaselectorairlock.Analysis) YukawaAirlock {
	return YukawaAirlock{
		Executed:                    true,
		Gate489AirlockAvailable:     true,
		YukawaNativeSelectors:       g.Ledger.NativeSelectorsPassing,
		YukawaEntriesEnvironmental:  g.Airlock.YukawaEntriesEnvironmental,
		YukawaRankThreeDerived:      g.Variational.RankThreeUpMatrixDerived || g.Variational.RankThreeDownMatrixDerived,
		CKMOrientationEnvironmental: g.Airlock.CKMOrientationEnvironmental,
		TraceAUsesYukawaSpectrum:    true,
		TraceASealedByFirewall:      g.Airlock.YukawaEntriesEnvironmental && !g.Variational.RankThreeUpMatrixDerived || g.Variational.RankThreeDownMatrixDerived && g.Ledger.NativeSelectorsPassing == 0,
		Verdict:                     StatusFailedYukawaTraceASealed,
		Reason:                      "The scalar kinetic coefficient depends on a=Tr(Y†Y). Gate489 closed the native Yukawa selector branch, so a is not a native numeric invariant unless a new finite-Dirac trace theorem is discovered.",
	}
}

func buildBoundary(k KineticProjection, y YukawaAirlock) Boundary {
	acceptSymbolic := k.SymbolicKineticProjectionReadOff && k.ProductActionContainsDphiSquared
	closed := acceptSymbolic && k.YukawaTraceANativelyNumeric && k.HeatKernelCoefficientNumeric && k.CanonicalI4MetricSelected && !y.TraceASealedByFirewall
	return Boundary{
		Executed:                               true,
		SymbolicProductKineticProjectionAccept: acceptSymbolic,
		DphiActionFormAccepted:                 acceptSymbolic,
		NativeScalarKineticCoefficientDerived:  closed,
		NativeCanonicalScalarMetricDerived:     k.CanonicalI4MetricSelected,
		NativeVacuumOrientationDerived:         false,
		NativeKappaSelected:                    false,
		NativeGaugeHessianSelected:             false,
		NativeWZMassMatrixDerived:              false,
		Verdict:                                strings.Join([]string{StatusSymbolicProjectionBridgeAccepted, StatusFailedYukawaTraceASealed, StatusFailedWZMassStillBlocked}, ";"),
		Reason:                                 "Gate500 accepts the symbolic product spectral-action scalar kinetic projection, but blocks native coefficient, metric, vacuum, kappa, gauge Hessian, and W/Z mass promotion because the normalization passes through sealed Yukawa trace data.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed:                  true,
		ObservedWMassImported:     false,
		ObservedZMassImported:     false,
		ObservedHiggsMassImported: false,
		FermiConstantImported:     false,
		WeakAngleImported:         false,
		FineStructureImported:     false,
		GaugeCouplingImported:     false,
		HiggsVEVImported:          false,
		YukawaImported:            false,
		CKMPMNSImported:           false,
		NativeKineticWritten:      false,
		NativeMetricWritten:       false,
		NativeVacuumWritten:       false,
		NativeWZMassWritten:       false,
		Verdict:                   StatusFirewallPreserved,
		Reason:                    "No observed W/Z/Higgs mass, Fermi constant, weak angle, alpha, gauge coupling, VEV, Yukawa, CKM, or PMNS datum is imported; no native kinetic coefficient, metric, vacuum, or W/Z registry write is made.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"No native scalar kinetic coefficient, canonical scalar metric, vacuum orientation, kappa, gauge Hessian, or W/Z mass entry is admitted at Gate500.",
		},
		BridgeEntries: []string{
			"The product almost-commutative spectral action symbolically contains the scalar kinetic channel K_phi |D_mu phi|^2.",
			"The structural Dphi socket from finite inner fluctuations is compatible with the CCM scalar kinetic channel.",
			"Canonical field rescaling is known only symbolically: H=(sqrt(f0 a)/pi) phi under the literal CCM convention.",
		},
		EnvironmentalEntries: []string{
			"The trace a=Tr(Y†Y), observed electroweak scale, W/Z/Higgs masses, Higgs VEV, gauge couplings, weak angle, Yukawa matrices, CKM, and PMNS remain sealed.",
		},
		FailedRoutes: []string{
			StatusFailedYukawaTraceASealed,
			StatusFailedHeatKernelCoefficientNotFixed,
			StatusFailedCanonicalI4MetricNotSelected,
			StatusFailedVacuumOrientationStillOpen,
			StatusFailedKappaStillBridge,
			StatusFailedWZMassStillBlocked,
		},
		OpenTheorems: []string{
			"Prove or formally seal whether the finite-Dirac trace a=Tr(Y†Y) has a native value independent of environmental Yukawa moduli.",
			"Derive a native canonical scalar metric and vacuum orientation before using the product kinetic channel to build a W/Z mass matrix.",
			"Prove whether kappa_U1=6 follows from the same normalized product-action Hessian after scalar normalization closes.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 501, Title: "Yukawa-Trace Scalar Normalization Airlock Audit", Reason: "Gate500 shows that the scalar kinetic coefficient is symbolic but depends on a=Tr(Y†Y), exactly where the flavor firewall is closed.", PrimaryTask: "prove or seal whether the finite Dirac/Yukawa trace a can be native without selecting Yukawa textures, and determine whether scalar kinetic normalization must remain bridge/environmental"}
}

func truth(a Analysis) string {
	if a.Boundary.SymbolicProductKineticProjectionAccept && !a.Boundary.NativeScalarKineticCoefficientDerived {
		return "Gate500 promotes the product spectral-action scalar kinetic term only at the symbolic bridge level.  The CCM ledger supplies the form f0 a |D_mu phi|^2 / pi^2 and is compatible with the Gate499 Higgs-doublet/Dphi socket, but the coefficient contains the sealed Yukawa trace a=Tr(Y†Y).  Therefore the canonical scalar metric, scalar normalization, vacuum orientation, kappa_U1 promotion, and physical W/Z mass matrix remain firewalled."
	}
	return "Gate500 did not even establish the symbolic product scalar kinetic projection."
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate499AuditDefined && a.Inheritance.StructuralDphiSocketFound, "Gate499 inheritance missing"},
		{a.ProductAction.Executed && a.ProductAction.CCMFormulaInstalled && a.ProductAction.ProductGeometryRecognized && a.ProductAction.StructuralClosure, "CCM product action ledger not inherited"},
		{a.KineticProjection.Executed && a.KineticProjection.SymbolicKineticProjectionReadOff && a.KineticProjection.ProductActionContainsDphiSquared, "symbolic scalar kinetic projection not read off"},
		{a.KineticProjection.CoefficientDependsOnYukawaTraceA && !a.KineticProjection.YukawaTraceANativelyNumeric && !a.KineticProjection.HeatKernelCoefficientNumeric, "Yukawa-trace scalar coefficient firewall not enforced"},
		{a.YukawaAirlock.Executed && a.YukawaAirlock.TraceASealedByFirewall && a.YukawaAirlock.YukawaNativeSelectors == 0, "Yukawa trace airlock not inherited"},
		{a.Boundary.SymbolicProductKineticProjectionAccept && !a.Boundary.NativeScalarKineticCoefficientDerived && !a.Boundary.NativeWZMassMatrixDerived, "boundary over-promoted scalar kinetic or W/Z data"},
		{a.Firewall.Executed && !a.Firewall.ObservedWMassImported && !a.Firewall.WeakAngleImported && !a.Firewall.HiggsVEVImported && !a.Firewall.YukawaImported && !a.Firewall.NativeKineticWritten && !a.Firewall.NativeWZMassWritten, "firewall violation"},
		{a.Next.Gate == 501, "Gate501 redirect missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate499=%t structural_Dphi=%t scalar_SU2_rep=%t kinetic_open=%t native_Dphi_open=%t heat_open=%t kappa_bridge=%t no_data=%t verdict=%s reason=%s", x.Gate499AuditDefined, x.StructuralDphiSocketFound, x.StructuralScalarSU2RepresentationFound, x.ProductKineticProjectionWasOpen, x.NativeDphiActionWasOpen, x.HeatKernelScalarCoefficientWasOpen, x.KappaStillBridge, x.NoElectroweakFlavorDataImported, x.Verdict, x.Reason)
}

func FormatProductAction(x ProductActionLedger) string {
	return fmt.Sprintf("CCM=%t product=%t kinetic=%q rescale=%q e_over_a2=%.15g gauge_trace=%t cosmo_open=%t Yukawa_symbols=%t structural=%t full_numeric=%t verdict=%s reason=%s", x.CCMFormulaInstalled, x.ProductGeometryRecognized, x.HiggsKineticFormula, x.CanonicalRescalingFormula, x.HiggsQuarticTraceRatio, x.GaugeLedgerRequiresTrace, x.CosmologicalOpen, x.YukawaTraceSymbolsKept, x.StructuralClosure, x.FullNumericalClosure, x.Verdict, x.Reason)
}

func FormatKineticProjection(x KineticProjection) string {
	return fmt.Sprintf("scalar_rep=%t Dphi2=%t symbolic_projection=%t coeff=%q f0=%t a=%t a_native=%t heat_numeric=%t rescale=%t I4=%t native_closed=%t verdict=%s reason=%s", x.ScalarOneFormRepresentationKnown, x.ProductActionContainsDphiSquared, x.SymbolicKineticProjectionReadOff, x.CoefficientSymbol, x.CoefficientDependsOnF0, x.CoefficientDependsOnYukawaTraceA, x.YukawaTraceANativelyNumeric, x.HeatKernelCoefficientNumeric, x.CanonicalScalarRescalingReadOff, x.CanonicalI4MetricSelected, x.NativeKineticProjectionClosed, x.Verdict, x.Reason)
}

func FormatYukawaAirlock(x YukawaAirlock) string {
	return fmt.Sprintf("gate489=%t selectors=%d environmental=%t rank3=%t CKM_env=%t trace_a_uses_yukawa=%t trace_a_sealed=%t verdict=%s reason=%s", x.Gate489AirlockAvailable, x.YukawaNativeSelectors, x.YukawaEntriesEnvironmental, x.YukawaRankThreeDerived, x.CKMOrientationEnvironmental, x.TraceAUsesYukawaSpectrum, x.TraceASealedByFirewall, x.Verdict, x.Reason)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("symbolic=%t Dphi_form=%t native_coeff=%t native_metric=%t vacuum=%t kappa=%t hessian=%t WZ=%t verdict=%s reason=%s", x.SymbolicProductKineticProjectionAccept, x.DphiActionFormAccepted, x.NativeScalarKineticCoefficientDerived, x.NativeCanonicalScalarMetricDerived, x.NativeVacuumOrientationDerived, x.NativeKappaSelected, x.NativeGaugeHessianSelected, x.NativeWZMassMatrixDerived, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("W=%t Z=%t H=%t Fermi=%t theta=%t alpha=%t gauge=%t v=%t Yukawa=%t CKM_PMNS=%t native_kinetic=%t native_metric=%t native_vacuum=%t native_WZ=%t verdict=%s reason=%s", x.ObservedWMassImported, x.ObservedZMassImported, x.ObservedHiggsMassImported, x.FermiConstantImported, x.WeakAngleImported, x.FineStructureImported, x.GaugeCouplingImported, x.HiggsVEVImported, x.YukawaImported, x.CKMPMNSImported, x.NativeKineticWritten, x.NativeMetricWritten, x.NativeVacuumWritten, x.NativeWZMassWritten, x.Verdict, x.Reason)
}

func FormatRegistry(x RegistryUpdate) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] failed=[%s] open=[%s]", strings.Join(x.NativeEntries, "; "), strings.Join(x.BridgeEntries, "; "), strings.Join(x.EnvironmentalEntries, "; "), strings.Join(x.FailedRoutes, "; "), strings.Join(x.OpenTheorems, "; "))
}
