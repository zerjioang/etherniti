<template>
</template>

<script type="text/javascript">

import defaultDialog from '@/components/dialog/generic';

export default {
  name: 'dialog-mixin',
  data () {
    return {
    }
  },
  methods: {
    // shows a popup message using vue-js-modal
    showPopup(title, message, icon, buttonText, onClose){
      this.$modal.show(defaultDialog,
        {
          title: title,
          message: message,
          icon: icon,
          buttonText: buttonText
        }, {
          draggable: false,
          height: "auto"
        },{
          'before-close': (event) => {
            log('modal closed');
            if (onClose) {
              onClose();
            }
          }
        });
    },
  },
  created(){
    log("dialog-mixin::created");
  },
  mounted(){
    log("dialog-mixin::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>