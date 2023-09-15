items: [
	{
		id:    string
		value: string
	},
	...,
]

#dict: {
	for item in items {
		"\( item.id )": {value: item.value}
	}
}
