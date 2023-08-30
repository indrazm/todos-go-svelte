<script>
	export let data;
	let task = '';

	const addTask = async () => {
		const res = await fetch(`http://localhost:8080/`, {
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
		const res = await fetch(`http://localhost:8080/${id}`, {
			method: 'DELETE'
		});
		const resData = await res.json();
		data = { data: resData };
	};
</script>

<h1>Todo list</h1>

<input type="text" placeholder="Add task" bind:value={task} />
<button on:click={addTask}>Add task</button>

{#each data.data as item}
	<div class="flex">
		<p>{item.title}</p>
		<button on:click={() => deleteTask(item.id)}>Delete</button>
	</div>
{/each}

<style>
	.flex {
		display: flex;
		gap: 4px;
		align-items: center;
	}
</style>
