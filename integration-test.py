import subprocess
from requests import get
from os import environ
import sys


def writeFile(filepath: str, content: str):
    f = open(filepath, "w")
    f.write(content)
    f.close()


def setTestEnvVaribles():
    environ["TECHNITIUM_URL"] = "http://localhost:5380"
    environ["TECHNITIUM_USER"] = "admin"
    environ["TECHNITIUM_PASSWORD"] = "admin"
    environ["ANSIBLE_INV_PATH"] = "./inventory.yaml"
    environ["ZONE_CONF_PATH"] = "zone-config.yaml"


def get_record_ip(records) -> str:
    for record in records:
        if record["name"] == "minastirith.gondor.middleearth":
            return record["rData"]["ipAddress"]
    return ""


def main():
    zoneConfigPath = "zone-config.yaml"
    zoneConfigContent = "zone: gondor.middleearth\ntype: Primary"

    ansibleInventoryPath = "inventory.yaml"
    ansibleInventoryContent = "minastirith:\n  hosts:\n    5.5.5.5:"

    writeFile(zoneConfigPath, zoneConfigContent)
    writeFile(ansibleInventoryPath, ansibleInventoryContent)

    setTestEnvVaribles()

    subprocess.run(["go", "run", "cmd/technapi/main.go"], check=True)

    x = get("http://localhost:5380/api/user/login?user=admin&pass=admin")
    sessionToken = x.json()["token"]

    x = get(f"http://localhost:5380/api/zones/list?token={sessionToken}")
    zones = x.json()["response"]["zones"]

    if not any("gondor.middleearth" in zone["name"] for zone in zones):
        print("The zone gondor.middleearth couldn't be find in technitium.")
        return 1

    x = get(f"http://localhost:5380/api/zones/records/get?domain=gondor.middleearth&listZone=true&token={sessionToken}")
    records = x.json()["response"]["records"]

    if not any("minastirith.gondor.middleearth" in record["name"] for record in records):
        print("The record minastirith.gondor.middleearth couldn't be found in the zone.")
        return 1

    ipAddress = get_record_ip(records)
    if ipAddress != "":
        if ipAddress != "5.5.5.5":
            print("The A record points to the wrong ip address")
            return 1
    else:
        print("Couldn't retrieve information of the record")
        return 1

    return 0


if __name__ == "__main__":
    sys.exit(main())
