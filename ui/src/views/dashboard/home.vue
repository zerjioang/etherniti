<template>
    <div>
      <pagetitle
        :title="title"/>
        <widgetsRow :config="layout.widgetRowConfig"/>
    </div>
</template>

<script>

import widgetRowConfig from '@/layout/widgetRowConfig';
import localstorage from '@/mixins/localstorage';
import wizardSlider from '@/components/wizardSlider';

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
        widgetRowConfig: widgetRowConfig
      }
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
</style>