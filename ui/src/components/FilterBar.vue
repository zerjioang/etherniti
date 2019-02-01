  // FilterBar.vue

  <template>
    <div class="row clearfix">
        <div class="col-lg-3 col-md-3 col-sm-3 col-xs-6">
            <div class="form-group form-float">
                <div class="form-line">
                    <input type="text" class="form-control" v-model="filterText" @keyup.enter="doFilter" placeholder="Search for...">
                </div>
            </div>
        </div>
        <div class="col-lg-6 col-md-6 col-sm-6 col-xs-12">
            <button type="button" class="btn btn-primary btn-lg m-l-15 waves-effect" @click="doFilter">Search</button>
            <button type="button" class="btn btn-primary btn-lg m-l-15 waves-effect" @click="resetFilter">Reset</button>
        </div>
    </div>
  </template>

  <script>
  export default {
    data () {
      return {
        filterText: ''
      }
    },
    methods: {
      doFilter () {
        console.log('doFilter:', this.filterText)
      },
      resetFilter () {
        this.filterText = ''
        console.log('resetFilter')
      }
    }
  }
</script>