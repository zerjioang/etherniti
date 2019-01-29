const contractCheckForm = {
	id:10,
	class: "col-lg-6 col-md-6 col-sm-12 col-xs-12",
	type: "card",
	api: {
		method: 'get',
		url: '/eth/hascontract/'
	},
	model: {
		address: undefined
	},
	header: {
		title: "Ethereum Contract Address Check",
		subtitle: "Verify if given address is a contract address or not",
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
				class: "col-md-12",
				inputgroup: [
					{
						id:101,
						title: "ETH Contract address",
						class: "input-group",
						items: [
							{
								id:102,
								type:"icon+text",
								class: "input-group",
								icon: "account_box",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "0x798abda6cc246d0edba912092a2a3dbd3d11191b",
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
						id: 11,
						type: "submit",
						class: "btn btn-lg bg-indigo m-t-15 waves-effect upper",
						text: "verify contract address"
					}
				]
			}
		]
	},
};

export default contractCheckForm;