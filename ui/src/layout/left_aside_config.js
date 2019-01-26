import routerNames from '@/layout/routerNames';

/*
icons are: https://feathericons.com/
*/
const leftAsideLayout = [
	{
		id:100,
		name:routerNames.home.name,
		icon: "home",
		class: "",
		to:routerNames.home.path,
		active: true,
		submenus: []
	},
	{
		id:101,
		name:routerNames.eth.name,
		icon: "storage",
		class: "",
		to:routerNames.eth.path,
		active: false,
		submenus: [
			{
				id:1011,
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
		id:101,
		name:routerNames.quorum.name,
		icon: "storage",
		class: "",
		to:routerNames.quorum.path,
		active: false,
		submenus: [
			{
				id:1011,
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
		id:1,
		name:routerNames.tools.name,
		icon: "layers",
		class: "",
		to:routerNames.tools.path,
		active: false,
		submenus: [
			{
				id:2,
				name:routerNames.addressChecker.name,
				icon: "done_all",
				class: "",
				to:routerNames.addressChecker.path,
				active: false,
				submenus: []
			},
			{
				id:3,
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
		id:96,
		name:"Profiles",
		icon: "group",
		class: "",
		to:"connectionProfile",
		active: false,
		submenus: [
			{
				id:2,
				name:routerNames.newProfile.name,
				icon: "add_circle_outline",
				class: "",
				to:routerNames.newProfile.path,
				active: false,
				submenus: []
			},
			{
				id:3,
				name:routerNames.manageProfile.name,
				icon: "assignment",
				class: "",
				to:routerNames.manageProfile.path,
				active: false,
				submenus: []
			}
		]
	},
	//library_books
	{
		id:96,
		name:routerNames.privateApi.name,
		icon: "lock",
		class: "",
		to:routerNames.privateApi.path,
		active: false,
		submenus: [
			{
				id:961,
				name:routerNames.manageProfile.name,
				icon: "art_track",
				class: "",
				to:routerNames.manageProfile.path,
				active: false,
				submenus: []
			},
			{
				id:962,
				name:routerNames.manageProfile.name,
				icon: "settings",
				class: "",
				to:routerNames.manageProfile.path,
				active: false,
				submenus: []
			},
			{
				id:963,
				name:routerNames.manageProfile.name,
				icon: "help_outline",
				class: "",
				to:routerNames.manageProfile.path,
				active: false,
				submenus: []
			}
		]
	},
	{
		id:97,
		name:routerNames.bugReport.name,
		icon: "report_problem",
		class: "",
		to:routerNames.bugReport.path,
		active: false,
		submenus: []
	},
	{
		id:98,
		name:routerNames.about.name,
		icon: "info",
		class: "",
		to:routerNames.about.path,
		active: false,
		submenus: []
	},
	{
		id:99,
		name:routerNames.license.name,
		icon: "public",
		class: "",
		to:routerNames.license.path,
		active: false,
		submenus: []
	},
];

export default leftAsideLayout;