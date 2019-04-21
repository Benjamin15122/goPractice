#!/usr/bin/python3
# Filename: performance.py
import json

def writeJson(dir,item):
    dirlen = len(dir)
    if dirlen!=0 and dir[dirlen-1]!="/":
        dir = dir + "/"
    with open(dir+"out.json","w") as f:
        json.dump(item.getObj(),f)

class var:
    key: str
    value: str
    def __init__(self,key:str,value:str):
        self.key = key
        self.value = value

class obj:
    __varList: dict = {}

    def addvar(self,item):
        self.__varList[item.key]=item.value
    def getObj(self)-> dict:
        return {"varlist":self.__varList}