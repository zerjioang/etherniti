<template>
    <div>
       <pagetitle
        title="Report an issue"
        subtitle="Have you found a bug? Do you think something can be improved? Let us know!"/>

        <div class="row clearfix">
          <apiForm :form="form.reportBug"></apiForm>
        </div>
    </div>
</template>

<script>

import reportBugForm from '@/layout/forms/reportBugForm';

export default {
  name: 'report-bug-view',
  data () {
    return {
      title: process.env.APP_NAME,
      form: {
        reportBug: reportBugForm
      }
    }
  },
  methods: {
  },
  created(){
    log("report-bug-view::created");
  },
  mounted(){
    log("report-bug-view::mounted");
  },
  components: {
    pagetitle: () => import('@/components/pagetitle'),
    apiForm: () => import('@/components/apiForm'),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped="true">
.centered {
  text-align: center;
}
</style>