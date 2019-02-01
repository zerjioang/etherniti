<template>
    <div>
      <pagetitle
        title="Alastria Blockchain Ecosystem"
        subtitle="Alastria network status"/>

        <widgetsRow :config="layout.counters"/>

        <div class="row clearfix">
          <div v-for="block in monitor" :key="block.name" class="col-md-12">
            <div style="text-align: center; padding: 20px">
              <preloader :visible="block.data==undefined"></preloader>
            </div>
            <datatable
              :searchEnabled=false
              :title="block.name"
              :subtitle="block.url"
              :rows="block.data"
              :fields="block.fields"/>
          </div> <!-- end of for -->
        </div>
    </div>
</template>

<script>

//import and register axios
import axios from 'axios';

import alastriaSummaryConfig from '@/layout/alastriaSummaryConfig';

export default {
  name: 'alastria-status-view',
  data () {
    return {
      layout: {
        counters: alastriaSummaryConfig
      },
      monitor: {
        constellations: {
          name: "Alastria Constellation nodes",
          url: "https://raw.githubusercontent.com/alastria/alastria-node/develop/data/constellation-nodes.json",
          fields: [
            {
              name: "id",
              sortField: "id",
              titleClass: 'text-left upper',
              dataClass: 'text-left bold',
              callback: 'formatEnode'
            },
            {
              name: "node",
              sortField: "node",
              titleClass: 'text-left upper',
              dataClass: 'text-left code',
              callback: 'formatEnode'
            }
          ],
          data: undefined
        },
        general: {
          name: "Alastria Permissioned nodes general",
          url: "https://raw.githubusercontent.com/alastria/alastria-node/develop/data/permissioned-nodes_general.json",
          fields: [
            {
              name: "id",
              sortField: "id",
              titleClass: 'text-left upper',
              dataClass: 'text-left bold',
              callback: 'formatEnode'
            },
            {
              name: "node",
              sortField: "node",
              titleClass: 'text-left upper',
              dataClass: 'text-left code',
              callback: 'formatEnode'
            }
          ],
          data: undefined
        },
        permissioned: {
          name: "Alastria Permissioned nodes validators",
          url: "https://raw.githubusercontent.com/alastria/alastria-node/develop/data/permissioned-nodes_validator.json",
          fields: [
            {
              name: "id",
              sortField: "id",
              titleClass: 'text-left upper',
              dataClass: 'text-left bold',
              callback: 'formatEnode'
            },
            {
              name: "node",
              sortField: "node",
              titleClass: 'text-left upper',
              dataClass: 'text-left code',
              callback: 'formatEnode'
            }
          ],
          data: undefined
        },
        static: {
          name: "Alastria static nodes",
          url: "https://raw.githubusercontent.com/alastria/alastria-node/develop/data/static-nodes.json",
          fields: [
            {
              name: "id",
              sortField: "id",
              titleClass: 'text-left upper',
              dataClass: 'text-left bold',
              callback: 'formatEnode'
            },
            {
              name: "node",
              sortField: "node",
              titleClass: 'text-left upper',
              dataClass: 'text-left code',
              callback: 'formatEnode'
            }
          ],
          data: undefined
        }
      }
    }
  },
  methods: {
    load: function (item, counterCard) {
      log("loading remote information...");
      axios.get(item.url)
      .then((response) => {
        if (response && response.data) {
          item.data = response.data;
          counterCard.value = response.data.length;
        }
      })
      .catch((err) => {
        error(err);
      }); 
    }
  },
  created(){
    log("alastria-status-view::created");
  },
  mounted(){
    log("alastria-status-view::mounted");
    this.load(this.monitor.constellations, this.layout.counters.cards[0]);
    this.load(this.monitor.general, this.layout.counters.cards[1]);
    this.load(this.monitor.permissioned, this.layout.counters.cards[2]);
    this.load(this.monitor.static, this.layout.counters.cards[3]);
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    preloader: () => import('@/components/preloader'),
    datatable: () => import('@/components/datatable'),
    widgetsRow: () => import('@/components/widgetsRow'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
</style>