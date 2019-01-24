<template>
  <div class="search-bar" :class="{ 'open': visible }">
        <div class="search-icon">
            <i class="material-icons">search</i>
        </div>
        <input
        type="text"
        ref="search"
        placeholder="START TYPING..."
        v-on:keyup.enter="querySearch"
        v-on:keyup.esc="hideSearchBar">
        <div class="close-search"
          v-on:click="hideSearchBar">
            <i class="material-icons">close</i>
        </div>
    </div>
</template>

<script>
export default {
  name: 'search-bar',
  props: {
    visible: {
      type: Boolean,
      default: false
    }
  },
  data () {
    return {
    }
  },
  watch: {
    visible: function(now, before){
      if(this.visible==true){
        this.showSearchBar();
      } else {
        this.hideSearchBar();
      }
    }
  },
  methods: {
    //Search button click event
    querySearch: function(){
      alert("searching");
      this.hideSearchBar();
    },
    showSearchBar: function () {
        this.$refs.search.focus();
        this.$emit('open');
    },
    hideSearchBar: function () {
        this.$emit('closed');
    }
  },
  created(){
    log("search-bar::created");
  },
  mounted(){
    log("search-bar::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>