<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <span>Login</span>
    </div>
    <div>
      <el-form ref="ruleForm2" :model="ruleForm2" status-icon label-width="80px">
        <el-form-item label="username">
          <el-input v-model="ruleForm2.username" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input v-model="ruleForm2.password" type="password" auto-complete="off">
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('ruleForm2')">Submmit</el-button>
          <el-button @click="resetForm('ruleForm2')">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
import ifetch from "jsappbase/ifetch"
import _ from "lodash"
import Cookies from 'js-cookie';

export default {
  data() {
    return {
      ruleForm2: {
        username: "",
        password: "",
      }
    }
  },
  mounted() {

  },
  methods: {
    postLogin() {
      ifetch("/api/v1/users/login", "JSONPOST", this.ruleForm2)
      .then((response) => {
        let sessionobj = response.data
        Cookies.set('auth_session', `{"username": "${sessionobj.username}", "sesssion": "${sessionobj.session}"}`, { expires: 1 })
        this.$message({
          message: "success",
          type: 'success'
        })
        setTimeout(() => {
          window.location.href = "/tokens"
        }, 1000)
      })
      .catch((response) => {
        this.$message.error('account or password not correct!');
      })
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.postLogin()
        } else {
          console.log('error submit!!')
          return false
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
  }
}
</script>

<style>
</style>
