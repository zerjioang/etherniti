const routerNames = {
	index: {
	  path: '/',
	  name: 'Index',
	},
	dashboardHome: {
	  path: '/dashboard',
	  name: 'Dashboard',
	},
		home: {
		  path: '/dashboard/',
		  name: 'Dashboard',
		},
		eth: {
		  path: '/dashboard/eth/',
		  name: 'Ethereum',
		},
			ethCreate: {
			  path: '/dashboard/eth/create',
			  name: 'New Account',
			},
		quorum: {
		  path: '/dashboard/quorum/',
		  name: 'Quorum',
		},
			quorumCreate: {
			  path: '/dashboard/quorum/create',
			  name: 'New Account',
			},
		tools: {
		  path: '/tools',
		  name: 'Useful tools',
		},
		addressChecker: {
		  path: '/address/verify',
		  name: 'Address Check',
		},
		balanceChecker: {
		  path: '/address/balance',
		  name: 'Balance Check',
		},
		newProfile: {
		  path: '/profiles/create',
		  name: 'New Profile',
		},
		manageProfile: {
		  path: '/profiles/manage',
		  name: 'Manage Profiles',
		},
		privateApi: {
		  path: '/private/api',
		  name: 'Private API',
		},
			privateApiConfiguration: {
			  path: '/private/api/configure',
			  name: 'Configuration',
			},
			privateApiManagement: {
			  path: '/private/api/manage',
			  name: 'Management',
			},
			privateApiInformation: {
			  path: '/private/api/info',
			  name: 'Information',
			},
		license: {
		  path: '/license',
		  name: 'License',
		},
		bugReport: {
		  path: '/report/issue',
		  name: 'Report a problem',
		},
		about: {
		  path: '/about',
		  name: 'About',
		},
	notfound: {
	  path: '*',
	  name: 'Not found',
	}
};

export default routerNames;