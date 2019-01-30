<template>
	<div>
		<h4>{{app}} is one of the best decentralized applications gateway</h4>
		<h5>Direct communication to Ethereum and Quorum based Distributed Ledgers or Blockchain networks.</h5>
		<img class="img-responsive thumbnail tech" :src="image">
		<p>
		    YOUR <b>PRIVATE</b> ACCESS TO THE ETHEREUM NETWORK.
		</p>
		<p>
		    Our easy to use WebAPI and UI provide secure, reliable, and private access to Ethereum and Quorum. We provide the infrastructure for your decentralized applications interaction so you can focus on the awesome features.
		</p>

		<h4>{{app}} Compatibility</h4>
		<p>{{app}} REST API implementation allows multiple services compatibility. This is because we follow ethereum original communication standards and protocols. In such a way, you can also use {{app}}  with following services:</p>
		<div class="row">
			<div class="col-md-4 center" v-for="item in compatibility" :key="item.title">
				<img class="img-responsive thumbnail compatibilityImage" :src="item.image" :alt="item.title">
				<p class="center"><b>{{item.title}}</b></p>
			</div>
		</div>
	</div>
</template>

<script>
export default {
  name: 'about-tab',
  data () {
    return {
    	app: process.env.APP_NAME,
  		image: require('@/assets/images/eth-utils.jpg'),
  		compatibility: [
  			{
	  			title: "Infura",
	  			image: require('@/assets/images/infura.png'),
	  			url: "https://infura.io"
	  		},
	  		{
	  			title: "QuikNode",
	  			image: require('@/assets/images/quiknode.png'),
	  			url: "https://quiknode.io/",
	  		},
	  		{
	  			title: "Alastria",
	  			image: require('@/assets/images/alastria.png'),
	  			url: "https://alastria.io"
	  		},
	  		
  		]
    }
  },
  methods: {
  },
  created(){
    log("about-tab::created");
  },
  mounted(){
    log("about-tab::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.tech {
    padding: 10px;
    background-color: #f7f7f7;
}
.compatibilityImage {
	border: transparent;
	width: 20%;
	max-width: 100px;
}
.center {
	text-align: -webkit-center;
}
h4, h5, p {
	text-align: justify;
}
h5 {
	font-size: 16px;
}
</style>