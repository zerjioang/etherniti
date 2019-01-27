<template>
  <body class="theme-gaethway">
    <!-- Page Loader -->
    <loader text="Please wait..." :visible="loaderVisible"/>
    <!-- Overlay For Sidebars -->
    <overlay :visible="searchBarVisible||isRightAsideVisible"/>

    <!--
    <div class="banner">test</div>
    -->

    <!-- Search Bar -->
    <searchbar
    :visible="searchBarVisible"
    v-on:open="searchBarVisible=true"
    v-on:closed="searchBarVisible=false"
    v-on:keyup.esc="searchBarVisible=false">
    </searchbar>

    <!-- #END# Search Bar -->
    <!-- Top Bar -->
    <navigationBar
        :title="navigationBarTitle"
        :showSearchBar="searchBarVisible"
        v-on:openSearch="searchBarVisible=true"
        v-on:toggleRightAside="toggleRightAside"
        v-on:toggleLeftAside="leftAsideVisible=!leftAsideVisible">
    </navigationBar>
    
    <section>
        <leftAside
        :visible="isRightAsideVisible"
        v-click-outside="onAsideClickOutside"
        v-on:open="rightAsideRequested=false"
        v-on:close=""
        v-on:keyup.esc="isRightAsideVisible=false"
        :footerLayout="layout.footer">
        </leftAside>

        <!-- Right Sidebar -->
        <rightAside
        :visible="isRightAsideVisible"
        v-click-outside="onAsideClickOutside"
        v-on:open="rightAsideRequested=false"
        v-on:close=""
        v-on:keyup.esc="isRightAsideVisible=false">
        </rightAside>
    </section>

    <section class="content">
        <div class="container-fluid">
            <div class="block-header">
                <h2>{{pagetitle}}</h2>
            </div>
            <transition
            name="fade"
            mode="out-in"
            @beforeLeave="beforeLeave"
            @enter="enter"
            @afterEnter="afterEnter">
              <router-view/>  
            </transition>
        </div>
    </section>
  </body>
</template>

<script>

import material from '@/mixins/material';
import footerLayout from '@/layout/footer_config';

export default {
  name: 'base-dashboard-view',
  mixins: [
    material
  ],
  props: {
    pagetitle: {
        type: String,
        default: "GAETHWAY DASHBOARD"
    }
  },
  directives: {
    'click-outside': {
        bind: function(el, binding, vNode) {
          // Provided expression must evaluate to a function.
          if (typeof binding.value !== 'function') {
            const compName = vNode.context.name
            let warn = `[Vue-click-outside:] provided expression '${binding.expression}' is not a function, but has to be`
            if (compName) { warn += `Found in component '${compName}'` }
            
            console.warn(warn)
          }
          // Define Handler and cache it on the element
          const bubble = binding.modifiers.bubble
          const handler = (e) => {
            if (bubble || (!el.contains(e.target) && el !== e.target)) {
              binding.value(e)
            }
          }
          el.__vueClickOutside__ = handler;

          // add Event Listeners
          document.addEventListener('click', handler)
        },
        unbind: function(el, binding) {
          // Remove Event Listeners
          document.removeEventListener('click', el.__vueClickOutside__)
          el.__vueClickOutside__ = null
        }
    }
  },
  watch: {
    '$route.meta' () {
      // update breadcrum
      // this.updateView();
    }
  },
  data () {
    return {
      //to allow router view animations
      prevHeight: 0,
      title: process.env.APP_NAME,
      navigationBarTitle: process.env.APP_TITLE,
      sidebarImage: require("@/assets/images/aside.png"),
      loaderVisible: false,
      searchBarVisible: false,
      isRightAsideVisible: false,
      leftAsideVisible: false,
      rightAsideRequested: false,
      layout: {
        footer: footerLayout
      }
    }
  },
  methods: {
    //to allow router view animations
    beforeLeave(element) {
      this.prevHeight = getComputedStyle(element).height;
    },
    enter(element) {
      const { height } = getComputedStyle(element);
      element.style.height = this.prevHeight;
      setTimeout(() => {
        element.style.height = height;
      });
    },
    afterEnter(element) {
      element.style.height = 'auto';
    },
    test: function(){
        this.searchBarVisible = false;
    },
    toggleRightAside: function() {
        log("toogling right aside");
        this.rightAsideRequested = true;
        this.isRightAsideVisible = !this.isRightAsideVisible;
    },
    //this function is called on every item click event 
    // happened on the screen
    onAsideClickOutside: function(e) {
      log("clicked outside right aside");
      // force right aside close only when is openene
      if(this.rightAsideRequested === false && this.isRightAsideVisible === true){
        log("closing right aside");
        this.isRightAsideVisible = false;
      }
    }
  },
  created(){
    log("base-view::created");
  },
  mounted(){
    log("base-view::mounted");
    document.title = process.env.APP_TITLE;
    this.materialLoad();
  },
  components: {
    loader: () => import('@/components/loader'),
    overlay: () => import('@/components/overlay'),
    searchbar: () => import('@/components/searchbar'),
    navigationBar: () => import('@/components/navigationBar'),
    rightAside: () => import('@/components/rightAside'),
    leftAside: () => import('@/components/leftAside')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style type="text/css" scoped="true">
.fade-enter-active,
.fade-leave-active {
  transition-duration: 0.3s;
  transition-property: opacity;
  transition-timing-function: ease;
}

.fade-enter,
.fade-leave-active {
  opacity: 0
}

.banner {
  align-items: center;
  background-color: #000;
  color: #fff;
  font-size: 14pt;
  height: 2.5rem;
  justify-content: center;
  position: fixed;
  top: 0;
  width: 100%;
  padding: 10px;
}
</style>