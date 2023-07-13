import os
import json

from dotenv import load_dotenv
load_dotenv("./fablo-target/fabric-docker/.env")

fablo_dict = json.load(open("./fablo-config.json", "r"))
json_config_file = {
    "orgs": {},
    "channels": {},
    "chaincodes": fablo_dict["chaincodes"],
}

# Add orgs
for index, org in enumerate(fablo_dict["orgs"]):
    if index == 0:
        continue

    org_name = org["organization"]["name"]
    domain = org["organization"]["domain"]

    connectionProfile = json.load(
        open(
            f"./fablo-target/fabric-config/connection-profiles/connection-profile-{org_name.lower()}.json",
            "r",
        )
    )

    connectionProfile["certificateAuthorities"]["ca." + domain]["tlsCACerts"] = str(
        open(
            f"./fablo-target/fabric-config/crypto-config/peerOrganizations/{domain}/tlsca/tlsca.{domain}-cert.pem",
            "r",
        ).read()
    )

    certificate = open(
        f"./fablo-target/fabric-config/crypto-config/peerOrganizations/{domain}/users/User1@{domain}/msp/signcerts/User1@{domain}-cert.pem",
        "r",
    )

    privateKey = open(
        f"./fablo-target/fabric-config/crypto-config/peerOrganizations/{domain}/users/User1@{domain}/msp/keystore/priv-key.pem",
        "r",
    )

    email = os.getenv("{}_CA_ADMIN_NAME".format(org_name.upper()))
    password = os.getenv("{}_CA_ADMIN_PASSWORD".format(org_name.upper()))
    msp = "{}MSP".format(org["organization"]["name"])
    connectionProfile = json.dumps(connectionProfile)
    certificate = str(certificate.read())
    privateKey = str(privateKey.read())

    json_config_file["orgs"][org_name] = {
        "email": email,
        "password": password,
        "msp": msp,
        "connectionProfile": connectionProfile,
        "certificate": certificate,
        "privateKey": privateKey,
    }

# Add channels
for channel in fablo_dict["channels"]:
    json_config_file["channels"][channel["name"]] = {
        "acceptOrgs": [org["name"] for org in channel["orgs"]]
    }

# Saving env file
with open("./env.json", "w") as outfile:
    outfile.write(json.dumps(json_config_file, indent=2))
