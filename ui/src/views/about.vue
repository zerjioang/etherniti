<template>
    <div>
        <h2>about gaethway</h2>
        <p>We provide secure, reliable, and high performance access to <b>Ethereum</b> and <b>Quorum</b> APIs.</p>

        <div class="row">
          <div class="col-md-12">
            <authorCard/>
          </div>
        </div>

        <div class="row">
            <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
                <div class="card">
                    <div class="header">
                        <h2>
                            About gaethway <b>WebAPI</b>
                            <small class="subtitle">General information about-view gaethway <b>WebAPI</b>, its service and this website</small>
                        </h2>
                        <ul class="header-dropdown m-r--5">
                            <li class="dropdown">
                                <a href="javascript:void(0);" class="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">
                                    <i class="material-icons">more_vert</i>
                                </a>
                                <ul class="dropdown-menu pull-right">
                                    <li><a href="javascript:void(0);" class=" waves-effect waves-block">Action</a></li>
                                    <li><a href="javascript:void(0);" class=" waves-effect waves-block">Another action</a></li>
                                    <li><a href="javascript:void(0);" class=" waves-effect waves-block">Something else here</a></li>
                                </ul>
                            </li>
                        </ul>
                    </div>

                    <div class="body">
                        <tabs :config="tabsConfig" :defaultTabId="tabsConfig[0].id">
                            <div :slot="tabsConfig[0].ref">
                                <aboutTab/>
                            </div>
                            <div :slot="tabsConfig[1].ref">
                                <termsTab/>
                            </div>
                            <div :slot="tabsConfig[2].ref">
                                <privacyTab/>
                            </div>
                            <div :slot="tabsConfig[3].ref">
                                <gdprTab/>
                            </div>
                        </tabs>
                    </div>
                </div>
            </div>
        </div> <!-- row finished -->

    </div>
</template>

<script>

export default {
  name: 'about-view',
  data () {
    return {
      tabsConfig: [
        {
            id: 0,
            title: "About us",
            icon: "",
            ref: "#about",
            active: true
        },
        {
            id: 1,
            title: "Terms & conditions",
            icon: "",
            ref: "#terms",
            active: false
        },
        {
            id: 2,
            title: "Privacy policy",
            icon: "",
            ref: "#privacy",
            active: false
        },
        {
            id: 3,
            title: "european gdpr",
            icon: "",
            ref: "#gdpr",
            active: false
        }
      ]
    }
  },
  methods: {
  },
  created(){
    log("about-view::created");
  },
  mounted(){
    log("about-view::mounted");
  },
  components: {
    tabs: () => import('@/components/tabs'),
    authorCard: () => import('@/components/authorCard'),
    aboutTab: () => import('@/components/tabs/about'),
    termsTab: () => import('@/components/tabs/terms'),
    privacyTab: () => import('@/components/tabs/privacy'),
    gdprTab: () => import('@/components/tabs/gdpr'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.centered {
  text-align: center;
}
.subtitle {
    font-size: 12pt !important;
}
</style>