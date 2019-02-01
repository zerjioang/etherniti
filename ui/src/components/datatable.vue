<template>
	<div class="card" v-if="rows!=undefined && rows.length > 0">
		<div class="header">
            <h2>
                {{title}}
                <small>{{subtitle}}</small>
            </h2>
            <ul v-if="false" class="header-dropdown m-r--5">
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
		<div class="body table-responsive">
		  <filterBar v-if="searchEnabled"></filterBar>
		  <vuetable
		  	class="table table-hover"
		    ref="vuetable"
		    :api-mode="false"
		    :css="table.css"
		    api-url="https://vuetable.ratiw.net/api/users"
		    :data=formattedRows
		    :data-manager="dataManager"
		    :fields="fields"
		    data-path="data"
		    pagination-path="links.pagination"

		    @vuetable:pagination-data="onPaginationData">
		    <div slot="actions" scope="props">
		        <div class="table-button-container">
		            <button class="ui button" @click="editRow(props.rowData)"><i class="fa fa-edit"></i> Edit</button>&nbsp;&nbsp;
		            <button class="ui basic red button" @click="deleteRow(props.rowData)"><i class="fa fa-remove"></i> Delete</button>&nbsp;&nbsp;
		        </div>
		    </div> <!-- end of custom actions -->
		  </vuetable>

		  <vuetable-pagination-info
		    ref="paginationInfo"
		  />
		  <vuetable-pagination
		  	class="pagination test"
		    ref="pagination"
		    @vuetable-pagination:change-page="onChangePage">
		  </vuetable-pagination>
		</div>
	</div>
</template>


<script>

//import and register axios
import axios from 'axios';

// data management
import _ from "lodash";

require('~/vuetable-2/dist/vuetable-2.css')

import VueTable, {
  VuetablePagination,
  VuetablePaginationInfo,
} from 'vuetable-2';

export default {
  name: 'datatable',
  props: {
  	title: {
  		type: String,
  		default: "Table"
  	},
  	subtitle: {
  		type: String,
  		default: "Visualize items ordered in pages"
  	},
  	searchEnabled: {
  		type: Boolean,
  		default: false
  	},
  	rows: {
  		type: Array,
  		default: () => []
  	},
  	fields: {
  		type: Array,
  		default: () => []
  	}
  },
  computed: {
  	formattedRows: function(){
  		return this.transformData(this.rows);
  	}
  },
  watch: {
  	rows: function(current, previous) {
  		Vue.nextTick( function() {
	  		this.$refs.vuetable.refresh();
		});
  	},
  	formattedRows: function(current, previous) {
  		Vue.nextTick( function() {
	  		this.$refs.vuetable.refresh();
		});
  	}
  },
  data () {
    return {
      table: {
        rowId: "",
        sortOrder: "",
        pagination: {
			current_page: 1,
			per_page: 10,
			from: 1,
			to: 1,
			last_page: 1,
			next_page_url: "",
			prev_page_url: "",
			total: 0,
		},
		css: {
          ascendingIcon: 'glyphicon glyphicon-chevron-up',
          descendingIcon: 'glyphicon glyphicon-chevron-down'
        }
      }
    }
  },
  methods: {
	isNextPressedAndNotTheLastPage: function(page, current_page, last_page) {
		return (page == 'next' && current_page < last_page);
	},
	isBackPressedAndNotTheFirstPage: function(page, current_page) {
		return (page == 'prev' && current_page > 1);
	},
	isPaginationNumberPressed: function(page) {
		return (typeof(page) == 'number');
	},
  	dataManager: function(sortOrder, pagination) {
  		if (!this.formattedRows) {
  			return
  		}
  		const local = this.rows;

		// set pagination data
		this.table.pagination.total = local.length;
		this.table.pagination.last_page = Math.ceil(this.table.pagination.total / this.table.pagination.per_page);

		const firstItem = (this.table.pagination.current_page-1) * this.table.pagination.per_page;
		const dataInPage = _.slice(local, firstItem, firstItem + this.table.pagination.per_page);

		return {
			data: dataInPage,
			links: { pagination: this.table.pagination },
		};

  	},	
    formatEnode: function (value) {
      return value;
    },
    transformData: function(raw) {
      const rawArray = raw;
      rawArray.forEach(function(part, index) {
        rawArray[index] = {
          id: index,
          node: rawArray[index]
        };
      });
      const formatted = {
        data: rawArray,
        links: {
          pagination: this.table.pagination
        },
      };
      return formatted;
    },
    //tables methods
    onPaginationData (paginationData) {
      // reference
      // this.$refs.pagination.setPaginationData(paginationData);
      // this.$refs.paginationInfo.setPaginationData(paginationData);
      this.$refs.pagination.setPaginationData(this.table.pagination);
      this.$refs.paginationInfo.setPaginationData(this.table.pagination);
    },
    onChangePage (page) {
      if (this.isNextPressedAndNotTheLastPage(page, this.table.pagination.current_page, this.table.pagination.last_page)) {
        this.table.pagination.current_page = this.table.pagination.current_page +1;
        this.$refs.vuetable.refresh();
      } else if (this.isBackPressedAndNotTheFirstPage(page, this.table.pagination.current_page)) {
        this.table.pagination.current_page = this.table.pagination.current_page -1;
        this.$refs.vuetable.refresh();
      }
      else if (this.isPaginationNumberPressed(page)) {
        this.table.pagination.current_page = page;
        this.$refs.vuetable.refresh();
      } 
    },
    editRow(rowData){
      alert("You clicked edit on"+ JSON.stringify(rowData));
    },
    deleteRow(rowData){
      alert("You clicked delete on"+ JSON.stringify(rowData));
    }
  },
  created(){
    log("datatable::created");
  },
  mounted(){
    log("datatable::mounted");
  },
  components: {
    preloader: () => import('@/components/preloader'),
    vuetable: VueTable,
    'vuetable-pagination': VuetablePagination,
    'vuetable-pagination-info': VuetablePaginationInfo,
    filterBar: () => import('@/components/FilterBar'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>

<style>
.upper {
	text-transform: uppercase;
}
.bold {
	font-weight: bold;
}
.code {
	font-family: 'IBM Plex Mono';
	color: #233286;
	font-weight: 500;
	overflow-y: auto;
	word-break: break-all;
}
</style>