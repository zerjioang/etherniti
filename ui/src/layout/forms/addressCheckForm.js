const addressCheckForm = {
	forms: [
		{
			id:10,
			class: "",
			type: "card",
			model: {
				address: undefined
			},
			header: {
				title: "Ethereum Address validation",
				subtitle: "Verify if given address is valid or not",
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
				submit: function (e, component) {
					e.preventDefault();
					alert('submit me');
					console.log(e)
					console.log(component)
				},
				columns: [
					{
						id:10,
						class: "col-md-12",
						inputgroup: [
							{
								id:101,
								title: "ETH Address",
								items: [
									{
										id:102,
										type:"icon+text",
										icon: "account_box",
										input: {
											type: "text",
											class: "form-control key",
											placeholder: "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
											required: true,
											disabled: false,
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
								text: "verify address"
							}
						]
					}
				]
			},
			columns: [
				{
					id:10,
					class: "col-lg-4 col-md-5 col-sm-7 col-xs-12",
					eventid: "verify"
				}
			]
		}
	]
};

export default addressCheckForm;