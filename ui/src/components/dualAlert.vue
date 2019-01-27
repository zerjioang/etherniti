<template>
<div v-show="result.visible" class="col-md-12">
  <div class="alert fade"
  :class="{
  'alert-success': result.valid,
  'alert-danger': !result.valid,
  'in': result.visible
  }">
      The address <strong>{{form.address}}</strong> {{result.message}}
  </div>
</div>
</template>