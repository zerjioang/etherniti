const reportBugForm = {
	id:10,
	class: "col-lg-12 col-md-12 col-sm-12 col-xs-12",
	type: "card",
	api: {
		method: 'post',
		url: '/ui/bug'
	},
	model: {
		username: undefined,
		email: undefined,
		details: undefined,
	},
	header: {
		title: "Bug reporting tool",
		subtitle: "Contact us",
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
		method: "POST",
		columns: [
			{
				id:10,
				class: "col-md-4",
				inputgroup: [
					{
						id:101,
						title: "Contact information",
						class: "form-group form-group-lg",
						items: [
							{
								id:102,
								type:"icon+text",
								icon: "account_circle",
								class: "input-group",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "Your name (Jhon Doe)",
									required: true,
									disabled: false,
									autocomplete: "off",
									modelKey: "username",
								}
							},
							{
								id:103,
								type:"icon+text",
								icon: "email",
								class: "input-group",
								input: {
									type: "email",
									pattern: ".+@globex.com",
									size: "60",
									class: "form-control key",
									placeholder: "Your email (username@domain.tld)",
									required: true,
									disabled: false,
									autocomplete: "off",
									modelKey: "email",
								}
							},
							{
								id:103,
								type:"icon+text",
								icon: "language",
								class: "input-group",
								input: {
									type: "text",
									class: "form-control key",
									placeholder: "Language (English, Spanish, etc)",
									required: true,
									disabled: false,
									autocomplete: "off",
									modelKey: "email",
								}
							}
						]
					},
				]
			},
			{
				id:20,
				class: "col-md-8",
				inputgroup: [
					{
						id:201,
						title: "Bug details",
						class: "form-group form-group-lg",
						items: [
							{
								id:202,
								type:"textarea",
								icon: "message",
								class: "input-group",
								input: {
									type: "textarea",
									class: "form-control key",
									placeholder: "Enter bug details...",
									cols: 30,
									rows: 7,
									required: true,
									disabled: false,
									autocomplete: "off",
									modelKey: "details",
								}
							}
						]
					},
				]
			},
			{
				id:30,
				class: "col-md-12",
				buttons: [
					{
						id: 31,
						type: "submit",
						class: "btn btn-lg bg-indigo m-t-15 waves-effect upper",
						text: "report bug"
					}
				]
			}
		]
	}
};

export default reportBugForm;