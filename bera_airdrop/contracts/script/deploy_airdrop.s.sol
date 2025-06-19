// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import {Script, console} from "forge-std/Script.sol";
import {Airdrop} from "../src/airdrop.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

/*

# prepare .env file
DEPLOYER=<deployer-account-name>
DEPLOYER_ADDRESS=<deployer-address>
PROXY_ADMIN=<proxy-admin-address>
ADMIN=<admin-address>
ACTIVATION_DELAY=<activation-delay-seconds>
TOKEN_ADDRESS=<token-address>
EVM_RPC=<evm-rpc>
ETHERSCAN_API_KEY=<etherscan-api-key>
ETHERSCAN_API_URL=<etherscan-api-url>

# source .env
# verify source code
forge script -vvvv \
    --account $DEPLOYER \
    --sender $DEPLOYER_ADDRESS \
    -f $EVM_RPC \
    --broadcast \
    --verify \
    --verifier custom \
    --verifier-api-key $ETHERSCAN_API_KEY \
    --verifier-url $ETHERSCAN_API_URL \
    script/deploy_airdrop.s.sol:DeployAirdrop

# verify source code using flatted code
forge script -vvvv \
    --account $DEPLOYER \
    --sender $DEPLOYER_ADDRESS \
    -f $EVM_RPC \
    --broadcast \
    script/deploy_airdrop.s.sol:DeployAirdrop
*/

contract DeployAirdrop is Script {
    function run() external {
        // Read all required parameters from environment variables
        address deployer = vm.envAddress("DEPLOYER_ADDRESS");
        address proxyAdmin = vm.envAddress("PROXY_ADMIN");
        address admin = vm.envAddress("ADMIN");
        uint32 activationDelay = uint32(vm.envUint("ACTIVATION_DELAY"));
        address tokenAddress = vm.envAddress("TOKEN_ADDRESS");

        vm.startBroadcast(deployer);

        // Print deployment parameters
        console.log("=== Deployment Parameters ===");
        console.log("Deployer:", deployer);
        console.log("ProxyAdmin:", proxyAdmin);
        console.log("Admin:", admin);
        console.log("ActivationDelay:", activationDelay);
        console.log("TokenAddress:", tokenAddress);

        if (proxyAdmin == address(0x0)) {
            ProxyAdmin _pa = new ProxyAdmin();
            _pa.transferOwnership(address(admin));
            proxyAdmin = address(_pa);
            console.log("\n=== Deploy new ProxyAdmin ===");
            console.log("ProxyAdmin deployed at:", proxyAdmin);
        }

        // Deploy implementation contract
        Airdrop implementation = new Airdrop();
        console.log("\n=== Deployment Results ===");
        console.log("Implementation deployed at:", address(implementation));

        // Prepare initialization data
        bytes memory initData =
            abi.encodeWithSelector(Airdrop.initialize.selector, activationDelay, admin, tokenAddress);

        // Deploy proxy contract using existing ProxyAdmin
        TransparentUpgradeableProxy proxy =
            new TransparentUpgradeableProxy(address(implementation), proxyAdmin, initData);
        console.log("Proxy deployed at:", address(proxy));
        vm.stopBroadcast();
    }
}
