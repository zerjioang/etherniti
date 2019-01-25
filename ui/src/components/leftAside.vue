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
	    <!-- Menu -->
	    <div class="menu">
	        <ul class="list">
	            <li class="header">MAIN NAVIGATION</li>
	            <li class="active">
	                <a href="index.html">
	                    <i class="material-icons">home</i>
	                    <span>Home</span>
	                </a>
	            </li>
	            <li>
	                <a href="pages/typography.html">
	                    <i class="material-icons">text_fields</i>
	                    <span>Typography</span>
	                </a>
	            </li>
	            <li>
	                <a href="pages/helper-classes.html">
	                    <i class="material-icons">layers</i>
	                    <span>Helper Classes</span>
	                </a>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">widgets</i>
	                    <span>Widgets</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="javascript:void(0);" class="menu-toggle">
	                            <span>Cards</span>
	                        </a>
	                        <ul class="ml-menu">
	                            <li>
	                                <a href="pages/widgets/cards/basic.html">Basic</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/cards/colored.html">Colored</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/cards/no-header.html">No Header</a>
	                            </li>
	                        </ul>
	                    </li>
	                    <li>
	                        <a href="javascript:void(0);" class="menu-toggle">
	                            <span>Infobox</span>
	                        </a>
	                        <ul class="ml-menu">
	                            <li>
	                                <a href="pages/widgets/infobox/infobox-1.html">Infobox-1</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/infobox/infobox-2.html">Infobox-2</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/infobox/infobox-3.html">Infobox-3</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/infobox/infobox-4.html">Infobox-4</a>
	                            </li>
	                            <li>
	                                <a href="pages/widgets/infobox/infobox-5.html">Infobox-5</a>
	                            </li>
	                        </ul>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">swap_calls</i>
	                    <span>User Interface (UI)</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/ui/alerts.html">Alerts</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/animations.html">Animations</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/badges.html">Badges</a>
	                    </li>

	                    <li>
	                        <a href="pages/ui/breadcrumbs.html">Breadcrumbs</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/buttons.html">Buttons</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/collapse.html">Collapse</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/colors.html">Colors</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/dialogs.html">Dialogs</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/icons.html">Icons</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/labels.html">Labels</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/list-group.html">List Group</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/media-object.html">Media Object</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/modals.html">Modals</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/notifications.html">Notifications</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/pagination.html">Pagination</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/preloaders.html">Preloaders</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/progressbars.html">Progress Bars</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/range-sliders.html">Range Sliders</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/sortable-nestable.html">Sortable & Nestable</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/tabs.html">Tabs</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/thumbnails.html">Thumbnails</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/tooltips-popovers.html">Tooltips & Popovers</a>
	                    </li>
	                    <li>
	                        <a href="pages/ui/waves.html">Waves</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">assignment</i>
	                    <span>Forms</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/forms/basic-form-elements.html">Basic Form Elements</a>
	                    </li>
	                    <li>
	                        <a href="pages/forms/advanced-form-elements.html">Advanced Form Elements</a>
	                    </li>
	                    <li>
	                        <a href="pages/forms/form-examples.html">Form Examples</a>
	                    </li>
	                    <li>
	                        <a href="pages/forms/form-validation.html">Form Validation</a>
	                    </li>
	                    <li>
	                        <a href="pages/forms/form-wizard.html">Form Wizard</a>
	                    </li>
	                    <li>
	                        <a href="pages/forms/editors.html">Editors</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">view_list</i>
	                    <span>Tables</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/tables/normal-tables.html">Normal Tables</a>
	                    </li>
	                    <li>
	                        <a href="pages/tables/jquery-datatable.html">Jquery Datatables</a>
	                    </li>
	                    <li>
	                        <a href="pages/tables/editable-table.html">Editable Tables</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">perm_media</i>
	                    <span>Medias</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/medias/image-gallery.html">Image Gallery</a>
	                    </li>
	                    <li>
	                        <a href="pages/medias/carousel.html">Carousel</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">pie_chart</i>
	                    <span>Charts</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/charts/morris.html">Morris</a>
	                    </li>
	                    <li>
	                        <a href="pages/charts/flot.html">Flot</a>
	                    </li>
	                    <li>
	                        <a href="pages/charts/chartjs.html">ChartJS</a>
	                    </li>
	                    <li>
	                        <a href="pages/charts/sparkline.html">Sparkline</a>
	                    </li>
	                    <li>
	                        <a href="pages/charts/jquery-knob.html">Jquery Knob</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">content_copy</i>
	                    <span>Example Pages</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/examples/profile.html">Profile</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/sign-in.html">Sign In</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/sign-up.html">Sign Up</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/forgot-password.html">Forgot Password</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/blank.html">Blank Page</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/404.html">404 - Not Found</a>
	                    </li>
	                    <li>
	                        <a href="pages/examples/500.html">500 - Server Error</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">map</i>
	                    <span>Maps</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="pages/maps/google.html">Google Map</a>
	                    </li>
	                    <li>
	                        <a href="pages/maps/yandex.html">YandexMap</a>
	                    </li>
	                    <li>
	                        <a href="pages/maps/jvectormap.html">jVectorMap</a>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="javascript:void(0);" class="menu-toggle">
	                    <i class="material-icons">trending_down</i>
	                    <span>Multi Level Menu</span>
	                </a>
	                <ul class="ml-menu">
	                    <li>
	                        <a href="javascript:void(0);">
	                            <span>Menu Item</span>
	                        </a>
	                    </li>
	                    <li>
	                        <a href="javascript:void(0);">
	                            <span>Menu Item - 2</span>
	                        </a>
	                    </li>
	                    <li>
	                        <a href="javascript:void(0);" class="menu-toggle">
	                            <span>Level - 2</span>
	                        </a>
	                        <ul class="ml-menu">
	                            <li>
	                                <a href="javascript:void(0);">
	                                    <span>Menu Item</span>
	                                </a>
	                            </li>
	                            <li>
	                                <a href="javascript:void(0);" class="menu-toggle">
	                                    <span>Level - 3</span>
	                                </a>
	                                <ul class="ml-menu">
	                                    <li>
	                                        <a href="javascript:void(0);">
	                                            <span>Level - 4</span>
	                                        </a>
	                                    </li>
	                                </ul>
	                            </li>
	                        </ul>
	                    </li>
	                </ul>
	            </li>
	            <li>
	                <a href="pages/changelogs.html">
	                    <i class="material-icons">update</i>
	                    <span>Changelogs</span>
	                </a>
	            </li>
	            <li class="header">LABELS</li>
	            <li>
	                <a href="javascript:void(0);">
	                    <i class="material-icons col-red">donut_large</i>
	                    <span>Important</span>
	                </a>
	            </li>
	            <li>
	                <a href="javascript:void(0);">
	                    <i class="material-icons col-amber">donut_large</i>
	                    <span>Warning</span>
	                </a>
	            </li>
	            <li>
	                <a href="javascript:void(0);">
	                    <i class="material-icons col-light-blue">donut_large</i>
	                    <span>Information</span>
	                </a>
	            </li>
	        </ul>
	    </div>
	    <!-- #Menu -->
	    <asideFooter
	        :version="footerLayout.version"
	        :copyright="footerLayout.copyright"
	        :years="footerLayout.years">
	    </asideFooter>
	    <!-- #Footer -->
	</aside>
</template>

<script>

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
    }
  },
  mixins: [
  ],
  data () {
    return {
    }
  },
  methods: {
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
</style>