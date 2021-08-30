import {BridgeBank, IbcToken, IbcToken__factory} from "../build";
import {DependencyContainer} from "tsyringe";
import fs from "fs";
import * as hardhat from "hardhat";
import web3 from "web3";

interface TokenAddress {
    address: string
}

async function attachIbcToken(bridgeBank: BridgeBank, token: IbcToken) {
    return await bridgeBank.addExistingBridgeToken(token.address)
}

export async function processTokenData(bridgeBank: BridgeBank, filename: string, container: DependencyContainer) {
    const fileContents = fs.readFileSync(filename, {encoding: "utf8"})

    for (const line of fileContents.split(/\r?\n+/)) {
        if ((line ?? "") === "")
            continue
        const data = JSON.parse(line) as TokenAddress
        const token = await hardhat.ethers.getContractAt("IbcToken", data.address) as IbcToken
        await attachIbcToken(bridgeBank, token)
        const result = {
            ownedByBridgeBank: token.address,
            addExistingBridgeTokenCalled: true
        }
        console.log(JSON.stringify(result))
    }
}

interface TokenData {
    symbol: string,
    name: string,
    decimals: number,
    cosmosDenom: string
}

const MINTER_ROLE: string = web3.utils.soliditySha3('MINTER_ROLE') ?? "0xBADBAD"  // this should never fail
if (MINTER_ROLE == "0xBADBAD")
    throw Error("failed to get MINTER_ROLE")
const DEFAULT_ADMIN_ROLE = '0x0000000000000000000000000000000000000000000000000000000000000000' // to bridgebank

async function buildIbcToken(tokenFactory: IbcToken__factory, tokenData: TokenData, bridgeBank: BridgeBank) {
    const newToken = await tokenFactory.deploy(tokenData.name, tokenData.symbol, tokenData.decimals, tokenData.cosmosDenom)
    console.log(JSON.stringify({deployed: await newToken.address, symbol: await newToken.symbol()}))
    await newToken.grantRole(DEFAULT_ADMIN_ROLE, bridgeBank.address)
    console.log(JSON.stringify({roleGrantedToBridgeBank: DEFAULT_ADMIN_ROLE}))
    await newToken.grantRole(MINTER_ROLE, bridgeBank.address)
    console.log(JSON.stringify({roleGrantedToBridgeBank: MINTER_ROLE}))
    await newToken.renounceRole(MINTER_ROLE, await tokenFactory.signer.getAddress())
    console.log(JSON.stringify({roleRenouncedByDeployer: MINTER_ROLE}))
    await newToken.renounceRole(DEFAULT_ADMIN_ROLE, await tokenFactory.signer.getAddress())
    console.log(JSON.stringify({roleRenouncedByDeployer: DEFAULT_ADMIN_ROLE}))
    await newToken.transferOwnership(bridgeBank.address)
    return newToken
}

export async function buildIbcTokens(ibcTokenFactory: IbcToken__factory, tokens: TokenData[], bridgeBank: BridgeBank) {
    const result = []
    for (const t of tokens) {
        const newToken = await buildIbcToken(ibcTokenFactory, t, bridgeBank)
        const tokenResult = {
            address: newToken.address,
            symbol: await newToken.symbol(),
        }
        console.log(JSON.stringify({...tokenResult, complete: true}))
        result.push(tokenResult)
    }
    return result
}

export async function readTokenData(filename: string): Promise<TokenData[]> {
    const result = fs.readFileSync(filename, {encoding: "utf8"})
    return JSON.parse(result) as TokenData[]
}
