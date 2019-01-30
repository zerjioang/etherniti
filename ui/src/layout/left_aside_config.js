import routerNames from '@/layout/routerNames';

/*
icons are: https://feathericons.com/
*/
const leftAsideLayout = [
	{
		id:100,
		name:routerNames.dashboardIndex.name,
		icon: "dashboard",
		class: "",
		to:routerNames.dashboardIndex.path,
		active: true,
		submenus: []
	},
	{
		// ganache
		id:101,
		name:routerNames.ganache.name,
		icon: "redeem",
		class: "",
		to:routerNames.ganache.path,
		active: false,
		submenus: [
			{
				//Accounts
				id:1011,
				name:routerNames.ganacheAccounts.name,
				icon: "contact_mail",
				class: "",
				to:routerNames.ganacheAccounts.path,
				active: false,
				submenus: []
			},
			{
				// transactions
				id:1011,
				name:routerNames.ganacheTransactions.name,
				icon: "reorder",
				class: "",
				to:routerNames.ganacheTransactions.path,
				active: false,
				submenus: []
			},
			{
				//blocks
				id:1011,
				name:routerNames.ganacheBlocks.name,
				icon: "view_day",
				class: "",
				to:routerNames.ganacheBlocks.path,
				active: false,
				submenus: []
			},
			{
				//settings
				id:1011,
				name:routerNames.ganacheSettings.name,
				icon: "settings",
				class: "",
				to:routerNames.ganacheSettings.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		id:102,
		name:routerNames.eth.name,
		icon: "storage",
		class: "",
		to:routerNames.eth.path,
		active: false,
		submenus: [
			{
				id:1023,
				name:routerNames.ethCreate.name,
				icon: "add",
				class: "",
				to:routerNames.ethCreate.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		id:103,
		name:routerNames.quorum.name,
		icon: "storage",
		class: "",
		to:routerNames.quorum.path,
		active: false,
		submenus: [
			{
				id:1031,
				name:routerNames.quorumCreate.name,
				icon: "add",
				class: "",
				to:routerNames.quorumCreate.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		id:104,
		name:routerNames.tools.name,
		icon: "layers",
		class: "",
		to:routerNames.tools.path,
		active: false,
		submenus: [
			{
				id:1041,
				name:routerNames.addressChecker.name,
				icon: "done_all",
				class: "",
				to:routerNames.addressChecker.path,
				active: false,
				submenus: []
			},
			{
				id:1042,
				name:routerNames.balanceChecker.name,
				icon: "attach_money",
				class: "",
				to:routerNames.balanceChecker.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		id:105,
		name:"Profiles",
		icon: "group",
		class: "",
		to:"connectionProfile",
		active: false,
		submenus: [
			{
				id:1051,
				name:routerNames.newProfile.name,
				icon: "add_circle_outline",
				class: "",
				to:routerNames.newProfile.path,
				active: false,
				submenus: []
			},
			{
				id:1052,
				name:routerNames.manageProfile.name,
				icon: "assignment",
				class: "",
				to:routerNames.manageProfile.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		// private api
		id:106,
		name:routerNames.privateApi.name,
		icon: "lock",
		class: "",
		to:routerNames.privateApi.path,
		active: false,
		submenus: [
			{
				// private api > configuration
				id:1061,
				name:routerNames.privateApiConfiguration.name,
				icon: "art_track",
				class: "",
				to:routerNames.privateApiConfiguration.path,
				active: false,
				submenus: []
			},
			{
				// private api > management
				id:1062,
				name:routerNames.privateApiManagement.name,
				icon: "settings",
				class: "",
				to:routerNames.privateApiManagement.path,
				active: false,
				submenus: []
			},
			{
				// private api > information
				id:1063,
				name:routerNames.privateApiInformation.name,
				icon: "help_outline",
				class: "",
				to:routerNames.privateApiInformation.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		// bug report
		id:107,
		name:routerNames.bugReport.name,
		icon: "report_problem",
		class: "",
		to:routerNames.bugReport.path,
		active: false,
		submenus: []
	},
	{
		// about
		id:108,
		name:routerNames.about.name,
		icon: "info",
		class: "",
		to:routerNames.about.path,
		active: false,
		submenus: []
	},
	{
		// license
		id:109,
		name:routerNames.license.name,
		icon: "public",
		class: "",
		to:routerNames.license.path,
		active: false,
		submenus: []
	},
];

export default leftAsideLayout;