<template>
	<div class="card">

	  <!--print card header if configured -->
	  <div v-if="form.header" class="header header-slim">
	      <h2 class="title">{{form.header.title}}
	        <small class="subtitle">{{form.header.subtitle}}</small>
	      </h2>
	      <ul class="header-dropdown m-r--5">
	          <li class="dropdown">
	              <a href="javascript:void(0);" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
	                  <i class="material-icons">{{form.header.dropdown.icon}}</i>
	              </a>
	              <ul class="dropdown-menu pull-right">
	                  <li v-for="dropitem in form.header.dropdown.items" :key="dropitem.id">
	                  	<a href="javascript:void(0);">{{dropitem.title}}</a></li>
	              </ul>
	          </li>
	      </ul>
	  </div> <!--header end -->

	  <slot name="card-top"></slot>

	  <div class="body slim">
	      <p v-if="form.body.titleInside" class="card-inside-title">{{form.body.titleInside}}</p>
	      <div :class="form.body.rowClass">
	        <form :method="form.body.method" v-on:submit="trigger($event)" id="submit">
	           <div
	           v-for="col in form.body.columns"
	           :key="col.id"
	           :class="col.class">
	           	  
	           	  <div v-for="inputGroup in col.inputgroup"
	           	  :key="inputGroup.id">
	              	<b>{{inputGroup.title}}</b>
	              	<jsonInput
	              	v-for="groupItem in inputGroup.items"
	              	:key="groupItem.id"
	              	:config="groupItem"
	              	:model="model"/>
	           	  </div> <!-- form group end -->
	           	  
	           	  <div
	           	  	v-for="btn in col.buttons"
	           	  	:key="btn.id"
	           	  	class="form-btn">
	           	  		<button
	           	  		type="btn.type"
	           	  		:class="btn.class"
	           	  		v-on:click="trigger($event)"
	           	  		:id="btn.eventid">{{btn.text}}
	           	  		</button> 
	              </div> <!-- buttons required in current col end -->
	          </div> <!-- col end -->
	        </form>
	      </div>
	  </div> <!--body end -->
	</div>
</template>

<script>

export default {
  name: 'card',
  props: {
  	form: {
  		type: Object,
  		default: () => {}
  	},
  	model: {
  		type: Object,
  		default: () => {}
  	},
  },
  data () {
    return {
    }
  },
  methods: {
    trigger: function (event) {
      const eventId = event.target.id;
      if (eventId!="") {
      	log("trigger event: "+eventId);
      	this.$emit('jsonevent', event, eventId) //submit, reset, etc
      }
    }
  },
  created(){
    log("card::created");
  },
  mounted(){
    log("card::mounted");
  },
  components: {
  	jsonInput: () => import('@/components/jsonInput')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">

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