<script>
	export let data;
	let task = '';

	const addTask = async () => {
		const res = await fetch(`http://localhost:8080/api/v1/task`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				id: data.data.length + 1,
				title: task,
				isCompleted: false
			})
		});
		const resData = await res.json();
		data = { data: resData };
		task = '';
	};

	const deleteTask = async (id) => {
		const res = await fetch(`http://localhost:8080/api/v1/task/${id}`, {
			method: 'DELETE'
		});
		const resData = await res.json();
		data = { data: resData };
	};
</script>

<h1 class="text-2xl font-bold">Todo list</h1>

<div class="flex items-center">
	<input type="text" placeholder="Add task" bind:value={task} />
	<button on:click={addTask}>Add task</button>
</div>

<div>
	{#each data.data as item}
		<div class="flex">
			<p>{item.title}</p>
			<button on:click={() => deleteTask(item.id)}>Delete</button>
		</div>
	{/each}
</div>

<style lang="postcss">
	:global(body) {
		@apply bg-black text-slate-300;
	}
	:global(input) {
		@apply w-auto py-2 px-3 rounded-lg bg-slate-800 text-slate-300;
	}
	:global(button) {
		@apply w-fit py-2 px-3 bg-indigo-600 hover:bg-indigo-400 text-white rounded-lg font-medium;
	}
</style>
