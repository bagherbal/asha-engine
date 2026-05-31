from __future__ import annotations

from math import cos, sin, sinh, cosh, pi
from typing import Dict, List

import numpy as np

from .contracts import COORDINATES, ETA_ASHA


def symplectic_matrix() -> np.ndarray:
    # Coordinate order: x0,x1,x2,x3,p0,p1,p2,p3.
    # Ω = Σ dp_mu ∧ dx^mu => matrix [[0,-I],[I,0]].
    I = np.eye(4)
    Z = np.zeros((4,4))
    return np.block([[Z, -I], [I, Z]])


def projector_x() -> np.ndarray:
    return np.diag([1,1,1,1,0,0,0,0])


def projector_p() -> np.ndarray:
    return np.diag([0,0,0,0,1,1,1,1])


def eta_matrix() -> np.ndarray:
    return np.diag(ETA_ASHA)


def one_flow(t_values: np.ndarray, omega=(1.0, 1.35, 1.8), hyperbolic_scale=0.42) -> np.ndarray:
    """Canonical ASHA-style phase flow for the seed 'One'.

    The time-energy pair is shown as hyperbolic, matching J_eta^2=+1 on (x0,p0).
    The three spatial pairs are elliptic oscillators, matching J_eta^2=-1 on (xi,pi).
    This is a visual law-space flow, not a physical measured trajectory.
    """
    pts = []
    for t in t_values:
        # Hyperbolic pair kept bounded for rendering through tanh-like normalization.
        u = hyperbolic_scale * (t - t_values[len(t_values)//2])
        x0 = sinh(u) / cosh(abs(u) + 0.7)
        p0 = cosh(u) / cosh(abs(u) + 0.7)
        x = [x0]
        p = [p0]
        for i, w in enumerate(omega, start=1):
            amp = 1.0 / (i ** 0.35)
            phase = i * pi / 7.0
            x.append(amp * cos(w*t + phase))
            p.append(-amp * w * sin(w*t + phase))
        pts.append(x + p)
    return np.array(pts)


def projected_paths(sample_count: int = 420) -> Dict:
    ts = np.linspace(0, 8*pi, sample_count)
    flow = one_flow(ts)
    return {
        "t_start": float(ts[0]),
        "t_end": float(ts[-1]),
        "sample_count": int(sample_count),
        "coordinates": COORDINATES,
        "eta": eta_matrix().tolist(),
        "omega": symplectic_matrix().tolist(),
        "Pi_X": projector_x().tolist(),
        "Pi_P": projector_p().tolist(),
        "trajectory_samples_sparse": flow[::max(1, sample_count//42)].round(8).tolist(),
        "phase_truth": {
            "time_energy_pair": "hyperbolic visual orbit on (x0,p0); J_eta^2=+1",
            "spatial_pairs": "elliptic phase-plane visual orbits on (xi,pi); J_eta^2=-1",
        },
    }
