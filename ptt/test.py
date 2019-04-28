#!/usr/bin/python3
from performance import performance

test = performance("/Users/wangguochang/Desktop/petWeb/src/lib")
test.pr_curve.append_node(0.0,1.00)
test.pr_curve.append_node(0.1,0.95)
test.pr_curve.append_node(0.15,0.94)
test.pr_curve.append_node(0.18,0.93)
test.pr_curve.append_node(0.2,0.91)
test.pr_curve.append_node(0.3,0.89)
test.pr_curve.append_node(0.4,0.85)
test.pr_curve.append_node(0.5,0.80)
test.pr_curve.append_node(0.6,0.70)
test.pr_curve.append_node(0.7,0.58)
test.pr_curve.append_node(0.8,0.53)
test.pr_curve.append_node(0.9,0.35)
test.pr_curve.append_node(1.0,0.00)

test.write_to_json()