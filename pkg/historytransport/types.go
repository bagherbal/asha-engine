// Package historytransport implements ASHA-HISTORY-TRANSPORT-END-CALCULATION-V1.
//
// The package is intentionally an observed-data airlock.  It computes a
// reproducible Standard-Model history-transport vector from version-pinned
// published endpoint data and ASHA-certified boundary normalizations.  It does
// not promote observed numbers, RG thresholds, Yukawa textures, CKM/PMNS data,
// cosmology, or scalar stability into native ASHA derivations.
package historytransport

const (
	TaskName = "ASHA-HISTORY-TRANSPORT-END-CALCULATION-V1"
	Mu0Name  = "M_Z"

	StatusEndVectorBuilt                    = "PASS_END_VECTOR_MZ_REPRODUCIBLY_BUILT_FROM_PINNED_OBSERVED_INPUTS"
	StatusBoundaryScaleSolved               = "PASS_G1_G2_BOUNDARY_SCALE_LAMBDA12_SOLVED_ONE_LOOP_SM"
	StatusBoundaryWeakAngleCertified        = "PASS_ASHA_BOUNDARY_WEAK_ANGLE_THREE_EIGHTHS_CERTIFIED"
	StatusWeakAngleTransportResidualVisible = "PASS_WEAK_ANGLE_TRANSPORT_RESIDUAL_VISIBLE"
	StatusStrongMismatchVisible             = "PASS_STRONG_COUPLING_MISMATCH_DELTA3_VISIBLE"
	StatusScalarTransportComputed           = "CONDITIONAL_SUPPORT_SCALAR_TRANSPORT_COMPUTED_ONE_LOOP_TOP_DOMINANT_APPROXIMATION"
	StatusFlavorTransportComputed           = "CONDITIONAL_SUPPORT_FLAVOR_TRANSPORT_COMPUTED_DIAGONAL_ONE_LOOP_APPROXIMATION"
	StatusCosmologyEndpointQuarantined      = "CONDITIONAL_SUPPORT_PLANCK_LCDM_ENDPOINT_INCLUDED_AS_COSMOLOGY_SEAL"
	StatusNoNativeDerivationClaim           = "FIREWALL_PRESERVED_NO_OBSERVED_INPUT_IMPORTED_AS_ASHA_NATIVE_DERIVATION"
	StatusNoPhysicalUnificationClaim        = "FIREWALL_PRESERVED_G1_G2_TEST_ONLY_NO_FULL_PHYSICAL_UNIFICATION_CLAIM"
	StatusThresholdsNotHidden               = "FIREWALL_PRESERVED_THRESHOLDS_AND_SCHEMES_EXPLICITLY_LABELED"
)

type SourceRef struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	URL     string `json:"url"`
	Version string `json:"version"`
	Note    string `json:"note"`
}

type MeasuredValue struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	Uncertainty float64 `json:"uncertainty,omitempty"`
	Unit        string  `json:"unit"`
	Scale       string  `json:"scale,omitempty"`
	Scheme      string  `json:"scheme,omitempty"`
	SourceID    string  `json:"source_id"`
	Role        string  `json:"role"`
	BridgeOnly  bool    `json:"bridge_only"`
}

type FermionMassInput struct {
	Name            string  `json:"name"`
	MassGeV         float64 `json:"mass_gev"`
	MassUncertainty float64 `json:"mass_uncertainty_gev,omitempty"`
	InputScaleGeV   float64 `json:"input_scale_gev"`
	TargetScaleGeV  float64 `json:"target_scale_gev"`
	MassAtMZGeV     float64 `json:"mass_at_mz_gev"`
	Scheme          string  `json:"scheme"`
	Transport       string  `json:"transport"`
	SourceID        string  `json:"source_id"`
	BridgeOnly      bool    `json:"bridge_only"`
}

type CKMInput struct {
	S12      float64 `json:"s12"`
	S13      float64 `json:"s13"`
	S23      float64 `json:"s23"`
	Delta    float64 `json:"delta_rad"`
	SourceID string  `json:"source_id"`
}

type CosmologyEndpoint struct {
	OmegaCH2   float64 `json:"omega_c_h2"`
	OmegaBH2   float64 `json:"omega_b_h2"`
	NS         float64 `json:"n_s"`
	Tau        float64 `json:"tau"`
	SourceID   string  `json:"source_id"`
	BridgeOnly bool    `json:"bridge_only"`
}

type ASHABoundaryLaw struct {
	KY                        float64 `json:"k_y"`
	Sin2ThetaBoundary         float64 `json:"sin2_theta_boundary"`
	CanonicalBoundaryRelation string  `json:"canonical_boundary_relation"`
	FiniteAlgebra             string  `json:"finite_algebra"`
	ScalarCarrier             string  `json:"scalar_carrier"`
	BridgeOnly                bool    `json:"bridge_only"`
}

type InputSet struct {
	TaskName     string                   `json:"task_name"`
	Mu0Name      string                   `json:"mu0_name"`
	Mu0GeV       float64                  `json:"mu0_gev"`
	ASHABoundary ASHABoundaryLaw          `json:"asha_boundary_law"`
	Measured     map[string]MeasuredValue `json:"measured"`
	Fermions     []FermionMassInput       `json:"fermions"`
	CKM          CKMInput                 `json:"ckm"`
	Cosmology    CosmologyEndpoint        `json:"cosmology"`
	Sources      []SourceRef              `json:"sources"`
	Warnings     []string                 `json:"warnings"`
}

type ComplexValue struct {
	Re  float64 `json:"re"`
	Im  float64 `json:"im"`
	Abs float64 `json:"abs"`
}

type YukawaSingularValues struct {
	UpQuarks       map[string]float64 `json:"up_quarks"`
	DownQuarks     map[string]float64 `json:"down_quarks"`
	ChargedLeptons map[string]float64 `json:"charged_leptons"`
}

type EndVector struct {
	Mu0GeV                 float64              `json:"mu0_gev"`
	VGeV                   float64              `json:"v_gev"`
	GY                     float64              `json:"g_y"`
	G1                     float64              `json:"g1_canonical"`
	G2                     float64              `json:"g2"`
	G3                     float64              `json:"g3"`
	AlphaS                 float64              `json:"alpha_s"`
	Sin2Theta              float64              `json:"sin2_theta_end_on_shell"`
	Lambda                 float64              `json:"lambda"`
	YukawaSingularValues   YukawaSingularValues `json:"yukawa_singular_values"`
	CKM                    [][]ComplexValue     `json:"ckm"`
	CKMMagnitudes          [][]float64          `json:"ckm_magnitudes"`
	CKMConvention          string               `json:"ckm_convention"`
	QuarkMassTransportNote string               `json:"quark_mass_transport_note"`
	Statuses               []string             `json:"statuses"`
}

type GaugeBoundary struct {
	Mu0GeV         float64  `json:"mu0_gev"`
	Lambda12GeV    float64  `json:"lambda_12_gev"`
	LogLambda12Mu0 float64  `json:"log_lambda12_over_mu0"`
	GStar          float64  `json:"g_star"`
	G1Lambda       float64  `json:"g1_lambda"`
	G2Lambda       float64  `json:"g2_lambda"`
	G3Lambda       float64  `json:"g3_lambda"`
	Delta3         float64  `json:"delta_3"`
	R3             float64  `json:"r_3"`
	Interpretation string   `json:"interpretation"`
	RGEConvention  string   `json:"rge_convention"`
	Statuses       []string `json:"statuses"`
}

type WeakAngleTransport struct {
	Sin2ThetaBoundary float64  `json:"sin2_theta_boundary"`
	Sin2ThetaEnd      float64  `json:"sin2_theta_end"`
	DeltaSin2         float64  `json:"delta_sin2"`
	TransportRequired bool     `json:"transport_required"`
	Statuses          []string `json:"statuses"`
}

type ScalarTransport struct {
	LambdaMZ              float64  `json:"lambda_mz"`
	YT_MZ                 float64  `json:"y_t_mz"`
	LambdaLambda12        float64  `json:"lambda_lambda12"`
	YT_Lambda12           float64  `json:"y_t_lambda12"`
	BetaLambdaMZ          float64  `json:"beta_lambda_mz"`
	ZeroCrossingScaleGeV  *float64 `json:"zero_crossing_scale_gev,omitempty"`
	VacuumStabilityStatus string   `json:"vacuum_stability_status"`
	Approximation         string   `json:"approximation"`
	Statuses              []string `json:"statuses"`
}

type YukawaInvariants struct {
	SpecYdagY  []float64 `json:"spec_y_dagger_y"`
	DetYdagY   float64   `json:"det_y_dagger_y"`
	TraceYdagY float64   `json:"trace_y_dagger_y"`
}

type FlavorTransport struct {
	Mu0GeV                       float64                     `json:"mu0_gev"`
	Lambda12GeV                  float64                     `json:"lambda12_gev"`
	YukawaSingularValuesMZ       YukawaSingularValues        `json:"yukawa_singular_values_mz"`
	YukawaSingularValuesLambda12 YukawaSingularValues        `json:"yukawa_singular_values_lambda12"`
	YukawaMatricesMZ             map[string][][]ComplexValue `json:"yukawa_matrices_mz"`
	YukawaMatricesLambda12       map[string][][]ComplexValue `json:"yukawa_matrices_lambda12"`
	InvariantsMZ                 map[string]YukawaInvariants `json:"texture_invariants_mz"`
	InvariantsLambda12           map[string]YukawaInvariants `json:"texture_invariants_lambda12"`
	CKM                          [][]ComplexValue            `json:"ckm_mz"`
	JCKM                         float64                     `json:"j_ckm"`
	KoideQe                      float64                     `json:"koide_qe"`
	Convention                   string                      `json:"convention"`
	ResidualPatterns             []string                    `json:"residual_patterns"`
	Statuses                     []string                    `json:"statuses"`
}

type HistoryResidual struct {
	Gauge struct {
		Delta3 float64 `json:"delta_3"`
		R3     float64 `json:"r_3"`
	} `json:"gauge"`
	WeakAngle struct {
		DeltaSin2 float64 `json:"delta_sin2"`
	} `json:"weak_angle"`
	Scalar struct {
		LambdaLambda12        float64  `json:"lambda_lambda12"`
		ZeroCrossingScaleGeV  *float64 `json:"zero_crossing_scale_gev,omitempty"`
		VacuumStabilityStatus string   `json:"vacuum_stability_status"`
	} `json:"scalar"`
	Flavor struct {
		YukawaHierarchyInvariants map[string]YukawaInvariants `json:"yukawa_hierarchy_invariants"`
		JCKM                      float64                     `json:"j_ckm"`
		KoideQe                   float64                     `json:"koide_qe"`
		ResidualPatterns          []string                    `json:"residual_patterns"`
	} `json:"flavor"`
	Cosmology      CosmologyEndpoint `json:"cosmology_optional"`
	Statuses       []string          `json:"statuses"`
	Interpretation string            `json:"interpretation"`
}

type Bundle struct {
	Inputs             InputSet           `json:"inputs"`
	EndVector          EndVector          `json:"end_vector_mz"`
	GaugeBoundary      GaugeBoundary      `json:"gauge_boundary"`
	WeakAngleTransport WeakAngleTransport `json:"weak_angle_transport"`
	ScalarTransport    ScalarTransport    `json:"scalar_transport"`
	FlavorTransport    FlavorTransport    `json:"flavor_transport"`
	HistoryResidual    HistoryResidual    `json:"history_residual"`
	Statuses           []string           `json:"statuses"`
}
