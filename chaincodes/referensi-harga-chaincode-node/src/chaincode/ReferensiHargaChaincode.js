const stringify = require("json-stringify-deterministic");
const sortKeysRecursive = require("sort-keys-recursive");
const { Contract } = require("fabric-contract-api");

class ReferensiHargaChaincode extends Contract {
  // CreateAsset issues a new asset to the world state with given details.
  async CreateAsset(ctx, id, umurTanam, harga) {
    // Demonstrate the use of Attribute-Based Access Control (ABAC) by checking
    // to see if the caller has the "abac.creator" attribute with a value of true;
    // if not, return an error.

    const isDinas = await ctx.getClientIdentity().getAttributeValue("dinas.DINAS");
    if (!isDinas) {
      throw new Error("Submitting client not authorized to create asset, does not have DINAS role");
    }

    const exists = await this.AssetExists(ctx, id);
    if (exists) {
      throw new Error(`The asset ${id} already exists`);
    }

    // Get ID of submitting client identity
    const idDinas = await this.GetSubmittingClientIdentity(ctx);

    const asset = {
      id,
      idDinas,
      umurTanam,
      harga,
      tanggalPembaruan: new Date().toISOString(),
    };

    // we insert data in alphabetic order using 'json-stringify-deterministic' and 'sort-keys-recursive'
    await ctx.stub.putState(id, Buffer.from(stringify(sortKeysRecursive(asset))));

    return JSON.stringify(asset);
  }

  // GetAllAssets returns all assets found in the world state.
  async GetAllAssets(ctx) {
    const allResults = [];
    // range query with empty string for startKey and endKey does an open-ended query of all assets in the chaincode namespace.
    const iterator = await ctx.stub.getStateByRange("", "");
    let result = await iterator.next();
    while (!result.done) {
      const strValue = Buffer.from(result.value.value.toString()).toString("utf8");
      let record;
      try {
        record = JSON.parse(strValue);
      } catch (err) {
        console.log(err);
        record = strValue;
      }
      allResults.push(record);
      result = await iterator.next();
    }
    return JSON.stringify(allResults);
  }

  // AssetExists returns true when asset with given ID exists in world state.
  async AssetExists(ctx, id) {
    const assetJSON = await ctx.stub.getState(id);
    return assetJSON && assetJSON.length > 0;
  }

  async GetSubmittingClientIdentity(ctx) {
    const clientId = ctx.clientIdentity.getID();
    return clientId;
  }
}

module.exports = ReferensiHargaChaincode;
