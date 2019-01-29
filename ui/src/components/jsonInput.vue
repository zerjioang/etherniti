<template>
<div
	v-if="config.type == 'icon+text'"
	:class="config.class">
	<span class="input-group-addon">
	  <i class="material-icons">{{config.icon}}</i>
	</span>
  <div
  class="form-line"
  :class="{'focused': focus}">
	  <input
    @focus="focus=true"
    @blur="focus=false"
	  :type="config.input.type"
	  :class="config.input.class"
	  :placeholder="config.input.placeholder"
    :required="config.input.required"
	  :autocomplete="config.input.autocomplete"
	  :disabled="config.input.disabled"
	  v-model="model[config.input.modelKey]">
	</div>
</div>
<div
  v-else-if="config.type == 'textarea'"
  :class="config.class">
  <span class="input-group-addon">
    <i class="material-icons">{{config.icon}}</i>
  </span>
  <div
  class="form-line"
  :class="{'focused': focus}">
    <textarea
    @focus="focus=true"
    @blur="focus=false"
    :type="config.input.type"
    :class="config.input.class"
    :placeholder="config.input.placeholder"
    :required="config.input.required"
    :disabled="config.input.disabled"
    :cols="config.input.cols"
    :rows="config.input.rows"
    v-model="model[config.input.modelKey]">
    </textarea>
  </div>
</div>
<div v-else>
	<p>Not supported json input</p>
</div>
</template>

<script>

export default {
  name: 'json-input',
  props: {
  	config: {
  		type: Object,
  		default: () => {}
  	},
  	model: {
  		type: Object,
  		default: () => {}
  	}
  },
  watch: {
    focus: function (to, from) {
      this.focus = to;
    }
  },
  data () {
    return {
      focus: false
    }
  },
  methods: {
  },
  created(){
    log("json-input::created");
  },
  mounted(){
    log("json-input::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">

.centered {
  text-align: center;
}

.title {
  color: #012282 !important;
  font-weight: bold !important;
}
.focus {
  color: #012282 !important;
  font-weight: bold !important;
  font-size: 25pt;
}
.subtitle {
  font-size: 12px !important;
}

.form-btn {
  text-align: right;
}

.slim {
  padding-left: 15px;
  padding-right: 15px;
  padding-bottom: 0px;
}

.header-slim {
  padding-left: 15px;
  padding-right: 15px;
  padding-bottom: 5px;
}

.upper {
  text-transform: uppercase;
}
</style>