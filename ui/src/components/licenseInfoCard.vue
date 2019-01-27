<template>
    <div class="card">
        <div class="header">
            <div class="centered">
                <img class="logo" :src="license.image">
            </div>
            <h2>
                <b>{{license.project}}</b>
                <small class="justify">{{license.shortDescription}}</small>
            </h2>
        </div>
        <div class="body license justify">
            {{license.full}}
        </div>
    </div>
</template>

<script>

export default {
  name: 'license-info-card',
  props: {
    license: {
        type: Object,
        default: () => {}
    }
  },
  data () {
    return {
    }
  },
  methods: {
  },
  created(){
    log("license-info-card::created");
  },
  mounted(){
    log("license-info-card::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.license {
    font-size: 10pt;
    font-family: 'IBM Plex Mono', sans-serif;
}
.logo {
    max-width: 100px;
    width: 50%;
}
.centered {
    text-align: center;
}
.justify {
  text-align: justify;
}
</style>