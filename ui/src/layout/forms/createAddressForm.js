const createAddressForm = {
	id:10,
	class: "col-lg-6 col-md-6 col-sm-12 col-xs-12",
	type: "card",
	api: {
		method: 'get',
		url: '/eth/create/'
	},
	model: {
		address: undefined,
		key: undefined
	},
	header: {
		title: "Create Ethereum/Quorum account",
		subtitle: "Create new account dynamically",
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
						title: "Account details",
						items: [
							{
								id:102,
								type:"icon+text",
								class: "input-group",
								icon: "account_box",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
									required: false,
									disabled: true,
									autocomplete: "on",
									modelKey: "address",
								}
							},
							{
								id:103,
								type:"icon+text",
								class: "input-group",
								icon: "account_box",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "secret",
									required: false,
									disabled: true,
									autocomplete: "on",
									modelKey: "key",
								}
							}
						]
					},
				],
				buttons: [
					{
						id: 11,
						type: "submit",
						class: "btn btn-lg bg-indigo m-t-15 waves-effect upper multiple-btn-row",
						text: "Create account"
					},
					{
						id: 12,
						type: "copy-to-clipboard",
						class: "btn btn-lg bg-indigo m-t-15 waves-effect upper multiple-btn-row",
						disabled: true,
						text: "Copy to clipboard"
					}
				]
			}
		]
	}
};

export default createAddressForm;