// Package cutoffmomentsource implements Gate 303:
// Cutoff Moment Source / Positive f0 Test-Function Class Audit.
//
// Gate 302 reduced scalar kinetic viability to the sign condition
// Z_H = N_4 f_0 K_H^raw with K_H^raw >= 0 and N_4 selectable positive.
// Gate 303 audits the mathematical source of f_0. It compares three lanes:
// a generic positive spectral-action test-function class, the sealed
// contact-spectral candidate zeta_contact(0)=7 from Gates 162/288, and a free
// phenomenological f_0. The gate deliberately proves only sign suitability and
// source obligations. It does not compute a Higgs mass, absolute Z_H, gauge
// couplings, branch selection, or B-gap instanton hierarchy.
package cutoffmomentsource

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE303-CUTOFF-MOMENT-SOURCE-POSITIVE-F0-TEST-FUNCTION-CLASS-AUDIT"

	StatusGate302Inherited                     = "CONDITIONAL_SUPPORT_GATE302_POSITIVE_F0_OBLIGATION_INHERITED"
	StatusGenericPositiveTestFunctionAudited   = "CONDITIONAL_SUPPORT_GENERIC_POSITIVE_TEST_FUNCTION_CLASS_AUDITED"
	StatusContactSpectralF0PositivePreflight   = "CONDITIONAL_SUPPORT_CONTACT_SPECTRAL_F0_EQUALS_7_POSITIVE_PREFLIGHT"
	StatusFreePhenomenologicalF0SieveCompleted = "CONDITIONAL_SUPPORT_FREE_PHENOMENOLOGICAL_F0_SIEVE_COMPLETED"
	StatusPositiveF0ClassFormalized            = "CONDITIONAL_SUPPORT_POSITIVE_F0_TEST_FUNCTION_CLASS_FORMALIZED"
	StatusSourceComparisonCompleted            = "CONDITIONAL_SUPPORT_CUTOFF_MOMENT_SOURCE_COMPARISON_COMPLETED"
	StatusFirewallsPreserved                   = "CONDITIONAL_SUPPORT_GATE303_NUMERICAL_FIREWALLS_PRESERVED"

	StatusFailedCutoffFunctionEqualsContactSpectrumNotDerived = "FAILED_ROUTE_CUTOFF_FUNCTION_EQUALS_CONTACT_SPECTRUM_NOT_DERIVED_AS_HEAT_KERNEL_THEOREM"
	StatusFailedFinalF0SourceNotUniquelySelected              = "FAILED_ROUTE_FINAL_F0_SOURCE_NOT_UNIQUELY_SELECTED"
	StatusFailedF0ScaleStillNotPhysicalPrediction             = "FAILED_ROUTE_F0_SCALE_STILL_NOT_PHYSICAL_PREDICTION"
	StatusFailedZHNumericalValueStillSealed                   = "FAILED_ROUTE_ZH_NUMERICAL_VALUE_STILL_SEALED"
	StatusFailedYukawaAmplitudesStillSealed                   = "FAILED_ROUTE_NUMERICAL_YUKAWA_AMPLITUDES_STILL_SEALED"
	StatusFailedHiggsMassQuarticStillFirewalled               = "FAILED_ROUTE_HIGGS_MASS_AND_QUARTIC_STILL_FIREWALLED"
	StatusFailedGaugeCouplingsAbsoluteStillFirewalled         = "FAILED_ROUTE_ABSOLUTE_GAUGE_COUPLINGS_STILL_FIREWALLED"
	StatusFailedBGAPInstantonStillSealed                      = "FAILED_ROUTE_BGAP_INSTANTON_ACTION_STILL_SEALED"
)

const (
	contactF0Exact = "7"
	contactF0      = 7
)

type Gate302Inheritance struct {
	PositivePrefactorLedgerFormalized bool
	KRawPositiveCarrierInherited      bool
	N4PositiveClassAvailable          bool
	F0PositiveRequired                bool
	F0NumericalValueDerived           bool
	CutoffGateActivated               bool
	NumericalZHComputed               bool
	YukawaNumbersInserted             bool
	Verdict                           string
}

type GenericTestFunctionAudit struct {
	SourceName           string
	FunctionClass        string
	MomentDefinition     string
	PositivityCondition  string
	SmoothnessCondition  string
	DecayCondition       string
	GuaranteesF0Positive bool
	FixesNumericalF0     bool
	ObservedInputUsed    bool
	PredictivePower      string
	Obligation           string
	Verdict              string
}

type ContactSpectralCutoffPreflight struct {
	SourceName                         string
	Gate162LedgerAvailable             bool
	Gate288CutoffIdentificationAudited bool
	ExactValue                         string
	IntegerValue                       int
	StrictlyPositive                   bool
	InternalAlgebraicSource            bool
	ObservedInputUsed                  bool
	MayBeActivatedAsSeal               bool
	ActivatedAsFinalSource             bool
	HeatKernelEqualityDerived          bool
	SatisfiesGate302SignRequirement    bool
	DoesNotSelectBranch                bool
	DoesNotDeriveHiggsPrediction       bool
	Obligation                         string
	Verdict                            string
}

type FreePhenomenologicalF0Sieve struct {
	SourceName                    string
	DomainRestriction             string
	GuaranteesF0PositiveIfImposed bool
	FixesNumericalF0              bool
	InternalPredictionLost        bool
	ExternalExperimentNeeded      bool
	PredictiveLosses              []string
	AdmissibleForStabilityOnly    bool
	Verdict                       string
}

type SourceCandidate struct {
	Name                        string
	Kind                        string
	SignGuarantee               string
	NumericalValue              string
	InternalToProject           bool
	ObservedInput               bool
	HeatKernelSourceDerived     bool
	CanSatisfyGate302           bool
	CanFixAbsoluteNormalization bool
	SelectedAsFinal             bool
	FirewallStatus              string
}

type SourceComparison struct {
	Candidates                []SourceCandidate
	AnyPositiveLaneAvailable  bool
	ContactLaneSatisfiesSign  bool
	GenericLaneSatisfiesSign  bool
	FreeLaneSatisfiesSign     bool
	UniqueFinalSourceSelected bool
	NoObservedInputRequired   bool
	BestCurrentPath           string
	Verdict                   string
}

type PositiveF0ClassSieve struct {
	RequiredCondition            string
	GenericClassCondition        string
	ContactSealCondition         string
	FreeParameterCondition       string
	StrictPositivityCanBeEnsured bool
	ContactValueCleanlySatisfies bool
	FinalNumericalF0Claimed      bool
	NumericalZHClaimed           bool
	HiggsPredictionClaimed       bool
	Verdict                      string
}

type RemainingObligation struct {
	Name, WhyRequired, Status string
	BlocksPrediction          bool
}

type FirewallAudit struct {
	NoObservedF0Inserted                  bool
	NoFinalCutoffSourceForced             bool
	NoYukawaNumbersInserted               bool
	NoNumericalZHComputed                 bool
	NoHiggsMassQuarticClaimed             bool
	NoGaugeCouplingAbsoluteClaimed        bool
	NoBGapInstantonClaimed                bool
	ContactValueUsedOnlyAsSealedPreflight bool
	FiniteCorePolluted                    bool
	Obligations                           []RemainingObligation
	Verdict                               string
}

type Summary struct {
	Gate302Inherited               bool
	GenericClassFormalized         bool
	ContactSealPositive            bool
	FreeF0SieveCompleted           bool
	PositiveF0ClassFormalized      bool
	UniqueFinalSourceSelected      bool
	NumericalF0Locked              bool
	NumericalZHComputed            bool
	PhysicalDynamicsDerived        bool
	FirewallPreserved              bool
	Status, DirectAnswer, NextGate string
}

type Analysis struct {
	Input      Gate302Inheritance
	Generic    GenericTestFunctionAudit
	Contact    ContactSpectralCutoffPreflight
	Free       FreePhenomenologicalF0Sieve
	Comparison SourceComparison
	Sieve      PositiveF0ClassSieve
	Firewalls  FirewallAudit
	Summary    Summary
	Truth      string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	i := inheritGate302()
	g := auditGenericTestFunction()
	c := auditContactSpectralCutoff()
	f := auditFreePhenomenologicalF0()
	cmp := compareSources(g, c, f)
	sieve := runPositiveF0ClassSieve(i, g, c, f, cmp)
	fw := auditFirewalls(i, c, cmp, sieve)
	s := buildSummary(i, g, c, f, cmp, sieve, fw)
	truth := "Gate 303 audits the source of the cutoff moment f_0 required by Gate 302. A generic admissible positive spectral-action test function guarantees f_0>0 as a class condition but leaves the scale free. The sealed contact-spectral candidate zeta_contact(0)=7 is internal, exact, and strictly positive, so it cleanly satisfies the Gate-302 sign obligation if explicitly activated as a sealed cutoff source. A free phenomenological f_0 can preserve stability only by imposing f_0>0 externally, losing internal normalization power. The gate formalizes positive f_0 source classes without selecting a final source or deriving numerical Z_H, Higgs coefficients, absolute gauge couplings, or B-gap instanton dynamics."
	return Analysis{Input: i, Generic: g, Contact: c, Free: f, Comparison: cmp, Sieve: sieve, Firewalls: fw, Summary: s, Truth: truth}, nil
}

func inheritGate302() Gate302Inheritance {
	return Gate302Inheritance{true, true, true, true, false, false, false, false, StatusGate302Inherited}
}

func auditGenericTestFunction() GenericTestFunctionAudit {
	return GenericTestFunctionAudit{
		SourceName:           "generic admissible spectral-action test function",
		FunctionClass:        "even/spectral profile f on non-negative Dirac eigenvalues, smooth or sufficiently regular for the heat-kernel expansion, non-negative on the spectrum, and nonzero in the a_4 moment/evaluation channel",
		MomentDefinition:     "f_0 is the dimension-four heat-kernel coefficient multiplier; in common four-dimensional conventions this is the zeroth/evaluation moment assigned to a_4",
		PositivityCondition:  "f_0>0; sufficient class condition: f(0)>0 or the convention-specific a_4 moment weight is strictly positive",
		SmoothnessCondition:  "smooth compact-support or rapidly decaying smoothed-step profile; discontinuous sharp cutoffs remain approximation data, not a proof object",
		DecayCondition:       "decay/compact support must make Tr(f(D_A/Lambda)) heat-kernel expandable without importing observed physics",
		GuaranteesF0Positive: true,
		FixesNumericalF0:     false,
		ObservedInputUsed:    false,
		PredictivePower:      "guarantees sign and admissibility but leaves absolute normalization and relative moment ratios unfixed",
		Obligation:           "select an explicit positive test-function class before numerical normalization claims",
		Verdict:              StatusGenericPositiveTestFunctionAudited,
	}
}

func auditContactSpectralCutoff() ContactSpectralCutoffPreflight {
	return ContactSpectralCutoffPreflight{
		SourceName:                         "sealed contact-spectral cutoff candidate",
		Gate162LedgerAvailable:             true,
		Gate288CutoffIdentificationAudited: true,
		ExactValue:                         "ζ_contact(0)=7",
		IntegerValue:                       contactF0,
		StrictlyPositive:                   contactF0 > 0,
		InternalAlgebraicSource:            true,
		ObservedInputUsed:                  false,
		MayBeActivatedAsSeal:               true,
		ActivatedAsFinalSource:             false,
		HeatKernelEqualityDerived:          false,
		SatisfiesGate302SignRequirement:    contactF0 > 0,
		DoesNotSelectBranch:                true,
		DoesNotDeriveHiggsPrediction:       true,
		Obligation:                         "prove or explicitly seal the identification cutoff moment f_0 := ζ_contact(0) before using 7 as the physical heat-kernel coefficient",
		Verdict:                            strings.Join([]string{StatusContactSpectralF0PositivePreflight, StatusFailedCutoffFunctionEqualsContactSpectrumNotDerived}, ";"),
	}
}

func auditFreePhenomenologicalF0() FreePhenomenologicalF0Sieve {
	losses := []string{
		"absolute scalar wave-function normalization remains arbitrary",
		"absolute gauge kinetic coefficients remain unpredicted",
		"cutoff-moment ratios cannot constrain the scalar potential",
		"internal contact-spectral normalization is not used as a prediction",
		"Higgs mass/quartic extraction remains a convention-plus-data fit rather than an algebraic output",
	}
	return FreePhenomenologicalF0Sieve{
		SourceName:                    "free phenomenological f_0",
		DomainRestriction:             "f_0 is allowed only in the open positive domain f_0>0 for scalar kinetic stability",
		GuaranteesF0PositiveIfImposed: true,
		FixesNumericalF0:              false,
		InternalPredictionLost:        true,
		ExternalExperimentNeeded:      true,
		PredictiveLosses:              losses,
		AdmissibleForStabilityOnly:    true,
		Verdict:                       strings.Join([]string{StatusFreePhenomenologicalF0SieveCompleted, StatusFailedF0ScaleStillNotPhysicalPrediction}, ";"),
	}
}

func compareSources(g GenericTestFunctionAudit, c ContactSpectralCutoffPreflight, f FreePhenomenologicalF0Sieve) SourceComparison {
	candidates := []SourceCandidate{
		{
			Name:                        "generic positive test-function class",
			Kind:                        "mathematical class condition",
			SignGuarantee:               g.PositivityCondition,
			NumericalValue:              "not fixed",
			InternalToProject:           true,
			ObservedInput:               g.ObservedInputUsed,
			HeatKernelSourceDerived:     true,
			CanSatisfyGate302:           g.GuaranteesF0Positive,
			CanFixAbsoluteNormalization: false,
			SelectedAsFinal:             false,
			FirewallStatus:              StatusGenericPositiveTestFunctionAudited,
		},
		{
			Name:                        "sealed contact-spectral candidate",
			Kind:                        "internal exact spectral seal candidate",
			SignGuarantee:               "ζ_contact(0)=7>0",
			NumericalValue:              contactF0Exact,
			InternalToProject:           c.InternalAlgebraicSource,
			ObservedInput:               c.ObservedInputUsed,
			HeatKernelSourceDerived:     c.HeatKernelEqualityDerived,
			CanSatisfyGate302:           c.SatisfiesGate302SignRequirement,
			CanFixAbsoluteNormalization: c.MayBeActivatedAsSeal,
			SelectedAsFinal:             c.ActivatedAsFinalSource,
			FirewallStatus:              c.Verdict,
		},
		{
			Name:                        "free phenomenological positive parameter",
			Kind:                        "external/open parameter domain",
			SignGuarantee:               f.DomainRestriction,
			NumericalValue:              "not fixed internally",
			InternalToProject:           false,
			ObservedInput:               f.ExternalExperimentNeeded,
			HeatKernelSourceDerived:     false,
			CanSatisfyGate302:           f.GuaranteesF0PositiveIfImposed,
			CanFixAbsoluteNormalization: false,
			SelectedAsFinal:             false,
			FirewallStatus:              f.Verdict,
		},
	}
	return SourceComparison{
		Candidates:                candidates,
		AnyPositiveLaneAvailable:  g.GuaranteesF0Positive || c.SatisfiesGate302SignRequirement || f.GuaranteesF0PositiveIfImposed,
		ContactLaneSatisfiesSign:  c.SatisfiesGate302SignRequirement,
		GenericLaneSatisfiesSign:  g.GuaranteesF0Positive,
		FreeLaneSatisfiesSign:     f.GuaranteesF0PositiveIfImposed,
		UniqueFinalSourceSelected: false,
		NoObservedInputRequired:   !g.ObservedInputUsed && !c.ObservedInputUsed,
		BestCurrentPath:           "retain the generic positive test-function class as the canonical sign theorem, and record ζ_contact(0)=7 as the strongest internal sealed candidate for a later explicit cutoff-source activation gate",
		Verdict:                   strings.Join([]string{StatusSourceComparisonCompleted, StatusFailedFinalF0SourceNotUniquelySelected}, ";"),
	}
}

func runPositiveF0ClassSieve(i Gate302Inheritance, g GenericTestFunctionAudit, c ContactSpectralCutoffPreflight, f FreePhenomenologicalF0Sieve, cmp SourceComparison) PositiveF0ClassSieve {
	ok := i.PositivePrefactorLedgerFormalized && i.F0PositiveRequired && cmp.AnyPositiveLaneAvailable && cmp.ContactLaneSatisfiesSign && cmp.GenericLaneSatisfiesSign && cmp.NoObservedInputRequired
	return PositiveF0ClassSieve{
		RequiredCondition:            "Gate 302 requires f_0>0 in Z_H=N_4 f_0 K_H^raw",
		GenericClassCondition:        g.PositivityCondition,
		ContactSealCondition:         "sealed candidate f_0=ζ_contact(0)=7 satisfies f_0>0 exactly",
		FreeParameterCondition:       f.DomainRestriction,
		StrictPositivityCanBeEnsured: ok,
		ContactValueCleanlySatisfies: c.SatisfiesGate302SignRequirement && c.StrictlyPositive && c.InternalAlgebraicSource && !c.ObservedInputUsed,
		FinalNumericalF0Claimed:      false,
		NumericalZHClaimed:           false,
		HiggsPredictionClaimed:       false,
		Verdict:                      strings.Join([]string{StatusPositiveF0ClassFormalized, StatusFailedFinalF0SourceNotUniquelySelected, StatusFailedZHNumericalValueStillSealed}, ";"),
	}
}

func auditFirewalls(i Gate302Inheritance, c ContactSpectralCutoffPreflight, cmp SourceComparison, s PositiveF0ClassSieve) FirewallAudit {
	obs := []RemainingObligation{
		{"final cutoff-source theorem", "needed to choose between generic positive class, sealed contact-spectral activation, or free parameter", StatusFailedFinalF0SourceNotUniquelySelected, true},
		{"contact cutoff equality theorem", "needed before treating ζ_contact(0)=7 as the physical heat-kernel coefficient rather than a sealed candidate", StatusFailedCutoffFunctionEqualsContactSpectrumNotDerived, true},
		{"absolute Z_H value", "requires final f_0 source, absolute N_4 convention, and nonzero numerical scalar amplitudes", StatusFailedZHNumericalValueStillSealed, true},
		{"Yukawa amplitude ledger", "K_H^raw strict magnitude depends on sealed Y_u,Y_d,Y_e,Y_ν amplitudes", StatusFailedYukawaAmplitudesStillSealed, true},
		{"Higgs mass/quartic extraction", "requires subtraction scheme, Z_H magnitude, scalar quadratic and quartic channels", StatusFailedHiggsMassQuarticStillFirewalled, true},
		{"absolute gauge couplings", "f_0 participates in a_4 gauge kinetic normalization but no absolute coupling is computed", StatusFailedGaugeCouplingsAbsoluteStillFirewalled, true},
		{"B-gap instanton action", "test-function positivity is polynomial heat-kernel bookkeeping and does not derive S_inst=(4/pi)/B_gap", StatusFailedBGAPInstantonStillSealed, true},
	}
	polluted := i.F0NumericalValueDerived || i.NumericalZHComputed || i.YukawaNumbersInserted || c.ActivatedAsFinalSource || s.FinalNumericalF0Claimed || s.NumericalZHClaimed || s.HiggsPredictionClaimed
	return FirewallAudit{
		NoObservedF0Inserted:                  cmp.NoObservedInputRequired,
		NoFinalCutoffSourceForced:             !cmp.UniqueFinalSourceSelected && !c.ActivatedAsFinalSource,
		NoYukawaNumbersInserted:               !i.YukawaNumbersInserted,
		NoNumericalZHComputed:                 !i.NumericalZHComputed && !s.NumericalZHClaimed,
		NoHiggsMassQuarticClaimed:             !s.HiggsPredictionClaimed,
		NoGaugeCouplingAbsoluteClaimed:        true,
		NoBGapInstantonClaimed:                true,
		ContactValueUsedOnlyAsSealedPreflight: c.StrictlyPositive && !c.ActivatedAsFinalSource && !c.ObservedInputUsed,
		FiniteCorePolluted:                    polluted,
		Obligations:                           obs,
		Verdict:                               strings.Join([]string{StatusFirewallsPreserved, StatusFailedFinalF0SourceNotUniquelySelected, StatusFailedCutoffFunctionEqualsContactSpectrumNotDerived, StatusFailedZHNumericalValueStillSealed, StatusFailedYukawaAmplitudesStillSealed, StatusFailedHiggsMassQuarticStillFirewalled, StatusFailedBGAPInstantonStillSealed}, ";"),
	}
}

func buildSummary(i Gate302Inheritance, g GenericTestFunctionAudit, c ContactSpectralCutoffPreflight, f FreePhenomenologicalF0Sieve, cmp SourceComparison, sieve PositiveF0ClassSieve, fw FirewallAudit) Summary {
	statuses := []string{
		StatusGate302Inherited,
		StatusGenericPositiveTestFunctionAudited,
		StatusContactSpectralF0PositivePreflight,
		StatusFreePhenomenologicalF0SieveCompleted,
		StatusPositiveF0ClassFormalized,
		StatusSourceComparisonCompleted,
		StatusFirewallsPreserved,
		StatusFailedCutoffFunctionEqualsContactSpectrumNotDerived,
		StatusFailedFinalF0SourceNotUniquelySelected,
		StatusFailedF0ScaleStillNotPhysicalPrediction,
		StatusFailedZHNumericalValueStillSealed,
		StatusFailedYukawaAmplitudesStillSealed,
		StatusFailedHiggsMassQuarticStillFirewalled,
		StatusFailedGaugeCouplingsAbsoluteStillFirewalled,
		StatusFailedBGAPInstantonStillSealed,
	}
	return Summary{
		Gate302Inherited:          i.PositivePrefactorLedgerFormalized && i.F0PositiveRequired && !i.F0NumericalValueDerived,
		GenericClassFormalized:    g.GuaranteesF0Positive && !g.FixesNumericalF0,
		ContactSealPositive:       c.SatisfiesGate302SignRequirement && c.IntegerValue == contactF0 && !c.ObservedInputUsed,
		FreeF0SieveCompleted:      f.GuaranteesF0PositiveIfImposed && f.InternalPredictionLost,
		PositiveF0ClassFormalized: sieve.StrictPositivityCanBeEnsured && sieve.ContactValueCleanlySatisfies,
		UniqueFinalSourceSelected: cmp.UniqueFinalSourceSelected,
		NumericalF0Locked:         false,
		NumericalZHComputed:       false,
		PhysicalDynamicsDerived:   false,
		FirewallPreserved:         !fw.FiniteCorePolluted && fw.NoFinalCutoffSourceForced && fw.NoYukawaNumbersInserted && fw.NoNumericalZHComputed && fw.ContactValueUsedOnlyAsSealedPreflight,
		Status:                    strings.Join(statuses, ";"),
		DirectAnswer:              "Gate 303 proves that the f_0 positivity obligation has mathematically valid source classes. A generic positive test-function class gives f_0>0 without fixing the value; the sealed contact-spectral value ζ_contact(0)=7 exactly satisfies the sign requirement; a free phenomenological f_0 is stable only when restricted to f_0>0 and loses internal predictive power. No final f_0 source or numerical dynamics is declared.",
		NextGate:                  "Gate 304 should audit whether the sealed contact-spectral cutoff can be promoted to a canonical heat-kernel source theorem by constructing an explicit positive test-function/profile or spectral measure whose a_4 moment equals ζ_contact(0)=7 without contaminating the continuum normalization or branch firewalls.",
	}
}

func FormatGate302Inheritance(i Gate302Inheritance) string {
	return fmt.Sprintf("prefactorLedger=%t Kraw=%t N4Positive=%t f0Required=%t f0Numeric=%t cutoffActivated=%t ZH=%t yukawas=%t verdict=%s", i.PositivePrefactorLedgerFormalized, i.KRawPositiveCarrierInherited, i.N4PositiveClassAvailable, i.F0PositiveRequired, i.F0NumericalValueDerived, i.CutoffGateActivated, i.NumericalZHComputed, i.YukawaNumbersInserted, i.Verdict)
}

func FormatGeneric(g GenericTestFunctionAudit) string {
	return fmt.Sprintf("source=%q class=%q moment=%q positivity=%q smooth=%q decay=%q guarantees=%t fixes=%t observed=%t predictive=%q obligation=%q verdict=%s", g.SourceName, g.FunctionClass, g.MomentDefinition, g.PositivityCondition, g.SmoothnessCondition, g.DecayCondition, g.GuaranteesF0Positive, g.FixesNumericalF0, g.ObservedInputUsed, g.PredictivePower, g.Obligation, g.Verdict)
}

func FormatContact(c ContactSpectralCutoffPreflight) string {
	return fmt.Sprintf("source=%q gate162=%t gate288=%t exact=%s value=%d positive=%t internal=%t observed=%t maySeal=%t activated=%t heatKernelEquality=%t sign=%t noBranch=%t noHiggs=%t obligation=%q verdict=%s", c.SourceName, c.Gate162LedgerAvailable, c.Gate288CutoffIdentificationAudited, c.ExactValue, c.IntegerValue, c.StrictlyPositive, c.InternalAlgebraicSource, c.ObservedInputUsed, c.MayBeActivatedAsSeal, c.ActivatedAsFinalSource, c.HeatKernelEqualityDerived, c.SatisfiesGate302SignRequirement, c.DoesNotSelectBranch, c.DoesNotDeriveHiggsPrediction, c.Obligation, c.Verdict)
}

func FormatFree(f FreePhenomenologicalF0Sieve) string {
	return fmt.Sprintf("source=%q domain=%q guarantees=%t fixes=%t predictionLost=%t experiment=%t losses=[%s] stabilityOnly=%t verdict=%s", f.SourceName, f.DomainRestriction, f.GuaranteesF0PositiveIfImposed, f.FixesNumericalF0, f.InternalPredictionLost, f.ExternalExperimentNeeded, strings.Join(f.PredictiveLosses, " | "), f.AdmissibleForStabilityOnly, f.Verdict)
}

func FormatCandidate(c SourceCandidate) string {
	return fmt.Sprintf("%s kind=%q sign=%q value=%q internal=%t observed=%t heatKernelSource=%t gate302=%t absolute=%t selected=%t firewall=%s", c.Name, c.Kind, c.SignGuarantee, c.NumericalValue, c.InternalToProject, c.ObservedInput, c.HeatKernelSourceDerived, c.CanSatisfyGate302, c.CanFixAbsoluteNormalization, c.SelectedAsFinal, c.FirewallStatus)
}

func FormatComparison(c SourceComparison) string {
	parts := []string{}
	for _, x := range c.Candidates {
		parts = append(parts, FormatCandidate(x))
	}
	return fmt.Sprintf("candidates=[%s] anyPositive=%t contact=%t generic=%t free=%t unique=%t noObserved=%t best=%q verdict=%s", strings.Join(parts, " || "), c.AnyPositiveLaneAvailable, c.ContactLaneSatisfiesSign, c.GenericLaneSatisfiesSign, c.FreeLaneSatisfiesSign, c.UniqueFinalSourceSelected, c.NoObservedInputRequired, c.BestCurrentPath, c.Verdict)
}

func FormatSieve(s PositiveF0ClassSieve) string {
	return fmt.Sprintf("required=%q generic=%q contact=%q free=%q ensured=%t contactClean=%t finalF0=%t numericalZH=%t higgs=%t verdict=%s", s.RequiredCondition, s.GenericClassCondition, s.ContactSealCondition, s.FreeParameterCondition, s.StrictPositivityCanBeEnsured, s.ContactValueCleanlySatisfies, s.FinalNumericalF0Claimed, s.NumericalZHClaimed, s.HiggsPredictionClaimed, s.Verdict)
}

func FormatObligation(o RemainingObligation) string {
	return fmt.Sprintf("%s required=%q status=%s blocks=%t", o.Name, o.WhyRequired, o.Status, o.BlocksPrediction)
}

func FormatFirewalls(f FirewallAudit) string {
	obs := []string{}
	for _, o := range f.Obligations {
		obs = append(obs, FormatObligation(o))
	}
	return fmt.Sprintf("noObservedF0=%t noFinalSource=%t noYukawa=%t noZH=%t noHiggs=%t noGauge=%t noBGap=%t contactPreflight=%t polluted=%t obligations=[%s] verdict=%s", f.NoObservedF0Inserted, f.NoFinalCutoffSourceForced, f.NoYukawaNumbersInserted, f.NoNumericalZHComputed, f.NoHiggsMassQuarticClaimed, f.NoGaugeCouplingAbsoluteClaimed, f.NoBGapInstantonClaimed, f.ContactValueUsedOnlyAsSealedPreflight, f.FiniteCorePolluted, strings.Join(obs, " | "), f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("gate302=%t generic=%t contact=%t free=%t class=%t unique=%t f0Locked=%t ZH=%t dynamics=%t firewall=%t status=%s answer=%q next=%q", s.Gate302Inherited, s.GenericClassFormalized, s.ContactSealPositive, s.FreeF0SieveCompleted, s.PositiveF0ClassFormalized, s.UniqueFinalSourceSelected, s.NumericalF0Locked, s.NumericalZHComputed, s.PhysicalDynamicsDerived, s.FirewallPreserved, s.Status, s.DirectAnswer, s.NextGate)
}
