'use strict'

const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  PRODUCTION: 'false',

  API_SCHEME: '"https"',
  DOMAIN: '"gaethway.org"',
  API_DOMAIN: '"localhost:4430"',
  WEBAPP_DOMAIN: '"gaethway.org"',
  
  DASHBOARD_BASE_PATH: '"/dashboard"',
  
  SHOW_DEBUG_DATA: 'false',
  DASHBOARD_PRIVATE_ACCESS: 'false'
})