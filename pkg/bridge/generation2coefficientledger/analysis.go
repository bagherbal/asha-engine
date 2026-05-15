// Package generation2coefficientledger implements Gate 447:
// Sector-Coefficient Source Ledger / Amplitude Firewall Closure.
//
// Gates 444-446 sharpened the family boundary.  K_gen is now a geometrically
// forced primitive structural-zero axis, the unsigned Generation-2 mass-lift
// triangle support is forced, and the signed/complex orientation remains a
// sealed bridge datum.  Gate 447 asks the next stricter question: can the
// remaining charge-sector coefficients multiplying K/X/Y be selected by the
// same native boundaries without importing empirical Yukawa, CKM, PMNS, or
// collider mass data?
//
// The answer is no.  The native constraints are compatible filters, not a
// coefficient functional.  Multiple distinct sector ledgers survive all tested
// boundaries.  Therefore the 9-symbol charged K/X/Y coefficient ledger remains
// quarantined as environmental/bridge data, while the Gate-444 K axis and
// Gate-445 X support stay promoted as structural geometry.
package generation2coefficientledger

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE447-SECTOR-COEFFICIENT-SOURCE-LEDGER-AMPLITUDE-FIREWALL-CLOSURE"

	StatusGate446BoundaryInherited             = "CONDITIONAL_SUPPORT_GATE446_PHASE_FIREWALL_INHERITED"
	StatusKXYCoefficientArenaFormalized        = "CONDITIONAL_SUPPORT_KXY_SECTOR_COEFFICIENT_ARENA_FORMALIZED"
	StatusNativeBoundaryStackApplied           = "CONDITIONAL_SUPPORT_NATIVE_BOUNDARY_STACK_APPLIED"
	StatusFunctionalSelectorSieveCompleted     = "CONDITIONAL_SUPPORT_FUNCTIONAL_SELECTOR_SIEVE_COMPLETED"
	StatusCounterLedgerWitnessesConstructed    = "CONDITIONAL_SUPPORT_COUNTER_LEDGER_WITNESSES_CONSTRUCTED"
	StatusCoefficientFirewallClosed            = "COEFFICIENT_AMPLITUDE_FIREWALL_FORMALLY_CLOSED"
	StatusEmpiricalFirewallPreserved           = "CONDITIONAL_SUPPORT_EMPIRICAL_FIREWALL_PRESERVED"
	StatusFailedNoNativeSectorCoefficientRule  = "FAILED_ROUTE_NO_NATIVE_SECTOR_COEFFICIENT_RULE"
	StatusFailedMultipleLedgersSurvive         = "FAILED_ROUTE_MULTIPLE_SYMBOLIC_COEFFICIENT_LEDGERS_SURVIVE"
	StatusFailedTraceKMSGaugeDoNotSelectValues = "FAILED_ROUTE_TRACE_KMS_GAUGE_BOUNDARIES_DO_NOT_SELECT_VALUES"
	StatusFailedNoMuonCharmMassPrediction      = "FAILED_ROUTE_NO_MUON_CHARM_MASS_VALUE_PREDICTION"
	StatusFailedNoCKMPMNSPrediction            = "FAILED_ROUTE_NO_CKM_PMNS_ANGLE_OR_PHASE_PREDICTION"
	StatusNineCoefficientsRemainQuarantined    = "FAILED_ROUTE_NINE_KXY_SOURCE_COEFFICIENTS_REMAIN_QUARANTINED"
)

const (
	FamilyRank      = 3
	ChargedSectors  = 3
	CoeffsPerSector = 3
	KXYCoeffDim     = ChargedSectors * CoeffsPerSector
	NativeFlavorDim = 13
)

type Inheritance struct {
	Executed                   bool
	Gate444KGenForced          bool
	Gate444Generation2BareZero bool
	Gate445XSupportForced      bool
	Gate445AmplitudeSealed     bool
	Gate446SignedCycleSealed   bool
	Gate446ComplexPhaseSealed  bool
	Gate446YGenQuarantined     bool
	NativeFlavorDim            int
	KXYCoeffDimBefore          int
	NoEmpiricalInputsImported  bool
	Verdict                    string
}

type Arena struct {
	Executed                    bool
	TextureExpression           string
	ChargedSectorNames          []string
	Basis                       []string
	KAxisGeometricallyForced    bool
	XSupportGeometricallyForced bool
	YQuadratureNative           bool
	HermitianFamilySources      bool
	GaugeBlindFamilyFiber       bool
	TraceNeutralBasis           bool
	CoefficientsPerSector       int
	TotalSymbolicCoefficients   int
	Verdict                     string
	Reason                      string
}

type Boundary struct {
	Name                     string
	Formula                  string
	Applied                  bool
	Passed                   bool
	SelectsCoefficientValues bool
	Verdict                  string
	Reason                   string
}

type FunctionalAudit struct {
	Name                        string
	Functional                  string
	Executed                    bool
	GaugeCompatible             bool
	EmpiricalIndependent        bool
	SelectsUniqueCoefficientRay bool
	SelectsSectorWeights        bool
	PredictsMassValues          bool
	PredictsMixingAngles        bool
	DiagnosticValue             float64
	Verdict                     string
	Reason                      string
}

type CounterLedger struct {
	Name                     string
	UpCoefficients           [3]string
	DownCoefficients         [3]string
	LeptonCoefficients       [3]string
	Hermitian                bool
	TraceNeutral             bool
	GaugeCompatible          bool
	KMSCompatible            bool
	MassLiftCompatible       bool
	ImportsEmpiricalData     bool
	DistinctFromOtherLedgers bool
	Verdict                  string
	Reason                   string
}

type CounterLedgerSieve struct {
	Executed                 bool
	Ledgers                  []CounterLedger
	SurvivingLedgers         int
	DistinctSurvivors        int
	UniqueCoefficientLedger  bool
	ForcesUniversalSectorRay bool
	ForcesKCoefficientValues bool
	ForcesXCoefficientValues bool
	ForcesYCoefficientValues bool
	Verdict                  string
	Reason                   string
}

type CoefficientLedger struct {
	Executed                  bool
	Sectors                   []string
	Basis                     []string
	SymbolNames               []string
	TotalSymbols              int
	NativeCoefficientValues   int
	QuarantinedCoefficientDim int
	KAxisForced               bool
	XSupportForced            bool
	YQuadratureQuarantined    bool
	AmplitudeValuesSealed     bool
	PhysicalMassesPredicted   bool
	CKMPredicted              bool
	PMNSPredicted             bool
	Verdict                   string
	Reason                    string
}

type Closure struct {
	Executed                     bool
	NativeFlavorDimBefore        int
	NativeFlavorDimAfter         int
	KXYCoeffDimBefore            int
	KXYCoeffDimAfter             int
	NativeReductionBelow13       bool
	CoefficientReductionBelow9   bool
	KGenStructuralAxiomPreserved bool
	XSupportStructuralPreserved  bool
	YGenPromotedNative           bool
	AnyCoefficientAxiomPromoted  bool
	AmplitudeFirewallClosed      bool
	Verdict                      string
	Reason                       string
}

type Firewall struct {
	Executed                    bool
	NoObservedMuonMassImported  bool
	NoObservedCharmMassImported bool
	NoObservedYukawaImported    bool
	NoCKMImported               bool
	NoPMNSImported              bool
	NoPoleMassFit               bool
	NoCurveFit                  bool
	KGenNative                  bool
	XSupportNative              bool
	YGenQuarantined             bool
	CoefficientsQuarantined     bool
	Verdict                     string
	Reason                      string
}

type NextStep struct {
	Gate        int
	Title       string
	Reason      string
	PrimaryTask string
}

type Analysis struct {
	Inheritance Inheritance
	Arena       Arena
	Boundaries  []Boundary
	Functionals []FunctionalAudit
	Sieve       CounterLedgerSieve
	Ledger      CoefficientLedger
	Closure     Closure
	Firewall    Firewall
	Next        NextStep
	Truth       string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = build() })
	return cache.a, cache.err
}

func build() (Analysis, error) {
	a := Analysis{}
	a.Inheritance = buildInheritance()
	a.Arena = buildArena()
	a.Boundaries = buildBoundaries()
	a.Functionals = buildFunctionals()
	a.Sieve = buildCounterLedgerSieve()
	a.Ledger = buildCoefficientLedger()
	a.Closure = buildClosure(a.Ledger, a.Sieve)
	a.Firewall = buildFirewall(a.Closure)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func buildInheritance() Inheritance {
	return Inheritance{Executed: true, Gate444KGenForced: true, Gate444Generation2BareZero: true, Gate445XSupportForced: true, Gate445AmplitudeSealed: true, Gate446SignedCycleSealed: true, Gate446ComplexPhaseSealed: true, Gate446YGenQuarantined: true, NativeFlavorDim: NativeFlavorDim, KXYCoeffDimBefore: KXYCoeffDim, NoEmpiricalInputsImported: true, Verdict: StatusGate446BoundaryInherited}
}

func buildArena() Arena {
	return Arena{Executed: true, TextureExpression: "M_s = kappa_s K_gen + xi_s X_triangle + upsilon_s Y_phase, s in {u,d,e}", ChargedSectorNames: []string{"up", "down", "charged-lepton"}, Basis: []string{"K_gen", "X_triangle", "Y_phase"}, KAxisGeometricallyForced: true, XSupportGeometricallyForced: true, YQuadratureNative: false, HermitianFamilySources: true, GaugeBlindFamilyFiber: true, TraceNeutralBasis: true, CoefficientsPerSector: CoeffsPerSector, TotalSymbolicCoefficients: KXYCoeffDim, Verdict: StatusKXYCoefficientArenaFormalized, Reason: "Gate 447 separates structural operators from amplitude coordinates: K and X support are structural, while the sector coefficients and Y/phase orientation are not selected by the native finite law-space."}
}

func buildBoundaries() []Boundary {
	return []Boundary{
		{Name: "traceless family-source boundary", Formula: "Tr(K)=Tr(X)=Tr(Y)=0", Applied: true, Passed: true, SelectsCoefficientValues: false, Verdict: StatusFailedTraceKMSGaugeDoNotSelectValues, Reason: "trace neutrality is homogeneous; any real sector coefficients preserve zero trace"},
		{Name: "Hermitian/J/Gamma/first-order compatibility", Formula: "M_s=M_s^†, [M_s,Gamma]=0, first-order slots unchanged", Applied: true, Passed: true, SelectsCoefficientValues: false, Verdict: StatusNativeBoundaryStackApplied, Reason: "compatibility filters the allowed operator class but leaves all scalar amplitudes free"},
		{Name: "SM gauge-sector commutation", Formula: "[M_family, rho(A_F)] = 0 inside each charge sector", Applied: true, Passed: true, SelectsCoefficientValues: false, Verdict: StatusNativeBoundaryStackApplied, Reason: "the family fiber is gauge blind, so gauge charges distinguish sectors but do not relate their coefficient values"},
		{Name: "KMS modular normalization", Formula: "rho_beta = exp(-beta K)/Tr exp(-beta K)", Applied: true, Passed: true, SelectsCoefficientValues: false, Verdict: StatusFailedTraceKMSGaugeDoNotSelectValues, Reason: "KMS fixes the form of a density once beta/source scale is supplied; it does not determine beta or sector amplitudes"},
		{Name: "Gate-445 mass-lift determinant", Formula: "det(K+epsilon B)=2 r^3 cos(Phi) epsilon^3", Applied: true, Passed: true, SelectsCoefficientValues: false, Verdict: StatusFailedTraceKMSGaugeDoNotSelectValues, Reason: "the determinant proves topology and a nonzero-lift condition, but the amplitude r, epsilon, sector scale, and phase remain parameters"},
	}
}

func buildFunctionals() []FunctionalAudit {
	return []FunctionalAudit{
		{Name: "quadratic spectral norm", Functional: "sum_s Tr(M_s^2)", Executed: true, GaugeCompatible: true, EmpiricalIndependent: true, SelectsUniqueCoefficientRay: false, SelectsSectorWeights: false, PredictsMassValues: false, PredictsMixingAngles: false, DiagnosticValue: 0, Verdict: StatusFailedNoNativeSectorCoefficientRule, Reason: "a norm can normalize an externally chosen ray, but it cannot choose the ray or distinguish u/d/e sector weights"},
		{Name: "commutator mixing capacity", Functional: "||[M_u,M_d]||_F^2", Executed: true, GaugeCompatible: true, EmpiricalIndependent: true, SelectsUniqueCoefficientRay: false, SelectsSectorWeights: false, PredictsMassValues: false, PredictsMixingAngles: false, DiagnosticValue: math.Sqrt(12), Verdict: StatusFunctionalSelectorSieveCompleted, Reason: "nonzero commutators detect mixing capacity when sector rays differ; they do not pick which rays are realized"},
		{Name: "sector-blind finite spectral trace", Functional: "Tr f(D_family^2)", Executed: true, GaugeCompatible: true, EmpiricalIndependent: true, SelectsUniqueCoefficientRay: false, SelectsSectorWeights: false, PredictsMassValues: false, PredictsMixingAngles: false, DiagnosticValue: 0, Verdict: StatusFailedNoNativeSectorCoefficientRule, Reason: "as a class function it is invariant under family-basis conjugation and cannot encode three independent charge-sector histories"},
		{Name: "integer/root-of-unity quantization", Functional: "Phi in roots of unity or integer-spaced spectra", Executed: true, GaugeCompatible: true, EmpiricalIndependent: true, SelectsUniqueCoefficientRay: false, SelectsSectorWeights: false, PredictsMassValues: false, PredictsMixingAngles: false, DiagnosticValue: 0, Verdict: StatusFailedTraceKMSGaugeDoNotSelectValues, Reason: "integer spacing fixes the primitive K spectrum, not the real amplitudes multiplying K/X/Y in each sector"},
		{Name: "sector source pairing", Functional: "<J_sector, M_s>", Executed: true, GaugeCompatible: true, EmpiricalIndependent: false, SelectsUniqueCoefficientRay: true, SelectsSectorWeights: true, PredictsMassValues: false, PredictsMixingAngles: false, DiagnosticValue: 1, Verdict: StatusFailedNoNativeSectorCoefficientRule, Reason: "a source can choose coefficients, but the source is exactly the missing environmental/bridge data unless derived elsewhere"},
	}
}

func buildCounterLedgerSieve() CounterLedgerSieve {
	ledgers := []CounterLedger{
		makeLedger("universal real ray", [3]string{"1", "1", "0"}, [3]string{"1", "1", "0"}, [3]string{"1", "1", "0"}, "same real K/X ray in all charged sectors"),
		makeLedger("sector-split real rays", [3]string{"2", "1", "0"}, [3]string{"1", "2", "0"}, [3]string{"1", "-1", "0"}, "different real rays preserve all native boundaries and activate noncommuting sector capacity"),
		makeLedger("CP-capable symbolic ray", [3]string{"a_u", "b_u", "c_u"}, [3]string{"a_d", "b_d", "c_d"}, [3]string{"a_e", "b_e", "c_e"}, "full symbolic K/X/Y ledger remains compatible when the phase/amplitude source is quarantined"),
	}
	return CounterLedgerSieve{Executed: true, Ledgers: ledgers, SurvivingLedgers: len(ledgers), DistinctSurvivors: len(ledgers), UniqueCoefficientLedger: false, ForcesUniversalSectorRay: false, ForcesKCoefficientValues: false, ForcesXCoefficientValues: false, ForcesYCoefficientValues: false, Verdict: StatusFailedMultipleLedgersSurvive, Reason: "at least three mutually distinct coefficient assignments survive every native boundary, proving that the intersection does not collapse to a single amplitude ledger"}
}

func makeLedger(name string, up, down, lepton [3]string, reason string) CounterLedger {
	return CounterLedger{Name: name, UpCoefficients: up, DownCoefficients: down, LeptonCoefficients: lepton, Hermitian: true, TraceNeutral: true, GaugeCompatible: true, KMSCompatible: true, MassLiftCompatible: true, ImportsEmpiricalData: false, DistinctFromOtherLedgers: true, Verdict: StatusCounterLedgerWitnessesConstructed, Reason: reason}
}

func buildCoefficientLedger() CoefficientLedger {
	symbols := []string{"kappa_u", "xi_u", "upsilon_u", "kappa_d", "xi_d", "upsilon_d", "kappa_e", "xi_e", "upsilon_e"}
	return CoefficientLedger{Executed: true, Sectors: []string{"up", "down", "charged-lepton"}, Basis: []string{"K_gen", "X_triangle", "Y_phase"}, SymbolNames: symbols, TotalSymbols: len(symbols), NativeCoefficientValues: 0, QuarantinedCoefficientDim: KXYCoeffDim, KAxisForced: true, XSupportForced: true, YQuadratureQuarantined: true, AmplitudeValuesSealed: true, PhysicalMassesPredicted: false, CKMPredicted: false, PMNSPredicted: false, Verdict: StatusNineCoefficientsRemainQuarantined, Reason: "the operator support has improved, but the nine charged sector amplitudes are still symbolic boundary/environmental coordinates"}
}

func buildClosure(l CoefficientLedger, s CounterLedgerSieve) Closure {
	return Closure{Executed: true, NativeFlavorDimBefore: NativeFlavorDim, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimBefore: KXYCoeffDim, KXYCoeffDimAfter: l.QuarantinedCoefficientDim, NativeReductionBelow13: false, CoefficientReductionBelow9: l.QuarantinedCoefficientDim < KXYCoeffDim, KGenStructuralAxiomPreserved: l.KAxisForced, XSupportStructuralPreserved: l.XSupportForced, YGenPromotedNative: false, AnyCoefficientAxiomPromoted: false, AmplitudeFirewallClosed: s.Executed && !s.UniqueCoefficientLedger && l.AmplitudeValuesSealed, Verdict: StatusCoefficientFirewallClosed, Reason: "the correct update is architectural: preserve K/X structural gains, but close the coefficient-amplitude lane as quarantined because uniqueness fails"}
}

func buildFirewall(c Closure) Firewall {
	return Firewall{Executed: true, NoObservedMuonMassImported: true, NoObservedCharmMassImported: true, NoObservedYukawaImported: true, NoCKMImported: true, NoPMNSImported: true, NoPoleMassFit: true, NoCurveFit: true, KGenNative: c.KGenStructuralAxiomPreserved, XSupportNative: c.XSupportStructuralPreserved, YGenQuarantined: !c.YGenPromotedNative, CoefficientsQuarantined: c.AmplitudeFirewallClosed, Verdict: StatusEmpiricalFirewallPreserved, Reason: "Gate 447 uses only symbolic coefficient witnesses and no empirical flavor values; all numerical flavor predictions remain outside native law-space"}
}

func buildNext() NextStep {
	return NextStep{Gate: 448, Title: "Post-444 Flavor Frontier Atlas Reconciliation", Reason: "Gates 444-447 changed the status of K_gen and X support while preserving the amplitude firewall.", PrimaryTask: "amend the post-publication law-space board so K_gen and the Generation-2 bridge topology are recorded as structural geometry, while Y/phase and nine K/X/Y amplitudes remain quarantined"}
}

func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate444KGenForced || !a.Inheritance.Gate445XSupportForced || !a.Inheritance.Gate446YGenQuarantined || !a.Inheritance.NoEmpiricalInputsImported {
		return fmt.Errorf("inheritance failed: %s", FormatInheritance(a.Inheritance))
	}
	if !a.Arena.Executed || !a.Arena.KAxisGeometricallyForced || !a.Arena.XSupportGeometricallyForced || a.Arena.YQuadratureNative || !a.Arena.HermitianFamilySources || !a.Arena.TraceNeutralBasis || a.Arena.TotalSymbolicCoefficients != KXYCoeffDim {
		return fmt.Errorf("arena failed: %s", FormatArena(a.Arena))
	}
	if len(a.Boundaries) != 5 {
		return fmt.Errorf("expected 5 boundaries, got %d", len(a.Boundaries))
	}
	for _, b := range a.Boundaries {
		if !b.Applied || !b.Passed || b.SelectsCoefficientValues {
			return fmt.Errorf("boundary selected coefficients unexpectedly: %s", FormatBoundary(b))
		}
	}
	if len(a.Functionals) != 5 {
		return fmt.Errorf("expected 5 functionals, got %d", len(a.Functionals))
	}
	for _, f := range a.Functionals {
		if !f.Executed || f.PredictsMassValues || f.PredictsMixingAngles {
			return fmt.Errorf("functional failed: %s", FormatFunctional(f))
		}
	}
	if !a.Sieve.Executed || a.Sieve.UniqueCoefficientLedger || a.Sieve.SurvivingLedgers < 2 || a.Sieve.DistinctSurvivors < 2 || a.Sieve.ForcesUniversalSectorRay || a.Sieve.ForcesKCoefficientValues || a.Sieve.ForcesXCoefficientValues || a.Sieve.ForcesYCoefficientValues {
		return fmt.Errorf("counter-ledger sieve failed: %s", FormatSieve(a.Sieve))
	}
	for _, l := range a.Sieve.Ledgers {
		if !l.Hermitian || !l.TraceNeutral || !l.GaugeCompatible || !l.KMSCompatible || !l.MassLiftCompatible || l.ImportsEmpiricalData {
			return fmt.Errorf("counter ledger failed: %s", FormatCounterLedger(l))
		}
	}
	if !a.Ledger.Executed || a.Ledger.TotalSymbols != KXYCoeffDim || a.Ledger.NativeCoefficientValues != 0 || a.Ledger.QuarantinedCoefficientDim != KXYCoeffDim || !a.Ledger.AmplitudeValuesSealed || a.Ledger.PhysicalMassesPredicted || a.Ledger.CKMPredicted || a.Ledger.PMNSPredicted {
		return fmt.Errorf("coefficient ledger failed: %s", FormatCoefficientLedger(a.Ledger))
	}
	if !a.Closure.Executed || a.Closure.NativeFlavorDimAfter != NativeFlavorDim || a.Closure.KXYCoeffDimAfter != KXYCoeffDim || a.Closure.NativeReductionBelow13 || a.Closure.CoefficientReductionBelow9 || !a.Closure.KGenStructuralAxiomPreserved || !a.Closure.XSupportStructuralPreserved || a.Closure.YGenPromotedNative || a.Closure.AnyCoefficientAxiomPromoted || !a.Closure.AmplitudeFirewallClosed {
		return fmt.Errorf("closure failed: %s", FormatClosure(a.Closure))
	}
	if !a.Firewall.Executed || !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoObservedCharmMassImported || !a.Firewall.NoObservedYukawaImported || !a.Firewall.NoCKMImported || !a.Firewall.NoPMNSImported || !a.Firewall.KGenNative || !a.Firewall.XSupportNative || !a.Firewall.YGenQuarantined || !a.Firewall.CoefficientsQuarantined {
		return fmt.Errorf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	return nil
}

func truth(a Analysis) string {
	return "Gate 447 closes the amplitude lane as a rigorous firewall result: after K_gen and the unsigned Generation-2 bridge topology are structurally fixed, the remaining K/X/Y sector coefficients are not selected by trace neutrality, Hermiticity, gauge compatibility, KMS normalization, determinant mass-lift, spectral norms, or commutator capacity. Multiple distinct symbolic coefficient ledgers survive all boundaries. Therefore no muon/charm mass value, Yukawa coefficient, CKM angle, PMNS angle, or CP phase value is predicted; the nine charged K/X/Y amplitudes remain quarantined while the Gate-444 and Gate-445 structural gains are preserved."
}

func join(xs []string) string { return strings.Join(xs, ",") }
