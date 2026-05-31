from __future__ import annotations

import html
from pathlib import Path
from math import pi

from .style import PALETTE, STYLE, FigureStyle
from .contact_vacuum import VacuumGeometry, validate_section2_geometry


def esc(s: str) -> str:
    return html.escape(s, quote=True)


def render_section2_svg(geometry: VacuumGeometry, out_path: str | Path, style: FigureStyle = STYLE) -> None:
    out_path = Path(out_path)
    p = PALETTE
    width, height = style.width, style.height
    checks = validate_section2_geometry(geometry)
    regions = {r.key: r for r in geometry.regions}
    pb = regions["P_B"]
    pg = regions["P_G"]
    k7 = regions["K_7"]

    parts: list[str] = []
    parts.append(f'''<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}" role="img" aria-labelledby="title desc">
<title id="title">ASHA Section 2 — Boolean/G2 Contact Vacuum K7</title>
<desc id="desc">Topological logarithmic visualization of two transparent projector lenses, P_B rank 56 and P_G rank 14, whose intersection K7 rank 7 is the bright payload. Visible areas scale as ln(rank).</desc>
<defs>
  <radialGradient id="bg" cx="50%" cy="45%" r="72%">
    <stop offset="0%" stop-color="#10101C"/>
    <stop offset="54%" stop-color="{p.abyss}"/>
    <stop offset="100%" stop-color="#010103"/>
  </radialGradient>
  <radialGradient id="payload" cx="50%" cy="48%" r="50%">
    <stop offset="0%" stop-color="#FFFFFF" stop-opacity="1"/>
    <stop offset="18%" stop-color="#FFF0B8" stop-opacity="0.92"/>
    <stop offset="48%" stop-color="#D9B45F" stop-opacity="0.34"/>
    <stop offset="100%" stop-color="#D9B45F" stop-opacity="0"/>
  </radialGradient>
  <radialGradient id="cyanGlass" cx="38%" cy="36%" r="72%">
    <stop offset="0%" stop-color="#E5FDFF" stop-opacity="0.38"/>
    <stop offset="58%" stop-color="#67E8F9" stop-opacity="0.095"/>
    <stop offset="100%" stop-color="#67E8F9" stop-opacity="0"/>
  </radialGradient>
  <radialGradient id="goldGlass" cx="62%" cy="34%" r="72%">
    <stop offset="0%" stop-color="#FFE8A9" stop-opacity="0.34"/>
    <stop offset="58%" stop-color="#D9B45F" stop-opacity="0.09"/>
    <stop offset="100%" stop-color="#D9B45F" stop-opacity="0"/>
  </radialGradient>
  <filter id="softGlow" x="-200%" y="-200%" width="500%" height="500%">
    <feGaussianBlur stdDeviation="7" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <filter id="largeGlow" x="-260%" y="-260%" width="620%" height="620%">
    <feGaussianBlur stdDeviation="28" result="blur"/>
    <feMerge><feMergeNode in="blur"/><feMergeNode in="SourceGraphic"/></feMerge>
  </filter>
  <filter id="glassBlur" x="-40%" y="-40%" width="180%" height="180%">
    <feGaussianBlur stdDeviation="0.55"/>
  </filter>
  <style><![CDATA[
    .title {{ font-family: {style.title_font}; font-size: 63px; letter-spacing: 0.4px; fill: {p.platinum}; }}
    .subtitle {{ font-family: {style.label_font}; font-size: 19px; letter-spacing: 4px; fill: {p.muted}; text-transform: uppercase; }}
    .math {{ font-family: {style.title_font}; font-size: 39px; fill: {p.platinum}; }}
    .label {{ font-family: {style.label_font}; font-size: 20px; letter-spacing: 3px; fill: #96A0B4; }}
    .caption {{ font-family: {style.label_font}; font-size: 22px; fill: #B6BFCE; }}
    .tiny {{ font-family: {style.label_font}; font-size: 15px; letter-spacing: 1.35px; fill: #8B93A5; }}
    .rank {{ font-family: {style.title_font}; font-size: 42px; fill: {p.gold}; }}
  ]]></style>
</defs>
<rect width="100%" height="100%" fill="url(#bg)"/>
<rect x="48" y="48" width="{width-96}" height="{height-96}" rx="52" fill="none" stroke="#2B2B36" stroke-opacity="0.34" stroke-width="1.4"/>
''')

    # Header.
    parts.append(f'''<text class="subtitle" x="{width/2:.1f}" y="132" text-anchor="middle">ASHA MANUSCRIPT · SECTION 2 · BOOLEAN/G2 CONTACT VACUUM</text>
<text class="title" x="{width/2:.1f}" y="214" text-anchor="middle">The Great Filtration into K7</text>
<text class="math" x="{width/2:.1f}" y="282" text-anchor="middle">U=Im(P_B), rank 56 · V=Im(P_G), rank 14 · K7=U∩V, rank 7</text>
''')

    # Faint Lambda4 chamber inherited from section 1.
    parts.append(f'<ellipse cx="{width/2:.1f}" cy="{k7.cy:.1f}" rx="700" ry="520" fill="none" stroke="{p.gold}" stroke-opacity="0.075" stroke-width="1.2" stroke-dasharray="4 18"/>' )
    parts.append(f'<text class="tiny" x="{width/2:.1f}" y="{k7.cy-575:.1f}" text-anchor="middle" fill="{p.gold}">LAMBDA^4 R^8 CHAMBER · 70-DIMENSIONAL BACKDROP</text>\n')

    # Transparent projector lenses.
    parts.append('<g id="projector_lenses">\n')
    parts.append(f'<ellipse cx="{pb.cx:.2f}" cy="{pb.cy:.2f}" rx="{pb.rx:.2f}" ry="{pb.ry:.2f}" fill="url(#cyanGlass)" stroke="{pb.stroke}" stroke-opacity="0.28" stroke-width="2.0" filter="url(#glassBlur)"/>\n')
    parts.append(f'<ellipse cx="{pg.cx:.2f}" cy="{pg.cy:.2f}" rx="{pg.rx:.2f}" ry="{pg.ry:.2f}" fill="url(#goldGlass)" stroke="{pg.stroke}" stroke-opacity="0.30" stroke-width="2.0" filter="url(#glassBlur)"/>\n')
    # Inner contour rings to give frosted refraction.
    for i, alpha in enumerate([0.13, 0.09, 0.055]):
        shrink = 1.0 - 0.08 * (i + 1)
        parts.append(f'<ellipse cx="{pb.cx:.2f}" cy="{pb.cy:.2f}" rx="{pb.rx*shrink:.2f}" ry="{pb.ry*shrink:.2f}" fill="none" stroke="{p.glass}" stroke-opacity="{alpha}" stroke-width="1.1"/>\n')
        parts.append(f'<ellipse cx="{pg.cx:.2f}" cy="{pg.cy:.2f}" rx="{pg.rx*shrink:.2f}" ry="{pg.ry*shrink:.2f}" fill="none" stroke="{p.gold}" stroke-opacity="{alpha}" stroke-width="1.1"/>\n')
    parts.append('</g>\n')

    # Payload K7: draw massive glow first, then exact log-scaled visible disk.
    parts.append(f'<circle cx="{k7.cx:.2f}" cy="{k7.cy:.2f}" r="{k7.rx*2.35:.2f}" fill="url(#payload)" opacity="0.98" filter="url(#largeGlow)"/>\n')
    parts.append(f'<circle cx="{k7.cx:.2f}" cy="{k7.cy:.2f}" r="{k7.rx:.2f}" fill="#F8F0C8" fill-opacity="0.86" stroke="#FFFFFF" stroke-opacity="0.96" stroke-width="2.3" filter="url(#softGlow)"/>\n')
    parts.append(f'<circle cx="{k7.cx:.2f}" cy="{k7.cy:.2f}" r="{k7.rx*0.56:.2f}" fill="#FFFFFF" fill-opacity="0.22"/>\n')

    # Thin convergence rays from projector identities to payload.
    for x1, y1, color, op in [
        (pb.cx-pb.rx*0.70, pb.cy-pb.ry*0.62, p.cyan, 0.18),
        (pb.cx-pb.rx*0.42, pb.cy+pb.ry*0.66, p.cyan, 0.12),
        (pg.cx+pg.rx*0.70, pg.cy-pg.ry*0.55, p.gold, 0.18),
        (pg.cx+pg.rx*0.46, pg.cy+pg.ry*0.67, p.gold, 0.12),
    ]:
        parts.append(f'<path d="M {x1:.1f},{y1:.1f} Q {k7.cx:.1f},{k7.cy:.1f} {k7.cx:.1f},{k7.cy:.1f}" fill="none" stroke="{color}" stroke-opacity="{op}" stroke-width="1.2"/>\n')

    # Labels connected with minimal leader lines.
    parts.append(f'''<g id="labels">
  <path d="M {pb.cx-pb.rx*0.55:.1f},{pb.cy-pb.ry*0.74:.1f} L 278,810" stroke="{p.cyan}" stroke-opacity="0.34" stroke-width="1"/>
  <text class="label" x="250" y="770" text-anchor="start" fill="{p.cyan}">BOOLEAN SUPPORT LENS</text>
  <text class="rank" x="250" y="826" text-anchor="start">P_B · 56</text>
  <text class="tiny" x="250" y="862" text-anchor="start">area ∝ ln(56)</text>

  <path d="M {pg.cx+pg.rx*0.54:.1f},{pg.cy-pg.ry*0.72:.1f} L 1320,830" stroke="{p.gold}" stroke-opacity="0.34" stroke-width="1"/>
  <text class="label" x="1360" y="790" text-anchor="start" fill="{p.gold}">G2 / OCTONIONIC SUPPORT LENS</text>
  <text class="rank" x="1360" y="846" text-anchor="start">P_G · 14</text>
  <text class="tiny" x="1360" y="882" text-anchor="start">area ∝ ln(14)</text>

  <path d="M {k7.cx:.1f},{k7.cy+k7.rx*0.98:.1f} L {width/2:.1f},1595" stroke="#FFFFFF" stroke-opacity="0.38" stroke-width="1.2"/>
  <text class="label" x="{width/2:.1f}" y="1650" text-anchor="middle" fill="#FFFFFF">CONTACT VACUUM PAYLOAD</text>
  <text class="rank" x="{width/2:.1f}" y="1710" text-anchor="middle" fill="#F8F0C8">K7 · 7</text>
  <text class="math" x="{width/2:.1f}" y="1765" text-anchor="middle" fill="#F8F0C8">K7 = Im(P_B) ∩ Im(P_G)</text>
</g>
''')

    # Lower proof strip / manifest values.
    area_pb = pi * pb.rx * pb.ry
    area_pg = pi * pg.rx * pg.ry
    area_k7 = pi * k7.rx * k7.ry
    parts.append(f'''<g transform="translate(220 2022)">
  <text class="caption" x="0" y="0">Exact rank chain:</text>
  <text class="math" x="0" y="58">56 → 14 → 7</text>
  <text class="caption" x="0" y="126">Visual scaling:</text>
  <text class="math" x="0" y="184">A(region) = S · ln(rank)</text>
  <text class="tiny" x="770" y="18">P_B/K7 area ratio = {area_pb/area_k7:.6f} = ln(56)/ln(7)</text>
  <text class="tiny" x="770" y="56">P_G/K7 area ratio = {area_pg/area_k7:.6f} = ln(14)/ln(7)</text>
  <text class="tiny" x="770" y="94">Boundary: topological/logarithmic visualization, not literal 56D/14D/7D geometry</text>
  <text class="tiny" x="770" y="132">Validation: {esc(checks['status'])}</text>
</g>
''')

    parts.append('</svg>')
    out_path.write_text(''.join(parts), encoding='utf-8')
