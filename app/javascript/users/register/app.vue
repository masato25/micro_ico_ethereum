<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix">
      <span>Register</span>
      <i class="el-icon-edit"></i>
    </div>
    <div>
      <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="160px">
        <el-form-item label="username">
          <el-input v-model="ruleForm2.username"></el-input>
        </el-form-item>
        <el-form-item label="password" prop="password">
          <el-input type="password" v-model="ruleForm2.password" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="confirm password" prop="checkPass">
          <el-input type="password" v-model="ruleForm2.checkPass" auto-complete="off"></el-input>
        </el-form-item>
        <el-form-item label="ethereum address">
          <el-input v-model="ruleForm2.eth_address"></el-input>
        </el-form-item>
        <el-form-item label="private key">
          <el-input v-model="ruleForm2.eth_private_key"></el-input>
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
import ifetch from 'jsappbase/ifetch'
import _ from "lodash"

export default {
   data() {
     var validatePass = (rule, value, callback) => {
       if (value === '') {
         callback(new Error('please input password'));
       } else {
         if (this.ruleForm2.checkPass !== '') {
           this.$refs.ruleForm2.validateField('checkPass');
         }
         callback();
       }
     };
     var validatePass2 = (rule, value, callback) => {
       if (value === '') {
         callback(new Error('please input password again'));
       } else if (value !== this.ruleForm2.password) {
         callback(new Error('password not matched!'));
       } else {
         callback();
       }
     };
     return {
       ruleForm2: {
         username: "",
         password: '',
         checkPass: '',
         eth_address: "",
         eth_private_key: "",
       },
       rules2: {
         pass: [
           { validator: validatePass, trigger: 'blur' }
         ],
         checkPass: [
           { validator: validatePass2, trigger: 'blur' }
         ]
       }
     };
   },
   methods: {
     postRegister() {
       ifetch("/api/v1/users/resigter", "JSONPOST", this.ruleForm2)
       .then((response) => {
         this.$message({
           message: "Account created",
           type: 'success'
         });
       })
       .catch((response) => {
         this.$message.error('Account registration failed!');
       })
     },
     submitForm(formName) {
       this.$refs[formName].validate((valid) => {
         if (valid) {
           this.postRegister()
         } else {
           console.log('error submit!!')
           return false
         }
       });
     },
     resetForm(formName) {
       this.$refs[formName].resetFields()
     }
   }
 }
</script>

<style>
</style>
