<template>
	<div class="panel panel-default panel-post">
	  <div class="panel-heading">
	      <div class="media">
	          <div class="media-left">
	              <a href="#">
	                  <img src="https://avatars3.githubusercontent.com/u/6706342?s=460&v=4">
	              </a>
	          </div>
	          <div class="media-body">
	              <h4 class="media-heading">
	                  <a href="#">@zerjioang</a>
	              </h4>
	              <q>
	              	Decentralized Ethereum & Quorum Blockchain Infrastructure | Serving 1,000+ Developers & Decentralized Apps and growing
	              </q>
	          </div>
	      </div>
	  </div>
	  <div class="panel-body">
	      <div class="post">
	          <div class="post-heading">
	              <p>
	                <b>Made with love and tons of coffe</b>
	            	</br>
	            	</br>
	            	Interested in custom solutions, sales, or have general infrastructure questions? Please, do not hesitate to contact us at <a href="mailto:info@gethway.org" target="_blank" rel="noopener"><tt class="email">info@gethway.org</tt></a>. We are glad to help!
	              </p>
	          </div>
	      </div>
	  </div>
	</div>
</template>

<script>

export default {
  name: 'author-card',
  props: {
  },
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("author-card::created");
  },
  mounted(){
    log("author-card::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.email {
	color: #0530ADFF;
	font-weight: bold;
}
</style>