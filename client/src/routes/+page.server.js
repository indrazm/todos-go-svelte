export async function load() {
	const res = await fetch('http://localhost:8080');
	const data = await res.json();
	return { data };
}

export const actions = {
	create: async ({ request }) => {
		console.log(request.body());
		// const res = await fetch(`http://localhost:8080/`, {
		// 	method: 'POST',
		// 	headers: {
		// 		'Content-Type': 'application/json'
		// 	},
		// 	body: JSON.stringify({
		// 		id: 5,
		// 		title: 'New task',
		// 		isCompleted: false
		// 	})
		// });
		// const resData = await res.json();
		// data = { data: resData };
	},

	delete: async ({ cookies, request }) => {
		const data = await request.formData();
		db.deleteTodo(cookies.get('userid'), data.get('id'));
	}
};
