#!/usr/bin/python3
# Filename: performance.py
import json

class graph:
    __data: list = []
    __options: dict = {}

    def __init__(self, options: dict):
        self.__options = options

    def append_node(self, x, y):
        self.__data.append({"x":x,"y":y})

    def get_dict(self) -> dict:
        graph = self.__options
        graph["data"] = [{
            "type": "spline",
            "dataPoints": self.__data
        }]
        return graph

class dictionary:
    __key: str = ""
    __value: dict = {}

    def __init__(self, key: str,value: dict):
        self.__key = key
        self.__value = value

    def get_dict(self) -> dict:
        return {self.__key:self.__value}
    
    def attach(self, key: str, value):
        self.__value[key] = value

    def remove(self, key: str):
        del self.__value[key]

    def load_from_json(self, path: str):
        with open(path,"r") as f:
            load_dict = json.load(f)
            self.__value.update(load_dict)
            print("load from file ",path," succeeded")

class performance:
    params = dictionary("parameters",{})
    evals = dictionary("evaluation",{})
    pr_curve = graph({
        "animationEnabled": "true",
		"title":{
			"text": "PR-Curve"
		},
		"axisX": {
            "title": "recall"
		},
		"axisY": {
			"title": "precision"
		},
    })
    __directory: str = ""

    def __init__(self, directory: str):
        self.__directory = directory

    def write_to_json(self):
        output = {}
        output.update(self.params.get_dict())
        output.update(self.evals.get_dict())
        output.update({"pr_curve":self.pr_curve.get_dict()})
        dirlen = len(self.__directory)
        if dirlen!=0 and self.__directory[dirlen-1]!="/":
            self.__directory = self.__directory + "/"
        with open(self.__directory+"out.json","w") as f:
            json.dump(output,f)