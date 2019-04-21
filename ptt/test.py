#!/usr/bin/python3
from performance import obj,var,writeJson

test =obj()
testVar = var("testK","testV")
test.addvar(testVar)
print(test)
writeJson("",test)