from asha_figures.boolean_lattice import build_nodes, build_edges, validate
from asha_figures.style import STYLE


def test_section1_exact_counts():
    nodes = build_nodes(STYLE.width, STYLE.margin_x, STYLE.tier_top, STYLE.tier_bottom,
                        STYLE.node_radius, STYLE.middle_node_radius, STYLE.terminal_node_radius)
    edges = build_edges(nodes)
    checks = validate(nodes, edges)
    assert checks["node_count_exact"]
    assert checks["tier_counts_exact"]
    assert checks["edge_count_exact"]
