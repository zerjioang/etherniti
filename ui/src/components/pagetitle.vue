<template>
	<div>
		<h2 v-if="title">{{title}}</h2>

    <p v-if="subtitle">{{subtitle}}</p>

    <div v-if="subtitleHtml">
      <p v-html="subtitleHtml"></p>
    </div>
    
	</div>
</template>

<script>

export default {
  name: 'pagetitle',
  props: {
  	title: {
  		type: String,
  		default: ""
  	},
    subtitleHtml: {
      type: String,
      default: ""
    },
  	subtitle: {
  		type: String,
  		default: ""
  	}
  },
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("pagetitle::created");
  },
  mounted(){
    log("pagetitle::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.focus {
  color: #012282 !important;
  font-weight: bold !important;
  font-size: 25pt;
}
</style>