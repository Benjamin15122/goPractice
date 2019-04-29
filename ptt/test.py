#!/usr/bin/python3
#coding:utf-8

# from performance import performance

# test = performance("/Users/wangguochang/Desktop/petWeb/src/lib")
# test.pr_curve.append_node(0.0,1.00)
# test.pr_curve.append_node(0.1,0.95)
# test.pr_curve.append_node(0.15,0.94)
# test.pr_curve.append_node(0.18,0.93)
# test.pr_curve.append_node(0.2,0.91)
# test.pr_curve.append_node(0.3,0.89)
# test.pr_curve.append_node(0.4,0.85)
# test.pr_curve.append_node(0.5,0.80)
# test.pr_curve.append_node(0.6,0.70)
# test.pr_curve.append_node(0.7,0.58)
# test.pr_curve.append_node(0.8,0.53)
# test.pr_curve.append_node(0.9,0.35)
# test.pr_curve.append_node(1.0,0.00)

# test.write_to_json()

import matplotlib
#matplotlib.use("Agg")
import numpy as np
import matplotlib.pyplot as plt
from sklearn.metrics import precision_recall_curve
plt.figure(1) # 创建图表1
plt.title('Precision/Recall Curve')# give plot a title
plt.xlabel('Recall')# make axis labels
plt.ylabel('Precision')
 
#y_true和y_scores分别是gt label和predict score
y_true = np.array([0, 0, 1, 1])
y_scores = np.array([0.1, 0.4, 0.35, 0.8])
precision, recall, thresholds = precision_recall_curve(y_true, y_scores)
plt.figure(1)
plt.plot(precision, recall)
plt.savefig('__out/auc.png')
plt.show()