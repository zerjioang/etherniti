<template>
</template>

<script type="text/javascript">

export default {
  name: 'localstorage-mixin',
  data () {
    return {
    }
  },
  methods: {
    //begin custom app helper methods
    shouldWelcomeWizardPopup: function () {
      log("checking welcome wizard status");
      return this.getStore(process.env.WELCOME_WIZARD_STATUS_KEY) != 'skip';
    },
    setWelcomeWizardNoMoreVisible(){
      log("setting welcome wizard as skipped");
      this.setStore(process.env.WELCOME_WIZARD_STATUS_KEY, 'skip');
    },

    // begin localstorage generic methods
    getBrowserName(){
      let ua= navigator.userAgent, tem,
      M= ua.match(/(opera|chrome|safari|firefox|msie|trident(?=\/))\/?\s*(\d+)/i) || [];
      if(/trident/i.test(M[1])){
          tem=  /\brv[ :]+(\d+)/g.exec(ua) || [];
          return 'IE '+(tem[1] || '');
      }
      if(M[1]=== 'Chrome'){
          tem= ua.match(/\b(OPR|Edge)\/(\d+)/);
          if(tem!= null) return tem.slice(1).join(' ').replace('OPR', 'Opera');
      }
      M= M[2]? [M[1], M[2]]: [navigator.appName, navigator.appVersion, '-?'];
      if((tem= ua.match(/version\/(\d+)/i))!= null) M.splice(1, 1, tem[1]);
      return M.join(' ');
    },
    showLocalStorageNotSupportedMessage: function(){
      log("local storage is not supported in your browser");
      this.$router.push({name: 'localStorageError'});
    },
    supportsLocalStorage: function(){
      /*
      // example code for logic
      if (typeof(Storage) !== "undefined") {
        // Code for localStorage/sessionStorage.
        return true;
      } else {
        // Sorry! No Web Storage support..
        return false;
      }
      */
      return typeof(Storage) !== "undefined";
    },
    //wrapped localstorage and session storage functions
    // so that they can be easily changed for other key-value storage methods
    getSession(key){
      return sessionStorage.getItem(key);
    },
    setSession(key, value){
      return sessionStorage.setItem(key, value);
    },
    removeSession(key){
      sessionStorage.removeItem(key);
    },
    getStore(key){
      return localStorage.getItem(key);
    },
    setStore(key, value){
      return localStorage.setItem(key, value);
    },
    removeStore(key){
      localStorage.removeItem(key);
    },
  },
  created(){
    log("localstorage-mixin::created");
  },
  mounted(){
    log("localstorage-mixin::mounted");
  },
  components: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>