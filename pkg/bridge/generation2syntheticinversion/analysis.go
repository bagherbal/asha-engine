// Package generation2syntheticinversion implements Gate 468:
// Common-Scale Synthetic Inversion Run / Uncertainty Propagation Harness.
package generation2syntheticinversion

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE468-COMMON-SCALE-SYNTHETIC-INVERSION-UNCERTAINTY-PROPAGATION-HARNESS"

	StatusGate467Inherited                = "CONDITIONAL_SUPPORT_GATE467_COMMON_SCALE_LEDGER_INHERITED"
	StatusSyntheticLedgerAccepted         = "CONDITIONAL_SUPPORT_SYNTHETIC_COMMON_SCALE_LEDGER_ACCEPTED"
	StatusSymbolicInverseExecuted         = "CONDITIONAL_SUPPORT_SYNTHETIC_RAY_INVERSION_EXECUTED"
	StatusIntervalPropagationExecuted     = "CONDITIONAL_SUPPORT_UNCERTAINTY_PROPAGATION_EXECUTED"
	StatusDUDSyntheticComputed            = "CONDITIONAL_SUPPORT_SYNTHETIC_DUD_BRIDGE_ONLY_COMPUTED"
	StatusSyntheticInversionValidated     = "CONDITIONAL_SUPPORT_COMMON_SCALE_SYNTHETIC_INVERSION_VALIDATED"
	StatusFirewallPreserved               = "CONDITIONAL_SUPPORT_13_MODULI_FIREWALL_PRESERVED"
	StatusFailedObservedDataRejected      = "FAILED_ROUTE_OBSERVED_DATA_REJECTED_IN_SYNTHETIC_INVERSION"
	StatusFailedMissingRankCompleteLedger = "FAILED_ROUTE_SYNTHETIC_INVERSION_REQUIRES_RANK_COMPLETE_LEDGER"
	StatusFailedProjectiveDomainRejected  = "FAILED_ROUTE_SYNTHETIC_INVERSION_PROJECTIVE_DOMAIN_REJECTED"
	StatusFailedPhaseDomainRejected       = "FAILED_ROUTE_SYNTHETIC_INVERSION_PHASE_DOMAIN_REJECTED"
	StatusFailedCausticRejected           = "FAILED_ROUTE_SYNTHETIC_INVERSION_CAUSTIC_REJECTED"
	StatusFailedBranchTagRejected         = "FAILED_ROUTE_SYNTHETIC_INVERSION_BRANCH_TAG_REJECTED"
	StatusFailedUncertaintyMissing        = "FAILED_ROUTE_SYNTHETIC_INVERSION_UNCERTAINTY_MISSING"
	StatusFailedCabibboAsRayRejected      = "FAILED_ROUTE_CABIBBO_USED_AS_SYNTHETIC_RAY_INPUT_REJECTED"
	StatusFailedNativePromotionRejected   = "FAILED_ROUTE_SYNTHETIC_DUD_NATIVE_PROMOTION_REJECTED"
	StatusFailedCKMNativePrediction       = "FAILED_ROUTE_SYNTHETIC_RESIDUAL_IS_NOT_CKM_PREDICTION"
)

const (
	NativeFlavorDim = 13
	KXYCoeffDim     = 9
)

type Inheritance struct {
	Executed, Gate444KGenForced, Gate445TriangleForced, Gate456InverseAvailable, Gate459BranchTagsRequired, Gate464DUDSocketAvailable, Gate465AirlockFailClosed, Gate466MassOnlyObstruction, Gate467CommonScaleLedger, Gate467RequiresISpecIK, Gate467RequiresBranchTags, Gate467RequiresUncertainty, Gate467DUDComputableIfNumeric, Gate467DidNotComputeDUD, NativeRegistryClean bool
	Verdict                                                                                                                                                                                                                                                                                                                                                                       string
}
type SyntheticComparator struct {
	Sector                                                                                                                                                            string
	IK, ISpec                                                                                                                                                         float64
	SigmaCP, C3Sheet                                                                                                                                                  int
	DeltaIK, DeltaISpec                                                                                                                                               float64
	CommonScale, CommonScheme, Source, SourceVersion, UncertaintyModel                                                                                                string
	BridgeOnly, SyntheticOnly, ObservedData, CabibboAsRayInput, NativePromotionClaim, HasISpec, HasIK, HasBranchTag, HasUncertainty, CommonScaleScheme, Dimensionless bool
}
type Ray struct {
	Sector                                                               string
	Alpha, CosThreePhi, Phi, IK, ISpec                                   float64
	SigmaCP, C3Sheet                                                     int
	AlphaMin, AlphaMax, PhiMin, PhiMax                                   float64
	InsideDomain, AtCaustic, BridgeOnly, SyntheticOnly, ExportsNativeRay bool
	Verdict, Reason                                                      string
}
type Distance struct {
	DeltaAlpha, DeltaPhi, DUD, DUDMin, DUDMax                                                                                   float64
	UncertaintyPropagated, BridgeOnly, SyntheticOnly, CabibboCompared, CKMMatrixConstructed, CKMEntryComputed, NativePrediction bool
	Verdict, Reason                                                                                                             string
}
type Case struct {
	Name            string
	U, D            SyntheticComparator
	Accepted        bool
	URay, DRay      Ray
	Distance        Distance
	Verdict, Reason string
	Failures        []string
}
type Harness struct {
	Executed                                                                                                                                                                                                                                                                                                                                                                            bool
	Cases                                                                                                                                                                                                                                                                                                                                                                               []Case
	AcceptedSyntheticCases, RejectedCases                                                                                                                                                                                                                                                                                                                                               int
	ValidSyntheticDUDComputed, UncertaintyPropagationExecuted, ObservedDataRejected, MissingRankLedgerRejected, ProjectiveDomainRejected, PhaseDomainRejected, CausticRejected, BranchTagRejected, MissingUncertaintyRejected, CabibboRayInputRejected, NativePromotionRejected, NoCKMMatrixConstructed, NoCKMEntryComputed, NoNativePredictionExported, AllAcceptedBridgeOnlySynthetic bool
	Verdict, Reason                                                                                                                                                                                                                                                                                                                                                                     string
}
type Firewall struct {
	Executed, SyntheticCoordinatesNative, SyntheticDUDNative, CKMNativePrediction, CKMMatrixConstructed, CKMEntryComputed, ObservedMassesImported, ObservedCKMImported, CabibboUsedAsRayInput, NativeRegistryWritten, KGenStillForced, XTriangleStillForced, YPhaseStillQuarantined, SectorCoefficientsStillSealed bool
	NativeFlavorDimAfter, KXYCoeffDimAfter                                                                                                                                                                                                                                                                         int
	Verdict, Reason                                                                                                                                                                                                                                                                                                string
}
type NextStep struct {
	Gate                       int
	Title, Reason, PrimaryTask string
}
type Analysis struct {
	Inheritance Inheritance
	Harness     Harness
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
	a.Harness = buildHarness()
	a.Firewall = buildFirewall(a)
	a.Next = buildNext()
	a.Truth = truth(a)
	if err := validate(a); err != nil {
		return a, err
	}
	return a, nil
}
func buildInheritance() Inheritance {
	return Inheritance{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, StatusGate467Inherited}
}
func buildHarness() Harness {
	cases := []Case{EvaluateCase("valid synthetic rank-complete u-d ledger", canonicalU(), canonicalD()), EvaluateCase("observed row rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.ObservedData = true }), canonicalD()), EvaluateCase("missing IK rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.HasIK = false }), canonicalD()), EvaluateCase("projective domain rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.IK = 1 }), canonicalD()), EvaluateCase("phase domain rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.ISpec = 9 }), canonicalD()), EvaluateCase("caustic rejected", causticComparator("u"), canonicalD()), EvaluateCase("branch tag rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.HasBranchTag = false; x.SigmaCP = 0 }), canonicalD()), EvaluateCase("missing uncertainty rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.HasUncertainty = false; x.DeltaIK = 0; x.DeltaISpec = 0 }), canonicalD()), EvaluateCase("Cabibbo as ray input rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.CabibboAsRayInput = true }), canonicalD()), EvaluateCase("native promotion rejected", mutate(canonicalU(), func(x *SyntheticComparator) { x.NativePromotionClaim = true }), canonicalD())}
	h := Harness{Executed: true, Cases: cases, NoCKMMatrixConstructed: true, NoCKMEntryComputed: true, NoNativePredictionExported: true, AllAcceptedBridgeOnlySynthetic: true}
	for _, c := range cases {
		if c.Accepted {
			h.AcceptedSyntheticCases++
			if c.Distance.DUD > 0 && c.Distance.UncertaintyPropagated && c.Distance.BridgeOnly && c.Distance.SyntheticOnly && !c.Distance.CabibboCompared && !c.Distance.NativePrediction {
				h.ValidSyntheticDUDComputed = true
				h.UncertaintyPropagationExecuted = true
			}
			if !c.Distance.BridgeOnly || !c.Distance.SyntheticOnly || c.Distance.CKMMatrixConstructed || c.Distance.CKMEntryComputed || c.Distance.NativePrediction {
				h.AllAcceptedBridgeOnlySynthetic = false
			}
		} else {
			h.RejectedCases++
			for _, f := range c.Failures {
				switch f {
				case StatusFailedObservedDataRejected:
					h.ObservedDataRejected = true
				case StatusFailedMissingRankCompleteLedger:
					h.MissingRankLedgerRejected = true
				case StatusFailedProjectiveDomainRejected:
					h.ProjectiveDomainRejected = true
				case StatusFailedPhaseDomainRejected:
					h.PhaseDomainRejected = true
				case StatusFailedCausticRejected:
					h.CausticRejected = true
				case StatusFailedBranchTagRejected:
					h.BranchTagRejected = true
				case StatusFailedUncertaintyMissing:
					h.MissingUncertaintyRejected = true
				case StatusFailedCabibboAsRayRejected:
					h.CabibboRayInputRejected = true
				case StatusFailedNativePromotionRejected:
					h.NativePromotionRejected = true
				}
			}
		}
	}
	h.Verdict = StatusSyntheticInversionValidated
	h.Reason = "one complete synthetic u-d ledger computes bridge-only alpha/phi rays, propagates comparator uncertainties to a d_ud interval, and every observed/incomplete/unsafe/native-promotion route fails closed"
	return h
}
func canonicalU() SyntheticComparator {
	ik, is := comparatorsFromAlphaPhi(1.00, 0.40)
	return SyntheticComparator{Sector: "u", IK: ik, ISpec: is, SigmaCP: 1, C3Sheet: 0, DeltaIK: 0.004, DeltaISpec: 0.003, CommonScale: "synthetic-Mstar", CommonScheme: "synthetic-MSbar", Source: "synthetic-redacted-ledger", SourceVersion: "gate468-dry-run-v1", UncertaintyModel: "box interval", BridgeOnly: true, SyntheticOnly: true, HasISpec: true, HasIK: true, HasBranchTag: true, HasUncertainty: true, CommonScaleScheme: true, Dimensionless: true}
}
func canonicalD() SyntheticComparator {
	ik, is := comparatorsFromAlphaPhi(1.22, 0.73)
	return SyntheticComparator{Sector: "d", IK: ik, ISpec: is, SigmaCP: 1, C3Sheet: 0, DeltaIK: 0.005, DeltaISpec: 0.003, CommonScale: "synthetic-Mstar", CommonScheme: "synthetic-MSbar", Source: "synthetic-redacted-ledger", SourceVersion: "gate468-dry-run-v1", UncertaintyModel: "box interval", BridgeOnly: true, SyntheticOnly: true, HasISpec: true, HasIK: true, HasBranchTag: true, HasUncertainty: true, CommonScaleScheme: true, Dimensionless: true}
}
func causticComparator(sector string) SyntheticComparator {
	ik, is := comparatorsFromAlphaPhi(1, 0)
	c := canonicalU()
	c.Sector = sector
	c.IK = ik
	c.ISpec = is
	c.SigmaCP = 1
	return c
}
func mutate(x SyntheticComparator, fn func(*SyntheticComparator)) SyntheticComparator {
	fn(&x)
	return x
}
func comparatorsFromAlphaPhi(alpha, phi float64) (float64, float64) {
	return alpha / math.Sqrt(alpha*alpha+3), 2 * math.Cos(3*phi) / math.Pow(alpha*alpha+3, 1.5)
}
func EvaluateCase(name string, u, d SyntheticComparator) Case {
	c := Case{Name: name, U: u, D: d}
	var fs []string
	fs = append(fs, validateComparator(u)...)
	fs = append(fs, validateComparator(d)...)
	if len(fs) > 0 {
		c.Failures = unique(fs)
		c.Verdict = strings.Join(c.Failures, ";")
		c.Reason = "one or both synthetic sector ledgers failed the rank-complete bridge-only preconditions"
		return c
	}
	ur, ue := invert(u)
	dr, de := invert(d)
	if ue != "" {
		fs = append(fs, ue)
	}
	if de != "" {
		fs = append(fs, de)
	}
	if len(fs) > 0 {
		c.URay = ur
		c.DRay = dr
		c.Failures = unique(fs)
		c.Verdict = strings.Join(c.Failures, ";")
		c.Reason = "Gate456 inverse rejected at least one sector"
		return c
	}
	dist := distance(ur, dr)
	c.Accepted = true
	c.URay = ur
	c.DRay = dr
	c.Distance = dist
	c.Verdict = StatusDUDSyntheticComputed
	c.Reason = "rank-complete synthetic ledgers invert to bridge rays and evaluate d_ud with interval propagation; no CKM matrix or observed Cabibbo comparison is performed"
	return c
}
func validateComparator(x SyntheticComparator) []string {
	var fs []string
	if x.ObservedData {
		fs = append(fs, StatusFailedObservedDataRejected)
	}
	if !x.HasIK || !x.HasISpec || !x.CommonScaleScheme || !x.Dimensionless || !x.BridgeOnly || !x.SyntheticOnly {
		fs = append(fs, StatusFailedMissingRankCompleteLedger)
	}
	if !x.HasBranchTag || (x.SigmaCP != -1 && x.SigmaCP != 1) || x.C3Sheet < 0 || x.C3Sheet > 2 {
		fs = append(fs, StatusFailedBranchTagRejected)
	}
	if !x.HasUncertainty || x.DeltaIK <= 0 || x.DeltaISpec <= 0 || x.UncertaintyModel == "" {
		fs = append(fs, StatusFailedUncertaintyMissing)
	}
	if x.CabibboAsRayInput {
		fs = append(fs, StatusFailedCabibboAsRayRejected)
	}
	if x.NativePromotionClaim {
		fs = append(fs, StatusFailedNativePromotionRejected)
	}
	return fs
}
func invert(x SyntheticComparator) (Ray, string) {
	r := Ray{Sector: x.Sector, IK: x.IK, ISpec: x.ISpec, SigmaCP: x.SigmaCP, C3Sheet: x.C3Sheet, BridgeOnly: true, SyntheticOnly: true}
	if math.Abs(x.IK) >= 1 {
		r.Verdict = StatusFailedProjectiveDomainRejected
		r.Reason = "I_K must lie in (-1,1)"
		return r, StatusFailedProjectiveDomainRejected
	}
	alpha := math.Sqrt(3) * x.IK / math.Sqrt(1-x.IK*x.IK)
	cos3 := (3 * math.Sqrt(3) / 2) * x.ISpec / math.Pow(1-x.IK*x.IK, 1.5)
	if math.Abs(cos3) > 1+1e-12 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Verdict = StatusFailedPhaseDomainRejected
		r.Reason = "derived cos(3phi) left [-1,1]"
		return r, StatusFailedPhaseDomainRejected
	}
	cos3 = clamp(cos3, -1, 1)
	phi := (float64(x.SigmaCP)*math.Acos(cos3) + 2*math.Pi*float64(x.C3Sheet)) / 3
	if math.Abs(math.Sin(3*phi)) < 1e-9 {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Phi = phi
		r.InsideDomain = true
		r.AtCaustic = true
		r.Verdict = StatusFailedCausticRejected
		r.Reason = "sin(3phi)=0 caustic"
		return r, StatusFailedCausticRejected
	}
	amin, amax, pmin, pmax, err := intervalRay(x)
	if err != "" {
		r.Alpha = alpha
		r.CosThreePhi = cos3
		r.Phi = phi
		r.InsideDomain = true
		r.Verdict = err
		r.Reason = "uncertainty box leaves valid inverse domain"
		return r, err
	}
	r.Alpha = alpha
	r.CosThreePhi = cos3
	r.Phi = phi
	r.AlphaMin = amin
	r.AlphaMax = amax
	r.PhiMin = pmin
	r.PhiMax = pmax
	r.InsideDomain = true
	r.Verdict = StatusSymbolicInverseExecuted
	r.Reason = "Gate456 inverse applied to a synthetic rank-complete sector ledger"
	return r, ""
}
func intervalRay(x SyntheticComparator) (float64, float64, float64, float64, string) {
	amin, amax, pmin, pmax := math.Inf(1), math.Inf(-1), math.Inf(1), math.Inf(-1)
	for _, ik := range []float64{x.IK - x.DeltaIK, x.IK + x.DeltaIK} {
		for _, is := range []float64{x.ISpec - x.DeltaISpec, x.ISpec + x.DeltaISpec} {
			if math.Abs(ik) >= 1 {
				return 0, 0, 0, 0, StatusFailedProjectiveDomainRejected
			}
			alpha := math.Sqrt(3) * ik / math.Sqrt(1-ik*ik)
			cos3 := (3 * math.Sqrt(3) / 2) * is / math.Pow(1-ik*ik, 1.5)
			if math.Abs(cos3) > 1+1e-12 {
				return 0, 0, 0, 0, StatusFailedPhaseDomainRejected
			}
			cos3 = clamp(cos3, -1, 1)
			phi := (float64(x.SigmaCP)*math.Acos(cos3) + 2*math.Pi*float64(x.C3Sheet)) / 3
			if math.Abs(math.Sin(3*phi)) < 1e-9 {
				return 0, 0, 0, 0, StatusFailedCausticRejected
			}
			amin = math.Min(amin, alpha)
			amax = math.Max(amax, alpha)
			pmin = math.Min(pmin, phi)
			pmax = math.Max(pmax, phi)
		}
	}
	return amin, amax, pmin, pmax, ""
}
func distance(u, d Ray) Distance {
	da := d.Alpha - u.Alpha
	dp := wrapPi(d.Phi - u.Phi)
	central := math.Sqrt(da*da + 4*math.Pow(math.Sin(dp/2), 2))
	mn, mx := math.Inf(1), math.Inf(-1)
	for _, au := range []float64{u.AlphaMin, u.AlphaMax} {
		for _, ad := range []float64{d.AlphaMin, d.AlphaMax} {
			for _, pu := range []float64{u.PhiMin, u.PhiMax} {
				for _, pd := range []float64{d.PhiMin, d.PhiMax} {
					v := math.Sqrt(math.Pow(ad-au, 2) + 4*math.Pow(math.Sin(wrapPi(pd-pu)/2), 2))
					mn = math.Min(mn, v)
					mx = math.Max(mx, v)
				}
			}
		}
	}
	return Distance{DeltaAlpha: da, DeltaPhi: dp, DUD: central, DUDMin: mn, DUDMax: mx, UncertaintyPropagated: true, BridgeOnly: true, SyntheticOnly: true, Verdict: StatusDUDSyntheticComputed, Reason: "synthetic u-d projective ray distance only; no Cabibbo/CKM comparator invoked"}
}
func buildFirewall(a Analysis) Firewall {
	return Firewall{Executed: true, KGenStillForced: a.Inheritance.Gate444KGenForced, XTriangleStillForced: a.Inheritance.Gate445TriangleForced, YPhaseStillQuarantined: true, SectorCoefficientsStillSealed: true, NativeFlavorDimAfter: NativeFlavorDim, KXYCoeffDimAfter: KXYCoeffDim, Verdict: StatusFirewallPreserved, Reason: "Gate468 computes only a synthetic bridge coordinate and interval; it writes no native ray, no CKM entry, and no theorem-registry observable"}
}
func buildNext() NextStep {
	return NextStep{469, "Observed Complete Comparator Dry-Run / Airlock Numerical Trial", "Gate468 proves the socket works for rank-complete synthetic ledgers; the next empirical step may run only if a real common-scale u/d ledger supplies I_spec, I_K, branch tags, and uncertainties without using CKM as an input.", "admit a fully provenanced observed comparator ledger through the Gate465 airlock, compute d_ud as bridge-only, and compare to Cabibbo solely as a residual target"}
}
func validate(a Analysis) error {
	if !a.Inheritance.Executed || !a.Inheritance.Gate467CommonScaleLedger || !a.Inheritance.Gate467RequiresISpecIK || !a.Inheritance.Gate467RequiresBranchTags || !a.Inheritance.Gate467RequiresUncertainty || !a.Inheritance.Gate467DUDComputableIfNumeric || !a.Inheritance.Gate467DidNotComputeDUD || !a.Inheritance.NativeRegistryClean {
		return fmt.Errorf("Gate468 inheritance incomplete")
	}
	if !a.Harness.Executed || a.Harness.AcceptedSyntheticCases != 1 || a.Harness.RejectedCases < 9 || !a.Harness.ValidSyntheticDUDComputed || !a.Harness.UncertaintyPropagationExecuted || !a.Harness.ObservedDataRejected || !a.Harness.MissingRankLedgerRejected || !a.Harness.ProjectiveDomainRejected || !a.Harness.PhaseDomainRejected || !a.Harness.CausticRejected || !a.Harness.BranchTagRejected || !a.Harness.MissingUncertaintyRejected || !a.Harness.CabibboRayInputRejected || !a.Harness.NativePromotionRejected || !a.Harness.NoCKMMatrixConstructed || !a.Harness.NoCKMEntryComputed || !a.Harness.NoNativePredictionExported || !a.Harness.AllAcceptedBridgeOnlySynthetic {
		return fmt.Errorf("synthetic inversion harness did not accept/reject expected routes")
	}
	if !a.Firewall.Executed || a.Firewall.SyntheticCoordinatesNative || a.Firewall.SyntheticDUDNative || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.CKMEntryComputed || a.Firewall.ObservedMassesImported || a.Firewall.ObservedCKMImported || a.Firewall.CabibboUsedAsRayInput || a.Firewall.NativeRegistryWritten || !a.Firewall.KGenStillForced || !a.Firewall.XTriangleStillForced || !a.Firewall.YPhaseStillQuarantined || !a.Firewall.SectorCoefficientsStillSealed || a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || a.Firewall.KXYCoeffDimAfter != KXYCoeffDim {
		return fmt.Errorf("native firewall violated by synthetic inversion")
	}
	return nil
}
func truth(a Analysis) string {
	if a.Harness.ValidSyntheticDUDComputed && a.Harness.UncertaintyPropagationExecuted && !a.Firewall.CKMNativePrediction && !a.Firewall.ObservedMassesImported {
		return "Gate 468 proves the Gate467 data product is sufficient in principle: a rank-complete synthetic u/d ledger can be inverted to bridge-only cylinder coordinates and propagated to a d_ud interval. This is a socket validation, not a CKM prediction; observed masses, CKM entries, and native coefficient values remain absent."
	}
	return "Gate 468 failed to validate the synthetic inversion socket."
}
func clamp(x, lo, hi float64) float64 {
	if x < lo {
		return lo
	}
	if x > hi {
		return hi
	}
	return x
}
func wrapPi(x float64) float64 {
	for x <= -math.Pi {
		x += 2 * math.Pi
	}
	for x > math.Pi {
		x -= 2 * math.Pi
	}
	return x
}
func unique(xs []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, x := range xs {
		if !seen[x] {
			seen[x] = true
			out = append(out, x)
		}
	}
	return out
}
func fmtFloat(x float64) string { return fmt.Sprintf("%.9g", x) }
func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branch_tags=%t socket=%t airlock=%t gate466_mass_only_obstruction=%t gate467_ledger=%t requires_I_spec_I_K=%t requires_branch=%t requires_uncertainty=%t computable_if_numeric=%t did_not_compute=%t native_clean=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseAvailable, x.Gate459BranchTagsRequired, x.Gate464DUDSocketAvailable, x.Gate465AirlockFailClosed, x.Gate466MassOnlyObstruction, x.Gate467CommonScaleLedger, x.Gate467RequiresISpecIK, x.Gate467RequiresBranchTags, x.Gate467RequiresUncertainty, x.Gate467DUDComputableIfNumeric, x.Gate467DidNotComputeDUD, x.NativeRegistryClean, x.Verdict)
}
func FormatComparator(x SyntheticComparator) string {
	return fmt.Sprintf("sector=%s I_K=%s I_spec=%s sigma_CP=%d n_C3=%d dI_K=%s dI_spec=%s scale=%s scheme=%s source=%s uncertainty=%s bridge_only=%t synthetic=%t observed=%t cabibbo_as_ray=%t native_claim=%t", x.Sector, fmtFloat(x.IK), fmtFloat(x.ISpec), x.SigmaCP, x.C3Sheet, fmtFloat(x.DeltaIK), fmtFloat(x.DeltaISpec), x.CommonScale, x.CommonScheme, x.Source, x.UncertaintyModel, x.BridgeOnly, x.SyntheticOnly, x.ObservedData, x.CabibboAsRayInput, x.NativePromotionClaim)
}
func FormatRay(x Ray) string {
	return fmt.Sprintf("sector=%s alpha=%s cos3phi=%s phi=%s alpha_interval=[%s,%s] phi_interval=[%s,%s] domain=%t caustic=%t bridge_only=%t synthetic=%t native_ray=%t verdict=%s", x.Sector, fmtFloat(x.Alpha), fmtFloat(x.CosThreePhi), fmtFloat(x.Phi), fmtFloat(x.AlphaMin), fmtFloat(x.AlphaMax), fmtFloat(x.PhiMin), fmtFloat(x.PhiMax), x.InsideDomain, x.AtCaustic, x.BridgeOnly, x.SyntheticOnly, x.ExportsNativeRay, x.Verdict)
}
func FormatDistance(x Distance) string {
	return fmt.Sprintf("Delta_alpha=%s Delta_phi=%s d_ud=%s interval=[%s,%s] uncertainty=%t bridge_only=%t synthetic=%t cabibbo_compared=%t ckm_matrix=%t ckm_entry=%t native=%t verdict=%s", fmtFloat(x.DeltaAlpha), fmtFloat(x.DeltaPhi), fmtFloat(x.DUD), fmtFloat(x.DUDMin), fmtFloat(x.DUDMax), x.UncertaintyPropagated, x.BridgeOnly, x.SyntheticOnly, x.CabibboCompared, x.CKMMatrixConstructed, x.CKMEntryComputed, x.NativePrediction, x.Verdict)
}
func FormatHarness(x Harness) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d valid_dud=%t uncertainty=%t observed_rejected=%t missing_rank=%t projective_rejected=%t phase_rejected=%t caustic_rejected=%t branch_rejected=%t uncertainty_missing=%t cabibbo_as_ray=%t native_promotion=%t no_ckm_matrix=%t no_ckm_entry=%t no_native=%t bridge_synthetic=%t verdict=%s reason=%s", x.Executed, x.AcceptedSyntheticCases, x.RejectedCases, x.ValidSyntheticDUDComputed, x.UncertaintyPropagationExecuted, x.ObservedDataRejected, x.MissingRankLedgerRejected, x.ProjectiveDomainRejected, x.PhaseDomainRejected, x.CausticRejected, x.BranchTagRejected, x.MissingUncertaintyRejected, x.CabibboRayInputRejected, x.NativePromotionRejected, x.NoCKMMatrixConstructed, x.NoCKMEntryComputed, x.NoNativePredictionExported, x.AllAcceptedBridgeOnlySynthetic, x.Verdict, x.Reason)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t coords_native=%t dud_native=%t ckm_native=%t ckm_matrix=%t ckm_entry=%t observed_masses=%t observed_ckm=%t cabibbo_as_ray=%t native_write=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.SyntheticCoordinatesNative, x.SyntheticDUDNative, x.CKMNativePrediction, x.CKMMatrixConstructed, x.CKMEntryComputed, x.ObservedMassesImported, x.ObservedCKMImported, x.CabibboUsedAsRayInput, x.NativeRegistryWritten, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}
func statuses() []string {
	return []string{StatusGate467Inherited, StatusSyntheticLedgerAccepted, StatusSymbolicInverseExecuted, StatusIntervalPropagationExecuted, StatusDUDSyntheticComputed, StatusSyntheticInversionValidated, StatusFirewallPreserved, StatusFailedObservedDataRejected, StatusFailedMissingRankCompleteLedger, StatusFailedProjectiveDomainRejected, StatusFailedPhaseDomainRejected, StatusFailedCausticRejected, StatusFailedBranchTagRejected, StatusFailedUncertaintyMissing, StatusFailedCabibboAsRayRejected, StatusFailedNativePromotionRejected, StatusFailedCKMNativePrediction}
}
func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 468 Registry Audit — Common-Scale Synthetic Inversion Run / Uncertainty Propagation Harness\n\n## Verdict\n\n`" + StatusSyntheticInversionValidated + "`\n\nGate 468 exercises the full ASHA cylinder socket on synthetic rank-complete u/d comparator ledgers. It computes bridge-only `(alpha,phi)` rays, propagates input uncertainty boxes, and returns a synthetic `d_ud` interval. It deliberately imports no observed PDG, CKM, PMNS, Yukawa, or mass values.\n\n")
	b.WriteString("## Inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Formulae executed\n\n```text\nalpha = sqrt(3) I_K / sqrt(1-I_K^2)\ncos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\nphi = (sigma_CP arccos(cos(3phi)) + 2pi n_C3)/3\nd_ud = sqrt((alpha_d-alpha_u)^2 + 4 sin^2((phi_d-phi_u)/2))\n```\n\n")
	b.WriteString("## Harness\n\n" + FormatHarness(a.Harness) + "\n\n| Case | Accepted | Verdict | `d_ud` / reason |\n|---|---:|---|---|\n")
	for _, c := range a.Harness.Cases {
		val := c.Reason
		if c.Accepted {
			val = FormatDistance(c.Distance)
		}
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(val)))
	}
	b.WriteString("\n")
	if len(a.Harness.Cases) > 0 && a.Harness.Cases[0].Accepted {
		c := a.Harness.Cases[0]
		b.WriteString("## Accepted synthetic dry-run detail\n\n- U ray: " + FormatRay(c.URay) + "\n- D ray: " + FormatRay(c.DRay) + "\n- Distance: " + FormatDistance(c.Distance) + "\n\n")
	}
	b.WriteString("## Native firewall proof\n\n" + FormatFirewall(a.Firewall) + "\n\nThe computed synthetic `d_ud` is a bridge socket validation only. It is not `V_us`, not a CKM entry, not a CKM matrix, not a physical prediction, and not a native ASHA theorem. Observed values remain outside this gate.\n\n## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n" + FormatNext(a.Next) + "\n\n## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
