<template>
	<!-- Left Sidebar -->
	<aside id="leftsidebar" class="sidebar">
	    <!-- User Info -->
	    <div class="user-info fit"
	     :style="{'background-image': 'url(' + sidebarImage + ')'}">
	        <div class="image">
	            <img src="@/assets/images/user.png" width="48" height="48" alt="User" />
	        </div>
	        <div class="info-container">
	            <div class="name" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">John Doe</div>
	            <div class="email">john.doe@example.com</div>
	            <div class="btn-group user-helper-dropdown">
	                <i class="material-icons" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">keyboard_arrow_down</i>
	                <ul class="dropdown-menu pull-right">
	                    <li><a href="javascript:void(0);"><i class="material-icons">person</i>Profile</a></li>
	                    <li role="separator" class="divider"></li>
	                    <li><a href="javascript:void(0);"><i class="material-icons">group</i>Followers</a></li>
	                    <li><a href="javascript:void(0);"><i class="material-icons">shopping_cart</i>Sales</a></li>
	                    <li><a href="javascript:void(0);"><i class="material-icons">favorite</i>Likes</a></li>
	                    <li role="separator" class="divider"></li>
	                    <li><a href="javascript:void(0);"><i class="material-icons">input</i>Sign Out</a></li>
	                </ul>
	            </div>
	        </div>
	    </div>
	    <!-- #User Info -->
	    <div class="menu">
	        <ul class="list">
	            <li class="header">{{navigationTitle}}</li>
	            <li v-for="rootMenuItem in menuLayout" :key="rootMenuItem.id" :class="{ 'active' : rootMenuItem.active==true }">
	                <router-link
	                	to=""
	                	@click.native="asideItemClick($event, rootMenuItem)"
	                	:class="{ 'menu-toggle' : rootMenuItem.submenus.length > 0 }">
	                    <i class="material-icons"
	                    :class="rootMenuItem.class">{{rootMenuItem.icon}}</i>
	                    <span>{{rootMenuItem.name}}</span>
	                </router-link>
	                <ul class="ml-menu">
	                	<li v-for="subItem in rootMenuItem.submenus">
  				            <router-link
                        to=""
                        @click.native="asideItemClick($event, subItem)"
                        :class="{ 'menu-toggle' : subItem.submenus.length > 0 }">
                          <i class="material-icons small"
                          :class="subItem.class">{{subItem.icon}}</i>
                          <span>{{subItem.name}}</span>
                      </router-link>
				            </li>
                	</ul>
	            </li>
	        </ul>
	    </div>
	    <asideFooter
	        :version="footerLayout.version"
	        :copyright="footerLayout.copyright"
	        :years="footerLayout.years">
	    </asideFooter>
	</aside>
</template>

<script>

import menuLayout from '@/layout/left_aside_config';

export default {
  name: 'left-aside',
  props: {
    footerLayout: {
      type: Object,
      default: () => {}
    },
    sidebarImage: {
      type: String,
      default: require("@/assets/images/aside.png")
    },
    navigationTitle: {
    	type: String,
    	default: "MAIN NAVIGATION"
    }
  },
  mixins: [
  ],
  data () {
    return {
    	menuLayout: menuLayout,
    }
  },
  methods: {
  	asideItemClick: function(e, itemClicked) {
  		e.preventDefault();
  		if(itemClicked.submenus.length > 0 ){
  			// an item was clicked to show submenus. do nothing here
  			menuLayout.forEach(function(item) {
		        item.active = false;
			});
	  		itemClicked.active = true;
  		} else {
  			// the item was clicked to navigate to requested view
  			menuLayout.forEach(function(item) {
		        item.active = false;
			});
	  		itemClicked.active = true;
	  		//same as {'name': item.to}
	  		this.$router.push(itemClicked.to);
  		}
  	},
  	activate: function () {
        var _this = this;
        var $body = $('body');
        var $overlay = $('.overlay');

        //Close sidebar
        $(window).click(function (e) {
            var $target = $(e.target);
            if (e.target.nodeName.toLowerCase() === 'i') { $target = $(e.target).parent(); }

            if (!$target.hasClass('bars') && _this.isOpen() && $target.parents('#leftsidebar').length === 0) {
                if (!$target.hasClass('js-right-sidebar')) $overlay.fadeOut();
                $body.removeClass('overlay-open');
            }
        });

        $.each($('.menu-toggle.toggled'), function (i, val) {
            $(val).next().slideToggle(0);
        });

        //When page load
        $.each($('.menu .list li.active'), function (i, val) {
            var $activeAnchors = $(val).find('a:eq(0)');

            $activeAnchors.addClass('toggled');
            $activeAnchors.next().show();
        });

        //Collapse or Expand Menu
        $('.menu-toggle').on('click', function (e) {
            var $this = $(this);
            var $content = $this.next();

            if ($($this.parents('ul')[0]).hasClass('list')) {
                var $not = $(e.target).hasClass('menu-toggle') ? e.target : $(e.target).parents('.menu-toggle');

                $.each($('.menu-toggle.toggled').not($not).next(), function (i, val) {
                    if ($(val).is(':visible')) {
                        $(val).prev().toggleClass('toggled');
                        $(val).slideUp();
                    }
                });
            }

            $this.toggleClass('toggled');
            $content.slideToggle(320);
        });

        //Set menu height
        _this.setMenuHeight(true);
        _this.checkStatusForResize(true);
        $(window).resize(function () {
            _this.setMenuHeight(false);
            _this.checkStatusForResize(false);
        });

        //Set Waves
        Waves.attach('.menu .list a', ['waves-block']);
        Waves.init();
    },
    setMenuHeight: function (isFirstTime) {
        if (typeof $.fn.slimScroll != 'undefined') {
            var configs = $.AdminBSB.options.leftSideBar;
            var height = ($(window).height() - ($('.legal').outerHeight() + $('.user-info').outerHeight() + $('.navbar').innerHeight()));
            var $el = $('.list');

            if (!isFirstTime) {
                $el.slimscroll({
                    destroy: true
                });
            }

            $el.slimscroll({
                height: height + "px",
                color: configs.scrollColor,
                size: configs.scrollWidth,
                alwaysVisible: configs.scrollAlwaysVisible,
                borderRadius: configs.scrollBorderRadius,
                railBorderRadius: configs.scrollRailBorderRadius
            });

            //Scroll active menu item when page load, if option set = true
            if ($.AdminBSB.options.leftSideBar.scrollActiveItemWhenPageLoad) {
                var item = $('.menu .list li.active')[0];
                if (item) {
                    var activeItemOffsetTop = item.offsetTop;
                    if (activeItemOffsetTop > 150) $el.slimscroll({ scrollTo: activeItemOffsetTop + 'px' });
                }
            }
        }
    },
    checkStatusForResize: function (firstTime) {
        var $body = $('body');
        var $openCloseBar = $('.navbar .navbar-header .bars');
        var width = $body.width();

        if (firstTime) {
            $body.find('.content, .sidebar').addClass('no-animate').delay(1000).queue(function () {
                $(this).removeClass('no-animate').dequeue();
            });
        }

        if (width < $.AdminBSB.options.leftSideBar.breakpointWidth) {
            $body.addClass('ls-closed');
            $openCloseBar.fadeIn();
        }
        else {
            $body.removeClass('ls-closed');
            $openCloseBar.fadeOut();
        }
    },
    isOpen: function () {
        return $('body').hasClass('overlay-open');
    }
  },
  created(){
    log("left-aside::created");
  },
  mounted(){
    log("left-aside::mounted");
    this.activate();
  },
  components: {
  	asideFooter: () => import('@/components/asideFooter')
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.fit {
    background-position: center;
    background-repeat: no-repeat;
}
.small {
  font-size: 17px;
}
</style>