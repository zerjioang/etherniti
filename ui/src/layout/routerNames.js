const routerNames = {
	index: {
	  path: '/',
	  name: 'Index',
	  component: () => import('@/views/index')
	},
	dashboardHome: {
	  path: '/dashboard',
	  name: 'Dashboard',
	  component: undefined
	},
		home: {
		  path: '/dashboard/',
		  name: 'Dashboard',
		  component: undefined
		},
		eth: {
		  path: '/dashboard/eth/',
		  name: 'Ethereum',
		  component: undefined
		},
			ethCreate: {
			  path: '/dashboard/eth/create',
			  name: 'New Account',
			  component: undefined
			},
		quorum: {
		  path: '/dashboard/quorum/',
		  name: 'Quorum',
		  component: undefined
		},
			quorumCreate: {
			  path: '/dashboard/quorum/create',
			  name: 'New Account',
			  component: undefined
			},
		tools: {
		  path: '/tools',
		  name: 'Useful tools',
		  component: undefined
		},
		addressChecker: {
		  path: '/address/verify',
		  name: 'Address Check',
		  component: undefined
		},
		balanceChecker: {
		  path: '/address/balance',
		  name: 'Balance Check',
		  component: undefined
		},
		newProfile: {
		  path: '/profiles/create',
		  name: 'New Profile',
		  component: undefined
		},
		manageProfile: {
		  path: '/profiles/manage',
		  name: 'Manage Profiles',
		  component: undefined
		},
		privateApi: {
		  path: '/private/api',
		  name: 'Private API',
		  component: undefined
		},
			privateApiConfiguration: {
			  path: '/private/api/configure',
			  name: 'Configuration',
			  component: undefined
			},
			privateApiManagement: {
			  path: '/private/api/manage',
			  name: 'Management',
			  component: undefined
			},
			privateApiInformation: {
			  path: '/private/api/info',
			  name: 'Information',
			  component: undefined
			},
		license: {
		  path: '/license',
		  name: 'License',
		  component: undefined
		},
		bugReport: {
		  path: '/report/issue',
		  name: 'Report a problem',
		  component: undefined
		},
		about: {
		  path: '/about',
		  name: 'About',
		  component: undefined
		},
	notfound: {
	  path: '*',
	  name: 'Not found',
	  component: () => import('@/views/notfound')
	},
	localstorage: {
	  path: '/browser/incompatibility/localstorage',
	  name: 'localstorage',
	  component: () => import('@/views/localstorage')
	},
	maintenance: {
	  path: '/',
	  name: 'Maintenance',
	  component: () => import('@/views/maintenance')
	},
	soon: {
	  path: '/',
	  name: 'Cooming soon',
	  component: () => import('@/views/soon')
	},
};

export default routerNames;