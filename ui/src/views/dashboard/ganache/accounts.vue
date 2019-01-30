<template>
    <div>
      <pagetitle/>
      <div class="row">
        <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
          <div class="card">
              <div class="header">
                  <h2>
                      Ganache accounts
                      <small>Add the badges component to any list group item and it will automatically be positioned on the right.</small>
                  </h2>
              </div>
              <div class="body">
                  <ul class="list-group">
                      <li
                        v-for="acc in accounts"
                        class="list-group-item"
                        v-if="acc.address && acc.key && acc.balance >= 0" >
                          <div class="row">
                            <div class="col-md-5 fit">
                              <b>Address: </b>
                              <tt>{{acc.address}}</tt>
                            </div>

                            <div class="col-md-6 fit">
                              <b>Private key: </b>
                              <tt>{{acc.key}}</tt>
                            </div>
                            <div class="col-md-1 fit">
                              <span class="badge bg-blue">{{acc.balance}} ETH</span>
                            </div>
                          </div> <!--row -->
                      </li>
                  </ul>
              </div>
          </div>
      </div>
      </div>
    </div>
</template>

<script>

import ganacheAccounts from '@/layout/ganacheAccounts';
export default {
  name: 'ganache-accounts-view',
  data () {
    return {
      app: process.env.APP_NAME,
      accounts: ganacheAccounts
    }
  },
  methods: {
  },
  created(){
    log("ganache-accounts-view::created");
  },
  mounted(){
    log("ganache-accounts-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.fit {
  margin: 0px !important;
  padding-left: 10px !important;
  padding-right: 10px !important;
  padding-top: 2px !important;
  padding-bottom: 2px !important;
  font-size: 16px;
  overflow-y: auto;
  word-wrap: break-word;
}
</style>