const routerNames = {
	home: {
	  path: '/',
	  name: 'Home',
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
	license: {
	  path: '/license',
	  name: 'License',
	},
	report: {
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