'use strict'

const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  PRODUCTION: 'false',

  API_SCHEME: '"http"',
  API_DOMAIN: '"apidev.methw.org:8080"',

  SWAGGER_URL: '"http://localhost:5555/api/docs/"',
  
  DASHBOARD_BASE_PATH: '"/dashboard"',
  
  SHOW_DEBUG_DATA: 'false',
  DASHBOARD_PRIVATE_ACCESS: 'false'
})