<template>
	<div class="row clearfix">
	    <div v-for="card in config.cards" :key="card.id" :class="card.class">
	        <div class="info-box hover-expand-effect" :class="card.color">
	            <div class="icon">
	                <i class="material-icons">{{card.icon}}</i>
	            </div>
	            <div class="content">
	                <div class="text upper">{{card.title}}</div>
	                <div class="number">{{card.value}}</div>
	            </div>
	        </div>
	    </div>
	</div>
</template>

<script>

export default {
  name: 'widgets-row',
  props: {
    config: {
      type: Object,
      default: () => {}
    },
  },
  mixins: [
  ],
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("widget-row::created");
  },
  mounted(){
    log("widget-row::mounted");
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.upper {
	text-transform: uppercase;
}
</style>
