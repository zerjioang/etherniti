<template>
    <div>
        <h2>about gaethway</h2>
        <p>We provide secure, reliable, and private access to Ethereum and Quorum APIs.</p>

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
                            About gaethway <b>webAPI</b>
                            <small>General information about-view gaethway <b>webAPI</b>, its service and this website</small>
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
                        <!-- Nav tabs -->
                        <ul class="nav nav-tabs tab-nav-right" role="tablist">
                            <li role="presentation" class="active"><a href="#home" data-toggle="tab" aria-expanded="true">ABOUT US</a></li>
                            <li role="presentation" class=""><a href="#profile" data-toggle="tab" aria-expanded="false">TERMS & CONDITIONS</a></li>
                            <li role="presentation" class=""><a href="#profile" data-toggle="tab" aria-expanded="false">PRIVACY POLICY</a></li>
                            <li role="presentation" class=""><a href="#profile" data-toggle="tab" aria-expanded="false">EUROPEAN GDPR</a></li>
                        </ul>

                        <!-- Tab panes -->
                        <div class="tab-content">
                            <div role="tabpanel" class="tab-pane fade" id="home">
                                <b>Home Content</b>
                                <p>
                                    Lorem ipsum dolor sit amet, ut duo atqui exerci dicunt, ius impedit mediocritatem an. Pri ut tation electram moderatius.
                                    Per te suavitate democritum. Duis nemore probatus ne quo, ad liber essent aliquid
                                    pro. Et eos nusquam accumsan, vide mentitum fabellas ne est, eu munere gubergren
                                    sadipscing mel.
                                </p>
                            </div>
                            <div role="tabpanel" class="tab-pane fade" id="profile">
                                <b>Profile Content</b>
                                <p>
                                    Lorem ipsum dolor sit amet, ut duo atqui exerci dicunt, ius impedit mediocritatem an. Pri ut tation electram moderatius.
                                    Per te suavitate democritum. Duis nemore probatus ne quo, ad liber essent aliquid
                                    pro. Et eos nusquam accumsan, vide mentitum fabellas ne est, eu munere gubergren
                                    sadipscing mel.
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>

export default {
  name: 'about-view',
  data () {
    return {
      title: process.env.UI_TITLE,
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
    authorCard: () => import('@/components/authorCard'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.centered {
  text-align: center;
}
</style>