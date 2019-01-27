<template>
	<div>
		<!-- Nav tabs -->
		<ul class="nav nav-tabs tab-nav-right">
		    <li
		    v-for="tab in config"
		    :key="tab.id"		    
		    :class="{ 'active': tab.active == true }">
		    	<a class="tabtitle" :href="tab.ref" v-on:click="changeTab($event, tab)">{{tab.title}}</a>
			</li>
		</ul>

		<!-- Tab panes -->
		<div class="tab-content">
		    <div v-for="tab in config"
		    :key="tab.id"
		    :class="{ 'in active': tab.active == true }"
		    class="tab-pane fade" :id="tab.ref">
		        <slot :name="tab.ref">
		        	<p style="text-align: center;">this tab has no content</p>
		        </slot>
		    </div>
		</div>
	</div>
</template>

<script>

export default {
  name: 'tabs',
  props: {
  	config: {
  		type: Array,
  		default: () =>  []
  	},
  	defaultTabId: {
  		type: Number,
  		default: 0
  	}
  },
  data () {
    return {
    	lastTabId: this.defaultTabId
    }
  },
  watch: {
  	lastTabId: function(current, previous){
  		this.config[current].active = true;
  		this.config[previous].active = false;
  	}
  },
  methods: {
  	changeTab: function (e, tab) {
  		e.preventDefault();
  		this.lastTabId = tab.id;
  	}
  },
  created(){
    log("tabs::created");
    //set the last clicked tab, the tab set as active
    if (!this.config) {
    	error('no tabs configuration provided for tabs components');
    }
  },
  mounted(){
    log("tabs::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.tabtitle {
	text-transform: uppercase;
	cursor: pointer !important;
}
</style>