/*
form elements
https://preview.tabler.io/form-elements.html
*/
const formModel = {
	ethAccountCreateForm: {
		title: "Create ethereum account",
		buttonText: "Create account",
		buttonClass: "btn btn-primary",
		formClass: "card",
		rows: [
			{
				id: 0,
				class: "row",
				columns: [
					{
						id: 0,
						class: "col-sm-12 col-md-6 col-lg-6",
						groups: [
							{
								id: 0,
								class: "form-group",
								elements: [
									{
										id: 0,
										type: "input+label",
										class: "form-label",
										label: "Account",
										placeholder: "ethereum account",
										required: false,
										disabled: true,
										value: undefined
									}
								]
							}
						]
					},
					{
						id: 1,
						class: "col-sm-12 col-md-6 col-lg-6",
						groups: [
							{
								id: 0,
								class: "form-group",
								elements: [
									{
										id: 0,
										type: "input+label",
										class: "form-label",
										label: "Private key",
										placeholder: "ethereum private key",
										required: false,
										disabled: true,
										value: undefined
									}
								]
							}
						]
					}
				]
			}
		]
	},
	ethSignCreateForm: {
		title: "Sign arbitrary data",
		buttonText: "Sign with my ETH account",
		buttonClass: "btn btn-primary",
		formClass: "card",
		rows: [
			{
				id: 0,
				class: "row",
				columns: [
					{
						id: 0,
						class: "col-sm-12 col-md-12 col-lg-12",
						groups: [
							{
								id: 0,
								class: "form-group",
								elements: [
									{
										id: 0,
										type: "input+label",
										class: "form-label",
										label: "Address",
										placeholder: "ethereum account",
										required: true,
										disabled: false,
										value: undefined
									},
									{
										id: 1,
										type: "input+label",
										class: "form-label",
										label: "Private key",
										placeholder: "ethereum private key",
										required: true,
										disabled: false,
										value: undefined
									}
								]
							}
						]
					},
					{
						id: 1,
						class: "col-sm-12 col-md-12 col-lg-12",
						groups: [
							{
								id: 0,
								class: "form-group",
								elements: [
									{
										id: 0,
										type: "input+label",
										class: "form-label",
										label: "Message",
										placeholder: "message",
										required: true,
										disabled: false,
										value: undefined
									},
									{
										id: 1,
										type: "input+label",
										class: "form-label",
										label: "Message hash",
										placeholder: "message hash",
										required: false,
										disabled: true,
										value: undefined
									},
									{
										id: 2,
										type: "input+label",
										class: "form-label",
										label: "Signature",
										placeholder: "signature",
										required: false,
										disabled: true,
										value: undefined
									}
								]
							}
						]
					}
				]
			}
		]
	}
};

export default formModel;