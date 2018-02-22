<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
           <span>GetBalance</span>
          </div>
          <div class="text item">
            <div>
              {{result.balanceResult}}
            </div>
            <el-button type="success" @click="getMyBalance()">Get My Balance</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
           <span>Get Rate</span>
          </div>
          <div class="text item">
            <div>
              {{result.converRate}}
            </div>
            <el-button type="success" @click="getConverRate()">Get Transfer Rate</el-button>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card class="box-card">
          <div slot="header" class="clearfix">
           <span>Token Transfer</span>
          </div>
          <div class="text item">
            <el-input v-model="ruleForm2.sendToAddress" auto-complete="off"></el-input>
            <el-input v-model="ruleForm2.amount" auto-complete="off"></el-input>
            <el-button type="success" @click="transferTo()">Transfer</el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from 'jsappbase/ifetch'
import _ from "lodash"
const Eth = require('ethjs-query')
const EthContract = require('ethjs-contract')
const BN = require('bn.js');
const $ = require("jQuery")

export default {
  data() {
    return {
      contractAbi: [],
      contractAddr: "",
      myContract: "",
      result: {
        balanceResult: "",
        converRate: "",
      },
      ruleForm2: {
        sendToAddress: "",
        amount: 0,
      },
      web3: {
        eth: null,
        contract: null,
        gas: 3000000,
      }
    }
  },
  mounted() {
    if (typeof web3 === 'undefined') {
      // set the provider you want from Web3.providers
      // web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
      alert("web3 not found")
      return
    }
    let contractAbiStr = $(".contract_abi").val()
    this.contractAbi = JSON.parse(contractAbiStr)
    this.contractAddr = $(".contract_addr").text()
    this.web3.eth = new Eth(web3.currentProvider)
    this.web3.contract = new EthContract(this.web3.eth)
  },
  methods: {
    // getContract() {
    //   let myContract = web3.eth.contract(this.contractAbi).at(this.contractAddr)
    //   return myContract
    // },
    getConverRate() {
      let self = this
      let myContract = this.getContract()
      myContract.getRate()
      .then((result) => {
        if(result.length != 0 ){
          self.result.converRate = result[0].words
          console.log(result)
        } else {
          self.result.converRate = "0"
          console.log(result)
        }
      })
      .catch((error) => {
        console.error(error)
        self.$message.error(error)
      })
    },
    getMyBalance() {
      let self = this
      let myContract = this.getContract()
      eval(`myContract.balanceOf(web3.eth.accounts[0])
      .then((result) => {
        if(result.length != 0 ){
          self.result.balanceResult = result[0].words
          console.log(result)
        } else {
          self.result.balanceResult = "0"
          console.log(result)
        }
      })
      .catch((error) => {
        console.error(error)
        self.$message.error(error)
      })`)
    },
    getContract () {
      const MiniToken = this.web3.contract(this.contractAbi)
      const miniToken = MiniToken.at(this.contractAddr)
      return miniToken
    },
    transferTo() {
      let ruleForm2 = this.ruleForm2
      let sendToAddress = ruleForm2.sendToAddress
      let amount = ruleForm2.amount
      if (sendToAddress === "" || amount === 0){
        this.$message.error("transfer error, please check your inputs")
        return
      }
      // amount = web3.toWei(amount, 'ether')
      console.log(sendToAddress, amount)
      let self = this
      let myContract = this.getContract()
      myContract.transfer(sendToAddress, amount, {from: web3.eth.accounts[0], gas: this.web3.gas}, (err, res) => {
        console.log(err, res)
      })
      // .then((txHash) => {
      //   self.result = txHash
      //   console.log(txHash)
      //   setTimeout(() => {
      //     self.waitForTxToBeMined(txHash)
      //   }, 3000)
      // })
      // .catch((error) => {
      //   console.error(error)
      //   self.$message.error(error)
      // })
    },
    async waitForTxToBeMined (txHash) {
      let txReceipt
      while (!txReceipt) {
        try {
          txReceipt = await web3.eth.getTransactionReceipt(txHash)
        } catch (err) {
          console.error(err)
          // return indicateFailure(err)
        }
      }
      // indicateSuccess()
      console.info("tx mined!")
    },
  }
}
</script>

<style>
  .box-card {
    height: 220px;
  }
</style>
