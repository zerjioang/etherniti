<!--this file should not be modified -->
<template>
  <div class="row err">
    <div class="col-md12">

        <slot name="image"></slot>
        <h3 class="message">{{title}}</h3>
        <div class="details text-muted">{{description}}</div>
        
        <slot name="content"></slot>

        <div class="button-place top right" v-if="buttonVisible===true">
            <a class="btn btn-default btn-lg waves-effect big" v-on:click="notShowMore">Do not show again</a>
            <a class="btn btn-default btn-lg waves-effect big" v-on:click="close">Close</a>
        </div>
    </div>
  </div>
</template>

<script>

import localstorage from '@/mixins/localstorage';

export default {
  name: 'wizard-tab',
  props: {
    title: {
      type: String,
      default: ""
    },
    description: {
      type: String,
      default: ""
    },
    buttonVisible: {
      type: Boolean,
      default: false
    }
  },
  mixins: [
    localstorage
  ],
  data () {
    return {
    }
  },
  methods: {
    notShowMore: function() {
      this.setWelcomeWizardNoMoreVisible();
      this.close();
    },
    close: function () {
      this.$emit('close');
    }
  },
  created(){
    log("wizard-tab::created");
  },
  mounted(){
    log("wizard-tab::mounted");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.err {
    padding-top: 5%;
    padding-bottom: 5%;
    margin: 0px;
    text-align: center;
}
.blue {
  color: #383A3F;
}
.title {
  font-size: 60pt;
}
.message {
  font-size: 25pt;
}
.details {
  font-size: 20pt;
}
.big {
  font-size: 16pt !important;
}
.top {
  padding-top: 30px;
}
.right {
  text-align: right;
}
body {
  background-color: #1d1a1a;
}
a {
  margin: 5px;
}
</style>