const balanceCheckForm = {
	id:10,
	class: "col-lg-12 col-md-12 col-sm-12 col-xs-12",
	type: "card",
	api: {
		method: 'get',
		url: '/eth/verify/'
	},
	model: {
		nodeAddress: undefined,
		profileId: undefined,
		address: undefined,
	},
	header: {
		title: "Ethereum Address balance",
		subtitle: "Get the balance of Ethereum account",
		dropdown: {
			icon: "more_vert",
			items: [
				{
					id:10,
					title:"What's this?",
					onclick: "a"
				},
				{
					id:11,
					title:"Show internals",
					onclick: "a"
				},
				{
					id:12,
					title:"View webAPI request",
					onclick: "a"
				}
			]
		}
	},
	body: {
		// titleInside: "Enter address to validate"
		rowClass: "row clearfix",
		method: "GET",
		columns: [
			{
				id:10,
				class: "col-md-6",
				inputgroup: [
					{
						id: 201,
						title: "Connection information",
						items: [
							{
								id: 202,
								type:"icon+text",
								class: "input-group",
								icon: "http",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "http://infura.io/XXX",
									required: true,
									disabled: false,
									autocomplete: "on",
									modelKey: "nodeAddress",
								}
							},
							{
								id: 202,
								type:"icon+text",
								class: "input-group",
								icon: "dns",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "Connection Profile ID or Token",
									required: true,
									disabled: false,
									autocomplete: "on",
									modelKey: "profileId",
								}
							}
						]
					}
				]
			},
			{
				id:10,
				class: "col-md-6",
				inputgroup: [
					{
						id: 201,
						title: "ETH Address",
						items: [
							{
								id: 202,
								type:"icon+text",
								class: "input-group",
								icon: "credit_card",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
									required: true,
									disabled: false,
									autocomplete: "on",
									modelKey: "address",
								}
							}
						]
					},
				],
				buttons: [
					{
						id: 21,
						type: "submit",
						class: "btn btn-lg bg-indigo m-t-15 waves-effect upper",
						text: "Get balance"
					}
				]
			}
		]
	}
};

export default balanceCheckForm;