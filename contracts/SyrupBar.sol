// SPDX-License-Identifier: MIT

pragma solidity ^0.8.5;

import "./ERC20.sol";
import "./NimbleToken.sol";

contract SyrupBar is ERC20('SyrupBar Token', 'SYRUP') {

    NimbleToken public nimble;


    constructor(NimbleToken _nimble) public {
        nimble = _nimble;

    }

    function safeNimbleTransfer(address _to, uint256 _amount) public onlyOwner {
        uint256 nimbleBal = nimble.balanceOf(address(this));
        if (_amount > nimbleBal) {
            nimble.transfer(_to, nimbleBal);
        } else {
            nimble.transfer(_to, _amount);
        }
    }

}