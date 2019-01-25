import routerNames from '@/layout/routerNames';

/*
icons are: https://feathericons.com/
*/
const leftAsideLayout = [
	{
		id:0,
		name:routerNames.home.name,
		icon: "home",
		class: "",
		to:routerNames.home.path,
		active: true,
		submenus: []
	},
	{
		id:1,
		name:"Useful Tools",
		icon: "layers",
		class: "",
		to:"connectionProfile",
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
		id:97,
		name:routerNames.report.name,
		icon: "report_problem",
		class: "",
		to:routerNames.report.path,
		active: false,
		submenus: []
	},
	{
		id:98,
		name:routerNames.about.name,
		icon: "info_outline",
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