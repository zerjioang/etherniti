<template>
	<section v-show="config!=undefined" class="container">
		<div v-for="form in config.forms"
			:key="form.id"
			:class="form.class">
			    <div v-for="col in form.columns"
			    :key="col.id"
			    :class="col.class">
			        <div v-if="form.type == 'card'">
			        	<card
                :config="form"
                v-on:jsonevent="emit"/>
			        </div>
			        <div v-else>
			        	<p>Unsupported form type requested</p>
			        </div>
			    </div>
			</div>
	</section>
</template>

<script>

export default {
  name: 'json-form',
  props: {
  	config: {
  		type: Object,
  		default: () => {}
  	}
  },
  data () {
    return {
    }
  },
  methods: {
    emit(event, eventId){
      this.$emit('jsonevent', event, eventId);
    }
  },
  created(){
    log("jsonForm::created");
  },
  mounted(){
    log("jsonForm::mounted");
  },
  components: {
  	card: () => import('@/components/card')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.container {
	padding: 0;
	margin: 0;
}
</style>