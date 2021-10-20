// SPDX-License-Identifier: MIT

pragma solidity ^0.8.5;

import "./Ownable.sol";
import "./NimbleToken.sol";
import "./SyrupBar.sol";
import "./EventBar.sol";

contract MasterChef is Ownable {

    NimbleToken public nimble;
    SyrupBar public syrup;
    EventBar public eventbar;

    uint256 timeDeploy;
    uint256 year = 0;
    uint256 public daysCheckin;

    constructor(NimbleToken _nimble,SyrupBar _syrup,EventBar _event) public {
        syrup = _syrup;
        nimble = _nimble;
        eventbar = _event;
        timeDeploy = block.timestamp;

    }

    uint256 public workday;
    uint256 poolRewardYear = 50000000;   //50000
    uint256 pointCheckin = 30;
    uint256 pointVote = 30;

    function configMint(uint256 value)public onlyOwner returns(uint256) {
        poolRewardYear = value;
        return poolRewardYear;
    }

    function addWorkday(uint256 value)public onlyOwner returns(uint256) {
        workday = value;
        return workday;
    }

    function mint()public {
        //time stamp
        uint256 timeLock = timeDeploy + ( year * 365 days);
        require(block.timestamp >= timeLock,"not time yet");
        nimble.mint(address(syrup), poolRewardYear);
        year++;
    }
    function autoClaimCheckin(address[] memory account)public onlyOwner{
        //timestamp
        uint256 timeLock = timeDeploy + ( daysCheckin * 1 days);
        require(block.timestamp >= timeLock,"not time yet");
        uint256 poolReward = poolRewardYear*pointCheckin/100;       // 50000 * 20/100 = 10000
        uint256 rewardDay = poolReward/workday;     // 10000/200 = 50
        uint256 rewardTotal = rewardDay/account.length; // 50/3  = 16.666/คน
        for(uint i=0;i<account.length;i++){
            safeTransfer(account[i],rewardTotal);
        }
    }

    //---------------- Vote ----------------
    struct Vote {
        uint256 usedVote;
        uint256 score;
    }
    event getScoreVote(address[] listWallet);
    event LogVote(address indexed from ,address indexed _to , uint256 score);

    mapping(address => Vote)  scoreWallet;



    function addVote (address[] memory listWallet) public onlyOwner{
        for(uint256 i=0 ;i<listWallet.length;i++){
            scoreWallet[listWallet[i]].usedVote = 4;
        }
        emit getScoreVote(listWallet);
    }

    function getScore(address wallet) external view returns (uint256) {
        return scoreWallet[wallet].score;
    }
    function getRightScore(address wallet) external view returns (uint256) {
        return scoreWallet[wallet].usedVote;
    }

    function vote(address _to , uint score) public {
        require(scoreWallet[msg.sender].usedVote >= score,"not enough score");
        scoreWallet[msg.sender].usedVote = scoreWallet[msg.sender].usedVote - score;
        scoreWallet[_to].score = scoreWallet[_to].score + score;
        emit LogVote(msg.sender,_to,score);
    }

    function claimRewardScoreVote(address[] memory account)public onlyOwner{
        uint256 totalScoreVote;
        for(uint i = 0;i<account.length;i++){
            totalScoreVote = totalScoreVote + scoreWallet[account[i]].score;
        }
        uint256 rewardPool = poolRewardYear*pointVote/100;
        uint256 rewardPoint = rewardPool/totalScoreVote;

        for(uint j =0;j<account.length;j++){
            uint256 reward = scoreWallet[account[j]].score*rewardPoint;
            safeTransfer(account[j],reward);
        }
    }


    //---------------- event ----------------
    struct Event {
        address createBy;
        string nameEvent;
        string detial;
        uint reward;
        uint status;    //uint
        uint timeStart;
    }

    struct Participant{
        address[] join;
    }

    struct UserEventInfo {
        uint[] idEvent;
    }

    Event[] public eventInfo;
    mapping(uint256 => Participant) participantInfo;    //mapping eid กับคนเข้าร่วม
    mapping(address => UserEventInfo) userInfo;     //mapping address กับประวัติการเข้าร่วม event

    event _createEvent(address _createBy ,string _nameEvent,string _detial, uint _reward,uint _timeStart);
    event _join(address from,uint eid);

    function createEvent(string memory _nameEvent,string memory _detial, uint _reward,uint _timeStart) public returns(uint){
        require(_reward <= nimble.balanceOf(msg.sender),"you don't have token");

        safeTransferFromEvent(msg.sender,_reward);

        eventInfo.push(Event({
        createBy: msg.sender,
        nameEvent: _nameEvent,
        detial: _detial,
        reward: _reward,
        status: 0,
        timeStart: _timeStart
        }));

        emit _createEvent(msg.sender,_nameEvent,_detial,_reward,_timeStart);

        return eventInfo.length-1;
    }

    function createEventAdmin(string memory _nameEvent,string memory _detial, uint _reward,uint _timeStart) public onlyOwner returns(uint){
        require(_reward <= nimble.balanceOf(address(syrup)),"you don't have token");
        eventInfo.push(Event({
        createBy: msg.sender,
        nameEvent: _nameEvent,
        detial: _detial,
        reward: _reward,
        status: 0,
        timeStart: _timeStart
        }));

        emit _createEvent(msg.sender,_nameEvent,_detial,_reward,_timeStart);

        return eventInfo.length-1;
    }

    function startEvent(uint eid) public{
        require(eventInfo[eid].createBy == msg.sender, "access denied");
        require(eventInfo[eid].status == 0, "status invalid");
        require(eventInfo[eid].timeStart <= block.timestamp,"It's not time yet");
        eventInfo[eid].status = 1;
    }


    function searchEvent()public returns(Event[] memory){
        return eventInfo;
    }

    function searchEventByAddress(address from)public returns(uint[] memory){
        return userInfo[from].idEvent;
    }

    function searchParticipant(uint256 eid)public returns(address[] memory){
        return participantInfo[eid].join;
    }

    function joinEvent(uint eid)public{
        require(eid<=eventInfo.length-1,"Event not real");
        require(eventInfo[eid].status == 1,"Event can't join");
        userInfo[msg.sender].idEvent.push(eid);
        participantInfo[eid].join.push(msg.sender);
        emit _join(msg.sender,eid);
    }

    function searchEventAdmin()public  returns(uint[] memory){
        uint id = 0;
        uint[] memory result;

        for(uint i=0;i<=eventInfo.length;i++){
            if(eventInfo[i].createBy == owner()){
                result[id] = i;
                id++;
            }
        }
        return result;
    }

    function closeEvent(uint eid)public{
        require(eventInfo[eid].createBy == msg.sender, "access denied");
        require(eventInfo[eid].status == 1, "status invalid");
        eventInfo[eid].status = 0;
    }

    function AcceptEvent(uint256 eid,address[] memory wallet)public{
        require(eventInfo[eid].createBy == msg.sender, "access denied");
        require(eventInfo[eid].status == 1, "status invalid");
        eventInfo[eid].status = 2;
        //transfer

        uint reward = eventInfo[eid].reward/wallet.length;
        for(uint i=0;i< wallet.length;i++){
            safeTransferEvent(wallet[i],reward);
        }
    }

    function AcceptEventAdmin(uint256 eid,address[] memory wallet)public{
        require(eventInfo[eid].createBy == msg.sender, "access denied");
        require(eventInfo[eid].status == 1, "status invalid");
        eventInfo[eid].status = 3;
        //transfer
        uint reward = eventInfo[eid].reward/wallet.length;
        for(uint i=0;i< wallet.length;i++){
            safeTransfer(wallet[i],reward);
        }
    }

    function safeTransfer(address _to, uint256 _amount) internal {
        syrup.safeNimbleTransfer(_to, _amount);
    }

    function safeTransferEvent(address _to, uint256 _amount) internal {
        eventbar.safeNimbleTransfer(_to, _amount);
    }

    function safeTransferFromEvent(address _from,uint256 _amount) internal{
        eventbar.safeTransferFrom(_from,_amount);
    }


}
