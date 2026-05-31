from __future__ import annotations

import argparse
import json
from math import atan2, cos, exp, log10, pi, sin, sqrt
from pathlib import Path
from typing import Dict, List, Tuple

import cairosvg
import numpy as np

from .contracts import (
    ACTION_SECTORS,
    COORDINATES,
    E,
    ETA_ASHA,
    GRADE_DIMENSIONS,
    L_HISTORY,
    N_Q,
    PHI,
    S_SPLIT,
    SOURCE_ALPHABET,
    SOURCE_ALPHABET_ROLES,
    THEOREM_LOCKS,
    WOUNDS,
    W_Q,
    contract_dict,
    formula_stack,
)
from .source_digest import write_source_digest
from .svg_primitives import SVG, polar, regular_polygon
from .trajectory import projected_paths, one_flow

Color = str
Point = Tuple[float, float]

BG = "#050508"
PLATINUM = "#F3EDE0"
GOLD = "#EAC66A"
GOLD2 = "#FFDD88"
CYAN = "#76F1FF"
BLUE = "#4E9CFF"
RED = "#FF8068"
MAGENTA = "#CFA3FF"
OBSIDIAN = "#050308"
DIM = "#9CA6B8"
GREEN = "#A8F0B5"
ORANGE = "#FFB36C"

W = 4200
H = 2600
CX = W / 2
CY = H / 2


def svg_defs() -> str:
    return f'''
<radialGradient id="bgGrad" cx="50%" cy="50%" r="70%">
  <stop offset="0%" stop-color="#11131A"/>
  <stop offset="55%" stop-color="#07070B"/>
  <stop offset="100%" stop-color="#010103"/>
</radialGradient>
<radialGradient id="coreGold" cx="50%" cy="45%" r="60%">
  <stop offset="0%" stop-color="#FFF8CF"/>
  <stop offset="38%" stop-color="#F2C86F"/>
  <stop offset="75%" stop-color="#B97833" stop-opacity="0.58"/>
  <stop offset="100%" stop-color="#000000" stop-opacity="0"/>
</radialGradient>
<radialGradient id="cyanLens" cx="45%" cy="50%" r="70%">
  <stop offset="0%" stop-color="#C7FBFF" stop-opacity="0.38"/>
  <stop offset="70%" stop-color="#55E8FF" stop-opacity="0.12"/>
  <stop offset="100%" stop-color="#0A151A" stop-opacity="0"/>
</radialGradient>
<radialGradient id="goldLens" cx="55%" cy="50%" r="70%">
  <stop offset="0%" stop-color="#FFE6AC" stop-opacity="0.36"/>
  <stop offset="72%" stop-color="#EAB45E" stop-opacity="0.10"/>
  <stop offset="100%" stop-color="#110A00" stop-opacity="0"/>
</radialGradient>
<linearGradient id="mpDrop" x1="0%" y1="0%" x2="0%" y2="100%">
  <stop offset="0%" stop-color="#FFF8D0"/>
  <stop offset="35%" stop-color="#F2C86F"/>
  <stop offset="70%" stop-color="#8AF6FF"/>
  <stop offset="100%" stop-color="#343B54"/>
</linearGradient>
<filter id="glow" x="-80%" y="-80%" width="260%" height="260%">
  <feGaussianBlur stdDeviation="7" result="blur"/>
  <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
</filter>
<filter id="softGlow" x="-80%" y="-80%" width="260%" height="260%">
  <feGaussianBlur stdDeviation="16" result="blur"/>
  <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
</filter>
<filter id="hardShadow" x="-40%" y="-40%" width="180%" height="180%">
  <feDropShadow dx="0" dy="18" stdDeviation="18" flood-color="#000000" flood-opacity="0.75"/>
</filter>
<filter id="obsidianGlow" x="-70%" y="-70%" width="240%" height="240%">
  <feGaussianBlur stdDeviation="5" result="blur"/>
  <feColorMatrix in="blur" type="matrix" values="0.6 0 0 0 0.15  0 0.2 0 0 0.06  0 0 0.25 0 0.22  0 0 0 1 0"/>
  <feMerge><feMergeNode/><feMergeNode in="SourceGraphic"/></feMerge>
</filter>
<marker id="arrowGold" markerWidth="12" markerHeight="12" refX="10" refY="6" orient="auto" markerUnits="strokeWidth">
  <path d="M2,2 L10,6 L2,10 Z" fill="{GOLD}" opacity="0.9"/>
</marker>
<marker id="arrowCyan" markerWidth="12" markerHeight="12" refX="10" refY="6" orient="auto" markerUnits="strokeWidth">
  <path d="M2,2 L10,6 L2,10 Z" fill="{CYAN}" opacity="0.9"/>
</marker>
<marker id="arrowRed" markerWidth="12" markerHeight="12" refX="10" refY="6" orient="auto" markerUnits="strokeWidth">
  <path d="M2,2 L10,6 L2,10 Z" fill="{RED}" opacity="0.9"/>
</marker>
<style>
  .label {{ font-family: Inter, DejaVu Sans, Arial, sans-serif; letter-spacing: 0.06em; }}
  .math {{ font-family: 'EB Garamond', Georgia, serif; }}
  .micro {{ font-family: Inter, DejaVu Sans, Arial, sans-serif; letter-spacing: 0.11em; }}
</style>
'''


def add_background(svg: SVG):
    svg.rect(0, 0, W, H, fill="url(#bgGrad)")
    # Subtle obsidian firewall ellipse and cosmic dust. Deterministic golden-angle stars.
    svg.ellipse(CX, CY, 1940, 1040, fill="none", stroke="#121017", width=62, opacity=0.92, filt="hardShadow")
    svg.ellipse(CX, CY, 1850, 970, fill="none", stroke="#2D2026", width=2, opacity=0.6)
    svg.ellipse(CX, CY, 1700, 890, fill="none", stroke="#141D25", width=2, opacity=0.8, dash="6 22")
    # outer chaos wisps
    for i in range(170):
        a = i * 2.399963229728653
        r = 1130 + (i * 37 % 940)
        x = CX + r*cos(a)*1.05
        y = CY + r*sin(a)*0.61
        if 0 < x < W and 0 < y < H:
            col = CYAN if i % 3 == 0 else (RED if i % 5 == 0 else GOLD)
            svg.circle(x, y, 1.2 + (i % 7)*0.28, fill=col, opacity=0.10 + (i % 5)*0.025)
    # exactly 13 external unresolved nodes
    for i in range(13):
        a = -pi/2 + i * 2*pi/13 + 0.11
        x = CX + 1990*cos(a)
        y = CY + 1110*sin(a)
        svg.circle(x, y, 24, fill=GOLD2, opacity=0.09, filt="softGlow")
        svg.circle(x, y, 8.5, fill=GOLD2, opacity=0.52, filt="glow")
        svg.ellipse(x, y, 42, 42, fill="none", stroke=GOLD2, width=1.2, opacity=0.18)


def add_title(svg: SVG):
    svg.text(CX, 92, "ASHA PROJECTOR UNIVERSE - SOURCE-TYPED ONTOLOGY", size=34, fill="#F4E9CC", weight="600", spacing="0.18em")
    svg.text(CX, 136, "from scalar identity to phase-space measurement, contact vacuum, finite spectral law-space, action sectors, and firewalled observation", size=18, fill="#AEB7C7", weight="400", spacing="0.06em")


def add_root_identity(svg: SVG):
    x, y = 230, CY
    svg.circle(x, y, 70, fill="url(#coreGold)", opacity=0.92, filt="softGlow")
    svg.circle(x, y, 10, fill="#FFFAD6", opacity=1.0, filt="glow")
    svg.text(x, y+130, "1", size=58, fill="#FFF6CE", family="Georgia, serif", weight="700")
    svg.text(x, y+174, "scalar identity", size=15, fill=DIM, spacing="0.12em")
    svg.path(f"M {x+70:.1f} {y:.1f} C 390 {y-170:.1f}, 500 {y-170:.1f}, 635 {y-110:.1f}", stroke=GOLD, width=3.0, opacity=0.74, marker="arrowGold", filt="glow")
    svg.path(f"M {x+70:.1f} {y:.1f} C 390 {y+170:.1f}, 500 {y+170:.1f}, 635 {y+110:.1f}", stroke=CYAN, width=2.4, opacity=0.6, marker="arrowCyan", filt="glow")


def add_phase_octave(svg: SVG, geometry: Dict):
    # V8 panel with X4/P4 split and correct time-energy/spatial phase pair behavior.
    px, py = 740, CY
    svg.rect(px-270, py-510, 520, 1020, fill="#070910", stroke="#273444", width=1.2, rx=28, opacity=0.82, filt="hardShadow")
    svg.text(px, py-455, "V8 = X4 + P4", size=32, fill=PLATINUM, family="Georgia, serif", weight="600")
    svg.text(px, py-416, "measurement octave, not 8 spatial dimensions", size=14, fill=DIM, spacing="0.08em")
    x_col = px - 115
    p_col = px + 115
    ys = [py-270, py-105, py+60, py+225]
    names_x = ["x0", "x1", "x2", "x3"]
    names_p = ["p0", "p1", "p2", "p3"]
    for i, y in enumerate(ys):
        time = i == 0
        col = RED if time else CYAN
        pcol = ORANGE if time else BLUE
        svg.circle(x_col, y, 28, fill=col, opacity=0.34, filt="glow")
        svg.circle(x_col, y, 7.5, fill=col, opacity=0.96)
        svg.circle(p_col, y, 28, fill=pcol, opacity=0.30, filt="glow")
        svg.circle(p_col, y, 7.5, fill=pcol, opacity=0.96)
        svg.text(x_col-46, y+6, names_x[i], size=22, fill=col, anchor="end", family="Georgia, serif")
        svg.text(p_col+46, y+6, names_p[i], size=22, fill=pcol, anchor="start", family="Georgia, serif")
        if time:
            # hyperbolic saddle pair, not a circular atom.
            svg.path(f"M {x_col+18} {y-48} C {px-35} {y-12}, {px+35} {y+12}, {p_col-18} {y+48}", stroke=RED, width=2.2, opacity=0.82, marker="arrowRed")
            svg.path(f"M {x_col+18} {y+48} C {px-35} {y+12}, {px+35} {y-12}, {p_col-18} {y-48}", stroke=ORANGE, width=2.0, opacity=0.72, marker="arrowGold")
            svg.text(px, y-74, "hyperbolic x0–p0", size=13, fill="#FFB7A9", spacing="0.08em")
        else:
            svg.ellipse(px, y, 126, 38, fill="none", stroke=col, width=1.8, opacity=0.38, rotate=0)
            svg.ellipse(px, y, 126, 38, fill="none", stroke=pcol, width=1.4, opacity=0.30, rotate=60)
            svg.ellipse(px, y, 126, 38, fill="none", stroke=PLATINUM, width=1.0, opacity=0.16, rotate=-60)
            svg.text(px, y-56, f"elliptic Π{i}", size=13, fill="#BDEEFF", spacing="0.08em")
        # pair axis
        svg.line(x_col+28, y, p_col-28, y, stroke="#EADFB6", width=0.8, opacity=0.20)
    # eta and omega signs side indicators.
    svg.text(px, py+392, "η = diag(+ − − − − − − −)", size=18, fill="#EBDDB9", family="Georgia, serif")
    svg.text(px, py+424, "Ω = Σ dpμ ∧ dxμ", size=18, fill="#B9F7FF", family="Georgia, serif")
    # Flow projection curve inside panel
    pts = one_flow(np.linspace(0, 8*pi, 180))
    # central phase trail projected using selected components x1,p1,x2 onto panel coords
    qpts = []
    for row in pts:
        xx = px + 72*row[1] + 38*row[2]
        yy = py + 62*row[5] - 26*row[6]
        qpts.append((xx, yy))
    svg.polyline(qpts, stroke="#FFF5C7", width=2.0, opacity=0.38, marker="arrowGold", filt="glow")


def add_contact_seven(svg: SVG):
    cx, cy = 1280, CY
    svg.rect(cx-315, cy-510, 610, 1020, fill="#070B0F", stroke="#273943", width=1.2, rx=34, opacity=0.75, filt="hardShadow")
    svg.text(cx, cy-452, "1 + 7  ->  1 + (1 + 6)", size=31, fill=PLATINUM, family="Georgia, serif", weight="600")
    svg.text(cx, cy-415, "x0 reference; contact seven = p0 plus three phase planes", size=14, fill=DIM, spacing="0.06em")
    # extracted x0 ray
    svg.path(f"M {cx-220} {cy-280} C {cx-170} {cy-380}, {cx-75} {cy-410}, {cx} {cy-360}", stroke=RED, width=2.2, opacity=0.60, marker="arrowRed", dash="8 8")
    svg.circle(cx-220, cy-280, 10, fill=RED, opacity=0.85, filt="glow")
    svg.text(cx-228, cy-304, "x0", size=20, fill=RED, family="Georgia, serif", anchor="end")
    svg.text(cx+68, cy-340, "observer-time reference", size=13, fill="#F6B0A4", spacing="0.08em", anchor="start")
    # contact hexagram
    R = 210
    verts = regular_polygon(cx, cy+30, R, 6, start=-pi/2)
    # Star of David as two triangles; geometric not decorative.
    tri1 = [verts[i] for i in [0,2,4]]
    tri2 = [verts[i] for i in [1,3,5]]
    svg.polygon(tri1, fill="#1C1010", stroke=RED, width=1.4, opacity=0.42, filt="glow")
    svg.polygon(tri2, fill="#081A1F", stroke=CYAN, width=1.4, opacity=0.48, filt="glow")
    svg.circle(cx, cy+30, 76, fill="url(#coreGold)", opacity=0.62, filt="softGlow")
    svg.circle(cx, cy+30, 14, fill=GOLD2, opacity=1, filt="glow")
    svg.text(cx, cy+36, "p0", size=18, fill=GOLD2, family="Georgia, serif")
    labels = ["x1", "p1", "x2", "p2", "x3", "p3"]
    colors = [CYAN, BLUE, CYAN, BLUE, CYAN, BLUE]
    for i, ((x,y), lab, col) in enumerate(zip(verts, labels, colors)):
        svg.circle(x, y, 34, fill=col, opacity=0.24, filt="glow")
        svg.circle(x, y, 8, fill=col, opacity=0.95)
        svg.text(x, y+7, lab, size=20, fill=col, family="Georgia, serif")
        svg.line(cx, cy+30, x, y, stroke=col, width=1.2, opacity=0.22)
    # rotating phase arrows around hexagram
    for k, r in enumerate([248, 286, 324]):
        col = [CYAN, GOLD, MAGENTA][k]
        svg.ellipse(cx, cy+30, r, r*0.58, fill="none", stroke=col, width=1.2, opacity=0.20, rotate=20+k*47, dash="5 13")
    # N_Q depth triple
    for i, (n, wq) in enumerate(zip(N_Q, W_Q)):
        yy = cy + 345 + i*40
        length = 230 * (1 - i*0.20)
        col = [GOLD2, CYAN, MAGENTA][i]
        svg.line(cx-length/2, yy, cx+length/2, yy, stroke=col, width=4.0, opacity=0.35, marker="arrowGold" if i==0 else "arrowCyan")
        svg.circle(cx-length/2, yy, 7, fill=col, opacity=0.7)
        svg.circle(cx+length/2, yy, 7, fill=col, opacity=0.7)
        svg.text(cx, yy-13, f"N={n:.3g}  W={wq:.3e}", size=13, fill=col, family="Georgia, serif")


def add_projector_chamber(svg: SVG):
    cx, cy = 1910, CY
    svg.rect(cx-400, cy-520, 780, 1040, fill="#06080D", stroke="#26323E", width=1.2, rx=38, opacity=0.78, filt="hardShadow")
    svg.text(cx, cy-457, "Boolean/G₂ projector chamber", size=28, fill=PLATINUM, family="Georgia, serif", weight="600")
    svg.text(cx, cy-421, "PB rank 56  AND  PG rank 14  ->  K7", size=18, fill="#C8D2DD", family="Georgia, serif")
    # Lenses: rank area logarithmic-ish visual sizes
    svg.ellipse(cx-110, cy+5, 285, 405, fill="url(#cyanLens)", stroke=CYAN, width=2.2, opacity=0.72, rotate=-13, filt="softGlow")
    svg.ellipse(cx+115, cy+5, 185, 330, fill="url(#goldLens)", stroke=GOLD, width=2.0, opacity=0.72, rotate=13, filt="softGlow")
    # Λ4 chamber background: 70 nodes in ring
    for i in range(70):
        a = -pi/2 + i * 2*pi/70
        rr = 362 + 16*sin(7*a)
        x = cx + rr*cos(a)*0.82
        y = cy + rr*sin(a)*0.92
        svg.circle(x, y, 2.8, fill=PLATINUM, opacity=0.18)
    # Boolean 56 flow strands
    for i in range(56):
        a = -pi/2 + i * 2*pi/56
        x1 = cx-265 + 250*cos(a)*0.73
        y1 = cy + 355*sin(a)*0.88
        x2 = cx-25 + 92*cos(a+0.2)
        y2 = cy + 105*sin(a+0.2)
        if i % 4 == 0:
            svg.path(f"M {x1:.1f} {y1:.1f} C {cx-170:.1f} {cy-160:.1f}, {cx-80:.1f} {cy+100:.1f}, {x2:.1f} {y2:.1f}", stroke=CYAN, width=0.65, opacity=0.13)
    # G2 14 rays
    for i in range(14):
        a = -pi/2 + i * 2*pi/14 + 0.13
        x1 = cx+215 + 170*cos(a)*0.7
        y1 = cy + 310*sin(a)*0.88
        x2 = cx+35 + 70*cos(a)
        y2 = cy + 95*sin(a)
        svg.path(f"M {x1:.1f} {y1:.1f} C {cx+200:.1f} {cy-100:.1f}, {cx+80:.1f} {cy+60:.1f}, {x2:.1f} {y2:.1f}", stroke=GOLD, width=0.9, opacity=0.21)
    # K7 payload: 7 nodes in seed flower
    svg.circle(cx, cy, 132, fill="url(#coreGold)", opacity=0.90, filt="softGlow")
    kverts = regular_polygon(cx, cy, 82, 7, start=-pi/2)
    for i, (x,y) in enumerate(kverts):
        svg.line(cx, cy, x, y, stroke=GOLD2, width=1, opacity=0.45)
        svg.circle(x, y, 11, fill=GOLD2, opacity=0.95, filt="glow")
    svg.circle(cx, cy, 17, fill="#FFFAD0", opacity=1, filt="glow")
    svg.text(cx, cy+170, "K7 zero-energy contact vacuum", size=18, fill=GOLD2, family="Georgia, serif")
    # q4 quarantine orbit subtle
    svg.ellipse(cx, cy, 158, 158, fill="none", stroke=MAGENTA, width=1.2, opacity=0.26, dash="4 10")
    svg.text(cx, cy+216, "q4 stays contact-internal", size=13, fill=MAGENTA, spacing="0.08em")


def add_finite_engine(svg: SVG):
    cx, cy = 2620, CY
    svg.rect(cx-430, cy-520, 850, 1040, fill="#06080D", stroke="#293545", width=1.2, rx=38, opacity=0.78, filt="hardShadow")
    svg.text(cx, cy-457, "finite spectral engine", size=29, fill=PLATINUM, family="Georgia, serif", weight="600")
    svg.text(cx, cy-421, "AF = C + H + M3(C), inner fluctuations, 12 gauge directions", size=15, fill=DIM, spacing="0.04em")
    # Islands C, H, M3
    c = (cx-245, cy-115)
    h0 = (cx, cy-115)
    m = (cx+260, cy-115)
    svg.circle(c[0], c[1], 40, fill=GOLD, opacity=0.32, filt="glow")
    svg.circle(c[0], c[1], 10, fill=GOLD2, opacity=0.95)
    svg.text(c[0], c[1]+86, "C", size=30, fill=GOLD2, family="Georgia, serif")
    # H 2x2
    for ix in range(2):
        for iy in range(2):
            x = h0[0] + (ix-0.5)*48
            y = h0[1] + (iy-0.5)*48
            svg.rect(x-14, y-14, 28, 28, fill=CYAN, stroke="#D8FFFF", width=1, rx=5, opacity=0.50, filt="glow")
    svg.text(h0[0], h0[1]+86, "H", size=30, fill=CYAN, family="Georgia, serif")
    # M3 3x3
    for ix in range(3):
        for iy in range(3):
            x = m[0] + (ix-1)*36
            y = m[1] + (iy-1)*36
            svg.rect(x-11, y-11, 22, 22, fill=GOLD if (ix+iy)%2 else BLUE, stroke="#FDF4D8", width=0.8, rx=4, opacity=0.46, filt="glow")
    svg.text(m[0], m[1]+86, "M3(C)", size=28, fill=GOLD2, family="Georgia, serif")
    # Morita bridge waves between islands
    for j, (a,b,col) in enumerate([(c,h0,CYAN),(h0,m,GOLD),(c,m,MAGENTA)]):
        x1,y1 = a; x2,y2 = b
        d = f"M {x1+40:.1f} {y1:.1f} C {(x1+x2)/2:.1f} {y1-80-20*j:.1f}, {(x1+x2)/2:.1f} {y2+80+20*j:.1f}, {x2-40:.1f} {y2:.1f}"
        svg.path(d, stroke=col, width=2.0, opacity=0.42, marker="arrowCyan" if col==CYAN else "arrowGold", filt="glow")
    # 16 matter states charge lattice
    lx, ly = cx-250, cy+235
    svg.text(lx, ly-120, "16 Fock states", size=18, fill="#D8DCE7", spacing="0.08em")
    for n in range(16):
        bits = [(n>>k)&1 for k in range(4)]
        N0,N1,N2,N3 = bits
        bl = -N0 + (N1+N2+N3)/3
        # place by B-L and occupation depth
        x = lx + bl*100
        y = ly + (sum(bits)-2)*42
        col = RED if N0 else CYAN
        if n == 0:
            col = GOLD2
        svg.circle(x, y, 8.5, fill=col, opacity=0.90, filt="glow")
        if n % 3 == 0:
            svg.line(x, y, lx, ly, stroke=col, width=0.6, opacity=0.18)
    svg.line(lx-185, ly, lx+185, ly, stroke="#D4E8FF", width=1, opacity=0.20)
    svg.line(lx, ly-130, lx, ly+130, stroke="#D4E8FF", width=1, opacity=0.20)
    # 12 gauge rays
    gcx, gcy = cx+230, cy+245
    svg.text(gcx, gcy-142, "12 gauge rays + Higgs edge-10", size=18, fill="#D8DCE7", spacing="0.04em")
    for i in range(12):
        a = i * 2*pi/12 - pi/2
        x2 = gcx + 130*cos(a)
        y2 = gcy + 92*sin(a)
        col = [GOLD, CYAN, BLUE][i % 3]
        svg.line(gcx, gcy, x2, y2, stroke=col, width=1.5, opacity=0.43, marker="arrowGold" if col==GOLD else "arrowCyan")
        svg.circle(x2, y2, 6, fill=col, opacity=0.76, filt="glow")
    # Edge-10 decagon for Higgs one-form
    dec = regular_polygon(gcx, gcy, 58, 10, start=-pi/2)
    svg.polygon(dec, fill="#160F08", stroke=GOLD2, width=1.7, opacity=0.58, filt="glow")
    for x,y in dec:
        svg.circle(x,y,4.8,fill=GOLD2,opacity=0.85)


def add_action_and_observation(svg: SVG):
    cx, cy = 3440, CY
    svg.rect(cx-470, cy-520, 910, 1040, fill="#06080D", stroke="#273545", width=1.2, rx=38, opacity=0.78, filt="hardShadow")
    svg.text(cx, cy-457, "scale, action, observation", size=29, fill=PLATINUM, family="Georgia, serif", weight="600")
    svg.text(cx, cy-421, "bridges and fillings are visible but firewalled", size=15, fill=DIM, spacing="0.08em")
    fs = formula_stack()
    # Planck-to-electroweak descent vertical axis
    sx = cx-315
    top = cy-320
    bottom = cy+300
    svg.line(sx, top, sx, bottom, stroke="url(#mpDrop)", width=10, opacity=0.72, marker="arrowGold", filt="glow")
    for i, yy in enumerate(np.linspace(top, bottom, 9)):
        svg.line(sx-36, yy, sx+36, yy, stroke="#E5E7EF", width=0.6, opacity=0.22)
    svg.text(sx, top-34, "MP", size=24, fill=GOLD2, family="Georgia, serif")
    svg.text(sx, bottom+46, "v", size=26, fill=CYAN, family="Georgia, serif")
    svg.text(sx+122, cy-65, "v/MP = exp[-12*pi + sqrt(3)/2 + 2S + 148S^2]", size=16, fill="#F4E7C5", family="Georgia, serif", anchor="start")
    svg.text(sx+122, cy-35, f"≈ {fs.v_over_mp:.3e}", size=16, fill=GOLD2, family="Georgia, serif", anchor="start")
    # Higgs basin / quartic potential
    bx, by = cx+105, cy-108
    svg.text(bx, by-188, "Higgs basin", size=18, fill=GOLD2, spacing="0.08em")
    # draw quartic Mexican hat-ish cross-section as two curves/shaded basin
    pts1=[]; pts2=[]
    for k in range(120):
        u=-1.8+3.6*k/119
        v=(u*u-1)**2
        x=bx+u*95
        y=by+v*48-36
        pts1.append((x,y))
        x2=bx+u*95
        y2=by-v*22+88
        pts2.append((x2,y2))
    svg.polyline(pts1, stroke=GOLD2, width=2.0, opacity=0.64, filt="glow")
    svg.polyline(pts2, stroke=CYAN, width=1.4, opacity=0.24)
    svg.text(bx, by+156, f"lambda={fs.lambda_asha:.6f}; mH=v*sqrt(2lambda)", size=15, fill="#EFD9A0", family="Georgia, serif")
    # Six action sector petals
    acx, acy = cx+145, cy+225
    svg.circle(acx, acy, 76, fill="url(#coreGold)", opacity=0.52, filt="softGlow")
    svg.text(acx, acy+6, "S", size=46, fill="#FFF6CA", family="Georgia, serif", weight="700")
    for i, sec in enumerate(ACTION_SECTORS):
        a = -pi/2 + i*2*pi/6
        x = acx + 185*cos(a)
        y = acy + 135*sin(a)
        col = [GOLD2,CYAN,GREEN,PLATINUM,MAGENTA,BLUE][i]
        svg.path(f"M {acx:.1f} {acy:.1f} C {acx+95*cos(a-0.2):.1f} {acy+70*sin(a-0.2):.1f}, {acx+120*cos(a+0.2):.1f} {acy+95*sin(a+0.2):.1f}, {x:.1f} {y:.1f}", stroke=col, width=2.2, opacity=0.42, marker="arrowGold" if i in [0,2] else "arrowCyan")
        svg.circle(x, y, 37, fill=col, opacity=0.20, filt="glow")
        svg.circle(x, y, 7.8, fill=col, opacity=0.93)
        svg.text(x, y+55, sec["symbol"], size=14, fill=col, family="Georgia, serif")
    # Wounds firewall strip
    fx = cx+330
    fy = cy-300
    svg.rect(fx-142, fy-20, 250, 310, fill="#10070A", stroke=RED, width=1.3, rx=14, opacity=0.66, filt="obsidianGlow")
    svg.text(fx-18, fy+12, "five wounds", size=17, fill="#FFB9AE", spacing="0.08em")
    for i, w in enumerate(WOUNDS):
        yy=fy+50+i*45
        svg.circle(fx-120, yy-4, 5, fill=RED, opacity=0.75, filt="glow")
        svg.text(fx-104, yy, w, size=11.5, fill="#E3B5B1", anchor="start")


def add_source_alphabet_and_locks(svg: SVG):
    # top theorem locks arc
    y = 300
    x0 = 510
    gap = 285
    svg.text(CX, 226, "theorem-level locks: actual ontology order, not discovery order", size=21, fill="#DCE3EC", spacing="0.08em")
    for i, lock in enumerate(THEOREM_LOCKS):
        x = x0 + i*gap
        if x > W-300:
            break
        col = GOLD2 if i < 6 else CYAN
        svg.circle(x, y, 18, fill=col, opacity=0.36, filt="glow")
        svg.circle(x, y, 5.5, fill=col, opacity=0.95)
        if i < len(THEOREM_LOCKS)-1 and i < 10:
            svg.line(x+20, y, x+gap-22, y, stroke=col, width=1.2, opacity=0.30, marker="arrowGold" if col==GOLD2 else "arrowCyan")
        # two-line label
        words = lock.split()
        line1 = " ".join(words[:2])
        line2 = " ".join(words[2:4])
        svg.text(x, y+48, line1, size=10.5, fill="#D5D9E0")
        svg.text(x, y+64, line2, size=10.5, fill="#AEB8C5")
    # source alphabet footer sigils
    yy = H-250
    svg.text(CX, yy-78, "locked source alphabet", size=22, fill="#F2E5C2", spacing="0.14em")
    start = CX - 570
    for i, num in enumerate(SOURCE_ALPHABET):
        x = start + i*190
        col = [GOLD2,CYAN,GOLD,PLATINUM,CYAN,GOLD2,MAGENTA][i]
        svg.circle(x, yy, 52, fill=col, opacity=0.16, filt="glow")
        svg.circle(x, yy, 30, fill="#090B0F", stroke=col, width=1.5, opacity=0.98)
        svg.text(x, yy+10, str(num), size=30, fill=col, family="Georgia, serif", weight="700")
        role = SOURCE_ALPHABET_ROLES[num].split(" /")[0]
        svg.text(x, yy+78, role[:28], size=11.2, fill="#C3C8D0")
    # constants pi phi e small orbit panel
    px, py = 350, H-260
    svg.rect(px-160, py-78, 320, 142, fill="#090A0F", stroke="#262B37", width=1.0, rx=22, opacity=0.72)
    svg.text(px, py-34, "pi · phi · e", size=30, fill=GOLD2, family="Georgia, serif")
    svg.text(px, py, "angle · proportion · growth", size=13, fill=DIM, spacing="0.08em")
    svg.ellipse(px, py+34, 92, 21, fill="none", stroke=GOLD2, width=1.1, opacity=0.35, rotate=0)
    svg.ellipse(px, py+34, 92, 21, fill="none", stroke=CYAN, width=1.1, opacity=0.28, rotate=60)
    svg.ellipse(px, py+34, 92, 21, fill="none", stroke=RED, width=1.1, opacity=0.22, rotate=-60)


def add_interpanel_arrows(svg: SVG):
    # clean ontology flow; not random central lines.
    y = CY
    coords = [(505,y-110),(965,y),(1515,y),(2190,y),(2970,y)]
    colors = [GOLD,CYAN,GOLD2,CYAN,GOLD]
    for (x1,y1),(x2,y2),col in zip(coords, coords[1:], colors):
        svg.path(f"M {x1:.1f} {y1:.1f} C {(x1+x2)/2:.1f} {y1-120:.1f}, {(x1+x2)/2:.1f} {y2+120:.1f}, {x2:.1f} {y2:.1f}", stroke=col, width=3.0, opacity=0.30, marker="arrowGold" if col==GOLD or col==GOLD2 else "arrowCyan", filt="glow")
    # lower bridge from law to observation
    svg.path(f"M 2050 {CY+515} C 2500 {CY+720}, 2920 {CY+700}, 3360 {CY+565}", stroke=RED, width=2.2, opacity=0.25, dash="10 12", marker="arrowRed")
    svg.text(2700, CY+705, "bridge/filling/firewall boundary — no native overwrite", size=17, fill="#D99891", spacing="0.08em")


def add_formula_sidebar(svg: SVG):
    # left-lower formula ledger panel from gate1349: readable but separated.
    x, y = 145, 360
    w, h = 270, 540
    svg.rect(x, y, w, h, fill="#07090E", stroke="#2C3440", width=1.0, rx=22, opacity=0.76)
    svg.text(x+w/2, y+42, "source-typed stack", size=19, fill=PLATINUM, spacing="0.08em")
    lines = [
        "Theorem != Bridge",
        "Bridge != Filling",
        "Filling != Fit",
        "Fit != Wound",
        "",
        "L = 1/(8π)",
        f"S = {S_SPLIT:.16f}",
        "",
        "V8 = X4 + P4",
        "K7 = Im(PB) AND Im(PG)",
        "AF = C + H + M3(C)",
        "Adepth = AF x Qcontact^3",
        "WQ = exp(-4πNQ)",
    ]
    yy = y+82
    for ln in lines:
        if ln == "":
            yy += 12
            continue
        col = GOLD2 if "≠" in ln else (CYAN if "V8" in ln or "K7" in ln else "#D5D8DD")
        svg.text(x+20, yy, ln, size=14, fill=col, anchor="start", family="Georgia, serif")
        yy += 31


def build_geometry_json(source_digest: Dict) -> Dict:
    return {
        "canvas": {"width": W, "height": H, "aspect": W/H},
        "design_intent": "Grand scientific/art-direction schematic of ASHA ontology. Built from formulas and source-type boundaries, not an AI image model.",
        "layout_truth": {
            "not_random_lines": "ontology flow is left-to-right; phase pair lines remain local to the V8 panel; bridges are separated from theorem roots",
            "phase_correction": "x0-p0 is hyperbolic; spatial (xi,pi) pairs are elliptic phase planes; contact seven is p0 plus six spatial phase coordinates",
            "middle_chamber": "K7 is shown as projector intersection payload, not a generic star; 7 is also shown as 1+6 in contact chamber",
            "six_S_terms": "six action terms are downstream action petals, not primitive source nodes",
        },
        "contract": contract_dict(),
        "trajectory": projected_paths(),
        "source_digest": source_digest,
    }


def validate(geometry: Dict) -> Dict:
    checks = []
    c = geometry["contract"]
    fs = c["formula_stack"]
    checks.append({"name": "grade dimensions sum to 256", "passed": sum(c["grade_dimensions"]) == 256, "detail": str(c["grade_dimensions"])})
    checks.append({"name": "source alphabet has seven numbers", "passed": c["source_alphabet"] == [3,4,7,27,56,70,72], "detail": str(c["source_alphabet"])})
    checks.append({"name": "six action sectors visible", "passed": len(c["action_sectors"]) == 6, "detail": ", ".join(s["symbol"] for s in c["action_sectors"])})
    checks.append({"name": "five wounds preserved", "passed": len(c["wounds"]) == 5, "detail": ", ".join(c["wounds"])})
    checks.append({"name": "N_Q / W_Q triple", "passed": len(c["N_Q"]) == 3 and all(w > 0 for w in c["W_Q"]), "detail": str(c["W_Q"])})
    checks.append({"name": "phase matrices dimensions", "passed": len(geometry["trajectory"]["eta"]) == 8 and len(geometry["trajectory"]["omega"]) == 8, "detail": "8x8 η and Ω matrices"})
    checks.append({"name": "read uploaded ASHA gates", "passed": geometry["source_digest"].get("gate_markdown_count", 0) > 700, "detail": f"gate markdown count={geometry['source_digest'].get('gate_markdown_count')}"})
    checks.append({"name": "formula values finite", "passed": 0 < fs["lambda_asha"] < 1 and fs["v_over_mp"] > 0, "detail": f"lambda={fs['lambda_asha']}, v/mp={fs['v_over_mp']}"})
    passed = all(ch["passed"] for ch in checks)
    return {
        "validation": "PASS_GRANDBREAKING_ASHA_ONTOLOGY_V3" if passed else "FAIL_GRANDBREAKING_ASHA_ONTOLOGY_V3",
        "passed": passed,
        "checks": checks,
    }


def render(svg_path: Path, png_path: Path, geometry_path: Path, manifest_path: Path, source_root: Path):
    outdir = svg_path.parent
    outdir.mkdir(parents=True, exist_ok=True)
    digest_path = outdir / "asha_source_digest.json"
    source_digest = write_source_digest(source_root, digest_path)
    geometry = build_geometry_json(source_digest)
    manifest = validate(geometry)
    svg = SVG(W, H)
    svg.add_defs(svg_defs())
    add_background(svg)
    add_title(svg)
    add_formula_sidebar(svg)
    add_root_identity(svg)
    add_phase_octave(svg, geometry)
    add_contact_seven(svg)
    add_projector_chamber(svg)
    add_finite_engine(svg)
    add_action_and_observation(svg)
    add_interpanel_arrows(svg)
    add_source_alphabet_and_locks(svg)
    svg_path.write_text(svg.render(), encoding="utf-8")
    cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=W, output_height=H)
    geometry_path.write_text(json.dumps(geometry, indent=2), encoding="utf-8")
    manifest_path.write_text(json.dumps(manifest, indent=2), encoding="utf-8")
    return manifest


def main():
    p = argparse.ArgumentParser()
    p.add_argument("--source-root", default="/mnt/data/asha_src_fresh")
    p.add_argument("--outdir", default="/mnt/data/asha_grand_viz_toolkit/outputs")
    args = p.parse_args()
    outdir = Path(args.outdir)
    stem = "asha_grandbreaking_projector_universe_v3"
    manifest = render(
        svg_path=outdir / f"{stem}.svg",
        png_path=outdir / f"{stem}.png",
        geometry_path=outdir / f"{stem}.geometry.json",
        manifest_path=outdir / f"{stem}.manifest.json",
        source_root=Path(args.source_root),
    )
    print(json.dumps(manifest, indent=2))

if __name__ == "__main__":
    main()
