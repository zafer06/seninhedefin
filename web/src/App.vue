<template>
  <div id="app">
    <appList/>
  </div>
</template>

<script>
import Targetlist from './components/Targetlist'
export default {
  components:{
    "appList" : Targetlist
  }
}



</script>

<style lang="scss">
  @import './assets/scss/targetlist.scss';
</style>
