<template>
  <hooper>
    <slide>
      <div class="bg">
        <wizardTab
          title="Welcome To Gaethway"
          description="A Multitenant High Performance Ethereum WebAPI.">
            <img slot="image" :src="require('@/assets/images/ethereum.png')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Web3 Compatible REST API"
          description="Start using in a few clicks our high performance Ethereum compatible WebAPI.">
            <img slot="image" :src="require('@/assets/images/tabs/api.svg')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Connect to ANY network"
          description="Compatible to any Ethereum and Quorum networks. Ready for Ropsten, Homestead, Rinkeby, etc.">
            <img slot="image" :src="require('@/assets/images/tabs/network.svg')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Built-in wallet support"
          description="Directly connect IIoT devices using our Highly secured built-in wallet implementation">
            <img slot="image" :src="require('@/assets/images/tabs/wallet.svg')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Security by design"
          description="We do not store your credentials anywhere, we work just as proxy (gateway) for Ethereum networks">
            <img slot="image" :src="require('@/assets/images/tabs/insurance.svg')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Compatible with private Ethereum"
          description="users can interact with their private Ethereum implementations in few seconds.">
            <img slot="image" :src="require('@/assets/images/tabs/insurance.svg')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Ganache ready"
          description="Working with Ganache? We are also compatible with it.">
            <img slot="image" :src="require('@/assets/images/ganache.png')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          title="Infura ready"
          description="Working with Infura? We are also compatible with it.">
            <img slot="image" :src="require('@/assets/images/infura.png')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <slide>
      <div class="bg">
        <wizardTab
          v-on:close="$emit('close')"
          v-bind:buttonVisible=true
          title="Alastria consortium ready"
          description="Partners of Alastria national Blockchain ecosystem are also welcome.">
            <img slot="image" :src="require('@/assets/images/alastria.png')" class="mini">
        </wizardTab>
      </div>
    </slide>
    <hooper-navigation slot="hooper-addons"></hooper-navigation>
  </hooper>
</template>

<script>

  //https://baianat.github.io/hooper/api.html#props

import {
  Hooper,
  Slide,
  //Pagination as HooperPagination 
  Navigation as HooperNavigation,
  Progress as HooperProgress
  } from 'hooper';

require('~/hooper/dist/hooper.css');

export default {
  name: 'hooper-wizard-slider',
  data() {
    var data = {
    }
    return data;
  },
  methods: {
  },
  created(){

  },
  mounted(){

  },
  components: {
    Hooper,
    Slide,
    HooperNavigation,
    HooperProgress,
    //HooperPagination,
    wizardTab: () => import('@/components/wizardTab')
  }
}
</script>

<style type="text/css" scoped="true">
.bg {
  background-color: white;
  height: 100%;
  width: auto;
  padding: 50px; 
  list-style-type: none;
}
</style>

<style type="text/css">
.hooper-slide {
    flex-shrink: 0;
    height: auto !important;
}
/*remove black dot from hooper slider*/
.hooper-track{
  list-style-type: none;
}
.mini {
  max-width: 20%;
  width: 150px;
}
</style>