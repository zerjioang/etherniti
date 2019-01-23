<template>
</template>

<script>
export default {
	name: "loader-mixin",
	methods: {
	},
	created(){
		log("loader-mixin::created");
		setTimeout(function () { $('.page-loader-wrapper').fadeOut(); }, 50);
	},
	mounted(){
		log("loader-mixin::mounted");
	},
	components: {
	}
}
</script>