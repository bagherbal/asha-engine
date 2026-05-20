package historytransport

import "math"

type yukawaBoundaryState struct {
	Up      map[string]float64
	Down    map[string]float64
	Leptons map[string]float64
}

func buildScalarAndYukawaTransport(in InputSet, end EndVector, gb GaugeBoundary) (ScalarTransport, yukawaBoundaryState) {
	initial := []float64{
		end.GY, end.G2, end.G3,
		end.YukawaSingularValues.UpQuarks["u"], end.YukawaSingularValues.UpQuarks["c"], end.YukawaSingularValues.UpQuarks["t"],
		end.YukawaSingularValues.DownQuarks["d"], end.YukawaSingularValues.DownQuarks["s"], end.YukawaSingularValues.DownQuarks["b"],
		end.YukawaSingularValues.ChargedLeptons["e"], end.YukawaSingularValues.ChargedLeptons["mu"], end.YukawaSingularValues.ChargedLeptons["tau"],
		end.Lambda,
	}
	beta0 := historyDerivatives(initial)
	state, zeroT := integrateHistory(initial, gb.LogLambda12Mu0, 20000)
	var zeroScale *float64
	status := "lambda_positive_to_lambda12_in_v1_approximation"
	if zeroT != nil {
		s := in.Mu0GeV * math.Exp(*zeroT)
		zeroScale = &s
		status = "lambda_crosses_zero_before_lambda12_in_v1_approximation"
	}
	yb := yukawaBoundaryState{
		Up:      map[string]float64{"u": state[3], "c": state[4], "t": state[5]},
		Down:    map[string]float64{"d": state[6], "s": state[7], "b": state[8]},
		Leptons: map[string]float64{"e": state[9], "mu": state[10], "tau": state[11]},
	}
	return ScalarTransport{LambdaMZ: end.Lambda, YT_MZ: end.YukawaSingularValues.UpQuarks["t"], LambdaLambda12: state[12], YT_Lambda12: state[5], BetaLambdaMZ: beta0[12], ZeroCrossingScaleGeV: zeroScale, VacuumStabilityStatus: status, Approximation: "one-loop SM v1; lambda beta uses top-dominant user formula; Yukawa transport uses diagonal one-loop SM approximation and fixed CKM", Statuses: []string{StatusScalarTransportComputed, StatusThresholdsNotHidden, StatusNoNativeDerivationClaim}}, yb
}

func historyDerivatives(y []float64) []float64 {
	gY, g2, g3 := y[0], y[1], y[2]
	yu, yc, yt := y[3], y[4], y[5]
	yd, ys, yb := y[6], y[7], y[8]
	ye, ymu, ytau := y[9], y[10], y[11]
	lambda := y[12]
	loop := 16.0 * pi * pi
	T := 3.0*(yu*yu+yc*yc+yt*yt) + 3.0*(yd*yd+ys*ys+yb*yb) + ye*ye + ymu*ymu + ytau*ytau
	out := make([]float64, len(y))
	out[0] = (41.0 / 6.0) * gY * gY * gY / loop
	out[1] = (-19.0 / 6.0) * g2 * g2 * g2 / loop
	out[2] = (-7.0) * g3 * g3 * g3 / loop
	gaugeU := (17.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8.0*g3*g3
	gaugeD := (5.0/12.0)*gY*gY + (9.0/4.0)*g2*g2 + 8.0*g3*g3
	gaugeE := (15.0/4.0)*gY*gY + (9.0/4.0)*g2*g2
	out[3] = yu * (1.5*(yu*yu-yd*yd) + T - gaugeU) / loop
	out[4] = yc * (1.5*(yc*yc-ys*ys) + T - gaugeU) / loop
	out[5] = yt * (1.5*(yt*yt-yb*yb) + T - gaugeU) / loop
	out[6] = yd * (1.5*(yd*yd-yu*yu) + T - gaugeD) / loop
	out[7] = ys * (1.5*(ys*ys-yc*yc) + T - gaugeD) / loop
	out[8] = yb * (1.5*(yb*yb-yt*yt) + T - gaugeD) / loop
	out[9] = ye * (1.5*ye*ye + T - gaugeE) / loop
	out[10] = ymu * (1.5*ymu*ymu + T - gaugeE) / loop
	out[11] = ytau * (1.5*ytau*ytau + T - gaugeE) / loop
	out[12] = (24.0*lambda*lambda - 6.0*yt*yt*yt*yt + (3.0/8.0)*(2.0*math.Pow(g2, 4)+math.Pow(g2*g2+gY*gY, 2)) + lambda*(-9.0*g2*g2-3.0*gY*gY+12.0*yt*yt)) / loop
	return out
}

func integrateHistory(initial []float64, tEnd float64, n int) ([]float64, *float64) {
	if n < 1 {
		n = 1
	}
	y := append([]float64(nil), initial...)
	dt := tEnd / float64(n)
	var zero *float64
	for i := 0; i < n; i++ {
		t := float64(i) * dt
		_ = t
		prevLambda := y[12]
		k1 := historyDerivatives(y)
		k2 := historyDerivatives(addScaled(y, k1, dt/2.0))
		k3 := historyDerivatives(addScaled(y, k2, dt/2.0))
		k4 := historyDerivatives(addScaled(y, k3, dt))
		for j := range y {
			y[j] += dt * (k1[j] + 2.0*k2[j] + 2.0*k3[j] + k4[j]) / 6.0
		}
		if zero == nil && prevLambda > 0 && y[12] <= 0 {
			frac := prevLambda / (prevLambda - y[12])
			z := (float64(i) + frac) * dt
			zero = &z
		}
	}
	return y, zero
}

func addScaled(y, k []float64, scale float64) []float64 {
	out := make([]float64, len(y))
	for i := range y {
		out[i] = y[i] + scale*k[i]
	}
	return out
}
