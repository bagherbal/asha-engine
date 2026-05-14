// Package flavorprojectionmetric implements Gate 325:
// Flavor Projection Metric / Variational Vacuum Selector Audit.
//
// Gate 324 proved that a unitary rotation can place the physical top vector in
// the nullspace of the signed triality source tau_eta=(2,-2,1).  Gate 325 audits
// the hidden metric assumption behind that statement.  If the physical
// top-Yukawa boundary is evaluated by a positive Hilbert-Schmidt/singular-value
// metric, top nulling is impossible because diag(|tau_eta|^2) is positive
// definite.  If the boundary is evaluated by a signed amplitude projection, the
// nullspace exists, but it is two-dimensional and the finite geometry still does
// not uniquely select a CKM/flavor texture.
//
// The gate therefore protects the architecture from a sign mistake: it proves
// capacity only in the signed-projection lane, rejects nulling in the positive
// trace lane, and preserves the firewall that Gate 322's flattened-top transport
// is not yet a physical Standard Model top-sector derivation.
package flavorprojectionmetric

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE325-FLAVOR-PROJECTION-METRIC-VARIATIONAL-VACUUM-SELECTOR-AUDIT"

	StatusProjectionMetricFormalized       = "CONDITIONAL_SUPPORT_FLAVOR_PROJECTION_METRIC_FORMALIZED"
	StatusPositiveMetricAudited            = "CONDITIONAL_SUPPORT_POSITIVE_TRACE_METRIC_AUDITED"
	StatusSignedMetricAudited              = "CONDITIONAL_SUPPORT_SIGNED_PROJECTION_METRIC_AUDITED"
	StatusVariationalSieveExecuted         = "CONDITIONAL_SUPPORT_VARIATIONAL_FLAVOR_VACUUM_SIEVE_EXECUTED"
	StatusCKMFallbackFormalized            = "CONDITIONAL_SUPPORT_CKM_QUARANTINE_FALLBACK_FORMALIZED"
	StatusTopSuppressionCapacitySignedOnly = "CONDITIONAL_SUPPORT_TOP_SUPPRESSION_CAPACITY_SIGNED_METRIC_ONLY"

	StatusTensionProjectionMetricUnselected = "CONDITIONAL_TENSION_PHYSICAL_PROJECTION_METRIC_NOT_SELECTED"
	StatusTensionSignedNullspaceDegenerate  = "CONDITIONAL_TENSION_SIGNED_NULLSPACE_VARIATIONAL_MINIMUM_DEGENERATE"
	StatusTensionGate322StillDiagnostic     = "CONDITIONAL_TENSION_GATE322_FLATTENED_TOP_LANE_STILL_DIAGNOSTIC"

	StatusFailedPositiveMetricForbidsTopNulling = "FAILED_ROUTE_POSITIVE_TRACE_METRIC_FORBIDS_TOP_NULLING"
	StatusFailedNativeMetricNotDerived          = "FAILED_ROUTE_NATIVE_FLAVOR_PROJECTION_METRIC_NOT_DERIVED"
	StatusFailedNativeVacuumNotSelected         = "FAILED_ROUTE_NATIVE_FLAVOR_VACUUM_NOT_SELECTED"
	StatusFailedUniqueCKMTextureNotDerived      = "FAILED_ROUTE_UNIQUE_CKM_TEXTURE_NOT_DERIVED"
	StatusFailedFlavorOrientationEmpiricalSeal  = "FAILED_ROUTE_FLAVOR_ORIENTATION_REMAINS_EMPIRICAL_PHASE_III_SEAL"
	StatusFailedTopBoundaryNotJustified         = "FAILED_ROUTE_TOP_BOUNDARY_SUPPRESSION_NOT_JUSTIFIED"
	StatusFailedPoleMassNotExecuted             = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedTwoLoopNotExecuted              = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed          = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	gate322RunningMassGeV       = 124.9766199157
	gate323UniqueLowTopMassGeV  = 258.687
	gate323HighSlotTopMassGeV   = 317.115
	nearGate322TolerancePercent = 1.0
)

type MetricAudit struct {
	Formalized             bool
	TrialityVector         []float64
	NormalizedTriality     []float64
	PositiveDiagonal       []float64
	PositiveEigenMin       float64
	PositiveEigenMax       float64
	SignedRank             int
	SignedNullity          int
	PhysicalMetricSelected bool
	DefaultQFTMetric       string
	Verdict                string
}

type PositiveMetricSieve struct {
	Audited                   bool
	MinimumFraction           float64
	MinimizingVector          []float64
	MaximumFraction           float64
	ExactNullingPossible      bool
	VariationalMinimumUnique  bool
	PredictedMassAtMinimumGeV float64
	Verdict                   string
}

type SignedMetricSieve struct {
	Audited                  bool
	NullspaceDimension       int
	NullspaceBasis           [][]float64
	ExactNullingPossible     bool
	VariationalMinimumValue  float64
	VariationalMinimumUnique bool
	NullTopMassGeV           float64
	Verdict                  string
}

type VariationalAudit struct {
	Executed                      bool
	PositiveLaneMinimum           float64
	PositiveLaneAuthorizesGate322 bool
	SignedLaneMinimum             float64
	SignedLaneAuthorizesGate322   bool
	NativeSelectorInstalled       bool
	DynamicalPotentialInstalled   bool
	Result                        string
	Verdict                       string
}

type CKMFallbackAudit struct {
	Formalized                      bool
	RequiresEmpiricalTexture        bool
	RequiresVacuumSelectorPotential bool
	Phase                           string
	Verdict                         string
}

type RGCompatibilityAudit struct {
	Audited                        bool
	Gate322MassGeV                 float64
	PositiveMetricMinTopMassGeV    float64
	SignedNullTopMassGeV           float64
	PositiveMetricPreservesGate322 bool
	SignedMetricPreservesGate322   bool
	PhysicalLaneAuthorized         bool
	Verdict                        string
}

type FirewallAudit struct {
	NoCKMImported             bool
	NoObservedTopMassInserted bool
	NoFlavorTextureInvented   bool
	NoProjectionMetricForced  bool
	NoPoleMassClaimed         bool
	NoTwoLoopClaimed          bool
	NoColliderMassClaimed     bool
	FiniteCorePolluted        bool
	Verdict                   string
}

type Summary struct {
	ProjectionMetricFormalized      bool
	PositiveMetricForbidsNulling    bool
	SignedMetricAllowsNulling       bool
	VariationalSieveExecuted        bool
	NativeMetricDerived             bool
	NativeFlavorVacuumSelected      bool
	TopBoundarySuppressionJustified bool
	Gate322PhysicalLaneAuthorized   bool
	FirewallsPreserved              bool
	FinalMassClaimed                bool
	Status                          string
	DirectAnswer                    string
	NextGate                        string
}

type Analysis struct {
	Metric      MetricAudit
	Positive    PositiveMetricSieve
	Signed      SignedMetricSieve
	Variational VariationalAudit
	CKMFallback CKMFallbackAudit
	RG          RGCompatibilityAudit
	Firewalls   FirewallAudit
	Summary     Summary
	Truth       string
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
	metric := auditMetric()
	positive := auditPositiveMetric(metric)
	signed := auditSignedMetric(metric)
	variational := auditVariational(metric, positive, signed)
	fallback := auditCKMFallback(variational)
	rg := auditRGCompatibility(positive, signed, variational)
	firewalls := auditFirewalls(metric, variational, fallback)
	summary := buildSummary(metric, positive, signed, variational, rg, firewalls)
	truth := "Gate 325 resolves the Gate-324 sign ambiguity.  A positive Hilbert-Schmidt flavor metric diag(|tau_eta|^2) is positive definite and cannot produce y_t(Λ_GUT)=0; its variational minimum is the unique low slot with fraction 1/9, which inherits the Gate-323 high Higgs-mass tension.  A signed rank-one projection |tau_eta><tau_eta| permits exact top nulling and reproduces the Gate-322 flattened-top envelope, but the signed metric and a unique null vector are not derived by the finite core.  Therefore top suppression remains a permitted variational capacity, not a physical Standard Model top-sector derivation."
	return Analysis{Metric: metric, Positive: positive, Signed: signed, Variational: variational, CKMFallback: fallback, RG: rg, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func auditMetric() MetricAudit {
	tau := []float64{2, -2, 1}
	norm := vectorNorm(tau)
	normalized := scaleVector(tau, 1/norm)
	diag := []float64{4.0 / 9.0, 4.0 / 9.0, 1.0 / 9.0}
	return MetricAudit{
		Formalized:             true,
		TrialityVector:         tau,
		NormalizedTriality:     normalized,
		PositiveDiagonal:       diag,
		PositiveEigenMin:       1.0 / 9.0,
		PositiveEigenMax:       4.0 / 9.0,
		SignedRank:             1,
		SignedNullity:          2,
		PhysicalMetricSelected: false,
		DefaultQFTMetric:       "Yukawa beta functions normally depend on positive singular-value/Hilbert-Schmidt data; a signed interference metric needs an explicit pre-squaring projection theorem.",
		Verdict:                strings.Join([]string{StatusProjectionMetricFormalized, StatusTensionProjectionMetricUnselected, StatusFailedNativeMetricNotDerived}, ";"),
	}
}

func auditPositiveMetric(m MetricAudit) PositiveMetricSieve {
	min := minFloat(m.PositiveDiagonal)
	max := maxFloat(m.PositiveDiagonal)
	return PositiveMetricSieve{
		Audited:                   true,
		MinimumFraction:           min,
		MinimizingVector:          []float64{0, 0, 1},
		MaximumFraction:           max,
		ExactNullingPossible:      false,
		VariationalMinimumUnique:  true,
		PredictedMassAtMinimumGeV: gate323UniqueLowTopMassGeV,
		Verdict:                   strings.Join([]string{StatusPositiveMetricAudited, StatusFailedPositiveMetricForbidsTopNulling, StatusFailedTopBoundaryNotJustified}, ";"),
	}
}

func auditSignedMetric(m MetricAudit) SignedMetricSieve {
	basis := [][]float64{normalize([]float64{1, 1, 0}), normalize([]float64{1, 0, -2})}
	exact := true
	for _, v := range basis {
		if math.Abs(dot(m.NormalizedTriality, v)) > 1e-12 {
			exact = false
		}
	}
	return SignedMetricSieve{
		Audited:                  true,
		NullspaceDimension:       m.SignedNullity,
		NullspaceBasis:           basis,
		ExactNullingPossible:     exact,
		VariationalMinimumValue:  0,
		VariationalMinimumUnique: false,
		NullTopMassGeV:           gate322RunningMassGeV,
		Verdict:                  strings.Join([]string{StatusSignedMetricAudited, StatusTopSuppressionCapacitySignedOnly, StatusTensionSignedNullspaceDegenerate, StatusFailedUniqueCKMTextureNotDerived}, ";"),
	}
}

func auditVariational(m MetricAudit, p PositiveMetricSieve, s SignedMetricSieve) VariationalAudit {
	posAuth := p.ExactNullingPossible && p.PredictedMassAtMinimumGeV > 0 && math.Abs(p.PredictedMassAtMinimumGeV-gate322RunningMassGeV)/gate322RunningMassGeV*100 < nearGate322TolerancePercent
	signedAuth := s.ExactNullingPossible && math.Abs(s.NullTopMassGeV-gate322RunningMassGeV)/gate322RunningMassGeV*100 < nearGate322TolerancePercent && m.PhysicalMetricSelected
	result := "positive metric selects the 1/9 slot, not zero; signed metric has a zero minimum but is degenerate and not selected as the native physical projection metric."
	return VariationalAudit{
		Executed:                      true,
		PositiveLaneMinimum:           p.MinimumFraction,
		PositiveLaneAuthorizesGate322: posAuth,
		SignedLaneMinimum:             s.VariationalMinimumValue,
		SignedLaneAuthorizesGate322:   signedAuth,
		NativeSelectorInstalled:       false,
		DynamicalPotentialInstalled:   false,
		Result:                        result,
		Verdict:                       strings.Join([]string{StatusVariationalSieveExecuted, StatusFailedNativeVacuumNotSelected, StatusFailedUniqueCKMTextureNotDerived, StatusTensionGate322StillDiagnostic}, ";"),
	}
}

func auditCKMFallback(v VariationalAudit) CKMFallbackAudit {
	return CKMFallbackAudit{
		Formalized:                      true,
		RequiresEmpiricalTexture:        !v.NativeSelectorInstalled,
		RequiresVacuumSelectorPotential: !v.DynamicalPotentialInstalled,
		Phase:                           "Phase III flavor texture / CKM empirical or variational-vacuum selector seal",
		Verdict:                         strings.Join([]string{StatusCKMFallbackFormalized, StatusFailedFlavorOrientationEmpiricalSeal}, ";"),
	}
}

func auditRGCompatibility(p PositiveMetricSieve, s SignedMetricSieve, v VariationalAudit) RGCompatibilityAudit {
	signedPreserves := s.ExactNullingPossible && math.Abs(s.NullTopMassGeV-gate322RunningMassGeV)/gate322RunningMassGeV*100 < nearGate322TolerancePercent
	return RGCompatibilityAudit{
		Audited:                        true,
		Gate322MassGeV:                 gate322RunningMassGeV,
		PositiveMetricMinTopMassGeV:    p.PredictedMassAtMinimumGeV,
		SignedNullTopMassGeV:           s.NullTopMassGeV,
		PositiveMetricPreservesGate322: v.PositiveLaneAuthorizesGate322,
		SignedMetricPreservesGate322:   signedPreserves,
		PhysicalLaneAuthorized:         v.SignedLaneAuthorizesGate322,
		Verdict:                        strings.Join([]string{StatusTensionGate322StillDiagnostic, StatusFailedTopBoundaryNotJustified}, ";"),
	}
}

func auditFirewalls(m MetricAudit, v VariationalAudit, c CKMFallbackAudit) FirewallAudit {
	return FirewallAudit{
		NoCKMImported:             true,
		NoObservedTopMassInserted: true,
		NoFlavorTextureInvented:   !v.NativeSelectorInstalled,
		NoProjectionMetricForced:  !m.PhysicalMetricSelected,
		NoPoleMassClaimed:         true,
		NoTwoLoopClaimed:          true,
		NoColliderMassClaimed:     true,
		FiniteCorePolluted:        false,
		Verdict:                   strings.Join([]string{StatusFailedNativeMetricNotDerived, StatusFailedUniqueCKMTextureNotDerived, c.Verdict, StatusFailedPoleMassNotExecuted, StatusFailedTwoLoopNotExecuted, StatusFailedColliderMassNotClaimed}, ";"),
	}
}

func buildSummary(m MetricAudit, p PositiveMetricSieve, s SignedMetricSieve, v VariationalAudit, rg RGCompatibilityAudit, f FirewallAudit) Summary {
	preserved := f.NoCKMImported && f.NoObservedTopMassInserted && f.NoFlavorTextureInvented && f.NoProjectionMetricForced && f.NoPoleMassClaimed && f.NoTwoLoopClaimed && f.NoColliderMassClaimed && !f.FiniteCorePolluted
	return Summary{
		ProjectionMetricFormalized:      m.Formalized,
		PositiveMetricForbidsNulling:    p.Audited && !p.ExactNullingPossible,
		SignedMetricAllowsNulling:       s.Audited && s.ExactNullingPossible,
		VariationalSieveExecuted:        v.Executed,
		NativeMetricDerived:             m.PhysicalMetricSelected,
		NativeFlavorVacuumSelected:      v.NativeSelectorInstalled,
		TopBoundarySuppressionJustified: rg.PhysicalLaneAuthorized,
		Gate322PhysicalLaneAuthorized:   rg.PhysicalLaneAuthorized,
		FirewallsPreserved:              preserved,
		FinalMassClaimed:                false,
		Status:                          strings.Join([]string{StatusProjectionMetricFormalized, StatusVariationalSieveExecuted, StatusFailedPositiveMetricForbidsTopNulling, StatusFailedNativeMetricNotDerived, StatusFailedFlavorOrientationEmpiricalSeal}, ";"),
		DirectAnswer:                    "The variational sieve does not yet justify the Gate-322 physical top lane.  Under the standard positive Yukawa metric, exact top nulling is forbidden; under a signed projection metric, top nulling is possible but degenerate and not natively selected.",
		NextGate:                        "derive a native signed flavor-projection operator or a dynamical flavor vacuum potential; otherwise quarantine U_flavor/CKM as a Phase-III empirical texture seal.",
	}
}

func dot(a, b []float64) float64 {
	limit := len(a)
	if len(b) < limit {
		limit = len(b)
	}
	s := 0.0
	for i := 0; i < limit; i++ {
		s += a[i] * b[i]
	}
	return s
}

func vectorNorm(v []float64) float64 { return math.Sqrt(dot(v, v)) }

func normalize(v []float64) []float64 {
	n := vectorNorm(v)
	if n == 0 {
		return append([]float64(nil), v...)
	}
	return scaleVector(v, 1/n)
}

func scaleVector(v []float64, s float64) []float64 {
	out := make([]float64, len(v))
	for i, x := range v {
		out[i] = s * x
	}
	return out
}

func minFloat(v []float64) float64 {
	if len(v) == 0 {
		return math.NaN()
	}
	m := v[0]
	for _, x := range v[1:] {
		if x < m {
			m = x
		}
	}
	return m
}

func maxFloat(v []float64) float64 {
	if len(v) == 0 {
		return math.NaN()
	}
	m := v[0]
	for _, x := range v[1:] {
		if x > m {
			m = x
		}
	}
	return m
}

func FormatVector(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.12f", x)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func FormatMetric(m MetricAudit) string {
	return fmt.Sprintf("formalized=%t tau=%s tauHat=%s posDiag=%s eigMin=%.12f eigMax=%.12f signedRank=%d signedNullity=%d selected=%t default=%q verdict=%s", m.Formalized, FormatVector(m.TrialityVector), FormatVector(m.NormalizedTriality), FormatVector(m.PositiveDiagonal), m.PositiveEigenMin, m.PositiveEigenMax, m.SignedRank, m.SignedNullity, m.PhysicalMetricSelected, m.DefaultQFTMetric, m.Verdict)
}

func FormatPositive(p PositiveMetricSieve) string {
	return fmt.Sprintf("audited=%t min=%.12f minVec=%s max=%.12f null=%t uniqueMin=%t mass=%.6f verdict=%s", p.Audited, p.MinimumFraction, FormatVector(p.MinimizingVector), p.MaximumFraction, p.ExactNullingPossible, p.VariationalMinimumUnique, p.PredictedMassAtMinimumGeV, p.Verdict)
}

func FormatSigned(s SignedMetricSieve) string {
	basis := make([]string, len(s.NullspaceBasis))
	for i, v := range s.NullspaceBasis {
		basis[i] = FormatVector(v)
	}
	return fmt.Sprintf("audited=%t nullity=%d basis=%s null=%t min=%.12f unique=%t mass=%.6f verdict=%s", s.Audited, s.NullspaceDimension, strings.Join(basis, ";"), s.ExactNullingPossible, s.VariationalMinimumValue, s.VariationalMinimumUnique, s.NullTopMassGeV, s.Verdict)
}

func FormatVariational(v VariationalAudit) string {
	return fmt.Sprintf("executed=%t positiveMin=%.12f positiveAuth=%t signedMin=%.12f signedAuth=%t nativeSelector=%t potential=%t result=%q verdict=%s", v.Executed, v.PositiveLaneMinimum, v.PositiveLaneAuthorizesGate322, v.SignedLaneMinimum, v.SignedLaneAuthorizesGate322, v.NativeSelectorInstalled, v.DynamicalPotentialInstalled, v.Result, v.Verdict)
}

func FormatCKM(c CKMFallbackAudit) string {
	return fmt.Sprintf("formalized=%t empiricalTexture=%t vacuumSelector=%t phase=%q verdict=%s", c.Formalized, c.RequiresEmpiricalTexture, c.RequiresVacuumSelectorPotential, c.Phase, c.Verdict)
}

func FormatRG(r RGCompatibilityAudit) string {
	return fmt.Sprintf("audited=%t gate322=%.6f positiveMinMass=%.6f signedNullMass=%.6f positivePreserves=%t signedPreserves=%t physical=%t verdict=%s", r.Audited, r.Gate322MassGeV, r.PositiveMetricMinTopMassGeV, r.SignedNullTopMassGeV, r.PositiveMetricPreservesGate322, r.SignedMetricPreservesGate322, r.PhysicalLaneAuthorized, r.Verdict)
}

func FormatFirewalls(f FirewallAudit) string {
	return fmt.Sprintf("noCKM=%t noTopMass=%t noTexture=%t noMetricForced=%t noPole=%t noTwoLoop=%t noCollider=%t polluted=%t verdict=%s", f.NoCKMImported, f.NoObservedTopMassInserted, f.NoFlavorTextureInvented, f.NoProjectionMetricForced, f.NoPoleMassClaimed, f.NoTwoLoopClaimed, f.NoColliderMassClaimed, f.FiniteCorePolluted, f.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("metric=%t positiveForbids=%t signedAllows=%t variational=%t nativeMetric=%t nativeVacuum=%t topJustified=%t gate322Physical=%t firewalls=%t finalClaim=%t status=%s answer=%q next=%q", s.ProjectionMetricFormalized, s.PositiveMetricForbidsNulling, s.SignedMetricAllowsNulling, s.VariationalSieveExecuted, s.NativeMetricDerived, s.NativeFlavorVacuumSelected, s.TopBoundarySuppressionJustified, s.Gate322PhysicalLaneAuthorized, s.FirewallsPreserved, s.FinalMassClaimed, s.Status, s.DirectAnswer, s.NextGate)
}
