<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>

require("@/assets/css/cleanscroll.css");

export default {
  name: 'app'
}
</script>

<style type="text/css" scoped="false">
/*set full width to all app
html, body, #app {
	height: 100%;
	width: 100%;
	min-width: 100%;
	max-width: 100%;
}
*/
</style>