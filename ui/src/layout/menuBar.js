import routerNames from '@/layout/routerNames';

/*
icons are: https://feathericons.com/
*/
const menuBarLayout = {
	items: [
		{
			id:0,
			name:routerNames.home.name,
			to:routerNames.home.name,
			class: "fe fe-home",
			submenus: []
		},
		{
			id:1,
			name:"Tools",
			class: "fe fe-calendar",
			to:routerNames.home.name,
			submenus: [
				{
					id:0,
					name:routerNames.ethereumTools.name,
					to:routerNames.ethereumTools.name,
					class: "fe fe-home",
					submenus: []
				},
				{
					id:1,
					name:routerNames.quorumTools.name,
					to:routerNames.quorumTools.name,
					class: "fe fe-home",
					submenus: []
				},
				{
					id:2,
					name:routerNames.fabricTools.name,
					to:routerNames.fabricTools.name,
					class: "fe fe-home",
					submenus: []
				},
			]
		},
		{
			id:2,
			name:"Connection",
			class: "fe fe-cloud-lightning",
			to:"connectionProfile",
			submenus: []
		},
		{
			id:3,
			name:"About",
			class: "fe fe-info",
			to:routerNames.about.name,
			submenus: []
		}
	]
};

export default menuBarLayout;