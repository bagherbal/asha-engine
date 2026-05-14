// Package spectralmomentledger implements Gate 344:
// Complete Spectral Moment Ledger / Cosmological Constant from Triple Hierarchy.
//
// Gate 343 showed that the Einstein-Hilbert channel fixes the product f2 Λ²,
// not f2 alone.  Gate 344 therefore treats the spectral moments as physical
// channel-products:
//
//	f0                         dimensionless gauge/kinetic channel
//	f2 Λ²                      gravitational channel
//	f4 Λ⁴ a0_eff               cosmological/vacuum channel
//
// The gate then audits whether the hierarchy rule from Gate 342 extends one
// more step to the cosmological constant.  It computes the moment ratios and
// rejects arbitrary exponential extensions when they fail to reach the observed
// dark-energy order of magnitude.
package spectralmomentledger

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE344-COMPLETE-SPECTRAL-MOMENT-LEDGER-COSMOLOGICAL-CONSTANT-TRIPLE-HIERARCHY"

	StatusGate343Inherited              = "CONDITIONAL_SUPPORT_GATE343_F2_LAMBDA_PRODUCT_INHERITED"
	StatusMomentLedgerFormalized        = "CONDITIONAL_SUPPORT_COMPLETE_SPECTRAL_MOMENT_LEDGER_FORMALIZED"
	StatusF2ProductInvariantAccepted    = "CONDITIONAL_SUPPORT_F2_LAMBDA_PRODUCT_ACCEPTED_AS_PHYSICAL_INVARIANT"
	StatusGaugeGravityRatioComputed     = "CONDITIONAL_SUPPORT_GAUGE_GRAVITY_MOMENT_RATIO_COMPUTED"
	StatusCosmologicalF4Audited         = "CONDITIONAL_SUPPORT_COSMOLOGICAL_F4_CHANNEL_AUDITED"
	StatusDarkEnergyTargetExtracted     = "CONDITIONAL_SUPPORT_DARK_ENERGY_SCALE_TARGET_EXTRACTED"
	StatusCosmologicalCandidatesAudited = "CONDITIONAL_SUPPORT_COSMOLOGICAL_SUPPRESSION_CANDIDATES_AUDITED"

	StatusTensionF2NotNeededButLambdaStillGauge = "CONDITIONAL_TENSION_F2_SEPARATION_UNNEEDED_BUT_CUTOFF_CONVENTION_REMAINS"
	StatusTensionDoubleHierarchyTooLarge        = "CONDITIONAL_TENSION_DOUBLE_HIERARCHY_TOO_LARGE_FOR_COSMOLOGICAL_CONSTANT"
	StatusTensionElectroweakVacuumTooLarge      = "CONDITIONAL_TENSION_ELECTROWEAK_VACUUM_SCALE_TOO_LARGE"
	StatusTensionRequiredExponentUnnatural      = "CONDITIONAL_TENSION_REQUIRED_COSMOLOGICAL_EXPONENT_NOT_CANONICAL"

	StatusFailedCosmologicalConstantNotDerived     = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_NOT_DERIVED"
	StatusFailedF4Lambda4NotLocked                 = "FAILED_ROUTE_F4_LAMBDA4_MOMENT_NOT_LOCKED"
	StatusFailedA0VacuumMultiplicityNotDerived     = "FAILED_ROUTE_A0_VACUUM_MULTIPLICITY_NOT_DERIVED"
	StatusFailedArbitraryExponentExtensionRejected = "FAILED_ROUTE_ARBITRARY_EXPONENT_EXTENSION_REJECTED"
	StatusFailedVacuumRenormalizationNotDerived    = "FAILED_ROUTE_VACUUM_RENORMALIZATION_SCHEME_NOT_DERIVED"
)

const (
	inheritedHighestGate = 343

	nGen                      = 3.0
	f0Contact                 = 7.0
	electroweakVEVGeV         = 246.22
	unreducedPlanckGeV        = 1.220890e19
	observedCosmologicalRatio = 1.0e-122 // nominal rho_Lambda / M_P^4 comparison target.
)

type Inputs struct {
	HighestInheritedGate      int
	NGen                      float64
	STop                      float64
	VEVGeV                    float64
	UnreducedPlanckGeV        float64
	HierarchyRho              float64
	ObservedHierarchy         float64
	ObservedCosmologicalRatio float64
	Status                    string
}

type SpectralMoment struct {
	Name    string
	Channel string
	Formula string
	Value   float64
	Units   string
	Derived bool
	Caveat  string
	Status  string
}

type MomentLedger struct {
	Gauge        SpectralMoment
	Gravity      SpectralMoment
	Cosmological SpectralMoment
	Status       string
}

type MomentRatio struct {
	Name           string
	Formula        string
	Value          float64
	Log10Value     float64
	Interpretation string
	Status         string
}

type CosmologicalTarget struct {
	TargetRatioToMP4                 float64
	TargetF4Lambda4AssumingA0OneGeV4 float64
	TargetF4Lambda4OverV4            float64
	RhoFourth                        float64
	ElectroweakVacuumExcessFactor    float64
	RequiredHalfActionCount          float64
	RequiredTopActionCount           float64
	Status                           string
}

type SuppressionCandidate struct {
	Name                    string
	Formula                 string
	RatioToMP4              float64
	Log10Ratio              float64
	RelativeToObserved      float64
	Log10RelativeToObserved float64
	Promoted                bool
	Reason                  string
}

type CosmologicalAudit struct {
	Candidates []SuppressionCandidate
	Best       SuppressionCandidate
	Derived    bool
	Status     string
}

type Firewall struct {
	F4Lambda4Locked              bool
	A0Derived                    bool
	VacuumRenormalizationDerived bool
	CosmologicalConstantDerived  bool
	Explanation                  string
	Status                       string
}

type Summary struct {
	DirectAnswer            string
	GaugeGravityMomentRatio string
	CosmologicalTarget      string
	CandidateVerdict        string
	NextGate                string
	Status                  string
}

type Analysis struct {
	Inputs            Inputs
	Ledger            MomentLedger
	GaugeGravityRatio MomentRatio
	Target            CosmologicalTarget
	Cosmological      CosmologicalAudit
	Firewall          Firewall
	Summary           Summary
	Truth             string
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
	inputs := compileInputs()
	ledger := compileMomentLedger(inputs)
	ratio := computeGaugeGravityRatio(inputs, ledger)
	target := extractCosmologicalTarget(inputs)
	cosmo := auditCosmologicalCandidates(inputs, target)
	firewall := compileFirewall(cosmo)
	summary := compileSummary(ratio, target, cosmo)
	truth := "Gate 344 accepts the Gate 343 lesson: the gravitational channel fixes the physical product f2Λ² rather than f2 alone.  The complete spectral moment ledger therefore has f0=7, f2Λ²=(π/64)M_P², and an unresolved cosmological product f4Λ⁴a0_eff.  The gauge-gravity moment ratio is enormous and exactly computable from the Pfaffian hierarchy.  However, extending the hierarchy one step to the cosmological constant does not reach the observed ~10^-122 M_P^4 scale: electroweak vacuum scaling is ~10^-67, double half-action/topological candidates are ~10^-34, and the required suppression corresponds to a noncanonical ~7.11 half-actions.  Thus f4Λ⁴ and vacuum renormalization remain firewalled."
	return Analysis{Inputs: inputs, Ledger: ledger, GaugeGravityRatio: ratio, Target: target, Cosmological: cosmo, Firewall: firewall, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	sTop := 8 * math.Pi * math.Pi
	rho := math.Pow(2, nGen/2) * math.Exp(-sTop/2)
	observed := electroweakVEVGeV / unreducedPlanckGeV
	return Inputs{
		HighestInheritedGate:      inheritedHighestGate,
		NGen:                      nGen,
		STop:                      sTop,
		VEVGeV:                    electroweakVEVGeV,
		UnreducedPlanckGeV:        unreducedPlanckGeV,
		HierarchyRho:              rho,
		ObservedHierarchy:         observed,
		ObservedCosmologicalRatio: observedCosmologicalRatio,
		Status:                    StatusGate343Inherited,
	}
}

func compileMomentLedger(i Inputs) MomentLedger {
	f2Lambda := math.Pi / 64 * i.UnreducedPlanckGeV * i.UnreducedPlanckGeV
	return MomentLedger{
		Gauge:        SpectralMoment{Name: "f0", Channel: "a4 gauge/scalar kinetic/quartic", Formula: "f0=7", Value: f0Contact, Units: "dimensionless", Derived: true, Caveat: "contact spectral cutoff promotion", Status: StatusMomentLedgerFormalized},
		Gravity:      SpectralMoment{Name: "f2Λ²", Channel: "a2 Einstein-Hilbert", Formula: "f2Λ²=(π/64)M_P²", Value: f2Lambda, Units: "GeV²", Derived: true, Caveat: "product invariant; f2 and Λ are not separately physical without a cutoff selector", Status: StatusF2ProductInvariantAccepted},
		Cosmological: SpectralMoment{Name: "f4Λ⁴a0_eff", Channel: "a0 cosmological/vacuum", Formula: "unresolved; compare to ρ_Λ≈10^-122 M_P⁴", Value: math.NaN(), Units: "GeV⁴", Derived: false, Caveat: "requires vacuum multiplicity a0_eff and vacuum renormalization", Status: StatusFailedF4Lambda4NotLocked},
		Status:       StatusMomentLedgerFormalized,
	}
}

func computeGaugeGravityRatio(i Inputs, ledger MomentLedger) MomentRatio {
	value := ledger.Gravity.Value / (ledger.Gauge.Value * i.VEVGeV * i.VEVGeV)
	formulaValue := math.Pi / (64 * f0Contact) / (i.HierarchyRho * i.HierarchyRho)
	// The two formulas should agree up to the small difference between predicted hierarchy and the numerical Planck input.
	interpretation := fmt.Sprintf("direct=%.12e; hierarchy_formula=π/(448ρ²)=%.12e", value, formulaValue)
	return MomentRatio{
		Name:           "(f2Λ²)/(f0 v²)",
		Formula:        "π/(448)·(M_P/v)² = π/(3584)·exp(8π²) using ρ²=2³e^{-8π²}",
		Value:          formulaValue,
		Log10Value:     math.Log10(formulaValue),
		Interpretation: interpretation,
		Status:         StatusGaugeGravityRatioComputed,
	}
}

func extractCosmologicalTarget(i Inputs) CosmologicalTarget {
	mp4 := math.Pow(i.UnreducedPlanckGeV, 4)
	targetGeV4 := i.ObservedCosmologicalRatio * mp4
	v4 := math.Pow(i.VEVGeV, 4)
	rho4 := math.Pow(i.HierarchyRho, 4)
	excess := rho4 / i.ObservedCosmologicalRatio
	requiredHalfActions := -math.Log(i.ObservedCosmologicalRatio) / (i.STop / 2)
	requiredTopActions := -math.Log(i.ObservedCosmologicalRatio) / i.STop
	return CosmologicalTarget{
		TargetRatioToMP4:                 i.ObservedCosmologicalRatio,
		TargetF4Lambda4AssumingA0OneGeV4: targetGeV4,
		TargetF4Lambda4OverV4:            targetGeV4 / v4,
		RhoFourth:                        rho4,
		ElectroweakVacuumExcessFactor:    excess,
		RequiredHalfActionCount:          requiredHalfActions,
		RequiredTopActionCount:           requiredTopActions,
		Status:                           StatusDarkEnergyTargetExtracted,
	}
}

func auditCosmologicalCandidates(i Inputs, t CosmologicalTarget) CosmologicalAudit {
	rho := i.HierarchyRho
	candidates := []SuppressionCandidate{
		cosCandidate("electroweak vacuum v^4/M_P^4", "ρ^4", math.Pow(rho, 4), i.ObservedCosmologicalRatio, false, "natural if vacuum energy is set by electroweak scale; still about 55 orders too large"),
		cosCandidate("single half-action", "exp(-S_top/2)=exp(-4π²)", math.Exp(-i.STop/2), i.ObservedCosmologicalRatio, false, "hierarchy factor component; about 104 orders too large"),
		cosCandidate("double half-action / full topological action", "exp(-S_top)=exp(-8π²)", math.Exp(-i.STop), i.ObservedCosmologicalRatio, false, "user-proposed one-more-step exponential; about 88 orders too large"),
		cosCandidate("squared Pfaffian hierarchy", "ρ²=2³exp(-8π²)", rho*rho, i.ObservedCosmologicalRatio, false, "square of the derived v/M_P hierarchy; about 88 orders too large"),
		cosCandidate("four topological actions", "exp(-4S_top)", math.Exp(-4*i.STop), i.ObservedCosmologicalRatio, false, "close in exponent class but uses an arbitrary action multiplicity not derived by the finite core"),
		cosCandidate("required target", "10^-122", i.ObservedCosmologicalRatio, i.ObservedCosmologicalRatio, false, "identity target; not an independent derivation"),
	}
	best := candidates[0]
	for _, c := range candidates[1:] {
		if c.Name == "required target" {
			continue
		}
		if math.Abs(c.Log10RelativeToObserved) < math.Abs(best.Log10RelativeToObserved) {
			best = c
		}
	}
	return CosmologicalAudit{Candidates: candidates, Best: best, Derived: false, Status: StatusCosmologicalCandidatesAudited}
}

func cosCandidate(name, formula string, value, observed float64, promoted bool, reason string) SuppressionCandidate {
	rel := value / observed
	return SuppressionCandidate{Name: name, Formula: formula, RatioToMP4: value, Log10Ratio: math.Log10(value), RelativeToObserved: rel, Log10RelativeToObserved: math.Log10(rel), Promoted: promoted, Reason: reason}
}

func compileFirewall(c CosmologicalAudit) Firewall {
	return Firewall{F4Lambda4Locked: false, A0Derived: false, VacuumRenormalizationDerived: false, CosmologicalConstantDerived: false, Explanation: "The ledger identifies f4Λ⁴a0_eff as the last spectral-moment product, but no native theorem selects its value. Candidate exponential extensions either overshoot by many orders or require arbitrary action powers; vacuum subtraction/renormalization is still mandatory.", Status: StatusFailedCosmologicalConstantNotDerived}
}

func compileSummary(r MomentRatio, t CosmologicalTarget, c CosmologicalAudit) Summary {
	return Summary{
		DirectAnswer:            "f2 need not be separated from Λ; f2Λ² is the physical gravitational moment product.",
		GaugeGravityMomentRatio: fmt.Sprintf("%s = %.12e (log10 %.6f)", r.Name, r.Value, r.Log10Value),
		CosmologicalTarget:      fmt.Sprintf("ρ_Λ/M_P⁴≈%.1e; required half-action count=%.6f; required S_top count=%.6f", t.TargetRatioToMP4, t.RequiredHalfActionCount, t.RequiredTopActionCount),
		CandidateVerdict:        fmt.Sprintf("best non-identity candidate=%s gives %.3e M_P⁴, log10 relative=%+.3f", c.Best.Name, c.Best.RatioToMP4, c.Best.Log10RelativeToObserved),
		NextGate:                "Derive a vacuum-energy subtraction/selection theorem or a native f4Λ⁴a0_eff invariant; do not extend exponent powers by fitting.",
		Status:                  StatusCosmologicalF4Audited,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Ledger.Status,
		a.Ledger.Gauge.Status,
		a.Ledger.Gravity.Status,
		a.GaugeGravityRatio.Status,
		a.Target.Status,
		a.Cosmological.Status,
		StatusCosmologicalF4Audited,
		StatusTensionF2NotNeededButLambdaStillGauge,
		StatusTensionDoubleHierarchyTooLarge,
		StatusTensionElectroweakVacuumTooLarge,
		StatusTensionRequiredExponentUnnatural,
		StatusFailedCosmologicalConstantNotDerived,
		StatusFailedF4Lambda4NotLocked,
		StatusFailedA0VacuumMultiplicityNotDerived,
		StatusFailedArbitraryExponentExtensionRejected,
		StatusFailedVacuumRenormalizationNotDerived,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; N_gen=%.0f; S_top=%.15f; v=%.12f GeV; M_P=%.12e GeV; rho_pred=%.15e; rho_obs=%.15e; rho_lambda_obs=%.3e; status=%s", i.HighestInheritedGate, i.NGen, i.STop, i.VEVGeV, i.UnreducedPlanckGeV, i.HierarchyRho, i.ObservedHierarchy, i.ObservedCosmologicalRatio, i.Status)
}

func FormatMoment(m SpectralMoment) string {
	return fmt.Sprintf("%s[%s]: %s = %.12e %s; derived=%t; caveat=%s; status=%s", m.Name, m.Channel, m.Formula, m.Value, m.Units, m.Derived, m.Caveat, m.Status)
}

func FormatLedger(l MomentLedger) string {
	return strings.Join([]string{FormatMoment(l.Gauge), FormatMoment(l.Gravity), FormatMoment(l.Cosmological), "status=" + l.Status}, "\n")
}

func FormatRatio(r MomentRatio) string {
	return fmt.Sprintf("%s; formula=%s; value=%.12e; log10=%.9f; interpretation=%s; status=%s", r.Name, r.Formula, r.Value, r.Log10Value, r.Interpretation, r.Status)
}

func FormatTarget(t CosmologicalTarget) string {
	return fmt.Sprintf("target=%.3e M_P^4; target_GeV4=%.12e; target/v^4=%.12e; rho^4=%.12e; electroweak_excess=%.12e; required_half_actions=%.9f; required_top_actions=%.9f; status=%s", t.TargetRatioToMP4, t.TargetF4Lambda4AssumingA0OneGeV4, t.TargetF4Lambda4OverV4, t.RhoFourth, t.ElectroweakVacuumExcessFactor, t.RequiredHalfActionCount, t.RequiredTopActionCount, t.Status)
}

func FormatCandidate(c SuppressionCandidate) string {
	return fmt.Sprintf("%s: %s = %.12e M_P^4; log10=%.6f; relative_to_obs=%.12e; log10_relative=%+.6f; promoted=%t; reason=%s", c.Name, c.Formula, c.RatioToMP4, c.Log10Ratio, c.RelativeToObserved, c.Log10RelativeToObserved, c.Promoted, c.Reason)
}

func FormatCosmologicalAudit(a CosmologicalAudit) string {
	parts := make([]string, 0, len(a.Candidates)+1)
	for _, c := range a.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	parts = append(parts, fmt.Sprintf("best=%s; derived=%t; status=%s", FormatCandidate(a.Best), a.Derived, a.Status))
	return strings.Join(parts, "\n")
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("f4_locked=%t; a0=%t; vacuum_renorm=%t; cosmological_constant=%t; explanation=%s; status=%s", f.F4Lambda4Locked, f.A0Derived, f.VacuumRenormalizationDerived, f.CosmologicalConstantDerived, f.Explanation, f.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("direct=%s; gauge_gravity=%s; cosmological_target=%s; candidate=%s; next=%s; status=%s", s.DirectAnswer, s.GaugeGravityMomentRatio, s.CosmologicalTarget, s.CandidateVerdict, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
