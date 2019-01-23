const routerNames = {
	home: {
	  path: '/',
	  name: 'Dashboard',
	},
	ethereumTools: {
	  path: '/accounts/ethereum',
	  name: 'Ethereum',
	},
	quorumTools: {
	  path: '/accounts/quorum',
	  name: 'Quorum',
	},
	fabricTools: {
	  path: '/accounts/fabric',
	  name: 'Hyperledger Fabric',
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