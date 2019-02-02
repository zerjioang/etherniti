<template>
    <div>
      <pagetitle
        :title="title"/>
        <widgetsRow :config="layout.widgetRowConfig"/>
        <p>
          For higher security, when executing transactions in Mainnets or sensitive environments, please verify the integrity of your connection using the provided message and following public key in order to avoid MITM attacks, web hijacking attacks, DNS based attacks, etc. We provide you below our <b>Public integrity key</b> for manual checking if needed:
          <div v-if="showDetails">
            <br>
            <small class="code">
              Following commands were used to generate the keys:
              <br>
              openssl genpkey -algorithm RSA -out {{app}}.pem -pkeyopt rsa_keygen_bits:4096
              <br>
              openssl rsa -pubout -in {{app}}.pem -out public_key.pem
            </small>
          </div>
        </p>
        
        <div class="row">
          <div class="col-md-6">
            <h4>Public {{app}} integrity key</h4>
            <p class="publicKey">{{layout.publicKey}}</p>
          </div>
          <div class="col-md-6">
            <h4>Integrity message returned by {{app}} webAPI</h4>
            <p class="publicKey">{{layout.publicKey}}</p>
          </div>
        </div>

    </div>
</template>

<script>

import widgetRowConfig from '@/layout/widgetRowConfig';
import localstorage from '@/mixins/localstorage';
import wizardSlider from '@/components/wizardSlider';

import publicKey from '@/layout/integrity/publicKey'

export default {
  name: 'home-view',
  mixins: [
    localstorage
  ],
  data () {
    return {
      app: process.env.APP_NAME,
      title: "Current "+process.env.APP_NAME+" stats",
      layout: {
        widgetRowConfig: widgetRowConfig,
        publicKey: publicKey
      },
      showDetails: false
    }
  },
  methods: {
    showWelcomeWizard: function () {
      this.$modal.show(wizardSlider,
        {
          title: "About Gaethway Project",
          message: 'Gaethway Project is a Multitenant High Performance Ethereum and Quorum compatible WebAPI',
          icon: "code",
          buttonText: "Close"
        }, {
          draggable: false,
          scrollable: true,
          adaptative: true,
          height: "auto"
        },{
          'before-close': (event) => {
            log('modal closed');
          }
        }
       );
    }
  },
  created(){
    log("home-view::created");
  },
  mounted(){
    log("home-view::mounted");
    if (this.shouldWelcomeWizardPopup() ) {
      log('showing welcome wizard');
      this.showWelcomeWizard();
    } else {
      //skip welcome wizard as user request
      log('skip welcome wizard as user request')
    }
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    widgetsRow: () => import('@/components/widgetsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.publicKey {
  font-family: 'IBM Plex Mono';
  font-weight: bold;
  padding-bottom: 10px;
  overflow-y: auto;
  word-break: break-word;
  background-color: #424242;
  color: #eeeeee;
  padding: 10px;
  border-style: solid;
  border-width: 5px;
  border: 2px solid #3747a0;
  border-radius: 6px;
  white-space: pre-line;
}

.code {
  font-family: 'IBM Plex Mono';
  font-weight: bold;
  overflow-y: auto;
  color: #424242;
}
</style>