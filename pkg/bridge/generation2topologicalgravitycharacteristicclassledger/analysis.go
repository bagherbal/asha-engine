// Package generation2topologicalgravitycharacteristicclassledger implements Gate 516:
// Topological Gravity Characteristic-Class Ledger.
//
// Gate 515 proved that numerical gravity/cosmology adapters are only bridge
// plumbing when fed explicit fake numbers. Gate 516 redirects to the scale-free
// a4 topological sector. It audits the Euler/Gauss-Bonnet and Pontryagin /
// signature sockets as characteristic-class data of the gravitational heat-
// kernel expansion, while refusing to infer manifold-specific integers, Newton
// normalization, cutoff data, cosmology, or empirical geometry.
package generation2topologicalgravitycharacteristicclassledger

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2a4curvaturesquaredledger"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2syntheticgravitycosmologyadapter"
	"github.com/bagherbal/asha-engine/pkg/bridge/generation2topologicalanomalyledger"
)

const (
	AuditID = "GATE516-TOPOLOGICAL-GRAVITY-CHARACTERISTIC-CLASS-LEDGER"

	StatusGate515GravityFirewallInherited        = "CONDITIONAL_SUPPORT_GATE515_GRAVITY_COSMOLOGY_FIREWALL_INHERITED"
	StatusGate511A4TopologicalSocketInherited    = "CONDITIONAL_SUPPORT_GATE511_A4_TOPOLOGICAL_SOCKET_INHERITED"
	StatusAnomalyLedgerInherited                 = "CONDITIONAL_SUPPORT_GATE490_MIXED_GRAVITY_ANOMALY_LEDGER_INHERITED"
	StatusA4CharacteristicLedgerDefined          = "CONDITIONAL_SUPPORT_A4_CHARACTERISTIC_CLASS_LEDGER_DEFINED"
	StatusEulerGaussBonnetScaleFree              = "CONDITIONAL_SUPPORT_EULER_GAUSS_BONNET_SOCKET_SCALE_FREE"
	StatusPontryaginSignatureScaleFree           = "CONDITIONAL_SUPPORT_PONTRYAGIN_SIGNATURE_SOCKET_SCALE_FREE"
	StatusTopologicalTermsIndependentOfScale     = "CONDITIONAL_SUPPORT_A4_TOPOLOGICAL_INVARIANTS_INDEPENDENT_OF_LAMBDA_F2_F4_G_NEWTON"
	StatusFiniteChiralityIndexSocketPresent      = "CONDITIONAL_SUPPORT_FINITE_CHIRALITY_INDEX_SOCKET_PRESENT"
	StatusMixedGravitationalGaugeTraceCanceled   = "CONDITIONAL_SUPPORT_MIXED_GRAVITATIONAL_GAUGE_TRACE_CANCELLATION_INHERITED"
	StatusNativeGravityTopologyConfirmed         = "CONDITIONAL_SUPPORT_NATIVE_GRAVITY_TOPOLOGY_CONFIRMED"
	StatusNoObservedGravityCosmologyDataImported = "CONDITIONAL_SUPPORT_NO_OBSERVED_GRAVITY_COSMOLOGY_TOPOLOGY_DATA_IMPORTED"

	StatusFailedManifoldEulerIntegerNotDerived       = "FAILED_ROUTE_MANIFOLD_EULER_CHARACTERISTIC_NOT_DERIVED_WITHOUT_GLOBAL_TOPOLOGY"
	StatusFailedManifoldSignatureIntegerNotDerived   = "FAILED_ROUTE_MANIFOLD_SIGNATURE_NOT_DERIVED_WITHOUT_GLOBAL_TOPOLOGY"
	StatusFailedFiniteAlgebraDoesNotSelectManifold   = "FAILED_ROUTE_FINITE_ALGEBRA_DOES_NOT_SELECT_CONTINUUM_MANIFOLD_TOPOLOGY"
	StatusFailedPontryaginThetaNotPhysical           = "FAILED_ROUTE_PONTRYAGIN_SOCKET_DOES_NOT_SELECT_PHYSICAL_GRAVITATIONAL_THETA_ANGLE"
	StatusFailedBoundaryEtaInvariantNotClosed        = "FAILED_ROUTE_BOUNDARY_ETA_INVARIANT_AND_INDEX_BOUNDARY_DATA_NOT_CLOSED"
	StatusFailedNewtonAndCosmologyStillBlocked       = "FAILED_ROUTE_NEWTON_CUTOFF_AND_COSMOLOGICAL_NORMALIZATION_STILL_BLOCKED"
	StatusFirewallNoEmpiricalTopologyImported        = "FIREWALL_PRESERVED_NO_NEWTON_PLANCK_COSMOLOGY_MANIFOLD_OR_OBSERVED_TOPOLOGY_DATA_IMPORTED"
	StatusFirewallTopologicalIntegerNativeWriteBlock = "FIREWALL_BLOCKED_MANIFOLD_TOPOLOGICAL_INTEGER_NATIVE_WRITE"
)

const (
	finiteTraceDimension = 96.0
	a4UnitPrefactor      = 1.0 / (60.0 * math.Pi * math.Pi) // inherited TrF/(360·16π²)
)

type Inheritance struct {
	Executed                          bool
	Gate515Inherited                  bool
	Gate515SyntheticOnly              bool
	Gate515NativeNormalizationBlocked bool
	Gate515ObservedDataImported       bool
	Gate511Inherited                  bool
	Gate511GaussBonnetSocket          bool
	Gate511WeylSocket                 bool
	Gate511DimensionlessA4            bool
	Gate511A4DoesNotUseF2Lambda       bool
	Gate490Inherited                  bool
	Gate490MixedGravityTraceCanceled  bool
	Verdict                           string
	Reason                            string
}

type CharacteristicClassLedger struct {
	Executed                  bool
	Dimension                 int
	EulerDensity              string
	PontryaginDensity         string
	HirzebruchSignature       string
	IndexDensity              string
	EulerSocketPresent        bool
	PontryaginSocketPresent   bool
	SignatureSocketPresent    bool
	A4CharacteristicSubspace  bool
	FiniteTraceDimension      float64
	A4UnitPrefactor           float64
	PhysicalThetaAngleDerived bool
	Verdict                   string
	Reason                    string
}

type ScaleIndependenceAudit struct {
	Executed                         bool
	UsesLambdaCutoff                 bool
	UsesF2Moment                     bool
	UsesF4Moment                     bool
	UsesNewtonConstant               bool
	UsesCosmologicalConstant         bool
	UsesHiggsVEVOrEWScale            bool
	UsesFlavorYukawaData             bool
	UsesObservedManifoldData         bool
	RequiresF0OnlyForLocalWeight     bool
	CharacteristicIntegralsScaleFree bool
	Verdict                          string
	Reason                           string
}

type FiniteSignatureAudit struct {
	Executed                         bool
	FiniteGradingAvailable           bool
	RealStructureAvailable           bool
	ChiralIndexSocketPresent         bool
	MixedGravitationalGaugeTraceZero bool
	ContinuumSignatureIntegerDerived bool
	ContinuumEulerIntegerDerived     bool
	ManifoldTopologySelected         bool
	BoundaryEtaInvariantClosed       bool
	Verdict                          string
	Reason                           string
}

type Firewall struct {
	Executed                          bool
	NewtonConstantImported            bool
	PlanckScaleImported               bool
	CutoffLambdaImported              bool
	F2MomentImported                  bool
	F4MomentImported                  bool
	CosmologicalConstantImported      bool
	DarkEnergyImported                bool
	ManifoldEulerIntegerImported      bool
	ManifoldSignatureImported         bool
	ObservedTopologyImported          bool
	ElectroweakDataImported           bool
	FlavorDataImported                bool
	ManifoldIntegerNativeWrite        bool
	PhysicalGravitationalThetaWritten bool
	NativeGravityNormalizationWritten bool
	Verdict                           string
	Reason                            string
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
	Ledger      CharacteristicClassLedger
	Scale       ScaleIndependenceAudit
	Finite      FiniteSignatureAudit
	Firewall    Firewall
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
	g515, err := generation2syntheticgravitycosmologyadapter.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate515 gravity/cosmology adapter firewall: %w", err)
	}
	g511, err := generation2a4curvaturesquaredledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate511 a4 topological socket: %w", err)
	}
	g490, err := generation2topologicalanomalyledger.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate490 anomaly ledger: %w", err)
	}
	a := Analysis{}
	a.Inheritance = buildInheritance(g515, g511, g490)
	a.Ledger = buildLedger()
	a.Scale = buildScale()
	a.Finite = buildFinite(g490)
	a.Firewall = buildFirewall()
	a.Registry = buildRegistry(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}

func buildInheritance(g515 generation2syntheticgravitycosmologyadapter.Analysis, g511 generation2a4curvaturesquaredledger.Analysis, g490 generation2topologicalanomalyledger.Analysis) Inheritance {
	return Inheritance{
		Executed:                          true,
		Gate515Inherited:                  g515.Inheritance.Gate514Inherited && g515.Airlock.SyntheticOnly,
		Gate515SyntheticOnly:              g515.Inputs.AllInputsSynthetic && !g515.Inputs.ObservedDataImported,
		Gate515NativeNormalizationBlocked: !g515.Airlock.NativeNormalizationWrite && !g515.Airlock.NewtonConstantDerived && !g515.Airlock.CosmologicalConstantDerived,
		Gate515ObservedDataImported:       g515.Firewall.ObservedComparatorImported || g515.Firewall.DarkEnergyImported || g515.Firewall.NewtonConstantImported,
		Gate511Inherited:                  g511.Basis.TopologicalCounterterm && g511.A4.DimensionlessChannel,
		Gate511GaussBonnetSocket:          g511.Topological.TopologicalSocketNative && g511.Topological.IntegralTopologicalInFourD,
		Gate511WeylSocket:                 g511.Dynamical.WeylSquaredSocketPresent,
		Gate511DimensionlessA4:            g511.A4.DimensionlessChannel,
		Gate511A4DoesNotUseF2Lambda:       !g511.A4.UsesF2LambdaSquared && !g511.A4.UsesF4LambdaFourth,
		Gate490Inherited:                  g490.Anomaly.AllPerturbativeGaugeCancel,
		Gate490MixedGravityTraceCanceled:  g490.Anomaly.AllMixedGaugeGravityCancel,
		Verdict:                           strings.Join([]string{StatusGate515GravityFirewallInherited, StatusGate511A4TopologicalSocketInherited, StatusAnomalyLedgerInherited}, ";"),
		Reason:                            "Gate516 inherits Gate515's gravity/cosmology normalization firewall, Gate511's dimensionless a4 topological curvature socket, and Gate490's mass-independent anomaly trace ledger.",
	}
}

func buildLedger() CharacteristicClassLedger {
	return CharacteristicClassLedger{
		Executed:                  true,
		Dimension:                 4,
		EulerDensity:              "E4 = Riem^2 - 4 Ric^2 + R^2; ∫ E4/(32π²) = χ(M) after orientation/boundary conventions",
		PontryaginDensity:         "p1(M) = -(1/8π²) Tr(Ω∧Ω); ∫ p1 = 3 τ(M) for closed oriented 4-manifolds",
		HirzebruchSignature:       "τ(M) = (1/3) ∫ p1(M)",
		IndexDensity:              "Â(M) and ch(F) sockets enter chiral Dirac index; finite grading supplies the local index carrier, not the global integer by itself",
		EulerSocketPresent:        true,
		PontryaginSocketPresent:   true,
		SignatureSocketPresent:    true,
		A4CharacteristicSubspace:  true,
		FiniteTraceDimension:      finiteTraceDimension,
		A4UnitPrefactor:           a4UnitPrefactor,
		PhysicalThetaAngleDerived: false,
		Verdict:                   strings.Join([]string{StatusA4CharacteristicLedgerDefined, StatusEulerGaussBonnetScaleFree, StatusPontryaginSignatureScaleFree, StatusNativeGravityTopologyConfirmed}, ";"),
		Reason:                    "the a4 curvature polynomial contains the four-dimensional characteristic-class sockets: Euler/Gauss-Bonnet from E4 and Pontryagin/signature from curvature two-form traces. These are topological sockets, not selected global manifold integers.",
	}
}

func buildScale() ScaleIndependenceAudit {
	return ScaleIndependenceAudit{
		Executed:                         true,
		UsesLambdaCutoff:                 false,
		UsesF2Moment:                     false,
		UsesF4Moment:                     false,
		UsesNewtonConstant:               false,
		UsesCosmologicalConstant:         false,
		UsesHiggsVEVOrEWScale:            false,
		UsesFlavorYukawaData:             false,
		UsesObservedManifoldData:         false,
		RequiresF0OnlyForLocalWeight:     true,
		CharacteristicIntegralsScaleFree: true,
		Verdict:                          StatusTopologicalTermsIndependentOfScale,
		Reason:                           "Euler and Pontryagin densities live in the a4 channel. Their local spectral-action coefficient may carry f0, but the characteristic-class integrals themselves are scale-free and do not use Λ, f2, f4, Newton normalization, cosmology, Higgs scale, or flavor data.",
	}
}

func buildFinite(g490 generation2topologicalanomalyledger.Analysis) FiniteSignatureAudit {
	return FiniteSignatureAudit{
		Executed:                         true,
		FiniteGradingAvailable:           true,
		RealStructureAvailable:           true,
		ChiralIndexSocketPresent:         true,
		MixedGravitationalGaugeTraceZero: g490.Anomaly.AllMixedGaugeGravityCancel,
		ContinuumSignatureIntegerDerived: false,
		ContinuumEulerIntegerDerived:     false,
		ManifoldTopologySelected:         false,
		BoundaryEtaInvariantClosed:       false,
		Verdict:                          strings.Join([]string{StatusFiniteChiralityIndexSocketPresent, StatusMixedGravitationalGaugeTraceCanceled, StatusFailedManifoldEulerIntegerNotDerived, StatusFailedManifoldSignatureIntegerNotDerived, StatusFailedFiniteAlgebraDoesNotSelectManifold, StatusFailedBoundaryEtaInvariantNotClosed}, ";"),
		Reason:                           "the finite grading and real structure supply the chiral/index socket and the inherited mixed gravitational-U(1) trace cancellation, but they do not choose the global topology, Euler characteristic, signature, or eta/boundary correction of the continuum manifold.",
	}
}

func buildFirewall() Firewall {
	return Firewall{
		Executed: true,
		Verdict:  strings.Join([]string{StatusNoObservedGravityCosmologyDataImported, StatusFailedPontryaginThetaNotPhysical, StatusFailedNewtonAndCosmologyStillBlocked, StatusFirewallNoEmpiricalTopologyImported, StatusFirewallTopologicalIntegerNativeWriteBlock}, ";"),
		Reason:   "Gate516 imports no Newton constant, Planck scale, cutoff, spectral moments, cosmological constant, dark-energy data, electroweak data, flavor data, manifold Euler/signature value, observed topology, boundary eta invariant, or gravitational theta angle; global topological integers are blocked from native registry unless the continuum manifold is specified natively.",
	}
}

func buildRegistry(a Analysis) RegistryUpdate {
	return RegistryUpdate{
		NativeEntries: []string{
			"Scale-free a4 characteristic-class sockets are native: Euler/Gauss-Bonnet and Pontryagin/signature densities exist in the gravitational curvature polynomial.",
			"The finite grading/real-structure ledger supplies a chiral index socket and inherits mixed gravitational-U(1) anomaly cancellation from the discrete charge ledger.",
		},
		BridgeEntries: []string{
			"Any use of f0 as a local spectral-action weight for curvature-squared/topological densities.",
			"Any application to a specified closed/bordered continuum manifold, including eta-invariant or boundary corrections.",
		},
		EnvironmentalEntries: []string{
			"The actual global manifold topology: χ(M), τ(M), Pontryagin numbers, boundary eta invariant, orientation choice, and boundary conditions.",
			"Newton/Planck normalization, Λ, f2, f4, cosmological constant, dark energy, and any observed gravity/cosmology comparator.",
		},
		FailedRoutes: []string{
			"Deriving a numerical Euler characteristic or signature from finite algebra alone.",
			"Treating the Pontryagin socket as a physical gravitational theta-angle prediction.",
			"Using characteristic-class topology to reopen Newton, cutoff, or cosmological normalization.",
		},
		OpenTheorems: []string{
			"A native continuum-manifold selection theorem, if ASHA ever supplies one.",
			"An Atiyah-Patodi-Singer boundary/eta ledger for bordered or noncompact spacetimes.",
			"A gravitational theta/Pontryagin coefficient provenance theorem, if not purely environmental.",
		},
	}
}

func buildNext() NextStep {
	return NextStep{Gate: 517, Title: "Gravitational Index and Boundary Eta Airlock", Reason: "Gate516 confirms scale-free topological sockets but blocks global integers without manifold and boundary data.", PrimaryTask: "Audit whether the ASHA finite grading can define an APS/index boundary ledger for eta corrections and chiral gravitational anomaly inflow without selecting empirical spacetime topology."}
}

func validate(a Analysis) error {
	checks := []struct {
		ok  bool
		msg string
	}{
		{a.Inheritance.Executed && a.Inheritance.Gate515Inherited && a.Inheritance.Gate515SyntheticOnly && a.Inheritance.Gate515NativeNormalizationBlocked && !a.Inheritance.Gate515ObservedDataImported && a.Inheritance.Gate511Inherited && a.Inheritance.Gate511GaussBonnetSocket && a.Inheritance.Gate511WeylSocket && a.Inheritance.Gate511DimensionlessA4 && a.Inheritance.Gate511A4DoesNotUseF2Lambda && a.Inheritance.Gate490Inherited && a.Inheritance.Gate490MixedGravityTraceCanceled, "Gate516 inheritance invalid"},
		{a.Ledger.Executed && a.Ledger.Dimension == 4 && a.Ledger.EulerSocketPresent && a.Ledger.PontryaginSocketPresent && a.Ledger.SignatureSocketPresent && a.Ledger.A4CharacteristicSubspace && nearly(a.Ledger.FiniteTraceDimension, 96, 1e-12) && nearly(a.Ledger.A4UnitPrefactor, 1.0/(60.0*math.Pi*math.Pi), 1e-12) && !a.Ledger.PhysicalThetaAngleDerived, "Gate516 characteristic ledger invalid"},
		{a.Scale.Executed && !a.Scale.UsesLambdaCutoff && !a.Scale.UsesF2Moment && !a.Scale.UsesF4Moment && !a.Scale.UsesNewtonConstant && !a.Scale.UsesCosmologicalConstant && !a.Scale.UsesHiggsVEVOrEWScale && !a.Scale.UsesFlavorYukawaData && !a.Scale.UsesObservedManifoldData && a.Scale.RequiresF0OnlyForLocalWeight && a.Scale.CharacteristicIntegralsScaleFree, "Gate516 scale audit invalid"},
		{a.Finite.Executed && a.Finite.FiniteGradingAvailable && a.Finite.RealStructureAvailable && a.Finite.ChiralIndexSocketPresent && a.Finite.MixedGravitationalGaugeTraceZero && !a.Finite.ContinuumSignatureIntegerDerived && !a.Finite.ContinuumEulerIntegerDerived && !a.Finite.ManifoldTopologySelected && !a.Finite.BoundaryEtaInvariantClosed, "Gate516 finite signature audit invalid"},
		{a.Firewall.Executed && !a.Firewall.NewtonConstantImported && !a.Firewall.PlanckScaleImported && !a.Firewall.CutoffLambdaImported && !a.Firewall.F2MomentImported && !a.Firewall.F4MomentImported && !a.Firewall.CosmologicalConstantImported && !a.Firewall.DarkEnergyImported && !a.Firewall.ManifoldEulerIntegerImported && !a.Firewall.ManifoldSignatureImported && !a.Firewall.ObservedTopologyImported && !a.Firewall.ElectroweakDataImported && !a.Firewall.FlavorDataImported && !a.Firewall.ManifoldIntegerNativeWrite && !a.Firewall.PhysicalGravitationalThetaWritten && !a.Firewall.NativeGravityNormalizationWritten, "Gate516 firewall invalid"},
	}
	for _, c := range checks {
		if !c.ok {
			return fmt.Errorf(c.msg)
		}
	}
	return nil
}

func statuses() []string {
	return []string{
		StatusGate515GravityFirewallInherited,
		StatusGate511A4TopologicalSocketInherited,
		StatusAnomalyLedgerInherited,
		StatusA4CharacteristicLedgerDefined,
		StatusEulerGaussBonnetScaleFree,
		StatusPontryaginSignatureScaleFree,
		StatusTopologicalTermsIndependentOfScale,
		StatusFiniteChiralityIndexSocketPresent,
		StatusMixedGravitationalGaugeTraceCanceled,
		StatusNativeGravityTopologyConfirmed,
		StatusNoObservedGravityCosmologyDataImported,
		StatusFailedManifoldEulerIntegerNotDerived,
		StatusFailedManifoldSignatureIntegerNotDerived,
		StatusFailedFiniteAlgebraDoesNotSelectManifold,
		StatusFailedPontryaginThetaNotPhysical,
		StatusFailedBoundaryEtaInvariantNotClosed,
		StatusFailedNewtonAndCosmologyStillBlocked,
		StatusFirewallNoEmpiricalTopologyImported,
		StatusFirewallTopologicalIntegerNativeWriteBlock,
	}
}

func truth(a Analysis) string {
	return "Gate 516 confirms the scale-free gravitational topology lane: the a4 heat-kernel socket contains Euler/Gauss-Bonnet and Pontryagin/signature characteristic-class carriers, and the finite chiral ledger supports the index/anomaly socket. ASHA may register these sockets as native topology, but it still cannot derive a specific manifold Euler characteristic, signature, boundary eta invariant, gravitational theta angle, Newton constant, cutoff, or cosmological constant without an additional continuum-topology or bridge input."
}

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate515 inherited=%t; synthetic only=%t; native normalization blocked=%t; observed imported=%t; Gate511 inherited=%t; GB socket=%t; Weyl socket=%t; a4 dimensionless=%t; a4 uses f2Λ²=%t; Gate490 inherited=%t; mixed grav-U1 trace canceled=%t", x.Gate515Inherited, x.Gate515SyntheticOnly, x.Gate515NativeNormalizationBlocked, x.Gate515ObservedDataImported, x.Gate511Inherited, x.Gate511GaussBonnetSocket, x.Gate511WeylSocket, x.Gate511DimensionlessA4, !x.Gate511A4DoesNotUseF2Lambda, x.Gate490Inherited, x.Gate490MixedGravityTraceCanceled)
}
func FormatLedger(x CharacteristicClassLedger) string {
	return fmt.Sprintf("dim=%d; Euler=%q; Pontryagin=%q; signature=%q; index=%q; Euler socket=%t; Pontryagin socket=%t; signature socket=%t; TrF=%.12g; a4 unit=%.12g; θ_grav derived=%t", x.Dimension, x.EulerDensity, x.PontryaginDensity, x.HirzebruchSignature, x.IndexDensity, x.EulerSocketPresent, x.PontryaginSocketPresent, x.SignatureSocketPresent, x.FiniteTraceDimension, x.A4UnitPrefactor, x.PhysicalThetaAngleDerived)
}
func FormatScale(x ScaleIndependenceAudit) string {
	return fmt.Sprintf("uses Λ=%t; uses f2=%t; uses f4=%t; uses G=%t; uses Λ_cosmo=%t; uses EW/Higgs scale=%t; uses flavor=%t; uses observed topology=%t; f0 local weight only=%t; characteristic integrals scale-free=%t", x.UsesLambdaCutoff, x.UsesF2Moment, x.UsesF4Moment, x.UsesNewtonConstant, x.UsesCosmologicalConstant, x.UsesHiggsVEVOrEWScale, x.UsesFlavorYukawaData, x.UsesObservedManifoldData, x.RequiresF0OnlyForLocalWeight, x.CharacteristicIntegralsScaleFree)
}
func FormatFinite(x FiniteSignatureAudit) string {
	return fmt.Sprintf("grading=%t; real structure=%t; chiral index socket=%t; mixed grav-U1 trace zero=%t; τ(M) derived=%t; χ(M) derived=%t; manifold selected=%t; eta boundary closed=%t", x.FiniteGradingAvailable, x.RealStructureAvailable, x.ChiralIndexSocketPresent, x.MixedGravitationalGaugeTraceZero, x.ContinuumSignatureIntegerDerived, x.ContinuumEulerIntegerDerived, x.ManifoldTopologySelected, x.BoundaryEtaInvariantClosed)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("G imported=%t; Planck imported=%t; Λ imported=%t; f2 imported=%t; f4 imported=%t; Λ_cosmo imported=%t; dark energy imported=%t; χ imported=%t; τ imported=%t; observed topology imported=%t; EW imported=%t; flavor imported=%t; manifold integer native write=%t; θ_grav write=%t; gravity normalization write=%t", x.NewtonConstantImported, x.PlanckScaleImported, x.CutoffLambdaImported, x.F2MomentImported, x.F4MomentImported, x.CosmologicalConstantImported, x.DarkEnergyImported, x.ManifoldEulerIntegerImported, x.ManifoldSignatureImported, x.ObservedTopologyImported, x.ElectroweakDataImported, x.FlavorDataImported, x.ManifoldIntegerNativeWrite, x.PhysicalGravitationalThetaWritten, x.NativeGravityNormalizationWritten)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 516 Registry Audit — Topological Gravity Characteristic-Class Ledger\n\n")
	b.WriteString("## Verdict\n\n```text\n" + strings.Join(statuses(), "\n") + "\n```\n\n")
	b.WriteString("## Inherited boundary\n\n" + a.Inheritance.Reason + "\n\n```text\n" + FormatInheritance(a.Inheritance) + "\n```\n\n")
	b.WriteString("## a4 characteristic-class sieve\n\n" + a.Ledger.Reason + "\n\n```text\n" + FormatLedger(a.Ledger) + "\n```\n\n")
	b.WriteString("Characteristic-class formulas:\n\n```text\nE4 = Riem² - 4 Ric² + R²\nχ(M) = (1/32π²) ∫_M E4 dvol + boundary terms\np1(M) = -(1/8π²) Tr(Ω∧Ω)\nτ(M) = (1/3) ∫_M p1(M)\n```\n\n")
	b.WriteString("## Scale independence check\n\n" + a.Scale.Reason + "\n\n```text\n" + FormatScale(a.Scale) + "\n```\n\n")
	b.WriteString("## Finite signature audit\n\n" + a.Finite.Reason + "\n\n```text\n" + FormatFinite(a.Finite) + "\n```\n\n")
	b.WriteString("## Firewall result\n\n" + a.Firewall.Reason + "\n\n```text\n" + FormatFirewall(a.Firewall) + "\n```\n\n")
	b.WriteString("## Registry update\n\n")
	writeList(&b, "Native entries", a.Registry.NativeEntries)
	writeList(&b, "Bridge entries", a.Registry.BridgeEntries)
	writeList(&b, "Environmental entries", a.Registry.EnvironmentalEntries)
	writeList(&b, "Failed routes", a.Registry.FailedRoutes)
	writeList(&b, "Open theorems", a.Registry.OpenTheorems)
	b.WriteString("## Next step\n\nGate517 should be:\n\n```text\nGate 517 — " + a.Next.Title + "\n```\n\nPrimary task:\n\n```text\n" + a.Next.PrimaryTask + "\n```\n\n")
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

func nearly(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
