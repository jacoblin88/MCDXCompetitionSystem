import subprocess
import yaml
import jinja2
import json
import os

keyGenPath = 'cmd/KeyGen/main.go'
machineInfoPath = "configs/AutoGenYaml/machineInfo.yaml"
masterKeyGoPath = "internal/pkg/SecretKey"
slavePath = "cmd/FlagAgent"
keyPath = "configs/AutoGenKey"
jinjaTemplatePath = "scripts/template"

def ReadYml(path):
    with open(path, "r") as stream:
        try:
            yml = yaml.safe_load(stream)
        except yaml.YAMLError as exc:
            print(exc)
    return yml

def GetKeyJson(path):
    keyJson = subprocess.check_output(['go', 'run', path])
    return keyJson.decode("utf8")

def GetKeyStruct(num):
    key = {}
    for i in range(num):
        keyJson = json.loads(GetKeyJson(keyGenPath))
        publicKey = {
            "N": keyJson["N"],
            "E": keyJson["E"]
            }
        privateKey = keyJson
        key[i] = {"publicKey": json.dumps(publicKey), "privateKey": json.dumps(privateKey)}
    return key

def main():
    yml = ReadYml(machineInfoPath)
    print("start generate key...")
    key = GetKeyStruct(len(yml)+1)
    print("finish generate key...")

    templateLoader = jinja2.FileSystemLoader(searchpath=jinjaTemplatePath)
    templateEnv = jinja2.Environment(loader=templateLoader)
    template = templateEnv.get_template("flagAgent.go.jinja")
    for i in yml:
        with open(os.path.join(slavePath, i["ip"] + ".go"), "w") as f:
            result = template.render(privateKey=key[i["machineid"]]["privateKey"].replace("\"","\\\""), publlicKey=key[0]["publicKey"].replace("\"","\\\""), ip=i["ip"], port=i["flag"]["port"] , tryTimes="3")
            f.write(result)
        with open(os.path.join(keyPath, i["ip"]) + ".key", "w") as f:
            f.write(key[i["machineid"]]["privateKey"])
    with open(os.path.join(keyPath, "Master") + ".key", "w") as f:
        f.write(key[0]["privateKey"])
    keyList = []
    for i in range(len(yml)+1):
        keyList.append(key[i]) 
    template = templateEnv.get_template("key.go.jinja")
    with open(os.path.join(masterKeyGoPath, "key.go"), "w") as f:
        result = template.render(items=keyList)
        f.write(result)

    


if __name__=="__main__":
    main()