<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <span>Create New Token</span>
    </div>
    <div>
      <el-form ref="form" :model="myform" label-width="180px">
        <el-form-item label="token name">
          <el-input v-model="myform.name"></el-input>
        </el-form-item>
        <el-form-item label="token symbol">
          <el-input v-model="myform.symbol">
          </el-input>
        </el-form-item>
        <el-form-item label="total supply">
          <el-input-number v-model="myform.totall_supply" :min="1" :max="10000">
          </el-input-number>
        </el-form-item>
        <el-form-item label="your ethereum address">
          <el-input v-model="myform.eth_address" :disabled="true">
          </el-input>
        </el-form-item>
        <el-form-item label="enter wallet unlock password">
          <el-input v-model="myform.ethereum_password">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('myform')">Submmit</el-button>
          <el-button @click="resetForm('myform')">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from "jsappbase/ifetch"
import _ from "lodash"

export default {
  data() {
    return {
      myform: {
        name: "",
        symbol: "",
        totall_supply: 0,
        eth_address: "",
        ethereum_password: "",
      }
    }
  },
  mounted() {
    this.getUser()
  },
  methods: {
    getUser() {
      ifetch("/api/v1/users/session", "GET")
      .then((response) => {
        this.myform.eth_address = response.data.EthereumAddress
      })
    },
    createToken() {
      ifetch("/api/v1/tokens/new", "JSONPOST", this.myform)
      .then((response) => {
        this.$message({
          message: "success",
          type: 'success'
        })
        console.log(response)
      })
      .catch((response) => {
        console.log(response)
        this.$message.error(response)
      })
    },
    submitForm(formName) {
      console.log(formName)
      this.createToken()
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  }
}
</script>

<style>
  .el-input-number__decrease {
    height: 98%;
  }
  .el-icon-minus {
    top: 25%;
    position: relative;
  }
  .el-input-number__increase {
    height: 98%;
  }
  .el-icon-plus {
    top: 25%;
    position: relative;
  }
</style>
