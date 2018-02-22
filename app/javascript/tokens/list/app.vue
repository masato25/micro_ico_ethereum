<template>
  <el-table
  :data="tableData"
  border
  style="width: 100%">
    <el-table-column
      prop="TokenName"
      label="name">
    </el-table-column>
    <el-table-column
      prop="TokenSymbol"
      label="symbol">
      <template slot-scope="scope">
        <el-tag type="success">
          {{scope.row.TokenSymbol}}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column
      prop="TotallSupply"
      label="TotallSupply">
    </el-table-column>
    <el-table-column
      prop="ContractAddress"
      label="contract address">
    </el-table-column>
    <el-table-column
      prop="OwnerAddress"
      label="owner address">
    </el-table-column>
    <el-table-column
      label="invoke">
      <template slot-scope="scope">
        <a :href="genContractUrl(scope.row.ID)">invoke</a>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import Vue from 'vue'
import ElementUI from 'element-ui'
import ifetch from 'jsappbase/ifetch'
import _ from "lodash"

export default {
  data() {
    return {
      tableData: [],
    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    getData() {
      ifetch("/api/v1/tokens", "GET")
      .then((response) => {
        this.tableData = response.data
      })
    },
    genContractUrl(id){
      console.log(id)
      return `/token_dapp/${id}`
    }
  }
}
</script>

<style>
</style>
