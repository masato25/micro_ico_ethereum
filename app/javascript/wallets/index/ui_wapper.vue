<template>
  <div>
    {{count}}
    <span v-for="(item, index) in items">
      <span>{{item.type}}</span>
      <el-input v-model="inputs[index]" placeholder=""></el-input>
    </span>
    <el-button type="success" @click="submit()">submit</el-button>
  </div>
</template>

<script>
import _ from "lodash"
import {store} from "../store/store"
export default {
  store,
  props: {
    items: Array,
  },
  data() {
    return {
      inputs: [],
    }
  },
  mounted() {
  },
  computed: {
   count () {
     return store.state.count
   }
  },
  methods: {
    submit(){
      console.log(this.inputs)
    }
  },
  watch:{
    items() {
      this.inputs = []
      store.commit('increment')
    }
  }
}
</script>

<style>
</style>
