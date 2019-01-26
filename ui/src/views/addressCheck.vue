<template>
    <div>
        <h2>Ethereum Address Check</h2>
        <p>Verify whether a given adress is valid or not</p>

        <div class="row clearfix">
            <div class="col-lg-4 col-md-4 col-sm-4 col-xs-4">
                <div class="card">
                  <div class="header header-slim">
                      <h2 class="title">Ethereum Address validation
                        <small class="subtitle">Verify if given address is valid or not</small>
                      </h2>
                      <ul class="header-dropdown m-r--5">
                          <li class="dropdown">
                              <a href="javascript:void(0);" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                                  <i class="material-icons">more_vert</i>
                              </a>
                              <ul class="dropdown-menu pull-right">
                                  <li><a href="javascript:void(0);">What's this?</a></li>
                                  <li><a href="javascript:void(0);">Show internals</a></li>
                                  <li><a href="javascript:void(0);">View webAPI request</a></li>
                              </ul>
                          </li>
                      </ul>
                  </div> <!--header end -->
                  <div class="body slim">
                      <p v-show="false" class="card-inside-title">Enter address to validate</p>
                      <div class="row clearfix">
                        <form method="GET" v-on:submit="submit">
                           <div class="col-md-12">
                              <b>ETH Address</b>
                              <div class="input-group">
                                  <span class="input-group-addon">
                                      <i class="material-icons">account_box</i>
                                  </span>
                                  <div class="form-line">
                                      <input
                                      type="text"
                                      class="form-control key"
                                      placeholder="0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae"
                                      :required="true"
                                      :disabled="false">
                                  </div>
                              </div> <!-- form group end -->
                              <div class="form-btn">
                                <button type="submit" class="btn btn-lg bg-indigo m-t-15 waves-effect upper">verify address</button>
                              </div>
                          </div> <!-- col end -->
                        </form>
                      </div>
                  </div> <!--body end -->
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
  name: 'eth-addr-check-view',
  data () {
    return {
      title: process.env.UI_TITLE,
    }
  },
  methods: {
    submit: function () {
      alert("test")
    }
  },
  created(){
    log("eth-addr-check-view::created");
  },
  mounted(){
    log("eth-addr-check-view::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.centered {
  text-align: center;
}

.title {
  color: #012282 !important;
  font-weight: bold !important;
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