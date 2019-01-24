<template>
  <body class="theme-methw">
    <!-- Page Loader -->
    <loader text="Please wait..." :visible="loaderVisible"/>
    <!-- Overlay For Sidebars -->
    <overlay :visible="searchBarVisible||isRightAsideVisible"/>
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
        title="METHW: A High Performance Ethereum WebAPI"
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
            <slot name="content">
                <p>This page has no content</p>
            </slot>
        </div>
    </section>

</body>
</template>

<script>

import material from '@/mixins/material';

export default {
  name: 'base-dashboard-view',
  mixins: [
    material
  ],
  props: {
    pagetitle: {
        type: String,
        default: "MethW DASHBOARD"
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
  data () {
    return {
      title: process.env.UI_TITLE,
      sidebarImage: require("@/assets/images/aside.png"),
      loaderVisible: false,
      searchBarVisible: false,
      isRightAsideVisible: false,
      leftAsideVisible: false,
      rightAsideRequested: false,
      layout: {
        footer: {
            version: "0.0.1",
            copyright: "METHW Project",
            years: "2018 - 2019"
        }
      }
    }
  },
  methods: {
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
    log("index-view::created");
  },
  mounted(){
    log("index-view::mounted");
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
<style scoped>
</style>