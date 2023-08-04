import fs from "fs/promises";
import dotenv from "dotenv";

const readFile = async (pathFilename) => {
  try {
    const result = await fs.readFile(pathFilename, "utf8");
    return result;
  } catch (error) {
    console.error(error);
  }
};

const readJsonFile = async (pathFilename) => {
  try {
    const json = await readFile(pathFilename);
    return JSON.parse(json);
  } catch (error) {
    console.error(error);
  }
};

const writeJsonFile = async (filename, content) => {
  try {
    await fs.writeFile(filename, JSON.stringify(content, null, "\t"), "utf8");
    console.log("Successfully created fabric-config.json file");
  } catch (error) {
    console.error(error);
  }
};

const main = async () => {
  const fabloConfig = await readJsonFile("fablo-config.json");
  dotenv.config({ path: "./fablo-target/fabric-docker/.env" });

  const fabricConfig = {
    orgs: {},
    channels: {},
    chaincodes: fabloConfig.chaincodes,
  };

  await Promise.all(
    fabloConfig.orgs.map(async (organization, index) => {
      if (index === 0) return;

      const { name, domain } = organization.organization;

      const connectionProfileObj = await readJsonFile(`fablo-target/fabric-config/connection-profiles/connection-profile-${name.toLowerCase()}.json`);

      const connectionProfileOrgs = {
        [name]: {
          mspid: connectionProfileObj.organizations[name].mspid,
          peers: connectionProfileObj.organizations[name].peers.filter((peer) => peer.includes(name.toLowerCase())),
          certificateAuthorities: connectionProfileObj.organizations[name].certificateAuthorities,
        },
      };

      const connectionProfilePeers = {};
      await Promise.all(
        connectionProfileOrgs[name].peers.map(async (peer) => {
          const pem = await readFile(`fablo-target/fabric-config/crypto-config/peerOrganizations/${domain}/tlsca/tlsca.${domain}-cert.pem`);

          connectionProfilePeers[peer] = {
            url: connectionProfileObj.peers[peer].url,
            tlsCACerts: { pem },
            grpcOptions: {
              "ssl-target-name-override": peer,
              hostnameOverride: peer,
            },
          };
        })
      );

      const connectionProfileCertificateAuthorities = connectionProfileObj.certificateAuthorities;

      connectionProfileCertificateAuthorities[`ca.${domain}`].tlsCACerts = {
        pem: await readFile(`fablo-target/fabric-config/crypto-config/peerOrganizations/${domain}/tlsca/tlsca.${domain}-cert.pem`),
      };

      connectionProfileObj.client.connection = {
        timeout: {
          peer: {
            endorser: "500",
          },
        },
      };

      const connectionProfile = {
        name: connectionProfileObj.name,
        version: connectionProfileObj.version,
        client: connectionProfileObj.client,
        organization: connectionProfileOrgs,
        peers: connectionProfilePeers,
        certificateAuthorities: connectionProfileCertificateAuthorities,
      };

      const email = process.env[`${name.toUpperCase()}_CA_ADMIN_NAME`];
      const password = process.env[`${name.toUpperCase()}_CA_ADMIN_PASSWORD`];
      const msp = connectionProfileObj.organizations[name].mspid;

      fabricConfig.orgs[name] = {
        email,
        password,
        msp,
        connectionProfile: connectionProfile,
      };
    })
  );

  const channels = fabloConfig.channels.reduce((result, channel) => {
    const { name, orgs } = channel;

    orgs.map((org) => {
      if (!result[name]) {
        result[name] = {};
      }

      result[name][org.name] = org.peers;
    });

    return result;
  }, {});

  fabricConfig.channels = channels;

  await writeJsonFile("fabric-config.json", fabricConfig);
};

main();
