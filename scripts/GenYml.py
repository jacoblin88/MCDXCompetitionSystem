from MCDXEnvConfig import *
import jinja2
import os

def genStruct():
  allMachine = []
  allService = []
  machineId = 0
  serviceId = 0
  for i in range(len(teamSubnet)):
    for j in range(len(teamMachine)):
      machineId += 1
      for k in machine:
        if k["name"] != teamMachine[j]["name"]:
          continue
        macInfo = {
            "machineid": machineId,
            "flag": {
                "path": teamMachine[j]["flagPath"],
                "port": 8443,
            },
            "teamid": i+1,
            "ip": teamSubnet[i]+teamMachine[j]["ip_ext"],
            "event": {
                "flagupdatefail": {
                    "atk": 0,
                    "def": 0,
                },
                "flagupdatesuc": {
                    "atk": 0,
                    "def": 0,
                },
                "flagloss": {
                    "atk": 0,
                    "def": sum([l["offline"] for l in k["Service"]]),
                },
                "flagexist": {
                    "atk": 0,
                    "def": 0,
                },
                "flagsubmit": {
                    "atk": 0,
                    "def": 0,
                },
                "resetsys": {
                    "atk": 0,
                    "def": sum([l["offline"] for l in k["Service"]])*3,
                },
            }
        }
        allMachine.append(macInfo)
        for l in k["Service"]:
            serviceId += 1
            servInfo = {
                "serviceid": serviceId,
                "teamid": i+1,
                "machineid": machineId,
                "ip": teamSubnet[i]+teamMachine[j]["ip_ext"],
                "port": l["port"],
                "checkpatch": l["checkpatch"],
                "checknormal": l["checknormal"],
                "event": {
                    "serviceloss":{
                        "atk": 0,
                        "def": l["offline"],
                    },
                    "servicenormal":{
                        "atk": 0,
                        "def": 0,
                    },
                    "servicepatch":{
                        "atk": 0,
                        "def": l["patch"],
                    },
                    "servicenotpatch":{
                        "atk": 0,
                        "def": 0,
                    },
                    "serviceblock":{
                        "atk": 0,
                        "def": l["block"],
                    },
                    "servicenotblock":{
                        "atk": 0,
                        "def": 0,
                    },
                }
            }
            allService.append(servInfo)
        break
  return (allMachine, allService)


def main():
    (allMachine, allService) = genStruct()
    configsPath = "configs/AutoGenYaml"
    configsTemplatePath = "scripts/template"

    templateLoader = jinja2.FileSystemLoader(searchpath=configsTemplatePath)
    templateEnv = jinja2.Environment(loader=templateLoader)
    template = templateEnv.get_template("machineInfo.yaml.jinja")
    with open(os.path.join(configsPath, "machineinfo.yaml"), "w") as f:
        result = template.render(machineInfo=allMachine)
        f.write(result)
    
    template = templateEnv.get_template("serviceInfo.yaml.jinja")
    with open(os.path.join(configsPath, "serviceinfo.yaml"), "w") as f:
        result = template.render(serviceInfo=allService)
        f.write(result)
    return 


if __name__ == "__main__":
  main()
