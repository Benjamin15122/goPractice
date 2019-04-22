#!/usr/bin/python3
from performance import obj,var,writeJson,graph

test =obj()
testVar = var("testK","testV")
testGraph = graph("testX","testY")
testGraph.addNode("2","3")
testGraph.addNode("1","4")
testGraph.addNode("5","7")
test.addvar(testVar)
test.addgraph(testGraph)
print(test)
writeJson("ptt.log",test)