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

class graph:
    __scaley: str = "y"
    __scalex: str = "x"
    __node: list = []
    def __init__(self,scalex:str,scaley:str):
        self.__scaley = scaley
        self.__scalex = scalex
    def addNode(self,x:str,y:str):
        self.__node.append({self.__scalex:x,self.__scaley:y})
    def getObj(self)-> dict:
        return {
            "scaley": self.__scaley,
            "scalex": self.__scalex,
            "node": self.__node
        }

class obj:
    __varList: dict = {}
    __graphList: list = []

    def addvar(self,item):
        self.__varList[item.key]=item.value
    def addgraph(self,item):
        self.__graphList.append(item.getObj())
    def getObj(self)-> dict:
        return {"varlist":self.__varList,"graphlist":self.__graphList}