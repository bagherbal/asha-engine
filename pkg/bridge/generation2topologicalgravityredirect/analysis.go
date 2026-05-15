// Package generation2topologicalgravityredirect implements Gate 509:
// Topological Anomalies & Gravitational Spectral Redirect.
//
// Gate 508 closed the electroweak residual-geometry adapter as bridge-only.
// Gate 509 redirects to native-facing invariants that do not depend on flavor,
// weak-angle, W/Z mass, VEV, or continuum comparator data.  The audit has two
// independent lanes:
//
//  1. inherit the discrete chiral charge ledger of Gate 490 and verify that all
//     local/mixed/global gauge-anomaly obstructions vanish exactly;
//  2. audit the product spectral-action D_total^2 / heat-kernel channel and
//     classify the Einstein-Hilbert scalar-curvature term as a structural
//     gravitational socket, while refusing Newton normalization, cutoff scale,
//     cosmological constant, and curvature-coupling promotion.
package generation2topologicalgravityredirect

import (
	"fmt"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2ewresidualgeometryairlock"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologicalanomalyledger"
	"github.com/bagherbal/asha-engine/pkg/bridge/productspectralactioncoefficients"
)

const (
	AuditID = "GATE509-TOPOLOGICAL-ANOMALIES-GRAVITATIONAL-SPECTRAL-REDIRECT"

	StatusGate508RedirectInherited                 = "CONDITIONAL_SUPPORT_GATE508_ELECTROWEAK_FIREWALL_INHERITED"
	StatusGate490AnomalyLedgerInherited            = "CONDITIONAL_SUPPORT_GATE490_TOPOLOGICAL_ANOMALY_LEDGER_INHERITED"
	StatusNativeAnomalyCancellationReaffirmed      = "CONDITIONAL_SUPPORT_NATIVE_ANOMALY_CANCELLATION_REAFFIRMED"
	StatusAnomalyFlavorMassIndependent             = "CONDITIONAL_SUPPORT_ANOMALY_LEDGER_FLAVOR_MASS_INDEPENDENT"
	StatusProductSpectralActionInherited           = "CONDITIONAL_SUPPORT_PRODUCT_SPECTRAL_ACTION_LEDGER_INHERITED"
	StatusDiracSquareCurvatureSocketDefined        = "CONDITIONAL_SUPPORT_DIRAC_SQUARE_CURVATURE_SOCKET_DEFINED"
	StatusEinsteinHilbertSocketStructurallyPresent = "CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_SCALAR_CURVATURE_SOCKET_PRESENT"
	StatusGravityAndAnomalySocketsDefined          = "CONDITIONAL_SUPPORT_NATIVE_GRAVITY_AND_ANOMALY_SOCKETS_DEFINED"
	StatusNoMassScaleDataImported                  = "CONDITIONAL_SUPPORT_NO_MASS_SCALE_OR_FLAVOR_DATA_IMPORTED"
	StatusNextCurvatureCoefficientGateDefined      = "CONDITIONAL_SUPPORT_GATE510_CURVATURE_COEFFICIENT_PROVENANCE_DEFINED"
	StatusFirewallPreserved                        = "FIREWALL_PRESERVED_NO_NEWTON_COSMOLOGY_ELECTROWEAK_OR_FLAVOR_DATA_IMPORTED"
	StatusFirewallNativeGravityWriteBlocked        = "FIREWALL_BLOCKED_GRAVITY_NORMALIZATION_NATIVE_WRITE"

	StatusFailedNewtonConstantNotDerived            = "FAILED_ROUTE_NEWTON_CONSTANT_NOT_DERIVED"
	StatusFailedCutoffScaleNotSelected              = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_NOT_SELECTED"
	StatusFailedF2MomentNotNativelySeparated        = "FAILED_ROUTE_F2_MOMENT_NOT_NATIVELY_SEPARATED_FROM_LAMBDA"
	StatusFailedCosmologicalConstantNotDerived      = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED"
	StatusFailedEHNormalizationNotClosed            = "FAILED_ROUTE_EINSTEIN_HILBERT_NORMALIZATION_NOT_CLOSED"
	StatusFailedCurvatureEndomorphismTermNotAudited = "FAILED_ROUTE_CURVATURE_ENDOMORPHISM_E_TERM_NOT_NATIVE_AUDITED_HERE"
	StatusFailedAnomalyDoesNotDeriveMassOrMixing    = "FAILED_ROUTE_ANOMALY_CANCELLATION_DOES_NOT_DERIVE_MASS_OR_MIXING"
	StatusFailedGravitySocketDoesNotDeriveFlavor    = "FAILED_ROUTE_GRAVITY_SOCKET_DOES_NOT_REOPEN_YUKAWA_CKM_PMNS"
)

type Inheritance struct {
	Executed                       bool
	Gate508EWFirewallClosed        bool
	Gate508NativeFrontierRedirect  bool
	Gate490AnomalyLedgerExecuted   bool
	Gate490AllAnomalyTracesCancel  bool
	Gate490FlavorMassIndependent   bool
	Gate377ProductActionExecuted   bool
	Gate377ProductTripleValid      bool
	Gate377HeatKernelDeclared      bool
	Gate377SMGravityStructural     bool
	NoElectroweakScaleDataImported bool
	NoObservedFlavorDataImported   bool
	Verdict                        string
	Reason                         string
}

type TopologicalAnomalySieve struct {
	Executed                      bool
	Multiplets                    int
	WeylStatesPerGeneration       int
	WeakDoubletCount              int
	WeakDoubletCountEven          bool
	PerturbativeGaugeCancel       bool
	MixedGaugeGravityCancel       bool
	WittenSU2GlobalCancels        bool
	BLTracesCancel                bool
	ExactRationalArithmetic       bool
	GenerationReplicationStable   bool
	FlavorMassIndependent         bool
	YukawaIndependent             bool
	CKMIndependent                bool
	PMNSIndependent               bool
	GaugeStabilityLedgerSatisfied bool
	DerivesYukawaTexture          bool
	DerivesCKMOrJarlskog          bool
	Verdict                       string
	Reason                        string
}

type GravitationalSpectralSocket struct {
	Executed                      bool
	ProductTripleValid            bool
	ProductAlgebra                string
	ProductHilbertSpace           string
	ProductDirac                  string
	HeatKernelDimension           int
	HeatKernelExpansionDeclared   bool
	A2ScalarCurvatureChannel      bool
	A4CurvatureSquaredChannel     bool
	EinsteinHilbertSocketPresent  bool
	RawEHCoefficientComputed      bool
	SkeletonEHCoefficientComputed bool
	RawEHCoefficientFullyPhysical bool
	SkeletonEHCoefficientPhysical bool
	SMGravityStructuralRecovered  bool
	AllCoefficientsDetermined     bool
	HardTOEClosure                bool
	Verdict                       string
	Reason                        string
}

type GravityFirewall struct {
	Executed                           bool
	NewtonConstantImported             bool
	NewtonConstantDerived              bool
	PlanckScaleImported                bool
	CosmologicalScaleImported          bool
	CutoffLambdaSelected               bool
	F2MomentSeparatedFromLambda        bool
	F2LambdaProductNativeOnly          bool
	EinsteinHilbertNormalizationClosed bool
	CosmologicalConstantDerived        bool
	CurvatureEndomorphismAudited       bool
	PhysicalMetricDynamicsDerived      bool
	NativeGravityRegistryWritten       bool
	Verdict                            string
	Reason                             string
}

type FrontierClassification struct {
	Executed                        bool
	GaugeStabilityNativeTopological bool
	GravitySocketStructural         bool
	GravityNormalizationBridge      bool
	FlavorMassBranchClosed          bool
	EWMassRatioBranchClosed         bool
	ReopensYukawaCKMPMNS            bool
	ImportsNewtonOrCosmology        bool
	Verdict                         string
	Reason                          string
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
	Inheritance Inheritance
	Anomaly     TopologicalAnomalySieve
	Gravity     GravitationalSpectralSocket
	Classify    FrontierClassification
	Firewall    GravityFirewall
	Registry    RegistryUpdate
	Next        NextStep
	Truth       string
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
	g508, err := generation2ewresidualgeometryairlock.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate508 electroweak residual-geometry airlock: %w", err)
	}
	g490, err := generation2topologicalanomalyledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate490 topological anomaly ledger: %w", err)
	}
	g377, err := productspectralactioncoefficients.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit product spectral-action coefficient ledger: %w", err)
	}

	a := Analysis{}
	a.Inheritance = buildInheritance(g508, g490, g377)
	a.Anomaly = buildAnomaly(g490)
	a.Gravity = buildGravity(g377)
	a.Classify = buildClassification(a.Anomaly, a.Gravity)
	a.Firewall = buildFirewall(a.Gravity)
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g508 generation2ewresidualgeometryairlock.Analysis, g490 generation2topologicalanomalyledger.Analysis, g377 productspectralactioncoefficients.Analysis) Inheritance {
	return Inheritance{
		Executed:                       true,
		Gate508EWFirewallClosed:        g508.Firewall.Executed && !g508.Firewall.NativeRegistryWritten && !g508.Firewall.PhysicalElectroweakPredictionMade,
		Gate508NativeFrontierRedirect:  g508.Next.Gate == 509,
		Gate490AnomalyLedgerExecuted:   g490.Anomaly.Executed,
		Gate490AllAnomalyTracesCancel:  g490.Anomaly.AllPerturbativeGaugeCancel && g490.Anomaly.AllMixedGaugeGravityCancel && g490.Anomaly.SU2GlobalWittenCancels,
		Gate490FlavorMassIndependent:   g490.Stability.FlavorMassIndependent && g490.Stability.YukawaIndependent && g490.Stability.CKMIndependent && g490.Stability.PMNSIndependent,
		Gate377ProductActionExecuted:   g377.Calculation.Executed,
		Gate377ProductTripleValid:      g377.Calculation.Product.Valid,
		Gate377HeatKernelDeclared:      g377.Calculation.Convention.Dimension == 4 && g377.Calculation.Convention.Expansion != "",
		Gate377SMGravityStructural:     g377.Calculation.StandardModelGravityStructural,
		NoElectroweakScaleDataImported: !g508.Firewall.ObservedNumbersImported && !g508.Firewall.WZMassNativeWritten && !g508.Firewall.WeakAngleNativeWritten,
		NoObservedFlavorDataImported:   !g490.Firewall.ObservedMassesImported && !g490.Firewall.ObservedYukawaImported && !g490.Firewall.ObservedCKMImported && !g490.Firewall.ObservedPMNSImported,
		Verdict:                        StatusGate508RedirectInherited,
		Reason:                         "Gate508 closed electroweak residual geometry as bridge-only; Gate509 may therefore redirect only to scale-free topological/spectral sockets.",
	}
}

func buildAnomaly(g490 generation2topologicalanomalyledger.Analysis) TopologicalAnomalySieve {
	return TopologicalAnomalySieve{
		Executed:                      true,
		Multiplets:                    len(g490.Ledger.Multiplets),
		WeylStatesPerGeneration:       g490.Ledger.LeftHandedWeylStates,
		WeakDoubletCount:              g490.Ledger.WeakDoubletCount,
		WeakDoubletCountEven:          g490.Ledger.WeakDoubletCountEven,
		PerturbativeGaugeCancel:       g490.Anomaly.AllPerturbativeGaugeCancel,
		MixedGaugeGravityCancel:       g490.Anomaly.AllMixedGaugeGravityCancel,
		WittenSU2GlobalCancels:        g490.Anomaly.SU2GlobalWittenCancels,
		BLTracesCancel:                bMinusLTracesCancel(g490.Anomaly.Moments),
		ExactRationalArithmetic:       g490.Anomaly.ExactRationalArithmetic,
		GenerationReplicationStable:   g490.Stability.FamilyReplicationPreservesZero,
		FlavorMassIndependent:         g490.Stability.FlavorMassIndependent,
		YukawaIndependent:             g490.Stability.YukawaIndependent,
		CKMIndependent:                g490.Stability.CKMIndependent,
		PMNSIndependent:               g490.Stability.PMNSIndependent,
		GaugeStabilityLedgerSatisfied: g490.Stability.GaugeStabilityLedgerSatisfied,
		DerivesYukawaTexture:          g490.Stability.YukawaTextureSelected,
		DerivesCKMOrJarlskog:          g490.Stability.CKMJarlskogDerived,
		Verdict:                       strings.Join([]string{StatusGate490AnomalyLedgerInherited, StatusNativeAnomalyCancellationReaffirmed, StatusAnomalyFlavorMassIndependent}, ";"),
		Reason:                        "the one-generation left-handed ledger cancels every local/mixed ABJ trace and clears the Witten SU(2) parity test using only discrete representation charges",
	}
}

func bMinusLTracesCancel(ms []generation2topologicalanomalyledger.Moment) bool {
	grav, cubic := false, false
	for _, m := range ms {
		if m.Symbol == "Tr(B-L)" && m.Cancels {
			grav = true
		}
		if m.Symbol == "Tr((B-L)^3)" && m.Cancels {
			cubic = true
		}
	}
	return grav && cubic
}

func buildGravity(g377 productspectralactioncoefficients.Analysis) GravitationalSpectralSocket {
	c := g377.Calculation
	return GravitationalSpectralSocket{
		Executed:                      true,
		ProductTripleValid:            c.Product.Valid,
		ProductAlgebra:                c.Product.Algebra,
		ProductHilbertSpace:           c.Product.HilbertSpace,
		ProductDirac:                  c.Product.Dirac,
		HeatKernelDimension:           c.Convention.Dimension,
		HeatKernelExpansionDeclared:   c.Convention.Expansion != "" && c.Convention.A2DiracRDensity != "",
		A2ScalarCurvatureChannel:      strings.Contains(c.Convention.A2DiracRDensity, "R") && c.A2RawEinsteinCoefficientPerMP2.Numeric > 0,
		A4CurvatureSquaredChannel:     strings.Contains(c.Convention.A4Channels, "curvature"),
		EinsteinHilbertSocketPresent:  c.A2RawEinsteinCoefficientPerMP2.Numeric > 0 && c.A2SkeletonEinsteinCoefficientPerMP2.Numeric > 0,
		RawEHCoefficientComputed:      c.A2RawEinsteinCoefficientPerMP2.Numeric > 0,
		SkeletonEHCoefficientComputed: c.A2SkeletonEinsteinCoefficientPerMP2.Numeric > 0,
		RawEHCoefficientFullyPhysical: c.A2RawEinsteinCoefficientPerMP2.FullyPhysical,
		SkeletonEHCoefficientPhysical: c.A2SkeletonEinsteinCoefficientPerMP2.FullyPhysical,
		SMGravityStructuralRecovered:  c.StandardModelGravityStructural,
		AllCoefficientsDetermined:     c.AllCoefficientsDetermined,
		HardTOEClosure:                c.HardTOEClosure,
		Verdict:                       strings.Join([]string{StatusProductSpectralActionInherited, StatusDiracSquareCurvatureSocketDefined, StatusEinsteinHilbertSocketStructurallyPresent}, ";"),
		Reason:                        "the product Dirac square/heat-kernel ledger contains an a2 scalar-curvature channel, so the Einstein-Hilbert socket is structurally present; its physical normalization remains unclosed",
	}
}

func buildClassification(an TopologicalAnomalySieve, gr GravitationalSpectralSocket) FrontierClassification {
	return FrontierClassification{
		Executed:                        true,
		GaugeStabilityNativeTopological: an.GaugeStabilityLedgerSatisfied && an.PerturbativeGaugeCancel && an.MixedGaugeGravityCancel && an.WittenSU2GlobalCancels,
		GravitySocketStructural:         gr.EinsteinHilbertSocketPresent && gr.SMGravityStructuralRecovered,
		GravityNormalizationBridge:      !gr.RawEHCoefficientFullyPhysical && !gr.SkeletonEHCoefficientPhysical,
		FlavorMassBranchClosed:          an.FlavorMassIndependent && !an.DerivesYukawaTexture && !an.DerivesCKMOrJarlskog,
		EWMassRatioBranchClosed:         true,
		ReopensYukawaCKMPMNS:            false,
		ImportsNewtonOrCosmology:        false,
		Verdict:                         StatusGravityAndAnomalySocketsDefined,
		Reason:                          "Gate509 has native leverage on topological zeroes and a structural curvature socket; it has no native license to write Newton normalization, cosmological constant, electroweak masses, or flavor moduli.",
	}
}

func buildFirewall(_ GravitationalSpectralSocket) GravityFirewall {
	return GravityFirewall{
		Executed:                           true,
		NewtonConstantImported:             false,
		NewtonConstantDerived:              false,
		PlanckScaleImported:                false,
		CosmologicalScaleImported:          false,
		CutoffLambdaSelected:               false,
		F2MomentSeparatedFromLambda:        false,
		F2LambdaProductNativeOnly:          true,
		EinsteinHilbertNormalizationClosed: false,
		CosmologicalConstantDerived:        false,
		CurvatureEndomorphismAudited:       false,
		PhysicalMetricDynamicsDerived:      false,
		NativeGravityRegistryWritten:       false,
		Verdict:                            strings.Join([]string{StatusFailedNewtonConstantNotDerived, StatusFailedCutoffScaleNotSelected, StatusFailedF2MomentNotNativelySeparated, StatusFailedCosmologicalConstantNotDerived, StatusFailedEHNormalizationNotClosed, StatusFirewallPreserved, StatusFirewallNativeGravityWriteBlocked}, ";"),
		Reason:                             "Gate509 audits the existence of the curvature socket only. It does not import G, choose Λ, separate f2 from Λ, solve the cosmological f4 channel, or normalize the Einstein-Hilbert coefficient as physical gravity.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Discrete one-generation chiral charge ledger cancels local gauge anomalies, mixed gauge-gravity anomaly traces, and the global SU(2) doublet-parity obstruction.",
			"Anomaly cancellation is generation-replication stable and independent of Yukawa amplitudes, CKM/PMNS orientation, masses, and Jarlskog data.",
		},
		BridgeEntries: []string{
			"The M×F product spectral-action ledger contains a structural a2 scalar-curvature / Einstein-Hilbert socket from D_total^2 and the heat-kernel expansion.",
			"The gravitational socket may be carried forward as structural spectral geometry, not as a native value of Newton's constant or a physical Planck-scale normalization.",
		},
		EnvironmentalEntries: []string{
			"Newton's constant, Planck/cutoff scale identification, cosmological constant, vacuum subtraction, curvature-endomorphism convention, and physical gravitational normalization remain external/bridge data unless separately proven.",
		},
		FailedRoutes: []string{
			StatusFailedNewtonConstantNotDerived,
			StatusFailedCutoffScaleNotSelected,
			StatusFailedF2MomentNotNativelySeparated,
			StatusFailedCosmologicalConstantNotDerived,
			StatusFailedEHNormalizationNotClosed,
			StatusFailedCurvatureEndomorphismTermNotAudited,
			StatusFailedAnomalyDoesNotDeriveMassOrMixing,
			StatusFailedGravitySocketDoesNotDeriveFlavor,
		},
		OpenTheorems: []string{
			"Derive or reject a native cutoff/moment theorem that closes the f2Λ² Einstein-Hilbert coefficient without importing G or M_P.",
			"Audit the curvature endomorphism E term and trace convention in D^2 before promoting any gravitational normalization.",
			"Keep cosmological f4/vacuum subtraction in a separate airlock.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 510, Title: "Curvature Coefficient Provenance and Heat-Kernel Trace Convention Audit", Reason: "Gate509 identifies the Einstein-Hilbert socket but blocks physical gravity normalization; the next valid gate must audit the exact a2 curvature coefficient, trace convention, and E-term provenance symbolically.", PrimaryTask: "prove whether the D^2 heat-kernel a2 coefficient is canonically normalized by finite spectral data, or formally quarantine Newton normalization and cutoff selection as bridge/environmental"}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate508EWFirewallClosed && a.Inheritance.Gate508NativeFrontierRedirect && a.Inheritance.Gate490AnomalyLedgerExecuted && a.Inheritance.Gate490AllAnomalyTracesCancel && a.Inheritance.Gate490FlavorMassIndependent && a.Inheritance.Gate377ProductActionExecuted && a.Inheritance.Gate377ProductTripleValid && a.Inheritance.Gate377HeatKernelDeclared && a.Inheritance.Gate377SMGravityStructural && a.Inheritance.NoElectroweakScaleDataImported && a.Inheritance.NoObservedFlavorDataImported, "Gate509 inheritance invalid"},
		{a.Anomaly.Executed && a.Anomaly.Multiplets == 6 && a.Anomaly.WeylStatesPerGeneration == 16 && a.Anomaly.WeakDoubletCount == 4 && a.Anomaly.WeakDoubletCountEven && a.Anomaly.PerturbativeGaugeCancel && a.Anomaly.MixedGaugeGravityCancel && a.Anomaly.WittenSU2GlobalCancels && a.Anomaly.BLTracesCancel && a.Anomaly.ExactRationalArithmetic && a.Anomaly.GenerationReplicationStable && a.Anomaly.FlavorMassIndependent && a.Anomaly.YukawaIndependent && a.Anomaly.CKMIndependent && a.Anomaly.PMNSIndependent && a.Anomaly.GaugeStabilityLedgerSatisfied && !a.Anomaly.DerivesYukawaTexture && !a.Anomaly.DerivesCKMOrJarlskog, "Gate509 anomaly sieve invalid"},
		{a.Gravity.Executed && a.Gravity.ProductTripleValid && a.Gravity.HeatKernelDimension == 4 && a.Gravity.HeatKernelExpansionDeclared && a.Gravity.A2ScalarCurvatureChannel && a.Gravity.A4CurvatureSquaredChannel && a.Gravity.EinsteinHilbertSocketPresent && a.Gravity.RawEHCoefficientComputed && a.Gravity.SkeletonEHCoefficientComputed && !a.Gravity.RawEHCoefficientFullyPhysical && !a.Gravity.SkeletonEHCoefficientPhysical && a.Gravity.SMGravityStructuralRecovered && !a.Gravity.AllCoefficientsDetermined && !a.Gravity.HardTOEClosure, "Gate509 gravity socket invalid"},
		{a.Classify.Executed && a.Classify.GaugeStabilityNativeTopological && a.Classify.GravitySocketStructural && a.Classify.GravityNormalizationBridge && a.Classify.FlavorMassBranchClosed && a.Classify.EWMassRatioBranchClosed && !a.Classify.ReopensYukawaCKMPMNS && !a.Classify.ImportsNewtonOrCosmology, "Gate509 classification invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.NewtonConstantDerived && !a.Firewall.PlanckScaleImported && !a.Firewall.CosmologicalScaleImported && !a.Firewall.CutoffLambdaSelected && !a.Firewall.F2MomentSeparatedFromLambda && a.Firewall.F2LambdaProductNativeOnly && !a.Firewall.EinsteinHilbertNormalizationClosed && !a.Firewall.CosmologicalConstantDerived && !a.Firewall.CurvatureEndomorphismAudited && !a.Firewall.PhysicalMetricDynamicsDerived && !a.Firewall.NativeGravityRegistryWritten, "Gate509 gravity firewall violated"},
		{a.Next.Gate == 510, "Gate510 next step missing"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func truth(a Analysis) string {
	if a.Anomaly.GaugeStabilityLedgerSatisfied && a.Gravity.EinsteinHilbertSocketPresent && !a.Firewall.NativeGravityRegistryWritten {
		return "Gate 509 redirects ASHA away from electroweak and flavor comparator arithmetic into native-facing topology and spectral geometry. The anomaly ledger is a true discrete stability theorem: the chiral charge traces vanish exactly and remain independent of masses, Yukawas, CKM, PMNS, and Jarlskog data. The gravitational lane is weaker but real: the product Dirac-square/heat-kernel expansion contains the Einstein-Hilbert scalar-curvature socket. What is not derived is equally important: no Newton constant, cutoff scale, f2 separation, cosmological constant, or physical gravitational normalization is written natively."
	}
	return "Gate 509 failed before separating the anomaly theorem from the gravitational normalization firewall."
}

func statuses() []string {
	return []string{
		StatusGate508RedirectInherited,
		StatusGate490AnomalyLedgerInherited,
		StatusNativeAnomalyCancellationReaffirmed,
		StatusAnomalyFlavorMassIndependent,
		StatusProductSpectralActionInherited,
		StatusDiracSquareCurvatureSocketDefined,
		StatusEinsteinHilbertSocketStructurallyPresent,
		StatusGravityAndAnomalySocketsDefined,
		StatusNoMassScaleDataImported,
		StatusNextCurvatureCoefficientGateDefined,
		StatusFailedNewtonConstantNotDerived,
		StatusFailedCutoffScaleNotSelected,
		StatusFailedF2MomentNotNativelySeparated,
		StatusFailedCosmologicalConstantNotDerived,
		StatusFailedEHNormalizationNotClosed,
		StatusFailedCurvatureEndomorphismTermNotAudited,
		StatusFailedAnomalyDoesNotDeriveMassOrMixing,
		StatusFailedGravitySocketDoesNotDeriveFlavor,
		StatusFirewallPreserved,
		StatusFirewallNativeGravityWriteBlocked,
	}
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate508_closed=%t redirect=%t gate490=%t anomalies_zero=%t flavor_independent=%t product_action=%t product_triple=%t heat_kernel=%t sm_gravity=%t no_EW_scale=%t no_flavor_data=%t verdict=%s reason=%s", x.Executed, x.Gate508EWFirewallClosed, x.Gate508NativeFrontierRedirect, x.Gate490AnomalyLedgerExecuted, x.Gate490AllAnomalyTracesCancel, x.Gate490FlavorMassIndependent, x.Gate377ProductActionExecuted, x.Gate377ProductTripleValid, x.Gate377HeatKernelDeclared, x.Gate377SMGravityStructural, x.NoElectroweakScaleDataImported, x.NoObservedFlavorDataImported, x.Verdict, x.Reason)
}

func FormatAnomaly(x TopologicalAnomalySieve) string {
	return fmt.Sprintf("executed=%t multiplets=%d Weyl=%d weak_doublets=%d even=%t perturbative=%t mixed_gravity=%t Witten=%t BL=%t exact=%t replicated=%t mass_independent=%t Yukawa=%t CKM=%t PMNS=%t stable=%t derives_Yukawa=%t derives_CKM_J=%t verdict=%s reason=%s", x.Executed, x.Multiplets, x.WeylStatesPerGeneration, x.WeakDoubletCount, x.WeakDoubletCountEven, x.PerturbativeGaugeCancel, x.MixedGaugeGravityCancel, x.WittenSU2GlobalCancels, x.BLTracesCancel, x.ExactRationalArithmetic, x.GenerationReplicationStable, x.FlavorMassIndependent, x.YukawaIndependent, x.CKMIndependent, x.PMNSIndependent, x.GaugeStabilityLedgerSatisfied, x.DerivesYukawaTexture, x.DerivesCKMOrJarlskog, x.Verdict, x.Reason)
}

func FormatGravity(x GravitationalSpectralSocket) string {
	return fmt.Sprintf("executed=%t product=%t algebra=%q H=%q D=%q dim=%d heat_kernel=%t a2_R=%t a4_curvature2=%t EH_socket=%t raw=%t skeleton=%t raw_physical=%t skeleton_physical=%t SM_gravity=%t all_coeffs=%t TOE=%t verdict=%s reason=%s", x.Executed, x.ProductTripleValid, x.ProductAlgebra, x.ProductHilbertSpace, x.ProductDirac, x.HeatKernelDimension, x.HeatKernelExpansionDeclared, x.A2ScalarCurvatureChannel, x.A4CurvatureSquaredChannel, x.EinsteinHilbertSocketPresent, x.RawEHCoefficientComputed, x.SkeletonEHCoefficientComputed, x.RawEHCoefficientFullyPhysical, x.SkeletonEHCoefficientPhysical, x.SMGravityStructuralRecovered, x.AllCoefficientsDetermined, x.HardTOEClosure, x.Verdict, x.Reason)
}

func FormatClassification(x FrontierClassification) string {
	return fmt.Sprintf("executed=%t gauge_stability_topological=%t gravity_socket=%t gravity_norm_bridge=%t flavor_closed=%t EW_mass_closed=%t reopens_flavor=%t imports_G_cosmo=%t verdict=%s reason=%s", x.Executed, x.GaugeStabilityNativeTopological, x.GravitySocketStructural, x.GravityNormalizationBridge, x.FlavorMassBranchClosed, x.EWMassRatioBranchClosed, x.ReopensYukawaCKMPMNS, x.ImportsNewtonOrCosmology, x.Verdict, x.Reason)
}

func FormatFirewall(x GravityFirewall) string {
	return fmt.Sprintf("executed=%t G_imported=%t G_derived=%t Planck_imported=%t cosmo_imported=%t Lambda_selected=%t f2_separated=%t f2Lambda_native_only=%t EH_norm_closed=%t cosmo_const=%t E_term=%t metric_dynamics=%t native_write=%t verdict=%s reason=%s", x.Executed, x.NewtonConstantImported, x.NewtonConstantDerived, x.PlanckScaleImported, x.CosmologicalScaleImported, x.CutoffLambdaSelected, x.F2MomentSeparatedFromLambda, x.F2LambdaProductNativeOnly, x.EinsteinHilbertNormalizationClosed, x.CosmologicalConstantDerived, x.CurvatureEndomorphismAudited, x.PhysicalMetricDynamicsDerived, x.NativeGravityRegistryWritten, x.Verdict, x.Reason)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 509 Registry Audit — Topological Anomalies & Gravitational Spectral Redirect\n\n")
	b.WriteString("## Verdict\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Inherited boundary\n\n")
	b.WriteString("Gate508 closed the electroweak residual-geometry adapter branch as bridge-only. Gate509 therefore forbids W/Z mass ratios, weak-angle values, VEVs, gauge-coupling values, and flavor data from entering the native registry. It redirects only to mass-independent topological traces and structural spectral-action sockets.\n\n")
	b.WriteString("```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## Topological Charge/Anomaly Sieve\n\n")
	b.WriteString("```text\n" + FormatAnomaly(a.Anomaly) + "\n```\n\n")
	b.WriteString("The anomaly lane is the strong part of Gate509: the discrete chiral ledger cancels the local ABJ/gauge traces, mixed gauge-gravity traces, and the global SU(2) doublet-parity obstruction per generation. Replication by families multiplies zero by the number of generations, so the cancellation is independent of Yukawa amplitudes, CKM/PMNS orientation, masses, and CP phases.\n\n")
	b.WriteString("## Gravitational Spectral Socket\n\n")
	b.WriteString("```text\n" + FormatGravity(a.Gravity) + "\n```\n\n")
	b.WriteString("The gravity lane is accepted only as a structural socket. The product spectral triple and heat-kernel expansion contain an `a2` scalar-curvature channel, so an Einstein-Hilbert term is present at the level of spectral geometry. But the physical coefficient is not native-closed: sign/trace convention, cutoff selection, `f2` separation, curvature-endomorphism terms, and matching to Newton normalization remain open.\n\n")
	b.WriteString("## Frontier classification\n\n")
	b.WriteString("```text\n" + FormatClassification(a.Classify) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n")
	b.WriteString("```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("No Newton constant, Planck scale, cosmological scale, electroweak mass/coupling, VEV, Yukawa, CKM, PMNS, or Jarlskog datum was imported. No physical gravitational coefficient was written natively.\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native", a.Registry.NativeEntries)
	writeList(&b, "Bridge", a.Registry.BridgeEntries)
	writeList(&b, "Environmental", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate510 should be:\n\n```text\nGate 510 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func writeList(b *strings.Builder, title string, xs []string) {
	b.WriteString("### " + title + "\n\n")
	if len(xs) == 0 {
		b.WriteString("- None.\n\n")
		return
	}
	for _, x := range xs {
		b.WriteString("- " + x + "\n")
	}
	b.WriteString("\n")
}
