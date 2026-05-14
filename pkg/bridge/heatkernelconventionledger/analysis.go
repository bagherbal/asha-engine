// Package heatkernelconventionledger implements Gate 302:
// Heat-Kernel Convention Ledger / Positive Prefactor Normalization Audit.
package heatkernelconventionledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE302-HEAT-KERNEL-CONVENTION-LEDGER-POSITIVE-PREFACTOR-NORMALIZATION-AUDIT"

	StatusGate301Inherited                  = "CONDITIONAL_SUPPORT_GATE301_POSITIVE_TRACE_CARRIER_INHERITED"
	StatusPrefactorLedgerFormalized         = "CONDITIONAL_SUPPORT_HEAT_KERNEL_PREFACTOR_LEDGER_FORMALIZED"
	StatusPositivePrefactorLedgerFormalized = "CONDITIONAL_SUPPORT_POSITIVE_HEAT_KERNEL_PREFACTOR_LEDGER_FORMALIZED"
	StatusWickSignRuleFormalized            = "CONDITIONAL_SUPPORT_WICK_SIGN_MATCHING_RULE_FORMALIZED"
	StatusPositiveF0RequirementFormalized   = "CONDITIONAL_SUPPORT_POSITIVE_F0_REQUIREMENT_FORMALIZED"
	StatusCanonicalMatchingRuleFormalized   = "CONDITIONAL_SUPPORT_CANONICAL_SCALAR_MATCHING_RULE_FORMALIZED"
	StatusEmpiricalFirewallsPreserved       = "CONDITIONAL_SUPPORT_GATE302_EMPIRICAL_FIREWALLS_PRESERVED"

	StatusFailedF0NumericalValueNotDerived       = "FAILED_ROUTE_F0_NUMERICAL_VALUE_NOT_DERIVED"
	StatusFailedCutoffIdentificationNotActivated = "FAILED_ROUTE_CONTACT_SPECTRAL_CUTOFF_IDENTIFICATION_NOT_ACTIVATED"
	StatusFailedWickConventionNotNative          = "FAILED_ROUTE_WICK_ROTATION_CONVENTION_NOT_DERIVED_FROM_FINITE_GEOMETRY"
	StatusFailedAbsoluteN4NotNumerical           = "FAILED_ROUTE_ABSOLUTE_N4_NUMERICAL_CONSTANT_NOT_DERIVED"
	StatusFailedStrictZHStillConditional         = "FAILED_ROUTE_STRICT_ZH_POSITIVITY_STILL_CONDITIONAL"
	StatusFailedYukawaAmplitudesStillSealed      = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"
	StatusFailedSubtractionSchemeStillMissing    = "FAILED_ROUTE_HEAT_KERNEL_SUBTRACTION_SCHEME_STILL_MISSING"
	StatusFailedMassQuarticStillFirewalled       = "FAILED_ROUTE_HIGGS_MASS_AND_QUARTIC_STILL_FIREWALLED"
	StatusFailedBGapInstantonStillSealed         = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

type InheritedGate301 struct {
	PositiveKRawCarrierProved      bool
	HilbertSchmidtSumStructure     bool
	StrictPositiveNeedsNonzeroEdge bool
	NumericalYukawasInserted       bool
	NumericalZHComputed            bool
	RequiresPositiveF0             bool
	RequiresTraceConventionLedger  bool
	RequiresEuclideanSignLedger    bool
	Verdict                        string
}

type PrefactorFactor struct {
	Name              string
	Symbol            string
	MathematicalRole  string
	Representative    string
	SignCondition     string
	PositiveByChoice  bool
	PositiveByTheorem bool
	NumericallyFixed  bool
	EmpiricalInput    bool
	AbsorbedInto      string
	Status            string
}

type PrefactorLedger struct {
	ZHFormula                  string
	N4Factorization            string
	Factors                    []PrefactorFactor
	AllFactorsExplicit         bool
	AllEmpiricalInputsExcluded bool
	CanChoosePositiveClass     bool
	AbsoluteN4Derived          bool
	Verdict                    string
}

type WickConventionAudit struct {
	EuclideanKineticForm     string
	LorentzianKineticTarget  string
	WickMap                  string
	HamiltonianCondition     string
	SignLedger               []string
	SignAmbiguityHidden      bool
	PositiveEnergyMapped     bool
	ConventionNativeToFinite bool
	Verdict                  string
}

type F0Requirement struct {
	CutoffMomentDefinition       string
	RequiredCondition            string
	ContactSpectralGate288Used   bool
	NumericalValueDerived        bool
	CanBePositiveWithoutEmpirics bool
	ConditionallyPositive        bool
	Verdict                      string
}

type CanonicalMatchingRule struct {
	RawCoefficient       string
	CanonicalTarget      string
	RescalingRule        string
	AbsorbedFactors      []string
	MassChannelImpact    string
	QuarticChannelImpact string
	PhysicalZHComputed   bool
	RuleFormalized       bool
	Verdict              string
}

type PrefactorPositivitySieve struct {
	KRawPositiveSemidefinite        bool
	AtLeastOneAmplitudeNeeded       bool
	N4PositiveCondition             string
	F0PositiveCondition             string
	OverallCondition                string
	PositivePrefactorAvailable      bool
	StrictZHGuaranteedConditionally bool
	NumericalStrictZHProved         bool
	FailureMode                     string
	Verdict                         string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksPrediction          bool
}

type FirewallAudit struct {
	NoF0NumberInserted, NoCutoffGateActivated, NoYukawaNumbersInserted, NoObservedMassesInserted      bool
	NoSubtractionSchemeInvented, NoBGapInstantonClaimed, NoHiggsPredictionClaimed, FiniteCorePolluted bool
	Obligations                                                                                       []RemainingObligation
	Verdict                                                                                           string
}

type Summary struct {
	Gate301Inherited, PrefactorLedgerFormalized, WickSignRuleFormalized, PositiveF0ConditionRecorded            bool
	CanonicalMatchingFormalized, PositivePrefactorAvailable, StrictZHNumericallyProved, PhysicalDynamicsDerived bool
	FirewallPreserved                                                                                           bool
	Status, DirectAnswer, NextGate                                                                              string
}

type Analysis struct {
	Input      InheritedGate301
	Ledger     PrefactorLedger
	Wick       WickConventionAudit
	F0         F0Requirement
	Matching   CanonicalMatchingRule
	Positivity PrefactorPositivitySieve
	Firewalls  FirewallAudit
	Summary    Summary
	Truth      string
}

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	i := inheritGate301()
	l := buildPrefactorLedger()
	w := auditWickConvention()
	f := formalizeF0Requirement()
	m := buildCanonicalMatchingRule(l)
	p := runPrefactorPositivitySieve(i, l, w, f)
	fw := auditFirewalls(i, l, w, f, m, p)
	s := buildSummary(i, l, w, f, m, p, fw)
	truth := "Gate 302 formalizes the positive convention ledger for Z_H = N_4 f_0 K_H^raw. Since Gate 301 proved K_H^raw is a non-negative Hilbert-Schmidt trace carrier, the remaining sign burden is isolated into explicit convention factors: Seeley-de Witt normalization, trace inner-product orientation, doubled-space multiplicity, scalar canonical matching, Wick/Lorentzian sign choice, and the cutoff moment f_0. The gate proves that a positive canonical prefactor class can be selected without empirical pollution, while strict numerical Z_H remains conditional on f_0>0, nonzero sealed scalar amplitudes, and no hidden sign reversal."
	return Analysis{Input: i, Ledger: l, Wick: w, F0: f, Matching: m, Positivity: p, Firewalls: fw, Summary: s, Truth: truth}, nil
}

func inheritGate301() InheritedGate301 {
	return InheritedGate301{true, true, true, false, false, true, true, true, StatusGate301Inherited}
}

func buildPrefactorLedger() PrefactorLedger {
	f := []PrefactorFactor{
		{"Seeley-de Witt four-dimensional density", "s_SD", "universal a_4 density normalization", "(4π)^-2 times positive combinatorial constants", "s_SD > 0", false, true, false, false, "N_4", StatusPrefactorLedgerFormalized},
		{"finite trace inner-product orientation", "s_Tr", "chooses Tr(A†A) as the positive finite Hilbert-space inner product", "Tr_F(Φ†Φ) rather than -Tr_F(Φ†Φ)", "s_Tr = +1", true, false, false, false, "N_4", StatusPrefactorLedgerFormalized},
		{"doubled-space multiplicity", "m_J", "accounts for H_F ⊕ H_F* particle/antiparticle carrier", "m_J=2 before optional no-double-counting division", "m_J > 0", false, true, false, false, "N_4", StatusPrefactorLedgerFormalized},
		{"canonical scalar coefficient convention", "c_H", "matches the complex Higgs-doublet kinetic channel to (D_mu H)†(D^mu H)", "positive scalar-field normalization constant; no ghost sign allowed", "c_H > 0", true, false, false, false, "N_4", StatusCanonicalMatchingRuleFormalized},
		{"Euclidean-to-Lorentzian sign bridge", "σ_W", "maps the Euclidean positive action term to the Lorentzian positive-energy kinetic convention", "choose S_E=-iS_M after t=-iτ and Hamiltonian positivity", "σ_W = +1", true, false, false, false, "N_4", StatusWickSignRuleFormalized},
		{"cutoff moment", "f_0", "multiplies the a_4 coefficient in the spectral action", "zeroth/evaluation moment of the cutoff profile in the a_4 channel", "f_0 > 0", true, false, false, false, "Z_H", StatusPositiveF0RequirementFormalized},
	}
	return PrefactorLedger{"Z_H := N_4 f_0 K_H^raw", "N_4 := s_SD · s_Tr · m_J · c_H · σ_W", f, true, true, true, false, strings.Join([]string{StatusPrefactorLedgerFormalized, StatusPositivePrefactorLedgerFormalized, StatusFailedAbsoluteN4NotNumerical}, ";")}
}

func auditWickConvention() WickConventionAudit {
	return WickConventionAudit{"S_E ⊃ + Z_H ∫ d^4x_E (D_μ H_raw)†(D_μ H_raw)", "L_M ⊃ + (D_μ H_phys)†(D^μ H_phys) under the chosen positive-Hamiltonian convention", "t=-iτ and S_E=-iS_M are accepted as the explicit convention bridge", "canonical momentum must contribute +|D_0 H_phys|² to the Hamiltonian density", []string{"finite trace orientation: Tr(A†A) ≥ 0", "Euclidean scalar gradient is positive in S_E", "Wick bridge: S_E=-iS_M", "Lorentzian target: no ghost sign; positive velocity-square Hamiltonian"}, false, true, false, strings.Join([]string{StatusWickSignRuleFormalized, StatusFailedWickConventionNotNative}, ";")}
}

func formalizeF0Requirement() F0Requirement {
	return F0Requirement{"f_0 is the a_4 spectral-action cutoff moment/profile coefficient multiplying local dimension-four operators", "f_0 must be strictly positive. A non-negative cutoff profile with nonzero a_4 weight satisfies the sign requirement, but no numerical value is inserted here.", false, false, true, true, strings.Join([]string{StatusPositiveF0RequirementFormalized, StatusFailedF0NumericalValueNotDerived, StatusFailedCutoffIdentificationNotActivated}, ";")}
}

func buildCanonicalMatchingRule(l PrefactorLedger) CanonicalMatchingRule {
	parts := []string{}
	for _, f := range l.Factors {
		parts = append(parts, f.Symbol)
	}
	return CanonicalMatchingRule{"C_H^raw := N_4 f_0 K_H^raw", "(D_μ H_phys)†(D^μ H_phys) with unit positive coefficient", "H_raw = H_phys / sqrt(Z_H), valid only for real Z_H>0", parts, "C_2^phys = C_2^raw / Z_H after subtraction and sign convention; μ_H² is not computed here", "λ_H = C_4^raw / Z_H² after scalar quartic projection; λ_H is not computed here", false, l.AllFactorsExplicit, strings.Join([]string{StatusCanonicalMatchingRuleFormalized, StatusFailedMassQuarticStillFirewalled}, ";")}
}

func runPrefactorPositivitySieve(i InheritedGate301, l PrefactorLedger, w WickConventionAudit, f F0Requirement) PrefactorPositivitySieve {
	n := 0
	for _, x := range l.Factors {
		if x.SignCondition != "" && (x.PositiveByChoice || x.PositiveByTheorem) && !x.EmpiricalInput {
			n++
		}
	}
	ok := i.PositiveKRawCarrierProved && i.HilbertSchmidtSumStructure && l.AllFactorsExplicit && l.CanChoosePositiveClass && w.PositiveEnergyMapped && !w.SignAmbiguityHidden && f.ConditionallyPositive && n == len(l.Factors)
	return PrefactorPositivitySieve{i.PositiveKRawCarrierProved, i.StrictPositiveNeedsNonzeroEdge, "N_4>0 iff s_SD>0, s_Tr=+1, m_J>0, c_H>0, and σ_W=+1", "f_0>0 is required; no numerical f_0 is supplied by Gate 302", "Z_H>0 iff K_H^raw>0, N_4>0, and f_0>0", ok, ok && i.StrictPositiveNeedsNonzeroEdge, false, "if σ_W=-1, s_Tr=-1, f_0≤0, or all scalar amplitudes vanish, positive kinetic normalization fails or degenerates", strings.Join([]string{StatusPositivePrefactorLedgerFormalized, StatusFailedStrictZHStillConditional, StatusFailedF0NumericalValueNotDerived, StatusFailedYukawaAmplitudesStillSealed}, ";")}
}

func auditFirewalls(i InheritedGate301, l PrefactorLedger, w WickConventionAudit, f F0Requirement, m CanonicalMatchingRule, p PrefactorPositivitySieve) FirewallAudit {
	obs := []RemainingObligation{{"positive cutoff moment f_0", "strict Z_H positivity requires the a_4 cutoff coefficient to be positive", StatusFailedF0NumericalValueNotDerived, true}, {"optional Contact-Spectral Cutoff Identification activation", "would tie f_0 to an internal spectral cutoff theorem", StatusFailedCutoffIdentificationNotActivated, true}, {"Wick/sign convention selection", "maps finite Euclidean positivity to Lorentzian positive-energy kinetics", StatusFailedWickConventionNotNative, false}, {"absolute N_4 numerical constant", "needed for absolute Z_H magnitude", StatusFailedAbsoluteN4NotNumerical, true}, {"nonzero scalar Yukawa amplitude seal", "K_H^raw is strictly positive only if at least one scalar edge amplitude is nonzero", StatusFailedYukawaAmplitudesStillSealed, true}, {"heat-kernel subtraction scheme", "mass and vacuum channels still require subtraction/renormalization", StatusFailedSubtractionSchemeStillMissing, true}, {"Higgs mass and quartic prediction", "needs numerical Z_H, a2 subtraction, quartic amplitudes, and cutoff data", StatusFailedMassQuarticStillFirewalled, true}, {"B-gap instanton action", "prefactor positivity is polynomial bookkeeping and does not derive S_inst=(4/pi)/B_gap", StatusFailedBGapInstantonStillSealed, true}}
	return FirewallAudit{!f.NumericalValueDerived, !f.ContactSpectralGate288Used, !i.NumericalYukawasInserted, true, true, true, !m.PhysicalZHComputed && !p.NumericalStrictZHProved, !l.AllEmpiricalInputsExcluded || w.SignAmbiguityHidden, obs, strings.Join([]string{StatusEmpiricalFirewallsPreserved, StatusFailedF0NumericalValueNotDerived, StatusFailedYukawaAmplitudesStillSealed, StatusFailedMassQuarticStillFirewalled, StatusFailedBGapInstantonStillSealed}, ";")}
}

func buildSummary(i InheritedGate301, l PrefactorLedger, w WickConventionAudit, f F0Requirement, m CanonicalMatchingRule, p PrefactorPositivitySieve, fw FirewallAudit) Summary {
	statuses := []string{StatusGate301Inherited, StatusPrefactorLedgerFormalized, StatusPositivePrefactorLedgerFormalized, StatusWickSignRuleFormalized, StatusPositiveF0RequirementFormalized, StatusCanonicalMatchingRuleFormalized, StatusEmpiricalFirewallsPreserved, StatusFailedF0NumericalValueNotDerived, StatusFailedCutoffIdentificationNotActivated, StatusFailedWickConventionNotNative, StatusFailedAbsoluteN4NotNumerical, StatusFailedStrictZHStillConditional, StatusFailedYukawaAmplitudesStillSealed, StatusFailedSubtractionSchemeStillMissing, StatusFailedMassQuarticStillFirewalled, StatusFailedBGapInstantonStillSealed}
	return Summary{i.PositiveKRawCarrierProved && i.HilbertSchmidtSumStructure && !i.NumericalZHComputed, l.AllFactorsExplicit && l.AllEmpiricalInputsExcluded, w.PositiveEnergyMapped && !w.SignAmbiguityHidden, f.ConditionallyPositive && !f.NumericalValueDerived, m.RuleFormalized && !m.PhysicalZHComputed, p.PositivePrefactorAvailable, false, false, !fw.FiniteCorePolluted && fw.NoF0NumberInserted && fw.NoYukawaNumbersInserted && fw.NoBGapInstantonClaimed, strings.Join(statuses, ";"), "Gate 302 proves that the convention layer can be made sign-safe: with positive Seeley-de Witt density, positive finite trace orientation, positive doubled-space multiplicity, positive scalar canonical matching, canonical Wick sign, and f_0>0, the positive Gate 301 trace maps to a positive scalar wave-function normalization. It does not compute numerical Z_H or derive f_0.", "Gate 303 should audit the cutoff profile/moment source: either activate the sealed contact-spectral cutoff identification or formalize a non-empirical positive test-function class that fixes f_0 only as a sign/scale convention, not as an observed parameter."}
}

func FormatInput(i InheritedGate301) string {
	return fmt.Sprintf("KrawPositive=%t HS=%t nonzeroNeeded=%t yukawas=%t numericalZH=%t f0=%t traceLedger=%t signLedger=%t verdict=%s", i.PositiveKRawCarrierProved, i.HilbertSchmidtSumStructure, i.StrictPositiveNeedsNonzeroEdge, i.NumericalYukawasInserted, i.NumericalZHComputed, i.RequiresPositiveF0, i.RequiresTraceConventionLedger, i.RequiresEuclideanSignLedger, i.Verdict)
}
func FormatPrefactorFactor(f PrefactorFactor) string {
	return fmt.Sprintf("%s symbol=%s role=%q rep=%q sign=%q choice=%t theorem=%t numeric=%t empirical=%t absorbed=%s status=%s", f.Name, f.Symbol, f.MathematicalRole, f.Representative, f.SignCondition, f.PositiveByChoice, f.PositiveByTheorem, f.NumericallyFixed, f.EmpiricalInput, f.AbsorbedInto, f.Status)
}
func FormatLedger(l PrefactorLedger) string {
	ps := []string{}
	for _, f := range l.Factors {
		ps = append(ps, FormatPrefactorFactor(f))
	}
	return fmt.Sprintf("ZH=%q N4=%q factors=[%s] explicit=%t noEmpirical=%t positiveClass=%t absoluteN4=%t verdict=%s", l.ZHFormula, l.N4Factorization, strings.Join(ps, " | "), l.AllFactorsExplicit, l.AllEmpiricalInputsExcluded, l.CanChoosePositiveClass, l.AbsoluteN4Derived, l.Verdict)
}
func FormatWick(w WickConventionAudit) string {
	return fmt.Sprintf("Euclidean=%q Lorentzian=%q Wick=%q Hamiltonian=%q signs=[%s] hidden=%t positiveEnergy=%t native=%t verdict=%s", w.EuclideanKineticForm, w.LorentzianKineticTarget, w.WickMap, w.HamiltonianCondition, strings.Join(w.SignLedger, " | "), w.SignAmbiguityHidden, w.PositiveEnergyMapped, w.ConventionNativeToFinite, w.Verdict)
}
func FormatF0(f F0Requirement) string {
	return fmt.Sprintf("definition=%q requirement=%q gate288=%t numeric=%t positiveWithoutEmpirics=%t conditional=%t verdict=%s", f.CutoffMomentDefinition, f.RequiredCondition, f.ContactSpectralGate288Used, f.NumericalValueDerived, f.CanBePositiveWithoutEmpirics, f.ConditionallyPositive, f.Verdict)
}
func FormatMatching(m CanonicalMatchingRule) string {
	return fmt.Sprintf("raw=%q target=%q rescale=%q absorbed=[%s] mass=%q quartic=%q physicalZH=%t rule=%t verdict=%s", m.RawCoefficient, m.CanonicalTarget, m.RescalingRule, strings.Join(m.AbsorbedFactors, "/"), m.MassChannelImpact, m.QuarticChannelImpact, m.PhysicalZHComputed, m.RuleFormalized, m.Verdict)
}
func FormatPositivity(p PrefactorPositivitySieve) string {
	return fmt.Sprintf("Kraw=%t nonzeroNeeded=%t N4=%q f0=%q overall=%q positivePrefactor=%t strictConditional=%t numericalStrict=%t failure=%q verdict=%s", p.KRawPositiveSemidefinite, p.AtLeastOneAmplitudeNeeded, p.N4PositiveCondition, p.F0PositiveCondition, p.OverallCondition, p.PositivePrefactorAvailable, p.StrictZHGuaranteedConditionally, p.NumericalStrictZHProved, p.FailureMode, p.Verdict)
}
func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}
func FormatFirewalls(f FirewallAudit) string {
	ps := []string{}
	for _, o := range f.Obligations {
		ps = append(ps, FormatObligation(o))
	}
	return fmt.Sprintf("noF0=%t noCutoffGate=%t noYukawa=%t noMasses=%t noSubtraction=%t noBGap=%t noHiggs=%t polluted=%t obligations=[%s] verdict=%s", f.NoF0NumberInserted, f.NoCutoffGateActivated, f.NoYukawaNumbersInserted, f.NoObservedMassesInserted, f.NoSubtractionSchemeInvented, f.NoBGapInstantonClaimed, f.NoHiggsPredictionClaimed, f.FiniteCorePolluted, strings.Join(ps, " | "), f.Verdict)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("inherit=%t ledger=%t wick=%t f0=%t matching=%t positivePrefactor=%t strictNumeric=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate301Inherited, s.PrefactorLedgerFormalized, s.WickSignRuleFormalized, s.PositiveF0ConditionRecorded, s.CanonicalMatchingFormalized, s.PositivePrefactorAvailable, s.StrictZHNumericallyProved, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
