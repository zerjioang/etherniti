<template>
	<div v-show="form!=undefined">
		<div :class="form.class">
	    <div v-if="form.type == 'card'">
        <card :form="form" :model="model" v-on:jsonevent="emit">
          <div slot="card-top">
            <slot name="form-top"></slot>
          </div>    
        </card>
      </div>
      <div v-else>
        <p>Unsupported form type requested</p>
      </div>
		</div>
	</div>
</template>

<script>

export default {
  name: 'json-form',
  props: {
  	form: {
  		type: Object,
  		default: () => {}
  	},
    model: {
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