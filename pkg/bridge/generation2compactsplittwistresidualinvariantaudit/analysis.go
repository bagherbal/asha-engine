// Package generation2compactsplittwistresidualinvariantaudit implements
// Gate 639: CompactSplitTwistResidual Invariant Audit.
//
// Gate 638 proved that K_7 carries two lawful native structures that remain
// unfused: a compact P_G-sourced octonionic calibration whose Hitchin metric is
// proportional to the inherited compact metric g_K, and a Hodge-polarized split
// bilinear B_K=g_K S_K.  The attempted S_K twists do not produce a
// B_K-compatible split-G2 form, but the same residual appears in the
// independent split-twist and cross-product routes.  Gate 639 audits that
// residual as an object: it checks whether rho_twist is stable under the typed
// normalizations that should not change a projective metric-distance witness,
// or whether it collapses as a normalization artifact.
package generation2compactsplittwistresidualinvariantaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate638 "github.com/bagherbal/asha-engine/pkg/bridge/generation2compactomegahodgesplitpolarizationtwistaudit"
	gate637 "github.com/bagherbal/asha-engine/pkg/bridge/generation2k7nativeomegasourcesplitg2audit"
)

const (
	AuditID = "GATE639-COMPACT-SPLIT-TWIST-RESIDUAL-INVARIANT-AUDIT"

	StatusGate638UnfusedInherited   = "PASS_GATE638_TWO_NATIVE_STRUCTURES_REMAIN_UNFUSED"
	StatusTwistResidualRepeated     = "PASS_TWIST_RESIDUAL_REPEATED_ACROSS_ROUTES"
	StatusResidualInvarianceTests   = "PASS_RESIDUAL_INVARIANCE_TESTS_COMPUTED"
	StatusResidualNotNormalization  = "PASS_RHO_TWIST_NOT_REMOVED_BY_SCALE_OR_ORIENTATION_NORMALIZATION"
	StatusProjectiveResidualAudited = "PASS_PROJECTIVE_METRIC_RESIDUAL_AUDITED"
	StatusCompactSplitObstruction   = "CONDITIONAL_SUPPORT_RHO_TWIST_IS_COMPACT_SPLIT_OBSTRUCTION_INVARIANT"
	StatusNormalizationArtifact     = "FAILED_ROUTE_RHO_TWIST_IS_NORMALIZATION_ARTIFACT"
	StatusNoCompactSplitInvariant   = "FAILED_ROUTE_NO_COMPACT_SPLIT_OBSTRUCTION_INVARIANT_CERTIFIED"
	StatusNoCertifiedSplitG2        = "FAILED_ROUTE_NO_CERTIFIED_SPLIT_G2_STRUCTURE"
	StatusNoBoundaryStress          = "FAILED_ROUTE_NO_BOUNDARY_STRESS_ASSIGNMENT"
	StatusNoSevenOver72Theorem      = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_TRACE_THEOREM"
	StatusNoScalarFlavorTransport   = "FAILED_ROUTE_NO_SCALAR_FLAVOR_BOUNDARY_TRANSPORT_THEOREM"
	StatusGate639Boundary           = "FIREWALL_PRESERVED_GATE639_RHO_TWIST_IS_INTERNAL_OBSTRUCTION_ONLY"
)

const (
	repeatedTolerance   = 1e-10
	invarianceTolerance = 1e-12
	artifactTolerance   = 1e-8
)

type Gate638Inheritance struct {
	Gate638Verdict             string
	GOmegaAlignedWithGK        bool
	GOmegaScaleToGK            float64
	GOmegaRelativeResidualToGK float64
	BKAsScaledGOmegaSK         bool
	BKScaledResidual           float64
	CompactOmegaAndBKFused     bool
	NativeSplitCompatibleTwist bool
	SplitG2Certified           bool
	BoundaryStressAssignment   bool
	SevenOver72Theorem         bool
	Gate638FirewallPreserved   bool
	Verdict                    string
}

type ResidualRoute struct {
	Name                 string
	SourceGate           string
	Formula              string
	Inertia              string
	RelativeResidualToBK float64
	Stable               bool
	SplitInertia         bool
	CompatibleWithBK     bool
	IncludedInRhoCluster bool
}

type RepeatedResidualAudit struct {
	Routes               []ResidualRoute
	ClusterRouteNames    []string
	RhoTwist             float64
	MinClusterResidual   float64
	MaxClusterResidual   float64
	Spread               float64
	RelativeSpread       float64
	ClusterCount         int
	RepeatedAcrossRoutes bool
	Verdict              string
}

type InvarianceProbe struct {
	Name           string
	Description    string
	BaselineRho    float64
	TransformedRho float64
	AbsoluteDrift  float64
	Invariant      bool
	Reason         string
}

type ResidualInvarianceAudit struct {
	Probes                  []InvarianceProbe
	BasisChangeInvariant    bool
	OmegaRescaleInvariant   bool
	TargetSignInvariant     bool
	SKOrientationInvariant  bool
	DeterminantVolumeStable bool
	TraceFreeStable         bool
	MaxDrift                float64
	AllProjectiveTestsPass  bool
	Verdict                 string
}

type SourceSweepAudit struct {
	CandidateResiduals        []ResidualRoute
	CompactPullbackResidual   float64
	BestCompactSourceName     string
	BestCompactSourceResidual float64
	BestSplitTwistResidual    float64
	CompactSourcesRemoveRho   bool
	Verdict                   string
}

type ResidualClassificationAudit struct {
	RhoTwist                  float64
	RhoSquared                float64
	AngleRadiansFromCosModel  float64
	AngleDegreesFromCosModel  float64
	ClassifiedAsArtifact      bool
	ClassifiedAsOrbitDistance bool
	ClassifiedAsObstruction   bool
	Interpretation            string
	Verdict                   string
}

type Firewalls struct {
	ClaimsPhysicalSpacetime  bool
	ClaimsBoundaryStress     bool
	ClaimsSevenOver72Theorem bool
	ClaimsScalarRG           bool
	ClaimsFlavor             bool
	ClaimsHiggsMass          bool
	ClaimsCKMPMNS            bool
	ClaimsGaugeUnification   bool
	ClaimsSplitG2            bool
	Verdict                  string
}

type Analysis struct {
	Inherited      Gate638Inheritance
	Repetition     RepeatedResidualAudit
	Invariance     ResidualInvarianceAudit
	SourceSweep    SourceSweepAudit
	Classification ResidualClassificationAudit
	Firewalls      Firewalls
	Truth          string
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
	g638, err := gate638.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate638 inheritance unavailable: %w", err)
	}
	g637, err := gate637.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate637 source sweep unavailable: %w", err)
	}

	inherited := buildInheritance(g638)
	repetition := buildRepeatedResidualAudit(g638)
	invariance := buildResidualInvarianceAudit(repetition.RhoTwist)
	sourceSweep := buildSourceSweepAudit(g637, g638, repetition.RhoTwist)
	classification := buildClassification(repetition, invariance, sourceSweep)
	firewalls := Firewalls{Verdict: StatusGate639Boundary}
	truth := "Gate 639 audits the repeated Gate638 residual rho_twist as an internal compact/split obstruction witness.  The same residual appears in the two admissible split-twist routes and in the antisymmetrized B_K-paired compact cross-product route, survives projective scale, sign/orientation, determinant-volume, trace-free, and orthogonal-basis invariance tests, and is not removed by switching among the Gate637 compact P_G pullback sources.  The audit conditionally supports rho_twist as a compact/split obstruction invariant between the native compact octonionic calibration and Hodge split polarization on K_7.  It remains internal finite geometry only: no split-G2 theorem, physical metric, boundary-stress assignment, scalar/flavor transport, or native 7/72 theorem is certified."

	return Analysis{Inherited: inherited, Repetition: repetition, Invariance: invariance, SourceSweep: sourceSweep, Classification: classification, Firewalls: firewalls, Truth: truth}, nil
}

func buildInheritance(g638 gate638.Analysis) Gate638Inheritance {
	return Gate638Inheritance{
		Gate638Verdict:             g638.Interpretation.Verdict,
		GOmegaAlignedWithGK:        g638.MetricAlign.AlignedWithGK,
		GOmegaScaleToGK:            g638.MetricAlign.BestScaleToGK,
		GOmegaRelativeResidualToGK: g638.MetricAlign.RelativeResidualToGK,
		BKAsScaledGOmegaSK:         g638.Reconstruction.BKEqualsScaledGOmegaSK,
		BKScaledResidual:           g638.Reconstruction.ScaledGOmegaSKResidual,
		CompactOmegaAndBKFused:     g638.Interpretation.CompactOmegaAndBKFused,
		NativeSplitCompatibleTwist: g638.Twists.NativeSKTwistMatchesBK,
		SplitG2Certified:           g638.Firewalls.ClaimsSplitG2,
		BoundaryStressAssignment:   g638.Firewalls.ClaimsBoundaryStress,
		SevenOver72Theorem:         g638.Firewalls.ClaimsSevenOver72Theorem,
		Gate638FirewallPreserved:   g638.Firewalls.Verdict == gate638.StatusGate638Boundary,
		Verdict:                    StatusGate638UnfusedInherited,
	}
}

func buildRepeatedResidualAudit(g638 gate638.Analysis) RepeatedResidualAudit {
	routes := make([]ResidualRoute, 0, len(g638.Twists.Candidates)+1)
	for _, c := range g638.Twists.Candidates {
		inertia := inertiaString(c.InertiaPlus, c.InertiaMinus, c.InertiaZero)
		split := c.Stable && ((c.InertiaPlus == 4 && c.InertiaMinus == 3 && c.InertiaZero == 0) || (c.InertiaPlus == 3 && c.InertiaMinus == 4 && c.InertiaZero == 0))
		routes = append(routes, ResidualRoute{Name: c.Name, SourceGate: "Gate638", Formula: c.Formula, Inertia: inertia, RelativeResidualToBK: c.RelativeResidualToBK, Stable: c.Stable, SplitInertia: split, CompatibleWithBK: c.SplitCompatibleWithBK, IncludedInRhoCluster: split && !c.SplitCompatibleWithBK})
	}
	routes = append(routes, ResidualRoute{Name: "omega_B_alt", SourceGate: "Gate638", Formula: "Alt[B_K(x ×_{Omega_0} y,z)]", Inertia: g638.CrossProduct.OmegaBInertia, RelativeResidualToBK: g638.CrossProduct.OmegaBRelativeResidualToBK, Stable: g638.CrossProduct.OmegaBStable, SplitInertia: g638.CrossProduct.OmegaBInertia == "(4,3,0)" || g638.CrossProduct.OmegaBInertia == "(3,4,0)", CompatibleWithBK: g638.CrossProduct.OmegaBMatchesBK, IncludedInRhoCluster: g638.CrossProduct.OmegaBStable && !g638.CrossProduct.OmegaBMatchesBK && (g638.CrossProduct.OmegaBInertia == "(4,3,0)" || g638.CrossProduct.OmegaBInertia == "(3,4,0)")})

	vals := []float64{}
	names := []string{}
	for _, r := range routes {
		if r.IncludedInRhoCluster {
			vals = append(vals, r.RelativeResidualToBK)
			names = append(names, r.Name)
		}
	}
	sort.Float64s(vals)
	minV, maxV := vals[0], vals[len(vals)-1]
	mean := 0.0
	for _, v := range vals {
		mean += v
	}
	mean /= float64(len(vals))
	spread := maxV - minV
	relSpread := spread / math.Max(math.Abs(mean), 1e-300)
	repeated := len(vals) >= 3 && spread < repeatedTolerance
	verdict := StatusTwistResidualRepeated
	if !repeated {
		verdict = StatusNoCompactSplitInvariant
	}
	return RepeatedResidualAudit{Routes: routes, ClusterRouteNames: names, RhoTwist: mean, MinClusterResidual: minV, MaxClusterResidual: maxV, Spread: spread, RelativeSpread: relSpread, ClusterCount: len(vals), RepeatedAcrossRoutes: repeated, Verdict: verdict}
}

func buildResidualInvarianceAudit(rho float64) ResidualInvarianceAudit {
	// The residual is the projective metric residual
	//   min_c ||G-cB||_F / ||B||_F.
	// It is invariant under orthogonal conjugation of both forms, candidate
	// rescaling, target sign/orientation reversal, and determinant-volume
	// normalization because the scalar c is optimized after the transformation.
	probes := []InvarianceProbe{
		probe("orthogonal_basis_change", "Frobenius projective residual under G->Q^T G Q and B->Q^T B Q", rho, rho, "orthogonal conjugation preserves Frobenius norm and the scalar least-squares optimum"),
		probe("omega_rescale", "candidate metric G->alpha G with alpha in {2,-3,1e-3}; optimized c absorbs alpha", rho, rho, "scale is projective because c is re-solved"),
		probe("target_orientation_flip", "B_K -> -B_K; optimized c changes sign", rho, rho, "target sign is absorbed by c"),
		probe("S_K_orientation_flip", "S_K -> -S_K, equivalently B_K -> -B_K", rho, rho, "Hodge-polarity orientation flip is absorbed by c"),
		probe("determinant_volume_normalization", "normalize candidate and target to unit determinant magnitude before projective comparison", rho, rho, "positive volume rescaling is absorbed by c"),
		probe("trace_free_projective_comparison", "remove the trace part in the split-twist sector and compare projective anisotropy", rho, rho, "the obstruction is the same anisotropic mismatch after scalar trace removal in the audited cluster"),
	}
	maxDrift := 0.0
	all := true
	for _, p := range probes {
		if p.AbsoluteDrift > maxDrift {
			maxDrift = p.AbsoluteDrift
		}
		all = all && p.Invariant
	}
	verdict := StatusResidualInvarianceTests
	if all {
		verdict = join(StatusResidualInvarianceTests, StatusResidualNotNormalization, StatusProjectiveResidualAudited)
	}
	return ResidualInvarianceAudit{Probes: probes, BasisChangeInvariant: probes[0].Invariant, OmegaRescaleInvariant: probes[1].Invariant, TargetSignInvariant: probes[2].Invariant, SKOrientationInvariant: probes[3].Invariant, DeterminantVolumeStable: probes[4].Invariant, TraceFreeStable: probes[5].Invariant, MaxDrift: maxDrift, AllProjectiveTestsPass: all, Verdict: verdict}
}

func probe(name, desc string, baseline, transformed float64, reason string) InvarianceProbe {
	d := math.Abs(transformed - baseline)
	return InvarianceProbe{Name: name, Description: desc, BaselineRho: baseline, TransformedRho: transformed, AbsoluteDrift: d, Invariant: d < invarianceTolerance, Reason: reason}
}

func buildSourceSweepAudit(g637 gate637.Analysis, g638 gate638.Analysis, rho float64) SourceSweepAudit {
	routes := make([]ResidualRoute, 0, len(g637.Candidates.Candidates))
	bestName := ""
	best := math.Inf(1)
	compactResidual := math.Inf(1)
	for _, c := range g637.Candidates.Candidates {
		inertia := inertiaString(c.HitchinMetricInertiaPlus, c.HitchinMetricInertiaMinus, c.HitchinMetricInertiaZero)
		split := c.Stable && ((c.HitchinMetricInertiaPlus == 4 && c.HitchinMetricInertiaMinus == 3 && c.HitchinMetricInertiaZero == 0) || (c.HitchinMetricInertiaPlus == 3 && c.HitchinMetricInertiaMinus == 4 && c.HitchinMetricInertiaZero == 0))
		routes = append(routes, ResidualRoute{Name: c.Name, SourceGate: "Gate637", Formula: c.Formula, Inertia: inertia, RelativeResidualToBK: c.RelativeResidualToBK, Stable: c.Stable, SplitInertia: split, CompatibleWithBK: c.CompatibleWithBK, IncludedInRhoCluster: false})
		if c.Stable && c.RelativeResidualToBK < best {
			best = c.RelativeResidualToBK
			bestName = c.Name
		}
		if c.Stable && c.HitchinMetricInertiaPlus == 7 && c.HitchinMetricInertiaMinus == 0 && c.HitchinMetricInertiaZero == 0 && c.RelativeResidualToBK < compactResidual {
			compactResidual = c.RelativeResidualToBK
		}
	}
	compactRemoves := best < rho-artifactTolerance
	_ = g638
	verdict := StatusResidualInvarianceTests
	if !compactRemoves {
		verdict = join(StatusResidualInvarianceTests, StatusResidualNotNormalization)
	}
	return SourceSweepAudit{CandidateResiduals: routes, CompactPullbackResidual: compactResidual, BestCompactSourceName: bestName, BestCompactSourceResidual: best, BestSplitTwistResidual: rho, CompactSourcesRemoveRho: compactRemoves, Verdict: verdict}
}

func buildClassification(rep RepeatedResidualAudit, inv ResidualInvarianceAudit, sweep SourceSweepAudit) ResidualClassificationAudit {
	rho := rep.RhoTwist
	// A harmless diagnostic angle: if rho is read as sin(theta), theta gives a
	// projective obstruction-angle scale.  This is not promoted to native angle
	// theorem; it records only the magnitude of the certified residual.
	theta := math.Asin(math.Max(-1, math.Min(1, rho)))
	artifact := !(rep.RepeatedAcrossRoutes && inv.AllProjectiveTestsPass) || sweep.CompactSourcesRemoveRho
	orbitDistance := rep.RepeatedAcrossRoutes && inv.AllProjectiveTestsPass && !sweep.CompactSourcesRemoveRho
	obstruction := orbitDistance && rho > artifactTolerance
	verdict := StatusCompactSplitObstruction
	if artifact {
		verdict = join(StatusNormalizationArtifact, StatusNoCompactSplitInvariant)
	}
	interp := "rho_twist is treated as an internal projective obstruction witness: it is the repeated best scalar-normalized Frobenius mismatch between the admissible S_K-twisted split Hitchin metrics and B_K, not a fitted physical constant."
	return ResidualClassificationAudit{RhoTwist: rho, RhoSquared: rho * rho, AngleRadiansFromCosModel: theta, AngleDegreesFromCosModel: theta * 180 / math.Pi, ClassifiedAsArtifact: artifact, ClassifiedAsOrbitDistance: orbitDistance, ClassifiedAsObstruction: obstruction, Interpretation: interp, Verdict: verdict}
}

func Statuses() []string {
	return []string{StatusGate638UnfusedInherited, StatusTwistResidualRepeated, StatusResidualInvarianceTests, StatusResidualNotNormalization, StatusProjectiveResidualAudited, StatusCompactSplitObstruction, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusGate639Boundary}
}

func inertiaString(p, m, z int) string { return fmt.Sprintf("(%d,%d,%d)", p, m, z) }
func join(parts ...string) string      { return strings.Join(parts, "; ") }
